package cliversion

import (
	"log/slog"
	"time"
)

// SlogAttr returns a slog.Group consisting of the relevant parts of a VersionInfo to
// add to a slog message for starting an application.
func (v *VersionInfo) SlogAttr(name string) slog.Attr {
	return slog.Group(name,
		slog.String("version", v.GetBuild().GetVersion()),
		slog.String("date", v.GetBuild().GetDate().AsTime().Format(time.RFC3339)),
		slog.String("commit", v.GetGit().GetCommit()),
		slog.String("built_by", v.GetBuild().GetMethod()),
		slog.String("repo", v.GetGit().GetRepo()),
	)
}
