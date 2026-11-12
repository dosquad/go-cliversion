package cliversion

import "time"

type VersionInfo struct {
	Build *BuildInfo `json:"build,omitempty"`
	Git   *GitInfo   `json:"git,omitempty"`
}

type BuildInfo struct {
	Date      time.Time `json:"date,omitempty"`
	Debug     bool      `json:"debug,omitempty"`
	Method    string    `json:"method,omitempty"`
	Version   string    `json:"version,omitempty"`
	GoVersion string    `json:"go_version,omitempty"`
}

type GitInfo struct {
	Repo     string `json:"repo,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Commit   string `json:"commit,omitempty"`
	Tag      string `json:"tag,omitempty"`
	ExactTag string `json:"exact_tag,omitempty"`
}

func (v *VersionInfo) HasBuild() bool {
	return v.Build != nil
}

func (v *VersionInfo) GetBuild() *BuildInfo {
	return v.Build
}

func (v *VersionInfo) HasGit() bool {
	return v.Git != nil
}

func (v *VersionInfo) GetGit() *GitInfo {
	return v.Git
}

func (vb *BuildInfo) SetDate(in time.Time) {
	vb.Date = in
}

func (vb *BuildInfo) GetDate() time.Time {
	return vb.Date
}

func (vb *BuildInfo) SetDebug(in bool) {
	vb.Debug = in
}

func (vb *BuildInfo) GetDebug() bool {
	return vb.Debug
}

func (vb *BuildInfo) SetMethod(in string) {
	vb.Method = in
}

func (vb *BuildInfo) GetMethod() string {
	return vb.Method
}

func (vb *BuildInfo) SetVersion(in string) {
	vb.Version = in
}

func (vb *BuildInfo) GetVersion() string {
	return vb.Version
}

func (vb *BuildInfo) SetGoVersion(in string) {
	vb.GoVersion = in
}

func (vb *BuildInfo) GetGoVersion() string {
	return vb.GoVersion
}

func (vg *GitInfo) SetRepo(in string) {
	vg.Repo = in
}

func (vg *GitInfo) GetRepo() string {
	return vg.Repo
}

func (vg *GitInfo) SetSlug(in string) {
	vg.Slug = in
}

func (vg *GitInfo) GetSlug() string {
	return vg.Slug
}

func (vg *GitInfo) SetCommit(in string) {
	vg.Commit = in
}

func (vg *GitInfo) GetCommit() string {
	return vg.Commit
}

func (vg *GitInfo) SetTag(in string) {
	vg.Tag = in
}

func (vg *GitInfo) GetTag() string {
	return vg.Tag
}

func (vg *GitInfo) SetExactTag(in string) {
	vg.ExactTag = in
}

func (vg *GitInfo) GetExactTag() string {
	return vg.ExactTag
}
