package main

import (
	"fmt"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core/proto"
)

// Finisher ability names the rogue specs' reports tally. Eviscerate and
// Slice and Dice are the ones a winner must keep casting.
const (
	abilityRupture    = "Rupture"
	abilityEviscerate = "Eviscerate"
	abilitySliceDice  = "Slice and Dice"
	abilityColdBlood  = "Cold Blood"
	abilityVenom      = "Venom"
)

// talentSpellNames names the talent spells the learned-ability table
// (names) does not carry: Cold Blood is a talent, not a trained rank.
var talentSpellNames = map[int]string{14177: abilityColdBlood}

// requiredFinishers are the abilities whose zero casts in a winner make
// it not adoptable.
var requiredFinishers = []string{abilityEviscerate, abilitySliceDice}

// finisherAbilities is the per-spec list of abilities the "Finisher
// casts" section tallies; a spec absent here has no such section.
var finisherAbilities = map[string][]string{
	"rogue-assassination": {abilityRupture, abilityEviscerate, abilitySliceDice, abilityColdBlood, abilityVenom},
	"rogue-combat":        {abilityRupture, abilityEviscerate, abilitySliceDice},
	"rogue-subtlety":      {abilityRupture, abilityEviscerate, abilitySliceDice},
}

// castTable is casts per iteration by ability name.
type castTable map[string]float64

// castsPerIteration folds the player's per-spell-id cast counts into
// per-iteration casts by ability name, for the wanted abilities only
// (every wanted name is present, zero when never cast).
func castsPerIteration(counts map[int64]int64, iterations int, names map[int]string, wanted []string) castTable {
	out := castTable{}
	for _, w := range wanted {
		out[w] = 0
	}
	for id, n := range counts {
		name, ok := names[int(id)]
		if !ok {
			name = talentSpellNames[int(id)]
		}
		if _, ok := out[name]; ok {
			out[name] += float64(n) / float64(iterations)
		}
	}
	return out
}

// tallyCasts reads a sim's player metrics into a castTable.
func tallyCasts(player *proto.UnitMetrics, iterations int, names map[int]string, wanted []string) castTable {
	counts := map[int64]int64{}
	for id, n := range request.CastSpellCounts(player) {
		counts[int64(id)] = n
	}
	return castsPerIteration(counts, iterations, names, wanted)
}

// notAdoptable lists the required finishers the table never casts.
func (t castTable) notAdoptable() []string {
	var out []string
	for _, name := range requiredFinishers {
		if t[name] == 0 {
			out = append(out, name)
		}
	}
	return out
}

// finisherCasts is the baseline's and the winner's cast tables.
type finisherCasts struct {
	abilities      []string
	baseline, best castTable
}

func (f finisherCasts) markdown() string {
	var b strings.Builder
	b.WriteString("## Finisher casts\n\n")
	b.WriteString("Casts per iteration, read from the sim's cast metrics at the confirm iteration count.\n\n")
	b.WriteString("| Ability | Baseline | Winner |\n|---|---|---|\n")
	for _, a := range f.abilities {
		fmt.Fprintf(&b, "| %s | %.2f | %.2f |\n", a, f.baseline[a], f.best[a])
	}
	b.WriteString("\n")
	if missing := f.best.notAdoptable(); len(missing) > 0 {
		fmt.Fprintf(&b, "**NOT ADOPTABLE: the winner never casts %s.** Its DPS gain comes from dropping a finisher, so it is not a rotation to paste.\n\n", strings.Join(missing, " or "))
	}
	return b.String()
}

// tallyFinishers sims the baseline and the winner once more each and
// tallies their finisher casts; nil when the spec has no finisher list.
func tallyFinishers(in inputs, base, best rotation, iterations int, names map[int]string) (*finisherCasts, error) {
	abilities, ok := finisherAbilities[in.setup.spec.Spec]
	if !ok {
		return nil, nil
	}
	tables := make([]castTable, 0, 2)
	for _, r := range []rotation{base, best} {
		_, player, err := simRotation(in.setup, r, iterations)
		if err != nil {
			return nil, fmt.Errorf("tallying finisher casts: %w", err)
		}
		tables = append(tables, tallyCasts(player, iterations, names, abilities))
	}
	return &finisherCasts{abilities: abilities, baseline: tables[0], best: tables[1]}, nil
}
