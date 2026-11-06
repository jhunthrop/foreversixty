package main

// The BiS page's source filters (owner MUST, 2026-11-05): Dungeon drops,
// Raid drops, World drops / BoE, Quest rewards, Crafted, Reputation,
// Vendor, PvP. A filter kind is loot.json's own source-kind word wherever
// one fits (web/src/lib/sim/loot.ts LOOT_KINDS), so the page's filter row,
// the Droptimizer's kind row and the row badges share one vocabulary; the
// three open-world kinds fold into "world". web/src/lib/bis/source-filter.ts
// mirrors filterKindOfSource and filterKindOrder, and both carry a test.
//
// Multi-source rule: an item belongs to EVERY filter kind it can be
// obtained from at the band (obtainableSources' own gates: faction, level,
// standing, content phase, class-restricted quests). A quest reward that a
// dungeon boss also drops stays listed while either "quest" or "dungeon"
// is on; its source label names the source of a kind that is on.
//
// World drops / BoE covers loot.json's world (a named open-world mob),
// world_drop (the level-banded world pool) and zone (a farm-anywhere zone
// drop). The item data carries no bind type, so a BoE dungeon drop is not
// counted under it.

const (
	filterKindDungeon = "dungeon"
	filterKindRaid    = "raid"
	filterKindWorld   = "world"
	filterKindQuest   = "quest"
	filterKindCrafted = "crafted"
	filterKindRep     = "rep"
	filterKindVendor  = "vendor"
	filterKindPvP     = "pvp"
)

// filterKindOrder is the filter row's order, and the order by_kind keys
// are written in (encoding/json sorts map keys anyway; this is for any
// code that walks the kinds).
var filterKindOrder = []string{
	filterKindDungeon, filterKindRaid, filterKindWorld, filterKindQuest,
	filterKindCrafted, filterKindRep, filterKindVendor, filterKindPvP,
}

// sourceKindToFilterKind is the one mapping from a loot.json / ranker
// source kind (itemSource.Kind) to its filter kind.
var sourceKindToFilterKind = map[string]string{
	"dungeon":    filterKindDungeon,
	"raid":       filterKindRaid,
	"world":      filterKindWorld,
	"world_drop": filterKindWorld,
	"zone":       filterKindWorld,
	"quest":      filterKindQuest,
	"crafted":    filterKindCrafted,
	"rep":        filterKindRep,
	"vendor":     filterKindVendor,
	"pvp":        filterKindPvP,
}

// filterKindOfSource is a source kind's filter kind; false for a kind no
// filter covers (none in this build's loot.json).
func filterKindOfSource(sourceKind string) (string, bool) {
	k, ok := sourceKindToFilterKind[sourceKind]
	return k, ok
}

// sourcesByFilterKind groups an item's obtainable sources by filter kind
// and keeps, per kind, the one preferredSource names - the same choice
// sourceFor makes among all of them. A class-restricted quest the class
// cannot take is no source of it. Only sources sourceFor could ever name
// count (pickableKind): a zone source alone never makes an item a pick.
func sourcesByFilterKind(obtainable []itemSource, classSlug string) map[string]itemSource {
	grouped := make(map[string][]itemSource)
	for _, s := range obtainable {
		kind, ok := filterKindOfSource(s.Kind)
		if !ok || !pickableKind(s.Kind) || !classAllowed(s.Classes, classSlug) {
			continue
		}
		grouped[kind] = append(grouped[kind], s)
	}
	out := make(map[string]itemSource, len(grouped))
	for kind, srcs := range grouped {
		if s, ok := preferredSource(srcs); ok {
			out[kind] = s
		}
	}
	return out
}

// filterKindSources is a candidate's sources by filter kind: KindSources,
// else its one published Source (a candidate built outside buildBandPool).
func filterKindSources(c scored) map[string]itemSource {
	if len(c.KindSources) > 0 {
		return c.KindSources
	}
	if !c.HasSource {
		return nil
	}
	kind, ok := filterKindOfSource(c.Source.Kind)
	if !ok {
		return nil
	}
	return map[string]itemSource{kind: c.Source}
}

// orderedFilterKinds is the filter kinds of a candidate, in filterKindOrder.
func orderedFilterKinds(sources map[string]itemSource) []string {
	out := make([]string, 0, len(sources))
	for _, k := range filterKindOrder {
		if _, ok := sources[k]; ok {
			out = append(out, k)
		}
	}
	return out
}
