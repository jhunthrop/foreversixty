// Command spec-breakdown sims one published BiS entry (the gear, race and
// talents the ranker published for a spec, band and preset) under the
// request the ranker builds for it, and prints where the damage comes from:
// damage by action with its hit table, aura uptimes and resource flow.
//
// It exists to compare two specs side by side under identical conditions
// and to measure one arrangement of consumables against another; it reads
// the published entry and layers the class kit and preset exactly as
// sim/cmd/leveling-bis does (withKit), so its total equals the ranker's
// set_dps for the same seed and iteration count.
//
//	go run ./cmd/spec-breakdown -repo-root .. -spec rogue-combat -preset raid
//	go run ./cmd/spec-breakdown -repo-root .. -spec rogue-combat -preset raid \
//	    -consumes main_hand_imbue:instant_poison,off_hand_imbue:deadly_poison
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// buildDirectory is data/builds/<active build> under the repo root: the build
// web/src/data/active-build.json names, never a build typed here.
func buildDirectory(repoRoot string) (string, error) {
	build, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		return "", err
	}
	return filepath.Join(repoRoot, "data", "builds", build), nil
}

type options struct {
	repoRoot   string
	spec       string
	preset     string
	band       int
	level      int
	faction    string
	race       string
	iterations int
	seed       int64
	consumes   string
	buffs      string
	talents    string
	// bandTalents keeps the band entry's talents at -level instead of the
	// ladder's at that level (the pre-round-24 behaviour).
	bandTalents bool
	rotation    string
	rank        string
	jsonOut     bool
	tank        bool
	heal        bool
	raidFight   bool
	duration    int
	targetType  string
	targets     int
	bisDir      string
	gear        string
}

func main() {
	var o options
	flag.StringVar(&o.repoRoot, "repo-root", "..", "site repository root")
	flag.StringVar(&o.spec, "spec", "", "spec slug, e.g. rogue-combat")
	flag.StringVar(&o.preset, "preset", "raid", "bare or raid")
	flag.IntVar(&o.band, "band", 60, "BiS band")
	flag.IntVar(&o.level, "level", 0, "character level inside the band (band to band+9; default the band), to judge a rank row at the band's top as well as its floor; on band 20 it reaches down to 10, wearing only the entry's items that level allows")
	flag.StringVar(&o.faction, "faction", "alliance", "faction of the published entry")
	flag.StringVar(&o.race, "race", "", "race slug replacing the published entry's (its gear and talents kept); the race's rank choices follow it")
	flag.IntVar(&o.iterations, "iterations", 20000, "iterations")
	flag.Int64Var(&o.seed, "seed", 1, "random seed")
	flag.StringVar(&o.consumes, "consumes", "", "comma list replacing the kit and preset consumables")
	flag.StringVar(&o.buffs, "buffs", "", "comma list added to the request's buffs")
	flag.StringVar(&o.talents, "talents", "", "talent string replacing the published one")
	flag.BoolVar(&o.bandTalents, "band-talents", false, "with -level, keep the band entry's talents instead of the ones the guide's leveling order gives at that level")
	flag.StringVar(&o.rank, "rank", "", "comma list of authoredSpellID=rank pairs forced on every reference to that spell (casts and conditions, every line), on -rotation or the spec's own curated rotation; the way to judge a rank row")
	flag.StringVar(&o.rotation, "rotation", "", "JSON file whose \"rotation\" object (the curated file's shape) replaces the spec's own priority list")
	flag.StringVar(&o.gear, "gear", "", "comma list of slot:item_id pairs replacing the published entry's items, e.g. off_hand:20688")
	flag.BoolVar(&o.heal, "heal", false, "print the healer fight: mana income and spend against data/curated/heal-profile.json (healer specs only)")
	flag.BoolVar(&o.raidFight, "raid-fight", false, "with -heal: run the raid-shaped fight (data/curated/heal-raid-profile.json, the /tiers healer ranking's) instead of heal-profile.json")
	flag.IntVar(&o.duration, "duration", 0, "fight length in seconds instead of the default (with -heal: instead of the profile's, where a shorter window shows a healer's mana rates before it runs dry)")
	flag.IntVar(&o.targets, "targets", 0, "number of targets instead of the ranker's one (a pack fight; threat /s and DPS sum over every target)")
	flag.StringVar(&o.targetType, "target-type", api.TargetTypeUnknown, "creature type of the target (beast, demon, dragonkin, elemental, giant, humanoid, mechanical, undead or unknown); the ranker's is unknown")
	flag.StringVar(&o.bisDir, "bis-dir", "", "directory holding the published <spec>.json entries; defaults to the build's bis directory (point it at a ranker run's -out to break that run down)")
	flag.BoolVar(&o.tank, "tank", false, "print the tank fight: final defensive stats, tank figures and the boss's swing outcomes (tank specs only)")
	flag.BoolVar(&o.jsonOut, "json", false, "print the totals as JSON")
	flag.Parse()
	if err := run(o); err != nil {
		fmt.Fprintln(os.Stderr, "spec-breakdown:", err)
		os.Exit(1)
	}
}

