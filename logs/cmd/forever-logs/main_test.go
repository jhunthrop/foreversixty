// logs/cmd/forever-logs/main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fixturePath(t *testing.T) string {
	t.Helper()
	p, err := filepath.Abs(filepath.Join("..", "..", "engine", "event", "testdata", "v16.log"))
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func exec(t *testing.T, args ...string) (string, string, error) {
	t.Helper()
	var out, errOut bytes.Buffer
	err := run(args, &out, &errOut)
	return out.String(), errOut.String(), err
}

func TestNoArgumentsIsAUsageError(t *testing.T) {
	if _, _, err := exec(t); err == nil {
		t.Fatal("want a usage error")
	}
	if _, _, err := exec(t, "bogus"); err == nil {
		t.Fatal("an unknown subcommand must be a usage error")
	}
}

func TestFightsListsTheFixtureFights(t *testing.T) {
	out, _, err := exec(t, "fights", fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "Warden Kelthas") {
		t.Fatalf("output:\n%s", out)
	}
	if !strings.Contains(out, "encounter") {
		t.Errorf("the encounter must be labelled:\n%s", out)
	}
}

func TestFightsJSONIsOneObjectPerLine(t *testing.T) {
	out, _, err := exec(t, "fights", "-json", fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) == 0 {
		t.Fatal("no fights")
	}
	found := false
	for _, line := range lines {
		var entry struct {
			Index int    `json:"index"`
			Name  string `json:"name"`
			Kind  string `json:"kind"`
		}
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			t.Fatalf("line %q: %v", line, err)
		}
		if entry.Name == "Warden Kelthas" && entry.Kind == "encounter" {
			found = true
		}
	}
	if !found {
		t.Fatalf("no encounter in:\n%s", out)
	}
}

func TestParseWritesTheStorageLayout(t *testing.T) {
	dir := t.TempDir()
	out, _, err := exec(t, "parse", "-out", dir, "-report", "abc123", fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "layout=retail-v16") {
		t.Errorf("output:\n%s", out)
	}
	for _, rel := range []string{
		"reports/abc123/report.json",
		"reports/abc123/fights/1/summary.json",
		"reports/abc123/fights/1/events.parquet",
	} {
		if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(rel))); err != nil {
			t.Errorf("%s was not written: %v", rel, err)
		}
	}
	b, err := os.ReadFile(filepath.Join(dir, "reports", "abc123", "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var report struct {
		ReportID string `json:"report_id"`
		Health   struct {
			Layout      string `json:"layout"`
			ParseErrors int    `json:"parse_errors"`
		} `json:"health"`
		Fights []struct {
			Name string `json:"name"`
		} `json:"fights"`
		Units []struct {
			GUID string `json:"guid"`
		} `json:"units"`
	}
	if err := json.Unmarshal(b, &report); err != nil {
		t.Fatal(err)
	}
	if report.ReportID != "abc123" || report.Health.Layout != "retail-v16" || report.Health.ParseErrors != 0 {
		t.Errorf("report = %+v", report)
	}
	if len(report.Fights) == 0 || len(report.Units) == 0 {
		t.Errorf("report has %d fights and %d units", len(report.Fights), len(report.Units))
	}
}

func TestTailOnceReadsToTheEnd(t *testing.T) {
	out, _, err := exec(t, "tail", "-once", fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "closed") {
		t.Fatalf("tail printed no closed fight:\n%s", out)
	}
}

func TestConformanceReportsLayoutsAndUnknownEvents(t *testing.T) {
	dir := t.TempDir()
	src, err := os.ReadFile(fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "good.txt"), src, 0o644); err != nil {
		t.Fatal(err)
	}
	odd := string(src) +
		"9/26 20:13:00.000  SPELL_EMPOWER_END,Player-4184-000000A1,\"Baelgrim-Nightslayer\",0x511,0x0," +
		"Player-4184-000000A1,\"Baelgrim-Nightslayer\",0x511,0x0,1,\"X\",0x1,3\n"
	if err := os.WriteFile(filepath.Join(dir, "odd.log"), []byte(odd), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "ignored.json"), []byte("{}"), 0o644); err != nil {
		t.Fatal(err)
	}

	out, _, err := exec(t, "conformance", dir)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "layout=retail-v16") {
		t.Errorf("output:\n%s", out)
	}
	if !strings.Contains(out, "SPELL_EMPOWER_END") {
		t.Errorf("the unknown event must be reported:\n%s", out)
	}
	if strings.Contains(out, "ignored.json") {
		t.Errorf("only .txt and .log files are scanned:\n%s", out)
	}
	if !strings.Contains(out, "2 files") {
		t.Errorf("the totals line is missing:\n%s", out)
	}
}

