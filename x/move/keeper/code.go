package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/initia/x/move/types"
	vmtypes "github.com/initia-labs/initiavm/types"
)

type CodeKeeper struct {
	*Keeper
}

// NewCodeKeeper create new CodeKeeper instance
func NewCodeKeeper(k *Keeper) CodeKeeper {
	return CodeKeeper{k}
}

// Load the allow_arbitrary flag from the move store
func (k CodeKeeper) GetAllowArbitrary(ctx sdk.Context) (bool, error) {
	bz, err := k.GetResourceBytes(ctx, vmtypes.StdAddress, vmtypes.StructTag{
		Address:  vmtypes.StdAddress,
		Module:   vmtypes.Identifier(types.MoveModuleNameCode),
		Name:     vmtypes.Identifier(types.ResourceNameModuleStore),
		TypeArgs: []vmtypes.TypeTag{},
	})
	if err != nil {
		return false, err
	}

	return vmtypes.NewDeserializer(bz).DeserializeBool()
}

// Store the allow_arbitrary flag to move store.
func (k CodeKeeper) SetAllowArbitrary(ctx sdk.Context, allow bool) error {
	ser := vmtypes.NewSerializer()
	if err := ser.SerializeBool(allow); err != nil {
		return err
	}

	return k.ExecuteEntryFunction(
		ctx,
		vmtypes.StdAddress,
		vmtypes.StdAddress,
		types.MoveModuleNameCode,
		types.FunctionNameCodeSetAllowArbitrary,
		[]vmtypes.TypeTag{},
		[][]byte{
			ser.GetBytes(),
		},
	)
}

// GetUpgradePolicy reads upgrade policy from the code module.
func (k CodeKeeper) GetUpgradePolicy(ctx sdk.Context, addr vmtypes.AccountAddress, name string) (types.UpgradePolicy, error) {
	st := vmtypes.StructTag{
		Address:  vmtypes.StdAddress,
		Module:   vmtypes.Identifier(types.MoveModuleNameCode),
		Name:     vmtypes.Identifier(types.ResourceNameMetadataStore),
		TypeArgs: []vmtypes.TypeTag{},
	}

	bz, err := k.GetResourceBytes(ctx, addr, st)
	if err != nil {
		return types.UpgradePolicy_ARBITRARY, err
	}

	tableHandle, err := types.ReadMetadataTableHandleFromMetadataStore(bz)
	if err != nil {
		return types.UpgradePolicy_ARBITRARY, err
	}

	vmAddr, err := vmtypes.NewAccountAddressFromBytes(addr[:])
	if err != nil {
		return types.UpgradePolicy_ARBITRARY, err
	}

	tableKey, err := vmtypes.SerializeString(vmtypes.NewModuleId(vmAddr, name).String())
	if err != nil {
		return types.UpgradePolicy_ARBITRARY, err
	}

	tableEntry, err := k.GetTableEntryBytes(ctx, tableHandle, tableKey)
	if err != nil {
		return types.UpgradePolicy_ARBITRARY, err
	}

	return types.ReadUpgradePolicyFromModuleMetadata(tableEntry.ValueBytes)
}