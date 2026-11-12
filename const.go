package cliversion

import (
	"strconv"
	"strings"
	"time"
)

//nolint:gochecknoglobals // These variables are set at compile time using the -ldflags option.
var (
	BuildDate      string
	BuildDebug     string
	BuildMethod    string
	BuildVersion   string
	BuildGoVersion string

	GitRepo     string
	GitSlug     string
	GitCommit   string
	GitTag      string
	GitExactTag string
)

func stringToBool(in string) bool {
	for _, v := range []string{
		"true", "t",
		"yes", "y",
		"1", "on",
	} {
		if strings.EqualFold(in, v) {
			return true
		}
	}

	for _, v := range []string{
		"false", "f",
		"no", "n",
		"0", "off",
	} {
		if strings.EqualFold(in, v) {
			return false
		}
	}

	return false
}

func stringToTime(in string) time.Time {
	for _, f := range []string{
		time.RFC1123, time.RFC1123Z,
		time.RFC3339, time.DateTime,
	} {
		if ts, err := time.Parse(f, in); err == nil {
			return ts
		}
	}

	if i, err := strconv.ParseInt(in, 10, 64); err == nil {
		return time.Unix(i, 0)
	}

	return time.Time{}
}

//nolint:gochecknoglobals // cached version info.
var versionInfoCache *VersionInfo

// Get returns the variables set at compile time.
func Get() *VersionInfo {
	if versionInfoCache == nil {
		versionInfoCache = &VersionInfo{
			Build: &BuildInfo{
				Date:      stringToTime(BuildDate),
				Debug:     stringToBool(BuildDebug),
				Method:    BuildMethod,
				Version:   BuildVersion,
				GoVersion: BuildGoVersion,
			},
			Git: &GitInfo{
				Repo:     GitRepo,
				Slug:     GitSlug,
				Commit:   GitCommit,
				Tag:      GitTag,
				ExactTag: GitExactTag,
			},
		}
	}

	return versionInfoCache
}
