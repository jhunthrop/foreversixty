package parse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// memObjects is an object store in a map: the Putter the publisher
// writes through and the reader the job streams the upload from.
type memObjects struct {
	mu      sync.Mutex
	objects map[string][]byte
	getErr  error
}

func newMemObjects() *memObjects { return &memObjects{objects: map[string][]byte{}} }

func (m *memObjects) Put(_ context.Context, key string, body []byte, _ store.PutOptions) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.objects[key] = append([]byte{}, body...)
	return nil
}

func (m *memObjects) Get(_ context.Context, key string) (io.ReadCloser, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.getErr != nil {
		return nil, m.getErr
	}
	b, ok := m.objects[key]
	if !ok {
		return nil, errors.New("no such object: " + key)
	}
	return io.NopCloser(bytes.NewReader(b)), nil
}

func (m *memObjects) get(key string) ([]byte, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	b, ok := m.objects[key]
	return b, ok
}

// fakeRanker records what the job ranked.
type fakeRanker struct {
	mu      sync.Mutex
	written []reports.RankedFight
	removed []string
}

func (f *fakeRanker) WriteFight(_ context.Context, rf reports.RankedFight) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.written = append(f.written, rf)
	return nil
}

func (f *fakeRanker) RemoveReport(_ context.Context, id, _ string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.removed = append(f.removed, id)
	return nil
}

func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set; start docker-compose.test.yml")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if _, err := pool.Exec(context.Background(), `truncate users, reports cascade`); err != nil {
		t.Fatal(err)
	}
	return pool
}

// fixtureUpload seeds a report whose upload holds the fixture log,
// compressed when packed is true.
func fixtureUpload(t *testing.T, packed bool) (Deps, string, *memObjects, *fakeRanker) {
	t.Helper()
	pool := testPool(t)
	ctx := context.Background()
	accounts := &auth.Store{Pool: pool}
	owner, err := accounts.UpsertEmailUser(ctx, "raider@example.com")
	if err != nil {
		t.Fatal(err)
	}
	rs := &reports.Store{Pool: pool}
	uploadID := auth.Base32ID(16)
	key := r2.UploadKey(uploadID)
	if err := rs.CreateUpload(ctx, reports.Upload{
		ID: uploadID, UserID: &owner.ID, ObjectKey: key, R2UploadID: "r2-1",
		SizeBytes: 1024, Filename: "WoWCombatLog.txt",
	}); err != nil {
		t.Fatal(err)
	}
	reportID := auth.NewReportID()
	character := "us/hardcore/baelgrim"
	if _, err := rs.Create(ctx, reports.Report{
		ID: reportID, OwnerID: &owner.ID, Title: "Tuesday", Visibility: reports.Public,
		Status: reports.StatusProcessing, UploadID: &uploadID, LoggingCharacter: &character,
	}); err != nil {
		t.Fatal(err)
	}
	objects := newMemObjects()
	body := engine.FixtureLog()
	if packed {
		body = pack(t, body)
	}
	if err := objects.Put(ctx, key, body, store.PutOptions{}); err != nil {
		t.Fatal(err)
	}
	rank := &fakeRanker{}
	return Deps{
		Reports: rs, Objects: objects, Rank: rank,
		Log: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}, reportID, objects, rank
}

