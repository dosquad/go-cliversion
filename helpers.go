package cliversion

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetBuildDebug returns the versionInfo.GetBuild().GetDebug() if versionInfo.GetBuild() is not nil,
// otherwise it returns true.
func GetBuildDebug(versionInfo *VersionInfo) bool {
	if versionInfo.HasBuild() {
		return versionInfo.GetBuild().GetDebug()
	}

	return true
}

// GetBuildMethod returns the versionInfo.GetBuild().GetMethod() if versionInfo.GetBuild() is not nil,
// otherwise it returns "not-set".
func GetBuildMethod(versionInfo *VersionInfo) string {
	if versionInfo.HasBuild() {
		return versionInfo.GetBuild().GetMethod()
	}

	return "not-set"
}

// GetBuildDate returns the versionInfo.GetBuild().GetDate().AsTime().Format(timeLayout)
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildDate(versionInfo *VersionInfo, timeLayout string) string {
	if versionInfo.HasBuild() {
		if !versionInfo.GetBuild().GetDate().IsZero() {
			return versionInfo.GetBuild().GetDate().Format(timeLayout)
		}
	}

	return ""
}

// GetBuildVersion returns the versionInfo.GetBuild().GetVersion()
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildVersion(versionInfo *VersionInfo) string {
	if versionInfo.HasBuild() {
		return versionInfo.GetBuild().GetVersion()
	}

	return ""
}

// GetBuildGoVersion returns the versionInfo.GetBuild().GetGoVersion()
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildGoVersion(versionInfo *VersionInfo) string {
	if versionInfo.HasBuild() {
		return versionInfo.GetBuild().GetGoVersion()
	}

	return ""
}

// message GitInfo {
// 	string repo = 1 [json_name = "repo"];
// 	string slug = 2 [json_name = "slug"];
// 	string commit = 3 [json_name = "commit"];
// 	string tag = 4 [json_name = "tag"];
// 	string exact_tag = 5 [json_name = "exact_tag"];
// }

// GetGitRepo returns the versionInfo.GetGit().GetRepo()
// if versionInfo.GetGit() is not nil, otherwise it returns an empty string.
func GetGitRepo(versionInfo *VersionInfo) string {
	if versionInfo.GetGit() != nil {
		return versionInfo.GetGit().GetRepo()
	}

	return ""
}

// GetGitSlug returns the versionInfo.GetGit().GetSlug()
// if versionInfo.GetGit() is not nil, otherwise it returns an empty string.
func GetGitSlug(versionInfo *VersionInfo) string {
	if versionInfo.GetGit() != nil {
		return versionInfo.GetGit().GetSlug()
	}

	return ""
}

// GetGitCommit returns the versionInfo.GetGit().GetCommit()
// if versionInfo.GetGit() is not nil, otherwise it returns an empty string.
func GetGitCommit(versionInfo *VersionInfo) string {
	if versionInfo.GetGit() != nil {
		return versionInfo.GetGit().GetCommit()
	}

	return ""
}

// GetGitTag returns the versionInfo.GetGit().GetTag()
// if versionInfo.GetGit() is not nil, otherwise it returns an empty string.
func GetGitTag(versionInfo *VersionInfo) string {
	if versionInfo.GetGit() != nil {
		return versionInfo.GetGit().GetTag()
	}

	return ""
}

// GetGitExactTag returns the versionInfo.GetGit().GetExactTag()
// if versionInfo.GetGit() is not nil, otherwise it returns an empty string.
func GetGitExactTag(versionInfo *VersionInfo) string {
	if versionInfo.GetGit() != nil {
		return versionInfo.GetGit().GetExactTag()
	}

	return ""
}

// JSON returns the VersionInfo formatted into JSON format.
func (v *VersionInfo) JSON() ([]byte, error) {
	return json.Marshal(v)
}

// VersionString returns the version as a string, compatible with display in a cli tool.
func (v *VersionInfo) VersionString() string {
	bld := v.GetBuild()
	git := v.GetGit()

	if bld != nil {
		dateStr := "()"
		if !bld.GetDate().IsZero() {
			dateStr = "(" + bld.GetDate().Format(time.RFC3339) + ")"
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
