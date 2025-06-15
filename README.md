# go-cliversion

**go-cliversion** is a Go library that helps CLI tools manage and display their version information. It provides a structured way to capture build details, Git repository information, and present it in various formats (formatted string or JSON).

This library is useful for:
- Standardizing version output across your Go CLI applications.
- Easily embedding detailed build and SCM information into your binaries.
- Providing version information in machine-readable (JSON) and human-readable formats.

## Features

- Captures Go version, build date, custom version strings, and build methods.
- Captures Git commit hash, tag, repository URL, and slug.
- Outputs version information as a formatted string, ideal for `mytool version` commands.
- Outputs version information as JSON for programmatic access.
- Provides a helper map for generating display headers.
- Uses Protocol Buffers for the underlying data structure.

## Installation

To add go-cliversion to your Go project, run:

```bash
go get github.com/dosquad/go-cliversion
```

## Usage

The primary way to use `go-cliversion` is to have version information embedded into your Go binary at build time using `-ldflags`. The library then provides a convenient way to access this information.

### 1. Populating Version Information via ldflags

`go-cliversion` exposes global string variables (e.g., `cliversion.BuildDate`, `cliversion.GitCommit`) that are intended to be populated by the Go linker.

In your application, you can then call `cliversion.Get()` to retrieve a `*cliversion.VersionInfo` struct containing all the embedded information.

```go
package main

import (
	"fmt"

	"github.com/dosquad/go-cliversion"
)

// AppVersionInfo will be populated by cliversion.Get()
var AppVersionInfo *cliversion.VersionInfo

func init() {
	// Get() reads the ldflags-populated variables and constructs VersionInfo
	AppVersionInfo = cliversion.Get()
}

func main() {
	// Your CLI app logic
	// ...

	// Example: Print version string
	fmt.Println("Version:", AppVersionInfo.VersionString())

	// Example: Print JSON
	jsonData, err := AppVersionInfo.JSON()
	if err == nil {
		fmt.Println("JSON Version:", string(jsonData))
	}
}

```

### 2. Generating ldflags

You need to pass the appropriate `-X` flags to `go build` to populate the variables that `cliversion.Get()` reads.

#### Option A: Using the `makever` helper (Recommended for Go-based build scripts)

The `go-cliversion/makever` sub-package provides helper functions to construct the `ldflags` array. This is useful if your build process is managed by a Go program (e.g., a `magefile`).

**Example (`magefile.go` or custom build program):**

