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

func isTitleCase(key string) bool {
	for i := range key {
		if i == 0 || key[i-1] == '-' {
			if key[i] >= 'a' && key[i] <= 'z' {
				return false
			}
		} else if key[i] >= 'A' && key[i] <= 'Z' {
			return false
		}
	}
	return true
}

func newTitleCase(key string) string {
	var b strings.Builder
	b.Grow(len(key))
	for i := range key {
		if i == 0 || key[i-1] == '-' {
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
	if b >= 'a' && b <= 'z' {
		return b + 'A' - 'a'
	}
	return b
}

func splitLines(s string) []string {
	if s == "" {
		return nil
	}
	var lines []string
	i := 0
	for {
		j := strings.Index(s[i:], "\r\n")
		if j == -1 {
			lines = append(lines, s[i:])
			return lines
		}
		lines = append(lines, s[i:i+j])
		i += j + 2
	}
}
