package versionpb

import (
	"github.com/dosquad/go-cliversion"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Get() *VersionInfo {
	v := cliversion.Get()
	return VersionInfo_builder{
		Bld: BuildInfo_builder{
			Debug:     &v.Build.Debug,
			Method:    &v.Build.Method,
			Date:      timestamppb.New(v.Build.Date),
			Version:   &v.Build.Version,
			GoVersion: &v.Build.GoVersion,
		}.Build(),
		Git: GitInfo_builder{
			Repo:     &v.Git.Repo,
			Slug:     &v.Git.Slug,
			Commit:   &v.Git.Commit,
			Tag:      &v.Git.Tag,
			ExactTag: &v.Git.ExactTag,
		}.Build(),
	}.Build()
}
