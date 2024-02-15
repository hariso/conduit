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
	"testing"

	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-processor-sdk"
	"github.com/matryer/is"
)

func TestExcludeFields_Process(t *testing.T) {
	is := is.New(t)
	proc := excludeFields{
		fields: []string{".Metadata", ".Payload.After.foo"},
	}
	records := []opencdc.Record{
		{
			Metadata: map[string]string{"key1": "val1", "key2": "val2"},
			Payload: opencdc.Change{
				After: opencdc.StructuredData{
					"foo":  "bar",
					"keep": "me",
				},
			},
		},
	}
	want := opencdc.Record{

		Metadata: map[string]string{},
		Payload: opencdc.Change{
			After: opencdc.StructuredData{
				"keep": "me",
			},
		},
	}
	output := proc.Process(context.Background(), records)
	is.True(len(output) == 1)
	res, ok := output[0].(sdk.SingleRecord)
	is.True(ok) // output record is not a sdk.SingleRecord type
	is.Equal(res.Metadata, want.Metadata)
	is.Equal(res.Payload.After, want.Payload.After)
}

func TestExcludeField_Configure(t *testing.T) {
	proc := excludeFields{}
	ctx := context.Background()
	testCases := []struct {
		name    string
		cfg     map[string]string
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     map[string]string{"fields": ".Metadata,.Payload"},
			wantErr: false,
		}, {
			name:    "missing parameter",
			cfg:     map[string]string{},
			wantErr: true,
		}, {
			name:    "cannot exclude .Operation",
			cfg:     map[string]string{"fields": ".Operation"},
			wantErr: true,
		}, {
			name:    "cannot exclude .Position",
			cfg:     map[string]string{"fields": ".Position"},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			err := proc.Configure(ctx, tc.cfg)
			if tc.wantErr {
				is.True(err != nil)
				return
			}
			is.NoErr(err)
		})
	}
}
