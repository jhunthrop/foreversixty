package request

// Small, test-independent helpers for the rotation ladder
// (ladder_test.go): talent truncation, gear picking, cast tallying, and
// golden rendering. Nothing here reaches into rotations_smoke_test.go's
// _test.go-only symbols (repoRoot, activeBuild, readBuildJSON, ...) -
// every function below takes the paths and build strings it needs as
// plain arguments, so this file compiles into the ordinary package as
// well as the test binary.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

// ladderLevels is Phase 1a's ladder: the seven character levels every
// written rotation is measured at.
var ladderLevels = []int{10, 20, 30, 38, 40, 50, 60}

// ladderIterations and ladderSeed are the ladder's fixed run: enough
// iterations that a rarely-reached line still gets its turn, one seed so
// a failure reproduces.
const (
	ladderIterations = 300
	ladderSeed       = 1
)

// ---------------------------------------------------------------------
// Talents: truncating a guide's level-60 build to a lower level.
//
// Moved to sim/leveling (lane bis-all) so the leveling BiS pipeline
// (sim/cmd/leveling-bis) can share the exact same rule instead of
// approximating its own. talentNode/talentTree/talentsFile/
// loadTalentTrees/guideBuildTalents/guideTalentTargets/
// ladderTalentString now live there as leveling.TalentNode/
// leveling.TalentTree/leveling.TalentsFile/leveling.LoadTalentTrees/
// leveling.GuideBuildTalents/leveling.GuideTalentTargets/
// leveling.LadderTalentString, with identical behaviour; only the call
// sites below were renamed to match.
// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Gear: a weapon set the engine's item database knows, chosen by level.
// ---------------------------------------------------------------------

// buildItem is the fields ladder gear picking needs from one row of
// data/builds/<build>/items/<class>.json.
type buildItem struct {
	ID            int     `json:"id"`
	Slot          string  `json:"slot"`
	RequiredLevel int     `json:"required_level"`
	ItemLevel     int     `json:"item_level"`
	TwoHand       bool    `json:"two_hand"`
	Speed         float64 `json:"speed"`
	DamageMax     int     `json:"damage_max"`
}

type classItemsFile struct {
	Items []buildItem `json:"items"`
}

// loadClassItems reads one class's item rows - already restricted to
// items that class can equip, which is the "class allow" the design
// asks for.
func loadClassItems(repoRoot, build, class string) ([]buildItem, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "items", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f classItemsFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	return f.Items, nil
}

type simItemsFile struct {
	Items []int `json:"items"`
}

// loadSimItemIDs reads simitems.json: the item ids the pinned engine's
// own item database carries. items/<class>.json is already a subset of
// this set on every build checked, but a gear pick still filters against
// it rather than trusting that silently, so a future build where a
// class item outruns the sim database fails to pick the item instead of
// handing the engine an id it panics on (see sim/adapter/golden_test.go).
func loadSimItemIDs(repoRoot, build string) (map[int]bool, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "simitems.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f simItemsFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := make(map[int]bool, len(f.Items))
	for _, id := range f.Items {
		out[id] = true
	}
	return out, nil
}

// handedness constrains a weapon pick to one-handed, two-handed, or
// either.
type handedness int

const (
	handAny handedness = iota
	handOne
	handTwo
)

// gearProfile is which weapon slots a spec's ladder character fills, and
// what handedness main_hand takes. The zero value is a caster: main hand
// only, any handedness (a caster's main_hand may be a one-hand dagger or
// a two-hand staff and either is fine, since no off-hand is ever filled
// alongside it), no off hand, no ranged - which is also why casters have
// no entry below.
type gearProfile struct {
	MainHand handedness
	OffHand  bool
	Ranged   bool
	// Skip is druid-feral: "none" in the design's own words. A feral
	// character fights shapeshifted and this ladder does not model
	// weapon-DPS-through-form-conversion, so it picks no weapon at all
	// rather than equip one that does nothing for the rotation measured.
	Skip bool
}

