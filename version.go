package version

type Version struct {
	Build BuildInfo `json:"build,omitempty" yaml:"build,omitempty"`
	Git   GitInfo   `json:"git,omitempty" yaml:"git,omitempty"`
}

type BuildInfo struct {
	Debug  bool   `json:"debug,omitempty" yaml:"debug,omitempty"`
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
	Date   Time   `json:"date,omitempty" yaml:"date,omitempty"`
}

type GitInfo struct {
	Repo     string `json:"repo,omitempty" yaml:"repo,omitempty"`
	Slug     string `json:"slug,omitempty" yaml:"slug,omitempty"`
	Commit   string `json:"commit,omitempty" yaml:"commit,omitempty"`
	Tag      string `json:"tag,omitempty" yaml:"tag,omitempty"`
	ExactTag string `json:"exact-tag,omitempty" yaml:"exact-tag,omitempty"`
}