func run(o options) error {
	if o.spec == "" {
		return fmt.Errorf("-spec is required")
	}
	if o.heal {
		out, err := runHealReport(o)
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	}
	if o.tank {
		out, err := runTankReport(o)
		if err != nil {
			return err
		}
		fmt.Print(out)
		return nil
	}
	req, err := legalRequest(o)
	if err != nil {
		return err
	}
	forced, err := parseForcedRanks(o.rank)
	if err != nil {
		return err
	}
	rotation, err := rotationFor(o, forced, specs.ByKey[o.spec].ClassSlug, req.Character.Race, req.Character.Level)
	if err != nil {
		return err
	}
	est, player, err := inproc.PlainRunWithRotation(req, rotation)
	if err != nil {
		return err
	}
	names, err := loadSpellNames(o.repoRoot)
	if err != nil {
		return err
	}
	secs, err := request.FightSeconds(req)
	if err != nil {
		return err
	}
	report := newReport(req, secs, est, player, names)
	if o.jsonOut {
		return json.NewEncoder(os.Stdout).Encode(report)
	}
	fmt.Print(report.markdown())
	return nil
}

// loadRotation reads the "rotation" object of a curated rotation file; an
// empty path keeps the spec's own rotation. Its ranks are rewritten
// exactly as the published rotation's are: the spec's rank choices under
// preset for race for the ids it names, otherwise (below 60) the highest rank
// learned by level that is not above the authored one, so a probe that
// names a lower rank keeps it. At 60 an id with no row stays as written.
func loadRotation(path, spec, class, preset, race string, level int) (*proto.APLRotation, error) {
	if path == "" {
		return nil, nil
	}
	raw, err := readRotationFile(path)
	if err != nil {
		return nil, err
	}
	return decodeRotation(raw, path, spec, class, preset, race, level)
}

