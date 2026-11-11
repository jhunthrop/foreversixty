// companion/internal/app/copy.go
// The player-facing sentences the Go side owns, in one place: the
// headline words (which the tray repeats exactly), the addon line, and
// the message behind every failure the window can show. Every word is a
// key of apps/shared/copy/companion.json, read through internal/copygen
// (generated); the window's own words come from the same file through
// ui/copy_gen.js. Nothing here ever carries an error's text: a failure's
// cause goes to companion.log and the player gets one of these. Times are
// the player's local time. Tests read these constants rather than repeat
// the words.
package app

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addon"
	"github.com/jhunthrop/foreversixty/companion/internal/copygen"
)

// The addon line's words.
const (
	CopyLastSynced = copygen.AddonLastSynced
	// CopyNeverSynced is the addon row before the site has acked an
	// export. The row is shown, not hidden: a missing addon is the
	// thing the player has to act on. CopyNeverSyncedHelp is the muted
	// line under it that says what to do, so the value column holds a
	// short state and not a sentence.
	CopyNeverSynced     = copygen.AddonNeverSynced
	CopyNeverSyncedHelp = copygen.AddonNeverSyncedHelp
	// CopyNoCharacterYetHelp is the same line when an install already has the
	// addon: nothing is missing but the first logout.
	CopyNoCharacterYetHelp = copygen.AddonNoCharacterYetHelp
	// CopyStaleHelp sits under a sync older than a day.
	CopyStaleHelp       = copygen.AddonStaleHelp
	CopySimQueued       = copygen.AddonSimQueued
	CopySimNeedsPremium = copygen.AddonSimNeedsPremium
	CopySimDuplicate    = copygen.AddonSimDuplicate
	CopySimInvalid      = copygen.AddonSimInvalid
	// CopySeparator joins the two halves of the line.
	CopySeparator = copygen.Joint
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

// AddonLine is the status page's addon row: "Last synced 4 min ago",
// "Last synced 2 h ago" or "Last synced 18 d ago", followed by the sim
// request's state when there is one. Before the site has acked an export
// it says so rather than going missing.
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

// agoCeiling is where the age turns into the date, as on the site.
const agoCeiling = 30 * 24 * time.Hour

// since is how long ago, the site's compact age (web/src/lib/dates.ts compactAgo, held to
// apps/shared/formats.json): "just now", "5 min ago", "2 h ago", "3 d ago", then the date past
// 30 days ("1 Oct", in UTC as the site writes it). A time ahead of the clock is "just now".
func since(seen, now time.Time) string {
	d := max(0, now.Sub(seen))
	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		return fmt.Sprintf("%d min ago", int(d/time.Minute))
	case d < 24*time.Hour:
		return fmt.Sprintf("%d h ago", int(d/time.Hour))
	case d < agoCeiling:
		return fmt.Sprintf("%d d ago", int(d/(24*time.Hour)))
	}
	return seen.UTC().Format("2 Jan")
}

