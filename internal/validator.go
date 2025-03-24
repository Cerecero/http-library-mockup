package internal

import "strings"

func AsTitle(key string) string {
	if key == "" {
		panic("empty header key")
	}
	if isTitleCase(key) {
		return key
	}

	return newTitleCase(key)
}

func newTitleCase(key string) string {
	var b strings.Builder
	b.Grow(len(key))
	for i := range key {
		if i == 0 || key[i-1] == "-" {
			b.WriteByte(upper(key[i]))
		} else {
			b.WriteByte(lower(key[i]))
		}
	}
	return b.String()
}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}

func upper(b byte) byte {

	panic("unimplemented")
}



