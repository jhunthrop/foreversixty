package guilds

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/contract"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// feedOf reads a calendar-feed answer.
func feedOf(t *testing.T, w raidWorld, method, path string) CalendarFeedView {
	t.Helper()
	var view CalendarFeedView
	w.h.data(w.h.do(method, path, ""), &view)
	return view
}

// icsPath is the public route of a feed URL.
func icsPath(t *testing.T, url *string) string {
	t.Helper()
	if url == nil {
		t.Fatal("the answer carries no url")
	}
	i := strings.Index(*url, CalendarFeedPrefix)
	if i < 0 {
		t.Fatalf("url %q is not a calendar link", *url)
	}
	return (*url)[i:]
}

func TestCalendarFeedIsShownOnceAndStoredHashed(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)

	first := feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed")
	path := icsPath(t, first.URL)
	token := strings.TrimSuffix(strings.TrimPrefix(path, CalendarFeedPrefix), calendarFeedSuffix)
	if !calendarToken.MatchString(token) {
		t.Fatalf("token %q is not 256 bits of base64url", token)
	}
	again := feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed")
	if again.URL != nil || !again.CreatedAt.Equal(first.CreatedAt) {
		t.Fatalf("second read = %+v, want the same link with no url", again)
	}

	var stored []byte
	var live int
	if err := w.h.pool.QueryRow(t.Context(),
		`select token_hash, (select count(*) from calendar_feeds where user_id = $1 and revoked_at is null)
		 from calendar_feeds where user_id = $1`, w.member).Scan(&stored, &live); err != nil {
		t.Fatal(err)
	}
	if live != 1 || len(stored) != 32 || bytes.Contains(stored, []byte(token)) {
		t.Fatalf("stored %x (%d live), want one SHA-256 and never the token", stored, live)
	}
}

