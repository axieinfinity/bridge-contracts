// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeManager

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

// BridgeManagerMetaData contains all meta data concerning the BridgeManager contract.
var BridgeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"roninChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"callbackRegisters\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"internalType\":\"uint96[]\",\"name\":\"voteWeights\",\"type\":\"uint96[]\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codehash\",\"type\":\"bytes32\"}],\"name\":\"ErrAddressIsNotCreatedEOA\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"ErrAlreadyVoted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"ErrBridgeOperatorAlreadyExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"ErrBridgeOperatorUpdateFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrCurrentProposalIsNotCompleted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrDuplicated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ErrInsufficientGas\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidArguments\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"ErrInvalidChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidExpiryTimestamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"}],\"name\":\"ErrInvalidProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidProposalNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrInvalidVoteWeight\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrOnlySelfCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrQueryForEmptyVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrUnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrUnsupportedVoteType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrVoteIsFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromBridgeOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toBridgeOperator\",\"type\":\"address\"}],\"name\":\"BridgeOperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"statuses\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint96[]\",\"name\":\"voteWeights\",\"type\":\"uint96[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"statuses\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"BridgeOperatorsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalProposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"GlobalProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"registers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"statuses\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"returnDatas\",\"type\":\"bytes[]\"}],\"name\":\"Notified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"successCalls\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"returnDatas\",\"type\":\"bytes[]\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ProposalExpiryDurationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumGlobalProposal.TargetOption\",\"name\":\"targetOption\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"TargetOptionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96[]\",\"name\":\"voteWeights\",\"type\":\"uint96[]\"},{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"addBridgeOperators\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"addeds\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"supports_\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"castGlobalProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"supports_\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"castProposalBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castProposalVoteForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"deleteExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"}],\"name\":\"getBridgeOperatorOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeOperator\",\"type\":\"address\"}],\"name\":\"getBridgeOperatorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallbackRegisters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"registers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFullBridgeOperatorInfos\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round_\",\"type\":\"uint256\"}],\"name\":\"getGlobalProposalSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"supports_\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"getGovernorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"}],\"name\":\"getGovernorWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"getGovernorsOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalExpiryDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getProposalSignatures\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_voters\",\"type\":\"address[]\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"num_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"globalProposalVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isBridgeOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"proposalVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"name\":\"proposeGlobal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structGlobalProposal.GlobalProposalDetail\",\"name\":\"globalProposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"supports_\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeGlobalProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumBallot.VoteType\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"proposeProposalForCurrentNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasAmounts\",\"type\":\"uint256[]\"}],\"internalType\":\"structProposal.ProposalDetail\",\"name\":\"_proposal\",\"type\":\"tuple\"},{\"internalType\":\"enumBallot.VoteType[]\",\"name\":\"_supports\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignatureConsumer.Signature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"}],\"name\":\"proposeProposalStructAndCastVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registers\",\"type\":\"address[]\"}],\"name\":\"registerCallbacks\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"registereds\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"bridgeOperators\",\"type\":\"address[]\"}],\"name\":\"removeBridgeOperators\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"removeds\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"}],\"name\":\"resolveTargets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"governors\",\"type\":\"address[]\"}],\"name\":\"sumGovernorsWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBridgeOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"registers\",\"type\":\"address[]\"}],\"name\":\"unregisterCallbacks\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"unregistereds\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBridgeOperator\",\"type\":\"address\"}],\"name\":\"updateBridgeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumGlobalProposal.TargetOption[]\",\"name\":\"targetOptions\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"updateManyTargetOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"enumVoteStatusConsumer.VoteStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"againstVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVoteWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeManagerMetaData.ABI instead.
var BridgeManagerABI = BridgeManagerMetaData.ABI

// BridgeManager is an auto generated Go binding around an Ethereum contract.
type BridgeManager struct {
	BridgeManagerCaller     // Read-only binding to the contract
	BridgeManagerTransactor // Write-only binding to the contract
	BridgeManagerFilterer   // Log filterer for contract events
}

// BridgeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeManagerSession struct {
	Contract     *BridgeManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeManagerCallerSession struct {
	Contract *BridgeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BridgeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeManagerTransactorSession struct {
	Contract     *BridgeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BridgeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeManagerRaw struct {
	Contract *BridgeManager // Generic contract binding to access the raw methods on
}

// BridgeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeManagerCallerRaw struct {
	Contract *BridgeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeManagerTransactorRaw struct {
	Contract *BridgeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeManager creates a new instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManager(address common.Address, backend bind.ContractBackend) (*BridgeManager, error) {
	contract, err := bindBridgeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeManager{BridgeManagerCaller: BridgeManagerCaller{contract: contract}, BridgeManagerTransactor: BridgeManagerTransactor{contract: contract}, BridgeManagerFilterer: BridgeManagerFilterer{contract: contract}}, nil
}

// NewBridgeManagerCaller creates a new read-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerCaller(address common.Address, caller bind.ContractCaller) (*BridgeManagerCaller, error) {
	contract, err := bindBridgeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerCaller{contract: contract}, nil
}

// NewBridgeManagerTransactor creates a new write-only instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeManagerTransactor, error) {
	contract, err := bindBridgeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerTransactor{contract: contract}, nil
}

// NewBridgeManagerFilterer creates a new log filterer instance of BridgeManager, bound to a specific deployed contract.
func NewBridgeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeManagerFilterer, error) {
	contract, err := bindBridgeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerFilterer{contract: contract}, nil
}

