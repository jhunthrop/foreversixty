package reports

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

func TestCreateAndReadAReport(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	if len(id) != 12 {
		t.Fatalf("report id = %q, want twelve characters", id)
	}

	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	var view View
	h.data(res, &view)
	if view.Title != "Tuesday" || view.Zone != "Blackrock Spire" || view.Status != StatusLive {
		t.Fatalf("view = %+v", view)
	}
	if view.Owner == nil || view.Owner.ID != h.owner {
		t.Fatalf("owner = %+v, want the creator", view.Owner)
	}
	if view.DataBaseURL != "https://foreversixty.gg/logs-data/reports/"+id {
		t.Fatalf("data_base_url = %q", view.DataBaseURL)
	}
	if view.Fights == nil || view.Players == nil {
		t.Fatalf("a new report should carry empty lists, not null: %+v", view)
	}
}

func TestCreateRejectsABadVisibilityAndCharacter(t *testing.T) {
	h := newHarness(t)
	res := h.json(http.MethodPost, "/v1/reports", `{"visibility":"secret"}`)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	if h.errorFields(res)["visibility"] == "" {
		t.Fatal("the 400 should name the visibility field")
	}
	res = h.json(http.MethodPost, "/v1/reports",
		`{"visibility":"public","logging_character":{"region":"us","ruleset":"nightslayer","name":"Baelgrim"}}`)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400 for a realm that is not a ruleset", res.StatusCode)
	}
	if h.errorFields(res)["logging_character"] == "" {
		t.Fatal("the 400 should name the character field")
	}
	res = h.json(http.MethodPost, "/v1/reports", `{`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400 for a broken body", res.StatusCode)
	}
}

func TestAPrivateReportIsInvisibleToEveryoneElse(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Private)

	h.actor = auth.Actor{UserID: h.owner + 99, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a stranger sees %d, want 404", res.StatusCode)
	}

	h.actor = auth.Actor{}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an anonymous reader sees %d, want 404", res.StatusCode)
	}

	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	var view View
	h.data(res, &view)
	if view.DataBaseURL != "https://api.foreversixty.gg/v1/reports/"+id+"/files" {
		t.Fatalf("a private report's files are served by the API: %q", view.DataBaseURL)
	}

	h.actor = auth.Actor{UserID: h.owner + 99, Role: "moderator", Method: "session"}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("a moderator sees %d, want 200", res.StatusCode)
	}
}

func TestAGuildReportIsVisibleToTheGuild(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(GuildTo)
	var guildID int64
	if err := h.store.Pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Forever') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`update reports set guild_id = $2 where id = $1`, id, guildID); err != nil {
		t.Fatal(err)
	}
	member := h.owner + 1
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into users (id, email) values ($1, 'member@example.com')
		 on conflict (id) do nothing`, member); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, 'member', now())`,
		guildID, member); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: member, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	var view View
	h.data(res, &view)
	if view.Guild == nil || view.Guild.Name != "Forever" || view.Guild.Ruleset != "hardcore" {
		t.Fatalf("guild = %+v", view.Guild)
	}

	h.actor = auth.Actor{UserID: member + 1, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("a non-member sees %d, want 404", res.StatusCode)
	}
}

