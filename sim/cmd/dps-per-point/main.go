// Command dps-per-point measures, for each spec, the damage per second
// (healing per second for a healer) one more point of the spec's
// reference stat is worth, and writes it to
// data/builds/<active build>/dps-per-point.json:
//
//	{ "warrior-arms": { "reference": "attack_power", "dps": 0.42, "healing": false }, ... }
//
// The data pipeline reads that file into Data.lua's dpsPerPoint
// (docs/contracts/addon-v2.md, contract 1); the addon's Progress
// multiplies a score delta, which is in reference-stat points, by it to
// say "+12 DPS".
//
// The figure is the stat-weights sweep's own raw derivative before it is
// normalised (sim/adapter.Weights divides every weight by exactly this
// number): one weights run per spec that weighs only the reference stat,
// on the level-60 character the raid preset's BiS band wears - the
// guide's raid build (leveling.GuideRaidBuildTalents), that band's gear
// and race, the class kit and the raid preset's buffs and consumables,
// against a level-63 target of no creature type. A healer is swept over
// the curated heal profile and reads the engine's healing table.
//
// A tank's reference stat is stamina, which does no damage, so tanks get
// no entry; neither does a spec whose figure is not positive beyond its
// own standard error. A spec with no entry is the addon's cue to fall
// back to "score" units.
//
//	go run ./cmd/dps-per-point -repo-root .. [-specs warrior-arms,priest-holy]
//
// -specs measures only the named specs and merges them into the
// existing file, so a partial run never drops another spec's figure.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/statid"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	// level is the only level measured: the raid build is a level-60 build.
	level = 60
	// bandFaction is whose raid band supplies the gear and race. Both
	// factions' raid bands wear the same set; the racial difference in a
	// per-point derivative is far below its own error.
	bandFaction = "alliance"
	// healingReference is the reference stat that makes a spec a healer
	// for this file: its figure is healing per second.
	healingReference = "healing_power"
	tankRole         = "tank"
	// decimals is how finely the file states a figure; the sweep's own
	// error is larger than the last kept digit.
	decimals = 4
	// healIterationScale multiplies a healer's sweep: effective healing
	// against the heal profile's fake raid is far noisier per point than
	// damage against a dummy (priest-holy measured 0.32 ± 0.14 at the
	// damage specs' 500 iterations, 0.31 ± 0.05 at 8x).
	healIterationScale = 8
)

