// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trusted_organization

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

// IRoninTrustedOrganizationTrustedOrganization is an auto generated low-level Go binding around an user-defined struct.
type IRoninTrustedOrganizationTrustedOrganization struct {
	ConsensusAddr common.Address
	Governor      common.Address
	BridgeVoter   common.Address
	Weight        *big.Int
	AddedBlock    *big.Int
}

// TrustedOrganizationMetaData contains all meta data concerning the TrustedOrganization contract.
var TrustedOrganizationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDenominator\",\"type\":\"uint256\"}],\"name\":\"ThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"orgs\",\"type\":\"tuple[]\"}],\"name\":\"TrustedOrganizationsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"orgs\",\"type\":\"address[]\"}],\"name\":\"TrustedOrganizationsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"orgs\",\"type\":\"tuple[]\"}],\"name\":\"TrustedOrganizationsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"name\":\"addTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteWeight\",\"type\":\"uint256\"}],\"name\":\"checkThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countTrustedOrganizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTrustedOrganizations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getBridgeVoterWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"getBridgeVoterWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"getConsensusWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"getConsensusWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"getGovernorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"getGovernorWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_res\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusAddr\",\"type\":\"address\"}],\"name\":\"getTrustedOrganization\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getTrustedOrganizationAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_trustedOrgs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"__num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"__denom\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumVoteWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"removeTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"sumBridgeVoterWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"sumConsensusWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_list\",\"type\":\"address[]\"}],\"name\":\"sumGovernorWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_res\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeVoter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structIRoninTrustedOrganization.TrustedOrganization[]\",\"name\":\"_list\",\"type\":\"tuple[]\"}],\"name\":\"updateTrustedOrganizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506123f8806100206000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063b505a07c116100b8578063d9d5dadb1161007c578063d9d5dadb146102d5578063dafae408146102e8578063db6693a21461030b578063e75235b81461032b578063e8c0685e14610336578063f09267c21461034957600080fd5b8063b505a07c14610256578063b9c3620914610269578063cacf8fb514610291578063cc7e6b3b14610299578063d78392f8146102ac57600080fd5b80635f14a1c31161010a5780635f14a1c3146101fa578063708236251461020d5780637c37103c146102205780637de5dedd14610233578063926323d51461023b578063a85c7d6e1461024357600080fd5b80630ed285df14610147578063150740051461015c57806341feed1c1461017a578063520fce62146101b157806356241911146101d1575b600080fd5b61015a610155366004611c5c565b61035c565b005b6101646103ab565b6040516101719190611cdb565b60405180910390f35b6101a3610188366004611d45565b6001600160a01b031660009081526005602052604090205490565b604051908152602001610171565b6101c46101bf366004611d60565b6105a1565b6040516101719190611dd5565b6101a36101df366004611d45565b6001600160a01b031660009081526007602052604090205490565b6101a3610208366004611d60565b610676565b6101c461021b366004611d60565b6106e5565b61015a61022e366004611e0d565b6107b3565b6101a36108dc565b6003546101a3565b61015a610251366004611d60565b610919565b61015a610264366004611c5c565b6109f8565b61027c610277366004611e5e565b610acc565b60408051928352602083019190915201610171565b6009546101a3565b6101c46102a7366004611d60565b610b1d565b6101a36102ba366004611d45565b6001600160a01b031660009081526006602052604090205490565b6101a36102e3366004611d60565b610beb565b6102fb6102f6366004611e80565b610c5a565b6040519015158152602001610171565b61031e610319366004611d45565b610c81565b6040516101719190611e99565b60015460025461027c565b6101a3610344366004611d60565b610d8a565b61031e610357366004611e80565b610df9565b610364610eff565b6001600160a01b0316336001600160a01b03161461039d5760405162461bcd60e51b815260040161039490611ea7565b60405180910390fd5b6103a78282610f2d565b5050565b60095460609067ffffffffffffffff8111156103c9576103c9611ee9565b60405190808252806020026020018201604052801561042257816020015b6040805160a0810182526000808252602080830182905292820181905260608201819052608082015282526000199092019101816103e75790505b5090506000805b825181101561059c576009818154811061044557610445611eff565b9060005260206000200160009054906101000a90046001600160a01b031691508183828151811061047857610478611eff565b60209081029190910101516001600160a01b039091169052600a8054829081106104a4576104a4611eff565b9060005260206000200160009054906101000a90046001600160a01b03168382815181106104d4576104d4611eff565b6020026020010151602001906001600160a01b031690816001600160a01b031681525050600b818154811061050b5761050b611eff565b9060005260206000200160009054906101000a90046001600160a01b031683828151811061053b5761053b611eff565b6020908102919091018101516001600160a01b03928316604091820152918416600090815260059091522054835184908390811061057b5761057b611eff565b6020908102919091010151606001528061059481611f2b565b915050610429565b505090565b60608167ffffffffffffffff8111156105bc576105bc611ee9565b6040519080825280602002602001820160405280156105e5578160200160208202803683370190505b50905060005b815181101561066f576005600085858481811061060a5761060a611eff565b905060200201602081019061061f9190611d45565b6001600160a01b03166001600160a01b031681526020019081526020016000205482828151811061065257610652611eff565b60209081029190910101528061066781611f2b565b9150506105eb565b5092915050565b6000805b8281101561066f576006600085858481811061069857610698611eff565b90506020020160208101906106ad9190611d45565b6001600160a01b031681526020810191909152604001600020546106d19083611f44565b9150806106dd81611f2b565b91505061067a565b60608167ffffffffffffffff81111561070057610700611ee9565b604051908082528060200260200182016040528015610729578160200160208202803683370190505b50905060005b815181101561066f576007600085858481811061074e5761074e611eff565b90506020020160208101906107639190611d45565b6001600160a01b03166001600160a01b031681526020019081526020016000205482828151811061079657610796611eff565b6020908102919091010152806107ab81611f2b565b91505061072f565b600054610100900460ff16158080156107d35750600054600160ff909116105b806107ed5750303b1580156107ed575060005460ff166001145b6108505760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610394565b6000805460ff191660011790558015610873576000805461ff0019166101001790555b8315610883576108838585610f2d565b61088d8383610fac565b505080156108d5576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b600060025460016002546003546001546108f69190611f57565b6109009190611f44565b61090a9190611f6e565b6109149190611f81565b905090565b610921610eff565b6001600160a01b0316336001600160a01b0316146109515760405162461bcd60e51b815260040161039490611ea7565b8061096e5760405162461bcd60e51b815260040161039490611fa3565b60005b818110156109ba576109a883838381811061098e5761098e611eff565b90506020020160208101906109a39190611d45565b61106d565b806109b281611f2b565b915050610971565b507f121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a782826040516109ec929190611fdf565b60405180910390a15050565b610a00610eff565b6001600160a01b0316336001600160a01b031614610a305760405162461bcd60e51b815260040161039490611ea7565b80610a4d5760405162461bcd60e51b815260040161039490611fa3565b60005b81811015610a9a57610a88838383818110610a6d57610a6d611eff565b905060a00201803603810190610a83919061202b565b6113ff565b80610a9281611f2b565b915050610a50565b507fe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda82826040516109ec9291906120bc565b600080610ad7610eff565b6001600160a01b0316336001600160a01b031614610b075760405162461bcd60e51b815260040161039490611ea7565b610b118484610fac565b915091505b9250929050565b60608167ffffffffffffffff811115610b3857610b38611ee9565b604051908082528060200260200182016040528015610b61578160200160208202803683370190505b50905060005b815181101561066f5760066000858584818110610b8657610b86611eff565b9050602002016020810190610b9b9190611d45565b6001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110610bce57610bce611eff565b602090810291909101015280610be381611f2b565b915050610b67565b6000805b8281101561066f5760076000858584818110610c0d57610c0d611eff565b9050602002016020810190610c229190611d45565b6001600160a01b03168152602081019190915260400160002054610c469083611f44565b915080610c5281611f2b565b915050610bef565b6000600354600154610c6c9190611f57565b600254610c799084611f57565b101592915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101829052905b600954811015610d1057826001600160a01b031660098281548110610cd457610cd4611eff565b6000918252602090912001546001600160a01b031603610cfe57610cf781610df9565b9392505050565b80610d0881611f2b565b915050610cad565b5060405162461bcd60e51b815260206004820152604260248201527f526f6e696e547275737465644f7267616e697a6174696f6e3a2071756572792060448201527f666f72206e6f6e2d6578697374656e7420636f6e73656e737573206164647265606482015261737360f01b608482015260a401610394565b6000805b8281101561066f5760056000858584818110610dac57610dac611eff565b9050602002016020810190610dc19190611d45565b6001600160a01b03168152602081019190915260400160002054610de59083611f44565b915080610df181611f2b565b915050610d8e565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152600060098381548110610e3957610e39611eff565b9060005260206000200160009054906101000a90046001600160a01b031690506040518060a00160405280826001600160a01b03168152602001600a8581548110610e8657610e86611eff565b600091825260209182902001546001600160a01b03168252600b8054929091019186908110610eb757610eb7611eff565b60009182526020808320909101546001600160a01b03908116845294909416808252600585526040808320548487015290825260089094528390205492019190915292915050565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103546001600160a01b031690565b60005b81811015610f7a57610f68838383818110610f4d57610f4d611eff565b905060a00201803603810190610f63919061202b565b6117e2565b80610f7281611f2b565b915050610f30565b507fc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c82826040516109ec9291906120bc565b600080828411156110015760405162461bcd60e51b815260206004820152602b60248201526000805160206123a383398151915260448201526a19081d1a1c995cda1bdb1960aa1b6064820152608401610394565b5050600180546002805492859055839055600480549192918491869190600061102983611f2b565b9091555060408051868152602081018690527f976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8910160405180910390a49250929050565b6001600160a01b038116600090815260056020526040812054908190036110d8576110a2826001600160a01b03166014611a72565b6040516020016110b291906121a7565b60408051601f198184030181529082905262461bcd60e51b8252610394916004016121e2565b600954600090815b8181101561113957846001600160a01b03166009828154811061110557611105611eff565b6000918252602090912001546001600160a01b03160361112757809250611139565b8061113181611f2b565b9150506110e0565b50826003600082825461114c9190611f6e565b90915550506001600160a01b038416600090815260086020908152604080832083905560059091528120556009611184600183611f6e565b8154811061119457611194611eff565b600091825260209091200154600980546001600160a01b0390921691849081106111c0576111c0611eff565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060098054806111ff576111ff612215565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560066000600a848154811061123e5761123e611eff565b60009182526020808320909101546001600160a01b03168352820192909252604001812055600a611270600183611f6e565b8154811061128057611280611eff565b600091825260209091200154600a80546001600160a01b0390921691849081106112ac576112ac611eff565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600a8054806112eb576112eb612215565b6001900381819060005260206000200160006101000a8154906001600160a01b030219169055905560076000600b848154811061132a5761132a611eff565b60009182526020808320909101546001600160a01b03168352820192909252604001812055600b61135c600183611f6e565b8154811061136c5761136c611eff565b600091825260209091200154600b80546001600160a01b03909216918490811061139857611398611eff565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600b8054806113d7576113d7612215565b600082815260209020810160001990810180546001600160a01b031916905501905550505050565b60008160600151116114235760405162461bcd60e51b81526004016103949061222b565b80516001600160a01b03166000908152600560205260408120549081900361145b5781516110a2906001600160a01b03166014611a72565b60095460005b818110156117dc5783600001516001600160a01b03166009828154811061148a5761148a611eff565b6000918252602090912001546001600160a01b0316036117ca5782600360008282546114b69190611f6e565b90915550506060840151600380546000906114d2908490611f44565b9250508190555083602001516001600160a01b0316600a82815481106114fa576114fa611eff565b6000918252602090912001546001600160a01b03161461162a576020808501516001600160a01b0316600090815260069091526040902054156115a55760405162461bcd60e51b815260206004820152603760248201527f526f6e696e547275737465644f7267616e697a6174696f6e3a2071756572792060448201527f666f72206475706c69636174656420676f7665726e6f720000000000000000006064820152608401610394565b60066000600a83815481106115bc576115bc611eff565b60009182526020808320909101546001600160a01b03168352828101939093526040909101812055840151600a8054839081106115fb576115fb611eff565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b83604001516001600160a01b0316600b828154811061164b5761164b611eff565b6000918252602090912001546001600160a01b031614611778576040808501516001600160a01b0316600090815260076020522054156116f35760405162461bcd60e51b815260206004820152603b60248201527f526f6e696e547275737465644f7267616e697a6174696f6e3a2071756572792060448201527f666f72206475706c6963617465642062726964676520766f74657200000000006064820152608401610394565b60076000600b838154811061170a5761170a611eff565b60009182526020808320909101546001600160a01b03168352828101939093526040909101812055840151600b80548390811061174957611749611eff565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b50505060608101805182516001600160a01b0390811660009081526005602090815260408083209490945584518187015184168352600682528483205593519483015190911681526007909252902055565b806117d481611f2b565b915050611461565b50505050565b6080810151156118345760405162461bcd60e51b815260206004820152602960248201526000805160206123a383398151915260448201526819081c995c5d595cdd60ba1b6064820152608401610394565b60008160600151116118585760405162461bcd60e51b81526004016103949061222b565b80516001600160a01b03166000908152600560205260409020541561189d57805161188d906001600160a01b03166014611a72565b6040516020016110b29190612261565b6020808201516001600160a01b0316600090815260069091526040902054156118e8576118d881602001516001600160a01b03166014611a72565b6040516020016110b291906122a0565b6040808201516001600160a01b0316600090815260076020522054156119305761192081604001516001600160a01b03166014611a72565b6040516020016110b29190612313565b8051600980546001808201835560009283527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af90910180546001600160a01b03199081166001600160a01b0395861617909155606085018051865186168552600560209081526040808720929092558088018051600a805480890182559089527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a80180548716918a16919091179055835190518816875260068252828720558188018051600b8054978801815588527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9909601805490951695881695909517909355815193518616855260078352808520939093558551909416835260089052812043905590516003805491929091611a6a908490611f44565b909155505050565b60606000611a81836002611f57565b611a8c906002611f44565b67ffffffffffffffff811115611aa457611aa4611ee9565b6040519080825280601f01601f191660200182016040528015611ace576020820181803683370190505b509050600360fc1b81600081518110611ae957611ae9611eff565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611b1857611b18611eff565b60200101906001600160f81b031916908160001a9053506000611b3c846002611f57565b611b47906001611f44565b90505b6001811115611bbf576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611b7b57611b7b611eff565b1a60f81b828281518110611b9157611b91611eff565b60200101906001600160f81b031916908160001a90535060049490941c93611bb88161238b565b9050611b4a565b508315611c0e5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610394565b90505b92915050565b60008083601f840112611c2957600080fd5b50813567ffffffffffffffff811115611c4157600080fd5b60208301915083602060a083028501011115610b1657600080fd5b60008060208385031215611c6f57600080fd5b823567ffffffffffffffff811115611c8657600080fd5b611c9285828601611c17565b90969095509350505050565b80516001600160a01b0390811683526020808301518216908401526040808301519091169083015260608082015190830152608090810151910152565b6020808252825182820181905260009190848201906040850190845b81811015611d1d57611d0a838551611c9e565b9284019260a09290920191600101611cf7565b50909695505050505050565b80356001600160a01b0381168114611d4057600080fd5b919050565b600060208284031215611d5757600080fd5b610cf782611d29565b60008060208385031215611d7357600080fd5b823567ffffffffffffffff80821115611d8b57600080fd5b818501915085601f830112611d9f57600080fd5b813581811115611dae57600080fd5b8660208260051b8501011115611dc357600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b81811015611d1d57835183529284019291840191600101611df1565b60008060008060608587031215611e2357600080fd5b843567ffffffffffffffff811115611e3a57600080fd5b611e4687828801611c17565b90989097506020870135966040013595509350505050565b60008060408385031215611e7157600080fd5b50508035926020909101359150565b600060208284031215611e9257600080fd5b5035919050565b60a08101611c118284611c9e565b60208082526022908201527f48617350726f787941646d696e3a20756e617574686f72697a65642073656e6460408201526132b960f11b606082015260800190565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611f3d57611f3d611f15565b5060010190565b80820180821115611c1157611c11611f15565b8082028115828204841417611c1157611c11611f15565b81810381811115611c1157611c11611f15565b600082611f9e57634e487b7160e01b600052601260045260246000fd5b500490565b6020808252602e908201526000805160206123a383398151915260408201526d0c840c2e4e4c2f240d8cadccee8d60931b606082015260800190565b60208082528181018390526000908460408401835b86811015612020576001600160a01b0361200d84611d29565b1682529183019190830190600101611ff4565b509695505050505050565b600060a0828403121561203d57600080fd5b60405160a0810181811067ffffffffffffffff8211171561206e57634e487b7160e01b600052604160045260246000fd5b60405261207a83611d29565b815261208860208401611d29565b602082015261209960408401611d29565b604082015260608301356060820152608083013560808201528091505092915050565b6020808252818101839052600090604080840186845b87811015612138576001600160a01b03806120ec84611d29565b168452806120fb878501611d29565b16868501528061210c868501611d29565b168486015250606082810135908401526080808301359084015260a092830192909101906001016120d2565b5090979650505050505050565b7f526f6e696e547275737465644f7267616e697a6174696f6e3a20636f6e73656e81526b039bab99030b2323932b9b9960a51b6020820152602c0190565b60005b8381101561219e578181015183820152602001612186565b50506000910152565b60006121b282612145565b83516121c2818360208801612183565b6c081a5cc81b9bdd081859191959609a1b9101908152600d019392505050565b6020815260008251806020840152612201816040850160208701612183565b601f01601f19169190910160400192915050565b634e487b7160e01b600052603160045260246000fd5b60208082526028908201526000805160206123a383398151915260408201526719081dd95a59da1d60c21b606082015260800190565b600061226c82612145565b835161227c818360208801612183565b7020697320616464656420616c726561647960781b91019081526011019392505050565b7f526f6e696e547275737465644f7267616e697a6174696f6e3a20676f76656e6f8152690391030b2323932b9b9960b51b6020820152600082516122eb81602a850160208701612183565b7020697320616464656420616c726561647960781b602a939091019283015250603b01919050565b7f526f6e696e547275737465644f7267616e697a6174696f6e3a2062726964676581526e0103b37ba32b91030b2323932b9b99608d1b60208201526000825161236381602f850160208701612183565b7020697320616464656420616c726561647960781b602f939091019283015250604001919050565b60008161239a5761239a611f15565b50600019019056fe526f6e696e547275737465644f7267616e697a6174696f6e3a20696e76616c69a2646970667358221220563f3cdfada3d382a0a435ace5a52ca372c22c02e576490ee81743d52ccecd0d64736f6c63430008110033",
}

// TrustedOrganizationABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustedOrganizationMetaData.ABI instead.
var TrustedOrganizationABI = TrustedOrganizationMetaData.ABI

// TrustedOrganizationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrustedOrganizationMetaData.Bin instead.
var TrustedOrganizationBin = TrustedOrganizationMetaData.Bin

// DeployTrustedOrganization deploys a new Ethereum contract, binding an instance of TrustedOrganization to it.
func DeployTrustedOrganization(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrustedOrganization, error) {
	parsed, err := TrustedOrganizationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrustedOrganizationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrustedOrganization{TrustedOrganizationCaller: TrustedOrganizationCaller{contract: contract}, TrustedOrganizationTransactor: TrustedOrganizationTransactor{contract: contract}, TrustedOrganizationFilterer: TrustedOrganizationFilterer{contract: contract}}, nil
}

// TrustedOrganization is an auto generated Go binding around an Ethereum contract.
type TrustedOrganization struct {
	TrustedOrganizationCaller     // Read-only binding to the contract
	TrustedOrganizationTransactor // Write-only binding to the contract
	TrustedOrganizationFilterer   // Log filterer for contract events
}

// TrustedOrganizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustedOrganizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedOrganizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustedOrganizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedOrganizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustedOrganizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedOrganizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustedOrganizationSession struct {
	Contract     *TrustedOrganization // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TrustedOrganizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustedOrganizationCallerSession struct {
	Contract *TrustedOrganizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TrustedOrganizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustedOrganizationTransactorSession struct {
	Contract     *TrustedOrganizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TrustedOrganizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustedOrganizationRaw struct {
	Contract *TrustedOrganization // Generic contract binding to access the raw methods on
}

// TrustedOrganizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustedOrganizationCallerRaw struct {
	Contract *TrustedOrganizationCaller // Generic read-only contract binding to access the raw methods on
}

// TrustedOrganizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustedOrganizationTransactorRaw struct {
	Contract *TrustedOrganizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustedOrganization creates a new instance of TrustedOrganization, bound to a specific deployed contract.
func NewTrustedOrganization(address common.Address, backend bind.ContractBackend) (*TrustedOrganization, error) {
	contract, err := bindTrustedOrganization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustedOrganization{TrustedOrganizationCaller: TrustedOrganizationCaller{contract: contract}, TrustedOrganizationTransactor: TrustedOrganizationTransactor{contract: contract}, TrustedOrganizationFilterer: TrustedOrganizationFilterer{contract: contract}}, nil
}

// NewTrustedOrganizationCaller creates a new read-only instance of TrustedOrganization, bound to a specific deployed contract.
func NewTrustedOrganizationCaller(address common.Address, caller bind.ContractCaller) (*TrustedOrganizationCaller, error) {
	contract, err := bindTrustedOrganization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationCaller{contract: contract}, nil
}

// NewTrustedOrganizationTransactor creates a new write-only instance of TrustedOrganization, bound to a specific deployed contract.
func NewTrustedOrganizationTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustedOrganizationTransactor, error) {
	contract, err := bindTrustedOrganization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationTransactor{contract: contract}, nil
}

// NewTrustedOrganizationFilterer creates a new log filterer instance of TrustedOrganization, bound to a specific deployed contract.
func NewTrustedOrganizationFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustedOrganizationFilterer, error) {
	contract, err := bindTrustedOrganization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationFilterer{contract: contract}, nil
}