// bindBridgeManager binds a generic wrapper to an already deployed contract.
func bindBridgeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.BridgeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.BridgeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeManager *BridgeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeManager *BridgeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeManager *BridgeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeManager.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeManager *BridgeManagerCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeManager *BridgeManagerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BridgeManager.Contract.DOMAINSEPARATOR(&_BridgeManager.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BridgeManager *BridgeManagerCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BridgeManager.Contract.DOMAINSEPARATOR(&_BridgeManager.CallOpts)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_BridgeManager *BridgeManagerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _BridgeManager.Contract.CheckThreshold(&_BridgeManager.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _BridgeManager.Contract.CheckThreshold(&_BridgeManager.CallOpts, _voteWeight)
}

// GetBridgeOperatorOf is a free data retrieval call binding the contract method 0xbc9182fd.
//
// Solidity: function getBridgeOperatorOf(address[] governors) view returns(address[] bridgeOperators)
func (_BridgeManager *BridgeManagerCaller) GetBridgeOperatorOf(opts *bind.CallOpts, governors []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getBridgeOperatorOf", governors)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperatorOf is a free data retrieval call binding the contract method 0xbc9182fd.
//
// Solidity: function getBridgeOperatorOf(address[] governors) view returns(address[] bridgeOperators)
func (_BridgeManager *BridgeManagerSession) GetBridgeOperatorOf(governors []common.Address) ([]common.Address, error) {
	return _BridgeManager.Contract.GetBridgeOperatorOf(&_BridgeManager.CallOpts, governors)
}

// GetBridgeOperatorOf is a free data retrieval call binding the contract method 0xbc9182fd.
//
// Solidity: function getBridgeOperatorOf(address[] governors) view returns(address[] bridgeOperators)
func (_BridgeManager *BridgeManagerCallerSession) GetBridgeOperatorOf(governors []common.Address) ([]common.Address, error) {
	return _BridgeManager.Contract.GetBridgeOperatorOf(&_BridgeManager.CallOpts, governors)
}

// GetBridgeOperatorWeight is a free data retrieval call binding the contract method 0x901979d5.
//
// Solidity: function getBridgeOperatorWeight(address bridgeOperator) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerCaller) GetBridgeOperatorWeight(opts *bind.CallOpts, bridgeOperator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getBridgeOperatorWeight", bridgeOperator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBridgeOperatorWeight is a free data retrieval call binding the contract method 0x901979d5.
//
// Solidity: function getBridgeOperatorWeight(address bridgeOperator) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerSession) GetBridgeOperatorWeight(bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetBridgeOperatorWeight(&_BridgeManager.CallOpts, bridgeOperator)
}

// GetBridgeOperatorWeight is a free data retrieval call binding the contract method 0x901979d5.
//
// Solidity: function getBridgeOperatorWeight(address bridgeOperator) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerCallerSession) GetBridgeOperatorWeight(bridgeOperator common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetBridgeOperatorWeight(&_BridgeManager.CallOpts, bridgeOperator)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[])
func (_BridgeManager *BridgeManagerCaller) GetBridgeOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getBridgeOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[])
func (_BridgeManager *BridgeManagerSession) GetBridgeOperators() ([]common.Address, error) {
	return _BridgeManager.Contract.GetBridgeOperators(&_BridgeManager.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[])
func (_BridgeManager *BridgeManagerCallerSession) GetBridgeOperators() ([]common.Address, error) {
	return _BridgeManager.Contract.GetBridgeOperators(&_BridgeManager.CallOpts)
}

// GetCallbackRegisters is a free data retrieval call binding the contract method 0x0f7c3189.
//
// Solidity: function getCallbackRegisters() view returns(address[] registers)
func (_BridgeManager *BridgeManagerCaller) GetCallbackRegisters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getCallbackRegisters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCallbackRegisters is a free data retrieval call binding the contract method 0x0f7c3189.
//
// Solidity: function getCallbackRegisters() view returns(address[] registers)
func (_BridgeManager *BridgeManagerSession) GetCallbackRegisters() ([]common.Address, error) {
	return _BridgeManager.Contract.GetCallbackRegisters(&_BridgeManager.CallOpts)
}

// GetCallbackRegisters is a free data retrieval call binding the contract method 0x0f7c3189.
//
// Solidity: function getCallbackRegisters() view returns(address[] registers)
func (_BridgeManager *BridgeManagerCallerSession) GetCallbackRegisters() ([]common.Address, error) {
	return _BridgeManager.Contract.GetCallbackRegisters(&_BridgeManager.CallOpts)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeManager *BridgeManagerCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeManager *BridgeManagerSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeManager.Contract.GetContract(&_BridgeManager.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeManager *BridgeManagerCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeManager.Contract.GetContract(&_BridgeManager.CallOpts, contractType)
}

// GetFullBridgeOperatorInfos is a free data retrieval call binding the contract method 0xc441c4a8.
//
// Solidity: function getFullBridgeOperatorInfos() view returns(address[] governors, address[] bridgeOperators, uint256[] weights)
func (_BridgeManager *BridgeManagerCaller) GetFullBridgeOperatorInfos(opts *bind.CallOpts) (struct {
	Governors       []common.Address
	BridgeOperators []common.Address
	Weights         []*big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getFullBridgeOperatorInfos")

	outstruct := new(struct {
		Governors       []common.Address
		BridgeOperators []common.Address
		Weights         []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Governors = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.BridgeOperators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Weights = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetFullBridgeOperatorInfos is a free data retrieval call binding the contract method 0xc441c4a8.
//
// Solidity: function getFullBridgeOperatorInfos() view returns(address[] governors, address[] bridgeOperators, uint256[] weights)
func (_BridgeManager *BridgeManagerSession) GetFullBridgeOperatorInfos() (struct {
	Governors       []common.Address
	BridgeOperators []common.Address
	Weights         []*big.Int
}, error) {
	return _BridgeManager.Contract.GetFullBridgeOperatorInfos(&_BridgeManager.CallOpts)
}

// GetFullBridgeOperatorInfos is a free data retrieval call binding the contract method 0xc441c4a8.
//
// Solidity: function getFullBridgeOperatorInfos() view returns(address[] governors, address[] bridgeOperators, uint256[] weights)
func (_BridgeManager *BridgeManagerCallerSession) GetFullBridgeOperatorInfos() (struct {
	Governors       []common.Address
	BridgeOperators []common.Address
	Weights         []*big.Int
}, error) {
	return _BridgeManager.Contract.GetFullBridgeOperatorInfos(&_BridgeManager.CallOpts)
}

// GetGlobalProposalSignatures is a free data retrieval call binding the contract method 0xbc4e068f.
//
// Solidity: function getGlobalProposalSignatures(uint256 round_) view returns(address[] voters, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures)
func (_BridgeManager *BridgeManagerCaller) GetGlobalProposalSignatures(opts *bind.CallOpts, round_ *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGlobalProposalSignatures", round_)

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

// GetGlobalProposalSignatures is a free data retrieval call binding the contract method 0xbc4e068f.
//
// Solidity: function getGlobalProposalSignatures(uint256 round_) view returns(address[] voters, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures)
func (_BridgeManager *BridgeManagerSession) GetGlobalProposalSignatures(round_ *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _BridgeManager.Contract.GetGlobalProposalSignatures(&_BridgeManager.CallOpts, round_)
}

// GetGlobalProposalSignatures is a free data retrieval call binding the contract method 0xbc4e068f.
//
// Solidity: function getGlobalProposalSignatures(uint256 round_) view returns(address[] voters, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures)
func (_BridgeManager *BridgeManagerCallerSession) GetGlobalProposalSignatures(round_ *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _BridgeManager.Contract.GetGlobalProposalSignatures(&_BridgeManager.CallOpts, round_)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address governor) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerCaller) GetGovernorWeight(opts *bind.CallOpts, governor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGovernorWeight", governor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address governor) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerSession) GetGovernorWeight(governor common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetGovernorWeight(&_BridgeManager.CallOpts, governor)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address governor) view returns(uint256 weight)
func (_BridgeManager *BridgeManagerCallerSession) GetGovernorWeight(governor common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.GetGovernorWeight(&_BridgeManager.CallOpts, governor)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] governors) view returns(uint256[] weights)
func (_BridgeManager *BridgeManagerCaller) GetGovernorWeights(opts *bind.CallOpts, governors []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGovernorWeights", governors)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] governors) view returns(uint256[] weights)
func (_BridgeManager *BridgeManagerSession) GetGovernorWeights(governors []common.Address) ([]*big.Int, error) {
	return _BridgeManager.Contract.GetGovernorWeights(&_BridgeManager.CallOpts, governors)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] governors) view returns(uint256[] weights)
func (_BridgeManager *BridgeManagerCallerSession) GetGovernorWeights(governors []common.Address) ([]*big.Int, error) {
	return _BridgeManager.Contract.GetGovernorWeights(&_BridgeManager.CallOpts, governors)
}

// GetGovernors is a free data retrieval call binding the contract method 0xf80b5352.
//
// Solidity: function getGovernors() view returns(address[])
func (_BridgeManager *BridgeManagerCaller) GetGovernors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGovernors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGovernors is a free data retrieval call binding the contract method 0xf80b5352.
//
// Solidity: function getGovernors() view returns(address[])
func (_BridgeManager *BridgeManagerSession) GetGovernors() ([]common.Address, error) {
	return _BridgeManager.Contract.GetGovernors(&_BridgeManager.CallOpts)
}

// GetGovernors is a free data retrieval call binding the contract method 0xf80b5352.
//
// Solidity: function getGovernors() view returns(address[])
func (_BridgeManager *BridgeManagerCallerSession) GetGovernors() ([]common.Address, error) {
	return _BridgeManager.Contract.GetGovernors(&_BridgeManager.CallOpts)
}

// GetGovernorsOf is a free data retrieval call binding the contract method 0xfdc4fa47.
//
// Solidity: function getGovernorsOf(address[] bridgeOperators) view returns(address[] governors)
func (_BridgeManager *BridgeManagerCaller) GetGovernorsOf(opts *bind.CallOpts, bridgeOperators []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getGovernorsOf", bridgeOperators)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGovernorsOf is a free data retrieval call binding the contract method 0xfdc4fa47.
//
// Solidity: function getGovernorsOf(address[] bridgeOperators) view returns(address[] governors)
func (_BridgeManager *BridgeManagerSession) GetGovernorsOf(bridgeOperators []common.Address) ([]common.Address, error) {
	return _BridgeManager.Contract.GetGovernorsOf(&_BridgeManager.CallOpts, bridgeOperators)
}

// GetGovernorsOf is a free data retrieval call binding the contract method 0xfdc4fa47.
//
// Solidity: function getGovernorsOf(address[] bridgeOperators) view returns(address[] governors)
func (_BridgeManager *BridgeManagerCallerSession) GetGovernorsOf(bridgeOperators []common.Address) ([]common.Address, error) {
	return _BridgeManager.Contract.GetGovernorsOf(&_BridgeManager.CallOpts, bridgeOperators)
}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetProposalExpiryDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getProposalExpiryDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetProposalExpiryDuration() (*big.Int, error) {
	return _BridgeManager.Contract.GetProposalExpiryDuration(&_BridgeManager.CallOpts)
}

// GetProposalExpiryDuration is a free data retrieval call binding the contract method 0xbc96180b.
//
// Solidity: function getProposalExpiryDuration() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetProposalExpiryDuration() (*big.Int, error) {
	return _BridgeManager.Contract.GetProposalExpiryDuration(&_BridgeManager.CallOpts)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x1c905e39.
//
// Solidity: function getProposalSignatures(uint256 _chainId, uint256 _round) view returns(address[] _voters, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures)
func (_BridgeManager *BridgeManagerCaller) GetProposalSignatures(opts *bind.CallOpts, _chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getProposalSignatures", _chainId, _round)

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
func (_BridgeManager *BridgeManagerSession) GetProposalSignatures(_chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _BridgeManager.Contract.GetProposalSignatures(&_BridgeManager.CallOpts, _chainId, _round)
}

// GetProposalSignatures is a free data retrieval call binding the contract method 0x1c905e39.
//
// Solidity: function getProposalSignatures(uint256 _chainId, uint256 _round) view returns(address[] _voters, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures)
func (_BridgeManager *BridgeManagerCallerSession) GetProposalSignatures(_chainId *big.Int, _round *big.Int) (struct {
	Voters     []common.Address
	Supports   []uint8
	Signatures []SignatureConsumerSignature
}, error) {
	return _BridgeManager.Contract.GetProposalSignatures(&_BridgeManager.CallOpts, _chainId, _round)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_BridgeManager *BridgeManagerCaller) GetThreshold(opts *bind.CallOpts) (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getThreshold")

	outstruct := new(struct {
		Num   *big.Int
		Denom *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Num = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Denom = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_BridgeManager *BridgeManagerSession) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _BridgeManager.Contract.GetThreshold(&_BridgeManager.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256 num_, uint256 denom_)
func (_BridgeManager *BridgeManagerCallerSession) GetThreshold() (struct {
	Num   *big.Int
	Denom *big.Int
}, error) {
	return _BridgeManager.Contract.GetThreshold(&_BridgeManager.CallOpts)
}

// GetTotalWeights is a free data retrieval call binding the contract method 0xada86b24.
//
// Solidity: function getTotalWeights() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) GetTotalWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "getTotalWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWeights is a free data retrieval call binding the contract method 0xada86b24.
//
// Solidity: function getTotalWeights() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) GetTotalWeights() (*big.Int, error) {
	return _BridgeManager.Contract.GetTotalWeights(&_BridgeManager.CallOpts)
}

// GetTotalWeights is a free data retrieval call binding the contract method 0xada86b24.
//
// Solidity: function getTotalWeights() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) GetTotalWeights() (*big.Int, error) {
	return _BridgeManager.Contract.GetTotalWeights(&_BridgeManager.CallOpts)
}

// GlobalProposalVoted is a free data retrieval call binding the contract method 0x828fc1a1.
//
// Solidity: function globalProposalVoted(uint256 round_, address voter) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) GlobalProposalVoted(opts *bind.CallOpts, round_ *big.Int, voter common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "globalProposalVoted", round_, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalProposalVoted is a free data retrieval call binding the contract method 0x828fc1a1.
//
// Solidity: function globalProposalVoted(uint256 round_, address voter) view returns(bool)
func (_BridgeManager *BridgeManagerSession) GlobalProposalVoted(round_ *big.Int, voter common.Address) (bool, error) {
	return _BridgeManager.Contract.GlobalProposalVoted(&_BridgeManager.CallOpts, round_, voter)
}

// GlobalProposalVoted is a free data retrieval call binding the contract method 0x828fc1a1.
//
// Solidity: function globalProposalVoted(uint256 round_, address voter) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) GlobalProposalVoted(round_ *big.Int, voter common.Address) (bool, error) {
	return _BridgeManager.Contract.GlobalProposalVoted(&_BridgeManager.CallOpts, round_, voter)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address addr) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) IsBridgeOperator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "isBridgeOperator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address addr) view returns(bool)
func (_BridgeManager *BridgeManagerSession) IsBridgeOperator(addr common.Address) (bool, error) {
	return _BridgeManager.Contract.IsBridgeOperator(&_BridgeManager.CallOpts, addr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address addr) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) IsBridgeOperator(addr common.Address) (bool, error) {
	return _BridgeManager.Contract.IsBridgeOperator(&_BridgeManager.CallOpts, addr)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) MinimumVoteWeight() (*big.Int, error) {
	return _BridgeManager.Contract.MinimumVoteWeight(&_BridgeManager.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _BridgeManager.Contract.MinimumVoteWeight(&_BridgeManager.CallOpts)
}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_BridgeManager *BridgeManagerCaller) ProposalVoted(opts *bind.CallOpts, _chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "proposalVoted", _chainId, _round, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_BridgeManager *BridgeManagerSession) ProposalVoted(_chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	return _BridgeManager.Contract.ProposalVoted(&_BridgeManager.CallOpts, _chainId, _round, _voter)
}

// ProposalVoted is a free data retrieval call binding the contract method 0x2c5e6520.
//
// Solidity: function proposalVoted(uint256 _chainId, uint256 _round, address _voter) view returns(bool)
func (_BridgeManager *BridgeManagerCallerSession) ProposalVoted(_chainId *big.Int, _round *big.Int, _voter common.Address) (bool, error) {
	return _BridgeManager.Contract.ProposalVoted(&_BridgeManager.CallOpts, _chainId, _round, _voter)
}

// ResolveTargets is a free data retrieval call binding the contract method 0x2d6d7d73.
//
// Solidity: function resolveTargets(uint8[] targetOptions) view returns(address[] targets)
func (_BridgeManager *BridgeManagerCaller) ResolveTargets(opts *bind.CallOpts, targetOptions []uint8) ([]common.Address, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "resolveTargets", targetOptions)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ResolveTargets is a free data retrieval call binding the contract method 0x2d6d7d73.
//
// Solidity: function resolveTargets(uint8[] targetOptions) view returns(address[] targets)
func (_BridgeManager *BridgeManagerSession) ResolveTargets(targetOptions []uint8) ([]common.Address, error) {
	return _BridgeManager.Contract.ResolveTargets(&_BridgeManager.CallOpts, targetOptions)
}

// ResolveTargets is a free data retrieval call binding the contract method 0x2d6d7d73.
//
// Solidity: function resolveTargets(uint8[] targetOptions) view returns(address[] targets)
func (_BridgeManager *BridgeManagerCallerSession) ResolveTargets(targetOptions []uint8) ([]common.Address, error) {
	return _BridgeManager.Contract.ResolveTargets(&_BridgeManager.CallOpts, targetOptions)
}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "round", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_BridgeManager *BridgeManagerSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _BridgeManager.Contract.Round(&_BridgeManager.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x34d5f37b.
//
// Solidity: function round(uint256 ) view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) Round(arg0 *big.Int) (*big.Int, error) {
	return _BridgeManager.Contract.Round(&_BridgeManager.CallOpts, arg0)
}

// SumGovernorsWeight is a free data retrieval call binding the contract method 0x0a44fa43.
//
// Solidity: function sumGovernorsWeight(address[] governors) view returns(uint256 sum)
func (_BridgeManager *BridgeManagerCaller) SumGovernorsWeight(opts *bind.CallOpts, governors []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "sumGovernorsWeight", governors)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumGovernorsWeight is a free data retrieval call binding the contract method 0x0a44fa43.
//
// Solidity: function sumGovernorsWeight(address[] governors) view returns(uint256 sum)
func (_BridgeManager *BridgeManagerSession) SumGovernorsWeight(governors []common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.SumGovernorsWeight(&_BridgeManager.CallOpts, governors)
}

// SumGovernorsWeight is a free data retrieval call binding the contract method 0x0a44fa43.
//
// Solidity: function sumGovernorsWeight(address[] governors) view returns(uint256 sum)
func (_BridgeManager *BridgeManagerCallerSession) SumGovernorsWeight(governors []common.Address) (*big.Int, error) {
	return _BridgeManager.Contract.SumGovernorsWeight(&_BridgeManager.CallOpts, governors)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_BridgeManager *BridgeManagerCaller) TotalBridgeOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "totalBridgeOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_BridgeManager *BridgeManagerSession) TotalBridgeOperators() (*big.Int, error) {
	return _BridgeManager.Contract.TotalBridgeOperators(&_BridgeManager.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256)
func (_BridgeManager *BridgeManagerCallerSession) TotalBridgeOperators() (*big.Int, error) {
	return _BridgeManager.Contract.TotalBridgeOperators(&_BridgeManager.CallOpts)
}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight, uint256 expiryTimestamp)
func (_BridgeManager *BridgeManagerCaller) Vote(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	var out []interface{}
	err := _BridgeManager.contract.Call(opts, &out, "vote", arg0, arg1)

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
func (_BridgeManager *BridgeManagerSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	return _BridgeManager.Contract.Vote(&_BridgeManager.CallOpts, arg0, arg1)
}

// Vote is a free data retrieval call binding the contract method 0xb384abef.
//
// Solidity: function vote(uint256 , uint256 ) view returns(uint8 status, bytes32 hash, uint256 againstVoteWeight, uint256 forVoteWeight, uint256 expiryTimestamp)
func (_BridgeManager *BridgeManagerCallerSession) Vote(arg0 *big.Int, arg1 *big.Int) (struct {
	Status            uint8
	Hash              [32]byte
	AgainstVoteWeight *big.Int
	ForVoteWeight     *big.Int
	ExpiryTimestamp   *big.Int
}, error) {
	return _BridgeManager.Contract.Vote(&_BridgeManager.CallOpts, arg0, arg1)
}

// AddBridgeOperators is a paid mutator transaction binding the contract method 0x01a5f43f.
//
// Solidity: function addBridgeOperators(uint96[] voteWeights, address[] governors, address[] bridgeOperators) returns(bool[] addeds)
func (_BridgeManager *BridgeManagerTransactor) AddBridgeOperators(opts *bind.TransactOpts, voteWeights []*big.Int, governors []common.Address, bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "addBridgeOperators", voteWeights, governors, bridgeOperators)
}

// AddBridgeOperators is a paid mutator transaction binding the contract method 0x01a5f43f.
//
// Solidity: function addBridgeOperators(uint96[] voteWeights, address[] governors, address[] bridgeOperators) returns(bool[] addeds)
func (_BridgeManager *BridgeManagerSession) AddBridgeOperators(voteWeights []*big.Int, governors []common.Address, bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.AddBridgeOperators(&_BridgeManager.TransactOpts, voteWeights, governors, bridgeOperators)
}

// AddBridgeOperators is a paid mutator transaction binding the contract method 0x01a5f43f.
//
// Solidity: function addBridgeOperators(uint96[] voteWeights, address[] governors, address[] bridgeOperators) returns(bool[] addeds)
func (_BridgeManager *BridgeManagerTransactorSession) AddBridgeOperators(voteWeights []*big.Int, governors []common.Address, bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.AddBridgeOperators(&_BridgeManager.TransactOpts, voteWeights, governors, bridgeOperators)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactor) CastGlobalProposalBySignatures(opts *bind.TransactOpts, globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "castGlobalProposalBySignatures", globalProposal, supports_, signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerSession) CastGlobalProposalBySignatures(globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastGlobalProposalBySignatures(&_BridgeManager.TransactOpts, globalProposal, supports_, signatures)
}

// CastGlobalProposalBySignatures is a paid mutator transaction binding the contract method 0x2faf925d.
//
// Solidity: function castGlobalProposalBySignatures((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactorSession) CastGlobalProposalBySignatures(globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastGlobalProposalBySignatures(&_BridgeManager.TransactOpts, globalProposal, supports_, signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactor) CastProposalBySignatures(opts *bind.TransactOpts, proposal ProposalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "castProposalBySignatures", proposal, supports_, signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerSession) CastProposalBySignatures(proposal ProposalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastProposalBySignatures(&_BridgeManager.TransactOpts, proposal, supports_, signatures)
}

// CastProposalBySignatures is a paid mutator transaction binding the contract method 0x0b881830.
//
// Solidity: function castProposalBySignatures((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactorSession) CastProposalBySignatures(proposal ProposalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastProposalBySignatures(&_BridgeManager.TransactOpts, proposal, supports_, signatures)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8 support) returns()
func (_BridgeManager *BridgeManagerTransactor) CastProposalVoteForCurrentNetwork(opts *bind.TransactOpts, proposal ProposalProposalDetail, support uint8) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "castProposalVoteForCurrentNetwork", proposal, support)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8 support) returns()
func (_BridgeManager *BridgeManagerSession) CastProposalVoteForCurrentNetwork(proposal ProposalProposalDetail, support uint8) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastProposalVoteForCurrentNetwork(&_BridgeManager.TransactOpts, proposal, support)
}

// CastProposalVoteForCurrentNetwork is a paid mutator transaction binding the contract method 0xa8a0e32c.
//
// Solidity: function castProposalVoteForCurrentNetwork((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, uint8 support) returns()
func (_BridgeManager *BridgeManagerTransactorSession) CastProposalVoteForCurrentNetwork(proposal ProposalProposalDetail, support uint8) (*types.Transaction, error) {
	return _BridgeManager.Contract.CastProposalVoteForCurrentNetwork(&_BridgeManager.TransactOpts, proposal, support)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 _chainId, uint256 _round) returns()
func (_BridgeManager *BridgeManagerTransactor) DeleteExpired(opts *bind.TransactOpts, _chainId *big.Int, _round *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "deleteExpired", _chainId, _round)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 _chainId, uint256 _round) returns()
func (_BridgeManager *BridgeManagerSession) DeleteExpired(_chainId *big.Int, _round *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.DeleteExpired(&_BridgeManager.TransactOpts, _chainId, _round)
}