// entry is one spec's figure as dps-per-point.json carries it.
type entry struct {
	Reference string  `json:"reference"`
	DPS       float64 `json:"dps"`
	Healing   bool    `json:"healing"`
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("dps-per-point", flag.ContinueOnError)
	repoRoot := fs.String("repo-root", ".", "the site repository root")
	only := fs.String("specs", "", "comma-separated spec keys to measure (default: every spec)")
	iterations := fs.Int("iterations", 500, "iterations per sweep direction, before api.WeightsIterationsFactor")
	seed := fs.Int64("seed", 7, "random seed for every sweep")
	if err := fs.Parse(args); err != nil {
		return err
	}
	active, err := leveling.ReadActiveBuild(*repoRoot)
	if err != nil {
		return err
	}
	picked, err := pickSpecs(*only)
	if err != nil {
		return err
	}
	profile, err := request.LoadHealProfile(filepath.Join(*repoRoot, "data", "curated", "heal-profile.json"))
	if err != nil {
		return err
	}
	characterFor := func(spec specs.Spec) (api.CharacterSpec, error) {
		return raidCharacter(*repoRoot, active, spec)
	}
	presetFor := func(spec specs.Spec) (*request.ResolvedPreset, error) {
		return request.ResolveFromFile(filepath.Join(*repoRoot, "data", "curated", "presets.json"), request.RaidPreset, spec)
	}
	got, skipped := measureAll(picked, characterFor, presetFor, engine{profile: profile}, *iterations, *seed)
	for _, s := range skipped {
		log.Printf("dps-per-point: %s", s)
	}
	path := filepath.Join(*repoRoot, "data", "builds", active, "dps-per-point.json")
	if err := writeMerged(path, got); err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func pickSpecs(only string) ([]specs.Spec, error) {
	if only == "" {
		return specs.All, nil
	}
	var out []specs.Spec
	for _, key := range strings.Split(only, ",") {
		spec, ok := specs.ByKey[strings.TrimSpace(key)]
		if !ok {
			return nil, fmt.Errorf("no spec %q in sim/specs", key)
		}
		out = append(out, spec)
	}
	return out, nil
}

// measured is false for a tank: stamina, its reference, does no damage.
func measured(spec specs.Spec) bool { return spec.Role != tankRole }

// healing is true when the spec's reference stat is a healing stat, so
// its figure is healing per second.
func healing(spec specs.Spec) bool { return spec.ReferenceStat == healingReference }

// measurer is the engine: the reference stat's raw per-point derivative
// and its standard error for one weights request.
type measurer interface {
	Measure(spec specs.Spec, req api.SimRequest) (perPoint, stdErr float64, err error)
}

type characterFunc func(specs.Spec) (api.CharacterSpec, error)

type presetFunc func(specs.Spec) (*request.ResolvedPreset, error)

// measureAll measures every picked spec, returning the trusted figures
// and one line per spec left out, saying why.
func measureAll(picked []specs.Spec, characterFor characterFunc, presetFor presetFunc, m measurer, iterations int, seed int64) (map[string]entry, []string) {
	out := map[string]entry{}
	var skipped []string
	for _, spec := range picked {
		if !measured(spec) {
			skipped = append(skipped, fmt.Sprintf("%s: a tank's reference stat (%s) does no damage; no entry", spec.Spec, spec.ReferenceStat))
			continue
		}
		start := time.Now()
		ch, err := characterFor(spec)
		if err != nil {
			skipped = append(skipped, fmt.Sprintf("%s: %v", spec.Spec, err))
			continue
		}
		var preset *request.ResolvedPreset
		if presetFor != nil {
			if preset, err = presetFor(spec); err != nil {
				skipped = append(skipped, fmt.Sprintf("%s: resolving the raid preset: %v", spec.Spec, err))
				continue
			}
		}
		perPoint, stdErr, err := m.Measure(spec, weightsRequest(spec, ch, preset, iterations, seed))
		if err != nil {
			skipped = append(skipped, fmt.Sprintf("%s: %v", spec.Spec, err))
			continue
		}
		if perPoint <= stdErr {
			skipped = append(skipped, fmt.Sprintf("%s: %s measured %.4f ± %.4f per point, not positive beyond its own error; no entry", spec.Spec, spec.ReferenceStat, perPoint, stdErr))
			continue
		}
		log.Printf("dps-per-point: %s: %.4f ± %.4f per %s in %s", spec.Spec, perPoint, stdErr, spec.ReferenceStat, time.Since(start).Round(time.Second))
		out[spec.Spec] = entry{Reference: spec.ReferenceStat, DPS: perPoint, Healing: healing(spec)}
	}
	return out, skipped
}

// weightsRequest is the sweep for one spec: only the reference stat is
// weighed (its derivative does not depend on which other stats are
// swept beside it, and one stat is a third of the cost of six), on ch
// wearing the class kit and, when preset is not nil, the preset's
// buffs and consumables.
func weightsRequest(spec specs.Spec, ch api.CharacterSpec, preset *request.ResolvedPreset, iterations int, seed int64) api.SimRequest {
	buffs, consumes := leveling.KitBuffs(spec.Spec, ch.Level), leveling.KitConsumes(spec.Spec, ch.Level)
	if preset != nil {
		buffs, consumes = preset.Layer(buffs, consumes)
	}
	ch.Buffs, ch.Consumes = buffs, consumes
	if leveling.NoMeleeAutoAttackSpecs[spec.Spec] {
		ch.DistanceFromTarget = leveling.CasterDistanceFromTarget
	}
	encounter := api.DefaultEncounter()
	encounter.TargetType = api.TargetTypeUnknown
	if healing(spec) {
		iterations *= healIterationScale
	}
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character:     ch,
		Encounter:     encounter,
		Iterations:    iterations,
		RandomSeed:    seed,
		Weights:       &api.WeightsSpec{Stats: []string{spec.ReferenceStat}, Reference: spec.ReferenceStat},
	}
}

