// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SfcV400ContractMetaData contains all meta data concerning the SfcV400Contract contract.
var SfcV400ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRedirected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSelfStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MalformedPubkey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDeactivatedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDriverAuth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughEpochsPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughTimePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToStash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PubkeyUsedByOtherValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Redirected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefundRatioTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameRedirectionAuthorizer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeIsFullySlashed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeSubscriberFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransfersNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDelegationLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotSlashed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewards\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AnnouncedRedirection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BurntFTM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ChangedValidatorStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"CreatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"}],\"name\":\"DeactivatedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundedSlashedLegacyDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"RestakedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"UpdatedSlashingRefundRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"syncPubkey\",\"type\":\"bool\"}],\"name\":\"_syncValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"announceRedirection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFTM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"constsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"createValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentSealedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedOriginatedTxsFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAccumulatedUptime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochAverageUptime\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochOfflineTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getEpochReceivedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRewardPerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getEpochValidatorIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getRedirection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getRedirectionRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deactivatedEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"}],\"name\":\"getValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"getValidatorPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"getWithdrawalRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sealedEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeDriver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"initiateRedirection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"isSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubkeyAddress\",\"type\":\"address\"}],\"name\":\"pubkeyAddressToValidatorID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"redirect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redirectionAuthorizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"restakeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"rewardsStash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"offlineTime\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"offlineBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"originatedTxsFee\",\"type\":\"uint256[]\"}],\"name\":\"sealEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nextValidatorIDs\",\"type\":\"uint256[]\"}],\"name\":\"sealEpochValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"setGenesisDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"auth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"createdTime\",\"type\":\"uint256\"}],\"name\":\"setGenesisValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"setRedirectionAuthorizer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"slashingRefundRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeSubscriberAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"}],\"name\":\"stashRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"}],\"name\":\"stashedRewardsUntilEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalActiveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateConstsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundRatio\",\"type\":\"uint256\"}],\"name\":\"updateSlashingRefundRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateStakeSubscriberAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"updateTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"toValidatorID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wrID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SfcV400ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SfcV400ContractMetaData.ABI instead.
var SfcV400ContractABI = SfcV400ContractMetaData.ABI

// SfcV400Contract is an auto generated Go binding around an Ethereum contract.
type SfcV400Contract struct {
	SfcV400ContractCaller     // Read-only binding to the contract
	SfcV400ContractTransactor // Write-only binding to the contract
	SfcV400ContractFilterer   // Log filterer for contract events
}

// SfcV400ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SfcV400ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SfcV400ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SfcV400ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SfcV400ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SfcV400ContractSession struct {
	Contract     *SfcV400Contract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SfcV400ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SfcV400ContractCallerSession struct {
	Contract *SfcV400ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SfcV400ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SfcV400ContractTransactorSession struct {
	Contract     *SfcV400ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SfcV400ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SfcV400ContractRaw struct {
	Contract *SfcV400Contract // Generic contract binding to access the raw methods on
}

// SfcV400ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SfcV400ContractCallerRaw struct {
	Contract *SfcV400ContractCaller // Generic read-only contract binding to access the raw methods on
}

// SfcV400ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SfcV400ContractTransactorRaw struct {
	Contract *SfcV400ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSfcV400Contract creates a new instance of SfcV400Contract, bound to a specific deployed contract.
func NewSfcV400Contract(address common.Address, backend bind.ContractBackend) (*SfcV400Contract, error) {
	contract, err := bindSfcV400Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SfcV400Contract{SfcV400ContractCaller: SfcV400ContractCaller{contract: contract}, SfcV400ContractTransactor: SfcV400ContractTransactor{contract: contract}, SfcV400ContractFilterer: SfcV400ContractFilterer{contract: contract}}, nil
}

// NewSfcV400ContractCaller creates a new read-only instance of SfcV400Contract, bound to a specific deployed contract.
func NewSfcV400ContractCaller(address common.Address, caller bind.ContractCaller) (*SfcV400ContractCaller, error) {
	contract, err := bindSfcV400Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractCaller{contract: contract}, nil
}

// NewSfcV400ContractTransactor creates a new write-only instance of SfcV400Contract, bound to a specific deployed contract.
func NewSfcV400ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SfcV400ContractTransactor, error) {
	contract, err := bindSfcV400Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractTransactor{contract: contract}, nil
}

// NewSfcV400ContractFilterer creates a new log filterer instance of SfcV400Contract, bound to a specific deployed contract.
func NewSfcV400ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SfcV400ContractFilterer, error) {
	contract, err := bindSfcV400Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractFilterer{contract: contract}, nil
}

// bindSfcV400Contract binds a generic wrapper to an already deployed contract.
func bindSfcV400Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SfcV400ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV400Contract *SfcV400ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV400Contract.Contract.SfcV400ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV400Contract *SfcV400ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SfcV400ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV400Contract *SfcV400ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SfcV400ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SfcV400Contract *SfcV400ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SfcV400Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SfcV400Contract *SfcV400ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SfcV400Contract *SfcV400ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SfcV400Contract *SfcV400ContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SfcV400Contract *SfcV400ContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SfcV400Contract.Contract.UPGRADEINTERFACEVERSION(&_SfcV400Contract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_SfcV400Contract *SfcV400ContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _SfcV400Contract.Contract.UPGRADEINTERFACEVERSION(&_SfcV400Contract.CallOpts)
}

// ConstsAddress is a free data retrieval call binding the contract method 0xd46fa518.
//
// Solidity: function constsAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCaller) ConstsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "constsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConstsAddress is a free data retrieval call binding the contract method 0xd46fa518.
//
// Solidity: function constsAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractSession) ConstsAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.ConstsAddress(&_SfcV400Contract.CallOpts)
}

// ConstsAddress is a free data retrieval call binding the contract method 0xd46fa518.
//
// Solidity: function constsAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCallerSession) ConstsAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.ConstsAddress(&_SfcV400Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) CurrentEpoch() (*big.Int, error) {
	return _SfcV400Contract.Contract.CurrentEpoch(&_SfcV400Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _SfcV400Contract.Contract.CurrentEpoch(&_SfcV400Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) CurrentSealedEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "currentSealedEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcV400Contract.Contract.CurrentSealedEpoch(&_SfcV400Contract.CallOpts)
}

// CurrentSealedEpoch is a free data retrieval call binding the contract method 0x7cacb1d6.
//
// Solidity: function currentSealedEpoch() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) CurrentSealedEpoch() (*big.Int, error) {
	return _SfcV400Contract.Contract.CurrentSealedEpoch(&_SfcV400Contract.CallOpts)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochAccumulatedOriginatedTxsFee(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochAccumulatedOriginatedTxsFee", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedOriginatedTxsFee is a free data retrieval call binding the contract method 0xdc31e1af.
//
// Solidity: function getEpochAccumulatedOriginatedTxsFee(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochAccumulatedOriginatedTxsFee(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedOriginatedTxsFee(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochAccumulatedRewardPerToken(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochAccumulatedRewardPerToken", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedRewardPerToken(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedRewardPerToken is a free data retrieval call binding the contract method 0x61e53fcc.
//
// Solidity: function getEpochAccumulatedRewardPerToken(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochAccumulatedRewardPerToken(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedRewardPerToken(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochAccumulatedUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochAccumulatedUptime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedUptime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAccumulatedUptime is a free data retrieval call binding the contract method 0xdf00c922.
//
// Solidity: function getEpochAccumulatedUptime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochAccumulatedUptime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochAccumulatedUptime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAverageUptime is a free data retrieval call binding the contract method 0xaa5d8292.
//
// Solidity: function getEpochAverageUptime(uint256 epoch, uint256 validatorID) view returns(uint64)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochAverageUptime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (uint64, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochAverageUptime", epoch, validatorID)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetEpochAverageUptime is a free data retrieval call binding the contract method 0xaa5d8292.
//
// Solidity: function getEpochAverageUptime(uint256 epoch, uint256 validatorID) view returns(uint64)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochAverageUptime(epoch *big.Int, validatorID *big.Int) (uint64, error) {
	return _SfcV400Contract.Contract.GetEpochAverageUptime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochAverageUptime is a free data retrieval call binding the contract method 0xaa5d8292.
//
// Solidity: function getEpochAverageUptime(uint256 epoch, uint256 validatorID) view returns(uint64)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochAverageUptime(epoch *big.Int, validatorID *big.Int) (uint64, error) {
	return _SfcV400Contract.Contract.GetEpochAverageUptime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochEndBlock is a free data retrieval call binding the contract method 0xdb5ca3e5.
//
// Solidity: function getEpochEndBlock(uint256 epoch) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochEndBlock(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochEndBlock", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochEndBlock is a free data retrieval call binding the contract method 0xdb5ca3e5.
//
// Solidity: function getEpochEndBlock(uint256 epoch) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochEndBlock(epoch *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochEndBlock(&_SfcV400Contract.CallOpts, epoch)
}

