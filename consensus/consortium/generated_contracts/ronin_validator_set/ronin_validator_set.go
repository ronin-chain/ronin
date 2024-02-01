// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package roninValidatorSet

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
)

// ICandidateManagerCommissionSchedule is an auto generated low-level Go binding around an user-defined struct.
type ICandidateManagerCommissionSchedule struct {
	EffectiveTimestamp *big.Int
	CommissionRate     *big.Int
}

// ICandidateManagerValidatorCandidate is an auto generated low-level Go binding around an user-defined struct.
type ICandidateManagerValidatorCandidate struct {
	ShadowedAdmin                common.Address
	ShadowedConsensus            common.Address
	ShadowedTreasury             common.Address
	DeprecatedBridgeOperatorAddr common.Address
	CommissionRate               *big.Int
	RevokingTimestamp            *big.Int
	TopupDeadline                *big.Int
}

// ICommonInfoEmergencyExitInfo is an auto generated low-level Go binding around an user-defined struct.
type ICommonInfoEmergencyExitInfo struct {
	LockedAmount *big.Int
	RecyclingAt  *big.Int
}

// RoninValidatorSetMetaData contains all meta data concerning the RoninValidatorSet contract.
var RoninValidatorSetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedEmergencyExit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedRevokingCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedUpdatingCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyWrappedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAtEndOfEpochOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallPrecompiled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeCoinbase\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ErrCannotBailout\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedsMaxNumberOfCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExistentCandidate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidateAdminAddr\",\"type\":\"address\"}],\"name\":\"ErrExistentCandidateAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"ErrExistentTreasury\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sendAmount\",\"type\":\"uint256\"}],\"name\":\"ErrInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidEffectiveDaysOnwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidMaxPrioritizedValidatorNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidMinEffectiveDaysOnwards\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"}],\"name\":\"ErrLockedFundMightBeRecycled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"}],\"name\":\"ErrLockedFundReleaseInfoNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNonExistentCandidate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrRecipientRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrTrustedOrgCannotRenounce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrUnauthorizedReceiveRON\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumContractType\",\"name\":\"expectedContractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ErrUnexpectedInternalCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonExistentRecyclingInfo\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"BlockProducerSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumICoinbaseExecution.BlockRewardDeprecatedType\",\"name\":\"deprecatedType\",\"type\":\"uint8\"}],\"name\":\"BlockRewardDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submittedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"CandidateGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"}],\"name\":\"CandidateRevokingTimestampUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"name\":\"CandidateTopupDeadlineUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"CandidatesRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateUpdateScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"DeprecatedRewardRecycleFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DeprecatedRewardRecycled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedFundReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedFundReleasingFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExpiryDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FastFinalityRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"FastFinalityRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numOfDays\",\"type\":\"uint256\"}],\"name\":\"MinEffectiveDaysOnwardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"StakingRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailedUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blockProducerRewardDeprecated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"bridgeOperatorRewardDeprecated\",\"type\":\"bool\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnjailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"periodEnding\",\"type\":\"bool\"}],\"name\":\"WrappedUpEpoch\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DEFAULT_ADDITION_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERIOD_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"checkJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"checkJailedAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus[]\",\"name\":\"consensusList\",\"type\":\"address[]\"}],\"name\":\"checkManyJailed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"candidateIds\",\"type\":\"address[]\"}],\"name\":\"checkManyJailedById\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"checkMiningRewardDeprecated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"checkMiningRewardDeprecatedAtPeriod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodStartAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExitLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"name\":\"execApplyValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"execBailOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"execChangeAdminAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"TConsensus\",\"name\":\"newConsensusAddr\",\"type\":\"address\"}],\"name\":\"execChangeConsensusAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"execChangeTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"execReleaseLockedFundForEmergencyExitRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secLeftToRevoke\",\"type\":\"uint256\"}],\"name\":\"execRequestEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secsLeft\",\"type\":\"uint256\"}],\"name\":\"execRequestRenounceCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cid\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"effectiveDaysOnwards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"name\":\"execRequestUpdateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newJailedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cannotBailout\",\"type\":\"bool\"}],\"name\":\"execSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockProducers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"getCandidateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"__shadowedAdmin\",\"type\":\"address\"},{\"internalType\":\"TConsensus\",\"name\":\"__shadowedConsensus\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"__shadowedTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"____deprecatedBridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"__shadowedAdmin\",\"type\":\"address\"},{\"internalType\":\"TConsensus\",\"name\":\"__shadowedConsensus\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"__shadowedTreasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"____deprecatedBridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate[]\",\"name\":\"list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"getCommissionChangeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.CommissionSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"getEmergencyExitInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recyclingAt\",\"type\":\"uint256\"}],\"internalType\":\"structICommonInfo.EmergencyExitInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"getJailedTimeLeft\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"getJailedTimeLeftAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingVestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__minEffectiveDaysOnwards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"__emergencyExitConfigs\",\"type\":\"uint256[2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fastFinalityTrackingContract\",\"type\":\"address\"}],\"name\":\"initializeV3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileContract\",\"type\":\"address\"}],\"name\":\"initializeV4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"}],\"name\":\"isBlockProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"TConsensus\",\"name\":\"consensus\",\"type\":\"address\"}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minEffectiveDaysOnward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompilePickValidatorSetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileSortValidatorsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setEmergencyExitLockedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setEmergencyExpiryDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxPrioritizedValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numOfDays\",\"type\":\"uint256\"}],\"name\":\"setMinEffectiveDaysOnwards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlockProducer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeprecatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"tryGetPeriodOfEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_filled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_periodNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RoninValidatorSetABI is the input ABI used to generate the binding from.
// Deprecated: Use RoninValidatorSetMetaData.ABI instead.
var RoninValidatorSetABI = RoninValidatorSetMetaData.ABI

// RoninValidatorSet is an auto generated Go binding around an Ethereum contract.
type RoninValidatorSet struct {
	RoninValidatorSetCaller     // Read-only binding to the contract
	RoninValidatorSetTransactor // Write-only binding to the contract
	RoninValidatorSetFilterer   // Log filterer for contract events
}

// RoninValidatorSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoninValidatorSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoninValidatorSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoninValidatorSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoninValidatorSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoninValidatorSetSession struct {
	Contract     *RoninValidatorSet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RoninValidatorSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoninValidatorSetCallerSession struct {
	Contract *RoninValidatorSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RoninValidatorSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoninValidatorSetTransactorSession struct {
	Contract     *RoninValidatorSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RoninValidatorSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoninValidatorSetRaw struct {
	Contract *RoninValidatorSet // Generic contract binding to access the raw methods on
}

// RoninValidatorSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoninValidatorSetCallerRaw struct {
	Contract *RoninValidatorSetCaller // Generic read-only contract binding to access the raw methods on
}

// RoninValidatorSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoninValidatorSetTransactorRaw struct {
	Contract *RoninValidatorSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoninValidatorSet creates a new instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSet(address common.Address, backend bind.ContractBackend) (*RoninValidatorSet, error) {
	contract, err := bindRoninValidatorSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSet{RoninValidatorSetCaller: RoninValidatorSetCaller{contract: contract}, RoninValidatorSetTransactor: RoninValidatorSetTransactor{contract: contract}, RoninValidatorSetFilterer: RoninValidatorSetFilterer{contract: contract}}, nil
}

// NewRoninValidatorSetCaller creates a new read-only instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetCaller(address common.Address, caller bind.ContractCaller) (*RoninValidatorSetCaller, error) {
	contract, err := bindRoninValidatorSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCaller{contract: contract}, nil
}

// NewRoninValidatorSetTransactor creates a new write-only instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetTransactor(address common.Address, transactor bind.ContractTransactor) (*RoninValidatorSetTransactor, error) {
	contract, err := bindRoninValidatorSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetTransactor{contract: contract}, nil
}

// NewRoninValidatorSetFilterer creates a new log filterer instance of RoninValidatorSet, bound to a specific deployed contract.
func NewRoninValidatorSetFilterer(address common.Address, filterer bind.ContractFilterer) (*RoninValidatorSetFilterer, error) {
	contract, err := bindRoninValidatorSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetFilterer{contract: contract}, nil
}

