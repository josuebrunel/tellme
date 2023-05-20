package tellme

import (
	"os"
	"runtime"
	"testing"
)

const (
	env          = "dev"
	name         = "myApp"
	version      = "0.0.1"
	buildStamp   = "2023-10-10T10:10:10"
	author       = "Josh Loking <jloking@example.org>"
	commitId     = "1e8a5e1c-5d52-4aa6-9eb5-03d0043cb663"
	commitTime   = "2023-10-10T08:10:10"
	commitBranch = "main"
	arch         = runtime.GOARCH
	archOS       = runtime.GOOS
)

func TestApp(t *testing.T) {
	setEnvVars := map[string]string{
		"env":     "dev",
		"appName": "myApp",
	}

	for key, val := range setEnvVars {
		os.Setenv(key, val)
	}
	app := NewApp(
		env,
		name,
		version,
		buildStamp,
		author,
		commitId,
		commitTime,
		commitBranch,
	)

	t.Run("Env", func(t *testing.T) {
		for key, val := range setEnvVars {
			if app.Env[key] != val {
				t.Fatalf("[ASSERT-ERROR] %s: expected %s but got %s\n", key, val, app.Env[key])
			}
		}
	})

	t.Run("Info", func(t *testing.T) {
		if app.Info.EnvName != env {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", env, app.Info.EnvName)
		}

		if app.Info.AppName != name {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", name, app.Info.AppName)
		}
		if app.Info.Version != version {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", version, app.Info.Version)
		}
		if app.Info.BuildStamp != buildStamp {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", buildStamp, app.Info.BuildStamp)
		}
		if app.Info.CommitAuthor != author {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", author, app.Info.CommitAuthor)
		}
		if app.Info.CommitID != commitId {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", commitId, app.Info.CommitID)
		}
		if app.Info.CommitTime != commitTime {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", commitTime, app.Info.CommitTime)
		}
		if app.Info.CommitBranch != commitBranch {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", commitBranch, app.Info.CommitBranch)
		}
		if app.Info.Arch != arch {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", arch, app.Info.Arch)
		}
		if app.Info.OS != archOS {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", archOS, app.Info.OS)
		}
		archVersion := runtime.Version()
		if app.Info.RuntimeVersion != archVersion {
			t.Fatalf("[ASSERT-ERROR] expected %s but got %s\n", archVersion, app.Info.RuntimeVersion)
		}
	})

	t.Run("All", func(t *testing.T) {
		app.GetInfo()
		app.GetEnv()
		app.GetMetrics()
		app.GetThreadDump()
	})
}
