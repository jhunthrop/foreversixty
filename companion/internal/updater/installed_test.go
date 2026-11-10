package updater

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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

type fakeKey struct {
	values map[string]string
	set    []string
	closed bool
}

func (k *fakeKey) GetString(name string) (string, error) {
	v, ok := k.values[name]
	if !ok {
		return "", errors.New("no such value")
	}
	return v, nil
}

func (k *fakeKey) SetString(name, value string) error {
	k.values[name] = value
	k.set = append(k.set, name+"="+value)
	return nil
}

func (k *fakeKey) Close() error { k.closed = true; return nil }

func opener(k *fakeKey) OpenKey {
	return func(path string) (RegistryKey, error) {
		if path != UninstallKey || k == nil {
			return nil, ErrNoRegistryKey
		}
		return k, nil
	}
}

// After a self-update the setup's Apps & Features entry names the running
// version, for the copy the setup installed and no other.
func TestTheSetupsEntryFollowsTheRunningVersion(t *testing.T) {
	install := filepath.Join("C:", "Users", "thrall", "AppData", "Local", "Programs", "Forever Sixty")
	exe := filepath.Join(install, "foreversixty-companion.exe")
	entry := func(values map[string]string) *fakeKey { return &fakeKey{values: values} }
	for _, tc := range []struct {
		name  string
		key   *fakeKey
		exe   string
		wrote bool
	}{
		{"no entry: not installed by the setup", nil, exe, false},
		{"an older version is brought up", entry(map[string]string{"InstallLocation": install + string(filepath.Separator), "DisplayVersion": "0.2.45"}), exe, true},
		{"the folder matches case aside", entry(map[string]string{"InstallLocation": strings.ToUpper(install), "DisplayVersion": "0.2.45"}), exe, true},
		{"already current", entry(map[string]string{"InstallLocation": install, "DisplayVersion": "0.2.46"}), exe, false},
		{"a bare exe elsewhere leaves the entry alone", entry(map[string]string{"InstallLocation": install, "DisplayVersion": "0.2.45"}), filepath.Join("C:", "Downloads", "foreversixty-companion_windows_amd64.exe"), false},
		{"an entry without a folder is left alone", entry(map[string]string{"DisplayVersion": "0.2.45"}), exe, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wrote, err := SyncInstalledVersion(opener(tc.key), tc.exe, "0.2.46")
			if err != nil || wrote != tc.wrote {
				t.Fatalf("SyncInstalledVersion = %v, %v; want %v", wrote, err, tc.wrote)
			}
			if tc.key == nil {
				return
			}
			if !tc.key.closed {
				t.Error("the key was left open")
			}
			want := 0
			if tc.wrote {
				want = 1
			}
			if len(tc.key.set) != want || (tc.wrote && tc.key.set[0] != "DisplayVersion=0.2.46") {
				t.Errorf("wrote %v", tc.key.set)
			}
		})
	}
}

// The entry's name is the setup's AppId: one GUID in two files.
func TestTheUninstallKeyIsTheSetupsAppID(t *testing.T) {
	iss, err := os.ReadFile(filepath.Join("..", "..", "packaging", "windows", "companion.iss"))
	if err != nil {
		t.Fatal(err)
	}
	// Inno writes a literal brace doubled: AppId={{GUID}.
	if want := "AppId={" + SetupAppID; !strings.Contains(string(iss), want) {
		t.Errorf("companion.iss has no %q", want)
	}
}
