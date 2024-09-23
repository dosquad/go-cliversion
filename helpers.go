package cliversion

// message BuildInfo {
// 	bool debug = 1 [json_name = "debug"];
// 	string method = 2 [json_name = "method"];
// 	google.protobuf.Timestamp date = 3 [json_name = "date"];
// 	string version = 4 [json_name = "version"];
// 	string go_version = 5 [json_name = "go_version"];
// }

// GetBuildDebug returns the versionInfo.GetBuild().GetDebug() if versionInfo.GetBuild() is not nil,
// otherwise it returns true.
func GetBuildDebug(versionInfo *VersionInfo) bool {
	if versionInfo.GetBuild() != nil {
		return versionInfo.GetBuild().GetDebug()
	}

	return true
}

// GetBuildMethod returns the versionInfo.GetBuild().GetMethod() if versionInfo.GetBuild() is not nil,
// otherwise it returns "not-set".
func GetBuildMethod(versionInfo *VersionInfo) string {
	if versionInfo.GetBuild() != nil {
		return versionInfo.GetBuild().GetMethod()
	}

	return "not-set"
}

// GetBuildDate returns the versionInfo.GetBuild().GetDate().AsTime().Format(timeLayout)
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildDate(versionInfo *VersionInfo, timeLayout string) string {
	if versionInfo.GetBuild() != nil {
		if versionInfo.GetBuild().GetDate() != nil {
			return versionInfo.GetBuild().GetDate().AsTime().Format(timeLayout)
		}
	}

	return ""
}

// GetBuildVersion returns the versionInfo.GetBuild().GetVersion()
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildVersion(versionInfo *VersionInfo) string {
	if versionInfo.GetBuild() != nil {
		return versionInfo.GetBuild().GetVersion()
	}

	return ""
}

// GetBuildGoVersion returns the versionInfo.GetBuild().GetGoVersion()
// if versionInfo.GetBuild() is not nil, otherwise it returns an empty string.
func GetBuildGoVersion(versionInfo *VersionInfo) string {
	if versionInfo.GetBuild() != nil {
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
