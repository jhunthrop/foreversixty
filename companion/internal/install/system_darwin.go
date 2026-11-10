//go:build darwin

// companion/internal/install/system_darwin.go
// The Mac's half: the system calls behind System, and Detect, which fills in
// the Env that Decide reads.
package install

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework ServiceManagement -framework CoreFoundation
#include <stdbool.h>
#include <stdlib.h>
char *fsOriginalPath(const char *path);
char *fsTrash(const char *path);
bool fsLoginItems(void);
char *fsSetLoginItem(bool on);
*/
import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"unsafe"
)

// Supported reports whether this platform moves itself into Applications.
const Supported = true

type macSystem struct{}

// NewSystem is the Mac's System.
func NewSystem() System { return macSystem{} }

// Detect is the Env of this Mac.
func Detect(home string) Env {
	return Env{Home: home, Original: originalPath, Images: mountedImages}
}

func originalPath(path string) (string, bool) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	out := C.fsOriginalPath(cpath)
	if out == nil {
		return "", false
	}
	defer C.free(unsafe.Pointer(out))
	return C.GoString(out), true
}

func mountedImages() []Image {
	out, err := exec.Command("/usr/bin/hdiutil", "info").Output()
	if err != nil {
		return nil
	}
	return ParseImages(bytes.NewReader(out))
}

func run(name string, args ...string) error {
	out, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w: %s", name, err, strings.TrimSpace(string(out)))
	}
	return nil
}

func (macSystem) Copy(src, dst string) error {
	if err := run("/usr/bin/ditto", src, dst); err != nil {
		return err
	}
	// The player opened this app already, so Gatekeeper has passed it; the
	// copy should not ask again. A copy without the attribute is no error.
	_ = run("/usr/bin/xattr", "-dr", "com.apple.quarantine", dst)
	return nil
}

func (macSystem) Relaunch(app string, args []string) error {
	return run("/usr/bin/open", append([]string{"-n", app, "--args"}, args...)...)
}

func (macSystem) LoginItems() bool { return bool(C.fsLoginItems()) }

func (macSystem) SetLoginItem(on bool) error { return cError(C.fsSetLoginItem(C.bool(on))) }

func (macSystem) Eject(volume string) error {
	return run("/usr/bin/hdiutil", "detach", volume, "-quiet")
}

func (macSystem) Trash(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return cError(C.fsTrash(cpath))
}

func cError(msg *C.char) error {
	if msg == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(msg))
	return errors.New(C.GoString(msg))
}
