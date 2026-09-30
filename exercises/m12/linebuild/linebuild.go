package linebuild

import (
	"strings"
)

func Plus(parts []string) string {
	var out string
	for _, part := range parts {
		out += part
	}
	return out
}

func Builder(parts []string) string {
	var b strings.Builder
	total := 0
	for _, part := range parts {
		total += len(part)
	}
	b.Grow(total)
	for _, part := range parts {
		b.WriteString(part)
	}
	return b.String()
}
