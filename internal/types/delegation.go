// Package types implements different core types of the API.
package types

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

const (
	// FiDelegationPk defines primary key column of the delegation table.
	FiDelegationPk = "_id"

	// FiDelegationOrdinal defines ordinal index column of the delegation table.
	FiDelegationOrdinal = "orx"

	// FiDelegationAddress defines delegation address column of the delegation table.
	FiDelegationAddress = "addr"

	// FiDelegationToValidator defines id of the validator column of the delegation table.
	FiDelegationToValidator = "to"

	// FiDelegationToValidatorAddress defines validator address column of the delegation table.
	FiDelegationToValidatorAddress = "to_addr"

	// FiDelegationAmountActive defines amount delegated column of the delegation table.
	FiDelegationAmountActive = "amount"

	// FiDelegationValue defines value of the delegation column of the delegation table.
	FiDelegationValue = "value"

	// FiDelegationTransaction defines transaction has column of the delegation table.
	FiDelegationTransaction = "trx"

	// FiDelegationStamp defines time stamp column of the delegation table.
	FiDelegationStamp = "stamp"
)

// Delegation represents a delegator in Opera blockchain.
type Delegation struct {
	ID                 string         `json:"id"`
	Transaction        common.Hash    `json:"trx"`
	Address            common.Address `json:"address"`
	ToValidatorID      *hexutil.Big   `json:"toStakerID"`
	ToValidatorAddress common.Address `json:"toStakerAddr"`
	CreatedTime        time.Time      `json:"createdTime"`
	Index              uint64         `json:"ordinalIndex"`

	// Amount represents the original staked amount
	Amount *hexutil.Big `json:"amountStaked"`
}

// DelegationDecimalsCorrection is used to adjust decimal precision of a delegation active value.
var DelegationDecimalsCorrection = new(big.Int).SetUint64(1000000000)

// OrdinalIndex returns an ordinal index for the given delegation.
// We construct the UID from the time the delegation was created (40 bits = 1099511627775s = 34000 years),
// a part of the creation transaction hash and part of the target validator index (12 bits = 4096).
func (dl *Delegation) OrdinalIndex() uint64 {
	return (uint64(dl.CreatedTime.Unix())&0x7FFFFFFFFF)<<24 | (dl.ToValidatorID.ToInt().Uint64()&0xFFF)<<12 |
		(binary.BigEndian.Uint64(dl.Transaction[:8]) & 0xFFF)
}

// MarshalBSON creates a BSON representation of the delegation record.
func (dl *Delegation) MarshalBSON() ([]byte, error) {
	// calculate the value to 9 digits (and 18 billions remain available)
	val := new(big.Int).Div(dl.Amount.ToInt(), DelegationDecimalsCorrection).Uint64()

	// do BSON encoding
	return bson.Marshal(struct {
		Orx             uint64    `bson:"orx"`
		Trx             string    `bson:"trx"`
		Delegator       string    `bson:"addr"`
		ToValidator     string    `bson:"to"`
		ToValidatorAddr string    `bson:"to_addr"`
		Amount          string    `bson:"amount"`
		Value           uint64    `bson:"value"`
		Stamp           time.Time `bson:"stamp"`
	}{
		Orx:             dl.OrdinalIndex(),
		Trx:             dl.Transaction.String(),
		Delegator:       dl.Address.String(),
		ToValidator:     dl.ToValidatorID.String(),
		ToValidatorAddr: dl.ToValidatorAddress.String(),
		Amount:          dl.Amount.String(),
		Value:           val,
		Stamp:           dl.CreatedTime,
	})
}

// UnmarshalBSON updates the value from BSON source.
func (dl *Delegation) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode BIG number in delegation unmarshal")
		}
	}()

	// try to decode the BSON data
	var row struct {
		ID     string    `bson:"_id"`
		Orx    uint64    `bson:"orx"`
		Trx    string    `bson:"trx"`
		Addr   string    `bson:"addr"`
		To     string    `bson:"to"`
		ToAddr string    `bson:"to_addr"`
		Staked string    `bson:"amount"`
		Value  uint64    `bson:"value"`
		Stamp  time.Time `bson:"stamp"`
	}
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer values
	dl.ID = row.ID
	dl.Transaction = common.HexToHash(row.Trx)
	dl.Address = common.HexToAddress(row.Addr)
	dl.ToValidatorID = (*hexutil.Big)(hexutil.MustDecodeBig(row.To))
	dl.ToValidatorAddress = common.HexToAddress(row.ToAddr)
	dl.CreatedTime = row.Stamp
	dl.Amount = (*hexutil.Big)(hexutil.MustDecodeBig(row.Staked))
	dl.Index = row.Orx
	return nil
}
