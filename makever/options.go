package makever

// Option functions used to build the LDFlags.
type Option func() string

func retfunc(key, value string) Option {
	return func() string {
		return "-X " + key + "=" + value
	}
}

// BuildDate specifies the build date for the application.
func BuildDate(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildDate",
		in,
	)
}

// BuildDebug specifies if the build was debug or not for the application.
func BuildDebug(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildDebug",
		in,
	)
}

// BuildMethod specifies the method the application was built with (eg. "magefile", "goreleaser", etc).
func BuildMethod(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildMethod",
		in,
	)
}

// BuildVersion specifies the version of the application.
func BuildVersion(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildVersion",
		in,
	)
}

// BuildGoVersion specifies go version used to build the application.
func BuildGoVersion(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.BuildGoVersion",
		in,
	)
}

// GitRepo specifies the git repo for the application.
func GitRepo(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitRepo",
		in,
	)
}

// GitSlug specifies the git slug when the application was built.
func GitSlug(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitSlug",
		in,
	)
}

// GitCommit specifies the git commit when the application was built.
func GitCommit(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitCommit",
		in,
	)
}

// GitTag specifies the last git tag in the tree when the application was built.
func GitTag(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitTag",
		in,
	)
}

// GitExactTag specifies the git tag for the commit when the application was built.
func GitExactTag(in string) Option {
	return retfunc(
		"github.com/dosquad/go-cliversion.GitExactTag",
		in,
	)
}
