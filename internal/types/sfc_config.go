package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// SfcConfig defines the current SFC contract configuration.
type SfcConfig struct {
	// minValidatorStake is the minimal amount of tokens required
	// to register a validator account with the default self stake.
	MinValidatorStake hexutil.Big

	// maxDelegatedRatio is the maximal ratio between a validator self stake
	// and the sum of all the received stakes of the validator.
	// The value is provided as a multiplier number with 18 decimals.
	MaxDelegatedRatio hexutil.Big

	// withdrawalPeriodEpochs is the minimal number of epochs
	// between an un-delegation and corresponding withdraw request.
	// The delay is enforced on withdraw call.
	WithdrawalPeriodEpochs hexutil.Big

	// withdrawalPeriodTime is the minimal number of seconds
	// between an un-delegation and corresponding withdraw request.
	// The delay is enforced on withdraw call.
	WithdrawalPeriodTime hexutil.Big
}

// Marshal encodes the config into bytes slice.
func (sc *SfcConfig) Marshal() ([]byte, error) {
	// we have 6x256bit numbers here
	buf := make([]byte, 4*32)

	// copy the bytes
	sc.MinValidatorStake.ToInt().FillBytes(buf[:32])
	sc.MaxDelegatedRatio.ToInt().FillBytes(buf[32:64])
	sc.WithdrawalPeriodEpochs.ToInt().FillBytes(buf[64:96])
	sc.WithdrawalPeriodTime.ToInt().FillBytes(buf[96:])
	return buf, nil
}

// Unmarshal decodes the buffer into the config set.
func (sc *SfcConfig) Unmarshal(buf []byte) error {
	// check for the buffer length, we expect 6*32 bytes
	if len(buf) != 128 {
		return fmt.Errorf("expected 128 bytes, %d received", len(buf))
	}

	// copy the data
	sc.MinValidatorStake.ToInt().SetBytes(buf[:32])
	sc.MaxDelegatedRatio.ToInt().SetBytes(buf[32:64])
	sc.WithdrawalPeriodEpochs.ToInt().SetBytes(buf[64:96])
	sc.WithdrawalPeriodTime.ToInt().SetBytes(buf[96:])
	return nil
}
