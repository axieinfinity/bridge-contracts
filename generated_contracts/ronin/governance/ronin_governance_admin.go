// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

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

// BridgeOperatorsBallotBridgeOperatorSet is an auto generated low-level Go binding around an user-defined struct.
type BridgeOperatorsBallotBridgeOperatorSet struct {
	Period    *big.Int
	Epoch     *big.Int
	Operators []common.Address
}

// GlobalProposalGlobalProposalDetail is an auto generated low-level Go binding around an user-defined struct.
type GlobalProposalGlobalProposalDetail struct {
	Nonce           *big.Int
	ExpiryTimestamp *big.Int
	TargetOptions   []uint8
	Values          []*big.Int
	Calldatas       [][]byte
	GasAmounts      []*big.Int
}

// ProposalProposalDetail is an auto generated low-level Go binding around an user-defined struct.
type ProposalProposalDetail struct {
	Nonce           *big.Int
	ChainId         *big.Int
	ExpiryTimestamp *big.Int
	Targets         []common.Address
	Values          []*big.Int
	Calldatas       [][]byte
	GasAmounts      []*big.Int
}

// SignatureConsumerSignature is an auto generated low-level Go binding around an user-defined struct.
type SignatureConsumerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalExpiryDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeBridgeContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeRoninTrustedOrgContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitPollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalProposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GlobalProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"successCalls\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"returnDatas\",\"type\":\"bytes[]\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"bridgeOperatorsVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"castProposalVoteForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"createEmergencyExitPoll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"deleteExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"emergencyPollVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"}],\"name\":\"getBridgeOperatorVotingSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getProposalSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSyncedBridgeOperatorSetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeVoter\",\"type\":\"address\"}],\"name\":\"lastVotedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"proposalVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"_targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"proposeGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeGlobalProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"proposeProposalForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryDuration\",\"type\":\"uint256\"}],\"name\":\"setProposalExpiryDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumVoteStatusConsumer.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"againstVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"_ballot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteBridgeOperatorsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"voteEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200614a3803806200614a8339810160408190526200003491620002e9565b838382806200004281600255565b50604080516020808201839052601660608301527f524f4e494e5f474f5645524e414e43455f41444d494e000000000000000000006080808401919091526107e4838501528351808403909101815260a0830184528051908201207f599a80fcaa47b95e2323ab4d34d34e0cc9feda4b843edafcc30c7bdf60ea15bf60c08401527f7e7935007966eb860f4a2ee3dcc9fd53fb3205ce2aa86b0126d4893d4d4c14b960e08401527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6610100840152610120808401919091528351808403909101815261014090920190925280519101207ff8704f8860d9e985bf6c52ec4738bd10fe31487599b36c0944f746ea09dc256b14620001a55760405162461bcd60e51b815260206004820152601f60248201527f476f7665726e616e636541646d696e3a20696e76616c696420646f6d61696e00604482015260640160405180910390fd5b620001b083620001d9565b620001bb826200022e565b505050620001cf826200027d60201b60201c565b505050506200033b565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7906020015b60405180910390a150565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae0489060200162000223565b600b80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b68361699060200162000223565b80516001600160a01b0381168114620002e457600080fd5b919050565b600080600080608085870312156200030057600080fd5b6200030b85620002cc565b93506200031b60208601620002cc565b92506200032b60408601620002cc565b6060959095015193969295505050565b615dff806200034b6000396000f3fe608060405234801561001057600080fd5b50600436106101fa5760003560e01c8063663ac0111161011a578063a8a0e32c116100ad578063cd5965831161007c578063cd596583146104d0578063cdf64a76146104e1578063dcc3eb19146104f4578063f3b7dead14610507578063fb4f63711461051a57600080fd5b8063a8a0e32c14610447578063b384abef1461045a578063b5e337de146104b5578063bc96180b146104c857600080fd5b80639a7d3382116100e95780639a7d3382146103fb5780639e0dc0b31461040e578063a1819f9a14610421578063a2fae5701461043457600080fd5b8063663ac0111461039b5780637eff275e146103ae578063988ef53c146103c157806399439089146103ea57600080fd5b80632e96a6fb116101925780633e2c314f116101615780633e2c314f146103425780635511cde11461036257806360911e8e1461037357806362e52e5f1461038657600080fd5b80632e96a6fb146102d95780632faf925d146102ec57806334d5f37b146102ff5780633644e5151461032d57600080fd5b80631c905e39116101ce5780631c905e391461024d578063204e1c7a146102785780632b5df351146102a35780632c5e6520146102c657600080fd5b80624054b8146101ff57806309fcd8c7146102145780630b26cf66146102275780630b8818301461023a575b600080fd5b61021261020d366004614888565b61052d565b005b61021261022236600461491c565b610582565b610212610235366004614a00565b61064b565b610212610248366004614888565b6106a0565b61026061025b366004614a1d565b6106bc565b60405161026f93929190614ab9565b60405180910390f35b61028b610286366004614a00565b610ab9565b6040516001600160a01b03909116815260200161026f565b6102b66102b1366004614b59565b610bab565b604051901515815260200161026f565b6102b66102d4366004614b59565b610be0565b6102126102e7366004614b92565b610c16565b6102126102fa366004614bab565b610c3e565b61031f61030d366004614b92565b60006020819052908152604090205481565b60405190815260200161026f565b61031f600080516020615daa83398151915281565b610355610350366004614c04565b610c78565b60405161026f9190614c56565b6003546001600160a01b031661028b565b610212610381366004614c69565b610db5565b61038e610e81565b60405161026f9190614d12565b6102126103a9366004614d59565b610f23565b6102126103bc366004614e36565b61102f565b61031f6103cf366004614a00565b6001600160a01b031660009081526009602052604090205490565b600b546001600160a01b031661028b565b610212610409366004614a1d565b611160565b61021261041c366004614e6f565b61116e565b61021261042f366004614ec1565b6113d4565b610212610442366004614f9a565b6114ba565b610212610455366004614fe0565b611586565b6104a4610468366004614a1d565b600160208181526000938452604080852090915291835291208054918101546002820154600383015460069093015460ff909416939192909185565b60405161026f959493929190615026565b6102126104c3366004614a00565b6115c1565b61031f611613565b6004546001600160a01b031661028b565b6102126104ef366004614a00565b611623565b6102b661050236600461505b565b6116b7565b61028b610515366004614a00565b6116e9565b610212610528366004614bab565b6117b6565b600061053833611821565b1161055e5760405162461bcd60e51b815260040161055590615080565b60405180910390fd5b61057b8585858585600080516020615daa83398151915233611995565b5050505050565b600061058d33611821565b116105aa5760405162461bcd60e51b815260040161055590615080565b61063f8989898989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506105f092508a91508b905061524f565b87878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061062d9250610da6915050565b6004546001600160a01b0316336119ff565b50505050505050505050565b33301461066a5760405162461bcd60e51b81526004016105559061525c565b6000816001600160a01b03163b116106945760405162461bcd60e51b8152600401610555906152a3565b61069d81611b34565b50565b61057b8585858585600080516020615daa833981519152611b89565b600082815260016020908152604080832084845290915281206004810154600582015460609384938493909290916106f482846152fd565b9050806001600160401b0381111561070e5761070e6150cc565b604051908082528060200260200182016040528015610737578160200160208202803683370190505b509550806001600160401b03811115610752576107526150cc565b60405190808252806020026020018201604052801561079d57816020015b60408051606081018252600080825260208083018290529282015282526000199092019101816107705790505b509450806001600160401b038111156107b8576107b86150cc565b6040519080825280602002602001820160405280156107e1578160200160208202803683370190505b50965060005b8381101561094057600087828151811061080357610803615310565b6020026020010190600181111561081c5761081c614a3f565b9081600181111561082f5761082f614a3f565b90525060008a81526001602090815260408083208c84529091528120600487018054600790920192918490811061086857610868615310565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286518790839081106108c6576108c6615310565b60200260200101819052508460040181815481106108e6576108e6615310565b9060005260206000200160009054906101000a90046001600160a01b031688828151811061091657610916615310565b6001600160a01b03909216602092830291909101909101528061093881615326565b9150506107e7565b5060005b82811015610aad5760018761095986846152fd565b8151811061096957610969615310565b6020026020010190600181111561098257610982614a3f565b9081600181111561099557610995614a3f565b90525060008a81526001602090815260408083208c8452909152812060058701805460079092019291849081106109ce576109ce615310565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286610a2386846152fd565b81518110610a3357610a33615310565b6020026020010181905250846005018181548110610a5357610a53615310565b6000918252602090912001546001600160a01b031688610a7386846152fd565b81518110610a8357610a83615310565b6001600160a01b039092166020928302919091019091015280610aa581615326565b915050610944565b50505050509250925092565b6000806000836001600160a01b0316604051610adf90635c60da1b60e01b815260040190565b600060405180830381855afa9150503d8060008114610b1a576040519150601f19603f3d011682016040523d82523d6000602084013e610b1f565b606091505b509150915081610b8f5760405162461bcd60e51b815260206004820152603560248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060696d706044820152741b195b595b9d185d1a5bdb8a0a580819985a5b1959605a1b6064820152608401610555565b80806020019051810190610ba3919061533f565b949350505050565b600083815260086020908152604080832085845282528083206001600160a01b03851684526002019091528120541515610ba3565b600083815260016020908152604080832085845282528083206001600160a01b038516845260080190915281205460ff16610ba3565b333014610c355760405162461bcd60e51b81526004016105559061525c565b61069d81600255565b61057b8585858585600080516020615daa833981519152610c676003546001600160a01b031690565b6004546001600160a01b0316611c10565b6060816001600160401b03811115610c9257610c926150cc565b604051908082528060200260200182016040528015610cdd57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610cb05790505b50905060005b82811015610d9d576000868152600a60209081526040808320888452909152812090858584818110610d1757610d17615310565b9050602002016020810190610d2c9190614a00565b6001600160a01b0316815260208082019290925260409081016000208151606081018352815460ff16815260018201549381019390935260020154908201528251839083908110610d7f57610d7f615310565b60200260200101819052508080610d9590615326565b915050610ce3565b50949350505050565b6003546001600160a01b031690565b610dd7838383610dc3611cc4565b600080516020615daa833981519152611e30565b8235600090815260086020908152604080832082870135845290915290206001815460ff166004811115610e0d57610e0d614a3f565b03610e7b57836005610e1f82826153c9565b507f7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d905084356020860135610e57604088018861535c565b604051610e67949392919061549c565b60405180910390a1805460ff191660021781555b50505050565b610ea560405180606001604052806000815260200160008152602001606081525090565b60408051606081018252600580548252600654602080840191909152600780548551818402810184018752818152949593949386019392830182828015610f1557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ef7575b505050505081525050905090565b6000610f2e33611821565b11610f4b5760405162461bcd60e51b815260040161055590615080565b60003390506000611014468d8d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610fd992508d91508e905061524f565b8a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508b925061215f915050565b9050611021828285612280565b505050505050505050505050565b33301461104e5760405162461bcd60e51b81526004016105559061525c565b604080516001600160a01b0383811660248084019190915283518084039091018152604490920183526020820180516001600160e01b03166308f2839760e41b17905291516000928516916110a29161551f565b6000604051808303816000865af19150503d80600081146110df576040519150601f19603f3d011682016040523d82523d6000602084013e6110e4565b606091505b505090508061115b5760405162461bcd60e51b815260206004820152603960248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206063686160448201527f6e676541646d696e28616464726573732960206661696c6564000000000000006064820152608401610555565b505050565b61116a82826123d9565b5050565b33600061117a82611821565b90506000811161119c5760405162461bcd60e51b815260040161055590615080565b60006111aa878787876123fb565b905080881461120b5760405162461bcd60e51b815260206004820152602760248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c696420766f6044820152660e8ca40d0c2e6d60cb1b6064820152608401610555565b6000818152600c6020526040902060058101546112845760405162461bcd60e51b815260206004820152603160248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f72206044820152706e6f6e2d6578697374656e7420766f746560781b6064820152608401610555565b6004815460ff16600481111561129c5761129c614a3f565b036112fe5760405162461bcd60e51b815260206004820152602c60248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f722060448201526b6578706972656420766f746560a01b6064820152608401610555565b600061131482868661130e611cc4565b8761246b565b9050600181600481111561132a5761132a614a3f565b0361137c57611339898961259a565b6040518381527fd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b9060200160405180910390a1815460ff1916600217825561063f565b600481600481111561139057611390614a3f565b0361063f576040518381527feecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe49060200160405180910390a150505050505050505050565b60006113df33611821565b116113fc5760405162461bcd60e51b815260040161055590615080565b6114ad8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c91829185019084908082843760009201919091525061147292508a91508b905061524f565b87878080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525033925061215f915050565b5050505050505050505050565b336114cd600b546001600160a01b031690565b6001600160a01b0316146114f457604051630e6444a160e31b815260040160405180910390fd5b6000611502858585856123fb565b6000818152600c60209081526040918290204260058201556004810186905582518481526001600160a01b038a811693820193909352918816828401526060820187905260808201869052915192935090917f18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de89181900360a00190a1505050505050565b600061159133611821565b116115ae5760405162461bcd60e51b815260040161055590615080565b61116a336115bb84615625565b83612280565b3330146115e05760405162461bcd60e51b81526004016105559061525c565b6000816001600160a01b03163b1161160a5760405162461bcd60e51b8152600401610555906152a3565b61069d81612736565b600061161e60025490565b905090565b3330146116425760405162461bcd60e51b81526004016105559061525c565b6000816001600160a01b03163b116116ae5760405162461bcd60e51b815260206004820152602960248201527f526f6e696e476f7665726e616e636541646d696e3a2073657420746f206e6f6e6044820152680b58dbdb9d1c9858dd60ba1b6064820152608401610555565b61069d81612784565b6000828152600c602090815260408083206001600160a01b038516845260020190915281205415155b90505b92915050565b6000806000836001600160a01b031660405161170f906303e1469160e61b815260040190565b600060405180830381855afa9150503d806000811461174a576040519150601f19603f3d011682016040523d82523d6000602084013e61174f565b606091505b509150915081610b8f5760405162461bcd60e51b815260206004820152602c60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206061646d60448201526b1a5b8a0a580819985a5b195960a21b6064820152608401610555565b60006117c133611821565b116117de5760405162461bcd60e51b815260040161055590615080565b6118198585858585600080516020615daa8339815191526118076003546001600160a01b031690565b6004546001600160a01b0316336127d2565b505050505050565b60008060006118386003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316631af0725f60e31b1790529251931692634bb5274a92611890929101615725565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516118c9919061551f565b600060405180830381855afa9150503d8060008114611904576040519150601f19603f3d011682016040523d82523d6000602084013e611909565b606091505b5091509150816119815760405162461bcd60e51b815260206004820152603f60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f476f7665726e6f7257656967687428616464726573732960206661696c6564006064820152608401610555565b80806020019051810190610ba39190615738565b6119a76119a188615625565b82612835565b5060006119bb6119b689615625565b61291d565b90506119f56119c989615625565b888888886119e1896119dc896000612abc565b612b12565b6119f08a6119dc8a6001612abc565b612b39565b5050505050505050565b6040805160c08101909152600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb554909182918190611a479060016152fd565b81526020018c81526020018b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250938552505050602082018b9052604082018a90526060909101889052909150611aab828787612e72565b9050611ac26002548261304090919063ffffffff16565b6000611acd8261291d565b9050611adb6000828f61313c565b935080847f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca84611b0a8761326a565b878a604051611b1c9493929190615856565b60405180910390a35050509998505050505050505050565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048906020015b60405180910390a150565b6000611b976119b688615625565b6020808901356000908152600180835260408083208c35845290935291902001549091508114611bd95760405162461bcd60e51b81526004016105559061593d565b611c07611be588615625565b87878787611bf8886119dc896000612abc565b6119f0896119dc8a6001612abc565b50505050505050565b6000611c278383611c208c6159f0565b9190612e72565b90506000611c3c611c378b6159f0565b61326a565b9050611c478261291d565b600080805260016020818152855183527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4990526040909120015414611c9e5760405162461bcd60e51b81526004016105559061593d565b61063f828a8a8a8a611cb58b6119dc896000612abc565b6119f08c6119dc8a6001612abc565b6000806000611cdb6003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b0316637de5dedd60e01b17905291516001600160a01b039390931692634bb5274a92611d24929101615725565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051611d5d919061551f565b600060405180830381855afa9150503d8060008114611d98576040519150601f19603f3d011682016040523d82523d6000602084013e611d9d565b606091505b509150915081611e155760405162461bcd60e51b815260206004820152603860248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c20606d696e60448201527f696d756d566f7465576569676874282960206661696c656400000000000000006064820152608401610555565b80806020019051810190611e299190615738565b9250505090565b600554853510801590611e495750600654602086013510155b611ebb5760405162461bcd60e51b815260206004820152603d60248201527f424f73476f7665726e616e636550726f706f73616c3a20717565727920666f7260448201527f206f7574646174656420627269646765206f70657261746f72207365740000006064820152608401610555565b611ec4856133c4565b82611f255760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b6064820152608401610555565b6000806000611f3388613584565b90506000611f418583612b12565b89356000818152600860209081526040808320828f0135808552908352818420948452600a83528184209084529091528120929350909190805b8a8110156120ff57368c8c83818110611f9657611f96615310565b606002919091019150611fc1905086611fb26020840184615ac9565b83602001358460400135613636565b9850886001600160a01b0316886001600160a01b0316106120385760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526a34b3b732b91037b93232b960a91b6064820152608401610555565b8897505060006120478961365e565b905080156120ec576001600160a01b0389166000908152600960205260409020439055600192508c8c8381811061208057612080615310565b9050606002018460008b6001600160a01b03166001600160a01b0316815260200190815260200160002081816120b69190615ae6565b50600190506120c8868b848f8c61246b565b60048111156120d9576120d9614a3f565b036120ec5750505050505050505061057b565b50806120f781615326565b915050611f7b565b50806110215760405162461bcd60e51b815260206004820152602960248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526869676e61747572657360b81b6064820152608401610555565b612167614772565b876000036121b75760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e2069646044820152606401610555565b6040805160e08101825260008a81526020819052919091205481906121dd9060016152fd565b815260200189815260200188815260200187815260200186815260200185815260200184815250905061221b6002548261304090919063ffffffff16565b60006122268261291d565b905060006122358a838b61313c565b905081818b7fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9868860405161226b929190615b17565b60405180910390a45050979650505050505050565b468260200151146122e25760405162461bcd60e51b815260206004820152602660248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c6964206368604482015265185a5b881a5960d21b6064820152608401610555565b6122eb8261291d565b602080840151600090815260018083526040808320875184529093529190200154146123765760405162461bcd60e51b815260206004820152603460248201527f526f6e696e476f7665726e616e636541646d696e3a206361737420766f746520604482015273199bdc881a5b9d985b1a59081c1c9bdc1bdcd85b60621b6064820152608401610555565b6000612380611cc4565b905060008161238d6137c9565b6123979190615b41565b6123a29060016152fd565b6040805160608101825260008082526020820181905291810191909152909150611c07858585858a866123d48d611821565b613910565b60008281526001602090815260408083208484529091529020610e7b81613d2b565b604080517f697acba4deaf1a718d8c2d93e42860488cb7812696f28ca10eed17bac41e70276020808301919091526001600160a01b0396871682840152949095166060860152608085019290925260a0808501919091528151808503909101815260c09093019052815191012090565b6000808660040154118015612484575042866004015411155b1561249c5750845460ff191660049081178655612591565b6001600160a01b038516600090815260028701602052604090205415612506576124d0856001600160a01b03166014613f4d565b6040516020016124e09190615b54565b60408051601f198184030181529082905262461bcd60e51b825261055591600401615725565b6001600160a01b03851660009081526002870160209081526040808320859055848352600389019091528120805486919083906125449084906152fd565b925050819055905083811015801561257157506000875460ff16600481111561256f5761256f614a3f565b145b1561258957865460ff19166001908117885587018390555b5050845460ff165b95945050505050565b60006125ae600b546001600160a01b031690565b604080516001600160a01b0386811660248084019190915286821660448085019190915284518085039091018152606490930184526020830180516001600160e01b03166361e45aeb60e11b1790529251931692634bb5274a92612613929101615725565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161264c919061551f565b6000604051808303816000865af19150503d8060008114612689576040519150601f19603f3d011682016040523d82523d6000602084013e61268e565b606091505b505090508061115b5760405162461bcd60e51b815260206004820152606260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206065786560448201527f6352656c656173654c6f636b656446756e64466f72456d657267656e6379457860648201527f69745265717565737428616464726573732c616464726573732960206661696c608482015261195960f21b60a482015260c401610555565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d790602001611b7e565b600b80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b683616990602001611b7e565b6127da614772565b6127ee6127e68b6159f0565b8585856140e8565b50905060006127ff611c378c6159f0565b9050612827828b8b8b8b6128188c6119dc896000612abc565b6119f08d6119dc8a6001612abc565b509998505050505050505050565b602082015160009080820361288c5760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e2069646044820152606401610555565b60025461289a908590613040565b60006128a58561291d565b90506128b68282876040015161313c565b855190935083146128d95760405162461bcd60e51b815260040161055590615ba8565b8083837fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9888860405161290d929190615b17565b60405180910390a4505092915050565b6000806000806000808660800151905060008760600151905060008860a00151516001600160401b03811115612955576129556150cc565b60405190808252806020026020018201604052801561297e578160200160208202803683370190505b5060c08a015190915060005b82518110156129e7578a60a0015181815181106129a9576129a9615310565b6020026020010151805190602001208382815181106129ca576129ca615310565b6020908102919091010152806129df81615326565b91505061298a565b506020835102602084012097506020845102602085012096506020825102602083012095506020815102602082012094507fd051578048e6ff0bbc9fca3b65a42088dbde10f36ca841de566711087ad9b08a60001b8a600001518b602001518c604001518b8b8b8b604051602001612a97989796959493929190978852602088019690965260408701949094526060860192909252608085015260a084015260c083015260e08201526101000190565b6040516020818303038152906040528051906020012098505050505050505050919050565b604051600090612af4907fd900570327c4c0df8dd6bdd522b7da7e39145dd049d2fd4602276adcd511e3c29085908590602001615bee565b60405160208183030381529060405280519060200120905092915050565b60405161190160f01b60208201526022810183905260428101829052600090606201612af4565b8415801590612b4757508483145b612ba45760405162461bcd60e51b815260206004820152602860248201527f476f7665726e616e636550726f706f73616c3a20696e76616c696420617272616044820152670f240d8cadccee8d60c31b6064820152608401610555565b6000612bae611cc4565b9050600081612bbb6137c9565b612bc59190615b41565b612bd09060016152fd565b9050600080366000805b89811015612e06578a8a82818110612bf457612bf4615310565b606002919091019350600090508d8d83818110612c1357612c13615310565b9050602002016020810190612c289190615c13565b6001811115612c3957612c39614a3f565b03612c6557612c5e89612c4f6020860186615ac9565b85602001358660400135613636565b9350612d19565b60018d8d83818110612c7957612c79615310565b9050602002016020810190612c8e9190615c13565b6001811115612c9f57612c9f614a3f565b03612cb557612c5e88612c4f6020860186615ac9565b60405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636550726f706f73616c3a20717565727920666f7220756e604482015272737570706f7274656420766f7465207479706560681b6064820152608401610555565b836001600160a01b0316856001600160a01b031610612d845760405162461bcd60e51b815260206004820152602160248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964206f7264656044820152603960f91b6064820152608401610555565b8394506000612d9285611821565b90508015612df35760019250612de18f8f8f85818110612db457612db4615310565b9050602002016020810190612dc99190615c13565b8a8a89612ddb368b90038b018b615c30565b87613910565b15612df3575050505050505050611c07565b5080612dfe81615326565b915050612bda565b5080612e635760405162461bcd60e51b815260206004820152602660248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964207369676e60448201526561747572657360d01b6064820152608401610555565b50505050505050505050505050565b612e7a614772565b83518152602080850151604080840191909152600091830191909152840151516001600160401b03811115612eb157612eb16150cc565b604051908082528060200260200182016040528015612eda578160200160208202803683370190505b5060608083019190915284015160808083019190915284015160a08083019190915284015160c082015260005b84604001515181101561303857600185604001518281518110612f2c57612f2c615310565b60200260200101516001811115612f4557612f45614a3f565b03612f86578282606001518281518110612f6157612f61615310565b60200260200101906001600160a01b031690816001600160a01b031681525050613026565b600085604001518281518110612f9e57612f9e615310565b60200260200101516001811115612fb757612fb7614a3f565b03612fd3578382606001518281518110612f6157612f61615310565b60405162461bcd60e51b815260206004820152602260248201527f476c6f62616c50726f706f73616c3a20756e737570706f727465642074617267604482015261195d60f21b6064820152608401610555565b8061303081615326565b915050612f07565b509392505050565b600082606001515111801561305e5750816080015151826060015151145b801561307357508160a0015151826060015151145b801561308857508160c0015151826060015151145b6130d45760405162461bcd60e51b815260206004820152601e60248201527f50726f706f73616c3a20696e76616c6964206172726179206c656e67746800006044820152606401610555565b6130de81426152fd565b8260400151111561116a5760405162461bcd60e51b815260206004820152602260248201527f50726f706f73616c3a20696e76616c6964206578706972792074696d6573746160448201526106d760f41b6064820152608401610555565b6000838152602081905260408120549081900361316d5750600083815260208190526040902060019081905561323d565b600084815260016020908152604080832084845290915281209061319082613d2b565b90508061323a576000825460ff1660048111156131af576131af614a3f565b036132165760405162461bcd60e51b815260206004820152603160248201527f436f7265476f7665726e616e63653a2063757272656e742070726f706f73616c604482015270081a5cc81b9bdd0818dbdb5c1b195d1959607a1b6064820152608401610555565b6000868152602081905260408120805490919061323290615326565b918290555092505b50505b60009384526001602081815260408087208488529091529094209384019290925560069092019190915590565b6000806000806000808660600151905060008760400151905060008860800151516001600160401b038111156132a2576132a26150cc565b6040519080825280602002602001820160405280156132cb578160200160208202803683370190505b5060a08a015190915060005b8251811015613334578a6080015181815181106132f6576132f6615310565b60200260200101518051906020012083828151811061331757613317615310565b60209081029190910101528061332c81615326565b9150506132d7565b5082516020908102848201208551820286830120845183028584012084518402858501208e518f860151604080517f1463f426c05aff2c1a7a0957a71c9898bc8b47142540538e79ee25ee911413509881019890985287019190915260608601526080850184905260a0850183905260c0850182905260e08501819052929b509099509750955061010001612a97565b60006133d3604083018361535c565b9050116134365760405162461bcd60e51b815260206004820152602b60248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b6064820152608401610555565b6000613445604083018361535c565b600081811061345657613456615310565b905060200201602081019061346b9190614a00565b905060015b61347d604084018461535c565b905081101561115b57613493604084018461535c565b828181106134a3576134a3615310565b90506020020160208101906134b89190614a00565b6001600160a01b0316826001600160a01b03161061353e5760405162461bcd60e51b815260206004820152603860248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206f60448201527f72646572206f6620627269646765206f70657261746f727300000000000000006064820152608401610555565b61354b604084018461535c565b8281811061355b5761355b615310565b90506020020160208101906135709190614a00565b91508061357c81615326565b915050613470565b60008080613595604085018561535c565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250508251602090810293810193909320604080517fd679a49e9e099fa9ed83a5446aaec83e746b03ec6723d6f5efb29d37d7f0b78a818701528935818301529885013560608a01526080808a01929092528051808a03909201825260a0909801909752505084519401939093209392505050565b6000806000613647878787876141a5565b9150915061365481614292565b5095945050505050565b60008060006136756003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316635624191160e01b1790529251931692634bb5274a926136cd929101615725565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051613706919061551f565b600060405180830381855afa9150503d8060008114613741576040519150601f19603f3d011682016040523d82523d6000602084013e613746565b606091505b5091509150816119815760405162461bcd60e51b815260206004820152604260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f427269646765566f74657257656967687428616464726573732960206661696c606482015261195960f21b608482015260a401610555565b60008060006137e06003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b031663926323d560e01b17905291516001600160a01b039390931692634bb5274a92613829929101615725565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051613862919061551f565b600060405180830381855afa9150503d806000811461389d576040519150601f19603f3d011682016040523d82523d6000602084013e6138a2565b606091505b509150915081611e155760405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060746f74604482015272185b15d95a59da1d1cca0a580819985a5b1959606a1b6064820152608401610555565b6020808801518851600082815260018452604080822083835290945292832061393881613d2b565b156139495760019350505050613d20565b6020808c0151600090815290819052604090205482146139c45760405162461bcd60e51b815260206004820152603060248201527f436f7265476f7665726e616e63653a20717565727920666f7220696e76616c6960448201526f642070726f706f73616c206e6f6e636560801b6064820152608401610555565b6000815460ff1660048111156139dc576139dc614a3f565b14613a375760405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a2074686520766f74652069732066696e616044820152641b1a5e995960da1b6064820152608401610555565b6001600160a01b038716600090815260088201602052604090205460ff1615613a7e57613a6e876001600160a01b03166014613f4d565b6040516020016124e09190615c91565b6001600160a01b03871660009081526008820160209081526040909120805460ff19166001179055860151151580613ab95750604086015115155b80613ac75750855160ff1615155b15613b0e576001600160a01b03871660009081526007820160209081526040918290208851815460ff191660ff909116178155908801516001820155908701516002909101555b866001600160a01b031681600101547f1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a238c88604051613b4e929190615ce1565b60405180910390a3600080808c6001811115613b6c57613b6c614a3f565b03613bc1576004830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600384018054899290613bb49084906152fd565b9250508190559150613c80565b60018c6001811115613bd557613bd5614a3f565b03613c2a576005830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600284018054899290613c1d9084906152fd565b9250508190559050613c80565b60405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a20756e737570706f7274656420766f7465604482015264207479706560d81b6064820152608401610555565b8a8210613cd457825460ff19166001908117845580840154604051919750907f5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b90600090a2613ccf838e614448565b613d1a565b898110613d1a57825460ff19166003178355600180840154604051919750907f55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f90600090a25b50505050505b979650505050505050565b600080825460ff166004811115613d4457613d44614a3f565b148015613d55575042826006015411155b90508015613f485760018201546040517f58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a5590600090a260005b6004830154811015613e4757826008016000846004018381548110613db557613db5615310565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556004840180546007860192919084908110613dfe57613dfe615310565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613e3f81615326565b915050613d8e565b5060005b6005830154811015613f0457826008016000846005018381548110613e7257613e72615310565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556005840180546007860192919084908110613ebb57613ebb615310565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613efc81615326565b915050613e4b565b50815460ff191682556000600183018190556002830181905560038301819055613f329060048401906147af565b613f406005830160006147af565b600060068301555b919050565b60606000613f5c8360026153a5565b613f679060026152fd565b6001600160401b03811115613f7e57613f7e6150cc565b6040519080825280601f01601f191660200182016040528015613fa8576020820181803683370190505b509050600360fc1b81600081518110613fc357613fc3615310565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613ff257613ff2615310565b60200101906001600160f81b031916908160001a90535060006140168460026153a5565b6140219060016152fd565b90505b6001811115614099576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061405557614055615310565b1a60f81b82828151811061406b5761406b615310565b60200101906001600160f81b031916908160001a90535060049490941c9361409281615cf8565b9050614024565b5083156116e05760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610555565b6140f0614772565b60006140fd868686612e72565b91506141146002548361304090919063ffffffff16565b600061411f8361291d565b9050614131600082896020015161313c565b835190925082146141545760405162461bcd60e51b815260040161055590615ba8565b80827f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca856141818b61326a565b8b896040516141939493929190615856565b60405180910390a35094509492505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156141dc5750600090506003614289565b8460ff16601b141580156141f457508460ff16601c14155b156142055750600090506004614289565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614259573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661428257600060019250925050614289565b9150600090505b94509492505050565b60008160048111156142a6576142a6614a3f565b036142ae5750565b60018160048111156142c2576142c2614a3f565b0361430f5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610555565b600281600481111561432357614323614a3f565b036143705760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610555565b600381600481111561438457614384614a3f565b036143dc5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610555565b60048160048111156143f0576143f0614a3f565b0361069d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610555565b614451816144b5565b1561116a57815460ff1916600217825560008061446d836144cf565b9150915083600101547fe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba83836040516144a7929190615d0f565b60405180910390a250505050565b60008160200151600014806116e357505060200151461490565b6060806144db836144b5565b6145335760405162461bcd60e51b815260206004820152602360248201527f50726f706f73616c3a20717565727920666f7220696e76616c696420636861696044820152621b925960ea1b6064820152608401610555565b8260600151516001600160401b03811115614550576145506150cc565b604051908082528060200260200182016040528015614579578160200160208202803683370190505b5091508260600151516001600160401b03811115614599576145996150cc565b6040519080825280602002602001820160405280156145cc57816020015b60608152602001906001900390816145b75790505b50905060005b83606001515181101561476c578360c0015181815181106145f5576145f5615310565b60200260200101515a1161464b5760405162461bcd60e51b815260206004820152601a60248201527f50726f706f73616c3a20696e73756666696369656e74206761730000000000006044820152606401610555565b8360600151818151811061466157614661615310565b60200260200101516001600160a01b03168460800151828151811061468857614688615310565b60200260200101518560c0015183815181106146a6576146a6615310565b6020026020010151908660a0015184815181106146c5576146c5615310565b60200260200101516040516146da919061551f565b600060405180830381858888f193505050503d8060008114614718576040519150601f19603f3d011682016040523d82523d6000602084013e61471d565b606091505b5084838151811061473057614730615310565b6020026020010184848151811061474957614749615310565b6020908102919091010191909152901515905261476581615326565b90506145d2565b50915091565b6040518060e00160405280600081526020016000815260200160008152602001606081526020016060815260200160608152602001606081525090565b508054600082559060005260206000209081019061069d91905b808211156147dd57600081556001016147c9565b5090565b600060e082840312156147f357600080fd5b50919050565b60008083601f84011261480b57600080fd5b5081356001600160401b0381111561482257600080fd5b6020830191508360208260051b850101111561483d57600080fd5b9250929050565b60008083601f84011261485657600080fd5b5081356001600160401b0381111561486d57600080fd5b60208301915083602060608302850101111561483d57600080fd5b6000806000806000606086880312156148a057600080fd5b85356001600160401b03808211156148b757600080fd5b6148c389838a016147e1565b965060208801359150808211156148d957600080fd5b6148e589838a016147f9565b909650945060408801359150808211156148fe57600080fd5b5061490b88828901614844565b969995985093965092949392505050565b600080600080600080600080600060a08a8c03121561493a57600080fd5b8935985060208a01356001600160401b038082111561495857600080fd5b6149648d838e016147f9565b909a50985060408c013591508082111561497d57600080fd5b6149898d838e016147f9565b909850965060608c01359150808211156149a257600080fd5b6149ae8d838e016147f9565b909650945060808c01359150808211156149c757600080fd5b506149d48c828d016147f9565b915080935050809150509295985092959850929598565b6001600160a01b038116811461069d57600080fd5b600060208284031215614a1257600080fd5b81356116e0816149eb565b60008060408385031215614a3057600080fd5b50508035926020909101359150565b634e487b7160e01b600052602160045260246000fd5b6002811061069d5761069d614a3f565b600081518084526020808501945080840160005b83811015614aae578151805160ff16885283810151848901526040908101519088015260609096019590820190600101614a79565b509495945050505050565b606080825284519082018190526000906020906080840190828801845b82811015614afb5781516001600160a01b031684529284019290840190600101614ad6565b5050508381038285015285518082528683019183019060005b81811015614b39578351614b2781614a55565b83529284019291840191600101614b14565b50508481036040860152614b4d8187614a65565b98975050505050505050565b600080600060608486031215614b6e57600080fd5b83359250602084013591506040840135614b87816149eb565b809150509250925092565b600060208284031215614ba457600080fd5b5035919050565b600080600080600060608688031215614bc357600080fd5b85356001600160401b0380821115614bda57600080fd5b9087019060c0828a031215614bee57600080fd5b909550602087013590808211156148d957600080fd5b60008060008060608587031215614c1a57600080fd5b843593506020850135925060408501356001600160401b03811115614c3e57600080fd5b614c4a878288016147f9565b95989497509550505050565b6020815260006116e06020830184614a65565b600080600060408486031215614c7e57600080fd5b83356001600160401b0380821115614c9557600080fd5b9085019060608288031215614ca957600080fd5b90935060208501359080821115614cbf57600080fd5b50614ccc86828701614844565b9497909650939450505050565b600081518084526020808501945080840160005b83811015614aae5781516001600160a01b031687529582019590820190600101614ced565b60208152815160208201526020820151604082015260006040830151606080840152610ba36080840182614cd9565b6002811061069d57600080fd5b8035613f4881614d41565b60008060008060008060008060008060c08b8d031215614d7857600080fd5b8a35995060208b01356001600160401b0380821115614d9657600080fd5b614da28e838f016147f9565b909b50995060408d0135915080821115614dbb57600080fd5b614dc78e838f016147f9565b909950975060608d0135915080821115614de057600080fd5b614dec8e838f016147f9565b909750955060808d0135915080821115614e0557600080fd5b50614e128d828e016147f9565b9094509250614e25905060a08c01614d4e565b90509295989b9194979a5092959850565b60008060408385031215614e4957600080fd5b8235614e54816149eb565b91506020830135614e64816149eb565b809150509250929050565b600080600080600060a08688031215614e8757600080fd5b853594506020860135614e99816149eb565b93506040860135614ea9816149eb565b94979396509394606081013594506080013592915050565b60008060008060008060008060008060c08b8d031215614ee057600080fd5b8a35995060208b0135985060408b01356001600160401b0380821115614f0557600080fd5b614f118e838f016147f9565b909a50985060608d0135915080821115614f2a57600080fd5b614f368e838f016147f9565b909850965060808d0135915080821115614f4f57600080fd5b614f5b8e838f016147f9565b909650945060a08d0135915080821115614f7457600080fd5b50614f818d828e016147f9565b915080935050809150509295989b9194979a5092959850565b60008060008060808587031215614fb057600080fd5b8435614fbb816149eb565b93506020850135614fcb816149eb565b93969395505050506040820135916060013590565b60008060408385031215614ff357600080fd5b82356001600160401b0381111561500957600080fd5b615015858286016147e1565b9250506020830135614e6481614d41565b60a081016005871061503a5761503a614a3f565b95815260208101949094526040840192909252606083015260809091015290565b6000806040838503121561506e57600080fd5b823591506020830135614e64816149eb565b6020808252602c908201527f526f6e696e476f7665726e616e636541646d696e3a2073656e6465722069732060408201526b3737ba1033b7bb32b93737b960a11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715615104576151046150cc565b60405290565b60405160c081016001600160401b0381118282101715615104576151046150cc565b604051601f8201601f191681016001600160401b0381118282101715615154576151546150cc565b604052919050565b60006001600160401b03821115615175576151756150cc565b5060051b60200190565b600061519261518d8461515c565b61512c565b8381529050602080820190600585901b8401868111156151b157600080fd5b845b818110156152445780356001600160401b03808211156151d35760008081fd5b8188019150601f8a818401126151e95760008081fd5b8235828111156151fb576151fb6150cc565b61520c818301601f1916880161512c565b92508083528b8782860101111561522557600091508182fd5b80878501888501376000908301870152508552509282019282016151b3565b505050509392505050565b60006116e036848461517f565b60208082526027908201527f476f7665726e616e636541646d696e3a206f6e6c7920616c6c6f7765642073656040820152661b198b58d85b1b60ca1b606082015260800190565b60208082526024908201527f476f7665726e616e636541646d696e3a2073657420746f206e6f6e2d636f6e746040820152631c9858dd60e21b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b808201808211156116e3576116e36152e7565b634e487b7160e01b600052603260045260246000fd5b600060018201615338576153386152e7565b5060010190565b60006020828403121561535157600080fd5b81516116e0816149eb565b6000808335601e1984360301811261537357600080fd5b8301803591506001600160401b0382111561538d57600080fd5b6020019150600581901b360382131561483d57600080fd5b80820281158282048414176116e3576116e36152e7565b600081356116e3816149eb565b81358155600160208084013582840155600283016040850135601e198636030181126153f457600080fd5b850180356001600160401b0381111561540c57600080fd5b83820191508060051b360382131561542357600080fd5b6801000000000000000081111561543c5761543c6150cc565b8254818455808210156154705760008481528581208381019083015b8082101561546c5782825590880190615458565b5050505b50600092835260208320925b818110156119f55761548d836153bc565b8482015591840191850161547c565b84815260208082018590526060604083018190528201839052600090849060808401835b868110156154ee5783356154d3816149eb565b6001600160a01b0316825292820192908201906001016154c0565b5098975050505050505050565b60005b838110156155165781810151838201526020016154fe565b50506000910152565b600082516155318184602087016154fb565b9190910192915050565b600082601f83011261554c57600080fd5b8135602061555c61518d8361515c565b82815260059290921b8401810191818101908684111561557b57600080fd5b8286015b8481101561559f578035615592816149eb565b835291830191830161557f565b509695505050505050565b600082601f8301126155bb57600080fd5b813560206155cb61518d8361515c565b82815260059290921b840181019181810190868411156155ea57600080fd5b8286015b8481101561559f57803583529183019183016155ee565b600082601f83011261561657600080fd5b6116e08383356020850161517f565b600060e0823603121561563757600080fd5b61563f6150e2565b82358152602083013560208201526040830135604082015260608301356001600160401b038082111561567157600080fd5b61567d3683870161553b565b6060840152608085013591508082111561569657600080fd5b6156a2368387016155aa565b608084015260a08501359150808211156156bb57600080fd5b6156c736838701615605565b60a084015260c08501359150808211156156e057600080fd5b506156ed368286016155aa565b60c08301525092915050565b600081518084526157118160208601602086016154fb565b601f01601f19169290920160200192915050565b6020815260006116e060208301846156f9565b60006020828403121561574a57600080fd5b5051919050565b600081518084526020808501945080840160005b83811015614aae57815187529582019590820190600101615765565b600081518084526020808501808196508360051b8101915082860160005b858110156157c95782840389526157b78483516156f9565b9885019893509084019060010161579f565b5091979650505050505050565b8051825260208101516020830152604081015160408301526000606082015160e0606085015261580960e0850182614cd9565b9050608083015184820360808601526158228282615751565b91505060a083015184820360a086015261583c8282615781565b91505060c083015184820360c08601526125918282615751565b60808152600061586960808301876157d6565b60208681850152838203604085015260c08201865183528187015182840152604087015160c0604085015281815180845260e0860191508483019350600092505b808310156158d35783516158bd81614a55565b82529284019260019290920191908401906158aa565b506060890151935084810360608601526158ed8185615751565b9350505050608086015182820360808401526159098282615781565b91505060a086015182820360a08401526159238282615751565b935050505061259160608301846001600160a01b03169052565b6020808252602f908201527f476f7665726e616e636541646d696e3a206361737420766f746520666f72206960408201526e1b9d985b1a59081c1c9bdc1bdcd85b608a1b606082015260800190565b600082601f83011261599d57600080fd5b813560206159ad61518d8361515c565b82815260059290921b840181019181810190868411156159cc57600080fd5b8286015b8481101561559f5780356159e381614d41565b83529183019183016159d0565b600060c08236031215615a0257600080fd5b615a0a61510a565b823581526020830135602082015260408301356001600160401b0380821115615a3257600080fd5b615a3e3683870161598c565b60408401526060850135915080821115615a5757600080fd5b615a63368387016155aa565b60608401526080850135915080821115615a7c57600080fd5b615a8836838701615605565b608084015260a0850135915080821115615aa157600080fd5b50615aae368286016155aa565b60a08301525092915050565b60ff8116811461069d57600080fd5b600060208284031215615adb57600080fd5b81356116e081615aba565b8135615af181615aba565b60ff811660ff198354161782555060208201356001820155604082013560028201555050565b604081526000615b2a60408301856157d6565b905060018060a01b03831660208301529392505050565b818103818111156116e3576116e36152e7565b73024b9b7b630ba32b223b7bb32b93730b731b29d160651b815260008251615b838160148501602087016154fb565b6d08185b1c9958591e481d9bdd195960921b6014939091019283015250602201919050565b60208082526026908201527f436f7265476f7665726e616e63653a20696e76616c69642070726f706f73616c604082015265206e6f6e636560d01b606082015260800190565b8381526020810183905260608101615c0583614a55565b826040830152949350505050565b600060208284031215615c2557600080fd5b81356116e081614d41565b600060608284031215615c4257600080fd5b604051606081018181106001600160401b0382111715615c6457615c646150cc565b6040528235615c7281615aba565b8152602083810135908201526040928301359281019290925250919050565b6f021b7b932a3b7bb32b93730b731b29d160851b815260008251615cbc8160108501602087016154fb565b6d08185b1c9958591e481d9bdd195960921b6010939091019283015250601e01919050565b60408101615cee84614a55565b9281526020015290565b600081615d0757615d076152e7565b506000190190565b604080825283519082018190526000906020906060840190828701845b82811015615d4a578151151584529284019290840190600101615d2c565b50505083810382850152845180825282820190600581901b8301840187850160005b83811015615d9a57601f19868403018552615d888383516156f9565b94870194925090860190600101615d6c565b5090999850505050505050505056fef8704f8860d9e985bf6c52ec4738bd10fe31487599b36c0944f746ea09dc256ba2646970667358221220892836b9b211104e964b2393142517a36af3685adea3db2948c6b93dbe0c46c564736f6c63430008110033",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// GovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceMetaData.Bin instead.
var GovernanceBin = GovernanceMetaData.Bin

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _roninTrustedOrganizationContract common.Address, _bridgeContract common.Address, _validatorContract common.Address, _proposalExpiryDuration *big.Int) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := GovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceBin), backend, _roninTrustedOrganizationContract, _bridgeContract, _validatorContract, _proposalExpiryDuration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Governance *GovernanceCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Governance *GovernanceSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Governance.Contract.DOMAINSEPARATOR(&_Governance.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Governance *GovernanceCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Governance.Contract.DOMAINSEPARATOR(&_Governance.CallOpts)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Governance *GovernanceCaller) BridgeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "bridgeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Governance *GovernanceSession) BridgeContract() (common.Address, error) {
	return _Governance.Contract.BridgeContract(&_Governance.CallOpts)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Governance *GovernanceCallerSession) BridgeContract() (common.Address, error) {
	return _Governance.Contract.BridgeContract(&_Governance.CallOpts)
}

