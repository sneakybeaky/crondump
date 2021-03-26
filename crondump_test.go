package crondump

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMinute(t *testing.T) {

	tests := map[string]struct {
		input       string
		errExpected bool
		want        string
	}{
		"single valid minute":   {input: "0", want: "0"},
		"minute on upper value": {input: "59", want: "59"},
		"minute too large":      {input: "60", errExpected: true},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			got, err := Minute(tc.input)

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