// ladderGearProfiles is every written spec whose weapon rule is not the
// caster default: the two-handers (warrior-arms, paladin-retribution),
// the dual-wielders (warrior-fury, rogue's three specs,
// shaman-enhancement, and hunter - hunters can dual-wield a melee
// stat-stick beside their bow in this build), hunter's ranged weapon,
// and feral's "none".
var ladderGearProfiles = map[string]gearProfile{
	"warrior-arms":         {MainHand: handTwo},
	"warrior-fury":         {MainHand: handOne, OffHand: true},
	"paladin-retribution":  {MainHand: handTwo},
	"shaman-enhancement":   {MainHand: handOne, OffHand: true},
	"druid-feral":          {Skip: true},
	"rogue-assassination":  {MainHand: handOne, OffHand: true},
	"rogue-combat":         {MainHand: handOne, OffHand: true},
	"rogue-subtlety":       {MainHand: handOne, OffHand: true},
	"hunter-beast-mastery": {MainHand: handOne, OffHand: true, Ranged: true},
	"hunter-marksmanship":  {MainHand: handOne, OffHand: true, Ranged: true},
	"hunter-survival":      {MainHand: handOne, OffHand: true, Ranged: true},
}

// pickGearItem is the highest item_level candidate in slot with
// required_level <= level, restricted to hand's handedness and to items
// this build's sim database knows (known). A candidate must be a real
// weapon: speed > 0 (which is what keeps a shield out of a
// dual-wielder's off hand: an off_hand row with speed == 0 is armor,
// not a weapon, in this item table) and damage_max > 0. The second
// check exists because this build's item table carries a large number
// of quality-3 "rare" weapon rows with damage_min = damage_max = 0 and
// no required_level (Bland Dagger, item 24071, is one) - unfinished or
// placeholder rows, not a character's real choice, and equipping one
// of them is what made the engine hang mid-sim rather than simulate a
// zero-damage weapon (see the report). Ties (equal item_level) break on
// the lower item id, so the pick is deterministic without depending on
// the source file's row order.
func pickGearItem(items []buildItem, known map[int]bool, slot string, level int, hand handedness) (buildItem, bool) {
	var best buildItem
	found := false
	for _, it := range items {
		if it.Slot != slot || it.Speed <= 0 || it.DamageMax <= 0 || !known[it.ID] {
			continue
		}
		if it.RequiredLevel > level {
			continue
		}
		if hand == handOne && it.TwoHand {
			continue
		}
		if hand == handTwo && !it.TwoHand {
			continue
		}
		if !found || it.ItemLevel > best.ItemLevel || (it.ItemLevel == best.ItemLevel && it.ID < best.ID) {
			best, found = it, true
		}
	}
	return best, found
}

// ladderGear is a level-appropriate weapon set for spec: main_hand,
// and off_hand/ranged where the spec's profile calls for them. Every
// other slot is bare, as the design's Phase 1a asks ("a bare character
// shows the rotation, not the raid").
func ladderGear(items []buildItem, known map[int]bool, spec string, level int) []api.GearSlot {
	profile := ladderGearProfiles[spec]
	if profile.Skip {
		return nil
	}
	var gear []api.GearSlot
	if it, ok := pickGearItem(items, known, "main_hand", level, profile.MainHand); ok {
		gear = append(gear, api.GearSlot{Slot: "main_hand", ItemID: it.ID})
	}
	if profile.OffHand {
		if it, ok := pickGearItem(items, known, "off_hand", level, handOne); ok {
			gear = append(gear, api.GearSlot{Slot: "off_hand", ItemID: it.ID})
		}
	}
	if profile.Ranged {
		if it, ok := pickGearItem(items, known, "ranged", level, handAny); ok {
			gear = append(gear, api.GearSlot{Slot: "ranged", ItemID: it.ID})
		}
	}
	return gear
}