// The Status headline: five states and the sign-in wait, first match
// wins (see Headline). The tray shows the title word for word.
const (
	CopyConnectingTitle = copygen.HeadlineConnectingTitle
	CopyConnectingSub   = copygen.HeadlineConnectingSub
	CopyConnectingBtn   = copygen.HeadlineConnectingButton

	CopyNotConnectedTitle = copygen.HeadlineNotConnectedTitle
	CopyNotConnectedSub   = copygen.HeadlineNotConnectedSub
	CopyNotConnectedBtn   = copygen.SignIn

	CopySignedOutTitle = copygen.HeadlineSignedOutTitle
	CopySignedOutSub   = copygen.HeadlineSignedOutSub

	CopyNoFolderTitle = copygen.HeadlineNoFolderTitle
	CopyNoFolderSub   = copygen.HeadlineNoFolderSub
	CopyNoFolderBtn   = copygen.HeadlineNoFolderButton

	CopyAdvancedOffTitle = copygen.HeadlineAdvancedOffTitle
	CopyAdvancedOffSub   = copygen.HeadlineAdvancedOffSub + " " + CopyAdvancedHow

	CopyReadyTitle = copygen.HeadlineReadyTitle
	CopyReadySub   = copygen.HeadlineReadySub

	CopyLoggingTitle = copygen.HeadlineLoggingTitle
	// CopyLoggingTitleHeld is the Logging title while a failed upload
	// holds the queue (withHeld); the Queued row names the fight.
	CopyLoggingTitleHeld = CopyLoggingTitle + CopySeparator + copygen.HeadlineLoggingHeld
	CopyLoggingSub       = copygen.HeadlineLoggingSub
	// CopyLoggingTitleNotCredited is the Logging title while the open report is
	// not credited to the player yet. This is the one Logging state that is
	// amber: the headline carries the one thing the player must do (log out
	// once with the addon loaded), which overrides R-C7 for this case only.
	CopyLoggingTitleNotCredited = CopyLoggingTitle + CopySeparator + copygen.HeadlineLoggingNotCredited

	// CopyAdvancedHow is the menu path, and what to do after.
	CopyAdvancedHow = copygen.HeadlineAdvancedHow
	// CopyAdvancedRowHelp is the short line a folder row keeps when the
	// headline is about something else, so the sentence is never on
	// screen twice.
	CopyAdvancedRowHelp = copygen.HeadlineAdvancedRowHelp
)

// Why a sign-in failed, when the site (not the network) said so.
const (
	CopyConnectDenied  = copygen.ConnectDenied
	CopyConnectExpired = copygen.ConnectExpired
)

// The tray's one sentence the window shows: what closing it did.
const (
	CopyTrayNotice = copygen.TrayNotice
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
	CodeNotPaired   = "not_paired"
	// The Install the addon button's three ways to fail.
	CodeGameRunning   = "game_running"
	CodeAddonDownload = "addon_download"
	CodeAddonInstall  = "addon_install"
)

var messages = map[string]string{
	CodeBadCode:       copygen.FailureBadCode,
	CodeEmptyCode:     copygen.FailureEmptyCode,
	CodeOffline:       copygen.FailureOffline,
	CodeServer:        copygen.FailureServer,
	CodeSave:          copygen.FailureSave,
	CodeNotWoW:        copygen.FailureNotWoW,
	CodeKnownFolder:   copygen.FailureKnownFolder,
	CodeNoPicker:      copygen.FailureNoPicker,
	CodeNoCharacter:   copygen.FailureNoCharacter,
	CodeSettings:      copygen.FailureSettings,
	CodeConnected:     copygen.FailureConnected,
	CodeNoConnect:     copygen.FailureNoConnect,
	CodeBadURL:        copygen.FailureBadURL,
	CodeNoBrowser:     copygen.FailureNoBrowser,
	CodeBadRequest:    copygen.FailureBadRequest,
	CodeNotAdded:      copygen.FailureNotAdded,
	CodeNotPaired:     copygen.FailureNotPaired,
	CodeInternal:      copygen.FailureInternal,
	CodeGameRunning:   copygen.FailureGameRunning,
	CodeAddonDownload: copygen.FailureAddonDownload,
	CodeAddonInstall:  copygen.FailureAddonInstall,
}

// Message is the sentence for a failure code.
func Message(code string) string {
	if m, ok := messages[code]; ok {
		return m
	}
	return messages[CodeInternal]
}

// CopyLoggingSubNotMatched is the Logging headline's sub line once the
// combat log has named its logger and no addon record fits yet: the fights
// are uploading, and the one step that credits them is on the card. The tray
// carries the title only.
const CopyLoggingSubNotMatched = copygen.HeadlineLoggingSubNotMatched

// CopyLoggingSubNoName is the sub line while the combat log has named nobody
// yet (credit none on an open report): nothing is wrong, the name arrives
// with the first fight.
const CopyLoggingSubNoName = copygen.HeadlineLoggingSubNoName

