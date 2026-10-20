// companion/internal/instance/instance.go
// Package instance keeps the companion to one running copy per home
// directory. Two copies would tail the same combat log and upload every
// fight twice, so the second launch finds the first, shows its window and
// exits. The claim is an operating-system lock on a file under the home
// directory: it is released by the system when the holder dies, however
// it dies, so a crash never leaves a stale claim behind.
package instance

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// LockFile and URLFile are the two files, under the home directory.
const (
	LockFile = "companion.lock"
	URLFile  = "ui-url"
)

// ErrRunning says another live process holds the claim.
var ErrRunning = errors.New("the companion is already running")

// Lock is the claim on a home directory.
type Lock struct {
	home string
	f    *os.File
}

// Acquire takes the claim on home. It returns ErrRunning when a live
// process already holds it. A lock file left by a dead process is not a
// claim and is taken over.
func Acquire(home string) (*Lock, error) {
	f, err := os.OpenFile(filepath.Join(home, LockFile), os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return nil, fmt.Errorf("open the instance lock: %w", err)
	}
	held, err := lockFile(f)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("lock the instance lock: %w", err)
	}
	if held {
		f.Close()
		return nil, ErrRunning
	}
	// The pid is for a person reading the file; the lock is the claim.
	if err := f.Truncate(0); err == nil {
		_, _ = f.WriteAt([]byte(strconv.Itoa(os.Getpid())+"\n"), 0)
	}
	return &Lock{home: home, f: f}, nil
}

// PublishURL records the address of the running instance's UI, token
// included, where only the user can read it, for a second launch to
// open.
func (l *Lock) PublishURL(url string) error {
	return writePrivate(filepath.Join(l.home, URLFile), url+"\n")
}

// Release gives the claim up and deletes the published address.
func (l *Lock) Release() error {
	_ = os.Remove(filepath.Join(l.home, URLFile))
	unlockFile(l.f)
	return l.f.Close()
}

// RunningURL is the address the running instance published, empty when
// it has not published one yet.
func RunningURL(home string) string {
	b, err := os.ReadFile(filepath.Join(home, URLFile))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(b))
}

func writePrivate(path, body string) error {
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, []byte(body), 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
