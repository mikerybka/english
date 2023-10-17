package english

import (
	"strings"
	"unicode"
)

// Name represents an English-language name, meaning a list of words.
type Name struct {
	Words []Word
}

// NewName parses a string into a Name.
func NewName(s string) *Name {
	words := strings.Split(s, " ")
	name := Name{}
	for _, word := range words {
		if word == "" {
			continue
		}
		name.Words = append(name.Words, Word(word))
	}
	return &name
}

// AppendRune appends a rune to the end of the last word in the name.
func (n *Name) AppendRune(r rune) {
	r = unicode.ToLower(r)
	if len(n.Words) == 0 {
		n.Words = []Word{Word(string(r))}
		return
	}
	n.Words[len(n.Words)-1] += Word(string(r))
}

// AppendWord appends a word to the end of the name.
func (n *Name) AppendWord(s string) {
	n.Words = append(n.Words, Word(s))
}

func (n *Name) Equal(other Name) bool {
	if n.Words == nil && other.Words == nil {
		return true
	}
	if n.Words == nil || other.Words == nil {
		return false
	}
	if len(n.Words) != len(other.Words) {
		return false
	}
	for i := range n.Words {
		if n.Words[i] != other.Words[i] {
			return false
		}
	}
	return true
}

// String returns the string representation of the name.
func (n *Name) String() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString(" ")
		}
		str.WriteString(word.String())
	}
	return str.String()
}

// TitleCase returns the name with the first letter of each word capitalized.
// If the word is an article, conjunction, or propostion it will be lowercased.
// Example: "Table of Contents".
func (n *Name) TitleCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString(" ")
		}
		if i > 0 && (word.isArtictle() || word.isConjunction() || word.isPreposition()) {
			str.WriteString(word.String())
		} else {
			str.WriteString(word.TitleCase())
		}
	}
	return str.String()
}

// PascalCase returns the name with the first letter of each word capitalized, no spaces between words, and all non-alphanumeric characters removed.
// Example: "TableOfContents".
func (n *Name) PascalCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString("")
		}
		str.WriteString(word.PascalCase())
	}
	return str.String()
}

// SnakeCase returns the name with all letters lowercased, an underscore (_) between words, and all non-alphanumeric characters removed.
// Example: "table_of_contents".
func (n *Name) SnakeCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString("_")
		}
		str.WriteString(word.LowerCase())
	}
	return str.String()
}

// CamelCase returns the name with the first letter of each word capitalized, no spaces between words, and all non-alphanumeric characters removed.
// Example: "tableOfContents".
func (n *Name) CamelCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString("")
		}
		if i == 0 {
			str.WriteString(word.LowerCase())
		} else {
			str.WriteString(word.PascalCase())
		}
	}
	return str.String()
}

// ScreamingSnakeCase returns the name with all letters capitalized, an underscore (_) between words, and all non-alphanumeric characters removed.
// Example: "TABLE_OF_CONTENTS".
func (n *Name) ScreamingSnakeCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString("_")
		}
		str.WriteString(word.UpperCase())
	}
	return str.String()
}

// KebabCase returns the name with all letters lowercased, a hyphen (-) between words, and all non-alphanumeric characters removed.
// Example: "table-of-contents".
func (n *Name) KebabCase() string {
	var str strings.Builder
	for i, word := range n.Words {
		if i != 0 {
			str.WriteString("-")
		}
		str.WriteString(word.LowerCase())
	}
	return str.String()
}
