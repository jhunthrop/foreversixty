// api/internal/rating/presence/evidence_source_test.go
package presence

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Task L6.12: one rule, no exceptions. Every reader that answers anything a player or guild
// sees picks its reports through EvidenceReport (not private, not flagged). This is the grep
// for it, over every non-test Go file of the api module: a SQL condition on a report's
// visibility either sits on the same line as EvidenceReport (a public-only reader states its
// stricter rule beside it) or is on the list below, each with the reason it is not a reader
// of evidence.

// visibilityCondition is a condition on a report's visibility: `visibility <> 'private'`,
// `r.visibility = 'public'`, `visibility in (...)`, `visibility != Private`. A column named
// default_visibility is not a report's visibility (the word boundary leaves it out).
var visibilityCondition = regexp.MustCompile(`\bvisibility\s*(<>|!=|=|in\b)`)

// notEvidence is every visibility condition that is not a reader of evidence: the file
// (slash-separated, from the api module root) and a substring of the line.
var notEvidence = []struct{ file, line, why string }{
	{"internal/rating/presence/presence.go", `".visibility <> 'private' and "`,
		"EvidenceReport itself"},
	{"internal/reports/report.go", "func Ranked(",
		"ingest's choice whether a report is ranked and rated at all, not a reader"},
	{"internal/reports/store.go", "visibility in ($2, $3)",
		"the guild's report list: which reports a viewer may open (access, not evidence)"},
	{"internal/reports/store.go", "r.visibility = $1 and r.status = $2",
		"the recent public reports list: which reports anyone may open (access, not evidence)"},
	{"internal/reports/store.go", "visibility = coalesce($3, visibility)",
		"the report edit's update, not a reader"},
	{"internal/guilds/home.go", "r.owner_id = $3 or r.visibility = 'public'",
		"HomeReports: which reports the viewer may open, its own included (access, not evidence)"},
	{"internal/guilds/raids.go", "r.owner_id = $2 or r.visibility = 'public'",
		"the Raids tab's report rows: which reports the viewer may open (access, not evidence); " +
			"what each row says about a night (present, top_parse) reads EvidenceReport"},
}

func TestEveryReportVisibilityConditionIsTheEvidenceRule(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", "..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		t.Fatalf("%s is not the api module root: %v", root, err)
	}
	used := make([]bool, len(notEvidence))
	var scanned int
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if name := d.Name(); name == "node_modules" || name == "testdata" || strings.HasPrefix(name, ".") && path != root {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		scanned++
		sc := bufio.NewScanner(f)
		sc.Buffer(make([]byte, 0, 64*1024), 1<<20)
		for n := 1; sc.Scan(); n++ {
			line := sc.Text()
			if trimmed := strings.TrimSpace(line); strings.HasPrefix(trimmed, "//") {
				continue
			}
			if !visibilityCondition.MatchString(line) || strings.Contains(line, "EvidenceReport(") {
				continue
			}
			allowed := false
			for i, a := range notEvidence {
				if a.file == rel && strings.Contains(line, a.line) {
					used[i], allowed = true, true
				}
			}
			if !allowed {
				t.Errorf("%s:%d reads reports by visibility without presence.EvidenceReport beside it:\n\t%s",
					rel, n, strings.TrimSpace(line))
			}
		}
		return sc.Err()
	})
	if err != nil {
		t.Fatal(err)
	}
	if scanned < 100 {
		t.Fatalf("scanned %d Go files under %s, want the whole api module", scanned, root)
	}
	for i, a := range notEvidence {
		if !used[i] {
			t.Errorf("the allowlist entry %s %q matches nothing: remove it", a.file, a.line)
		}
	}
}
