// Command dps-per-point measures, for each spec, the damage per second
// (healing per second for a healer) one more point of the spec's
// reference stat is worth, and writes it to
// data/builds/<active build>/dps-per-point.json:
//
//	{ "warrior-arms": { "reference": "attack_power", "dps": 0.2222, "healing": false,
//	    "weapon_dps": 3.1536,
//	    "bands": { "20": { "level": 25, "dps": 0.0385, "dps_error": 0,
//	                       "weapon_dps": 0.5432, "weapon_dps_error": 0.0007 }, ... } }, ... }
//
// The data pipeline reads that file into Data.lua's dpsPerPoint
// (docs/contracts/addon-v2.md, contract 1); the addon's Progress
// multiplies a score delta, which is in reference-stat points, by it to
// say "+12 DPS".
//
// Every sweep also weighs the weapon the reference stat swings (the main
// hand, or the ranged weapon for a ranged-attack-power spec) as the
// engine's own weapon-DPS pseudo stat, so "weapon_dps" is the measured
// DPS one more point of that weapon's damage per second is worth (task
// L1.5): the addon no longer has to assume 14 attack power per weapon DPS
// at the level-60 rate. And every leveling band (-bands, 20 through 60)
// is measured too, on that band's own character: the bare band's BiS set
// and race from bis/<spec>.json, the guide's leveling build at the
// band's level, the class kit and no preset, at the middle of the levels
// the addon reads the band for (bandLevel). Those figures go under
// "bands", and the data pipeline ships them as Data.lua's
// dpsPerPoint[spec][band].
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
// -specs measures only the named specs and -bands only the named bands;
// either way the run is merged into the existing file band by band, so a
// partial run never drops another spec's figure or another band's
// (-bands "" measures the level-60 figure alone and keeps every band).
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
	"strconv"
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
	// level is the raid character's level (the raid build is a level-60
	// build) and the level band 60 is measured at; the other bands are
	// measured at bandLevel.
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
	// defaultBands is leveling-bis's own band list (its defaultBandsFlag).
	defaultBands = "20,30,40,50,60"
	// bandStep is the width of a band: Gear.bandFor reads band N for
	// levels N through N+9.
	bandStep = 10
	// barePreset is the bis/<spec>.json band measured without a preset.
	barePreset = "bare"
)

// entry is one spec's figure as dps-per-point.json carries it.
type entry struct {
	Reference string  `json:"reference"`
	DPS       float64 `json:"dps"`
	Healing   bool    `json:"healing"`
	// WeaponDPS is the DPS (HPS for a healer) one more point of the swung
	// weapon's damage per second is worth on the level-60 raid character;
	// 0 when the spec does not swing it (a caster, a healer).
	WeaponDPS float64 `json:"weapon_dps"`
	// Bands is the same pair of figures per leveling band, keyed by the
	// band as a string ("20" ... "60").
	Bands map[string]bandEntry `json:"bands,omitempty"`
}

// bandEntry is one leveling band's measurement, with each figure's
// standard error.
type bandEntry struct {
	Level       int     `json:"level"`
	DPS         float64 `json:"dps"`
	DPSError    float64 `json:"dps_error"`
	WeaponDPS   float64 `json:"weapon_dps"`
	WeaponError float64 `json:"weapon_dps_error"`
}

// measurement is one sweep's two figures: the reference stat's raw
// per-point derivative and the swung weapon's per-weapon-DPS derivative,
// each with its standard error.
type measurement struct {
	PerPoint, StdErr  float64
	Weapon, WeaponErr float64
}

// bandLevel is the level band is measured at: the middle of the levels
// the addon reads it for (Gear.bandFor rounds a level down to its band,
// so band 20 serves 20-29 and is measured at 25), and 60 for band 60,
// the level cap.
func bandLevel(band int) int {
	if band >= level {
		return level
	}
	return band + bandStep/2
}

// weaponPseudoStat is the weapon a spec's reference stat swings: the
// ranged weapon for a ranged-attack-power spec, the main hand for every
// other (a caster's or a healer's measures zero, which is the point).
func weaponPseudoStat(spec specs.Spec) proto.PseudoStat {
	if spec.ReferenceStat == "ranged_attack_power" {
		return proto.PseudoStat_PseudoStatRangedDps
	}
	return proto.PseudoStat_PseudoStatMainHandDps
}