// TestAFrozenClaimantsReportEditRightIsSuspendedButOnlyForOthers is D's
// regression net (third security review response): while a guild's
// claim is contested and frozen, the disputed claimant's
// officer-derived edit right over ANOTHER member's guild report is
// refused; their own report (ownership always wins first) still edits
// fine; GET is entirely unaffected by the freeze (mayView never
// consults it); and a different, unrelated verified officer of the
// same guild is unaffected.
func TestAFrozenClaimantsReportEditRightIsSuspendedButOnlyForOthers(t *testing.T) {
	h := newHarness(t)
	ctx := t.Context()

	var gid int64
	if err := h.store.Pool.QueryRow(ctx,
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Frozen Freehold') returning id`).
		Scan(&gid); err != nil {
		t.Fatal(err)
	}

	claimant := h.owner
	if _, err := h.store.Pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/claimant', $2, 'leader')`,
		gid, claimant); err != nil {
		t.Fatal(err)
	}
	if _, err := h.guilds.Claim(ctx, gid, claimant, true); err != nil {
		t.Fatal(err)
	}

	contester := claimant + 1
	if _, err := h.store.Pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'frozen-contester@example.com') on conflict (id) do nothing`, contester); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank) values ($1, 'us/hardcore/contester', $2, 'officer')`,
		gid, contester); err != nil {
		t.Fatal(err)
	}
	if err := h.guilds.ContestClaim(ctx, gid, contester, true); err != nil {
		t.Fatal(err)
	}
	// The claim was established moments ago by Claim() above, so it is
	// young (well under 14 days) - the contest freezes without needing
	// a separate corroboration setup.

	otherOfficer := claimant + 2
	if _, err := h.store.Pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'frozen-other-officer@example.com') on conflict (id) do nothing`, otherOfficer); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(ctx,
		`insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, 'officer', now())`,
		gid, otherOfficer); err != nil {
		t.Fatal(err)
	}

	otherMember := claimant + 3
	if _, err := h.store.Pool.Exec(ctx,
		`insert into users (id, email) values ($1, 'frozen-other-member@example.com') on conflict (id) do nothing`, otherMember); err != nil {
		t.Fatal(err)
	}
	h.actor = auth.Actor{UserID: otherMember, Role: "user", Method: "session"}
	othersReport := h.createReport(GuildTo)
	if _, err := h.store.Pool.Exec(ctx, `update reports set guild_id = $2 where id = $1`, othersReport, gid); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: claimant, Role: "user", Method: "session"}
	ownReport := h.createReport(GuildTo)
	if _, err := h.store.Pool.Exec(ctx, `update reports set guild_id = $2 where id = $1`, ownReport, gid); err != nil {
		t.Fatal(err)
	}

	// The frozen claimant may not edit another member's guild report.
	res := h.json(http.MethodPatch, "/v1/reports/"+othersReport, `{"title":"Claimant edits someone else's"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a frozen claimant patching another member's report = %d, want 403", res.StatusCode)
	}

	// The frozen claimant may still edit their OWN report - ownership
	// wins before the rank/freeze check is even reached.
	res = h.json(http.MethodPatch, "/v1/reports/"+ownReport, `{"title":"Claimant edits their own"}`)
	var view View
	h.data(res, &view)
	if view.Title != "Claimant edits their own" {
		t.Fatalf("a frozen claimant patching their own report failed: %+v", view)
	}

	// GET is entirely unaffected by the freeze - mayView never consults it.
	res = h.do(http.MethodGet, "/v1/reports/"+othersReport, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("a frozen claimant reading another member's guild report = %d, want 200 (GET is unaffected)", res.StatusCode)
	}

	// A different, unrelated verified officer of the same guild is unaffected.
	h.actor = auth.Actor{UserID: otherOfficer, Role: "user", Method: "session"}
	res = h.json(http.MethodPatch, "/v1/reports/"+othersReport, `{"title":"Another officer edits fine"}`)
	h.data(res, &view)
	if view.Title != "Another officer edits fine" {
		t.Fatalf("an unrelated verified officer should be unaffected by the freeze: %+v", view)
	}
}

