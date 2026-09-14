// logs/engine/session/session_test.go
package session

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/layout"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

var base = time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC)

func fixture(t *testing.T) []byte {
	t.Helper()
	b, err := os.ReadFile("../event/testdata/v16.log")
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func opts() Options {
	o := Options{
		ReportID:   "report-1",
		Base:       base,
		KeepEvents: true,
		Units:      units.Options{ClassBySpec: units.RetailSpecClass},
		Fight:      fight.DefaultOptions(),
		Summary:    summary.DefaultOptions(),
	}
	o.Summary.SpecNames = units.RetailSpecName
	return o
}

// feed runs the whole input through a session in chunks of the given size
// and returns everything it produced.
func feed(t *testing.T, o Options, src []byte, size int) (Result, Health) {
	t.Helper()
	s := New(o)
	var all Result
	for i := 0; i < len(src); i += size {
		end := min(i+size, len(src))
		r, err := s.Feed(src[i:end], int64(i))
		if err != nil {
			t.Fatalf("Feed at %d: %v", i, err)
		}
		all.Events = append(all.Events, r.Events...)
		all.Closed = append(all.Closed, r.Closed...)
		all.Bytes += r.Bytes
	}
	r, err := s.Close()
	if err != nil {
		t.Fatalf("Close: %v", err)
	}
	all.Events = append(all.Events, r.Events...)
	all.Closed = append(all.Closed, r.Closed...)
	return all, s.Health()
}

// digest is the comparable shape of a run: what came out, not how.
type digest struct {
	Events []string
	Closed []Closed
}

func digestOf(r Result) digest {
	d := digest{Closed: r.Closed}
	for _, e := range r.Events {
		d.Events = append(d.Events, e.Name+"@"+e.Time.Format(time.RFC3339Nano))
	}
	return d
}

func jsonOf(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestStreamingEquivalenceAtEveryChunkSize(t *testing.T) {
	src := fixture(t)
	want, wantHealth := feed(t, opts(), src, len(src))
	wantJSON := jsonOf(t, digestOf(want))

	sizes := []int{1, 7, 64, 4096}
	rng := rand.New(rand.NewPCG(7, 11))
	for range 5 {
		sizes = append(sizes, 1+rng.IntN(2048))
	}
	for _, size := range sizes {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			got, gotHealth := feed(t, opts(), src, size)
			if jsonOf(t, digestOf(got)) != wantJSON {
				t.Fatalf("chunk size %d produced different output", size)
			}
			if jsonOf(t, gotHealth) != jsonOf(t, wantHealth) {
				t.Fatalf("chunk size %d produced different health", size)
			}
		})
	}
}

func TestDuplicateAndOverlappingChunksAreIgnored(t *testing.T) {
	src := fixture(t)
	want, _ := feed(t, opts(), src, len(src))

	s := New(opts())
	var got Result
	add := func(r Result) {
		got.Events = append(got.Events, r.Events...)
		got.Closed = append(got.Closed, r.Closed...)
	}

	half := len(src) / 2
	r, err := s.Feed(src[:half], 0)
	if err != nil {
		t.Fatal(err)
	}
	add(r)
	// The same range again, and a range that overlaps backwards.
	if r, err = s.Feed(src[:half], 0); err != nil {
		t.Fatal(err)
	}
	add(r)
	if r, err = s.Feed(src[half/2:], int64(half/2)); err != nil {
		t.Fatal(err)
	}
	add(r)
	if r, err = s.Close(); err != nil {
		t.Fatal(err)
	}
	add(r)

	if jsonOf(t, digestOf(got)) != jsonOf(t, digestOf(want)) {
		t.Fatal("re-sent bytes changed the output")
	}
}

func TestAGapIsRefused(t *testing.T) {
	s := New(opts())
	if _, err := s.Feed(fixture(t)[:100], 5000); err == nil {
		t.Fatal("a chunk past the end of the stream must be refused")
	}
}