// trustedWeapon is the weapon term when it is positive beyond its own
// error, else 0: a term inside its error is noise, not a number.
func trustedWeapon(m measurement) float64 {
	if m.Weapon > m.WeaponErr {
		return m.Weapon
	}
	return 0
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
	bandsFlag := fs.String("bands", defaultBands, "comma-separated leveling bands to measure as well (empty: the level-60 raid figure only)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	bands, err := parseBands(*bandsFlag)
	if err != nil {
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
	m := engine{profile: profile}
	got, skipped := measureAll(picked, characterFor, presetFor, m, *iterations, *seed)
	bandCharacterFor := func(spec specs.Spec, band int) (api.CharacterSpec, error) {
		return bandCharacter(*repoRoot, active, spec, band)
	}
	byBand, bandSkipped := measureBands(picked, bands, bandCharacterFor, m, *iterations, *seed)
	for spec, e := range got {
		e.Bands = byBand[spec]
		got[spec] = e
	}
	for _, s := range append(skipped, bandSkipped...) {
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

// measurer is the engine: the reference stat's and the swung weapon's raw
// derivatives, with their standard errors, for one weights request.
type measurer interface {
	Measure(spec specs.Spec, req api.SimRequest) (measurement, error)
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
		got, err := m.Measure(spec, weightsRequest(spec, ch, preset, iterations, seed))
		if err != nil {
			skipped = append(skipped, fmt.Sprintf("%s: %v", spec.Spec, err))
			continue
		}
		if got.PerPoint <= got.StdErr {
			skipped = append(skipped, fmt.Sprintf("%s: %s measured %.4f ± %.4f per point, not positive beyond its own error; no entry", spec.Spec, spec.ReferenceStat, got.PerPoint, got.StdErr))
			continue
		}
		log.Printf("dps-per-point: %s: %.4f ± %.4f per %s, %.4f ± %.4f per weapon DPS, in %s", spec.Spec, got.PerPoint, got.StdErr, spec.ReferenceStat, got.Weapon, got.WeaponErr, time.Since(start).Round(time.Second))
		out[spec.Spec] = entry{Reference: spec.ReferenceStat, DPS: got.PerPoint, Healing: healing(spec), WeaponDPS: trustedWeapon(got)}
	}
	return out, skipped
}

type bandCharacterFunc func(spec specs.Spec, band int) (api.CharacterSpec, error)

// measureBands measures every picked spec at every band on that band's
// own character (no preset), returning spec -> band -> figures and one
// line per spec or band left out. A band whose reference figure is not
// positive beyond its error is left out; a weapon term inside its error
// is kept as 0 (trustedWeapon), with its error, since a caster's zero is
// a finding.
func measureBands(picked []specs.Spec, bands []int, characterFor bandCharacterFunc, m measurer, iterations int, seed int64) (map[string]map[string]bandEntry, []string) {
	out := map[string]map[string]bandEntry{}
	var skipped []string
	for _, spec := range picked {
		if !measured(spec) {
			skipped = append(skipped, fmt.Sprintf("%s: a tank's reference stat (%s) does no damage; no bands", spec.Spec, spec.ReferenceStat))
			continue
		}
		for _, band := range bands {
			ch, err := characterFor(spec, band)
			if err != nil {
				skipped = append(skipped, fmt.Sprintf("%s band %d: %v", spec.Spec, band, err))
				continue
			}
			got, err := m.Measure(spec, weightsRequest(spec, ch, nil, iterations, seed))
			if err != nil {
				skipped = append(skipped, fmt.Sprintf("%s band %d: %v", spec.Spec, band, err))
				continue
			}
			if got.PerPoint <= got.StdErr {
				skipped = append(skipped, fmt.Sprintf("%s band %d: %s measured %.4f ± %.4f per point, not positive beyond its own error", spec.Spec, band, spec.ReferenceStat, got.PerPoint, got.StdErr))
				continue
			}
			log.Printf("dps-per-point: %s band %d (level %d): %.4f ± %.4f per %s, %.4f ± %.4f per weapon DPS", spec.Spec, band, ch.Level, got.PerPoint, got.StdErr, spec.ReferenceStat, got.Weapon, got.WeaponErr)
			if out[spec.Spec] == nil {
				out[spec.Spec] = map[string]bandEntry{}
			}
			out[spec.Spec][strconv.Itoa(band)] = bandEntry{
				Level:       ch.Level,
				DPS:         got.PerPoint,
				DPSError:    got.StdErr,
				WeaponDPS:   trustedWeapon(got),
				WeaponError: got.WeaponErr,
			}
		}
	}
	return out, skipped
}

func parseBands(s string) ([]int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	var out []int
	for _, part := range strings.Split(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("-bands: %q is not a number: %w", part, err)
		}
		out = append(out, n)
	}
	return out, nil
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
	encounter := leveling.RankerEncounter()
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
	band, err := loadBand(filepath.Join(repoRoot, "data", "builds", active), spec.Spec, level, request.RaidPreset)
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

// bandCharacter is spec's character for one leveling band: the guide's
// leveling build spent to bandLevel(band), in the compiled engine's
// layout, wearing the bare band's alliance BiS set and race.
func bandCharacter(repoRoot, active string, spec specs.Spec, band int) (api.CharacterSpec, error) {
	stamp, digits, err := leveling.GuideBuildTalents(repoRoot, spec.ClassSlug, spec.SpecSlug)
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
	lvl := bandLevel(band)
	positional := leveling.LadderTalentString(trees, leveling.GuideTalentTargets(trees, digits), spec.TreeIndex, lvl)
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
	b, err := loadBand(filepath.Join(repoRoot, "data", "builds", active), spec.Spec, band, barePreset)
	if err != nil {
		return api.CharacterSpec{}, err
	}
	return api.CharacterSpec{
		Name:    "dps-per-point",
		Race:    b.Race,
		Class:   spec.ClassSlug,
		Level:   lvl,
		Talents: talents,
		Gear:    b.gear(),
	}, nil
}

// loadBand is one band of bis/<spec>.json: the alliance entry for band
// under preset.
func loadBand(buildDir, spec string, band int, preset string) (bisBand, error) {
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
	for _, b := range f.Bands {
		if b.Band == band && b.Faction == bandFaction && b.Preset == preset {
			return b, nil
		}
	}
	return bisBand{}, fmt.Errorf("%s has no level-%d %s %s band", path, band, bandFaction, preset)
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

func (e engine) Measure(spec specs.Spec, req api.SimRequest) (measurement, error) {
	inproc.Register()
	if healing(spec) {
		// A healer is measured over the profile's one stated fight
		// (leveling-bis's healEngine.forProfile does the same).
		req.Encounter.DurationSec = e.profile.DurationSec
		req.Encounter.Variation = 0
	}
	engineReq, err := request.BuildWeights(req, request.Options{OpenIterations: true})
	if err != nil {
		return measurement{}, fmt.Errorf("building the weights request: %w", err)
	}
	weapon := weaponPseudoStat(spec)
	engineReq.PseudoStatsToWeigh = []proto.PseudoStat{weapon}
	if healing(spec) {
		e.profile.AttachWeights(engineReq)
	}
	if err := simdb.AttachWeights(engineReq); err != nil {
		return measurement{}, fmt.Errorf("attaching the item database: %w", err)
	}
	res := core.StatWeights(engineReq)
	if res.GetError() != nil && res.GetError().GetMessage() != "" {
		return measurement{}, errors.New(res.GetError().GetMessage())
	}
	values := res.GetDps()
	if healing(spec) {
		values = res.GetHps()
	}
	perPoint, stdErr, err := referenceDerivative(values, req)
	if err != nil {
		return measurement{}, err
	}
	w, wErr := pseudoDerivative(values, weapon, req)
	return measurement{PerPoint: perPoint, StdErr: stdErr, Weapon: w, WeaponErr: wErr}, nil
}

// pseudoDerivative is referenceDerivative for a pseudo stat (a weapon's
// damage per second); 0 when the engine skipped it, which it does for a
// pseudo stat that moves nothing (a weapon nobody swings).
func pseudoDerivative(values *proto.StatWeightValues, ps proto.PseudoStat, req api.SimRequest) (float64, float64) {
	raw := values.GetWeights().GetPseudoStats()
	stdev := values.GetWeightsStdev().GetPseudoStats()
	if int(ps) >= len(raw) || int(ps) >= len(stdev) {
		return 0, 0
	}
	n := float64(req.Iterations * api.WeightsIterationsFactor)
	return raw[ps], stdev[ps] / math.Sqrt(n)
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
// run keeps every spec it did not measure, and every band of a spec it
// did measure that this run did not), one spec per line, keys sorted,
// figures rounded to decimals.
//
// Rounding to four decimals writes an error below 0.00005 as 0: a
// warrior's paired-stream sweep reads "dps_error": 0 for that reason,
// not because it has no error. And Data.lua's heap stood at 1,992.3 KB
// under Lua 5.1 after task L1.5, 55.7 KB under the addon's 2 MB budget:
// the next figure added here needs a size decision first.
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
		// A band this run did not measure keeps its existing figures.
		bands := map[string]bandEntry{}
		for band, b := range merged[k].Bands {
			bands[band] = b
		}
		for band, b := range v.Bands {
			bands[band] = b
		}
		if len(bands) == 0 {
			bands = nil
		}
		v.Bands = bands
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
	round := func(v float64) float64 { return math.Round(v*scale) / scale }
	for i, k := range keys {
		e := merged[k]
		e.DPS = round(e.DPS)
		e.WeaponDPS = round(e.WeaponDPS)
		if e.Bands != nil {
			bands := make(map[string]bandEntry, len(e.Bands))
			for band, b := range e.Bands {
				b.DPS, b.DPSError = round(b.DPS), round(b.DPSError)
				b.WeaponDPS, b.WeaponError = round(b.WeaponDPS), round(b.WeaponError)
				bands[band] = b
			}
			e.Bands = bands
		}
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
