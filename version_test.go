package cliversion_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dosquad/go-cliversion"
	"github.com/dosquad/go-cliversion/makever"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestVersionString(t *testing.T) {
	testDateStr := "2023-01-15T10:30:00Z"
	testDate, _ := time.Parse(time.RFC3339, testDateStr)
	tsProto := timestamppb.New(testDate)

	tests := []struct {
		name     string
		setupVI  func() *cliversion.VersionInfo
		expected string
	}{
		{
			name: "Full VersionInfo",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("v1.0.0")
				bld.SetMethod("test-method")
				bld.SetDate(tsProto)
				bld.SetGoVersion("go1.20")

				git := &cliversion.GitInfo{}
				git.SetCommit("abcdef")
				git.SetRepo("example.com/repo")
				git.SetTag("v1.0.0")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(git)
				return vi
			},
			expected: "v1.0.0 [abcdef] (2023-01-15T10:30:00Z) <test-method> <example.com/repo>",
		},
		{
			name: "BuildInfo only, with Date",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("v1.1.0")
				bld.SetMethod("build-only")
				bld.SetDate(tsProto)
				bld.SetGoVersion("go1.21")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(nil)
				return vi
			},
			expected: "v1.1.0 [] (2023-01-15T10:30:00Z) <build-only> <>",
		},
		{
			name: "BuildInfo only, nil Date",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("v1.2.0")
				bld.SetMethod("build-nil-date")
				bld.SetDate(nil)
				bld.SetGoVersion("go1.22")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(nil)
				return vi
			},
			expected: "v1.2.0 [] () <build-nil-date> <>",
		},
		{
			name: "GitInfo only, BuildInfo nil",
			setupVI: func() *cliversion.VersionInfo {
				git := &cliversion.GitInfo{}
				git.SetCommit("fedcba")
				git.SetRepo("example.com/git-only")
				git.SetTag("v0.9.0")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(nil)
				vi.SetGit(git)
				return vi
			},
			expected: "v0.9.0 [fedcba] () <> <example.com/git-only>",
		},
		{
			name: "Both BuildInfo and GitInfo nil",
			setupVI: func() *cliversion.VersionInfo {
				vi := &cliversion.VersionInfo{}
				vi.SetBld(nil)
				vi.SetGit(nil)
				return vi
			},
			expected: "v0.0.0+unknown [] () <> <>",
		},
		{
			name: "BuildInfo and GitInfo present, some fields empty",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("") // Empty version
				bld.SetMethod("partial-method")
				bld.SetDate(tsProto)
				bld.SetGoVersion("go1.19")

				git := &cliversion.GitInfo{}
				git.SetCommit("") // Empty commit
				git.SetRepo("example.com/partial-repo")
				git.SetTag("v0.8.0")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(git)
				return vi
			},
			expected: " [] (2023-01-15T10:30:00Z) <partial-method> <example.com/partial-repo>",
		},
		{
			name: "BuildInfo (non-nil Date) and GitInfo (non-empty commit/repo)",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("v1.0.1")
				bld.SetMethod("test-case-7")
				bld.SetDate(tsProto)
				bld.SetGoVersion("go1.20.1")

				git := &cliversion.GitInfo{}
				git.SetCommit("commit7")
				git.SetRepo("example.com/repo7")
				git.SetTag("v1.0.1")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(git)
				return vi
			},
			expected: "v1.0.1 [commit7] (2023-01-15T10:30:00Z) <test-case-7> <example.com/repo7>",
		},
		{
			name: "BuildInfo (nil Date) and GitInfo (non-empty commit/repo)",
			setupVI: func() *cliversion.VersionInfo {
				bld := &cliversion.BuildInfo{}
				bld.SetVersion("v1.0.2")
				bld.SetMethod("test-case-8")
				bld.SetDate(nil) // nil Date
				bld.SetGoVersion("go1.20.2")

				git := &cliversion.GitInfo{}
				git.SetCommit("commit8")
				git.SetRepo("example.com/repo8")
				git.SetTag("v1.0.2")

				vi := &cliversion.VersionInfo{}
				vi.SetBld(bld)
				vi.SetGit(git)
				return vi
			},
			expected: "v1.0.2 [commit8] () <test-case-8> <example.com/repo8>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vi := tt.setupVI()
			got := vi.VersionString()
			if got != tt.expected {
				t.Errorf("VersionString() got = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestBuildVersion(t *testing.T) {
	tests := []struct {
		args   []string
		expect string
	}{
		{nil, "{\n    \"build\": {\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"git\": {}\n}"},
		{
			[]string{`-ldflags`, strings.Join(makever.LDFlags(
				makever.BuildVersion("0.0.0"),
			), " ")},
			readTestData("test-2.json"),
		},
		{
			[]string{`-ldflags`, `-X main.hash=ffffffff`},
			readTestData("empty.json"),
		},
		{
			[]string{"-ldflags", strings.Join(makever.LDFlags(
				makever.BuildDate("2024-06-04T14:43:17+10:00"),
				makever.BuildDebug("true"),
				makever.BuildMethod("testfile"),
				makever.BuildVersion("1.2.3"),
				makever.GitCommit("824bfc6a0f1b4314ea260276a4ce0099b94d6ccd"),
				makever.GitExactTag(""),
				makever.GitTag("v0.3.0-8-gf749871"),
				makever.GitRepo("git@github.com:dosquad/go-cliversion.git"),
				makever.GitSlug("dosquad/go-cliversion"),
			), " ")},
			readTestData("test-1.json"),
		},
	}

	for _, tt := range tests {
		t.Run(strings.ReplaceAll(tt.expect, "    ", ""), func(t *testing.T) {
			blankVersion, err := buildTestBinary(tt.args)
			if err != nil {
				t.Errorf("buildTestBinary(): error got '%s', expect 'nil'", err)
			}

			if diff := cmp.Diff(blankVersion, tt.expect); diff != "" {
				t.Errorf("JSON: version display -got +want:\n%s", diff)
			}
		})
	}
}
