package english

func ParseSnakeCase(s string) *Name {
	name := &Name{}
	for _, r := range s {
		if r == '_' {
			name.AppendWord("")
		} else {
			name.AppendRune(r)
		}
	}
	return name
}
