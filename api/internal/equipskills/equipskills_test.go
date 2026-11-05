package equipskills

import (
	"strings"
	"testing"
)

const sample = `{
  "build": "1.60.1.1",
  "classes": {
    "warrior": {
      "armor": [{"subclass": 4, "skill": "Plate Mail", "level": 40}, {"subclass": 1, "skill": "Cloth", "level": 1}],
      "weapons": [{"subclass": 0, "skill": "Axes", "level": 1}],
      "dual_wield_level": 20
    },
    "shaman": {
      "armor": [{"subclass": 3, "skill": "Mail", "level": 40}],
      "weapons": [{"subclass": 13, "skill": "Fist Weapons", "level": 1}],
      "dual_wield_level": null
    }
  }
}`

func TestParseRefusesAnEmptyFile(t *testing.T) {
	for _, body := range []string{`{}`, `{"build": "x"}`, `{"classes": {"warrior": {}}}`, `not json`} {
		if _, err := Parse([]byte(body)); err == nil {
			t.Errorf("Parse(%q) took it, want an error", body)
		}
	}
}

func TestGoSourceRendersTheTable(t *testing.T) {
	doc, err := Parse([]byte(sample))
	if err != nil {
		t.Fatal(err)
	}
	src, err := GoSource(doc, SourcePath(doc.Build))
	if err != nil {
		t.Fatal(err)
	}
	got := string(src)
	for _, want := range []string{
		"from data/builds/1.60.1.1/equip-skills.json. DO NOT EDIT.",
		`const equipSkillsBuild = "1.60.1.1"`,
		"4: 40, // Plate Mail",
		"13: 1, // Fist Weapons",
		"dualWieldLevel: 20,",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("source lacks %q:\n%s", want, got)
		}
	}
	// Armour in subclass order, and a class that never dual wields has no dualWieldLevel at all.
	if strings.Index(got, "// Cloth") > strings.Index(got, "// Plate Mail") {
		t.Errorf("armour out of subclass order:\n%s", got)
	}
	if strings.Count(got, "dualWieldLevel") != 1 {
		t.Errorf("want one dualWieldLevel:\n%s", got)
	}
}
