package main

// An item's own reputation requirement.
//
// loot.json gates a reputation reward through its `rep` source's standing,
// and a vendor row inherits it (vendorInheritsRepStandingGate). A vendor
// whose stock has no `rep` source states neither a standing nor a side:
// Creeg Bothunk, the Darkspear Islands quartermaster, is the one that
// matters at launch. The client states both on the item itself, in
// ItemSparse.MinFactionID and MinReputation, which the pipeline carries
// into items/<class>.json as `reputation`. Every source of such an item is
// held to that standing (repStandingObtainable: friendly and honored while
// leveling, revered and exalted at 60) and, for a faction-side reputation,
// to that side.

// itemReputation is items/<class>.json's `reputation`: the reputation
// faction id and the lowest standing the item needs.
type itemReputation struct {
	FactionID int    `json:"faction_id"`
	Standing  string `json:"standing"`
}

// standingRank orders the standings a reward can need.
var standingRank = map[string]int{"neutral": 0, "friendly": 1, "honored": 2, "revered": 3, "exalted": 4}

// repFactionSide is the side a faction-side reputation belongs to: the
// client's Faction table puts these under Alliance Forces (891) or Horde
// Forces (892). Any other reputation (Timbermaw Hold, Argent Dawn, the
// Forever zone factions) is open to both sides.
var repFactionSide = map[int]string{
	509:  "alliance", // The League of Arathor
	510:  "horde",    // The Defilers
	729:  "horde",    // Frostwolf Clan
	730:  "alliance", // Stormpike Guard
	889:  "alliance", // Silverwing Sentinels
	890:  "horde",    // Warsong Outriders
	2798: "horde",    // Darkspear Raiders
	2799: "alliance", // Theramore Expeditionary Force
}

// stricterStanding is the higher of two standings; "" is no requirement.
func stricterStanding(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" || standingRank[a] >= standingRank[b] {
		return a
	}
	return b
}

// applyItemReputations returns idx with every source of a reputation item
// held to the item's standing and side. idx itself is left unchanged.
func applyItemReputations(idx lootIndex, items []candidate) lootIndex {
	reps := make(map[int]itemReputation)
	for _, c := range items {
		// Neutral is where every character starts, so it gates nothing.
		if c.Reputation != nil && standingRank[c.Reputation.Standing] > standingRank["neutral"] {
			reps[c.ID] = *c.Reputation
		}
	}
	out := make(lootIndex, len(idx))
	for id, srcs := range idx {
		rep, ok := reps[id]
		if !ok {
			out[id] = srcs
			continue
		}
		gated := make([]itemSource, len(srcs))
		for i, s := range srcs {
			s.Standing = stricterStanding(s.Standing, rep.Standing)
			if s.Side == "" {
				s.Side = repFactionSide[rep.FactionID]
			}
			gated[i] = s
		}
		out[id] = gated
	}
	return out
}
