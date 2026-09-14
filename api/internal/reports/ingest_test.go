package reports

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/klauspost/compress/zstd"
)

// fakeRanker records the fights handed to the rankings store.
type fakeRanker struct {
	mu      sync.Mutex
	written []RankedFight
	removed []string
}

func (f *fakeRanker) WriteFight(_ context.Context, rf RankedFight) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.written = append(f.written, rf)
	return nil
}

func (f *fakeRanker) RemoveReport(_ context.Context, id string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.removed = append(f.removed, id)
	return nil
}

func (f *fakeRanker) fights() []RankedFight {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]RankedFight{}, f.written...)
}

// fakeSampler records the reports queued for a raw-sample check.
type fakeSampler struct {
	mu        sync.Mutex
	scheduled []string
}

func (f *fakeSampler) Schedule(id string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.scheduled = append(f.scheduled, id)
}

func (f *fakeSampler) all() []string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]string{}, f.scheduled...)
}

// bundle is a fight bundle exactly as the companion posts one.
type bundle struct {
	body        *bytes.Buffer
	contentType string
	rows        []metrics.Row
	sha         string
}

// makeBundle builds a multipart bundle from the engine fixture,
// letting a test bend the metrics before they are posted.
func makeBundle(t *testing.T, index int, bend func([]metrics.Row)) bundle {
	t.Helper()
	fx, err := engine.NewFixture("fixture")
	if err != nil {
		t.Fatal(err)
	}
	rows := metrics.Derive(fx.Fight, fx.Summary)
	if bend != nil {
		bend(rows)
	}
	sum := sha256.Sum256(fx.Log)
	raw := RawRange{StartOffset: 0, EndOffset: int64(len(fx.Log)), SHA256: hex.EncodeToString(sum[:])}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	writeJSON := func(name string, v any) {
		part, err := w.CreateFormFile(name, name+".json")
		if err != nil {
			t.Fatal(err)
		}
		if err := json.NewEncoder(part).Encode(v); err != nil {
			t.Fatal(err)
		}
	}
	writeJSON("summary", fx.Summary)
	writeJSON("metrics", rows)
	writeJSON("raw_range", raw)
	events, err := w.CreateFormFile("events", "events.parquet")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := events.Write(fx.Parquet); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return bundle{body: &buf, contentType: w.FormDataContentType(), rows: rows, sha: raw.SHA256}
}

// makeBundleWithout builds a bundle with one part left out.
func makeBundleWithout(t *testing.T, drop string) bundle {
	t.Helper()
	return buildBundle(t, func(parts map[string][]byte) { delete(parts, drop) })
}

// makeBundleWith builds a bundle with one part replaced.
func makeBundleWith(t *testing.T, bend func(map[string][]byte)) bundle {
	t.Helper()
	return buildBundle(t, bend)
}