// bindRoninValidatorSet binds a generic wrapper to an already deployed contract.
func bindRoninValidatorSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoninValidatorSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninValidatorSet *RoninValidatorSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninValidatorSet.Contract.RoninValidatorSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninValidatorSet *RoninValidatorSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RoninValidatorSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninValidatorSet *RoninValidatorSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.RoninValidatorSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoninValidatorSet *RoninValidatorSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoninValidatorSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoninValidatorSet *RoninValidatorSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoninValidatorSet *RoninValidatorSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) DEFAULTADDITIONGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "DEFAULT_ADDITION_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _RoninValidatorSet.Contract.DEFAULTADDITIONGAS(&_RoninValidatorSet.CallOpts)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _RoninValidatorSet.Contract.DEFAULTADDITIONGAS(&_RoninValidatorSet.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) PERIODDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "PERIOD_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) PERIODDURATION() (*big.Int, error) {
	return _RoninValidatorSet.Contract.PERIODDURATION(&_RoninValidatorSet.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PERIODDURATION() (*big.Int, error) {
	return _RoninValidatorSet.Contract.PERIODDURATION(&_RoninValidatorSet.CallOpts)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckJailed(opts *bind.CallOpts, consensus common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkJailed", consensus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) CheckJailed(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.CheckJailed(&_RoninValidatorSet.CallOpts, consensus)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckJailed(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.CheckJailed(&_RoninValidatorSet.CallOpts, consensus)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address addr, uint256 blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckJailedAtBlock(opts *bind.CallOpts, addr common.Address, blockNum *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkJailedAtBlock", addr, blockNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address addr, uint256 blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) CheckJailedAtBlock(addr common.Address, blockNum *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.CheckJailedAtBlock(&_RoninValidatorSet.CallOpts, addr, blockNum)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address addr, uint256 blockNum) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckJailedAtBlock(addr common.Address, blockNum *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.CheckJailedAtBlock(&_RoninValidatorSet.CallOpts, addr, blockNum)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] consensusList) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckManyJailed(opts *bind.CallOpts, consensusList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkManyJailed", consensusList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] consensusList) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetSession) CheckManyJailed(consensusList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.CheckManyJailed(&_RoninValidatorSet.CallOpts, consensusList)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] consensusList) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckManyJailed(consensusList []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.CheckManyJailed(&_RoninValidatorSet.CallOpts, consensusList)
}

// CheckManyJailedById is a free data retrieval call binding the contract method 0x7e1dc16f.
//
// Solidity: function checkManyJailedById(address[] candidateIds) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckManyJailedById(opts *bind.CallOpts, candidateIds []common.Address) ([]bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkManyJailedById", candidateIds)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyJailedById is a free data retrieval call binding the contract method 0x7e1dc16f.
//
// Solidity: function checkManyJailedById(address[] candidateIds) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetSession) CheckManyJailedById(candidateIds []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.CheckManyJailedById(&_RoninValidatorSet.CallOpts, candidateIds)
}

// CheckManyJailedById is a free data retrieval call binding the contract method 0x7e1dc16f.
//
// Solidity: function checkManyJailedById(address[] candidateIds) view returns(bool[])
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckManyJailedById(candidateIds []common.Address) ([]bool, error) {
	return _RoninValidatorSet.Contract.CheckManyJailedById(&_RoninValidatorSet.CallOpts, candidateIds)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckMiningRewardDeprecated(opts *bind.CallOpts, consensus common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkMiningRewardDeprecated", consensus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) CheckMiningRewardDeprecated(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.CheckMiningRewardDeprecated(&_RoninValidatorSet.CallOpts, consensus)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckMiningRewardDeprecated(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.CheckMiningRewardDeprecated(&_RoninValidatorSet.CallOpts, consensus)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address consensus, uint256 period) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) CheckMiningRewardDeprecatedAtPeriod(opts *bind.CallOpts, consensus common.Address, period *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "checkMiningRewardDeprecatedAtPeriod", consensus, period)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address consensus, uint256 period) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) CheckMiningRewardDeprecatedAtPeriod(consensus common.Address, period *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.CheckMiningRewardDeprecatedAtPeriod(&_RoninValidatorSet.CallOpts, consensus, period)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address consensus, uint256 period) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CheckMiningRewardDeprecatedAtPeriod(consensus common.Address, period *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.CheckMiningRewardDeprecatedAtPeriod(&_RoninValidatorSet.CallOpts, consensus, period)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) CurrentPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriod(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CurrentPeriod() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriod(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) CurrentPeriodStartAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "currentPeriodStartAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriodStartAtBlock(&_RoninValidatorSet.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.CurrentPeriodStartAtBlock(&_RoninValidatorSet.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) EmergencyExitLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "emergencyExitLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.EmergencyExitLockedAmount(&_RoninValidatorSet.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.EmergencyExitLockedAmount(&_RoninValidatorSet.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) EmergencyExpiryDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "emergencyExpiryDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _RoninValidatorSet.Contract.EmergencyExpiryDuration(&_RoninValidatorSet.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _RoninValidatorSet.Contract.EmergencyExpiryDuration(&_RoninValidatorSet.CallOpts)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) EpochEndingAt(opts *bind.CallOpts, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "epochEndingAt", _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.EpochEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _RoninValidatorSet.Contract.EpochEndingAt(&_RoninValidatorSet.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) EpochOf(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "epochOf", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.EpochOf(&_RoninValidatorSet.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _RoninValidatorSet.Contract.EpochOf(&_RoninValidatorSet.CallOpts, _block)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] result)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetBlockProducers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getBlockProducers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] result)
func (_RoninValidatorSet *RoninValidatorSetSession) GetBlockProducers() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBlockProducers(&_RoninValidatorSet.CallOpts)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] result)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetBlockProducers() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetBlockProducers(&_RoninValidatorSet.CallOpts)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address consensus) view returns((address,address,address,address,uint256,uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetCaller) GetCandidateInfo(opts *bind.CallOpts, consensus common.Address) (ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getCandidateInfo", consensus)

	if err != nil {
		return *new(ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new(ICandidateManagerValidatorCandidate)).(*ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address consensus) view returns((address,address,address,address,uint256,uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetSession) GetCandidateInfo(consensus common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfo(&_RoninValidatorSet.CallOpts, consensus)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address consensus) view returns((address,address,address,address,uint256,uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetCandidateInfo(consensus common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfo(&_RoninValidatorSet.CallOpts, consensus)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] list)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetCandidateInfos(opts *bind.CallOpts) ([]ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getCandidateInfos")

	if err != nil {
		return *new([]ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICandidateManagerValidatorCandidate)).(*[]ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] list)
func (_RoninValidatorSet *RoninValidatorSetSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfos(&_RoninValidatorSet.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] list)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _RoninValidatorSet.Contract.GetCandidateInfos(&_RoninValidatorSet.CallOpts)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address consensus) view returns((uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetCaller) GetCommissionChangeSchedule(opts *bind.CallOpts, consensus common.Address) (ICandidateManagerCommissionSchedule, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getCommissionChangeSchedule", consensus)

	if err != nil {
		return *new(ICandidateManagerCommissionSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(ICandidateManagerCommissionSchedule)).(*ICandidateManagerCommissionSchedule)

	return out0, err

}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address consensus) view returns((uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetSession) GetCommissionChangeSchedule(consensus common.Address) (ICandidateManagerCommissionSchedule, error) {
	return _RoninValidatorSet.Contract.GetCommissionChangeSchedule(&_RoninValidatorSet.CallOpts, consensus)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address consensus) view returns((uint256,uint256))
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetCommissionChangeSchedule(consensus common.Address) (ICandidateManagerCommissionSchedule, error) {
	return _RoninValidatorSet.Contract.GetCommissionChangeSchedule(&_RoninValidatorSet.CallOpts, consensus)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninValidatorSet *RoninValidatorSetSession) GetContract(contractType uint8) (common.Address, error) {
	return _RoninValidatorSet.Contract.GetContract(&_RoninValidatorSet.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _RoninValidatorSet.Contract.GetContract(&_RoninValidatorSet.CallOpts, contractType)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address consensus) view returns((uint256,uint256) _info)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetEmergencyExitInfo(opts *bind.CallOpts, consensus common.Address) (ICommonInfoEmergencyExitInfo, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getEmergencyExitInfo", consensus)

	if err != nil {
		return *new(ICommonInfoEmergencyExitInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICommonInfoEmergencyExitInfo)).(*ICommonInfoEmergencyExitInfo)

	return out0, err

}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address consensus) view returns((uint256,uint256) _info)
func (_RoninValidatorSet *RoninValidatorSetSession) GetEmergencyExitInfo(consensus common.Address) (ICommonInfoEmergencyExitInfo, error) {
	return _RoninValidatorSet.Contract.GetEmergencyExitInfo(&_RoninValidatorSet.CallOpts, consensus)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address consensus) view returns((uint256,uint256) _info)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetEmergencyExitInfo(consensus common.Address) (ICommonInfoEmergencyExitInfo, error) {
	return _RoninValidatorSet.Contract.GetEmergencyExitInfo(&_RoninValidatorSet.CallOpts, consensus)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address consensus) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetJailedTimeLeft(opts *bind.CallOpts, consensus common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getJailedTimeLeft", consensus)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address consensus) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetSession) GetJailedTimeLeft(consensus common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.GetJailedTimeLeft(&_RoninValidatorSet.CallOpts, consensus)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address consensus) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetJailedTimeLeft(consensus common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.GetJailedTimeLeft(&_RoninValidatorSet.CallOpts, consensus)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address consensus, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetJailedTimeLeftAtBlock(opts *bind.CallOpts, consensus common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getJailedTimeLeftAtBlock", consensus, _blockNum)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address consensus, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetSession) GetJailedTimeLeftAtBlock(consensus common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.GetJailedTimeLeftAtBlock(&_RoninValidatorSet.CallOpts, consensus, _blockNum)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address consensus, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetJailedTimeLeftAtBlock(consensus common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _RoninValidatorSet.Contract.GetJailedTimeLeftAtBlock(&_RoninValidatorSet.CallOpts, consensus, _blockNum)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetLastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getLastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.GetLastUpdatedBlock(&_RoninValidatorSet.CallOpts)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _RoninValidatorSet.Contract.GetLastUpdatedBlock(&_RoninValidatorSet.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetCaller) GetValidatorCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getValidatorCandidates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetSession) GetValidatorCandidates() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidatorCandidates(&_RoninValidatorSet.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetValidatorCandidates() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidatorCandidates(&_RoninValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] validatorList_)
func (_RoninValidatorSet *RoninValidatorSetCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] validatorList_)
func (_RoninValidatorSet *RoninValidatorSetSession) GetValidators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidators(&_RoninValidatorSet.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] validatorList_)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) GetValidators() ([]common.Address, error) {
	return _RoninValidatorSet.Contract.GetValidators(&_RoninValidatorSet.CallOpts)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address consensusAddr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsBlockProducer(opts *bind.CallOpts, consensusAddr common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isBlockProducer", consensusAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address consensusAddr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsBlockProducer(consensusAddr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBlockProducer(&_RoninValidatorSet.CallOpts, consensusAddr)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address consensusAddr) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsBlockProducer(consensusAddr common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsBlockProducer(&_RoninValidatorSet.CallOpts, consensusAddr)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address consensusAddr, address admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsCandidateAdmin(opts *bind.CallOpts, consensusAddr common.Address, admin common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isCandidateAdmin", consensusAddr, admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address consensusAddr, address admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsCandidateAdmin(consensusAddr common.Address, admin common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsCandidateAdmin(&_RoninValidatorSet.CallOpts, consensusAddr, admin)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address consensusAddr, address admin) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsCandidateAdmin(consensusAddr common.Address, admin common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsCandidateAdmin(&_RoninValidatorSet.CallOpts, consensusAddr, admin)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsPeriodEnding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isPeriodEnding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsPeriodEnding() (bool, error) {
	return _RoninValidatorSet.Contract.IsPeriodEnding(&_RoninValidatorSet.CallOpts)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsPeriodEnding() (bool, error) {
	return _RoninValidatorSet.Contract.IsPeriodEnding(&_RoninValidatorSet.CallOpts)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCaller) IsValidatorCandidate(opts *bind.CallOpts, consensus common.Address) (bool, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "isValidatorCandidate", consensus)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetSession) IsValidatorCandidate(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidatorCandidate(&_RoninValidatorSet.CallOpts, consensus)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address consensus) view returns(bool)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) IsValidatorCandidate(consensus common.Address) (bool, error) {
	return _RoninValidatorSet.Contract.IsValidatorCandidate(&_RoninValidatorSet.CallOpts, consensus)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxPrioritizedValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxPrioritizedValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxPrioritizedValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxPrioritizedValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxValidatorCandidate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxValidatorCandidate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxValidatorCandidate() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorCandidate(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxValidatorCandidate() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorCandidate(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "maxValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetSession) MaxValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MaxValidatorNumber(&_RoninValidatorSet.CallOpts)
}

// MinEffectiveDaysOnward is a free data retrieval call binding the contract method 0x612c8d98.
//
// Solidity: function minEffectiveDaysOnward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) MinEffectiveDaysOnward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "minEffectiveDaysOnward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinEffectiveDaysOnward is a free data retrieval call binding the contract method 0x612c8d98.
//
// Solidity: function minEffectiveDaysOnward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) MinEffectiveDaysOnward() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MinEffectiveDaysOnward(&_RoninValidatorSet.CallOpts)
}

// MinEffectiveDaysOnward is a free data retrieval call binding the contract method 0x612c8d98.
//
// Solidity: function minEffectiveDaysOnward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) MinEffectiveDaysOnward() (*big.Int, error) {
	return _RoninValidatorSet.Contract.MinEffectiveDaysOnward(&_RoninValidatorSet.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetCaller) NumberOfBlocksInEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "numberOfBlocksInEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfBlocksInEpoch(&_RoninValidatorSet.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _RoninValidatorSet.Contract.NumberOfBlocksInEpoch(&_RoninValidatorSet.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) PrecompilePickValidatorSetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "precompilePickValidatorSetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompilePickValidatorSetAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompilePickValidatorSetAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCaller) PrecompileSortValidatorsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "precompileSortValidatorsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompileSortValidatorsAddress(&_RoninValidatorSet.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _RoninValidatorSet.Contract.PrecompileSortValidatorsAddress(&_RoninValidatorSet.CallOpts)
}

// TotalBlockProducer is a free data retrieval call binding the contract method 0xafc69d73.
//
// Solidity: function totalBlockProducer() view returns(uint256 total)
func (_RoninValidatorSet *RoninValidatorSetCaller) TotalBlockProducer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "totalBlockProducer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBlockProducer is a free data retrieval call binding the contract method 0xafc69d73.
//
// Solidity: function totalBlockProducer() view returns(uint256 total)
func (_RoninValidatorSet *RoninValidatorSetSession) TotalBlockProducer() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBlockProducer(&_RoninValidatorSet.CallOpts)
}

// TotalBlockProducer is a free data retrieval call binding the contract method 0xafc69d73.
//
// Solidity: function totalBlockProducer() view returns(uint256 total)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) TotalBlockProducer() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalBlockProducer(&_RoninValidatorSet.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) TotalDeprecatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "totalDeprecatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) TotalDeprecatedReward() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalDeprecatedReward(&_RoninValidatorSet.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) TotalDeprecatedReward() (*big.Int, error) {
	return _RoninValidatorSet.Contract.TotalDeprecatedReward(&_RoninValidatorSet.CallOpts)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_RoninValidatorSet *RoninValidatorSetCaller) TryGetPeriodOfEpoch(opts *bind.CallOpts, _epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "tryGetPeriodOfEpoch", _epoch)

	outstruct := new(struct {
		Filled       bool
		PeriodNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Filled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PeriodNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_RoninValidatorSet *RoninValidatorSetSession) TryGetPeriodOfEpoch(_epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	return _RoninValidatorSet.Contract.TryGetPeriodOfEpoch(&_RoninValidatorSet.CallOpts, _epoch)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) TryGetPeriodOfEpoch(_epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	return _RoninValidatorSet.Contract.TryGetPeriodOfEpoch(&_RoninValidatorSet.CallOpts, _epoch)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoninValidatorSet.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetSession) ValidatorCount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.ValidatorCount(&_RoninValidatorSet.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_RoninValidatorSet *RoninValidatorSetCallerSession) ValidatorCount() (*big.Int, error) {
	return _RoninValidatorSet.Contract.ValidatorCount(&_RoninValidatorSet.CallOpts)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xf883afaf.
//
// Solidity: function execApplyValidatorCandidate(address candidateAdmin, address cid, address treasuryAddr, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecApplyValidatorCandidate(opts *bind.TransactOpts, candidateAdmin common.Address, cid common.Address, treasuryAddr common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execApplyValidatorCandidate", candidateAdmin, cid, treasuryAddr, commissionRate)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xf883afaf.
//
// Solidity: function execApplyValidatorCandidate(address candidateAdmin, address cid, address treasuryAddr, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecApplyValidatorCandidate(candidateAdmin common.Address, cid common.Address, treasuryAddr common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecApplyValidatorCandidate(&_RoninValidatorSet.TransactOpts, candidateAdmin, cid, treasuryAddr, commissionRate)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xf883afaf.
//
// Solidity: function execApplyValidatorCandidate(address candidateAdmin, address cid, address treasuryAddr, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecApplyValidatorCandidate(candidateAdmin common.Address, cid common.Address, treasuryAddr common.Address, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecApplyValidatorCandidate(&_RoninValidatorSet.TransactOpts, candidateAdmin, cid, treasuryAddr, commissionRate)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address validatorId, uint256 period) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecBailOut(opts *bind.TransactOpts, validatorId common.Address, period *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execBailOut", validatorId, period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address validatorId, uint256 period) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecBailOut(validatorId common.Address, period *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecBailOut(&_RoninValidatorSet.TransactOpts, validatorId, period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address validatorId, uint256 period) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecBailOut(validatorId common.Address, period *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecBailOut(&_RoninValidatorSet.TransactOpts, validatorId, period)
}

// ExecChangeAdminAddress is a paid mutator transaction binding the contract method 0xea80d67d.
//
// Solidity: function execChangeAdminAddress(address cid, address newAdmin) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecChangeAdminAddress(opts *bind.TransactOpts, cid common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execChangeAdminAddress", cid, newAdmin)
}

// ExecChangeAdminAddress is a paid mutator transaction binding the contract method 0xea80d67d.
//
// Solidity: function execChangeAdminAddress(address cid, address newAdmin) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecChangeAdminAddress(cid common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeAdminAddress(&_RoninValidatorSet.TransactOpts, cid, newAdmin)
}

// ExecChangeAdminAddress is a paid mutator transaction binding the contract method 0xea80d67d.
//
// Solidity: function execChangeAdminAddress(address cid, address newAdmin) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecChangeAdminAddress(cid common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeAdminAddress(&_RoninValidatorSet.TransactOpts, cid, newAdmin)
}

