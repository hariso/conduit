// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package field

import (
	"github.com/conduitio/conduit-commons/config"
)

func (renameConfig) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		"mapping": {
			Default:     "",
			Description: "Mapping is a comma separated list of keys and values for fields and their new names (keys and values\nare separated by colons \":\"). For example: `.Metadata.key:id,.Payload.After.foo:bar`.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
	}
}