// bindTrustedOrganization binds a generic wrapper to an already deployed contract.
func bindTrustedOrganization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrustedOrganizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedOrganization *TrustedOrganizationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustedOrganization.Contract.TrustedOrganizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedOrganization *TrustedOrganizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.TrustedOrganizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedOrganization *TrustedOrganizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.TrustedOrganizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedOrganization *TrustedOrganizationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustedOrganization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedOrganization *TrustedOrganizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedOrganization *TrustedOrganizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.contract.Transact(opts, method, params...)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_TrustedOrganization *TrustedOrganizationCaller) CheckThreshold(opts *bind.CallOpts, _voteWeight *big.Int) (bool, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "checkThreshold", _voteWeight)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_TrustedOrganization *TrustedOrganizationSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _TrustedOrganization.Contract.CheckThreshold(&_TrustedOrganization.CallOpts, _voteWeight)
}

// CheckThreshold is a free data retrieval call binding the contract method 0xdafae408.
//
// Solidity: function checkThreshold(uint256 _voteWeight) view returns(bool)
func (_TrustedOrganization *TrustedOrganizationCallerSession) CheckThreshold(_voteWeight *big.Int) (bool, error) {
	return _TrustedOrganization.Contract.CheckThreshold(&_TrustedOrganization.CallOpts, _voteWeight)
}

