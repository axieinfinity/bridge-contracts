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

// ICandidateManagerValidatorCandidate is an auto generated low-level Go binding around an user-defined struct.
type ICandidateManagerValidatorCandidate struct {
	Admin              common.Address
	ConsensusAddr      common.Address
	TreasuryAddr       common.Address
	BridgeOperatorAddr common.Address
	CommissionRate     *big.Int
	RevokedTimestamp   *big.Int
	ExtraData          []byte
}

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"BlockProducerSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIRoninValidatorSet.BlockRewardDeprecatedType\",\"name\":\"deprecatedType\",\"type\":\"uint8\"}],\"name\":\"BlockRewardDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbaseAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submittedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bonusAmount\",\"type\":\"uint256\"}],\"name\":\"BlockRewardSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"BridgeOperatorRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeTrackingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"CandidateGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedTimestamp\",\"type\":\"uint256\"}],\"name\":\"CandidateRevokedTimestampUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"CandidatesRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"MaintenanceContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorCandidateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"MaxValidatorNumberUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"MiningRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"SlashIndicatorContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contractBalance\",\"type\":\"uint256\"}],\"name\":\"StakingRewardDistributionFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakingVestingContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jailedUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blockProducerRewardDeprecated\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"bridgeOperatorRewardDeprecated\",\"type\":\"bool\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"consensusAddrs\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"ValidatorUnjailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"periodNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"periodEnding\",\"type\":\"bool\"}],\"name\":\"WrappedUpEpoch\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"_isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTrackingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"bulkJailed\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodStartAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_block\",\"type\":\"uint256\"}],\"name\":\"epochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"execBailOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newJailedUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"execSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockProducers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_bridgeOperatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revokedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structICandidateManager.ValidatorCandidate[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionRate\",\"type\":\"uint256\"}],\"name\":\"grantValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__stakingVestingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__maintenanceContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__bridgeTrackingContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBlockProducer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\"}],\"name\":\"isBridgeOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_candidate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPeriodEnding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"jailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"jailedAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"jailedTimeLeft\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"jailedTimeLeftAtBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isJailed_\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"blockLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"}],\"name\":\"miningRewardDeprecated\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_blockProducers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"miningRewardDeprecatedAtPeriod\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"_result\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompilePickValidatorSetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"precompileSortValidatorsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_secsLeft\",\"type\":\"uint256\"}],\"name\":\"requestRevokeCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeTrackingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxPrioritizedValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setSlashIndicatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setStakingVestingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashIndicatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingVestingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBlockProducers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBridgeOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000e4565b603c54610100900460ff16156200008f5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b603c5460ff9081161015620000e257603c805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b614dfe80620000f46000396000f3fe60806040526004361061036f5760003560e01c80637043e5dd116101c65780639e94b9ec116100f7578063ba77b06c11610095578063d2cb215e1161006f578063d2cb215e146109b0578063ee99205c146109ce578063eeb629a8146109ec578063facd743b14610a015761037e565b8063ba77b06c14610966578063c94aaa021461097b578063d09f1ab41461099b5761037e565b8063ad295783116100d1578063ad295783146108f1578063b405aaf214610911578063b5e337de14610931578063b7ab4db5146109515761037e565b80639e94b9ec14610884578063a0c3f2d214610899578063a3d545f5146108d15761037e565b806387c891bd116101645780639b19dbfd1161013e5780639b19dbfd1461080d5780639b8c334b146108225780639c8d98da146108445780639dd373b9146108645761037e565b806387c891bd146107c45780638d559c38146107d957806392a8c2e8146107ed5761037e565b80637593ff71116101a05780637593ff711461072757806381f9535f14610747578063823a7b9c1461078457806385ad5aec146107a45761037e565b80637043e5dd146106df57806372e46810146106ff578063733ec970146107075761037e565b806346fe9311116102a0578063562d53041161023e578063605239a111610218578063605239a11461067557806365244ece1461068a5780636aa1c2ef146106aa5780636efa12bd146106bf5761037e565b8063562d530414610622578063570ccb13146106375780635a08482d146106575761037e565b80634f2a693f1161027a5780634f2a693f146105ba57806352091f17146105da5780635248184a146105e25780635511cde1146106045761037e565b806346fe93111461055857806349096d26146105785780634a68f8c61461059a5761037e565b8063297a8fca1161030d5780633986de6a116102e75780633986de6a146104d95780633b3159b6146104f9578063428483c31461050d5780634493421e1461053a5761037e565b8063297a8fca146104725780632bcf3d15146104875780633529214b146104a75761037e565b806315b5ebde1161034957806315b5ebde146103f0578063217f35c2146104105780632607d9191461042557806328bde1e1146104455761037e565b806304d971ab1461038657806306040618146103bb5780630f43a677146103da5761037e565b3661037e5761037c610a21565b005b61037c610a21565b34801561039257600080fd5b506103a66103a1366004614194565b610ab3565b60405190151581526020015b60405180910390f35b3480156103c757600080fd5b506040545b6040519081526020016103b2565b3480156103e657600080fd5b506103cc60425481565b3480156103fc57600080fd5b5061037c61040b3660046141cd565b610ada565b34801561041c57600080fd5b506103a6610bc5565b34801561043157600080fd5b506103a66104403660046141cd565b610bd8565b34801561045157600080fd5b506104656104603660046141f9565b610bfc565b6040516103b291906142b8565b34801561047e57600080fd5b506041546103cc565b34801561049357600080fd5b5061037c6104a23660046141f9565b610d3f565b3480156104b357600080fd5b506001546001600160a01b03165b6040516001600160a01b0390911681526020016103b2565b3480156104e557600080fd5b5061037c6104f43660046142cb565b610df4565b34801561050557600080fd5b5060686104c1565b34801561051957600080fd5b5061052d6105283660046143c0565b610f5d565b6040516103b29190614401565b34801561054657600080fd5b506005546001600160a01b03166104c1565b34801561056457600080fd5b5061037c6105733660046141f9565b611018565b34801561058457600080fd5b5061058d6110c7565b6040516103b2919061448b565b3480156105a657600080fd5b5061052d6105b53660046143c0565b6111b0565b3480156105c657600080fd5b5061037c6105d536600461449e565b61127a565b61037c6112bb565b3480156105ee57600080fd5b506105f7611624565b6040516103b291906144b7565b34801561061057600080fd5b506004546001600160a01b03166104c1565b34801561062e57600080fd5b506042546103cc565b34801561064357600080fd5b5061037c610652366004614519565b6117e0565b34801561066357600080fd5b506002546001600160a01b03166104c1565b34801561068157600080fd5b506006546103cc565b34801561069657600080fd5b506103a66106a53660046141f9565b61196d565b3480156106b657600080fd5b50603e546103cc565b3480156106cb57600080fd5b5061037c6106da3660046141cd565b6119a7565b3480156106eb57600080fd5b506103a66106fa3660046141f9565b611aff565b61037c611b0b565b34801561071357600080fd5b5061037c61072236600461454e565b611d56565b34801561073357600080fd5b506103a661074236600461449e565b612059565b34801561075357600080fd5b506107676107623660046141f9565b61207e565b6040805193151584526020840192909252908201526060016103b2565b34801561079057600080fd5b5061037c61079f36600461449e565b61209a565b3480156107b057600080fd5b506107676107bf3660046141cd565b6120db565b3480156107d057600080fd5b50603f546103cc565b3480156107e557600080fd5b5060666104c1565b3480156107f957600080fd5b5061052d6108083660046145b2565b61215e565b34801561081957600080fd5b5061058d61221b565b34801561082e57600080fd5b506103a661083d36600461449e565b6040541090565b34801561085057600080fd5b5061037c61085f3660046141f9565b6122db565b34801561087057600080fd5b5061037c61087f3660046141f9565b61238d565b34801561089057600080fd5b506103cc612438565b3480156108a557600080fd5b506103a66108b43660046141f9565b6001600160a01b0316600090815260086020526040902054151590565b3480156108dd57600080fd5b506103cc6108ec36600461449e565b61248c565b3480156108fd57600080fd5b5061037c61090c3660046141f9565b6124a7565b34801561091d57600080fd5b506103a661092c3660046141f9565b612559565b34801561093d57600080fd5b5061037c61094c3660046141f9565b6125bd565b34801561095d57600080fd5b5061058d61267e565b34801561097257600080fd5b5061058d61272b565b34801561098757600080fd5b5061037c61099636600461449e565b61278d565b3480156109a757600080fd5b50603d546103cc565b3480156109bc57600080fd5b506003546001600160a01b03166104c1565b3480156109da57600080fd5b506000546001600160a01b03166104c1565b3480156109f857600080fd5b506045546103cc565b348015610a0d57600080fd5b506103a6610a1c3660046141f9565b6127ce565b6001546001600160a01b03163314610ab15760405162461bcd60e51b815260206004820152604260248201527f526f6e696e56616c696461746f725365743a206f6e6c7920726563656976657360448201527f20524f4e2066726f6d207374616b696e672076657374696e6720636f6e74726160648201526118dd60f21b608482015260a4015b60405180910390fd5b565b6001600160a01b038281166000908152600960205260409020548116908216145b92915050565b33610aed6002546001600160a01b031690565b6001600160a01b031614610b135760405162461bcd60e51b8152600401610aa8906145fd565b6001600160a01b038216600081815260476020908152604080832085845282528083208054600160ff199182168117909255948452604683528184208685529092529091208054909216909155610b6a9043614682565b6001600160a01b038316600081815260496020526040908190209290925590517f6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e90610bb99084815260200190565b60405180910390a25050565b6000610bd361083d4261280b565b905090565b6001600160a01b0382166000908152604960205260408120548211155b9392505050565b610c046140d4565b6001600160a01b038216600090815260086020526040902054610c395760405162461bcd60e51b8152600401610aa890614695565b6001600160a01b03808316600090815260096020908152604091829020825160e081018452815485168152600182015485169281019290925260028101548416928201929092526003820154909216606083015260048101546080830152600581015460a083015260068101805460c084019190610cb6906146e7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce2906146e7565b8015610d2f5780601f10610d0457610100808354040283529160200191610d2f565b820191906000526020600020905b815481529060010190602001808311610d1257829003601f168201915b5050505050815250509050919050565b610d47612818565b6001600160a01b0316336001600160a01b031614610d775760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b11610de85760405162461bcd60e51b815260206004820152602e60248201527f486173536c617368496e64696361746f72436f6e74726163743a20736574207460448201526d1bc81b9bdb8b58dbdb9d1c9858dd60921b6064820152608401610aa8565b610df181612846565b50565b603c54610100900460ff1615808015610e145750603c54600160ff909116105b80610e2e5750303b158015610e2e5750603c5460ff166001145b610e915760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610aa8565b603c805460ff191660011790558015610eb457603c805461ff0019166101001790555b610ebd8b612846565b610ec68a61289b565b610ecf896128e9565b610ed888612937565b610ee186612985565b610eea876129d3565b610ef385612a21565b610efc84612a56565b610f0583612a8b565b603e8290558015610f5057603c805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050505050505050565b6060816001600160401b03811115610f7757610f7761475d565b604051908082528060200260200182016040528015610fa0578160200160208202803683370190505b50905060005b8281101561101157610fdd848483818110610fc357610fc3614773565b9050602002016020810190610fd891906141f9565b612b5e565b828281518110610fef57610fef614773565b911515602092830291909101909101528061100981614789565b915050610fa6565b5092915050565b611020612818565b6001600160a01b0316336001600160a01b0316146110505760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b116110be5760405162461bcd60e51b815260206004820152602b60248201527f4861734d61696e74656e616e6365436f6e74726163743a2073657420746f206e60448201526a1bdb8b58dbdb9d1c9858dd60aa1b6064820152608401610aa8565b610df181612937565b60606042546001600160401b038111156110e3576110e361475d565b60405190808252806020026020018201604052801561110c578160200160208202803683370190505b5090506000805b82518110156111aa5760008181526043602052604090205461113d906001600160a01b031661196d565b15611198576000818152604360205260409020546001600160a01b0316838361116581614789565b94508151811061117757611177614773565b60200260200101906001600160a01b031690816001600160a01b0316815250505b806111a281614789565b915050611113565b50815290565b6060816001600160401b038111156111ca576111ca61475d565b6040519080825280602002602001820160405280156111f3578160200160208202803683370190505b509050600061120160405490565b905060005b838110156112725761123e85858381811061122357611223614773565b905060200201602081019061123891906141f9565b83612b7f565b83828151811061125057611250614773565b911515602092830291909101909101528061126a81614789565b915050611206565b505092915050565b611282612818565b6001600160a01b0316336001600160a01b0316146112b25760405162461bcd60e51b8152600401610aa89061471b565b610df181612a56565b3341146112da5760405162461bcd60e51b8152600401610aa8906147a2565b343360006112e78261196d565b80156112f957506112f782612b5e565b155b801561131457506113128261130d60405490565b612b7f565b155b60018054604051630634f5b960e01b8152831515600482015260248101839052929350909160009182916001600160a01b0390911690630634f5b9906044016060604051808303816000875af1158015611372573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113969190614808565b925092505080604c60008282546113ad919061483d565b9091555084905061140457846001600160a01b03167f4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e8760016040516113f4929190614866565b60405180910390a2505050505050565b60408051878152602081018490526001600160a01b038716917f0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1910160405180910390a2600061145360405490565b90506000611461848961483d565b6001600160a01b03881660009081526047602090815260408083208684529091528120549192509060ff1615611565576002546040805163631c8fd160e11b815290516000926001600160a01b03169163c6391fa29160048083019260809291908290030181865afa1580156114db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ff9190614898565b9350505050612710818461151391906148ce565b61151d91906148fb565b9150886001600160a01b03167f4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e83600260405161155b929190614866565b60405180910390a2505b61156f8183614682565b6001600160a01b03891660009081526009602052604081206004015491935061271061159b85846148ce565b6115a591906148fb565b6001600160a01b038b166000908152604a60205260408120805492935083929091906115d290849061483d565b90915550600090506115e48286614682565b6001600160a01b038c166000908152604b602052604081208054929350839290919061161190849061483d565b9091555050505050505050505050505050565b6007546060906001600160401b038111156116415761164161475d565b60405190808252806020026020018201604052801561167a57816020015b6116676140d4565b81526020019060019003908161165f5790505b50905060005b81518110156117dc5760096000600783815481106116a0576116a0614773565b60009182526020808320909101546001600160a01b039081168452838201949094526040928301909120825160e081018452815485168152600182015485169281019290925260028101548416928201929092526003820154909216606083015260048101546080830152600581015460a083015260068101805460c08401919061172a906146e7565b80601f0160208091040260200160405190810160405280929190818152602001828054611756906146e7565b80156117a35780601f10611778576101008083540402835291602001916117a3565b820191906000526020600020905b81548152906001019060200180831161178657829003601f168201915b5050505050815250508282815181106117be576117be614773565b602002602001018190525080806117d490614789565b915050611680565b5090565b336117f36002546001600160a01b031690565b6001600160a01b0316146118195760405162461bcd60e51b8152600401610aa8906145fd565b600061182460405490565b6001600160a01b03851660008181526046602090815260408083208584528252808320805460ff19166001179055928252604a8152828220829055604b8152828220829055604990522054909150831115611895576001600160a01b03841660009081526049602052604090208390555b81156119025760005460405163c905bb3560e01b81526001600160a01b038681166004830152602482018590529091169063c905bb3590604401600060405180830381600087803b1580156118e957600080fd5b505af11580156118fd573d6000803e3d6000fd5b505050505b6001600160a01b038416600081815260496020908152604080832054815190815291820186905260019082015260608101919091528291907f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a9060800160405180910390a350505050565b6001600160a01b038116600090815260446020526040812054610ad49060019060ff1660038111156119a1576119a1614850565b90612baa565b336119ba6000546001600160a01b031690565b6001600160a01b0316146119e05760405162461bcd60e51b8152600401610aa89061490f565b6001600160a01b038216600090815260086020526040902054611a155760405162461bcd60e51b8152600401610aa890614695565b6000611a21824261483d565b6001600160a01b0384166000908152600960205260409020600501549091508110611aa25760405162461bcd60e51b815260206004820152602b60248201527f43616e6469646174654d616e616765723a20696e76616c6964207265766f6b6560448201526a0642074696d657374616d760ac1b6064820152608401610aa8565b6001600160a01b03831660008181526009602052604090819020600501839055517fdb451f2c533d44367eeca766bcee562bfc473b8d3f5d34f7b87befe026434aaa90611af29084815260200190565b60405180910390a2505050565b6000610ad48243610bd8565b334114611b2a5760405162461bcd60e51b8152600401610aa8906147a2565b611b3343612059565b611b9b5760405162461bcd60e51b815260206004820152603360248201527f526f6e696e56616c696461746f725365743a206f6e6c7920616c6c6f776564206044820152720c2e840e8d0ca40cadcc840decc40cae0dec6d606b1b6064820152608401610aa8565b611ba44361248c565b611baf603f5461248c565b10611c1a5760405162461bcd60e51b815260206004820152603560248201527f526f6e696e56616c696461746f725365743a20717565727920666f7220616c726044820152740cac2c8f240eee4c2e0e0cac840eae040cae0dec6d605b1b6064820152608401610aa8565b43603f556000611c294261280b565b90506000611c38826040541090565b9050611c4543600161483d565b6041556000611c5261267e565b90506000611c5f4361248c565b90506000611c6c60405490565b90508315611d0357600080611c818386612bdd565b91509150611c9183868484612fd7565b60025460405163129fccc160e01b81526001600160a01b039091169063129fccc190611cc3908890879060040161496c565b600060405180830381600087803b158015611cdd57600080fd5b505af1158015611cf1573d6000803e3d6000fd5b50505050611cfe876130da565b945050505b611d0d858461329c565b81817f0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce786604051611d42911515815260200190565b60405180910390a350505060409190915550565b33611d696000546001600160a01b031690565b6001600160a01b031614611d8f5760405162461bcd60e51b8152600401610aa89061490f565b6007546006548110611e025760405162461bcd60e51b815260206004820152603660248201527f43616e6469646174654d616e616765723a2065786365656473206d6178696d756044820152756d206e756d626572206f662063616e6469646174657360501b6064820152608401610aa8565b6001600160a01b03851660009081526008602052604090205415611e875760405162461bcd60e51b815260206004820152603660248201527f43616e6469646174654d616e616765723a20717565727920666f7220616c7265604482015275616479206578697374656e742063616e64696461746560501b6064820152608401610aa8565b612710821115611eea5760405162461bcd60e51b815260206004820152602860248201527f43616e6469646174654d616e616765723a20696e76616c696420636f6d697373604482015267696f6e207261746560c01b6064820152608401610aa8565b6001600160a01b038581166000818152600860209081526040808320861990556007805460018082019092557fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b03199081168717909155825160e0810184528d881681528085018781528c89168286019081528c8a1660608401908152608084018d815260001960a0860190815288518b8152808b018a5260c087019081529b8b52600990995296909820835181548616908c16178155915194820180548516958b169590951790945592516002840180548416918a1691909117905594516003830180549092169716969096179095555160048501555160058401559051909190600682019061200590826149d4565b50506040516001600160a01b038581168252808916925086811691908816907fd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd9060200160405180910390a4505050505050565b60006001603e5461206a9190614682565b603e546120779084614a93565b1492915050565b600080600061208d84436120db565b9250925092509193909250565b6120a2612818565b6001600160a01b0316336001600160a01b0316146120d25760405162461bcd60e51b8152600401610aa89061471b565b610df181612a21565b6001600160a01b038216600090815260496020526040812054819081908481101561211157600080600093509350935050612157565b6001935061211f8582614682565b61212a90600161483d565b92506121358561248c565b61213e8261248c565b6121489190614682565b61215390600161483d565b9150505b9250925092565b6060826001600160401b038111156121785761217861475d565b6040519080825280602002602001820160405280156121a1578160200160208202803683370190505b50905060005b83811015612213576121df8585838181106121c4576121c4614773565b90506020020160208101906121d991906141f9565b84612b7f565b8282815181106121f1576121f1614773565b911515602092830291909101909101528061220b81614789565b9150506121a7565b509392505050565b60606042546001600160401b038111156122375761223761475d565b604051908082528060200260200182016040528015612260578160200160208202803683370190505b50905060005b81518110156117dc576000818152604360209081526040808320546001600160a01b03908116845260099092529091206003015483519116908390839081106122b1576122b1614773565b6001600160a01b0390921660209283029190910190910152806122d381614789565b915050612266565b6122e3612818565b6001600160a01b0316336001600160a01b0316146123135760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b116123845760405162461bcd60e51b815260206004820152602e60248201527f486173427269646765547261636b696e67436f6e74726163743a20736574207460448201526d1bc81b9bdb8b58dbdb9d1c9858dd60921b6064820152608401610aa8565b610df181612985565b612395612818565b6001600160a01b0316336001600160a01b0316146123c55760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b1161242f5760405162461bcd60e51b815260206004820152602760248201527f4861735374616b696e67436f6e74726163743a2073657420746f206e6f6e2d636044820152661bdb9d1c9858dd60ca1b6064820152608401610aa8565b610df18161289b565b6000805b6042548110156117dc57600081815260436020526040902054612467906001600160a01b031661196d565b1561247a578161247681614789565b9250505b8061248481614789565b91505061243c565b6000603e548261249c91906148fb565b610ad490600161483d565b6124af612818565b6001600160a01b0316336001600160a01b0316146124df5760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b116125505760405162461bcd60e51b815260206004820152602e60248201527f4861735374616b696e6756657374696e67436f6e74726163743a20736574207460448201526d1bc81b9bdb8b58dbdb9d1c9858dd60921b6064820152608401610aa8565b610df1816128e9565b6000805b6042548110156125b7576000818152604360209081526040808320546001600160a01b0390811684526009909252909120600301548185169116036125a557600191506125b7565b806125af81614789565b91505061255d565b50919050565b6125c5612818565b6001600160a01b0316336001600160a01b0316146125f55760405162461bcd60e51b8152600401610aa89061471b565b6000816001600160a01b03163b116126755760405162461bcd60e51b815260206004820152603860248201527f486173526f6e696e547275737465644f7267616e697a6174696f6e436f6e747260448201527f6163743a2073657420746f206e6f6e2d636f6e747261637400000000000000006064820152608401610aa8565b610df1816129d3565b60606042546001600160401b0381111561269a5761269a61475d565b6040519080825280602002602001820160405280156126c3578160200160208202803683370190505b50905060005b81518110156117dc5760008181526043602052604090205482516001600160a01b039091169083908390811061270157612701614773565b6001600160a01b03909216602092830291909101909101528061272381614789565b9150506126c9565b6060600780548060200260200160405190810160405280929190818152602001828054801561278357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612765575b5050505050905090565b612795612818565b6001600160a01b0316336001600160a01b0316146127c55760405162461bcd60e51b8152600401610aa89061471b565b610df181612a8b565b6001600160a01b0381166000908152604460205260408120546128049060ff1660038111156127ff576127ff614850565b6134e0565b1592915050565b6000610ad460b4836148fb565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103546001600160a01b031690565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527faa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f906020015b60405180910390a150565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf90602001612890565b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527fc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d90602001612890565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527f31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b90602001612890565b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527f034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a611162190602001612890565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d790602001612890565b603d8190556040518181527fb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b590602001612890565b60068190556040518181527f82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab490602001612890565b603d54811115612b295760405162461bcd60e51b815260206004820152605960248201527f526f6e696e56616c696461746f725365743a2063616e6e6f7420736574206e7560448201527f6d626572206f66207072696f726974697a65642067726561746572207468616e60648201527f206e756d626572206f66206d61782076616c696461746f727300000000000000608482015260a401610aa8565b60458190556040518181527fa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e90602001612890565b6001600160a01b038116600090815260496020526040812054431115610ad4565b6001600160a01b03919091166000908152604660209081526040808320938352929052205460ff1690565b6000816003811115612bbe57612bbe614850565b836003811115612bd057612bd0614850565b1660ff1615159392505050565b6000606082516001600160401b03811115612bfa57612bfa61475d565b604051908082528060200260200182016040528015612c23578160200160208202803683370190505b5060055460405163889998ef60e01b81526004810187905291925060009182916001600160a01b0316908290829063889998ef90602401602060405180830381865afa158015612c77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c9b9190614aa7565b60405163033cdc2b60e31b8152600481018a90529091506000906001600160a01b038416906319e6e15890602401602060405180830381865afa158015612ce6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d0a9190614aa7565b90506000836001600160a01b03166357daa1708b8b6040518363ffffffff1660e01b8152600401612d3c929190614ac0565b600060405180830381865afa158015612d59573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612d819190810190614b34565b90506000806000600260009054906101000a90046001600160a01b03166001600160a01b0316631079402a6040518163ffffffff1660e01b8152600401606060405180830381865afa158015612ddb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dff9190614bc9565b92509250925060005b8c51811015612fc0578c8181518110612e2357612e23614773565b60200260200101519950600960008b6001600160a01b03166001600160a01b0316815260200190815260200160002060020160009054906101000a90046001600160a01b03169850612e948e8b878481518110612e8257612e82614773565b6020026020010151898b8989896134fe565b612ec48a8f6001600160a01b03919091166000908152604860209081526040808320938352929052205460ff1690565b612ef2576001600160a01b03808b16600090815260096020526040902060030154612ef2918c91168b61371a565b612efb8a612b5e565b158015612f0f5750612f0d8a8f612b7f565b155b15612f7e576001600160a01b038a166000908152604b6020526040902054612f37908d61483d565b6001600160a01b038b166000908152604b60205260409020548c51919d50908c9083908110612f6857612f68614773565b602002602001018181525050612f7e8a8a6137fd565b6001600160a01b038a166000908152604b60209081526040808320839055604a8252808320839055604d90915281205580612fb881614789565b915050612e08565b50604c600090555050505050505050509250929050565b6000546001600160a01b031682156130d257612ff381846138c7565b1561309957604051633b8cb16b60e01b81526001600160a01b03821690633b8cb16b9061302890879086908a90600401614c27565b600060405180830381600087803b15801561304257600080fd5b505af1158015613056573d6000803e3d6000fd5b505050507feb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf1398360405161308b91815260200190565b60405180910390a1506130d4565b604080518481524760208201527f0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5910160405180910390a15b505b50505050565b60606130e4613923565b600080546040516357e5970160e11b81526001600160a01b039091169063afcb2e029061311690600790600401614c9b565b600060405180830381865afa158015613133573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261315b9190810190614b34565b60048054604051632907e73160e11b81529293506000926001600160a01b039091169163520fce62916131919160079101614c9b565b600060405180830381865afa1580156131ae573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526131d69190810190614b34565b90506000613245600780548060200260200160405190810160405280929190818152602001828054801561323357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613215575b50505050508484603d54604554613c21565b9094509050613255848287613d40565b847f8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff61327f61221b565b60405161328c919061448b565b60405180910390a2505050919050565b6003546000906001600160a01b031663f0a4670960076132bd43600161483d565b6040518363ffffffff1660e01b81526004016132da929190614cae565b600060405180830381865afa1580156132f7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261331f9190810190614cc1565b905060005b82518110156134a857600083828151811061334157613341614773565b6020026020010151905060006133568261196d565b9050600061336383612b5e565b80613384575084848151811061337b5761337b614773565b60200260200101515b159050811580156133925750805b15613410576001600160a01b0383166000908152604460205260409020546133d19060019060ff1660038111156133cb576133cb614850565b90613ec6565b6001600160a01b0384166000908152604460205260409020805460ff1916600183600381111561340357613403614850565b0217905550505050613496565b81801561341b575080155b15613492576001600160a01b03831660009081526044602052604090205461345a9060019060ff16600381111561345457613454614850565b90613f01565b6001600160a01b0384166000908152604460205260409020805460ff1916600183600381111561348c5761348c614850565b02179055505b5050505b806134a081614789565b915050613324565b50827f60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef6134d36110c7565b604051611af2919061448b565b60008160038111156134f4576134f4614850565b60ff161592915050565b8315801561350a575084155b1561353d57604254604c5461351f91906148fb565b6001600160a01b0388166000908152604d6020526040902055613710565b60008561354c612710896148ce565b61355691906148fb565b9050600061356682612710614682565b905083811115613654576001600160a01b03891660008181526048602090815260408083208e845282528083208054600160ff199182168117909255948452604683528184208f855290925290912080549092161790556135e86135ca844361483d565b6001600160a01b038b16600090815260496020526040902054613f3d565b6001600160a01b038a166000818152604960209081526040808320859055805194855290840191909152600190830181905260608301528b917f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a906080015b60405180910390a361370d565b848111156136df576001600160a01b03891660008181526048602090815260408083208e84528252808320805460ff19166001908117909155848452604983528184205482519081529283018490529082019290925260608101919091528b91907f54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a90608001613647565b851561370d5785604c546136f391906148fb565b6001600160a01b038a166000908152604d60205260409020555b50505b5050505050505050565b6001600160a01b0383166000908152604d602052604090205480156130d45761374382826138c7565b156137a557816001600160a01b0316836001600160a01b0316856001600160a01b03167f72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c8460405161379791815260200190565b60405180910390a450505050565b816001600160a01b0316836001600160a01b0316856001600160a01b03167fd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b728447604051613797929190918252602082015260400190565b6001600160a01b0382166000908152604a602052604090205480156138c25761382682826138c7565b1561387d57816001600160a01b0316836001600160a01b03167f1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec8360405161387091815260200190565b60405180910390a3505050565b604080518281524760208201526001600160a01b0380851692908616917f6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e9101613870565b505050565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114613914576040519150601f19603f3d011682016040523d82523d6000602084013e613919565b606091505b5090949350505050565b600080546040805163909791dd60e01b815290516001600160a01b039092169291839163909791dd9160048083019260209291908290030181865afa158015613970573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906139949190614aa7565b90506000826001600160a01b031663017dd95060076040518263ffffffff1660e01b81526004016139c59190614c9b565b600060405180830381865afa1580156139e2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613a0a9190810190614b34565b600754909150600080826001600160401b03811115613a2b57613a2b61475d565b604051908082528060200260200182016040528015613a54578160200160208202803683370190505b5090506000805b84811015613b7d5760078181548110613a7657613a76614773565b9060005260206000200160009054906101000a90046001600160a01b0316915086868281518110613aa957613aa9614773565b60200260200101511080613ad857506001600160a01b0382166000908152600960205260409020600501544210155b15613b6b5785613ae786614d4d565b95508581518110613afa57613afa614773565b6020026020010151868281518110613b1457613b14614773565b6020908102919091010152818385613b2b81614789565b965081518110613b3d57613b3d614773565b60200260200101906001600160a01b031690816001600160a01b031681525050613b6682613f54565b613a5b565b80613b7581614789565b915050613a5b565b50508115613c19578181527f4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c9181604051613bb7919061448b565b60405180910390a16040516374d62f0360e11b81526001600160a01b0387169063e9ac5e0690613beb90849060040161448b565b600060405180830381600087803b158015613c0557600080fd5b505af115801561370d573d6000803e3d6000fd5b505050505050565b60606000806068905060008888888888604051602401613c45959493929190614d64565b60408051601f19818403018152919052602080820180516001600160e01b0316633bca0a8960e11b17905281518b519293506001929091600091613c88916148ce565b613c9390604061483d565b90506020840181888483895afa613ca957600093505b503d613cb457600092505b60208701965082613d2d5760405162461bcd60e51b815260206004820152603960248201527f507265636f6d70696c6555736167655069636b56616c696461746f725365743a60448201527f2063616c6c20746f20707265636f6d70696c65206661696c73000000000000006064820152608401610aa8565b8651955050505050509550959350505050565b815b604254811015613d9e57600081815260436020818152604080842080546001600160a01b0316855260448352908420805460ff19169055928490525280546001600160a01b031916905580613d9681614789565b915050613d42565b506000805b83811015613e80576000858281518110613dbf57613dbf614773565b602090810291909101810151600085815260439092526040909120549091506001600160a01b0390811690821603613e045782613dfb81614789565b93505050613e6e565b600083815260436020818152604080842080546001600160a01b03908116865260448452828620805460ff19908116909155908716808752928620805490911660031790559387905291905281546001600160a01b03191617905582613e6981614789565b935050505b80613e7881614789565b915050613da3565b5080604281905550817f3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e85604051613eb8919061448b565b60405180910390a250505050565b6000816003811115613eda57613eda614850565b836003811115613eec57613eec614850565b1760ff166003811115610bf557610bf5614850565b6000816003811115613f1557613f15614850565b19836003811115613f2857613f28614850565b1660ff166003811115610bf557610bf5614850565b600081831015613f4d5781610bf5565b5090919050565b6001600160a01b03811660009081526008602052604081205490819003613f79575050565b6001600160a01b038216600090815260096020526040812080546001600160a01b0319908116825560018201805482169055600282018054821690556003820180549091169055600481018290556005810182905590613fdc6006830182614135565b50506001600160a01b03821660009081526008602052604081208190556007805461400990600190614682565b8154811061401957614019614773565b6000918252602090912001546001600160a01b0390811691508316811461409c576001600160a01b038116600090815260086020526040902082905560078054829190841990811061406d5761406d614773565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60078054806140ad576140ad614db2565b600082815260209020810160001990810180546001600160a01b0319169055019055505050565b6040518060e0016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b508054614141906146e7565b6000825580601f10614151575050565b601f016020900490600052602060002090810190610df191905b808211156117dc576000815560010161416b565b6001600160a01b0381168114610df157600080fd5b600080604083850312156141a757600080fd5b82356141b28161417f565b915060208301356141c28161417f565b809150509250929050565b600080604083850312156141e057600080fd5b82356141eb8161417f565b946020939093013593505050565b60006020828403121561420b57600080fd5b8135610bf58161417f565b600060018060a01b038083511684526020818185015116818601528160408501511660408601528160608501511660608601526080840151608086015260a084015160a086015260c0840151915060e060c086015281518060e087015260005b818110156142935783810183015187820161010001528201614276565b506101009250600083828801015282601f19601f830116870101935050505092915050565b602081526000610bf56020830184614216565b6000806000806000806000806000806101408b8d0312156142eb57600080fd5b8a356142f68161417f565b995060208b01356143068161417f565b985060408b01356143168161417f565b975060608b01356143268161417f565b965060808b01356143368161417f565b955060a08b01356143468161417f565b999c989b5096999598949794965050505060c08301359260e08101359261010082013592506101209091013590565b60008083601f84011261438757600080fd5b5081356001600160401b0381111561439e57600080fd5b6020830191508360208260051b85010111156143b957600080fd5b9250929050565b600080602083850312156143d357600080fd5b82356001600160401b038111156143e957600080fd5b6143f585828601614375565b90969095509350505050565b6020808252825182820181905260009190848201906040850190845b8181101561443b57835115158352928401929184019160010161441d565b50909695505050505050565b600081518084526020808501945080840160005b838110156144805781516001600160a01b03168752958201959082019060010161445b565b509495945050505050565b602081526000610bf56020830184614447565b6000602082840312156144b057600080fd5b5035919050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561450c57603f198886030184526144fa858351614216565b945092850192908501906001016144de565b5092979650505050505050565b60008060006060848603121561452e57600080fd5b83356145398161417f565b95602085013595506040909401359392505050565b600080600080600060a0868803121561456657600080fd5b85356145718161417f565b945060208601356145818161417f565b935060408601356145918161417f565b925060608601356145a18161417f565b949793965091946080013592915050565b6000806000604084860312156145c757600080fd5b83356001600160401b038111156145dd57600080fd5b6145e986828701614375565b909790965060209590950135949350505050565b60208082526049908201527f486173536c617368496e64696361746f72436f6e74726163743a206d6574686f60408201527f642063616c6c6572206d75737420626520736c61736820696e64696361746f726060820152680818dbdb9d1c9858dd60ba1b608082015260a00190565b634e487b7160e01b600052601160045260246000fd5b81810381811115610ad457610ad461466c565b60208082526032908201527f43616e6469646174654d616e616765723a20717565727920666f72206e6f6e2d6040820152716578697374656e742063616e64696461746560701b606082015260800190565b600181811c908216806146fb57607f821691505b6020821081036125b757634e487b7160e01b600052602260045260246000fd5b60208082526022908201527f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460408201526132b960f11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161479b5761479b61466c565b5060010190565b60208082526031908201527f526f6e696e56616c696461746f725365743a206d6574686f642063616c6c6572604082015270206d75737420626520636f696e6261736560781b606082015260800190565b8051801515811461480357600080fd5b919050565b60008060006060848603121561481d57600080fd5b614826846147f3565b925060208401519150604084015190509250925092565b80820180821115610ad457610ad461466c565b634e487b7160e01b600052602160045260246000fd5b828152604081016003831061488b57634e487b7160e01b600052602160045260246000fd5b8260208301529392505050565b600080600080608085870312156148ae57600080fd5b505082516020840151604085015160609095015191969095509092509050565b8082028115828204841417610ad457610ad461466c565b634e487b7160e01b600052601260045260246000fd5b60008261490a5761490a6148e5565b500490565b6020808252603a908201527f4861735374616b696e67436f6e74726163743a206d6574686f642063616c6c6560408201527f72206d757374206265207374616b696e6720636f6e7472616374000000000000606082015260800190565b60408152600061497f6040830185614447565b90508260208301529392505050565b601f8211156138c257600081815260208120601f850160051c810160208610156149b55750805b601f850160051c820191505b81811015613c19578281556001016149c1565b81516001600160401b038111156149ed576149ed61475d565b614a01816149fb84546146e7565b8461498e565b602080601f831160018114614a365760008415614a1e5750858301515b600019600386901b1c1916600185901b178555613c19565b600085815260208120601f198616915b82811015614a6557888601518255948401946001909101908401614a46565b5085821015614a835787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082614aa257614aa26148e5565b500690565b600060208284031215614ab957600080fd5b5051919050565b828152604060208201526000614ad96040830184614447565b949350505050565b604051601f8201601f191681016001600160401b0381118282101715614b0957614b0961475d565b604052919050565b60006001600160401b03821115614b2a57614b2a61475d565b5060051b60200190565b60006020808385031215614b4757600080fd5b82516001600160401b03811115614b5d57600080fd5b8301601f81018513614b6e57600080fd5b8051614b81614b7c82614b11565b614ae1565b81815260059190911b82018301908381019087831115614ba057600080fd5b928401925b82841015614bbe57835182529284019290840190614ba5565b979650505050505050565b600080600060608486031215614bde57600080fd5b8351925060208401519150604084015190509250925092565b600081518084526020808501945080840160005b8381101561448057815187529582019590820190600101614c0b565b606081526000614c3a6060830186614447565b8281036020840152614c4c8186614bf7565b915050826040830152949350505050565b6000815480845260208085019450836000528060002060005b838110156144805781546001600160a01b031687529582019560019182019101614c76565b602081526000610bf56020830184614c5d565b60408152600061497f6040830185614c5d565b60006020808385031215614cd457600080fd5b82516001600160401b03811115614cea57600080fd5b8301601f81018513614cfb57600080fd5b8051614d09614b7c82614b11565b81815260059190911b82018301908381019087831115614d2857600080fd5b928401925b82841015614bbe57614d3e846147f3565b82529284019290840190614d2d565b600081614d5c57614d5c61466c565b506000190190565b60a081526000614d7760a0830188614447565b8281036020840152614d898188614bf7565b90508281036040840152614d9d8187614bf7565b60608401959095525050608001529392505050565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220f8662e4eca34546e81764a66cb592c557c4c18753c150599f5fda6bf38a3056464736f6c63430008110033",
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

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCaller) BulkJailed(opts *bind.CallOpts, _addrList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "bulkJailed", _addrList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorSession) BulkJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.BulkJailed(&_Validator.CallOpts, _addrList)
}