// BridgeOperatorsVoted is a free data retrieval call binding the contract method 0x2b5df351.
//
// Solidity: function bridgeOperatorsVoted(uint256 _period, uint256 _epoch, address _voter) view returns(bool)
func (_Governance *GovernanceCaller) BridgeOperatorsVoted(opts *bind.CallOpts, _period *big.Int, _epoch *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "bridgeOperatorsVoted", _period, _epoch, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeOperatorsVoted is a free data retrieval call binding the contract method 0x2b5df351.
//
// Solidity: function bridgeOperatorsVoted(uint256 _period, uint256 _epoch, address _voter) view returns(bool)
func (_Governance *GovernanceSession) BridgeOperatorsVoted(_period *big.Int, _epoch *big.Int, _voter common.Address) (bool, error) {
	return _Governance.Contract.BridgeOperatorsVoted(&_Governance.CallOpts, _period, _epoch, _voter)
}

// BridgeOperatorsVoted is a free data retrieval call binding the contract method 0x2b5df351.
//
// Solidity: function bridgeOperatorsVoted(uint256 _period, uint256 _epoch, address _voter) view returns(bool)
func (_Governance *GovernanceCallerSession) BridgeOperatorsVoted(_period *big.Int, _epoch *big.Int, _voter common.Address) (bool, error) {
	return _Governance.Contract.BridgeOperatorsVoted(&_Governance.CallOpts, _period, _epoch, _voter)
}

// EmergencyPollVoted is a free data retrieval call binding the contract method 0xdcc3eb19.
//
// Solidity: function emergencyPollVoted(bytes32 _voteHash, address _voter) view returns(bool)
func (_Governance *GovernanceCaller) EmergencyPollVoted(opts *bind.CallOpts, _voteHash [32]byte, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "emergencyPollVoted", _voteHash, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyPollVoted is a free data retrieval call binding the contract method 0xdcc3eb19.
//
// Solidity: function emergencyPollVoted(bytes32 _voteHash, address _voter) view returns(bool)
func (_Governance *GovernanceSession) EmergencyPollVoted(_voteHash [32]byte, _voter common.Address) (bool, error) {
	return _Governance.Contract.EmergencyPollVoted(&_Governance.CallOpts, _voteHash, _voter)
}

// EmergencyPollVoted is a free data retrieval call binding the contract method 0xdcc3eb19.
//
// Solidity: function emergencyPollVoted(bytes32 _voteHash, address _voter) view returns(bool)
func (_Governance *GovernanceCallerSession) EmergencyPollVoted(_voteHash [32]byte, _voter common.Address) (bool, error) {
	return _Governance.Contract.EmergencyPollVoted(&_Governance.CallOpts, _voteHash, _voter)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x3e2c314f.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch, address[] _voters) view returns((uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCaller) GetBridgeOperatorVotingSignatures(opts *bind.CallOpts, _period *big.Int, _epoch *big.Int, _voters []common.Address) ([]SignatureConsumerSignature, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getBridgeOperatorVotingSignatures", _period, _epoch, _voters)

	if err != nil {
		return *new([]SignatureConsumerSignature), err
	}

	out0 := *abi.ConvertType(out[0], new([]SignatureConsumerSignature)).(*[]SignatureConsumerSignature)

	return out0, err

}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x3e2c314f.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch, address[] _voters) view returns((uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _epoch *big.Int, _voters []common.Address) ([]SignatureConsumerSignature, error) {
	return _Governance.Contract.GetBridgeOperatorVotingSignatures(&_Governance.CallOpts, _period, _epoch, _voters)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x3e2c314f.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch, address[] _voters) view returns((uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCallerSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _epoch *big.Int, _voters []common.Address) ([]SignatureConsumerSignature, error) {
	return _Governance.Contract.GetBridgeOperatorVotingSignatures(&_Governance.CallOpts, _period, _epoch, _voters)
}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_Governance *GovernanceCaller) GetProposalExpiryDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalExpiryDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_Governance *GovernanceSession) GetProposalExpiryDuration() (*big.Int, error) {
	return _Governance.Contract.GetProposalExpiryDuration(&_Governance.CallOpts)
}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetProposalExpiryDuration() (*big.Int, error) {
	return _Governance.Contract.GetProposalExpiryDuration(&_Governance.CallOpts)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x1c905e39.
//
// Solidity: function getProposalSignatures(uint256 _chainId, uint256 _round) view returns(address[] _voters, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCaller) GetProposalSignatures(opts *bind.CallOpts, _chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalSignatures", _chainId, _round)

	outstruct := new(struct {
		Voters     []common.Address
		Supports   []uint8
		Signatures []SignatureConsumerSignature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voters = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Supports = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]SignatureConsumerSignature)).(*[]SignatureConsumerSignature)

	return *outstruct, err

}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x1c905e39.
//
// Solidity: function getProposalSignatures(uint256 _chainId, uint256 _round) view returns(address[] _voters, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceSession) GetProposalSignatures(_chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _Governance.Contract.GetProposalSignatures(&_Governance.CallOpts, _chainId, _round)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x1c905e39.
//
// Solidity: function getProposalSignatures(uint256 _chainId, uint256 _round) view returns(address[] _voters, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCallerSession) GetProposalSignatures(_chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _Governance.Contract.GetProposalSignatures(&_Governance.CallOpts, _chainId, _round)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Governance *GovernanceCaller) GetProxyAdmin(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProxyAdmin", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Governance *GovernanceSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _Governance.Contract.GetProxyAdmin(&_Governance.CallOpts, _proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address _proxy) view returns(address)
func (_Governance *GovernanceCallerSession) GetProxyAdmin(_proxy common.Address) (common.Address, error) {
	return _Governance.Contract.GetProxyAdmin(&_Governance.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Governance *GovernanceCaller) GetProxyImplementation(opts *bind.CallOpts, _proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProxyImplementation", _proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Governance *GovernanceSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _Governance.Contract.GetProxyImplementation(&_Governance.CallOpts, _proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address _proxy) view returns(address)
func (_Governance *GovernanceCallerSession) GetProxyImplementation(_proxy common.Address) (common.Address, error) {
	return _Governance.Contract.GetProxyImplementation(&_Governance.CallOpts, _proxy)
}

// LastSyncedBridgeOperatorSetInfo is a free data retrieval call binding the contract method 0x62e52e5f.
//
// Solidity: function lastSyncedBridgeOperatorSetInfo() view returns((uint256,uint256,address[]))
func (_Governance *GovernanceCaller) LastSyncedBridgeOperatorSetInfo(opts *bind.CallOpts) (BridgeOperatorsBallotBridgeOperatorSet, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "lastSyncedBridgeOperatorSetInfo")

	if err != nil {
		return *new(BridgeOperatorsBallotBridgeOperatorSet), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeOperatorsBallotBridgeOperatorSet)).(*BridgeOperatorsBallotBridgeOperatorSet)

	return out0, err

}

// LastSyncedBridgeOperatorSetInfo is a free data retrieval call binding the contract method 0x62e52e5f.
//
// Solidity: function lastSyncedBridgeOperatorSetInfo() view returns((uint256,uint256,address[]))
func (_Governance *GovernanceSession) LastSyncedBridgeOperatorSetInfo() (BridgeOperatorsBallotBridgeOperatorSet, error) {
	return _Governance.Contract.LastSyncedBridgeOperatorSetInfo(&_Governance.CallOpts)
}

// LastSyncedBridgeOperatorSetInfo is a free data retrieval call binding the contract method 0x62e52e5f.
//
// Solidity: function lastSyncedBridgeOperatorSetInfo() view returns((uint256,uint256,address[]))
func (_Governance *GovernanceCallerSession) LastSyncedBridgeOperatorSetInfo() (BridgeOperatorsBallotBridgeOperatorSet, error) {
	return _Governance.Contract.LastSyncedBridgeOperatorSetInfo(&_Governance.CallOpts)
}

// LastVotedBlock is a free data retrieval call binding the contract method 0x988ef53c.
//
// Solidity: function lastVotedBlock(address _bridgeVoter) view returns(uint256)
func (_Governance *GovernanceCaller) LastVotedBlock(opts *bind.CallOpts, _bridgeVoter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "lastVotedBlock", _bridgeVoter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVotedBlock is a free data retrieval call binding the contract method 0x988ef53c.
//
// Solidity: function lastVotedBlock(address _bridgeVoter) view returns(uint256)
func (_Governance *GovernanceSession) LastVotedBlock(_bridgeVoter common.Address) (*big.Int, error) {
	return _Governance.Contract.LastVotedBlock(&_Governance.CallOpts, _bridgeVoter)
}

// LastVotedBlock is a free data retrieval call binding the contract method 0x988ef53c.
//
// Solidity: function lastVotedBlock(address _bridgeVoter) view returns(uint256)
func (_Governance *GovernanceCallerSession) LastVotedBlock(_bridgeVoter common.Address) (*big.Int, error) {
	return _Governance.Contract.LastVotedBlock(&_Governance.CallOpts, _bridgeVoter)
}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_Governance *GovernanceCaller) ProposalVoted(opts *bind.CallOpts, _chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalVoted", _chainId, _round, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_Governance *GovernanceSession) ProposalVoted(_chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	return _Governance.Contract.ProposalVoted(&_Governance.CallOpts, _chainId, _round, _voter)
}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_Governance *GovernanceCallerSession) ProposalVoted(_chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	return _Governance.Contract.ProposalVoted(&_Governance.CallOpts, _chainId, _round, _voter)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Governance *GovernanceCaller) RoninTrustedOrganizationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "roninTrustedOrganizationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Governance *GovernanceSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Governance.Contract.RoninTrustedOrganizationContract(&_Governance.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Governance *GovernanceCallerSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Governance.Contract.RoninTrustedOrganizationContract(&_Governance.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Governance *GovernanceCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "round", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Governance *GovernanceSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Round(&_Governance.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_Governance *GovernanceCallerSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Round(&_Governance.CallOpts, arg0)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Governance *GovernanceCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Governance *GovernanceSession) ValidatorContract() (common.Address, error) {
	return _Governance.Contract.ValidatorContract(&_Governance.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Governance *GovernanceCallerSession) ValidatorContract() (common.Address, error) {
	return _Governance.Contract.ValidatorContract(&_Governance.CallOpts)
}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight, uint256 expiryTimestamp)
func (_Governance *GovernanceCaller) Vote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "vote", arg0, arg1)

	outstruct := new(struct {
		Status            uint8
		Hash              [32]byte
		AgainstVoteWeight *big.Int
		ForVoteWeight     *big.Int
		ExpiryTimestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Hash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.AgainstVoteWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ForVoteWeight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExpiryTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight, uint256 expiryTimestamp)
func (_Governance *GovernanceSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	return _Governance.Contract.Vote(&_Governance.CallOpts, arg0, arg1)
}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight, uint256 expiryTimestamp)
func (_Governance *GovernanceCallerSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	return _Governance.Contract.Vote(&_Governance.CallOpts, arg0, arg1)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactor) CastGlobalProposalBySignatures(opts *bind.TransactOpts, _globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castGlobalProposalBySignatures", _globalProposal, _supports, _signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceSession) CastGlobalProposalBySignatures(_globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.CastGlobalProposalBySignatures(&_Governance.TransactOpts, _globalProposal, _supports, _signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactorSession) CastGlobalProposalBySignatures(_globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.CastGlobalProposalBySignatures(&_Governance.TransactOpts, _globalProposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactor) CastProposalBySignatures(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castProposalBySignatures", _proposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceSession) CastProposalBySignatures(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.CastProposalBySignatures(&_Governance.TransactOpts, _proposal, _supports, _signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactorSession) CastProposalBySignatures(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.CastProposalBySignatures(&_Governance.TransactOpts, _proposal, _supports, _signatures)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8 _support) returns()
func (_Governance *GovernanceTransactor) CastProposalVoteForCurrentNetwork(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _support uint8) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "castProposalVoteForCurrentNetwork", _proposal, _support)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8 _support) returns()
func (_Governance *GovernanceSession) CastProposalVoteForCurrentNetwork(_proposal ProposalProposalDetail, _support uint8) (*types.Transaction, error) {
	return _Governance.Contract.CastProposalVoteForCurrentNetwork(&_Governance.TransactOpts, _proposal, _support)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8 _support) returns()