func TestSnapshotDuringAFightThenTheClosedSummaryAgree(t *testing.T) {
	src := fixture(t)
	s := New(opts())
	if _, err := s.Feed(src, 0); err != nil {
		t.Fatal(err)
	}
	open, snap, ok := s.Snapshot()
	if !ok {
		t.Fatal("expected a fight in progress at the end of the fixture")
	}
	if !open.InProgress {
		t.Error("the open fight must be marked in progress")
	}
	if snap.EngineVersion != Version {
		t.Errorf("snapshot engine version = %q", snap.EngineVersion)
	}
	r, err := s.Close()
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Closed) == 0 {
		t.Fatal("Close must flush the open fight")
	}
	last := r.Closed[len(r.Closed)-1]
	if last.Summary.DurationMS < snap.DurationMS {
		t.Errorf("the closed summary is shorter than the snapshot: %d < %d",
			last.Summary.DurationMS, snap.DurationMS)
	}
}

func TestTheFixtureProducesTheEncounterWithItsMetrics(t *testing.T) {
	got, health := feed(t, opts(), fixture(t), 512)
	var enc *Closed
	for i := range got.Closed {
		if got.Closed[i].Fight.Kind == fight.Encounter {
			enc = &got.Closed[i]
		}
	}
	if enc == nil {
		t.Fatalf("no encounter among %d fights", len(got.Closed))
	}
	if enc.Fight.Name != "Warden Kelthas" || !enc.Fight.Kill {
		t.Errorf("encounter = %+v", enc.Fight)
	}
	if len(enc.Metrics) == 0 {
		t.Error("a boss kill must produce ranking metrics rows")
	}
	for _, m := range enc.Metrics {
		if m.ReportID != "report-1" || m.EngineVersion != Version || m.EncounterID != 9001 {
			t.Fatalf("metric row = %+v", m)
		}
	}
	if len(enc.Events) == 0 {
		t.Error("KeepEvents must retain the fight's events")
	}
	if health.ParseErrors != 0 || len(health.UnknownEvents) != 0 {
		t.Errorf("health = %+v", health)
	}
	if health.Layout != "retail-v16" || !health.LayoutVerified || !health.AdvancedLogging {
		t.Errorf("layout health = %+v", health)
	}
}

func TestSerialiseAndRestoreMidFight(t *testing.T) {
	src := fixture(t)
	want, wantHealth := feed(t, opts(), src, len(src))

	s := New(opts())
	var got Result
	add := func(r Result) {
		got.Events = append(got.Events, r.Events...)
		got.Closed = append(got.Closed, r.Closed...)
	}

	// Feed up to somewhere inside the encounter, then serialise.
	cut := len(src) - 200
	r, err := s.Feed(src[:cut], 0)
	if err != nil {
		t.Fatal(err)
	}
	add(r)
	blob, err := s.State()
	if err != nil {
		t.Fatal(err)
	}
	replay, err := ReplayOffset(blob)
	if err != nil {
		t.Fatal(err)
	}
	if replay > int64(cut) {
		t.Fatalf("replay offset %d is past the bytes already fed (%d)", replay, cut)
	}
	// Drop everything the restored session will replay.
	got.Events = keepBefore(got.Events, replay)

	revived, err := Restore(opts(), blob)
	if err != nil {
		t.Fatal(err)
	}
	if r, err = revived.Feed(src[replay:], replay); err != nil {
		t.Fatal(err)
	}
	add(r)
	if r, err = revived.Close(); err != nil {
		t.Fatal(err)
	}
	add(r)
	gotHealth := revived.Health()

	if jsonOf(t, digestOf(got)) != jsonOf(t, digestOf(want)) {
		t.Fatal("a serialise and restore mid-fight changed the output")
	}
	if jsonOf(t, gotHealth) != jsonOf(t, wantHealth) {
		t.Fatalf("a serialise and restore mid-fight changed health: got %+v, want %+v", gotHealth, wantHealth)
	}
}