// DeleteExpired is a paid mutator transaction binding the contract method 0x9a7d3382.
//
// Solidity: function deleteExpired(uint256 _chainId, uint256 _round) returns()
func (_BridgeManager *BridgeManagerTransactorSession) DeleteExpired(_chainId *big.Int, _round *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.DeleteExpired(&_BridgeManager.TransactOpts, _chainId, _round)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_BridgeManager *BridgeManagerTransactor) Propose(opts *bind.TransactOpts, _chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "propose", _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_BridgeManager *BridgeManagerSession) Propose(_chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.Propose(&_BridgeManager.TransactOpts, _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// Propose is a paid mutator transaction binding the contract method 0xa1819f9a.
//
// Solidity: function propose(uint256 _chainId, uint256 _expiryTimestamp, address[] _targets, uint256[] _values, bytes[] _calldatas, uint256[] _gasAmounts) returns()
func (_BridgeManager *BridgeManagerTransactorSession) Propose(_chainId *big.Int, _expiryTimestamp *big.Int, _targets []common.Address, _values []*big.Int, _calldatas [][]byte, _gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.Propose(&_BridgeManager.TransactOpts, _chainId, _expiryTimestamp, _targets, _values, _calldatas, _gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 expiryTimestamp, uint8[] targetOptions, uint256[] values, bytes[] calldatas, uint256[] gasAmounts) returns()
func (_BridgeManager *BridgeManagerTransactor) ProposeGlobal(opts *bind.TransactOpts, expiryTimestamp *big.Int, targetOptions []uint8, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "proposeGlobal", expiryTimestamp, targetOptions, values, calldatas, gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 expiryTimestamp, uint8[] targetOptions, uint256[] values, bytes[] calldatas, uint256[] gasAmounts) returns()
func (_BridgeManager *BridgeManagerSession) ProposeGlobal(expiryTimestamp *big.Int, targetOptions []uint8, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeGlobal(&_BridgeManager.TransactOpts, expiryTimestamp, targetOptions, values, calldatas, gasAmounts)
}

// ProposeGlobal is a paid mutator transaction binding the contract method 0x09fcd8c7.
//
// Solidity: function proposeGlobal(uint256 expiryTimestamp, uint8[] targetOptions, uint256[] values, bytes[] calldatas, uint256[] gasAmounts) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ProposeGlobal(expiryTimestamp *big.Int, targetOptions []uint8, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeGlobal(&_BridgeManager.TransactOpts, expiryTimestamp, targetOptions, values, calldatas, gasAmounts)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactor) ProposeGlobalProposalStructAndCastVotes(opts *bind.TransactOpts, globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "proposeGlobalProposalStructAndCastVotes", globalProposal, supports_, signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerSession) ProposeGlobalProposalStructAndCastVotes(globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeGlobalProposalStructAndCastVotes(&_BridgeManager.TransactOpts, globalProposal, supports_, signatures)
}

// ProposeGlobalProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0xfb4f6371.
//
// Solidity: function proposeGlobalProposalStructAndCastVotes((uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, uint8[] supports_, (uint8,bytes32,bytes32)[] signatures) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ProposeGlobalProposalStructAndCastVotes(globalProposal GlobalProposalGlobalProposalDetail, supports_ []uint8, signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeGlobalProposalStructAndCastVotes(&_BridgeManager.TransactOpts, globalProposal, supports_, signatures)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 expiryTimestamp, address[] targets, uint256[] values, bytes[] calldatas, uint256[] gasAmounts, uint8 support) returns()
func (_BridgeManager *BridgeManagerTransactor) ProposeProposalForCurrentNetwork(opts *bind.TransactOpts, expiryTimestamp *big.Int, targets []common.Address, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int, support uint8) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "proposeProposalForCurrentNetwork", expiryTimestamp, targets, values, calldatas, gasAmounts, support)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 expiryTimestamp, address[] targets, uint256[] values, bytes[] calldatas, uint256[] gasAmounts, uint8 support) returns()
func (_BridgeManager *BridgeManagerSession) ProposeProposalForCurrentNetwork(expiryTimestamp *big.Int, targets []common.Address, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int, support uint8) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeProposalForCurrentNetwork(&_BridgeManager.TransactOpts, expiryTimestamp, targets, values, calldatas, gasAmounts, support)
}

