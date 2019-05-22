package document

import (
	"github.com/irisnet/irishub-sync/store"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleDelegator = "stake_role_delegator"
)

type Delegator struct {
	Address        string  `bson:"address"`
	ValidatorAddr  string  `bson:"validator_addr"` // validatorAddr
	Shares         float64 `bson:"shares"`
	OriginalShares string  `bson:"original_shares"`
	BondedHeight   int64   `bson:"height"`

	UnbondingDelegation UnbondingDelegation `bson:"unbonding_delegation"`
}

// UnbondingDelegation reflects a delegation's passive unbonding queue.
type UnbondingDelegation struct {
	CreationHeight int64       `bson:"creation_height"` // height which the unbonding took place
	MinTime        int64       `bson:"min_time"`        // unix time for unbonding completion
	InitialBalance store.Coins `bson:"initial_balance"` // atoms initially scheduled to receive at completion
	Balance        store.Coins `bson:"balance"`         // atoms to receive at completion
}

func (d Delegator) Name() string {
	return CollectionNmStakeRoleDelegator
}

func (d Delegator) PkKvPair() map[string]interface{} {
	return bson.M{"address": d.Address, "validator_addr": d.ValidatorAddr}
}
