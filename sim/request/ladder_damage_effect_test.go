package request

import "testing"

// TestIsDamageSpellIDRecognisesAPeriodicTriggerGroundEffect is this
// wave's harness rule 1 (this lane's brief): isDamageEffect's old rule -
// effect 6 with aura 3 only - missed the ground-effect pattern
// Consecration, Blizzard, Rain of Fire, Hurricane and Volley all use in
// this build's own spellconst (data/builds/<build>/spellconst/<class>.json):
// effect 6 with aura 226 (a periodic TRIGGER, not a periodic damage aura
// itself - the container spell applies an aura whose own tick casts the
// real damage spell), which a learned-but-unused Consecration surfaced
// as a false positive on the ladder's "learned but unused" report. Every
// row below is read from the real, committed build rather than a
// synthetic fixture, so the test fails the moment the pipeline's own
// spellconst shape for these spells changes out from under the rule.
func TestIsDamageSpellIDRecognisesAPeriodicTriggerGroundEffect(t *testing.T) {
	build := activeBuild(t)

	tests := []struct {
		class string
		id    int
		name  string
		want  bool
	}{
		// Effect 6 / aura 226 (periodic trigger spell): the AoE
		// ground-effect pattern this lane's brief names explicitly.
		{"paladin", 20116, "Consecration", true},
		{"mage", 6141, "Blizzard", true},
		{"warlock", 5740, "Rain of Fire", true},
		{"druid", 16914, "Hurricane", true},
		{"hunter", 1510, "Volley", true},
		// Effect 6 / aura 3 (a direct periodic-damage aura): already
		// handled before this lane, kept here so the same table proves
		// the old rule did not regress.
		{"warlock", 348, "Immolate", true},
		{"warlock", 172, "Corruption", true},
		// A negative control: Arcane Intellect is effect 6 / aura 29 (a
		// stat-modifier aura), which must NOT read as a damage effect -
		// otherwise every buff in the game would show up as "learned but
		// unused" on the ladder's informational report.
		{"mage", 1459, "Arcane Intellect", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			consts, err := loadSpellConst(repoRoot, build, tc.class)
			if err != nil {
				t.Fatal(err)
			}
			entry, ok := consts[tc.id]
			if !ok {
				t.Fatalf("spellconst/%s.json has no spell %d (%s) - has the build's data moved?", tc.class, tc.id, tc.name)
			}
			if entry.Name != tc.name {
				t.Fatalf("spellconst/%s.json spell %d is %q, want %q - the id no longer names the spell this table expects",
					tc.class, tc.id, entry.Name, tc.name)
			}
			if got := isDamageSpellID(consts, tc.id); got != tc.want {
				t.Errorf("isDamageSpellID(%s %d %q) = %v, want %v (effects=%+v)",
					tc.class, tc.id, tc.name, got, tc.want, entry.Effects)
			}
		})
	}
}
