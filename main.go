package bridge_contracts

import (
	ethereumGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var ABIMaps = map[string]*bind.MetaData{
	"RoninGateway":    roninGateway.GatewayMetaData,
	"EthereumGateway": ethereumGateway.GatewayMetaData,
	"RoninValidator":  validator.ValidatorMetaData,
}