```go
// magefile.go
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/dosquad/go-cliversion/makever"
	"github.com/magefile/mage/mg"
)

func getGitCommit() string {
	out, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	return strings.TrimSpace(string(out))
}

func getGitTag() string {
	out, _ := exec.Command("git", "describe", "--tags", "--abbrev=0", "--exact-match").Output()
	if len(out) > 0 {
		return strings.TrimSpace(string(out))
	}
	out, _ = exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	return strings.TrimSpace(string(out))
}

func getGitRepo() string {
    out, _ := exec.Command("git", "config", "--get", "remote.origin.url").Output()
    return strings.TrimSpace(string(out))
}

func getGitSlug() string {
    repoURL := getGitRepo()
    base := strings.TrimSuffix(repoURL, ".git")
    parts := strings.Split(base, "/")
    if len(parts) >= 2 {
        return parts[len(parts)-2] + "/" + parts[len(parts)-1]
    }
    return "" // Or handle error
}


func getGoVersion() string {
    out, _ := exec.Command("go", "version").Output()
    parts := strings.Fields(string(out))
    if len(parts) >= 3 {
        return parts[2]
    }
    return "unknown"
}


// Build compiles the application.
func Build() error {
	mg.Deps(InstallDeps) // Example dependency

	ldflags := makever.LDFlags(
		makever.BuildVersion("v1.2.4"), // Or dynamically determine
		makever.BuildDate(time.Now().Format(time.RFC3339)),
		makever.BuildMethod("magefile"),
		makever.BuildGoVersion(getGoVersion()),
		makever.GitCommit(getGitCommit()),
		makever.GitRepo(getGitRepo()),
		makever.GitTag(getGitTag()),
		makever.GitExactTag(getGitTag()), // Simplified for example, might need specific logic for exact match
		makever.GitSlug(getGitSlug()),
	)

	fmt.Println("Building with ldflags:", strings.Join(ldflags, " "))

	cmd := exec.Command("go", "build", "-ldflags", strings.Join(ldflags, " "), "-o", "myapp", "yourpackage/main")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// InstallDeps ensures project dependencies are installed.
func InstallDeps() error {
	fmt.Println("Installing dependencies...")
	cmd := exec.Command("go", "mod", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	os.Setenv("GO111MODULE", "on")
}
```
*(Note: The `getGitSlug` and other helper functions in the magefile example are illustrative. You'd implement them robustly for your needs.)*

The `makever.LDFlags` function generates an array of strings like:
`["-X github.com/dosquad/go-cliversion.BuildVersion=v1.2.4", "-X github.com/dosquad/go-cliversion.GitCommit=deadbeef", ...]`
which is then passed to `go build`.

#### Option B: Manually specifying ldflags

You can also construct the `ldflags` string manually. This is useful for simpler shell scripts or Makefiles.

**Build command with manual ldflags:**

```bash
COMMIT=$(git rev-parse HEAD)
TAG=$(git describe --tags --abbrev=0 --exact-match 2>/dev/null || git describe --tags --abbrev=0 2>/dev/null)
EXACT_TAG=$(git describe --tags --exact-match 2>/dev/null)
REPO=$(git config --get remote.origin.url)
# A more robust slug extraction might be needed depending on URL format (git@ vs https)
SLUG=$(basename -s .git $(git config --get remote.origin.url) | sed 's|.*:||' | sed 's|ssh://git@||' | sed 's|https://||' | sed 's|http://||' | tr -d '
')
DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GO_VER=$(go version | awk '{print $3}')
VERSION_PKG="github.com/dosquad/go-cliversion"

LD_FLAGS="-X ${VERSION_PKG}.BuildVersion=v1.2.3"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.BuildDate=${DATE}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.BuildMethod=manual-script"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.BuildGoVersion=${GO_VER}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.GitCommit=${COMMIT}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.GitRepo=${REPO}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.GitTag=${TAG}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.GitExactTag=${EXACT_TAG}"
LD_FLAGS="${LD_FLAGS} -X ${VERSION_PKG}.GitSlug=${SLUG}"

go build -ldflags="${LD_FLAGS}" myapp.go
```

### 3. Displaying Version Information

Once `AppVersionInfo` is populated by `cliversion.Get()`, you can use its methods:

#### Formatted String

The `VersionString()` method returns a human-readable string, suitable for CLI output.

```go
// Assuming AppVersionInfo is populated by cliversion.Get()
versionStr := AppVersionInfo.VersionString()
fmt.Println(versionStr)
```

**Typical Output (will vary based on ldflags provided):**

`v1.2.3 [a1b2c3d] (2023-10-27T10:30:00Z) <manual-script> <git@github.com:user/repo.git>`

If some information is missing (e.g., an ldflag was not set), the output gracefully degrades. For example, if `GitCommit` was not set:
`v1.2.3 [] (2023-10-27T10:30:00Z) <manual-script> <git@github.com:user/repo.git>`


#### JSON Output

The `JSON()` method returns the version information as a JSON byte slice.

```go
jsonData, err := AppVersionInfo.JSON()
if err != nil {
    // handle error
}
fmt.Println(string(jsonData))
```

**Typical Output (will vary based on ldflags provided):**

```json
{
  "bld": { // Note: "bld" instead of "build" due to protobuf field name
    "method": "manual-script",
    "date": "2023-10-27T10:30:00Z",
    "version": "v1.2.3",
    "go_version": "go1.21.3"
  },
  "git": {
    "repo": "git@github.com:user/repo.git",
    "slug": "user/repo",
    "commit": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0",
    "tag": "v1.2.3",
    "exact_tag": "v1.2.3"
  }
}
```
*(Note: The JSON field name for BuildInfo is `bld` as defined in the protobuf (`json:"bld,omitempty"`). The README previously showed `"build"`, this has been corrected.)*


### 4. Using `VersionHeaderMap()`

The `VersionHeaderMap()` function provides a map that can be useful for generating headers when displaying version information in a tabular format.

```go
headers := cliversion.VersionHeaderMap()
// Example:
// buildHeaders := headers["Build"].(map[string]any)
// fmt.Println(buildHeaders["Version"]) // Output: Version
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes, including tests if applicable.
4. Ensure your code passes linting and tests (`golangci-lint run` and `go test ./...`).
5. Submit a pull request.

## License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.
