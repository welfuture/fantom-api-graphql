/*
Package rpc implements bridge to Opera full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Opera node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Opera RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Opera RPC interface with connection limited to specified endpoints.

We strongly discourage opening Opera RPC interface for unrestricted Internet access.
*/
package rpc

//go:generate tools/abigen.sh --abi ./contracts/abi/sfc_4.0.0.json --pkg contracts --type SfcV400Contract --out ./contracts/sfc-v400.go

import (
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// SfcVersion returns current version of the SFC contract as a single number.
func (ftm *FtmBridge) SfcVersion() (hexutil.Uint64, error) {
	// get the version information from the contract
	var ver [3]byte
	var err error
	ver, err = ftm.SfcContract().Version(nil)
	if err != nil {
		ftm.log.Criticalf("failed to get the SFC version; %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64((uint64(ver[0]) << 16) | (uint64(ver[1]) << 8) | uint64(ver[2])), nil
}

// CurrentEpoch extract the current epoch id from SFC smart contract.
func (ftm *FtmBridge) CurrentEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := ftm.SfcContract().CurrentEpoch(ftm.DefaultCallOpts())
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// CurrentSealedEpoch extract the current sealed epoch id from SFC smart contract.
func (ftm *FtmBridge) CurrentSealedEpoch() (hexutil.Uint64, error) {
	// get the value from the contract
	epoch, err := ftm.SfcContract().CurrentSealedEpoch(ftm.DefaultCallOpts())
	if err != nil {
		ftm.log.Errorf("failed to get the current sealed epoch: %s", err.Error())
		return 0, err
	}
	return hexutil.Uint64(epoch.Uint64()), nil
}

// Epoch extract information about an epoch from SFC smart contract.
func (ftm *FtmBridge) Epoch(id hexutil.Uint64) (*types.Epoch, error) {
	// extract epoch snapshot
	epo, err := ftm.SfcContract().GetEpochSnapshot(nil, big.NewInt(int64(id)))
	if err != nil {
		ftm.log.Errorf("failed to extract epoch information: %s", err.Error())
		return nil, err
	}

	return &types.Epoch{
		Id:                  id,
		EndTime:             hexutil.Uint64(epo.EndTime.Uint64()),
		EndBlock:            hexutil.Uint64(epo.EndBlock.Uint64()),
		EpochFee:            (hexutil.Big)(*epo.EpochFee),
		BaseRewardPerSecond: (hexutil.Big)(*epo.BaseRewardPerSecond),
		StakeTotalAmount:    (hexutil.Big)(*epo.TotalStake),
		TotalSupply:         (hexutil.Big)(*epo.TotalSupply),
	}, nil
}

// RewardsAllowed returns if the rewards can be manipulated with.
func (ftm *FtmBridge) RewardsAllowed() (bool, error) {
	ftm.log.Debug("rewards lock always open")
	return true, nil
}

// TotalStaked returns the total amount of staked tokens.
func (ftm *FtmBridge) TotalStaked() (*big.Int, error) {
	return ftm.SfcContract().TotalStake(ftm.DefaultCallOpts())
}

// SfcMinValidatorStake extracts a value of minimal validator self stake.
func (ftm *FtmBridge) SfcMinValidatorStake() (*big.Int, error) {
	return ftm.sfcShards.minSelfStake()
}

// SfcMaxDelegatedRatio extracts a ratio between self delegation and received stake.
func (ftm *FtmBridge) SfcMaxDelegatedRatio() (*big.Int, error) {
	return ftm.sfcShards.maxDelegatedRatio()
}

// SfcWithdrawalPeriodEpochs extracts a minimal number of epochs between un-delegate and withdraw.
func (ftm *FtmBridge) SfcWithdrawalPeriodEpochs() (*big.Int, error) {
	return ftm.sfcShards.withdrawalPeriodEpochs()
}

// SfcWithdrawalPeriodTime extracts a minimal number of seconds between un-delegate and withdraw.
func (ftm *FtmBridge) SfcWithdrawalPeriodTime() (*big.Int, error) {
	return ftm.sfcShards.withdrawalPeriodTime()
}
