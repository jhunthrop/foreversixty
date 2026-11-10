package install

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeWin struct {
	calls   []string
	entry   Entry
	failOn  string
	started string
	startup bool
}

func (f *fakeWin) note(call string) error {
	f.calls = append(f.calls, call)
	if f.failOn != "" && strings.HasPrefix(call, f.failOn) {
		return errors.New(f.failOn + " failed")
	}
	return nil
}

func (f *fakeWin) StopCompanions() error { return f.note("stop") }
func (f *fakeWin) CopyExe(src, dst string) (int64, error) {
	return 16 << 20, f.note("copy " + src + " -> " + dst)
}
func (f *fakeWin) CreateShortcut(lnk, target string) error {
	return f.note("shortcut " + lnk + " -> " + target)
}
func (f *fakeWin) SetUninstallEntry(e Entry) error  { f.entry = e; return f.note("entry") }
func (f *fakeWin) DeleteUninstallEntry() error      { return f.note("delete entry") }
func (f *fakeWin) SetStartup(command string) error  { return f.note("startup " + command) }
func (f *fakeWin) DeleteStartup() error             { return f.note("no startup") }
func (f *fakeWin) StartupEnabled() bool             { return f.startup }
func (f *fakeWin) Start(exe string) error           { f.started = exe; return f.note("start " + exe) }
func (f *fakeWin) Remove(path string) error         { return f.note("remove " + path) }
func (f *fakeWin) RemoveAfterExit(dir string) error { return f.note("remove after exit " + dir) }
func (f *fakeWin) ForgetToken() error               { return f.note("forget token") }

var folders = Folders{
	Install:   filepath.Join("C:", "Users", "thrall", "AppData", "Local", "Programs", "Forever Sixty"),
	StartMenu: filepath.Join("C:", "Users", "thrall", "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs"),
	Data:      filepath.Join("C:", "Users", "thrall", "AppData", "Roaming", "ForeverSixty"),
}

func TestSetupMode(t *testing.T) {
	for exe, want := range map[string]bool{
		filepath.Join("C:", "Users", "thrall", "Downloads", "ForeverSixtySetup.exe"):              true,
		filepath.Join("C:", "Users", "thrall", "Downloads", "ForeverSixtySetup (1).exe"):          true,
		filepath.Join("C:", "Users", "thrall", "Downloads", "foreversixty-companion-setup.exe"):   true,
		filepath.Join("C:", "Users", "thrall", "Downloads", "foreversixty-companion_windows.exe"): false,
		filepath.Join(folders.Install, ExeName):                                                   false,
		filepath.Join(strings.ToUpper(folders.Install), "ForeverSixtySetup.exe"):                  false,
	} {
		if got := SetupMode(exe, folders.Install); got != want {
			t.Errorf("SetupMode(%q) = %v, want %v", exe, got, want)
		}
	}
}