func TestConformanceJSON(t *testing.T) {
	dir := t.TempDir()
	src, err := os.ReadFile(fixturePath(t))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "good.txt"), src, 0o644); err != nil {
		t.Fatal(err)
	}
	out, _, err := exec(t, "conformance", "-json", dir)
	if err != nil {
		t.Fatal(err)
	}
	var row ConformanceRow
	if err := json.Unmarshal([]byte(strings.TrimSpace(out)), &row); err != nil {
		t.Fatalf("output %q: %v", out, err)
	}
	if row.Layout != "retail-v16" || !row.Verified || row.ParseErrors != 0 || row.Fights == 0 {
		t.Fatalf("row = %+v", row)
	}
}

func TestConformanceOnAnEmptyDirectorySaysSo(t *testing.T) {
	_, errOut, err := exec(t, "conformance", t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(errOut, "no .txt or .log files") {
		t.Fatalf("stderr = %q", errOut)
	}
}

func TestAnUnknownLayoutNameIsRejected(t *testing.T) {
	if _, _, err := exec(t, "fights", "-layout", "nope", fixturePath(t)); err == nil {
		t.Fatal("an unknown layout name must be rejected")
	}
}

func TestAMissingFileIsAnError(t *testing.T) {
	if _, _, err := exec(t, "fights", filepath.Join(t.TempDir(), "absent.txt")); err == nil {
		t.Fatal("want an error for a missing file")
	}
}

// TestFightsOnAV22ExcerptReportsPlayers runs the committed v22 excerpt
// through the fights command the way a user would, so the acceptance
// numbers in the plan are checked by something that runs in CI rather than
// only by the corpus sweep, which needs logs that are not in the repo.
//
// The plain-text fights report carries no layout field — only `parse` and
// `conformance` print one — so the "not inferred" half of the acceptance
// is covered by TestConformanceOnTheV22ExcerptIsClean below; this test
// covers the half fights can actually show: the excerpt's player count.
func TestFightsOnAV22ExcerptReportsPlayers(t *testing.T) {
	var out strings.Builder
	if err := run([]string{"fights", filepath.Join("..", "..", "engine", "event", "testdata", "v22.log")}, &out, io.Discard); err != nil {
		t.Fatal(err)
	}
	got := out.String()
	if !strings.Contains(got, "players=6") {
		t.Errorf("the report does not show the excerpt's 6 players:\n%s", got)
	}
}

// TestConformanceOnTheV22ExcerptIsClean is the in-repo half of the
// acceptance: the committed excerpts must show the verified row, no parse
// errors and no unknown events.
func TestConformanceOnTheV22ExcerptIsClean(t *testing.T) {
	var out strings.Builder
	dir := filepath.Join("..", "..", "engine", "event", "testdata")
	if err := run([]string{"conformance", "-json", dir}, &out, io.Discard); err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		var row ConformanceRow
		if err := json.Unmarshal([]byte(line), &row); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(row.Path, "v22") {
			continue // the v16 fixture is covered by its own tests
		}
		if row.Layout != "retail-v22" || !row.Verified || row.Inferred {
			t.Errorf("%s: layout=%s verified=%v inferred=%v", row.Path, row.Layout, row.Verified, row.Inferred)
		}
		if row.ParseErrors != 0 {
			t.Errorf("%s: %d parse errors", row.Path, row.ParseErrors)
		}
		if len(row.UnknownEvents) != 0 {
			t.Errorf("%s: unknown events %v", row.Path, row.UnknownEvents)
		}
	}
}
