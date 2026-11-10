//go:build windows

// companion/internal/install/setup_windows.go
// The Windows system behind RunSetup and RunUninstall: the files, the
// registry (the startup entry and the Apps & Features entry, both under
// HKEY_CURRENT_USER), the Start menu shortcut through the shell's own
// IShellLink, and the processes. Nothing here opens a console window.
package install

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"

	"github.com/jhunthrop/foreversixty/companion/internal/paths"
	"github.com/jhunthrop/foreversixty/companion/internal/secret"
	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

// RunKey is the startup list, and the value the companion is under there.
const (
	RunKey   = `Software\Microsoft\Windows\CurrentVersion\Run`
	RunValue = ProductName
)

// SetupSupported reports whether this platform has the setup.
const SetupSupported = true

// How long a copy waits for a stopped companion to let go of its exe.
const (
	copyTries = 30
	copyEvery = 100 * time.Millisecond
)

type winSystem struct{}

// NewWinSystem is the real Windows system.
func NewWinSystem() WinSystem { return winSystem{} }

// WindowsFolders are this account's folders.
func WindowsFolders() (Folders, error) {
	local, err := windows.KnownFolderPath(windows.FOLDERID_LocalAppData, 0)
	if err != nil {
		return Folders{}, fmt.Errorf("find the local app data folder: %w", err)
	}
	programs, err := windows.KnownFolderPath(windows.FOLDERID_Programs, 0)
	if err != nil {
		return Folders{}, fmt.Errorf("find the Start menu folder: %w", err)
	}
	data, err := paths.Home()
	if err != nil {
		return Folders{}, err
	}
	return Folders{Install: InstallDir(local), StartMenu: programs, Data: data}, nil
}

// hidden runs a console tool without a window.
func hidden(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true, CreationFlags: windows.CREATE_NO_WINDOW}
	return cmd
}

func (winSystem) StopCompanions() error {
	// taskkill exits 128 when nothing matched: no companion was running.
	err := hidden(exec.Command("taskkill", "/F", "/IM", ExeName, "/FI", "PID ne "+strconv.Itoa(os.Getpid()))).Run()
	var exit *exec.ExitError
	if errors.As(err, &exit) && exit.ExitCode() == 128 {
		return nil
	}
	return err
}

func (winSystem) CopyExe(src, dst string) (int64, error) {
	body, err := os.ReadFile(src)
	if err != nil {
		return 0, err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return 0, err
	}
	staged := dst + ".new"
	if err := os.WriteFile(staged, body, 0o755); err != nil {
		return 0, err
	}
	for range copyTries {
		if err = os.Rename(staged, dst); err == nil {
			return int64(len(body)), nil
		}
		time.Sleep(copyEvery) // a companion just stopped can hold its exe a moment longer
	}
	os.Remove(staged)
	return 0, err
}

