package main

import (
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// ladderWeapon is the highest item-level ranged weapon this class,
// level and faction can equip, or nil. It exists because a hunter
// with no ranged weapon at all hits a real engine artifact: the
// weapon's SwingSpeed is the Weapon{} zero value, which collapses the
// ranged swing timer to its cast-time floor and saturates the
// character's action economy so completely the priority list never
// gets a decision window (recorded in data/curated/apl/hunter-
// beast-mastery.json's own notes, and true of every hunter spec that
// shares its rotation package). The rotation-accuracy ladder design
// this lane's brief points at states the same fix for the same
// reason: "a weapon set the engine's item database knows, chosen by
// level... Bare otherwise."
//
// Faction is deliberately not a filter here: the weights run answers
// "how much does this spec value a point of X stat", which is not a
// faction-specific question, and restricting the ladder weapon to one
// faction would need two weights runs per band instead of one - the
// brief's own budget ("one weights run per band") assumes a single,
// faction-agnostic ladder character. required_level is still checked,
// since an unwearable-at-this-level weapon would defeat the point.
func ladderWeapon(items []candidate, level int) *candidate {
	var best *candidate
	for i := range items {
		c := items[i]
		if c.RequiredLevel > level {
			continue
		}
		ranged := false
		for _, s := range c.Slots {
			if s == "ranged" {
				ranged = true
				break
			}
		}
		if !ranged {
			continue
		}
		if best == nil || c.ItemLevel > best.ItemLevel {
			best = &c
		}
	}
	return best
}

// ladderCharacter is the bare-but-armed character the weights run
// measures: no armor, the truncated talent build, and the best
// ranged weapon this level and faction allow (see ladderWeapon).
func ladderCharacter(race, classSlug string, level int, talents string, weapon *candidate) api.CharacterSpec {
	ch := api.CharacterSpec{
		Name:    "ladder",
		Race:    race,
		Class:   classSlug,
		Level:   level,
		Talents: talents,
	}
	if weapon != nil {
		ch.Gear = []api.GearSlot{{Slot: "ranged", ItemID: weapon.ID}}
	}
	return ch
}

// referenceStatOverride works around a real finding this lane's run
// turned up: data/curated/specs.json's reference_stat for both hunter
// specs on the canonical list (marksmanship here; beast-mastery
// shares the value) is "attack_power" - melee attack power - but a
// bare ranged-weapon-only ladder character (no melee weapon, never in
// melee range) measures its melee attack_power weight as EXACTLY
// zero: sim/adapter.Weights refuses to normalise against a reference
// stat that weighs nothing (ErrNoWeights: "%s weighs nothing"), which
// makes every weights run for this spec fail outright with the
// engine's own reference_stat. ranged_attack_power is what a
// ranged-primary hunter spec should normalise against, and it does
// measure nonzero here; this map is the prototype's workaround until
// the curated data is corrected (see this lane's report).
var referenceStatOverride = map[string]string{
	"hunter-marksmanship":  "ranged_attack_power",
	"hunter-beast-mastery": "ranged_attack_power",
}

// weightsRequest builds the SimRequest runWeights takes: the spec's
// own weight_stats (data/curated/specs.json), and its reference_stat
// unless referenceStatOverride names a different one for this spec.
func weightsRequest(spec specInfo, ch api.CharacterSpec, iterations int, seed int64) api.SimRequest {
	reference := spec.ReferenceStat
	if override, ok := referenceStatOverride[spec.Spec]; ok {
		reference = override
	}
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     api.DefaultEncounter(),
		Iterations:    iterations,
		RandomSeed:    seed,
		Weights:       &api.WeightsSpec{Stats: spec.WeightStats, Reference: reference},
	}
}

// plainRequest builds an ordinary DPS-run SimRequest for ch, at
// iterations/seed - used by verify.go for the baseline and each
// per-slot swap.
func plainRequest(spec specInfo, ch api.CharacterSpec, iterations int, seed int64) api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     api.DefaultEncounter(),
		Iterations:    iterations,
		RandomSeed:    seed,
	}
}
