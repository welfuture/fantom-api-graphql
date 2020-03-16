// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// accMaxTransactionsPerRequest maximal number of transaction end-client can request in one query.
const accMaxTransactionsPerRequest = 50

// Account represents resolvable blockchain account structure.
type Account struct {
	repo       repository.Repository
	staker     *types.Staker
	delegation *types.Delegator
	balance    *hexutil.Big
	types.Account
}

// NewAccount builds new resolvable account structure.
func NewAccount(acc *types.Account, repo repository.Repository) *Account {
	return &Account{
		repo:    repo,
		Account: *acc,
	}
}

// Account resolves blockchain account by address.
func (rs *rootResolver) Account(args struct{ Address common.Address }) (*Account, error) {
	// simply pull the block by hash
	acc, err := rs.repo.Account(&args.Address)
	if err != nil {
		rs.log.Errorf("could not get the specified account")
		return nil, err
	}

	return NewAccount(acc, rs.repo), nil
}

// Resolves total number of active accounts on the blockchain.
func (rs *rootResolver) AccountsActive() (hexutil.Uint64, error) {
	return rs.repo.AccountsActive()
}

// Balance resolves total balance of the account.
func (acc *Account) Balance() (hexutil.Big, error) {
	if acc.balance == nil {
		// get the sender by address
		bal, err := acc.repo.AccountBalance(&acc.Account)
		if err != nil {
			return hexutil.Big{}, err
		}

		acc.balance = bal
	}

	return *acc.balance, nil
}

// TotalValue resolves address total value including delegated amount and pending rewards.
func (acc *Account) TotalValue() (hexutil.Big, error) {
	// get the balance
	balance, err := acc.Balance()
	if err != nil {
		return hexutil.Big{}, err
	}

	// try to get delegation
	del, err := acc.getDelegation()
	if err != nil {
		return balance, err
	}

	// do we have a delegation?
	if del != nil {
		// add delegated amount to the balance
		val := big.NewInt(0).Add(balance.ToInt(), del.Amount.ToInt())

		// get pending rewards
		rw, err := acc.repo.DelegationRewards(acc.Address)
		if err != nil {
			return hexutil.Big(*val), err
		}

		// add pending rewards to the final value
		val = big.NewInt(0).Add(val, rw.Amount.ToInt())
		return hexutil.Big(*val), err
	}

	return balance, nil
}

// TxCount resolves the number of transaction sent by the account, also known as nonce.
func (acc *Account) TxCount() (hexutil.Uint64, error) {
	// get the sender by address
	bal, err := acc.repo.AccountNonce(&acc.Account)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	return *bal, nil
}

// TxList resolves list of transaction associated with the account.
func (acc *Account) TxList(args struct {
	Cursor *Cursor
	Count  int32
}) (*TransactionList, error) {
	// limit count
	if args.Count > accMaxTransactionsPerRequest {
		args.Count = accMaxTransactionsPerRequest
	}

	// get the transaction hash list from repository
	bl, err := acc.repo.AccountTransactions(&acc.Account, (*string)(args.Cursor), args.Count)
	if err != nil {
		return nil, err
	}

	return NewTransactionList(bl, acc.repo), nil
}

// Staker resolves the account staker detail, if the account is a staker.
func (acc *Account) Staker() (*Staker, error) {
	// get the staker
	st, err := acc.getStaker()
	if err != nil {
		return nil, err
	}

	// not staker
	if st == nil {
		return nil, nil
	}

	return NewStaker(st, acc.repo), nil
}

// Delegation resolves the account delegator detail, if the account is a delegater.
func (acc *Account) Delegation() (*Delegator, error) {
	// try to get the delegator info
	dl, err := acc.repo.Delegation(acc.Address)
	if err != nil {
		return nil, err
	}

	// not delegator
	if dl == nil {
		return nil, nil
	}

	return NewDelegator(dl, acc.repo), nil
}

// getStaker returns lazy loaded staker information.
func (acc *Account) getStaker() (*types.Staker, error) {
	// try to get the staker info
	if acc.staker == nil {
		st, err := acc.repo.StakerByAddress(acc.Address)
		if err != nil {
			return nil, err
		}

		// is this a valid staker info?
		if st.Id <= 0 {
			return nil, nil
		}

		// keep the staker info
		acc.staker = st
	}

	return acc.staker, nil
}

// getDelegation return lazy loaded delegation detail for the account.
func (acc *Account) getDelegation() (*types.Delegator, error) {
	// try to get the staker info
	if acc.delegation == nil {
		// try to get the staker info
		dl, err := acc.repo.Delegation(acc.Address)
		if err != nil {
			return nil, err
		}

		// is this a valid staker info?
		if dl.ToStakerId <= 0 {
			return nil, nil
		}

		acc.delegation = dl
	}

	return acc.delegation, nil
}
