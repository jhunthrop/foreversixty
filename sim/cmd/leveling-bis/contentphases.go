package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// contentphases.go applies data/curated/content-phases.json - which
// content is open when, keyed by game id - to the loot index the ranker
// picks from. The pipeline (data/pipeline/loot/content_phases.py) writes the
// same gates into loot.json; applying them here too means a gate holds for
// any loot.json, including one generated before the gate was added, and
// that Python and Go read one statement of the rules rather than two.
//
// The key vocabulary, and the house rule behind the gates, are in the
// curated file's own notes.

const (
	contentPhasesFile = "content-phases.json"
	phasesFile        = "phases.json"
	// launchPhase is what every gate is measured against: a gate whose
	// phase is at or before it is open.
	launchPhase = "launch"
	// opensLater is the sentinel for a phase with no announced date.
	opensLater = "later"
)

// standingOrder is the reputation standings, lowest first, for a faction
// gate's from_standing.
var standingOrder = []string{"hated", "hostile", "unfriendly", "neutral", "friendly", "honored", "revered", "exalted"}

type contentGateID struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	FromStanding string `json:"from_standing,omitempty"`
}

type contentGate struct {
	Key           string          `json:"key"`
	Label         string          `json:"label"`
	OriginalPatch string          `json:"original_patch"`
	Phase         string          `json:"phase"`
	Date          string          `json:"date"`
	Confidence    string          `json:"confidence"`
	Why           string          `json:"why"`
	Npcs          []contentGateID `json:"npcs"`
	Quests        []contentGateID `json:"quests"`
	Factions      []contentGateID `json:"factions"`
	Zones         []contentGateID `json:"zones"`
	Kinds         []string        `json:"kinds"`
	Items         []contentGateID `json:"items"`
}

type contentPhasesDocument struct {
	Sources []json.RawMessage `json:"sources"`
	Notes   string            `json:"notes"`
	Gates   []contentGate     `json:"gates"`
}

// gateRef is what one listed id gates to.
type gateRef struct {
	opens        string
	key          string
	fromStanding string
}

// listedID is one id a closed gate lists, for reporting the ones a
// build's loot.json does not name.
type listedID struct {
	keyType string
	id      int
	kind    string
	name    string
	gate    string
}

// contentGates is every CLOSED gate, by key type.
type contentGates struct {
	rank       map[string]int
	npcs       map[int]gateRef
	npcNames   map[string]gateRef
	quests     map[int]gateRef
	factions   map[int]gateRef
	zones      map[int]gateRef
	items      map[int]gateRef
	kinds      map[string]gateRef
	gateClosed map[string]bool
	listed     []listedID
}

func loadContentGates(repoRoot string) (contentGates, error) {
	dir := filepath.Join(repoRoot, "data", "curated")
	file, err := os.ReadFile(filepath.Join(dir, contentPhasesFile))
	if err != nil {
		return contentGates{}, fmt.Errorf("reading %s: %w", contentPhasesFile, err)
	}
	phases, err := os.ReadFile(filepath.Join(dir, phasesFile))
	if err != nil {
		return contentGates{}, fmt.Errorf("reading %s: %w", phasesFile, err)
	}
	g, err := parseContentGates(file, phases)
	if err != nil {
		return contentGates{}, fmt.Errorf("%s: %w", contentPhasesFile, err)
	}
	return g, nil
}

func phaseRanks(phases []byte) (map[string]int, error) {
	var rows []struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(phases, &rows); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", phasesFile, err)
	}
	rank := make(map[string]int, len(rows)+1)
	for i, row := range rows {
		rank[row.Name] = i
	}
	if _, ok := rank[launchPhase]; !ok {
		return nil, fmt.Errorf("%s names no %q phase to gate against", phasesFile, launchPhase)
	}
	rank[opensLater] = len(rows)
	return rank, nil
}

func parseContentGates(file, phases []byte) (contentGates, error) {
	rank, err := phaseRanks(phases)
	if err != nil {
		return contentGates{}, err
	}
	dec := json.NewDecoder(bytes.NewReader(file))
	dec.DisallowUnknownFields()
	var doc contentPhasesDocument
	if err := dec.Decode(&doc); err != nil {
		return contentGates{}, fmt.Errorf("decoding: %w", err)
	}
	g := contentGates{
		rank: rank, npcs: map[int]gateRef{}, npcNames: map[string]gateRef{}, quests: map[int]gateRef{},
		factions: map[int]gateRef{}, zones: map[int]gateRef{}, items: map[int]gateRef{}, kinds: map[string]gateRef{},
		gateClosed: map[string]bool{},
	}
	seen := map[string]string{}
	for _, gate := range doc.Gates {
		if err := g.addGate(gate, seen); err != nil {
			return contentGates{}, err
		}
	}
	return g, nil
}