// buildBundle assembles a multipart body from named parts, after the
// caller has bent them.
func buildBundle(t *testing.T, bend func(map[string][]byte)) bundle {
	t.Helper()
	fx, err := engine.NewFixture("fixture")
	if err != nil {
		t.Fatal(err)
	}
	rows := metrics.Derive(fx.Fight, fx.Summary)
	sum := sha256.Sum256(fx.Log)
	raw := RawRange{StartOffset: 0, EndOffset: int64(len(fx.Log)), SHA256: hex.EncodeToString(sum[:])}
	parts := map[string][]byte{
		"summary":   mustJSON(t, fx.Summary),
		"metrics":   mustJSON(t, rows),
		"raw_range": mustJSON(t, raw),
		"events":    fx.Parquet,
	}
	if bend != nil {
		bend(parts)
	}
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	for _, name := range []string{"summary", "metrics", "raw_range", "events"} {
		body, ok := parts[name]
		if !ok {
			continue
		}
		part, err := w.CreateFormFile(name, name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := part.Write(body); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return bundle{body: &buf, contentType: w.FormDataContentType(), rows: rows, sha: raw.SHA256}
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// brokenPutter is an object store that is down.
type brokenPutter struct{}

func (brokenPutter) Put(context.Context, string, []byte, store.PutOptions) error {
	return errors.New("r2 is down")
}

func TestAVerifiedBundleIsStoredPublishedAndRanked(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	b := makeBundle(t, 1, nil)

	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	var out struct {
		FightIndex int  `json:"fight_index"`
		Verified   bool `json:"verified"`
	}
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", res.StatusCode)
	}
	h.data(res, &out)
	if out.FightIndex != 1 || !out.Verified {
		t.Fatalf("out = %+v", out)
	}

	keys := store.Keys{ReportID: id}
	for _, key := range []string{keys.FightSummary(1), keys.FightEvents(1), keys.Report()} {
		if !h.fileExists(key) {
			t.Errorf("object %s was not written", key)
		}
	}
	var report store.Report
	decodeInto(t, h.readFile(keys.Report()), &report)
	if len(report.Fights) != 1 || report.Fights[0].Index != 1 {
		t.Fatalf("report.json = %+v", report)
	}

	fights, err := h.store.Fights(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if len(fights) != 1 || !fights[0].Verified || !fights[0].Kill {
		t.Fatalf("fights = %+v", fights)
	}
	if fights[0].EncounterID != 9001 || fights[0].Name == "" {
		t.Fatalf("fight = %+v, want the encounter named", fights[0])
	}

	ranked := h.ranker.fights()
	if len(ranked) != 1 {
		t.Fatalf("ranked = %+v", ranked)
	}
	if ranked[0].Region != "us" || ranked[0].Ruleset != "hardcore" {
		t.Fatalf("ranked under %s/%s, want the logging character's", ranked[0].Region, ranked[0].Ruleset)
	}
	if len(ranked[0].Rows) != 3 {
		t.Fatalf("ranked rows = %d, want one per player", len(ranked[0].Rows))
	}
}

func TestResendingTheSameBundleIsAccepted(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	first := makeBundle(t, 1, nil)
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", first.contentType, first.body)
	res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("first = %d", res.StatusCode)
	}
	again := makeBundle(t, 1, nil)
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", again.contentType, again.body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("second = %d, want 200", res.StatusCode)
	}
	if got := len(h.ranker.fights()); got != 1 {
		t.Fatalf("the fight was ranked %d times, want once", got)
	}
}

func TestAnInflatedMetricIsRefusedAndFlagsTheReport(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	b := makeBundle(t, 1, func(rows []metrics.Row) {
		for i := range rows {
			if rows[i].MetricDPS > 0 {
				rows[i].MetricDPS *= 3
			}
		}
	})
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("status = %d, want 409", res.StatusCode)
	}
	if got := h.errorFields(res)["metrics"]; got == "" || !strings.Contains(got, "metric_dps") {
		t.Fatalf("fields.metrics = %q, want the field that disagreed", got)
	}

	fights, err := h.store.Fights(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if len(fights) != 1 || fights[0].Verified {
		t.Fatalf("the fight should be stored unverified: %+v", fights)
	}
	rep, err := h.store.Get(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Flagged == nil || *rep.Flagged != "metrics_mismatch" {
		t.Fatalf("flagged = %v, want metrics_mismatch", rep.Flagged)
	}
	if got := len(h.ranker.fights()); got != 0 {
		t.Fatalf("an unverified fight was ranked %d times", got)
	}
	if h.fileExists(store.Keys{ReportID: id}.FightEvents(1)) {
		t.Fatal("an unverified fight's events must not be published")
	}
}

func TestABundleForSomeoneElsesReportIsRefused(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.actor = auth.Actor{UserID: h.owner + 7, Method: "device", DeviceID: "other"}
	b := makeBundle(t, 1, nil)
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", res.StatusCode)
	}
}

func TestABrokenBundleIsRefused(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", "multipart/form-data; boundary=x",
		strings.NewReader("not multipart at all"))
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}
	b := makeBundle(t, 1, nil)
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/notanumber", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a non-numeric index = %d, want 400", res.StatusCode)
	}
}

func TestALiveSnapshotIsWrittenAndTheIndexTouched(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	fx, err := engine.NewFixture(id)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(LiveInput{Summary: fx.Summary, ElapsedMS: 12000, UpdatedAt: fx.Fight.Start})
	if err != nil {
		t.Fatal(err)
	}
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1/live", "application/json", bytes.NewReader(body))
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	keys := store.Keys{ReportID: id}
	if !h.fileExists(keys.FightLive(1)) || !h.fileExists(keys.Report()) {
		t.Fatal("live.json and report.json should both have been written")
	}
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1/live", "application/json", strings.NewReader("{"))
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a broken snapshot = %d, want 400", res.StatusCode)
	}
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/notanumber/live", "application/json",
		bytes.NewReader(body))
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a non-numeric index = %d, want 400", res.StatusCode)
	}
}

