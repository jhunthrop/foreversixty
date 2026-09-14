// logs/engine/store/store_test.go
package store

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// recorder is the in-memory Putter the tests assert against, standing in
// for the R2 client a later phase will write.
type recorder struct {
	keys []string
	body map[string][]byte
	opts map[string]PutOptions
}

func newRecorder() *recorder {
	return &recorder{body: map[string][]byte{}, opts: map[string]PutOptions{}}
}

func (r *recorder) Put(_ context.Context, key string, body []byte, o PutOptions) error {
	r.keys = append(r.keys, key)
	r.body[key] = append([]byte(nil), body...)
	r.opts[key] = o
	return nil
}

func TestKeysMatchTheStorageLayout(t *testing.T) {
	k := Keys{ReportID: "abc123"}
	for got, want := range map[string]string{
		k.Report():        "reports/abc123/report.json",
		k.FightSummary(3): "reports/abc123/fights/3/summary.json",
		k.FightEvents(3):  "reports/abc123/fights/3/events.parquet",
		k.FightLive(3):    "reports/abc123/fights/3/live.json",
		k.Raw(1048576):    "reports/abc123/raw/1048576.zst",
	} {
		if got != want {
			t.Errorf("key = %q, want %q", got, want)
		}
	}
}

func TestPublisherWritesEveryFileWithTheRightCaching(t *testing.T) {
	rec := newRecorder()
	p := Publisher{Keys: Keys{ReportID: "abc123"}, Put: rec}
	ctx := t.Context()

	f := fight.Fight{Index: 1, Kind: fight.Encounter, Name: "Warden Kelthas",
		EncounterID: 9001, Kill: true,
		Start: time.Unix(1000, 0).UTC(), End: time.Unix(1040, 0).UTC(),
		Players: []string{"Player-4184-000000A1"}}
	rep := Report{ReportID: "abc123", EngineVersion: session.Version,
		Health: session.Health{Layout: "retail-v16"},
		Fights: []FightEntry{EntryOf(f)}}
	if err := p.WriteReport(ctx, rep); err != nil {
		t.Fatal(err)
	}
	sum := summary.Summary{EngineVersion: session.Version, FightIndex: 1, DurationMS: 40000}
	evs := []event.Event{{Time: time.Unix(1001, 0).UTC(), Line: 1, Name: "SPELL_DAMAGE", Kind: event.Damage}}
	if err := p.WriteFight(ctx, 1, sum, evs); err != nil {
		t.Fatal(err)
	}
	if err := p.WriteLive(ctx, 2, sum); err != nil {
		t.Fatal(err)
	}
	if err := p.WriteRaw(ctx, 4096, []byte("9/26 20:10:00.000  ZONE_CHANGE,1,\"Z\",0\n")); err != nil {
		t.Fatal(err)
	}

	want := map[string]PutOptions{
		"reports/abc123/report.json": {ContentType: "application/json", CacheControl: cacheMutable},
		"reports/abc123/fights/1/summary.json": {
			ContentType: "application/json", CacheControl: cacheImmutable},
		"reports/abc123/fights/1/events.parquet": {
			ContentType: "application/vnd.apache.parquet", CacheControl: cacheImmutable},
		"reports/abc123/fights/2/live.json": {
			ContentType: "application/json", CacheControl: cacheMutable},
		"reports/abc123/raw/4096.zst": {
			ContentType: "application/zstd", CacheControl: cachePrivate, ContentEncoding: "zstd"},
	}
	if len(rec.keys) != len(want) {
		t.Fatalf("wrote %v, want %d objects", rec.keys, len(want))
	}
	for key, o := range want {
		got, ok := rec.opts[key]
		if !ok {
			t.Errorf("%s was not written", key)
			continue
		}
		if got != o {
			t.Errorf("%s options = %+v, want %+v", key, got, o)
		}
	}

	var back Report
	if err := json.Unmarshal(rec.body["reports/abc123/report.json"], &back); err != nil {
		t.Fatal(err)
	}
	if len(back.Fights) != 1 || back.Fights[0].Name != "Warden Kelthas" || back.Fights[0].DurationMS != 40000 {
		t.Errorf("report = %+v", back)
	}

	readBack, err := parquet.Unmarshal(rec.body["reports/abc123/fights/1/events.parquet"])
	if err != nil {
		t.Fatal(err)
	}
	if len(readBack) != 1 || readBack[0].Name != "SPELL_DAMAGE" {
		t.Errorf("events = %+v", readBack)
	}
}