// formatGear renders a gear set for the golden: "bare" for none,
// otherwise each slot's item id, sorted so the line never reorders
// itself between regenerations.
func formatGear(gear []api.GearSlot) string {
	if len(gear) == 0 {
		return "bare"
	}
	parts := make([]string, len(gear))
	for i, g := range gear {
		parts[i] = fmt.Sprintf("%s:%d", g.Slot, g.ItemID)
	}
	sort.Strings(parts)
	return strings.Join(parts, " ")
}

// ---------------------------------------------------------------------
// Spell data: which abilities exist, when they are learned, and which
// of them deal damage.
// ---------------------------------------------------------------------

// spellRankEntry is one rank of one named ability, from
// data/builds/<build>/spellranks.json.
type spellRankEntry struct {
	ID    int `json:"id"`
	Level int `json:"level"`
	Rank  int `json:"rank"`
}

type spellRanksFile struct {
	Classes map[string]map[string][]spellRankEntry `json:"classes"`
}

// loadSpellRanks reads the whole spellranks.json once; callers index
// f.Classes[class] themselves so the (large) file is read only once per
// ladder run rather than once per spec.
func loadSpellRanks(repoRoot, build string) (spellRanksFile, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "spellranks.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return spellRanksFile{}, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f spellRanksFile
	if err := json.Unmarshal(b, &f); err != nil {
		return spellRanksFile{}, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	return f, nil
}

// junkAbilityName matches spellranks.json keys that are not a named
// ability at all: item-set-bonus rows key on a version-stamped label
// like "1.60.0 - Item - Tier 1 - Warrior - Arms/Fury 2P Bonus - Haste".
var junkAbilityName = regexp.MustCompile(`^\d`)

// rankTier is every id that shares one rank number of one ability, and
// the level that rank is learned at - the ladder's own mirror of
// sim/internal/spellranks' buildRankChain (unexported there, so this is
// a second, read-only copy of the same grouping rather than a call into
// it), used only for the informational "learned but unused" list
// (rule 3). A tier's level is the LEVEL OF THE FIRST ROW SEEN at that
// rank, matching that package exactly: a later duplicate at the same
// rank number and a different level (typically 0, an NPC/internal
// copy) widens the tier's id set without ever lowering its level.
type rankTier struct {
	Level int
	IDs   []int
}

// classAbilities is one class's spellranks.json rows grouped by ability
// name into rank tiers. A name whose rows are ALL rank 0 has NO entry
// here at all - exactly like sim/internal/spellranks' table, which is
// how Bloodrage (two rank-0 rows, no real progression) and Judgement
// (three rank-0 rows) are represented: they are not "ranked spells"
// this ladder - or the engine's own rewrite - tracks at all.
type classAbilities struct {
	Tiers map[string][]rankTier
}

func buildClassAbilities(all spellRanksFile, class string) classAbilities {
	out := classAbilities{Tiers: map[string][]rankTier{}}
	for name, entries := range all.Classes[class] {
		if junkAbilityName.MatchString(name) {
			continue
		}
		levelByRank := map[int]int{}
		idsByRank := map[int][]int{}
		var ranks []int
		for _, e := range entries {
			if e.Rank <= 0 {
				continue
			}
			if _, seen := levelByRank[e.Rank]; !seen {
				levelByRank[e.Rank] = e.Level
				ranks = append(ranks, e.Rank)
			}
			idsByRank[e.Rank] = append(idsByRank[e.Rank], e.ID)
		}
		if len(ranks) == 0 {
			continue
		}
		sort.Ints(ranks)
		tiers := make([]rankTier, len(ranks))
		for i, r := range ranks {
			tiers[i] = rankTier{Level: levelByRank[r], IDs: idsByRank[r]}
		}
		out.Tiers[name] = tiers
	}
	return out
}

// learnedTierAtLevel is the highest-rank tier a character of level has
// reached, from an ability's own rank-ascending tier list (tiers are
// already sorted by rank, which in this table is also ascending by
// level in every case checked, so walking forward and keeping the last
// tier at or under level picks the same tier
// sim/internal/spellranks.highestLearnedIn would, walking from the top
// down).
func learnedTierAtLevel(tiers []rankTier, level int) (rankTier, bool) {
	best := -1
	for i, t := range tiers {
		if t.Level <= level {
			best = i
		}
	}
	if best == -1 {
		return rankTier{}, false
	}
	return tiers[best], true
}

// abilityNameByID is every tracked id's ability name, for labeling a
// zero_casts violation (rule 1 resolves the id sim/internal/spellranks'
// own HighestLearnedSpellID would cast, which carries no name). An id
// this table does not track - an unranked ability like Bloodrage, or
// one spellranks.json's generator dropped entirely - has no entry, and
// callers fall back to the bare id.
func abilityNameByID(abilities classAbilities) map[int]string {
	out := map[int]string{}
	for name, tiers := range abilities.Tiers {
		for _, tier := range tiers {
			for _, id := range tier.IDs {
				out[id] = name
			}
		}
	}
	return out
}

// spellEffectConst is one effect of one spell, from
// data/builds/<build>/spellconst/<class>.json.
type spellEffectConst struct {
	Effect int `json:"effect"`
	Aura   int `json:"aura"`
}

type spellConstEntry struct {
	Name    string             `json:"name"`
	Effects []spellEffectConst `json:"effects"`
}

type spellConstFile struct {
	Spells map[string]spellConstEntry `json:"spells"`
}

// loadSpellConst reads one class's spell constants, keyed by spell id.
func loadSpellConst(repoRoot, build, class string) (map[int]spellConstEntry, error) {
	path := filepath.Join(repoRoot, "data", "builds", build, "spellconst", class+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var f spellConstFile
	if err := json.Unmarshal(b, &f); err != nil {
		return nil, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	out := make(map[int]spellConstEntry, len(f.Spells))
	for idStr, entry := range f.Spells {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		out[id] = entry
	}
	return out, nil
}

// isDamageEffect is the design's own rule for "this effect deals
// damage": type 2 (school damage), type 6 with aura 3 (periodic
// damage), or type 121, 31 or 58 (the three weapon-damage effect types).
// It is deliberately narrow - Heroic Strike's own bonus-weapon-damage
// effect is type 17 and does not match, which the report calls out as a
// gap in the rule rather than something this code quietly widens to
// cover.
func isDamageEffect(e spellEffectConst) bool {
	if e.Effect == 2 {
		return true
	}
	if e.Effect == 6 && e.Aura == 3 {
		return true
	}
	switch e.Effect {
	case 121, 31, 58:
		return true
	}
	return false
}

// isDamageSpellID reports whether id, looked up in consts, deals damage
// by isDamageEffect. An id the table does not carry is not a damage
// spell as far as this ladder can tell.
func isDamageSpellID(consts map[int]spellConstEntry, id int) bool {
	entry, ok := consts[id]
	if !ok {
		return false
	}
	for _, e := range entry.Effects {
		if isDamageEffect(e) {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------
// Cast tally.
// ---------------------------------------------------------------------

// castSpellCounts sums every plain-spell action's casts across the run,
// keyed by the base spell id (a tag or rank variant of the same spell -
// see sim/adapter's ActionName - folds into the one id, since what this
// ladder asks is "was this ability cast at all", not which of its
// variants). Items, pets and "other" actions are not spells and are
// left out.
func castSpellCounts(player *proto.UnitMetrics) map[int]int64 {
	out := map[int]int64{}
	for _, action := range player.GetActions() {
		id := action.GetId()
		if id == nil {
			continue
		}
		spellID, ok := id.RawId.(*proto.ActionID_SpellId)
		if !ok {
			continue
		}
		var casts int64
		for _, target := range action.GetTargets() {
			casts += int64(target.GetCasts())
		}
		if casts == 0 {
			continue
		}
		out[int(spellID.SpellId)] += casts
	}
	return out
}

// castTally is one action and how many times the player performed it
// per iteration - the ladder's own version of rotations_smoke_test.go's
// castCount, normalized against THIS run's own iteration count rather
// than that file's smokeIterations constant (200; the ladder runs 300 -
// reusing castSet here would silently divide by the wrong number).
type castTally struct {
	Name         string
	PerIteration float64
}

// ladderCastSet is every action the player actually performed, most
// first, normalized by iterations. The same source rotations_smoke_test.go's
// castSet reads (player.GetActions()), computed against this run's own
// iteration count.
func ladderCastSet(player *proto.UnitMetrics, iterations int) []castTally {
	var out []castTally
	for _, action := range player.GetActions() {
		var casts int32
		for _, target := range action.GetTargets() {
			casts += target.GetCasts()
		}
		if casts == 0 {
			continue
		}
		_, name := adapter.ActionName(action.GetId())
		out = append(out, castTally{Name: name, PerIteration: float64(casts) / float64(iterations)})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].PerIteration != out[j].PerIteration {
			return out[i].PerIteration > out[j].PerIteration
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// ladderDistinctSpellCasts counts how many distinct "spell:"-named
// actions were cast at all - autoattacks and resource gains the engine
// records alongside them are not one, which is what makes this the
// design's rule-4 check ("no cast but auto-attack").
func ladderDistinctSpellCasts(tallies []castTally) int {
	n := 0
	for _, c := range tallies {
		if strings.HasPrefix(c.Name, "spell:") {
			n++
		}
	}
	return n
}

// ladderTopCasts renders the n most-performed actions for the golden.
func ladderTopCasts(tallies []castTally, n int) string {
	if n > len(tallies) {
		n = len(tallies)
	}
	parts := make([]string, n)
	for i, c := range tallies[:n] {
		parts[i] = fmt.Sprintf("%s=%.1f", c.Name, c.PerIteration)
	}
	return strings.Join(parts, ", ")
}

// extractCastSpellIDs walks a decoded APL JSON tree (any of
// map[string]any, []any, or a scalar) and collects the spell id of
// every castSpell action it finds, at any depth - a compound action
// (sequence, resetSequence, ...) can hold cast actions nested inside
// it, so this does not stop at the top level of prepullActions or
// priorityList.
func extractCastSpellIDs(node any, out map[int]bool) {
	switch v := node.(type) {
	case map[string]any:
		if cs, ok := v["castSpell"].(map[string]any); ok {
			if sid, ok := cs["spellId"].(map[string]any); ok {
				if raw, ok := sid["spellId"].(float64); ok {
					out[int(raw)] = true
				}
			}
		}
		for _, child := range v {
			extractCastSpellIDs(child, out)
		}
	case []any:
		for _, child := range v {
			extractCastSpellIDs(child, out)
		}
	}
}

// authoredCastSpellIDs is every spell id a curated rotation's own JSON
// (its "rotation" field, the same source apl/<spec>.apl.json embeds)
// names in a castSpell action - the lines the rotation lane wrote down,
// at the ranks it wrote them at, before any level rewrite.
func authoredCastSpellIDs(rotation json.RawMessage) (map[int]bool, error) {
	var tree any
	if err := json.Unmarshal(rotation, &tree); err != nil {
		return nil, fmt.Errorf("ladder: parsing the curated rotation: %w", err)
	}
	out := map[int]bool{}
	extractCastSpellIDs(tree, out)
	return out, nil
}

// ---------------------------------------------------------------------
// The curated file's ladder-relevant fields.
// ---------------------------------------------------------------------

// expectedIdleEntry is one line data/curated/apl/<spec>.json's
// expected_idle array may carry: an engine-resolvable spell id this
// ladder's bare run cannot cast, and why.
type expectedIdleEntry struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}

// ladderCurated is the half of data/curated/apl/<spec>.json the ladder
// reads: rotations_smoke_test.go's own curatedRotation reads spec/
// state/inert for its own purpose; this is a second, ladder-only
// reader carrying expected_idle beside inert (the new key this task
// adds, per the design) and the raw rotation JSON authoredCastSpellIDs
// needs.
type ladderCurated struct {
	Spec         string              `json:"spec"`
	State        string              `json:"state"`
	Inert        []int               `json:"inert"`
	ExpectedIdle []expectedIdleEntry `json:"expected_idle"`
	Rotation     json.RawMessage     `json:"rotation"`
}

func loadLadderCurated(repoRoot, spec string) (ladderCurated, error) {
	path := filepath.Join(repoRoot, "data", "curated", "apl", spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return ladderCurated{}, fmt.Errorf("ladder: reading %s: %w", path, err)
	}
	var doc ladderCurated
	if err := json.Unmarshal(b, &doc); err != nil {
		return ladderCurated{}, fmt.Errorf("ladder: parsing %s: %w", path, err)
	}
	if doc.Spec != spec {
		return ladderCurated{}, fmt.Errorf("ladder: %s declares spec %q", path, doc.Spec)
	}
	return doc, nil
}

// ---------------------------------------------------------------------
// Golden rendering.
// ---------------------------------------------------------------------

// ladderRow is one level's line in a spec's golden table.
type ladderRow struct {
	Level         int
	Talents       string
	Gear          string
	DPS           float64
	DistinctCasts int
	TopCasts      string
	Unresolved    []string
}

// unusedEntry is one (level, ability) the class had learned and could
// deal damage with, that this level's cast set never touched.
// Informational: the design's rule 3, not a failure.
type unusedEntry struct {
	Level int
	Name  string
	ID    int
}

// ladderRulesHeader documents, once, the rules every golden in this
// directory was generated under. It is repeated verbatim into each
// file rather than referenced, so a reader of one golden never has to
// go find the rule that produced it.
const ladderRulesHeader = `Rules this ladder runs under (Phase 1a,
docs/superpowers/specs/2026-09-28-rotation-accuracy-program-design.md):

- Levels: 10, 20, 30, 38, 40, 50, 60. 300 iterations, seed 1, the
  default encounter (a stationary target three levels above the
  character).
- Talents: the guide's level-60 FS1 build code
  (web/src/content/guides/<class>/<spec>.md), truncated to level-9
  points (0 below level 10). Each talent's guide-assigned rank is read
  by the talent's own stable id against the client build the guide
  names, then re-resolved to that talent's row in the ACTIVE build
  (data/builds/<active>/talents/<class>.json) - a talent the active
  build no longer carries is dropped rather than misaligning every
  digit after it. Points are spent walking the active build's trees
  top row down, the spec's own tree first, then the other two in the
  build code's own order (0, 1, 2, skipping the spec's own). A row the
  budget runs out before reaching is simply left at 0. If the guide
  build itself spends fewer than level-9 points, the remainder is left
  unspent.
- Gear: main_hand always, off_hand for the classes that dual-wield in
  this build (rogue, warrior-fury, shaman-enhancement, hunter),
  ranged for hunter only. Each slot picks the highest item_level
  weapon (speed > 0, so a shield never fills an off hand; damage_max >
  0, so an unfinished/placeholder weapon row this build's item table
  still carries - equipping one hangs the engine mid-sim rather than
  simulating a zero-damage weapon, see the report - is never picked)
  with required_level <= the character's level, from
  data/builds/<active>/items/<class>.json filtered to
  data/builds/<active>/simitems.json's known ids, honoring the spec's
  handedness (warrior-arms and paladin-retribution two-hand only;
  warrior-fury, rogue and shaman-enhancement one-hand only for both
  hands). druid-feral picks no weapon at all. Every other slot is
  bare. Buffs and consumables: none.
- DPS regression: each level's DPS is compared against the ladder's own
  PREVIOUS rung (not literally level-10, since the ladder's own gaps
  are uneven - 30 to 38 is 8 levels, 38 to 40 is 2). A level scoring
  lower than the rung before it is a violation.
- Unresolved: an id the engine's ComputeStats warns it cannot resolve.
  Expected when data/curated/apl/<spec>.json's own inert array names
  it; otherwise a violation.
- Zero casts: one of the curated rotation's own castSpell lines, resolved
  to the id sim/internal/spellranks.HighestLearnedSpellID says the
  engine's OWN rank rewrite actually casts at this level (not a second,
  approximate copy of that resolution - this ladder calls the same
  function sim/request's rewriteRotationRanks calls), that never fires
  in the run - unless the engine could not resolve that id at all
  (already counted as unresolved) or data/curated/apl/<spec>.json's
  expected_idle array names the LINE'S AUTHORED id with a reason.
  HighestLearnedSpellID returning "not learned" means the engine's own
  rewrite already dropped the line before the request was built, which
  is not a violation to report twice.
- Learned but unused (informational, not a violation): every damage
  ability (spellconst effect 2, 6 with aura 3, or 121/31/58 - see
  ladder.go's isDamageEffect) the class has learned by this level that
  this level's cast set never touched, regardless of whether the
  rotation names it at all. An ability whose spellranks.json rows are
  ALL rank 0 (Bloodrage, Judgement - a single always-known ability, not
  a rank progression) is not tracked by this rule at all, the same way
  it is invisible to the engine's own rank rewrite. Heroic
  Strike-shaped bonus-weapon-damage effects (type 17) are NOT covered
  by the damage-effect rule and so never appear here even when
  genuinely unused - see the report for specs where that matters.
- Strict failures (FOREVER_LADDER_STRICT=1) are the four rules above;
  see the run's own "Violations" section below for what this file's own
  run found.
`

// renderLadderGolden is the markdown for one spec's golden file.
func renderLadderGolden(spec string, rows []ladderRow, unused []unusedEntry, violations []string) []byte {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s rotation ladder\n\n", spec)
	b.WriteString(ladderRulesHeader)
	b.WriteString("\n## Ladder\n\n")
	b.WriteString("| Level | Talents | Gear | DPS | Distinct casts | Top casts | Unresolved |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, r := range rows {
		unresolved := "-"
		if len(r.Unresolved) > 0 {
			unresolved = strings.Join(r.Unresolved, ", ")
		}
		topCasts := r.TopCasts
		if topCasts == "" {
			topCasts = "-"
		}
		fmt.Fprintf(&b, "| %d | %s | %s | %.1f | %d | %s | %s |\n",
			r.Level, r.Talents, r.Gear, r.DPS, r.DistinctCasts, topCasts, unresolved)
	}

	b.WriteString("\n## Learned but unused (informational)\n\n")
	if len(unused) == 0 {
		b.WriteString("None at any ladder level.\n")
	} else {
		sorted := append([]unusedEntry(nil), unused...)
		sort.Slice(sorted, func(i, j int) bool {
			if sorted[i].Level != sorted[j].Level {
				return sorted[i].Level < sorted[j].Level
			}
			if sorted[i].Name != sorted[j].Name {
				return sorted[i].Name < sorted[j].Name
			}
			return sorted[i].ID < sorted[j].ID
		})
		lastLevel := -1
		for _, u := range sorted {
			if u.Level != lastLevel {
				fmt.Fprintf(&b, "\n### Level %d\n\n", u.Level)
				lastLevel = u.Level
			}
			fmt.Fprintf(&b, "- %s (spell %d)\n", u.Name, u.ID)
		}
	}

	b.WriteString("\n## Violations found in this run\n\n")
	if len(violations) == 0 {
		b.WriteString("None.\n")
	} else {
		sorted := append([]string(nil), violations...)
		sort.Strings(sorted)
		for _, v := range sorted {
			fmt.Fprintf(&b, "- %s\n", v)
		}
	}

	return []byte(b.String())
}
