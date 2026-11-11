// companion/internal/install/install.go
// Package install puts the Mac companion where it belongs on its first
// launch: a player opens the download (the zip's app in Downloads, or the
// app on the mounted disk image) and the companion offers to move itself to
// Applications, replaces an older copy there, starts again from the new
// place, ejects the disk image and puts the download in the Trash. It is the
// LetsMove pattern, in Go with a few lines of Objective-C (install_darwin.m).
//
// The decision is pure and runs on every platform (Decide, the tests); the
// moves and the system calls are the darwin build's. On Windows the setup
// program installs the companion and none of this runs.
package install

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// AppName is the bundle's name in Applications, whatever the download was
// called: a browser that saved a second copy as "Forever Sixty 2.app" still
// installs "Forever Sixty.app".
const AppName = "Forever Sixty.app"

// SystemApplications is the Applications folder every account shares.
const SystemApplications = "/Applications"

// Plan is what a launch should do about where the app sits.
type Plan struct {
	// Move is true when the app runs from anywhere but an Applications folder.
	Move bool
	// Source is the bundle to copy: the running one, which App Translocation
	// may have mounted somewhere random and read-only.
	Source string
	// Original is where the player's copy really is (the translocated
	// bundle's own path), the same as Source when nothing was translocated.
	Original string
	// Volume is the mounted disk image the app runs from, to eject after the
	// move; empty when it does not run from one.
	Volume string
	// Trash is the downloaded file to put in the Trash after the move: the
	// disk image, or the app the zip unpacked. Empty when it is nothing the
	// player downloaded into their own folders.
	Trash string
}

// Image is one mounted disk image: its file and where it is mounted.
type Image struct {
	Path       string
	MountPoint string
}

// Env is what Decide reads from the system; the darwin build fills it in
// (System) and the tests make it up.
type Env struct {
	// Home is the player's home folder.
	Home string
	// Original returns the path a translocated bundle was opened from, and
	// false when the path is not translocated.
	Original func(path string) (string, bool)
	// Images lists the mounted disk images.
	Images func() []Image
}

// Decide says whether the app running as exe should move, and what to clean
// up after. An executable outside an app bundle (a developer's build, the
// headless binary) never moves.
func Decide(bundle string, env Env) Plan {
	if bundle == "" {
		return Plan{}
	}
	original := bundle
	if env.Original != nil {
		if o, ok := env.Original(bundle); ok && o != "" {
			original = o
		}
	}
	if InApplications(original, env.Home) {
		return Plan{}
	}
	plan := Plan{Move: true, Source: bundle, Original: original}
	if env.Images != nil {
		for _, img := range env.Images() {
			if img.MountPoint != "" && within(original, img.MountPoint) {
				plan.Volume = img.MountPoint
				if within(img.Path, env.Home) {
					plan.Trash = img.Path
				}
				return plan
			}
		}
	}
	if within(original, env.Home) {
		plan.Trash = original
	}
	return plan
}

// InApplications reports whether the bundle sits directly in the shared or
// the player's own Applications folder.
func InApplications(bundle, home string) bool {
	parent := filepath.Dir(filepath.Clean(bundle))
	return parent == SystemApplications || (home != "" && parent == filepath.Join(home, "Applications"))
}

// Destinations are the folders the app is installed into, in order: the
// shared Applications folder, and the player's own when that one is not
// writable (a standard account).
func Destinations(home string) []string {
	return []string{SystemApplications, filepath.Join(home, "Applications")}
}

// within reports whether path is dir or under it.
func within(path, dir string) bool {
	if dir == "" {
		return false
	}
	path, dir = filepath.Clean(path), filepath.Clean(dir)
	return path == dir || strings.HasPrefix(path, dir+string(filepath.Separator))
}

// ParseImages reads `hdiutil info`: each image's image-path line, then its
// device lines, the mounted one ending in its mount point after a tab.
func ParseImages(r io.Reader) []Image {
	var (
		out []Image
		cur string
	)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if rest, ok := strings.CutPrefix(line, "image-path"); ok {
			if _, path, ok := strings.Cut(rest, ":"); ok {
				cur = strings.TrimSpace(path)
			}
			continue
		}
		if !strings.HasPrefix(line, "/dev/") || cur == "" {
			continue
		}
		fields := strings.Split(line, "\t")
		if mount := strings.TrimSpace(fields[len(fields)-1]); len(fields) >= 3 && strings.HasPrefix(mount, "/") {
			out = append(out, Image{Path: cur, MountPoint: mount})
		}
	}
	return out
}

// Copier copies the bundle at src to dst, which does not exist yet. The
// darwin build uses ditto, which keeps the signature and the stapled ticket.
type Copier func(src, dst string) error

// ErrNoDestination says no Applications folder took the app.
var ErrNoDestination = errors.New("install: no Applications folder could take the app")

// Install copies the bundle at src into the first destination folder that
// takes it, as AppName, replacing an older copy there, and returns the new
// bundle's path. A folder that cannot be written (or created, for the
// player's own) is skipped for the next.
func Install(src string, dests []string, copyBundle Copier) (string, error) {
	var errs []error
	for _, dir := range dests {
		target, err := installInto(src, dir, copyBundle)
		if err == nil {
			return target, nil
		}
		errs = append(errs, fmt.Errorf("%s: %w", dir, err))
	}
	return "", errors.Join(append([]error{ErrNoDestination}, errs...)...)
}

// installInto stages the copy beside the target, so the swap is two renames
// on one volume, and keeps the older copy until the new one is in place.
func installInto(src, dir string, copyBundle Copier) (string, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	stage, err := os.MkdirTemp(dir, "."+AppName+".install-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(stage)
	fresh := filepath.Join(stage, AppName)
	if err := copyBundle(src, fresh); err != nil {
		return "", fmt.Errorf("copy the app: %w", err)
	}
	target := filepath.Join(dir, AppName)
	previous := filepath.Join(stage, "previous.app")
	hadPrevious := false
	if _, err := os.Lstat(target); err == nil {
		if err := os.Rename(target, previous); err != nil {
			return "", fmt.Errorf("move the older copy aside: %w", err)
		}
		hadPrevious = true
	}
	if err := os.Rename(fresh, target); err != nil {
		if hadPrevious {
			if restoreErr := os.Rename(previous, target); restoreErr != nil {
				return "", fmt.Errorf("put the new copy in place: %w; restoring the older copy also failed: %w", err, restoreErr)
			}
		}
		return "", fmt.Errorf("put the new copy in place: %w", err)
	}
	return target, nil
}
