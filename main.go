package bridge_contracts

import (
	ethereumGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	ethereumGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	roninTrustedOrganization "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_organization"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/validator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var ABIMaps = map[string]*bind.MetaData{
	"RoninGateway":             roninGateway.GatewayMetaData,
	"EthereumGateway":          ethereumGateway.GatewayMetaData,
	"RoninValidator":           validator.ValidatorMetaData,
	"RoninGovernanceAdmin":     roninGovernance.GovernanceMetaData,
	"EthereumGovernanceAdmin":  ethereumGovernance.GovernanceMetaData,
	"RoninTrustedOrganization": roninTrustedOrganization.TrustedOrganizationMetaData,
}
