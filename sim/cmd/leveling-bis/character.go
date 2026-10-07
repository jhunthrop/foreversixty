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

// ladderMeleeWeapons arms the weights character for melee the way
// ladderWeapon arms it for ranged: the best melee weapons this level
// allows, by weapon DPS, so crit, haste and hit are weighed on real
// swings instead of fists (the 2026-10-07 survival-raid review found the
// fist-swinging character's 1.0 s swing stepping on Raptor Strike's
// 6 s cooldown, which made melee haste weigh NEGATIVE for Survival).
// Dual-wield specs get a one-hand pair; a two-hander wins the main hand
// only where no pair is offered; shamans and the rest take the single
// strongest weapon. Casters (leveling.NoMeleeAutoAttackSpecs) never swing,
// so they stay unarmed in melee.
func ladderMeleeWeapons(items []candidate, level int, spec string) []api.GearSlot {
	if leveling.NoMeleeAutoAttackSpecs[spec] {
		return nil
	}
	// bestOne is the strongest main-hand-capable one-hander; bestOff the
	// strongest off-hand-capable one-hander that is not bestOne (an item
	// the engine marks main-hand-only, such as Andonisus, would overwrite
	// the main hand instead of pairing with it: the 2026-10-07 ratings
	// review found the Fury weights character swinging one weapon).
	var bestOne, bestOff, bestTwo *candidate
	for i := range items {
		c := items[i]
		if c.RequiredLevel > level || c.DPS <= 0 {
			continue
		}
		main, off := false, false
		for _, s := range c.Slots {
			switch s {
			case "main_hand":
				main = true
			case "off_hand":
				off = true
			}
		}
		if !main && !off {
			continue
		}
		if c.TwoHand {
			if bestTwo == nil || c.DPS > bestTwo.DPS {
				bestTwo = &c
			}
			continue
		}
		if main && (bestOne == nil || c.DPS > bestOne.DPS) {
			bestOne = &c
		}
	}
	if leveling.DualWieldSpecs[spec] && bestOne != nil {
		for i := range items {
			c := items[i]
			if c.ID == bestOne.ID || c.RequiredLevel > level || c.DPS <= 0 || c.TwoHand {
				continue
			}
			off := false
			for _, s := range c.Slots {
				if s == "off_hand" {
					off = true
				}
			}
			if off && (bestOff == nil || c.DPS > bestOff.DPS) {
				bestOff = &c
			}
		}
		if bestOff != nil {
			return []api.GearSlot{{Slot: "main_hand", ItemID: bestOne.ID}, {Slot: "off_hand", ItemID: bestOff.ID}}
		}
	}
	best := bestOne
	if bestTwo != nil && (best == nil || bestTwo.DPS > best.DPS) {
		best = bestTwo
	}
	if best == nil {
		return nil
	}
	return []api.GearSlot{{Slot: "main_hand", ItemID: best.ID}}
}

// ladderCharacter is the bare-but-armed character the weights run
// measures: no armor, the truncated talent build, and the best
// ranged weapon this level and faction allow (see ladderWeapon).
//
// Deliberately NOT given leveling.NoMeleeAutoAttackSpecs' own DistanceFromTarget
// treatment (bandCharacter's own doc, this lane's brief item 7): this
// lane's own brief scopes the fix to "the ranker's tournament
// character" - the weights sweep this character measures is a
// separate, much larger-blast-radius concern (every score() weight for
// every item in the band reads it) this lane did not verify by hand.
// specSlug "" is bandCharacter's own signal for "not applicable".
func ladderCharacter(race, classSlug string, level int, talents string, weapon *candidate, melee ...api.GearSlot) api.CharacterSpec {
	var gear []api.GearSlot
	if weapon != nil {
		gear = []api.GearSlot{{Slot: "ranged", ItemID: weapon.ID}}
	}
	gear = append(gear, melee...)
	return bandCharacter("ladder", race, classSlug, "", level, talents, gear)
}

