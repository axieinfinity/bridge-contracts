contract_path = "../ronin-dpos-contracts"
generate:
	solc --abi --bin "${contract_path}/contracts/ronin/RoninGovernanceAdmin.sol" -o "./contracts/ronin" --include-path "${contract_path}/node_modules/" --base-path ${contract_path} --overwrite --optimize
	abigen --abi "./contracts/ronin/RoninGovernanceAdmin.abi" --bin "./contracts/ronin/RoninGovernanceAdmin.bin" --pkg "governance" --out "./generated_contracts/ronin/governance/ronin_governance_admin.go"

	solc --abi --bin "${contract_path}/contracts/mainchain/MainchainGovernanceAdmin.sol" -o "./contracts/ethereum" --include-path "${contract_path}/node_modules/" --base-path ${contract_path} --overwrite --optimize
	abigen --abi "./contracts/ethereum/MainchainGovernanceAdmin.abi" --bin "./contracts/ethereum/MainchainGovernanceAdmin.bin" --pkg "governance" --out "./generated_contracts/ethereum/governance/mainchain_governance_admin.go"

	solc --abi --bin "${contract_path}/contracts/multi-chains/RoninTrustedOrganization.sol" -o "./contracts/ronin" --include-path "${contract_path}/node_modules/" --base-path ${contract_path} --overwrite --optimize
	abigen --abi "./contracts/ronin/RoninTrustedOrganization.abi" --bin "./contracts/ronin/RoninTrustedOrganization.bin" --pkg "trusted_organization" --out "./generated_contracts/ronin/trusted_organization/ronin_trusted_organization.go"

	solc --abi --bin "${contract_path}/contracts/ronin/validator/RoninValidatorSet.sol" -o "./contracts/ronin" --include-path "${contract_path}/node_modules/" --base-path ${contract_path} --overwrite --optimize
	abigen --abi "./contracts/ronin/RoninValidatorSet.abi" --bin "./contracts/ronin/RoninValidatorSet.bin" --pkg "validator" --out "./generated_contracts/ronin/validator/ronin_validator.go" --alias _isPeriodEnding=IsPeriodEndingInternal