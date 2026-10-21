// companion/internal/app/version.go
// Knowing that this launch is the first of a new version, so the footer
// can say "Updated to 0.2.12" for a little while. Updates apply on the
// next launch, so nothing is clicked; this is the only trace the player
// gets that one happened.
package app

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/jhunthrop/foreversixty/companion/internal/updater"
)

// lastVersionFile records, beside config.json, the version that last ran.
const lastVersionFile = "last-version"

// noteVersion compares the running version with the one recorded and
// records this one. A home that has never run a companion (nothing
// recorded, and the tray notice was never shown either) is a first
// install, not an update. The tray notice stands in for "an older
// companion has run here": it was the one marker releases before this
// file left behind, so the update that introduces this file is still told.
func (a *App) noteVersion() {
	cur := updater.Version
	p := filepath.Join(a.o.Dirs.Home, lastVersionFile)
	b, err := os.ReadFile(p)
	last := strings.TrimSpace(string(b))
	switch {
	case err == nil && last == cur:
		return
	case err == nil && last != "":
		a.updatedTo = cur
	case errors.Is(err, fs.ErrNotExist):
		if _, seen := os.Stat(filepath.Join(a.o.Dirs.Home, trayNoticeFile)); seen == nil {
			a.updatedTo = cur
		}
	}
	// A record that cannot be written costs the next launch its line,
	// never this one its window.
	if err := os.WriteFile(p, []byte(cur+"\n"), 0o600); err != nil {
		a.log.Warn("could not record the running version", "component", "app", "err", err.Error())
	}
}