// ExecChangeConsensusAddress is a paid mutator transaction binding the contract method 0xd93fa306.
//
// Solidity: function execChangeConsensusAddress(address cid, address newConsensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecChangeConsensusAddress(opts *bind.TransactOpts, cid common.Address, newConsensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execChangeConsensusAddress", cid, newConsensusAddr)
}

// ExecChangeConsensusAddress is a paid mutator transaction binding the contract method 0xd93fa306.
//
// Solidity: function execChangeConsensusAddress(address cid, address newConsensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecChangeConsensusAddress(cid common.Address, newConsensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeConsensusAddress(&_RoninValidatorSet.TransactOpts, cid, newConsensusAddr)
}

// ExecChangeConsensusAddress is a paid mutator transaction binding the contract method 0xd93fa306.
//
// Solidity: function execChangeConsensusAddress(address cid, address newConsensusAddr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecChangeConsensusAddress(cid common.Address, newConsensusAddr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeConsensusAddress(&_RoninValidatorSet.TransactOpts, cid, newConsensusAddr)
}

// ExecChangeTreasuryAddress is a paid mutator transaction binding the contract method 0xecd850cc.
//
// Solidity: function execChangeTreasuryAddress(address cid, address newTreasury) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecChangeTreasuryAddress(opts *bind.TransactOpts, cid common.Address, newTreasury common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execChangeTreasuryAddress", cid, newTreasury)
}

// ExecChangeTreasuryAddress is a paid mutator transaction binding the contract method 0xecd850cc.
//
// Solidity: function execChangeTreasuryAddress(address cid, address newTreasury) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecChangeTreasuryAddress(cid common.Address, newTreasury common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeTreasuryAddress(&_RoninValidatorSet.TransactOpts, cid, newTreasury)
}

// ExecChangeTreasuryAddress is a paid mutator transaction binding the contract method 0xecd850cc.
//
// Solidity: function execChangeTreasuryAddress(address cid, address newTreasury) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecChangeTreasuryAddress(cid common.Address, newTreasury common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecChangeTreasuryAddress(&_RoninValidatorSet.TransactOpts, cid, newTreasury)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address cid, address recipient) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecReleaseLockedFundForEmergencyExitRequest(opts *bind.TransactOpts, cid common.Address, recipient common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execReleaseLockedFundForEmergencyExitRequest", cid, recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address cid, address recipient) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecReleaseLockedFundForEmergencyExitRequest(cid common.Address, recipient common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_RoninValidatorSet.TransactOpts, cid, recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address cid, address recipient) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecReleaseLockedFundForEmergencyExitRequest(cid common.Address, recipient common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_RoninValidatorSet.TransactOpts, cid, recipient)
}