// GetEpochEndBlock is a free data retrieval call binding the contract method 0xdb5ca3e5.
//
// Solidity: function getEpochEndBlock(uint256 epoch) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochEndBlock(epoch *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochEndBlock(&_SfcV400Contract.CallOpts, epoch)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochOfflineBlocks(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochOfflineBlocks", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochOfflineBlocks(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineBlocks is a free data retrieval call binding the contract method 0xa198d229.
//
// Solidity: function getEpochOfflineBlocks(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochOfflineBlocks(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochOfflineBlocks(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochOfflineTime(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochOfflineTime", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochOfflineTime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochOfflineTime is a free data retrieval call binding the contract method 0xe261641a.
//
// Solidity: function getEpochOfflineTime(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochOfflineTime(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochOfflineTime(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochReceivedStake(opts *bind.CallOpts, epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochReceivedStake", epoch, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochReceivedStake(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochReceivedStake is a free data retrieval call binding the contract method 0x58f95b80.
//
// Solidity: function getEpochReceivedStake(uint256 epoch, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochReceivedStake(epoch *big.Int, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochReceivedStake(&_SfcV400Contract.CallOpts, epoch, validatorID)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 epoch) view returns(uint256 endTime, uint256 endBlock, uint256 epochFee, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochSnapshot(opts *bind.CallOpts, epoch *big.Int) (struct {
	EndTime             *big.Int
	EndBlock            *big.Int
	EpochFee            *big.Int
	BaseRewardPerSecond *big.Int
	TotalStake          *big.Int
	TotalSupply         *big.Int
}, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochSnapshot", epoch)

	outstruct := new(struct {
		EndTime             *big.Int
		EndBlock            *big.Int
		EpochFee            *big.Int
		BaseRewardPerSecond *big.Int
		TotalStake          *big.Int
		TotalSupply         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BaseRewardPerSecond = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalStake = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 epoch) view returns(uint256 endTime, uint256 endBlock, uint256 epochFee, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_SfcV400Contract *SfcV400ContractSession) GetEpochSnapshot(epoch *big.Int) (struct {
	EndTime             *big.Int
	EndBlock            *big.Int
	EpochFee            *big.Int
	BaseRewardPerSecond *big.Int
	TotalStake          *big.Int
	TotalSupply         *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetEpochSnapshot(&_SfcV400Contract.CallOpts, epoch)
}

// GetEpochSnapshot is a free data retrieval call binding the contract method 0x39b80c00.
//
// Solidity: function getEpochSnapshot(uint256 epoch) view returns(uint256 endTime, uint256 endBlock, uint256 epochFee, uint256 baseRewardPerSecond, uint256 totalStake, uint256 totalSupply)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochSnapshot(epoch *big.Int) (struct {
	EndTime             *big.Int
	EndBlock            *big.Int
	EpochFee            *big.Int
	BaseRewardPerSecond *big.Int
	TotalStake          *big.Int
	TotalSupply         *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetEpochSnapshot(&_SfcV400Contract.CallOpts, epoch)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_SfcV400Contract *SfcV400ContractCaller) GetEpochValidatorIDs(opts *bind.CallOpts, epoch *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getEpochValidatorIDs", epoch)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_SfcV400Contract *SfcV400ContractSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochValidatorIDs(&_SfcV400Contract.CallOpts, epoch)
}

// GetEpochValidatorIDs is a free data retrieval call binding the contract method 0xb88a37e2.
//
// Solidity: function getEpochValidatorIDs(uint256 epoch) view returns(uint256[])
func (_SfcV400Contract *SfcV400ContractCallerSession) GetEpochValidatorIDs(epoch *big.Int) ([]*big.Int, error) {
	return _SfcV400Contract.Contract.GetEpochValidatorIDs(&_SfcV400Contract.CallOpts, epoch)
}

// GetRedirection is a free data retrieval call binding the contract method 0x736de9ae.
//
// Solidity: function getRedirection(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractCaller) GetRedirection(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getRedirection", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRedirection is a free data retrieval call binding the contract method 0x736de9ae.
//
// Solidity: function getRedirection(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractSession) GetRedirection(delegator common.Address) (common.Address, error) {
	return _SfcV400Contract.Contract.GetRedirection(&_SfcV400Contract.CallOpts, delegator)
}

// GetRedirection is a free data retrieval call binding the contract method 0x736de9ae.
//
// Solidity: function getRedirection(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetRedirection(delegator common.Address) (common.Address, error) {
	return _SfcV400Contract.Contract.GetRedirection(&_SfcV400Contract.CallOpts, delegator)
}

// GetRedirectionRequest is a free data retrieval call binding the contract method 0x468f35ee.
//
// Solidity: function getRedirectionRequest(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractCaller) GetRedirectionRequest(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getRedirectionRequest", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRedirectionRequest is a free data retrieval call binding the contract method 0x468f35ee.
//
// Solidity: function getRedirectionRequest(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractSession) GetRedirectionRequest(delegator common.Address) (common.Address, error) {
	return _SfcV400Contract.Contract.GetRedirectionRequest(&_SfcV400Contract.CallOpts, delegator)
}

// GetRedirectionRequest is a free data retrieval call binding the contract method 0x468f35ee.
//
// Solidity: function getRedirectionRequest(address delegator) view returns(address receiver)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetRedirectionRequest(delegator common.Address) (common.Address, error) {
	return _SfcV400Contract.Contract.GetRedirectionRequest(&_SfcV400Contract.CallOpts, delegator)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) GetSelfStake(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getSelfStake", validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetSelfStake(&_SfcV400Contract.CallOpts, validatorID)
}

