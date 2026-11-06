package main

// The race sweep (sim-accuracy round 25). A band's published set is
// measured on the guide's race for its faction; the sweep runs the same
// gear, talents, count and seed on every race the class can be in that
// faction, from the client's race-class combinations (combos.json, Forever's
// added Skyborne, Human hunter, Troll mage and Gnome shaman included), and
// publishes the best race, its margin over the guide race and the margin's
// error next to the pick. DPS specs are judged on damage, healers on the heal
// score the ranker ranks by, tanks on threat (owner rule).
//
// The standalone -race-sweep mode runs the same sweep over a frozen copy of
// the published reports without ranking any gear (runRaceSweep).

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
)

// The units a race sweep's scores are in.
const (
	raceScoreDPS    = "dps"
	raceScoreHeal   = "heal_score"
	raceScoreThreat = "threat_per_second"
)

// raceScore is one race's measurement of the band's set.
type raceScore struct {
	Race  string  `json:"race"`
	Score float64 `json:"score"`
	Error float64 `json:"error"`
}

// raceSweep is the published sweep of one band and faction. Margin is the
// best race's score minus the guide race's, MarginPct the same over the
// guide race's score, Error the margin's combined standard error, and
// BeyondError whether the margin exceeds it.
type raceSweep struct {
	GuideRace   string      `json:"guide_race"`
	BestRace    string      `json:"best_race"`
	Margin      float64     `json:"margin"`
	MarginPct   float64     `json:"margin_pct"`
	Error       float64     `json:"error"`
	BeyondError bool        `json:"beyond_error"`
	ScoreUnit   string      `json:"score_unit"`
	Iterations  int         `json:"iterations"`
	Seed        int64       `json:"seed"`
	Races       []raceScore `json:"races"`
}

type comboFile struct {
	RaceID  int `json:"race_id"`
	ClassID int `json:"class_id"`
}

type idSlug struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Faction string `json:"faction"`
}

func readJSONFile(path string, into any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}
	if err := json.Unmarshal(b, into); err != nil {
		return fmt.Errorf("decoding %s: %w", path, err)
	}
	return nil
}

// legalRaces is every race classSlug can be in faction, in races.json order,
// read from the build's client race-class combinations.
func legalRaces(buildDir, classSlug, faction string) ([]string, error) {
	var races, classes []idSlug
	var combos []comboFile
	for name, into := range map[string]any{"races.json": &races, "classes.json": &classes, "combos.json": &combos} {
		if err := readJSONFile(filepath.Join(buildDir, name), into); err != nil {
			return nil, err
		}
	}
	classID := 0
	for _, c := range classes {
		if c.Slug == classSlug {
			classID = c.ID
		}
	}
	if classID == 0 {
		return nil, fmt.Errorf("class %q is not in %s/classes.json", classSlug, buildDir)
	}
	allowed := make(map[int]bool)
	for _, c := range combos {
		if c.ClassID == classID {
			allowed[c.RaceID] = true
		}
	}
	var out []string
	for _, r := range races {
		if allowed[r.ID] && r.Faction == faction {
			out = append(out, r.Slug)
		}
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no %s race can be a %s in %s/combos.json", faction, classSlug, buildDir)
	}
	return out, nil
}

// requestIterationsFor is what to ask spec's runner for so a run lands at
// target iterations: a healer's runner multiplies every request by
// healIterationScale, so a healer asks for the quotient, rounded up. The
// second return is the count the run really lands at.
func requestIterationsFor(spec specInfo, target int) (asked, effective int) {
	if isHealer(spec) {
		asked = (target + healIterationScale - 1) / healIterationScale
		return asked, asked * healIterationScale
	}
	return target, target
}

// raceScorer measures one request in the unit the sweep judges spec by.
type raceScorer func(req api.SimRequest) (mean, stdErr float64, err error)

