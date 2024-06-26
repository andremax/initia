package types

// staking module event types
const (
	EventTypeCompleteUnbonding         = "complete_unbonding"
	EventTypeCompleteRedelegation      = "complete_redelegation"
	EventTypeCreateValidator           = "create_validator"
	EventTypeEditValidator             = "edit_validator"
	EventTypeDelegate                  = "delegate"
	EventTypeUnbond                    = "unbond"
	EventTypeCancelUnbondingDelegation = "cancel_unbonding_delegation"
	EventTypeRedelegate                = "redelegate"

	AttributeKeyValidator      = "validator"
	AttributeKeyCommissionRate = "commission_rate"
	AttributeKeySrcValidator   = "source_validator"
	AttributeKeyDstValidator   = "destination_validator"
	AttributeKeyDelegator      = "delegator"
	AttributeKeyCompletionTime = "completion_time"
	AttributeKeyCreationHeight = "creation_height"
	AttributeKeyNewShares      = "new_shares"
	AttributeValueCategory     = ModuleName
)
