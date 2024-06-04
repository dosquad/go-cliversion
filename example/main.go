package main

import (
	"encoding/json"
	"os"

	version "github.com/dosquad/go-cliversion"
)

// var Version version.Version

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(version.Get())
}