// An update keeps the player's Start with Windows choice: the window starts from it.
func TestTheSetupWindowReportsAnInstallAndItsStartupChoice(t *testing.T) {
	install := t.TempDir()
	if err := os.WriteFile(filepath.Join(install, ExeName), []byte("old"), 0o755); err != nil {
		t.Fatal(err)
	}
	prompt := func(_ string, bind func(Window) map[string]any) error {
		info := bind(fakeWindow{})["fsSetupInfo"].(func() map[string]any)()
		if info["upgrade"] != true || info["startup"] != false {
			t.Errorf("info %v", info)
		}
		return nil
	}
	err := ShowSetup(SetupOptions{
		Mode: ModeSetup, System: &fakeWin{startup: false}, Folders: Folders{Install: install},
		Log: quiet, Prompt: prompt,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestInstallDirIsTheAccountsOwnPrograms(t *testing.T) {
	if got := InstallDir(filepath.Join("C:", "L")); got != filepath.Join("C:", "L", "Programs", "Forever Sixty") {
		t.Errorf("InstallDir = %q", got)
	}
}

func TestTheSetupInstallsRegistersAndStarts(t *testing.T) {
	for _, startup := range []bool{true, false} {
		t.Run(fmt.Sprint("startup ", startup), func(t *testing.T) {
			sys := &fakeWin{}
			var steps []string
			self := filepath.Join("C:", "Downloads", "ForeverSixtySetup.exe")
			err := RunSetup(sys, folders, self, "0.2.46", startup, func(step, of int, key string) {
				steps = append(steps, fmt.Sprintf("%d/%d %s", step, of, key))
			}, quiet)
			if err != nil {
				t.Fatal(err)
			}
			exe := folders.Exe()
			startCall := "no startup"
			if startup {
				startCall = `startup "` + exe + `"`
			}
			want := []string{
				"stop", "copy " + self + " -> " + exe,
				"shortcut " + filepath.Join(folders.StartMenu, "Forever Sixty.lnk") + " -> " + exe,
				"entry", startCall, "start " + exe,
			}
			if strings.Join(sys.calls, "|") != strings.Join(want, "|") {
				t.Errorf("calls\n %q\nwant\n %q", sys.calls, want)
			}
			if got := strings.Join(steps, ","); got != "1/5 stop,2/5 copy,3/5 shortcut,4/5 register,5/5 start" {
				t.Errorf("steps %s", got)
			}
		})
	}
}

func TestTheAppsAndFeaturesEntry(t *testing.T) {
	exe := folders.Exe()
	e := NewEntry(exe, "0.2.46", 16<<20+1)
	want := Entry{
		DisplayName: "Forever Sixty", DisplayIcon: exe, DisplayVersion: "0.2.46", Publisher: "Commish LLC",
		InstallLocation: folders.Install, UninstallString: `"` + exe + `" -uninstall`,
		URLInfoAbout: "https://foreversixty.gg", EstimatedSize: 16*1024 + 1,
	}
	if e != want {
		t.Errorf("entry\n %+v\nwant\n %+v", e, want)
	}
}

func TestAFailedStepStopsTheSetup(t *testing.T) {
	sys := &fakeWin{failOn: "shortcut"}
	err := RunSetup(sys, folders, "setup.exe", "0.2.46", true, nil, quiet)
	if err == nil || !strings.Contains(err.Error(), "Start menu shortcut") {
		t.Fatalf("err = %v", err)
	}
	if sys.started != "" {
		t.Error("a half-done install was started")
	}
}

func TestTheUninstallRemovesEverythingAndKeepsSettingsWhenAsked(t *testing.T) {
	for _, keep := range []bool{true, false} {
		t.Run(fmt.Sprint("keep ", keep), func(t *testing.T) {
			sys := &fakeWin{}
			if err := RunUninstall(sys, folders, keep, nil, quiet); err != nil {
				t.Fatal(err)
			}
			want := []string{"stop", "no startup", "remove " + folders.Shortcut(), "delete entry"}
			if !keep {
				want = append(want, "remove "+folders.Data, "forget token")
			}
			want = append(want, "remove after exit "+folders.Install)
			if strings.Join(sys.calls, "|") != strings.Join(want, "|") {
				t.Errorf("calls\n %q\nwant\n %q", sys.calls, want)
			}
		})
	}
}

func TestTheSetupWindowRunsTheStepsAndReportsThem(t *testing.T) {
	sys := &fakeWin{}
	var evals []string
	closed := make(chan struct{})
	prompt := func(url string, bind func(Window) map[string]any) error {
		if !strings.HasSuffix(url, "install.html?mode=setup") {
			t.Errorf("url %s", url)
		}
		fns := bind(fakeWindow{closed: closed, evals: &evals})
		if info := fns["fsSetupInfo"].(func() map[string]any)(); info["upgrade"] != false || info["startup"] != false {
			t.Errorf("info %v", info)
		}
		fns["fsSetupRun"].(func(bool))(true)
		<-closed
		return nil
	}
	err := ShowSetup(SetupOptions{
		Mode: ModeSetup, System: sys, Folders: Folders{Install: t.TempDir(), StartMenu: "S", Data: "D"},
		Self: "setup.exe", Version: "0.2.46", Log: quiet, Prompt: prompt,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(evals) != 6 || evals[0] != `window.fsProgress(1, 5, "stop")` || evals[4] != `window.fsProgress(5, 5, "start")` || evals[5] != "window.fsDone()" {
		t.Errorf("evals %q", evals)
	}
}
