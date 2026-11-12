package versionpb

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

// JSON returns the VersionInfo formatted into JSON format.
func (v *VersionInfo) JSON() ([]byte, error) {
	return protojson.Marshal(v)
}

// VersionString returns the version as a string, compatible with display in a cli tool.
func (v *VersionInfo) VersionString() string {
	bld := v.GetBld()
	git := v.GetGit()

	if bld != nil {
		dateStr := "()"
		if bld.GetDate() != nil {
			dateStr = "(" + bld.GetDate().AsTime().Format(time.RFC3339) + ")"
		}

		if git != nil {
			return fmt.Sprintf("%s [%s] %s <%s> <%s>",
				bld.GetVersion(), git.GetCommit(), dateStr, bld.GetMethod(), git.GetRepo())
		}
		return fmt.Sprintf("%s [] %s <%s> <>",
			bld.GetVersion(), dateStr, bld.GetMethod())
	}

	if git != nil {
		return fmt.Sprintf("%s [%s] () <> <%s>",
			git.GetTag(), git.GetCommit(), git.GetRepo())
	}

	return "v0.0.0+unknown [] () <> <>"
}
