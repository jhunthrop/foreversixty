package wow

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
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

func TestCombatLogsAndSavedVariablesArePerInstall(t *testing.T) {
	root := t.TempDir()
	in := fakeInstall(t, root, "_classic_era_", "ACCOUNT#1")
	fakeInstall(t, root, "_classic_era_", "ACCOUNT#2") // second account, same install
	for _, n := range []string{"WoWCombatLog.txt", "WoWCombatLog-120926_200000.txt", "notes.md"} {
		if err := os.WriteFile(filepath.Join(in.Logs, n), nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	logs, err := in.CombatLogs()
	if err != nil {
		t.Fatal(err)
	}
	if len(logs) != 2 {
		t.Fatalf("CombatLogs = %v", logs)
	}
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

func TestEnsureLogsCreatesTheDirectory(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "_classic_era_")
	if err := os.MkdirAll(filepath.Join(dir, "WTF"), 0o755); err != nil {
		t.Fatal(err)
	}
	in := installAt(dir)
	if err := in.EnsureLogs(); err != nil {
		t.Fatal(err)
	}
	if fi, err := os.Stat(in.Logs); err != nil || !fi.IsDir() {
		t.Fatalf("Logs = %v, %v", fi, err)
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