func TestPatchIsForTheOwnerAndGuildOfficers(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)

	res := h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Wednesday","visibility":"unlisted"}`)
	var view View
	h.data(res, &view)
	if view.Title != "Wednesday" || view.Visibility != Unlisted {
		t.Fatalf("view = %+v", view)
	}

	h.actor = auth.Actor{UserID: h.owner + 50, Role: "user", Method: "session"}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Mine now"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a stranger patching = %d, want 403", res.StatusCode)
	}

	var guildID int64
	if err := h.store.Pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'pvp', 'Officers') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`update reports set guild_id = $2 where id = $1`, id, guildID); err != nil {
		t.Fatal(err)
	}
	officer := h.owner + 50
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into users (id, email) values ($1, 'officer@example.com') on conflict (id) do nothing`,
		officer); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, 'officer', now())`,
		guildID, officer); err != nil {
		t.Fatal(err)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Guild night"}`)
	h.data(res, &view)
	if view.Title != "Guild night" {
		t.Fatalf("an officer could not patch: %+v", view)
	}

	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"visibility":"invisible"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a bad visibility = %d, want 400", res.StatusCode)
	}
}

// A private report answers 404 to a read it refuses; PATCH must refuse
// the same way rather than confirming existence with a 403.
func TestPatchingAPrivateReportIsInvisibleToAStranger(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Private)

	h.actor = auth.Actor{UserID: h.owner + 99, Role: "user", Method: "session"}
	stranger := h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Not yours"}`)
	stranger.Body.Close()
	unknown := h.json(http.MethodPatch, "/v1/reports/nosuchreport", `{"title":"x"}`)
	unknown.Body.Close()

	if stranger.StatusCode != http.StatusNotFound {
		t.Fatalf("a stranger patching a private report = %d, want 404", stranger.StatusCode)
	}
	if stranger.StatusCode != unknown.StatusCode {
		t.Fatalf("a private report and one that does not exist should answer the same: %d vs %d",
			stranger.StatusCode, unknown.StatusCode)
	}
}

