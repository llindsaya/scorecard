// Copyright 2022 OpenSSF Scorecard Authors
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
package evaluation

import (
	"testing"

	"github.com/ossf/scorecard/v5/finding"
	scut "github.com/ossf/scorecard/v5/utests"
)

// Tip: If you add new findings to this test, else
// add a unit test to the probes with the same findings.
func TestCITests(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		findings []finding.Finding
		result   scut.TestReturn
	}{
		{
			name: "Has CI tests. 1 tested out of 1 merged",
			findings: []finding.Finding{
				{
					Outcome:  finding.OutcomeTrue,
					Probe:    "testsRunInCI",
					Message:  "CI test found: pr: 1, context: e2e",
					Location: &finding.Location{Type: 4},
				},
			},
			result: scut.TestReturn{
				Score:         10,
				NumberOfDebug: 1,
			},
		},
		{
			name: "Has CI tests. 3 tested out of 4 merged",
			findings: []finding.Finding{
				{
					Outcome:  finding.OutcomeTrue,
					Probe:    "testsRunInCI",
					Message:  "CI test found: pr: 1, context: e2e",
					Location: &finding.Location{Type: 4},
				},
				{
					Outcome:  finding.OutcomeTrue,
					Probe:    "testsRunInCI",
					Message:  "CI test found: pr: 1, context: e2e",
					Location: &finding.Location{Type: 4},
				},
				{
					Outcome:  finding.OutcomeTrue,
					Probe:    "testsRunInCI",
					Message:  "CI test found: pr: 1, context: e2e",
					Location: &finding.Location{Type: 4},
				},
				{
					Outcome:  finding.OutcomeFalse,
					Probe:    "testsRunInCI",
					Message:  "CI test found: pr: 1, context: e2e",
					Location: &finding.Location{Type: 4},
				},
			},
			result: scut.TestReturn{
				Score:         7,
				NumberOfDebug: 4,
			},
		},
		{
			name: "Tests debugging",
			findings: []finding.Finding{
				{
					Outcome:  finding.OutcomeFalse,
					Probe:    "testsRunInCI",
					Message:  "merged PR 1 without CI test at HEAD: 1",
					Location: &finding.Location{Type: 4},
				},
				{
					Outcome:  finding.OutcomeFalse,
					Probe:    "testsRunInCI",
					Message:  "merged PR 1 without CI test at HEAD: 1",
					Location: &finding.Location{Type: 4},
				},
				{
					Outcome:  finding.OutcomeFalse,
					Probe:    "testsRunInCI",
					Message:  "merged PR 1 without CI test at HEAD: 1",
					Location: &finding.Location{Type: 4},
				},
			},
			result: scut.TestReturn{
				NumberOfDebug: 3,
				Score:         0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dl := scut.TestDetailLogger{}
			got := CITests(tt.name, tt.findings, &dl)
			scut.ValidateTestReturn(t, tt.name, &tt.result, &got, &dl)
		})
	}
}