func (g *contentGates) addGate(gate contentGate, seen map[string]string) error {
	if _, dup := g.gateClosed[gate.Key]; dup {
		return fmt.Errorf("gate %q is named twice", gate.Key)
	}
	phaseRank, ok := g.rank[gate.Phase]
	if !ok {
		return fmt.Errorf("gate %q names unknown phase %q", gate.Key, gate.Phase)
	}
	closed := phaseRank > g.rank[launchPhase]
	g.gateClosed[gate.Key] = closed
	lists := map[string][]contentGateID{"npc": gate.Npcs, "quest": gate.Quests, "faction": gate.Factions, "zone": gate.Zones, "item": gate.Items}
	count := len(gate.Kinds)
	for keyType, ids := range lists {
		count += len(ids)
		for _, entry := range ids {
			if err := claim(seen, fmt.Sprintf("%s %d", keyType, entry.ID), gate.Key); err != nil {
				return err
			}
			if entry.FromStanding != "" && (keyType != "faction" || indexOf(standingOrder, entry.FromStanding) < 0) {
				return fmt.Errorf("gate %q %s %d: bad from_standing %q", gate.Key, keyType, entry.ID, entry.FromStanding)
			}
		}
	}
	for _, kind := range gate.Kinds {
		if err := claim(seen, "kind "+kind, gate.Key); err != nil {
			return err
		}
	}
	if count == 0 {
		return fmt.Errorf("gate %q names no id", gate.Key)
	}
	if !closed {
		return nil
	}
	ref := gateRef{opens: gate.Phase, key: gate.Key}
	for _, keyType := range []string{"npc", "quest", "faction", "zone", "kind", "item"} {
		if keyType == "kind" {
			for _, kind := range gate.Kinds {
				g.kinds[kind] = ref
				g.listed = append(g.listed, listedID{keyType: keyType, kind: kind, gate: gate.Key})
			}
			continue
		}
		for _, entry := range lists[keyType] {
			g.index(keyType, entry, ref)
			g.listed = append(g.listed, listedID{keyType: keyType, id: entry.ID, name: entry.Name, gate: gate.Key})
		}
	}
	return nil
}

func (g *contentGates) index(keyType string, entry contentGateID, ref gateRef) {
	switch keyType {
	case "npc":
		g.npcs[entry.ID] = ref
		g.npcNames[entry.Name] = ref
	case "quest":
		g.quests[entry.ID] = ref
	case "faction":
		ref.fromStanding = entry.FromStanding
		g.factions[entry.ID] = ref
	case "zone":
		g.zones[entry.ID] = ref
	case "item":
		g.items[entry.ID] = ref
	}
}

// claim records that key belongs to gate, failing when another gate
// already lists it: one id in two gates would leave its phase to order.
func claim(seen map[string]string, key, gate string) error {
	if other, ok := seen[key]; ok && other != gate {
		return fmt.Errorf("%s is listed in both %q and %q", key, other, gate)
	}
	seen[key] = gate
	return nil
}

func indexOf(values []string, v string) int {
	for i, each := range values {
		if each == v {
			return i
		}
	}
	return -1
}

func (g contentGates) rankOf(opens string) int {
	if opens == "" {
		return -1
	}
	if r, ok := g.rank[opens]; ok {
		return r
	}
	return g.rank[opensLater]
}

func (g contentGates) strictest(values ...string) string {
	best := ""
	for _, v := range values {
		if g.rankOf(v) > g.rankOf(best) {
			best = v
		}
	}
	return best
}

func standingGated(ref gateRef, standing string) bool {
	if ref.fromStanding == "" || standing == "" {
		return true
	}
	return indexOf(standingOrder, standing) >= indexOf(standingOrder, ref.fromStanding)
}

// matchedIDs records which listed ids a build's loot index names.
type matchedIDs map[string]bool

func (m matchedIDs) add(keyType string, id any) { m[fmt.Sprintf("%s %v", keyType, id)] = true }