// readRotationFile is the "rotation" object of a curated rotation file.
func readRotationFile(path string) (json.RawMessage, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var file struct {
		Rotation json.RawMessage `json:"rotation"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	if len(file.Rotation) == 0 {
		return nil, fmt.Errorf("%s has no \"rotation\" object", path)
	}
	return file.Rotation, nil
}

// decodeRotation rewrites a rotation's ranks as the published one's are
// (see loadRotation) and parses it.
func decodeRotation(raw []byte, path, spec, class, preset, race string, level int) (*proto.APLRotation, error) {
	raw, err := request.RewriteRotationRanks(raw, spec, class, preset, race, level)
	if err != nil {
		return nil, fmt.Errorf("rewriting %s to level %d: %w", path, level, err)
	}
	return parseRotation(raw, path)
}

// parseRotation parses a rotation whose ranks are already rewritten.
func parseRotation(raw []byte, path string) (*proto.APLRotation, error) {
	rotation := &proto.APLRotation{}
	if err := protojson.Unmarshal(raw, rotation); err != nil {
		return nil, fmt.Errorf("decoding the rotation in %s: %w", path, err)
	}
	return rotation, nil
}

// forcedRank is one -rank pair: the authored spell id and the rank forced on
// every reference to it.
type forcedRank struct {
	authored int32
	rank     int
}

// parseForcedRanks reads -rank's comma list of authoredID=rank pairs.
func parseForcedRanks(list string) ([]forcedRank, error) {
	if list == "" {
		return nil, nil
	}
	var out []forcedRank
	for _, pair := range strings.Split(list, ",") {
		idText, rankText, ok := strings.Cut(strings.TrimSpace(pair), "=")
		id, idErr := strconv.ParseInt(idText, 10, 32)
		rank, rankErr := strconv.Atoi(rankText)
		if !ok || idErr != nil || rankErr != nil || id <= 0 || rank <= 0 {
			return nil, fmt.Errorf("-rank wants authoredSpellID=rank pairs, got %q", pair)
		}
		out = append(out, forcedRank{authored: int32(id), rank: rank})
	}
	return out, nil
}

// rotationFor is the rotation a run uses: -rotation's file, or with -rank the
// spec's own curated rotation when -rotation is not given, with each forced
// rank on every reference to its spell (request.ForceRotationRank) before
// the ordinary rank rewrite, which skips the rows on a forced spell
// (request.RewriteForcedRotationRanks): a forced rank always wins.
func rotationFor(o options, forced []forcedRank, class, race string, level int) (*proto.APLRotation, error) {
	if len(forced) == 0 {
		return loadRotation(o.rotation, o.spec, class, o.preset, race, level)
	}
	path := o.rotation
	if path == "" {
		path = filepath.Join(o.repoRoot, "data", "curated", "apl", o.spec+".json")
	}
	raw, err := readRotationFile(path)
	if err != nil {
		return nil, err
	}
	authored := make([]int32, 0, len(forced))
	for _, f := range forced {
		if raw, err = request.ForceRotationRank(raw, class, f.authored, f.rank, level); err != nil {
			return nil, fmt.Errorf("-rank %d=%d: %w", f.authored, f.rank, err)
		}
		authored = append(authored, f.authored)
	}
	if raw, err = request.RewriteForcedRotationRanks(raw, o.spec, class, o.preset, race, level, authored); err != nil {
		return nil, fmt.Errorf("rewriting %s to level %d: %w", path, level, err)
	}
	return parseRotation(raw, path)
}

type bisEntry struct {
	Band    int       `json:"band"`
	Preset  string    `json:"preset"`
	Faction string    `json:"faction"`
	Race    string    `json:"race"`
	Talents string    `json:"talents"`
	Slots   []bisSlot `json:"slots"`
}

// bisSlot is one published slot: its pick, the permanent enchant the raid
// pass ranked it with (leveling-bis enchants.go; 0 on a bare entry) and the
// band's runners-up for it.
type bisSlot struct {
	Slot         string    `json:"slot"`
	ItemID       int       `json:"item_id"`
	EnchantID    int       `json:"enchant_id"`
	Alternatives []bisItem `json:"alternatives"`
	RunnersUp    []bisItem `json:"runners_up"`
}

type bisItem struct {
	ItemID int `json:"item_id"`
}

// alternatives is each slot's runners-up, best first: the alternatives the
// ranker verified, then its other runners-up.
func (e bisEntry) alternatives() map[string][]int {
	out := make(map[string][]int, len(e.Slots))
	for _, s := range e.Slots {
		for _, list := range [][]bisItem{s.Alternatives, s.RunnersUp} {
			for _, a := range list {
				out[s.Slot] = append(out[s.Slot], a.ItemID)
			}
		}
	}
	return out
}

// bisDirectory is where the published entries are read from: the active
// build's bis/, or (leveling.BisDir) the newest other build's until the
// nightly has ranked the active one.
func (o options) bisDirectory() (string, error) {
	if o.bisDir != "" {
		return o.bisDir, nil
	}
	dir, err := buildDirectory(o.repoRoot)
	if err != nil {
		return "", err
	}
	return leveling.BisDir(dir), nil
}

func loadEntry(o options) (bisEntry, error) {
	bisDir, err := o.bisDirectory()
	if err != nil {
		return bisEntry{}, err
	}
	raw, err := os.ReadFile(filepath.Join(bisDir, o.spec+".json"))
	if err != nil {
		return bisEntry{}, err
	}
	var file struct {
		Bands []bisEntry `json:"bands"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return bisEntry{}, err
	}
	for _, e := range file.Bands {
		if e.Band == o.band && e.Preset == o.preset && e.Faction == o.faction {
			return e, nil
		}
	}
	return bisEntry{}, fmt.Errorf("no %s band %d %s %s entry", o.spec, o.band, o.preset, o.faction)
}

// lowestBand is the lowest band the ranker publishes; levelFloor is the lowest
// level a talent build exists at. -level on the lowest band reaches down to
// levelFloor on that band's entry, since no lower entry exists.
const (
	lowestBand = 20
	levelFloor = 10
)

// characterLevel is the level the band's entry is simmed at: the band itself,
// or -level when it falls inside the band's ten levels (a rank row spans them),
// or, on the lowest band, below it down to levelFloor.
func (o options) characterLevel() (int, error) {
	if o.level == 0 {
		return o.band, nil
	}
	top := min(o.band+9, 60)
	floor := o.band
	if o.band == lowestBand {
		floor = levelFloor
	}
	if o.level < floor || o.level > top {
		return 0, fmt.Errorf("-level %d is outside band %d (%d to %d)", o.level, o.band, floor, top)
	}
	return o.level, nil
}

