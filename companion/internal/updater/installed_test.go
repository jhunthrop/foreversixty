package updater

import (
	"os"
	"path/filepath"
	"testing"
)

// The Windows setup installs the companion for one account, under
// %LOCALAPPDATA%\Programs\Forever Sixty, and the companion keeps its update
// stage in %APPDATA%\ForeverSixty\update (internal/paths): two folders, maybe
// two volumes, neither needing admin rights. The swap must happen beside the
// installed exe, so its Start menu shortcut and Start with Windows entry keep
// pointing at the new build, and never through the stage folder.
func TestAnUpdateSwapsTheSetupsInstalledExeInPlace(t *testing.T) {
	sk, pub := keypair(t)
	root := t.TempDir()
	stage := filepath.Join(root, "AppData", "Roaming", "ForeverSixty", "update")
	install := filepath.Join(root, "AppData", "Local", "Programs", "Forever Sixty")
	for _, dir := range []string{stage, install} {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	exe := filepath.Join(install, "foreversixty-companion.exe")
	if err := os.WriteFile(exe, []byte("0.2.44"), 0o755); err != nil {
		t.Fatal(err)
	}
	if BundleRoot(exe) != "" {
		t.Fatalf("the installed exe was taken for an app bundle")
	}
	stagePending(t, stage, sk, []byte("0.2.45"))

	swapped, err := ApplyPending(stage, pub, exe)
	if err != nil || !swapped {
		t.Fatalf("ApplyPending = %v, %v", swapped, err)
	}
	if b, _ := os.ReadFile(exe); string(b) != "0.2.45" {
		t.Errorf("the installed exe holds %q", b)
	}
	// Only the exe and the copy it replaced are in the install folder: the
	// uninstaller removes exactly these (packaging/windows/companion.iss).
	entries, _ := os.ReadDir(install)
	names := map[string]bool{}
	for _, e := range entries {
		names[e.Name()] = true
	}
	if len(names) != 2 || !names["foreversixty-companion.exe"] || !names["foreversixty-companion.exe.previous"] {
		t.Errorf("the install folder holds %v", names)
	}
	if left, _ := os.ReadDir(stage); len(left) != 0 {
		t.Errorf("the stage still holds %v", left)
	}
}

// The Mac app the first launch installs (internal/install) is updated where
// it sits, in /Applications or the player's own Applications folder.
func TestTheInstalledMacAppIsItsOwnBundle(t *testing.T) {
	for _, exe := range []string{
		"/Applications/Forever Sixty.app/Contents/MacOS/foreversixty-companion",
		"/Users/thrall/Applications/Forever Sixty.app/Contents/MacOS/foreversixty-companion",
	} {
		want := filepath.Dir(filepath.Dir(filepath.Dir(exe)))
		if got := BundleRoot(filepath.FromSlash(exe)); got != filepath.FromSlash(want) {
			t.Errorf("BundleRoot(%q) = %q, want %q", exe, got, want)
		}
	}
}