// packRaw compresses a raw chunk the way the companion does.
func packRaw(t *testing.T, text string) ([]byte, string) {
	t.Helper()
	var buf bytes.Buffer
	w, err := zstd.NewWriter(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write([]byte(text)); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	return buf.Bytes(), hex.EncodeToString(sum[:])
}

func (h *harness) putRaw(id string, offset int, packed []byte, sha string) *http.Response {
	h.t.Helper()
	r, err := http.NewRequest(http.MethodPut,
		h.server.URL+"/v1/reports/"+id+"/raw?offset="+itoaFull(offset), bytes.NewReader(packed))
	if err != nil {
		h.t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/zstd")
	r.Header.Set("X-Raw-SHA256", sha)
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		h.t.Fatal(err)
	}
	return res
}

func itoaFull(n int) string {
	if n == 0 {
		return "0"
	}
	digits := ""
	for n > 0 {
		digits = string(rune('0'+n%10)) + digits
		n /= 10
	}
	return digits
}

func TestRawChunksAreVerifiedStoredAndIdempotent(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	text := "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1\n"
	packed, sha := packRaw(t, text)

	res := h.putRaw(id, 0, packed, sha)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	if !h.fileExists(store.Keys{ReportID: id}.Raw(0)) {
		t.Fatal("the chunk was not stored")
	}

	res = h.putRaw(id, 0, packed, sha)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("re-sending = %d, want 200", res.StatusCode)
	}

	other, otherSHA := packRaw(t, "9/26 20:10:01.000  ZONE_CHANGE,2284,\"Elsewhere\",8\n")
	res = h.putRaw(id, 0, other, otherSHA)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("a different chunk at the same offset = %d, want 409", res.StatusCode)
	}

	res = h.putRaw(id, 0, packed, "not-a-hash")
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a bad hash header = %d, want 400", res.StatusCode)
	}

	wrongSHA := strings.Repeat("ab", 32)
	res = h.putRaw(id, 1000, packed, wrongSHA)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a chunk that does not hash to its header = %d, want 400", res.StatusCode)
	}

	res = h.do(http.MethodPut, "/v1/reports/"+id+"/raw?offset=x", "application/zstd", bytes.NewReader(packed))
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a bad offset = %d, want 400", res.StatusCode)
	}

	res = h.putRaw(id, 2000, []byte("not zstd"), sha)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a body that is not zstd = %d, want 400", res.StatusCode)
	}
}

func TestAnOverlappingRangeIsRefused(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	first, firstSHA := packRaw(t, strings.Repeat("a", 100))
	res := h.putRaw(id, 0, first, firstSHA)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d", res.StatusCode)
	}
	second, secondSHA := packRaw(t, strings.Repeat("b", 100))
	res = h.putRaw(id, 50, second, secondSHA)
	res.Body.Close()
	if res.StatusCode != http.StatusConflict {
		t.Fatalf("an overlapping range = %d, want 409", res.StatusCode)
	}
}

func TestCompleteFinishesTheReportAndSchedulesTheCheck(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	res := h.json(http.MethodPost, "/v1/reports/"+id+"/complete",
		`{"final_offset":1024,"engine_version":"0.1.0","health":{"lines":33,"parse_errors":0}}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", res.StatusCode)
	}
	rep, err := h.store.Get(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != StatusComplete || rep.EngineVersion != "0.1.0" || rep.CompletedAt == nil {
		t.Fatalf("report = %+v", rep)
	}
	if len(rep.Health) == 0 || !strings.Contains(string(rep.Health), `"lines": 33`) {
		t.Fatalf("health = %s", rep.Health)
	}
	if got := h.sampler.all(); len(got) != 1 || got[0] != id {
		t.Fatalf("scheduled = %v, want the report", got)
	}
	var report store.Report
	decodeInto(t, h.readFile(store.Keys{ReportID: id}.Report()), &report)
	if report.Health.Lines != 33 {
		t.Fatalf("report.json health = %+v", report.Health)
	}

	res = h.json(http.MethodPost, "/v1/reports/"+id+"/complete", `{`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a broken body = %d, want 400", res.StatusCode)
	}
}

func TestIngestRoutesNeedADeviceToken(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.actor = auth.Actor{UserID: h.owner, Role: "user", Method: "session"}
	res := h.json(http.MethodPost, "/v1/reports/"+id+"/complete", `{"final_offset":1}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 without a device token", res.StatusCode)
	}
}

func TestAFailingObjectStoreIs500(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.ingest.Put = brokenPutter{}
	b := makeBundle(t, 1, nil)
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", res.StatusCode)
	}
}