// BulkJailed is a free data retrieval call binding the contract method 0x428483c3.
//
// Solidity: function bulkJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) BulkJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.BulkJailed(&_Validator.CallOpts, _addrList)
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
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
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
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
func (_Validator *ValidatorSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _bridgeOperatorList)
func (_Validator *ValidatorCallerSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
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
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
func (_Validator *ValidatorSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,bytes))
func (_Validator *ValidatorCallerSession) GetCandidateInfo(_candidate common.Address) (ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
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
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
func (_Validator *ValidatorSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,bytes)[] _list)
func (_Validator *ValidatorCallerSession) GetCandidateInfos() ([]ICandidateManagerValidatorCandidate, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
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

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) Jailed(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "jailed", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_Validator *ValidatorSession) Jailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.Jailed(&_Validator.CallOpts, _addr)
}

// Jailed is a free data retrieval call binding the contract method 0x7043e5dd.
//
// Solidity: function jailed(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) Jailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.Jailed(&_Validator.CallOpts, _addr)
}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCaller) JailedAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "jailedAtBlock", _addr, _blockNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorSession) JailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.JailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// JailedAtBlock is a free data retrieval call binding the contract method 0x2607d919.
//
// Solidity: function jailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCallerSession) JailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.JailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) JailedTimeLeft(opts *bind.CallOpts, _addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "jailedTimeLeft", _addr)

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

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) JailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.JailedTimeLeft(&_Validator.CallOpts, _addr)
}

