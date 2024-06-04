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

// import "time"

// type VersionInfo struct {
// 	Build BuildInfo `json:"build,omitempty" yaml:"build,omitempty"`
// 	Git   GitInfo   `json:"git,omitempty" yaml:"git,omitempty"`
// }

// type BuildInfo struct {
// 	Debug     bool      `json:"debug,omitempty" yaml:"debug,omitempty"`
// 	Method    string    `json:"method,omitempty" yaml:"method,omitempty"`
// 	Date      time.Time `json:"date,omitempty" yaml:"date,omitempty"`
// 	Version   string    `json:"version,omitempty" yaml:"version,omitempty"`
// 	GoVersion string    `json:"go-version,omitempty" yaml:"go-version,omitempty"`
// }

// type GitInfo struct {
// 	Repo     string `json:"repo,omitempty" yaml:"repo,omitempty"`
// 	Slug     string `json:"slug,omitempty" yaml:"slug,omitempty"`
// 	Commit   string `json:"commit,omitempty" yaml:"commit,omitempty"`
// 	Tag      string `json:"tag,omitempty" yaml:"tag,omitempty"`
// 	ExactTag string `json:"exact-tag,omitempty" yaml:"exact-tag,omitempty"`
// }
