/*
Package repository implements repository for handling fast and efficient access to data required
by the resolvers of the API server.

Internally it utilizes RPC to access Opera full node for blockchain interaction. Mongo database
for fast, robust and scalable off-chain data storage, especially for aggregated and pre-calculated data mining
results. BigCache for in-memory object storage to speed up loading of frequently accessed entities.
*/
package repository

import (
	"errors"
	"fantom-api-graphql/internal/repository/db"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
)

// IsDelegating returns if the given address is an SFC delegator.
func (p *proxy) IsDelegating(addr *common.Address) (bool, error) {
	// count only active delegations (with non-zero value)
	count, err := p.db.DelegationsCountFiltered(&bson.D{
		{Key: types.FiDelegationAddress, Value: addr.String()},
		{Key: types.FiDelegationValue, Value: bson.D{{Key: "$gt", Value: 0}}},
	})
	if err != nil {
		p.log.Errorf("can not check delegation by address; %s", addr.String())
		return false, err
	}
	return 0 < count, nil
}

// StoreDelegation stores the delegation in persistent database.
func (p *proxy) StoreDelegation(dl *types.Delegation) error {
	return p.db.AddDelegation(dl)
}

// UpdateDelegationBalance updates active balance of the given delegation.
func (p *proxy) UpdateDelegationBalance(addr *common.Address, valID *hexutil.Big, unknownDelegation func(*big.Int) error) error {
	// pull the current value
	val, err := p.DelegationAmountStaked(addr, valID)
	if err != nil {
		p.log.Errorf("delegation balance not available for %s to %d; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
		return err
	}

	// do the update
	err = p.updateDelegationBalance(addr, valID, val)
	if err == nil {
		return nil
	}

	// unknown delegation detected?
	if errors.Is(err, db.ErrUnknownDelegation) {
		p.log.Debugf("delegation %s to #%d missing", addr.String(), valID.ToInt().Uint64())
		return unknownDelegation(val)
	}

	// some other error
	p.log.Errorf("delegation %s to %d update failed; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
	return err
}

// updateDelegationBalance performs delegation balance update if needed.
func (p *proxy) updateDelegationBalance(delegator *common.Address, toValidatorID *hexutil.Big, amount *big.Int) error {
	// get the delegation detail
	dlg, err := p.Delegation(delegator, toValidatorID)
	if err != nil {
		return err
	}

	// do we need to update? if the amount did not change, skip the update
	if dlg.Amount.ToInt().Cmp(amount) == 0 {
		return nil
	}

	// update the delegation in DB and memory
	dlg.Amount = (*hexutil.Big)(amount)
	err = p.db.UpdateDelegationBalance(delegator, toValidatorID, dlg.Amount)
	if nil == err {
		p.cache.PushDelegation(dlg)
	}
	return err
}

// Delegation returns a detail of delegation for the given address.
func (p *proxy) Delegation(adr *common.Address, valID *hexutil.Big) (*types.Delegation, error) {
	// log what we do
	p.log.Debugf("accessing delegation of %s to #%d", adr.String(), valID.ToInt().Uint64())

	// try cache first
	dlg := p.cache.PullDelegation(*adr, valID)
	if dlg != nil {
		return dlg, nil
	}

	// pull from DB instead; do we actually have it?
	dlg, err := p.db.Delegation(adr, valID)
	if err != nil {
		return nil, err
	}

	// store to cache for future reference
	p.cache.PushDelegation(dlg)
	return dlg, nil
}

// DelegationAmountStaked returns the current amount of staked tokens for the given delegation.
func (p *proxy) DelegationAmountStaked(addr *common.Address, valID *hexutil.Big) (*big.Int, error) {
	val, err := p.rpc.AmountStaked(addr, (*big.Int)(valID))
	if err != nil {
		p.log.Errorf("can not get amount delegated by %s to %d; %s", addr.String(), valID.ToInt().Uint64(), err.Error())
		return nil, err
	}
	// log and return
	p.log.Debugf("%s delegated %d to %d", addr.String(), val.Uint64(), valID.ToInt().Uint64())
	return val, nil
}

// DelegationsByAddress returns a list of all delegations of a given delegator address.
func (p *proxy) DelegationsByAddress(addr *common.Address, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of %s", addr.String())
	return p.db.Delegations(cursor, count, &bson.D{{Key: types.FiDelegationAddress, Value: addr.String()}})
}

// DelegationsByAddressAll returns a list of all delegations of the given address un-paged.
func (p *proxy) DelegationsByAddressAll(addr *common.Address) ([]*types.Delegation, error) {
	p.log.Debugf("loading all delegations of %s", addr.String())
	return p.db.DelegationsAll(&bson.D{{Key: types.FiDelegationAddress, Value: addr.String()}})
}

// DelegationsOfValidator extract a list of delegations for a given validator.
func (p *proxy) DelegationsOfValidator(valID *hexutil.Big, cursor *string, count int32) (*types.DelegationList, error) {
	p.log.Debugf("loading delegations of #%d", valID.ToInt().Uint64())
	return p.db.Delegations(cursor, count, &bson.D{{Key: types.FiDelegationToValidator, Value: valID.String()}})
}

// PendingRewards returns a detail of pending rewards for the given delegation address and validator ID.
func (p *proxy) PendingRewards(addr *common.Address, valID *hexutil.Big) (*types.PendingRewards, error) {
	p.log.Debugf("loading pending rewards of %s to #%d", addr.String(), valID.ToInt().Uint64())
	return p.rpc.PendingRewards(addr, valID.ToInt())
}

// DelegationFluidStakingActive signals if the delegation is upgraded to Fluid Staking model.
func (p *proxy) DelegationFluidStakingActive(_ *common.Address, _ *hexutil.Big) (bool, error) {
	return true, nil
}