// ProposeProposalForCurrentNetwork is a paid mutator transaction binding the contract method 0x663ac011.
//
// Solidity: function proposeProposalForCurrentNetwork(uint256 expiryTimestamp, address[] targets, uint256[] values, bytes[] calldatas, uint256[] gasAmounts, uint8 support) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ProposeProposalForCurrentNetwork(expiryTimestamp *big.Int, targets []common.Address, values []*big.Int, calldatas [][]byte, gasAmounts []*big.Int, support uint8) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeProposalForCurrentNetwork(&_BridgeManager.TransactOpts, expiryTimestamp, targets, values, calldatas, gasAmounts, support)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_BridgeManager *BridgeManagerTransactor) ProposeProposalStructAndCastVotes(opts *bind.TransactOpts, _proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "proposeProposalStructAndCastVotes", _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_BridgeManager *BridgeManagerSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeProposalStructAndCastVotes(&_BridgeManager.TransactOpts, _proposal, _supports, _signatures)
}

// ProposeProposalStructAndCastVotes is a paid mutator transaction binding the contract method 0x004054b8.
//
// Solidity: function proposeProposalStructAndCastVotes((uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) _proposal, uint8[] _supports, (uint8,bytes32,bytes32)[] _signatures) returns()
func (_BridgeManager *BridgeManagerTransactorSession) ProposeProposalStructAndCastVotes(_proposal ProposalProposalDetail, _supports []uint8, _signatures []SignatureConsumerSignature) (*types.Transaction, error) {
	return _BridgeManager.Contract.ProposeProposalStructAndCastVotes(&_BridgeManager.TransactOpts, _proposal, _supports, _signatures)
}

