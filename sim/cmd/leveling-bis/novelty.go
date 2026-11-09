package main

// novelty.go keeps fishing catches and novelty items off the BiS lists
// (player feedback 2026-11-09: a level-40 Protection list picked Rockhide
// Strongfish, a 0.1% catch from Stranglethorn's tastyfish pools, as its
// main hand).

// fishingSourceID is loot.json's one fishing bucket: every catch
// classic-db's fishing tables and fishing pools name
// (pipeline.loot.classicdb). A catch is luck at a rod, not a place a
// player goes for gear, so it is never a source the ranker plans around.
const fishingSourceID = "world:fishing"

// isPlannableSource reports whether a loot.json source can source a pick.
func isPlannableSource(src lootSource) bool {
	return src.ID != fishingSourceID
}

// noveltyItems are items that are real and obtainable but are fishing
// novelties no player wears as gear. A small cited deny-list: items.json
// carries no "requires Fishing" flag to write a rule on (the client's
// ItemSparse.RequiredSkill 356, Fishing, is not carried through).
var noveltyItems = map[int]string{
	// Requires Fishing (1) to equip. The reward for turning in Keefer's
	// Angelfish (19805), a rare catch from the Stranglethorn Fishing
	// Extravaganza pools (quest "Rare Fish - Keefer's Angelfish").
	19972: "Lucky Fishing Hat: a fishing-tournament novelty that requires Fishing to equip",
}

// isNoveltyItem reports whether id is on the novelty deny-list.
func isNoveltyItem(id int) bool {
	_, denied := noveltyItems[id]
	return denied
}