// Setting guild_id hands that guild's officers edit rights on the
// report and its members read rights once visibility is guild, so the
// caller must have standing - officer or leader - in the guild they
// are naming, not just ownership of the report.
func TestPatchingGuildIDRequiresStandingInTheTargetGuild(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)

	var guildID int64
	if err := h.store.Pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'normal', 'Targets') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}

	// The owner has no standing at all in that guild: refused, and the
	// report is left untouched.
	res := h.json(http.MethodPatch, "/v1/reports/"+id, fmt.Sprintf(`{"guild_id":%d}`, guildID))
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("attaching to a guild with no standing = %d, want 403", res.StatusCode)
	}
	rep, err := h.store.Get(t.Context(), id)
	if err != nil || rep.GuildID != nil {
		t.Fatalf("a refused patch should not change guild_id: %+v, %v", rep.GuildID, err)
	}

	// Plain membership is not enough either.
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into guild_members (guild_id, user_id, rank, verified_at) values ($1, $2, 'member', now())`,
		guildID, h.owner); err != nil {
		t.Fatal(err)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, fmt.Sprintf(`{"guild_id":%d}`, guildID))
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("attaching to a guild as a plain member = %d, want 403", res.StatusCode)
	}

	// An officer of the target guild may attach the report to it.
	if _, err := h.store.Pool.Exec(t.Context(),
		`update guild_members set rank = 'officer' where guild_id = $1 and user_id = $2`,
		guildID, h.owner); err != nil {
		t.Fatal(err)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, fmt.Sprintf(`{"guild_id":%d}`, guildID))
	var view View
	h.data(res, &view)
	if view.Guild == nil || view.Guild.ID != guildID {
		t.Fatalf("an officer could not attach the report to their guild: %+v", view.Guild)
	}
}

// TestAnUnverifiedGuildCharacterCannotSeeOrEditAGuildReport is the
// spoofing regression test (design §3.3): a bare guild_characters row
// from an export - forged or not - must never grant access on its own.
// A single appearance in the guild's own uploaded report does not
// verify it; a second appearance on a distinct report date within 30
// days does; a second appearance more than 30 days after the first
// does not.
func TestAnUnverifiedGuildCharacterCannotSeeOrEditAGuildReport(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(GuildTo)
	var guildID int64
	if err := h.store.Pool.QueryRow(t.Context(),
		`insert into guilds (region, ruleset, name) values ('us', 'hardcore', 'Spoofed') returning id`).
		Scan(&guildID); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`update reports set guild_id = $2 where id = $1`, id, guildID); err != nil {
		t.Fatal(err)
	}
	spoofer := h.owner + 500
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into users (id, email) values ($1, 'spoofer@example.com') on conflict (id) do nothing`,
		spoofer); err != nil {
		t.Fatal(err)
	}
	// A bare, unverified guild_characters row - exactly what a forged
	// export alone can produce.
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into guild_characters (guild_id, character_key, user_id, rank)
		 values ($1, 'us/hardcore/spoofer', $2, 'member')`, guildID, spoofer); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`select 1 from guild_characters`); err != nil { // sanity: table reachable from this package
		t.Fatal(err)
	}
	guildStore := &guilds.Store{Pool: h.store.Pool}
	tx, err := h.store.Pool.Begin(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	uid := spoofer
	if err := guilds.RecomputeMembership(t.Context(), tx, guildID, &uid); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(t.Context()); err != nil {
		t.Fatal(err)
	}

	h.actor = auth.Actor{UserID: spoofer, Role: "user", Method: "session"}
	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unverified spoofed member reading a guild report = %d, want 404", res.StatusCode)
	}
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"stolen"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden && res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unverified spoofed member editing a guild report = %d, want 403 or 404", res.StatusCode)
	}

	// One appearance in the guild's own uploaded report does not verify it.
	first := time.Now().Add(-20 * 24 * time.Hour)
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ('spoofnight1x', $1, $2, 'guild', 'complete', $3)`, spoofer, guildID, first); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into fights (report_id, fight_index, players) values ('spoofnight1x', 0, $1)`,
		[]string{"us/hardcore/spoofer"}); err != nil {
		t.Fatal(err)
	}
	if err := guildStore.VerifyByLogs(t.Context()); err != nil {
		t.Fatal(err)
	}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("one guild-log appearance should not verify the spoofer: status = %d, want 404", res.StatusCode)
	}

	// A second appearance whose own date falls outside the trailing
	// 30-day window (VerifyByLogs counts distinct report dates with
	// r2.created_at >= now() - 30 days; a report older than that is not
	// counted at all) still leaves only one countable date, so it still
	// does not verify.
	tooLate := time.Now().Add(-35 * 24 * time.Hour)
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ('spoofnight2late', $1, $2, 'guild', 'complete', $3)`, spoofer, guildID, tooLate); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into fights (report_id, fight_index, players) values ('spoofnight2late', 0, $1)`,
		[]string{"us/hardcore/spoofer"}); err != nil {
		t.Fatal(err)
	}
	if err := guildStore.VerifyByLogs(t.Context()); err != nil {
		t.Fatal(err)
	}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("both appearances fall outside a shared 30-day window from `first`: status = %d, want 404", res.StatusCode)
	}

	// A second appearance within 30 days of the first verifies it.
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into reports (id, owner_id, guild_id, visibility, status, created_at)
		 values ('spoofnight2ok', $1, $2, 'guild', 'complete', $3)`,
		spoofer, guildID, first.Add(5*24*time.Hour)); err != nil {
		t.Fatal(err)
	}
	if _, err := h.store.Pool.Exec(t.Context(),
		`insert into fights (report_id, fight_index, players) values ('spoofnight2ok', 0, $1)`,
		[]string{"us/hardcore/spoofer"}); err != nil {
		t.Fatal(err)
	}
	if err := guildStore.VerifyByLogs(t.Context()); err != nil {
		t.Fatal(err)
	}
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	var view View
	h.data(res, &view)
	if view.Guild == nil {
		t.Fatal("two distinct report dates within 30 days should have verified the spoofer")
	}
}

