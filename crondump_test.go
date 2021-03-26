package crondump

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMinute(t *testing.T) {

	tests := map[string]struct {
		input string
		want  string
	}{
		"single valid minute": {input: "0", want: "0"},
	}

	for desc, tc := range tests {
		t.Run(desc, func(t *testing.T) {
			got := Minute(tc.input)

			if !cmp.Equal(tc.want, got) {
				t.Error(cmp.Diff(tc.want, got))
			}

		})
	}

}
