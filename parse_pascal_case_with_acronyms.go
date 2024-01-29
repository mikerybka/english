package english

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ParsePascalCaseWithAcronyms(s string) *Name {
	for _, word := range acronyms {
		replacement := cases.Title(language.AmericanEnglish).String(word)
		s = strings.ReplaceAll(s, word, replacement)
	}
	name := ParsePascalCase(s)
	for i, word := range name.Words {
		if _, ok := fiveLetterAcronyms[word.UpperCase()]; ok {
			name.Words[i] = Word(strings.ToUpper(string(word)))
		}
		if _, ok := fourLetterAcronyms[word.UpperCase()]; ok {
			name.Words[i] = Word(strings.ToUpper(string(word)))
		}
		if _, ok := threeLetterAcronyms[word.UpperCase()]; ok {
			name.Words[i] = Word(strings.ToUpper(string(word)))
		}
		if _, ok := twoLetterAcronyms[word.UpperCase()]; ok {
			name.Words[i] = Word(strings.ToUpper(string(word)))
		}
	}
	return name
}

var acronyms = []string{}

func init() {
	for k, _ := range fiveLetterAcronyms {
		acronyms = append(acronyms, k)
	}
	for k, _ := range fourLetterAcronyms {
		acronyms = append(acronyms, k)
	}
	for k, _ := range threeLetterAcronyms {
		acronyms = append(acronyms, k)
	}
	for k, _ := range twoLetterAcronyms {
		acronyms = append(acronyms, k)
	}
}