func TestABundleMissingAPartIsRefused(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	for _, drop := range []string{"summary", "metrics", "raw_range", "events"} {
		b := makeBundleWithout(t, drop)
		res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("a bundle with no %s part = %d, want 400", drop, res.StatusCode)
		}
	}
}

func TestABundleWithUnreadableEventsIsRefused(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	b := makeBundleWith(t, func(parts map[string][]byte) {
		parts["events"] = []byte("not parquet")
	})
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", res.StatusCode)
	}

	b = makeBundleWith(t, func(parts map[string][]byte) {
		parts["metrics"] = []byte("[]")
	})
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("an empty metrics part = %d, want 400", res.StatusCode)
	}

	b = makeBundleWith(t, func(parts map[string][]byte) {
		parts["raw_range"] = []byte(`{"start_offset":0,"end_offset":1,"sha256":"nope"}`)
	})
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a bad raw hash = %d, want 400", res.StatusCode)
	}

	b = makeBundleWith(t, func(parts map[string][]byte) {
		parts["summary"] = []byte("{{{")
	})
	res = h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a broken summary part = %d, want 400", res.StatusCode)
	}
}

func TestReportRealmReadsTheLoggingCharacter(t *testing.T) {
	key := "eu/pvp/baelgrim"
	region, ruleset := ReportRealm(Report{LoggingCharacter: &key})
	if region != "eu" || ruleset != "pvp" {
		t.Fatalf("realm = %q/%q", region, ruleset)
	}
	if region, ruleset := ReportRealm(Report{}); region != "us" || ruleset != "normal" {
		t.Fatalf("default = %q/%q, want us/normal", region, ruleset)
	}
	bad := "nonsense"
	if region, ruleset := ReportRealm(Report{LoggingCharacter: &bad}); region != "us" || ruleset != "normal" {
		t.Fatalf("a broken key = %q/%q, want the defaults", region, ruleset)
	}
}

// ingestCalls is one request to each of the four ingest routes, for the
// tests that check how they all answer the same broken world.
func (h *harness) ingestCalls(id string) map[string]*http.Response {
	h.t.Helper()
	b := makeBundle(h.t, 1, nil)
	packed, sha := packRaw(h.t, "9/26 20:10:00.000  ZONE_CHANGE,2284,\"Nowhere\",8\n")
	return map[string]*http.Response{
		"fight":    h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body),
		"live":     h.json(http.MethodPut, "/v1/reports/"+id+"/fights/1/live", `{"elapsed_ms":1}`),
		"raw":      h.putRaw(id, 0, packed, sha),
		"complete": h.json(http.MethodPost, "/v1/reports/"+id+"/complete", `{"final_offset":1}`),
	}
}

func TestEveryIngestRouteIs404ForAnUnknownReport(t *testing.T) {
	h := newHarness(t)
	for name, res := range h.ingestCalls("nosuchreport") {
		res.Body.Close()
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("%s = %d, want 404", name, res.StatusCode)
		}
	}
}

func TestEveryIngestRouteAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Public)
	h.store.Pool.Close()
	for name, res := range h.ingestCalls(id) {
		res.Body.Close()
		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("%s = %d, want 500", name, res.StatusCode)
		}
	}
}

func TestAPrivateReportsFightIsVerifiedAndStoredButNeverRanked(t *testing.T) {
	h := newHarness(t)
	id := h.createReport(Private)
	b := makeBundle(t, 1, nil)
	res := h.do(http.MethodPut, "/v1/reports/"+id+"/fights/1", b.contentType, b.body)
	res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		t.Fatalf("status = %d, want 201", res.StatusCode)
	}
	fights, err := h.store.Fights(t.Context(), id)
	if err != nil {
		t.Fatal(err)
	}
	if len(fights) != 1 || !fights[0].Verified {
		t.Fatalf("fights = %+v, want the fight stored verified", fights)
	}
	if got := h.ranker.fights(); len(got) != 0 {
		t.Fatalf("a private report was ranked: %+v", got)
	}
}

func TestMismatchFieldFallsBackWhenTheErrorIsNotAMismatch(t *testing.T) {
	if got := mismatchField(errors.New("something else went wrong")); got != "rows" {
		t.Fatalf("field = %q, want the rows fallback", got)
	}
}
