// Package golinks is the tracked short links, foreversixty.gg/go/<slug>: the owner's list of
// links (a slug, a path on this site, a label), one click counter per link per day with bots
// apart, the first link an account arrived from, and the funnel that joins them. Nothing
// about a visitor is stored: no IP, no user agent, no cookie. The site's Worker resolves and
// counts a click; the site sends the arrival with the first sign-in.
package golinks

import (
	"regexp"
	"strings"
	"time"
)

// The limits migration 0066's checks hold; the handler refuses anything past them first, so a
// bad link is a 400 with a reason rather than a constraint error.
const (
	MaxTargetLen = 200
	MaxLabelLen  = 80
)

// ArrivalWindow is how new an account must be for its arrival to count as how it came: the
// email link takes up to 20 minutes, so an hour covers every sign-up and no returning account.
const ArrivalWindow = time.Hour

// ClickWindow is how long the site keeps a click before sign-up (the web's ARRIVAL_TTL_MS): a
// creator's link counts for 30 days (creators.LinkWindow).
const ClickWindow = 30 * 24 * time.Hour

var slugPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{0,39}$`)

// ValidSlug reports whether s is a link's slug: lowercase letters, digits and dashes, 1 to 40.
func ValidSlug(s string) bool { return slugPattern.MatchString(s) }

// ValidTarget reports whether s is a path on this site: one leading slash, never two, no
// backslash and no whitespace, so a link can never send a visitor anywhere else.
func ValidTarget(s string) bool {
	if len(s) == 0 || len(s) > MaxTargetLen || s[0] != '/' || strings.HasPrefix(s, "//") {
		return false
	}
	for _, r := range s {
		if r == '\\' || r <= ' ' || r == 0x7f {
			return false
		}
	}
	return true
}

// ValidLabel reports whether s is a label the owner's table can show.
func ValidLabel(s string) bool {
	trimmed := strings.TrimSpace(s)
	return trimmed != "" && len(s) <= MaxLabelLen
}

// Link is one tracked link.
type Link struct {
	Slug      string    `json:"slug"`
	Target    string    `json:"target"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
}

// FunnelRow is one link's funnel: people's clicks (bots apart), the accounts that arrived
// through it, how many of those linked a character, and how many hold Premium now.
type FunnelRow struct {
	Link
	Clicks    int64 `json:"clicks"`
	BotClicks int64 `json:"bot_clicks"`
	SignIns   int64 `json:"sign_ins"`
	Linked    int64 `json:"linked"`
	Premium   int64 `json:"premium"`
}
