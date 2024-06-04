package cliversion_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const gomodTemplate = `module testmodule

go 1.22.3

replace github.com/dosquad/go-cliversion => %s

require github.com/dosquad/go-cliversion v0.0.0-00010101000000-000000000000
`

const codeTemplate = `package main

import (
	"encoding/json"
	"os"

	version "github.com/dosquad/go-cliversion"
)

func main() {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	_ = enc.Encode(version.Get())
}
`

func readTestData(name string) string {
	if out, err := os.ReadFile("testdata/" + name); err == nil {
		return string(out)
	}

	return ""
}

// workDir get working directory or panic if unable to.
func workDir(path ...string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(append([]string{wd}, path...)...)
}

func buildTestBinary(args []string) (string, error) {
	var tempPath string
	{
		f, err := os.CreateTemp("", "testfile-*")
		if err != nil {
			return "", fmt.Errorf("unable to create tempPath: %w", err)
		}
		tempPath = f.Name()
		f.Close()
	}

	if err := os.Remove(tempPath); err != nil {
		return "", errors.New("unable to delete tempPath")
	}

	if err := os.Mkdir(tempPath, 0o0755); err != nil {
		return "", errors.New("unable to make tempPathdir")
	}

	if err := os.WriteFile(
		filepath.Join(tempPath, "main.go"),
		[]byte(codeTemplate),
		0o0666,
	); err != nil {
		return "", fmt.Errorf("unable to write temp source file: %w", err)
	}

	if err := os.WriteFile(
		filepath.Join(tempPath, "go.mod"),
		[]byte(fmt.Sprintf(gomodTemplate, workDir())),
		0o0666,
	); err != nil {
		return "", fmt.Errorf("unable to write temp go.mod: %w", err)
	}

	defer func() { _ = os.RemoveAll(tempPath) }()

	buf := &strings.Builder{}
	// targetFile := filepath.Join(tempPath, "main")
	{
		buildArgs := []string{"run"}
		buildArgs = append(buildArgs, args...)
		buildArgs = append(buildArgs, "main.go")

		buildCmd := exec.Command("go", buildArgs...)
		buildCmd.Dir = tempPath
		buildCmd.Stdout = buf
		buildCmd.Stderr = os.Stderr

		if err := buildCmd.Run(); err != nil {
			return "", fmt.Errorf("unable to go run command [%s](%s): %w", buildCmd.Dir, buildCmd.Args, err)
		}
	}

	// targetFile := filepath.Join(tempPath, "main")
	// {
	// 	// buildArgs := []string{"build", "-o", targetFile}
	// 	buildArgs := []string{"run"}
	// 	buildArgs = append(buildArgs, args...)
	// 	// buildArgs = append(buildArgs, tempPath)
	// 	buildArgs = append(buildArgs, targetFile+".go")

	// 	buildCmd := exec.Command("go", buildArgs...)
	// 	buildCmd.Dir = filepath.Dir(targetFile)

	// 	if err := buildCmd.Run(); err != nil {
	// 		return "", fmt.Errorf("unable to run build command (%s): %w", buildCmd.Args, err)
	// 	}
	// }

	// buf := &strings.Builder{}
	// {
	// 	cmd := exec.Command("./" + filepath.Base(targetFile))
	// 	cmd.Dir = filepath.Dir(targetFile)
	// 	cmd.Stdout = buf

	// 	if err := cmd.Run(); err != nil {
	// 		return "", fmt.Errorf("unable to run compiled command (%s): %w", cmd.Args, err)
	// 	}
	// }

	return strings.TrimSpace(buf.String()), nil
}
