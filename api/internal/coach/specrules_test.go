package coach

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const validRules = `{
  "spec": "warrior-fury", "build": "1.60.1.70630",
  "core": [{"family": "Bloodthirst", "spell_ids": [23881, 23894], "talent": true, "verified": "1.60.1.70630"}],
  "auras": [{"spell_id": 6673, "name": "Battle Shout", "kind": "self_buff", "verified": "1.60.1.70630"}],
  "downranks": [], "defensives": [], "efficiency": []
}`

func TestParseSpecRules_ReadsAStampedFile(t *testing.T) {
	r, err := ParseSpecRules([]byte(validRules))
	if err != nil {
		t.Fatal(err)
	}
	if r.Spec != "warrior-fury" || len(r.Core) != 1 || !r.Core[0].Talent || r.Core[0].SpellIDs[1] != 23894 {
		t.Fatalf("core %+v", r)
	}
	if len(r.Auras) != 1 || r.Auras[0].Kind != AuraSelfBuff {
		t.Fatalf("auras %+v", r.Auras)
	}
}

func TestParseSpecRules_RefusesUnstamped(t *testing.T) {
	cases := []struct{ name, from, to, wantErr string }{
		{"unstamped core entry", `"talent": true, "verified": "1.60.1.70630"`, `"talent": true, "verified": ""`, "Bloodthirst has no verified stamp"},
		{"core entry with no spell id", `"spell_ids": [23881, 23894]`, `"spell_ids": []`, "Bloodthirst has no spell id"},
		{"unstamped aura", `"kind": "self_buff", "verified": "1.60.1.70630"`, `"kind": "self_buff", "verified": ""`, "Battle Shout has no verified stamp"},
		{"unknown aura kind", `"kind": "self_buff"`, `"kind": "buffish"`, `Battle Shout has unknown kind "buffish"`},
		{"aura ranks without its spell id", `"spell_id": 6673,`, `"spell_id": 6673, "spell_ids": [5242],`, "Battle Shout lists ranks without its spell id 6673"},
		{"unknown field", `"downranks": []`, `"downrank": []`, "unknown field"},
		{"no spec", `"spec": "warrior-fury"`, `"spec": ""`, "no spec"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data := strings.Replace(validRules, c.from, c.to, 1)
			if data == validRules {
				t.Fatalf("the case did not change the file: %q", c.from)
			}
			if _, err := ParseSpecRules([]byte(data)); err == nil || !strings.Contains(err.Error(), c.wantErr) {
				t.Fatalf("err = %v, want one naming %q", err, c.wantErr)
			}
		})
	}
}

func TestLoadSpecRules(t *testing.T) {
	dir := t.TempDir()
	if _, ok, err := LoadSpecRules(dir, "warrior-fury"); ok || err != nil {
		t.Fatalf("a missing file is ok=false with no error, got ok=%v err=%v", ok, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "warrior-fury.json"), []byte(validRules), 0o600); err != nil {
		t.Fatal(err)
	}
	r, ok, err := LoadSpecRules(dir, "warrior-fury")
	if err != nil || !ok || r.Spec != "warrior-fury" {
		t.Fatalf("LoadSpecRules = %+v, %v, %v", r, ok, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "warrior-arms.json"), []byte(validRules), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, _, err := LoadSpecRules(dir, "warrior-arms"); err == nil {
		t.Fatal("a file whose spec field names another spec must be refused")
	}
}

// The fixture rule files obey the same rules as the shipped ones.
func TestFixtureRulesParse(t *testing.T) {
	for _, r := range []SpecRules{furyRules, protRules, holyRules} {
		for _, e := range append(append(append(append([]RuleEntry{}, r.Core...), r.Downranks...), r.Defensives...), r.Efficiency...) {
			if err := checkRuleEntry(e); err != nil {
				t.Errorf("%s: %v", r.Spec, err)
			}
		}
		for _, a := range r.Auras {
			if err := checkAuraRule(a); err != nil {
				t.Errorf("%s: %v", r.Spec, err)
			}
		}
	}
}
