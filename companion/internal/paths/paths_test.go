package paths

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHomeHonoursTheOverride(t *testing.T) {
	dir := t.TempDir()
	t.Setenv(HomeEnv, dir)
	got, err := Home()
	if err != nil {
		t.Fatal(err)
	}
	if got != dir {
		t.Fatalf("Home() = %q, want %q", got, dir)
	}
}

func TestHomeIsTheOSApplicationDirectory(t *testing.T) {
	t.Setenv(HomeEnv, "")
	got, err := Home()
	if err != nil {
		t.Fatal(err)
	}
	base, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Dir(got) != base {
		t.Fatalf("Home() = %q, want a child of %q", got, base)
	}
	if name := filepath.Base(got); name != "ForeverSixty" && name != "foreversixty" {
		t.Fatalf("directory name = %q", name)
	}
}

func TestResolveCreatesEveryDirectory(t *testing.T) {
	root := t.TempDir()
	t.Setenv(HomeEnv, filepath.Join(root, "app"))
	d, err := Resolve()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{d.Home, d.Queue, d.State, d.Logs, d.Update} {
		fi, err := os.Stat(p)
		if err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		if !fi.IsDir() {
			t.Errorf("%s is not a directory", p)
		}
	}
	if want := filepath.Join(d.Home, "config.json"); d.ConfigFile() != want {
		t.Errorf("ConfigFile() = %q, want %q", d.ConfigFile(), want)
	}
	if want := filepath.Join(d.State, "abc.json"); d.StateFile("abc") != want {
		t.Errorf("StateFile() = %q, want %q", d.StateFile("abc"), want)
	}
}
