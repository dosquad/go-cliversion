package makever

func retfunc(key, value string) Option {
	return func() string {
		return "-X " + key + "=" + value
	}
}

func BuildDate(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildDate",
		in,
	)
}
func BuildDebug(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildDebug",
		in,
	)
}
func BuildMethod(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildMethod",
		in,
	)
}
func BuildVersion(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildVersion",
		in,
	)
}

func GitRepo(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitRepo",
		in,
	)
}
func GitSlug(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitSlug",
		in,
	)
}
func GitCommit(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitCommit",
		in,
	)
}
func GitTag(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitTag",
		in,
	)
}
func GitExactTag(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitExactTag",
		in,
	)
}
