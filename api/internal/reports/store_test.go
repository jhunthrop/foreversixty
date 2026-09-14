package reports

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

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
	if fights[0].Deaths != 1 || fights[0].NPCKills != 2 {
		t.Fatalf("counts = %+v", fights[0])
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
