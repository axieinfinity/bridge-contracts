package bridge_contracts

import (
	ethereumGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	ethereumGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/katana"
	bridgeTracking "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/bridge_tracking"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	"github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/maintenance"
	ronStaking "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/ron_staking"
	slashIndicator "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/slash_indicator"
	stakingVesting "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/staking_vesting"
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

	"StakingVesting": stakingVesting.StakingVestingMetaData,
	"Maintenance":    maintenance.MaintenanceMetaData,
	"SlashIndicator": slashIndicator.SlashIndicatorMetaData,
	"BridgeTracking": bridgeTracking.BridgeTrackingMetaData,
	"Validator":      validator.ValidatorMetaData,
	"Staking":        ronStaking.RonStakingMetaData,
	"Katana":         katana.KatanaMetaData,
}