func TestVisibilityRouteIsCachedForTheWorker(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Unlisted)
	res := h.do(http.MethodGet, "/v1/reports/"+id+"/visibility", "", nil)
	var out struct {
		Visibility string `json:"visibility"`
	}
	if got := res.Header.Get("Cache-Control"); got != "public, max-age=60" {
		t.Fatalf("cache-control = %q", got)
	}
	h.data(res, &out)
	if out.Visibility != Unlisted {
		t.Fatalf("visibility = %q", out.Visibility)
	}
	res = h.do(http.MethodGet, "/v1/reports/nosuchreport/visibility", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", res.StatusCode)
	}
}

func TestAccessHandsOutSignedRedirectsForAPrivateReport(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	id := h.createReport(Private)

	res := h.do(http.MethodGet, "/v1/reports/"+id+"/access", "", nil)
	var access struct {
		DataBaseURL string `json:"data_base_url"`
		ExpiresIn   int    `json:"expires_in"`
	}
	h.data(res, &access)
	if !strings.HasSuffix(access.DataBaseURL, "/v1/reports/"+id+"/files") || access.ExpiresIn != 600 {
		t.Fatalf("access = %+v", access)
	}

	client := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}}
	r, _ := http.NewRequest(http.MethodGet, h.server.URL+"/v1/reports/"+id+"/files/report.json", nil)
	redirect, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer redirect.Body.Close()
	if redirect.StatusCode != http.StatusFound {
		t.Fatalf("status = %d, want a redirect", redirect.StatusCode)
	}
	if got := redirect.Header.Get("Location"); !strings.Contains(got, "reports/"+id+"/report.json") {
		t.Fatalf("redirected to %q", got)
	}

	// A raw chunk is never served here, whoever asks.
	r, _ = http.NewRequest(http.MethodGet, h.server.URL+"/v1/reports/"+id+"/files/raw/0.zst", nil)
	refused, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	defer refused.Body.Close()
	if refused.StatusCode != http.StatusNotFound {
		t.Fatalf("a raw chunk = %d, want 404", refused.StatusCode)
	}
}

func TestReadablePathAllowsOnlyThePageFiles(t *testing.T) {
	for path, want := range map[string]bool{
		"report.json":              true,
		"fights/0/summary.json":    true,
		"fights/12/events.parquet": true,
		"fights/3/live.json":       true,
		"raw/0.zst":                false,
		"fights/0/raw.zst":         false,
		"fights/x/summary.json":    false,
		"../../other/report.json":  false,
		"":                         false,
		"fights/0":                 false,
		"report.json/../raw/0.zst": false,
	} {
		if got := readablePath(path); got != want {
			t.Errorf("readablePath(%q) = %v, want %v", path, got, want)
		}
	}
}

func TestTheCardRendersAPNG(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	res := h.do(http.MethodGet, "/reports/"+id+"/card.png", "", nil)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "image/png" {
		t.Fatalf("content-type = %q", got)
	}
	if got := res.Header.Get("Cache-Control"); got != "public, max-age=300" {
		t.Fatalf("cache-control = %q", got)
	}
	if got := res.Header.Get("Vary"); got != "Cookie" {
		t.Fatalf("vary = %q, want Cookie", got)
	}
	head := make([]byte, 8)
	if _, err := res.Body.Read(head); err != nil {
		t.Fatal(err)
	}
	if string(head[1:4]) != "PNG" {
		t.Fatalf("body does not start with a PNG header: %q", head)
	}
}

func TestTheCardOfAnUnknownReportIsTheStandIn(t *testing.T) {
	h := newHarness(t)
	res := h.do(http.MethodGet, "/reports/nosuchreport/card.png", "", nil)
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "image/png" {
		t.Fatalf("content-type = %q, want a stand-in card", got)
	}
}

func TestValidVisibilityAndRanked(t *testing.T) {
	for _, v := range Visibilities {
		if !ValidVisibility(v) {
			t.Errorf("%q should be valid", v)
		}
	}
	if ValidVisibility("secret") {
		t.Error("secret is not a visibility")
	}
	if Ranked(Private) || !Ranked(Public) || !Ranked(GuildTo) {
		t.Error("a private report is the only one that is never ranked")
	}
}

