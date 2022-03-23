package params

import (
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

var (
	// BypassMinFeeMsgTypesKey defines the configuration key for the
	// BypassMinFeeMsgTypes value.
	// nolint: gosec
	BypassMinFeeMsgTypesKey = "bypass-min-fee-msg-types"

	// CustomConfigTemplate defines Gaia's custom application configuration TOML
	// template. It extends the core SDK template.
	CustomConfigTemplate = serverconfig.DefaultConfigTemplate + `
	###############################################################################
	###                        Custom Gaia Configuration                        ###
	###############################################################################
	
	# bypass-min-fee-msg-types defines custom message types the operator may set that
	# will bypass minimum fee checks during CheckTx.
	bypass-min-fee-msg-types = {{ .BypassMinFeeMsgTypes }}
	`
)

// CustomAppConfig defines Gaia's custom application configuration.
type CustomAppConfig struct {
	serverconfig.Config

	// BypassMinFeeMsgTypes defines custom message types the operator may set that
	// will bypass minimum fee checks during CheckTx.
	BypassMinFeeMsgTypes []string `mapstructure:"bypass-min-fee-msg-types"`
}
