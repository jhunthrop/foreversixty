package state

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var t0 = time.Date(2026, 12, 9, 20, 0, 0, 0, time.UTC)

func TestAReportStateRoundTrips(t *testing.T) {
	dir := t.TempDir()
	want := Report{
		Key: "20261209-200000-deadbeef", ReportID: "rpt000000001",
		LogPath: "/w/Logs/WoWCombatLog.txt", StartOffset: 10, Offset: 4096, RawSent: 4096,
		Visibility: "public", StartedAt: t0, LastAppend: t0.Add(time.Minute),
		Fights: 3, Session: []byte{1, 2, 3}, EngineVersion: "0.1.0",
	}
	if err := Save(dir, want); err != nil {
		t.Fatal(err)
	}
	got, err := Load(dir, want.Key)
	if err != nil {
		t.Fatal(err)
	}
	if got.ReportID != want.ReportID || got.Offset != want.Offset || got.RawSent != want.RawSent ||
		got.Fights != want.Fights || string(got.Session) != string(want.Session) ||
		!got.LastAppend.Equal(want.LastAppend) {
		t.Fatalf("Load = %+v", got)
	}
}

func TestListIsSortedAndSkipsRubbish(t *testing.T) {
	dir := t.TempDir()
	for _, k := range []string{"20261209-210000-bb", "20261209-200000-aa"} {
		if err := Save(dir, Report{Key: k}); err != nil {
			t.Fatal(err)
		}
	}
	got, err := List(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Key != "20261209-200000-aa" {
		t.Fatalf("List = %+v", got)
	}
}

func TestSaveRefusesAnEmptyKeyAndRemoveIsIdempotent(t *testing.T) {
	dir := t.TempDir()
	if err := Save(dir, Report{}); err == nil {
		t.Fatal("Save accepted an empty key")
	}
	if err := Remove(dir, "never-existed"); err != nil {
		t.Fatalf("Remove of a missing report = %v", err)
	}
}

func TestNewKeyIsSortableAndUnique(t *testing.T) {
	a, b := NewKey(t0), NewKey(t0)
	if a == b {
		t.Fatal("two keys collided")
	}
	if len(a) != len("20261209-200000-deadbeef") {
		t.Fatalf("key = %q", a)
	}
	if NewKey(t0) >= NewKey(t0.Add(time.Hour)) {
		t.Fatal("keys do not sort by time")
	}
}

func TestLoadingAMissingOrBrokenReportIsAnError(t *testing.T) {
	dir := t.TempDir()
	if _, err := Load(dir, "never-existed"); err == nil {
		t.Fatal("Load accepted a missing report")
	}
	if err := os.WriteFile(filepath.Join(dir, "broken.json"), []byte("{not json"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(dir, "broken"); err == nil {
		t.Fatal("Load accepted a broken report")
	}
	// List skips it rather than failing: one bad file must not stop
	// tonight's raid from being logged.
	got, err := List(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("List = %+v", got)
	}
}

func TestListingADirectoryThatIsNotThereIsEmpty(t *testing.T) {
	got, err := List(filepath.Join(t.TempDir(), "state"))
	if err != nil || got != nil {
		t.Fatalf("List = %+v, %v", got, err)
	}
}

func TestRemoveDeletesOneReport(t *testing.T) {
	dir := t.TempDir()
	if err := Save(dir, Report{Key: "k"}); err != nil {
		t.Fatal(err)
	}
	if err := Remove(dir, "k"); err != nil {
		t.Fatal(err)
	}
	if got, _ := List(dir); len(got) != 0 {
		t.Fatalf("List = %+v", got)
	}
}
