package cliversion_test

import (
	"strings"
	"testing"

	"github.com/dosquad/go-cliversion/makever"
	"github.com/google/go-cmp/cmp"
)

func TestBuildVersion(t *testing.T) {
	tests := []struct {
		args   []string
		expect string
	}{
		{nil, "{\n    \"build\": {\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"git\": {}\n}"},
		{
			[]string{`-ldflags`, makever.LDFlags(
				makever.BuildVersion("0.0.0"),
			)},
			readTestData("test-2.json"),
		},
		{
			[]string{`-ldflags`, `-X main.hash=ffffffff`},
			readTestData("empty.json"),
		},
		{
			[]string{"-ldflags", makever.LDFlags(
				makever.BuildDate("2024-06-04T14:43:17+10:00"),
				makever.BuildDebug("true"),
				makever.BuildMethod("testfile"),
				makever.BuildVersion("1.2.3"),
				makever.GitCommit("824bfc6a0f1b4314ea260276a4ce0099b94d6ccd"),
				makever.GitExactTag(""),
				makever.GitTag("v0.3.0-8-gf749871"),
				makever.GitRepo("git@github.com:dosquad/go-cliversion.git"),
				makever.GitSlug("dosquad/go-cliversion"),
			)},
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
			// if !strings.EqualFold(blankVersion, tt.expect) {
			// 	t.Errorf("Output: got '%s', want '%s'", blankVersion, tt.expect)
			// }
		})
	}
}