// raidCharacter is spec's level-60 raid character: the guide's raid
// build, re-read by stable talent id against the active trees and
// written in the compiled engine's own talent layout, wearing the raid
// preset band's BiS set and race.
func raidCharacter(repoRoot, active string, spec specs.Spec) (api.CharacterSpec, error) {
	stamp, digits, err := leveling.GuideRaidBuildTalents(repoRoot, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	if err := leveling.RequireGuideBuildMatchesActive(stamp, active); err != nil {
		return api.CharacterSpec{}, err
	}
	trees, err := leveling.LoadTalentTrees(repoRoot, active, spec.ClassSlug)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	positional := leveling.LadderTalentString(trees, leveling.GuideTalentTargets(trees, digits), spec.TreeIndex, level)
	engineDir, err := enginetalents.SourceDir(filepath.Join(repoRoot, "sim"))
	if err != nil {
		return api.CharacterSpec{}, err
	}
	layout, err := enginetalents.ForClass(engineDir, spec.ClassSlug)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	talents, err := layout.Reposition(trees, positional)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	band, err := loadRaidBand(filepath.Join(repoRoot, "data", "builds", active), spec.Spec)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	return api.CharacterSpec{
		Name:    "dps-per-point",
		Race:    band.Race,
		Class:   spec.ClassSlug,
		Level:   level,
		Talents: talents,
		Gear:    band.gear(),
	}, nil
}

// bisBand is the part of one data/builds/<build>/bis/<spec>.json band
// this command wears.
type bisBand struct {
	Band    int    `json:"band"`
	Preset  string `json:"preset"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Slots   []struct {
		Slot   string `json:"slot"`
		ItemID int    `json:"item_id"`
	} `json:"slots"`
}

func loadRaidBand(buildDir, spec string) (bisBand, error) {
	path := filepath.Join(leveling.BisDir(buildDir), spec+".json")
	b, err := os.ReadFile(path)
	if err != nil {
		return bisBand{}, err
	}
	var f struct {
		Bands []bisBand `json:"bands"`
	}
	if err := json.Unmarshal(b, &f); err != nil {
		return bisBand{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	for _, band := range f.Bands {
		if band.Band == level && band.Faction == bandFaction && band.Preset == request.RaidPreset {
			return band, nil
		}
	}
	return bisBand{}, fmt.Errorf("%s has no level-%d %s %s band", path, level, bandFaction, request.RaidPreset)
}

func (b bisBand) gear() []api.GearSlot {
	var out []api.GearSlot
	for _, s := range b.Slots {
		if s.ItemID != 0 {
			out = append(out, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
		}
	}
	return out
}

// engine is the real measurer: the wowsims engine in this process.
type engine struct {
	profile request.HealProfile
}

func (e engine) Measure(spec specs.Spec, req api.SimRequest) (float64, float64, error) {
	inproc.Register()
	if healing(spec) {
		// A healer is measured over the profile's one stated fight
		// (leveling-bis's healEngine.forProfile does the same).
		req.Encounter.DurationSec = e.profile.DurationSec
		req.Encounter.Variation = 0
	}
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return 0, 0, fmt.Errorf("building the weights request: %w", err)
	}
	if healing(spec) {
		e.profile.AttachWeights(engineReq)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return 0, 0, fmt.Errorf("attaching the item database: %w", err)
	}
	res := core.StatWeights(engineReq)
	if res.GetError() != nil && res.GetError().GetMessage() != "" {
		return 0, 0, errors.New(res.GetError().GetMessage())
	}
	values := res.GetDps()
	if healing(spec) {
		values = res.GetHps()
	}
	return referenceDerivative(values, req)
}

// referenceDerivative reads the reference stat's raw, un-normalised
// weight - the per-point figure sim/adapter.Weights normalises every
// other weight by - and converts the engine's population stdev to a
// standard error the way adapter.Weights does (divided by the square
// root of the merged sample count).
func referenceDerivative(values *proto.StatWeightValues, req api.SimRequest) (float64, float64, error) {
	reference, ok := statid.Parse(req.Weights.Reference)
	if !ok {
		return 0, 0, fmt.Errorf("reference stat %q is not a known stat id", req.Weights.Reference)
	}
	raw := values.GetWeights().GetStats()
	stdev := values.GetWeightsStdev().GetStats()
	if int(reference) >= len(raw) || int(reference) >= len(stdev) {
		return 0, 0, fmt.Errorf("the engine returned no weight for %s", req.Weights.Reference)
	}
	n := float64(req.Iterations * api.WeightsIterationsFactor)
	return raw[reference], stdev[reference] / math.Sqrt(n), nil
}

// writeMerged writes got over the figures already at path (a partial
// run keeps every spec it did not measure), one spec per line, keys
// sorted, figures rounded to decimals.
func writeMerged(path string, got map[string]entry) error {
	merged := map[string]entry{}
	if b, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(b, &merged); err != nil {
			return fmt.Errorf("decoding %s: %w", path, err)
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	for k, v := range got {
		merged[k] = v
	}
	keys := make([]string, 0, len(merged))
	for k := range merged {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	scale := math.Pow(10, decimals)
	var buf bytes.Buffer
	buf.WriteString("{\n")
	for i, k := range keys {
		e := merged[k]
		e.DPS = math.Round(e.DPS*scale) / scale
		key, _ := json.Marshal(k)
		value, err := json.Marshal(e)
		if err != nil {
			return err
		}
		sep := ","
		if i == len(keys)-1 {
			sep = ""
		}
		fmt.Fprintf(&buf, "  %s: %s%s\n", key, value, sep)
	}
	buf.WriteString("}\n")
	return os.WriteFile(path, buf.Bytes(), 0o644)
}