func raceScorerFor(runner engineRunner, spec specInfo) (raceScorer, string, error) {
	switch {
	case spec.Role == roleTank:
		tank, ok := unwrapRunner(runner).(tankRunner)
		if !ok {
			return nil, "", fmt.Errorf("%s is a tank but its runner is not a tank runner", spec.Spec)
		}
		return func(req api.SimRequest) (float64, float64, error) {
			f, err := tank.RunTankFigures(req)
			return f.TPS, f.TPSError, err
		}, raceScoreThreat, nil
	case isHealer(spec):
		return runner.RunPlainDPSWithError, raceScoreHeal, nil
	default:
		return runner.RunPlainDPSWithError, raceScoreDPS, nil
	}
}

// sweepRaces measures gear and talents at level on every race in races, at
// iterations (the effective count) on seed, and names the best one against
// guideRace, which must be among them.
func sweepRaces(runner engineRunner, spec specInfo, guideRace string, races []string, level int, talents string, gear []api.GearSlot, iterations int, seed int64) (raceSweep, error) {
	score, unit, err := raceScorerFor(runner, spec)
	if err != nil {
		return raceSweep{}, err
	}
	asked, effective := requestIterationsFor(spec, iterations)
	out := raceSweep{GuideRace: guideRace, ScoreUnit: unit, Iterations: effective, Seed: seed}
	guide := -1
	for _, race := range races {
		ch := bandCharacter("final-set", race, spec.ClassSlug, spec.Spec, level, talents, gear)
		mean, stdErr, err := score(plainRequest(spec, ch, asked, seed))
		if err != nil {
			return raceSweep{}, fmt.Errorf("race sweep %s band %d as %s: %w", spec.Spec, level, race, err)
		}
		if race == guideRace {
			guide = len(out.Races)
		}
		out.Races = append(out.Races, raceScore{Race: race, Score: mean, Error: stdErr})
	}
	if guide < 0 {
		return raceSweep{}, fmt.Errorf("race sweep %s: the guide race %q is not a legal race here (%v)", spec.Spec, guideRace, races)
	}
	g := out.Races[guide]
	best := g
	for _, r := range out.Races {
		if r.Score > best.Score {
			best = r
		}
	}
	out.BestRace = best.Race
	if best.Race != g.Race {
		out.Margin = best.Score - g.Score
		out.Error = math.Hypot(best.Error, g.Error)
		out.BeyondError = out.Margin > out.Error
		if g.Score != 0 {
			out.MarginPct = 100 * out.Margin / g.Score
		}
	}
	return out, nil
}

// raceSweepConfig is the standalone -race-sweep mode's inputs.
type raceSweepConfig struct {
	RepoRoot, BuildDir, BISDir, OutDir, Spec string
	Iterations                               int
	Seed                                     int64
}

// bandRaceSweep is one swept band entry of the standalone mode's output.
type bandRaceSweep struct {
	Band          int       `json:"band"`
	Preset        string    `json:"preset"`
	Faction       string    `json:"faction"`
	PublishedRace string    `json:"published_race"`
	Sweep         raceSweep `json:"sweep"`
}

// frozenReport is the part of a published report the standalone sweep reads.
type frozenReport struct {
	Bands []struct {
		Band    int    `json:"band"`
		Preset  string `json:"preset"`
		Faction string `json:"faction"`
		Race    string `json:"race"`
		Talents string `json:"talents"`
		Slots   []struct {
			Slot   string `json:"slot"`
			ItemID int    `json:"item_id"`
		} `json:"slots"`
	} `json:"bands"`
}

