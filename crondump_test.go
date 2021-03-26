package crondump_test

import (
	"crondump"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

func TestMinute(t *testing.T) {

	tests := map[string]struct {
		input       string
		errExpected bool
		want        string
	}{
		"single valid minute":               {input: "0", want: "0"},
		"minute on upper value":             {input: "59", want: "59"},
		"minute too large":                  {input: "60", errExpected: true},
		"minute too small":                  {input: "-1", errExpected: true},
		"range of two minutes":              {input: "0-1", want: "0 1"},
		"range of all minutes":              {input: "0-59", want: allMinutes()},
		"range ends too high":               {input: "0-60", errExpected: true},
		"range starts too low":              {input: "-1-59", errExpected: true},
		"range too high":                    {input: "60-61", errExpected: true},
		"range too low":                     {input: "-10--1", errExpected: true},
		"list of two minutes":               {input: "1,2", want: "1 2"},
		"list of three minutes":             {input: "1,2,3", want: "1 2 3"},
		"list of a range and single minute": {input: "1-2,10", want: "1 2 10"},
		"Unrestricted range":                {input: "*", want: allMinutes()},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			got, err := crondump.ExpandMinute(tc.input)

			errorReceived := err != nil

			if errorReceived != tc.errExpected {
				t.Fatalf("Unexpected error status %v ", err)
			}

			if !cmp.Equal(tc.want, got) {
				t.Fatalf(cmp.Diff(tc.want, got))
			}

		})
	}

}

func allMinutes() string {
	var sb strings.Builder

	for i := 0; i < 60; i++ {
		sb.WriteString(fmt.Sprintf("%d ", i))
	}
	return strings.TrimSpace(sb.String())
}
