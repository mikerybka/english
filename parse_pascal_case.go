package english

import "unicode"

func ParsePascalCase(s string) *Name {
	name := &Name{}
	for _, r := range s {
		if unicode.IsUpper(r) {
			name.AppendWord(string(r))
		} else {
			name.AppendRune(r)
		}
	}
	return name
}