// JailedTimeLeft is a free data retrieval call binding the contract method 0x81f9535f.
//
// Solidity: function jailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) JailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.JailedTimeLeft(&_Validator.CallOpts, _addr)
}

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) JailedTimeLeftAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "jailedTimeLeftAtBlock", _addr, _blockNum)

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

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) JailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.JailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// JailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x85ad5aec.
//
// Solidity: function jailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) JailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.JailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
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

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorCaller) MiningRewardDeprecated(opts *bind.CallOpts, _blockProducers []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "miningRewardDeprecated", _blockProducers)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorSession) MiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _Validator.Contract.MiningRewardDeprecated(&_Validator.CallOpts, _blockProducers)
}

// MiningRewardDeprecated is a free data retrieval call binding the contract method 0x4a68f8c6.
//
// Solidity: function miningRewardDeprecated(address[] _blockProducers) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) MiningRewardDeprecated(_blockProducers []common.Address) ([]bool, error) {
	return _Validator.Contract.MiningRewardDeprecated(&_Validator.CallOpts, _blockProducers)
}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorCaller) MiningRewardDeprecatedAtPeriod(opts *bind.CallOpts, _blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "miningRewardDeprecatedAtPeriod", _blockProducers, _period)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorSession) MiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _Validator.Contract.MiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducers, _period)
}

// MiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x92a8c2e8.
//
// Solidity: function miningRewardDeprecatedAtPeriod(address[] _blockProducers, uint256 _period) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) MiningRewardDeprecatedAtPeriod(_blockProducers []common.Address, _period *big.Int) ([]bool, error) {
	return _Validator.Contract.MiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducers, _period)
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
// Solidity: function totalBridgeOperators() view returns(uint256)
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
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_Validator *ValidatorSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_Validator *ValidatorCallerSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
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
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactor) GrantValidatorCandidate(opts *bind.TransactOpts, _admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "grantValidatorCandidate", _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorSession) GrantValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.GrantValidatorCandidate(&_Validator.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// GrantValidatorCandidate is a paid mutator transaction binding the contract method 0x733ec970.
//
// Solidity: function grantValidatorCandidate(address _admin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactorSession) GrantValidatorCandidate(_admin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.GrantValidatorCandidate(&_Validator.TransactOpts, _admin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_Validator *ValidatorTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_Validator *ValidatorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
}

// Initialize is a paid mutator transaction binding the contract method 0x3986de6a.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __numberOfBlocksInEpoch) returns()
func (_Validator *ValidatorTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __numberOfBlocksInEpoch *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __numberOfBlocksInEpoch)
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
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerSetUpdated is a free log retrieval operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterBlockProducerSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*ValidatorBlockProducerSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockProducerSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockProducerSetUpdatedIterator{contract: _Validator.contract, event: "BlockProducerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerSetUpdated is a free log subscription operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchBlockProducerSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBlockProducerSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockProducerSetUpdated", periodRule)
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

