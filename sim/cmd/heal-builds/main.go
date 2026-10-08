// Command heal-builds sims talent builds for one healing spec against the
// heal profile and prints them side by side.
//
// The gear is a band-60 raid-ready entry of a BiS file (the ranker's
// output, committed or scratch), so only the talents vary. A build is a
// JSON object of talent name -> rank per label; it is checked against the
// tree rules a player is held to (tier gates, prerequisites, 51 points)
// before it is simmed, and every build runs the same seed over the same
// fight so the numbers are paired. This is how a healer guide's talent
// build is chosen on evidence instead of by recall: no healer has a
// talent search of its own (sim/cmd/talent-search skips healers).
//
//	go run ./sim/cmd/heal-builds -spec priest-holy -bis out/priest-holy.json -builds builds.json
package main

import (
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
	"text/tabwriter"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/enginetalents"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/specs"
)

const (
	pointsPerTier = 5
	// talentPoints is what a level-60 character spends.
	talentPoints = 51
	level        = 60
	seed         = 7
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

type options struct {
	repoRoot, spec, bisPath, buildsPath, faction string
	iterations                                   int
}

func run(args []string) error {
	var o options
	fs := flag.NewFlagSet("heal-builds", flag.ContinueOnError)
	fs.StringVar(&o.repoRoot, "repo-root", ".", "the site repository root")
	fs.StringVar(&o.spec, "spec", "", "healing spec, e.g. priest-holy (required)")
	fs.StringVar(&o.bisPath, "bis", "", "BiS file whose level-60 raid band supplies the gear (required)")
	fs.StringVar(&o.buildsPath, "builds", "", "JSON file: label -> {talent name: rank} (required)")
	fs.StringVar(&o.faction, "faction", "alliance", "faction whose band to wear")
	fs.IntVar(&o.iterations, "iterations", 1000, "iterations per build")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if o.spec == "" || o.bisPath == "" || o.buildsPath == "" {
		return errors.New("-spec, -bis and -builds are required")
	}
	return report(o)
}

type bisBand struct {
	Preset  string `json:"preset"`
	Band    int    `json:"band"`
	Faction string `json:"faction"`
	Race    string `json:"race"`
	Slots   []struct {
		Slot   string `json:"slot"`
		ItemID int    `json:"item_id"`
	} `json:"slots"`
}

func loadBand(path, faction string) (bisBand, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return bisBand{}, err
	}
	var file struct {
		Bands []bisBand `json:"bands"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		return bisBand{}, fmt.Errorf("decoding %s: %w", path, err)
	}
	for _, band := range file.Bands {
		if band.Band == level && band.Faction == faction && band.Preset == request.RaidPreset {
			return band, nil
		}
	}
	return bisBand{}, fmt.Errorf("%s has no level-%d %s raid band", path, level, faction)
}

func report(o options) error {
	spec, ok := specs.ByKey[o.spec]
	if !ok || spec.Role != "healer" {
		return fmt.Errorf("%q is not a healing spec", o.spec)
	}
	band, err := loadBand(o.bisPath, o.faction)
	if err != nil {
		return err
	}
	buildDir, err := activeBuildDir(o.repoRoot)
	if err != nil {
		return err
	}
	trees, err := leveling.LoadTalentTrees(o.repoRoot, filepath.Base(buildDir), spec.ClassSlug)
	if err != nil {
		return err
	}
	engineDir, err := enginetalents.SourceDir(filepath.Join(o.repoRoot, "sim"))
	if err != nil {
		return err
	}
	layout, err := enginetalents.ForClass(engineDir, spec.ClassSlug)
	if err != nil {
		return err
	}
	profile, err := request.LoadHealProfile(filepath.Join(o.repoRoot, "data", "curated", "heal-profile.json"))
	if err != nil {
		return err
	}
	buffs, consumes, err := raidLoadout(o.repoRoot, spec)
	if err != nil {
		return err
	}
	builds, err := loadBuilds(o.buildsPath)
	if err != nil {
		return err
	}

	var rows []row
	for _, label := range sortedKeys(builds) {
		ranks, err := ranksByID(trees, builds[label])
		if err != nil {
			return fmt.Errorf("build %q: %w", label, err)
		}
		if err := legal(trees, ranks); err != nil {
			return fmt.Errorf("build %q: %w", label, err)
		}
		site := siteString(trees, ranks)
		engine, err := layout.Reposition(trees, site)
		if err != nil {
			return fmt.Errorf("build %q: %w", label, err)
		}
		result, err := inproc.HealingRun(buildRequest(o, spec, band, engine, buffs, consumes), profile)
		if err != nil {
			return fmt.Errorf("build %q: %w", label, err)
		}
		rows = append(rows, row{label: label, site: site, engine: engine, result: result, fight: profile.Duration().Seconds()})
	}
	printRows(rows)
	return nil
}

func buildRequest(o options, spec specs.Spec, band bisBand, talents string, buffs, consumes []string) api.SimRequest {
	gear := make([]api.GearSlot, 0, len(band.Slots))
	for _, s := range band.Slots {
		if s.ItemID != 0 {
			gear = append(gear, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
		}
	}
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          spec.Spec,
		Source:        api.CharacterSource{Kind: api.SourceBuild},
		Character: api.CharacterSpec{
			Name: "heal-builds", Race: band.Race, Class: spec.ClassSlug, Level: level,
			Talents: talents, Gear: gear, Buffs: buffs, Consumes: consumes,
			DistanceFromTarget: leveling.CasterDistanceFromTarget,
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: o.iterations,
		RandomSeed: seed,
	}
}

// raidLoadout is the raid preset's buffs and consumables for spec with
// the class kit on top, as the ranker builds a raid-ready character.
func raidLoadout(repoRoot string, spec specs.Spec) (buffs, consumes []string, err error) {
	presets, err := request.LoadPresets(filepath.Join(repoRoot, "data", "curated", "presets.json"))
	if err != nil {
		return nil, nil, err
	}
	resolved, err := presets.Resolve(request.RaidPreset, spec)
	if err != nil {
		return nil, nil, err
	}
	buffs, consumes = resolved.Layer(leveling.KitBuffs(spec.Spec, level), leveling.KitConsumes(spec.Spec, level))
	return buffs, consumes, nil
}

func activeBuildDir(repoRoot string) (string, error) {
	build, err := leveling.ReadActiveBuild(repoRoot)
	if err != nil {
		return "", err
	}
	return filepath.Join(repoRoot, "data", "builds", build), nil
}

func loadBuilds(path string) (map[string]map[string]int, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var builds map[string]map[string]int
	if err := json.Unmarshal(raw, &builds); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", path, err)
	}
	return builds, nil
}

func sortedKeys(m map[string]map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ranksByID resolves a build's talent names to stable talent ids.
func ranksByID(trees []leveling.TalentTree, byName map[string]int) (map[int]int, error) {
	index := map[string]leveling.TalentNode{}
	for _, tree := range trees {
		for _, node := range tree.Talents {
			index[strings.ToLower(node.Name)] = node
		}
	}
	out := map[int]int{}
	for name, rank := range byName {
		node, ok := index[strings.ToLower(name)]
		if !ok {
			return nil, fmt.Errorf("no talent named %q in this class", name)
		}
		out[node.ID] = rank
	}
	return out, nil
}

// legal holds ranks to the rules a player is held to: ranks within the
// maximum, each tier unlocked by five points per tier beneath it in its
// tree, every prerequisite met, and exactly talentPoints spent.
func legal(trees []leveling.TalentTree, ranks map[int]int) error {
	return legalFor(trees, ranks, talentPoints)
}

// legalFor is legal for a build of total points.
func legalFor(trees []leveling.TalentTree, ranks map[int]int, total int) error {
	treeOf := map[int]int{}
	nodes := map[int]leveling.TalentNode{}
	for ti, tree := range trees {
		for _, node := range tree.Talents {
			treeOf[node.ID] = ti
			nodes[node.ID] = node
		}
	}
	below := func(ti, tier int) int {
		n := 0
		for _, node := range trees[ti].Talents {
			if node.Tier < tier {
				n += ranks[node.ID]
			}
		}
		return n
	}
	spent := 0
	for id, rank := range ranks {
		node := nodes[id]
		spent += rank
		switch {
		case rank < 0 || rank > node.MaxRank:
			return fmt.Errorf("%s has %d of %d points", node.Name, rank, node.MaxRank)
		case rank == 0:
		case below(treeOf[id], node.Tier) < pointsPerTier*node.Tier:
			return fmt.Errorf("%s (tier %d) needs %d points above it", node.Name, node.Tier, pointsPerTier*node.Tier)
		case node.PrereqTalentID != 0 && ranks[node.PrereqTalentID] < node.PrereqRank:
			return fmt.Errorf("%s needs %d points in %s", node.Name, node.PrereqRank, nodes[node.PrereqTalentID].Name)
		}
	}
	if spent != total {
		return fmt.Errorf("spends %d points, want %d", spent, total)
	}
	return nil
}

// siteString is the build in the active build's positional order, one
// digit run per tree joined with "-", the layout Reposition reads.
func siteString(trees []leveling.TalentTree, ranks map[int]int) string {
	parts := make([]string, len(trees))
	for ti, tree := range trees {
		var sb strings.Builder
		for _, node := range tree.Talents {
			sb.WriteByte(byte('0' + ranks[node.ID]))
		}
		parts[ti] = sb.String()
	}
	return strings.Join(parts, "-")
}

type row struct {
	label, site, engine string
	result              inproc.HealingResult
	fight               float64
}

// score is the ranker's: effective healing per second scaled by the
// squared share of the fight the mana lasted.
func (r row) score() float64 {
	share := math.Min(1, r.result.ManaLastsSec/r.fight)
	return r.result.Effective.Mean * share * share
}

func printRows(rows []row) {
	sort.Slice(rows, func(i, j int) bool { return rows[i].score() > rows[j].score() })
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "build\tscore\tHPS\t+/-\traw\toverheal\tmana lasts\tHPM\tsite talents\tengine talents")
	for _, r := range rows {
		fmt.Fprintf(w, "%s\t%.1f\t%.1f\t%.1f\t%.1f\t%.1f%%\t%.0f s\t%.2f\t%s\t%s\n",
			r.label, r.score(), r.result.Effective.Mean, r.result.Effective.Error, r.result.Raw.Mean,
			100*r.result.OverhealShare(), r.result.ManaLastsSec, r.result.HealingPerMana(), r.site, r.engine)
	}
	w.Flush()
}
