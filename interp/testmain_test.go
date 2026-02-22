package interp_test

import (
	"os"
	"path/filepath"
	"testing"
)

// testGOPATH is a temporary GOPATH set up by TestMain so that test files
// importing packages from _test/ (e.g. "github.com/GoCodeAlone/yaegi/_test/foo")
// can be resolved. A symlink is created at $tmp/src/github.com/GoCodeAlone/yaegi
// pointing to the repository root.
var testGOPATH string

func TestMain(m *testing.M) {
	tmp, err := os.MkdirTemp("", "yaegi-test-gopath-*")
	if err != nil {
		panic(err)
	}

	// Get repo root (parent of interp/).
	repoRoot, err := filepath.Abs("..")
	if err != nil {
		os.RemoveAll(tmp)
		panic(err)
	}

	// Create symlink: $tmp/src/github.com/GoCodeAlone/yaegi → repoRoot
	linkDir := filepath.Join(tmp, "src", "github.com", "GoCodeAlone")
	if err := os.MkdirAll(linkDir, 0o755); err != nil {
		os.RemoveAll(tmp)
		panic(err)
	}
	if err := os.Symlink(repoRoot, filepath.Join(linkDir, "yaegi")); err != nil {
		os.RemoveAll(tmp)
		panic(err)
	}

	testGOPATH = tmp

	code := m.Run()
	os.RemoveAll(tmp)
	os.Exit(code)
}
