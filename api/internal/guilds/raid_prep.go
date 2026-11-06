// api/internal/guilds/raid_prep.go
//
// The raid-prep readings (owner decision D14) from the export's school=, supplies=, locks=
// and quests= sections (fs1/prep.go): flask and food on hand, resistance totals and
// attunement. One computation read by the member's checklist (me_checklist.go) and the
// officer's Readiness rows (readiness.go), so the two never disagree. An export from an addon
// too old to write a section reads as not tracked, never as missing.
package guilds

import (
	"sort"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// exportPrep is the raid-prep part of one roster row's export, already decoded.
type exportPrep struct {
	bank        []int
	resistances *fs1.Resistances
	supplies    map[int]int
	hasSupplies bool
	keyring     []int
	hasKeyring  bool
	quests      map[int]bool
	hasQuests   bool
}

// prepOf is the raid-prep part of a decoded export.
func prepOf(d fs1.Decoded) exportPrep {
	return exportPrep{
		bank: d.Bank, resistances: d.Resistances,
		supplies: d.Supplies, hasSupplies: d.HasSupplies,
		keyring: d.Keyring, hasKeyring: d.HasKeyring,
		quests: d.Quests, hasQuests: d.HasQuests,
	}
}

// What a flask-and-food reading can be missing.
const (
	missingFlaskOrElixir = "flask_or_elixir"
	missingFood          = "food"
)

// SupplyStack is one counted consumable on hand.
type SupplyStack struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Kind  string `json:"kind"`
	Count int    `json:"count"`
}

// suppliesReading is what the bags hold that a raid leader counts. The minimum is one flask
// or one raid elixir, and one buff food: a single flask with no food is not ready, and
// neither is a stack of food with no flask or elixir.
type suppliesReading struct {
	Stacks  []SupplyStack
	Missing []string
}

// ready reports whether the reading meets the minimum.
func (r suppliesReading) ready() bool { return len(r.Missing) == 0 }

// readSupplies names the counted stacks in supplies, flasks first, then elixirs, then food,
// each by item id, and what the minimum still lacks.
func readSupplies(supplies map[int]int) suppliesReading {
	r := suppliesReading{Stacks: []SupplyStack{}, Missing: []string{}}
	for id, count := range supplies {
		if s, ok := raidSupplies[id]; ok && count > 0 {
			r.Stacks = append(r.Stacks, SupplyStack{ID: id, Name: s.Name, Kind: s.Kind, Count: count})
		}
	}
	sort.Slice(r.Stacks, func(i, j int) bool {
		a, b := r.Stacks[i], r.Stacks[j]
		if a.Kind != b.Kind {
			return supplyOrder[a.Kind] < supplyOrder[b.Kind]
		}
		return a.ID < b.ID
	})
	if !hasKind(r.Stacks, supplyFlask) && !hasKind(r.Stacks, supplyElixir) {
		r.Missing = append(r.Missing, missingFlaskOrElixir)
	}
	if !hasKind(r.Stacks, supplyFood) {
		r.Missing = append(r.Missing, missingFood)
	}
	return r
}

// supplyOrder is the order stacks list in.
var supplyOrder = map[string]int{supplyFlask: 0, supplyElixir: 1, supplyFood: 2}

func hasKind(stacks []SupplyStack, kind string) bool {
	for _, s := range stacks {
		if s.Kind == kind {
			return true
		}
	}
	return false
}

// suppliesOf is the row's reading, ok false when it cannot be read: the bags are not shared
// or the export predates supplies=.
func suppliesOf(row RosterRow) (suppliesReading, bool) {
	if row.Consent != "gear_bags" || !row.prep.hasSupplies {
		return suppliesReading{}, false
	}
	return readSupplies(row.prep.supplies), true
}

// resistancesOf is the row's resistance totals, nil without gear consent or a school= section.
func resistancesOf(row RosterRow) *fs1.Resistances {
	if !hasGearConsent(row.Consent) {
		return nil
	}
	return row.prep.resistances
}

// Attunement readings: the checklist states, with why one is not tracked.
type attunementReading struct {
	State  string
	Reason string
}

// attunementOf is the row's attunement for zone: its key item (KeyItem) or its quest chain
// (attunementFor). The zero reading (State "") means the zone needs neither.
func attunementOf(zone Zone, row RosterRow) attunementReading {
	if zone.KeyItem != nil {
		return keyAttunement(*zone.KeyItem, row)
	}
	if a, ok := attunementFor(zone); ok {
		return questAttunement(a, row)
	}
	return attunementReading{}
}

// keyAttunement is ok when the key is carried, banked or worn. Missing is a fail only once
// the export lists the keys carried (locks=) and the bags are shared; before that a key not
// seen proves nothing, since bags= lists only gear.
func keyAttunement(key KeyItem, row RosterRow) attunementReading {
	if carries(row, key.ID) {
		return attunementReading{State: CheckOK}
	}
	if row.Consent == "gear_bags" && row.prep.hasKeyring {
		return attunementReading{State: CheckFail}
	}
	return attunementReading{State: CheckNotTracked, Reason: reasonKeysNotInExport}
}

// questAttunement is ok when the client flags any of the chain's quests done or the chain's
// item is carried or worn; a fail only when the export asked about every quest of the chain
// and none is done.
func questAttunement(a Attunement, row RosterRow) attunementReading {
	if !hasGearConsent(row.Consent) {
		return attunementReading{State: CheckNotTracked, Reason: reasonGearNotShared}
	}
	asked := row.prep.hasQuests
	for _, id := range a.Quests {
		done, ok := row.prep.quests[id]
		if done {
			return attunementReading{State: CheckOK}
		}
		asked = asked && ok
	}
	switch {
	case carries(row, a.Item.ID):
		return attunementReading{State: CheckOK}
	case asked:
		return attunementReading{State: CheckFail}
	}
	return attunementReading{State: CheckNotTracked, Reason: reasonQuestsNotInExport}
}

// carries reports whether the item is worn (gear consent) or in the bags, the bank or among
// the keys carried (bags consent).
func carries(row RosterRow, itemID int) bool {
	if hasGearConsent(row.Consent) {
		for _, id := range row.gear {
			if id == itemID {
				return true
			}
		}
	}
	if row.Consent != "gear_bags" {
		return false
	}
	for _, list := range [][]int{row.bags, row.prep.bank, row.prep.keyring} {
		for _, id := range list {
			if id == itemID {
				return true
			}
		}
	}
	return false
}

// attunedFlag is an attunement reading as the Readiness row's attuned: true or false when
// read, nil when not tracked or the zone needs none.
func attunedFlag(r attunementReading) *bool {
	switch r.State {
	case CheckOK, CheckFail:
		v := r.State == CheckOK
		return &v
	}
	return nil
}
