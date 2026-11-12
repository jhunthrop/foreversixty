package watch

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

// appendTo appends to a log and stamps its modification time, because
// the watcher picks the newest file and a test must control that.
func appendTo(t *testing.T, path, text string, mtime time.Time) {
	t.Helper()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString(text); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, mtime, mtime); err != nil {
		t.Fatal(err)
	}
}

func kinds(evs []Event) []string {
	out := make([]string, 0, len(evs))
	for _, e := range evs {
		out = append(out, e.Kind.String())
	}
	return out
}

func same(got []string, want ...string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func TestAppendedBytesArriveWithTheirOffset(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})

	if evs, err := w.Poll(t0); err != nil || len(evs) != 0 {
		t.Fatalf("empty directory = %v, %v", kinds(evs), err)
	}
	appendTo(t, log, "line one\n", t0)
	evs, err := w.Poll(t0.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[1].Offset != 0 || string(evs[1].Data) != "line one\n" {
		t.Fatalf("append = %+v", evs[1])
	}
	appendTo(t, log, "line two\n", t0.Add(2*time.Second))
	evs, err = w.Poll(t0.Add(3 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "append") || evs[0].Offset != 9 || string(evs[0].Data) != "line two\n" {
		t.Fatalf("second append = %v %+v", kinds(evs), evs)
	}
	if w.Offset() != 18 {
		t.Errorf("offset = %d, want 18", w.Offset())
	}
}

func TestAPreExistingLogIsNotReplayedFromItsStart(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "last night\n", t0.Add(-24*time.Hour))

	w := New(Options{Dir: dir})
	if evs, err := w.Poll(t0); err != nil || len(evs) != 0 {
		t.Fatalf("first poll = %v, %v", kinds(evs), err)
	}
	if w.Offset() != 11 {
		t.Fatalf("offset = %d, want the file's size 11", w.Offset())
	}
	appendTo(t, log, "tonight\n", t0.Add(time.Second))
	evs, err := w.Poll(t0.Add(2 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "start", "append") || string(evs[1].Data) != "tonight\n" {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
}

func TestRotationToATimestampedFileClosesTheReportAndOpensANewOne(t *testing.T) {
	dir := t.TempDir()
	live := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, live, "first night\n", t0)
	if evs, _ := w.Poll(t0.Add(time.Second)); !same(kinds(evs), "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	// The game rotates: the live file is renamed and a new one starts.
	rotated := filepath.Join(dir, "WoWCombatLog-120926_200000.txt")
	if err := os.Rename(live, rotated); err != nil {
		t.Fatal(err)
	}
	appendTo(t, live, "second night\n", t0.Add(time.Minute))
	evs, err := w.Poll(t0.Add(time.Minute + time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[0].Offset != 12 {
		t.Errorf("the completed report's final offset = %d, want 12", evs[0].Offset)
	}
	if evs[2].Offset != 0 || string(evs[2].Data) != "second night\n" {
		t.Errorf("new report append = %+v", evs[2])
	}
}

func TestATruncatedFileIsTreatedAsANewReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, log, "aaaaaaaaaa\n", t0)
	w.Poll(t0.Add(time.Second))

	if err := os.WriteFile(log, []byte("b\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(log, t0.Add(2*time.Second), t0.Add(2*time.Second)); err != nil {
		t.Fatal(err)
	}
	evs, err := w.Poll(t0.Add(3 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if string(evs[2].Data) != "b\n" {
		t.Errorf("append = %q", evs[2].Data)
	}
}

func TestFifteenQuietMinutesCompleteTheReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir})
	w.Poll(t0)
	appendTo(t, log, "pull\n", t0)
	w.Poll(t0.Add(time.Second))

	if evs, _ := w.Poll(t0.Add(14 * time.Minute)); len(evs) != 0 {
		t.Fatalf("completed after fourteen minutes: %v", kinds(evs))
	}
	evs, err := w.Poll(t0.Add(16 * time.Minute))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete") || evs[0].Offset != 5 {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
	if evs2, _ := w.Poll(t0.Add(17 * time.Minute)); len(evs2) != 0 {
		t.Fatalf("completed twice: %v", kinds(evs2))
	}
}

func TestAGapLongerThanThirtyMinutesStartsANewReport(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	// CompleteIdle is set past the gap so the only rule under test is
	// the thirty-minute one.
	w := New(Options{Dir: dir, CompleteIdle: 4 * time.Hour})
	w.Poll(t0)
	appendTo(t, log, "first raid\n", t0)
	w.Poll(t0.Add(time.Second))

	late := t0.Add(31 * time.Minute)
	appendTo(t, log, "second raid\n", late)
	evs, err := w.Poll(late.Add(time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "complete", "start", "append") {
		t.Fatalf("events = %v", kinds(evs))
	}
	if evs[1].Offset != 11 {
		t.Errorf("the new report starts at %d, want 11", evs[1].Offset)
	}
}

func TestResumeContinuesAReportAcrossARestart(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "before the restart\n", t0)

	w := New(Options{Dir: dir})
	if err := w.Resume(log, 19, t0); err != nil {
		t.Fatal(err)
	}
	appendTo(t, log, "after\n", t0.Add(time.Second))
	evs, err := w.Poll(t0.Add(2 * time.Second))
	if err != nil {
		t.Fatal(err)
	}
	if !same(kinds(evs), "append") || string(evs[0].Data) != "after\n" {
		t.Fatalf("events = %v %+v", kinds(evs), evs)
	}
}

func TestResumeRefusesAnOffsetPastTheEndOfTheFile(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "short\n", t0)
	w := New(Options{Dir: dir})
	if err := w.Resume(log, 9999, t0); err == nil {
		t.Fatal("Resume accepted an offset past the end of the file")
	}
	if err := w.Resume(filepath.Join(dir, "missing.txt"), 0, t0); err == nil {
		t.Fatal("Resume accepted a file that is not there")
	}
}

func TestALargeAppendIsDeliveredInChunks(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	w := New(Options{Dir: dir, ReadChunk: 4})
	w.Poll(t0)
	appendTo(t, log, "abcdefghij", t0)

	var got string
	for i := range 3 {
		evs, err := w.Poll(t0.Add(time.Duration(i+1) * time.Second))
		if err != nil {
			t.Fatal(err)
		}
		for _, e := range evs {
			if e.Kind == Append {
				got += string(e.Data)
			}
		}
	}
	if got != "abcdefghij" {
		t.Fatalf("reassembled %q", got)
	}
}

func TestReadRangeReturnsAFightsBytes(t *testing.T) {
	dir := t.TempDir()
	log := filepath.Join(dir, "WoWCombatLog.txt")
	appendTo(t, log, "0123456789", t0)
	got, err := ReadRange(log, 2, 6)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "2345" {
		t.Fatalf("ReadRange = %q", got)
	}
	if _, err := ReadRange(log, 6, 2); err == nil {
		t.Error("ReadRange accepted a backwards range")
	}
	if _, err := ReadRange(log, 8, 99); err == nil {
		t.Error("ReadRange accepted a range past the end of the file")
	}
}

const (
	header1 = "9/26 20:10:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,1,BUILD_VERSION,9.0.2,PROJECT_ID,1\n"
	header2 = "9/27 19:00:00.000  COMBAT_LOG_VERSION,16,ADVANCED_LOG_ENABLED,0,BUILD_VERSION,9.0.2,PROJECT_ID,1\n"
	filler  = "9/26 20:10:00.500  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n"
)

// A companion that starts over a log the game is already writing begins at the file's end,
// past the header; the header is still in the file, and the latest one before the start is
// the one that governs the lines that follow.
func TestLastHeaderBeforeFindsTheLatestHeaderBehindTheOffset(t *testing.T) {
	log := filepath.Join(t.TempDir(), "WoWCombatLog.txt")
	body := header1 + strings.Repeat(filler, 50) + header2 + strings.Repeat(filler, 20)
	appendTo(t, log, body, t0)

	line, ok, err := LastHeaderBefore(log, int64(len(body)))
	if err != nil || !ok {
		t.Fatalf("header = %q, %v, %v", line, ok, err)
	}
	if line != strings.TrimSuffix(header2, "\n") {
		t.Fatalf("header = %q, want the second header", line)
	}
	// An offset inside the second header belongs to the report itself, so the one before it is
	// the answer; and before the first header there is none.
	line, ok, err = LastHeaderBefore(log, int64(len(header1)+50*len(filler)+10))
	if err != nil || !ok || line != strings.TrimSuffix(header1, "\n") {
		t.Fatalf("mid-header offset = %q, %v, %v; want the first header", line, ok, err)
	}
	if _, ok, err := LastHeaderBefore(log, 5); err != nil || ok {
		t.Fatalf("an offset before any header found one: %v, %v", ok, err)
	}
}

// The scan steps back through the file in windows; a header cut by a window boundary is
// still read whole.
func TestLastHeaderBeforeReadsAHeaderSplitAcrossWindows(t *testing.T) {
	saved := headerScanWindow
	headerScanWindow = 40
	t.Cleanup(func() { headerScanWindow = saved })
	log := filepath.Join(t.TempDir(), "WoWCombatLog.txt")
	body := strings.Repeat(filler, 3) + header1 + strings.Repeat(filler, 7)
	appendTo(t, log, body, t0)
	line, ok, err := LastHeaderBefore(log, int64(len(body)))
	if err != nil || !ok || line != strings.TrimSuffix(header1, "\n") {
		t.Fatalf("header = %q, %v, %v", line, ok, err)
	}
}

func TestLastHeaderBeforeIsFalseForAFileWithoutOne(t *testing.T) {
	log := filepath.Join(t.TempDir(), "WoWCombatLog.txt")
	appendTo(t, log, strings.Repeat(filler, 9), t0)
	if line, ok, err := LastHeaderBefore(log, int64(9*len(filler))); err != nil || ok {
		t.Fatalf("header = %q, %v, %v", line, ok, err)
	}
}