func (_Governance *GovernanceTransactorSession) CastProposalVoteForCurrentNetwork(_proposal ProposalProposalDetail, _support uint8) (*types.Transaction, error) {
	return _Governance.Contract.CastProposalVoteForCurrentNetwork(&_Governance.TransactOpts, _proposal, _support)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Governance *GovernanceTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, _proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "changeProxyAdmin", _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Governance *GovernanceSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeProxyAdmin(&_Governance.TransactOpts, _proxy, _newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address _proxy, address _newAdmin) returns()
func (_Governance *GovernanceTransactorSession) ChangeProxyAdmin(_proxy common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.ChangeProxyAdmin(&_Governance.TransactOpts, _proxy, _newAdmin)
}

// CreateEmergencyExitPoll is a paid mutator transaction binding the contract method 0xa2fae570.
//
// Solidity: function createEmergencyExitPoll(address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceTransactor) CreateEmergencyExitPoll(opts *bind.TransactOpts, _consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "createEmergencyExitPoll", _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// CreateEmergencyExitPoll is a paid mutator transaction binding the contract method 0xa2fae570.
//
// Solidity: function createEmergencyExitPoll(address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceSession) CreateEmergencyExitPoll(_consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CreateEmergencyExitPoll(&_Governance.TransactOpts, _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// CreateEmergencyExitPoll is a paid mutator transaction binding the contract method 0xa2fae570.
//
// Solidity: function createEmergencyExitPoll(address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceTransactorSession) CreateEmergencyExitPoll(_consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.CreateEmergencyExitPoll(&_Governance.TransactOpts, _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 chainId, uint256 round) returns()
func (_Governance *GovernanceTransactor) DeleteExpired(opts *bind.TransactOpts, chainId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "deleteExpired", chainId, round)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 chainId, uint256 round) returns()
func (_Governance *GovernanceSession) DeleteExpired(chainId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.DeleteExpired(&_Governance.TransactOpts, chainId, round)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 chainId, uint256 round) returns()
func (_Governance *GovernanceTransactorSession) DeleteExpired(chainId *big.Int, round *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.DeleteExpired(&_Governance.TransactOpts, chainId, round)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceTransactor) Propose(opts *bind.TransactOpts, _chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose", _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceSession) Propose(_chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceTransactorSession) Propose(_chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 _expiryTimestamp, uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceTransactor) ProposeGlobal(opts *bind.TransactOpts, _expiryTimestamp *big.Int, _targetOptions []uint8, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeGlobal", _expiryTimestamp, _targetOptions, _values, _calldatas, _gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 _expiryTimestamp, uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceSession) ProposeGlobal(_expiryTimestamp *big.Int, _targetOptions []uint8, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.ProposeGlobal(&_Governance.TransactOpts, _expiryTimestamp, _targetOptions, _values, _calldatas, _gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 _expiryTimestamp, uint8[] _targetOptions, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_Governance *GovernanceTransactorSession) ProposeGlobal(_expiryTimestamp *big.Int, _targetOptions []uint8, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _Governance.Contract.ProposeGlobal(&_Governance.TransactOpts, _expiryTimestamp, _targetOptions, _values, _calldatas, _gasAmounts)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactor) ProposeGlobalProposalStructAndCastVotes(opts *bind.TransactOpts, _globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeGlobalProposalStructAndCastVotes", _globalProposal, _supports, _signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceSession) ProposeGlobalProposalStructAndCastVotes(_globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.ProposeGlobalProposalStructAndCastVotes(&_Governance.TransactOpts, _globalProposal, _supports, _signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) _globalProposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactorSession) ProposeGlobalProposalStructAndCastVotes(_globalProposal GlobalProposalGlobalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.ProposeGlobalProposalStructAndCastVotes(&_Governance.TransactOpts, _globalProposal, _supports, _signatures)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts, uint8 _support) returns()
func (_Governance *GovernanceTransactor) ProposeProposalForCurrentNetwork(opts *bind.TransactOpts, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int, _support uint8) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeProposalForCurrentNetwork", _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts, _support)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts, uint8 _support) returns()
func (_Governance *GovernanceSession) ProposeProposalForCurrentNetwork(_expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int, _support uint8) (*types.Transaction, error) {
	return _Governance.Contract.ProposeProposalForCurrentNetwork(&_Governance.TransactOpts, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts, _support)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts, uint8 _support) returns()
func (_Governance *GovernanceTransactorSession) ProposeProposalForCurrentNetwork(_expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int, _support uint8) (*types.Transaction, error) {
	return _Governance.Contract.ProposeProposalForCurrentNetwork(&_Governance.TransactOpts, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts, _support)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactor) ProposeProposalStructAndCastVotes(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "proposeProposalStructAndCastVotes", _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.ProposeProposalStructAndCastVotes(&_Governance.TransactOpts, _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactorSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.ProposeProposalStructAndCastVotes(&_Governance.TransactOpts, _proposal, _supports, _signatures)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_Governance *GovernanceTransactor) SetBridgeContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setBridgeContract", _addr)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_Governance *GovernanceSession) SetBridgeContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetBridgeContract(&_Governance.TransactOpts, _addr)
}

// SetBridgeContract is a paid mutator transaction binding the contract method 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _addr) returns()
func (_Governance *GovernanceTransactorSession) SetBridgeContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetBridgeContract(&_Governance.TransactOpts, _addr)
}

// SetProposalExpiryDuration is a paid mutator transaction binding the contract method 0x2e96a6fb.
//
// Solidity: function setProposalExpiryDuration(uint256 _expiryDuration) returns()
func (_Governance *GovernanceTransactor) SetProposalExpiryDuration(opts *bind.TransactOpts, _expiryDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setProposalExpiryDuration", _expiryDuration)
}

// SetProposalExpiryDuration is a paid mutator transaction binding the contract method 0x2e96a6fb.
//
// Solidity: function setProposalExpiryDuration(uint256 _expiryDuration) returns()
func (_Governance *GovernanceSession) SetProposalExpiryDuration(_expiryDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetProposalExpiryDuration(&_Governance.TransactOpts, _expiryDuration)
}

// SetProposalExpiryDuration is a paid mutator transaction binding the contract method 0x2e96a6fb.
//
// Solidity: function setProposalExpiryDuration(uint256 _expiryDuration) returns()
func (_Governance *GovernanceTransactorSession) SetProposalExpiryDuration(_expiryDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetProposalExpiryDuration(&_Governance.TransactOpts, _expiryDuration)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Governance *GovernanceTransactor) SetRoninTrustedOrganizationContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setRoninTrustedOrganizationContract", _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Governance *GovernanceSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRoninTrustedOrganizationContract(&_Governance.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Governance *GovernanceTransactorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRoninTrustedOrganizationContract(&_Governance.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Governance *GovernanceTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Governance *GovernanceSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetValidatorContract(&_Governance.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Governance *GovernanceTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetValidatorContract(&_Governance.TransactOpts, _addr)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x60911e8e.
//
// Solidity: function voteBridgeOperatorsBySignatures((uint256,uint256,address[]) _ballot, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactor) VoteBridgeOperatorsBySignatures(opts *bind.TransactOpts, _ballot BridgeOperatorsBallotBridgeOperatorSet, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "voteBridgeOperatorsBySignatures", _ballot, _signatures)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x60911e8e.
//
// Solidity: function voteBridgeOperatorsBySignatures((uint256,uint256,address[]) _ballot, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceSession) VoteBridgeOperatorsBySignatures(_ballot BridgeOperatorsBallotBridgeOperatorSet, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.VoteBridgeOperatorsBySignatures(&_Governance.TransactOpts, _ballot, _signatures)
}

