// companion/internal/app/folders.go
// The game folders the player adds and removes one at a time, through
// the system's folder dialog. Only what the player added is stored in
// wow_paths: the folders found by looking are found again on every
// launch, so a frozen copy of them in config.json would go stale.
package app

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/jhunthrop/foreversixty/companion/internal/picker"
	"github.com/jhunthrop/foreversixty/companion/internal/wow"
)

// PickFolder shows the folder dialog and adds the choice. A dialog the
// player closes is not an error: nothing changes. The folder may be a
// flavour folder or the folder above them; every install found in it
// is stored by its own path. A Fault says what was wrong with the
// choice.
func (a *App) PickFolder() error {
	dir, err := a.o.PickFolder()
	switch {
	case errors.Is(err, picker.ErrCancelled):
		return nil
	case errors.Is(err, picker.ErrUnavailable):
		return fault(CodeNoPicker, err)
	case err != nil:
		a.log.Warn("the folder dialog failed", "component", "app", "err", err.Error())
		return fault(CodeNoPicker, err)
	}
	return a.AddFolder(dir)
}

// AddFolder validates a folder with wow.Pick and appends the installs
// in it to wow_paths. A folder already in the list is refused with a
// sentence, not silently ignored.
func (a *App) AddFolder(dir string) error {
	found, err := wow.Pick(dir)
	if err != nil {
		return fault(CodeNotWoW, err)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	cfg := a.cfg
	cfg.WoWPaths = slices.Clone(cfg.WoWPaths)
	known := map[string]bool{}
	for _, in := range a.installs {
		known[in.Path] = true
	}
	added := 0
	for _, in := range found {
		if known[in.Path] {
			continue
		}
		cfg.WoWPaths = append(cfg.WoWPaths, in.Path)
		added++
	}
	if added == 0 {
		return fault(CodeKnownFolder, nil)
	}
	if err := a.persist(cfg); err != nil {
		return err
	}
	a.log.Info("a game folder was added", "component", "app", "folder", dir, "installs", added)
	return nil
}

// RemoveFolder removes a folder the player added. A found folder has
// nothing to remove, and an unknown one is refused the same way.
func (a *App) RemoveFolder(path string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	cfg := a.cfg
	i := slices.Index(cfg.WoWPaths, path)
	if i < 0 {
		return fault(CodeNotAdded, nil)
	}
	cfg.WoWPaths = slices.Delete(slices.Clone(cfg.WoWPaths), i, i+1)
	if err := a.persist(cfg); err != nil {
		return err
	}
	a.log.Info("a game folder was removed", "component", "app", "folder", path)
	return nil
}

// trayNoticeFile marks that the player has been told, once, that closing
// the window leaves the companion running in the tray.
const trayNoticeFile = "tray-notice-seen"

// FirstClose reports whether this is the first time the window has been
// closed, and remembers that it was, in a file beside config.json. A
// marker that cannot be written answers false after the first call in a
// process rather than nagging on every close.
func (a *App) FirstClose() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.noticeShown {
		return false
	}
	a.noticeShown = true
	f, err := os.OpenFile(filepath.Join(a.o.Dirs.Home, trayNoticeFile), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		return errors.Is(err, fs.ErrNotExist) // no home yet: nothing recorded, say it once
	}
	_ = f.Close()
	return true
}
