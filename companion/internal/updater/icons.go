// companion/internal/updater/icons.go
// After an update swaps the companion in, the system's icon caches still
// hold the old build's icon: Finder, LaunchServices and the Dock on a Mac,
// Explorer, shortcuts and the taskbar on Windows. refreshIcons tells each
// system to look again, so a release that changes the icon shows it at
// once. Every step is best effort: it logs and carries on, and it never
// fails an update that has already been installed.
package updater

import (
	"log/slog"
	"os"
	"time"
)

// notifyIconChange tells the desktop shell that icons may have changed:
// SHChangeNotify(SHCNE_ASSOCCHANGED) on Windows (icons_windows.go), nothing
// elsewhere. A variable so a test can stand in for the system call.
var notifyIconChange = shellNotifyIconChange

// touchBundle sets the bundle folder's modification time to now, as
// `touch` would: LaunchServices and the Dock re-read a bundle's icon when
// the bundle is newer than what they cached.
func touchBundle(bundle string, now time.Time) error {
	return os.Chtimes(bundle, now, now)
}

// refreshIcons runs after a successful swap of the build at exe.
func refreshIcons(exe string, log *slog.Logger, now time.Time) {
	if bundle := BundleRoot(exe); bundle != "" {
		if err := touchBundle(bundle, now); err != nil {
			log.Warn("the update was installed; the Dock may show the old icon until the next login",
				"component", "updater", "err", err.Error())
		}
	}
	if err := notifyIconChange(); err != nil {
		log.Warn("the update was installed; Explorer may show the old icon until it refreshes",
			"component", "updater", "err", err.Error())
	}
}
