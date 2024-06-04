package makever

type Option func() string

func LDFlags(items ...Option) []string {
	out := []string{}
	for _, item := range items {
		out = append(out, item())
	}

	return out
}
