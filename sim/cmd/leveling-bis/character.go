package main

import (
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/leveling"
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

// weightsRequest builds the SimRequest runWeights takes: the spec's own
// weight_stats and reference_stat (data/curated/specs.json). All three
// hunter dps specs now carry "ranged_attack_power" there (the rotation
// accuracy program's data-weapons lane, 2026-09-28) rather than melee
// "attack_power" - a bare ranged-weapon-only ladder character (see
// ladderCharacter above) never enters melee, so its melee attack_power
// weight measures EXACTLY zero and sim/adapter.Weights refuses to
// normalise against a reference stat that weighs nothing (ErrNoWeights:
// "%s weighs nothing"). A referenceStatOverride map used to work around
// this here; the curated data is the fix now, so this function just reads
// the spec's own reference_stat.
func weightsRequest(spec specInfo, ch api.CharacterSpec, iterations int, seed int64) api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     api.DefaultEncounter(),
		Iterations:    iterations,
		RandomSeed:    seed,
		Weights:       &api.WeightsSpec{Stats: spec.WeightStats, Reference: spec.ReferenceStat},
	}
}

// plainRequest builds an ordinary DPS-run SimRequest for ch, at
// iterations/seed - used by verify.go for the baseline and each
// per-slot swap.
func plainRequest(spec specInfo, ch api.CharacterSpec, iterations int, seed int64) api.SimRequest {
	if ch.Consumes == nil {
		// The class kit the ladder's character carries too (rogue
		// poisons from 20, shaman-enhancement's weapon imbues
		// throughout): a rogue verified without poisons ranked its
		// level-20 set at a tenth of a hunter's. Keyed by the full spec
		// slug, not ch.Class, because the kit is a spec property.
		ch.Consumes = leveling.KitConsumes(spec.Spec, ch.Level)
	}
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