// bandCharacter is the one constructor every plain-DPS site in this
// band's rank/verify passes (trinkets.go, rank.go, sets.go, verify.go)
// builds its api.CharacterSpec from. Before this existed, each of those
// sites built its own literal and none of them set Talents, so every
// verify/rank/set-completion sim ran the band's gear on a bare,
// talent-less character while ladderCharacter's own weights run (above)
// carried the band's real talent string - a direct sim of the published
// BM level-60 set found this: 227 DPS with talents equipped, 201
// without, yet the published set_dps was 200.15 for BOTH BM and MM at
// every band (two different talent builds producing an identical DPS to
// 14 decimals is the tell that talents were never in the request at
// all). One constructor used by every site closes the gap the same way
// everywhere, and gives the guard in character_test.go one place to
// assert Talents is always set from the band's own string.
//
// specSlug (this lane's brief, item 7) is the full spec key
// (data/curated/specs.json's own "spec", e.g. "shaman-elemental") -
// leveling.NoMeleeAutoAttackSpecs' own doc explains why a caster spec's
// tournament character stands out of melee range instead of at the
// engine's own zero-value default. Every caller already carries a
// specInfo with this exact value in scope (spec.Spec) at its own
// bandCharacter call site; ladderCharacter (above) deliberately passes
// "" instead - out of this lane's own scope.
func bandCharacter(name, race, classSlug, specSlug string, level int, talents string, gear []api.GearSlot) api.CharacterSpec {
	ch := api.CharacterSpec{
		Name:    name,
		Race:    race,
		Class:   classSlug,
		Level:   level,
		Talents: talents,
		Gear:    gear,
	}
	if leveling.NoMeleeAutoAttackSpecs[specSlug] {
		ch.DistanceFromTarget = leveling.CasterDistanceFromTarget
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
	ch = withKit(spec, ch)
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     rankerEncounter(),
		Iterations:    iterations,
		RandomSeed:    seed,
		Weights:       &api.WeightsSpec{Stats: spec.WeightStats, Reference: spec.ReferenceStat},
	}
}

// withKit gives ch the class kit the ladder's character carries too:
// rogue poisons from 20, shaman-enhancement's weapon imbues throughout
// (a rogue verified without poisons ranked its level-20 set at a tenth
// of a hunter's), and the self buffs leveling.KitBuffs names. Keyed by
// the full spec slug, not ch.Class, because the kit is a spec property.
// Every request the ranker builds - the DPS runs and the stat-weight
// sweep alike - goes through it, so the published weights come from the
// same buffed character the picks are scored on. A character that
// already carries its own consumes or buffs keeps them. A spec ranked
// under a preset (specInfo.Applied) carries the preset's buffs, debuffs
// and consumables with the kit on top (request.ResolvedPreset.Layer).
func withKit(spec specInfo, ch api.CharacterSpec) api.CharacterSpec {
	kitConsumes := leveling.KitConsumes(spec.Spec, ch.Level)
	kitBuffs := leveling.KitBuffs(spec.Spec, ch.Level)
	if spec.Applied != nil {
		kitBuffs, kitConsumes = spec.Applied.Layer(kitBuffs, kitConsumes)
	}
	if ch.Consumes == nil {
		ch.Consumes = kitConsumes
	}
	if ch.Buffs == nil {
		ch.Buffs = kitBuffs
	}
	return ch
}

// plainRequest builds an ordinary DPS-run SimRequest for ch, at
// iterations/seed - used by verify.go for the baseline and each
// per-slot swap.
func plainRequest(spec specInfo, ch api.CharacterSpec, iterations int, seed int64) api.SimRequest {
	ch = withKit(spec, ch)
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     rankerEncounter(),
		Iterations:    iterations,
		RandomSeed:    seed,
	}
}

// rankerEncounter is the default encounter with no creature type: a raid
// boss never grants a slaying bonus (the engine's own default target type
// is Humanoid, which would credit humanoid-slaying gear), so every
// published number is against a level-63 target of unknown type and a
// "+45 Attack Power against Undead" bracer counts for exactly nothing.
func rankerEncounter() api.EncounterSpec {
	e := api.DefaultEncounter()
	e.TargetType = api.TargetTypeUnknown
	return e
}
