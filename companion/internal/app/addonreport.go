// companion/internal/app/addonreport.go
// Which Forever Sixty addons the game folders hold, told to the site (PUT
// /v1/devices/self/addons) so /setup can tick its addon step for a player
// whose characters come from Battle.net and whose addon has never sent an
// export. The scan reads the TOCs the addon updater already judges by; a
// report goes up only when what was found changed, or the device signed in
// again, and a failed one is tried again at the next scan.
package app

import (
	"context"
	"slices"
	"sync"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addonupdater"
	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// AddonScanEvery is how often the AddOns folders are read for the report.
const AddonScanEvery = time.Minute

// addonReport is the last scan and what the site was last told, guarded by mu.
type addonReport struct {
	mu       sync.Mutex
	scanned  time.Time
	sentTok  string
	sent     []client.InstalledAddon
	sentOnce bool
}

// ScanAddons lists every folder the releases ship that an install holds, in
// the families' order, each with the newest version any install has.
func ScanAddons(installs []string) []client.InstalledAddon {
	found := []client.InstalledAddon{}
	for _, f := range addonupdater.Families {
		for _, folder := range f.Folders {
			if addon, ok := newestInstalled(f, folder, installs); ok {
				found = append(found, addon)
			}
		}
	}
	return found
}

func newestInstalled(f addonupdater.Family, folder string, installs []string) (client.InstalledAddon, bool) {
	best, seen := "", false
	for _, in := range installs {
		v, ok := addonupdater.ReadVersion(addonupdater.AddonsDir(in), folder)
		if !ok {
			continue
		}
		if !seen || f.Newer(v, best) {
			best = v
		}
		seen = true
	}
	return client.InstalledAddon{Folder: folder, Version: best}, seen
}

// reportAddons scans at most once a minute and tells the site when the
// result differs from what it was last told under this token.
func (a *App) reportAddons(ctx context.Context, now time.Time, tok string) error {
	r := &a.addonReport
	r.mu.Lock()
	defer r.mu.Unlock()
	// A new sign-in reports at once; otherwise the folders are read once a minute.
	if tok == r.sentTok && !r.scanned.IsZero() && now.Sub(r.scanned) < AddonScanEvery {
		return nil
	}
	r.scanned = now
	found := ScanAddons(a.installPaths())
	if r.sentOnce && r.sentTok == tok && slices.Equal(r.sent, found) {
		return nil
	}
	if err := a.client.ReportAddons(ctx, found); err != nil {
		return err
	}
	r.sent, r.sentTok, r.sentOnce = found, tok, true
	return nil
}
