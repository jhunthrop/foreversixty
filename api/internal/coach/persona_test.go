package coach

import (
	"strings"
	"testing"
)

const fixtureDir = "testdata/persona"

func fixture(t *testing.T) string {
	t.Helper()
	p := fixtureDir + "/warrior-fury.md"
	data, err := readTestFile(p)
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

func TestParsePersona_Refuses(t *testing.T) {
	for name, mutate := range map[string]func(string) string{
		"digit in prose":     func(s string) string { return strings.Replace(s, "## Voice\n", "## Voice\nPress it 3 times.\n", 1) },
		"exclamation":        func(s string) string { return strings.Replace(s, "## Voice\n", "## Voice\nGo!\n", 1) },
		"heal more in prose": func(s string) string { return strings.Replace(s, "## Voice\n", "## Voice\nJust heal more.\n", 1) },
		"heal more in drills": func(s string) string {
			return strings.Replace(s, "## Drills\n", "## Drills\n- drill.ranks: Heal more.\n", 1)
		},
		"sections reordered":  swapSections("Drills", "Voice"),
		"unknown drill id":    func(s string) string { return strings.Replace(s, "drill.death:", "drill.dying:", 1) },
		"missing a DPS drill": dropLine("drill.adherence:"),
		"no reviewed field":   dropLine("reviewed:"),
		"unknown status":      func(s string) string { return strings.Replace(s, "status: full", "status: draft", 1) },
		"unknown role":        func(s string) string { return strings.Replace(s, "role: dps", "role: bard", 1) },
		"drill over the cap": func(s string) string {
			return strings.Replace(s, "## Drills\n", "## Drills\n- drill.ranks: "+strings.Repeat("x", 241)+"\n", 1)
		},
	} {
		if _, err := ParsePersona([]byte(mutate(fixture(t)))); err == nil {
			t.Errorf("%s: accepted", name)
		}
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
	ok := strings.Replace(fixture(t), "## Voice\n", "## Voice\n- Never tell a healer to heal more.\n", 1)
	if _, err := ParsePersona([]byte(ok)); err != nil {
		t.Fatalf("a Voice rule line states the rule (§8.1): %v", err)
	}
}

func TestPersonas_CachesBySpec(t *testing.T) {
	ps := &Personas{Dir: fixtureDir}
	a, err := ps.For("warrior-fury")
	if err != nil {
		t.Fatal(err)
	}
	b, _ := ps.For("warrior-fury")
	if a.Version != b.Version {
		t.Fatal("cache changed the version")
	}
	if _, err := ps.For("nonesuch"); err == nil {
		t.Fatal("missing spec accepted")
	}
}
