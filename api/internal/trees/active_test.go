package trees

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// twoBuilds is a builds directory with an older and a newer client build, the shape the
// nightly leaves when it fetched a build no activation has moved to yet.
func twoBuilds(t *testing.T) string {
	t.Helper()
	files := map[string]string{}
	for _, v := range []string{"1.60.1.70658", "1.60.1.70685"} {
		files[v+"/classes.json"] = oneClass
		files[v+"/races.json"] = oneRace
		files[v+"/combos.json"] = oneCombo
		files[v+"/talents/warrior.json"] = oneTree
	}
	return writeBuild(t, files)
}

func activeOf(t *testing.T, dir string) (string, *Data) {
	t.Helper()
	d, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	b, ok := d.Active()
	if !ok {
		t.Fatal("no active build")
	}
	return b.Version, d
}

func TestActiveIsTheBuildTheFileNamesNotTheNewestFolder(t *testing.T) {
	root := twoBuilds(t)
	if err := WriteActiveBuildFile(root, "1.60.1.70658"); err != nil {
		t.Fatal(err)
	}
	got, d := activeOf(t, root)
	if got != "1.60.1.70658" {
		t.Fatalf("active = %q with 70685 fetched but not activated, want 1.60.1.70658", got)
	}
	if d.ActiveFrom() != filepath.Join(root, ActiveBuildFile) {
		t.Fatalf("ActiveFrom = %q", d.ActiveFrom())
	}
	// The fetched build stays loadable by version; it just is not the active one.
	if _, ok := d.Build("1.60.1.70685"); !ok {
		t.Fatal("the fetched build must still load by version")
	}

	// Activating it (the file now names it) switches the active build.
	if err := WriteActiveBuildFile(root, "1.60.1.70685"); err != nil {
		t.Fatal(err)
	}
	if got, _ := activeOf(t, root); got != "1.60.1.70685" {
		t.Fatalf("active after activation = %q, want 1.60.1.70685", got)
	}
}

func TestActiveReadsTheRepositoryFileForTheRepositoryBuildsDir(t *testing.T) {
	repo := t.TempDir()
	builds := filepath.Join(repo, "data", "builds")
	if err := os.MkdirAll(filepath.Dir(builds), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(twoBuilds(t), builds); err != nil {
		t.Fatal(err)
	}
	webData := filepath.Join(repo, "web", "src", "data")
	if err := os.MkdirAll(webData, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := WriteActiveBuildFile(webData, "1.60.1.70658"); err != nil {
		t.Fatal(err)
	}
	if got, d := activeOf(t, builds); got != "1.60.1.70658" || !strings.HasSuffix(d.ActiveFrom(), filepath.Join("web", "src", "data", ActiveBuildFile)) {
		t.Fatalf("active = %q from %q, want 1.60.1.70658 from web/src/data", got, d.ActiveFrom())
	}
	// A file beside the builds (the image's /data/active-build.json) wins.
	if err := WriteActiveBuildFile(builds, "1.60.1.70685"); err != nil {
		t.Fatal(err)
	}
	if got, _ := activeOf(t, builds); got != "1.60.1.70685" {
		t.Fatalf("active = %q, want the file beside the builds to win", got)
	}
}

func TestActiveFallsBackToTheNewestFolderOnlyWithoutAFile(t *testing.T) {
	got, d := activeOf(t, twoBuilds(t))
	if got != "1.60.1.70685" || d.ActiveFrom() != "" {
		t.Fatalf("active = %q from %q, want the newest folder and no file", got, d.ActiveFrom())
	}
}

func TestLoadFailsOnABadActiveBuildFile(t *testing.T) {
	for name, body := range map[string]string{
		"names a build not on disk": `{"build": "1.60.1.99999"}`,
		"names no build":            `{"build": ""}`,
		"is not json":               `1.60.1.70658`,
	} {
		t.Run(name, func(t *testing.T) {
			root := twoBuilds(t)
			if err := os.WriteFile(filepath.Join(root, ActiveBuildFile), []byte(body), 0o644); err != nil {
				t.Fatal(err)
			}
			if _, err := Load(root); err == nil {
				t.Fatal("a bad active-build file must fail the load, not fall back")
			}
		})
	}
}

func TestReadActiveBuild(t *testing.T) {
	dir := t.TempDir()
	if _, err := ReadActiveBuild(filepath.Join(dir, ActiveBuildFile)); err == nil {
		t.Fatal("an absent file has no fallback here")
	}
	if err := WriteActiveBuildFile(dir, "1.60.1.70658"); err != nil {
		t.Fatal(err)
	}
	if got, err := ReadActiveBuild(filepath.Join(dir, ActiveBuildFile)); err != nil || got != "1.60.1.70658" {
		t.Fatalf("ReadActiveBuild = %q, %v", got, err)
	}
}

// The repository's own data: whatever builds data.yml has fetched, the API's active build is
// the one web/src/data/active-build.json names, the file the site, the addon and the guides
// follow.
func TestRepositoryActiveBuildIsActiveBuildJSON(t *testing.T) {
	want, err := ReadActiveBuild(filepath.Join("..", "..", "..", "web", "src", "data", ActiveBuildFile))
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := activeOf(t, filepath.Join("..", "..", "..", "data", "builds")); got != want {
		t.Fatalf("repository active = %q, want active-build.json's %q", got, want)
	}
}