// CountTrustedOrganizations is a free data retrieval call binding the contract method 0xcacf8fb5.
//
// Solidity: function countTrustedOrganizations() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) CountTrustedOrganizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "countTrustedOrganizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountTrustedOrganizations is a free data retrieval call binding the contract method 0xcacf8fb5.
//
// Solidity: function countTrustedOrganizations() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) CountTrustedOrganizations() (*big.Int, error) {
	return _TrustedOrganization.Contract.CountTrustedOrganizations(&_TrustedOrganization.CallOpts)
}

// CountTrustedOrganizations is a free data retrieval call binding the contract method 0xcacf8fb5.
//
// Solidity: function countTrustedOrganizations() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) CountTrustedOrganizations() (*big.Int, error) {
	return _TrustedOrganization.Contract.CountTrustedOrganizations(&_TrustedOrganization.CallOpts)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] _list)
func (_TrustedOrganization *TrustedOrganizationCaller) GetAllTrustedOrganizations(opts *bind.CallOpts) ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getAllTrustedOrganizations")

	if err != nil {
		return *new([]IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRoninTrustedOrganizationTrustedOrganization)).(*[]IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] _list)
func (_TrustedOrganization *TrustedOrganizationSession) GetAllTrustedOrganizations() ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetAllTrustedOrganizations(&_TrustedOrganization.CallOpts)
}

// GetAllTrustedOrganizations is a free data retrieval call binding the contract method 0x15074005.
//
// Solidity: function getAllTrustedOrganizations() view returns((address,address,address,uint256,uint256)[] _list)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetAllTrustedOrganizations() ([]IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetAllTrustedOrganizations(&_TrustedOrganization.CallOpts)
}

// GetBridgeVoterWeight is a free data retrieval call binding the contract method 0x56241911.
//
// Solidity: function getBridgeVoterWeight(address _addr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) GetBridgeVoterWeight(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getBridgeVoterWeight", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBridgeVoterWeight is a free data retrieval call binding the contract method 0x56241911.
//
// Solidity: function getBridgeVoterWeight(address _addr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) GetBridgeVoterWeight(_addr common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetBridgeVoterWeight(&_TrustedOrganization.CallOpts, _addr)
}

// GetBridgeVoterWeight is a free data retrieval call binding the contract method 0x56241911.
//
// Solidity: function getBridgeVoterWeight(address _addr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetBridgeVoterWeight(_addr common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetBridgeVoterWeight(&_TrustedOrganization.CallOpts, _addr)
}

// GetBridgeVoterWeights is a free data retrieval call binding the contract method 0x70823625.
//
// Solidity: function getBridgeVoterWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCaller) GetBridgeVoterWeights(opts *bind.CallOpts, _list []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getBridgeVoterWeights", _list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBridgeVoterWeights is a free data retrieval call binding the contract method 0x70823625.
//
// Solidity: function getBridgeVoterWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationSession) GetBridgeVoterWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetBridgeVoterWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetBridgeVoterWeights is a free data retrieval call binding the contract method 0x70823625.
//
// Solidity: function getBridgeVoterWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetBridgeVoterWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetBridgeVoterWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address _consensusAddr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) GetConsensusWeight(opts *bind.CallOpts, _consensusAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getConsensusWeight", _consensusAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address _consensusAddr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) GetConsensusWeight(_consensusAddr common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetConsensusWeight(&_TrustedOrganization.CallOpts, _consensusAddr)
}

// GetConsensusWeight is a free data retrieval call binding the contract method 0x41feed1c.
//
// Solidity: function getConsensusWeight(address _consensusAddr) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetConsensusWeight(_consensusAddr common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetConsensusWeight(&_TrustedOrganization.CallOpts, _consensusAddr)
}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCaller) GetConsensusWeights(opts *bind.CallOpts, _list []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getConsensusWeights", _list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationSession) GetConsensusWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetConsensusWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetConsensusWeights is a free data retrieval call binding the contract method 0x520fce62.
//
// Solidity: function getConsensusWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetConsensusWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetConsensusWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) GetGovernorWeight(opts *bind.CallOpts, _governor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getGovernorWeight", _governor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) GetGovernorWeight(_governor common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetGovernorWeight(&_TrustedOrganization.CallOpts, _governor)
}

// GetGovernorWeight is a free data retrieval call binding the contract method 0xd78392f8.
//
// Solidity: function getGovernorWeight(address _governor) view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetGovernorWeight(_governor common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.GetGovernorWeight(&_TrustedOrganization.CallOpts, _governor)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCaller) GetGovernorWeights(opts *bind.CallOpts, _list []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getGovernorWeights", _list)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationSession) GetGovernorWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetGovernorWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetGovernorWeights is a free data retrieval call binding the contract method 0xcc7e6b3b.
//
// Solidity: function getGovernorWeights(address[] _list) view returns(uint256[] _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetGovernorWeights(_list []common.Address) ([]*big.Int, error) {
	return _TrustedOrganization.Contract.GetGovernorWeights(&_TrustedOrganization.CallOpts, _list)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _TrustedOrganization.Contract.GetThreshold(&_TrustedOrganization.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetThreshold() (*big.Int, *big.Int, error) {
	return _TrustedOrganization.Contract.GetThreshold(&_TrustedOrganization.CallOpts)
}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationCaller) GetTrustedOrganization(opts *bind.CallOpts, _consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getTrustedOrganization", _consensusAddr)

	if err != nil {
		return *new(IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new(IRoninTrustedOrganizationTrustedOrganization)).(*IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationSession) GetTrustedOrganization(_consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetTrustedOrganization(&_TrustedOrganization.CallOpts, _consensusAddr)
}

// GetTrustedOrganization is a free data retrieval call binding the contract method 0xdb6693a2.
//
// Solidity: function getTrustedOrganization(address _consensusAddr) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetTrustedOrganization(_consensusAddr common.Address) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetTrustedOrganization(&_TrustedOrganization.CallOpts, _consensusAddr)
}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationCaller) GetTrustedOrganizationAt(opts *bind.CallOpts, _idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "getTrustedOrganizationAt", _idx)

	if err != nil {
		return *new(IRoninTrustedOrganizationTrustedOrganization), err
	}

	out0 := *abi.ConvertType(out[0], new(IRoninTrustedOrganizationTrustedOrganization)).(*IRoninTrustedOrganizationTrustedOrganization)

	return out0, err

}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationSession) GetTrustedOrganizationAt(_idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetTrustedOrganizationAt(&_TrustedOrganization.CallOpts, _idx)
}