// ParseBlockProducerSetUpdated is a log parse operation binding the contract event 0x60324bb9c8b0d077621d76762c52d6cc937043427992a2f6a602b449315922ef.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, address[] consensusAddrs)
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
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSetUpdated is a free log retrieval operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*ValidatorBridgeOperatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorSetUpdatedIterator{contract: _Validator.contract, event: "BridgeOperatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSetUpdated is a free log subscription operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorSetUpdated", periodRule)
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

// ParseBridgeOperatorSetUpdated is a log parse operation binding the contract event 0x8d7d519e81c2b8dc67b44fd645fd2c8805110d9ab1d643e3dd68b622bde331ff.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, address[] bridgeOperators)
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

// ValidatorCandidateRevokedTimestampUpdatedIterator is returned from FilterCandidateRevokedTimestampUpdated and is used to iterate over the raw logs and unpacked data for CandidateRevokedTimestampUpdated events raised by the Validator contract.
type ValidatorCandidateRevokedTimestampUpdatedIterator struct {
	Event *ValidatorCandidateRevokedTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateRevokedTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateRevokedTimestampUpdated)
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
		it.Event = new(ValidatorCandidateRevokedTimestampUpdated)
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
func (it *ValidatorCandidateRevokedTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateRevokedTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateRevokedTimestampUpdated represents a CandidateRevokedTimestampUpdated event raised by the Validator contract.
type ValidatorCandidateRevokedTimestampUpdated struct {
	ConsensusAddr    common.Address
	RevokedTimestamp *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCandidateRevokedTimestampUpdated is a free log retrieval operation binding the contract event 0xdb451f2c533d44367eeca766bcee562bfc473b8d3f5d34f7b87befe026434aaa.
//
// Solidity: event CandidateRevokedTimestampUpdated(address indexed consensusAddr, uint256 revokedTimestamp)
func (_Validator *ValidatorFilterer) FilterCandidateRevokedTimestampUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCandidateRevokedTimestampUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateRevokedTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateRevokedTimestampUpdatedIterator{contract: _Validator.contract, event: "CandidateRevokedTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateRevokedTimestampUpdated is a free log subscription operation binding the contract event 0xdb451f2c533d44367eeca766bcee562bfc473b8d3f5d34f7b87befe026434aaa.
//
// Solidity: event CandidateRevokedTimestampUpdated(address indexed consensusAddr, uint256 revokedTimestamp)
func (_Validator *ValidatorFilterer) WatchCandidateRevokedTimestampUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateRevokedTimestampUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateRevokedTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateRevokedTimestampUpdated)
				if err := _Validator.contract.UnpackLog(event, "CandidateRevokedTimestampUpdated", log); err != nil {
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

// ParseCandidateRevokedTimestampUpdated is a log parse operation binding the contract event 0xdb451f2c533d44367eeca766bcee562bfc473b8d3f5d34f7b87befe026434aaa.
//
// Solidity: event CandidateRevokedTimestampUpdated(address indexed consensusAddr, uint256 revokedTimestamp)
func (_Validator *ValidatorFilterer) ParseCandidateRevokedTimestampUpdated(log types.Log) (*ValidatorCandidateRevokedTimestampUpdated, error) {
	event := new(ValidatorCandidateRevokedTimestampUpdated)
	if err := _Validator.contract.UnpackLog(event, "CandidateRevokedTimestampUpdated", log); err != nil {
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
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributed is a free log retrieval operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributedIterator{contract: _Validator.contract, event: "StakingRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributed is a free log subscription operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
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

// ParseStakingRewardDistributed is a log parse operation binding the contract event 0xeb09b8cc1cefa77cd4ec30003e6364cf60afcedd20be8c09f26e717788baf139.
//
// Solidity: event StakingRewardDistributed(uint256 amount)
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
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributionFailed is a free log retrieval operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributionFailed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributionFailedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributionFailedIterator{contract: _Validator.contract, event: "StakingRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributionFailed is a free log subscription operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
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

// ParseStakingRewardDistributionFailed is a log parse operation binding the contract event 0x0752cb1e4b6fb7b2beb1cf423d908acaec7acfb7782e67a88d158351b1c0c4a5.
//
// Solidity: event StakingRewardDistributionFailed(uint256 amount, uint256 contractBalance)
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
