package bridge_contracts

import (
	ethereumGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	ethereumGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	bridgeTracking "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/bridge_tracking"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/katana"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/maintenance"
	ronStaking "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/ron_staking"
	slashIndicator "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/slash_indicator"
	stakingVesting "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/staking_vesting"
	roninTrustedOrganization "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_organization"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/validator"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/wron"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var ABIMaps = map[string]*bind.MetaData{
	"RoninGateway":             roninGateway.GatewayMetaData,
	"EthereumGateway":          ethereumGateway.GatewayMetaData,
	"RoninValidator":           validator.ValidatorMetaData,
	"RoninGovernanceAdmin":     roninGovernance.GovernanceMetaData,
	"EthereumGovernanceAdmin":  ethereumGovernance.GovernanceMetaData,
	"RoninTrustedOrganization": roninTrustedOrganization.TrustedOrganizationMetaData,

	"StakingVesting": stakingVesting.StakingVestingMetaData,
	"Maintenance":    maintenance.MaintenanceMetaData,
	"SlashIndicator": slashIndicator.SlashIndicatorMetaData,
	"BridgeTracking": bridgeTracking.BridgeTrackingMetaData,
	"Validator":      validator.ValidatorMetaData,
	"Staking":        ronStaking.RonStakingMetaData,
	"KatanaPair":     katana.KatanaPairMetaData,
	"KatanaRouter":   katana.KatanaRouterMetaData,
	"KatanaPairV3":   katana.KatanaPairV3,
	"KatanaRouterV3": katana.KatanaRouterV3,
	"WRON":           wron.WRONMetaData,
}
