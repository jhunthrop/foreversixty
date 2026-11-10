package install

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const home = "/Users/thrall"

func env(original map[string]string, images ...Image) Env {
	return Env{
		Home: home,
		Original: func(p string) (string, bool) {
			o, ok := original[p]
			return o, ok
		},
		Images: func() []Image { return images },
	}
}

func TestDecide(t *testing.T) {
	const translocated = "/private/var/folders/xy/T/AppTranslocation/1A2B/d/Forever Sixty.app"
	dmg := Image{Path: home + "/Downloads/foreversixty-companion_darwin_universal.dmg", MountPoint: "/Volumes/Forever Sixty"}
	for _, tc := range []struct {
		name   string
		bundle string
		env    Env
		want   Plan
	}{
		{"not an app bundle", "", env(nil), Plan{}},
		{"in Applications", "/Applications/Forever Sixty.app", env(nil), Plan{}},
		{"in the player's Applications", home + "/Applications/Forever Sixty.app", env(nil), Plan{}},
		{"in a folder inside Applications", "/Applications/Games/Forever Sixty.app", env(nil), Plan{
			Move: true, Source: "/Applications/Games/Forever Sixty.app", Original: "/Applications/Games/Forever Sixty.app",
		}},
		{"unzipped in Downloads", home + "/Downloads/Forever Sixty.app", env(nil), Plan{
			Move: true, Source: home + "/Downloads/Forever Sixty.app", Original: home + "/Downloads/Forever Sixty.app",
			Trash: home + "/Downloads/Forever Sixty.app",
		}},
		{"translocated from Downloads", translocated, env(map[string]string{translocated: home + "/Downloads/Forever Sixty.app"}), Plan{
			Move: true, Source: translocated, Original: home + "/Downloads/Forever Sixty.app",
			Trash: home + "/Downloads/Forever Sixty.app",
		}},
		{"on the mounted disk image", "/Volumes/Forever Sixty/Forever Sixty.app", env(nil, dmg), Plan{
			Move: true, Source: "/Volumes/Forever Sixty/Forever Sixty.app", Original: "/Volumes/Forever Sixty/Forever Sixty.app",
			Volume: "/Volumes/Forever Sixty", Trash: dmg.Path,
		}},
		{"translocated from the disk image", translocated, env(map[string]string{translocated: "/Volumes/Forever Sixty/Forever Sixty.app"}, dmg), Plan{
			Move: true, Source: translocated, Original: "/Volumes/Forever Sixty/Forever Sixty.app",
			Volume: "/Volumes/Forever Sixty", Trash: dmg.Path,
		}},
		{"a disk image outside the home folder is ejected, never trashed", "/Volumes/Forever Sixty/Forever Sixty.app",
			env(nil, Image{Path: "/Volumes/Share/fs.dmg", MountPoint: "/Volumes/Forever Sixty"}), Plan{
				Move: true, Source: "/Volumes/Forever Sixty/Forever Sixty.app", Original: "/Volumes/Forever Sixty/Forever Sixty.app",
				Volume: "/Volumes/Forever Sixty",
			}},
		{"a volume of a similar name is not the image", "/Volumes/Forever Sixty 1/Forever Sixty.app", env(nil, dmg), Plan{
			Move: true, Source: "/Volumes/Forever Sixty 1/Forever Sixty.app", Original: "/Volumes/Forever Sixty 1/Forever Sixty.app",
		}},
		{"translocated from Applications stays", translocated, env(map[string]string{translocated: "/Applications/Forever Sixty.app"}), Plan{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Decide(tc.bundle, tc.env); got != tc.want {
				t.Errorf("Decide(%q)\n got %+v\nwant %+v", tc.bundle, got, tc.want)
			}
		})
	}
}

func TestParseImagesReadsHdiutilInfo(t *testing.T) {
	const info = "framework       : 683.160.3\n" +
		"================================================\n" +
		"image-path      : /Users/thrall/Downloads/foreversixty-companion_darwin_universal.dmg\n" +
		"image-type      : read-only disk image\n" +
		"/dev/disk4\tGUID_partition_scheme\t\n" +
		"/dev/disk4s1\t48465300-0000-11AA-AA11-00306543ECAC\t/Volumes/Forever Sixty\n" +
		"================================================\n" +
		"image-path      : /Library/Simulator.dmg\n" +
		"/dev/disk5\tGUID_partition_scheme\t\n" +
		"/dev/disk5s1\t41504653-0000-11AA-AA11-00306543ECAC\t\n"
	got := ParseImages(strings.NewReader(info))
	want := []Image{{Path: "/Users/thrall/Downloads/foreversixty-companion_darwin_universal.dmg", MountPoint: "/Volumes/Forever Sixty"}}
	if len(got) != 1 || got[0] != want[0] {
		t.Errorf("ParseImages = %+v, want %+v", got, want)
	}
}

