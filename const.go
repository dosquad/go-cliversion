package cliversion

import (
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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

	return false
}

func stringToTime(in string) *timestamppb.Timestamp {
	for _, f := range []string{
		time.RFC1123, time.RFC1123Z,
		time.RFC3339, time.DateTime,
	} {
		if ts, err := time.Parse(f, in); err == nil {
			return timestamppb.New(ts)
		}
	}

	if i, err := strconv.ParseInt(in, 10, 64); err == nil {
		return timestamppb.New(time.Unix(i, 0))
	}

	return timestamppb.New(time.Time{})
}

func Get() *VersionInfo {
	return &VersionInfo{
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