// characterTalents is -talents, else with -level (and not -band-talents) the
// talents the guide's published order gives at that level (the ladder string
// the site and the addon show for it, leveling.GuideLevelTalents; the raid
// build under the raid preset), else the entry's. Round 24 hunter: rank rows
// judged at band+N used to run without the points a player has by then.
func (o options) characterTalents(entry bisEntry, level int) (string, error) {
	if o.talents != "" {
		return o.talents, nil
	}
	if o.level == 0 || o.bandTalents {
		return entry.Talents, nil
	}
	build, err := leveling.ReadActiveBuild(o.repoRoot)
	if err != nil {
		return "", err
	}
	return leveling.GuideLevelTalents(o.repoRoot, build, specs.ByKey[o.spec], level, o.preset == request.RaidPreset)
}

// levelGear is gear as a level below the band holds it (request.LevelGear.Fit):
// a slot the level cannot wear takes a runner-up or a weapon it can.
func levelGear(o options, entry bisEntry, gear []api.GearSlot, level int) ([]api.GearSlot, error) {
	build, err := leveling.ReadActiveBuild(o.repoRoot)
	if err != nil {
		return nil, err
	}
	fit, err := request.LoadLevelGear(o.repoRoot, build, o.spec, o.faction)
	if err != nil {
		return nil, err
	}
	return fit.Fit(gear, entry.alternatives(), level), nil
}

func buildRequest(o options) (api.SimRequest, error) {
	level, err := o.characterLevel()
	if err != nil {
		return api.SimRequest{}, err
	}
	entry, err := loadEntry(o)
	if err != nil {
		return api.SimRequest{}, err
	}
	info, ok := specs.ByKey[o.spec]
	if !ok {
		return api.SimRequest{}, fmt.Errorf("unknown spec %q", o.spec)
	}
	// The entry's enchants travel with its items: the ranker measured the
	// published set wearing them, and a breakdown that dropped them read a
	// raid caster 8 to 10% under its own page (2026-11-11 casters lane).
	gear := make([]api.GearSlot, 0, len(entry.Slots))
	for _, s := range entry.Slots {
		gear = append(gear, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID, Enchant: s.EnchantID})
	}
	gear, err = replaceGear(gear, o.gear)
	if err != nil {
		return api.SimRequest{}, err
	}
	if level < o.band {
		if gear, err = levelGear(o, entry, gear, level); err != nil {
			return api.SimRequest{}, err
		}
	}
	talents, err := o.characterTalents(entry, level)
	if err != nil {
		return api.SimRequest{}, err
	}
	kitBuffs := leveling.KitBuffs(o.spec, min(level, o.band))
	kitConsumes := leveling.KitConsumes(o.spec, min(level, o.band))
	if o.preset == request.RaidPreset {
		presets, err := request.LoadPresets(filepath.Join(o.repoRoot, "data", "curated", "presets.json"))
		if err != nil {
			return api.SimRequest{}, err
		}
		resolved, err := presets.Resolve(request.RaidPreset, info)
		if err != nil {
			return api.SimRequest{}, err
		}
		kitBuffs, kitConsumes = resolved.Layer(kitBuffs, kitConsumes)
	}
	if o.consumes != "" {
		kitConsumes = strings.Split(o.consumes, ",")
	}
	if o.buffs != "" {
		kitBuffs = append(kitBuffs, strings.Split(o.buffs, ",")...)
	}
	encounter := leveling.RankerEncounter()
	if o.targetType != "" {
		encounter.TargetType = o.targetType
	}
	if o.duration > 0 {
		encounter.DurationSec = o.duration
	}
	if o.targets > 0 {
		encounter.Targets = o.targets
	}
	race := entry.Race
	if o.race != "" {
		race = o.race
	}
	character := api.CharacterSpec{
		Name: "breakdown", Race: race, Class: info.ClassSlug, Level: level,
		Talents: talents, Gear: gear, Buffs: kitBuffs, Consumes: kitConsumes,
	}
	// Where the ranker stands the spec (leveling-bis's bandCharacter), so
	// a range-gated spell behaves the same in both.
	if leveling.NoMeleeAutoAttackSpecs[o.spec] {
		character.DistanceFromTarget = leveling.CasterDistanceFromTarget
	}
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Preset:        o.preset,
		Spec:          o.spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     character,
		Encounter:     encounter,
		Iterations:    o.iterations,
		RandomSeed:    o.seed,
	}, nil
}

