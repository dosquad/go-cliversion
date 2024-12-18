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

func stringToBool(in string) *bool {
	for _, v := range []string{
		"true", "t",
		"yes", "y",
		"1", "on",
	} {
		if strings.EqualFold(in, v) {
			b := true
			return &b
		}
	}

	for _, v := range []string{
		"false", "f",
		"no", "n",
		"0", "off",
	} {
		if strings.EqualFold(in, v) {
			b := false
			return &b
		}
	}

	return nil
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

func stringToPtr(in string) *string {
	if in != "" {
		return &in
	}

	return nil
}

// Get returns the variables set at compile time.
func Get() *VersionInfo {
	return VersionInfo_builder{
		Bld: BuildInfo_builder{
			Date:      stringToTime(BuildDate),
			Debug:     stringToBool(BuildDebug),
			Method:    stringToPtr(BuildMethod),
			Version:   stringToPtr(BuildVersion),
			GoVersion: stringToPtr(BuildGoVersion),
		}.Build(),
		Git: GitInfo_builder{
			Repo:     stringToPtr(GitRepo),
			Slug:     stringToPtr(GitSlug),
			Commit:   stringToPtr(GitCommit),
			Tag:      stringToPtr(GitTag),
			ExactTag: stringToPtr(GitExactTag),
		}.Build(),
	}.Build()
}