// runRaceSweep sweeps every band entry of cfg.BISDir/<spec>.json on its
// published gear and talents, judging the guide's current race line, and
// writes cfg.OutDir/<spec>.races.json. It ranks no gear.
func runRaceSweep(runner engineRunner, cfg raceSweepConfig, resolveLayout talentLayoutResolver) error {
	spec, err := loadSpec(cfg.RepoRoot, cfg.Spec)
	if err != nil {
		return err
	}
	runner, _, err = withHealerEngine(runner, cfg.RepoRoot, spec)
	if err != nil {
		return err
	}
	runner = roleRunner(runner, spec)
	raidPreset, err := resolveRaidPreset(cfg.RepoRoot, spec)
	if err != nil {
		return fmt.Errorf("resolving the %s preset for %s: %w", presetRaid, cfg.Spec, err)
	}
	if spec.RaidEnchants, err = loadRaidEnchants(cfg.RepoRoot, spec); err != nil {
		return err
	}
	guide, err := loadGuideRaces(cfg.RepoRoot, cfg.BuildDir, spec.ClassSlug, spec.SpecSlug)
	if err != nil {
		return err
	}
	layout, err := resolveLayout(cfg.RepoRoot, spec.ClassSlug)
	if err != nil {
		return err
	}
	trees, err := leveling.LoadTalentTrees(cfg.RepoRoot, filepath.Base(cfg.BuildDir), spec.ClassSlug)
	if err != nil {
		return err
	}
	var frozen frozenReport
	if err := readJSONFile(filepath.Join(cfg.BISDir, cfg.Spec+".json"), &frozen); err != nil {
		return err
	}
	guideRace := map[string]string{"alliance": guide.AllianceRace, "horde": guide.HordeRace}
	var out []bandRaceSweep
	for _, b := range frozen.Bands {
		pass := passSpec(spec, b.Preset, raidPreset)
		races, err := legalRaces(cfg.BuildDir, spec.ClassSlug, b.Faction)
		if err != nil {
			return err
		}
		talents, err := layout.Reposition(trees, b.Talents)
		if err != nil {
			return fmt.Errorf("band %d %s: converting talents to the engine's layout: %w", b.Band, b.Faction, err)
		}
		var gear []api.GearSlot
		for _, s := range b.Slots {
			if s.ItemID != 0 {
				gear = append(gear, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
			}
		}
		sweep, err := sweepRaces(runner, pass, guideRace[b.Faction], races, b.Band, talents, gear, cfg.Iterations, cfg.Seed)
		if err != nil {
			return fmt.Errorf("band %d %s %s: %w", b.Band, b.Preset, b.Faction, err)
		}
		log.Printf("leveling-bis: race sweep %s band %d %s %s: %s", cfg.Spec, b.Band, b.Preset, b.Faction, formatRaceSweep(sweep))
		out = append(out, bandRaceSweep{Band: b.Band, Preset: b.Preset, Faction: b.Faction, PublishedRace: b.Race, Sweep: sweep})
	}
	raw, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.OutDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(cfg.OutDir, cfg.Spec+".races.json"), append(raw, '\n'), 0o644)
}

// passSpec is spec as the named pass ranks it: the raid pass wears the raid
// preset, every other pass the class kit only (passesFor).
func passSpec(spec specInfo, preset string, raid request.ResolvedPreset) specInfo {
	spec.Applied = nil
	if preset == presetRaid {
		spec.Applied = &raid
	}
	return spec
}

// formatRaceSweep is the one-line log and markdown form of a sweep: every
// race's score, best first, then the verdict against the guide race.
func formatRaceSweep(s raceSweep) string {
	rows := append([]raceScore(nil), s.Races...)
	sort.SliceStable(rows, func(i, j int) bool { return rows[i].Score > rows[j].Score })
	line := ""
	for i, r := range rows {
		if i > 0 {
			line += ", "
		}
		line += fmt.Sprintf("%s %.1f ± %.1f", r.Race, r.Score, r.Error)
	}
	if s.BestRace == s.GuideRace {
		return fmt.Sprintf("%s (%s); the guide race %s is best", line, s.ScoreUnit, s.GuideRace)
	}
	verdict := "inside error"
	if s.BeyondError {
		verdict = "beyond error"
	}
	return fmt.Sprintf("%s (%s); %s beats the guide race %s by %.1f (%.2f%%) ± %.1f, %s", line, s.ScoreUnit, s.BestRace, s.GuideRace, s.Margin, s.MarginPct, s.Error, verdict)
}