// GetSelfStake is a free data retrieval call binding the contract method 0x5601fe01.
//
// Solidity: function getSelfStake(uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetSelfStake(validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetSelfStake(&_SfcV400Contract.CallOpts, validatorID)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address delegator, uint256 validatorID) view returns(uint256 stake)
func (_SfcV400Contract *SfcV400ContractCaller) GetStake(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getStake", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address delegator, uint256 validatorID) view returns(uint256 stake)
func (_SfcV400Contract *SfcV400ContractSession) GetStake(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetStake(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// GetStake is a free data retrieval call binding the contract method 0xcfd47663.
//
// Solidity: function getStake(address delegator, uint256 validatorID) view returns(uint256 stake)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetStake(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetStake(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorID) view returns(uint256 status, uint256 receivedStake, address auth, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedTime, uint256 deactivatedEpoch)
func (_SfcV400Contract *SfcV400ContractCaller) GetValidator(opts *bind.CallOpts, validatorID *big.Int) (struct {
	Status           *big.Int
	ReceivedStake    *big.Int
	Auth             common.Address
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
}, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getValidator", validatorID)

	outstruct := new(struct {
		Status           *big.Int
		ReceivedStake    *big.Int
		Auth             common.Address
		CreatedEpoch     *big.Int
		CreatedTime      *big.Int
		DeactivatedTime  *big.Int
		DeactivatedEpoch *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReceivedStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CreatedEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DeactivatedEpoch = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorID) view returns(uint256 status, uint256 receivedStake, address auth, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedTime, uint256 deactivatedEpoch)
func (_SfcV400Contract *SfcV400ContractSession) GetValidator(validatorID *big.Int) (struct {
	Status           *big.Int
	ReceivedStake    *big.Int
	Auth             common.Address
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetValidator(&_SfcV400Contract.CallOpts, validatorID)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorID) view returns(uint256 status, uint256 receivedStake, address auth, uint256 createdEpoch, uint256 createdTime, uint256 deactivatedTime, uint256 deactivatedEpoch)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetValidator(validatorID *big.Int) (struct {
	Status           *big.Int
	ReceivedStake    *big.Int
	Auth             common.Address
	CreatedEpoch     *big.Int
	CreatedTime      *big.Int
	DeactivatedTime  *big.Int
	DeactivatedEpoch *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetValidator(&_SfcV400Contract.CallOpts, validatorID)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address auth) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractCaller) GetValidatorID(opts *bind.CallOpts, auth common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getValidatorID", auth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address auth) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractSession) GetValidatorID(auth common.Address) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetValidatorID(&_SfcV400Contract.CallOpts, auth)
}

// GetValidatorID is a free data retrieval call binding the contract method 0x0135b1db.
//
// Solidity: function getValidatorID(address auth) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetValidatorID(auth common.Address) (*big.Int, error) {
	return _SfcV400Contract.Contract.GetValidatorID(&_SfcV400Contract.CallOpts, auth)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 validatorID) view returns(bytes pubkey)
func (_SfcV400Contract *SfcV400ContractCaller) GetValidatorPubkey(opts *bind.CallOpts, validatorID *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getValidatorPubkey", validatorID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 validatorID) view returns(bytes pubkey)
func (_SfcV400Contract *SfcV400ContractSession) GetValidatorPubkey(validatorID *big.Int) ([]byte, error) {
	return _SfcV400Contract.Contract.GetValidatorPubkey(&_SfcV400Contract.CallOpts, validatorID)
}

// GetValidatorPubkey is a free data retrieval call binding the contract method 0x854873e1.
//
// Solidity: function getValidatorPubkey(uint256 validatorID) view returns(bytes pubkey)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetValidatorPubkey(validatorID *big.Int) ([]byte, error) {
	return _SfcV400Contract.Contract.GetValidatorPubkey(&_SfcV400Contract.CallOpts, validatorID)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address delegator, uint256 validatorID, uint256 wrID) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_SfcV400Contract *SfcV400ContractCaller) GetWithdrawalRequest(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int, wrID *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "getWithdrawalRequest", delegator, validatorID, wrID)

	outstruct := new(struct {
		Epoch  *big.Int
		Time   *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epoch = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address delegator, uint256 validatorID, uint256 wrID) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_SfcV400Contract *SfcV400ContractSession) GetWithdrawalRequest(delegator common.Address, validatorID *big.Int, wrID *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetWithdrawalRequest(&_SfcV400Contract.CallOpts, delegator, validatorID, wrID)
}

// GetWithdrawalRequest is a free data retrieval call binding the contract method 0x1f270152.
//
// Solidity: function getWithdrawalRequest(address delegator, uint256 validatorID, uint256 wrID) view returns(uint256 epoch, uint256 time, uint256 amount)
func (_SfcV400Contract *SfcV400ContractCallerSession) GetWithdrawalRequest(delegator common.Address, validatorID *big.Int, wrID *big.Int) (struct {
	Epoch  *big.Int
	Time   *big.Int
	Amount *big.Int
}, error) {
	return _SfcV400Contract.Contract.GetWithdrawalRequest(&_SfcV400Contract.CallOpts, delegator, validatorID, wrID)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_SfcV400Contract *SfcV400ContractCaller) IsSlashed(opts *bind.CallOpts, validatorID *big.Int) (bool, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "isSlashed", validatorID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_SfcV400Contract *SfcV400ContractSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _SfcV400Contract.Contract.IsSlashed(&_SfcV400Contract.CallOpts, validatorID)
}

// IsSlashed is a free data retrieval call binding the contract method 0xc3de580e.
//
// Solidity: function isSlashed(uint256 validatorID) view returns(bool)
func (_SfcV400Contract *SfcV400ContractCallerSession) IsSlashed(validatorID *big.Int) (bool, error) {
	return _SfcV400Contract.Contract.IsSlashed(&_SfcV400Contract.CallOpts, validatorID)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) LastValidatorID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "lastValidatorID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) LastValidatorID() (*big.Int, error) {
	return _SfcV400Contract.Contract.LastValidatorID(&_SfcV400Contract.CallOpts)
}

// LastValidatorID is a free data retrieval call binding the contract method 0xc7be95de.
//
// Solidity: function lastValidatorID() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) LastValidatorID() (*big.Int, error) {
	return _SfcV400Contract.Contract.LastValidatorID(&_SfcV400Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Contract *SfcV400ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Contract *SfcV400ContractSession) Owner() (common.Address, error) {
	return _SfcV400Contract.Contract.Owner(&_SfcV400Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SfcV400Contract *SfcV400ContractCallerSession) Owner() (common.Address, error) {
	return _SfcV400Contract.Contract.Owner(&_SfcV400Contract.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) PendingRewards(opts *bind.CallOpts, delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "pendingRewards", delegator, toValidatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.PendingRewards(&_SfcV400Contract.CallOpts, delegator, toValidatorID)
}

// PendingRewards is a free data retrieval call binding the contract method 0x6099ecb2.
//
// Solidity: function pendingRewards(address delegator, uint256 toValidatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) PendingRewards(delegator common.Address, toValidatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.PendingRewards(&_SfcV400Contract.CallOpts, delegator, toValidatorID)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SfcV400Contract *SfcV400ContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SfcV400Contract *SfcV400ContractSession) ProxiableUUID() ([32]byte, error) {
	return _SfcV400Contract.Contract.ProxiableUUID(&_SfcV400Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_SfcV400Contract *SfcV400ContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _SfcV400Contract.Contract.ProxiableUUID(&_SfcV400Contract.CallOpts)
}

// PubkeyAddressToValidatorID is a free data retrieval call binding the contract method 0xfb36025f.
//
// Solidity: function pubkeyAddressToValidatorID(address pubkeyAddress) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractCaller) PubkeyAddressToValidatorID(opts *bind.CallOpts, pubkeyAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "pubkeyAddressToValidatorID", pubkeyAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PubkeyAddressToValidatorID is a free data retrieval call binding the contract method 0xfb36025f.
//
// Solidity: function pubkeyAddressToValidatorID(address pubkeyAddress) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractSession) PubkeyAddressToValidatorID(pubkeyAddress common.Address) (*big.Int, error) {
	return _SfcV400Contract.Contract.PubkeyAddressToValidatorID(&_SfcV400Contract.CallOpts, pubkeyAddress)
}

// PubkeyAddressToValidatorID is a free data retrieval call binding the contract method 0xfb36025f.
//
// Solidity: function pubkeyAddressToValidatorID(address pubkeyAddress) view returns(uint256 validatorID)
func (_SfcV400Contract *SfcV400ContractCallerSession) PubkeyAddressToValidatorID(pubkeyAddress common.Address) (*big.Int, error) {
	return _SfcV400Contract.Contract.PubkeyAddressToValidatorID(&_SfcV400Contract.CallOpts, pubkeyAddress)
}

// RedirectionAuthorizer is a free data retrieval call binding the contract method 0xe9a505a7.
//
// Solidity: function redirectionAuthorizer() view returns(address)
func (_SfcV400Contract *SfcV400ContractCaller) RedirectionAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "redirectionAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedirectionAuthorizer is a free data retrieval call binding the contract method 0xe9a505a7.
//
// Solidity: function redirectionAuthorizer() view returns(address)
func (_SfcV400Contract *SfcV400ContractSession) RedirectionAuthorizer() (common.Address, error) {
	return _SfcV400Contract.Contract.RedirectionAuthorizer(&_SfcV400Contract.CallOpts)
}

// RedirectionAuthorizer is a free data retrieval call binding the contract method 0xe9a505a7.
//
// Solidity: function redirectionAuthorizer() view returns(address)
func (_SfcV400Contract *SfcV400ContractCallerSession) RedirectionAuthorizer() (common.Address, error) {
	return _SfcV400Contract.Contract.RedirectionAuthorizer(&_SfcV400Contract.CallOpts)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) RewardsStash(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "rewardsStash", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.RewardsStash(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// RewardsStash is a free data retrieval call binding the contract method 0x6f498663.
//
// Solidity: function rewardsStash(address delegator, uint256 validatorID) view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) RewardsStash(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.RewardsStash(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 validatorID) view returns(uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractCaller) SlashingRefundRatio(opts *bind.CallOpts, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "slashingRefundRatio", validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 validatorID) view returns(uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractSession) SlashingRefundRatio(validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.SlashingRefundRatio(&_SfcV400Contract.CallOpts, validatorID)
}

// SlashingRefundRatio is a free data retrieval call binding the contract method 0xc65ee0e1.
//
// Solidity: function slashingRefundRatio(uint256 validatorID) view returns(uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractCallerSession) SlashingRefundRatio(validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.SlashingRefundRatio(&_SfcV400Contract.CallOpts, validatorID)
}

// StakeSubscriberAddress is a free data retrieval call binding the contract method 0x093b41d0.
//
// Solidity: function stakeSubscriberAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCaller) StakeSubscriberAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "stakeSubscriberAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeSubscriberAddress is a free data retrieval call binding the contract method 0x093b41d0.
//
// Solidity: function stakeSubscriberAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractSession) StakeSubscriberAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.StakeSubscriberAddress(&_SfcV400Contract.CallOpts)
}

// StakeSubscriberAddress is a free data retrieval call binding the contract method 0x093b41d0.
//
// Solidity: function stakeSubscriberAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCallerSession) StakeSubscriberAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.StakeSubscriberAddress(&_SfcV400Contract.CallOpts)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address delegator, uint256 validatorID) view returns(uint256 epoch)
func (_SfcV400Contract *SfcV400ContractCaller) StashedRewardsUntilEpoch(opts *bind.CallOpts, delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "stashedRewardsUntilEpoch", delegator, validatorID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address delegator, uint256 validatorID) view returns(uint256 epoch)
func (_SfcV400Contract *SfcV400ContractSession) StashedRewardsUntilEpoch(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.StashedRewardsUntilEpoch(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// StashedRewardsUntilEpoch is a free data retrieval call binding the contract method 0xa86a056f.
//
// Solidity: function stashedRewardsUntilEpoch(address delegator, uint256 validatorID) view returns(uint256 epoch)
func (_SfcV400Contract *SfcV400ContractCallerSession) StashedRewardsUntilEpoch(delegator common.Address, validatorID *big.Int) (*big.Int, error) {
	return _SfcV400Contract.Contract.StashedRewardsUntilEpoch(&_SfcV400Contract.CallOpts, delegator, validatorID)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) TotalActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "totalActiveStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) TotalActiveStake() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalActiveStake(&_SfcV400Contract.CallOpts)
}

// TotalActiveStake is a free data retrieval call binding the contract method 0x28f73148.
//
// Solidity: function totalActiveStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) TotalActiveStake() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalActiveStake(&_SfcV400Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) TotalStake() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalStake(&_SfcV400Contract.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) TotalStake() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalStake(&_SfcV400Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractSession) TotalSupply() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalSupply(&_SfcV400Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SfcV400Contract *SfcV400ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _SfcV400Contract.Contract.TotalSupply(&_SfcV400Contract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractSession) TreasuryAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.TreasuryAddress(&_SfcV400Contract.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_SfcV400Contract *SfcV400ContractCallerSession) TreasuryAddress() (common.Address, error) {
	return _SfcV400Contract.Contract.TreasuryAddress(&_SfcV400Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_SfcV400Contract *SfcV400ContractCaller) Version(opts *bind.CallOpts) ([3]byte, error) {
	var out []interface{}
	err := _SfcV400Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new([3]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([3]byte)).(*[3]byte)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_SfcV400Contract *SfcV400ContractSession) Version() ([3]byte, error) {
	return _SfcV400Contract.Contract.Version(&_SfcV400Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(bytes3)
func (_SfcV400Contract *SfcV400ContractCallerSession) Version() ([3]byte, error) {
	return _SfcV400Contract.Contract.Version(&_SfcV400Contract.CallOpts)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SyncValidator(opts *bind.TransactOpts, validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "_syncValidator", validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_SfcV400Contract *SfcV400ContractSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SyncValidator(&_SfcV400Contract.TransactOpts, validatorID, syncPubkey)
}

// SyncValidator is a paid mutator transaction binding the contract method 0xcc8343aa.
//
// Solidity: function _syncValidator(uint256 validatorID, bool syncPubkey) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SyncValidator(validatorID *big.Int, syncPubkey bool) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SyncValidator(&_SfcV400Contract.TransactOpts, validatorID, syncPubkey)
}

// AnnounceRedirection is a paid mutator transaction binding the contract method 0x46f1ca35.
//
// Solidity: function announceRedirection(address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) AnnounceRedirection(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "announceRedirection", to)
}

// AnnounceRedirection is a paid mutator transaction binding the contract method 0x46f1ca35.
//
// Solidity: function announceRedirection(address to) returns()
func (_SfcV400Contract *SfcV400ContractSession) AnnounceRedirection(to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.AnnounceRedirection(&_SfcV400Contract.TransactOpts, to)
}

// AnnounceRedirection is a paid mutator transaction binding the contract method 0x46f1ca35.
//
// Solidity: function announceRedirection(address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) AnnounceRedirection(to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.AnnounceRedirection(&_SfcV400Contract.TransactOpts, to)
}

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) BurnFTM(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "burnFTM", amount)
}

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractSession) BurnFTM(amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.BurnFTM(&_SfcV400Contract.TransactOpts, amount)
}

// BurnFTM is a paid mutator transaction binding the contract method 0x90a6c475.
//
// Solidity: function burnFTM(uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) BurnFTM(amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.BurnFTM(&_SfcV400Contract.TransactOpts, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) ClaimRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "claimRewards", toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.ClaimRewards(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x0962ef79.
//
// Solidity: function claimRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) ClaimRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.ClaimRewards(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "createValidator", pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_SfcV400Contract *SfcV400ContractSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.CreateValidator(&_SfcV400Contract.TransactOpts, pubkey)
}