// gatesFor is every gate one source of item falls under. repFaction and
// repStanding are the item's first reputation source, the faction a quest
// reward that states none of its own needs (Earthstrike).
func (g contentGates) gatesFor(item int, s itemSource, repFaction int, repStanding string, m matchedIDs) []string {
	var out []string
	hit := func(ref gateRef, ok bool, keyType string, id any) {
		if ok {
			m.add(keyType, id)
			out = append(out, ref.opens)
		}
	}
	if s.NpcID != 0 && (s.Kind == "vendor" || s.Kind == "dungeon" || s.Kind == "raid") {
		ref, ok := g.npcs[s.NpcID]
		hit(ref, ok, "npc", s.NpcID)
	}
	if s.Kind == "world" {
		ref, ok := g.npcNames[s.SourceName]
		hit(ref, ok, "npc-name", s.SourceName)
	}
	if s.ZoneID != 0 {
		ref, ok := g.zones[s.ZoneID]
		hit(ref, ok, "zone", s.ZoneID)
	}
	ref, ok := g.kinds[s.Kind]
	hit(ref, ok, "kind", s.Kind)
	if s.Kind != "dungeon" && s.Kind != "raid" {
		// An item gate covers flat sources and quests, as the pipeline's
		// does; a boss drop is gated by its npc instead.
		ref, ok := g.items[item]
		hit(ref, ok, "item", item)
	}
	faction, standing := 0, ""
	switch s.Kind {
	case "rep":
		faction, standing = s.FactionID, s.Standing
	case "quest":
		ref, ok := g.quests[s.QuestID]
		hit(ref, ok, "quest", s.QuestID)
		faction, standing = s.RepFaction, s.RepStanding
		if faction == 0 {
			faction, standing = repFaction, repStanding
		}
	}
	if ref, ok := g.factions[faction]; ok && faction != 0 {
		m.add("faction", faction)
		if standingGated(ref, standing) {
			out = append(out, ref.opens)
		}
	}
	return out
}

// apply returns a copy of idx with every closed gate's phase on the
// sources it covers, and the listed ids idx does not name. A missing id
// is only reported: the gate stays and closes the source the night it
// returns.
func (g contentGates) apply(idx lootIndex) (lootIndex, []string) {
	m := matchedIDs{}
	out := make(lootIndex, len(idx))
	for item, srcs := range idx {
		repFaction, repStanding := 0, ""
		for _, s := range srcs {
			if s.Kind == "rep" {
				repFaction, repStanding = s.FactionID, s.Standing
				break
			}
		}
		gated := make([]itemSource, len(srcs))
		for i, s := range srcs {
			s.Opens = g.strictest(append([]string{s.Opens}, g.gatesFor(item, s, repFaction, repStanding, m)...)...)
			gated[i] = s
		}
		out[item] = g.closeDuplicateQuests(gated)
	}
	return out, g.missing(m)
}

// closeDuplicateQuests gives every copy of one quest the strictest phase
// any copy carries. One ungated way to an item is enough for the ranker,
// so a bare copy used to defeat the gate its twin carried (Treasure of
// the Timeless One).
func (g contentGates) closeDuplicateQuests(srcs []itemSource) []itemSource {
	strictest := map[int]string{}
	for _, s := range srcs {
		if s.Kind == "quest" && s.QuestID != 0 {
			strictest[s.QuestID] = g.strictest(strictest[s.QuestID], s.Opens)
		}
	}
	for i, s := range srcs {
		if s.Kind == "quest" && s.QuestID != 0 {
			srcs[i].Opens = strictest[s.QuestID]
		}
	}
	return srcs
}

func (g contentGates) missing(m matchedIDs) []string {
	var out []string
	for _, l := range g.listed {
		var found bool
		var label string
		switch l.keyType {
		case "kind":
			found, label = m["kind "+l.kind], fmt.Sprintf("%s: kind %s", l.gate, l.kind)
		case "npc":
			found = m[fmt.Sprintf("npc %d", l.id)] || m["npc-name "+l.name]
			label = fmt.Sprintf("%s: npc %d %s", l.gate, l.id, l.name)
		default:
			found = m[fmt.Sprintf("%s %d", l.keyType, l.id)]
			label = fmt.Sprintf("%s: %s %d %s", l.gate, l.keyType, l.id, l.name)
		}
		if !found {
			out = append(out, label)
		}
	}
	return out
}

// loadGatedLootIndex is loadLootIndex with data/curated/content-phases.json
// applied: the index the ranker picks from. Any listed id the build's
// loot.json does not name is written to warn, and its gate is kept.
func loadGatedLootIndex(repoRoot, buildDir string, itemFactionRestriction map[int]string, warn io.Writer) (lootIndex, map[int]int, error) {
	idx, questFloors, err := loadLootIndex(buildDir, itemFactionRestriction)
	if err != nil {
		return nil, nil, err
	}
	gates, err := loadContentGates(repoRoot)
	if err != nil {
		return nil, nil, err
	}
	gated, missing := gates.apply(idx)
	if len(missing) > 0 {
		if _, err := fmt.Fprintf(warn, "warning: content-phases: %d gated ids are not in %s's loot.json; their gates are kept: %s\n",
			len(missing), filepath.Base(buildDir), strings.Join(missing, "; ")); err != nil {
			return nil, nil, fmt.Errorf("writing the content-phases warning: %w", err)
		}
	}
	return gated, questFloors, nil
}