func TestCalendarFeedResetRevokesTheOldLink(t *testing.T) {
	w := newRaidWorld(t, true)
	w.createRaid(t)
	w.as(w.member)
	old := icsPath(t, feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed").URL)
	guildBody(t, w.h.do(http.MethodGet, old, ""), http.StatusOK)

	reset := feedOf(t, w, http.MethodPost, "/v1/me/calendar-feed/reset")
	fresh := icsPath(t, reset.URL)
	if fresh == old {
		t.Fatal("a reset issues a new link")
	}
	w.as(0) // the feed is public by token
	expectError(t, w.h.do(http.MethodGet, old, ""), http.StatusNotFound, "not_found")
	guildBody(t, w.h.do(http.MethodGet, fresh, ""), http.StatusOK)
}

func TestCalendarICSListsYourGuildsRaidsWithYourAnswer(t *testing.T) {
	w := newRaidWorld(t, true)
	raid := w.createRaid(t)
	cancelled := w.createRaid(t)
	w.as(w.officer)
	guildBody(t, w.h.do(http.MethodPost, w.path("/raids/%d/cancel", cancelled.ID), ""), http.StatusOK)
	w.as(w.member)
	w.signup(t, raid.ID, memberKeyA, `{"state":"late","role":"tank"}`)
	memberFeed := icsPath(t, feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed").URL)
	w.as(w.member2)
	member2Feed := icsPath(t, feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed").URL)
	stranger := seedUser(t, w.h.pool, "calendar-stranger@example.com")
	w.as(stranger)
	strangerFeed := icsPath(t, feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed").URL)
	w.as(0)

	res := w.h.do(http.MethodGet, memberFeed, "")
	body := string(guildBody(t, res, http.StatusOK))
	for header, want := range map[string]string{
		"Content-Type": "text/calendar; charset=utf-8", "Cache-Control": "private, max-age=300",
		"X-Content-Type-Options": "nosniff", "Referrer-Policy": "no-referrer",
	} {
		if got := res.Header.Get(header); got != want {
			t.Errorf("%s = %q, want %q", header, got, want)
		}
	}
	uid := "UID:raid-" + itoa(raid.ID) + "@foreversixty.gg"
	for _, want := range []string{"BEGIN:VCALENDAR\r\n", uid, "SUMMARY:Onyxia's Lair · Late", "LOCATION:Olympus · Onyxia's Lair",
		"of 40 signed up.", "X-WR-CALNAME:Forever Sixty raids"} {
		if !strings.Contains(body, want) {
			t.Errorf("feed lacks %q:\n%s", want, body)
		}
	}
	for _, never := range []string{"VALARM", "raid-" + itoa(cancelled.ID) + "@", "flasks"} {
		if strings.Contains(body, never) {
			t.Errorf("feed carries %q", never)
		}
	}
	if b := string(guildBody(t, w.h.do(http.MethodGet, member2Feed, ""), http.StatusOK)); !strings.Contains(b, "SUMMARY:Onyxia's Lair · Not signed up") {
		t.Errorf("a member who has not answered reads Not signed up:\n%s", b)
	}
	if b := string(guildBody(t, w.h.do(http.MethodGet, strangerFeed, ""), http.StatusOK)); strings.Contains(b, "BEGIN:VEVENT") {
		t.Errorf("someone outside the guild sees its raids:\n%s", b)
	}
}

func TestCalendarICSRefusesEveryBadLinkTheSameWay(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(0)
	contract.AssertFixture(t, "error-unauthorized",
		guildBody(t, w.h.do(http.MethodGet, "/v1/me/calendar-feed", ""), http.StatusUnauthorized), contract.Options{})
	unknown := guildBody(t, w.h.do(http.MethodGet, CalendarFeedPrefix+strings.Repeat("A", 43)+".ics", ""), http.StatusNotFound)
	contract.AssertFixture(t, "calendar.ics.error-not-found", unknown, contract.Options{})
	for _, bad := range []string{"short.ics", strings.Repeat("A", 43), strings.Repeat("A", 43) + ".txt", strings.Repeat("A", 44) + ".ics"} {
		got := guildBody(t, w.h.do(http.MethodGet, CalendarFeedPrefix+bad, ""), http.StatusNotFound)
		var a, b map[string]any
		_ = json.Unmarshal(got, &a)
		_ = json.Unmarshal(unknown, &b)
		if a["error"].(map[string]any)["message"] != b["error"].(map[string]any)["message"] {
			t.Errorf("%s answers differently from an unknown token: %s", bad, got)
		}
	}
}

func TestCalendarICSBudgetIsSixtyAnHourPerLink(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)
	path := icsPath(t, feedOf(t, w, http.MethodGet, "/v1/me/calendar-feed").URL)
	w.as(0)
	for i := 0; i < calendarReadsPerTokenHour; i++ {
		guildBody(t, w.h.do(http.MethodGet, path, ""), http.StatusOK)
	}
	expectError(t, w.h.do(http.MethodGet, path, ""), http.StatusTooManyRequests, "rate_limited")
}

func TestCalendarFeedResetBudget(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)
	for i := 0; i < calendarResetsPerHour; i++ {
		guildBody(t, w.h.do(http.MethodPost, "/v1/me/calendar-feed/reset", ""), http.StatusOK)
	}
	expectError(t, w.h.do(http.MethodPost, "/v1/me/calendar-feed/reset", ""), http.StatusTooManyRequests, "rate_limited")
}

func TestAppContractCalendarFeed(t *testing.T) {
	w := newRaidWorld(t, true)
	w.as(w.member)
	res := w.h.do(http.MethodGet, "/v1/me/calendar-feed", "")
	if got := res.Header.Get("Cache-Control"); got != "private, no-store" {
		t.Errorf("Cache-Control = %q, want private, no-store", got)
	}
	contract.AssertFixture(t, "calendar.feed", guildBody(t, res, http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "calendar.feed.shown",
		guildBody(t, w.h.do(http.MethodGet, "/v1/me/calendar-feed", ""), http.StatusOK), contract.Options{})
	contract.AssertFixture(t, "calendar.feed",
		guildBody(t, w.h.do(http.MethodPost, "/v1/me/calendar-feed/reset", ""), http.StatusOK), contract.Options{})
}

// The request log redacts the route this package serves the token on.
func TestCalendarRouteIsRedactedInTheLogs(t *testing.T) {
	if got := httpx.LogPath(CalendarFeedPrefix + "token.ics"); strings.Contains(got, "token") {
		t.Fatalf("LogPath keeps the token: %q", got)
	}
}

// RenderCalendar is held to the golden feed in the contract fixtures, the one the docs show.
func TestRenderCalendarGolden(t *testing.T) {
	size := time.Date(2026, 12, 10, 0, 0, 0, 0, time.UTC)
	raids := []CalendarRaid{
		{
			Raid: Raid{ID: 31, Zone: raidZoneOny, Title: "Onyxia's Lair", StartsAt: time.Date(2026, 12, 11, 1, 45, 0, 0, time.UTC),
				DurationMinutes: 195, Status: RaidPlanned, UpdatedAt: size},
			Guild: Guild{Region: "us", Ruleset: "pvp", Name: "OLYMPUS XXVII"}, Timezone: "America/Chicago", SignedUp: 24, State: "in",
		},
		{
			Raid: Raid{ID: 32, Zone: raidZoneOny, Title: "Onyxia, alts; bring resists", StartsAt: time.Date(2026, 12, 18, 1, 45, 0, 0, time.UTC),
				DurationMinutes: 120, Status: RaidPlanned, UpdatedAt: size},
			Guild: Guild{Region: "us", Ruleset: "pvp", Name: "OLYMPUS XXVII"}, Timezone: "America/Chicago", SignedUp: 0,
		},
	}
	got := RenderCalendar(raids, "https://foreversixty.gg", size)
	path := filepath.Join(filepath.Dir(contract.FixturePath("x")), "calendar.ics")
	if os.Getenv("UPDATE_GOLDEN") != "" {
		if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	want, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != string(want) {
		t.Fatalf("feed differs from %s:\n%s", path, got)
	}
	for _, line := range strings.Split(strings.TrimSuffix(got, "\r\n"), "\r\n") {
		if len(line) > icsLineOctets {
			t.Errorf("line over 75 octets: %q", line)
		}
	}
}

func TestICSTextEscapesEverySeparator(t *testing.T) {
	if got := escapeICSText("a\\b;c,d\ne"); got != `a\\b\;c\,d\ne` {
		t.Fatalf("escapeICSText = %q", got)
	}
	long := strings.Repeat("·", 60)
	folded := foldICSLine("SUMMARY:" + long)
	for _, part := range strings.Split(strings.TrimSuffix(folded, "\r\n"), "\r\n ") {
		if len(part) > icsLineOctets || !strings.HasPrefix(part+"·", part) {
			t.Fatalf("bad fold %q", part)
		}
	}
	if strings.ReplaceAll(strings.TrimSuffix(folded, "\r\n"), "\r\n ", "") != "SUMMARY:"+long {
		t.Fatal("unfolding gives the line back")
	}
}

func itoa(id int64) string { return strconv.FormatInt(id, 10) }
