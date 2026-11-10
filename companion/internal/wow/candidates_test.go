package wow

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCandidatesOnAMacAreBothApplicationsFolders(t *testing.T) {
	got := candidatesFor("darwin", "/Users/jh", func(string) string { return "" })
	want := []string{
		"/Applications/World of Warcraft",
		filepath.Join("/Users/jh", "Applications", "World of Warcraft"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("darwin candidates = %q, want %q", got, want)
	}
}

func TestCandidatesOnWindowsFollowTheProgramFilesVariables(t *testing.T) {
	env := map[string]string{"ProgramFiles(x86)": `D:\Games`, "ProgramFiles": ""}
	got := candidatesFor("windows", `C:\Users\jh`, func(k string) string { return env[k] })
	want := []string{filepath.Join(`D:\Games`, "World of Warcraft"), `C:\Program Files (x86)\World of Warcraft`}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("windows candidates = %q, want %q", got, want)
	}
}

func TestCandidatesOnLinuxAreTheWinePrefixes(t *testing.T) {
	got := candidatesFor("linux", "/home/jh", func(string) string { return "" })
	if len(got) != 2 || filepath.Base(got[0]) != "World of Warcraft" || filepath.Base(got[1]) != "World of Warcraft" {
		t.Fatalf("linux candidates = %q", got)
	}
}

// The Mac layout: the Battle.net folder holds the launcher's .app, Data and
// one flavour folder per game, and each flavour folder holds its client's
// .app beside WTF and Logs (the owner's /Applications/World of Warcraft).
func TestDetectFindsTheBetaInAMacLayout(t *testing.T) {
	root := filepath.Join(t.TempDir(), "World of Warcraft")
	for _, p := range []string{
		filepath.Join(root, "World of Warcraft Launcher.app", "Contents", "MacOS"),
		filepath.Join(root, "Data"),
		filepath.Join(root, "_classic_beta_", "World of Warcraft Beta.app", "Contents", "MacOS"),
		filepath.Join(root, "_retail_", "World of Warcraft.app", "Contents", "MacOS"),
	} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	beta := fakeInstall(t, root, "_classic_beta_", "ACCOUNT#1")
	fakeInstall(t, root, "_retail_", "ACCOUNT#1")

	got := detectIn([]string{filepath.Join(t.TempDir(), "absent"), root, root})
	if len(got) != 1 || got[0].Path != beta.Path {
		t.Fatalf("detect = %+v, want only %s", got, beta.Path)
	}
	if got[0].Logs != filepath.Join(beta.Path, "Logs") || got[0].WTF != filepath.Join(beta.Path, "WTF") {
		t.Errorf("install = %+v", got[0])
	}
	sv, err := got[0].SavedVariables("ForeverSixty")
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(beta.Path, "WTF", "Account", "ACCOUNT#1", "SavedVariables", "ForeverSixty.lua")
	if len(sv) != 1 || sv[0] != want {
		t.Errorf("saved variables = %q, want %q", sv, want)
	}
}

// The live game's folder is found by its WTF and Logs, not by the name of
// the client's .app inside it, so this layout carries no .app in _forever_
// at all. With the beta beside it, the live game comes first.
func TestDetectFindsTheLiveGameFirstInAMacLayout(t *testing.T) {
	root := filepath.Join(t.TempDir(), "World of Warcraft")
	for _, p := range []string{
		filepath.Join(root, "World of Warcraft Launcher.app", "Contents", "MacOS"),
		filepath.Join(root, "_classic_beta_", "World of Warcraft Beta.app", "Contents", "MacOS"),
	} {
		if err := os.MkdirAll(p, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	beta := fakeInstall(t, root, "_classic_beta_", "ACCOUNT#1")
	live := fakeInstall(t, root, "_forever_", "ACCOUNT#1")
	fakeInstall(t, root, "_retail_", "ACCOUNT#1")

	got := detectIn([]string{root})
	if len(got) != 2 || got[0].Path != live.Path || got[1].Path != beta.Path {
		t.Fatalf("detect = %+v, want %s then %s", got, live.Path, beta.Path)
	}
	if FlavorName(got[0].Flavor) != "WoW Forever" {
		t.Errorf("live install named %q", FlavorName(got[0].Flavor))
	}
}
