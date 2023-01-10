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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roninChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_roninTrustedOrganizationContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalExpiryDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeBridgeContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeRoninTrustedOrgContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCallerMustBeValidatorContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"BridgeContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"EmergencyExitPollCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"}],\"name\":\"EmergencyExitPollExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalProposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GlobalProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"successCalls\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"returnDatas\",\"type\":\"bytes[]\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorContractUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"bridgeOperatorsVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"castProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"castProposalVoteForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"createEmergencyExitPoll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"deleteExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"emergencyPollVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"}],\"name\":\"getBridgeOperatorVotingSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getProposalSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSyncedBridgeOperatorSetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeVoter\",\"type\":\"address\"}],\"name\":\"lastVotedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"proposalVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"_targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"proposeGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"_globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeGlobalProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"_support\",\"type\":\"uint8\"}],\"name\":\"proposeProposalForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setBridgeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_expiryDuration\",\"type\":\"uint256\"}],\"name\":\"setProposalExpiryDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setValidatorContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumVoteStatusConsumer.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"againstVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"internalType\":\"structBridgeOperatorsBallot.BridgeOperatorSet\",\"name\":\"_ballot\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"voteBridgeOperatorsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipientAfterUnlockedFund\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiredAt\",\"type\":\"uint256\"}],\"name\":\"voteEmergencyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200609538038062006095833981016040819052620000349162000280565b84848483806200004381600255565b506005849055604080516020808201839052601660608301527f524f4e494e5f474f5645524e414e43455f41444d494e000000000000000000006080808401919091528284018890528351808403909101815260a0830184528051908201207f599a80fcaa47b95e2323ab4d34d34e0cc9feda4b843edafcc30c7bdf60ea15bf60c08401527f7e7935007966eb860f4a2ee3dcc9fd53fb3205ce2aa86b0126d4893d4d4c14b960e08401527fad7c5bef027816a800da1736444fb58a807ef4c9603b7848673f7e3a68eb14a561010084015261012080840191909152835180840390910181526101409092019092528051910120600655620001458362000170565b6200015082620001c5565b5050505062000165826200021460201b60201c565b5050505050620002de565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7906020015b60405180910390a150565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae04890602001620001ba565b600d80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b683616990602001620001ba565b80516001600160a01b03811681146200027b57600080fd5b919050565b600080600080600060a086880312156200029957600080fd5b85519450620002ab6020870162000263565b9350620002bb6040870162000263565b9250620002cb6060870162000263565b9150608086015190509295509295909350565b615da780620002ee6000396000f3fe608060405234801561001057600080fd5b50600436106102055760003560e01c8063663ac0111161011a578063a8a0e32c116100ad578063cd5965831161007c578063cd596583146104d4578063cdf64a76146104e5578063dcc3eb19146104f8578063f3b7dead1461050b578063fb4f63711461051e57600080fd5b8063a8a0e32c1461044b578063b384abef1461045e578063b5e337de146104b9578063bc96180b146104cc57600080fd5b80639a7d3382116100e95780639a7d3382146103ff5780639e0dc0b314610412578063a1819f9a14610425578063a2fae5701461043857600080fd5b8063663ac0111461039f5780637eff275e146103b2578063988ef53c146103c557806399439089146103ee57600080fd5b80632c5e65201161019d5780633644e5151161016c5780633644e5151461033d5780633e2c314f146103465780635511cde11461036657806360911e8e1461037757806362e52e5f1461038a57600080fd5b80632c5e6520146102e45780632e96a6fb146102f75780632faf925d1461030a57806334d5f37b1461031d57600080fd5b806317ce2dd4116101d957806317ce2dd4146102585780631c905e3914610274578063204e1c7a146102965780632b5df351146102c157600080fd5b80624054b81461020a57806309fcd8c71461021f5780630b26cf66146102325780630b88183014610245575b600080fd5b61021d610218366004614850565b610531565b005b61021d61022d3660046148e4565b61057a565b61021d6102403660046149c8565b610643565b61021d610253366004614850565b610698565b61026160055481565b6040519081526020015b60405180910390f35b6102876102823660046149e5565b6106a8565b60405161026b93929190614a81565b6102a96102a43660046149c8565b610aa5565b6040516001600160a01b03909116815260200161026b565b6102d46102cf366004614b21565b610b97565b604051901515815260200161026b565b6102d46102f2366004614b21565b610bcc565b61021d610305366004614b5a565b610c02565b61021d610318366004614b73565b610c2a565b61026161032b366004614b5a565b60006020819052908152604090205481565b61026160065481565b610359610354366004614bcc565b610c58565b60405161026b9190614c1e565b6003546001600160a01b03166102a9565b61021d610385366004614c31565b610d95565b610392610e55565b60405161026b9190614cda565b61021d6103ad366004614d21565b610ef7565b61021d6103c0366004614dfe565b611003565b6102616103d33660046149c8565b6001600160a01b03166000908152600b602052604090205490565b600d546001600160a01b03166102a9565b61021d61040d3660046149e5565b611134565b61021d610420366004614e37565b611142565b61021d610433366004614e89565b6113a8565b61021d610446366004614f62565b61148e565b61021d610459366004614fa8565b61155a565b6104a861046c3660046149e5565b600160208181526000938452604080852090915291835291208054918101546002820154600383015460069093015460ff909416939192909185565b60405161026b959493929190614fee565b61021d6104c73660046149c8565b611595565b6102616115e7565b6004546001600160a01b03166102a9565b61021d6104f33660046149c8565b6115f7565b6102d4610506366004615023565b61168b565b6102a96105193660046149c8565b6116bd565b61021d61052c366004614b73565b61178a565b600061053c336117e9565b116105625760405162461bcd60e51b815260040161055990615048565b60405180910390fd5b61057385858585856006543361195d565b5050505050565b6000610585336117e9565b116105a25760405162461bcd60e51b815260040161055990615048565b6106378989898989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506105e892508a91508b9050615217565b8787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506106259250610d86915050565b6004546001600160a01b0316336119c7565b50505050505050505050565b3330146106625760405162461bcd60e51b815260040161055990615224565b6000816001600160a01b03163b1161068c5760405162461bcd60e51b81526004016105599061526b565b61069581611afc565b50565b6105738585858585600654611b51565b600082815260016020908152604080832084845290915281206004810154600582015460609384938493909290916106e082846152c5565b9050806001600160401b038111156106fa576106fa615094565b604051908082528060200260200182016040528015610723578160200160208202803683370190505b509550806001600160401b0381111561073e5761073e615094565b60405190808252806020026020018201604052801561078957816020015b604080516060810182526000808252602080830182905292820152825260001990920191018161075c5790505b509450806001600160401b038111156107a4576107a4615094565b6040519080825280602002602001820160405280156107cd578160200160208202803683370190505b50965060005b8381101561092c5760008782815181106107ef576107ef6152d8565b6020026020010190600181111561080857610808614a07565b9081600181111561081b5761081b614a07565b90525060008a81526001602090815260408083208c845290915281206004870180546007909201929184908110610854576108546152d8565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286518790839081106108b2576108b26152d8565b60200260200101819052508460040181815481106108d2576108d26152d8565b9060005260206000200160009054906101000a90046001600160a01b0316888281518110610902576109026152d8565b6001600160a01b039092166020928302919091019091015280610924816152ee565b9150506107d3565b5060005b82811015610a995760018761094586846152c5565b81518110610955576109556152d8565b6020026020010190600181111561096e5761096e614a07565b9081600181111561098157610981614a07565b90525060008a81526001602090815260408083208c8452909152812060058701805460079092019291849081106109ba576109ba6152d8565b60009182526020808320909101546001600160a01b0316835282810193909352604091820190208151606081018352815460ff168152600182015493810193909352600201549082015286610a0f86846152c5565b81518110610a1f57610a1f6152d8565b6020026020010181905250846005018181548110610a3f57610a3f6152d8565b6000918252602090912001546001600160a01b031688610a5f86846152c5565b81518110610a6f57610a6f6152d8565b6001600160a01b039092166020928302919091019091015280610a91816152ee565b915050610930565b50505050509250925092565b6000806000836001600160a01b0316604051610acb90635c60da1b60e01b815260040190565b600060405180830381855afa9150503d8060008114610b06576040519150601f19603f3d011682016040523d82523d6000602084013e610b0b565b606091505b509150915081610b7b5760405162461bcd60e51b815260206004820152603560248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060696d706044820152741b195b595b9d185d1a5bdb8a0a580819985a5b1959605a1b6064820152608401610559565b80806020019051810190610b8f9190615307565b949350505050565b6000838152600a6020908152604080832085845282528083206001600160a01b03851684526002019091528120541515610b8f565b600083815260016020908152604080832085845282528083206001600160a01b038516845260080190915281205460ff16610b8f565b333014610c215760405162461bcd60e51b815260040161055990615224565b61069581600255565b6105738585858585600654610c476003546001600160a01b031690565b6004546001600160a01b0316611bd8565b6060816001600160401b03811115610c7257610c72615094565b604051908082528060200260200182016040528015610cbd57816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610c905790505b50905060005b82811015610d7d576000868152600c60209081526040808320888452909152812090858584818110610cf757610cf76152d8565b9050602002016020810190610d0c91906149c8565b6001600160a01b0316815260208082019290925260409081016000208151606081018352815460ff16815260018201549381019390935260020154908201528251839083908110610d5f57610d5f6152d8565b60200260200101819052508080610d75906152ee565b915050610cc3565b50949350505050565b6003546001600160a01b031690565b610dab838383610da3611c8c565b600654611df8565b82356000908152600a6020908152604080832082870135845290915290206001815460ff166004811115610de157610de1614a07565b03610e4f57836007610df38282615391565b507f7c45875370690698791a915954b9c69729cc5f9373edc5a2e04436c07589f30d905084356020860135610e2b6040880188615324565b604051610e3b9493929190615464565b60405180910390a1805460ff191660021781555b50505050565b610e7960405180606001604052806000815260200160008152602001606081525090565b60408051606081018252600780548252600854602080840191909152600980548551818402810184018752818152949593949386019392830182828015610ee957602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610ecb575b505050505081525050905090565b6000610f02336117e9565b11610f1f5760405162461bcd60e51b815260040161055990615048565b60003390506000610fe8468d8d8d80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250610fad92508d91508e9050615217565b8a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508b9250612127915050565b9050610ff5828285612248565b505050505050505050505050565b3330146110225760405162461bcd60e51b815260040161055990615224565b604080516001600160a01b0383811660248084019190915283518084039091018152604490920183526020820180516001600160e01b03166308f2839760e41b1790529151600092851691611076916154e7565b6000604051808303816000865af19150503d80600081146110b3576040519150601f19603f3d011682016040523d82523d6000602084013e6110b8565b606091505b505090508061112f5760405162461bcd60e51b815260206004820152603960248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206063686160448201527f6e676541646d696e28616464726573732960206661696c6564000000000000006064820152608401610559565b505050565b61113e82826123a1565b5050565b33600061114e826117e9565b9050600081116111705760405162461bcd60e51b815260040161055990615048565b600061117e878787876123c3565b90508088146111df5760405162461bcd60e51b815260206004820152602760248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c696420766f6044820152660e8ca40d0c2e6d60cb1b6064820152608401610559565b6000818152600e6020526040902060058101546112585760405162461bcd60e51b815260206004820152603160248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f72206044820152706e6f6e2d6578697374656e7420766f746560781b6064820152608401610559565b6004815460ff16600481111561127057611270614a07565b036112d25760405162461bcd60e51b815260206004820152602c60248201527f526f6e696e476f7665726e616e636541646d696e3a20717565727920666f722060448201526b6578706972656420766f746560a01b6064820152608401610559565b60006112e88286866112e2611c8c565b87612433565b905060018160048111156112fe576112fe614a07565b036113505761130d8989612562565b6040518381527fd3500576a0d4923326fbb893cf2169273e0df93f3cb6b94b83f2ca2e0ecb681b9060200160405180910390a1815460ff19166002178255610637565b600481600481111561136457611364614a07565b03610637576040518381527feecb3148acc573548e89cb64eb5f2023a61171f1c413ed8bf0fe506c19aeebe49060200160405180910390a150505050505050505050565b60006113b3336117e9565b116113d05760405162461bcd60e51b815260040161055990615048565b6114818a8a8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c91829185019084908082843760009201919091525061144692508a91508b9050615217565b878780806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250339250612127915050565b5050505050505050505050565b336114a1600d546001600160a01b031690565b6001600160a01b0316146114c857604051630e6444a160e31b815260040160405180910390fd5b60006114d6858585856123c3565b6000818152600e60209081526040918290204260058201556004810186905582518481526001600160a01b038a811693820193909352918816828401526060820187905260808201869052915192935090917f18ea835340bb2973a31996158138f109e9c5b9cfdb2424e999e6b1a9ce565de89181900360a00190a1505050505050565b6000611565336117e9565b116115825760405162461bcd60e51b815260040161055990615048565b61113e3361158f846155ed565b83612248565b3330146115b45760405162461bcd60e51b815260040161055990615224565b6000816001600160a01b03163b116115de5760405162461bcd60e51b81526004016105599061526b565b610695816126fe565b60006115f260025490565b905090565b3330146116165760405162461bcd60e51b815260040161055990615224565b6000816001600160a01b03163b116116825760405162461bcd60e51b815260206004820152602960248201527f526f6e696e476f7665726e616e636541646d696e3a2073657420746f206e6f6e6044820152680b58dbdb9d1c9858dd60ba1b6064820152608401610559565b6106958161274c565b6000828152600e602090815260408083206001600160a01b038516845260020190915281205415155b90505b92915050565b6000806000836001600160a01b03166040516116e3906303e1469160e61b815260040190565b600060405180830381855afa9150503d806000811461171e576040519150601f19603f3d011682016040523d82523d6000602084013e611723565b606091505b509150915081610b7b5760405162461bcd60e51b815260206004820152602c60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206061646d60448201526b1a5b8a0a580819985a5b195960a21b6064820152608401610559565b6000611795336117e9565b116117b25760405162461bcd60e51b815260040161055990615048565b6117e185858585856006546117cf6003546001600160a01b031690565b6004546001600160a01b03163361279a565b505050505050565b60008060006118006003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316631af0725f60e31b1790529251931692634bb5274a926118589291016156ed565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161189191906154e7565b600060405180830381855afa9150503d80600081146118cc576040519150601f19603f3d011682016040523d82523d6000602084013e6118d1565b606091505b5091509150816119495760405162461bcd60e51b815260206004820152603f60248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f476f7665726e6f7257656967687428616464726573732960206661696c6564006064820152608401610559565b80806020019051810190610b8f9190615700565b61196f611969886155ed565b826127fd565b50600061198361197e896155ed565b6128e5565b90506119bd611991896155ed565b888888886119a9896119a4896000612a84565b612ada565b6119b88a6119a48a6001612a84565b612b01565b5050505050505050565b6040805160c08101909152600080805260208190527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb554909182918190611a0f9060016152c5565b81526020018c81526020018b8b808060200260200160405190810160405280939291908181526020018383602002808284376000920182905250938552505050602082018b9052604082018a90526060909101889052909150611a73828787612e3a565b9050611a8a6002548261300890919063ffffffff16565b6000611a95826128e5565b9050611aa36000828f613104565b935080847f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca84611ad287613232565b878a604051611ae4949392919061581e565b60405180910390a35050509998505050505050505050565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527f5cbd8a0bb00196365d5eb3457c6734e7f06666c3c78e5469b4c9deec7edae048906020015b60405180910390a150565b6000611b5f61197e886155ed565b6020808901356000908152600180835260408083208c35845290935291902001549091508114611ba15760405162461bcd60e51b815260040161055990615905565b611bcf611bad886155ed565b87878787611bc0886119a4896000612a84565b6119b8896119a48a6001612a84565b50505050505050565b6000611bef8383611be88c6159b8565b9190612e3a565b90506000611c04611bff8b6159b8565b613232565b9050611c0f826128e5565b600080805260016020818152855183527fa6eef7e35abe7026729641147f7915573c7e97b47efa546f5f6e3230263bcb4990526040909120015414611c665760405162461bcd60e51b815260040161055990615905565b610637828a8a8a8a611c7d8b6119a4896000612a84565b6119b88c6119a48a6001612a84565b6000806000611ca36003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b0316637de5dedd60e01b17905291516001600160a01b039390931692634bb5274a92611cec9291016156ed565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051611d2591906154e7565b600060405180830381855afa9150503d8060008114611d60576040519150601f19603f3d011682016040523d82523d6000602084013e611d65565b606091505b509150915081611ddd5760405162461bcd60e51b815260206004820152603860248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c20606d696e60448201527f696d756d566f7465576569676874282960206661696c656400000000000000006064820152608401610559565b80806020019051810190611df19190615700565b9250505090565b600754853510801590611e115750600854602086013510155b611e835760405162461bcd60e51b815260206004820152603d60248201527f424f73476f7665726e616e636550726f706f73616c3a20717565727920666f7260448201527f206f7574646174656420627269646765206f70657261746f72207365740000006064820152608401610559565b611e8c8561338c565b82611eed5760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b6064820152608401610559565b6000806000611efb8861354c565b90506000611f098583612ada565b89356000818152600a60209081526040808320828f0135808552908352818420948452600c83528184209084529091528120929350909190805b8a8110156120c757368c8c83818110611f5e57611f5e6152d8565b606002919091019150611f89905086611f7a6020840184615a91565b836020013584604001356135fe565b9850886001600160a01b0316886001600160a01b0316106120005760405162461bcd60e51b815260206004820152602b60248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526a34b3b732b91037b93232b960a91b6064820152608401610559565b88975050600061200f89613626565b905080156120b4576001600160a01b0389166000908152600b60205260409020439055600192508c8c83818110612048576120486152d8565b9050606002018460008b6001600160a01b03166001600160a01b03168152602001908152602001600020818161207e9190615aae565b5060019050612090868b848f8c612433565b60048111156120a1576120a1614a07565b036120b457505050505050505050610573565b50806120bf816152ee565b915050611f43565b5080610ff55760405162461bcd60e51b815260206004820152602960248201527f424f73476f7665726e616e636550726f706f73616c3a20696e76616c6964207360448201526869676e61747572657360b81b6064820152608401610559565b61212f61473a565b8760000361217f5760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e2069646044820152606401610559565b6040805160e08101825260008a81526020819052919091205481906121a59060016152c5565b81526020018981526020018881526020018781526020018681526020018581526020018481525090506121e36002548261300890919063ffffffff16565b60006121ee826128e5565b905060006121fd8a838b613104565b905081818b7fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd98688604051612233929190615adf565b60405180910390a45050979650505050505050565b468260200151146122aa5760405162461bcd60e51b815260206004820152602660248201527f526f6e696e476f7665726e616e636541646d696e3a20696e76616c6964206368604482015265185a5b881a5960d21b6064820152608401610559565b6122b3826128e5565b6020808401516000908152600180835260408083208751845290935291902001541461233e5760405162461bcd60e51b815260206004820152603460248201527f526f6e696e476f7665726e616e636541646d696e3a206361737420766f746520604482015273199bdc881a5b9d985b1a59081c1c9bdc1bdcd85b60621b6064820152608401610559565b6000612348611c8c565b9050600081612355613791565b61235f9190615b09565b61236a9060016152c5565b6040805160608101825260008082526020820181905291810191909152909150611bcf858585858a8661239c8d6117e9565b6138d8565b60008281526001602090815260408083208484529091529020610e4f81613cf3565b604080517f697acba4deaf1a718d8c2d93e42860488cb7812696f28ca10eed17bac41e70276020808301919091526001600160a01b0396871682840152949095166060860152608085019290925260a0808501919091528151808503909101815260c09093019052815191012090565b600080866004015411801561244c575042866004015411155b156124645750845460ff191660049081178655612559565b6001600160a01b0385166000908152600287016020526040902054156124ce57612498856001600160a01b03166014613f15565b6040516020016124a89190615b1c565b60408051601f198184030181529082905262461bcd60e51b8252610559916004016156ed565b6001600160a01b038516600090815260028701602090815260408083208590558483526003890190915281208054869190839061250c9084906152c5565b925050819055905083811015801561253957506000875460ff16600481111561253757612537614a07565b145b1561255157865460ff19166001908117885587018390555b5050845460ff165b95945050505050565b6000612576600d546001600160a01b031690565b604080516001600160a01b0386811660248084019190915286821660448085019190915284518085039091018152606490930184526020830180516001600160e01b03166361e45aeb60e11b1790529251931692634bb5274a926125db9291016156ed565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161261491906154e7565b6000604051808303816000865af19150503d8060008114612651576040519150601f19603f3d011682016040523d82523d6000602084013e612656565b606091505b505090508061112f5760405162461bcd60e51b815260206004820152606260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206065786560448201527f6352656c656173654c6f636b656446756e64466f72456d657267656e6379457860648201527f69745265717565737428616464726573732c616464726573732960206661696c608482015261195960f21b60a482015260c401610559565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d790602001611b46565b600d80546001600160a01b0319166001600160a01b0383169081179091556040519081527fef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b683616990602001611b46565b6127a261473a565b6127b66127ae8b6159b8565b8585856140b0565b50905060006127c7611bff8c6159b8565b90506127ef828b8b8b8b6127e08c6119a4896000612a84565b6119b88d6119a48a6001612a84565b509998505050505050505050565b60208201516000908082036128545760405162461bcd60e51b815260206004820181905260248201527f436f7265476f7665726e616e63653a20696e76616c696420636861696e2069646044820152606401610559565b600254612862908590613008565b600061286d856128e5565b905061287e82828760400151613104565b855190935083146128a15760405162461bcd60e51b815260040161055990615b70565b8083837fa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd988886040516128d5929190615adf565b60405180910390a4505092915050565b6000806000806000808660800151905060008760600151905060008860a00151516001600160401b0381111561291d5761291d615094565b604051908082528060200260200182016040528015612946578160200160208202803683370190505b5060c08a015190915060005b82518110156129af578a60a001518181518110612971576129716152d8565b602002602001015180519060200120838281518110612992576129926152d8565b6020908102919091010152806129a7816152ee565b915050612952565b506020835102602084012097506020845102602085012096506020825102602083012095506020815102602082012094507fd051578048e6ff0bbc9fca3b65a42088dbde10f36ca841de566711087ad9b08a60001b8a600001518b602001518c604001518b8b8b8b604051602001612a5f989796959493929190978852602088019690965260408701949094526060860192909252608085015260a084015260c083015260e08201526101000190565b6040516020818303038152906040528051906020012098505050505050505050919050565b604051600090612abc907fd900570327c4c0df8dd6bdd522b7da7e39145dd049d2fd4602276adcd511e3c29085908590602001615bb6565b60405160208183030381529060405280519060200120905092915050565b60405161190160f01b60208201526022810183905260428101829052600090606201612abc565b8415801590612b0f57508483145b612b6c5760405162461bcd60e51b815260206004820152602860248201527f476f7665726e616e636550726f706f73616c3a20696e76616c696420617272616044820152670f240d8cadccee8d60c31b6064820152608401610559565b6000612b76611c8c565b9050600081612b83613791565b612b8d9190615b09565b612b989060016152c5565b9050600080366000805b89811015612dce578a8a82818110612bbc57612bbc6152d8565b606002919091019350600090508d8d83818110612bdb57612bdb6152d8565b9050602002016020810190612bf09190615bdb565b6001811115612c0157612c01614a07565b03612c2d57612c2689612c176020860186615a91565b856020013586604001356135fe565b9350612ce1565b60018d8d83818110612c4157612c416152d8565b9050602002016020810190612c569190615bdb565b6001811115612c6757612c67614a07565b03612c7d57612c2688612c176020860186615a91565b60405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636550726f706f73616c3a20717565727920666f7220756e604482015272737570706f7274656420766f7465207479706560681b6064820152608401610559565b836001600160a01b0316856001600160a01b031610612d4c5760405162461bcd60e51b815260206004820152602160248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964206f7264656044820152603960f91b6064820152608401610559565b8394506000612d5a856117e9565b90508015612dbb5760019250612da98f8f8f85818110612d7c57612d7c6152d8565b9050602002016020810190612d919190615bdb565b8a8a89612da3368b90038b018b615bf8565b876138d8565b15612dbb575050505050505050611bcf565b5080612dc6816152ee565b915050612ba2565b5080612e2b5760405162461bcd60e51b815260206004820152602660248201527f476f7665726e616e636550726f706f73616c3a20696e76616c6964207369676e60448201526561747572657360d01b6064820152608401610559565b50505050505050505050505050565b612e4261473a565b83518152602080850151604080840191909152600091830191909152840151516001600160401b03811115612e7957612e79615094565b604051908082528060200260200182016040528015612ea2578160200160208202803683370190505b5060608083019190915284015160808083019190915284015160a08083019190915284015160c082015260005b84604001515181101561300057600185604001518281518110612ef457612ef46152d8565b60200260200101516001811115612f0d57612f0d614a07565b03612f4e578282606001518281518110612f2957612f296152d8565b60200260200101906001600160a01b031690816001600160a01b031681525050612fee565b600085604001518281518110612f6657612f666152d8565b60200260200101516001811115612f7f57612f7f614a07565b03612f9b578382606001518281518110612f2957612f296152d8565b60405162461bcd60e51b815260206004820152602260248201527f476c6f62616c50726f706f73616c3a20756e737570706f727465642074617267604482015261195d60f21b6064820152608401610559565b80612ff8816152ee565b915050612ecf565b509392505050565b60008260600151511180156130265750816080015151826060015151145b801561303b57508160a0015151826060015151145b801561305057508160c0015151826060015151145b61309c5760405162461bcd60e51b815260206004820152601e60248201527f50726f706f73616c3a20696e76616c6964206172726179206c656e67746800006044820152606401610559565b6130a681426152c5565b8260400151111561113e5760405162461bcd60e51b815260206004820152602260248201527f50726f706f73616c3a20696e76616c6964206578706972792074696d6573746160448201526106d760f41b6064820152608401610559565b6000838152602081905260408120549081900361313557506000838152602081905260409020600190819055613205565b600084815260016020908152604080832084845290915281209061315882613cf3565b905080613202576000825460ff16600481111561317757613177614a07565b036131de5760405162461bcd60e51b815260206004820152603160248201527f436f7265476f7665726e616e63653a2063757272656e742070726f706f73616c604482015270081a5cc81b9bdd0818dbdb5c1b195d1959607a1b6064820152608401610559565b600086815260208190526040812080549091906131fa906152ee565b918290555092505b50505b60009384526001602081815260408087208488529091529094209384019290925560069092019190915590565b6000806000806000808660600151905060008760400151905060008860800151516001600160401b0381111561326a5761326a615094565b604051908082528060200260200182016040528015613293578160200160208202803683370190505b5060a08a015190915060005b82518110156132fc578a6080015181815181106132be576132be6152d8565b6020026020010151805190602001208382815181106132df576132df6152d8565b6020908102919091010152806132f4816152ee565b91505061329f565b5082516020908102848201208551820286830120845183028584012084518402858501208e518f860151604080517f1463f426c05aff2c1a7a0957a71c9898bc8b47142540538e79ee25ee911413509881019890985287019190915260608601526080850184905260a0850183905260c0850182905260e08501819052929b509099509750955061010001612a5f565b600061339b6040830183615324565b9050116133fe5760405162461bcd60e51b815260206004820152602b60248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206160448201526a0e4e4c2f240d8cadccee8d60ab1b6064820152608401610559565b600061340d6040830183615324565b600081811061341e5761341e6152d8565b905060200201602081019061343391906149c8565b905060015b6134456040840184615324565b905081101561112f5761345b6040840184615324565b8281811061346b5761346b6152d8565b905060200201602081019061348091906149c8565b6001600160a01b0316826001600160a01b0316106135065760405162461bcd60e51b815260206004820152603860248201527f4272696467654f70657261746f727342616c6c6f743a20696e76616c6964206f60448201527f72646572206f6620627269646765206f70657261746f727300000000000000006064820152608401610559565b6135136040840184615324565b82818110613523576135236152d8565b905060200201602081019061353891906149c8565b915080613544816152ee565b915050613438565b6000808061355d6040850185615324565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250508251602090810293810193909320604080517fd679a49e9e099fa9ed83a5446aaec83e746b03ec6723d6f5efb29d37d7f0b78a818701528935818301529885013560608a01526080808a01929092528051808a03909201825260a0909801909752505084519401939093209392505050565b600080600061360f8787878761416d565b9150915061361c8161425a565b5095945050505050565b600080600061363d6003546001600160a01b031690565b604080516001600160a01b03878116602480840191909152835180840382018152604490930184526020830180516001600160e01b0316635624191160e01b1790529251931692634bb5274a926136959291016156ed565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516136ce91906154e7565b600060405180830381855afa9150503d8060008114613709576040519150601f19603f3d011682016040523d82523d6000602084013e61370e565b606091505b5091509150816119495760405162461bcd60e51b815260206004820152604260248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c206067657460448201527f427269646765566f74657257656967687428616464726573732960206661696c606482015261195960f21b608482015260a401610559565b60008060006137a86003546001600160a01b031690565b6040805160048152602480820183526020820180516001600160e01b031663926323d560e01b17905291516001600160a01b039390931692634bb5274a926137f19291016156ed565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161382a91906154e7565b600060405180830381855afa9150503d8060008114613865576040519150601f19603f3d011682016040523d82523d6000602084013e61386a565b606091505b509150915081611ddd5760405162461bcd60e51b815260206004820152603360248201527f476f7665726e616e636541646d696e3a2070726f78792063616c6c2060746f74604482015272185b15d95a59da1d1cca0a580819985a5b1959606a1b6064820152608401610559565b6020808801518851600082815260018452604080822083835290945292832061390081613cf3565b156139115760019350505050613ce8565b6020808c01516000908152908190526040902054821461398c5760405162461bcd60e51b815260206004820152603060248201527f436f7265476f7665726e616e63653a20717565727920666f7220696e76616c6960448201526f642070726f706f73616c206e6f6e636560801b6064820152608401610559565b6000815460ff1660048111156139a4576139a4614a07565b146139ff5760405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a2074686520766f74652069732066696e616044820152641b1a5e995960da1b6064820152608401610559565b6001600160a01b038716600090815260088201602052604090205460ff1615613a4657613a36876001600160a01b03166014613f15565b6040516020016124a89190615c59565b6001600160a01b03871660009081526008820160209081526040909120805460ff19166001179055860151151580613a815750604086015115155b80613a8f5750855160ff1615155b15613ad6576001600160a01b03871660009081526007820160209081526040918290208851815460ff191660ff909116178155908801516001820155908701516002909101555b866001600160a01b031681600101547f1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a238c88604051613b16929190615ca9565b60405180910390a3600080808c6001811115613b3457613b34614a07565b03613b89576004830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600384018054899290613b7c9084906152c5565b9250508190559150613c48565b60018c6001811115613b9d57613b9d614a07565b03613bf2576005830180546001810182556000918252602082200180546001600160a01b0319166001600160a01b038c16179055600284018054899290613be59084906152c5565b9250508190559050613c48565b60405162461bcd60e51b815260206004820152602560248201527f436f7265476f7665726e616e63653a20756e737570706f7274656420766f7465604482015264207479706560d81b6064820152608401610559565b8a8210613c9c57825460ff19166001908117845580840154604051919750907f5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b90600090a2613c97838e614410565b613ce2565b898110613ce257825460ff19166003178355600180840154604051919750907f55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f90600090a25b50505050505b979650505050505050565b600080825460ff166004811115613d0c57613d0c614a07565b148015613d1d575042826006015411155b90508015613f105760018201546040517f58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a5590600090a260005b6004830154811015613e0f57826008016000846004018381548110613d7d57613d7d6152d8565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556004840180546007860192919084908110613dc657613dc66152d8565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613e07816152ee565b915050613d56565b5060005b6005830154811015613ecc57826008016000846005018381548110613e3a57613e3a6152d8565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff191690556005840180546007860192919084908110613e8357613e836152d8565b60009182526020808320909101546001600160a01b031683528201929092526040018120805460ff19168155600181018290556002015580613ec4816152ee565b915050613e13565b50815460ff191682556000600183018190556002830181905560038301819055613efa906004840190614777565b613f08600583016000614777565b600060068301555b919050565b60606000613f2483600261536d565b613f2f9060026152c5565b6001600160401b03811115613f4657613f46615094565b6040519080825280601f01601f191660200182016040528015613f70576020820181803683370190505b509050600360fc1b81600081518110613f8b57613f8b6152d8565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613fba57613fba6152d8565b60200101906001600160f81b031916908160001a9053506000613fde84600261536d565b613fe99060016152c5565b90505b6001811115614061576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061401d5761401d6152d8565b1a60f81b828281518110614033576140336152d8565b60200101906001600160f81b031916908160001a90535060049490941c9361405a81615cc0565b9050613fec565b5083156116b45760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610559565b6140b861473a565b60006140c5868686612e3a565b91506140dc6002548361300890919063ffffffff16565b60006140e7836128e5565b90506140f96000828960200151613104565b8351909250821461411c5760405162461bcd60e51b815260040161055990615b70565b80827f771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca856141498b613232565b8b8960405161415b949392919061581e565b60405180910390a35094509492505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156141a45750600090506003614251565b8460ff16601b141580156141bc57508460ff16601c14155b156141cd5750600090506004614251565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614221573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661424a57600060019250925050614251565b9150600090505b94509492505050565b600081600481111561426e5761426e614a07565b036142765750565b600181600481111561428a5761428a614a07565b036142d75760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610559565b60028160048111156142eb576142eb614a07565b036143385760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610559565b600381600481111561434c5761434c614a07565b036143a45760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610559565b60048160048111156143b8576143b8614a07565b036106955760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610559565b6144198161447d565b1561113e57815460ff1916600217825560008061443583614497565b9150915083600101547fe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba838360405161446f929190615cd7565b60405180910390a250505050565b60008160200151600014806116b757505060200151461490565b6060806144a38361447d565b6144fb5760405162461bcd60e51b815260206004820152602360248201527f50726f706f73616c3a20717565727920666f7220696e76616c696420636861696044820152621b925960ea1b6064820152608401610559565b8260600151516001600160401b0381111561451857614518615094565b604051908082528060200260200182016040528015614541578160200160208202803683370190505b5091508260600151516001600160401b0381111561456157614561615094565b60405190808252806020026020018201604052801561459457816020015b606081526020019060019003908161457f5790505b50905060005b836060015151811015614734578360c0015181815181106145bd576145bd6152d8565b60200260200101515a116146135760405162461bcd60e51b815260206004820152601a60248201527f50726f706f73616c3a20696e73756666696369656e74206761730000000000006044820152606401610559565b83606001518181518110614629576146296152d8565b60200260200101516001600160a01b031684608001518281518110614650576146506152d8565b60200260200101518560c00151838151811061466e5761466e6152d8565b6020026020010151908660a00151848151811061468d5761468d6152d8565b60200260200101516040516146a291906154e7565b600060405180830381858888f193505050503d80600081146146e0576040519150601f19603f3d011682016040523d82523d6000602084013e6146e5565b606091505b508483815181106146f8576146f86152d8565b60200260200101848481518110614711576147116152d8565b6020908102919091010191909152901515905261472d816152ee565b905061459a565b50915091565b6040518060e00160405280600081526020016000815260200160008152602001606081526020016060815260200160608152602001606081525090565b508054600082559060005260206000209081019061069591905b808211156147a55760008155600101614791565b5090565b600060e082840312156147bb57600080fd5b50919050565b60008083601f8401126147d357600080fd5b5081356001600160401b038111156147ea57600080fd5b6020830191508360208260051b850101111561480557600080fd5b9250929050565b60008083601f84011261481e57600080fd5b5081356001600160401b0381111561483557600080fd5b60208301915083602060608302850101111561480557600080fd5b60008060008060006060868803121561486857600080fd5b85356001600160401b038082111561487f57600080fd5b61488b89838a016147a9565b965060208801359150808211156148a157600080fd5b6148ad89838a016147c1565b909650945060408801359150808211156148c657600080fd5b506148d38882890161480c565b969995985093965092949392505050565b600080600080600080600080600060a08a8c03121561490257600080fd5b8935985060208a01356001600160401b038082111561492057600080fd5b61492c8d838e016147c1565b909a50985060408c013591508082111561494557600080fd5b6149518d838e016147c1565b909850965060608c013591508082111561496a57600080fd5b6149768d838e016147c1565b909650945060808c013591508082111561498f57600080fd5b5061499c8c828d016147c1565b915080935050809150509295985092959850929598565b6001600160a01b038116811461069557600080fd5b6000602082840312156149da57600080fd5b81356116b4816149b3565b600080604083850312156149f857600080fd5b50508035926020909101359150565b634e487b7160e01b600052602160045260246000fd5b6002811061069557610695614a07565b600081518084526020808501945080840160005b83811015614a76578151805160ff16885283810151848901526040908101519088015260609096019590820190600101614a41565b509495945050505050565b606080825284519082018190526000906020906080840190828801845b82811015614ac35781516001600160a01b031684529284019290840190600101614a9e565b5050508381038285015285518082528683019183019060005b81811015614b01578351614aef81614a1d565b83529284019291840191600101614adc565b50508481036040860152614b158187614a2d565b98975050505050505050565b600080600060608486031215614b3657600080fd5b83359250602084013591506040840135614b4f816149b3565b809150509250925092565b600060208284031215614b6c57600080fd5b5035919050565b600080600080600060608688031215614b8b57600080fd5b85356001600160401b0380821115614ba257600080fd5b9087019060c0828a031215614bb657600080fd5b909550602087013590808211156148a157600080fd5b60008060008060608587031215614be257600080fd5b843593506020850135925060408501356001600160401b03811115614c0657600080fd5b614c12878288016147c1565b95989497509550505050565b6020815260006116b46020830184614a2d565b600080600060408486031215614c4657600080fd5b83356001600160401b0380821115614c5d57600080fd5b9085019060608288031215614c7157600080fd5b90935060208501359080821115614c8757600080fd5b50614c948682870161480c565b9497909650939450505050565b600081518084526020808501945080840160005b83811015614a765781516001600160a01b031687529582019590820190600101614cb5565b60208152815160208201526020820151604082015260006040830151606080840152610b8f6080840182614ca1565b6002811061069557600080fd5b8035613f1081614d09565b60008060008060008060008060008060c08b8d031215614d4057600080fd5b8a35995060208b01356001600160401b0380821115614d5e57600080fd5b614d6a8e838f016147c1565b909b50995060408d0135915080821115614d8357600080fd5b614d8f8e838f016147c1565b909950975060608d0135915080821115614da857600080fd5b614db48e838f016147c1565b909750955060808d0135915080821115614dcd57600080fd5b50614dda8d828e016147c1565b9094509250614ded905060a08c01614d16565b90509295989b9194979a5092959850565b60008060408385031215614e1157600080fd5b8235614e1c816149b3565b91506020830135614e2c816149b3565b809150509250929050565b600080600080600060a08688031215614e4f57600080fd5b853594506020860135614e61816149b3565b93506040860135614e71816149b3565b94979396509394606081013594506080013592915050565b60008060008060008060008060008060c08b8d031215614ea857600080fd5b8a35995060208b0135985060408b01356001600160401b0380821115614ecd57600080fd5b614ed98e838f016147c1565b909a50985060608d0135915080821115614ef257600080fd5b614efe8e838f016147c1565b909850965060808d0135915080821115614f1757600080fd5b614f238e838f016147c1565b909650945060a08d0135915080821115614f3c57600080fd5b50614f498d828e016147c1565b915080935050809150509295989b9194979a5092959850565b60008060008060808587031215614f7857600080fd5b8435614f83816149b3565b93506020850135614f93816149b3565b93969395505050506040820135916060013590565b60008060408385031215614fbb57600080fd5b82356001600160401b03811115614fd157600080fd5b614fdd858286016147a9565b9250506020830135614e2c81614d09565b60a081016005871061500257615002614a07565b95815260208101949094526040840192909252606083015260809091015290565b6000806040838503121561503657600080fd5b823591506020830135614e2c816149b3565b6020808252602c908201527f526f6e696e476f7665726e616e636541646d696e3a2073656e6465722069732060408201526b3737ba1033b7bb32b93737b960a11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156150cc576150cc615094565b60405290565b60405160c081016001600160401b03811182821017156150cc576150cc615094565b604051601f8201601f191681016001600160401b038111828210171561511c5761511c615094565b604052919050565b60006001600160401b0382111561513d5761513d615094565b5060051b60200190565b600061515a61515584615124565b6150f4565b8381529050602080820190600585901b84018681111561517957600080fd5b845b8181101561520c5780356001600160401b038082111561519b5760008081fd5b8188019150601f8a818401126151b15760008081fd5b8235828111156151c3576151c3615094565b6151d4818301601f191688016150f4565b92508083528b878286010111156151ed57600091508182fd5b808785018885013760009083018701525085525092820192820161517b565b505050509392505050565b60006116b4368484615147565b60208082526027908201527f476f7665726e616e636541646d696e3a206f6e6c7920616c6c6f7765642073656040820152661b198b58d85b1b60ca1b606082015260800190565b60208082526024908201527f476f7665726e616e636541646d696e3a2073657420746f206e6f6e2d636f6e746040820152631c9858dd60e21b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b808201808211156116b7576116b76152af565b634e487b7160e01b600052603260045260246000fd5b600060018201615300576153006152af565b5060010190565b60006020828403121561531957600080fd5b81516116b4816149b3565b6000808335601e1984360301811261533b57600080fd5b8301803591506001600160401b0382111561535557600080fd5b6020019150600581901b360382131561480557600080fd5b80820281158282048414176116b7576116b76152af565b600081356116b7816149b3565b81358155600160208084013582840155600283016040850135601e198636030181126153bc57600080fd5b850180356001600160401b038111156153d457600080fd5b83820191508060051b36038213156153eb57600080fd5b6801000000000000000081111561540457615404615094565b8254818455808210156154385760008481528581208381019083015b808210156154345782825590880190615420565b5050505b50600092835260208320925b818110156119bd5761545583615384565b84820155918401918501615444565b84815260208082018590526060604083018190528201839052600090849060808401835b868110156154b657833561549b816149b3565b6001600160a01b031682529282019290820190600101615488565b5098975050505050505050565b60005b838110156154de5781810151838201526020016154c6565b50506000910152565b600082516154f98184602087016154c3565b9190910192915050565b600082601f83011261551457600080fd5b8135602061552461515583615124565b82815260059290921b8401810191818101908684111561554357600080fd5b8286015b8481101561556757803561555a816149b3565b8352918301918301615547565b509695505050505050565b600082601f83011261558357600080fd5b8135602061559361515583615124565b82815260059290921b840181019181810190868411156155b257600080fd5b8286015b8481101561556757803583529183019183016155b6565b600082601f8301126155de57600080fd5b6116b483833560208501615147565b600060e082360312156155ff57600080fd5b6156076150aa565b82358152602083013560208201526040830135604082015260608301356001600160401b038082111561563957600080fd5b61564536838701615503565b6060840152608085013591508082111561565e57600080fd5b61566a36838701615572565b608084015260a085013591508082111561568357600080fd5b61568f368387016155cd565b60a084015260c08501359150808211156156a857600080fd5b506156b536828601615572565b60c08301525092915050565b600081518084526156d98160208601602086016154c3565b601f01601f19169290920160200192915050565b6020815260006116b460208301846156c1565b60006020828403121561571257600080fd5b5051919050565b600081518084526020808501945080840160005b83811015614a765781518752958201959082019060010161572d565b600081518084526020808501808196508360051b8101915082860160005b8581101561579157828403895261577f8483516156c1565b98850198935090840190600101615767565b5091979650505050505050565b8051825260208101516020830152604081015160408301526000606082015160e060608501526157d160e0850182614ca1565b9050608083015184820360808601526157ea8282615719565b91505060a083015184820360a08601526158048282615749565b91505060c083015184820360c08601526125598282615719565b608081526000615831608083018761579e565b60208681850152838203604085015260c08201865183528187015182840152604087015160c0604085015281815180845260e0860191508483019350600092505b8083101561589b57835161588581614a1d565b8252928401926001929092019190840190615872565b506060890151935084810360608601526158b58185615719565b9350505050608086015182820360808401526158d18282615749565b91505060a086015182820360a08401526158eb8282615719565b935050505061255960608301846001600160a01b03169052565b6020808252602f908201527f476f7665726e616e636541646d696e3a206361737420766f746520666f72206960408201526e1b9d985b1a59081c1c9bdc1bdcd85b608a1b606082015260800190565b600082601f83011261596557600080fd5b8135602061597561515583615124565b82815260059290921b8401810191818101908684111561599457600080fd5b8286015b848110156155675780356159ab81614d09565b8352918301918301615998565b600060c082360312156159ca57600080fd5b6159d26150d2565b823581526020830135602082015260408301356001600160401b03808211156159fa57600080fd5b615a0636838701615954565b60408401526060850135915080821115615a1f57600080fd5b615a2b36838701615572565b60608401526080850135915080821115615a4457600080fd5b615a50368387016155cd565b608084015260a0850135915080821115615a6957600080fd5b50615a7636828601615572565b60a08301525092915050565b60ff8116811461069557600080fd5b600060208284031215615aa357600080fd5b81356116b481615a82565b8135615ab981615a82565b60ff811660ff198354161782555060208201356001820155604082013560028201555050565b604081526000615af2604083018561579e565b905060018060a01b03831660208301529392505050565b818103818111156116b7576116b76152af565b73024b9b7b630ba32b223b7bb32b93730b731b29d160651b815260008251615b4b8160148501602087016154c3565b6d08185b1c9958591e481d9bdd195960921b6014939091019283015250602201919050565b60208082526026908201527f436f7265476f7665726e616e63653a20696e76616c69642070726f706f73616c604082015265206e6f6e636560d01b606082015260800190565b8381526020810183905260608101615bcd83614a1d565b826040830152949350505050565b600060208284031215615bed57600080fd5b81356116b481614d09565b600060608284031215615c0a57600080fd5b604051606081018181106001600160401b0382111715615c2c57615c2c615094565b6040528235615c3a81615a82565b8152602083810135908201526040928301359281019290925250919050565b6f021b7b932a3b7bb32b93730b731b29d160851b815260008251615c848160108501602087016154c3565b6d08185b1c9958591e481d9bdd195960921b6010939091019283015250601e01919050565b60408101615cb684614a1d565b9281526020015290565b600081615ccf57615ccf6152af565b506000190190565b604080825283519082018190526000906020906060840190828701845b82811015615d12578151151584529284019290840190600101615cf4565b50505083810382850152845180825282820190600581901b8301840187850160005b83811015615d6257601f19868403018552615d508383516156c1565b94870194925090860190600101615d34565b5090999850505050505050505056fea264697066735822122064d69f78167fa956c292fdd8e2758cb35cd84f98e422c782673c424c7cd867df64736f6c63430008110033",
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
