package reports

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

func TestStoreReadsBackFightsPlayersAndRanges(t *testing.T) {
	h := newHarness(t)
	h.reportID = h.createReport(Public)
	h.seedFight(1, 9001, 0, 500)
	h.seedFight(2, 0, 500, 900)

	fights, err := h.store.Fights(t.Context(), h.reportID)
	if err != nil {
		t.Fatal(err)
	}
	if len(fights) != 2 {
		t.Fatalf("fights = %d", len(fights))
	}
	if fights[0].Kind != "encounter" || fights[1].Kind != "trash" {
		t.Fatalf("kinds = %q, %q", fights[0].Kind, fights[1].Kind)
	}
	if c := fights[0]; c.Deaths != 1 || deref(c.MobDeaths) != 3 || deref(c.GroupKills) != 2 ||
		deref(c.GroupDeaths) != 1 || len(c.Group) != 1 || c.Group[0] != "Player-1" {
		t.Fatalf("counts = %+v", c)
	}
	if !fights[0].End.After(fights[0].Start) {
		t.Fatalf("end should follow start: %+v", fights[0])
	}

	players, err := h.store.Players(t.Context(), h.reportID)
	if err != nil {
		t.Fatal(err)
	}
	if len(players) != 2 || players[0] != "Player-1" {
		t.Fatalf("players = %v", players)
	}

	start, end, err := h.store.FightRawRange(t.Context(), h.reportID, 1)
	if err != nil || start == nil || end == nil || *end != 500 {
		t.Fatalf("range = %v, %v, %v", start, end, err)
	}
	if _, _, err := h.store.FightRawRange(t.Context(), h.reportID, 99); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}

	// Writing the same fight again updates rather than duplicating.
	h.seedFight(1, 9001, 0, 600)
	again, err := h.store.Fights(t.Context(), h.reportID)
	if err != nil || len(again) != 2 {
		t.Fatalf("fights = %d, %v", len(again), err)
	}
}

// A fight indexed before engine 0.6.0 has no counts. The store answers
// them as absent, never as zero, so the page shows no kill count rather
// than a wrong one.
func TestStoreLeavesOutTheCountsOfAFightFromBeforeThem(t *testing.T) {
	h := newHarness(t)
	h.reportID = h.createReport(Public)
	h.seedFight(1, 0, 0, 500)
	if _, err := h.store.Pool.Exec(t.Context(),
		`update fights set mob_deaths = null, group_kills = null, group_deaths = null,
		   group_players = null, npc_kills = 7 where report_id = $1`, h.reportID); err != nil {
		t.Fatal(err)
	}
	fights, err := h.store.Fights(t.Context(), h.reportID)
	if err != nil || len(fights) != 1 {
		t.Fatalf("fights = %v, %v", fights, err)
	}
	f := fights[0]
	if f.MobDeaths != nil || f.GroupKills != nil || f.GroupDeaths != nil || f.Group != nil {
		t.Fatalf("an old fight's counts = %+v, want all absent", f)
	}
	b, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"mob_deaths", "group_kills", "group_deaths", "group", "npc_kills"} {
		if bytes.Contains(b, []byte(`"`+k+`"`)) {
			t.Errorf("an old fight's JSON carries %q: %s", k, b)
		}
	}
}

func deref(p *int) int {
	if p == nil {
		return -1
	}
	return *p
}

func TestStoreRecordsRawChunksInOrder(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	for _, c := range []RawChunk{
		{Start: 0, End: 100, SHA256: []byte("aaaa")},
		{Start: 100, End: 200, SHA256: []byte("bbbb")},
	} {
		stored, err := h.store.PutRawChunk(t.Context(), id, c)
		if err != nil || !stored {
			t.Fatalf("chunk at %d: %v, %v", c.Start, stored, err)
		}
	}
	chunks, err := h.store.RawChunks(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if len(chunks) != 2 || chunks[0].Start != 0 || chunks[1].Start != 100 {
		t.Fatalf("chunks = %+v", chunks)
	}
	if !bytes.Equal(chunks[0].SHA256, []byte("aaaa")) {
		t.Fatalf("hash = %q", chunks[0].SHA256)
	}
}

func TestFlagSetsAndClears(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	if err := h.store.Flag(t.Context(), id, "tampered"); err != nil {
		t.Fatal(err)
	}
	rep, err := h.store.Get(t.Context(), id)
	if err != nil || rep.Flagged == nil || *rep.Flagged != "tampered" {
		t.Fatalf("flagged = %v, %v", rep.Flagged, err)
	}
	if err := h.store.Flag(t.Context(), id, ""); err != nil {
		t.Fatal(err)
	}
	rep, err = h.store.Get(t.Context(), id)
	if err != nil || rep.Flagged != nil {
		t.Fatalf("flagged = %v, %v", rep.Flagged, err)
	}
}

func TestSetStatusOnAnUnknownReport(t *testing.T) {
	h := newHarness(t)
	if err := h.store.SetStatus(t.Context(), "nosuchreport", StatusFailed, "", nil); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestEveryHandlerAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.store.Pool.Close()
	// "complete" is Task 11's ingest-completion route, not one of the
	// eight this task's Mount registers (see the brief's interfaces
	// list) - it belongs here once the ingest handlers land beside
	// these, so it is left out rather than invented ahead of that task.
	for name, res := range map[string]*http.Response{
		"get":        h.do(http.MethodGet, "/v1/reports/"+id, "", nil),
		"visibility": h.do(http.MethodGet, "/v1/reports/"+id+"/visibility", "", nil),
		"create":     h.json(http.MethodPost, "/v1/reports", `{"visibility":"public"}`),
	} {
		res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, res.StatusCode)
		}
	}
	h.asSession()
	for name, res := range map[string]*http.Response{
		"patch":  h.json(http.MethodPatch, "/v1/reports/"+id, `{"title":"x"}`),
		"access": h.do(http.MethodGet, "/v1/reports/"+id+"/access", "", nil),
		"file":   h.do(http.MethodGet, "/v1/reports/"+id+"/files/report.json", "", nil),
		"card":   h.do(http.MethodGet, "/reports/"+id+"/card.png", "", nil),
	} {
		res.Body.Close()
		if name == "card" {
			if res.StatusCode != http.StatusInternalServerError {
				t.Errorf("card = %d, want 500", res.StatusCode)
			}
			continue
		}
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, res.StatusCode)
		}
	}
}

