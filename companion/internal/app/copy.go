// companion/internal/app/copy.go
// The player-facing sentences the Go side owns, in one place: the
// headline words (which the tray repeats exactly), the addon line, and
// the message behind every failure the window can show. The window's
// own words are in ui/copy.js. Nothing here ever carries an error's
// text: a failure's cause goes to companion.log and the player gets one
// of these. Times are the player's local time. Tests read these
// constants rather than repeat the words.
package app

import (
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
)

// The addon line's words.
const (
	CopyLastSynced = "Last synced"
	// CopyNeverSynced is the addon row before the site has acked an
	// export. The row is shown, not hidden: a missing addon is the
	// thing the player has to act on.
	CopyNeverSynced = "Addon not synced yet. Install ForeverSixty, then log out once."
	// CopyStaleHelp sits under a sync older than a day.
	CopyStaleHelp       = "Log out of the game with the addon installed to sync."
	CopySimQueued       = "Sim queued"
	CopySimNeedsPremium = "Sim needs premium"
	CopySimDuplicate    = "Sim already running"
	CopySimInvalid      = "Sim request invalid"
	// CopySeparator joins the two halves of the line.
	CopySeparator = " · "
)

// StaleAfter is how old the last sync may be before the row warns.
const StaleAfter = 24 * time.Hour

// simCopy is the words for each sim request state the ack can carry.
var simCopy = map[string]string{
	addon.SimQueued:       CopySimQueued,
	addon.SimNeedsPremium: CopySimNeedsPremium,
	addon.SimDuplicate:    CopySimDuplicate,
	addon.SimInvalid:      CopySimInvalid,
}

// AddonLine is the status page's addon row: "Last synced 4 minutes
// ago" under an hour, "Last synced 14:03" later the same day, then
// "Last synced yesterday" and "Last synced 18 days ago", followed by
// the sim request's state when there is one. Before the site has acked
// an export it says so rather than going missing. "Today" and the
// clock are the player's, in now's location.
func AddonLine(seen time.Time, simState string, now time.Time) string {
	if seen.IsZero() {
		return CopyNeverSynced
	}
	line := CopyLastSynced + " " + since(seen, now)
	if words := simCopy[simState]; words != "" {
		line += CopySeparator + words
	}
	return line
}

// AddonStale reports whether the last sync is old enough to warn
// about. No sync at all is not stale: it has its own sentence.
func AddonStale(seen, now time.Time) bool {
	return !seen.IsZero() && now.Sub(seen) > StaleAfter
}

// since is how long ago, in the words the row uses.
func since(seen, now time.Time) string {
	local := seen.In(now.Location())
	d := now.Sub(seen)
	switch {
	case d < time.Minute:
		return "less than a minute ago"
	case d < 2*time.Minute:
		return "1 minute ago"
	case d < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(d/time.Minute))
	}
	y, m, day := local.Date()
	ny, nm, nday := now.Date()
	if y == ny && m == nm && day == nday {
		return local.Format("15:04")
	}
	days := int(time.Date(ny, nm, nday, 12, 0, 0, 0, time.UTC).
		Sub(time.Date(y, m, day, 12, 0, 0, 0, time.UTC)).Hours() / 24)
	if days <= 1 {
		return "yesterday"
	}
	return fmt.Sprintf("%d days ago", days)
}

// The Status headline: five states and the sign-in wait, first match
// wins (see Headline). The tray shows the title word for word.
const (
	CopyConnectingTitle = "Waiting for you to sign in..."
	CopyConnectingSub   = "Finish in your browser. The companion connects by itself."
	CopyConnectingBtn   = "Show the sign-in"

	CopyNotConnectedTitle = "Not connected"
	CopyNotConnectedSub   = "Nothing is uploading until you sign in."
	CopyNotConnectedBtn   = "Sign in with Battle.net"

	CopyNoFolderTitle = "No game folder found"
	CopyNoFolderSub   = "Add your World of Warcraft folder."
	CopyNoFolderBtn   = "Add a folder"

	CopyAdvancedOffTitle = "Advanced logging is off"
	CopyAdvancedOffSub   = "Reports will miss deaths with health, item level, resources and positions. " +
		CopyAdvancedHow

	CopyReadyTitle = "Ready"
	CopyReadySub   = "Waiting for a combat log. Type /combatlog in game, or leave combat logging on."

	CopyLoggingTitle = "Logging"
	CopyLoggingSub   = "Uploading your fights as they end."

	// CopyAdvancedHow is the menu path, and what to do after.
	CopyAdvancedHow = "Turn it on in Options > Network > Advanced Combat Logging, then restart the game."
	// CopyAdvancedRowHelp is the short line a folder row keeps when the
	// headline is about something else, so the sentence is never on
	// screen twice.
	CopyAdvancedRowHelp = "Turn it on in Options > Network > Advanced Combat Logging."
)

// Why a sign-in failed, when the site (not the network) said so.
const (
	CopyConnectDenied  = "You pressed Cancel on the site. Nothing was connected."
	CopyConnectExpired = "The sign-in link expired after 10 minutes."
)

// The tray's one sentence the window shows: what closing it did.
const (
	CopyTrayNotice = "Still running in the tray. Use Quit from the tray icon to stop logging."
)

// The message behind each failure code the window can receive. The
// window shows these as they are.
const (
	CodeBadCode     = "bad_code"
	CodeEmptyCode   = "empty_code"
	CodeOffline     = "offline"
	CodeServer      = "server"
	CodeSave        = "save"
	CodeNotWoW      = "not_wow_folder"
	CodeKnownFolder = "known_folder"
	CodeNoPicker    = "no_picker"
	CodeNoCharacter = "bad_character"
	CodeSettings    = "settings"
	CodeConnected   = "already_connected"
	CodeNoConnect   = "no_connect"
	CodeBadURL      = "bad_url"
	CodeNoBrowser   = "no_browser"
	CodeBadRequest  = "bad_request"
	CodeInternal    = "internal"
	CodeNotAdded    = "not_added"
)

var messages = map[string]string{
	CodeBadCode: "That code is not valid. A code works once and expires after 10 minutes. " +
		"Get a new one on the site.",
	CodeEmptyCode: "Paste the code from the site first.",
	CodeOffline:   "Could not reach foreversixty.gg. Check your connection and try again.",
	CodeServer:    "The site had a problem. Try again in a minute.",
	CodeSave: "Connected, but this PC would not let the companion save the sign-in. " +
		"Nothing will upload. Details are in logs/companion.log.",
	CodeNotWoW: "That folder is not a World of Warcraft folder. " +
		"Choose the one that contains WTF and Logs, such as _classic_era_.",
	CodeKnownFolder: "That folder is already in the list.",
	CodeNoPicker:    "The companion could not open a folder dialog on this computer.",
	CodeNoCharacter: "Choose a character from the list.",
	CodeSettings:    "Could not save the settings. Details are in logs/companion.log.",
	CodeConnected:   "This device is already connected.",
	CodeNoConnect:   "There is no sign-in in progress.",
	CodeBadURL:      "That link cannot be opened from here.",
	CodeNoBrowser:   "Could not open your browser.",
	CodeBadRequest:  "The companion could not read that request.",
	CodeNotAdded:    "That folder was found automatically, so there is nothing to remove.",
	CodeInternal:    "The companion hit a problem. Details are in logs/companion.log.",
}

// Message is the sentence for a failure code.
func Message(code string) string {
	if m, ok := messages[code]; ok {
		return m
	}
	return messages[CodeInternal]
}