// CreateValidator is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) CreateValidator(pubkey []byte) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.CreateValidator(&_SfcV400Contract.TransactOpts, pubkey)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) DeactivateValidator(opts *bind.TransactOpts, validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "deactivateValidator", validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_SfcV400Contract *SfcV400ContractSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.DeactivateValidator(&_SfcV400Contract.TransactOpts, validatorID, status)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x1e702f83.
//
// Solidity: function deactivateValidator(uint256 validatorID, uint256 status) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) DeactivateValidator(validatorID *big.Int, status *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.DeactivateValidator(&_SfcV400Contract.TransactOpts, validatorID, status)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Delegate(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "delegate", toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_SfcV400Contract *SfcV400ContractSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Delegate(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// Delegate is a paid mutator transaction binding the contract method 0x9fa6dd35.
//
// Solidity: function delegate(uint256 toValidatorID) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Delegate(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Delegate(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fbfd4df.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address _c, address owner) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Initialize(opts *bind.TransactOpts, sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, _c common.Address, owner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "initialize", sealedEpoch, _totalSupply, nodeDriver, _c, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fbfd4df.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address _c, address owner) returns()
func (_SfcV400Contract *SfcV400ContractSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, _c common.Address, owner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Initialize(&_SfcV400Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, _c, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fbfd4df.
//
// Solidity: function initialize(uint256 sealedEpoch, uint256 _totalSupply, address nodeDriver, address _c, address owner) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Initialize(sealedEpoch *big.Int, _totalSupply *big.Int, nodeDriver common.Address, _c common.Address, owner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Initialize(&_SfcV400Contract.TransactOpts, sealedEpoch, _totalSupply, nodeDriver, _c, owner)
}

// InitiateRedirection is a paid mutator transaction binding the contract method 0xcc172784.
//
// Solidity: function initiateRedirection(address from, address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) InitiateRedirection(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "initiateRedirection", from, to)
}

