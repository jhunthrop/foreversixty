//go:build darwin

package install

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// The Mac's real calls, on throwaway files: ditto keeps the bundle whole,
// hdiutil's own listing finds a mounted image under the app, the volume
// ejects, and a file goes to the Trash. Applications and the login item are
// left alone; Decide and Install cover them with fakes.
func TestTheMacSystemOnThrowawayFiles(t *testing.T) {
	if testing.Short() {
		t.Skip("makes and mounts a disk image")
	}
	sys := NewSystem()
	work := t.TempDir()
	app := newApp(t, work, "body")

	copied := filepath.Join(work, "copy", AppName)
	if err := os.MkdirAll(filepath.Dir(copied), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := sys.Copy(app, copied); err != nil {
		t.Fatal(err)
	}
	if readExe(t, copied) != "body" {
		t.Error("ditto did not copy the bundle")
	}

	dmg := filepath.Join(work, "fs.dmg")
	mount := filepath.Join(work, "mnt")
	if out, err := exec.Command("hdiutil", "create", "-volname", "FS Test", "-srcfolder", app, "-format", "UDZO", dmg).CombinedOutput(); err != nil {
		t.Fatalf("hdiutil create: %v %s", err, out)
	}
	if out, err := exec.Command("hdiutil", "attach", "-nobrowse", "-mountpoint", mount, dmg).CombinedOutput(); err != nil {
		t.Fatalf("hdiutil attach: %v %s", err, out)
	}
	real, err := filepath.EvalSymlinks(mount) // hdiutil names the real path, /private/var for /var
	if err != nil {
		t.Fatal(err)
	}
	plan := Decide(filepath.Join(real, "Forever Sixty 2.app"), Detect(work))
	if !plan.Move || plan.Volume == "" {
		t.Errorf("a bundle on a mounted image: %+v", plan)
	}
	if err := sys.Eject(mount); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(mount, "Contents")); err == nil {
		t.Error("the volume is still mounted")
	}

	if _, translocated := originalPath(app); translocated {
		t.Error("a bundle in a temp folder reads as translocated")
	}
	junk := filepath.Join(work, "trash-me.dmg")
	if err := os.WriteFile(junk, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := sys.Trash(junk); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(junk); !os.IsNotExist(err) {
		t.Error("the file is still there")
	}
}