// GetTrustedOrganizationAt is a free data retrieval call binding the contract method 0xf09267c2.
//
// Solidity: function getTrustedOrganizationAt(uint256 _idx) view returns((address,address,address,uint256,uint256))
func (_TrustedOrganization *TrustedOrganizationCallerSession) GetTrustedOrganizationAt(_idx *big.Int) (IRoninTrustedOrganizationTrustedOrganization, error) {
	return _TrustedOrganization.Contract.GetTrustedOrganizationAt(&_TrustedOrganization.CallOpts, _idx)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) MinimumVoteWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "minimumVoteWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) MinimumVoteWeight() (*big.Int, error) {
	return _TrustedOrganization.Contract.MinimumVoteWeight(&_TrustedOrganization.CallOpts)
}

// MinimumVoteWeight is a free data retrieval call binding the contract method 0x7de5dedd.
//
// Solidity: function minimumVoteWeight() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) MinimumVoteWeight() (*big.Int, error) {
	return _TrustedOrganization.Contract.MinimumVoteWeight(&_TrustedOrganization.CallOpts)
}

// SumBridgeVoterWeights is a free data retrieval call binding the contract method 0xd9d5dadb.
//
// Solidity: function sumBridgeVoterWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCaller) SumBridgeVoterWeights(opts *bind.CallOpts, _list []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "sumBridgeVoterWeights", _list)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumBridgeVoterWeights is a free data retrieval call binding the contract method 0xd9d5dadb.
//
// Solidity: function sumBridgeVoterWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationSession) SumBridgeVoterWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumBridgeVoterWeights(&_TrustedOrganization.CallOpts, _list)
}

// SumBridgeVoterWeights is a free data retrieval call binding the contract method 0xd9d5dadb.
//
// Solidity: function sumBridgeVoterWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) SumBridgeVoterWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumBridgeVoterWeights(&_TrustedOrganization.CallOpts, _list)
}

// SumConsensusWeights is a free data retrieval call binding the contract method 0xe8c0685e.
//
// Solidity: function sumConsensusWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCaller) SumConsensusWeights(opts *bind.CallOpts, _list []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "sumConsensusWeights", _list)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumConsensusWeights is a free data retrieval call binding the contract method 0xe8c0685e.
//
// Solidity: function sumConsensusWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationSession) SumConsensusWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumConsensusWeights(&_TrustedOrganization.CallOpts, _list)
}

// SumConsensusWeights is a free data retrieval call binding the contract method 0xe8c0685e.
//
// Solidity: function sumConsensusWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) SumConsensusWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumConsensusWeights(&_TrustedOrganization.CallOpts, _list)
}

// SumGovernorWeights is a free data retrieval call binding the contract method 0x5f14a1c3.
//
// Solidity: function sumGovernorWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCaller) SumGovernorWeights(opts *bind.CallOpts, _list []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "sumGovernorWeights", _list)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SumGovernorWeights is a free data retrieval call binding the contract method 0x5f14a1c3.
//
// Solidity: function sumGovernorWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationSession) SumGovernorWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumGovernorWeights(&_TrustedOrganization.CallOpts, _list)
}

// SumGovernorWeights is a free data retrieval call binding the contract method 0x5f14a1c3.
//
// Solidity: function sumGovernorWeights(address[] _list) view returns(uint256 _res)
func (_TrustedOrganization *TrustedOrganizationCallerSession) SumGovernorWeights(_list []common.Address) (*big.Int, error) {
	return _TrustedOrganization.Contract.SumGovernorWeights(&_TrustedOrganization.CallOpts, _list)
}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCaller) TotalWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustedOrganization.contract.Call(opts, &out, "totalWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationSession) TotalWeights() (*big.Int, error) {
	return _TrustedOrganization.Contract.TotalWeights(&_TrustedOrganization.CallOpts)
}

// TotalWeights is a free data retrieval call binding the contract method 0x926323d5.
//
// Solidity: function totalWeights() view returns(uint256)
func (_TrustedOrganization *TrustedOrganizationCallerSession) TotalWeights() (*big.Int, error) {
	return _TrustedOrganization.Contract.TotalWeights(&_TrustedOrganization.CallOpts)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactor) AddTrustedOrganizations(opts *bind.TransactOpts, _list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.contract.Transact(opts, "addTrustedOrganizations", _list)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationSession) AddTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.AddTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// AddTrustedOrganizations is a paid mutator transaction binding the contract method 0x0ed285df.
//
// Solidity: function addTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactorSession) AddTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.AddTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] _trustedOrgs, uint256 __num, uint256 __denom) returns()
func (_TrustedOrganization *TrustedOrganizationTransactor) Initialize(opts *bind.TransactOpts, _trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, __num *big.Int, __denom *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.contract.Transact(opts, "initialize", _trustedOrgs, __num, __denom)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] _trustedOrgs, uint256 __num, uint256 __denom) returns()
func (_TrustedOrganization *TrustedOrganizationSession) Initialize(_trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, __num *big.Int, __denom *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.Initialize(&_TrustedOrganization.TransactOpts, _trustedOrgs, __num, __denom)
}

// Initialize is a paid mutator transaction binding the contract method 0x7c37103c.
//
// Solidity: function initialize((address,address,address,uint256,uint256)[] _trustedOrgs, uint256 __num, uint256 __denom) returns()
func (_TrustedOrganization *TrustedOrganizationTransactorSession) Initialize(_trustedOrgs []IRoninTrustedOrganizationTrustedOrganization, __num *big.Int, __denom *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.Initialize(&_TrustedOrganization.TransactOpts, _trustedOrgs, __num, __denom)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactor) RemoveTrustedOrganizations(opts *bind.TransactOpts, _list []common.Address) (*types.Transaction, error) {
	return _TrustedOrganization.contract.Transact(opts, "removeTrustedOrganizations", _list)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationSession) RemoveTrustedOrganizations(_list []common.Address) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.RemoveTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// RemoveTrustedOrganizations is a paid mutator transaction binding the contract method 0xa85c7d6e.
