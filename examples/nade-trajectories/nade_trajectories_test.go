package main

import (
	"os"
	"testing"

	examples "github.com/visual42/demoinfocs-golang/examples"
)

// Just make sure the example runs
func TestBouncyNades(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test")
	}

	os.Args = []string{"cmd", "-demo", "../../cs-demos/default.dem"}

	examples.RedirectStdout(func() {
		main()
	})
}
