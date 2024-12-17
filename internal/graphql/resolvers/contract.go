// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"regexp"
)

const (
	// scValidationMinSourceCodeLength is the minimum length of a validated
	// smart contract source code.
	scMinSourceCodeLength = 32

	// scMaxNameLength is the maximum accepted length of a smart contract name.
	scMaxNameLength = 64

	// scMaxNameLength is the maximum accepted length of a smart contract version.
	scMaxVersionLength = 14

	// scMaxSupportLinkLength is the maximum accepted length of smart contract
	// support link.
	scMaxSupportLinkLength = 64
)

// scVersionSyntaxRegexp represents a regular expression for testing smart contract
// version string syntax. We enforce specific syntax on provided contract versions.
var scVersionSyntaxRegexp = regexp.MustCompile("^\\w?(\\d+\\.)+\\d+$")

// Contract represents resolvable blockchain smart contract structure.
type Contract struct {
	types.Contract
}

// ContractValidationInput represents an input structure used
// to validate contract source code against deployed contract byte code.
type ContractValidationInput struct {
	// Address represents the deployment address of the contract being validated.
	Address common.Address `json:"address"`

	// Name represents an optional name of the contract.
	Name *string `json:"name,omitempty"`

	// Version represents an optional version of the contract.
	// We assume version to be constructed from numbers and dots
	// with optional character at the beginning.
	// I.e. "v1.5.17"
	Version *string `json:"version,omitempty"`

	// SupportContact represents an optional contact information
	// the contract validator wants to publish with the contract
	// details.
	SupportContact *string `json:"supportContact,omitempty"`

	// License represents an optional contact open source license
	// being used.
	License *string `json:"license,omitempty"`

	// IsOptimized signals that the contract byte code was optimized
	// during compilation.
	Optimized bool `json:"optimized"`

	// OptimizeRuns represents number of optimization runs used
	// during the contract compilation.
	OptimizeRuns int32 `json:"optimizeRuns"`

	// SourceCode represents the Solidity source code to be validated.
	SourceCode string `json:"sourceCode"`
}

// NewContract builds new resolvable smart contract structure.
func NewContract(con *types.Contract) *Contract {
	if con == nil {
		return nil
	}
	return &Contract{Contract: *con}
}

// DeployedBy resolves the deployment transaction of the contract.
func (con *Contract) DeployedBy() (*Transaction, error) {
	tr, err := repository.R().Transaction(&con.TransactionHash)
	return NewTransaction(tr), err
}