// ExecRequestEmergencyExit is a paid mutator transaction binding the contract method 0xcdaa4e81.
//
// Solidity: function execRequestEmergencyExit(address cid, uint256 secLeftToRevoke) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecRequestEmergencyExit(opts *bind.TransactOpts, cid common.Address, secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execRequestEmergencyExit", cid, secLeftToRevoke)
}

// ExecRequestEmergencyExit is a paid mutator transaction binding the contract method 0xcdaa4e81.
//
// Solidity: function execRequestEmergencyExit(address cid, uint256 secLeftToRevoke) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecRequestEmergencyExit(cid common.Address, secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestEmergencyExit(&_RoninValidatorSet.TransactOpts, cid, secLeftToRevoke)
}

// ExecRequestEmergencyExit is a paid mutator transaction binding the contract method 0xcdaa4e81.
//
// Solidity: function execRequestEmergencyExit(address cid, uint256 secLeftToRevoke) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecRequestEmergencyExit(cid common.Address, secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestEmergencyExit(&_RoninValidatorSet.TransactOpts, cid, secLeftToRevoke)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address cid, uint256 secsLeft) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecRequestRenounceCandidate(opts *bind.TransactOpts, cid common.Address, secsLeft *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execRequestRenounceCandidate", cid, secsLeft)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address cid, uint256 secsLeft) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecRequestRenounceCandidate(cid common.Address, secsLeft *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestRenounceCandidate(&_RoninValidatorSet.TransactOpts, cid, secsLeft)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address cid, uint256 secsLeft) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecRequestRenounceCandidate(cid common.Address, secsLeft *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestRenounceCandidate(&_RoninValidatorSet.TransactOpts, cid, secsLeft)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address cid, uint256 effectiveDaysOnwards, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecRequestUpdateCommissionRate(opts *bind.TransactOpts, cid common.Address, effectiveDaysOnwards *big.Int, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execRequestUpdateCommissionRate", cid, effectiveDaysOnwards, commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address cid, uint256 effectiveDaysOnwards, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecRequestUpdateCommissionRate(cid common.Address, effectiveDaysOnwards *big.Int, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestUpdateCommissionRate(&_RoninValidatorSet.TransactOpts, cid, effectiveDaysOnwards, commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address cid, uint256 effectiveDaysOnwards, uint256 commissionRate) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecRequestUpdateCommissionRate(cid common.Address, effectiveDaysOnwards *big.Int, commissionRate *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecRequestUpdateCommissionRate(&_RoninValidatorSet.TransactOpts, cid, effectiveDaysOnwards, commissionRate)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address validatorId, uint256 newJailedUntil, uint256 slashAmount, bool cannotBailout) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) ExecSlash(opts *bind.TransactOpts, validatorId common.Address, newJailedUntil *big.Int, slashAmount *big.Int, cannotBailout bool) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "execSlash", validatorId, newJailedUntil, slashAmount, cannotBailout)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address validatorId, uint256 newJailedUntil, uint256 slashAmount, bool cannotBailout) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) ExecSlash(validatorId common.Address, newJailedUntil *big.Int, slashAmount *big.Int, cannotBailout bool) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecSlash(&_RoninValidatorSet.TransactOpts, validatorId, newJailedUntil, slashAmount, cannotBailout)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address validatorId, uint256 newJailedUntil, uint256 slashAmount, bool cannotBailout) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) ExecSlash(validatorId common.Address, newJailedUntil *big.Int, slashAmount *big.Int, cannotBailout bool) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.ExecSlash(&_RoninValidatorSet.TransactOpts, validatorId, newJailedUntil, slashAmount, cannotBailout)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address , uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, arg5 common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, arg5, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address , uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, arg5 common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, arg5, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address , uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, arg5 common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Initialize(&_RoninValidatorSet.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, arg5, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RoninValidatorSet *RoninValidatorSetSession) InitializeV2() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV2(&_RoninValidatorSet.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV2(&_RoninValidatorSet.TransactOpts)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address fastFinalityTrackingContract) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) InitializeV3(opts *bind.TransactOpts, fastFinalityTrackingContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initializeV3", fastFinalityTrackingContract)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address fastFinalityTrackingContract) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) InitializeV3(fastFinalityTrackingContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV3(&_RoninValidatorSet.TransactOpts, fastFinalityTrackingContract)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0x3101cfcb.
//
// Solidity: function initializeV3(address fastFinalityTrackingContract) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) InitializeV3(fastFinalityTrackingContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV3(&_RoninValidatorSet.TransactOpts, fastFinalityTrackingContract)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x110a8308.
//
// Solidity: function initializeV4(address profileContract) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) InitializeV4(opts *bind.TransactOpts, profileContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "initializeV4", profileContract)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x110a8308.
//
// Solidity: function initializeV4(address profileContract) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) InitializeV4(profileContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV4(&_RoninValidatorSet.TransactOpts, profileContract)
}

