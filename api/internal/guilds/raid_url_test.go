// api/internal/guilds/raid_url_test.go
package guilds

import (
	"net/url"
	"regexp"
	"slices"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

// The site's rules for a guild raid path, mirrored from web/src/lib/characters.ts.
var (
	siteRegions        = []string{"us", "eu", "kr", "tw", "cn"}
	siteRulesets       = []string{"normal", "pvp", "rp", "hardcore"}
	siteRaidID         = regexp.MustCompile(`^\d+$`)
	siteUnsafeInSlug   = regexp.MustCompile(`[/\\.?#%\s]`)
	siteMaxSlugLength  = 64
	siteMaxEncodedSlug = siteMaxSlugLength * 4
)

// parseGuildRaidPathGo mirrors parseGuildRaidPath: exactly six segments, guild, a region, a
// ruleset, a character-safe slug, raids, and a digits-only id.
func parseGuildRaidPathGo(path string) bool {
	var parts []string
	for _, p := range strings.Split(strings.TrimRight(path, "/"), "/") {
		if p != "" {
			parts = append(parts, p)
		}
	}
	if len(parts) != 6 || parts[0] != "guild" || parts[4] != "raids" {
		return false
	}
	return slices.Contains(siteRegions, parts[1]) && slices.Contains(siteRulesets, parts[2]) &&
		isSiteSlug(parts[3]) && siteRaidID.MatchString(parts[5])
}

// isSiteSlug mirrors isCharacterSlug: decoded, bounded, nothing that traverses or re-targets.
func isSiteSlug(slug string) bool {
	if slug == "" || len(slug) > siteMaxEncodedSlug {
		return false
	}
	decoded, err := url.PathUnescape(slug)
	if err != nil || utf8.RuneCountInString(decoded) > siteMaxSlugLength || siteUnsafeInSlug.MatchString(decoded) {
		return false
	}
	for _, r := range decoded {
		if r < 0x20 {
			return false
		}
	}
	return true
}

func TestParseGuildRaidPathGoMirrorsTheSite(t *testing.T) {
	for path, want := range map[string]bool{
		"/guild/us/normal/olympus-xxvii/raids/1":   true,
		"/guild/us/normal/olympus-xxvii/raids/1/":  true,
		"/guild/us/normal/olympus-xxvii/raids/x1":  false,
		"/guild/us/normal/olympus-xxvii/raid/1":    false,
		"/guild/xx/normal/olympus-xxvii/raids/1":   false,
		"/guild/us/classic/olympus-xxvii/raids/1":  false,
		"/guild/us/normal/%2e%2e/raids/1":          false,
		"/guild/us/normal/olympus/xxvii/raids/1":   false,
		"/guild/us/normal/olympus-xxvii/raids/1/2": false,
	} {
		if got := parseGuildRaidPathGo(path); got != want {
			t.Errorf("%s = %v, want %v", path, got, want)
		}
	}
}

func TestEveryPostedRaidURLIsTheSitePath(t *testing.T) {
	const example = "https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/1"
	u, err := url.Parse(example)
	if err != nil || !parseGuildRaidPathGo(u.EscapedPath()) {
		t.Fatalf("the example %s fails the site rule", example)
	}
	g := Guild{Region: "US", Ruleset: "Normal", Name: "Olympus XXVII"}
	if got := testSiteBase + RaidPagePath(GuildPagePath(g), 1); got != example {
		t.Fatalf("Olympus XXVII raid 1 = %s, want %s", got, example)
	}

	w, events, _ := postingWorld(t)
	views := w.createRaids(t, scheduleBody(startIn(2*day), raidRepeatWeeks, "true", "true"))
	now := time.Now()
	for _, at := range []time.Duration{6 * day, 13 * day, 20 * day} {
		w.run(t, now.Add(at))
	}
	posted, _, _, _ := events.snapshot()
	if len(posted) != len(views) {
		t.Fatalf("posted %d of %d raids", len(posted), len(views))
	}
	for _, e := range posted {
		u, err := url.Parse(e.RaidURL)
		if err != nil || !strings.HasPrefix(e.RaidURL, testSiteBase+"/") || !parseGuildRaidPathGo(u.EscapedPath()) {
			t.Fatalf("posted raid URL %q is not a site path", e.RaidURL)
		}
	}
}
