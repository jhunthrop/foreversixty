// Package slopguard is the api half of the AI-slop copy guard (owner 2026-11-09): every pattern
// in tools/slop/patterns.txt over the api copy trees listed in tools/slop/sources.txt. It mirrors
// tools/slop/scan.mjs, which the web vitest runs over everything outside api/. A real game name
// or a legal phrase goes in tools/slop/allowlist.txt, never a looser pattern.
package slopguard

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// repoRoot is the repository root, relative to this package directory.
const repoRoot = "../../.."

const generatedHeaderLines = 5

var generatedHeader = regexp.MustCompile(`(?i)generated|do not edit`)

type pattern struct {
	id string
	re *regexp.Regexp
}

type allowed struct {
	id   string
	path *regexp.Regexp
	text string
}

// dataLines reads a tools/slop file without its comments and blank lines.
func dataLines(t *testing.T, name string) []string {
	t.Helper()
	f, err := os.Open(filepath.Join(repoRoot, "tools", "slop", name))
	if err != nil {
		t.Fatalf("open %s: %v", name, err)
	}
	defer f.Close()
	var out []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimRight(sc.Text(), "\r")
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		out = append(out, line)
	}
	if err := sc.Err(); err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return out
}

func fields(t *testing.T, name, line string, n int) []string {
	t.Helper()
	parts := strings.Split(line, "\t")
	if len(parts) != n {
		t.Fatalf("%s: want %d tab-separated fields, got %d: %s", name, n, len(parts), line)
	}
	return parts
}

func loadPatterns(t *testing.T) []pattern {
	t.Helper()
	var out []pattern
	for _, line := range dataLines(t, "patterns.txt") {
		p := fields(t, "patterns.txt", line, 3)
		src := p[2]
		if p[1] == "i" {
			src = "(?i)" + src
		}
		out = append(out, pattern{id: p[0], re: regexp.MustCompile(src)})
	}
	return out
}

func loadAllowlist(t *testing.T) []allowed {
	t.Helper()
	var out []allowed
	for _, line := range dataLines(t, "allowlist.txt") {
		p := fields(t, "allowlist.txt", line, 3)
		out = append(out, allowed{id: p[0], path: regexp.MustCompile(p[1]), text: p[2]})
	}
	return out
}

// apiCopyFiles lists the .go copy files under the api/ entries of sources.txt.
func apiCopyFiles(t *testing.T) []string {
	t.Helper()
	var include []string
	var exclude []*regexp.Regexp
	for _, line := range dataLines(t, "sources.txt") {
		switch {
		case strings.HasPrefix(line, "+ api/"):
			include = append(include, strings.TrimSpace(line[2:]))
		case strings.HasPrefix(line, "- "):
			exclude = append(exclude, regexp.MustCompile(strings.TrimSpace(line[2:])))
		}
	}
	if len(include) == 0 {
		t.Fatal("sources.txt lists no api/ tree")
	}
	var files []string
	for _, root := range include {
		err := filepath.WalkDir(filepath.Join(repoRoot, root), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(repoRoot, path)
			rel = filepath.ToSlash(rel)
			if d.IsDir() || filepath.Ext(rel) != ".go" || excluded(rel, exclude) {
				return nil
			}
			files = append(files, rel)
			return nil
		})
		if err != nil {
			t.Fatalf("walk %s: %v", root, err)
		}
	}
	return files
}

func excluded(path string, exclude []*regexp.Regexp) bool {
	for _, re := range exclude {
		if re.MatchString(path) {
			return true
		}
	}
	return false
}

var lineComment = regexp.MustCompile(`\s//(\s|$)`)

// copyLines drops /* */ blocks and // comments, as scan.mjs does for C-style sources.
func copyLines(text string) []string {
	var out []string
	inBlock := false
	for _, raw := range strings.Split(text, "\n") {
		line, kept := raw, ""
		for line != "" {
			if inBlock {
				end := strings.Index(line, "*/")
				if end == -1 {
					line = ""
					break
				}
				line, inBlock = line[end+2:], false
				continue
			}
			open := strings.Index(line, "/*")
			if open == -1 {
				kept += line
				break
			}
			kept, line, inBlock = kept+line[:open], line[open+2:], true
		}
		if strings.HasPrefix(strings.TrimSpace(kept), "//") {
			kept = ""
		} else if at := lineComment.FindStringIndex(kept); at != nil {
			kept = kept[:at[0]]
		}
		out = append(out, kept)
	}
	return out
}

func isGenerated(text string) bool {
	lines := strings.SplitN(text, "\n", generatedHeaderLines+1)
	for i := 0; i < len(lines) && i < generatedHeaderLines; i++ {
		if generatedHeader.MatchString(lines[i]) {
			return true
		}
	}
	return false
}

func isAllowed(list []allowed, id, path, line string) bool {
	for _, a := range list {
		if a.id == id && a.path.MatchString(path) && strings.Contains(line, a.text) {
			return true
		}
	}
	return false
}

func TestAPICopyHasNoSlop(t *testing.T) {
	patterns, allow := loadPatterns(t), loadAllowlist(t)
	for _, path := range apiCopyFiles(t) {
		body, err := os.ReadFile(filepath.Join(repoRoot, path))
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		if isGenerated(string(body)) {
			continue
		}
		for i, line := range copyLines(string(body)) {
			for _, p := range patterns {
				if p.re.MatchString(line) && !isAllowed(allow, p.id, path, line) {
					t.Errorf("%s:%d: [%s] %s", path, i+1, p.id, strings.TrimSpace(line))
				}
			}
		}
	}
}

func TestPatternsStillCatchTheTells(t *testing.T) {
	patterns := loadPatterns(t)
	samples := map[string]string{
		"em-dash":      "not enough logs yet — 4 kills",
		"nudge-prefix": "Next: spend your 1 talent point.",
		"here-is-what": "Here is what your log shows for Ragnaros.",
		"grade-hype":   "Fury is graded D this week",
	}
	for id, text := range samples {
		hit := false
		for _, p := range patterns {
			hit = hit || (p.id == id && p.re.MatchString(text))
		}
		if !hit {
			t.Errorf("pattern %s no longer catches %q", id, text)
		}
	}
	for _, p := range patterns {
		if p.re.MatchString("const noFigure = \"—\"") {
			t.Errorf("pattern %s flags an empty-cell dash", p.id)
		}
	}
}