func TestDestinationsFallBackToThePlayersApplications(t *testing.T) {
	got := Destinations(home)
	if len(got) != 2 || got[0] != "/Applications" || got[1] != home+"/Applications" {
		t.Errorf("Destinations = %v", got)
	}
}

// copyTree stands in for ditto.
func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, p)
		out := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(out, 0o755)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		return os.WriteFile(out, b, 0o755)
	})
}

func newApp(t *testing.T, dir, body string) string {
	t.Helper()
	app := filepath.Join(dir, "Forever Sixty 2.app")
	exe := filepath.Join(app, "Contents", "MacOS", "foreversixty-companion")
	if err := os.MkdirAll(filepath.Dir(exe), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(exe, []byte(body), 0o755); err != nil {
		t.Fatal(err)
	}
	return app
}

func readExe(t *testing.T, app string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(app, "Contents", "MacOS", "foreversixty-companion"))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestInstallReplacesAnOlderCopyUnderTheFixedName(t *testing.T) {
	src := newApp(t, t.TempDir(), "new")
	apps := t.TempDir()
	older := filepath.Join(apps, AppName)
	if err := copyTree(newApp(t, t.TempDir(), "old"), older); err != nil {
		t.Fatal(err)
	}
	got, err := Install(src, []string{apps}, copyTree)
	if err != nil {
		t.Fatal(err)
	}
	if got != older || readExe(t, got) != "new" {
		t.Errorf("Install = %q holding %q", got, readExe(t, got))
	}
	if entries, _ := os.ReadDir(apps); len(entries) != 1 {
		t.Errorf("left beside the app: %v", entries)
	}
}

func TestInstallFallsBackWhenAFolderCannotBeWritten(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("root writes anywhere")
	}
	src := newApp(t, t.TempDir(), "new")
	locked := t.TempDir()
	if err := os.Chmod(locked, 0o555); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chmod(locked, 0o755) })
	own := filepath.Join(t.TempDir(), "Applications") // not there yet: it is made
	got, err := Install(src, []string{locked, own}, copyTree)
	if err != nil {
		t.Fatal(err)
	}
	if got != filepath.Join(own, AppName) {
		t.Errorf("Install = %q", got)
	}
}

func TestInstallKeepsTheOlderCopyWhenTheCopyFails(t *testing.T) {
	apps := t.TempDir()
	older := filepath.Join(apps, AppName)
	if err := copyTree(newApp(t, t.TempDir(), "old"), older); err != nil {
		t.Fatal(err)
	}
	failing := func(string, string) error { return errors.New("disk full") }
	if _, err := Install("/nowhere.app", []string{apps}, failing); !errors.Is(err, ErrNoDestination) {
		t.Fatalf("err = %v", err)
	}
	if readExe(t, older) != "old" {
		t.Error("the older copy was lost")
	}
	if entries, _ := os.ReadDir(apps); len(entries) != 1 {
		t.Errorf("left beside the app: %v", entries)
	}
}

type fakeSystem struct {
	loginItems bool
	calls      []string
	ejectFails int
	relaunched []string
}

func (f *fakeSystem) Copy(src, dst string) error { return copyTree(src, dst) }
func (f *fakeSystem) Relaunch(app string, args []string) error {
	f.relaunched = append([]string{app}, args...)
	return nil
}
func (f *fakeSystem) LoginItems() bool { return f.loginItems }
func (f *fakeSystem) SetLoginItem(on bool) error {
	f.calls = append(f.calls, "login "+map[bool]string{true: "on", false: "off"}[on])
	return nil
}
func (f *fakeSystem) Eject(v string) error {
	if f.ejectFails > 0 {
		f.ejectFails--
		return errors.New("resource busy")
	}
	f.calls = append(f.calls, "eject "+v)
	return nil
}
func (f *fakeSystem) Trash(p string) error { f.calls = append(f.calls, "trash "+p); return nil }

var quiet = slog.New(slog.NewTextHandler(io.Discard, nil))

