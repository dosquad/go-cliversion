package cliversion

import (
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

func (v *VersionInfo) JSON() ([]byte, error) {
	return protojson.Marshal(v)
}

func (v *VersionInfo) VersionString() string {
	if b := v.GetBuild(); b != nil {
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

func VersionHeaderMap() map[string]any {
	return versionHeaderMap
}
