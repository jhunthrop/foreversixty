// companion/internal/addonupdater/replace.go
package addonupdater

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// Suffixes of the two sibling folders a swap uses. Both are the family's
// own folder name plus a suffix: nothing else under AddOns is read or
// changed.
const (
	previousSuffix = ".fs-previous"
	incomingSuffix = ".fs-new"
)

// rename and removeAll are the file operations a swap makes; a test
// replaces rename to make one half of a swap fail.
var (
	rename    = os.Rename
	removeAll = os.RemoveAll
)

// ErrDevCopy says a folder is a link or a checkout of the addon's source,
// which belongs to a developer and is never replaced.
var ErrDevCopy = errors.New("the folder is a link or a source checkout, so it is left alone")

// isDevCopy reports a folder that is a symlink or junction, or that holds a
// .git of its own.
func isDevCopy(dir string) bool {
	fi, err := os.Lstat(dir)
	if err != nil {
		return false
	}
	if fi.Mode()&(os.ModeSymlink|os.ModeIrregular) != 0 {
		return true
	}
	_, err = os.Lstat(filepath.Join(dir, ".git"))
	return err == nil
}

// swap is one folder replaced, kept until the whole family is in so a
// failure on the second folder can put the first back.
type swap struct {
	target, previous string
	hadPrevious      bool
}

// replace installs every folder of a family from src (a verified, unpacked
// zip) into addons. Per folder: the new tree is copied beside the old as
// <folder>.fs-new, the old is renamed to <folder>.fs-previous, the new is
// renamed into place. Once every folder is in, the first folder's TOC must
// read back as want; only then are the previous folders deleted. On any
// failure the folders already swapped are put back, and the error says so.
func replace(addons, src string, folders []string, want string, log *slog.Logger) error {
	if err := os.MkdirAll(addons, 0o755); err != nil {
		return err
	}
	for _, folder := range folders {
		if isDevCopy(filepath.Join(addons, folder)) {
			return fmt.Errorf("%s: %w", folder, ErrDevCopy)
		}
	}
	var done []swap
	rollback := func() error {
		var errs []error
		for i := len(done) - 1; i >= 0; i-- {
			s := done[i]
			if err := removeAll(s.target); err != nil {
				errs = append(errs, err)
				continue
			}
			if s.hadPrevious {
				if err := rename(s.previous, s.target); err != nil {
					errs = append(errs, err)
				}
			}
		}
		return errors.Join(errs...)
	}
	fail := func(err error) error {
		if rerr := rollback(); rerr != nil {
			return fmt.Errorf("%w; putting the previous folders back also failed: %w", err, rerr)
		}
		if len(done) > 0 {
			return fmt.Errorf("%w; the previous version was put back", err)
		}
		return err
	}

	for _, folder := range folders {
		target := filepath.Join(addons, folder)
		previous := target + previousSuffix
		incoming := target + incomingSuffix

		// An earlier run that died between its two renames left the old
		// folder aside and nothing in its place: put it back first.
		if _, err := os.Lstat(target); errors.Is(err, os.ErrNotExist) {
			if _, perr := os.Lstat(previous); perr == nil {
				if err := rename(previous, target); err != nil {
					return fail(fmt.Errorf("restore %s from an interrupted update: %w", folder, err))
				}
			}
		}
		if err := removeAll(previous); err != nil {
			return fail(fmt.Errorf("clear %s: %w", previous, err))
		}
		if err := removeAll(incoming); err != nil {
			return fail(fmt.Errorf("clear %s: %w", incoming, err))
		}
		if err := copyTree(filepath.Join(src, folder), incoming); err != nil {
			_ = removeAll(incoming)
			return fail(fmt.Errorf("copy %s: %w", folder, err))
		}
		_, statErr := os.Lstat(target)
		had := statErr == nil
		if had {
			if err := rename(target, previous); err != nil {
				_ = removeAll(incoming)
				return fail(fmt.Errorf("move %s aside: %w", folder, err))
			}
		}
		if err := rename(incoming, target); err != nil {
			_ = removeAll(incoming)
			if had {
				if rerr := rename(previous, target); rerr != nil {
					return fail(fmt.Errorf("move the new %s in: %w; putting the old one back also failed: %w", folder, err, rerr))
				}
			}
			return fail(fmt.Errorf("move the new %s in: %w", folder, err))
		}
		done = append(done, swap{target: target, previous: previous, hadPrevious: had})
	}

	got, found := ReadVersion(addons, folders[0])
	if !found || got != want {
		return fail(fmt.Errorf("the new %s reads back as version %q, not %q", folders[0], got, want))
	}
	for _, s := range done {
		if !s.hadPrevious {
			continue
		}
		if err := removeAll(s.previous); err != nil {
			// The new version is in and working; the leftover is cleared
			// by the next update.
			log.Warn("the previous addon folder could not be deleted", "component", "addonupdater",
				"folder", s.previous, "err", err.Error())
		}
	}
	return nil
}

// leftovers lists the swap siblings a crash could have left in addons, for
// a test to assert there are none.
func leftovers(addons string) []string {
	entries, _ := os.ReadDir(addons)
	var out []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), previousSuffix) || strings.HasSuffix(e.Name(), incomingSuffix) {
			out = append(out, e.Name())
		}
	}
	return out
}
