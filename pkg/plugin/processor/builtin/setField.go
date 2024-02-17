// Copyright © 2024 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builtin

import (
	"context"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-processor-sdk"
	"github.com/conduitio/conduit/pkg/foundation/cerrors"
)

type setField struct {
	referenceResolver sdk.ReferenceResolver
	value             string

	sdk.UnimplementedProcessor
}

func (p *setField) Specification() (sdk.Specification, error) {
	return sdk.Specification{
		Name:    "field.set",
		Summary: "Set the value of a certain field.",
		Description: `Set the value of a certain field to any value. 
Note that this processor only runs on structured data, if the record contains JSON data, then use the processor "decode.json" to parse it into structured data first.`,
		Version: "v0.1.0",
		Author:  "Meroxa, Inc.",
		Parameters: map[string]sdk.Parameter{
			"field": {
				Default:     "",
				Type:        sdk.ParameterTypeString,
				Description: "The target field, as it would be addressed in a Go template.",
				Validations: []sdk.Validation{
					{
						Type: sdk.ValidationTypeRequired,
					},
				},
			},
			"value": {
				Default:     "",
				Type:        sdk.ParameterTypeString,
				Description: `A Go template expression which will be evaluated and stored in "field".`,
				Validations: []sdk.Validation{
					{
						Type: sdk.ValidationTypeRequired,
					},
				},
			},
		},
	}, nil
}

func (p *setField) Configure(_ context.Context, cfg map[string]string) error {
	field, ok := cfg["field"]
	if !ok {
		return cerrors.Errorf("%w (%q)", ErrRequiredParamMissing, "field")
	}
	value, ok := cfg["value"]
	if !ok {
		return cerrors.Errorf("%w (%q)", ErrRequiredParamMissing, "value")
	}
	resolver, err := sdk.NewReferenceResolver(field)
	if err != nil {
		return err
	}
	p.referenceResolver = resolver
	p.value = value

	return nil
}

func (p *setField) Open(context.Context) error {
	return nil
}

func (p *setField) Process(_ context.Context, records []opencdc.Record) []sdk.ProcessedRecord {
	out := make([]sdk.ProcessedRecord, 0, len(records))
	for _, record := range records {
		ref, err := p.referenceResolver.Resolve(&record)
		if err != nil {
			return append(out, sdk.ErrorRecord{Error: err})
		}
		err = ref.Set(p.value) // todo: evaluate value
		if err != nil {
			return append(out, sdk.ErrorRecord{Error: err})
		}
		out = append(out, sdk.SingleRecord(record))
	}
	return out
}

func (p *setField) Teardown(context.Context) error {
	return nil
}