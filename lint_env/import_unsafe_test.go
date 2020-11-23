package lint_env

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}

func TestMain(m *testing.M) {
	tempDir := os.Getenv("TEST_TMPDIR")
	srcDir := os.Getenv("TEST_SRCDIR")
	if len(tempDir) == 0 || len(srcDir) == 0 || runtime.GOROOT() != "GOROOT" {
		os.Exit(m.Run())
	}
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	// assuming the Bazel rule has `data = ["@go_sdk//:files"]`
	goRoot := filepath.Join(srcDir, "__main__/external/go_sdk")
	// Go executable requires a GOCACHE to be set after go1.12.
	cmd.Env = append(os.Environ(),
		"GOCACHE="+filepath.Join(tempDir, "cache"),
		"GOROOT="+goRoot,
		fmt.Sprintf("PATH=%s:%s", filepath.Join(goRoot, "bin"), os.Getenv("PATH")),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		os.Exit(0)
	}
	if xerr, ok := err.(*exec.ExitError); ok {
		os.Exit(xerr.ExitCode())
	}
	fmt.Fprintln(os.Stderr, err)
	// Exist with 6, in line with Bazel's RUN_FAILURE.
	os.Exit(6)
}
