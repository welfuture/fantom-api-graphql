// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

// handleDelegationLog handles a new delegation event from logs.
func handleNewDelegation(lr *types.LogRecord, delegator common.Address, toValidatorID *big.Int, amount *big.Int) {
	// get the validator address
	val, err := repo.ValidatorAddress((*hexutil.Big)(toValidatorID))
	if err != nil {
		log.Errorf("unknown validator #%d; %s", toValidatorID.Uint64(), err.Error())
		return
	}

	// make the delegation record
	dl := types.Delegation{
		Transaction:        lr.Trx.Hash,
		Address:            delegator,
		ToValidatorID:      (*hexutil.Big)(toValidatorID),
		ToValidatorAddress: *val,
		Amount:             (*hexutil.Big)(amount),
		CreatedTime:        time.Unix(int64(lr.Block.TimeStamp), 0),
	}

	// store the delegation
	if err = repo.StoreDelegation(&dl); err != nil {
		log.Errorf("failed to store delegation; %s", err.Error())
	}
}

// handleSfcCreatedDelegation handles a new delegation event from SFC v1 and SFC v2 contract
// and also the new delegation event from SFC3 contract with the same structure.
// event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func handleSfcCreatedDelegation(lr *types.LogRecord) {
	handleNewDelegation(
		lr,
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Data),
	)
}

// handleSfcUndelegated handles new withdrawal request from SFCv3 contract.
// We ignore withdrawals from previous SFC versions since after the upgrade all the pending
// withdrawals will be settled.
// event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcUndelegated(lr *types.LogRecord) {
	// sanity check for data (1x uint256 = 32 bytes)
	if len(lr.Data) != 32 {
		log.Criticalf("%s lr invalid data length; expected 32 bytes, %d bytes given, %d topics given",
			lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// create withdraw request
	handleNewWithdrawRequest(
		lr,
		types.WithdrawTypeUndelegated,
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		new(big.Int).SetBytes(lr.Data[:]),
	)
}

// handleNewWithdrawRequest will create a new withdrawal request for the given stake.
func handleNewWithdrawRequest(lr *types.LogRecord, wrt string, delegator common.Address, toValidatorID *big.Int,
	requestID *big.Int, amount *big.Int) {
	// make the request
	wr := types.WithdrawRequest{
		Type:              wrt,
		RequestTrx:        lr.TxHash,
		WithdrawRequestID: (*hexutil.Big)(requestID),
		Address:           delegator,
		StakerID:          (*hexutil.Big)(toValidatorID),
		CreatedTime:       lr.Block.TimeStamp,
		Amount:            (*hexutil.Big)(amount),
	}

	// lr what we do
	log.Debugf("new withdrawal of type %s by %s to #%d, req %s at %s",
		wrt,
		delegator.String(),
		toValidatorID.Uint64(),
		((*hexutil.Big)(requestID)).String(),
		lr.TxHash.String(),
	)

	// store the request
	if err := repo.StoreWithdrawRequest(&wr); err != nil {
		log.Errorf("failed to store new withdraw request; %s", err.Error())
	}

	// check active amount on the delegation
	if err := repo.UpdateDelegationBalance(&wr.Address, wr.StakerID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &wr.Address, wr.StakerID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleFinishedWithdrawRequest handles withdrawal request finalisation event.
func handleFinishedWithdrawRequest(lr *types.LogRecord, delegator common.Address, toValidatorID *big.Int,
	requestID *big.Int, amount *big.Int, penalty *big.Int) {
	// make sure the delegation balance will be updated
	defer func() {
		// check active amount on the delegation
		if err := repo.UpdateDelegationBalance(&delegator, (*hexutil.Big)(toValidatorID), func(amo *big.Int) error {
			return makeAdHocDelegation(lr, &delegator, (*hexutil.Big)(toValidatorID), amo)
		}); err != nil {
			log.Errorf("failed to update delegation; %s", err.Error())
		}
	}()

	// lr what we do
	log.Debugf("closing withdrawal by %s to #%d, req %s at %s",
		delegator.String(),
		toValidatorID.Uint64(),
		((*hexutil.Big)(requestID)).String(),
		lr.TxHash.String(),
	)

	// try to get the request from database
	req, err := repo.WithdrawRequest(&delegator, (*hexutil.Big)(toValidatorID), (*hexutil.Big)(requestID))
	if err != nil {
		log.Errorf("can not load withdraw requests to finalise; %s", err.Error())
		return
	}

	// update the request to have the finalization details
	req.WithdrawTime = &lr.Block.TimeStamp
	req.WithdrawTrx = &lr.TxHash
	req.Amount = (*hexutil.Big)(amount)
	req.Penalty = (*hexutil.Big)(penalty)

	// store the updated request
	if err = repo.UpdateWithdrawRequest(req); err != nil {
		log.Errorf("failed to store finalized withdraw request; %s", err.Error())
	}
}

// handleSfcWithdrawn handles a withdrawal request finalization event.
// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount, uint256 penalty)
func handleSfcWithdrawn(lr *types.LogRecord) {
	// sanity check for data (4x topic + 1x + 2 x uint256 = 64 bytes)
	if len(lr.Topics) != 4 || len(lr.Data) != 64 {
		log.Criticalf("%s is not event Withdrawn; expected 64 bytes, %d bytes given; expected 4 topics, %d given",
			lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		lr,
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		new(big.Int).SetBytes(lr.Data[:32]),
		new(big.Int).SetBytes(lr.Data[32:]),
	)
}

// event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func handleSfcWithdrawnWoPenalty(lr *types.LogRecord) {
	// sanity check for data (4x topic + 1x + 1 x uint256 = 32 bytes)
	if len(lr.Topics) != 4 || len(lr.Data) != 32 {
		log.Criticalf("%s is not event Withdrawn; expected 32 bytes, %d bytes given; expected 4 topics, %d given",
			lr.TxHash.String(), len(lr.Data), len(lr.Topics))
		return
	}

	// finish the request
	handleFinishedWithdrawRequest(
		lr,
		common.BytesToAddress(lr.Topics[1].Bytes()),
		new(big.Int).SetBytes(lr.Topics[2].Bytes()),
		new(big.Int).SetBytes(lr.Topics[3].Bytes()),
		new(big.Int).SetBytes(lr.Data[:32]),
		new(big.Int),
	)
}

// makeAdHocDelegation creates a new delegation in case an expected existing delegation
// could not be found on a new lr event processing.
func makeAdHocDelegation(lr *types.LogRecord, delegator *common.Address, toValidatorID *hexutil.Big,
	amount *big.Int) error {
	val, err := repo.ValidatorAddress(toValidatorID)
	if err != nil {
		return err
	}

	// do the insert
	log.Noticef("creating ad-hoc delegation of %s to #%d", delegator.String(), toValidatorID.ToInt().Uint64())
	return repo.StoreDelegation(&types.Delegation{
		Transaction:        lr.TxHash,
		Address:            *delegator,
		ToValidatorID:      toValidatorID,
		ToValidatorAddress: *val,
		Amount:             (*hexutil.Big)(amount),
		CreatedTime:        time.Unix(int64(lr.Block.TimeStamp), 0),
	})
}
