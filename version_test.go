package version_test

import "testing"

func TestBuildVersion(t *testing.T) {
	blankVersion, err := buildTestBinary([]string{})
	if err != nil {
		t.Errorf("buildTestBinary(): error got '%s', expect 'nil'", err)
	}

	t.Logf("blankVersion: %s", blankVersion)
	t.Fail()
}
