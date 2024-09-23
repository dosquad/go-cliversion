package makever

// LDFlags produces an -ldflags compatible output made up from specified input Option functions.
func LDFlags(items ...Option) []string {
	out := []string{}
	for _, item := range items {
		out = append(out, item())
	}

	return out
}