func (winSystem) SetUninstallEntry(e Entry) error {
	k, _, err := registry.CreateKey(registry.CURRENT_USER, updater.UninstallKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()
	for name, value := range map[string]string{
		"DisplayName": e.DisplayName, "DisplayIcon": e.DisplayIcon, "DisplayVersion": e.DisplayVersion,
		"Publisher": e.Publisher, "InstallLocation": e.InstallLocation,
		"UninstallString": e.UninstallString, "URLInfoAbout": e.URLInfoAbout,
	} {
		if err := k.SetStringValue(name, value); err != nil {
			return err
		}
	}
	for name, value := range map[string]uint32{"EstimatedSize": e.EstimatedSize, "NoModify": 1, "NoRepair": 1} {
		if err := k.SetDWordValue(name, value); err != nil {
			return err
		}
	}
	return nil
}

func (winSystem) DeleteUninstallEntry() error {
	err := registry.DeleteKey(registry.CURRENT_USER, updater.UninstallKey)
	if errors.Is(err, registry.ErrNotExist) {
		return nil
	}
	return err
}

func (winSystem) SetStartup(command string) error {
	k, _, err := registry.CreateKey(registry.CURRENT_USER, RunKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()
	return k.SetStringValue(RunValue, command)
}

func (winSystem) DeleteStartup() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, RunKey, registry.SET_VALUE)
	if errors.Is(err, registry.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer k.Close()
	if err := k.DeleteValue(RunValue); err != nil && !errors.Is(err, registry.ErrNotExist) {
		return err
	}
	return nil
}

func (winSystem) Start(exe string) error {
	cmd := exec.Command(exe)
	cmd.Dir = filepath.Dir(exe)
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Process.Release()
}

func (winSystem) Remove(path string) error { return os.RemoveAll(path) }

// RemoveAfterExit leaves a hidden cmd behind that waits for this process to
// end (a few seconds of ping), then deletes the folder, this exe with it.
func (winSystem) RemoveAfterExit(dir string) error {
	cmd := exec.Command(filepath.Join(os.Getenv("SystemRoot"), "System32", "cmd.exe"))
	cmd.Dir = os.TempDir() // a process's working folder cannot be deleted
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: windows.CREATE_NO_WINDOW, // not DETACHED_PROCESS: ping would then open a console of its own
		CmdLine:       `cmd.exe /c ping -n 4 127.0.0.1 >nul & rmdir /s /q "` + dir + `"`,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	return cmd.Process.Release()
}

func (winSystem) ForgetToken() error { return secret.Keyring{}.Clear() }

// The shell's link object (ShlObj_core.h): CLSID_ShellLink, IShellLinkW and IPersistFile.
var (
	clsidShellLink  = windows.GUID{Data1: 0x00021401, Data4: [8]byte{0xC0, 0, 0, 0, 0, 0, 0, 0x46}}
	iidIShellLinkW  = windows.GUID{Data1: 0x000214F9, Data4: [8]byte{0xC0, 0, 0, 0, 0, 0, 0, 0x46}}
	iidIPersistFile = windows.GUID{Data1: 0x0000010B, Data4: [8]byte{0xC0, 0, 0, 0, 0, 0, 0, 0x46}}

	ole32            = windows.NewLazySystemDLL("ole32.dll")
	coCreateInstance = ole32.NewProc("CoCreateInstance")
)

// The vtable slots used: IUnknown's, then IShellLinkW's SetDescription,
// SetWorkingDirectory, SetIconLocation and SetPath, and IPersistFile's Save.
const (
	slotQueryInterface   = 0
	slotRelease          = 2
	slotSetDescription   = 7
	slotSetWorkingDir    = 9
	slotSetIconLocation  = 17
	slotSetPath          = 20
	slotPersistFileSave  = 6
	clsctxInprocServer   = 1
	coinitApartmentThred = 0x2
)

// comObject is an interface pointer: its first word points at the vtable.
type comObject struct{ ptr unsafe.Pointer }

// vtableSlots bounds the vtable view; IShellLinkW, the longest used, has 21.
const vtableSlots = 32

func (o comObject) call(slot int, args ...uintptr) error {
	vtbl := (*[vtableSlots]uintptr)(*(*unsafe.Pointer)(o.ptr))
	hr, _, _ := syscall.SyscallN(vtbl[slot], append([]uintptr{uintptr(o.ptr)}, args...)...)
	if int32(hr) < 0 {
		return fmt.Errorf("COM call %d: HRESULT 0x%08X", slot, uint32(hr))
	}
	return nil
}

func (o comObject) release() { _ = o.call(slotRelease) }

// utf16 holds Windows strings for the COM calls and keeps them alive until
// the calls are done (the collector never moves them, but must not free them).
type utf16 struct{ kept []*uint16 }

func (u *utf16) of(s string) uintptr {
	p, _ := windows.UTF16PtrFromString(s)
	u.kept = append(u.kept, p)
	return uintptr(unsafe.Pointer(p))
}

func (winSystem) CreateShortcut(lnk, target string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	// S_FALSE (1): COM was already started on this thread, which still wants its uninitialize.
	if err := windows.CoInitializeEx(0, coinitApartmentThred); err != nil && err != syscall.Errno(windows.S_FALSE) {
		return fmt.Errorf("start COM: %w", err)
	}
	defer windows.CoUninitialize()
	var link comObject
	hr, _, _ := coCreateInstance.Call(uintptr(unsafe.Pointer(&clsidShellLink)), 0, clsctxInprocServer,
		uintptr(unsafe.Pointer(&iidIShellLinkW)), uintptr(unsafe.Pointer(&link.ptr)))
	if int32(hr) < 0 {
		return fmt.Errorf("create the shell link: HRESULT 0x%08X", uint32(hr))
	}
	defer link.release()
	var s utf16
	defer runtime.KeepAlive(&s)
	for _, step := range []struct {
		slot int
		args []uintptr
	}{
		{slotSetPath, []uintptr{s.of(target)}},
		{slotSetWorkingDir, []uintptr{s.of(filepath.Dir(target))}},
		{slotSetDescription, []uintptr{s.of(ProductName)}},
		{slotSetIconLocation, []uintptr{s.of(target), 0}},
	} {
		if err := link.call(step.slot, step.args...); err != nil {
			return err
		}
	}
	var file comObject
	if err := link.call(slotQueryInterface, uintptr(unsafe.Pointer(&iidIPersistFile)), uintptr(unsafe.Pointer(&file.ptr))); err != nil {
		return err
	}
	defer file.release()
	if err := os.MkdirAll(filepath.Dir(lnk), 0o755); err != nil {
		return err
	}
	return file.call(slotPersistFileSave, s.of(lnk), 1)
}