func TestWriteReportSortsWithoutMutatingTheCallersSlice(t *testing.T) {
	rec := newRecorder()
	p := Publisher{Keys: Keys{ReportID: "abc123"}, Put: rec}

	original := []FightEntry{
		{Index: 3, Name: "Third"},
		{Index: 1, Name: "First"},
		{Index: 2, Name: "Second"},
	}
	fights := append([]FightEntry(nil), original...)
	rep := Report{ReportID: "abc123", Fights: fights}
	if err := p.WriteReport(t.Context(), rep); err != nil {
		t.Fatal(err)
	}

	// Finding 1's regression test: the caller's own slice must be left
	// exactly as it was, in the order the caller built it.
	if !reflect.DeepEqual(fights, original) {
		t.Fatalf("caller's slice was mutated: got %+v, want %+v", fights, original)
	}

	// The determinism requirement: published output is sorted by index
	// regardless of the order the caller supplied.
	var back Report
	if err := json.Unmarshal(rec.body["reports/abc123/report.json"], &back); err != nil {
		t.Fatal(err)
	}
	wantOrder := []int{1, 2, 3}
	if len(back.Fights) != len(wantOrder) {
		t.Fatalf("fights = %+v, want %d entries", back.Fights, len(wantOrder))
	}
	for i, idx := range wantOrder {
		if back.Fights[i].Index != idx {
			t.Errorf("fights[%d].Index = %d, want %d", i, back.Fights[i].Index, idx)
		}
	}
}

func TestRawChunksRoundTripThroughZstd(t *testing.T) {
	rec := newRecorder()
	p := Publisher{Keys: Keys{ReportID: "abc123"}, Put: rec}
	original := []byte("9/26 20:10:00.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n")
	if err := p.WriteRaw(t.Context(), 0, original); err != nil {
		t.Fatal(err)
	}
	got, err := Decompress(rec.body["reports/abc123/raw/0.zst"])
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(original) {
		t.Fatalf("round trip = %q", got)
	}
}

func TestTheLocalDirectoryStoreWritesRealFiles(t *testing.T) {
	root := t.TempDir()
	p := Publisher{Keys: Keys{ReportID: "abc123"}, Put: NewDir(root)}
	if err := p.WriteReport(t.Context(), Report{ReportID: "abc123"}); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, "reports", "abc123", "report.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var back Report
	if err := json.Unmarshal(b, &back); err != nil {
		t.Fatal(err)
	}
	if back.ReportID != "abc123" {
		t.Fatalf("report = %+v", back)
	}
}

func TestACancelledContextStopsTheLocalStore(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	if err := NewDir(t.TempDir()).Put(ctx, "reports/x/report.json", []byte("{}"), PutOptions{}); err == nil {
		t.Fatal("a cancelled context must stop the write")
	}
}

func TestTheLocalStoreReportsAWriteFailure(t *testing.T) {
	// A file where a directory needs to be makes MkdirAll fail.
	root := t.TempDir()
	blocker := filepath.Join(root, "reports")
	if err := os.WriteFile(blocker, []byte("not a directory"), 0o644); err != nil {
		t.Fatal(err)
	}
	err := NewDir(root).Put(t.Context(), "reports/abc123/report.json", []byte("{}"), PutOptions{})
	if err == nil {
		t.Fatal("want an error when the path cannot be created")
	}
}

func TestDirPutRefusesAKeyThatEscapesRoot(t *testing.T) {
	root := t.TempDir()
	escaped := filepath.Join(filepath.Dir(root), "escape.json")
	if err := NewDir(root).Put(t.Context(), "../escape.json", []byte("{}"), PutOptions{}); err == nil {
		t.Fatal("want an error for a key that escapes the root")
	}
	if _, err := os.Stat(escaped); !os.IsNotExist(err) {
		t.Fatalf("a key that escapes the root wrote to %s", escaped)
	}
}

func TestDirPutAcceptsEveryKeyTheStorageLayoutProduces(t *testing.T) {
	root := t.TempDir()
	d := NewDir(root)
	k := Keys{ReportID: "abc123"}
	for _, key := range []string{
		k.Report(), k.FightSummary(1), k.FightEvents(1), k.FightLive(1), k.Raw(0),
	} {
		if err := d.Put(t.Context(), key, []byte("x"), PutOptions{}); err != nil {
			t.Errorf("Put(%q) = %v, want nil", key, err)
		}
	}
}

func TestDecompressRejectsGarbage(t *testing.T) {
	if _, err := Decompress([]byte("not zstd")); err == nil {
		t.Fatal("want an error for a non-zstd payload")
	}
}
