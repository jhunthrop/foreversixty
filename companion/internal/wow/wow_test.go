package wow

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
)

// fakeInstall lays out a flavour directory the way the game does.
func fakeInstall(t *testing.T, root, flavor, account string) Install {
	t.Helper()
	dir := filepath.Join(root, flavor)
	for _, p := range []string{
		filepath.Join(dir, "Logs"),
		filepath.Join(dir, "WTF", "Account", account, "SavedVariables"),
	} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	return installAt(dir)
}

func TestScanFindsFlavourDirectoriesUnderARoot(t *testing.T) {
	root := t.TempDir()
	fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	fakeInstall(t, root, "_forever_", "ACCOUNT#1")
	if err := os.MkdirAll(filepath.Join(root, "Utils"), 0o755); err != nil {
		t.Fatal(err)
	}
	got := Scan(root)
	if len(got) != 2 {
		t.Fatalf("Scan found %d installs: %+v", len(got), got)
	}
	if got[0].Flavor != "_classic_era_" || got[1].Flavor != "_forever_" {
		t.Fatalf("flavours = %q, %q", got[0].Flavor, got[1].Flavor)
	}
	if got[0].Logs != filepath.Join(root, "_classic_era_", "Logs") {
		t.Errorf("logs = %q", got[0].Logs)
	}
}

func TestPickAcceptsAFlavourDirectoryOrItsParent(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")

	fromParent, err := Pick(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(fromParent) != 1 || fromParent[0].Path != in.Path {
		t.Fatalf("Pick(parent) = %+v", fromParent)
	}
	fromFlavor, err := Pick(in.Path)
	if err != nil {
		t.Fatal(err)
	}
	if len(fromFlavor) != 1 || fromFlavor[0].Path != in.Path {
		t.Fatalf("Pick(flavour) = %+v", fromFlavor)
	}
}

func TestPickRefusesAFolderThatIsNotTheGame(t *testing.T) {
	if _, err := Pick(t.TempDir()); !errors.Is(err, ErrNotAnInstall) {
		t.Fatalf("err = %v, want ErrNotAnInstall", err)
	}
	if _, err := Pick(filepath.Join(t.TempDir(), "nope")); err == nil {
		t.Fatal("Pick accepted a path that does not exist")
	}
	file := filepath.Join(t.TempDir(), "WoW.exe")
	if err := os.WriteFile(file, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Pick(file); !errors.Is(err, ErrNotAnInstall) {
		t.Fatalf("Pick(file) = %v", err)
	}
}

func TestAdvancedLoggingIsReadFromConfigWTF(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	write := func(body string) {
		t.Helper()
		if err := os.WriteFile(in.AdvancedLoggingPath(), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := in.AdvancedLogging(); err == nil {
		t.Error("a missing Config.wtf was not reported")
	}
	write("SET gxWindow \"1\"\n")
	if on, err := in.AdvancedLogging(); err != nil || on {
		t.Errorf("with no setting: %v, %v", on, err)
	}
	write("SET gxWindow \"1\"\nSET advancedCombatLogging \"0\"\n")
	if on, err := in.AdvancedLogging(); err != nil || on {
		t.Errorf("with the box off: %v, %v", on, err)
	}
	write("SET advancedCombatLogging \"1\"\nSET gxWindow \"1\"\n")
	if on, err := in.AdvancedLogging(); err != nil || !on {
		t.Errorf("with the box on: %v, %v", on, err)
	}
}

func TestSavedVariablesAreOnePerAccount(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	fakeInstall(t, root, "_classic_era_", "ACCOUNT#2") // second account, same install
	sv, err := in.SavedVariables("ForeverSixty")
	if err != nil {
		t.Fatal(err)
	}
	if len(sv) != 2 {
		t.Fatalf("SavedVariables = %v", sv)
	}
	want := filepath.Join(in.WTF, "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if sv[0] != want {
		t.Errorf("SavedVariables[0] = %q, want %q", sv[0], want)
	}
}

func TestCandidatesAreAbsolutePathsForThisOS(t *testing.T) {
	got := Candidates()
	if len(got) == 0 {
		t.Fatal("no candidate paths")
	}
	for _, p := range got {
		if !filepath.IsAbs(p) {
			t.Errorf("%q is not absolute", p)
		}
	}
	// Detect must not panic on a machine with no game installed.
	Detect()
}

func TestFlavourFoldersAreNamedInThePlayersWords(t *testing.T) {
	for folder, want := range map[string]string{
		"_anniversary_":  "Classic Anniversary",
		"_classic_era_":  "Classic Era",
		"_classic_beta_": "Classic beta",
		"_retail_":       "Retail",
		"_forever_":      "_forever_",
		"MyServer":       "MyServer",
	} {
		if got := FlavorName(folder); got != want {
			t.Errorf("FlavorName(%q) = %q, want %q", folder, got, want)
		}
	}
}

func touch(t *testing.T, path string, at time.Time) {
	t.Helper()
	if err := os.WriteFile(path, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, at, at); err != nil {
		t.Fatal(err)
	}
}

func TestLastActivityIsTheNewestFileDirectlyUnderLogs(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_beta_", "ACCOUNT#1")
	old := time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)
	newest := time.Date(2026, 10, 15, 19, 42, 0, 0, time.UTC)
	touch(t, filepath.Join(in.Logs, "General.log"), old)
	touch(t, filepath.Join(in.Logs, "WoWCombatLog.txt"), newest)
	touch(t, filepath.Join(in.Logs, "gx.log"), old.Add(time.Hour))
	// A file below Logs is not the game running.
	sub := filepath.Join(in.Logs, "Crashes")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	touch(t, filepath.Join(sub, "crash.txt"), newest.Add(24*time.Hour))
	got, ok := LastActivity(in)
	if !ok || !got.Equal(newest) {
		t.Fatalf("LastActivity = %v, %v; want %v", got, ok, newest)
	}
}

func TestLastActivityIsUnknownWithoutLogs(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_retail_", "ACCOUNT#1")
	if _, ok := LastActivity(in); ok {
		t.Fatal("an empty Logs directory has activity")
	}
	if err := os.RemoveAll(in.Logs); err != nil {
		t.Fatal(err)
	}
	if _, ok := LastActivity(in); ok {
		t.Fatal("a missing Logs directory has activity")
	}
	if err := os.MkdirAll(filepath.Join(in.Logs, "only-a-dir"), 0o755); err != nil {
		t.Fatal(err)
	}
	if _, ok := LastActivity(in); ok {
		t.Fatal("a Logs directory holding only a folder has activity")
	}
}
