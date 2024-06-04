package makever

import "strings"

type Option func() string

func LDFlags(items ...Option) string {
	sb := []string{}
	for _, item := range items {
		sb = append(sb, item())
	}

	return strings.Join(sb, " ")
}
