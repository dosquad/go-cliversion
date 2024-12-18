package cliversion

import (
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
	if b := v.GetBld(); b != nil {
		if g := v.GetGit(); g != nil {
			return b.GetVersion() + " [" +
				g.GetCommit() + "] (" +
				b.GetDate().AsTime().Format(time.RFC3339) + ") <" +
				b.GetMethod() + "> <" + g.GetRepo() + ">"
		}

		return b.GetVersion() + " [] (" +
			b.GetDate().AsTime().Format(time.RFC3339) + ") <" +
			b.GetMethod() + "> <>"
	}

	if g := v.GetGit(); g != nil {
		return g.GetTag() + " [" +
			g.GetCommit() + "] () <> <" + g.GetRepo() + ">"
	}

	return "v0.0.0+unknown [] () <> <>"
}

// single-allocation for the version properties (for use with headers in displaying lists).
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
