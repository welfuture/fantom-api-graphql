// Package svc implements blockchain data processing services.
package svc

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// handleSfcRewardClaim handles a rewards claim event.
func handleSfcRewardClaim(lr *types.LogRecord, delegator common.Address, toValidatorID *hexutil.Big,
	amount *big.Int, isRestaked bool) {
	log.Debugf("%s claimed %d in stake to #%d", delegator.String(), amount.Uint64(),
		toValidatorID.ToInt().Uint64())

	// add the rewards claim into the repository
	if err := repo.StoreRewardClaim(&types.RewardClaim{
		Delegator:     delegator,
		ToValidatorId: *toValidatorID,
		Claimed:       lr.Block.TimeStamp,
		ClaimTrx:      lr.TxHash,
		Amount:        (hexutil.Big)(*amount),
		IsDelegated:   isRestaked,
	}); err != nil {
		log.Criticalf("can not store rewards claim; %s", err.Error())
		return
	}

	// check active amount on the delegation
	if err := repo.UpdateDelegationBalance(&delegator, toValidatorID, func(amo *big.Int) error {
		return makeAdHocDelegation(lr, &delegator, toValidatorID, amo)
	}); err != nil {
		log.Errorf("failed to update delegation; %s", err.Error())
	}
}

// handleSfcCommonRewardClaim handles the common reward claim on SFC contract.
// event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func handleSfcCommonRewardClaim(lr *types.LogRecord, isRestaked bool) {
	if len(lr.Data) != 32 {
		log.Criticalf("%s lr invalid data length; expected 32 bytes, given %d bytes",
			lr.TxHash.String(), len(lr.Data))
		return
	}

	delegator := common.BytesToAddress(lr.Topics[1].Bytes())
	toValidatorID := (*hexutil.Big)(new(big.Int).SetBytes(lr.Topics[2].Bytes()))
	amount := new(big.Int).SetBytes(lr.Data[:32])

	handleSfcRewardClaim(lr, delegator, toValidatorID, amount, isRestaked)
}

// handleSfcRestakedRewards handles a rewards re-stake event.
// event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func handleSfcRestakedRewards(lr *types.LogRecord) {
	handleSfcCommonRewardClaim(lr, true)
}

// handleSfcClaimedRewards handles a rewards re-stake event.
// event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func handleSfcClaimedRewards(lr *types.LogRecord) {
	handleSfcCommonRewardClaim(lr, false)
}
