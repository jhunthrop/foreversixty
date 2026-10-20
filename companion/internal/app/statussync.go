// companion/internal/app/statussync.go
// Two things the Status page says that the log used to keep to itself:
// what the last character sync put on the site, and who the current
// report is credited to.
package app

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/state"
)

// The states of the Last sync row.
const (
	// SyncSent: the site holds Count characters, sent At.
	SyncSent = "sent"
	// SyncNone: the addon is installed and nothing has gone up yet.
	SyncNone = "none"
	// SyncRefused: the site refused the last list; Reason is its message.
	SyncRefused = "refused"
	// SyncAbsent: no addon folder, so the Addon row's install line says it all.
	SyncAbsent = "absent"
)

// LastSyncView is the Last sync row. The window words the sent state
// itself (count and a relative time) so the time keeps moving between
// snapshots; the other lines are the Go side's (copy.go).
type LastSyncView struct {
	State string `json:"state"`
	Count int    `json:"count,omitempty"`
	// At is when the list went up, in local time; Title is the same
	// as a tooltip.
	At    *time.Time `json:"at,omitempty"`
	Title string     `json:"title,omitempty"`
	// Line is the whole sentence for the none and refused states.
	Line string `json:"line,omitempty"`
}

// The states of a report's credit.
const (
	CreditCredited  = "credited"
	CreditUnmatched = "unmatched"
	CreditNone      = "none"
)

// Credit is who a report is credited to. The window reads two states and
// no reason: whether the addon lacks the name or only the realm is the
// resolver's to log (logger.Reason, in companion.log), because the player's
// step is the same either way. Name is in its space form.
type Credit struct {
	Name  string `json:"name,omitempty"`
	State string `json:"state"`
}

// creditOf reads the credit off a report: the character it carries
// (detected or chosen), else the name the log gave that no addon record
// matched, else nothing.
func creditOf(r state.Report) Credit {
	switch {
	case r.LoggingCharacter != nil && r.LoggingCharacter.Name != "":
		return Credit{Name: SpacedName(r.LoggingCharacter.Name), State: CreditCredited}
	case r.Auto && r.DetectedName != "":
		return Credit{Name: SpacedName(r.DetectedName), State: CreditUnmatched}
	}
	return Credit{State: CreditNone}
}

// currentCredit is the credit of the report the pipeline has open, none
// when it has none or the report cannot be read. The tray's headline asks
// on a timer, so it loads the one report rather than listing them all.
func (a *App) currentCredit(key string) Credit {
	if key == "" {
		return Credit{State: CreditNone}
	}
	r, err := state.Load(a.o.Dirs.State, key)
	if err != nil {
		return Credit{State: CreditNone}
	}
	return creditOf(r)
}

// addonInstalled reports whether any install has the addon's folder or
// a saved-variables file for it.
func (a *App) addonInstalled() bool {
	a.mu.Lock()
	installs := a.installs
	a.mu.Unlock()
	for _, in := range installs {
		if st, err := os.Stat(filepath.Join(in.Path, "Interface", "AddOns", addon.SavedVariablesName)); err == nil && st.IsDir() {
			return true
		}
		if paths, err := in.SavedVariables(addon.SavedVariablesName); err == nil {
			for _, p := range paths {
				if _, err := os.Stat(p); err == nil {
					return true
				}
			}
		}
	}
	return false
}

// lastSyncView is the Last sync row for the sync's status.
func (a *App) lastSyncView(st addon.Status, now time.Time) LastSyncView {
	switch {
	case st.Refusal != "" && st.RefusedAt.After(st.SentAt):
		return LastSyncView{State: SyncRefused, Line: CopySiteRefused(st.Refusal)}
	case !st.SentAt.IsZero():
		local := st.SentAt.In(now.Location())
		return LastSyncView{State: SyncSent, Count: st.SentCount, At: &local,
			Title: local.Format("Mon 2 Jan 2006, 15:04")}
	case !a.addonInstalled():
		return LastSyncView{State: SyncAbsent}
	}
	return LastSyncView{State: SyncNone, Line: CopyNoCharactersSent}
}

// BuildsChecked is the Addon row's second line after the Check for builds
// now button: when it ran, how many builds it found that the inbox did
// not hold before, and for whom. The window words it, so the time keeps
// moving between snapshots.
type BuildsChecked struct {
	At  time.Time `json:"at"`
	New int       `json:"new_builds"`
	For []string  `json:"for,omitempty"`
}

// CheckBuilds fetches the inbox once, at the player's asking, and writes
// it as soon as the game folder is quiet. It is the Check for builds now
// button. A failure is a Fault the window words (offline, or the site's
// problem); an unpaired device has nothing to ask.
func (a *App) CheckBuilds(ctx context.Context) error {
	a.mu.Lock()
	signedOut := a.signedOut != nil
	a.mu.Unlock()
	if !a.paired(signedOut) {
		return fault(CodeNotPaired, nil)
	}
	tok := a.deviceToken()
	found, err := a.sync.CheckNow(ctx, a.now())
	if err != nil {
		if a.signOutIfRefused(err, tok) {
			return fault(CodeNotPaired, err)
		}
		a.log.Warn("checking for builds failed", "component", "app", "err", err.Error())
		return fault(netCode(err, CodeServer), err)
	}
	a.log.Info("checked for builds", "component", "app", "new", found.New)
	return nil
}
