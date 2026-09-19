package sims

import (
	"log/slog"
	"slices"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// buffVocabulary maps a recorded buff's spell id onto the engine's own
// id for it, the vocabulary generated into sim/request/IDS.md.
//
// Two different things are being joined here: a combat log records what
// was on the player (a spell id, per rank), and the engine names a buff
// by a protobuf field. Only ids verified in both directions belong in
// this table - the spell id against the build's spells.json, the id
// against IDS.md, which buffs_test.go re-checks on every run. An id
// missing from here is dropped and logged, which is a sim run without
// that buff and a line saying so; an id guessed wrong would be a wrong
// number with nothing saying anything.
var buffVocabulary = map[int64]string{
	// Paladin blessings, greater and lesser, every rank the client has.
	20217: "blessing_of_kings",
	25898: "blessing_of_kings",
	19838: "blessing_of_might",
	25291: "blessing_of_might",
	25916: "blessing_of_might",
	19979: "blessing_of_wisdom",
	25290: "blessing_of_wisdom",
	25918: "blessing_of_wisdom",
	// Warrior shouts.
	25289: "battle_shout",
	2048:  "battle_shout",
	// Priest stamina and shadow protection.
	10938: "power_word_fortitude",
	21562: "power_word_fortitude",
	10958: "shadow_protection",
	27683: "shadow_protection",
	// Druid: IDS.md has no individual-target id for this, only the
	// Forever ruleset's raid-wide field, which is what a cast of this
	// buff means for sim purposes regardless of who was targeted.
	9885:  "gift_of_the_wild",
	21850: "gift_of_the_wild",
	// Mage.
	10157: "arcane_brilliance",
	23028: "arcane_brilliance",
}

// BuffIDs turns a fight's recorded auras into the engine's buff ids,
// sorted and deduplicated. An aura this table does not know is dropped
// with a line naming it, so the gap is visible in the logs and the
// number is honest about what it ran with.
func BuffIDs(refs []summary.AuraRef, log *slog.Logger) []string {
	if log == nil {
		log = slog.Default()
	}
	seen := map[string]bool{}
	out := []string{}
	for _, ref := range refs {
		id, ok := buffVocabulary[ref.SpellID]
		if !ok {
			log.Debug("sims", "op", "sim input", "dropped_buff", ref.SpellID,
				"name", ref.Name, "reason", "no engine id for that spell")
			continue
		}
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	slices.Sort(out)
	return out
}