func keepBefore(evs []event.Event, offset int64) []event.Event {
	out := evs[:0]
	for _, e := range evs {
		if e.Offset < offset {
			out = append(out, e)
		}
	}
	return out
}

func TestRestoreRefusesStateFromAnotherEngineVersion(t *testing.T) {
	s := New(opts())
	if _, err := s.Feed(fixture(t), 0); err != nil {
		t.Fatal(err)
	}
	blob, err := s.State()
	if err != nil {
		t.Fatal(err)
	}
	var raw map[string]any
	if err := json.Unmarshal(blob, &raw); err != nil {
		t.Fatal(err)
	}
	raw["version"] = "0.0.0-old"
	stale, err := json.Marshal(raw)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Restore(opts(), stale); err == nil {
		t.Fatal("state from another engine version must be refused")
	}
}

func TestAYearRolloverAndAClockJumpAreReported(t *testing.T) {
	log := "12/31 23:59:59.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n" +
		"1/1 00:00:01.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n" +
		"1/1 00:00:00.500  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"
	o := opts()
	o.Base = time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	res, health := feed(t, o, []byte(log), 64)
	if health.YearRollovers != 1 {
		t.Errorf("year rollovers = %d, want 1", health.YearRollovers)
	}
	if health.ClockJumps != 1 {
		t.Errorf("clock jumps = %d, want 1", health.ClockJumps)
	}
	if res.Events[1].Time.Year() != 2027 {
		t.Errorf("the line after new year is %s", res.Events[1].Time)
	}
}

func TestAMalformedLineIsOneParseErrorNotAFailedFile(t *testing.T) {
	log := "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n" +
		"9/26 20:10:01.000  SPELL_DAMAGE,Player-1-A,\"A\",0x512,0x0\n" +
		"9/26 20:10:02.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"
	res, health := feed(t, opts(), []byte(log), 16)
	if health.ParseErrors != 1 {
		t.Fatalf("parse errors = %d, want 1", health.ParseErrors)
	}
	if len(res.Events) != 3 {
		t.Fatalf("events = %d, want 3: the bad line costs one event, not the file", len(res.Events))
	}
	if res.Events[1].Kind != event.ParseError || res.Events[1].Raw == "" {
		t.Errorf("bad line = %+v", res.Events[1])
	}
}

func TestAnUnknownEventIsCountedAndKeptRaw(t *testing.T) {
	log := "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n" +
		"9/26 20:10:01.000  SPELL_EMPOWER_END,Player-1-A,\"A\",0x512,0x0,Player-1-A,\"A\",0x512,0x0,1,\"X\",0x1,3\n"
	res, health := feed(t, opts(), []byte(log), 4096)
	if health.UnknownEvents["SPELL_EMPOWER_END"] != 1 {
		t.Fatalf("unknown events = %v", health.UnknownEvents)
	}
	if res.Events[1].Raw == "" {
		t.Error("an unknown event must keep its raw text")
	}
}

func TestALogWithNoHeaderInfersALayout(t *testing.T) {
	log := "9/26 20:10:01.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n" +
		"9/26 20:10:02.000  SPELL_AURA_APPLIED,Player-1-A,\"A\",0x512,0x0,Player-1-A,\"A\",0x512,0x0,17,\"Shield\",0x2,BUFF\n"
	o := opts()
	o.Infer = true
	_, health := feed(t, o, []byte(log), 4096)
	if !health.MissingHeader || !health.LayoutInferred {
		t.Fatalf("health = %+v", health)
	}
	if health.LayoutVerified {
		t.Error("an inferred layout is never verified")
	}
}

