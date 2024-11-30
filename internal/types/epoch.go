// Package types implements different core types of the API.
package types

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"time"
)

// Epoch represents epoch detail.
type Epoch struct {
	Id                    hexutil.Uint64 `json:"id"`
	EndTime               hexutil.Uint64 `json:"end"`
	EndBlock              hexutil.Uint64 `json:"block"`
	EpochFee              hexutil.Big    `json:"fee"`
	EpochFeeBurn          hexutil.Big    `json:"burn"`
	EpochFeeTreasury      hexutil.Big    `json:"treasury"`
	TotalBaseRewardWeight hexutil.Big    `json:"trw"`
	TotalTxRewardWeight   hexutil.Big    `json:"txw"`
	BaseRewardPerSecond   hexutil.Big    `json:"brw"`
	StakeTotalAmount      hexutil.Big    `json:"stake"`
	TotalSupply           hexutil.Big    `json:"supply"`
}

// BsonEpoch represents the epoch data structure for BSON formatting.
type BsonEpoch struct {
	ID                  int64     `bson:"_id"`
	EndTime             int64     `bson:"et"`
	EndBlock            int64     `bson:"block"`
	End                 time.Time `bson:"end"`
	Fee                 string    `bson:"fee"`
	FeeBurn             string    `bson:"feb"`
	FeeTreasury         string    `bson:"fet"`
	BaseRewardWeight    string    `bson:"brw"`
	TrxRewardWeight     string    `bson:"trw"`
	BaseRewardPerSecond string    `bson:"rew"`
	TotalStake          string    `bson:"stake"`
	TotalSupply         string    `bson:"supply"`
	Burned              int64     `bson:"burned"`
	Treasured           int64     `bson:"treasured"`
}

// UnmarshalEpoch parses the JSON-encoded Epoch data.
func UnmarshalEpoch(data []byte) (*Epoch, error) {
	var ep Epoch
	err := json.Unmarshal(data, &ep)
	return &ep, err
}

// Marshal returns the JSON encoding of Epoch.
func (e *Epoch) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// MarshalBSON creates a BSON representation of the Epoch record.
func (e *Epoch) MarshalBSON() ([]byte, error) {
	amBurned := new(big.Int).Div(e.EpochFeeBurn.ToInt(), BurnDecimalsCorrection)
	amTreasured := new(big.Int).Div(e.EpochFeeTreasury.ToInt(), BurnDecimalsCorrection)

	// prep the structure for saving
	return bson.Marshal(BsonEpoch{
		ID:                  int64(e.Id),
		EndTime:             int64(e.EndTime),
		EndBlock:            int64(e.EndBlock),
		End:                 time.Unix(int64(e.EndTime), 0),
		Fee:                 e.EpochFee.String(),
		FeeBurn:             e.EpochFeeBurn.String(),
		FeeTreasury:         e.EpochFeeTreasury.String(),
		BaseRewardWeight:    e.TotalBaseRewardWeight.String(),
		TrxRewardWeight:     e.TotalTxRewardWeight.String(),
		BaseRewardPerSecond: e.BaseRewardPerSecond.String(),
		TotalStake:          e.StakeTotalAmount.String(),
		TotalSupply:         e.TotalSupply.String(),
		Burned:              amBurned.Int64(),
		Treasured:           amTreasured.Int64(),
	})
}

// UnmarshalBSON updates the value from BSON source.
func (e *Epoch) UnmarshalBSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("can not decode stored epoch")
		}
	}()

	// try to decode BSON data
	var row BsonEpoch
	if err = bson.Unmarshal(data, &row); err != nil {
		return err
	}

	// transfer the data points
	e.Id = (hexutil.Uint64)(row.ID)
	e.EndTime = (hexutil.Uint64)(row.EndTime)
	e.EndBlock = (hexutil.Uint64)(row.EndBlock)
	e.EpochFee = (hexutil.Big)(*hexutil.MustDecodeBig(row.Fee))
	e.EpochFeeBurn = (hexutil.Big)(*hexutil.MustDecodeBig(row.FeeBurn))
	e.EpochFeeTreasury = (hexutil.Big)(*hexutil.MustDecodeBig(row.FeeTreasury))
	e.TotalBaseRewardWeight = (hexutil.Big)(*hexutil.MustDecodeBig(row.BaseRewardWeight))
	e.TotalTxRewardWeight = (hexutil.Big)(*hexutil.MustDecodeBig(row.TrxRewardWeight))
	e.BaseRewardPerSecond = (hexutil.Big)(*hexutil.MustDecodeBig(row.BaseRewardPerSecond))
	e.StakeTotalAmount = (hexutil.Big)(*hexutil.MustDecodeBig(row.TotalStake))
	e.TotalSupply = (hexutil.Big)(*hexutil.MustDecodeBig(row.TotalSupply))
	return nil
}
