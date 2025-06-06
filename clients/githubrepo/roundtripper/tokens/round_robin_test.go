// Copyright 2023 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package tokens

import (
	"testing"
)

//nolint:paralleltest // order dependent
func TestNext(t *testing.T) {
	tokens := []string{"token1", "token2", "token3", "token4", "token5"}
	rr := makeRoundRobinAccessor(tokens)

	tests := []struct {
		name      string
		releaseID *uint64 // nil if no token is released
		want      string
	}{
		{"First call", nil, "token2"},
		{"Second call", nil, "token3"},
		{"Third call", nil, "token4"},
		{"Fourth call", nil, "token5"},
		{"After release", func() *uint64 { v := uint64(0); return &v }(), "token1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.releaseID != nil {
				rr.Release(*tt.releaseID)
			}
			_, got := rr.Next()
			if got != tt.want {
				t.Errorf("Next() = %s, want %s", got, tt.want)
			}
		})
	}
}