// InitiateRedirection is a paid mutator transaction binding the contract method 0xcc172784.
//
// Solidity: function initiateRedirection(address from, address to) returns()
func (_SfcV400Contract *SfcV400ContractSession) InitiateRedirection(from common.Address, to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.InitiateRedirection(&_SfcV400Contract.TransactOpts, from, to)
}

// InitiateRedirection is a paid mutator transaction binding the contract method 0xcc172784.
//
// Solidity: function initiateRedirection(address from, address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) InitiateRedirection(from common.Address, to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.InitiateRedirection(&_SfcV400Contract.TransactOpts, from, to)
}

// Redirect is a paid mutator transaction binding the contract method 0xd725e91f.
//
// Solidity: function redirect(address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Redirect(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "redirect", to)
}

// Redirect is a paid mutator transaction binding the contract method 0xd725e91f.
//
// Solidity: function redirect(address to) returns()
func (_SfcV400Contract *SfcV400ContractSession) Redirect(to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Redirect(&_SfcV400Contract.TransactOpts, to)
}

// Redirect is a paid mutator transaction binding the contract method 0xd725e91f.
//
// Solidity: function redirect(address to) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Redirect(to common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Redirect(&_SfcV400Contract.TransactOpts, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Contract *SfcV400ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Contract *SfcV400ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV400Contract.Contract.RenounceOwnership(&_SfcV400Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SfcV400Contract.Contract.RenounceOwnership(&_SfcV400Contract.TransactOpts)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) RestakeRewards(opts *bind.TransactOpts, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "restakeRewards", toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.RestakeRewards(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// RestakeRewards is a paid mutator transaction binding the contract method 0x08c36874.
//
// Solidity: function restakeRewards(uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) RestakeRewards(toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.RestakeRewards(&_SfcV400Contract.TransactOpts, toValidatorID)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SealEpoch(opts *bind.TransactOpts, offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "sealEpoch", offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_SfcV400Contract *SfcV400ContractSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SealEpoch(&_SfcV400Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpoch is a paid mutator transaction binding the contract method 0xebdf104c.
//
// Solidity: function sealEpoch(uint256[] offlineTime, uint256[] offlineBlocks, uint256[] uptimes, uint256[] originatedTxsFee) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SealEpoch(offlineTime []*big.Int, offlineBlocks []*big.Int, uptimes []*big.Int, originatedTxsFee []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SealEpoch(&_SfcV400Contract.TransactOpts, offlineTime, offlineBlocks, uptimes, originatedTxsFee)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SealEpochValidators(opts *bind.TransactOpts, nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "sealEpochValidators", nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_SfcV400Contract *SfcV400ContractSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SealEpochValidators(&_SfcV400Contract.TransactOpts, nextValidatorIDs)
}

// SealEpochValidators is a paid mutator transaction binding the contract method 0xe08d7e66.
//
// Solidity: function sealEpochValidators(uint256[] nextValidatorIDs) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SealEpochValidators(nextValidatorIDs []*big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SealEpochValidators(&_SfcV400Contract.TransactOpts, nextValidatorIDs)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0xa8ab09ba.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SetGenesisDelegation(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int, stake *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "setGenesisDelegation", delegator, toValidatorID, stake)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0xa8ab09ba.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake) returns()
func (_SfcV400Contract *SfcV400ContractSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetGenesisDelegation(&_SfcV400Contract.TransactOpts, delegator, toValidatorID, stake)
}

// SetGenesisDelegation is a paid mutator transaction binding the contract method 0xa8ab09ba.
//
// Solidity: function setGenesisDelegation(address delegator, uint256 toValidatorID, uint256 stake) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SetGenesisDelegation(delegator common.Address, toValidatorID *big.Int, stake *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetGenesisDelegation(&_SfcV400Contract.TransactOpts, delegator, toValidatorID, stake)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x76fed43a.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 createdTime) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SetGenesisValidator(opts *bind.TransactOpts, auth common.Address, validatorID *big.Int, pubkey []byte, createdTime *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "setGenesisValidator", auth, validatorID, pubkey, createdTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x76fed43a.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 createdTime) returns()
func (_SfcV400Contract *SfcV400ContractSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, createdTime *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetGenesisValidator(&_SfcV400Contract.TransactOpts, auth, validatorID, pubkey, createdTime)
}

// SetGenesisValidator is a paid mutator transaction binding the contract method 0x76fed43a.
//
// Solidity: function setGenesisValidator(address auth, uint256 validatorID, bytes pubkey, uint256 createdTime) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SetGenesisValidator(auth common.Address, validatorID *big.Int, pubkey []byte, createdTime *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetGenesisValidator(&_SfcV400Contract.TransactOpts, auth, validatorID, pubkey, createdTime)
}

// SetRedirectionAuthorizer is a paid mutator transaction binding the contract method 0xb0ef386c.
//
// Solidity: function setRedirectionAuthorizer(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) SetRedirectionAuthorizer(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "setRedirectionAuthorizer", v)
}

// SetRedirectionAuthorizer is a paid mutator transaction binding the contract method 0xb0ef386c.
//
// Solidity: function setRedirectionAuthorizer(address v) returns()
func (_SfcV400Contract *SfcV400ContractSession) SetRedirectionAuthorizer(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetRedirectionAuthorizer(&_SfcV400Contract.TransactOpts, v)
}

// SetRedirectionAuthorizer is a paid mutator transaction binding the contract method 0xb0ef386c.
//
// Solidity: function setRedirectionAuthorizer(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) SetRedirectionAuthorizer(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.SetRedirectionAuthorizer(&_SfcV400Contract.TransactOpts, v)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) StashRewards(opts *bind.TransactOpts, delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "stashRewards", delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.StashRewards(&_SfcV400Contract.TransactOpts, delegator, toValidatorID)
}

// StashRewards is a paid mutator transaction binding the contract method 0x8cddb015.
//
// Solidity: function stashRewards(address delegator, uint256 toValidatorID) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) StashRewards(delegator common.Address, toValidatorID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.StashRewards(&_SfcV400Contract.TransactOpts, delegator, toValidatorID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Contract *SfcV400ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.TransferOwnership(&_SfcV400Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.TransferOwnership(&_SfcV400Contract.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Undelegate(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "undelegate", toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Undelegate(&_SfcV400Contract.TransactOpts, toValidatorID, wrID, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4f864df4.
//
// Solidity: function undelegate(uint256 toValidatorID, uint256 wrID, uint256 amount) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Undelegate(toValidatorID *big.Int, wrID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Undelegate(&_SfcV400Contract.TransactOpts, toValidatorID, wrID, amount)
}

// UpdateConstsAddress is a paid mutator transaction binding the contract method 0x860c2750.
//
// Solidity: function updateConstsAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) UpdateConstsAddress(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "updateConstsAddress", v)
}

// UpdateConstsAddress is a paid mutator transaction binding the contract method 0x860c2750.
//
// Solidity: function updateConstsAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractSession) UpdateConstsAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateConstsAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpdateConstsAddress is a paid mutator transaction binding the contract method 0x860c2750.
//
// Solidity: function updateConstsAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) UpdateConstsAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateConstsAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) UpdateSlashingRefundRatio(opts *bind.TransactOpts, validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "updateSlashingRefundRatio", validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_SfcV400Contract *SfcV400ContractSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateSlashingRefundRatio(&_SfcV400Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateSlashingRefundRatio is a paid mutator transaction binding the contract method 0x4f7c4efb.
//
// Solidity: function updateSlashingRefundRatio(uint256 validatorID, uint256 refundRatio) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) UpdateSlashingRefundRatio(validatorID *big.Int, refundRatio *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateSlashingRefundRatio(&_SfcV400Contract.TransactOpts, validatorID, refundRatio)
}

