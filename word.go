package english

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Word is an English word.
// If the word is a proper noun, it should be capitalized.
// Any other grammer is allowed.
// The word should contain no spaces.
// The word should not be blank.
type Word string

// Filter returns the word with all non-alphanumeric characters removed.
func (w Word) Filter() Word {
	numbers := "0123456789"
	letters := "abcdefghijklmnopqrstuvwxyz"
	capitals := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	filteredWord := ""
	for _, char := range w {
		if strings.ContainsRune(numbers, char) || strings.ContainsRune(letters, char) || strings.ContainsRune(capitals, char) {
			filteredWord += string(char)
		}
	}
	return Word(filteredWord)
}

// String returns the string representation of the word.
func (w Word) String() string {
	return string(w)
}

// TitleCase returns the word with the first letter capitalized.
func (w Word) TitleCase() string {
	return cases.Title(language.English).String(w.String())
}

// PascalCase returns the word with the first letter capitalized and all non-alphanumeric characters removed.
func (w Word) PascalCase() string {
	return w.Filter().TitleCase()
}

// LowerCase returns the word with all letters lowercased.
func (w Word) LowerCase() string {
	return strings.ToLower(w.Filter().String())
}

// UpperCase returns the word with all letters capitalized.
func (w Word) UpperCase() string {
	return strings.ToUpper(w.Filter().String())
}

func (w Word) isArtictle() bool {
	articles := []string{
		"a",
		"an",
		"the",
	}
	for _, article := range articles {
		if w.LowerCase() == article {
			return true
		}
	}
	return false
}

func (w Word) isConjunction() bool {
	conjunctions := []string{
		"and",
		"but",
		"or",
		"nor",
		"for",
		"yet",
		"so",
	}
	for _, conjunction := range conjunctions {
		if w.LowerCase() == conjunction {
			return true
		}
	}
	return false
}

func (w Word) isPreposition() bool {
	prepositions := []string{
		"to",
		"in",
		"into",
		"on",
		"onto",
		"at",
		"by",
	}
	for _, preposition := range prepositions {
		if w.LowerCase() == preposition {
			return true
		}
	}
	return false
}