func TestACardWithBossesAndNoSigner(t *testing.T) {
	h := newHarness(t)
	h.reportID = h.createReport(Public)
	h.seedFight(1, 9001, 0, 100)
	h.seedFight(2, 0, 100, 200)
	res := h.do(http.MethodGet, "/reports/"+h.reportID+"/card.png", "", nil)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK || res.Header.Get("Cache-Control") != "public, max-age=300" {
		t.Fatalf("status = %d, cache = %q", res.StatusCode, res.Header.Get("Cache-Control"))
	}

	h.service.Signer = nil
	h.asSession()
	private := h.createReport(Private)
	res = h.do(http.MethodGet, "/v1/reports/"+private+"/files/report.json", "", nil)
	res.Body.Close()
	if res.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("with no signer = %d, want 503", res.StatusCode)
	}
}

func TestAnonymousCallersCannotCreateAReport(t *testing.T) {
	h := newHarness(t)
	h.actor = auth.Actor{}
	res := h.json(http.MethodPost, "/v1/reports", `{"visibility":"public"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", res.StatusCode)
	}
}

// The tests below cover the Store methods this task writes but whose
// callers - the ingest route and the whole-file upload route - are
// Tasks 11 and 12. Direct store-level coverage here means this task's
// own surface is tested without reaching into those routes.

func TestFightSHAReadsBackTheStoredHashAndWhetherItVerified(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	upsert := func(verified bool) {
		t.Helper()
		if _, err := h.store.UpsertFight(t.Context(), FightRecord{
			ReportID: id, Index: 1, Name: "Trash", Players: []string{},
			RawSHA256: []byte("hash-bytes"), Verified: verified,
		}); err != nil {
			t.Fatal(err)
		}
	}

	upsert(false)
	sha, verified, err := h.store.FightSHA(t.Context(), id, 1)
	if err != nil || string(sha) != "hash-bytes" {
		t.Fatalf("sha = %q, %v", sha, err)
	}
	if verified {
		t.Fatal("a fight stored unverified must not read back verified")
	}

	upsert(true)
	if _, verified, err = h.store.FightSHA(t.Context(), id, 1); err != nil || !verified {
		t.Fatalf("verified = %v, %v, want true", verified, err)
	}

	if _, _, err := h.store.FightSHA(t.Context(), id, 99); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestPutRawChunkDetectsOverlapsAndHashMismatches(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	if _, err := h.store.PutRawChunk(t.Context(), id, RawChunk{Start: 0, End: 100, SHA256: []byte("aaaa")}); err != nil {
		t.Fatal(err)
	}
	// Same start, a different hash: the stored chunk does not match.
	if _, err := h.store.PutRawChunk(t.Context(), id, RawChunk{Start: 0, End: 100, SHA256: []byte("bbbb")}); !errors.Is(err, ErrOverlap) {
		t.Fatalf("err = %v, want ErrOverlap for a hash mismatch", err)
	}
	// A different start that still overlaps the stored range.
	if _, err := h.store.PutRawChunk(t.Context(), id, RawChunk{Start: 50, End: 150, SHA256: []byte("cccc")}); !errors.Is(err, ErrOverlap) {
		t.Fatalf("err = %v, want ErrOverlap for an overlapping range", err)
	}
}

func TestDeleteRawChunkFreesTheOffsetAgain(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	if _, err := h.store.PutRawChunk(t.Context(), id, RawChunk{Start: 0, End: 100, SHA256: []byte("aaaa")}); err != nil {
		t.Fatal(err)
	}
	if err := h.store.DeleteRawChunk(t.Context(), id, 0); err != nil {
		t.Fatal(err)
	}
	chunks, err := h.store.RawChunks(t.Context(), id)
	if err != nil || len(chunks) != 0 {
		t.Fatalf("chunks = %+v, %v, want the row gone", chunks, err)
	}
	// The offset is free again, and free for different bytes: the whole
	// point is that nothing is recorded as held.
	if stored, err := h.store.PutRawChunk(t.Context(), id,
		RawChunk{Start: 0, End: 100, SHA256: []byte("bbbb")}); err != nil || !stored {
		t.Fatalf("stored = %v, %v, want the offset writable again", stored, err)
	}
	// Removing a chunk that is not there is not an error.
	if err := h.store.DeleteRawChunk(t.Context(), id, 9999); err != nil {
		t.Fatalf("err = %v, want removing an absent chunk to be a no-op", err)
	}
}

func TestUploadLifecycle(t *testing.T) {
	h := newHarness(t)
	u := Upload{
		ID: "up1", UserID: &h.owner, ObjectKey: "uploads/up1/raw.txt.zst",
		R2UploadID: "r2-upload-1", SizeBytes: 4096, Filename: "WoWCombatLog.txt",
	}
	if err := h.store.CreateUpload(t.Context(), u); err != nil {
		t.Fatal(err)
	}
	got, err := h.store.Upload(t.Context(), u.ID)
	if err != nil || got.ObjectKey != u.ObjectKey || got.ReportID != nil {
		t.Fatalf("upload = %+v, %v", got, err)
	}
	reportID := h.createReport(Public)
	if err := h.store.FinishUpload(t.Context(), u.ID, reportID); err != nil {
		t.Fatal(err)
	}
	got, err = h.store.Upload(t.Context(), u.ID)
	if err != nil || got.ReportID == nil || *got.ReportID != reportID {
		t.Fatalf("finished upload = %+v, %v", got, err)
	}
	if _, err := h.store.Upload(t.Context(), "nosuchupload"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestKeysBuildsTheReportPrefix(t *testing.T) {
	k := Keys("abc123def456")
	if got := k.Report(); got != "reports/abc123def456/report.json" {
		t.Fatalf("report key = %q", got)
	}
}

// TestForGuildListsAnAttributedReport is the guild tab's rule for whose
// report a guild's is: the one an officer set by hand (guild_id) or the
// one its logger's verified standing named (matched_guild_id), read as
// the matcher reads it, coalesce(guild_id, matched_guild_id). A member's
// public upload nobody has claimed by hand is the guild's report all the
// same; a report with neither column set is nobody's.
func TestForGuildListsAnAttributedReport(t *testing.T) {
	h := newHarness(t)
	guildID := h.seedGuild("Attributed", "hardcore")
	now := time.Now().UTC()

	h.seedReport("byhand", "Set by an officer", "Molten Core", GuildTo, StatusComplete, now.Add(-2*time.Hour))
	h.attachReportToGuild("byhand", guildID)
	h.seedReport("bystanding", "A member's own upload", "Molten Core", Public, StatusComplete, now.Add(-time.Hour))
	h.attributeReportToGuild("bystanding", guildID)
	// Attributed but private: the visibility rule still hides it.
	h.seedReport("standingpriv", "A member's private upload", "Molten Core", Private, StatusComplete, now.Add(-30*time.Minute))
	h.attributeReportToGuild("standingpriv", guildID)
	// Neither column set: nobody's.
	h.seedReport("nobodys", "Unattributed", "Molten Core", Public, StatusComplete, now)

	rows, total, err := h.store.ForGuild(t.Context(), guildID, 1, MinePerPage)
	if err != nil {
		t.Fatal(err)
	}
	if total != 2 {
		t.Fatalf("total = %d, want 2 (one by hand, one by standing)", total)
	}
	if len(rows) != 2 || rows[0].ID != "bystanding" || rows[1].ID != "byhand" {
		t.Fatalf("rows = %+v, want bystanding then byhand", rows)
	}
}

// TestForGuildReadsAnOfficersGuildOverTheMatchersGuess is the coalesce
// order: a report an officer set to one guild never lists under the
// guild its logger's standing named.
func TestForGuildReadsAnOfficersGuildOverTheMatchersGuess(t *testing.T) {
	h := newHarness(t)
	officers, standing := h.seedGuild("Officers", "hardcore"), h.seedGuild("Standing", "hardcore")
	h.seedReport("moved", "Moved by an officer", "Molten Core", Public, StatusComplete, time.Now().UTC())
	h.attachReportToGuild("moved", officers)
	h.attributeReportToGuild("moved", standing)

	rows, total, err := h.store.ForGuild(t.Context(), standing, 1, MinePerPage)
	if err != nil {
		t.Fatal(err)
	}
	if total != 0 || len(rows) != 0 {
		t.Fatalf("the standing guild lists %d rows (%+v), want none", total, rows)
	}
	rows, total, err = h.store.ForGuild(t.Context(), officers, 1, MinePerPage)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 || len(rows) != 1 || rows[0].ID != "moved" {
		t.Fatalf("the officers' guild lists %d rows (%+v), want moved", total, rows)
	}
}