// UpdateStakeSubscriberAddress is a paid mutator transaction binding the contract method 0xe880a159.
//
// Solidity: function updateStakeSubscriberAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) UpdateStakeSubscriberAddress(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "updateStakeSubscriberAddress", v)
}

// UpdateStakeSubscriberAddress is a paid mutator transaction binding the contract method 0xe880a159.
//
// Solidity: function updateStakeSubscriberAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractSession) UpdateStakeSubscriberAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateStakeSubscriberAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpdateStakeSubscriberAddress is a paid mutator transaction binding the contract method 0xe880a159.
//
// Solidity: function updateStakeSubscriberAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) UpdateStakeSubscriberAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateStakeSubscriberAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) UpdateTreasuryAddress(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "updateTreasuryAddress", v)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractSession) UpdateTreasuryAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateTreasuryAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpdateTreasuryAddress is a paid mutator transaction binding the contract method 0x841e4561.
//
// Solidity: function updateTreasuryAddress(address v) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) UpdateTreasuryAddress(v common.Address) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpdateTreasuryAddress(&_SfcV400Contract.TransactOpts, v)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SfcV400Contract *SfcV400ContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpgradeToAndCall(&_SfcV400Contract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.UpgradeToAndCall(&_SfcV400Contract.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Withdraw(opts *bind.TransactOpts, toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.contract.Transact(opts, "withdraw", toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_SfcV400Contract *SfcV400ContractSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Withdraw(&_SfcV400Contract.TransactOpts, toValidatorID, wrID)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 toValidatorID, uint256 wrID) returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Withdraw(toValidatorID *big.Int, wrID *big.Int) (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Withdraw(&_SfcV400Contract.TransactOpts, toValidatorID, wrID)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SfcV400Contract *SfcV400ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SfcV400Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SfcV400Contract *SfcV400ContractSession) Receive() (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Receive(&_SfcV400Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SfcV400Contract *SfcV400ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _SfcV400Contract.Contract.Receive(&_SfcV400Contract.TransactOpts)
}

// SfcV400ContractAnnouncedRedirectionIterator is returned from FilterAnnouncedRedirection and is used to iterate over the raw logs and unpacked data for AnnouncedRedirection events raised by the SfcV400Contract contract.
type SfcV400ContractAnnouncedRedirectionIterator struct {
	Event *SfcV400ContractAnnouncedRedirection // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractAnnouncedRedirectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractAnnouncedRedirection)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractAnnouncedRedirection)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractAnnouncedRedirectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractAnnouncedRedirectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractAnnouncedRedirection represents a AnnouncedRedirection event raised by the SfcV400Contract contract.
type SfcV400ContractAnnouncedRedirection struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAnnouncedRedirection is a free log retrieval operation binding the contract event 0x857125196131cfcd709c738c6d1fd2701ce70f2a03785aeadae6f4b47fe73c1d.
//
// Solidity: event AnnouncedRedirection(address indexed from, address indexed to)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterAnnouncedRedirection(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SfcV400ContractAnnouncedRedirectionIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "AnnouncedRedirection", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractAnnouncedRedirectionIterator{contract: _SfcV400Contract.contract, event: "AnnouncedRedirection", logs: logs, sub: sub}, nil
}

