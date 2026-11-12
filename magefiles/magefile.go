//go:build mage

package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"

	//mage:import
	"github.com/dosquad/mage"
	"github.com/dosquad/mage/helper/paths"
)

// Install update, protoc, format, tidy, lint & test.
func Install(ctx context.Context) {
	mg.CtxDeps(ctx, mage.Protobuf.GenGo)
	// mg.CtxDeps(ctx, mage.Protobuf.GenGoGRPC)
	// versionpb runs first so its coverage does not overwrite the root module's cover.out.
	mg.CtxDeps(ctx, VersionPB.Lint)
	mg.CtxDeps(ctx, VersionPB.Test)
	mg.CtxDeps(ctx, mage.Golang.Lint)
	mg.CtxDeps(ctx, mage.Golang.Test)
}

// VersionPB namespace groups targets for the versionpb module.
type VersionPB mg.Namespace

// Lint run golangci-lint in ./versionpb.
func (VersionPB) Lint(ctx context.Context) error {
	return inDir("versionpb", func() error { return mage.Golang{}.Lint(ctx) })
}

// Test run the test suite in ./versionpb.
func (VersionPB) Test(ctx context.Context) error {
	if err := inDir("versionpb", func() error { return mage.Golang{}.Test(ctx) }); err != nil {
		return err
	}

	coverDir := paths.MustGetArtifactPath("coverage")
	return os.Rename(
		filepath.Join(coverDir, "cover.out"),
		filepath.Join(coverDir, "cover-versionpb.out"),
	)
}

// inDir runs fn from dir; mage.Golang targets only accept packages relative to the working directory.
func inDir(dir string, fn func() error) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return err
	}
	defer func() { _ = os.Chdir(wd) }()

	return fn()
}

var Default = Install