// RegisterCallbacks is a paid mutator transaction binding the contract method 0x35da8121.
//
// Solidity: function registerCallbacks(address[] registers) returns(bool[] registereds)
func (_BridgeManager *BridgeManagerTransactor) RegisterCallbacks(opts *bind.TransactOpts, registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "registerCallbacks", registers)
}

// RegisterCallbacks is a paid mutator transaction binding the contract method 0x35da8121.
//
// Solidity: function registerCallbacks(address[] registers) returns(bool[] registereds)
func (_BridgeManager *BridgeManagerSession) RegisterCallbacks(registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RegisterCallbacks(&_BridgeManager.TransactOpts, registers)
}

// RegisterCallbacks is a paid mutator transaction binding the contract method 0x35da8121.
//
// Solidity: function registerCallbacks(address[] registers) returns(bool[] registereds)
func (_BridgeManager *BridgeManagerTransactorSession) RegisterCallbacks(registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RegisterCallbacks(&_BridgeManager.TransactOpts, registers)
}

// RemoveBridgeOperators is a paid mutator transaction binding the contract method 0xe9c03498.
//
// Solidity: function removeBridgeOperators(address[] bridgeOperators) returns(bool[] removeds)
func (_BridgeManager *BridgeManagerTransactor) RemoveBridgeOperators(opts *bind.TransactOpts, bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "removeBridgeOperators", bridgeOperators)
}

// RemoveBridgeOperators is a paid mutator transaction binding the contract method 0xe9c03498.
//
// Solidity: function removeBridgeOperators(address[] bridgeOperators) returns(bool[] removeds)
func (_BridgeManager *BridgeManagerSession) RemoveBridgeOperators(bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RemoveBridgeOperators(&_BridgeManager.TransactOpts, bridgeOperators)
}

// RemoveBridgeOperators is a paid mutator transaction binding the contract method 0xe9c03498.
//
// Solidity: function removeBridgeOperators(address[] bridgeOperators) returns(bool[] removeds)
func (_BridgeManager *BridgeManagerTransactorSession) RemoveBridgeOperators(bridgeOperators []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.RemoveBridgeOperators(&_BridgeManager.TransactOpts, bridgeOperators)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeManager *BridgeManagerTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeManager *BridgeManagerSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetContract(&_BridgeManager.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeManager *BridgeManagerTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetContract(&_BridgeManager.TransactOpts, contractType, addr)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 numerator, uint256 denominator) returns(uint256, uint256)
func (_BridgeManager *BridgeManagerTransactor) SetThreshold(opts *bind.TransactOpts, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "setThreshold", numerator, denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 numerator, uint256 denominator) returns(uint256, uint256)
func (_BridgeManager *BridgeManagerSession) SetThreshold(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetThreshold(&_BridgeManager.TransactOpts, numerator, denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 numerator, uint256 denominator) returns(uint256, uint256)
func (_BridgeManager *BridgeManagerTransactorSession) SetThreshold(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _BridgeManager.Contract.SetThreshold(&_BridgeManager.TransactOpts, numerator, denominator)
}

// UnregisterCallbacks is a paid mutator transaction binding the contract method 0x1f425338.
//
// Solidity: function unregisterCallbacks(address[] registers) returns(bool[] unregistereds)
func (_BridgeManager *BridgeManagerTransactor) UnregisterCallbacks(opts *bind.TransactOpts, registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "unregisterCallbacks", registers)
}

// UnregisterCallbacks is a paid mutator transaction binding the contract method 0x1f425338.
//
// Solidity: function unregisterCallbacks(address[] registers) returns(bool[] unregistereds)
func (_BridgeManager *BridgeManagerSession) UnregisterCallbacks(registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UnregisterCallbacks(&_BridgeManager.TransactOpts, registers)
}

// UnregisterCallbacks is a paid mutator transaction binding the contract method 0x1f425338.
//
// Solidity: function unregisterCallbacks(address[] registers) returns(bool[] unregistereds)
func (_BridgeManager *BridgeManagerTransactorSession) UnregisterCallbacks(registers []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UnregisterCallbacks(&_BridgeManager.TransactOpts, registers)
}

// UpdateBridgeOperator is a paid mutator transaction binding the contract method 0x9b2ee437.
//
// Solidity: function updateBridgeOperator(address newBridgeOperator) returns()
func (_BridgeManager *BridgeManagerTransactor) UpdateBridgeOperator(opts *bind.TransactOpts, newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "updateBridgeOperator", newBridgeOperator)
}

// UpdateBridgeOperator is a paid mutator transaction binding the contract method 0x9b2ee437.
//
// Solidity: function updateBridgeOperator(address newBridgeOperator) returns()
func (_BridgeManager *BridgeManagerSession) UpdateBridgeOperator(newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateBridgeOperator(&_BridgeManager.TransactOpts, newBridgeOperator)
}

// UpdateBridgeOperator is a paid mutator transaction binding the contract method 0x9b2ee437.
//
// Solidity: function updateBridgeOperator(address newBridgeOperator) returns()
func (_BridgeManager *BridgeManagerTransactorSession) UpdateBridgeOperator(newBridgeOperator common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateBridgeOperator(&_BridgeManager.TransactOpts, newBridgeOperator)
}

// UpdateManyTargetOption is a paid mutator transaction binding the contract method 0x800eaab3.
//
// Solidity: function updateManyTargetOption(uint8[] targetOptions, address[] targets) returns()
func (_BridgeManager *BridgeManagerTransactor) UpdateManyTargetOption(opts *bind.TransactOpts, targetOptions []uint8, targets []common.Address) (*types.Transaction, error) {
	return _BridgeManager.contract.Transact(opts, "updateManyTargetOption", targetOptions, targets)
}

// UpdateManyTargetOption is a paid mutator transaction binding the contract method 0x800eaab3.
//
// Solidity: function updateManyTargetOption(uint8[] targetOptions, address[] targets) returns()
func (_BridgeManager *BridgeManagerSession) UpdateManyTargetOption(targetOptions []uint8, targets []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateManyTargetOption(&_BridgeManager.TransactOpts, targetOptions, targets)
}

// UpdateManyTargetOption is a paid mutator transaction binding the contract method 0x800eaab3.
//
// Solidity: function updateManyTargetOption(uint8[] targetOptions, address[] targets) returns()
func (_BridgeManager *BridgeManagerTransactorSession) UpdateManyTargetOption(targetOptions []uint8, targets []common.Address) (*types.Transaction, error) {
	return _BridgeManager.Contract.UpdateManyTargetOption(&_BridgeManager.TransactOpts, targetOptions, targets)
}