func loadSpellNames(repoRoot string) (map[int32]string, error) {
	dir, err := buildDirectory(repoRoot)
	if err != nil {
		return nil, err
	}
	raw, err := os.ReadFile(filepath.Join(dir, "spells.json"))
	if err != nil {
		return nil, err
	}
	var rows []struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(raw, &rows); err != nil {
		return nil, err
	}
	out := make(map[int32]string, len(rows))
	for _, r := range rows {
		out[r.ID] = r.Name
	}
	return out, nil
}

// actionRow is one engine action's totals, per second of the fight.
type actionRow struct {
	Name                                  string
	Passive                               bool
	DPS, Share                            float64
	CastsPerIter, HitsPerIter             float64
	CritPct, MissPct, DodgePct, GlancePct float64
	GlanceDPS, AvgHit                     float64
}

type auraRow struct {
	Name         string
	UptimePct    float64
	ProcsPerIter float64
}

type resourceRow struct {
	Name          string
	Type          string
	GainPerIter   float64
	ActualPerIter float64
}

type report struct {
	Spec, Preset       string
	Consumes, Buffs    []string
	DPS, Error         float64
	FightSeconds       float64
	SecondsOOM         float64 // average seconds per iteration spent out of mana
	Actions            []actionRow
	Auras              []auraRow
	Resources          []resourceRow
	GlancePenaltyShare float64
	MissShareOfSwings  float64
}

func newReport(req api.SimRequest, secs float64, est api.Estimate, p *proto.UnitMetrics, names map[int32]string) report {
	iters := float64(req.Iterations)
	r := report{Spec: req.Spec, Consumes: req.Character.Consumes, Buffs: req.Character.Buffs, DPS: est.Mean, Error: est.Error,
		FightSeconds: secs, SecondsOOM: p.GetSecondsOomAvg()}
	for _, a := range p.GetActions() {
		r.Actions = append(r.Actions, foldAction(a, iters, secs, est.Mean, names))
	}
	// Pet damage counts toward the player's DPS, so its rows sit in the same
	// table under the pet's name.
	for _, pet := range p.GetPets() {
		for _, a := range pet.GetActions() {
			row := foldAction(a, iters, secs, est.Mean, names)
			row.Name = pet.GetName() + ": " + row.Name
			r.Actions = append(r.Actions, row)
		}
	}
	sort.Slice(r.Actions, func(i, j int) bool { return r.Actions[i].DPS > r.Actions[j].DPS })
	for _, a := range p.GetAuras() {
		r.Auras = append(r.Auras, auraRow{
			Name:         actionName(a.GetId(), names),
			UptimePct:    100 * a.GetUptimeSecondsAvg() / secs,
			ProcsPerIter: a.GetProcsAvg(),
		})
	}
	sort.Slice(r.Auras, func(i, j int) bool { return r.Auras[i].UptimePct > r.Auras[j].UptimePct })
	r.Resources = resourceRows(p, iters, names)
	return r
}

// resourceRows is every resource gain of the unit, per iteration.
func resourceRows(p *proto.UnitMetrics, iters float64, names map[int32]string) []resourceRow {
	rows := make([]resourceRow, 0, len(p.GetResources()))
	for _, rs := range p.GetResources() {
		rows = append(rows, resourceRow{
			Name: actionName(rs.GetId(), names), Type: rs.GetType().String(),
			GainPerIter: rs.GetGain() / iters, ActualPerIter: rs.GetActualGain() / iters,
		})
	}
	return rows
}

// writeResourceTable prints the resource gains as the report's markdown table.
func writeResourceTable(b *strings.Builder, rows []resourceRow) {
	b.WriteString("\n| Resource | Type | Gain/iter | Actual/iter |\n|---|---|---|---|\n")
	for _, a := range rows {
		fmt.Fprintf(b, "| %s | %s | %.1f | %.1f |\n", a.Name, a.Type, a.GainPerIter, a.ActualPerIter)
	}
}