func TestMoveInstallsReleasesTheClaimAndRelaunchesWithTheCleanUp(t *testing.T) {
	src := newApp(t, t.TempDir(), "new")
	h := t.TempDir()
	sys := &fakeSystem{}
	released := false
	plan := Plan{Move: true, Source: src, Original: src, Volume: "/Volumes/Forever Sixty", Trash: h + "/Downloads/fs.dmg"}
	o := OfferOptions{
		Plan: plan, Home: h, Log: quiet, System: sys, BeforeRelaunch: func() { released = true },
		Destinations: []string{filepath.Join(h, "Applications")},
	}
	target, err := Move(o, true)
	if err != nil {
		t.Fatal(err)
	}
	if !released {
		t.Error("the one-copy claim was kept across the relaunch")
	}
	want := []string{target, "-installed", "-open-at-login=true", "-eject=/Volumes/Forever Sixty", "-trash-download=" + h + "/Downloads/fs.dmg"}
	if strings.Join(sys.relaunched, "|") != strings.Join(want, "|") {
		t.Errorf("relaunched %q, want %q", sys.relaunched, want)
	}
	if readExe(t, target) != "new" {
		t.Error("the installed copy is not the running one")
	}
}

func TestCompleteDoesWhatTheMoveAsked(t *testing.T) {
	sys := &fakeSystem{loginItems: true, ejectFails: 2}
	f := Finish{Moved: true, OpenAtLogin: true, Eject: "/Volumes/Forever Sixty", Trash: home + "/Downloads/fs.dmg"}
	if err := f.Complete(context.Background(), sys, home, quiet); err != nil {
		t.Fatal(err)
	}
	want := "login on|eject /Volumes/Forever Sixty|trash " + home + "/Downloads/fs.dmg"
	if got := strings.Join(sys.calls, "|"); got != want {
		t.Errorf("calls %q, want %q", got, want)
	}
}

func TestCompleteSkipsTheLoginItemWhereTheSystemHasNone(t *testing.T) {
	sys := &fakeSystem{}
	if err := (Finish{Moved: true, OpenAtLogin: true}).Complete(context.Background(), sys, home, quiet); err != nil {
		t.Fatal(err)
	}
	if len(sys.calls) != 0 {
		t.Errorf("calls %v", sys.calls)
	}
}

func TestCompleteTrashesOnlyADownload(t *testing.T) {
	for _, path := range []string{"/Applications/Forever Sixty.app", home + "/Applications/Forever Sixty.app", home + "/Documents/notes.txt", "/etc/hosts.dmg"} {
		sys := &fakeSystem{}
		err := (Finish{Moved: true, Trash: path}).Complete(context.Background(), sys, home, quiet)
		if !errors.Is(err, ErrNotADownload) || len(sys.calls) != 0 {
			t.Errorf("%s: err %v, calls %v", path, err, sys.calls)
		}
	}
}

func TestOfferServesTheWindowAndMovesOnTheButton(t *testing.T) {
	src := newApp(t, t.TempDir(), "new")
	h := t.TempDir()
	sys := &fakeSystem{loginItems: true}
	closed := make(chan struct{})
	prompt := func(url string, bind func(func()) map[string]any) error {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		page, _ := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode != http.StatusOK || !strings.Contains(string(page), "install.js") {
			t.Errorf("the window's page: %d %.80s", res.StatusCode, page)
		}
		fns := bind(func() { close(closed) })
		if info := fns["fsInstallInfo"].(func() map[string]any)(); info["loginItems"] != true {
			t.Errorf("info = %v", info)
		}
		if _, err := fns["fsInstallMove"].(func(bool) (map[string]any, error))(false); err != nil {
			return err
		}
		<-closed // the binding closes the window once it has answered
		return nil
	}
	moved, err := Offer(OfferOptions{
		Plan: Plan{Move: true, Source: src, Original: src}, Home: h, Log: quiet, System: sys,
		Prompt: prompt, Destinations: []string{filepath.Join(h, "Applications")},
	})
	if err != nil || !moved {
		t.Fatalf("Offer = %v, %v", moved, err)
	}
	if !strings.Contains(strings.Join(sys.relaunched, " "), "-open-at-login=false") {
		t.Errorf("relaunched %v", sys.relaunched)
	}
}

func TestALaunchThatWasNotAMoveDoesNothing(t *testing.T) {
	sys := &fakeSystem{loginItems: true}
	if err := (Finish{Trash: home + "/Downloads/fs.dmg"}).Complete(context.Background(), sys, home, quiet); err != nil || len(sys.calls) != 0 {
		t.Errorf("err %v, calls %v", err, sys.calls)
	}
}
