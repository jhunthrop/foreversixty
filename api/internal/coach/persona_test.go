package coach

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const fixtureDir = "testdata/persona"

func fixture(t *testing.T) string {
	t.Helper()
	data, err := readTestFile(fixtureDir + "/warrior-fury.md")
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func dropLine(prefix string) func(string) string {
	return func(s string) string {
		var kept []string
		for _, l := range strings.Split(s, "\n") {
			if !strings.HasPrefix(strings.TrimPrefix(l, "- "), prefix) {
				kept = append(kept, l)
			}
		}
		return strings.Join(kept, "\n")
	}
}

func swapSections(a, b string) func(string) string {
	return func(s string) string {
		s = strings.Replace(s, "## "+a+"\n", "## @@A\n", 1)
		s = strings.Replace(s, "## "+b+"\n", "## "+a+"\n", 1)
		return strings.Replace(s, "## @@A\n", "## "+b+"\n", 1)
	}
}

func insertAfter(heading, line string) func(string) string {
	return func(s string) string { return strings.Replace(s, heading+"\n", heading+"\n"+line+"\n", 1) }
}

func replace(old, repl string) func(string) string {
	return func(s string) string { return strings.Replace(s, old, repl, 1) }
}

func compose(fs ...func(string) string) func(string) string {
	return func(s string) string {
		for _, f := range fs {
			s = f(s)
		}
		return s
	}
}

func fixturePersona(t *testing.T) Persona {
	t.Helper()
	p, err := LoadPersona(fixtureDir, "warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func TestParsePersona_TheFixture(t *testing.T) {
	p := fixturePersona(t)
	if p.Spec != "warrior-fury" || p.Status != PersonaFull || !strings.HasPrefix(p.Version, "persona-") || p.Drills["drill.core_casts"] == "" {
		t.Fatalf("%+v", p)
	}
}

func TestPersonaVersion_IsTheContentHash(t *testing.T) {
	base := fixture(t)
	a, _ := ParsePersona([]byte(base))
	same, _ := ParsePersona([]byte(base))
	changed, err := ParsePersona([]byte(strings.Replace(base, "Short sentences.", "Short sentence.", 1)))
	if err != nil {
		t.Fatal(err)
	}
	if a.Version != same.Version || a.Version == changed.Version {
		t.Fatalf("versions %s %s %s", a.Version, same.Version, changed.Version)
	}
}

func TestPersonaVersion_StableAcrossCheckouts(t *testing.T) {
	base := fixture(t)
	want, _ := ParsePersona([]byte(base))
	crlf := strings.ReplaceAll(base, "\n", "\r\n")
	padded := strings.Replace(base, "## Voice\n", "## Voice  \n", 1) + "\n\n"
	for name, text := range map[string]string{"crlf": crlf, "trailing whitespace": padded} {
		got, err := ParsePersona([]byte(text))
		if err != nil || got.Version != want.Version {
			t.Errorf("%s: version %s (err %v), want %s", name, got.Version, err, want.Version)
		}
	}
}

// refusal is a mutation of the fixture, the line the loader must name and a message part.
// Fixture lines: spec 2, Drills heading 19, drill.death 27, What they know 15, Voice 30.
type refusal struct {
	mutate func(string) string
	line   int
	want   string
}

func TestParsePersona_Refuses(t *testing.T) {
	const drillsHeading = 19
	cases := map[string]refusal{
		"digit in prose":      {insertAfter("## Voice", "Press it 3 times."), 31, "digit"},
		"digit in What":       {insertAfter("## What they know", "Press it 3 times."), 16, "digit"},
		"digit in Who":        {insertAfter("## Who they are", "Press it 3 times."), 12, "digit"},
		"exclamation":         {insertAfter("## Voice", "Go!"), 31, "exclamation"},
		"exclamation drills":  {replace("Before the pull, find", "Go! Before the pull, find"), 27, "exclamation"},
		"heal more in prose":  {insertAfter("## Voice", "Just heal more."), 31, "heal more"},
		"heal more spaced":    {insertAfter("## Voice", "Just heal   more."), 31, "heal"},
		"heal more in drills": {replace("Before the pull, find", "Heal more. Before the pull, find"), 27, "heal more"},
		"sections reordered":  {swapSections("Drills", "Voice"), drillsHeading, "section"},
		"unknown drill id":    {replace("drill.death:", "drill.dying:"), 27, "drill.dying"},
		"duplicate drill id":  {insertAfter("## Drills", "- drill.death: Twice."), 28, "twice"},
		"missing a DPS drill": {dropLine("drill.adherence:"), drillsHeading, "drill.adherence"},
		"no reviewed field":   {dropLine("reviewed:"), noLine, "reviewed"},
		"unknown status":      {replace("status: full", "status: draft"), 4, "status"},
		"unknown role":        {replace("role: dps", "role: bard"), 3, "role"},
		"missing a tank drill": {compose(replace("role: dps", "role: tank"),
			dropLine("drill.adherence:")), drillsHeading, "drill.adherence"},
		"missing a healer drill": {compose(replace("role: dps", "role: healer"),
			dropLine("drill.ranks:")), drillsHeading, "drill.ranks"},
		"drill over the cap": {insertAfter("## Drills", "- drill.ranks: "+strings.Repeat("x", 241)), 20, "241"},
	}
	for name, c := range cases {
		_, err := ParsePersona([]byte(c.mutate(fixture(t))))
		var pe *PersonaError
		if !errors.As(err, &pe) {
			t.Errorf("%s: want a *PersonaError, got %v", name, err)
			continue
		}
		if pe.Line != c.line || !strings.Contains(strings.ToLower(pe.Msg), strings.ToLower(c.want)) {
			t.Errorf("%s: line %d %q, want line %d containing %q", name, pe.Line, pe.Msg, c.line, c.want)
		}
	}
}

func TestLoadPersona_ErrorNamesFileAndLine(t *testing.T) {
	dir := t.TempDir()
	text := insertAfter("## Voice", "Go!")(fixture(t))
	if err := os.WriteFile(filepath.Join(dir, "warrior-fury.md"), []byte(text), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := LoadPersona(dir, "warrior-fury")
	var pe *PersonaError
	if !errors.As(err, &pe) || pe.File != filepath.Join(dir, "warrior-fury.md") || pe.Line != 31 {
		t.Fatalf("got %+v (%v)", pe, err)
	}
	if !strings.Contains(err.Error(), "warrior-fury.md:31:") {
		t.Fatalf("message %q", err)
	}
}

func TestLoadPersona_SpecMismatchAndBadKey(t *testing.T) {
	dir := t.TempDir()
	text := replace("spec: warrior-fury", "spec: warrior-arms")(fixture(t))
	if err := os.WriteFile(filepath.Join(dir, "warrior-fury.md"), []byte(text), 0o600); err != nil {
		t.Fatal(err)
	}
	_, err := LoadPersona(dir, "warrior-fury")
	var pe *PersonaError
	if !errors.As(err, &pe) || pe.Line != 2 || !strings.Contains(pe.Msg, "does not match") {
		t.Fatalf("mismatch: %v", err)
	}
	for _, bad := range []string{"../warrior-fury", "a/b", "", "Warrior"} {
		if _, err := LoadPersona(fixtureDir, bad); err == nil {
			t.Errorf("spec %q accepted", bad)
		}
	}
}

func TestParsePersona_FrontMatterMayHoldDigitsAndBangs(t *testing.T) {
	s := replace("version: fixture", "version: build 1.60! edition")(fixture(t))
	if _, err := ParsePersona([]byte(s)); err != nil {
		t.Fatalf("front matter is not prose: %v", err)
	}
}

func TestParsePersona_HealerNeedsNoIdleOrAdherenceDrill(t *testing.T) {
	s := strings.Replace(fixture(t), "role: dps", "role: healer", 1)
	s = dropLine("drill.adherence:")(dropLine("drill.idle:")(s))
	if _, err := ParsePersona([]byte(s)); err != nil {
		t.Fatal(err)
	}
}

func TestParsePersona_VoiceRuleLineMayNameTheRule(t *testing.T) {
	ok := insertAfter("## Voice", "- Never tell a healer to heal more.")(fixture(t))
	if _, err := ParsePersona([]byte(ok)); err != nil {
		t.Fatalf("a Voice rule line states the rule (§8.1): %v", err)
	}
}

func TestPersonas_CachesBySpecAndReturnsCopies(t *testing.T) {
	ps := &Personas{Dir: fixtureDir}
	a, err := ps.For("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	a.Drills["drill.death"] = "mutated"
	b, _ := ps.For("warrior-fury")
	if b.Drills["drill.death"] == "mutated" {
		t.Fatal("a caller mutated the cached persona")
	}
	if a.Version != b.Version {
		t.Fatal("cache changed the version")
	}
	if _, err := ps.For("nonesuch"); err == nil {
		t.Fatal("missing spec accepted")
	}
}

func TestParsePersona_BannedPhraseSplitAcrossLinesIsCaught(t *testing.T) {
	bad := insertAfter("## Voice", "Just heal\nmore.")(fixture(t))
	if _, err := ParsePersona([]byte(bad)); err == nil || !strings.Contains(err.Error(), "heal") {
		t.Fatalf("a phrase split across a line break passed: %v", err)
	}
	ok := insertAfter("## Voice", "- Never tell a healer to heal more.")(fixture(t))
	if _, err := ParsePersona([]byte(ok)); err != nil {
		t.Fatalf("a Voice rule line is still exempt: %v", err)
	}
}