func foldAction(a *proto.ActionMetrics, iters, secs, totalDPS float64, names map[int32]string) actionRow {
	row := actionRow{Name: actionName(a.GetId(), names), Passive: a.GetIsPassive()}
	var damage, glance, hits, crits, misses, dodges, glances, casts, blocks float64
	for _, t := range a.GetTargets() {
		damage += t.GetDamage()
		glance += t.GetGlanceDamage()
		hits += float64(t.GetHits())
		crits += float64(t.GetCrits() + t.GetCritTicks())
		blocks += float64(t.GetBlocks() + t.GetParries())
		misses += float64(t.GetMisses())
		dodges += float64(t.GetDodges())
		glances += float64(t.GetGlances())
		casts += float64(t.GetCasts())
		hits += float64(t.GetTicks())
	}
	row.DPS = damage / iters / secs
	row.GlanceDPS = glance / iters / secs
	row.CastsPerIter = casts / iters
	if totalDPS > 0 {
		row.Share = 100 * row.DPS / totalDPS
	}
	// The engine counts a white hit as hit, crit, glance, miss or dodge
	// separately (Hits excludes crits and glances), so the outcome total
	// is the denominator of every percentage.
	attempts := hits + crits + glances + misses + dodges + blocks
	if attempts > 0 {
		row.CritPct = 100 * crits / attempts
		row.MissPct = 100 * misses / attempts
		row.DodgePct = 100 * dodges / attempts
		row.GlancePct = 100 * glances / attempts
	}
	if landed := hits + crits + glances; landed > 0 {
		row.AvgHit = damage / landed
	}
	row.HitsPerIter = (hits + crits + glances) / iters
	return row
}

func actionName(id *proto.ActionID, names map[int32]string) string {
	if id == nil {
		return "?"
	}
	tag := ""
	if id.Tag != 0 {
		tag = fmt.Sprintf("#%d", id.Tag)
	}
	switch raw := id.RawId.(type) {
	case *proto.ActionID_SpellId:
		name := names[raw.SpellId]
		if name == "" {
			name = "spell"
		}
		return fmt.Sprintf("%s (%d)%s", name, raw.SpellId, tag)
	case *proto.ActionID_ItemId:
		return fmt.Sprintf("item %d%s", raw.ItemId, tag)
	case *proto.ActionID_OtherId:
		return raw.OtherId.String() + tag
	}
	return "?"
}

func (r report) markdown() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s  DPS %.2f +/- %.2f\nout of mana: %.1f s of %.0f s\nconsumes: %v\nbuffs: %v\n\n",
		r.Spec, r.DPS, r.Error, r.SecondsOOM, r.FightSeconds, r.Consumes, r.Buffs)
	b.WriteString("| Action | DPS | Share % | Casts/iter | Hits/iter | Avg hit | Crit % | Miss % | Dodge % | Glance % | Glance DPS |\n|---|---|---|---|---|---|---|---|---|---|---|\n")
	for _, a := range r.Actions {
		if a.DPS == 0 && a.CastsPerIter == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %.1f | %.1f | %.2f | %.2f | %.0f | %.1f | %.1f | %.1f | %.1f | %.1f |\n",
			a.Name, a.DPS, a.Share, a.CastsPerIter, a.HitsPerIter, a.AvgHit, a.CritPct, a.MissPct, a.DodgePct, a.GlancePct, a.GlanceDPS)
	}
	b.WriteString("\n| Aura | Uptime % | Procs/iter |\n|---|---|---|\n")
	for _, a := range r.Auras {
		if a.UptimePct == 0 && a.ProcsPerIter == 0 {
			continue
		}
		fmt.Fprintf(&b, "| %s | %.1f | %.2f |\n", a.Name, a.UptimePct, a.ProcsPerIter)
	}
	writeResourceTable(&b, r.Resources)
	return b.String()
}

// replaceGear applies a "slot:item_id,slot:item_id" list over the published
// entry's gear: a slot the entry has takes the new item, a slot it lacks is
// added.
func replaceGear(gear []api.GearSlot, list string) ([]api.GearSlot, error) {
	if list == "" {
		return gear, nil
	}
	out := append([]api.GearSlot(nil), gear...)
	for _, pair := range strings.Split(list, ",") {
		slot, id, found := strings.Cut(pair, ":")
		itemID, err := strconv.Atoi(id)
		if !found || err != nil || slot == "" {
			return nil, fmt.Errorf("-gear %q: want slot:item_id", pair)
		}
		replaced := false
		for i := range out {
			if out[i].Slot == slot {
				out[i].ItemID = itemID
				replaced = true
			}
		}
		if !replaced {
			out = append(out, api.GearSlot{Slot: slot, ItemID: itemID})
		}
	}
	return out, nil
}
