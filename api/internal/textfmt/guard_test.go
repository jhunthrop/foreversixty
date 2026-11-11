package textfmt

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// apiRoot is the api module's root, from this package.
const apiRoot = "../.."

// guardRules are the copies this package replaced (component audit 2026-11-10, cluster 21). A
// reader-facing date or clock layout is written here only, and an embed colour is packed by
// brand.RGB only.
var guardRules = []struct {
	name    string
	pattern *regexp.Regexp
	home    string // the one package directory allowed to hold it
}{
	{
		name: "a reader-facing date or clock layout",
		// Mon 2 Jan, Monday 2 January, 2 Jan 2006, 2 January, 3:04 pm, 15:04 MST; not the ISO
		// stamps (2006-01-02 15:04) logs and job keys use.
		pattern: regexp.MustCompile(`"[^"\n]*(Mon(day)? 2 Jan|\b2 Jan(uary)?\b|\b3:04\b|[^-\d]15:04 MST)[^"\n]*"`),
		home:    "internal/textfmt",
	},
	{
		name:    "an RGB colour packed by hand",
		pattern: regexp.MustCompile(`\.R\)\s*<<\s*16`),
		home:    "internal/brand",
	},
}

// TestNoFormattingCopiesOutsideTheirHome fails when a copy reappears outside its home.
func TestNoFormattingCopiesOutsideTheirHome(t *testing.T) {
	err := filepath.WalkDir(apiRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == "testdata" || d.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(apiRoot, path)
		for _, hit := range guardHits(filepath.ToSlash(rel), string(src)) {
			t.Errorf("%s: %s; use textfmt or brand.RGB", rel, hit)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

// guardHits are the rules a file at rel breaks.
func guardHits(rel, src string) []string {
	var out []string
	for _, r := range guardRules {
		if strings.HasPrefix(rel, r.home+"/") {
			continue
		}
		if m := r.pattern.FindString(src); m != "" {
			out = append(out, r.name+" ("+m+")")
		}
	}
	return out
}

// TestTheGuardCatchesWhatItWasWrittenFor: each copy the audit found is caught outside its home
// and allowed inside it, and the ISO stamps stay allowed everywhere.
func TestTheGuardCatchesWhatItWasWrittenFor(t *testing.T) {
	caught := []string{
		`const raidChoiceDate = "Mon 2 Jan"`,
		`const confirmationTimeLayout = "3:04 pm MST"`,
		`const raidWhenLayout = "Mon 2 Jan, 15:04 MST"`,
		`const guildDateLayout = "Monday 2 January"`,
		`const sheetDateLayout = "2 Jan 2006"`,
		`const dayMonthLayout = "2 January"`,
		`return int(c.R)<<16 | int(c.G)<<8 | int(c.B)`,
	}
	for _, src := range caught {
		if len(guardHits("internal/discordbot/x.go", src)) == 0 {
			t.Errorf("the guard missed %s", src)
		}
	}
	if hits := guardHits("internal/textfmt/textfmt.go", caught[0]); len(hits) != 0 {
		t.Errorf("textfmt is the layouts' home, got %v", hits)
	}
	if hits := guardHits("internal/brand/paint.go", caught[len(caught)-1]); len(hits) != 0 {
		t.Errorf("brand is the colour packing's home, got %v", hits)
	}
	for _, src := range []string{`const RequestedAtLayout = "2006-01-02 15:04"`, `Format("2006-01-02T15:04Z")`, `Perm = 1 << 16`} {
		if hits := guardHits("internal/sims/x.go", src); len(hits) != 0 {
			t.Errorf("%s is not a reader-facing copy, got %v", src, hits)
		}
	}
}
