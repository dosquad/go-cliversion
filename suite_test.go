package version_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const gomodTemplate = `module github.com/dosquad/go-cliversion

go 1.22.3


`

const codeTemplate = `package main

import (
	"os"
	"encoding/json"
	version "github.com/dosquad/go-cliversion"
)

var Version version.VersionInfo

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(Version)
}
`

func buildTestBinary(args []string) (string, error) {
	var tempPath string
	{
		f, err := os.CreateTemp("", "testfile-*")
		if err != nil {
			return "", fmt.Errorf("unable to create tempPath: %w", err)
		}
		tempPath = f.Name()
		f.Close()
		if err := os.Remove(tempPath); err != nil {
			return "", fmt.Errorf("unable to delete tempPath")
		}

		if err := os.Mkdir(tempPath, 0o0755); err != nil {
			return "", fmt.Errorf("unable to make tempPathdir")
		}
	}

	if err := os.WriteFile(
		filepath.Join(tempPath, "main.go"),
		[]byte(codeTemplate),
		0o0666,
	); err != nil {
		return "", fmt.Errorf("unable to write temp source file: %s", err)
	}

	targetFile := filepath.Join(tempPath, "main"),
	{
		buildArgs := []string{"build", "-o", targetFile}
		buildArgs = append(buildArgs, args...)
		buildArgs = append(buildArgs, tempPath)

		buildCmd := exec.Command("go", buildArgs...)
		buildCmd.Dir = filepath.Dir(targetFile)

		if err := buildCmd.Run(); err != nil {
			return "", fmt.Errorf("unable to run build command (%s): %w", buildCmd.Args, err)
		}
	}

	buf := &strings.Builder{}
	{
		cmd := exec.Command("./" + filepath.Base(targetFile))
		cmd.Dir = filepath.Dir(targetFile)
		cmd.Stdout = buf

		if err := cmd.Run(); err != nil {
			return "", fmt.Errorf("unable to run compiled command (%s): %w", cmd.Args, err)
		}
	}

	return buf.String(), nil
}