// InitializeV4 is a paid mutator transaction binding the contract method 0x110a8308.
//
// Solidity: function initializeV4(address profileContract) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) InitializeV4(profileContract common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.InitializeV4(&_RoninValidatorSet.TransactOpts, profileContract)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetContract(&_RoninValidatorSet.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetContract(&_RoninValidatorSet.TransactOpts, contractType, addr)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 amount) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetEmergencyExitLockedAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setEmergencyExitLockedAmount", amount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 amount) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetEmergencyExitLockedAmount(amount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetEmergencyExitLockedAmount(&_RoninValidatorSet.TransactOpts, amount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 amount) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetEmergencyExitLockedAmount(amount *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetEmergencyExitLockedAmount(&_RoninValidatorSet.TransactOpts, amount)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 duration) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetEmergencyExpiryDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setEmergencyExpiryDuration", duration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 duration) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetEmergencyExpiryDuration(duration *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetEmergencyExpiryDuration(&_RoninValidatorSet.TransactOpts, duration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 duration) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetEmergencyExpiryDuration(duration *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetEmergencyExpiryDuration(&_RoninValidatorSet.TransactOpts, duration)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxPrioritizedValidatorNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxPrioritizedValidatorNumber", _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxPrioritizedValidatorNumber(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxPrioritizedValidatorNumber(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxValidatorCandidate(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxValidatorCandidate", _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorCandidate(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorCandidate(&_RoninValidatorSet.TransactOpts, _number)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMaxValidatorNumber(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMaxValidatorNumber", _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMaxValidatorNumber(&_RoninValidatorSet.TransactOpts, _max)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SetMinEffectiveDaysOnwards(opts *bind.TransactOpts, _numOfDays *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "setMinEffectiveDaysOnwards", _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMinEffectiveDaysOnwards(&_RoninValidatorSet.TransactOpts, _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SetMinEffectiveDaysOnwards(&_RoninValidatorSet.TransactOpts, _numOfDays)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) SubmitBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "submitBlockReward")
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) SubmitBlockReward() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SubmitBlockReward(&_RoninValidatorSet.TransactOpts)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.SubmitBlockReward(&_RoninValidatorSet.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) WrapUpEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.Transact(opts, "wrapUpEpoch")
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) WrapUpEpoch() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.WrapUpEpoch(&_RoninValidatorSet.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.WrapUpEpoch(&_RoninValidatorSet.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Fallback(&_RoninValidatorSet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Fallback(&_RoninValidatorSet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoninValidatorSet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetSession) Receive() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Receive(&_RoninValidatorSet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RoninValidatorSet *RoninValidatorSetTransactorSession) Receive() (*types.Transaction, error) {
	return _RoninValidatorSet.Contract.Receive(&_RoninValidatorSet.TransactOpts)
}

// RoninValidatorSetBlockProducerSetUpdatedIterator is returned from FilterBlockProducerSetUpdated and is used to iterate over the raw logs and unpacked data for BlockProducerSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockProducerSetUpdatedIterator struct {
	Event *RoninValidatorSetBlockProducerSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockProducerSetUpdated)
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
		it.Event = new(RoninValidatorSetBlockProducerSetUpdated)
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
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockProducerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockProducerSetUpdated represents a BlockProducerSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockProducerSetUpdated struct {
	Period         *big.Int
	Epoch          *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerSetUpdated is a free log retrieval operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockProducerSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*RoninValidatorSetBlockProducerSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockProducerSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "BlockProducerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerSetUpdated is a free log subscription operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockProducerSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockProducerSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockProducerSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
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

// ParseBlockProducerSetUpdated is a log parse operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockProducerSetUpdated(log types.Log) (*RoninValidatorSetBlockProducerSetUpdated, error) {
	event := new(RoninValidatorSetBlockProducerSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBlockRewardDeprecatedIterator is returned from FilterBlockRewardDeprecated and is used to iterate over the raw logs and unpacked data for BlockRewardDeprecated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardDeprecatedIterator struct {
	Event *RoninValidatorSetBlockRewardDeprecated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockRewardDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockRewardDeprecated)
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
		it.Event = new(RoninValidatorSetBlockRewardDeprecated)
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
func (it *RoninValidatorSetBlockRewardDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockRewardDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockRewardDeprecated represents a BlockRewardDeprecated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardDeprecated struct {
	CoinbaseAddr   common.Address
	RewardAmount   *big.Int
	DeprecatedType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardDeprecated is a free log retrieval operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockRewardDeprecated(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*RoninValidatorSetBlockRewardDeprecatedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockRewardDeprecatedIterator{contract: _RoninValidatorSet.contract, event: "BlockRewardDeprecated", logs: logs, sub: sub}, nil
}

// WatchBlockRewardDeprecated is a free log subscription operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockRewardDeprecated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockRewardDeprecated, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockRewardDeprecated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
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

// ParseBlockRewardDeprecated is a log parse operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockRewardDeprecated(log types.Log) (*RoninValidatorSetBlockRewardDeprecated, error) {
	event := new(RoninValidatorSetBlockRewardDeprecated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBlockRewardSubmittedIterator is returned from FilterBlockRewardSubmitted and is used to iterate over the raw logs and unpacked data for BlockRewardSubmitted events raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardSubmittedIterator struct {
	Event *RoninValidatorSetBlockRewardSubmitted // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBlockRewardSubmitted)
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
		it.Event = new(RoninValidatorSetBlockRewardSubmitted)
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
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBlockRewardSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBlockRewardSubmitted represents a BlockRewardSubmitted event raised by the RoninValidatorSet contract.
type RoninValidatorSetBlockRewardSubmitted struct {
	CoinbaseAddr    common.Address
	SubmittedAmount *big.Int
	BonusAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardSubmitted is a free log retrieval operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBlockRewardSubmitted(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*RoninValidatorSetBlockRewardSubmittedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBlockRewardSubmittedIterator{contract: _RoninValidatorSet.contract, event: "BlockRewardSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockRewardSubmitted is a free log subscription operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBlockRewardSubmitted(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBlockRewardSubmitted, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBlockRewardSubmitted)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
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

// ParseBlockRewardSubmitted is a log parse operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBlockRewardSubmitted(log types.Log) (*RoninValidatorSetBlockRewardSubmitted, error) {
	event := new(RoninValidatorSetBlockRewardSubmitted)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorRewardDistributedIterator is returned from FilterBridgeOperatorRewardDistributed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributedIterator struct {
	Event *RoninValidatorSetBridgeOperatorRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributed)
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
		it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributed)
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
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorRewardDistributed represents a BridgeOperatorRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributed struct {
	ConsensusAddr  common.Address
	BridgeOperator common.Address
	RecipientAddr  common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributed is a free log retrieval operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (*RoninValidatorSetBridgeOperatorRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributed is a free log subscription operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorRewardDistributed, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributed is a log parse operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorRewardDistributed(log types.Log) (*RoninValidatorSetBridgeOperatorRewardDistributed, error) {
	event := new(RoninValidatorSetBridgeOperatorRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator is returned from FilterBridgeOperatorRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetBridgeOperatorRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
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
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorRewardDistributionFailed represents a BridgeOperatorRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	BridgeOperator  common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributionFailed is a free log retrieval operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (*RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributionFailed is a free log subscription operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorRewardDistributionFailed, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributionFailed is a log parse operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorRewardDistributionFailed(log types.Log) (*RoninValidatorSetBridgeOperatorRewardDistributionFailed, error) {
	event := new(RoninValidatorSetBridgeOperatorRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetBridgeOperatorSetUpdatedIterator is returned from FilterBridgeOperatorSetUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorSetUpdatedIterator struct {
	Event *RoninValidatorSetBridgeOperatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetBridgeOperatorSetUpdated)
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
		it.Event = new(RoninValidatorSetBridgeOperatorSetUpdated)
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
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetBridgeOperatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetBridgeOperatorSetUpdated represents a BridgeOperatorSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetBridgeOperatorSetUpdated struct {
	Period          *big.Int
	Epoch           *big.Int
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSetUpdated is a free log retrieval operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterBridgeOperatorSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*RoninValidatorSetBridgeOperatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetBridgeOperatorSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "BridgeOperatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSetUpdated is a free log subscription operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchBridgeOperatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetBridgeOperatorSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetBridgeOperatorSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
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

// ParseBridgeOperatorSetUpdated is a log parse operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseBridgeOperatorSetUpdated(log types.Log) (*RoninValidatorSetBridgeOperatorSetUpdated, error) {
	event := new(RoninValidatorSetBridgeOperatorSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidateGrantedIterator is returned from FilterCandidateGranted and is used to iterate over the raw logs and unpacked data for CandidateGranted events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateGrantedIterator struct {
	Event *RoninValidatorSetCandidateGranted // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidateGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidateGranted)
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
		it.Event = new(RoninValidatorSetCandidateGranted)
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
func (it *RoninValidatorSetCandidateGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidateGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidateGranted represents a CandidateGranted event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateGranted struct {
	ConsensusAddr common.Address
	TreasuryAddr  common.Address
	Admin         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCandidateGranted is a free log retrieval operation binding the contract event 0x1ca451a9920472b99355a9cf74185bf017604a7849c113f020888ecec9db9366.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidateGranted(opts *bind.FilterOpts, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (*RoninValidatorSetCandidateGrantedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidateGrantedIterator{contract: _RoninValidatorSet.contract, event: "CandidateGranted", logs: logs, sub: sub}, nil
}

// WatchCandidateGranted is a free log subscription operation binding the contract event 0x1ca451a9920472b99355a9cf74185bf017604a7849c113f020888ecec9db9366.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidateGranted(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidateGranted, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidateGranted)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
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

// ParseCandidateGranted is a log parse operation binding the contract event 0x1ca451a9920472b99355a9cf74185bf017604a7849c113f020888ecec9db9366.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidateGranted(log types.Log) (*RoninValidatorSetCandidateGranted, error) {
	event := new(RoninValidatorSetCandidateGranted)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidateRevokingTimestampUpdatedIterator is returned from FilterCandidateRevokingTimestampUpdated and is used to iterate over the raw logs and unpacked data for CandidateRevokingTimestampUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateRevokingTimestampUpdatedIterator struct {
	Event *RoninValidatorSetCandidateRevokingTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidateRevokingTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidateRevokingTimestampUpdated)
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
		it.Event = new(RoninValidatorSetCandidateRevokingTimestampUpdated)
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
func (it *RoninValidatorSetCandidateRevokingTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidateRevokingTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidateRevokingTimestampUpdated represents a CandidateRevokingTimestampUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateRevokingTimestampUpdated struct {
	Cid               common.Address
	RevokingTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCandidateRevokingTimestampUpdated is a free log retrieval operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed cid, uint256 revokingTimestamp)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidateRevokingTimestampUpdated(opts *bind.FilterOpts, cid []common.Address) (*RoninValidatorSetCandidateRevokingTimestampUpdatedIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidateRevokingTimestampUpdated", cidRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidateRevokingTimestampUpdatedIterator{contract: _RoninValidatorSet.contract, event: "CandidateRevokingTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateRevokingTimestampUpdated is a free log subscription operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed cid, uint256 revokingTimestamp)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidateRevokingTimestampUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidateRevokingTimestampUpdated, cid []common.Address) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidateRevokingTimestampUpdated", cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidateRevokingTimestampUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
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

// ParseCandidateRevokingTimestampUpdated is a log parse operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed cid, uint256 revokingTimestamp)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidateRevokingTimestampUpdated(log types.Log) (*RoninValidatorSetCandidateRevokingTimestampUpdated, error) {
	event := new(RoninValidatorSetCandidateRevokingTimestampUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidateTopupDeadlineUpdatedIterator is returned from FilterCandidateTopupDeadlineUpdated and is used to iterate over the raw logs and unpacked data for CandidateTopupDeadlineUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateTopupDeadlineUpdatedIterator struct {
	Event *RoninValidatorSetCandidateTopupDeadlineUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidateTopupDeadlineUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidateTopupDeadlineUpdated)
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
		it.Event = new(RoninValidatorSetCandidateTopupDeadlineUpdated)
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
func (it *RoninValidatorSetCandidateTopupDeadlineUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidateTopupDeadlineUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidateTopupDeadlineUpdated represents a CandidateTopupDeadlineUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidateTopupDeadlineUpdated struct {
	Cid           common.Address
	TopupDeadline *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCandidateTopupDeadlineUpdated is a free log retrieval operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed cid, uint256 topupDeadline)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidateTopupDeadlineUpdated(opts *bind.FilterOpts, cid []common.Address) (*RoninValidatorSetCandidateTopupDeadlineUpdatedIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidateTopupDeadlineUpdated", cidRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidateTopupDeadlineUpdatedIterator{contract: _RoninValidatorSet.contract, event: "CandidateTopupDeadlineUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateTopupDeadlineUpdated is a free log subscription operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed cid, uint256 topupDeadline)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidateTopupDeadlineUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidateTopupDeadlineUpdated, cid []common.Address) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidateTopupDeadlineUpdated", cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidateTopupDeadlineUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
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

// ParseCandidateTopupDeadlineUpdated is a log parse operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed cid, uint256 topupDeadline)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidateTopupDeadlineUpdated(log types.Log) (*RoninValidatorSetCandidateTopupDeadlineUpdated, error) {
	event := new(RoninValidatorSetCandidateTopupDeadlineUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCandidatesRevokedIterator is returned from FilterCandidatesRevoked and is used to iterate over the raw logs and unpacked data for CandidatesRevoked events raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidatesRevokedIterator struct {
	Event *RoninValidatorSetCandidatesRevoked // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCandidatesRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCandidatesRevoked)
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
		it.Event = new(RoninValidatorSetCandidatesRevoked)
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
func (it *RoninValidatorSetCandidatesRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCandidatesRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCandidatesRevoked represents a CandidatesRevoked event raised by the RoninValidatorSet contract.
type RoninValidatorSetCandidatesRevoked struct {
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidatesRevoked is a free log retrieval operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCandidatesRevoked(opts *bind.FilterOpts) (*RoninValidatorSetCandidatesRevokedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCandidatesRevokedIterator{contract: _RoninValidatorSet.contract, event: "CandidatesRevoked", logs: logs, sub: sub}, nil
}

// WatchCandidatesRevoked is a free log subscription operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCandidatesRevoked(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCandidatesRevoked) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCandidatesRevoked)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
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

// ParseCandidatesRevoked is a log parse operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCandidatesRevoked(log types.Log) (*RoninValidatorSetCandidatesRevoked, error) {
	event := new(RoninValidatorSetCandidatesRevoked)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCommissionRateUpdateScheduledIterator is returned from FilterCommissionRateUpdateScheduled and is used to iterate over the raw logs and unpacked data for CommissionRateUpdateScheduled events raised by the RoninValidatorSet contract.
type RoninValidatorSetCommissionRateUpdateScheduledIterator struct {
	Event *RoninValidatorSetCommissionRateUpdateScheduled // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCommissionRateUpdateScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCommissionRateUpdateScheduled)
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
		it.Event = new(RoninValidatorSetCommissionRateUpdateScheduled)
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
func (it *RoninValidatorSetCommissionRateUpdateScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCommissionRateUpdateScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCommissionRateUpdateScheduled represents a CommissionRateUpdateScheduled event raised by the RoninValidatorSet contract.
type RoninValidatorSetCommissionRateUpdateScheduled struct {
	ConsensusAddr      common.Address
	EffectiveTimestamp *big.Int
	Rate               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdateScheduled is a free log retrieval operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCommissionRateUpdateScheduled(opts *bind.FilterOpts, consensusAddr []common.Address) (*RoninValidatorSetCommissionRateUpdateScheduledIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCommissionRateUpdateScheduledIterator{contract: _RoninValidatorSet.contract, event: "CommissionRateUpdateScheduled", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdateScheduled is a free log subscription operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCommissionRateUpdateScheduled(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCommissionRateUpdateScheduled, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCommissionRateUpdateScheduled)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
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

// ParseCommissionRateUpdateScheduled is a log parse operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCommissionRateUpdateScheduled(log types.Log) (*RoninValidatorSetCommissionRateUpdateScheduled, error) {
	event := new(RoninValidatorSetCommissionRateUpdateScheduled)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetCommissionRateUpdatedIterator is returned from FilterCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for CommissionRateUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetCommissionRateUpdatedIterator struct {
	Event *RoninValidatorSetCommissionRateUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetCommissionRateUpdated)
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
		it.Event = new(RoninValidatorSetCommissionRateUpdated)
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
func (it *RoninValidatorSetCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetCommissionRateUpdated represents a CommissionRateUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetCommissionRateUpdated struct {
	ConsensusAddr common.Address
	Rate          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdated is a free log retrieval operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterCommissionRateUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*RoninValidatorSetCommissionRateUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetCommissionRateUpdatedIterator{contract: _RoninValidatorSet.contract, event: "CommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdated is a free log subscription operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetCommissionRateUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetCommissionRateUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
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

// ParseCommissionRateUpdated is a log parse operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseCommissionRateUpdated(log types.Log) (*RoninValidatorSetCommissionRateUpdated, error) {
	event := new(RoninValidatorSetCommissionRateUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetContractUpdatedIterator struct {
	Event *RoninValidatorSetContractUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetContractUpdated)
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
		it.Event = new(RoninValidatorSetContractUpdated)
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
func (it *RoninValidatorSetContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetContractUpdated represents a ContractUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*RoninValidatorSetContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetContractUpdatedIterator{contract: _RoninValidatorSet.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetContractUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseContractUpdated(log types.Log) (*RoninValidatorSetContractUpdated, error) {
	event := new(RoninValidatorSetContractUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetDeprecatedRewardRecycleFailedIterator is returned from FilterDeprecatedRewardRecycleFailed and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycleFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetDeprecatedRewardRecycleFailedIterator struct {
	Event *RoninValidatorSetDeprecatedRewardRecycleFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetDeprecatedRewardRecycleFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetDeprecatedRewardRecycleFailed)
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
		it.Event = new(RoninValidatorSetDeprecatedRewardRecycleFailed)
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
func (it *RoninValidatorSetDeprecatedRewardRecycleFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetDeprecatedRewardRecycleFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetDeprecatedRewardRecycleFailed represents a DeprecatedRewardRecycleFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetDeprecatedRewardRecycleFailed struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycleFailed is a free log retrieval operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterDeprecatedRewardRecycleFailed(opts *bind.FilterOpts, recipientAddr []common.Address) (*RoninValidatorSetDeprecatedRewardRecycleFailedIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetDeprecatedRewardRecycleFailedIterator{contract: _RoninValidatorSet.contract, event: "DeprecatedRewardRecycleFailed", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycleFailed is a free log subscription operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchDeprecatedRewardRecycleFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetDeprecatedRewardRecycleFailed, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetDeprecatedRewardRecycleFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
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

// ParseDeprecatedRewardRecycleFailed is a log parse operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseDeprecatedRewardRecycleFailed(log types.Log) (*RoninValidatorSetDeprecatedRewardRecycleFailed, error) {
	event := new(RoninValidatorSetDeprecatedRewardRecycleFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetDeprecatedRewardRecycledIterator is returned from FilterDeprecatedRewardRecycled and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycled events raised by the RoninValidatorSet contract.
type RoninValidatorSetDeprecatedRewardRecycledIterator struct {
	Event *RoninValidatorSetDeprecatedRewardRecycled // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetDeprecatedRewardRecycledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetDeprecatedRewardRecycled)
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
		it.Event = new(RoninValidatorSetDeprecatedRewardRecycled)
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
func (it *RoninValidatorSetDeprecatedRewardRecycledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetDeprecatedRewardRecycledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetDeprecatedRewardRecycled represents a DeprecatedRewardRecycled event raised by the RoninValidatorSet contract.
type RoninValidatorSetDeprecatedRewardRecycled struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycled is a free log retrieval operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterDeprecatedRewardRecycled(opts *bind.FilterOpts, recipientAddr []common.Address) (*RoninValidatorSetDeprecatedRewardRecycledIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetDeprecatedRewardRecycledIterator{contract: _RoninValidatorSet.contract, event: "DeprecatedRewardRecycled", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycled is a free log subscription operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchDeprecatedRewardRecycled(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetDeprecatedRewardRecycled, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetDeprecatedRewardRecycled)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
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

// ParseDeprecatedRewardRecycled is a log parse operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseDeprecatedRewardRecycled(log types.Log) (*RoninValidatorSetDeprecatedRewardRecycled, error) {
	event := new(RoninValidatorSetDeprecatedRewardRecycled)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator is returned from FilterEmergencyExitLockedAmountUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedAmountUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator struct {
	Event *RoninValidatorSetEmergencyExitLockedAmountUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetEmergencyExitLockedAmountUpdated)
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
		it.Event = new(RoninValidatorSetEmergencyExitLockedAmountUpdated)
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
func (it *RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetEmergencyExitLockedAmountUpdated represents a EmergencyExitLockedAmountUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedAmountUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedAmountUpdated is a free log retrieval operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterEmergencyExitLockedAmountUpdated(opts *bind.FilterOpts) (*RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetEmergencyExitLockedAmountUpdatedIterator{contract: _RoninValidatorSet.contract, event: "EmergencyExitLockedAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedAmountUpdated is a free log subscription operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchEmergencyExitLockedAmountUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetEmergencyExitLockedAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetEmergencyExitLockedAmountUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
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

// ParseEmergencyExitLockedAmountUpdated is a log parse operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseEmergencyExitLockedAmountUpdated(log types.Log) (*RoninValidatorSetEmergencyExitLockedAmountUpdated, error) {
	event := new(RoninValidatorSetEmergencyExitLockedAmountUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetEmergencyExitLockedFundReleasedIterator is returned from FilterEmergencyExitLockedFundReleased and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleased events raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedFundReleasedIterator struct {
	Event *RoninValidatorSetEmergencyExitLockedFundReleased // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetEmergencyExitLockedFundReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetEmergencyExitLockedFundReleased)
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
		it.Event = new(RoninValidatorSetEmergencyExitLockedFundReleased)
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
func (it *RoninValidatorSetEmergencyExitLockedFundReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetEmergencyExitLockedFundReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetEmergencyExitLockedFundReleased represents a EmergencyExitLockedFundReleased event raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedFundReleased struct {
	ConsensusAddr  common.Address
	Recipient      common.Address
	UnlockedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleased is a free log retrieval operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterEmergencyExitLockedFundReleased(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetEmergencyExitLockedFundReleasedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetEmergencyExitLockedFundReleasedIterator{contract: _RoninValidatorSet.contract, event: "EmergencyExitLockedFundReleased", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleased is a free log subscription operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchEmergencyExitLockedFundReleased(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetEmergencyExitLockedFundReleased, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetEmergencyExitLockedFundReleased)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
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

// ParseEmergencyExitLockedFundReleased is a log parse operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseEmergencyExitLockedFundReleased(log types.Log) (*RoninValidatorSetEmergencyExitLockedFundReleased, error) {
	event := new(RoninValidatorSetEmergencyExitLockedFundReleased)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator is returned from FilterEmergencyExitLockedFundReleasingFailed and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleasingFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator struct {
	Event *RoninValidatorSetEmergencyExitLockedFundReleasingFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetEmergencyExitLockedFundReleasingFailed)
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
		it.Event = new(RoninValidatorSetEmergencyExitLockedFundReleasingFailed)
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
func (it *RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetEmergencyExitLockedFundReleasingFailed represents a EmergencyExitLockedFundReleasingFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitLockedFundReleasingFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	UnlockedAmount  *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleasingFailed is a free log retrieval operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterEmergencyExitLockedFundReleasingFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetEmergencyExitLockedFundReleasingFailedIterator{contract: _RoninValidatorSet.contract, event: "EmergencyExitLockedFundReleasingFailed", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleasingFailed is a free log subscription operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchEmergencyExitLockedFundReleasingFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetEmergencyExitLockedFundReleasingFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetEmergencyExitLockedFundReleasingFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
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

// ParseEmergencyExitLockedFundReleasingFailed is a log parse operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseEmergencyExitLockedFundReleasingFailed(log types.Log) (*RoninValidatorSetEmergencyExitLockedFundReleasingFailed, error) {
	event := new(RoninValidatorSetEmergencyExitLockedFundReleasingFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetEmergencyExitRequestedIterator is returned from FilterEmergencyExitRequested and is used to iterate over the raw logs and unpacked data for EmergencyExitRequested events raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitRequestedIterator struct {
	Event *RoninValidatorSetEmergencyExitRequested // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetEmergencyExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetEmergencyExitRequested)
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
		it.Event = new(RoninValidatorSetEmergencyExitRequested)
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
func (it *RoninValidatorSetEmergencyExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetEmergencyExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetEmergencyExitRequested represents a EmergencyExitRequested event raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExitRequested struct {
	ConsensusAddr common.Address
	LockedAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitRequested is a free log retrieval operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterEmergencyExitRequested(opts *bind.FilterOpts, consensusAddr []common.Address) (*RoninValidatorSetEmergencyExitRequestedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetEmergencyExitRequestedIterator{contract: _RoninValidatorSet.contract, event: "EmergencyExitRequested", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitRequested is a free log subscription operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchEmergencyExitRequested(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetEmergencyExitRequested, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetEmergencyExitRequested)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
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

// ParseEmergencyExitRequested is a log parse operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseEmergencyExitRequested(log types.Log) (*RoninValidatorSetEmergencyExitRequested, error) {
	event := new(RoninValidatorSetEmergencyExitRequested)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetEmergencyExpiryDurationUpdatedIterator is returned from FilterEmergencyExpiryDurationUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExpiryDurationUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExpiryDurationUpdatedIterator struct {
	Event *RoninValidatorSetEmergencyExpiryDurationUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetEmergencyExpiryDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetEmergencyExpiryDurationUpdated)
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
		it.Event = new(RoninValidatorSetEmergencyExpiryDurationUpdated)
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
func (it *RoninValidatorSetEmergencyExpiryDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetEmergencyExpiryDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetEmergencyExpiryDurationUpdated represents a EmergencyExpiryDurationUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetEmergencyExpiryDurationUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExpiryDurationUpdated is a free log retrieval operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterEmergencyExpiryDurationUpdated(opts *bind.FilterOpts) (*RoninValidatorSetEmergencyExpiryDurationUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetEmergencyExpiryDurationUpdatedIterator{contract: _RoninValidatorSet.contract, event: "EmergencyExpiryDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExpiryDurationUpdated is a free log subscription operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchEmergencyExpiryDurationUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetEmergencyExpiryDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetEmergencyExpiryDurationUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
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

// ParseEmergencyExpiryDurationUpdated is a log parse operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseEmergencyExpiryDurationUpdated(log types.Log) (*RoninValidatorSetEmergencyExpiryDurationUpdated, error) {
	event := new(RoninValidatorSetEmergencyExpiryDurationUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetFastFinalityRewardDistributedIterator is returned from FilterFastFinalityRewardDistributed and is used to iterate over the raw logs and unpacked data for FastFinalityRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetFastFinalityRewardDistributedIterator struct {
	Event *RoninValidatorSetFastFinalityRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetFastFinalityRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetFastFinalityRewardDistributed)
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
		it.Event = new(RoninValidatorSetFastFinalityRewardDistributed)
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
func (it *RoninValidatorSetFastFinalityRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetFastFinalityRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetFastFinalityRewardDistributed represents a FastFinalityRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetFastFinalityRewardDistributed struct {
	ConsensusAddr common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFastFinalityRewardDistributed is a free log retrieval operation binding the contract event 0x0c4d6a43fc8470dee97db74874b5685e412cc517d9bdecfde1623c5e835b18e4.
//
// Solidity: event FastFinalityRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterFastFinalityRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetFastFinalityRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "FastFinalityRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetFastFinalityRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "FastFinalityRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchFastFinalityRewardDistributed is a free log subscription operation binding the contract event 0x0c4d6a43fc8470dee97db74874b5685e412cc517d9bdecfde1623c5e835b18e4.
//
// Solidity: event FastFinalityRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchFastFinalityRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetFastFinalityRewardDistributed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "FastFinalityRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetFastFinalityRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "FastFinalityRewardDistributed", log); err != nil {
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

// ParseFastFinalityRewardDistributed is a log parse operation binding the contract event 0x0c4d6a43fc8470dee97db74874b5685e412cc517d9bdecfde1623c5e835b18e4.
//
// Solidity: event FastFinalityRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseFastFinalityRewardDistributed(log types.Log) (*RoninValidatorSetFastFinalityRewardDistributed, error) {
	event := new(RoninValidatorSetFastFinalityRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "FastFinalityRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetFastFinalityRewardDistributionFailedIterator is returned from FilterFastFinalityRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for FastFinalityRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetFastFinalityRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetFastFinalityRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetFastFinalityRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetFastFinalityRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetFastFinalityRewardDistributionFailed)
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
func (it *RoninValidatorSetFastFinalityRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetFastFinalityRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetFastFinalityRewardDistributionFailed represents a FastFinalityRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetFastFinalityRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFastFinalityRewardDistributionFailed is a free log retrieval operation binding the contract event 0x98697ee35f04a599a814432016fff3968c483d2d88dacb484926b9358f8e7cf9.
//
// Solidity: event FastFinalityRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterFastFinalityRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetFastFinalityRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "FastFinalityRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetFastFinalityRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "FastFinalityRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchFastFinalityRewardDistributionFailed is a free log subscription operation binding the contract event 0x98697ee35f04a599a814432016fff3968c483d2d88dacb484926b9358f8e7cf9.
//
// Solidity: event FastFinalityRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchFastFinalityRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetFastFinalityRewardDistributionFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "FastFinalityRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetFastFinalityRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "FastFinalityRewardDistributionFailed", log); err != nil {
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

// ParseFastFinalityRewardDistributionFailed is a log parse operation binding the contract event 0x98697ee35f04a599a814432016fff3968c483d2d88dacb484926b9358f8e7cf9.
//
// Solidity: event FastFinalityRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseFastFinalityRewardDistributionFailed(log types.Log) (*RoninValidatorSetFastFinalityRewardDistributionFailed, error) {
	event := new(RoninValidatorSetFastFinalityRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "FastFinalityRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RoninValidatorSet contract.
type RoninValidatorSetInitializedIterator struct {
	Event *RoninValidatorSetInitialized // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetInitialized)
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
		it.Event = new(RoninValidatorSetInitialized)
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
func (it *RoninValidatorSetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetInitialized represents a Initialized event raised by the RoninValidatorSet contract.
type RoninValidatorSetInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoninValidatorSetInitializedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetInitializedIterator{contract: _RoninValidatorSet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetInitialized) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetInitialized)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseInitialized(log types.Log) (*RoninValidatorSetInitialized, error) {
	event := new(RoninValidatorSetInitialized)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator is returned from FilterMaxPrioritizedValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxPrioritizedValidatorNumberUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator struct {
	Event *RoninValidatorSetMaxPrioritizedValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
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
		it.Event = new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
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
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxPrioritizedValidatorNumberUpdated represents a MaxPrioritizedValidatorNumberUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxPrioritizedValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxPrioritizedValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxPrioritizedValidatorNumberUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxPrioritizedValidatorNumberUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxPrioritizedValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPrioritizedValidatorNumberUpdated is a free log subscription operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxPrioritizedValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxPrioritizedValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
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

// ParseMaxPrioritizedValidatorNumberUpdated is a log parse operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxPrioritizedValidatorNumberUpdated(log types.Log) (*RoninValidatorSetMaxPrioritizedValidatorNumberUpdated, error) {
	event := new(RoninValidatorSetMaxPrioritizedValidatorNumberUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxValidatorCandidateUpdatedIterator is returned from FilterMaxValidatorCandidateUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorCandidateUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorCandidateUpdatedIterator struct {
	Event *RoninValidatorSetMaxValidatorCandidateUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxValidatorCandidateUpdated)
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
		it.Event = new(RoninValidatorSetMaxValidatorCandidateUpdated)
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
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxValidatorCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxValidatorCandidateUpdated represents a MaxValidatorCandidateUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorCandidateUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorCandidateUpdated is a free log retrieval operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxValidatorCandidateUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxValidatorCandidateUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxValidatorCandidateUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxValidatorCandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorCandidateUpdated is a free log subscription operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxValidatorCandidateUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxValidatorCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxValidatorCandidateUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
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

// ParseMaxValidatorCandidateUpdated is a log parse operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxValidatorCandidateUpdated(log types.Log) (*RoninValidatorSetMaxValidatorCandidateUpdated, error) {
	event := new(RoninValidatorSetMaxValidatorCandidateUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMaxValidatorNumberUpdatedIterator is returned from FilterMaxValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorNumberUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorNumberUpdatedIterator struct {
	Event *RoninValidatorSetMaxValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMaxValidatorNumberUpdated)
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
		it.Event = new(RoninValidatorSetMaxValidatorNumberUpdated)
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
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMaxValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMaxValidatorNumberUpdated represents a MaxValidatorNumberUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMaxValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMaxValidatorNumberUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMaxValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMaxValidatorNumberUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MaxValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorNumberUpdated is a free log subscription operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMaxValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMaxValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMaxValidatorNumberUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
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

// ParseMaxValidatorNumberUpdated is a log parse operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMaxValidatorNumberUpdated(log types.Log) (*RoninValidatorSetMaxValidatorNumberUpdated, error) {
	event := new(RoninValidatorSetMaxValidatorNumberUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator is returned from FilterMinEffectiveDaysOnwardsUpdated and is used to iterate over the raw logs and unpacked data for MinEffectiveDaysOnwardsUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator struct {
	Event *RoninValidatorSetMinEffectiveDaysOnwardsUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMinEffectiveDaysOnwardsUpdated)
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
		it.Event = new(RoninValidatorSetMinEffectiveDaysOnwardsUpdated)
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
func (it *RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMinEffectiveDaysOnwardsUpdated represents a MinEffectiveDaysOnwardsUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetMinEffectiveDaysOnwardsUpdated struct {
	NumOfDays *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinEffectiveDaysOnwardsUpdated is a free log retrieval operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMinEffectiveDaysOnwardsUpdated(opts *bind.FilterOpts) (*RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMinEffectiveDaysOnwardsUpdatedIterator{contract: _RoninValidatorSet.contract, event: "MinEffectiveDaysOnwardsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinEffectiveDaysOnwardsUpdated is a free log subscription operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMinEffectiveDaysOnwardsUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMinEffectiveDaysOnwardsUpdated) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMinEffectiveDaysOnwardsUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
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

// ParseMinEffectiveDaysOnwardsUpdated is a log parse operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMinEffectiveDaysOnwardsUpdated(log types.Log) (*RoninValidatorSetMinEffectiveDaysOnwardsUpdated, error) {
	event := new(RoninValidatorSetMinEffectiveDaysOnwardsUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMiningRewardDistributedIterator is returned from FilterMiningRewardDistributed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributedIterator struct {
	Event *RoninValidatorSetMiningRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMiningRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMiningRewardDistributed)
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
		it.Event = new(RoninValidatorSetMiningRewardDistributed)
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
func (it *RoninValidatorSetMiningRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMiningRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMiningRewardDistributed represents a MiningRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributed struct {
	ConsensusAddr common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributed is a free log retrieval operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMiningRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetMiningRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMiningRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "MiningRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributed is a free log subscription operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMiningRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMiningRewardDistributed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMiningRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
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

// ParseMiningRewardDistributed is a log parse operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMiningRewardDistributed(log types.Log) (*RoninValidatorSetMiningRewardDistributed, error) {
	event := new(RoninValidatorSetMiningRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetMiningRewardDistributionFailedIterator is returned from FilterMiningRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetMiningRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetMiningRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetMiningRewardDistributionFailed)
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
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetMiningRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetMiningRewardDistributionFailed represents a MiningRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetMiningRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributionFailed is a free log retrieval operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterMiningRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*RoninValidatorSetMiningRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetMiningRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "MiningRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributionFailed is a free log subscription operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchMiningRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetMiningRewardDistributionFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetMiningRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
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

// ParseMiningRewardDistributionFailed is a log parse operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseMiningRewardDistributionFailed(log types.Log) (*RoninValidatorSetMiningRewardDistributionFailed, error) {
	event := new(RoninValidatorSetMiningRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetStakingRewardDistributedIterator is returned from FilterStakingRewardDistributed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributed events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributedIterator struct {
	Event *RoninValidatorSetStakingRewardDistributed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingRewardDistributed)
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
		it.Event = new(RoninValidatorSetStakingRewardDistributed)
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
func (it *RoninValidatorSetStakingRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingRewardDistributed represents a StakingRewardDistributed event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributed struct {
	TotalAmount    *big.Int
	ConsensusAddrs []common.Address
	Amounts        []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributed is a free log retrieval operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingRewardDistributed(opts *bind.FilterOpts) (*RoninValidatorSetStakingRewardDistributedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingRewardDistributedIterator{contract: _RoninValidatorSet.contract, event: "StakingRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributed is a free log subscription operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingRewardDistributed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingRewardDistributed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
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

// ParseStakingRewardDistributed is a log parse operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingRewardDistributed(log types.Log) (*RoninValidatorSetStakingRewardDistributed, error) {
	event := new(RoninValidatorSetStakingRewardDistributed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetStakingRewardDistributionFailedIterator is returned from FilterStakingRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributionFailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributionFailedIterator struct {
	Event *RoninValidatorSetStakingRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetStakingRewardDistributionFailed)
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
		it.Event = new(RoninValidatorSetStakingRewardDistributionFailed)
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
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetStakingRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetStakingRewardDistributionFailed represents a StakingRewardDistributionFailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetStakingRewardDistributionFailed struct {
	TotalAmount     *big.Int
	ConsensusAddrs  []common.Address
	Amounts         []*big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributionFailed is a free log retrieval operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterStakingRewardDistributionFailed(opts *bind.FilterOpts) (*RoninValidatorSetStakingRewardDistributionFailedIterator, error) {

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetStakingRewardDistributionFailedIterator{contract: _RoninValidatorSet.contract, event: "StakingRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributionFailed is a free log subscription operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchStakingRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetStakingRewardDistributionFailed) (event.Subscription, error) {

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetStakingRewardDistributionFailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
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

// ParseStakingRewardDistributionFailed is a log parse operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseStakingRewardDistributionFailed(log types.Log) (*RoninValidatorSetStakingRewardDistributionFailed, error) {
	event := new(RoninValidatorSetStakingRewardDistributionFailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorPunishedIterator is returned from FilterValidatorPunished and is used to iterate over the raw logs and unpacked data for ValidatorPunished events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorPunishedIterator struct {
	Event *RoninValidatorSetValidatorPunished // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorPunished)
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
		it.Event = new(RoninValidatorSetValidatorPunished)
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
func (it *RoninValidatorSetValidatorPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorPunished represents a ValidatorPunished event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorPunished struct {
	Cid                            common.Address
	Period                         *big.Int
	JailedUntil                    *big.Int
	DeductedStakingAmount          *big.Int
	BlockProducerRewardDeprecated  bool
	BridgeOperatorRewardDeprecated bool
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed cid, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorPunished(opts *bind.FilterOpts, cid []common.Address, period []*big.Int) (*RoninValidatorSetValidatorPunishedIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorPunished", cidRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorPunishedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed cid, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorPunished, cid []common.Address, period []*big.Int) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorPunished", cidRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorPunished)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
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

// ParseValidatorPunished is a log parse operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed cid, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorPunished(log types.Log) (*RoninValidatorSetValidatorPunished, error) {
	event := new(RoninValidatorSetValidatorPunished)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorSetUpdatedIterator struct {
	Event *RoninValidatorSetValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorSetUpdated)
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
		it.Event = new(RoninValidatorSetValidatorSetUpdated)
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
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorSetUpdated represents a ValidatorSetUpdated event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorSetUpdated struct {
	Period         *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*RoninValidatorSetValidatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorSetUpdatedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorSetUpdated)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorSetUpdated(log types.Log) (*RoninValidatorSetValidatorSetUpdated, error) {
	event := new(RoninValidatorSetValidatorSetUpdated)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorUnjailedIterator struct {
	Event *RoninValidatorSetValidatorUnjailed // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetValidatorUnjailed)
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
		it.Event = new(RoninValidatorSetValidatorUnjailed)
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
func (it *RoninValidatorSetValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetValidatorUnjailed represents a ValidatorUnjailed event raised by the RoninValidatorSet contract.
type RoninValidatorSetValidatorUnjailed struct {
	Cid    common.Address
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed cid, uint256 period)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts, cid []common.Address) (*RoninValidatorSetValidatorUnjailedIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "ValidatorUnjailed", cidRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetValidatorUnjailedIterator{contract: _RoninValidatorSet.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed cid, uint256 period)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetValidatorUnjailed, cid []common.Address) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "ValidatorUnjailed", cidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetValidatorUnjailed)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
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

// ParseValidatorUnjailed is a log parse operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed cid, uint256 period)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseValidatorUnjailed(log types.Log) (*RoninValidatorSetValidatorUnjailed, error) {
	event := new(RoninValidatorSetValidatorUnjailed)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoninValidatorSetWrappedUpEpochIterator is returned from FilterWrappedUpEpoch and is used to iterate over the raw logs and unpacked data for WrappedUpEpoch events raised by the RoninValidatorSet contract.
type RoninValidatorSetWrappedUpEpochIterator struct {
	Event *RoninValidatorSetWrappedUpEpoch // Event containing the contract specifics and raw log

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
func (it *RoninValidatorSetWrappedUpEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoninValidatorSetWrappedUpEpoch)
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
		it.Event = new(RoninValidatorSetWrappedUpEpoch)
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
func (it *RoninValidatorSetWrappedUpEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoninValidatorSetWrappedUpEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoninValidatorSetWrappedUpEpoch represents a WrappedUpEpoch event raised by the RoninValidatorSet contract.
type RoninValidatorSetWrappedUpEpoch struct {
	PeriodNumber *big.Int
	EpochNumber  *big.Int
	PeriodEnding bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWrappedUpEpoch is a free log retrieval operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) FilterWrappedUpEpoch(opts *bind.FilterOpts, periodNumber []*big.Int, epochNumber []*big.Int) (*RoninValidatorSetWrappedUpEpochIterator, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.FilterLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return &RoninValidatorSetWrappedUpEpochIterator{contract: _RoninValidatorSet.contract, event: "WrappedUpEpoch", logs: logs, sub: sub}, nil
}

// WatchWrappedUpEpoch is a free log subscription operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) WatchWrappedUpEpoch(opts *bind.WatchOpts, sink chan<- *RoninValidatorSetWrappedUpEpoch, periodNumber []*big.Int, epochNumber []*big.Int) (event.Subscription, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _RoninValidatorSet.contract.WatchLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoninValidatorSetWrappedUpEpoch)
				if err := _RoninValidatorSet.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
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

// ParseWrappedUpEpoch is a log parse operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_RoninValidatorSet *RoninValidatorSetFilterer) ParseWrappedUpEpoch(log types.Log) (*RoninValidatorSetWrappedUpEpoch, error) {
	event := new(RoninValidatorSetWrappedUpEpoch)
	if err := _RoninValidatorSet.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
