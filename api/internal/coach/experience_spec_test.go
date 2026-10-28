package coach

import (
	"os"
	"strings"
	"testing"
)

// experienceSpecPath is the experience spec whose §8 copy table fixes the pack's words.
const (
	experienceSpecPath = "../../../design/specs/2026-10-28-log-coach.md"
	packCopyHeading    = "### 8.1 Pack copy"
)

// experienceRows is every row of the table under heading, by its first column's key: the
// backticked segments of the second column, whitespace collapsed (the specCell pattern of
// api/internal/discordbot/spectable_test.go).
func experienceRows(t *testing.T, heading string) map[string][]string {
	t.Helper()
	raw, err := os.ReadFile(experienceSpecPath)
	if err != nil {
		t.Fatal(err)
	}
	spec := string(raw)
	start := strings.Index(spec, heading)
	if start == -1 {
		t.Fatalf("experience spec has no heading %q", heading)
	}
	section := spec[start+len(heading):]
	if end := strings.Index(section, "\n#"); end != -1 {
		section = section[:end]
	}
	rows := map[string][]string{}
	for _, line := range strings.Split(section, "\n") {
		cells := strings.Split(line, "|")
		if !strings.HasPrefix(line, "| ") || len(cells) < 4 {
			continue
		}
		key := strings.TrimSpace(cells[1])
		if key == "Key" || strings.Trim(key, "-") == "" {
			continue
		}
		for _, m := range inlineSpan.FindAllStringSubmatch(cells[2], -1) {
			rows[key] = append(rows[key], collapse(m[1]))
		}
	}
	return rows
}

func TestExperienceSpecCopyIsPinnedWordForWord(t *testing.T) {
	rows := experienceRows(t, packCopyHeading)
	if len(rows) < 50 {
		t.Fatalf("read %d rows from §8.1; the table has more", len(rows))
	}
	for key, cell := range rows {
		got, ok := experienceCopy[key]
		switch {
		case !ok:
			t.Errorf("§8.1 key %s has no copy.go entry", key)
		case len(cell) == 0:
			t.Errorf("§8.1 key %s has no backticked words", key)
		case collapse(got) != cell[0]:
			t.Errorf("%s = %q, §8.1 says %q", key, got, cell[0])
		}
	}
	for key := range experienceCopy {
		if _, ok := rows[key]; !ok {
			t.Errorf("copy.go entry %s has no §8.1 row", key)
		}
	}
}

func TestExperienceRowsReadsOneCell(t *testing.T) {
	rows := experienceRows(t, packCopyHeading)
	if got := rows["DetailCoreMore"]; len(got) != 1 || got[0] != "Also lost: {more}." {
		t.Fatalf("DetailCoreMore = %q: the Where column's backticks are not the words", got)
	}
}

// Drafts holds only what §8.1 left: no draft repeats a pinned string.
func TestDraftsAreOnlyWhatTheTableLeft(t *testing.T) {
	pinned := map[string]string{}
	for key, s := range experienceCopy {
		pinned[s] = key
	}
	for key, d := range Drafts {
		if k, ok := pinned[d]; ok {
			t.Errorf("draft %s repeats §8.1's %s: read the constant", key, k)
		}
	}
}
