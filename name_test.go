package english_test

import (
	"testing"

	"github.com/mikerybka/english"
)

var name = english.NewName("one two three")

func TestString(t *testing.T) {
	expected := "one two three"
	got := name.String()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestTitleCase(t *testing.T) {
	expected := "One Two Three"
	got := name.TitleCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestPascalCase(t *testing.T) {
	expected := "OneTwoThree"
	got := name.PascalCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestSnakeCase(t *testing.T) {
	expected := "one_two_three"
	got := name.SnakeCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestCamelCase(t *testing.T) {
	expected := "oneTwoThree"
	got := name.CamelCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestScreamingSnakeCase(t *testing.T) {
	expected := "ONE_TWO_THREE"
	got := name.ScreamingSnakeCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestKebabCase(t *testing.T) {
	expected := "one-two-three"
	got := name.KebabCase()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
