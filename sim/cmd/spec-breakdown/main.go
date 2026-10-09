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
	faction    string
	iterations int
	seed       int64
	consumes   string
	buffs      string
	talents    string
	rotation   string
	jsonOut    bool
	tank       bool
	heal       bool
	duration   int
	bisDir     string
	gear       string
}

func main() {
	var o options
	flag.StringVar(&o.repoRoot, "repo-root", "..", "site repository root")
	flag.StringVar(&o.spec, "spec", "", "spec slug, e.g. rogue-combat")
	flag.StringVar(&o.preset, "preset", "raid", "bare or raid")
	flag.IntVar(&o.band, "band", 60, "BiS band")
	flag.StringVar(&o.faction, "faction", "alliance", "faction of the published entry")
	flag.IntVar(&o.iterations, "iterations", 20000, "iterations")
	flag.Int64Var(&o.seed, "seed", 1, "random seed")
	flag.StringVar(&o.consumes, "consumes", "", "comma list replacing the kit and preset consumables")
	flag.StringVar(&o.buffs, "buffs", "", "comma list added to the request's buffs")
	flag.StringVar(&o.talents, "talents", "", "talent string replacing the published one")
	flag.StringVar(&o.rotation, "rotation", "", "JSON file whose \"rotation\" object (the curated file's shape) replaces the spec's own priority list")
	flag.StringVar(&o.gear, "gear", "", "comma list of slot:item_id pairs replacing the published entry's items, e.g. off_hand:20688")
	flag.BoolVar(&o.heal, "heal", false, "print the healer fight: mana income and spend against data/curated/heal-profile.json (healer specs only)")
	flag.IntVar(&o.duration, "duration", 0, "with -heal: fight length in seconds, instead of the profile's (a shorter window shows a healer's mana rates before it runs dry)")
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
	req, err := buildRequest(o)
	if err != nil {
		return err
	}
	rotation, err := loadRotation(o.rotation)
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
	report := newReport(req, est, player, names)
	if o.jsonOut {
		return json.NewEncoder(os.Stdout).Encode(report)
	}
	fmt.Print(report.markdown())
	return nil
}

// loadRotation reads the "rotation" object of a curated rotation file; an
// empty path keeps the spec's own rotation.
func loadRotation(path string) (*proto.APLRotation, error) {
	if path == "" {
		return nil, nil
	}
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
	rotation := &proto.APLRotation{}
	if err := protojson.Unmarshal(file.Rotation, rotation); err != nil {
		return nil, fmt.Errorf("decoding the rotation in %s: %w", path, err)
	}
	return rotation, nil
}

type bisEntry struct {
	Band    int    `json:"band"`
	Preset  string `json:"preset"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Talents string `json:"talents"`
	Slots   []struct {
		Slot   string `json:"slot"`
		ItemID int    `json:"item_id"`
	} `json:"slots"`
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

func buildRequest(o options) (api.SimRequest, error) {
	entry, err := loadEntry(o)
	if err != nil {
		return api.SimRequest{}, err
	}
	info, ok := specs.ByKey[o.spec]
	if !ok {
		return api.SimRequest{}, fmt.Errorf("unknown spec %q", o.spec)
	}
	gear := make([]api.GearSlot, 0, len(entry.Slots))
	for _, s := range entry.Slots {
		gear = append(gear, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
	}
	gear, err = replaceGear(gear, o.gear)
	if err != nil {
		return api.SimRequest{}, err
	}
	talents := entry.Talents
	if o.talents != "" {
		talents = o.talents
	}
	kitBuffs := leveling.KitBuffs(o.spec, o.band)
	kitConsumes := leveling.KitConsumes(o.spec, o.band)
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
	encounter := api.DefaultEncounter()
	encounter.TargetType = api.TargetTypeUnknown
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          o.spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character: api.CharacterSpec{
			Name: "breakdown", Race: entry.Race, Class: info.ClassSlug, Level: o.band,
			Talents: talents, Gear: gear, Buffs: kitBuffs, Consumes: kitConsumes,
		},
		Encounter:  encounter,
		Iterations: o.iterations,
		RandomSeed: o.seed,
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
	Actions            []actionRow
	Auras              []auraRow
	Resources          []resourceRow
	GlancePenaltyShare float64
	MissShareOfSwings  float64
}

func newReport(req api.SimRequest, est api.Estimate, p *proto.UnitMetrics, names map[int32]string) report {
	iters := float64(req.Iterations)
	secs := float64(req.Encounter.DurationSec)
	r := report{Spec: req.Spec, Consumes: req.Character.Consumes, Buffs: req.Character.Buffs, DPS: est.Mean, Error: est.Error}
	for _, a := range p.GetActions() {
		r.Actions = append(r.Actions, foldAction(a, iters, secs, est.Mean, names))
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
	for _, rs := range p.GetResources() {
		r.Resources = append(r.Resources, resourceRow{
			Name: actionName(rs.GetId(), names), Type: rs.GetType().String(),
			GainPerIter: rs.GetGain() / iters, ActualPerIter: rs.GetActualGain() / iters,
		})
	}
	return r
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
	fmt.Fprintf(&b, "%s  DPS %.1f +/- %.1f\nconsumes: %v\nbuffs: %v\n\n", r.Spec, r.DPS, r.Error, r.Consumes, r.Buffs)
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
	b.WriteString("\n| Resource | Type | Gain/iter | Actual/iter |\n|---|---|---|---|\n")
	for _, a := range r.Resources {
		fmt.Fprintf(&b, "| %s | %s | %.1f | %.1f |\n", a.Name, a.Type, a.GainPerIter, a.ActualPerIter)
	}
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
