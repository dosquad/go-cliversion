package makever

// LDFlags produces an -ldflags compatible output made up from specified input Option functions.
func LDFlags(items ...Option) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, item())
	}

	return out
}