// VoteBridgeOperatorsBySignatures is a paid mutator transaction binding the contract method 0x60911e8e.
//
// Solidity: function voteBridgeOperatorsBySignatures((uint256,uint256,address[]) _ballot, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_Governance *GovernanceTransactorSession) VoteBridgeOperatorsBySignatures(_ballot BridgeOperatorsBallotBridgeOperatorSet, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _Governance.Contract.VoteBridgeOperatorsBySignatures(&_Governance.TransactOpts, _ballot, _signatures)
}

// VoteEmergencyExit is a paid mutator transaction binding the contract method 0x9e0dc0b3.
//
// Solidity: function voteEmergencyExit(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceTransactor) VoteEmergencyExit(opts *bind.TransactOpts, _voteHash [32]byte, _consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "voteEmergencyExit", _voteHash, _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// VoteEmergencyExit is a paid mutator transaction binding the contract method 0x9e0dc0b3.
//
// Solidity: function voteEmergencyExit(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceSession) VoteEmergencyExit(_voteHash [32]byte, _consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.VoteEmergencyExit(&_Governance.TransactOpts, _voteHash, _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// VoteEmergencyExit is a paid mutator transaction binding the contract method 0x9e0dc0b3.
//
// Solidity: function voteEmergencyExit(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt) returns()
func (_Governance *GovernanceTransactorSession) VoteEmergencyExit(_voteHash [32]byte, _consensusAddr common.Address, _recipientAfterUnlockedFund common.Address, _requestedAt *big.Int, _expiredAt *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.VoteEmergencyExit(&_Governance.TransactOpts, _voteHash, _consensusAddr, _recipientAfterUnlockedFund, _requestedAt, _expiredAt)
}

