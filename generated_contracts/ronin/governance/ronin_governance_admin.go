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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roninChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalExpiryDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeBridgeContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeRoninTrustedOrgContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitPollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalProposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GlobalProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"successCalls\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"returnDatas\",\"type\":\"bytes[]\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"bridgeOperatorsVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"castProposalVoteForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"createEmergencyExitPoll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"deleteExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"emergencyPollVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"getBridgeOperatorVotingSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getProposalSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSyncedBridgeOperatorSetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeVoter\",\"type\":\"address\"}],\"name\":\"lastVotedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"proposalVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"_targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"proposeGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeGlobalProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"proposeProposalForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryDuration\",\"type\":\"uint256\"}],\"name\":\"setProposalExpiryDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumVoteStatusConsumer.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"againstVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"_ballot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteBridgeOperatorsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"voteEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620060a7380380620060a7833981016040819052620000349162000280565b84848483806200004381600255565b506005849055604080516020808201839052601660608301527f524f4e494e5f474f5645524e414e43455f41444d494e000000000000000000006080808401919091528284018890528351808403909101815260a0830184528051908201207f599a80fcaa47b95e2323ab4d34d34e0cc9feda4b843edafcc30c7bdf60ea15bf60c08401527f7e7935007966eb860f4a2ee3dcc9fd53fb3205ce2aa86b0126d4893d4d4c14b960e08401527fad7c5bef027816a800da1736444fb58a807ef4c9603b7848673f7e3a68eb14a561010084015261012080840191909152835180840390910181526101409092019092528051910120600655620001458362000170565b6200015082620001c5565b5050505062000165826200021460201b60201c565b5050505050620002de565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7906020015b60405180910390a150565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae04890602001620001ba565b600d80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b683616990602001620001ba565b80516001600160a01b03811681146200027b57600080fd5b919050565b600080600080600060a086880312156200029957600080fd5b85519450620002ab6020870162000263565b9350620002bb6040870162000263565b9250620002cb6060870162000263565b9150608086015190509295509295909350565b615db980620002ee6000396000f3fe608060405234801561001057600080fd5b50600436106102055760003560e01c8063663ac0111161011a578063a8a0e32c116100ad578063cd5965831161007c578063cd596583146104d5578063cdf64a76146104e6578063dcc3eb19146104f9578063f3b7dead1461050c578063fb4f63711461051f57600080fd5b8063a8a0e32c1461044c578063b384abef1461045f578063b5e337de146104ba578063bc96180b146104cd57600080fd5b80639a7d3382116100e95780639a7d3382146104005780639e0dc0b314610413578063a1819f9a14610426578063a2fae5701461043957600080fd5b8063663ac011146103a05780637eff275e146103b3578063988ef53c146103c657806399439089146103ef57600080fd5b80632b5df3511161019d57806334d5f37b1161016c57806334d5f37b1461033e5780633644e5151461035e5780635511cde11461036757806360911e8e1461037857806362e52e5f1461038b57600080fd5b80632b5df351146102e25780632c5e6520146103055780632e96a6fb146103185780632faf925d1461032b57600080fd5b806317ce2dd4116101d957806317ce2dd4146102585780631c905e39146102745780631e23e04814610296578063204e1c7a146102b757600080fd5b80624054b81461020a57806309fcd8c71461021f5780630b26cf66146102325780630b88183014610245575b600080fd5b61021d6102183660046148cf565b610532565b005b61021d61022d366004614963565b61057b565b61021d610240366004614a47565b610644565b61021d6102533660046148cf565b610699565b61026160055481565b6040519081526020015b60405180910390f35b610287610282366004614a64565b6106a9565b60405161026b93929190614b39565b6102a96102a4366004614a64565b610aa6565b60405161026b929190614bac565b6102ca6102c5366004614a47565b610c2d565b6040516001600160a01b03909116815260200161026b565b6102f56102f0366004614bd1565b610d1f565b604051901515815260200161026b565b6102f5610313366004614bd1565b610d54565b61021d610326366004614c0a565b610d8a565b61021d610339366004614c23565b610db2565b61026161034c366004614c0a565b60006020819052908152604090205481565b61026160065481565b6003546001600160a01b03166102ca565b61021d610386366004614c7c565b610def565b610393610eaf565b60405161026b9190614cec565b61021d6103ae366004614d33565b610f51565b61021d6103c1366004614e10565b61105d565b6102616103d4366004614a47565b6001600160a01b03166000908152600b602052604090205490565b600d546001600160a01b03166102ca565b61021d61040e366004614a64565b61118e565b61021d610421366004614e49565b61119c565b61021d610434366004614e9b565b611402565b61021d610447366004614f74565b6114e8565b61021d61045a366004614fba565b6115b4565b6104a961046d366004614a64565b600160208181526000938452604080852090915291835291208054918101546002820154600383015460069093015460ff909416939192909185565b60405161026b959493929190615000565b61021d6104c8366004614a47565b6115ef565b610261611641565b6004546001600160a01b03166102ca565b61021d6104f4366004614a47565b611651565b6102f5610507366004615035565b6116e5565b6102ca61051a366004614a47565b611717565b61021d61052d366004614c23565b6117e4565b600061053d33611843565b116105635760405162461bcd60e51b815260040161055a9061505a565b60405180910390fd5b6105748585858585600654336119b7565b5050505050565b600061058633611843565b116105a35760405162461bcd60e51b815260040161055a9061505a565b6106388989898989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506105e992508a91508b9050615229565b8787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506106269250610de0915050565b6004546001600160a01b031633611a21565b50505050505050505050565b3330146106635760405162461bcd60e51b815260040161055a90615236565b6000816001600160a01b03163b1161068d5760405162461bcd60e51b815260040161055a9061527d565b61069681611b56565b50565b6105748585858585600654611bab565b600082815260016020908152604080832084845290915281206004810154600582015460609384938493909290916106e182846152d7565b9050806001600160401b038111156106fb576106fb6150a6565b604051908082528060200260200182016040528015610724578160200160208202803683370190505b509550806001600160401b0381111561073f5761073f6150a6565b60405190808252806020026020018201604052801561078a57816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161075d5790505b509450806001600160401b038111156107a5576107a56150a6565b6040519080825280602002602001820160405280156107ce578160200160208202803683370190505b50965060005b8381101561092d5760008782815181106107f0576107f06152ea565b6020026020010190600181111561080957610809614aca565b9081600181111561081c5761081c614aca565b90525060008a81526001602090815260408083208c845290915281206004870180546007909201929184908110610855576108556152ea565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286518790839081106108b3576108b36152ea565b60200260200101819052508460040181815481106108d3576108d36152ea565b9060005260206000200160009054906101000a90046001600160a01b0316888281518110610903576109036152ea565b6001600160a01b03909216602092830291909101909101528061092581615300565b9150506107d4565b5060005b82811015610a9a5760018761094686846152d7565b81518110610956576109566152ea565b6020026020010190600181111561096f5761096f614aca565b9081600181111561098257610982614aca565b90525060008a81526001602090815260408083208c8452909152812060058701805460079092019291849081106109bb576109bb6152ea565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286610a1086846152d7565b81518110610a2057610a206152ea565b6020026020010181905250846005018181548110610a4057610a406152ea565b6000918252602090912001546001600160a01b031688610a6086846152d7565b81518110610a7057610a706152ea565b6001600160a01b039092166020928302919091019091015280610a9281615300565b915050610931565b50505050509250925092565b6000828152600c60209081526040808320848452825291829020600181018054845181850281018501909552808552606094859490929190830182828015610b1757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610af9575b5050505050925082516001600160401b03811115610b3757610b376150a6565b604051908082528060200260200182016040528015610b8257816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610b555790505b50915060005b8351811015610c2457816000016000858381518110610ba957610ba96152ea565b6020908102919091018101516001600160a01b031682528181019290925260409081016000208151606081018352815460ff16815260018201549381019390935260020154908201528351849083908110610c0657610c066152ea565b60200260200101819052508080610c1c90615300565b915050610b88565b50509250929050565b6000806000836001600160a01b0316604051610c5390635c60da1b60e01b815260040190565b600060405180830381855afa9150503d8060008114610c8e576040519150601f19603f3d011682016040523d82523d6000602084013e610c93565b606091505b509150915081610d035760405162461bcd60e51b815260206004820152603560248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060696d706044820152741b195b595b9d185d1a5bdb8a0a580819985a5b1959605a1b606482015260840161055a565b80806020019051810190610d179190615319565b949350505050565b6000838152600a6020908152604080832085845282528083206001600160a01b03851684526002019091528120541515610d17565b600083815260016020908152604080832085845282528083206001600160a01b038516845260080190915281205460ff16610d17565b333014610da95760405162461bcd60e51b815260040161055a90615236565b61069681600255565b6105748585858585600654610dcf6003546001600160a01b031690565b6004546001600160a01b0316611c32565b6003546001600160a01b031690565b610e05838383610dfd611ce6565b600654611e52565b82356000908152600a6020908152604080832082870135845290915290206001815460ff166004811115610e3b57610e3b614aca565b03610ea957836007610e4d82826153a3565b507f7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d905084356020860135610e856040880188615336565b604051610e959493929190615476565b60405180910390a1805460ff191660021781555b50505050565b610ed360405180606001604052806000815260200160008152602001606081525090565b60408051606081018252600780548252600854602080840191909152600980548551818402810184018752818152949593949386019392830182828015610f4357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f25575b505050505081525050905090565b6000610f5c33611843565b11610f795760405162461bcd60e51b815260040161055a9061505a565b60003390506000611042468d8d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061100792508d91508e9050615229565b8a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508b92506121a6915050565b905061104f8282856122c7565b505050505050505050505050565b33301461107c5760405162461bcd60e51b815260040161055a90615236565b604080516001600160a01b0383811660248084019190915283518084039091018152604490920183526020820180516001600160e01b03166308f2839760e41b17905291516000928516916110d0916154f9565b6000604051808303816000865af19150503d806000811461110d576040519150601f19603f3d011682016040523d82523d6000602084013e611112565b606091505b50509050806111895760405162461bcd60e51b815260206004820152603960248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206063686160448201527f6e676541646d696e28616464726573732960206661696c656400000000000000606482015260840161055a565b505050565b6111988282612420565b5050565b3360006111a882611843565b9050600081116111ca5760405162461bcd60e51b815260040161055a9061505a565b60006111d887878787612442565b90508088146112395760405162461bcd60e51b815260206004820152602760248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c696420766f6044820152660e8ca40d0c2e6d60cb1b606482015260840161055a565b6000818152600e6020526040902060058101546112b25760405162461bcd60e51b815260206004820152603160248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f72206044820152706e6f6e2d6578697374656e7420766f746560781b606482015260840161055a565b6004815460ff1660048111156112ca576112ca614aca565b0361132c5760405162461bcd60e51b815260206004820152602c60248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f722060448201526b6578706972656420766f746560a01b606482015260840161055a565b600061134282868661133c611ce6565b876124b2565b9050600181600481111561135857611358614aca565b036113aa5761136789896125e1565b6040518381527fd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b9060200160405180910390a1815460ff19166002178255610638565b60048160048111156113be576113be614aca565b03610638576040518381527feecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe49060200160405180910390a150505050505050505050565b600061140d33611843565b1161142a5760405162461bcd60e51b815260040161055a9061505a565b6114db8a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152506114a092508a91508b9050615229565b8787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152503392506121a6915050565b5050505050505050505050565b336114fb600d546001600160a01b031690565b6001600160a01b03161461152257604051630e6444a160e31b815260040160405180910390fd5b600061153085858585612442565b6000818152600e60209081526040918290204260058201556004810186905582518481526001600160a01b038a811693820193909352918816828401526060820187905260808201869052915192935090917f18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de89181900360a00190a1505050505050565b60006115bf33611843565b116115dc5760405162461bcd60e51b815260040161055a9061505a565b611198336115e9846155ff565b836122c7565b33301461160e5760405162461bcd60e51b815260040161055a90615236565b6000816001600160a01b03163b116116385760405162461bcd60e51b815260040161055a9061527d565b6106968161277d565b600061164c60025490565b905090565b3330146116705760405162461bcd60e51b815260040161055a90615236565b6000816001600160a01b03163b116116dc5760405162461bcd60e51b815260206004820152602960248201527f526f6e696e476f7665726e616e636541646d696e3a2073657420746f206e6f6e6044820152680b58dbdb9d1c9858dd60ba1b606482015260840161055a565b610696816127cb565b6000828152600e602090815260408083206001600160a01b038516845260020190915281205415155b90505b92915050565b6000806000836001600160a01b031660405161173d906303e1469160e61b815260040190565b600060405180830381855afa9150503d8060008114611778576040519150601f19603f3d011682016040523d82523d6000602084013e61177d565b606091505b509150915081610d035760405162461bcd60e51b815260206004820152602c60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206061646d60448201526b1a5b8a0a580819985a5b195960a21b606482015260840161055a565b60006117ef33611843565b1161180c5760405162461bcd60e51b815260040161055a9061505a565b61183b85858585856006546118296003546001600160a01b031690565b6004546001600160a01b031633612819565b505050505050565b600080600061185a6003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316631af0725f60e31b1790529251931692634bb5274a926118b29291016156ff565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516118eb91906154f9565b600060405180830381855afa9150503d8060008114611926576040519150601f19603f3d011682016040523d82523d6000602084013e61192b565b606091505b5091509150816119a35760405162461bcd60e51b815260206004820152603f60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f476f7665726e6f7257656967687428616464726573732960206661696c656400606482015260840161055a565b80806020019051810190610d179190615712565b6119c96119c3886155ff565b8261287c565b5060006119dd6119d8896155ff565b612964565b9050611a176119eb896155ff565b88888888611a03896119fe896000612b03565b612b59565b611a128a6119fe8a6001612b03565b612b80565b5050505050505050565b6040805160c08101909152600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb554909182918190611a699060016152d7565b81526020018c81526020018b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250938552505050602082018b9052604082018a90526060909101889052909150611acd828787612eb9565b9050611ae46002548261308790919063ffffffff16565b6000611aef82612964565b9050611afd6000828f613183565b935080847f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca84611b2c876132b1565b878a604051611b3e9493929190615830565b60405180910390a35050509998505050505050505050565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048906020015b60405180910390a150565b6000611bb96119d8886155ff565b6020808901356000908152600180835260408083208c35845290935291902001549091508114611bfb5760405162461bcd60e51b815260040161055a90615917565b611c29611c07886155ff565b87878787611c1a886119fe896000612b03565b611a12896119fe8a6001612b03565b50505050505050565b6000611c498383611c428c6159ca565b9190612eb9565b90506000611c5e611c598b6159ca565b6132b1565b9050611c6982612964565b600080805260016020818152855183527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4990526040909120015414611cc05760405162461bcd60e51b815260040161055a90615917565b610638828a8a8a8a611cd78b6119fe896000612b03565b611a128c6119fe8a6001612b03565b6000806000611cfd6003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b0316637de5dedd60e01b17905291516001600160a01b039390931692634bb5274a92611d469291016156ff565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051611d7f91906154f9565b600060405180830381855afa9150503d8060008114611dba576040519150601f19603f3d011682016040523d82523d6000602084013e611dbf565b606091505b509150915081611e375760405162461bcd60e51b815260206004820152603860248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c20606d696e60448201527f696d756d566f7465576569676874282960206661696c65640000000000000000606482015260840161055a565b80806020019051810190611e4b9190615712565b9250505090565b600754853510801590611e6b5750600854602086013510155b611edd5760405162461bcd60e51b815260206004820152603d60248201527f424f73476f7665726e616e636550726f706f73616c3a20717565727920666f7260448201527f206f7574646174656420627269646765206f70657261746f7220736574000000606482015260840161055a565b611ee68561340b565b82611f475760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b606482015260840161055a565b6000806000611f55886135cb565b90506000611f638583612b59565b89356000818152600a60209081526040808320828f0135808552908352818420948452600c83528184209084529091528120929350909190805b8a81101561214657368c8c83818110611fb857611fb86152ea565b606002919091019150611fe3905086611fd46020840184615aa3565b8360200135846040013561367d565b9850886001600160a01b0316886001600160a01b03161061205a5760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526a34b3b732b91037b93232b960a91b606482015260840161055a565b889750506000612069896136a5565b90508015612133576001600160a01b0389166000908152600b60205260409020439055600192508c8c838181106120a2576120a26152ea565b6001600160a01b038c1660009081526020889052604090206060909102929092019190506120d08282615ac0565b5050600184810180548083018255600091825260209091200180546001600160a01b0319166001600160a01b038c1617905561210f868b848f8c6124b2565b600481111561212057612120614aca565b0361213357505050505050505050610574565b508061213e81615300565b915050611f9d565b508061104f5760405162461bcd60e51b815260206004820152602960248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526869676e61747572657360b81b606482015260840161055a565b6121ae6147b9565b876000036121fe5760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e206964604482015260640161055a565b6040805160e08101825260008a81526020819052919091205481906122249060016152d7565b81526020018981526020018881526020018781526020018681526020018581526020018481525090506122626002548261308790919063ffffffff16565b600061226d82612964565b9050600061227c8a838b613183565b905081818b7fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd986886040516122b2929190615af1565b60405180910390a45050979650505050505050565b468260200151146123295760405162461bcd60e51b815260206004820152602660248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c6964206368604482015265185a5b881a5960d21b606482015260840161055a565b61233282612964565b602080840151600090815260018083526040808320875184529093529190200154146123bd5760405162461bcd60e51b815260206004820152603460248201527f526f6e696e476f7665726e616e636541646d696e3a206361737420766f746520604482015273199bdc881a5b9d985b1a59081c1c9bdc1bdcd85b60621b606482015260840161055a565b60006123c7611ce6565b90506000816123d4613810565b6123de9190615b1b565b6123e99060016152d7565b6040805160608101825260008082526020820181905291810191909152909150611c29858585858a8661241b8d611843565b613957565b60008281526001602090815260408083208484529091529020610ea981613d72565b604080517f697acba4deaf1a718d8c2d93e42860488cb7812696f28ca10eed17bac41e70276020808301919091526001600160a01b0396871682840152949095166060860152608085019290925260a0808501919091528151808503909101815260c09093019052815191012090565b60008086600401541180156124cb575042866004015411155b156124e35750845460ff1916600490811786556125d8565b6001600160a01b03851660009081526002870160205260409020541561254d57612517856001600160a01b03166014613f94565b6040516020016125279190615b2e565b60408051601f198184030181529082905262461bcd60e51b825261055a916004016156ff565b6001600160a01b038516600090815260028701602090815260408083208590558483526003890190915281208054869190839061258b9084906152d7565b92505081905590508381101580156125b857506000875460ff1660048111156125b6576125b6614aca565b145b156125d057865460ff19166001908117885587018390555b5050845460ff165b95945050505050565b60006125f5600d546001600160a01b031690565b604080516001600160a01b0386811660248084019190915286821660448085019190915284518085039091018152606490930184526020830180516001600160e01b03166361e45aeb60e11b1790529251931692634bb5274a9261265a9291016156ff565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161269391906154f9565b6000604051808303816000865af19150503d80600081146126d0576040519150601f19603f3d011682016040523d82523d6000602084013e6126d5565b606091505b50509050806111895760405162461bcd60e51b815260206004820152606260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206065786560448201527f6352656c656173654c6f636b656446756e64466f72456d657267656e6379457860648201527f69745265717565737428616464726573732c616464726573732960206661696c608482015261195960f21b60a482015260c40161055a565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d790602001611ba0565b600d80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b683616990602001611ba0565b6128216147b9565b61283561282d8b6159ca565b85858561412f565b5090506000612846611c598c6159ca565b905061286e828b8b8b8b61285f8c6119fe896000612b03565b611a128d6119fe8a6001612b03565b509998505050505050505050565b60208201516000908082036128d35760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e206964604482015260640161055a565b6002546128e1908590613087565b60006128ec85612964565b90506128fd82828760400151613183565b855190935083146129205760405162461bcd60e51b815260040161055a90615b82565b8083837fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd98888604051612954929190615af1565b60405180910390a4505092915050565b6000806000806000808660800151905060008760600151905060008860a00151516001600160401b0381111561299c5761299c6150a6565b6040519080825280602002602001820160405280156129c5578160200160208202803683370190505b5060c08a015190915060005b8251811015612a2e578a60a0015181815181106129f0576129f06152ea565b602002602001015180519060200120838281518110612a1157612a116152ea565b602090810291909101015280612a2681615300565b9150506129d1565b506020835102602084012097506020845102602085012096506020825102602083012095506020815102602082012094507fd051578048e6ff0bbc9fca3b65a42088dbde10f36ca841de566711087ad9b08a60001b8a600001518b602001518c604001518b8b8b8b604051602001612ade989796959493929190978852602088019690965260408701949094526060860192909252608085015260a084015260c083015260e08201526101000190565b6040516020818303038152906040528051906020012098505050505050505050919050565b604051600090612b3b907fd900570327c4c0df8dd6bdd522b7da7e39145dd049d2fd4602276adcd511e3c29085908590602001615bc8565b60405160208183030381529060405280519060200120905092915050565b60405161190160f01b60208201526022810183905260428101829052600090606201612b3b565b8415801590612b8e57508483145b612beb5760405162461bcd60e51b815260206004820152602860248201527f476f7665726e616e636550726f706f73616c3a20696e76616c696420617272616044820152670f240d8cadccee8d60c31b606482015260840161055a565b6000612bf5611ce6565b9050600081612c02613810565b612c0c9190615b1b565b612c179060016152d7565b9050600080366000805b89811015612e4d578a8a82818110612c3b57612c3b6152ea565b606002919091019350600090508d8d83818110612c5a57612c5a6152ea565b9050602002016020810190612c6f9190615bed565b6001811115612c8057612c80614aca565b03612cac57612ca589612c966020860186615aa3565b8560200135866040013561367d565b9350612d60565b60018d8d83818110612cc057612cc06152ea565b9050602002016020810190612cd59190615bed565b6001811115612ce657612ce6614aca565b03612cfc57612ca588612c966020860186615aa3565b60405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636550726f706f73616c3a20717565727920666f7220756e604482015272737570706f7274656420766f7465207479706560681b606482015260840161055a565b836001600160a01b0316856001600160a01b031610612dcb5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964206f7264656044820152603960f91b606482015260840161055a565b8394506000612dd985611843565b90508015612e3a5760019250612e288f8f8f85818110612dfb57612dfb6152ea565b9050602002016020810190612e109190615bed565b8a8a89612e22368b90038b018b615c0a565b87613957565b15612e3a575050505050505050611c29565b5080612e4581615300565b915050612c21565b5080612eaa5760405162461bcd60e51b815260206004820152602660248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964207369676e60448201526561747572657360d01b606482015260840161055a565b50505050505050505050505050565b612ec16147b9565b83518152602080850151604080840191909152600091830191909152840151516001600160401b03811115612ef857612ef86150a6565b604051908082528060200260200182016040528015612f21578160200160208202803683370190505b5060608083019190915284015160808083019190915284015160a08083019190915284015160c082015260005b84604001515181101561307f57600185604001518281518110612f7357612f736152ea565b60200260200101516001811115612f8c57612f8c614aca565b03612fcd578282606001518281518110612fa857612fa86152ea565b60200260200101906001600160a01b031690816001600160a01b03168152505061306d565b600085604001518281518110612fe557612fe56152ea565b60200260200101516001811115612ffe57612ffe614aca565b0361301a578382606001518281518110612fa857612fa86152ea565b60405162461bcd60e51b815260206004820152602260248201527f476c6f62616c50726f706f73616c3a20756e737570706f727465642074617267604482015261195d60f21b606482015260840161055a565b8061307781615300565b915050612f4e565b509392505050565b60008260600151511180156130a55750816080015151826060015151145b80156130ba57508160a0015151826060015151145b80156130cf57508160c0015151826060015151145b61311b5760405162461bcd60e51b815260206004820152601e60248201527f50726f706f73616c3a20696e76616c6964206172726179206c656e6774680000604482015260640161055a565b61312581426152d7565b826040015111156111985760405162461bcd60e51b815260206004820152602260248201527f50726f706f73616c3a20696e76616c6964206578706972792074696d6573746160448201526106d760f41b606482015260840161055a565b600083815260208190526040812054908190036131b457506000838152602081905260409020600190819055613284565b60008481526001602090815260408083208484529091528120906131d782613d72565b905080613281576000825460ff1660048111156131f6576131f6614aca565b0361325d5760405162461bcd60e51b815260206004820152603160248201527f436f7265476f7665726e616e63653a2063757272656e742070726f706f73616c604482015270081a5cc81b9bdd0818dbdb5c1b195d1959607a1b606482015260840161055a565b6000868152602081905260408120805490919061327990615300565b918290555092505b50505b60009384526001602081815260408087208488529091529094209384019290925560069092019190915590565b6000806000806000808660600151905060008760400151905060008860800151516001600160401b038111156132e9576132e96150a6565b604051908082528060200260200182016040528015613312578160200160208202803683370190505b5060a08a015190915060005b825181101561337b578a60800151818151811061333d5761333d6152ea565b60200260200101518051906020012083828151811061335e5761335e6152ea565b60209081029190910101528061337381615300565b91505061331e565b5082516020908102848201208551820286830120845183028584012084518402858501208e518f860151604080517f1463f426c05aff2c1a7a0957a71c9898bc8b47142540538e79ee25ee911413509881019890985287019190915260608601526080850184905260a0850183905260c0850182905260e08501819052929b509099509750955061010001612ade565b600061341a6040830183615336565b90501161347d5760405162461bcd60e51b815260206004820152602b60248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b606482015260840161055a565b600061348c6040830183615336565b600081811061349d5761349d6152ea565b90506020020160208101906134b29190614a47565b905060015b6134c46040840184615336565b9050811015611189576134da6040840184615336565b828181106134ea576134ea6152ea565b90506020020160208101906134ff9190614a47565b6001600160a01b0316826001600160a01b0316106135855760405162461bcd60e51b815260206004820152603860248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206f60448201527f72646572206f6620627269646765206f70657261746f72730000000000000000606482015260840161055a565b6135926040840184615336565b828181106135a2576135a26152ea565b90506020020160208101906135b79190614a47565b9150806135c381615300565b9150506134b7565b600080806135dc6040850185615336565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250508251602090810293810193909320604080517fd679a49e9e099fa9ed83a5446aaec83e746b03ec6723d6f5efb29d37d7f0b78a818701528935818301529885013560608a01526080808a01929092528051808a03909201825260a0909801909752505084519401939093209392505050565b600080600061368e878787876141ec565b9150915061369b816142d9565b5095945050505050565b60008060006136bc6003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316635624191160e01b1790529251931692634bb5274a926137149291016156ff565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161374d91906154f9565b600060405180830381855afa9150503d8060008114613788576040519150601f19603f3d011682016040523d82523d6000602084013e61378d565b606091505b5091509150816119a35760405162461bcd60e51b815260206004820152604260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f427269646765566f74657257656967687428616464726573732960206661696c606482015261195960f21b608482015260a40161055a565b60008060006138276003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b031663926323d560e01b17905291516001600160a01b039390931692634bb5274a926138709291016156ff565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516138a991906154f9565b600060405180830381855afa9150503d80600081146138e4576040519150601f19603f3d011682016040523d82523d6000602084013e6138e9565b606091505b509150915081611e375760405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060746f74604482015272185b15d95a59da1d1cca0a580819985a5b1959606a1b606482015260840161055a565b6020808801518851600082815260018452604080822083835290945292832061397f81613d72565b156139905760019350505050613d67565b6020808c015160009081529081905260409020548214613a0b5760405162461bcd60e51b815260206004820152603060248201527f436f7265476f7665726e616e63653a20717565727920666f7220696e76616c6960448201526f642070726f706f73616c206e6f6e636560801b606482015260840161055a565b6000815460ff166004811115613a2357613a23614aca565b14613a7e5760405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a2074686520766f74652069732066696e616044820152641b1a5e995960da1b606482015260840161055a565b6001600160a01b038716600090815260088201602052604090205460ff1615613ac557613ab5876001600160a01b03166014613f94565b6040516020016125279190615c6b565b6001600160a01b03871660009081526008820160209081526040909120805460ff19166001179055860151151580613b005750604086015115155b80613b0e5750855160ff1615155b15613b55576001600160a01b03871660009081526007820160209081526040918290208851815460ff191660ff909116178155908801516001820155908701516002909101555b866001600160a01b031681600101547f1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a238c88604051613b95929190615cbb565b60405180910390a3600080808c6001811115613bb357613bb3614aca565b03613c08576004830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600384018054899290613bfb9084906152d7565b9250508190559150613cc7565b60018c6001811115613c1c57613c1c614aca565b03613c71576005830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600284018054899290613c649084906152d7565b9250508190559050613cc7565b60405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a20756e737570706f7274656420766f7465604482015264207479706560d81b606482015260840161055a565b8a8210613d1b57825460ff19166001908117845580840154604051919750907f5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b90600090a2613d16838e61448f565b613d61565b898110613d6157825460ff19166003178355600180840154604051919750907f55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f90600090a25b50505050505b979650505050505050565b600080825460ff166004811115613d8b57613d8b614aca565b148015613d9c575042826006015411155b90508015613f8f5760018201546040517f58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a5590600090a260005b6004830154811015613e8e57826008016000846004018381548110613dfc57613dfc6152ea565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556004840180546007860192919084908110613e4557613e456152ea565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613e8681615300565b915050613dd5565b5060005b6005830154811015613f4b57826008016000846005018381548110613eb957613eb96152ea565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556005840180546007860192919084908110613f0257613f026152ea565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613f4381615300565b915050613e92565b50815460ff191682556000600183018190556002830181905560038301819055613f799060048401906147f6565b613f876005830160006147f6565b600060068301555b919050565b60606000613fa383600261537f565b613fae9060026152d7565b6001600160401b03811115613fc557613fc56150a6565b6040519080825280601f01601f191660200182016040528015613fef576020820181803683370190505b509050600360fc1b8160008151811061400a5761400a6152ea565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110614039576140396152ea565b60200101906001600160f81b031916908160001a905350600061405d84600261537f565b6140689060016152d7565b90505b60018111156140e0576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061409c5761409c6152ea565b1a60f81b8282815181106140b2576140b26152ea565b60200101906001600160f81b031916908160001a90535060049490941c936140d981615cd2565b905061406b565b50831561170e5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161055a565b6141376147b9565b6000614144868686612eb9565b915061415b6002548361308790919063ffffffff16565b600061416683612964565b90506141786000828960200151613183565b8351909250821461419b5760405162461bcd60e51b815260040161055a90615b82565b80827f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca856141c88b6132b1565b8b896040516141da9493929190615830565b60405180910390a35094509492505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561422357506000905060036142d0565b8460ff16601b1415801561423b57508460ff16601c14155b1561424c57506000905060046142d0565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156142a0573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166142c9576000600192509250506142d0565b9150600090505b94509492505050565b60008160048111156142ed576142ed614aca565b036142f55750565b600181600481111561430957614309614aca565b036143565760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161055a565b600281600481111561436a5761436a614aca565b036143b75760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161055a565b60038160048111156143cb576143cb614aca565b036144235760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161055a565b600481600481111561443757614437614aca565b036106965760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161055a565b614498816144fc565b1561119857815460ff191660021782556000806144b483614516565b9150915083600101547fe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba83836040516144ee929190615ce9565b60405180910390a250505050565b600081602001516000148061171157505060200151461490565b606080614522836144fc565b61457a5760405162461bcd60e51b815260206004820152602360248201527f50726f706f73616c3a20717565727920666f7220696e76616c696420636861696044820152621b925960ea1b606482015260840161055a565b8260600151516001600160401b03811115614597576145976150a6565b6040519080825280602002602001820160405280156145c0578160200160208202803683370190505b5091508260600151516001600160401b038111156145e0576145e06150a6565b60405190808252806020026020018201604052801561461357816020015b60608152602001906001900390816145fe5790505b50905060005b8360600151518110156147b3578360c00151818151811061463c5761463c6152ea565b60200260200101515a116146925760405162461bcd60e51b815260206004820152601a60248201527f50726f706f73616c3a20696e73756666696369656e7420676173000000000000604482015260640161055a565b836060015181815181106146a8576146a86152ea565b60200260200101516001600160a01b0316846080015182815181106146cf576146cf6152ea565b60200260200101518560c0015183815181106146ed576146ed6152ea565b6020026020010151908660a00151848151811061470c5761470c6152ea565b602002602001015160405161472191906154f9565b600060405180830381858888f193505050503d806000811461475f576040519150601f19603f3d011682016040523d82523d6000602084013e614764565b606091505b50848381518110614777576147776152ea565b60200260200101848481518110614790576147906152ea565b602090810291909101019190915290151590526147ac81615300565b9050614619565b50915091565b6040518060e00160405280600081526020016000815260200160008152602001606081526020016060815260200160608152602001606081525090565b508054600082559060005260206000209081019061069691905b808211156148245760008155600101614810565b5090565b600060e0828403121561483a57600080fd5b50919050565b60008083601f84011261485257600080fd5b5081356001600160401b0381111561486957600080fd5b6020830191508360208260051b850101111561488457600080fd5b9250929050565b60008083601f84011261489d57600080fd5b5081356001600160401b038111156148b457600080fd5b60208301915083602060608302850101111561488457600080fd5b6000806000806000606086880312156148e757600080fd5b85356001600160401b03808211156148fe57600080fd5b61490a89838a01614828565b9650602088013591508082111561492057600080fd5b61492c89838a01614840565b9096509450604088013591508082111561494557600080fd5b506149528882890161488b565b969995985093965092949392505050565b600080600080600080600080600060a08a8c03121561498157600080fd5b8935985060208a01356001600160401b038082111561499f57600080fd5b6149ab8d838e01614840565b909a50985060408c01359150808211156149c457600080fd5b6149d08d838e01614840565b909850965060608c01359150808211156149e957600080fd5b6149f58d838e01614840565b909650945060808c0135915080821115614a0e57600080fd5b50614a1b8c828d01614840565b915080935050809150509295985092959850929598565b6001600160a01b038116811461069657600080fd5b600060208284031215614a5957600080fd5b813561170e81614a32565b60008060408385031215614a7757600080fd5b50508035926020909101359150565b600081518084526020808501945080840160005b83811015614abf5781516001600160a01b031687529582019590820190600101614a9a565b509495945050505050565b634e487b7160e01b600052602160045260246000fd5b6002811061069657610696614aca565b600081518084526020808501945080840160005b83811015614abf578151805160ff16885283810151848901526040908101519088015260609096019590820190600101614b04565b606081526000614b4c6060830186614a86565b82810360208481019190915285518083528682019282019060005b81811015614b8c578451614b7a81614ae0565b83529383019391830191600101614b67565b50508481036040860152614ba08187614af0565b98975050505050505050565b604081526000614bbf6040830185614a86565b82810360208401526125d88185614af0565b600080600060608486031215614be657600080fd5b83359250602084013591506040840135614bff81614a32565b809150509250925092565b600060208284031215614c1c57600080fd5b5035919050565b600080600080600060608688031215614c3b57600080fd5b85356001600160401b0380821115614c5257600080fd5b9087019060c0828a031215614c6657600080fd5b9095506020870135908082111561492057600080fd5b600080600060408486031215614c9157600080fd5b83356001600160401b0380821115614ca857600080fd5b9085019060608288031215614cbc57600080fd5b90935060208501359080821115614cd257600080fd5b50614cdf8682870161488b565b9497909650939450505050565b60208152815160208201526020820151604082015260006040830151606080840152610d176080840182614a86565b6002811061069657600080fd5b8035613f8f81614d1b565b60008060008060008060008060008060c08b8d031215614d5257600080fd5b8a35995060208b01356001600160401b0380821115614d7057600080fd5b614d7c8e838f01614840565b909b50995060408d0135915080821115614d9557600080fd5b614da18e838f01614840565b909950975060608d0135915080821115614dba57600080fd5b614dc68e838f01614840565b909750955060808d0135915080821115614ddf57600080fd5b50614dec8d828e01614840565b9094509250614dff905060a08c01614d28565b90509295989b9194979a5092959850565b60008060408385031215614e2357600080fd5b8235614e2e81614a32565b91506020830135614e3e81614a32565b809150509250929050565b600080600080600060a08688031215614e6157600080fd5b853594506020860135614e7381614a32565b93506040860135614e8381614a32565b94979396509394606081013594506080013592915050565b60008060008060008060008060008060c08b8d031215614eba57600080fd5b8a35995060208b0135985060408b01356001600160401b0380821115614edf57600080fd5b614eeb8e838f01614840565b909a50985060608d0135915080821115614f0457600080fd5b614f108e838f01614840565b909850965060808d0135915080821115614f2957600080fd5b614f358e838f01614840565b909650945060a08d0135915080821115614f4e57600080fd5b50614f5b8d828e01614840565b915080935050809150509295989b9194979a5092959850565b60008060008060808587031215614f8a57600080fd5b8435614f9581614a32565b93506020850135614fa581614a32565b93969395505050506040820135916060013590565b60008060408385031215614fcd57600080fd5b82356001600160401b03811115614fe357600080fd5b614fef85828601614828565b9250506020830135614e3e81614d1b565b60a081016005871061501457615014614aca565b95815260208101949094526040840192909252606083015260809091015290565b6000806040838503121561504857600080fd5b823591506020830135614e3e81614a32565b6020808252602c908201527f526f6e696e476f7665726e616e636541646d696e3a2073656e6465722069732060408201526b3737ba1033b7bb32b93737b960a11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156150de576150de6150a6565b60405290565b60405160c081016001600160401b03811182821017156150de576150de6150a6565b604051601f8201601f191681016001600160401b038111828210171561512e5761512e6150a6565b604052919050565b60006001600160401b0382111561514f5761514f6150a6565b5060051b60200190565b600061516c61516784615136565b615106565b8381529050602080820190600585901b84018681111561518b57600080fd5b845b8181101561521e5780356001600160401b03808211156151ad5760008081fd5b8188019150601f8a818401126151c35760008081fd5b8235828111156151d5576151d56150a6565b6151e6818301601f19168801615106565b92508083528b878286010111156151ff57600091508182fd5b808785018885013760009083018701525085525092820192820161518d565b505050509392505050565b600061170e368484615159565b60208082526027908201527f476f7665726e616e636541646d696e3a206f6e6c7920616c6c6f7765642073656040820152661b198b58d85b1b60ca1b606082015260800190565b60208082526024908201527f476f7665726e616e636541646d696e3a2073657420746f206e6f6e2d636f6e746040820152631c9858dd60e21b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b80820180821115611711576117116152c1565b634e487b7160e01b600052603260045260246000fd5b600060018201615312576153126152c1565b5060010190565b60006020828403121561532b57600080fd5b815161170e81614a32565b6000808335601e1984360301811261534d57600080fd5b8301803591506001600160401b0382111561536757600080fd5b6020019150600581901b360382131561488457600080fd5b8082028115828204841417611711576117116152c1565b6000813561171181614a32565b81358155600160208084013582840155600283016040850135601e198636030181126153ce57600080fd5b850180356001600160401b038111156153e657600080fd5b83820191508060051b36038213156153fd57600080fd5b68010000000000000000811115615416576154166150a6565b82548184558082101561544a5760008481528581208381019083015b808210156154465782825590880190615432565b5050505b50600092835260208320925b81811015611a175761546783615396565b84820155918401918501615456565b84815260208082018590526060604083018190528201839052600090849060808401835b868110156154c85783356154ad81614a32565b6001600160a01b03168252928201929082019060010161549a565b5098975050505050505050565b60005b838110156154f05781810151838201526020016154d8565b50506000910152565b6000825161550b8184602087016154d5565b9190910192915050565b600082601f83011261552657600080fd5b8135602061553661516783615136565b82815260059290921b8401810191818101908684111561555557600080fd5b8286015b8481101561557957803561556c81614a32565b8352918301918301615559565b509695505050505050565b600082601f83011261559557600080fd5b813560206155a561516783615136565b82815260059290921b840181019181810190868411156155c457600080fd5b8286015b8481101561557957803583529183019183016155c8565b600082601f8301126155f057600080fd5b61170e83833560208501615159565b600060e0823603121561561157600080fd5b6156196150bc565b82358152602083013560208201526040830135604082015260608301356001600160401b038082111561564b57600080fd5b61565736838701615515565b6060840152608085013591508082111561567057600080fd5b61567c36838701615584565b608084015260a085013591508082111561569557600080fd5b6156a1368387016155df565b60a084015260c08501359150808211156156ba57600080fd5b506156c736828601615584565b60c08301525092915050565b600081518084526156eb8160208601602086016154d5565b601f01601f19169290920160200192915050565b60208152600061170e60208301846156d3565b60006020828403121561572457600080fd5b5051919050565b600081518084526020808501945080840160005b83811015614abf5781518752958201959082019060010161573f565b600081518084526020808501808196508360051b8101915082860160005b858110156157a35782840389526157918483516156d3565b98850198935090840190600101615779565b5091979650505050505050565b8051825260208101516020830152604081015160408301526000606082015160e060608501526157e360e0850182614a86565b9050608083015184820360808601526157fc828261572b565b91505060a083015184820360a0860152615816828261575b565b91505060c083015184820360c08601526125d8828261572b565b60808152600061584360808301876157b0565b60208681850152838203604085015260c08201865183528187015182840152604087015160c0604085015281815180845260e0860191508483019350600092505b808310156158ad57835161589781614ae0565b8252928401926001929092019190840190615884565b506060890151935084810360608601526158c7818561572b565b9350505050608086015182820360808401526158e3828261575b565b91505060a086015182820360a08401526158fd828261572b565b93505050506125d860608301846001600160a01b03169052565b6020808252602f908201527f476f7665726e616e636541646d696e3a206361737420766f746520666f72206960408201526e1b9d985b1a59081c1c9bdc1bdcd85b608a1b606082015260800190565b600082601f83011261597757600080fd5b8135602061598761516783615136565b82815260059290921b840181019181810190868411156159a657600080fd5b8286015b848110156155795780356159bd81614d1b565b83529183019183016159aa565b600060c082360312156159dc57600080fd5b6159e46150e4565b823581526020830135602082015260408301356001600160401b0380821115615a0c57600080fd5b615a1836838701615966565b60408401526060850135915080821115615a3157600080fd5b615a3d36838701615584565b60608401526080850135915080821115615a5657600080fd5b615a62368387016155df565b608084015260a0850135915080821115615a7b57600080fd5b50615a8836828601615584565b60a08301525092915050565b60ff8116811461069657600080fd5b600060208284031215615ab557600080fd5b813561170e81615a94565b8135615acb81615a94565b60ff811660ff198354161782555060208201356001820155604082013560028201555050565b604081526000615b0460408301856157b0565b905060018060a01b03831660208301529392505050565b81810381811115611711576117116152c1565b73024b9b7b630ba32b223b7bb32b93730b731b29d160651b815260008251615b5d8160148501602087016154d5565b6d08185b1c9958591e481d9bdd195960921b6014939091019283015250602201919050565b60208082526026908201527f436f7265476f7665726e616e63653a20696e76616c69642070726f706f73616c604082015265206e6f6e636560d01b606082015260800190565b8381526020810183905260608101615bdf83614ae0565b826040830152949350505050565b600060208284031215615bff57600080fd5b813561170e81614d1b565b600060608284031215615c1c57600080fd5b604051606081018181106001600160401b0382111715615c3e57615c3e6150a6565b6040528235615c4c81615a94565b8152602083810135908201526040928301359281019290925250919050565b6f021b7b932a3b7bb32b93730b731b29d160851b815260008251615c968160108501602087016154d5565b6d08185b1c9958591e481d9bdd195960921b6010939091019283015250601e01919050565b60408101615cc884614ae0565b9281526020015290565b600081615ce157615ce16152c1565b506000190190565b604080825283519082018190526000906020906060840190828701845b82811015615d24578151151584529284019290840190600101615d06565b50505083810382850152845180825282820190600581901b8301840187850160005b83811015615d7457601f19868403018552615d628383516156d3565b94870194925090860190600101615d46565b5090999850505050505050505056fea2646970667358221220ec51525c88016fdf6b976948df679c2a91d6c3781ed03f2be0d5ea8c803c864664736f6c63430008110033",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// GovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceMetaData.Bin instead.
var GovernanceBin = GovernanceMetaData.Bin

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _roninChainId *big.Int, _roninTrustedOrganizationContract common.Address, _bridgeContract common.Address, _validatorContract common.Address, _proposalExpiryDuration *big.Int) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := GovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceBin), backend, _roninChainId, _roninTrustedOrganizationContract, _bridgeContract, _validatorContract, _proposalExpiryDuration)
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

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x1e23e048.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch) view returns(address[] _voters, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCaller) GetBridgeOperatorVotingSignatures(opts *bind.CallOpts, _period *big.Int, _epoch *big.Int) (struct {
	Voters     []common.Address
	Signatures []SignatureConsumerSignature
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getBridgeOperatorVotingSignatures", _period, _epoch)

	outstruct := new(struct {
		Voters     []common.Address
		Signatures []SignatureConsumerSignature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voters = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Signatures = *abi.ConvertType(out[1], new([]SignatureConsumerSignature)).(*[]SignatureConsumerSignature)

	return *outstruct, err

}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x1e23e048.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch) view returns(address[] _voters, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _epoch *big.Int) (struct {
	Voters     []common.Address
	Signatures []SignatureConsumerSignature
}, error) {
	return _Governance.Contract.GetBridgeOperatorVotingSignatures(&_Governance.CallOpts, _period, _epoch)
}

// GetBridgeOperatorVotingSignatures is a free data retrieval call binding the contract method 0x1e23e048.
//
// Solidity: function getBridgeOperatorVotingSignatures(uint256 _period, uint256 _epoch) view returns(address[] _voters, (uint8,bytes32,bytes32)[] _signatures)
func (_Governance *GovernanceCallerSession) GetBridgeOperatorVotingSignatures(_period *big.Int, _epoch *big.Int) (struct {
	Voters     []common.Address
	Signatures []SignatureConsumerSignature
}, error) {
	return _Governance.Contract.GetBridgeOperatorVotingSignatures(&_Governance.CallOpts, _period, _epoch)
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

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Governance *GovernanceCaller) RoninChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "roninChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Governance *GovernanceSession) RoninChainId() (*big.Int, error) {
	return _Governance.Contract.RoninChainId(&_Governance.CallOpts)
}

// RoninChainId is a free data retrieval call binding the contract method 0x17ce2dd4.
//
// Solidity: function roninChainId() view returns(uint256)
func (_Governance *GovernanceCallerSession) RoninChainId() (*big.Int, error) {
	return _Governance.Contract.RoninChainId(&_Governance.CallOpts)
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
