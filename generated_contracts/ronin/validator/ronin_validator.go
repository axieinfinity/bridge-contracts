// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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
	Admin              common.Address
	ConsensusAddr      common.Address
	TreasuryAddr       common.Address
	BridgeOperatorAddr common.Address
	CommissionRate     *big.Int
	RevokingTimestamp  *big.Int
	TopupDeadline      *big.Int
}

// ICommonInfoEmergencyExitInfo is an auto generated low-level Go binding around an user-defined struct.
type ICommonInfoEmergencyExitInfo struct {
	LockedAmount *big.Int
	RecyclingAt  *big.Int
}

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedEmergencyExit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedRevokingCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedUpdatingCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAlreadyWrappedEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrAtEndOfEpochOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallPrecompiled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeBridgeTrackingContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeCoinbase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeMaintenanceContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeRoninTrustedOrgContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeSlashIndicatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeStakingContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeStakingVestingContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExceedsMaxNumberOfCandidate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"}],\"name\":\"ErrExistentBridgeOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrExistentCandidate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidateAdminAddr\",\"type\":\"address\"}],\"name\":\"ErrExistentCandidateAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"ErrExistentTreasury\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidEffectiveDaysOnwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidMinEffectiveDaysOnwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNonExistentCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRecipientRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxPrioitizedValidatorNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedReceiveRON\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"BlockProducerSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumICoinbaseExecution.BlockRewardDeprecatedType\",\"name\":\"deprecatedType\",\"type\":\"uint8\"}],\"name\":\"BlockRewardDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submittedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeTrackingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"CandidateGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"}],\"name\":\"CandidateRevokingTimestampUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"name\":\"CandidateTopupDeadlineUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"CandidatesRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateUpdateScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"CommissionRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"DeprecatedRewardRecycleFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DeprecatedRewardRecycled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedFundReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitLockedFundReleasingFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyExpiryDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numOfDays\",\"type\":\"uint256\"}],\"name\":\"MinEffectiveDaysOnwardsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SlashIndicatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"StakingRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingVestingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailedUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blockProducerRewardDeprecated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"bridgeOperatorRewardDeprecated\",\"type\":\"bool\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnjailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"periodEnding\",\"type\":\"bool\"}],\"name\":\"WrappedUpEpoch\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"_isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTrackingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"checkJailedAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"checkManyJailed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"}],\"name\":\"checkMiningRewardDeprecated\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"checkMiningRewardDeprecatedAtPeriod\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodStartAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExitLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"execBailOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_secLeftToRevoke\",\"type\":\"uint256\"}],\"name\":\"execEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"execReleaseLockedFundForEmergencyExitRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_effectiveDaysOnwards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"execRequestUpdateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newJailedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"execSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockProducers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCommissionChangeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"internalType\":\"structICandidateManager.CommissionSchedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"getEmergencyExitInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recyclingAt\",\"type\":\"uint256\"}],\"internalType\":\"structICommonInfo.EmergencyExitInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getJailedTimeLeft\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"getJailedTimeLeftAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidateAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"grantValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingVestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__bridgeTrackingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__minEffectiveDaysOnwards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"__emergencyExitConfigs\",\"type\":\"uint256[2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBlockProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"}],\"name\":\"isBridgeOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"isOperatingBridge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minEffectiveDaysOnwards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompilePickValidatorSetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileSortValidatorsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_secsLeft\",\"type\":\"uint256\"}],\"name\":\"requestRevokeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeTrackingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emergencyExitLockedAmount\",\"type\":\"uint256\"}],\"name\":\"setEmergencyExitLockedAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emergencyExpiryDuration\",\"type\":\"uint256\"}],\"name\":\"setEmergencyExpiryDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxPrioritizedValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numOfDays\",\"type\":\"uint256\"}],\"name\":\"setMinEffectiveDaysOnwards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSlashIndicatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingVestingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashIndicatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingVestingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlockProducers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBridgeOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDeprecatedReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"tryGetPeriodOfEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000e4565b600054610100900460ff16156200008f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161015620000e2576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b615b5c80620000f46000396000f3fe6080604052600436106104095760003560e01c8063690b753611610213578063a66c0f7711610123578063cba44de9116100ab578063edb194bb1161007a578063edb194bb14610be4578063ee99205c14610c46578063eeb629a814610c64578063f2811bcc14610c79578063facd743b14610c9957610418565b8063cba44de914610b7c578063d09f1ab414610b91578063d2cb215e14610ba6578063e5125a1d14610bc457610418565b8063b5e337de116100f2578063b5e337de14610af2578063b7ab4db514610b12578063ba77b06c14610b27578063c3c8b5d614610b3c578063c94aaa0214610b5c57610418565b8063a66c0f7714610a7d578063a7c2f11914610a92578063ad29578314610ab2578063b405aaf214610ad257610418565b80638d559c38116101a65780639c8d98da116101755780639c8d98da146109d05780639dd373b9146109f05780639e94b9ec14610a10578063a0c3f2d214610a25578063a3d545f514610a5d57610418565b80638d559c381461096557806396585fc2146109795780639b19dbfd146109995780639b8c334b146109ae57610418565b8063733ec970116101e2578063733ec970146108f05780637593ff7114610910578063823a7b9c1461093057806387c891bd1461095057610418565b8063690b75361461089e5780636aa1c2ef146108b35780636efa12bd146108c857806372e46810146108e857610418565b80633b3159b61161031957806352091f17116102a1578063570ccb1311610270578063570ccb131461080b5780635a08482d1461082b578063605239a11461084957806365244ece1461085e5780636611f8431461087e57610418565b806352091f17146107ae5780635248184a146107b65780635511cde1146107d8578063562d5304146107f657610418565b806349096d26116102e857806349096d26146107175780634d8df063146107395780634de2b735146107595780634ee4d72b146107795780634f2a693f1461078e57610418565b80633b3159b61461068e5780634493421e146106a2578063468c96ae146106c057806346fe9311146106f757610418565b8063217f35c21161039c578063297a8fca1161036b578063297a8fca146105da5780632bcf3d15146105ef5780632d784a981461060f5780633529214b1461063c578063367ec12b1461066e57610418565b8063217f35c21461055857806323c65eb01461056d57806328bde1e11461058d5780632924de71146105ba57610418565b80631196ab66116103d85780631196ab66146104cb57806315b5ebde146104eb5780631b6e0a991461050b5780631f6288011461053857610418565b806304d971ab1461042057806306040618146104555780630f43a6771461047857806311662dc21461048e57610418565b3661041857610416610cb9565b005b610416610cb9565b34801561042c57600080fd5b5061044061043b36600461515b565b610cff565b60405190151581526020015b60405180910390f35b34801561046157600080fd5b5061046a610d26565b60405190815260200161044c565b34801561048457600080fd5b5061046a60aa5481565b34801561049a57600080fd5b506104ae6104a9366004615194565b610d36565b60408051931515845260208401929092529082015260600161044c565b3480156104d757600080fd5b506104166104e63660046151c0565b610db9565b3480156104f757600080fd5b50610416610506366004615194565b610e06565b34801561051757600080fd5b5061052b610526366004615224565b610ef2565b60405161044c919061526f565b34801561054457600080fd5b506104406105533660046152b5565b610faf565b34801561056457600080fd5b50610440610fe9565b34801561057957600080fd5b50610440610588366004615194565b610ff7565b34801561059957600080fd5b506105ad6105a83660046152b5565b61101b565b60405161044c9190615326565b3480156105c657600080fd5b506104406105d53660046152b5565b6110fb565b3480156105e657600080fd5b5060045461046a565b3480156105fb57600080fd5b5061041661060a3660046152b5565b611107565b34801561061b57600080fd5b5061062f61062a3660046152b5565b611173565b60405161044c9190615334565b34801561064857600080fd5b50606d546001600160a01b03165b6040516001600160a01b03909116815260200161044c565b34801561067a57600080fd5b5061041661068936600461535c565b611203565b34801561069a57600080fd5b506068610656565b3480156106ae57600080fd5b50606e546001600160a01b0316610656565b3480156106cc57600080fd5b506106e06106db3660046151c0565b61138e565b60408051921515835260208301919091520161044c565b34801561070357600080fd5b506104166107123660046152b5565b6113cc565b34801561072357600080fd5b5061072c611438565b60405161044c9190615466565b34801561074557600080fd5b506104166107543660046151c0565b611521565b34801561076557600080fd5b5061052b610774366004615479565b611562565b34801561078557600080fd5b5060e45461046a565b34801561079a57600080fd5b506104166107a93660046151c0565b61161d565b61041661165e565b3480156107c257600080fd5b506107cb6119f7565b60405161044c91906154ba565b3480156107e457600080fd5b5060a8546001600160a01b0316610656565b34801561080257600080fd5b5061046a611b55565b34801561081757600080fd5b506104166108263660046154fc565b611ba9565b34801561083757600080fd5b506070546001600160a01b0316610656565b34801561085557600080fd5b5060725461046a565b34801561086a57600080fd5b506104406108793660046152b5565b611da9565b34801561088a57600080fd5b506104166108993660046151c0565b611ddd565b3480156108aa57600080fd5b5060e55461046a565b3480156108bf57600080fd5b5060015461046a565b3480156108d457600080fd5b506104166108e3366004615194565b611e1e565b610416611eab565b3480156108fc57600080fd5b5061041661090b366004615531565b612096565b34801561091c57600080fd5b5061044061092b3660046151c0565b612345565b34801561093c57600080fd5b5061041661094b3660046151c0565b612369565b34801561095c57600080fd5b5060025461046a565b34801561097157600080fd5b506066610656565b34801561098557600080fd5b506104ae6109943660046152b5565b6123aa565b3480156109a557600080fd5b5061072c6123c6565b3480156109ba57600080fd5b506104406109c93660046151c0565b6003541090565b3480156109dc57600080fd5b506104166109eb3660046152b5565b6124b2565b3480156109fc57600080fd5b50610416610a0b3660046152b5565b61251e565b348015610a1c57600080fd5b5061046a61258a565b348015610a3157600080fd5b50610440610a403660046152b5565b6001600160a01b0316600090815260746020526040902054151590565b348015610a6957600080fd5b5061046a610a783660046151c0565b6125de565b348015610a8957600080fd5b5060e65461046a565b348015610a9e57600080fd5b50610416610aad366004615194565b6125e9565b348015610abe57600080fd5b50610416610acd3660046152b5565b612865565b348015610ade57600080fd5b50610440610aed3660046152b5565b6128d1565b348015610afe57600080fd5b50610416610b0d3660046152b5565b61295c565b348015610b1e57600080fd5b5061072c6129c8565b348015610b3357600080fd5b5061072c612a75565b348015610b4857600080fd5b50610416610b5736600461515b565b612ad7565b348015610b6857600080fd5b50610416610b773660046151c0565b612d69565b348015610b8857600080fd5b5060765461046a565b348015610b9d57600080fd5b5060a95461046a565b348015610bb257600080fd5b50606f546001600160a01b0316610656565b348015610bd057600080fd5b50610416610bdf3660046154fc565b612daa565b348015610bf057600080fd5b5061062f610bff3660046152b5565b6040805180820190915260008082526020820152506001600160a01b0316600090815260776020908152604091829020825180840190935280548352600101549082015290565b348015610c5257600080fd5b506071546001600160a01b0316610656565b348015610c7057600080fd5b5060ad5461046a565b348015610c8557600080fd5b5061052b610c94366004615479565b612eed565b348015610ca557600080fd5b50610440610cb43660046152b5565b612fb6565b606d546001600160a01b03163314801590610cdf57506071546001600160a01b03163314155b15610cfd57604051635ff7650760e01b815260040160405180910390fd5b565b6001600160a01b038281166000908152607560205260409020548116908216145b92915050565b6000610d3160035490565b905090565b6001600160a01b0382166000908152603a60205260408120548190819084811015610d6c57600080600093509350935050610db2565b60019350610d7a85826155ab565b610d859060016155be565b9250610d90856125de565b610d99826125de565b610da391906155ab565b610dae9060016155be565b9150505b9250925092565b610dc1612ff3565b6001600160a01b0316336001600160a01b031614610dfa5760405162461bcd60e51b8152600401610df1906155d1565b60405180910390fd5b610e0381613021565b50565b33610e196070546001600160a01b031690565b6001600160a01b031614610e40576040516328b9c24b60e21b815260040160405180910390fd5b6001600160a01b038216600081815260386020908152604080832085845282528083208054600160ff199182168117909255948452603783528184208685529092529091208054909216909155610e9790436155ab565b6001600160a01b0383166000818152603a6020526040908190209290925590517f6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e90610ee69084815260200190565b60405180910390a25050565b6060826001600160401b03811115610f0c57610f0c615613565b604051908082528060200260200182016040528015610f35578160200160208202803683370190505b50905060005b83811015610fa757610f73858583818110610f5857610f58615629565b9050602002016020810190610f6d91906152b5565b8461307f565b828281518110610f8557610f85615629565b9115156020928302919091019091015280610f9f8161563f565b915050610f3b565b509392505050565b6001600160a01b038116600090815260ac6020526040812054610d209060029060ff166003811115610fe357610fe3615658565b906130aa565b6000610d316109c9426130dd565b6001600160a01b0382166000908152603a60205260408120548211155b9392505050565b6040805160e08101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905260c083018290526001600160a01b03851682526074905291909120546110865760405163a64b34ad60e01b815260040160405180910390fd5b506001600160a01b03908116600090815260756020908152604091829020825160e081018452815485168152600182015485169281019290925260028101548416928201929092526003820154909216606083015260048101546080830152600581015460a08301526006015460c082015290565b6000610d208243610ff7565b61110f612ff3565b6001600160a01b0316336001600160a01b03161461113f5760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b60000361116a57604051637bcd509160e01b815260040160405180910390fd5b610e03816130ec565b604080518082018252600080825260209182018190526001600160a01b038416815260e88252829020825180840190935280548352600101549082018190526111fe5760405162461bcd60e51b815260206004820181905260248201527f436f6d6d6f6e53746f726167653a206e6f6e2d6578697374656e7420696e666f6044820152606401610df1565b919050565b600054610100900460ff16158080156112235750600054600160ff909116105b8061123d5750303b15801561123d575060005460ff166001145b6112a05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610df1565b6000805460ff1916600117905580156112c3576000805461ff0019166101001790555b6112cc8d6130ec565b6112d58c61313a565b6112de8b613188565b6112e78a6131d6565b6112f088613224565b6112f989613272565b611302876132c0565b61130b866132f5565b6113148561332a565b61131d84613021565b6113278235613382565b61133460208301356133b7565b6001839055801561137f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050505050565b60008061139a436125de565b831115806113b5575060008381526005602052604090205415155b600093845260056020526040909320549293915050565b6113d4612ff3565b6001600160a01b0316336001600160a01b0316146114045760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b60000361142f57604051637bcd509160e01b815260040160405180910390fd5b610e03816131d6565b606060aa546001600160401b0381111561145457611454615613565b60405190808252806020026020018201604052801561147d578160200160208202803683370190505b5090506000805b825181101561151b57600081815260ab60205260409020546114ae906001600160a01b0316611da9565b1561150957600081815260ab60205260409020546001600160a01b031683836114d68161563f565b9450815181106114e8576114e8615629565b60200260200101906001600160a01b031690816001600160a01b0316815250505b806115138161563f565b915050611484565b50815290565b611529612ff3565b6001600160a01b0316336001600160a01b0316146115595760405162461bcd60e51b8152600401610df1906155d1565b610e03816133b7565b6060816001600160401b0381111561157c5761157c615613565b6040519080825280602002602001820160405280156115a5578160200160208202803683370190505b50905060005b82811015611616576115e28484838181106115c8576115c8615629565b90506020020160208101906115dd91906152b5565b6133ec565b8282815181106115f4576115f4615629565b911515602092830291909101909101528061160e8161563f565b9150506115ab565b5092915050565b611625612ff3565b6001600160a01b0316336001600160a01b0316146116555760405162461bcd60e51b8152600401610df1906155d1565b610e03816132f5565b33411461167e576040516309f358fd60e01b815260040160405180910390fd5b3433600061168b82611da9565b801561169d575061169b826133ec565b155b80156116b757506116b5826116b0610d26565b61307f565b155b606d54604051630634f5b960e01b8152821515600482015260016024820181905292935060009182916001600160a01b0390911690630634f5b9906044016060604051808303816000875af1158015611714573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611738919061567e565b92509250508060e2600082825461174f91906155be565b909155508490506117bf578560e4600082825461176c91906155be565b92505081905550846001600160a01b03167f4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e8760016040516117af9291906156b3565b60405180910390a2505050505050565b60408051878152602081018490526001600160a01b038716917f0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1910160405180910390a2600061180d610d26565b9050600061181b84896155be565b6001600160a01b03881660009081526038602090815260408083208684529091528120549192509060ff1615611938576070546040805163631c8fd160e11b815290516000926001600160a01b03169163c6391fa29160048083019260809291908290030181865afa158015611895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118b991906156e5565b935050505061271081846118cd919061571b565b6118d79190615748565b91508160e460008282546118eb91906155be565b92505081905550886001600160a01b03167f4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e83600260405161192e9291906156b3565b60405180910390a2505b61194281836155ab565b6001600160a01b03891660009081526075602052604081206004015491935061271061196e858461571b565b6119789190615748565b6001600160a01b038b16600090815260e060205260408120805492935083929091906119a59084906155be565b90915550600090506119b782866155ab565b6001600160a01b038c16600090815260e160205260408120805492935083929091906119e49084906155be565b9091555050505050505050505050505050565b6073546060906001600160401b03811115611a1457611a14615613565b604051908082528060200260200182016040528015611a7b57816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c08201528252600019909201910181611a325790505b50905060005b8151811015611b51576075600060738381548110611aa157611aa1615629565b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160e081018452815485168152600182015485169281019290925260028101548416928201929092526003820154909216606083015260048101546080830152600581015460a08301526006015460c08201528251839083908110611b3357611b33615629565b60200260200101819052508080611b499061563f565b915050611a81565b5090565b6000805b60aa54811015611b5157600081815260ab6020526040902054611b84906001600160a01b0316610faf565b15611b975781611b938161563f565b9250505b80611ba18161563f565b915050611b59565b33611bbc6070546001600160a01b031690565b6001600160a01b031614611be3576040516328b9c24b60e21b815260040160405180910390fd5b6000611bed610d26565b6001600160a01b03851660008181526037602090815260408083208584528252808320805460ff1916600117905592825260e181528282205460e090915291902054919250611c3b916155be565b60e46000828254611c4c91906155be565b90915550506001600160a01b038416600090815260e06020908152604080832083905560e18252808320839055603a909152902054831115611ca4576001600160a01b0384166000908152603a602052604090208390555b8115611d3d5760715460405163138ac02f60e11b81526001600160a01b038681166004830152602482018590526000921690632715805e906044016020604051808303816000875af1158015611cfe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d22919061575c565b90508060e46000828254611d3691906155be565b9091555050505b6001600160a01b0384166000818152603a6020908152604080832054815190815291820186905260019082015260608101919091528291907f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a906080015b60405180910390a350505050565b6001600160a01b038116600090815260ac6020526040812054610d209060019060ff166003811115610fe357610fe3615658565b611de5612ff3565b6001600160a01b0316336001600160a01b031614611e155760405162461bcd60e51b8152600401610df1906155d1565b610e0381613382565b33611e316071546001600160a01b031690565b6001600160a01b031614611e5857604051638aaf4a0760e01b815260040160405180910390fd5b6001600160a01b0382166000908152607560205260409020600581015415611e935760405163fab9167360e01b815260040160405180910390fd5b611ea681611ea184426155be565b61340d565b505050565b334114611ecb576040516309f358fd60e01b815260040160405180910390fd5b611ed443612345565b611ef157604051636c74eecf60e01b815260040160405180910390fd5b611efa436125de565b611f056002546125de565b10611f2357604051632458f64160e01b815260040160405180910390fd5b436002556000611f32426130dd565b90506000611f41826003541090565b9050611f4e4360016155be565b6004556000611f5b6129c8565b90506000611f68436125de565b90506000611f778260016155be565b90506000611f83610d26565b9050841561203457611f95818561348f565b600080611fa283876136dc565b91509150611fb283878484613919565b611fba613a22565b611fc2613b7d565b607054604051631da0214360e21b81526001600160a01b0390911690637680850c90611ff49089908790600401615775565b600060405180830381600087803b15801561200e57600080fd5b505af1158015612022573d6000803e3d6000fd5b5050505061202f88613cbb565b955050505b61203f868386613e3f565b82817f0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce787604051612074911515815260200190565b60405180910390a3506000908152600560205260409020849055505050600355565b336120a96071546001600160a01b031690565b6001600160a01b0316146120d057604051638aaf4a0760e01b815260040160405180910390fd5b60735460725481106120f557604051638616841b60e01b815260040160405180910390fd5b6001600160a01b0385166000908152607460205260409020541561212c57604051638ad9cdf960e01b815260040160405180910390fd5b61271082111561214f57604051631b8454a360e21b815260040160405180910390fd5b60005b60735481101561225c576000607560006073848154811061217557612175615629565b60009182526020808320909101546001600160a01b039081168452908301939093526040909101902080549092508116908916036121d15760405163fc3d8c7560e01b81526001600160a01b0389166004820152602401610df1565b60028101546001600160a01b039081169087160361220d57604051632d33a7e760e11b81526001600160a01b0387166004820152602401610df1565b60038101546001600160a01b0390811690861603612249576040516350e1263b60e01b81526001600160a01b0386166004820152602401610df1565b50806122548161563f565b915050612152565b506001600160a01b038581166000818152607460209081526040808320861990556073805460018082019092557ff79bde9ddd17963ebce6f7d021d60de7c2bd0db944d23c900c0c0e775f5300520180546001600160a01b031990811687179091556075845293829020805485168d88169081178255918101805486168717905560028101805486168c8916908117909155600382018054909616978b1697881790955560048101899055915195865290949093917fd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd910160405180910390a450505050505050565b60006001805461235591906155ab565b6001546123629084615797565b1492915050565b612371612ff3565b6001600160a01b0316336001600160a01b0316146123a15760405162461bcd60e51b8152600401610df1906155d1565b610e03816132c0565b60008060006123b98443610d36565b9250925092509193909250565b606060aa546001600160401b038111156123e2576123e2615613565b60405190808252806020026020018201604052801561240b578160200160208202803683370190505b5090506000805b825181101561151b57600081815260ab602052604090205461243c906001600160a01b0316611da9565b156124a057600081815260ab6020526040902054612462906001600160a01b03166141f7565b838361246d8161563f565b94508151811061247f5761247f615629565b60200260200101906001600160a01b031690816001600160a01b0316815250505b806124aa8161563f565b915050612412565b6124ba612ff3565b6001600160a01b0316336001600160a01b0316146124ea5760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b60000361251557604051637bcd509160e01b815260040160405180910390fd5b610e0381613224565b612526612ff3565b6001600160a01b0316336001600160a01b0316146125565760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b60000361258157604051637bcd509160e01b815260040160405180910390fd5b610e038161313a565b6000805b60aa54811015611b5157600081815260ab60205260409020546125b9906001600160a01b0316611da9565b156125cc57816125c88161563f565b9250505b806125d68161563f565b91505061258e565b6000610d2082614202565b336125fc6071546001600160a01b031690565b6001600160a01b03161461262357604051638aaf4a0760e01b815260040160405180910390fd5b6001600160a01b038216600090815260e86020526040902060018101541561265e5760405163057aab3160e31b815260040160405180910390fd5b600061266a83426155be565b6001600160a01b038516600090815260756020526040902090915061268f908261340d565b6001600160a01b038481166000818152603b602052604080822085905560715460e554915163138ac02f60e11b81526004810194909452602484019190915290921690632715805e906044016020604051808303816000875af11580156126fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061271e919061575c565b9050801561281b57600060e6544261273691906155be565b60e78054600180820183556000929092527f6cb0db1d7354dfb4a1464318006df0643cafe2002a86a29ff8560f900fef28a10180546001600160a01b0319166001600160a01b038a1617905583865585018190559050612794612ff3565b6001600160a01b0387811660008181526075602052604090819020600201549051630a2fae5760e41b81526004810192909252821660248201524260448201526064810184905291169063a2fae57090608401600060405180830381600087803b15801561280157600080fd5b505af1158015612815573d6000803e3d6000fd5b50505050505b846001600160a01b03167f77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c48260405161285691815260200190565b60405180910390a25050505050565b61286d612ff3565b6001600160a01b0316336001600160a01b03161461289d5760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b6000036128c857604051637bcd509160e01b815260040160405180910390fd5b610e0381613188565b6000805b60aa5481101561295657600081815260ab60205260409020546001600160a01b038085169161290491166141f7565b6001600160a01b03161480156129365750600081815260ab6020526040902054612936906001600160a01b0316610faf565b156129445760019150612956565b8061294e8161563f565b9150506128d5565b50919050565b612964612ff3565b6001600160a01b0316336001600160a01b0316146129945760405162461bcd60e51b8152600401610df1906155d1565b806001600160a01b03163b6000036129bf57604051637bcd509160e01b815260040160405180910390fd5b610e0381613272565b606060aa546001600160401b038111156129e4576129e4615613565b604051908082528060200260200182016040528015612a0d578160200160208202803683370190505b50905060005b8151811015611b5157600081815260ab602052604090205482516001600160a01b0390911690839083908110612a4b57612a4b615629565b6001600160a01b039092166020928302919091019091015280612a6d8161563f565b915050612a13565b60606073805480602002602001604051908101604052809291908181526020018280548015612acd57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612aaf575b5050505050905090565b612adf612ff3565b6001600160a01b0316336001600160a01b031614612b0f5760405162461bcd60e51b8152600401610df1906155d1565b6001600160a01b038216600090815260e8602052604090206001015415612d655760e7548060005b82811015612b9057846001600160a01b031660e78281548110612b5c57612b5c615629565b6000918252602090912001546001600160a01b031603612b7e57809150612b90565b80612b888161563f565b915050612b37565b50818103612b9e5750505050565b6001600160a01b038416600090815260e860205260409020548015612d61576001600160a01b038516600090815260e860205260408120818155600190810191909155831115612c605760e7612bf56001856155ab565b81548110612c0557612c05615629565b60009182526020909120015460e780546001600160a01b039092169184908110612c3157612c31615629565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60e7805480612c7157612c716157ab565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038716825260e9905260409020805460ff19166001179055612cc38482610dac61421d565b15612d1c57836001600160a01b0316856001600160a01b03167f7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e483604051612d0d91815260200190565b60405180910390a35050505050565b604080518281524760208201526001600160a01b0380871692908816917f3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a19101612d0d565b5050505b5050565b612d71612ff3565b6001600160a01b0316336001600160a01b031614612da15760405162461bcd60e51b8152600401610df1906155d1565b610e038161332a565b33612dbd6071546001600160a01b031690565b6001600160a01b031614612de457604051638aaf4a0760e01b815260040160405180910390fd5b6001600160a01b03831660009081526077602052604090205415612e1b57604051632f32dcdd60e11b815260040160405180910390fd5b612710811115612e3e57604051631b8454a360e21b815260040160405180910390fd5b607654821015612e615760405163fa0ae69360e01b815260040160405180910390fd5b6001600160a01b03831660009081526077602052604081209083612e886201518042615748565b612e9291906155be565b612e9f906201518061571b565b8083556001830184905560408051828152602081018690529192506001600160a01b038716917f6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f219101612856565b6060816001600160401b03811115612f0757612f07615613565b604051908082528060200260200182016040528015612f30578160200160208202803683370190505b5090506000612f3d610d26565b905060005b83811015612fae57612f7a858583818110612f5f57612f5f615629565b9050602002016020810190612f7491906152b5565b8361307f565b838281518110612f8c57612f8c615629565b9115156020928302919091019091015280612fa68161563f565b915050612f42565b505092915050565b6001600160a01b038116600090815260ac6020526040812054612fec9060ff166003811115612fe757612fe7615658565b61427d565b1592915050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103546001600160a01b031690565b6001811015613043576040516317b8970f60e01b815260040160405180910390fd5b60768190556040518181527f266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267906020015b60405180910390a150565b6001600160a01b03919091166000908152603760209081526040808320938352929052205460ff1690565b60008160038111156130be576130be615658565b8360038111156130d0576130d0615658565b1660ff1615159392505050565b6000610d206201518083615748565b607080546001600160a01b0319166001600160a01b0383169081179091556040519081527faa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f90602001613074565b607180546001600160a01b0319166001600160a01b0383169081179091556040519081527f6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf90602001613074565b606d80546001600160a01b0319166001600160a01b0383169081179091556040519081527fc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d90602001613074565b606f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b90602001613074565b606e80546001600160a01b0319166001600160a01b0383169081179091556040519081527f034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a611162190602001613074565b60a880546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d790602001613074565b60a98190556040518181527fb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b590602001613074565b60728190556040518181527f82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab490602001613074565b60a95481111561334d5760405163fa9546fb60e01b815260040160405180910390fd5b60ad8190556040518181527fa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e90602001613074565b60e58190556040518181527f17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a90602001613074565b60e68190556040518181527f0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b90602001613074565b6001600160a01b0381166000908152603a6020526040812054431115610d20565b60018201546001600160a01b03166000908152607460205260409020546134475760405163a64b34ad60e01b815260040160405180910390fd5b6005820181905560018201546040518281526001600160a01b03909116907fb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e790602001610ee6565b606e5460405163889998ef60e01b8152600481018490526001600160a01b0390911690600090829063889998ef90602401602060405180830381865afa1580156134dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613501919061575c565b60405163033cdc2b60e31b8152600481018690529091506000906001600160a01b038416906319e6e15890602401602060405180830381865afa15801561354c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613570919061575c565b90506000836001600160a01b031663f67e815287876040518363ffffffff1660e01b81526004016135a29291906157c1565b600060405180830381865afa1580156135bf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526135e79190810190615835565b9050600080600080607060009054906101000a90046001600160a01b03166001600160a01b0316631079402a6040518163ffffffff1660e01b8152600401608060405180830381865afa158015613642573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061366691906156e5565b935093509350935060005b89518110156136cf576136bd8b8b838151811061369057613690615629565b60200260200101518884815181106136aa576136aa615629565b60200260200101518a8c8a8a8a8a61429b565b806136c78161563f565b915050613671565b5050505050505050505050565b6000606060008084516001600160401b038111156136fc576136fc615613565b604051908082528060200260200182016040528015613725578160200160208202803683370190505b50925060005b85518110156139095785818151811061374657613746615629565b6020908102919091018101516001600160a01b03808216600090815260758452604080822060020154603986528183208d845290955290205491955091909116925060ff166137be576001600160a01b038084166000908152607560205260409020600301546137b99185911684614597565b6137f0565b6001600160a01b038316600090815260e3602052604081205460e48054919290916137ea9084906155be565b90915550505b6137f9836133ec565b15801561380d575061380b838861307f565b155b15613881576001600160a01b038316600090815260e1602052604090205461383590866155be565b6001600160a01b038416600090815260e1602052604090205485519196509085908390811061386657613866615629565b60200260200101818152505061387c838361467d565b6138c7565b6001600160a01b038316600090815260e1602090815260408083205460e0909252909120546138b091906155be565b60e460008282546138c191906155be565b90915550505b6001600160a01b038316600090815260e16020908152604080832083905560e0825280832083905560e3909152812055806139018161563f565b91505061372b565b5060e26000905550509250929050565b6071546001600160a01b03168215612d61576139358184614745565b156139dd5760405163566bce2360e11b81526001600160a01b0382169063acd79c469061396a90879086908a906004016158fa565b600060405180830381600087803b15801561398457600080fd5b505af1158015613998573d6000803e3d6000fd5b505050507f9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d0878385846040516139cf93929190615930565b60405180910390a150613a1c565b7fe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f8983858447604051613a129493929190615965565b60405180910390a1505b50505050565b60e754600080805b83831015613a1c5760e78381548110613a4557613a45615629565b60009182526020808320909101546001600160a01b031680835260e89091526040909120600181015491935091504210613b6b57805460e48054600090613a8d9084906155be565b90915550506001600160a01b038216600090815260e860205260408120818155600101819055613abc856159a2565b9450841115613b335760e78481548110613ad857613ad8615629565b60009182526020909120015460e780546001600160a01b039092169185908110613b0457613b04615629565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60e7805480613b4457613b446157ab565b600082815260209020810160001990810180546001600160a01b0319169055019055613a2a565b82613b758161563f565b935050613a2a565b60e4548015610e03576000613b9a606d546001600160a01b031690565b600060e481905560408051600481526024810182526020810180516001600160e01b03166359f778df60e01b179052905192935090916001600160a01b038416918591613be791906159b9565b60006040518083038185875af1925050503d8060008114613c24576040519150601f19603f3d011682016040523d82523d6000602084013e613c29565b606091505b505090508015613c7b57816001600160a01b03167fc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc884604051613c6e91815260200190565b60405180910390a2505050565b604080518481524760208201526001600160a01b038416917fa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b739101613c6e565b6060613cc56147a1565b6071546040516391f8723f60e01b81526000916001600160a01b0316906391f8723f90613cf790607390600401615a26565b600060405180830381865afa158015613d14573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613d3c9190810190615835565b60a854604051632907e73160e11b81529192506000916001600160a01b039091169063520fce6290613d7390607390600401615a26565b600060405180830381865afa158015613d90573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613db89190810190615835565b90506000613e276073805480602002602001604051908101604052809291908181526020018280548015613e1557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613df7575b5050505050848460a95460ad54614cb1565b9094509050613e37848287614d7b565b505050919050565b606f546000906001600160a01b031663fdadda816073613e604360016155be565b6040518363ffffffff1660e01b8152600401613e7d929190615a39565b600060405180830381865afa158015613e9a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613ec29190810190615a4c565b905060005b825181101561417e576000838281518110613ee457613ee4615629565b6020908102919091018101516001600160a01b0381166000908152603b909252604082205490925042111590613f1983611da9565b90506000613f26846133ec565b80613f475750858581518110613f3e57613f3e615629565b60200260200101515b80613f4f5750825b15905081158015613f5d5750805b15613fd8576001600160a01b038416600090815260ac6020526040902054613f9c9060019060ff166003811115613f9657613f96615658565b90614f01565b6001600160a01b038516600090815260ac60205260409020805460ff19166001836003811115613fce57613fce615658565b021790555061405a565b818015613fe3575080155b1561405a576001600160a01b038416600090815260ac60205260409020546140229060019060ff16600381111561401c5761401c615658565b90614f3c565b6001600160a01b038516600090815260ac60205260409020805460ff1916600183600381111561405457614054615658565b02179055505b600061406585610faf565b90508315811580156140745750805b156140e9576001600160a01b038616600090815260ac60205260409020546140ad9060029060ff166003811115613f9657613f96615658565b6001600160a01b038716600090815260ac60205260409020805460ff191660018360038111156140df576140df615658565b0217905550614165565b8180156140f4575080155b15614165576001600160a01b038616600090815260ac602052604090205461412d9060029060ff16600381111561401c5761401c615658565b6001600160a01b038716600090815260ac60205260409020805460ff1916600183600381111561415f5761415f615658565b02179055505b50505050505080806141769061563f565b915050613ec7565b5082847f283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b826141aa611438565b6040516141b79190615466565b60405180910390a382847f773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f6141ea6123c6565b604051611d9b9190615466565b6000610d2082614f78565b6000600154826142129190615748565b610d209060016155be565b6000836001600160a01b0316838390604051600060405180830381858888f193505050503d806000811461426d576040519150601f19603f3d011682016040523d82523d6000602084013e614272565b606091505b509095945050505050565b600081600381111561429157614291615658565b60ff161592915050565b84158080156142a8575086155b156142e1576142b5611b55565b60e2546142c29190615748565b6001600160a01b038a16600090815260e360205260409020555061458c565b80156142ed575061458c565b81871161430d57858860e254614303919061571b565b6142c29190615748565b60008761431c6127108b61571b565b6143269190615748565b90506000614336826127106155ab565b905085811115614477576001603960008d6001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002060006101000a81548160ff0219169083151502179055506001603760008d6001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002060006101000a81548160ff02191690831515021790555061440b85436143ed91906155be565b6001600160a01b038d166000908152603a6020526040902054614f9d565b6001600160a01b038c166000818152603a60209081526040808320859055805194855290840191909152600190830181905260608301528d917f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a906080015b60405180910390a3614588565b8681111561454f576001603960008d6001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002060006101000a81548160ff0219169083151502179055508b8b6001600160a01b03167f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a603a60008f6001600160a01b03166001600160a01b0316815260200190815260200160002054600080600160405161446a94939291909384526020840192909252151560408301521515606082015260800190565b871561458857878a60e254614564919061571b565b61456e9190615748565b6001600160a01b038c16600090815260e360205260409020555b5050505b505050505050505050565b6001600160a01b038316600090815260e360205260409020548015613a1c576145c38282610dac61421d565b1561462557816001600160a01b0316836001600160a01b0316856001600160a01b03167f72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c8460405161461791815260200190565b60405180910390a450505050565b816001600160a01b0316836001600160a01b0316856001600160a01b03167fd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b728447604051614617929190918252602082015260400190565b6001600160a01b038216600090815260e060205260409020548015611ea6576146a98282610dac61421d565b1561470057816001600160a01b0316836001600160a01b03167f1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec836040516146f391815260200190565b60405180910390a3505050565b604080518281524760208201526001600160a01b0380851692908616917f6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e91016146f3565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114614792576040519150601f19603f3d011682016040523d82523d6000602084013e614797565b606091505b5090949350505050565b6071546040805163af24542960e01b815290516001600160a01b0390921691600091839163af245429916004808201926020929091908290030181865afa1580156147f0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614814919061575c565b90506000826001600160a01b031663909791dd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614856573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061487a919061575c565b90506000836001600160a01b03166342ef3c3460736040518263ffffffff1660e01b81526004016148ab9190615a26565b600060405180830381865afa1580156148c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526148f09190810190615835565b607354909150600080826001600160401b0381111561491157614911615613565b60405190808252806020026020018201604052801561493a578160200160208202803683370190505b50905060008060005b85831015614c0a576073838154811061495e5761495e615629565b60009182526020808320909101546001600160a01b0316808352607590915260409091206006810154895192945090925015159089908990869081106149a6576149a6615629565b60200260200101511015614a145780614a0f5760006149c58b426155be565b600684018190556040518181529091506001600160a01b038516907f88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a342759060200160405180910390a2505b614a67565b8015614a67578160060160009055826001600160a01b03167f88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a342756000604051614a5e91815260200190565b60405180910390a25b60008260050154600014158015614a82575042836005015411155b80614aa557506001600160a01b038416600090815260e9602052604090205460ff165b905060008360060154600014158015614ac2575042846006015411155b90508180614acd5750805b15614b635789614adc8a6159a2565b99508981518110614aef57614aef615629565b60200260200101518a8781518110614b0957614b09615629565b6020908102919091010152848789614b208161563f565b9a5081518110614b3257614b32615629565b60200260200101906001600160a01b031690816001600160a01b031681525050614b5b85614fb4565b505050614943565b6001600160a01b0385166000908152607760205260409020548015801590614b8b5750428111155b15614bf4576001600160a01b0386166000818152607760209081526040808320600181018054918590559390935560048901839055518281529192917f86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9910160405180910390a2505b86614bfe8161563f565b97505050505050614943565b505082159050614ca8578181527f4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c9181604051614c469190615466565b60405180910390a16040516374d62f0360e11b81526001600160a01b0388169063e9ac5e0690614c7a908490600401615466565b600060405180830381600087803b158015614c9457600080fd5b505af11580156136cf573d6000803e3d6000fd5b50505050505050565b60606000806068905060008888888888604051602401614cd5959493929190615ad8565b60408051601f19818403018152919052602080820180516001600160e01b0316633bca0a8960e11b17905281518b519293506001929091600091614d189161571b565b614d239060406155be565b90506020840181888483895afa614d3957600093505b503d614d4457600092505b60208701965082614d6857604051630fc2632160e01b815260040160405180910390fd5b8651955050505050509550959350505050565b815b60aa54811015614dd957600081815260ab6020818152604080842080546001600160a01b0316855260ac8352908420805460ff19169055928490525280546001600160a01b031916905580614dd18161563f565b915050614d7d565b506000805b83811015614ebb576000858281518110614dfa57614dfa615629565b602090810291909101810151600085815260ab9092526040909120549091506001600160a01b0390811690821603614e3f5782614e368161563f565b93505050614ea9565b600083815260ab6020818152604080842080546001600160a01b03908116865260ac8452828620805460ff19908116909155908716808752928620805490911660031790559387905291905281546001600160a01b03191617905582614ea48161563f565b935050505b80614eb38161563f565b915050614dde565b508060aa81905550817f3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e85604051614ef39190615466565b60405180910390a250505050565b6000816003811115614f1557614f15615658565b836003811115614f2757614f27615658565b1760ff16600381111561101457611014615658565b6000816003811115614f5057614f50615658565b19836003811115614f6357614f63615658565b1660ff16600381111561101457611014615658565b6001600160a01b03808216600090815260756020526040812060030154909116610d20565b600081831015614fad5781611014565b5090919050565b6001600160a01b038116600090815260e960209081526040808320805460ff191690556074909152812054610e0391839190819003614ff1575050565b6001600160a01b038216600090815260756020908152604080832080546001600160a01b03199081168255600180830180548316905560028301805483169055600383018054909216909155600482018590556005820185905560069091018490556074835281842084905560779092528220828155810182905560738054909161507b916155ab565b8154811061508b5761508b615629565b6000918252602090912001546001600160a01b0390811691508316811461510e576001600160a01b03811660009081526074602052604090208290556073805482919084199081106150df576150df615629565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b607380548061511f5761511f6157ab565b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b6001600160a01b0381168114610e0357600080fd5b6000806040838503121561516e57600080fd5b823561517981615146565b9150602083013561518981615146565b809150509250929050565b600080604083850312156151a757600080fd5b82356151b281615146565b946020939093013593505050565b6000602082840312156151d257600080fd5b5035919050565b60008083601f8401126151eb57600080fd5b5081356001600160401b0381111561520257600080fd5b6020830191508360208260051b850101111561521d57600080fd5b9250929050565b60008060006040848603121561523957600080fd5b83356001600160401b0381111561524f57600080fd5b61525b868287016151d9565b909790965060209590950135949350505050565b6020808252825182820181905260009190848201906040850190845b818110156152a957835115158352928401929184019160010161528b565b50909695505050505050565b6000602082840312156152c757600080fd5b813561101481615146565b60018060a01b03808251168352806020830151166020840152806040830151166040840152806060830151166060840152506080810151608083015260a081015160a083015260c081015160c08301525050565b60e08101610d2082846152d2565b815181526020808301519082015260408101610d20565b8060408101831015610d2057600080fd5b6000806000806000806000806000806000806101a08d8f03121561537f57600080fd5b8c3561538a81615146565b9b5060208d013561539a81615146565b9a5060408d01356153aa81615146565b995060608d01356153ba81615146565b985060808d01356153ca81615146565b975060a08d01356153da81615146565b965060c08d0135955060e08d013594506101008d013593506101208d013592506101408d013591506154108e6101608f0161534b565b90509295989b509295989b509295989b565b600081518084526020808501945080840160005b8381101561545b5781516001600160a01b031687529582019590820190600101615436565b509495945050505050565b6020815260006110146020830184615422565b6000806020838503121561548c57600080fd5b82356001600160401b038111156154a257600080fd5b6154ae858286016151d9565b90969095509350505050565b6020808252825182820181905260009190848201906040850190845b818110156152a9576154e98385516152d2565b9284019260e092909201916001016154d6565b60008060006060848603121561551157600080fd5b833561551c81615146565b95602085013595506040909401359392505050565b600080600080600060a0868803121561554957600080fd5b853561555481615146565b9450602086013561556481615146565b9350604086013561557481615146565b9250606086013561558481615146565b949793965091946080013592915050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610d2057610d20615595565b80820180821115610d2057610d20615595565b60208082526022908201527f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460408201526132b960f11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161565157615651615595565b5060010190565b634e487b7160e01b600052602160045260246000fd5b805180151581146111fe57600080fd5b60008060006060848603121561569357600080fd5b61569c8461566e565b925060208401519150604084015190509250925092565b82815260408101600383106156d857634e487b7160e01b600052602160045260246000fd5b8260208301529392505050565b600080600080608085870312156156fb57600080fd5b505082516020840151604085015160609095015191969095509092509050565b8082028115828204841417610d2057610d20615595565b634e487b7160e01b600052601260045260246000fd5b60008261575757615757615732565b500490565b60006020828403121561576e57600080fd5b5051919050565b6040815260006157886040830185615422565b90508260208301529392505050565b6000826157a6576157a6615732565b500690565b634e487b7160e01b600052603160045260246000fd5b8281526040602082015260006157da6040830184615422565b949350505050565b604051601f8201601f191681016001600160401b038111828210171561580a5761580a615613565b604052919050565b60006001600160401b0382111561582b5761582b615613565b5060051b60200190565b6000602080838503121561584857600080fd5b82516001600160401b0381111561585e57600080fd5b8301601f8101851361586f57600080fd5b805161588261587d82615812565b6157e2565b81815260059190911b820183019083810190878311156158a157600080fd5b928401925b828410156158bf578351825292840192908401906158a6565b979650505050505050565b600081518084526020808501945080840160005b8381101561545b578151875295820195908201906001016158de565b60608152600061590d6060830186615422565b828103602084015261591f81866158ca565b915050826040830152949350505050565b8381526060602082015260006159496060830185615422565b828103604084015261595b81856158ca565b9695505050505050565b84815260806020820152600061597e6080830186615422565b828103604084015261599081866158ca565b91505082606083015295945050505050565b6000816159b1576159b1615595565b506000190190565b6000825160005b818110156159da57602081860181015185830152016159c0565b506000920191825250919050565b6000815480845260208085019450836000528060002060005b8381101561545b5781546001600160a01b031687529582019560019182019101615a01565b60208152600061101460208301846159e8565b60408152600061578860408301856159e8565b60006020808385031215615a5f57600080fd5b82516001600160401b03811115615a7557600080fd5b8301601f81018513615a8657600080fd5b8051615a9461587d82615812565b81815260059190911b82018301908381019087831115615ab357600080fd5b928401925b828410156158bf57615ac98461566e565b82529284019290840190615ab8565b60a081526000615aeb60a0830188615422565b8281036020840152615afd81886158ca565b90508281036040840152615b1181876158ca565b6060840195909552505060800152939250505056fea2646970667358221220d22ea8190a3508d1a26e1dd9c4d48d3b7fc424473f029a17f030860ac96ee69b64736f6c63430008110033",
}

// ValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMetaData.ABI instead.
var ValidatorABI = ValidatorMetaData.ABI

// ValidatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorMetaData.Bin instead.
var ValidatorBin = ValidatorMetaData.Bin

// DeployValidator deploys a new Ethereum contract, binding an instance of Validator to it.
func DeployValidator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validator, error) {
	parsed, err := ValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// IsPeriodEndingInternal is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_Validator *ValidatorCaller) IsPeriodEndingInternal(opts *bind.CallOpts, _newPeriod *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "_isPeriodEnding", _newPeriod)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeriodEndingInternal is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_Validator *ValidatorSession) IsPeriodEndingInternal(_newPeriod *big.Int) (bool, error) {
	return _Validator.Contract.IsPeriodEndingInternal(&_Validator.CallOpts, _newPeriod)
}

// IsPeriodEndingInternal is a free data retrieval call binding the contract method 0x9b8c334b.
//
// Solidity: function _isPeriodEnding(uint256 _newPeriod) view returns(bool)
func (_Validator *ValidatorCallerSession) IsPeriodEndingInternal(_newPeriod *big.Int) (bool, error) {
	return _Validator.Contract.IsPeriodEndingInternal(&_Validator.CallOpts, _newPeriod)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorCaller) BridgeTrackingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "bridgeTrackingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorSession) BridgeTrackingContract() (common.Address, error) {
	return _Validator.Contract.BridgeTrackingContract(&_Validator.CallOpts)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorCallerSession) BridgeTrackingContract() (common.Address, error) {
	return _Validator.Contract.BridgeTrackingContract(&_Validator.CallOpts)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) CheckJailed(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkJailed", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorSession) CheckJailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.CheckJailed(&_Validator.CallOpts, _addr)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) CheckJailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.CheckJailed(&_Validator.CallOpts, _addr)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCaller) CheckJailedAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkJailedAtBlock", _addr, _blockNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorSession) CheckJailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.CheckJailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCallerSession) CheckJailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.CheckJailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCaller) CheckManyJailed(opts *bind.CallOpts, _addrList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkManyJailed", _addrList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorSession) CheckManyJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckManyJailed(&_Validator.CallOpts, _addrList)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) CheckManyJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckManyJailed(&_Validator.CallOpts, _addrList)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0xf2811bcc.
//
// Solidity: function checkMiningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorCaller) CheckMiningRewardDeprecated(opts *bind.CallOpts, _blockProducers []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkMiningRewardDeprecated", _blockProducers)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0xf2811bcc.
//
// Solidity: function checkMiningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorSession) CheckMiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecated(&_Validator.CallOpts, _blockProducers)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0xf2811bcc.
//
// Solidity: function checkMiningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) CheckMiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecated(&_Validator.CallOpts, _blockProducers)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x1b6e0a99.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorCaller) CheckMiningRewardDeprecatedAtPeriod(opts *bind.CallOpts, _blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkMiningRewardDeprecatedAtPeriod", _blockProducers, _period)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x1b6e0a99.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorSession) CheckMiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducers, _period)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x1b6e0a99.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) CheckMiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducers, _period)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorSession) CurrentPeriod() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriod(&_Validator.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorCallerSession) CurrentPeriod() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriod(&_Validator.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorCaller) CurrentPeriodStartAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "currentPeriodStartAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriodStartAtBlock(&_Validator.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorCallerSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriodStartAtBlock(&_Validator.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorCaller) EmergencyExitLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "emergencyExitLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _Validator.Contract.EmergencyExitLockedAmount(&_Validator.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorCallerSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _Validator.Contract.EmergencyExitLockedAmount(&_Validator.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorCaller) EmergencyExpiryDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "emergencyExpiryDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _Validator.Contract.EmergencyExpiryDuration(&_Validator.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorCallerSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _Validator.Contract.EmergencyExpiryDuration(&_Validator.CallOpts)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorCaller) EpochEndingAt(opts *bind.CallOpts, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "epochEndingAt", _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _Validator.Contract.EpochEndingAt(&_Validator.CallOpts, _block)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorCallerSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _Validator.Contract.EpochEndingAt(&_Validator.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorCaller) EpochOf(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "epochOf", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _Validator.Contract.EpochOf(&_Validator.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorCallerSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _Validator.Contract.EpochOf(&_Validator.CallOpts, _block)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorCaller) GetBlockProducers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getBlockProducers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorSession) GetBlockProducers() ([]common.Address, error) {
	return _Validator.Contract.GetBlockProducers(&_Validator.CallOpts)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorCallerSession) GetBlockProducers() ([]common.Address, error) {
	return _Validator.Contract.GetBlockProducers(&_Validator.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorCaller) GetBridgeOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getBridgeOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorCallerSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorCaller) GetCandidateInfo(opts *bind.CallOpts, _candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCandidateInfo", _candidate)

	if err != nil {
		return *new(ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new(ICandidateManagerValidatorCandidate)).(*ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorCallerSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorCaller) GetCandidateInfos(opts *bind.CallOpts) ([]ICandidateManagerValidatorCandidate, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCandidateInfos")

	if err != nil {
		return *new([]ICandidateManagerValidatorCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICandidateManagerValidatorCandidate)).(*[]ICandidateManagerValidatorCandidate)

	return out0, err

}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorCallerSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorCaller) GetCommissionChangeSchedule(opts *bind.CallOpts, _candidate common.Address) (ICandidateManagerCommissionSchedule, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCommissionChangeSchedule", _candidate)

	if err != nil {
		return *new(ICandidateManagerCommissionSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(ICandidateManagerCommissionSchedule)).(*ICandidateManagerCommissionSchedule)

	return out0, err

}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorSession) GetCommissionChangeSchedule(_candidate common.Address) (ICandidateManagerCommissionSchedule, error) {
	return _Validator.Contract.GetCommissionChangeSchedule(&_Validator.CallOpts, _candidate)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorCallerSession) GetCommissionChangeSchedule(_candidate common.Address) (ICandidateManagerCommissionSchedule, error) {
	return _Validator.Contract.GetCommissionChangeSchedule(&_Validator.CallOpts, _candidate)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorCaller) GetEmergencyExitInfo(opts *bind.CallOpts, _consensusAddr common.Address) (ICommonInfoEmergencyExitInfo, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getEmergencyExitInfo", _consensusAddr)

	if err != nil {
		return *new(ICommonInfoEmergencyExitInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ICommonInfoEmergencyExitInfo)).(*ICommonInfoEmergencyExitInfo)

	return out0, err

}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorSession) GetEmergencyExitInfo(_consensusAddr common.Address) (ICommonInfoEmergencyExitInfo, error) {
	return _Validator.Contract.GetEmergencyExitInfo(&_Validator.CallOpts, _consensusAddr)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorCallerSession) GetEmergencyExitInfo(_consensusAddr common.Address) (ICommonInfoEmergencyExitInfo, error) {
	return _Validator.Contract.GetEmergencyExitInfo(&_Validator.CallOpts, _consensusAddr)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) GetJailedTimeLeft(opts *bind.CallOpts, _addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getJailedTimeLeft", _addr)

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
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) GetJailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeft(&_Validator.CallOpts, _addr)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) GetJailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeft(&_Validator.CallOpts, _addr)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) GetJailedTimeLeftAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getJailedTimeLeftAtBlock", _addr, _blockNum)

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
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) GetJailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) GetJailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorCaller) GetLastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getLastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _Validator.Contract.GetLastUpdatedBlock(&_Validator.CallOpts)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorCallerSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _Validator.Contract.GetLastUpdatedBlock(&_Validator.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorCaller) GetValidatorCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getValidatorCandidates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorSession) GetValidatorCandidates() ([]common.Address, error) {
	return _Validator.Contract.GetValidatorCandidates(&_Validator.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorCallerSession) GetValidatorCandidates() ([]common.Address, error) {
	return _Validator.Contract.GetValidatorCandidates(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_Validator *ValidatorCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_Validator *ValidatorSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList)
func (_Validator *ValidatorCallerSession) GetValidators() ([]common.Address, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsBlockProducer(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isBlockProducer", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsBlockProducer(&_Validator.CallOpts, _addr)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsBlockProducer(&_Validator.CallOpts, _addr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_Validator *ValidatorCaller) IsBridgeOperator(opts *bind.CallOpts, _bridgeOperatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isBridgeOperator", _bridgeOperatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_Validator *ValidatorSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _Validator.Contract.IsBridgeOperator(&_Validator.CallOpts, _bridgeOperatorAddr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _result)
func (_Validator *ValidatorCallerSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _Validator.Contract.IsBridgeOperator(&_Validator.CallOpts, _bridgeOperatorAddr)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorCaller) IsCandidateAdmin(opts *bind.CallOpts, _candidate common.Address, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isCandidateAdmin", _candidate, _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _Validator.Contract.IsCandidateAdmin(&_Validator.CallOpts, _candidate, _admin)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorCallerSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _Validator.Contract.IsCandidateAdmin(&_Validator.CallOpts, _candidate, _admin)
}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorCaller) IsOperatingBridge(opts *bind.CallOpts, _consensusAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isOperatingBridge", _consensusAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorSession) IsOperatingBridge(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.IsOperatingBridge(&_Validator.CallOpts, _consensusAddr)
}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsOperatingBridge(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.IsOperatingBridge(&_Validator.CallOpts, _consensusAddr)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorCaller) IsPeriodEnding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isPeriodEnding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorSession) IsPeriodEnding() (bool, error) {
	return _Validator.Contract.IsPeriodEnding(&_Validator.CallOpts)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorCallerSession) IsPeriodEnding() (bool, error) {
	return _Validator.Contract.IsPeriodEnding(&_Validator.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsValidator(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsValidatorCandidate(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isValidatorCandidate", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorCandidate(&_Validator.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorCandidate(&_Validator.CallOpts, _addr)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorCaller) MaintenanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maintenanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorSession) MaintenanceContract() (common.Address, error) {
	return _Validator.Contract.MaintenanceContract(&_Validator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorCallerSession) MaintenanceContract() (common.Address, error) {
	return _Validator.Contract.MaintenanceContract(&_Validator.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorCaller) MaxPrioritizedValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxPrioritizedValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxPrioritizedValidatorNumber(&_Validator.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorCallerSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxPrioritizedValidatorNumber(&_Validator.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorCaller) MaxValidatorCandidate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxValidatorCandidate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorSession) MaxValidatorCandidate() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorCandidate(&_Validator.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorCallerSession) MaxValidatorCandidate() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorCandidate(&_Validator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorNumber(&_Validator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorNumber(&_Validator.CallOpts)
}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorCaller) MinEffectiveDaysOnwards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "minEffectiveDaysOnwards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorSession) MinEffectiveDaysOnwards() (*big.Int, error) {
	return _Validator.Contract.MinEffectiveDaysOnwards(&_Validator.CallOpts)
}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorCallerSession) MinEffectiveDaysOnwards() (*big.Int, error) {
	return _Validator.Contract.MinEffectiveDaysOnwards(&_Validator.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorCaller) NumberOfBlocksInEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "numberOfBlocksInEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _Validator.Contract.NumberOfBlocksInEpoch(&_Validator.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorCallerSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _Validator.Contract.NumberOfBlocksInEpoch(&_Validator.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorCaller) PrecompilePickValidatorSetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "precompilePickValidatorSetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _Validator.Contract.PrecompilePickValidatorSetAddress(&_Validator.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorCallerSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _Validator.Contract.PrecompilePickValidatorSetAddress(&_Validator.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorCaller) PrecompileSortValidatorsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "precompileSortValidatorsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _Validator.Contract.PrecompileSortValidatorsAddress(&_Validator.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorCallerSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _Validator.Contract.PrecompileSortValidatorsAddress(&_Validator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorCaller) RoninTrustedOrganizationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "roninTrustedOrganizationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Validator.Contract.RoninTrustedOrganizationContract(&_Validator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorCallerSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Validator.Contract.RoninTrustedOrganizationContract(&_Validator.CallOpts)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorCaller) SlashIndicatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "slashIndicatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorSession) SlashIndicatorContract() (common.Address, error) {
	return _Validator.Contract.SlashIndicatorContract(&_Validator.CallOpts)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorCallerSession) SlashIndicatorContract() (common.Address, error) {
	return _Validator.Contract.SlashIndicatorContract(&_Validator.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorSession) StakingContract() (common.Address, error) {
	return _Validator.Contract.StakingContract(&_Validator.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorCallerSession) StakingContract() (common.Address, error) {
	return _Validator.Contract.StakingContract(&_Validator.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorCaller) StakingVestingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "stakingVestingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorSession) StakingVestingContract() (common.Address, error) {
	return _Validator.Contract.StakingVestingContract(&_Validator.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorCallerSession) StakingVestingContract() (common.Address, error) {
	return _Validator.Contract.StakingVestingContract(&_Validator.CallOpts)
}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorCaller) TotalBlockProducers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalBlockProducers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorSession) TotalBlockProducers() (*big.Int, error) {
	return _Validator.Contract.TotalBlockProducers(&_Validator.CallOpts)
}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorCallerSession) TotalBlockProducers() (*big.Int, error) {
	return _Validator.Contract.TotalBlockProducers(&_Validator.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorCaller) TotalBridgeOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalBridgeOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorCallerSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorCaller) TotalDeprecatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalDeprecatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorSession) TotalDeprecatedReward() (*big.Int, error) {
	return _Validator.Contract.TotalDeprecatedReward(&_Validator.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorCallerSession) TotalDeprecatedReward() (*big.Int, error) {
	return _Validator.Contract.TotalDeprecatedReward(&_Validator.CallOpts)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool, uint256)
func (_Validator *ValidatorCaller) TryGetPeriodOfEpoch(opts *bind.CallOpts, _epoch *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "tryGetPeriodOfEpoch", _epoch)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool, uint256)
func (_Validator *ValidatorSession) TryGetPeriodOfEpoch(_epoch *big.Int) (bool, *big.Int, error) {
	return _Validator.Contract.TryGetPeriodOfEpoch(&_Validator.CallOpts, _epoch)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool, uint256)
func (_Validator *ValidatorCallerSession) TryGetPeriodOfEpoch(_epoch *big.Int) (bool, *big.Int, error) {
	return _Validator.Contract.TryGetPeriodOfEpoch(&_Validator.CallOpts, _epoch)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorSession) ValidatorCount() (*big.Int, error) {
	return _Validator.Contract.ValidatorCount(&_Validator.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorCallerSession) ValidatorCount() (*big.Int, error) {
	return _Validator.Contract.ValidatorCount(&_Validator.CallOpts)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorTransactor) ExecBailOut(opts *bind.TransactOpts, _validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execBailOut", _validatorAddr, _period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorSession) ExecBailOut(_validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecBailOut(&_Validator.TransactOpts, _validatorAddr, _period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorTransactorSession) ExecBailOut(_validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecBailOut(&_Validator.TransactOpts, _validatorAddr, _period)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorTransactor) ExecEmergencyExit(opts *bind.TransactOpts, _consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execEmergencyExit", _consensusAddr, _secLeftToRevoke)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorSession) ExecEmergencyExit(_consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecEmergencyExit(&_Validator.TransactOpts, _consensusAddr, _secLeftToRevoke)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorTransactorSession) ExecEmergencyExit(_consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecEmergencyExit(&_Validator.TransactOpts, _consensusAddr, _secLeftToRevoke)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorTransactor) ExecReleaseLockedFundForEmergencyExitRequest(opts *bind.TransactOpts, _consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execReleaseLockedFundForEmergencyExitRequest", _consensusAddr, _recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorSession) ExecReleaseLockedFundForEmergencyExitRequest(_consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_Validator.TransactOpts, _consensusAddr, _recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorTransactorSession) ExecReleaseLockedFundForEmergencyExitRequest(_consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_Validator.TransactOpts, _consensusAddr, _recipient)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactor) ExecRequestUpdateCommissionRate(opts *bind.TransactOpts, _consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execRequestUpdateCommissionRate", _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorSession) ExecRequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestUpdateCommissionRate(&_Validator.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactorSession) ExecRequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestUpdateCommissionRate(&_Validator.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x570ccb13.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_Validator *ValidatorTransactor) ExecSlash(opts *bind.TransactOpts, _validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execSlash", _validatorAddr, _newJailedUntil, _slashAmount)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x570ccb13.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_Validator *ValidatorSession) ExecSlash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecSlash(&_Validator.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x570ccb13.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount) returns()
func (_Validator *ValidatorTransactorSession) ExecSlash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecSlash(&_Validator.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactor) GrantValidatorCandidate(opts *bind.TransactOpts, _candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "grantValidatorCandidate", _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorSession) GrantValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.GrantValidatorCandidate(&_Validator.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactorSession) GrantValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.GrantValidatorCandidate(&_Validator.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x6efa12bd.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorTransactor) RequestRevokeCandidate(opts *bind.TransactOpts, _consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "requestRevokeCandidate", _consensusAddr, _secsLeft)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x6efa12bd.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorSession) RequestRevokeCandidate(_consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.RequestRevokeCandidate(&_Validator.TransactOpts, _consensusAddr, _secsLeft)
}

// RequestRevokeCandidate is a paid mutator transaction binding the contract method 0x6efa12bd.
//
// Solidity: function requestRevokeCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorTransactorSession) RequestRevokeCandidate(_consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.RequestRevokeCandidate(&_Validator.TransactOpts, _consensusAddr, _secsLeft)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetBridgeTrackingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setBridgeTrackingContract", _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetBridgeTrackingContract(&_Validator.TransactOpts, _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetBridgeTrackingContract(&_Validator.TransactOpts, _addr)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorTransactor) SetEmergencyExitLockedAmount(opts *bind.TransactOpts, _emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setEmergencyExitLockedAmount", _emergencyExitLockedAmount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorSession) SetEmergencyExitLockedAmount(_emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExitLockedAmount(&_Validator.TransactOpts, _emergencyExitLockedAmount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorTransactorSession) SetEmergencyExitLockedAmount(_emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExitLockedAmount(&_Validator.TransactOpts, _emergencyExitLockedAmount)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorTransactor) SetEmergencyExpiryDuration(opts *bind.TransactOpts, _emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setEmergencyExpiryDuration", _emergencyExpiryDuration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorSession) SetEmergencyExpiryDuration(_emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExpiryDuration(&_Validator.TransactOpts, _emergencyExpiryDuration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorTransactorSession) SetEmergencyExpiryDuration(_emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExpiryDuration(&_Validator.TransactOpts, _emergencyExpiryDuration)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetMaintenanceContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaintenanceContract", _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetMaintenanceContract(&_Validator.TransactOpts, _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetMaintenanceContract(&_Validator.TransactOpts, _addr)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorTransactor) SetMaxPrioritizedValidatorNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxPrioritizedValidatorNumber", _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxPrioritizedValidatorNumber(&_Validator.TransactOpts, _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorTransactorSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxPrioritizedValidatorNumber(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorTransactor) SetMaxValidatorCandidate(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxValidatorCandidate", _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorCandidate(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorTransactorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorCandidate(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorTransactor) SetMaxValidatorNumber(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxValidatorNumber", _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorNumber(&_Validator.TransactOpts, _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorTransactorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorNumber(&_Validator.TransactOpts, _max)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorTransactor) SetMinEffectiveDaysOnwards(opts *bind.TransactOpts, _numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMinEffectiveDaysOnwards", _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMinEffectiveDaysOnwards(&_Validator.TransactOpts, _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorTransactorSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMinEffectiveDaysOnwards(&_Validator.TransactOpts, _numOfDays)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetRoninTrustedOrganizationContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setRoninTrustedOrganizationContract", _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetRoninTrustedOrganizationContract(&_Validator.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetRoninTrustedOrganizationContract(&_Validator.TransactOpts, _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetSlashIndicatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setSlashIndicatorContract", _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetSlashIndicatorContract(&_Validator.TransactOpts, _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetSlashIndicatorContract(&_Validator.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetStakingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setStakingContract", _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetStakingVestingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setStakingVestingContract", _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingVestingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingVestingContract(&_Validator.TransactOpts, _addr)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorTransactor) SubmitBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "submitBlockReward")
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlockReward(&_Validator.TransactOpts)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorTransactorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlockReward(&_Validator.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorTransactor) WrapUpEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "wrapUpEpoch")
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _Validator.Contract.WrapUpEpoch(&_Validator.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorTransactorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _Validator.Contract.WrapUpEpoch(&_Validator.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Validator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validator.Contract.Fallback(&_Validator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validator.Contract.Fallback(&_Validator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorSession) Receive() (*types.Transaction, error) {
	return _Validator.Contract.Receive(&_Validator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Validator.Contract.Receive(&_Validator.TransactOpts)
}

// ValidatorBlockProducerSetUpdatedIterator is returned from FilterBlockProducerSetUpdated and is used to iterate over the raw logs and unpacked data for BlockProducerSetUpdated events raised by the Validator contract.
type ValidatorBlockProducerSetUpdatedIterator struct {
	Event *ValidatorBlockProducerSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockProducerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockProducerSetUpdated)
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
		it.Event = new(ValidatorBlockProducerSetUpdated)
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
func (it *ValidatorBlockProducerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockProducerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockProducerSetUpdated represents a BlockProducerSetUpdated event raised by the Validator contract.
type ValidatorBlockProducerSetUpdated struct {
	Period         *big.Int
	Epoch          *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerSetUpdated is a free log retrieval operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterBlockProducerSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*ValidatorBlockProducerSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockProducerSetUpdatedIterator{contract: _Validator.contract, event: "BlockProducerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerSetUpdated is a free log subscription operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchBlockProducerSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBlockProducerSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockProducerSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBlockProducerSetUpdated(log types.Log) (*ValidatorBlockProducerSetUpdated, error) {
	event := new(ValidatorBlockProducerSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBlockRewardDeprecatedIterator is returned from FilterBlockRewardDeprecated and is used to iterate over the raw logs and unpacked data for BlockRewardDeprecated events raised by the Validator contract.
type ValidatorBlockRewardDeprecatedIterator struct {
	Event *ValidatorBlockRewardDeprecated // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockRewardDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockRewardDeprecated)
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
		it.Event = new(ValidatorBlockRewardDeprecated)
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
func (it *ValidatorBlockRewardDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockRewardDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockRewardDeprecated represents a BlockRewardDeprecated event raised by the Validator contract.
type ValidatorBlockRewardDeprecated struct {
	CoinbaseAddr   common.Address
	RewardAmount   *big.Int
	DeprecatedType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardDeprecated is a free log retrieval operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_Validator *ValidatorFilterer) FilterBlockRewardDeprecated(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*ValidatorBlockRewardDeprecatedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockRewardDeprecatedIterator{contract: _Validator.contract, event: "BlockRewardDeprecated", logs: logs, sub: sub}, nil
}

// WatchBlockRewardDeprecated is a free log subscription operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_Validator *ValidatorFilterer) WatchBlockRewardDeprecated(opts *bind.WatchOpts, sink chan<- *ValidatorBlockRewardDeprecated, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockRewardDeprecated)
				if err := _Validator.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBlockRewardDeprecated(log types.Log) (*ValidatorBlockRewardDeprecated, error) {
	event := new(ValidatorBlockRewardDeprecated)
	if err := _Validator.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBlockRewardSubmittedIterator is returned from FilterBlockRewardSubmitted and is used to iterate over the raw logs and unpacked data for BlockRewardSubmitted events raised by the Validator contract.
type ValidatorBlockRewardSubmittedIterator struct {
	Event *ValidatorBlockRewardSubmitted // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockRewardSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockRewardSubmitted)
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
		it.Event = new(ValidatorBlockRewardSubmitted)
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
func (it *ValidatorBlockRewardSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockRewardSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockRewardSubmitted represents a BlockRewardSubmitted event raised by the Validator contract.
type ValidatorBlockRewardSubmitted struct {
	CoinbaseAddr    common.Address
	SubmittedAmount *big.Int
	BonusAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardSubmitted is a free log retrieval operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_Validator *ValidatorFilterer) FilterBlockRewardSubmitted(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*ValidatorBlockRewardSubmittedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockRewardSubmittedIterator{contract: _Validator.contract, event: "BlockRewardSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockRewardSubmitted is a free log subscription operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_Validator *ValidatorFilterer) WatchBlockRewardSubmitted(opts *bind.WatchOpts, sink chan<- *ValidatorBlockRewardSubmitted, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockRewardSubmitted)
				if err := _Validator.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBlockRewardSubmitted(log types.Log) (*ValidatorBlockRewardSubmitted, error) {
	event := new(ValidatorBlockRewardSubmitted)
	if err := _Validator.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorRewardDistributedIterator is returned from FilterBridgeOperatorRewardDistributed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributed events raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributedIterator struct {
	Event *ValidatorBridgeOperatorRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorRewardDistributed)
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
		it.Event = new(ValidatorBridgeOperatorRewardDistributed)
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
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorRewardDistributed represents a BridgeOperatorRewardDistributed event raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributed struct {
	ConsensusAddr  common.Address
	BridgeOperator common.Address
	RecipientAddr  common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributed is a free log retrieval operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (*ValidatorBridgeOperatorRewardDistributedIterator, error) {

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

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorRewardDistributedIterator{contract: _Validator.contract, event: "BridgeOperatorRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributed is a free log subscription operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorRewardDistributed, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBridgeOperatorRewardDistributed(log types.Log) (*ValidatorBridgeOperatorRewardDistributed, error) {
	event := new(ValidatorBridgeOperatorRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorRewardDistributionFailedIterator is returned from FilterBridgeOperatorRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributionFailed events raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributionFailedIterator struct {
	Event *ValidatorBridgeOperatorRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorRewardDistributionFailed)
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
		it.Event = new(ValidatorBridgeOperatorRewardDistributionFailed)
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
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorRewardDistributionFailed represents a BridgeOperatorRewardDistributionFailed event raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributionFailed struct {
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
func (_Validator *ValidatorFilterer) FilterBridgeOperatorRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (*ValidatorBridgeOperatorRewardDistributionFailedIterator, error) {

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

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorRewardDistributionFailedIterator{contract: _Validator.contract, event: "BridgeOperatorRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributionFailed is a free log subscription operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorRewardDistributionFailed, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBridgeOperatorRewardDistributionFailed(log types.Log) (*ValidatorBridgeOperatorRewardDistributionFailed, error) {
	event := new(ValidatorBridgeOperatorRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorSetUpdatedIterator is returned from FilterBridgeOperatorSetUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorSetUpdated events raised by the Validator contract.
type ValidatorBridgeOperatorSetUpdatedIterator struct {
	Event *ValidatorBridgeOperatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorSetUpdated)
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
		it.Event = new(ValidatorBridgeOperatorSetUpdated)
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
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorSetUpdated represents a BridgeOperatorSetUpdated event raised by the Validator contract.
type ValidatorBridgeOperatorSetUpdated struct {
	Period          *big.Int
	Epoch           *big.Int
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSetUpdated is a free log retrieval operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*ValidatorBridgeOperatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorSetUpdatedIterator{contract: _Validator.contract, event: "BridgeOperatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSetUpdated is a free log subscription operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseBridgeOperatorSetUpdated(log types.Log) (*ValidatorBridgeOperatorSetUpdated, error) {
	event := new(ValidatorBridgeOperatorSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeTrackingContractUpdatedIterator is returned from FilterBridgeTrackingContractUpdated and is used to iterate over the raw logs and unpacked data for BridgeTrackingContractUpdated events raised by the Validator contract.
type ValidatorBridgeTrackingContractUpdatedIterator struct {
	Event *ValidatorBridgeTrackingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeTrackingContractUpdated)
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
		it.Event = new(ValidatorBridgeTrackingContractUpdated)
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
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeTrackingContractUpdated represents a BridgeTrackingContractUpdated event raised by the Validator contract.
type ValidatorBridgeTrackingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingContractUpdated is a free log retrieval operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterBridgeTrackingContractUpdated(opts *bind.FilterOpts) (*ValidatorBridgeTrackingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeTrackingContractUpdatedIterator{contract: _Validator.contract, event: "BridgeTrackingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingContractUpdated is a free log subscription operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchBridgeTrackingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeTrackingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeTrackingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
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

// ParseBridgeTrackingContractUpdated is a log parse operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseBridgeTrackingContractUpdated(log types.Log) (*ValidatorBridgeTrackingContractUpdated, error) {
	event := new(ValidatorBridgeTrackingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateGrantedIterator is returned from FilterCandidateGranted and is used to iterate over the raw logs and unpacked data for CandidateGranted events raised by the Validator contract.
type ValidatorCandidateGrantedIterator struct {
	Event *ValidatorCandidateGranted // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateGranted)
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
		it.Event = new(ValidatorCandidateGranted)
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
func (it *ValidatorCandidateGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateGranted represents a CandidateGranted event raised by the Validator contract.
type ValidatorCandidateGranted struct {
	ConsensusAddr  common.Address
	TreasuryAddr   common.Address
	Admin          common.Address
	BridgeOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidateGranted is a free log retrieval operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) FilterCandidateGranted(opts *bind.FilterOpts, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (*ValidatorCandidateGrantedIterator, error) {

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

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateGrantedIterator{contract: _Validator.contract, event: "CandidateGranted", logs: logs, sub: sub}, nil
}

// WatchCandidateGranted is a free log subscription operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) WatchCandidateGranted(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateGranted, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateGranted)
				if err := _Validator.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
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

// ParseCandidateGranted is a log parse operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) ParseCandidateGranted(log types.Log) (*ValidatorCandidateGranted, error) {
	event := new(ValidatorCandidateGranted)
	if err := _Validator.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateRevokingTimestampUpdatedIterator is returned from FilterCandidateRevokingTimestampUpdated and is used to iterate over the raw logs and unpacked data for CandidateRevokingTimestampUpdated events raised by the Validator contract.
type ValidatorCandidateRevokingTimestampUpdatedIterator struct {
	Event *ValidatorCandidateRevokingTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateRevokingTimestampUpdated)
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
		it.Event = new(ValidatorCandidateRevokingTimestampUpdated)
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
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateRevokingTimestampUpdated represents a CandidateRevokingTimestampUpdated event raised by the Validator contract.
type ValidatorCandidateRevokingTimestampUpdated struct {
	ConsensusAddr     common.Address
	RevokingTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCandidateRevokingTimestampUpdated is a free log retrieval operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) FilterCandidateRevokingTimestampUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCandidateRevokingTimestampUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateRevokingTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateRevokingTimestampUpdatedIterator{contract: _Validator.contract, event: "CandidateRevokingTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateRevokingTimestampUpdated is a free log subscription operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) WatchCandidateRevokingTimestampUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateRevokingTimestampUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateRevokingTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateRevokingTimestampUpdated)
				if err := _Validator.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
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
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) ParseCandidateRevokingTimestampUpdated(log types.Log) (*ValidatorCandidateRevokingTimestampUpdated, error) {
	event := new(ValidatorCandidateRevokingTimestampUpdated)
	if err := _Validator.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateTopupDeadlineUpdatedIterator is returned from FilterCandidateTopupDeadlineUpdated and is used to iterate over the raw logs and unpacked data for CandidateTopupDeadlineUpdated events raised by the Validator contract.
type ValidatorCandidateTopupDeadlineUpdatedIterator struct {
	Event *ValidatorCandidateTopupDeadlineUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateTopupDeadlineUpdated)
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
		it.Event = new(ValidatorCandidateTopupDeadlineUpdated)
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
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateTopupDeadlineUpdated represents a CandidateTopupDeadlineUpdated event raised by the Validator contract.
type ValidatorCandidateTopupDeadlineUpdated struct {
	ConsensusAddr common.Address
	TopupDeadline *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCandidateTopupDeadlineUpdated is a free log retrieval operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) FilterCandidateTopupDeadlineUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCandidateTopupDeadlineUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateTopupDeadlineUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateTopupDeadlineUpdatedIterator{contract: _Validator.contract, event: "CandidateTopupDeadlineUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateTopupDeadlineUpdated is a free log subscription operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) WatchCandidateTopupDeadlineUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateTopupDeadlineUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateTopupDeadlineUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateTopupDeadlineUpdated)
				if err := _Validator.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
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
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) ParseCandidateTopupDeadlineUpdated(log types.Log) (*ValidatorCandidateTopupDeadlineUpdated, error) {
	event := new(ValidatorCandidateTopupDeadlineUpdated)
	if err := _Validator.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidatesRevokedIterator is returned from FilterCandidatesRevoked and is used to iterate over the raw logs and unpacked data for CandidatesRevoked events raised by the Validator contract.
type ValidatorCandidatesRevokedIterator struct {
	Event *ValidatorCandidatesRevoked // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidatesRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidatesRevoked)
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
		it.Event = new(ValidatorCandidatesRevoked)
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
func (it *ValidatorCandidatesRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidatesRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidatesRevoked represents a CandidatesRevoked event raised by the Validator contract.
type ValidatorCandidatesRevoked struct {
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidatesRevoked is a free log retrieval operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterCandidatesRevoked(opts *bind.FilterOpts) (*ValidatorCandidatesRevokedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidatesRevokedIterator{contract: _Validator.contract, event: "CandidatesRevoked", logs: logs, sub: sub}, nil
}

// WatchCandidatesRevoked is a free log subscription operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchCandidatesRevoked(opts *bind.WatchOpts, sink chan<- *ValidatorCandidatesRevoked) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidatesRevoked)
				if err := _Validator.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseCandidatesRevoked(log types.Log) (*ValidatorCandidatesRevoked, error) {
	event := new(ValidatorCandidatesRevoked)
	if err := _Validator.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCommissionRateUpdateScheduledIterator is returned from FilterCommissionRateUpdateScheduled and is used to iterate over the raw logs and unpacked data for CommissionRateUpdateScheduled events raised by the Validator contract.
type ValidatorCommissionRateUpdateScheduledIterator struct {
	Event *ValidatorCommissionRateUpdateScheduled // Event containing the contract specifics and raw log

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
func (it *ValidatorCommissionRateUpdateScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCommissionRateUpdateScheduled)
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
		it.Event = new(ValidatorCommissionRateUpdateScheduled)
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
func (it *ValidatorCommissionRateUpdateScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCommissionRateUpdateScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCommissionRateUpdateScheduled represents a CommissionRateUpdateScheduled event raised by the Validator contract.
type ValidatorCommissionRateUpdateScheduled struct {
	ConsensusAddr      common.Address
	EffectiveTimestamp *big.Int
	Rate               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdateScheduled is a free log retrieval operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_Validator *ValidatorFilterer) FilterCommissionRateUpdateScheduled(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCommissionRateUpdateScheduledIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCommissionRateUpdateScheduledIterator{contract: _Validator.contract, event: "CommissionRateUpdateScheduled", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdateScheduled is a free log subscription operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_Validator *ValidatorFilterer) WatchCommissionRateUpdateScheduled(opts *bind.WatchOpts, sink chan<- *ValidatorCommissionRateUpdateScheduled, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCommissionRateUpdateScheduled)
				if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseCommissionRateUpdateScheduled(log types.Log) (*ValidatorCommissionRateUpdateScheduled, error) {
	event := new(ValidatorCommissionRateUpdateScheduled)
	if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCommissionRateUpdatedIterator is returned from FilterCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for CommissionRateUpdated events raised by the Validator contract.
type ValidatorCommissionRateUpdatedIterator struct {
	Event *ValidatorCommissionRateUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCommissionRateUpdated)
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
		it.Event = new(ValidatorCommissionRateUpdated)
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
func (it *ValidatorCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCommissionRateUpdated represents a CommissionRateUpdated event raised by the Validator contract.
type ValidatorCommissionRateUpdated struct {
	ConsensusAddr common.Address
	Rate          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdated is a free log retrieval operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_Validator *ValidatorFilterer) FilterCommissionRateUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCommissionRateUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCommissionRateUpdatedIterator{contract: _Validator.contract, event: "CommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdated is a free log subscription operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_Validator *ValidatorFilterer) WatchCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCommissionRateUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCommissionRateUpdated)
				if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseCommissionRateUpdated(log types.Log) (*ValidatorCommissionRateUpdated, error) {
	event := new(ValidatorCommissionRateUpdated)
	if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorDeprecatedRewardRecycleFailedIterator is returned from FilterDeprecatedRewardRecycleFailed and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycleFailed events raised by the Validator contract.
type ValidatorDeprecatedRewardRecycleFailedIterator struct {
	Event *ValidatorDeprecatedRewardRecycleFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDeprecatedRewardRecycleFailed)
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
		it.Event = new(ValidatorDeprecatedRewardRecycleFailed)
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
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDeprecatedRewardRecycleFailed represents a DeprecatedRewardRecycleFailed event raised by the Validator contract.
type ValidatorDeprecatedRewardRecycleFailed struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycleFailed is a free log retrieval operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_Validator *ValidatorFilterer) FilterDeprecatedRewardRecycleFailed(opts *bind.FilterOpts, recipientAddr []common.Address) (*ValidatorDeprecatedRewardRecycleFailedIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDeprecatedRewardRecycleFailedIterator{contract: _Validator.contract, event: "DeprecatedRewardRecycleFailed", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycleFailed is a free log subscription operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_Validator *ValidatorFilterer) WatchDeprecatedRewardRecycleFailed(opts *bind.WatchOpts, sink chan<- *ValidatorDeprecatedRewardRecycleFailed, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDeprecatedRewardRecycleFailed)
				if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseDeprecatedRewardRecycleFailed(log types.Log) (*ValidatorDeprecatedRewardRecycleFailed, error) {
	event := new(ValidatorDeprecatedRewardRecycleFailed)
	if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorDeprecatedRewardRecycledIterator is returned from FilterDeprecatedRewardRecycled and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycled events raised by the Validator contract.
type ValidatorDeprecatedRewardRecycledIterator struct {
	Event *ValidatorDeprecatedRewardRecycled // Event containing the contract specifics and raw log

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
func (it *ValidatorDeprecatedRewardRecycledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDeprecatedRewardRecycled)
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
		it.Event = new(ValidatorDeprecatedRewardRecycled)
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
func (it *ValidatorDeprecatedRewardRecycledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDeprecatedRewardRecycledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDeprecatedRewardRecycled represents a DeprecatedRewardRecycled event raised by the Validator contract.
type ValidatorDeprecatedRewardRecycled struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycled is a free log retrieval operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) FilterDeprecatedRewardRecycled(opts *bind.FilterOpts, recipientAddr []common.Address) (*ValidatorDeprecatedRewardRecycledIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDeprecatedRewardRecycledIterator{contract: _Validator.contract, event: "DeprecatedRewardRecycled", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycled is a free log subscription operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) WatchDeprecatedRewardRecycled(opts *bind.WatchOpts, sink chan<- *ValidatorDeprecatedRewardRecycled, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDeprecatedRewardRecycled)
				if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseDeprecatedRewardRecycled(log types.Log) (*ValidatorDeprecatedRewardRecycled, error) {
	event := new(ValidatorDeprecatedRewardRecycled)
	if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedAmountUpdatedIterator is returned from FilterEmergencyExitLockedAmountUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedAmountUpdated events raised by the Validator contract.
type ValidatorEmergencyExitLockedAmountUpdatedIterator struct {
	Event *ValidatorEmergencyExitLockedAmountUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedAmountUpdated)
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
		it.Event = new(ValidatorEmergencyExitLockedAmountUpdated)
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
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedAmountUpdated represents a EmergencyExitLockedAmountUpdated event raised by the Validator contract.
type ValidatorEmergencyExitLockedAmountUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedAmountUpdated is a free log retrieval operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedAmountUpdated(opts *bind.FilterOpts) (*ValidatorEmergencyExitLockedAmountUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedAmountUpdatedIterator{contract: _Validator.contract, event: "EmergencyExitLockedAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedAmountUpdated is a free log subscription operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedAmountUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedAmountUpdated)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedAmountUpdated(log types.Log) (*ValidatorEmergencyExitLockedAmountUpdated, error) {
	event := new(ValidatorEmergencyExitLockedAmountUpdated)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedFundReleasedIterator is returned from FilterEmergencyExitLockedFundReleased and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleased events raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasedIterator struct {
	Event *ValidatorEmergencyExitLockedFundReleased // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedFundReleased)
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
		it.Event = new(ValidatorEmergencyExitLockedFundReleased)
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
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedFundReleased represents a EmergencyExitLockedFundReleased event raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleased struct {
	ConsensusAddr  common.Address
	Recipient      common.Address
	UnlockedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleased is a free log retrieval operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedFundReleased(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorEmergencyExitLockedFundReleasedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedFundReleasedIterator{contract: _Validator.contract, event: "EmergencyExitLockedFundReleased", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleased is a free log subscription operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedFundReleased(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedFundReleased, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedFundReleased)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedFundReleased(log types.Log) (*ValidatorEmergencyExitLockedFundReleased, error) {
	event := new(ValidatorEmergencyExitLockedFundReleased)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedFundReleasingFailedIterator is returned from FilterEmergencyExitLockedFundReleasingFailed and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleasingFailed events raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasingFailedIterator struct {
	Event *ValidatorEmergencyExitLockedFundReleasingFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedFundReleasingFailed)
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
		it.Event = new(ValidatorEmergencyExitLockedFundReleasingFailed)
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
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedFundReleasingFailed represents a EmergencyExitLockedFundReleasingFailed event raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasingFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	UnlockedAmount  *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleasingFailed is a free log retrieval operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedFundReleasingFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorEmergencyExitLockedFundReleasingFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedFundReleasingFailedIterator{contract: _Validator.contract, event: "EmergencyExitLockedFundReleasingFailed", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleasingFailed is a free log subscription operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedFundReleasingFailed(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedFundReleasingFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedFundReleasingFailed)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedFundReleasingFailed(log types.Log) (*ValidatorEmergencyExitLockedFundReleasingFailed, error) {
	event := new(ValidatorEmergencyExitLockedFundReleasingFailed)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitRequestedIterator is returned from FilterEmergencyExitRequested and is used to iterate over the raw logs and unpacked data for EmergencyExitRequested events raised by the Validator contract.
type ValidatorEmergencyExitRequestedIterator struct {
	Event *ValidatorEmergencyExitRequested // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitRequested)
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
		it.Event = new(ValidatorEmergencyExitRequested)
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
func (it *ValidatorEmergencyExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitRequested represents a EmergencyExitRequested event raised by the Validator contract.
type ValidatorEmergencyExitRequested struct {
	ConsensusAddr common.Address
	LockedAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitRequested is a free log retrieval operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitRequested(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorEmergencyExitRequestedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitRequestedIterator{contract: _Validator.contract, event: "EmergencyExitRequested", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitRequested is a free log subscription operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitRequested(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitRequested, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitRequested)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseEmergencyExitRequested(log types.Log) (*ValidatorEmergencyExitRequested, error) {
	event := new(ValidatorEmergencyExitRequested)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExpiryDurationUpdatedIterator is returned from FilterEmergencyExpiryDurationUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExpiryDurationUpdated events raised by the Validator contract.
type ValidatorEmergencyExpiryDurationUpdatedIterator struct {
	Event *ValidatorEmergencyExpiryDurationUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExpiryDurationUpdated)
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
		it.Event = new(ValidatorEmergencyExpiryDurationUpdated)
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
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExpiryDurationUpdated represents a EmergencyExpiryDurationUpdated event raised by the Validator contract.
type ValidatorEmergencyExpiryDurationUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExpiryDurationUpdated is a free log retrieval operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) FilterEmergencyExpiryDurationUpdated(opts *bind.FilterOpts) (*ValidatorEmergencyExpiryDurationUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExpiryDurationUpdatedIterator{contract: _Validator.contract, event: "EmergencyExpiryDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExpiryDurationUpdated is a free log subscription operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) WatchEmergencyExpiryDurationUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExpiryDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExpiryDurationUpdated)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseEmergencyExpiryDurationUpdated(log types.Log) (*ValidatorEmergencyExpiryDurationUpdated, error) {
	event := new(ValidatorEmergencyExpiryDurationUpdated)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Validator contract.
type ValidatorInitializedIterator struct {
	Event *ValidatorInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorInitialized)
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
		it.Event = new(ValidatorInitialized)
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
func (it *ValidatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorInitialized represents a Initialized event raised by the Validator contract.
type ValidatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Validator *ValidatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorInitializedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorInitializedIterator{contract: _Validator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Validator *ValidatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorInitialized)
				if err := _Validator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseInitialized(log types.Log) (*ValidatorInitialized, error) {
	event := new(ValidatorInitialized)
	if err := _Validator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaintenanceContractUpdatedIterator is returned from FilterMaintenanceContractUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceContractUpdated events raised by the Validator contract.
type ValidatorMaintenanceContractUpdatedIterator struct {
	Event *ValidatorMaintenanceContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaintenanceContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaintenanceContractUpdated)
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
		it.Event = new(ValidatorMaintenanceContractUpdated)
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
func (it *ValidatorMaintenanceContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaintenanceContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaintenanceContractUpdated represents a MaintenanceContractUpdated event raised by the Validator contract.
type ValidatorMaintenanceContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceContractUpdated is a free log retrieval operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterMaintenanceContractUpdated(opts *bind.FilterOpts) (*ValidatorMaintenanceContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaintenanceContractUpdatedIterator{contract: _Validator.contract, event: "MaintenanceContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceContractUpdated is a free log subscription operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchMaintenanceContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaintenanceContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaintenanceContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
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

// ParseMaintenanceContractUpdated is a log parse operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseMaintenanceContractUpdated(log types.Log) (*ValidatorMaintenanceContractUpdated, error) {
	event := new(ValidatorMaintenanceContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxPrioritizedValidatorNumberUpdatedIterator is returned from FilterMaxPrioritizedValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxPrioritizedValidatorNumberUpdated events raised by the Validator contract.
type ValidatorMaxPrioritizedValidatorNumberUpdatedIterator struct {
	Event *ValidatorMaxPrioritizedValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxPrioritizedValidatorNumberUpdated)
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
		it.Event = new(ValidatorMaxPrioritizedValidatorNumberUpdated)
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
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxPrioritizedValidatorNumberUpdated represents a MaxPrioritizedValidatorNumberUpdated event raised by the Validator contract.
type ValidatorMaxPrioritizedValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxPrioritizedValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) FilterMaxPrioritizedValidatorNumberUpdated(opts *bind.FilterOpts) (*ValidatorMaxPrioritizedValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxPrioritizedValidatorNumberUpdatedIterator{contract: _Validator.contract, event: "MaxPrioritizedValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPrioritizedValidatorNumberUpdated is a free log subscription operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) WatchMaxPrioritizedValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxPrioritizedValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxPrioritizedValidatorNumberUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMaxPrioritizedValidatorNumberUpdated(log types.Log) (*ValidatorMaxPrioritizedValidatorNumberUpdated, error) {
	event := new(ValidatorMaxPrioritizedValidatorNumberUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxValidatorCandidateUpdatedIterator is returned from FilterMaxValidatorCandidateUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorCandidateUpdated events raised by the Validator contract.
type ValidatorMaxValidatorCandidateUpdatedIterator struct {
	Event *ValidatorMaxValidatorCandidateUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxValidatorCandidateUpdated)
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
		it.Event = new(ValidatorMaxValidatorCandidateUpdated)
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
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxValidatorCandidateUpdated represents a MaxValidatorCandidateUpdated event raised by the Validator contract.
type ValidatorMaxValidatorCandidateUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorCandidateUpdated is a free log retrieval operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_Validator *ValidatorFilterer) FilterMaxValidatorCandidateUpdated(opts *bind.FilterOpts) (*ValidatorMaxValidatorCandidateUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxValidatorCandidateUpdatedIterator{contract: _Validator.contract, event: "MaxValidatorCandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorCandidateUpdated is a free log subscription operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_Validator *ValidatorFilterer) WatchMaxValidatorCandidateUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxValidatorCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxValidatorCandidateUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMaxValidatorCandidateUpdated(log types.Log) (*ValidatorMaxValidatorCandidateUpdated, error) {
	event := new(ValidatorMaxValidatorCandidateUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxValidatorNumberUpdatedIterator is returned from FilterMaxValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorNumberUpdated events raised by the Validator contract.
type ValidatorMaxValidatorNumberUpdatedIterator struct {
	Event *ValidatorMaxValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxValidatorNumberUpdated)
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
		it.Event = new(ValidatorMaxValidatorNumberUpdated)
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
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxValidatorNumberUpdated represents a MaxValidatorNumberUpdated event raised by the Validator contract.
type ValidatorMaxValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) FilterMaxValidatorNumberUpdated(opts *bind.FilterOpts) (*ValidatorMaxValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxValidatorNumberUpdatedIterator{contract: _Validator.contract, event: "MaxValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorNumberUpdated is a free log subscription operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) WatchMaxValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxValidatorNumberUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMaxValidatorNumberUpdated(log types.Log) (*ValidatorMaxValidatorNumberUpdated, error) {
	event := new(ValidatorMaxValidatorNumberUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMinEffectiveDaysOnwardsUpdatedIterator is returned from FilterMinEffectiveDaysOnwardsUpdated and is used to iterate over the raw logs and unpacked data for MinEffectiveDaysOnwardsUpdated events raised by the Validator contract.
type ValidatorMinEffectiveDaysOnwardsUpdatedIterator struct {
	Event *ValidatorMinEffectiveDaysOnwardsUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMinEffectiveDaysOnwardsUpdated)
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
		it.Event = new(ValidatorMinEffectiveDaysOnwardsUpdated)
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
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMinEffectiveDaysOnwardsUpdated represents a MinEffectiveDaysOnwardsUpdated event raised by the Validator contract.
type ValidatorMinEffectiveDaysOnwardsUpdated struct {
	NumOfDays *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinEffectiveDaysOnwardsUpdated is a free log retrieval operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_Validator *ValidatorFilterer) FilterMinEffectiveDaysOnwardsUpdated(opts *bind.FilterOpts) (*ValidatorMinEffectiveDaysOnwardsUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMinEffectiveDaysOnwardsUpdatedIterator{contract: _Validator.contract, event: "MinEffectiveDaysOnwardsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinEffectiveDaysOnwardsUpdated is a free log subscription operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_Validator *ValidatorFilterer) WatchMinEffectiveDaysOnwardsUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMinEffectiveDaysOnwardsUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMinEffectiveDaysOnwardsUpdated)
				if err := _Validator.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMinEffectiveDaysOnwardsUpdated(log types.Log) (*ValidatorMinEffectiveDaysOnwardsUpdated, error) {
	event := new(ValidatorMinEffectiveDaysOnwardsUpdated)
	if err := _Validator.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMiningRewardDistributedIterator is returned from FilterMiningRewardDistributed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributed events raised by the Validator contract.
type ValidatorMiningRewardDistributedIterator struct {
	Event *ValidatorMiningRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorMiningRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMiningRewardDistributed)
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
		it.Event = new(ValidatorMiningRewardDistributed)
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
func (it *ValidatorMiningRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMiningRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMiningRewardDistributed represents a MiningRewardDistributed event raised by the Validator contract.
type ValidatorMiningRewardDistributed struct {
	ConsensusAddr common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributed is a free log retrieval operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_Validator *ValidatorFilterer) FilterMiningRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorMiningRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMiningRewardDistributedIterator{contract: _Validator.contract, event: "MiningRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributed is a free log subscription operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_Validator *ValidatorFilterer) WatchMiningRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorMiningRewardDistributed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMiningRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMiningRewardDistributed(log types.Log) (*ValidatorMiningRewardDistributed, error) {
	event := new(ValidatorMiningRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMiningRewardDistributionFailedIterator is returned from FilterMiningRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributionFailed events raised by the Validator contract.
type ValidatorMiningRewardDistributionFailedIterator struct {
	Event *ValidatorMiningRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorMiningRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMiningRewardDistributionFailed)
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
		it.Event = new(ValidatorMiningRewardDistributionFailed)
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
func (it *ValidatorMiningRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMiningRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMiningRewardDistributionFailed represents a MiningRewardDistributionFailed event raised by the Validator contract.
type ValidatorMiningRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributionFailed is a free log retrieval operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterMiningRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorMiningRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMiningRewardDistributionFailedIterator{contract: _Validator.contract, event: "MiningRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributionFailed is a free log subscription operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchMiningRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorMiningRewardDistributionFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMiningRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseMiningRewardDistributionFailed(log types.Log) (*ValidatorMiningRewardDistributionFailed, error) {
	event := new(ValidatorMiningRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRoninTrustedOrganizationContractUpdatedIterator is returned from FilterRoninTrustedOrganizationContractUpdated and is used to iterate over the raw logs and unpacked data for RoninTrustedOrganizationContractUpdated events raised by the Validator contract.
type ValidatorRoninTrustedOrganizationContractUpdatedIterator struct {
	Event *ValidatorRoninTrustedOrganizationContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRoninTrustedOrganizationContractUpdated)
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
		it.Event = new(ValidatorRoninTrustedOrganizationContractUpdated)
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
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRoninTrustedOrganizationContractUpdated represents a RoninTrustedOrganizationContractUpdated event raised by the Validator contract.
type ValidatorRoninTrustedOrganizationContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninTrustedOrganizationContractUpdated is a free log retrieval operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterRoninTrustedOrganizationContractUpdated(opts *bind.FilterOpts) (*ValidatorRoninTrustedOrganizationContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorRoninTrustedOrganizationContractUpdatedIterator{contract: _Validator.contract, event: "RoninTrustedOrganizationContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninTrustedOrganizationContractUpdated is a free log subscription operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchRoninTrustedOrganizationContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorRoninTrustedOrganizationContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRoninTrustedOrganizationContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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

// ParseRoninTrustedOrganizationContractUpdated is a log parse operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseRoninTrustedOrganizationContractUpdated(log types.Log) (*ValidatorRoninTrustedOrganizationContractUpdated, error) {
	event := new(ValidatorRoninTrustedOrganizationContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSlashIndicatorContractUpdatedIterator is returned from FilterSlashIndicatorContractUpdated and is used to iterate over the raw logs and unpacked data for SlashIndicatorContractUpdated events raised by the Validator contract.
type ValidatorSlashIndicatorContractUpdatedIterator struct {
	Event *ValidatorSlashIndicatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSlashIndicatorContractUpdated)
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
		it.Event = new(ValidatorSlashIndicatorContractUpdated)
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
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSlashIndicatorContractUpdated represents a SlashIndicatorContractUpdated event raised by the Validator contract.
type ValidatorSlashIndicatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSlashIndicatorContractUpdated is a free log retrieval operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterSlashIndicatorContractUpdated(opts *bind.FilterOpts) (*ValidatorSlashIndicatorContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorSlashIndicatorContractUpdatedIterator{contract: _Validator.contract, event: "SlashIndicatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashIndicatorContractUpdated is a free log subscription operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchSlashIndicatorContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorSlashIndicatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSlashIndicatorContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
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

// ParseSlashIndicatorContractUpdated is a log parse operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseSlashIndicatorContractUpdated(log types.Log) (*ValidatorSlashIndicatorContractUpdated, error) {
	event := new(ValidatorSlashIndicatorContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingContractUpdatedIterator is returned from FilterStakingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingContractUpdated events raised by the Validator contract.
type ValidatorStakingContractUpdatedIterator struct {
	Event *ValidatorStakingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingContractUpdated)
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
		it.Event = new(ValidatorStakingContractUpdated)
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
func (it *ValidatorStakingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingContractUpdated represents a StakingContractUpdated event raised by the Validator contract.
type ValidatorStakingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingContractUpdated is a free log retrieval operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterStakingContractUpdated(opts *bind.FilterOpts) (*ValidatorStakingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingContractUpdatedIterator{contract: _Validator.contract, event: "StakingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingContractUpdated is a free log subscription operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchStakingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
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

// ParseStakingContractUpdated is a log parse operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseStakingContractUpdated(log types.Log) (*ValidatorStakingContractUpdated, error) {
	event := new(ValidatorStakingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingRewardDistributedIterator is returned from FilterStakingRewardDistributed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributed events raised by the Validator contract.
type ValidatorStakingRewardDistributedIterator struct {
	Event *ValidatorStakingRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingRewardDistributed)
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
		it.Event = new(ValidatorStakingRewardDistributed)
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
func (it *ValidatorStakingRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingRewardDistributed represents a StakingRewardDistributed event raised by the Validator contract.
type ValidatorStakingRewardDistributed struct {
	TotalAmount    *big.Int
	ConsensusAddrs []common.Address
	Amounts        []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributed is a free log retrieval operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributedIterator{contract: _Validator.contract, event: "StakingRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributed is a free log subscription operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_Validator *ValidatorFilterer) WatchStakingRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorStakingRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseStakingRewardDistributed(log types.Log) (*ValidatorStakingRewardDistributed, error) {
	event := new(ValidatorStakingRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingRewardDistributionFailedIterator is returned from FilterStakingRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributionFailed events raised by the Validator contract.
type ValidatorStakingRewardDistributionFailedIterator struct {
	Event *ValidatorStakingRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingRewardDistributionFailed)
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
		it.Event = new(ValidatorStakingRewardDistributionFailed)
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
func (it *ValidatorStakingRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingRewardDistributionFailed represents a StakingRewardDistributionFailed event raised by the Validator contract.
type ValidatorStakingRewardDistributionFailed struct {
	TotalAmount     *big.Int
	ConsensusAddrs  []common.Address
	Amounts         []*big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributionFailed is a free log retrieval operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributionFailed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributionFailedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributionFailedIterator{contract: _Validator.contract, event: "StakingRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributionFailed is a free log subscription operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchStakingRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorStakingRewardDistributionFailed) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseStakingRewardDistributionFailed(log types.Log) (*ValidatorStakingRewardDistributionFailed, error) {
	event := new(ValidatorStakingRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingVestingContractUpdatedIterator is returned from FilterStakingVestingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingVestingContractUpdated events raised by the Validator contract.
type ValidatorStakingVestingContractUpdatedIterator struct {
	Event *ValidatorStakingVestingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingVestingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingVestingContractUpdated)
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
		it.Event = new(ValidatorStakingVestingContractUpdated)
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
func (it *ValidatorStakingVestingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingVestingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingVestingContractUpdated represents a StakingVestingContractUpdated event raised by the Validator contract.
type ValidatorStakingVestingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingVestingContractUpdated is a free log retrieval operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterStakingVestingContractUpdated(opts *bind.FilterOpts) (*ValidatorStakingVestingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingVestingContractUpdatedIterator{contract: _Validator.contract, event: "StakingVestingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingVestingContractUpdated is a free log subscription operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchStakingVestingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingVestingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingVestingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
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

// ParseStakingVestingContractUpdated is a log parse operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseStakingVestingContractUpdated(log types.Log) (*ValidatorStakingVestingContractUpdated, error) {
	event := new(ValidatorStakingVestingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorPunishedIterator is returned from FilterValidatorPunished and is used to iterate over the raw logs and unpacked data for ValidatorPunished events raised by the Validator contract.
type ValidatorValidatorPunishedIterator struct {
	Event *ValidatorValidatorPunished // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorPunished)
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
		it.Event = new(ValidatorValidatorPunished)
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
func (it *ValidatorValidatorPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorPunished represents a ValidatorPunished event raised by the Validator contract.
type ValidatorValidatorPunished struct {
	ConsensusAddr                  common.Address
	Period                         *big.Int
	JailedUntil                    *big.Int
	DeductedStakingAmount          *big.Int
	BlockProducerRewardDeprecated  bool
	BridgeOperatorRewardDeprecated bool
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) FilterValidatorPunished(opts *bind.FilterOpts, consensusAddr []common.Address, period []*big.Int) (*ValidatorValidatorPunishedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorPunishedIterator{contract: _Validator.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorPunished, consensusAddr []common.Address, period []*big.Int) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorPunished)
				if err := _Validator.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
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
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) ParseValidatorPunished(log types.Log) (*ValidatorValidatorPunished, error) {
	event := new(ValidatorValidatorPunished)
	if err := _Validator.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Validator contract.
type ValidatorValidatorSetUpdatedIterator struct {
	Event *ValidatorValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorSetUpdated)
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
		it.Event = new(ValidatorValidatorSetUpdated)
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
func (it *ValidatorValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Validator contract.
type ValidatorValidatorSetUpdated struct {
	Period         *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*ValidatorValidatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorSetUpdatedIterator{contract: _Validator.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseValidatorSetUpdated(log types.Log) (*ValidatorValidatorSetUpdated, error) {
	event := new(ValidatorValidatorSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the Validator contract.
type ValidatorValidatorUnjailedIterator struct {
	Event *ValidatorValidatorUnjailed // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorUnjailed)
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
		it.Event = new(ValidatorValidatorUnjailed)
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
func (it *ValidatorValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorUnjailed represents a ValidatorUnjailed event raised by the Validator contract.
type ValidatorValidatorUnjailed struct {
	Validator common.Address
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorValidatorUnjailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorUnjailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorUnjailedIterator{contract: _Validator.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorUnjailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorUnjailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorUnjailed)
				if err := _Validator.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
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
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) ParseValidatorUnjailed(log types.Log) (*ValidatorValidatorUnjailed, error) {
	event := new(ValidatorValidatorUnjailed)
	if err := _Validator.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWrappedUpEpochIterator is returned from FilterWrappedUpEpoch and is used to iterate over the raw logs and unpacked data for WrappedUpEpoch events raised by the Validator contract.
type ValidatorWrappedUpEpochIterator struct {
	Event *ValidatorWrappedUpEpoch // Event containing the contract specifics and raw log

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
func (it *ValidatorWrappedUpEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWrappedUpEpoch)
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
		it.Event = new(ValidatorWrappedUpEpoch)
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
func (it *ValidatorWrappedUpEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWrappedUpEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWrappedUpEpoch represents a WrappedUpEpoch event raised by the Validator contract.
type ValidatorWrappedUpEpoch struct {
	PeriodNumber *big.Int
	EpochNumber  *big.Int
	PeriodEnding bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWrappedUpEpoch is a free log retrieval operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_Validator *ValidatorFilterer) FilterWrappedUpEpoch(opts *bind.FilterOpts, periodNumber []*big.Int, epochNumber []*big.Int) (*ValidatorWrappedUpEpochIterator, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWrappedUpEpochIterator{contract: _Validator.contract, event: "WrappedUpEpoch", logs: logs, sub: sub}, nil
}

// WatchWrappedUpEpoch is a free log subscription operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_Validator *ValidatorFilterer) WatchWrappedUpEpoch(opts *bind.WatchOpts, sink chan<- *ValidatorWrappedUpEpoch, periodNumber []*big.Int, epochNumber []*big.Int) (event.Subscription, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWrappedUpEpoch)
				if err := _Validator.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseWrappedUpEpoch(log types.Log) (*ValidatorWrappedUpEpoch, error) {
	event := new(ValidatorWrappedUpEpoch)
	if err := _Validator.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