// BridgeManagerBridgeOperatorUpdatedIterator is returned from FilterBridgeOperatorUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorUpdated events raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorUpdatedIterator struct {
	Event *BridgeManagerBridgeOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeOperatorUpdated)
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
		it.Event = new(BridgeManagerBridgeOperatorUpdated)
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
func (it *BridgeManagerBridgeOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeOperatorUpdated represents a BridgeOperatorUpdated event raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorUpdated struct {
	Governor           common.Address
	FromBridgeOperator common.Address
	ToBridgeOperator   common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorUpdated is a free log retrieval operation binding the contract event 0xcef34cd748f30a1b7a2f214fd1651779f79bc6c1be02785cad5c1f0ee877213d.
//
// Solidity: event BridgeOperatorUpdated(address indexed governor, address indexed fromBridgeOperator, address indexed toBridgeOperator)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeOperatorUpdated(opts *bind.FilterOpts, governor []common.Address, fromBridgeOperator []common.Address, toBridgeOperator []common.Address) (*BridgeManagerBridgeOperatorUpdatedIterator, error) {

	var governorRule []interface{}
	for _, governorItem := range governor {
		governorRule = append(governorRule, governorItem)
	}
	var fromBridgeOperatorRule []interface{}
	for _, fromBridgeOperatorItem := range fromBridgeOperator {
		fromBridgeOperatorRule = append(fromBridgeOperatorRule, fromBridgeOperatorItem)
	}
	var toBridgeOperatorRule []interface{}
	for _, toBridgeOperatorItem := range toBridgeOperator {
		toBridgeOperatorRule = append(toBridgeOperatorRule, toBridgeOperatorItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeOperatorUpdated", governorRule, fromBridgeOperatorRule, toBridgeOperatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeOperatorUpdatedIterator{contract: _BridgeManager.contract, event: "BridgeOperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorUpdated is a free log subscription operation binding the contract event 0xcef34cd748f30a1b7a2f214fd1651779f79bc6c1be02785cad5c1f0ee877213d.
//
// Solidity: event BridgeOperatorUpdated(address indexed governor, address indexed fromBridgeOperator, address indexed toBridgeOperator)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeOperatorUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeOperatorUpdated, governor []common.Address, fromBridgeOperator []common.Address, toBridgeOperator []common.Address) (event.Subscription, error) {

	var governorRule []interface{}
	for _, governorItem := range governor {
		governorRule = append(governorRule, governorItem)
	}
	var fromBridgeOperatorRule []interface{}
	for _, fromBridgeOperatorItem := range fromBridgeOperator {
		fromBridgeOperatorRule = append(fromBridgeOperatorRule, fromBridgeOperatorItem)
	}
	var toBridgeOperatorRule []interface{}
	for _, toBridgeOperatorItem := range toBridgeOperator {
		toBridgeOperatorRule = append(toBridgeOperatorRule, toBridgeOperatorItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeOperatorUpdated", governorRule, fromBridgeOperatorRule, toBridgeOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeOperatorUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorUpdated", log); err != nil {
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

// ParseBridgeOperatorUpdated is a log parse operation binding the contract event 0xcef34cd748f30a1b7a2f214fd1651779f79bc6c1be02785cad5c1f0ee877213d.
//
// Solidity: event BridgeOperatorUpdated(address indexed governor, address indexed fromBridgeOperator, address indexed toBridgeOperator)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeOperatorUpdated(log types.Log) (*BridgeManagerBridgeOperatorUpdated, error) {
	event := new(BridgeManagerBridgeOperatorUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerBridgeOperatorsAddedIterator is returned from FilterBridgeOperatorsAdded and is used to iterate over the raw logs and unpacked data for BridgeOperatorsAdded events raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorsAddedIterator struct {
	Event *BridgeManagerBridgeOperatorsAdded // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeOperatorsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeOperatorsAdded)
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
		it.Event = new(BridgeManagerBridgeOperatorsAdded)
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
func (it *BridgeManagerBridgeOperatorsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeOperatorsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeOperatorsAdded represents a BridgeOperatorsAdded event raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorsAdded struct {
	Statuses        []bool
	VoteWeights     []*big.Int
	Governors       []common.Address
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorsAdded is a free log retrieval operation binding the contract event 0x897810999654e525e272b5909785c4d0ceaee1bbf9c87d9091a37558b0423b78.
//
// Solidity: event BridgeOperatorsAdded(bool[] statuses, uint96[] voteWeights, address[] governors, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeOperatorsAdded(opts *bind.FilterOpts) (*BridgeManagerBridgeOperatorsAddedIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeOperatorsAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeOperatorsAddedIterator{contract: _BridgeManager.contract, event: "BridgeOperatorsAdded", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorsAdded is a free log subscription operation binding the contract event 0x897810999654e525e272b5909785c4d0ceaee1bbf9c87d9091a37558b0423b78.
//
// Solidity: event BridgeOperatorsAdded(bool[] statuses, uint96[] voteWeights, address[] governors, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeOperatorsAdded(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeOperatorsAdded) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeOperatorsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeOperatorsAdded)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorsAdded", log); err != nil {
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

// ParseBridgeOperatorsAdded is a log parse operation binding the contract event 0x897810999654e525e272b5909785c4d0ceaee1bbf9c87d9091a37558b0423b78.
//
// Solidity: event BridgeOperatorsAdded(bool[] statuses, uint96[] voteWeights, address[] governors, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeOperatorsAdded(log types.Log) (*BridgeManagerBridgeOperatorsAdded, error) {
	event := new(BridgeManagerBridgeOperatorsAdded)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerBridgeOperatorsRemovedIterator is returned from FilterBridgeOperatorsRemoved and is used to iterate over the raw logs and unpacked data for BridgeOperatorsRemoved events raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorsRemovedIterator struct {
	Event *BridgeManagerBridgeOperatorsRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeManagerBridgeOperatorsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerBridgeOperatorsRemoved)
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
		it.Event = new(BridgeManagerBridgeOperatorsRemoved)
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
func (it *BridgeManagerBridgeOperatorsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerBridgeOperatorsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerBridgeOperatorsRemoved represents a BridgeOperatorsRemoved event raised by the BridgeManager contract.
type BridgeManagerBridgeOperatorsRemoved struct {
	Statuses        []bool
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorsRemoved is a free log retrieval operation binding the contract event 0xdf3dcd7987202f64648f3acdbf12401e3a2bb23e77e19f99826b5475cbb86369.
//
// Solidity: event BridgeOperatorsRemoved(bool[] statuses, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) FilterBridgeOperatorsRemoved(opts *bind.FilterOpts) (*BridgeManagerBridgeOperatorsRemovedIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "BridgeOperatorsRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerBridgeOperatorsRemovedIterator{contract: _BridgeManager.contract, event: "BridgeOperatorsRemoved", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorsRemoved is a free log subscription operation binding the contract event 0xdf3dcd7987202f64648f3acdbf12401e3a2bb23e77e19f99826b5475cbb86369.
//
// Solidity: event BridgeOperatorsRemoved(bool[] statuses, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) WatchBridgeOperatorsRemoved(opts *bind.WatchOpts, sink chan<- *BridgeManagerBridgeOperatorsRemoved) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "BridgeOperatorsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerBridgeOperatorsRemoved)
				if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorsRemoved", log); err != nil {
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

// ParseBridgeOperatorsRemoved is a log parse operation binding the contract event 0xdf3dcd7987202f64648f3acdbf12401e3a2bb23e77e19f99826b5475cbb86369.
//
// Solidity: event BridgeOperatorsRemoved(bool[] statuses, address[] bridgeOperators)
func (_BridgeManager *BridgeManagerFilterer) ParseBridgeOperatorsRemoved(log types.Log) (*BridgeManagerBridgeOperatorsRemoved, error) {
	event := new(BridgeManagerBridgeOperatorsRemoved)
	if err := _BridgeManager.contract.UnpackLog(event, "BridgeOperatorsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the BridgeManager contract.
type BridgeManagerContractUpdatedIterator struct {
	Event *BridgeManagerContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerContractUpdated)
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
		it.Event = new(BridgeManagerContractUpdated)
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
func (it *BridgeManagerContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerContractUpdated represents a ContractUpdated event raised by the BridgeManager contract.
type BridgeManagerContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeManager *BridgeManagerFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*BridgeManagerContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerContractUpdatedIterator{contract: _BridgeManager.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeManager *BridgeManagerFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerContractUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseContractUpdated(log types.Log) (*BridgeManagerContractUpdated, error) {
	event := new(BridgeManagerContractUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerGlobalProposalCreatedIterator is returned from FilterGlobalProposalCreated and is used to iterate over the raw logs and unpacked data for GlobalProposalCreated events raised by the BridgeManager contract.
type BridgeManagerGlobalProposalCreatedIterator struct {
	Event *BridgeManagerGlobalProposalCreated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerGlobalProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerGlobalProposalCreated)
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
		it.Event = new(BridgeManagerGlobalProposalCreated)
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
func (it *BridgeManagerGlobalProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerGlobalProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerGlobalProposalCreated represents a GlobalProposalCreated event raised by the BridgeManager contract.
type BridgeManagerGlobalProposalCreated struct {
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
func (_BridgeManager *BridgeManagerFilterer) FilterGlobalProposalCreated(opts *bind.FilterOpts, round []*big.Int, proposalHash [][32]byte) (*BridgeManagerGlobalProposalCreatedIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerGlobalProposalCreatedIterator{contract: _BridgeManager.contract, event: "GlobalProposalCreated", logs: logs, sub: sub}, nil
}

// WatchGlobalProposalCreated is a free log subscription operation binding the contract event 0x771d78ae9e5fca95a532fb0971d575d0ce9b59d14823c063e08740137e0e0eca.
//
// Solidity: event GlobalProposalCreated(uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, bytes32 globalProposalHash, (uint256,uint256,uint8[],uint256[],bytes[],uint256[]) globalProposal, address creator)
func (_BridgeManager *BridgeManagerFilterer) WatchGlobalProposalCreated(opts *bind.WatchOpts, sink chan<- *BridgeManagerGlobalProposalCreated, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "GlobalProposalCreated", roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerGlobalProposalCreated)
				if err := _BridgeManager.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseGlobalProposalCreated(log types.Log) (*BridgeManagerGlobalProposalCreated, error) {
	event := new(BridgeManagerGlobalProposalCreated)
	if err := _BridgeManager.contract.UnpackLog(event, "GlobalProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerNotifiedIterator is returned from FilterNotified and is used to iterate over the raw logs and unpacked data for Notified events raised by the BridgeManager contract.
type BridgeManagerNotifiedIterator struct {
	Event *BridgeManagerNotified // Event containing the contract specifics and raw log

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
func (it *BridgeManagerNotifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerNotified)
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
		it.Event = new(BridgeManagerNotified)
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
func (it *BridgeManagerNotifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerNotifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerNotified represents a Notified event raised by the BridgeManager contract.
type BridgeManagerNotified struct {
	CallData    []byte
	Registers   []common.Address
	Statuses    []bool
	ReturnDatas [][]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNotified is a free log retrieval operation binding the contract event 0xc0b07a27e66788f39cc91405f012f34066b16f31b4bda9438c52f2dae0cc5b63.
//
// Solidity: event Notified(bytes callData, address[] registers, bool[] statuses, bytes[] returnDatas)
func (_BridgeManager *BridgeManagerFilterer) FilterNotified(opts *bind.FilterOpts) (*BridgeManagerNotifiedIterator, error) {

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "Notified")
	if err != nil {
		return nil, err
	}
	return &BridgeManagerNotifiedIterator{contract: _BridgeManager.contract, event: "Notified", logs: logs, sub: sub}, nil
}

// WatchNotified is a free log subscription operation binding the contract event 0xc0b07a27e66788f39cc91405f012f34066b16f31b4bda9438c52f2dae0cc5b63.
//
// Solidity: event Notified(bytes callData, address[] registers, bool[] statuses, bytes[] returnDatas)
func (_BridgeManager *BridgeManagerFilterer) WatchNotified(opts *bind.WatchOpts, sink chan<- *BridgeManagerNotified) (event.Subscription, error) {

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "Notified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerNotified)
				if err := _BridgeManager.contract.UnpackLog(event, "Notified", log); err != nil {
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

// ParseNotified is a log parse operation binding the contract event 0xc0b07a27e66788f39cc91405f012f34066b16f31b4bda9438c52f2dae0cc5b63.
//
// Solidity: event Notified(bytes callData, address[] registers, bool[] statuses, bytes[] returnDatas)
func (_BridgeManager *BridgeManagerFilterer) ParseNotified(log types.Log) (*BridgeManagerNotified, error) {
	event := new(BridgeManagerNotified)
	if err := _BridgeManager.contract.UnpackLog(event, "Notified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the BridgeManager contract.
type BridgeManagerProposalApprovedIterator struct {
	Event *BridgeManagerProposalApproved // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalApproved)
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
		it.Event = new(BridgeManagerProposalApproved)
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
func (it *BridgeManagerProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalApproved represents a ProposalApproved event raised by the BridgeManager contract.
type BridgeManagerProposalApproved struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalApproved(opts *bind.FilterOpts, proposalHash [][32]byte) (*BridgeManagerProposalApprovedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalApproved", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalApprovedIterator{contract: _BridgeManager.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x5c819725ea53655a3b898f3df59b66489761935454e9212ca1e5ebd759953d0b.
//
// Solidity: event ProposalApproved(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalApproved, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalApproved", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalApproved)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalApproved(log types.Log) (*BridgeManagerProposalApproved, error) {
	event := new(BridgeManagerProposalApproved)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the BridgeManager contract.
type BridgeManagerProposalCreatedIterator struct {
	Event *BridgeManagerProposalCreated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalCreated)
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
		it.Event = new(BridgeManagerProposalCreated)
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
func (it *BridgeManagerProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalCreated represents a ProposalCreated event raised by the BridgeManager contract.
type BridgeManagerProposalCreated struct {
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
func (_BridgeManager *BridgeManagerFilterer) FilterProposalCreated(opts *bind.FilterOpts, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (*BridgeManagerProposalCreatedIterator, error) {

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

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalCreatedIterator{contract: _BridgeManager.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xa57d40f1496988cf60ab7c9d5ba4ff83647f67d3898d441a3aaf21b651678fd9.
//
// Solidity: event ProposalCreated(uint256 indexed chainId, uint256 indexed round, bytes32 indexed proposalHash, (uint256,uint256,uint256,address[],uint256[],bytes[],uint256[]) proposal, address creator)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalCreated, chainId []*big.Int, round []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalCreated", chainIdRule, roundRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalCreated)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalCreated(log types.Log) (*BridgeManagerProposalCreated, error) {
	event := new(BridgeManagerProposalCreated)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the BridgeManager contract.
type BridgeManagerProposalExecutedIterator struct {
	Event *BridgeManagerProposalExecuted // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalExecuted)
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
		it.Event = new(BridgeManagerProposalExecuted)
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
func (it *BridgeManagerProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalExecuted represents a ProposalExecuted event raised by the BridgeManager contract.
type BridgeManagerProposalExecuted struct {
	ProposalHash [32]byte
	SuccessCalls []bool
	ReturnDatas  [][]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0xe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash, bool[] successCalls, bytes[] returnDatas)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalHash [][32]byte) (*BridgeManagerProposalExecutedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalExecuted", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalExecutedIterator{contract: _BridgeManager.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0xe134987599ae266ec90edeff1b26125b287dbb57b10822649432d1bb26537fba.
//
// Solidity: event ProposalExecuted(bytes32 indexed proposalHash, bool[] successCalls, bytes[] returnDatas)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalExecuted, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalExecuted", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalExecuted)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalExecuted(log types.Log) (*BridgeManagerProposalExecuted, error) {
	event := new(BridgeManagerProposalExecuted)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalExpiredIterator is returned from FilterProposalExpired and is used to iterate over the raw logs and unpacked data for ProposalExpired events raised by the BridgeManager contract.
type BridgeManagerProposalExpiredIterator struct {
	Event *BridgeManagerProposalExpired // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalExpired)
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
		it.Event = new(BridgeManagerProposalExpired)
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
func (it *BridgeManagerProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalExpired represents a ProposalExpired event raised by the BridgeManager contract.
type BridgeManagerProposalExpired struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalExpired is a free log retrieval operation binding the contract event 0x58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a55.
//
// Solidity: event ProposalExpired(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalExpired(opts *bind.FilterOpts, proposalHash [][32]byte) (*BridgeManagerProposalExpiredIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalExpired", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalExpiredIterator{contract: _BridgeManager.contract, event: "ProposalExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExpired is a free log subscription operation binding the contract event 0x58f98006a7f2f253f8ae8f8b7cec9008ca05359633561cd7c22f3005682d4a55.
//
// Solidity: event ProposalExpired(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalExpired(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalExpired, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalExpired", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalExpired)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalExpired(log types.Log) (*BridgeManagerProposalExpired, error) {
	event := new(BridgeManagerProposalExpired)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalExpiryDurationChangedIterator is returned from FilterProposalExpiryDurationChanged and is used to iterate over the raw logs and unpacked data for ProposalExpiryDurationChanged events raised by the BridgeManager contract.
type BridgeManagerProposalExpiryDurationChangedIterator struct {
	Event *BridgeManagerProposalExpiryDurationChanged // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalExpiryDurationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalExpiryDurationChanged)
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
		it.Event = new(BridgeManagerProposalExpiryDurationChanged)
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
func (it *BridgeManagerProposalExpiryDurationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalExpiryDurationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalExpiryDurationChanged represents a ProposalExpiryDurationChanged event raised by the BridgeManager contract.
type BridgeManagerProposalExpiryDurationChanged struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposalExpiryDurationChanged is a free log retrieval operation binding the contract event 0xe5cd1c123a8cf63fa1b7229678db61fe8ae99dbbd27889370b6667c8cae97da1.
//
// Solidity: event ProposalExpiryDurationChanged(uint256 indexed duration)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalExpiryDurationChanged(opts *bind.FilterOpts, duration []*big.Int) (*BridgeManagerProposalExpiryDurationChangedIterator, error) {

	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalExpiryDurationChanged", durationRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalExpiryDurationChangedIterator{contract: _BridgeManager.contract, event: "ProposalExpiryDurationChanged", logs: logs, sub: sub}, nil
}

// WatchProposalExpiryDurationChanged is a free log subscription operation binding the contract event 0xe5cd1c123a8cf63fa1b7229678db61fe8ae99dbbd27889370b6667c8cae97da1.
//
// Solidity: event ProposalExpiryDurationChanged(uint256 indexed duration)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalExpiryDurationChanged(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalExpiryDurationChanged, duration []*big.Int) (event.Subscription, error) {

	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalExpiryDurationChanged", durationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalExpiryDurationChanged)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalExpiryDurationChanged", log); err != nil {
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

// ParseProposalExpiryDurationChanged is a log parse operation binding the contract event 0xe5cd1c123a8cf63fa1b7229678db61fe8ae99dbbd27889370b6667c8cae97da1.
//
// Solidity: event ProposalExpiryDurationChanged(uint256 indexed duration)
func (_BridgeManager *BridgeManagerFilterer) ParseProposalExpiryDurationChanged(log types.Log) (*BridgeManagerProposalExpiryDurationChanged, error) {
	event := new(BridgeManagerProposalExpiryDurationChanged)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalExpiryDurationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalRejectedIterator is returned from FilterProposalRejected and is used to iterate over the raw logs and unpacked data for ProposalRejected events raised by the BridgeManager contract.
type BridgeManagerProposalRejectedIterator struct {
	Event *BridgeManagerProposalRejected // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalRejected)
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
		it.Event = new(BridgeManagerProposalRejected)
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
func (it *BridgeManagerProposalRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalRejected represents a ProposalRejected event raised by the BridgeManager contract.
type BridgeManagerProposalRejected struct {
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalRejected is a free log retrieval operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalRejected(opts *bind.FilterOpts, proposalHash [][32]byte) (*BridgeManagerProposalRejectedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalRejectedIterator{contract: _BridgeManager.contract, event: "ProposalRejected", logs: logs, sub: sub}, nil
}

// WatchProposalRejected is a free log subscription operation binding the contract event 0x55295d4ce992922fa2e5ffbf3a3dcdb367de0a15e125ace083456017fd22060f.
//
// Solidity: event ProposalRejected(bytes32 indexed proposalHash)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalRejected(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalRejected, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalRejected", proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalRejected)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalRejected(log types.Log) (*BridgeManagerProposalRejected, error) {
	event := new(BridgeManagerProposalRejected)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the BridgeManager contract.
type BridgeManagerProposalVotedIterator struct {
	Event *BridgeManagerProposalVoted // Event containing the contract specifics and raw log

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
func (it *BridgeManagerProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerProposalVoted)
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
		it.Event = new(BridgeManagerProposalVoted)
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
func (it *BridgeManagerProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerProposalVoted represents a ProposalVoted event raised by the BridgeManager contract.
type BridgeManagerProposalVoted struct {
	ProposalHash [32]byte
	Voter        common.Address
	Support      uint8
	Weight       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_BridgeManager *BridgeManagerFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalHash [][32]byte, voter []common.Address) (*BridgeManagerProposalVotedIterator, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerProposalVotedIterator{contract: _BridgeManager.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0x1203f9e81c814a35f5f4cc24087b2a24c6fb7986a9f1406b68a9484882c93a23.
//
// Solidity: event ProposalVoted(bytes32 indexed proposalHash, address indexed voter, uint8 support, uint256 weight)
func (_BridgeManager *BridgeManagerFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *BridgeManagerProposalVoted, proposalHash [][32]byte, voter []common.Address) (event.Subscription, error) {

	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ProposalVoted", proposalHashRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerProposalVoted)
				if err := _BridgeManager.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_BridgeManager *BridgeManagerFilterer) ParseProposalVoted(log types.Log) (*BridgeManagerProposalVoted, error) {
	event := new(BridgeManagerProposalVoted)
	if err := _BridgeManager.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerTargetOptionUpdatedIterator is returned from FilterTargetOptionUpdated and is used to iterate over the raw logs and unpacked data for TargetOptionUpdated events raised by the BridgeManager contract.
type BridgeManagerTargetOptionUpdatedIterator struct {
	Event *BridgeManagerTargetOptionUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerTargetOptionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerTargetOptionUpdated)
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
		it.Event = new(BridgeManagerTargetOptionUpdated)
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
func (it *BridgeManagerTargetOptionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerTargetOptionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerTargetOptionUpdated represents a TargetOptionUpdated event raised by the BridgeManager contract.
type BridgeManagerTargetOptionUpdated struct {
	TargetOption uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTargetOptionUpdated is a free log retrieval operation binding the contract event 0x356c8c57e9e84b99b1cb58b13c985b2c979f78cbdf4d0fa70fe2a98bb09a099d.
//
// Solidity: event TargetOptionUpdated(uint8 indexed targetOption, address indexed addr)
func (_BridgeManager *BridgeManagerFilterer) FilterTargetOptionUpdated(opts *bind.FilterOpts, targetOption []uint8, addr []common.Address) (*BridgeManagerTargetOptionUpdatedIterator, error) {

	var targetOptionRule []interface{}
	for _, targetOptionItem := range targetOption {
		targetOptionRule = append(targetOptionRule, targetOptionItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "TargetOptionUpdated", targetOptionRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerTargetOptionUpdatedIterator{contract: _BridgeManager.contract, event: "TargetOptionUpdated", logs: logs, sub: sub}, nil
}

// WatchTargetOptionUpdated is a free log subscription operation binding the contract event 0x356c8c57e9e84b99b1cb58b13c985b2c979f78cbdf4d0fa70fe2a98bb09a099d.
//
// Solidity: event TargetOptionUpdated(uint8 indexed targetOption, address indexed addr)
func (_BridgeManager *BridgeManagerFilterer) WatchTargetOptionUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerTargetOptionUpdated, targetOption []uint8, addr []common.Address) (event.Subscription, error) {

	var targetOptionRule []interface{}
	for _, targetOptionItem := range targetOption {
		targetOptionRule = append(targetOptionRule, targetOptionItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "TargetOptionUpdated", targetOptionRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerTargetOptionUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "TargetOptionUpdated", log); err != nil {
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

// ParseTargetOptionUpdated is a log parse operation binding the contract event 0x356c8c57e9e84b99b1cb58b13c985b2c979f78cbdf4d0fa70fe2a98bb09a099d.
//
// Solidity: event TargetOptionUpdated(uint8 indexed targetOption, address indexed addr)
func (_BridgeManager *BridgeManagerFilterer) ParseTargetOptionUpdated(log types.Log) (*BridgeManagerTargetOptionUpdated, error) {
	event := new(BridgeManagerTargetOptionUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "TargetOptionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeManagerThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the BridgeManager contract.
type BridgeManagerThresholdUpdatedIterator struct {
	Event *BridgeManagerThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeManagerThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeManagerThresholdUpdated)
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
		it.Event = new(BridgeManagerThresholdUpdated)
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
func (it *BridgeManagerThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeManagerThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeManagerThresholdUpdated represents a ThresholdUpdated event raised by the BridgeManager contract.
type BridgeManagerThresholdUpdated struct {
	Nonce               *big.Int
	Numerator           *big.Int
	Denominator         *big.Int
	PreviousNumerator   *big.Int
	PreviousDenominator *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterThresholdUpdated is a free log retrieval operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_BridgeManager *BridgeManagerFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*BridgeManagerThresholdUpdatedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _BridgeManager.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &BridgeManagerThresholdUpdatedIterator{contract: _BridgeManager.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_BridgeManager *BridgeManagerFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *BridgeManagerThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var numeratorRule []interface{}
	for _, numeratorItem := range numerator {
		numeratorRule = append(numeratorRule, numeratorItem)
	}
	var denominatorRule []interface{}
	for _, denominatorItem := range denominator {
		denominatorRule = append(denominatorRule, denominatorItem)
	}

	logs, sub, err := _BridgeManager.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeManagerThresholdUpdated)
				if err := _BridgeManager.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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

// ParseThresholdUpdated is a log parse operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_BridgeManager *BridgeManagerFilterer) ParseThresholdUpdated(log types.Log) (*BridgeManagerThresholdUpdated, error) {
	event := new(BridgeManagerThresholdUpdated)
	if err := _BridgeManager.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