// GovernanceBridgeContractUpdatedIterator is returned from FilterBridgeContractUpdated and is used to iterate over the raw logs and unpacked data for BridgeContractUpdated events raised by the Governance contract.
type GovernanceBridgeContractUpdatedIterator struct {
	Event *GovernanceBridgeContractUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceBridgeContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceBridgeContractUpdated)
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
		it.Event = new(GovernanceBridgeContractUpdated)
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
func (it *GovernanceBridgeContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceBridgeContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceBridgeContractUpdated represents a BridgeContractUpdated event raised by the Governance contract.
type GovernanceBridgeContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeContractUpdated is a free log retrieval operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) FilterBridgeContractUpdated(opts *bind.FilterOpts) (*GovernanceBridgeContractUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "BridgeContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceBridgeContractUpdatedIterator{contract: _Governance.contract, event: "BridgeContractUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeContractUpdated is a free log subscription operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) WatchBridgeContractUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceBridgeContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "BridgeContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceBridgeContractUpdated)
				if err := _Governance.contract.UnpackLog(event, "BridgeContractUpdated", log); err != nil {
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

// ParseBridgeContractUpdated is a log parse operation binding the contract event 0x5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048.
//
// Solidity: event BridgeContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) ParseBridgeContractUpdated(log types.Log) (*GovernanceBridgeContractUpdated, error) {
	event := new(GovernanceBridgeContractUpdated)
	if err := _Governance.contract.UnpackLog(event, "BridgeContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceBridgeOperatorsApprovedIterator is returned from FilterBridgeOperatorsApproved and is used to iterate over the raw logs and unpacked data for BridgeOperatorsApproved events raised by the Governance contract.
type GovernanceBridgeOperatorsApprovedIterator struct {
	Event *GovernanceBridgeOperatorsApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceBridgeOperatorsApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceBridgeOperatorsApproved)
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
		it.Event = new(GovernanceBridgeOperatorsApproved)
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
func (it *GovernanceBridgeOperatorsApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceBridgeOperatorsApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceBridgeOperatorsApproved represents a BridgeOperatorsApproved event raised by the Governance contract.
type GovernanceBridgeOperatorsApproved struct {
	Period    *big.Int
	Epoch     *big.Int
	Operators []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorsApproved is a free log retrieval operation binding the contract event 0x7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d.
//
// Solidity: event BridgeOperatorsApproved(uint256 _period, uint256 _epoch, address[] _operators)
func (_Governance *GovernanceFilterer) FilterBridgeOperatorsApproved(opts *bind.FilterOpts) (*GovernanceBridgeOperatorsApprovedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "BridgeOperatorsApproved")
	if err != nil {
		return nil, err
	}
	return &GovernanceBridgeOperatorsApprovedIterator{contract: _Governance.contract, event: "BridgeOperatorsApproved", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorsApproved is a free log subscription operation binding the contract event 0x7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d.
//
// Solidity: event BridgeOperatorsApproved(uint256 _period, uint256 _epoch, address[] _operators)
func (_Governance *GovernanceFilterer) WatchBridgeOperatorsApproved(opts *bind.WatchOpts, sink chan<- *GovernanceBridgeOperatorsApproved) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "BridgeOperatorsApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceBridgeOperatorsApproved)
				if err := _Governance.contract.UnpackLog(event, "BridgeOperatorsApproved", log); err != nil {
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

// ParseBridgeOperatorsApproved is a log parse operation binding the contract event 0x7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d.
//
// Solidity: event BridgeOperatorsApproved(uint256 _period, uint256 _epoch, address[] _operators)
func (_Governance *GovernanceFilterer) ParseBridgeOperatorsApproved(log types.Log) (*GovernanceBridgeOperatorsApproved, error) {
	event := new(GovernanceBridgeOperatorsApproved)
	if err := _Governance.contract.UnpackLog(event, "BridgeOperatorsApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceEmergencyExitPollApprovedIterator is returned from FilterEmergencyExitPollApproved and is used to iterate over the raw logs and unpacked data for EmergencyExitPollApproved events raised by the Governance contract.
type GovernanceEmergencyExitPollApprovedIterator struct {
	Event *GovernanceEmergencyExitPollApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceEmergencyExitPollApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceEmergencyExitPollApproved)
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
		it.Event = new(GovernanceEmergencyExitPollApproved)
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
func (it *GovernanceEmergencyExitPollApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceEmergencyExitPollApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceEmergencyExitPollApproved represents a EmergencyExitPollApproved event raised by the Governance contract.
type GovernanceEmergencyExitPollApproved struct {
	VoteHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitPollApproved is a free log retrieval operation binding the contract event 0xd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b.
//
// Solidity: event EmergencyExitPollApproved(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) FilterEmergencyExitPollApproved(opts *bind.FilterOpts) (*GovernanceEmergencyExitPollApprovedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "EmergencyExitPollApproved")
	if err != nil {
		return nil, err
	}
	return &GovernanceEmergencyExitPollApprovedIterator{contract: _Governance.contract, event: "EmergencyExitPollApproved", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitPollApproved is a free log subscription operation binding the contract event 0xd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b.
//
// Solidity: event EmergencyExitPollApproved(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) WatchEmergencyExitPollApproved(opts *bind.WatchOpts, sink chan<- *GovernanceEmergencyExitPollApproved) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "EmergencyExitPollApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceEmergencyExitPollApproved)
				if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollApproved", log); err != nil {
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

// ParseEmergencyExitPollApproved is a log parse operation binding the contract event 0xd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b.
//
// Solidity: event EmergencyExitPollApproved(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) ParseEmergencyExitPollApproved(log types.Log) (*GovernanceEmergencyExitPollApproved, error) {
	event := new(GovernanceEmergencyExitPollApproved)
	if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceEmergencyExitPollCreatedIterator is returned from FilterEmergencyExitPollCreated and is used to iterate over the raw logs and unpacked data for EmergencyExitPollCreated events raised by the Governance contract.
type GovernanceEmergencyExitPollCreatedIterator struct {
	Event *GovernanceEmergencyExitPollCreated // Event containing the contract specifics and raw log

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
func (it *GovernanceEmergencyExitPollCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceEmergencyExitPollCreated)
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
		it.Event = new(GovernanceEmergencyExitPollCreated)
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
func (it *GovernanceEmergencyExitPollCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceEmergencyExitPollCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceEmergencyExitPollCreated represents a EmergencyExitPollCreated event raised by the Governance contract.
type GovernanceEmergencyExitPollCreated struct {
	VoteHash                   [32]byte
	ConsensusAddr              common.Address
	RecipientAfterUnlockedFund common.Address
	RequestedAt                *big.Int
	ExpiredAt                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitPollCreated is a free log retrieval operation binding the contract event 0x18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de8.
//
// Solidity: event EmergencyExitPollCreated(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt)
func (_Governance *GovernanceFilterer) FilterEmergencyExitPollCreated(opts *bind.FilterOpts) (*GovernanceEmergencyExitPollCreatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "EmergencyExitPollCreated")
	if err != nil {
		return nil, err
	}
	return &GovernanceEmergencyExitPollCreatedIterator{contract: _Governance.contract, event: "EmergencyExitPollCreated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitPollCreated is a free log subscription operation binding the contract event 0x18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de8.
//
// Solidity: event EmergencyExitPollCreated(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt)
func (_Governance *GovernanceFilterer) WatchEmergencyExitPollCreated(opts *bind.WatchOpts, sink chan<- *GovernanceEmergencyExitPollCreated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "EmergencyExitPollCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceEmergencyExitPollCreated)
				if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollCreated", log); err != nil {
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

// ParseEmergencyExitPollCreated is a log parse operation binding the contract event 0x18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de8.
//
// Solidity: event EmergencyExitPollCreated(bytes32 _voteHash, address _consensusAddr, address _recipientAfterUnlockedFund, uint256 _requestedAt, uint256 _expiredAt)
func (_Governance *GovernanceFilterer) ParseEmergencyExitPollCreated(log types.Log) (*GovernanceEmergencyExitPollCreated, error) {
	event := new(GovernanceEmergencyExitPollCreated)
	if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceEmergencyExitPollExpiredIterator is returned from FilterEmergencyExitPollExpired and is used to iterate over the raw logs and unpacked data for EmergencyExitPollExpired events raised by the Governance contract.
type GovernanceEmergencyExitPollExpiredIterator struct {
	Event *GovernanceEmergencyExitPollExpired // Event containing the contract specifics and raw log

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
func (it *GovernanceEmergencyExitPollExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceEmergencyExitPollExpired)
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
		it.Event = new(GovernanceEmergencyExitPollExpired)
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
func (it *GovernanceEmergencyExitPollExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceEmergencyExitPollExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceEmergencyExitPollExpired represents a EmergencyExitPollExpired event raised by the Governance contract.
type GovernanceEmergencyExitPollExpired struct {
	VoteHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitPollExpired is a free log retrieval operation binding the contract event 0xeecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe4.
//
// Solidity: event EmergencyExitPollExpired(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) FilterEmergencyExitPollExpired(opts *bind.FilterOpts) (*GovernanceEmergencyExitPollExpiredIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "EmergencyExitPollExpired")
	if err != nil {
		return nil, err
	}
	return &GovernanceEmergencyExitPollExpiredIterator{contract: _Governance.contract, event: "EmergencyExitPollExpired", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitPollExpired is a free log subscription operation binding the contract event 0xeecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe4.
//
// Solidity: event EmergencyExitPollExpired(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) WatchEmergencyExitPollExpired(opts *bind.WatchOpts, sink chan<- *GovernanceEmergencyExitPollExpired) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "EmergencyExitPollExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceEmergencyExitPollExpired)
				if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollExpired", log); err != nil {
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

// ParseEmergencyExitPollExpired is a log parse operation binding the contract event 0xeecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe4.
//
// Solidity: event EmergencyExitPollExpired(bytes32 _voteHash)
func (_Governance *GovernanceFilterer) ParseEmergencyExitPollExpired(log types.Log) (*GovernanceEmergencyExitPollExpired, error) {
	event := new(GovernanceEmergencyExitPollExpired)
	if err := _Governance.contract.UnpackLog(event, "EmergencyExitPollExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceGlobalProposalCreatedIterator is returned from FilterGlobalProposalCreated and is used to iterate over the raw logs and unpacked data for GlobalProposalCreated events raised by the Governance contract.
type GovernanceGlobalProposalCreatedIterator struct {
	Event *GovernanceGlobalProposalCreated // Event containing the contract specifics and raw log

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
func (it *GovernanceGlobalProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceGlobalProposalCreated)
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
		it.Event = new(GovernanceGlobalProposalCreated)
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
func (it *GovernanceGlobalProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceGlobalProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceGlobalProposalCreated represents a GlobalProposalCreated event raised by the Governance contract.
type GovernanceGlobalProposalCreated struct {
	Round              *big.Int
	ProposalHash       [32]byte
	Proposal           ProposalProposalDetail
	GlobalProposalHash [32]byte
	GlobalProposal     GlobalProposalGlobalProposalDetail
	Creator            common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGlobalProposalCreated is a free log retrieval operation binding the contract event 0x771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, bytes32 globalProposalHash, (uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, address creator)
func (_Governance *GovernanceFilterer) FilterGlobalProposalCreated(opts *bind.FilterOpts, round []*big.Int, proposalHash [][32]byte) (*GovernanceGlobalProposalCreatedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceGlobalProposalCreatedIterator{contract: _Governance.contract, event: "GlobalProposalCreated", logs: logs, sub: sub}, nil
}

// WatchGlobalProposalCreated is a free log subscription operation binding the contract event 0x771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, bytes32 globalProposalHash, (uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, address creator)
func (_Governance *GovernanceFilterer) WatchGlobalProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernanceGlobalProposalCreated, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceGlobalProposalCreated)
				if err := _Governance.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
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

// ParseGlobalProposalCreated is a log parse operation binding the contract event 0x771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, bytes32 globalProposalHash, (uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, address creator)
func (_Governance *GovernanceFilterer) ParseGlobalProposalCreated(log types.Log) (*GovernanceGlobalProposalCreated, error) {
	event := new(GovernanceGlobalProposalCreated)
	if err := _Governance.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the Governance contract.
type GovernanceProposalApprovedIterator struct {
	Event *GovernanceProposalApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalApproved)
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
		it.Event = new(GovernanceProposalApproved)
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
func (it *GovernanceProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalApproved represents a ProposalApproved event raised by the Governance contract.
type GovernanceProposalApproved struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) FilterProposalApproved(opts *bind.FilterOpts, proposalHash [][32]byte) (*GovernanceProposalApprovedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalApproved", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalApprovedIterator{contract: _Governance.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *GovernanceProposalApproved, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalApproved", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalApproved)
				if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
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

// ParseProposalApproved is a log parse operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) ParseProposalApproved(log types.Log) (*GovernanceProposalApproved, error) {
	event := new(GovernanceProposalApproved)
	if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Governance contract.
type GovernanceProposalCreatedIterator struct {
	Event *GovernanceProposalCreated // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalCreated)
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
		it.Event = new(GovernanceProposalCreated)
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
func (it *GovernanceProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalCreated represents a ProposalCreated event raised by the Governance contract.
type GovernanceProposalCreated struct {
	ChainId      *big.Int
	Round        *big.Int
	ProposalHash [32]byte
	Proposal     ProposalProposalDetail
	Creator      common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, address creator)
func (_Governance *GovernanceFilterer) FilterProposalCreated(opts *bind.FilterOpts, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (*GovernanceProposalCreatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalCreatedIterator{contract: _Governance.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, address creator)
func (_Governance *GovernanceFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GovernanceProposalCreated, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalCreated)
				if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, address creator)
func (_Governance *GovernanceFilterer) ParseProposalCreated(log types.Log) (*GovernanceProposalCreated, error) {
	event := new(GovernanceProposalCreated)
	if err := _Governance.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governance contract.
type GovernanceProposalExecutedIterator struct {
	Event *GovernanceProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecuted)
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
		it.Event = new(GovernanceProposalExecuted)
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
func (it *GovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecuted represents a ProposalExecuted event raised by the Governance contract.
type GovernanceProposalExecuted struct {
	ProposalHash [32]byte
	SuccessCalls []bool
	ReturnDatas  [][]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0xe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash, bool[] successCalls, bytes[] returnDatas)
func (_Governance *GovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalHash [][32]byte) (*GovernanceProposalExecutedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecuted", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutedIterator{contract: _Governance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0xe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash, bool[] successCalls, bytes[] returnDatas)
func (_Governance *GovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecuted, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecuted", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0xe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash, bool[] successCalls, bytes[] returnDatas)
func (_Governance *GovernanceFilterer) ParseProposalExecuted(log types.Log) (*GovernanceProposalExecuted, error) {
	event := new(GovernanceProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExpiredIterator is returned from FilterProposalExpired and is used to iterate over the raw logs and unpacked data for ProposalExpired events raised by the Governance contract.
type GovernanceProposalExpiredIterator struct {
	Event *GovernanceProposalExpired // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExpired)
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
		it.Event = new(GovernanceProposalExpired)
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
func (it *GovernanceProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExpired represents a ProposalExpired event raised by the Governance contract.
type GovernanceProposalExpired struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalExpired is a free log retrieval operation binding the contract event 0x58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a55.
//
// Solidity: event ProposalExpired(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) FilterProposalExpired(opts *bind.FilterOpts, proposalHash [][32]byte) (*GovernanceProposalExpiredIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExpired", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExpiredIterator{contract: _Governance.contract, event: "ProposalExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExpired is a free log subscription operation binding the contract event 0x58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a55.
//
// Solidity: event ProposalExpired(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) WatchProposalExpired(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExpired, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExpired", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExpired)
				if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
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

// ParseProposalExpired is a log parse operation binding the contract event 0x58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a55.
//
// Solidity: event ProposalExpired(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) ParseProposalExpired(log types.Log) (*GovernanceProposalExpired, error) {
	event := new(GovernanceProposalExpired)
	if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the Governance contract.
type GovernanceProposalRejectedIterator struct {
	Event *GovernanceProposalRejected // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalRejected)
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
		it.Event = new(GovernanceProposalRejected)
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
func (it *GovernanceProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalRejected represents a ProposalRejected event raised by the Governance contract.
type GovernanceProposalRejected struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) FilterProposalRejected(opts *bind.FilterOpts, proposalHash [][32]byte) (*GovernanceProposalRejectedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalRejectedIterator{contract: _Governance.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *GovernanceProposalRejected, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalRejected)
				if err := _Governance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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

// ParseProposalRejected is a log parse operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_Governance *GovernanceFilterer) ParseProposalRejected(log types.Log) (*GovernanceProposalRejected, error) {
	event := new(GovernanceProposalRejected)
	if err := _Governance.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Governance contract.
type GovernanceProposalVotedIterator struct {
	Event *GovernanceProposalVoted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVoted)
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
		it.Event = new(GovernanceProposalVoted)
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
func (it *GovernanceProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVoted represents a ProposalVoted event raised by the Governance contract.
type GovernanceProposalVoted struct {
	ProposalHash [32]byte
	Voter        common.Address
	Support      uint8
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_Governance *GovernanceFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalHash [][32]byte, voter []common.Address) (*GovernanceProposalVotedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVotedIterator{contract: _Governance.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_Governance *GovernanceFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVoted, proposalHash [][32]byte, voter []common.Address) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVoted)
				if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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

// ParseProposalVoted is a log parse operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_Governance *GovernanceFilterer) ParseProposalVoted(log types.Log) (*GovernanceProposalVoted, error) {
	event := new(GovernanceProposalVoted)
	if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceRoninTrustedOrganizationContractUpdatedIterator is returned from FilterRoninTrustedOrganizationContractUpdated and is used to iterate over the raw logs and unpacked data for RoninTrustedOrganizationContractUpdated events raised by the Governance contract.
type GovernanceRoninTrustedOrganizationContractUpdatedIterator struct {
	Event *GovernanceRoninTrustedOrganizationContractUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceRoninTrustedOrganizationContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRoninTrustedOrganizationContractUpdated)
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
		it.Event = new(GovernanceRoninTrustedOrganizationContractUpdated)
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
func (it *GovernanceRoninTrustedOrganizationContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRoninTrustedOrganizationContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRoninTrustedOrganizationContractUpdated represents a RoninTrustedOrganizationContractUpdated event raised by the Governance contract.
type GovernanceRoninTrustedOrganizationContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninTrustedOrganizationContractUpdated is a free log retrieval operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) FilterRoninTrustedOrganizationContractUpdated(opts *bind.FilterOpts) (*GovernanceRoninTrustedOrganizationContractUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceRoninTrustedOrganizationContractUpdatedIterator{contract: _Governance.contract, event: "RoninTrustedOrganizationContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninTrustedOrganizationContractUpdated is a free log subscription operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) WatchRoninTrustedOrganizationContractUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceRoninTrustedOrganizationContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRoninTrustedOrganizationContractUpdated)
				if err := _Governance.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseRoninTrustedOrganizationContractUpdated(log types.Log) (*GovernanceRoninTrustedOrganizationContractUpdated, error) {
	event := new(GovernanceRoninTrustedOrganizationContractUpdated)
	if err := _Governance.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Governance contract.
type GovernanceValidatorContractUpdatedIterator struct {
	Event *GovernanceValidatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceValidatorContractUpdated)
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
		it.Event = new(GovernanceValidatorContractUpdated)
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
func (it *GovernanceValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Governance contract.
type GovernanceValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*GovernanceValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceValidatorContractUpdatedIterator{contract: _Governance.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceValidatorContractUpdated)
				if err := _Governance.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
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

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Governance *GovernanceFilterer) ParseValidatorContractUpdated(log types.Log) (*GovernanceValidatorContractUpdated, error) {
	event := new(GovernanceValidatorContractUpdated)
	if err := _Governance.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