func pack(t *testing.T, b []byte) []byte {
	t.Helper()
	var buf bytes.Buffer
	w, err := zstd.NewWriter(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write(b); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func TestParseReportFillsInTheReportProgressively(t *testing.T) {
	for _, packed := range []bool{false, true} {
		d, reportID, objects, rank := fixtureUpload(t, packed)
		ctx := context.Background()
		if err := Report(ctx, d, reportID); err != nil {
			t.Fatalf("packed=%v: %v", packed, err)
		}
		rep, err := d.Reports.Get(ctx, reportID)
		if err != nil {
			t.Fatal(err)
		}
		if rep.Status != reports.StatusComplete || rep.EngineVersion != engine.Version {
			t.Fatalf("report = %+v", rep)
		}
		if !strings.Contains(string(rep.Health), `"lines"`) {
			t.Fatalf("health = %s", rep.Health)
		}
		fights, err := d.Reports.Fights(ctx, reportID)
		if err != nil {
			t.Fatal(err)
		}
		if len(fights) != 1 || !fights[0].Verified || fights[0].Name != "Warden Kelthas" {
			t.Fatalf("fights = %+v", fights)
		}
		keys := store.Keys{ReportID: reportID}
		for _, key := range []string{keys.Report(), keys.FightSummary(fights[0].Index), keys.FightEvents(fights[0].Index)} {
			if _, ok := objects.get(key); !ok {
				t.Errorf("object %s was not written", key)
			}
		}
		var report store.Report
		body, _ := objects.get(keys.Report())
		decodeJSON(t, body, &report)
		if len(report.Units) == 0 {
			t.Fatal("the final report.json should carry the unit list")
		}
		if len(rank.written) != 1 || rank.written[0].Ruleset != "hardcore" {
			t.Fatalf("ranked = %+v", rank.written)
		}
	}
}

func TestParseReportFailsLoudlyWhenTheUploadIsMissing(t *testing.T) {
	d, reportID, objects, _ := fixtureUpload(t, false)
	objects.getErr = errors.New("no such object")
	if err := Report(context.Background(), d, reportID); err == nil {
		t.Fatal("a missing upload must fail")
	}
	rep, err := d.Reports.Get(context.Background(), reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != reports.StatusFailed {
		t.Fatalf("status = %q, want failed", rep.Status)
	}
}

func TestParseReportNeedsAnUpload(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	if _, err := d.Reports.Pool.Exec(context.Background(),
		`update reports set upload_id = null where id = $1`, reportID); err != nil {
		t.Fatal(err)
	}
	if err := Report(context.Background(), d, reportID); err == nil {
		t.Fatal("a report with no upload cannot be parsed")
	}
	if err := Report(context.Background(), d, "nosuchreport"); err == nil {
		t.Fatal("an unknown report cannot be parsed")
	}
}

func TestMaybeDecompressPassesPlainTextThrough(t *testing.T) {
	r, err := maybeDecompress(io.NopCloser(strings.NewReader("plain text")))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	got, err := io.ReadAll(r)
	if err != nil || string(got) != "plain text" {
		t.Fatalf("read %q, %v", got, err)
	}
}

// decodeJSON is a small helper for the objects the job writes.
func decodeJSON(t *testing.T, b []byte, into any) {
	t.Helper()
	if err := json.Unmarshal(b, into); err != nil {
		t.Fatal(err)
	}
}

func TestACancelledContextStopsTheParse(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Report(ctx, d, reportID); err == nil {
		t.Fatal("a cancelled parse must report it")
	}
}

// failingRanker refuses everything, which is what a database outage
// looks like to the job.
type failingRanker struct{}

func (failingRanker) WriteFight(context.Context, reports.RankedFight) error {
	return errors.New("rankings are down")
}

func (failingRanker) RemoveReport(context.Context, string, string) error {
	return errors.New("rankings are down")
}

func TestAFailingRankerFailsTheParse(t *testing.T) {
	d, reportID, _, _ := fixtureUpload(t, false)
	d.Rank = failingRanker{}
	err := Report(context.Background(), d, reportID)
	if err == nil || !strings.Contains(err.Error(), "failed") {
		t.Fatalf("err = %v, want the report marked failed", err)
	}
	rep, err := d.Reports.Get(context.Background(), reportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != reports.StatusFailed {
		t.Fatalf("status = %q, want failed", rep.Status)
	}
}

// pacedReader hands the log out a few hundred bytes at a time and
// records how much of it the job has taken, so a test can tell whether
// a fight was published before the file had been read to the end.
type pacedReader struct {
	data []byte
	step int
	read int
}

func (p *pacedReader) Read(b []byte) (int, error) {
	if p.read >= len(p.data) {
		return 0, io.EOF
	}
	n := copy(b, p.data[p.read:min(p.read+p.step, len(p.data))])
	p.read += n
	return n, nil
}

// twoEncounters is the fixture log twice over, the second copy under a
// different encounter id, so a stream holds two fights that close at
// different points in the file.
func twoEncounters() []byte {
	log := engine.FixtureLog()
	return append(append([]byte{}, log...), bytes.ReplaceAll(log, []byte("9001"), []byte("9002"))...)
}

func TestStreamPublishesEachFightBeforeTheFileIsRead(t *testing.T) {
	src := &pacedReader{data: twoEncounters(), step: 512}
	s := session.New(engine.SessionOptions("rpt", engine.FixtureBase, true))
	var readWhenClosed []int
	if err := stream(context.Background(), s, src, func(session.Closed) error {
		readWhenClosed = append(readWhenClosed, src.read)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	if len(readWhenClosed) != 2 {
		t.Fatalf("closed %d fights, want the log's 2", len(readWhenClosed))
	}
	if readWhenClosed[0] >= len(src.data) {
		t.Fatalf("the first fight was published only after all %d bytes were read: "+
			"the parse is not progressive", len(src.data))
	}
}

// errReader is a body that dies on the first read, which is what a
// dropped connection to the bucket looks like to the job.
type errReader struct{ err error }

func (e errReader) Read([]byte) (int, error) { return 0, e.err }

func TestStreamReportsAReadError(t *testing.T) {
	s := session.New(engine.SessionOptions("rpt", engine.FixtureBase, true))
	err := stream(context.Background(), s, errReader{errors.New("connection reset")},
		func(session.Closed) error { return nil })
	if err == nil || !strings.Contains(err.Error(), "connection reset") {
		t.Fatalf("err = %v, want the read error", err)
	}
}

func TestStreamStopsOnACancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	s := session.New(engine.SessionOptions("rpt", engine.FixtureBase, true))
	err := stream(ctx, s, bytes.NewReader(engine.FixtureLog()), func(session.Closed) error { return nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
}
