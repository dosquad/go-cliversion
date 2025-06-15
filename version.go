package cliversion

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

// GetBuild is the pre-protobuf editions compatibility.
//
// Deprecated: GetBuild exists for historical compatibility
// and should not be used. To access the the BuildInfo use
// GetBld instead.
func (v *VersionInfo) GetBuild() *BuildInfo {
	return v.GetBld()
}

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

// single-allocation for the version properties (for use with headers in displaying lists).
//
//nolint:gochecknoglobals // This map is a constant and is used to provide a consistent header format
var versionHeaderMap = map[string]any{
	"Build": map[string]any{
		"Debug":     "Debug",
		"Method":    "Method",
		"Date":      "Date",
		"Version":   "Version",
		"GoVersion": "Go Version",
	},
	"Git": map[string]any{
		"Repo":     "Repository",
		"Slug":     "Repo Slug",
		"Commit":   "Commit",
		"Tag":      "Tag",
		"ExactTag": "Exact Tag",
	},
}

// VersionHeaderMap returns the versionHeaderMap.
func VersionHeaderMap() map[string]any {
	return versionHeaderMap
}