//
// Solidity: function removeTrustedOrganizations(address[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactorSession) RemoveTrustedOrganizations(_list []common.Address) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.RemoveTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationTransactor) SetThreshold(opts *bind.TransactOpts, _numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.contract.Transact(opts, "setThreshold", _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.SetThreshold(&_TrustedOrganization.TransactOpts, _numerator, _denominator)
}

// SetThreshold is a paid mutator transaction binding the contract method 0xb9c36209.
//
// Solidity: function setThreshold(uint256 _numerator, uint256 _denominator) returns(uint256, uint256)
func (_TrustedOrganization *TrustedOrganizationTransactorSession) SetThreshold(_numerator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.SetThreshold(&_TrustedOrganization.TransactOpts, _numerator, _denominator)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactor) UpdateTrustedOrganizations(opts *bind.TransactOpts, _list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.contract.Transact(opts, "updateTrustedOrganizations", _list)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationSession) UpdateTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.UpdateTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// UpdateTrustedOrganizations is a paid mutator transaction binding the contract method 0xb505a07c.
//
// Solidity: function updateTrustedOrganizations((address,address,address,uint256,uint256)[] _list) returns()
func (_TrustedOrganization *TrustedOrganizationTransactorSession) UpdateTrustedOrganizations(_list []IRoninTrustedOrganizationTrustedOrganization) (*types.Transaction, error) {
	return _TrustedOrganization.Contract.UpdateTrustedOrganizations(&_TrustedOrganization.TransactOpts, _list)
}

// TrustedOrganizationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TrustedOrganization contract.
type TrustedOrganizationInitializedIterator struct {
	Event *TrustedOrganizationInitialized // Event containing the contract specifics and raw log

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
func (it *TrustedOrganizationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedOrganizationInitialized)
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
		it.Event = new(TrustedOrganizationInitialized)
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
func (it *TrustedOrganizationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedOrganizationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedOrganizationInitialized represents a Initialized event raised by the TrustedOrganization contract.
type TrustedOrganizationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TrustedOrganization *TrustedOrganizationFilterer) FilterInitialized(opts *bind.FilterOpts) (*TrustedOrganizationInitializedIterator, error) {

	logs, sub, err := _TrustedOrganization.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationInitializedIterator{contract: _TrustedOrganization.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TrustedOrganization *TrustedOrganizationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TrustedOrganizationInitialized) (event.Subscription, error) {

	logs, sub, err := _TrustedOrganization.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedOrganizationInitialized)
				if err := _TrustedOrganization.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TrustedOrganization *TrustedOrganizationFilterer) ParseInitialized(log types.Log) (*TrustedOrganizationInitialized, error) {
	event := new(TrustedOrganizationInitialized)
	if err := _TrustedOrganization.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedOrganizationThresholdUpdatedIterator is returned from FilterThresholdUpdated and is used to iterate over the raw logs and unpacked data for ThresholdUpdated events raised by the TrustedOrganization contract.
type TrustedOrganizationThresholdUpdatedIterator struct {
	Event *TrustedOrganizationThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *TrustedOrganizationThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedOrganizationThresholdUpdated)
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
		it.Event = new(TrustedOrganizationThresholdUpdated)
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
func (it *TrustedOrganizationThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedOrganizationThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedOrganizationThresholdUpdated represents a ThresholdUpdated event raised by the TrustedOrganization contract.
type TrustedOrganizationThresholdUpdated struct {
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
func (_TrustedOrganization *TrustedOrganizationFilterer) FilterThresholdUpdated(opts *bind.FilterOpts, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (*TrustedOrganizationThresholdUpdatedIterator, error) {

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

	logs, sub, err := _TrustedOrganization.contract.FilterLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationThresholdUpdatedIterator{contract: _TrustedOrganization.contract, event: "ThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdUpdated is a free log subscription operation binding the contract event 0x976f8a9c5bdf8248dec172376d6e2b80a8e3df2f0328e381c6db8e1cf138c0f8.
//
// Solidity: event ThresholdUpdated(uint256 indexed nonce, uint256 indexed numerator, uint256 indexed denominator, uint256 previousNumerator, uint256 previousDenominator)
func (_TrustedOrganization *TrustedOrganizationFilterer) WatchThresholdUpdated(opts *bind.WatchOpts, sink chan<- *TrustedOrganizationThresholdUpdated, nonce []*big.Int, numerator []*big.Int, denominator []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TrustedOrganization.contract.WatchLogs(opts, "ThresholdUpdated", nonceRule, numeratorRule, denominatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedOrganizationThresholdUpdated)
				if err := _TrustedOrganization.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
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
func (_TrustedOrganization *TrustedOrganizationFilterer) ParseThresholdUpdated(log types.Log) (*TrustedOrganizationThresholdUpdated, error) {
	event := new(TrustedOrganizationThresholdUpdated)
	if err := _TrustedOrganization.contract.UnpackLog(event, "ThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedOrganizationTrustedOrganizationsAddedIterator is returned from FilterTrustedOrganizationsAdded and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsAdded events raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsAddedIterator struct {
	Event *TrustedOrganizationTrustedOrganizationsAdded // Event containing the contract specifics and raw log

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
func (it *TrustedOrganizationTrustedOrganizationsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedOrganizationTrustedOrganizationsAdded)
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
		it.Event = new(TrustedOrganizationTrustedOrganizationsAdded)
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
func (it *TrustedOrganizationTrustedOrganizationsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedOrganizationTrustedOrganizationsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedOrganizationTrustedOrganizationsAdded represents a TrustedOrganizationsAdded event raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsAdded struct {
	Orgs []IRoninTrustedOrganizationTrustedOrganization
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsAdded is a free log retrieval operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) FilterTrustedOrganizationsAdded(opts *bind.FilterOpts) (*TrustedOrganizationTrustedOrganizationsAddedIterator, error) {

	logs, sub, err := _TrustedOrganization.contract.FilterLogs(opts, "TrustedOrganizationsAdded")
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationTrustedOrganizationsAddedIterator{contract: _TrustedOrganization.contract, event: "TrustedOrganizationsAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsAdded is a free log subscription operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) WatchTrustedOrganizationsAdded(opts *bind.WatchOpts, sink chan<- *TrustedOrganizationTrustedOrganizationsAdded) (event.Subscription, error) {

	logs, sub, err := _TrustedOrganization.contract.WatchLogs(opts, "TrustedOrganizationsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedOrganizationTrustedOrganizationsAdded)
				if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsAdded", log); err != nil {
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

// ParseTrustedOrganizationsAdded is a log parse operation binding the contract event 0xc753dbf7952c70ff6b9fa7b626403aa1d2230d97136b635bd5e85bec72bcca6c.
//
// Solidity: event TrustedOrganizationsAdded((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) ParseTrustedOrganizationsAdded(log types.Log) (*TrustedOrganizationTrustedOrganizationsAdded, error) {
	event := new(TrustedOrganizationTrustedOrganizationsAdded)
	if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedOrganizationTrustedOrganizationsRemovedIterator is returned from FilterTrustedOrganizationsRemoved and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsRemoved events raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsRemovedIterator struct {
	Event *TrustedOrganizationTrustedOrganizationsRemoved // Event containing the contract specifics and raw log

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
func (it *TrustedOrganizationTrustedOrganizationsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedOrganizationTrustedOrganizationsRemoved)
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
		it.Event = new(TrustedOrganizationTrustedOrganizationsRemoved)
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
func (it *TrustedOrganizationTrustedOrganizationsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedOrganizationTrustedOrganizationsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedOrganizationTrustedOrganizationsRemoved represents a TrustedOrganizationsRemoved event raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsRemoved struct {
	Orgs []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsRemoved is a free log retrieval operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) FilterTrustedOrganizationsRemoved(opts *bind.FilterOpts) (*TrustedOrganizationTrustedOrganizationsRemovedIterator, error) {

	logs, sub, err := _TrustedOrganization.contract.FilterLogs(opts, "TrustedOrganizationsRemoved")
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationTrustedOrganizationsRemovedIterator{contract: _TrustedOrganization.contract, event: "TrustedOrganizationsRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsRemoved is a free log subscription operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) WatchTrustedOrganizationsRemoved(opts *bind.WatchOpts, sink chan<- *TrustedOrganizationTrustedOrganizationsRemoved) (event.Subscription, error) {

	logs, sub, err := _TrustedOrganization.contract.WatchLogs(opts, "TrustedOrganizationsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedOrganizationTrustedOrganizationsRemoved)
				if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsRemoved", log); err != nil {
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

// ParseTrustedOrganizationsRemoved is a log parse operation binding the contract event 0x121945697ac30ee0fc67821492cb685c65f0ea4d7f1b710fde44d6e2237f43a7.
//
// Solidity: event TrustedOrganizationsRemoved(address[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) ParseTrustedOrganizationsRemoved(log types.Log) (*TrustedOrganizationTrustedOrganizationsRemoved, error) {
	event := new(TrustedOrganizationTrustedOrganizationsRemoved)
	if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedOrganizationTrustedOrganizationsUpdatedIterator is returned from FilterTrustedOrganizationsUpdated and is used to iterate over the raw logs and unpacked data for TrustedOrganizationsUpdated events raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsUpdatedIterator struct {
	Event *TrustedOrganizationTrustedOrganizationsUpdated // Event containing the contract specifics and raw log

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
func (it *TrustedOrganizationTrustedOrganizationsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedOrganizationTrustedOrganizationsUpdated)
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
		it.Event = new(TrustedOrganizationTrustedOrganizationsUpdated)
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
func (it *TrustedOrganizationTrustedOrganizationsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedOrganizationTrustedOrganizationsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedOrganizationTrustedOrganizationsUpdated represents a TrustedOrganizationsUpdated event raised by the TrustedOrganization contract.
type TrustedOrganizationTrustedOrganizationsUpdated struct {
	Orgs []IRoninTrustedOrganizationTrustedOrganization
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTrustedOrganizationsUpdated is a free log retrieval operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) FilterTrustedOrganizationsUpdated(opts *bind.FilterOpts) (*TrustedOrganizationTrustedOrganizationsUpdatedIterator, error) {

	logs, sub, err := _TrustedOrganization.contract.FilterLogs(opts, "TrustedOrganizationsUpdated")
	if err != nil {
		return nil, err
	}
	return &TrustedOrganizationTrustedOrganizationsUpdatedIterator{contract: _TrustedOrganization.contract, event: "TrustedOrganizationsUpdated", logs: logs, sub: sub}, nil
}

// WatchTrustedOrganizationsUpdated is a free log subscription operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) WatchTrustedOrganizationsUpdated(opts *bind.WatchOpts, sink chan<- *TrustedOrganizationTrustedOrganizationsUpdated) (event.Subscription, error) {

	logs, sub, err := _TrustedOrganization.contract.WatchLogs(opts, "TrustedOrganizationsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedOrganizationTrustedOrganizationsUpdated)
				if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsUpdated", log); err != nil {
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

// ParseTrustedOrganizationsUpdated is a log parse operation binding the contract event 0xe887c8106c09d1770c0ef0bf8ca62c54766f18b07506801865501783376cbeda.
//
// Solidity: event TrustedOrganizationsUpdated((address,address,address,uint256,uint256)[] orgs)
func (_TrustedOrganization *TrustedOrganizationFilterer) ParseTrustedOrganizationsUpdated(log types.Log) (*TrustedOrganizationTrustedOrganizationsUpdated, error) {
	event := new(TrustedOrganizationTrustedOrganizationsUpdated)
	if err := _TrustedOrganization.contract.UnpackLog(event, "TrustedOrganizationsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