func TestAMidFileHeaderIsAHardBoundary(t *testing.T) {
	header := "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n"
	log := header +
		"9/26 20:10:01.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n" +
		"9/26 20:20:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n" +
		"9/26 20:20:01.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"
	_, health := feed(t, opts(), []byte(log), 32)
	if health.HeaderRestarts != 1 {
		t.Fatalf("header restarts = %d, want 1", health.HeaderRestarts)
	}
}

func TestAForcedLayoutSkipsHeaderDetection(t *testing.T) {
	o := opts()
	o.Layout = layout.RetailV16()
	_, health := feed(t, o, fixture(t), 1024)
	if health.Layout != "retail-v16" {
		t.Fatalf("layout = %q", health.Layout)
	}
}

// stampAt returns a yearless "M/D HH:MM:SS.mmm" timestamp n seconds after
// 9/26 20:10:00.000, for synthetic logs that need many distinct timestamps
// without risking an hour or day rollover.
func stampAt(n int) string {
	total := 10*60 + n // seconds since 20:00:00
	hh := 20 + total/3600
	mm := (total % 3600) / 60
	ss := total % 60
	return fmt.Sprintf("9/26 %02d:%02d:%02d.000", hh, mm, ss)
}

// swingDamageLog builds n identically shaped SWING_DAMAGE lines, one second
// apart, a hostile creature swinging on a friendly player. That is real
// combat: it opens and sustains a trash fight, and its field count is
// exactly what a real retail-v16 log carries, so a layout inferred from a
// sample of these lines decodes every later one the same way.
func swingDamageLog(n int) []byte {
	const line = `SWING_DAMAGE,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,41320,44000,612,0,1955,0,1,0,1000,0,-1487.02,6409.71,1675,1.2044,45,812,1290,-1,1,0,0,0,nil,nil,nil`
	var b strings.Builder
	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, "%s  %s\n", stampAt(i), line)
	}
	return []byte(b.String())
}

func TestSerialiseAndRestoreAfterAnInferredLayout(t *testing.T) {
	src := swingDamageLog(inferWindow + 500)
	o := opts()
	o.Infer = true

	want, _ := feed(t, o, src, len(src))

	s := New(o)
	var got Result
	add := func(r Result) {
		got.Events = append(got.Events, r.Events...)
		got.Closed = append(got.Closed, r.Closed...)
	}

	// Feed far enough to cross the inference window and open a fight, then
	// serialise while still short of the end of the file.
	cut := len(src) - 200
	r, err := s.Feed(src[:cut], 0)
	if err != nil {
		t.Fatal(err)
	}
	add(r)
	if health := s.Health(); health.Layout != "inferred" || !health.LayoutInferred {
		t.Fatalf("health before serialising = %+v, want an inferred layout", health)
	}
	if _, _, ok := s.Snapshot(); !ok {
		t.Fatal("expected a fight in progress before serialising")
	}

	blob, err := s.State()
	if err != nil {
		t.Fatal(err)
	}
	replay, err := ReplayOffset(blob)
	if err != nil {
		t.Fatal(err)
	}
	if replay > int64(cut) {
		t.Fatalf("replay offset %d is past the bytes already fed (%d)", replay, cut)
	}
	got.Events = keepBefore(got.Events, replay)

	revived, err := Restore(o, blob)
	if err != nil {
		t.Fatal(err)
	}
	if health := revived.Health(); health.Layout != "inferred" || !health.LayoutInferred {
		t.Fatalf("restored health = %+v, want an inferred layout", health)
	}
	if r, err = revived.Feed(src[replay:], replay); err != nil {
		t.Fatal(err)
	}
	add(r)
	if r, err = revived.Close(); err != nil {
		t.Fatal(err)
	}
	add(r)

	if jsonOf(t, digestOf(got)) != jsonOf(t, digestOf(want)) {
		t.Fatal("a serialise and restore after an inferred layout changed the output")
	}
}