// WatchAnnouncedRedirection is a free log subscription operation binding the contract event 0x857125196131cfcd709c738c6d1fd2701ce70f2a03785aeadae6f4b47fe73c1d.
//
// Solidity: event AnnouncedRedirection(address indexed from, address indexed to)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchAnnouncedRedirection(opts *bind.WatchOpts, sink chan<- *SfcV400ContractAnnouncedRedirection, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "AnnouncedRedirection", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractAnnouncedRedirection)
				if err := _SfcV400Contract.contract.UnpackLog(event, "AnnouncedRedirection", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAnnouncedRedirection is a log parse operation binding the contract event 0x857125196131cfcd709c738c6d1fd2701ce70f2a03785aeadae6f4b47fe73c1d.
//
// Solidity: event AnnouncedRedirection(address indexed from, address indexed to)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseAnnouncedRedirection(log types.Log) (*SfcV400ContractAnnouncedRedirection, error) {
	event := new(SfcV400ContractAnnouncedRedirection)
	if err := _SfcV400Contract.contract.UnpackLog(event, "AnnouncedRedirection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractBurntFTMIterator is returned from FilterBurntFTM and is used to iterate over the raw logs and unpacked data for BurntFTM events raised by the SfcV400Contract contract.
type SfcV400ContractBurntFTMIterator struct {
	Event *SfcV400ContractBurntFTM // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractBurntFTMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractBurntFTM)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractBurntFTM)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractBurntFTMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractBurntFTMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractBurntFTM represents a BurntFTM event raised by the SfcV400Contract contract.
type SfcV400ContractBurntFTM struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurntFTM is a free log retrieval operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterBurntFTM(opts *bind.FilterOpts) (*SfcV400ContractBurntFTMIterator, error) {

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "BurntFTM")
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractBurntFTMIterator{contract: _SfcV400Contract.contract, event: "BurntFTM", logs: logs, sub: sub}, nil
}

// WatchBurntFTM is a free log subscription operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchBurntFTM(opts *bind.WatchOpts, sink chan<- *SfcV400ContractBurntFTM) (event.Subscription, error) {

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "BurntFTM")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractBurntFTM)
				if err := _SfcV400Contract.contract.UnpackLog(event, "BurntFTM", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurntFTM is a log parse operation binding the contract event 0x8918bd6046d08b314e457977f29562c5d76a7030d79b1edba66e8a5da0b77ae8.
//
// Solidity: event BurntFTM(uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseBurntFTM(log types.Log) (*SfcV400ContractBurntFTM, error) {
	event := new(SfcV400ContractBurntFTM)
	if err := _SfcV400Contract.contract.UnpackLog(event, "BurntFTM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractChangedValidatorStatusIterator is returned from FilterChangedValidatorStatus and is used to iterate over the raw logs and unpacked data for ChangedValidatorStatus events raised by the SfcV400Contract contract.
type SfcV400ContractChangedValidatorStatusIterator struct {
	Event *SfcV400ContractChangedValidatorStatus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractChangedValidatorStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractChangedValidatorStatus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractChangedValidatorStatus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractChangedValidatorStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractChangedValidatorStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractChangedValidatorStatus represents a ChangedValidatorStatus event raised by the SfcV400Contract contract.
type SfcV400ContractChangedValidatorStatus struct {
	ValidatorID *big.Int
	Status      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChangedValidatorStatus is a free log retrieval operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterChangedValidatorStatus(opts *bind.FilterOpts, validatorID []*big.Int) (*SfcV400ContractChangedValidatorStatusIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractChangedValidatorStatusIterator{contract: _SfcV400Contract.contract, event: "ChangedValidatorStatus", logs: logs, sub: sub}, nil
}

// WatchChangedValidatorStatus is a free log subscription operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchChangedValidatorStatus(opts *bind.WatchOpts, sink chan<- *SfcV400ContractChangedValidatorStatus, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "ChangedValidatorStatus", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractChangedValidatorStatus)
				if err := _SfcV400Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedValidatorStatus is a log parse operation binding the contract event 0xcd35267e7654194727477d6c78b541a553483cff7f92a055d17868d3da6e953e.
//
// Solidity: event ChangedValidatorStatus(uint256 indexed validatorID, uint256 status)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseChangedValidatorStatus(log types.Log) (*SfcV400ContractChangedValidatorStatus, error) {
	event := new(SfcV400ContractChangedValidatorStatus)
	if err := _SfcV400Contract.contract.UnpackLog(event, "ChangedValidatorStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the SfcV400Contract contract.
type SfcV400ContractClaimedRewardsIterator struct {
	Event *SfcV400ContractClaimedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractClaimedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractClaimedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractClaimedRewards represents a ClaimedRewards event raised by the SfcV400Contract contract.
type SfcV400ContractClaimedRewards struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Rewards       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0x70de20a533702af05c8faf1637846c4586a021bbc71b6928b089b6d555e4fbc2.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterClaimedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*SfcV400ContractClaimedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractClaimedRewardsIterator{contract: _SfcV400Contract.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0x70de20a533702af05c8faf1637846c4586a021bbc71b6928b089b6d555e4fbc2.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *SfcV400ContractClaimedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "ClaimedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractClaimedRewards)
				if err := _SfcV400Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimedRewards is a log parse operation binding the contract event 0x70de20a533702af05c8faf1637846c4586a021bbc71b6928b089b6d555e4fbc2.
//
// Solidity: event ClaimedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseClaimedRewards(log types.Log) (*SfcV400ContractClaimedRewards, error) {
	event := new(SfcV400ContractClaimedRewards)
	if err := _SfcV400Contract.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractCreatedValidatorIterator is returned from FilterCreatedValidator and is used to iterate over the raw logs and unpacked data for CreatedValidator events raised by the SfcV400Contract contract.
type SfcV400ContractCreatedValidatorIterator struct {
	Event *SfcV400ContractCreatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractCreatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractCreatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractCreatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractCreatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractCreatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractCreatedValidator represents a CreatedValidator event raised by the SfcV400Contract contract.
type SfcV400ContractCreatedValidator struct {
	ValidatorID  *big.Int
	Auth         common.Address
	CreatedEpoch *big.Int
	CreatedTime  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatedValidator is a free log retrieval operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterCreatedValidator(opts *bind.FilterOpts, validatorID []*big.Int, auth []common.Address) (*SfcV400ContractCreatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractCreatedValidatorIterator{contract: _SfcV400Contract.contract, event: "CreatedValidator", logs: logs, sub: sub}, nil
}

// WatchCreatedValidator is a free log subscription operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchCreatedValidator(opts *bind.WatchOpts, sink chan<- *SfcV400ContractCreatedValidator, validatorID []*big.Int, auth []common.Address) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}
	var authRule []interface{}
	for _, authItem := range auth {
		authRule = append(authRule, authItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "CreatedValidator", validatorIDRule, authRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractCreatedValidator)
				if err := _SfcV400Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreatedValidator is a log parse operation binding the contract event 0x49bca1ed2666922f9f1690c26a569e1299c2a715fe57647d77e81adfabbf25bf.
//
// Solidity: event CreatedValidator(uint256 indexed validatorID, address indexed auth, uint256 createdEpoch, uint256 createdTime)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseCreatedValidator(log types.Log) (*SfcV400ContractCreatedValidator, error) {
	event := new(SfcV400ContractCreatedValidator)
	if err := _SfcV400Contract.contract.UnpackLog(event, "CreatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractDeactivatedValidatorIterator is returned from FilterDeactivatedValidator and is used to iterate over the raw logs and unpacked data for DeactivatedValidator events raised by the SfcV400Contract contract.
type SfcV400ContractDeactivatedValidatorIterator struct {
	Event *SfcV400ContractDeactivatedValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractDeactivatedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractDeactivatedValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractDeactivatedValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractDeactivatedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractDeactivatedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractDeactivatedValidator represents a DeactivatedValidator event raised by the SfcV400Contract contract.
type SfcV400ContractDeactivatedValidator struct {
	ValidatorID      *big.Int
	DeactivatedEpoch *big.Int
	DeactivatedTime  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeactivatedValidator is a free log retrieval operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterDeactivatedValidator(opts *bind.FilterOpts, validatorID []*big.Int) (*SfcV400ContractDeactivatedValidatorIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractDeactivatedValidatorIterator{contract: _SfcV400Contract.contract, event: "DeactivatedValidator", logs: logs, sub: sub}, nil
}

// WatchDeactivatedValidator is a free log subscription operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchDeactivatedValidator(opts *bind.WatchOpts, sink chan<- *SfcV400ContractDeactivatedValidator, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "DeactivatedValidator", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractDeactivatedValidator)
				if err := _SfcV400Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivatedValidator is a log parse operation binding the contract event 0xac4801c32a6067ff757446524ee4e7a373797278ac3c883eac5c693b4ad72e47.
//
// Solidity: event DeactivatedValidator(uint256 indexed validatorID, uint256 deactivatedEpoch, uint256 deactivatedTime)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseDeactivatedValidator(log types.Log) (*SfcV400ContractDeactivatedValidator, error) {
	event := new(SfcV400ContractDeactivatedValidator)
	if err := _SfcV400Contract.contract.UnpackLog(event, "DeactivatedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the SfcV400Contract contract.
type SfcV400ContractDelegatedIterator struct {
	Event *SfcV400ContractDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractDelegated represents a Delegated event raised by the SfcV400Contract contract.
type SfcV400ContractDelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*SfcV400ContractDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractDelegatedIterator{contract: _SfcV400Contract.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *SfcV400ContractDelegated, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "Delegated", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractDelegated)
				if err := _SfcV400Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0x9a8f44850296624dadfd9c246d17e47171d35727a181bd090aa14bbbe00238bb.
//
// Solidity: event Delegated(address indexed delegator, uint256 indexed toValidatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseDelegated(log types.Log) (*SfcV400ContractDelegated, error) {
	event := new(SfcV400ContractDelegated)
	if err := _SfcV400Contract.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SfcV400Contract contract.
type SfcV400ContractInitializedIterator struct {
	Event *SfcV400ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractInitialized represents a Initialized event raised by the SfcV400Contract contract.
type SfcV400ContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*SfcV400ContractInitializedIterator, error) {

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractInitializedIterator{contract: _SfcV400Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SfcV400ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractInitialized)
				if err := _SfcV400Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseInitialized(log types.Log) (*SfcV400ContractInitialized, error) {
	event := new(SfcV400ContractInitialized)
	if err := _SfcV400Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SfcV400Contract contract.
type SfcV400ContractOwnershipTransferredIterator struct {
	Event *SfcV400ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractOwnershipTransferred represents a OwnershipTransferred event raised by the SfcV400Contract contract.
type SfcV400ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SfcV400ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractOwnershipTransferredIterator{contract: _SfcV400Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SfcV400ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractOwnershipTransferred)
				if err := _SfcV400Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseOwnershipTransferred(log types.Log) (*SfcV400ContractOwnershipTransferred, error) {
	event := new(SfcV400ContractOwnershipTransferred)
	if err := _SfcV400Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractRefundedSlashedLegacyDelegationIterator is returned from FilterRefundedSlashedLegacyDelegation and is used to iterate over the raw logs and unpacked data for RefundedSlashedLegacyDelegation events raised by the SfcV400Contract contract.
type SfcV400ContractRefundedSlashedLegacyDelegationIterator struct {
	Event *SfcV400ContractRefundedSlashedLegacyDelegation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractRefundedSlashedLegacyDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractRefundedSlashedLegacyDelegation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractRefundedSlashedLegacyDelegation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractRefundedSlashedLegacyDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractRefundedSlashedLegacyDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractRefundedSlashedLegacyDelegation represents a RefundedSlashedLegacyDelegation event raised by the SfcV400Contract contract.
type SfcV400ContractRefundedSlashedLegacyDelegation struct {
	Delegator   common.Address
	ValidatorID *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRefundedSlashedLegacyDelegation is a free log retrieval operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterRefundedSlashedLegacyDelegation(opts *bind.FilterOpts, delegator []common.Address, validatorID []*big.Int) (*SfcV400ContractRefundedSlashedLegacyDelegationIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractRefundedSlashedLegacyDelegationIterator{contract: _SfcV400Contract.contract, event: "RefundedSlashedLegacyDelegation", logs: logs, sub: sub}, nil
}

// WatchRefundedSlashedLegacyDelegation is a free log subscription operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchRefundedSlashedLegacyDelegation(opts *bind.WatchOpts, sink chan<- *SfcV400ContractRefundedSlashedLegacyDelegation, delegator []common.Address, validatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "RefundedSlashedLegacyDelegation", delegatorRule, validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractRefundedSlashedLegacyDelegation)
				if err := _SfcV400Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundedSlashedLegacyDelegation is a log parse operation binding the contract event 0x172fdfaf5222519d28d2794b7617be033f46d954f9b6c3896e7d2611ff444252.
//
// Solidity: event RefundedSlashedLegacyDelegation(address indexed delegator, uint256 indexed validatorID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseRefundedSlashedLegacyDelegation(log types.Log) (*SfcV400ContractRefundedSlashedLegacyDelegation, error) {
	event := new(SfcV400ContractRefundedSlashedLegacyDelegation)
	if err := _SfcV400Contract.contract.UnpackLog(event, "RefundedSlashedLegacyDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractRestakedRewardsIterator is returned from FilterRestakedRewards and is used to iterate over the raw logs and unpacked data for RestakedRewards events raised by the SfcV400Contract contract.
type SfcV400ContractRestakedRewardsIterator struct {
	Event *SfcV400ContractRestakedRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractRestakedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractRestakedRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractRestakedRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractRestakedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractRestakedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractRestakedRewards represents a RestakedRewards event raised by the SfcV400Contract contract.
type SfcV400ContractRestakedRewards struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	Rewards       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRestakedRewards is a free log retrieval operation binding the contract event 0x663e0f63f4fc6b01be195c4b56111fd6f14b947d6264497119b08daf77e26da5.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterRestakedRewards(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int) (*SfcV400ContractRestakedRewardsIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractRestakedRewardsIterator{contract: _SfcV400Contract.contract, event: "RestakedRewards", logs: logs, sub: sub}, nil
}

// WatchRestakedRewards is a free log subscription operation binding the contract event 0x663e0f63f4fc6b01be195c4b56111fd6f14b947d6264497119b08daf77e26da5.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchRestakedRewards(opts *bind.WatchOpts, sink chan<- *SfcV400ContractRestakedRewards, delegator []common.Address, toValidatorID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "RestakedRewards", delegatorRule, toValidatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractRestakedRewards)
				if err := _SfcV400Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRestakedRewards is a log parse operation binding the contract event 0x663e0f63f4fc6b01be195c4b56111fd6f14b947d6264497119b08daf77e26da5.
//
// Solidity: event RestakedRewards(address indexed delegator, uint256 indexed toValidatorID, uint256 rewards)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseRestakedRewards(log types.Log) (*SfcV400ContractRestakedRewards, error) {
	event := new(SfcV400ContractRestakedRewards)
	if err := _SfcV400Contract.contract.UnpackLog(event, "RestakedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the SfcV400Contract contract.
type SfcV400ContractUndelegatedIterator struct {
	Event *SfcV400ContractUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractUndelegated represents a Undelegated event raised by the SfcV400Contract contract.
type SfcV400ContractUndelegated struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*SfcV400ContractUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractUndelegatedIterator{contract: _SfcV400Contract.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *SfcV400ContractUndelegated, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "Undelegated", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractUndelegated)
				if err := _SfcV400Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0xd3bb4e423fbea695d16b982f9f682dc5f35152e5411646a8a5a79a6b02ba8d57.
//
// Solidity: event Undelegated(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseUndelegated(log types.Log) (*SfcV400ContractUndelegated, error) {
	event := new(SfcV400ContractUndelegated)
	if err := _SfcV400Contract.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractUpdatedSlashingRefundRatioIterator is returned from FilterUpdatedSlashingRefundRatio and is used to iterate over the raw logs and unpacked data for UpdatedSlashingRefundRatio events raised by the SfcV400Contract contract.
type SfcV400ContractUpdatedSlashingRefundRatioIterator struct {
	Event *SfcV400ContractUpdatedSlashingRefundRatio // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractUpdatedSlashingRefundRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractUpdatedSlashingRefundRatio)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractUpdatedSlashingRefundRatio)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractUpdatedSlashingRefundRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractUpdatedSlashingRefundRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractUpdatedSlashingRefundRatio represents a UpdatedSlashingRefundRatio event raised by the SfcV400Contract contract.
type SfcV400ContractUpdatedSlashingRefundRatio struct {
	ValidatorID *big.Int
	RefundRatio *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSlashingRefundRatio is a free log retrieval operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterUpdatedSlashingRefundRatio(opts *bind.FilterOpts, validatorID []*big.Int) (*SfcV400ContractUpdatedSlashingRefundRatioIterator, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractUpdatedSlashingRefundRatioIterator{contract: _SfcV400Contract.contract, event: "UpdatedSlashingRefundRatio", logs: logs, sub: sub}, nil
}

// WatchUpdatedSlashingRefundRatio is a free log subscription operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchUpdatedSlashingRefundRatio(opts *bind.WatchOpts, sink chan<- *SfcV400ContractUpdatedSlashingRefundRatio, validatorID []*big.Int) (event.Subscription, error) {

	var validatorIDRule []interface{}
	for _, validatorIDItem := range validatorID {
		validatorIDRule = append(validatorIDRule, validatorIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "UpdatedSlashingRefundRatio", validatorIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractUpdatedSlashingRefundRatio)
				if err := _SfcV400Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatedSlashingRefundRatio is a log parse operation binding the contract event 0x047575f43f09a7a093d94ec483064acfc61b7e25c0de28017da442abf99cb917.
//
// Solidity: event UpdatedSlashingRefundRatio(uint256 indexed validatorID, uint256 refundRatio)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseUpdatedSlashingRefundRatio(log types.Log) (*SfcV400ContractUpdatedSlashingRefundRatio, error) {
	event := new(SfcV400ContractUpdatedSlashingRefundRatio)
	if err := _SfcV400Contract.contract.UnpackLog(event, "UpdatedSlashingRefundRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the SfcV400Contract contract.
type SfcV400ContractUpgradedIterator struct {
	Event *SfcV400ContractUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractUpgraded represents a Upgraded event raised by the SfcV400Contract contract.
type SfcV400ContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SfcV400ContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractUpgradedIterator{contract: _SfcV400Contract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SfcV400ContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractUpgraded)
				if err := _SfcV400Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseUpgraded(log types.Log) (*SfcV400ContractUpgraded, error) {
	event := new(SfcV400ContractUpgraded)
	if err := _SfcV400Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SfcV400ContractWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the SfcV400Contract contract.
type SfcV400ContractWithdrawnIterator struct {
	Event *SfcV400ContractWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SfcV400ContractWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SfcV400ContractWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SfcV400ContractWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SfcV400ContractWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SfcV400ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SfcV400ContractWithdrawn represents a Withdrawn event raised by the SfcV400Contract contract.
type SfcV400ContractWithdrawn struct {
	Delegator     common.Address
	ToValidatorID *big.Int
	WrID          *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (*SfcV400ContractWithdrawnIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.FilterLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return &SfcV400ContractWithdrawnIterator{contract: _SfcV400Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *SfcV400ContractWithdrawn, delegator []common.Address, toValidatorID []*big.Int, wrID []*big.Int) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var toValidatorIDRule []interface{}
	for _, toValidatorIDItem := range toValidatorID {
		toValidatorIDRule = append(toValidatorIDRule, toValidatorIDItem)
	}
	var wrIDRule []interface{}
	for _, wrIDItem := range wrID {
		wrIDRule = append(wrIDRule, wrIDItem)
	}

	logs, sub, err := _SfcV400Contract.contract.WatchLogs(opts, "Withdrawn", delegatorRule, toValidatorIDRule, wrIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SfcV400ContractWithdrawn)
				if err := _SfcV400Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x75e161b3e824b114fc1a33274bd7091918dd4e639cede50b78b15a4eea956a21.
//
// Solidity: event Withdrawn(address indexed delegator, uint256 indexed toValidatorID, uint256 indexed wrID, uint256 amount)
func (_SfcV400Contract *SfcV400ContractFilterer) ParseWithdrawn(log types.Log) (*SfcV400ContractWithdrawn, error) {
	event := new(SfcV400ContractWithdrawn)
	if err := _SfcV400Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