// CopyLoggingSubAs is the Logging headline's sub line for a report
// credited to a character: the same words, then who it is. The name is
// printed in its space form, whatever form it was stored in, because the
// hyphenated one is the log's and never the player's.
func CopyLoggingSubAs(name string) string {
	return CopyLoggingSub + CopySeparator + copygen.Fill(copygen.HeadlineLoggingAs, "{name}", SpacedName(name))
}

// SpacedName is a character name as a player reads it: "Bow Jackzon",
// never the log's "Bow-Jackzon". A WoW name has no hyphen of its own, so
// every hyphen is a joint between words.
func SpacedName(name string) string {
	return strings.Join(strings.FieldsFunc(name, func(r rune) bool { return r == '-' || r == ' ' }), " ")
}

// The Characters sent row's lines.
const (
	// CopyNoCharactersSent is the row when the addon is installed and
	// no character list has reached the site.
	CopyNoCharactersSent = copygen.AddonNoCharactersSent
)

// CopySiteRefused is the row when the site read the character list and
// refused it; reason is the API's own message.
func CopySiteRefused(reason string) string {
	return copygen.Fill(copygen.AddonSiteRefused, "{reason}", reason)
}

// The Addon row's version lines (addon_update in the snapshot). One line
// per family per install, in the same words every time; none of them ever
// shows a bare number, and none says manifest, asset or stage.
const (
	// CopyAddonName and CopyDataName start the lines.
	CopyAddonName = copygen.VersionAddon
	CopyDataName  = copygen.VersionData
	// CopyUpToDate ends a line for a version that is the newest.
	CopyUpToDate = copygen.VersionUpToDate
	// CopyTakesEffect ends the line for a version the companion just put
	// in, until the player's next login is seen in the log.
	CopyTakesEffect = copygen.VersionTakesEffect
	// CopyInstallsWhenClosed ends the line for a version waiting for the
	// game to close.
	CopyInstallsWhenClosed = copygen.VersionInstallsWhenClosed
)

// CopyVersionLine words one family in one install: "Addon 2.0.15 · up to
// date", "Addon 2.0.16 installed · takes effect at your next login" or
// "Addon 2.0.16 ready · installs when the game closes". name is
// CopyAddonName or CopyDataName.
func CopyVersionLine(name, version string, state AddonState) string {
	switch state {
	case AddonInstalled:
		return copygen.Fill(copygen.VersionInstalled, "{name}", name, "{version}", version) + CopySeparator + CopyTakesEffect
	case AddonReady:
		return copygen.Fill(copygen.VersionReady, "{name}", name, "{version}", version) + CopySeparator + CopyInstallsWhenClosed
	}
	return name + " " + version + CopySeparator + CopyUpToDate
}

// CopyAddonNotFound is the Addon row's line for an install with no addon.
func CopyAddonNotFound(folder string) string {
	return copygen.Fill(copygen.AddonNotFound, "{folder}", folder)
}

// CopyInstallAddon is the label of the button that puts the addon into an
// install that has none.
func CopyInstallAddon(folder string) string {
	return copygen.Fill(copygen.AddonInstall, "{folder}", folder)
}

// CopyForInstall puts an install's name before its lines when the Addon row
// speaks for more than one.
func CopyForInstall(folder, line string) string { return folder + ": " + line }

// CopyBuildsWaiting is the Addon row's line while a sent build is in (or
// on its way to) the addon's inbox and the addon has not shown it yet:
// "A build is waiting for Bow Jackzon. Log out to the character screen
// and back in." A /reload does not show it, so the line does not offer one.
func CopyBuildsWaiting(n int, who []string) string {
	forWho := ""
	if len(who) > 0 {
		forWho = copygen.Fill(copygen.AddonBuildsFor, "{who}", strings.Join(who, ", "))
	}
	if n == 1 {
		return copygen.Fill(copygen.AddonBuildWaiting, "{forWho}", forWho)
	}
	return copygen.Fill(copygen.AddonBuildsWaiting, "{n}", strconv.Itoa(n), "{forWho}", forWho)
}
