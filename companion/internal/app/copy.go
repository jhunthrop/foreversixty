// companion/internal/app/copy.go
// The words the status page shows for the addon sync, in one place
// (spec §4.10 and §5.2): when the site last saw the player's export,
// and what became of the newest in-game sim request. Times are the
// player's local time. Tests read these constants rather than repeat
// the words.
package app

import (
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// The addon line's words.
const (
	CopySeenBySite      = "Seen by the site"
	CopySimQueued       = "Sim queued"
	CopySimNeedsPremium = "Sim needs premium"
	CopySimDuplicate    = "Sim already running"
	CopySimInvalid      = "Sim request invalid"
	// CopySeparator joins the two halves of the line.
	CopySeparator = " · "
)

// simCopy is the words for each sim request state the ack can carry.
var simCopy = map[string]string{
	addon.SimQueued:       CopySimQueued,
	addon.SimNeedsPremium: CopySimNeedsPremium,
	addon.SimDuplicate:    CopySimDuplicate,
	addon.SimInvalid:      CopySimInvalid,
}

// AddonLine is the status page's addon line: "Seen by the site 14:03"
// when the site saw the export today, "Seen by the site 4 Nov 14:03"
// on another day, followed by the sim request's state when there is
// one. "Today" and the clock are the player's, in now's location. It
// is empty until the site has acked an export: with no ack there is
// nothing to say, and the line is absent.
func AddonLine(seen time.Time, simState string, now time.Time) string {
	if seen.IsZero() {
		return ""
	}
	local := seen.In(now.Location())
	layout := "15:04"
	if y, m, d := local.Date(); y != now.Year() || m != now.Month() || d != now.Day() {
		layout = "2 Jan 15:04"
	}
	line := CopySeenBySite + " " + local.Format(layout)
	if words := simCopy[simState]; words != "" {
		line += CopySeparator + words
	}
	return line
}
