package specs

import (
	"slices"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
)

// physicalForbidden and casterForbidden are the two content rules task
// 5(c) pins: a physical spec (ReferenceStat attack_power) is never
// asked about a stat only a caster cares about, and a caster spec
// (ReferenceStat spell_power) is never asked about a melee-only stat.
// Named as tables, per spec, so the test states the rule rather than
// re-deriving it - and so a stat added to either list here is a
// deliberate content decision, not a typo.
var physicalForbidden = []string{
	"spirit", "mp5", "intellect", "spell_power", "spell_haste",
	"spell_penetration", "arcane_power", "fire_power", "frost_power",
	"holy_power", "nature_power", "shadow_power",
}

var casterForbidden = []string{"strength", "expertise", "armor_penetration", "feral_attack_power"}

func TestEverySpecHasANonEmptyWeightStats(t *testing.T) {
	for _, spec := range All {
		if len(spec.WeightStats) == 0 {
			t.Errorf("%s has an empty WeightStats", spec.Spec)
		}
	}
}

func TestEveryWeightStatIsAKnownStat(t *testing.T) {
	for _, spec := range All {
		for _, stat := range spec.WeightStats {
			if !slices.Contains(api.KnownStats, stat) {
				t.Errorf("%s.WeightStats has %q, not in api.KnownStats", spec.Spec, stat)
			}
		}
	}
}

// TestTheReferenceStatIsAlwaysInItsOwnWeightStats pins the rule
// api.WeightsSpec.validate depends on: a reference stat that is not
// among the stats being weighed makes a spec's own default weights
// request illegal.
func TestTheReferenceStatIsAlwaysInItsOwnWeightStats(t *testing.T) {
	for _, spec := range All {
		if !slices.Contains(spec.WeightStats, spec.ReferenceStat) {
			t.Errorf("%s: reference stat %q is not in its own WeightStats %v",
				spec.Spec, spec.ReferenceStat, spec.WeightStats)
		}
	}
}

func TestNoPhysicalSpecCarriesACasterStatAndNoCasterSpecCarriesAMeleeStat(t *testing.T) {
	for _, spec := range All {
		forbidden := casterForbidden
		if spec.ReferenceStat == "attack_power" {
			forbidden = physicalForbidden
		}
		for _, stat := range spec.WeightStats {
			if slices.Contains(forbidden, stat) {
				t.Errorf("%s (reference %s) carries forbidden stat %q",
					spec.Spec, spec.ReferenceStat, stat)
			}
		}
	}
}

// TestFeralAttackPowerBelongsToDruidFeralAndNowhereElse pins the
// brief's one named-stat rule directly, rather than leaving it to
// fall out of the forbidden-pair table alone.
func TestFeralAttackPowerBelongsToDruidFeralAndNowhereElse(t *testing.T) {
	for _, spec := range All {
		carries := slices.Contains(spec.WeightStats, "feral_attack_power")
		want := spec.Spec == "druid-feral"
		if carries != want {
			t.Errorf("%s carries feral_attack_power=%v, want %v", spec.Spec, carries, want)
		}
	}
}