func TestUnknownReportsAre404Everywhere(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for _, path := range []string{
		"/v1/reports/nosuchreport",
		"/v1/reports/nosuchreport/access",
		"/v1/reports/nosuchreport/files/report.json",
	} {
		res := h.do(http.MethodGet, path, "", nil)
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("GET %s = %d, want 404", path, res.StatusCode)
		}
	}
	res := h.json(http.MethodPatch, "/v1/reports/nosuchreport", `{"title":"x"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("PATCH = %d, want 404", res.StatusCode)
	}
}

// decodeInto is a small helper for tests that read a raw JSON body.
func decodeInto(t *testing.T, body []byte, into any) {
	t.Helper()
	if err := json.Unmarshal(body, into); err != nil {
		t.Fatal(err)
	}
}

func TestTheOwnReportsListIsPagedAndCounted(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	first := h.createReport(Public)
	second := h.createReport(Private)
	h.reportID = second
	h.seedFight(1, 9001, 0, 100)
	h.seedFight(2, 0, 100, 200)

	res := h.do(http.MethodGet, "/v1/reports?mine=1", "", nil)
	var page MinePage
	h.data(res, &page)
	if page.Total != 2 || page.PerPage != MinePerPage || page.Page != 1 {
		t.Fatalf("page = %+v", page)
	}
	if len(page.Rows) != 2 {
		t.Fatalf("rows = %+v", page.Rows)
	}
	newest := page.Rows[0]
	if newest.ID != second || newest.Status != StatusLive || newest.Visibility != Private {
		t.Fatalf("newest = %+v", newest)
	}
	if newest.FightCount != 2 || newest.KillCount != 1 {
		t.Fatalf("counts = %d fights and %d kills, want 2 and 1 (the trash fight is not a kill)",
			newest.FightCount, newest.KillCount)
	}
	if page.Rows[1].ID != first || page.Rows[1].FightCount != 0 || page.Rows[1].KillCount != 0 {
		t.Fatalf("oldest = %+v", page.Rows[1])
	}

	res = h.do(http.MethodGet, "/v1/reports?mine=1&page=2", "", nil)
	h.data(res, &page)
	if len(page.Rows) != 0 || page.Page != 2 {
		t.Fatalf("second page = %+v", page)
	}

	// Someone else's list is their own, not this one.
	h.actor = auth.Actor{UserID: h.owner + 42, Role: "user", Method: "session"}
	res = h.do(http.MethodGet, "/v1/reports?mine=1", "", nil)
	h.data(res, &page)
	if page.Total != 0 {
		t.Fatalf("a stranger sees %d reports, want none", page.Total)
	}
}

func TestTheOwnReportsListNeedsMineAndAValidPage(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	for name, path := range map[string]string{
		"no mine":  "/v1/reports",
		"bad page": "/v1/reports?mine=1&page=0",
	} {
		res := h.do(http.MethodGet, path, "", nil)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
	h.actor = auth.Actor{UserID: h.owner, Method: "device", DeviceID: "d"}
	res := h.do(http.MethodGet, "/v1/reports?mine=1", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("a device listing = %d, want 401", res.StatusCode)
	}
}

// setOwner changes a column on the harness owner's account row.
func (h *harness) setOwner(column string, value any) {
	h.t.Helper()
	if _, err := h.store.Pool.Exec(context.Background(),
		`update users set `+column+` = $2 where id = $1`, h.owner, value); err != nil {
		h.t.Fatal(err)
	}
}

// A public report is readable by anyone holding the link, and the
// contract puts its owner in the page's Open Graph tags. The harness
// owner signed in by email and so has no battletag: they must read as
// a pseudonym, and their address must appear nowhere in the body.
func TestAPublicReportNamesItsOwnerWithoutPublishingTheirEmail(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)

	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "raider@example.com") {
		t.Fatalf("the owner's email address is in a public report body: %s", body)
	}
	var env struct {
		Data View `json:"data"`
	}
	if err := json.Unmarshal(body, &env); err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("user-%d", h.owner)
	if env.Data.Owner == nil || env.Data.Owner.Battletag != want {
		t.Fatalf("owner = %+v, want the pseudonym %q", env.Data.Owner, want)
	}
}

// The anonymize flag governs this one read path, and it is the only
// place the API promises it applies.
func TestAnAnonymizedOwnerReadsAsThePseudonym(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.setOwner("battletag", "Baelgrim#1234")

	res := h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	var view View
	h.data(res, &view)
	if view.Owner == nil || view.Owner.Battletag != "Baelgrim#1234" {
		t.Fatalf("owner = %+v, want the battletag", view.Owner)
	}

	h.setOwner("anonymize", true)
	res = h.do(http.MethodGet, "/v1/reports/"+id, "", nil)
	h.data(res, &view)
	want := fmt.Sprintf("user-%d", h.owner)
	if view.Owner == nil || view.Owner.Battletag != want {
		t.Fatalf("owner = %+v, want the pseudonym %q", view.Owner, want)
	}
}

// Making a report private takes the report body away but used to leave
// every fight_metrics row in place, so /v1/rankings kept publishing the
// player keys, names, guild and numbers of a report that now 404s.
func TestPatchingAReportOutOfTheRankingsWithdrawsItsRows(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	id := h.createReport(Public)

	// Unlisted still ranks, so nothing is withdrawn on the way there.
	res := h.json(http.MethodPatch, "/v1/reports/"+id, `{"visibility":"unlisted"}`)
	h.data(res, nil)
	if removed, _ := h.ranker.withdrawals(); len(removed) != 0 {
		t.Fatalf("withdrawn = %v, want nothing: unlisted still ranks", removed)
	}

	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"visibility":"private"}`)
	h.data(res, nil)
	removed, reasons := h.ranker.withdrawals()
	if len(removed) != 1 || removed[0] != id {
		t.Fatalf("withdrawn = %v, want the report once", removed)
	}
	if reasons[0] != ReasonNotRankable {
		t.Fatalf("reason = %q, want %q", reasons[0], ReasonNotRankable)
	}

	// Only the crossing withdraws: patching a report that is already
	// private must not write a second withdrawal.
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"Wednesday"}`)
	h.data(res, nil)
	res = h.json(http.MethodPatch, "/v1/reports/"+id, `{"visibility":"private"}`)
	h.data(res, nil)
	if removed, _ := h.ranker.withdrawals(); len(removed) != 1 {
		t.Fatalf("withdrawn = %v, want exactly one withdrawal", removed)
	}
}

// A card drawn for a private or guild report passed mayView for one
// caller. api.foreversixty.gg sits behind Cloudflare, which caches
// .png by extension, so a public Cache-Control would let the owner's
// own fetch fill a shared entry served to everyone else.
func TestTheCardOfAGatedReportIsNotSharedCacheable(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	for _, c := range []struct {
		visibility string
		want       string
	}{
		{Public, "public, max-age=300"},
		{Unlisted, "public, max-age=300"},
		{Private, "private, max-age=300"},
		{GuildTo, "private, max-age=300"},
	} {
		t.Run(c.visibility, func(t *testing.T) {
			id := h.createReport(c.visibility)
			res := h.do(http.MethodGet, "/reports/"+id+"/card.png", "", nil)
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				t.Fatalf("status = %d", res.StatusCode)
			}
			if got := res.Header.Get("Cache-Control"); got != c.want {
				t.Fatalf("cache-control = %q, want %q", got, c.want)
			}
			if got := res.Header.Get("Vary"); got != "Cookie" {
				t.Fatalf("vary = %q, want Cookie", got)
			}
		})
	}
}
