package english_test

import (
	"testing"

	"github.com/mikerybka/english"
)

func TestParsePascalCaseWithAcronyms(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"API", "API"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			name := english.ParsePascalCaseWithAcronyms(tt.in)
			if name.String() != tt.want {
				t.Errorf("got %q", name.String())
			}
		})
	}
}
