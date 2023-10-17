package english_test

import (
	"testing"

	"github.com/mikerybka/english"
)

func TestParsePascalCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"P", "P"},
		{"AB", "A B"},
		{"ABC", "A B C"},
		{"ABCd", "A B Cd"},
		{"AbCd", "Ab Cd"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			name := english.ParsePascalCase(tt.in)
			if name.String() != tt.want {
				t.Errorf("got %q", name.String())
			}
		})
	}
}