func TestInferenceWindowIsChunkSizeInvariant(t *testing.T) {
	var b strings.Builder
	for i := 0; i < inferWindow; i++ {
		fmt.Fprintf(&b, "%s  TEST_INFER_SPECIAL,111,222\n", stampAt(i))
	}
	// A wider shape for the same event name, seen only past the window: a
	// sample that ran past the window by chunking accident would infer a
	// wider Special and decode these differently than a sample cut off
	// exactly at the window.
	for i := inferWindow; i < inferWindow+500; i++ {
		fmt.Fprintf(&b, "%s  TEST_INFER_SPECIAL,111,222,333\n", stampAt(i))
	}
	log := []byte(b.String())

	o := opts()
	o.Infer = true

	small, smallHealth := feed(t, o, log, 1)
	whole, wholeHealth := feed(t, o, log, len(log))

	if jsonOf(t, small.Events) != jsonOf(t, whole.Events) {
		t.Fatal("the inference window's sample depended on how the input was chunked")
	}
	if jsonOf(t, smallHealth) != jsonOf(t, wholeHealth) {
		t.Fatalf("health differs by chunk size: chunk-1=%+v whole=%+v", smallHealth, wholeHealth)
	}
}

// TestSerialiseAndRestoreAcrossAYearRollover proves State/Restore reseed the
// decoder's clock with the value as of the open fight's start, not as of
// whenever State() happened to be called: a fight that opens before a year
// rollover and is still open afterwards must replay its pre-rollover lines
// with the pre-rollover year, not the year the clock had advanced to by the
// time it was serialised.
func TestSerialiseAndRestoreAcrossAYearRollover(t *testing.T) {
	const swing = `SWING_DAMAGE,Creature-0-2085-2284-7855-169753-0000AA0001,"Hollow Sentinel",0xa48,0x0,Player-4184-000000A1,"Baelgrim-Nightslayer",0x511,0x0,Creature-0-2085-2284-7855-169753-0000AA0001,0000000000000000,41320,44000,612,0,1955,0,1,0,1000,0,-1487.02,6409.71,1675,1.2044,45,812,1290,-1,1,0,0,0,nil,nil,nil`
	stamps := []string{
		"12/31 23:59:58.000",
		"12/31 23:59:59.000",
		"1/1 00:00:00.000",
		"1/1 00:00:01.000",
		"1/1 00:00:02.000",
	}
	var b strings.Builder
	for _, stamp := range stamps {
		fmt.Fprintf(&b, "%s  %s\n", stamp, swing)
	}
	src := []byte(b.String())

	o := opts()
	o.Layout = layout.RetailV16()
	o.Base = time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)

	want, _ := feed(t, o, src, len(src))

	// Feed everything to one session without closing, so the trash fight
	// opened by the first line is still open, straddling the rollover,
	// when State() is called.
	s := New(o)
	if _, err := s.Feed(src, 0); err != nil {
		t.Fatal(err)
	}
	if _, _, ok := s.Snapshot(); !ok {
		t.Fatal("expected the trash fight to still be open across the rollover")
	}

	blob, err := s.State()
	if err != nil {
		t.Fatal(err)
	}
	replay, err := ReplayOffset(blob)
	if err != nil {
		t.Fatal(err)
	}
	if replay != 0 {
		t.Fatalf("replay offset = %d, want 0: the fight opened on the first line", replay)
	}

	revived, err := Restore(o, blob)
	if err != nil {
		t.Fatal(err)
	}
	r, err := revived.Feed(src[replay:], replay)
	if err != nil {
		t.Fatal(err)
	}
	closeR, err := revived.Close()
	if err != nil {
		t.Fatal(err)
	}
	got := Result{
		Events: append(append([]event.Event{}, r.Events...), closeR.Events...),
		Closed: append(append([]Closed{}, r.Closed...), closeR.Closed...),
	}

	if jsonOf(t, digestOf(got)) != jsonOf(t, digestOf(want)) {
		t.Fatal("a serialise and restore across a year rollover produced the wrong timestamps")
	}
}
