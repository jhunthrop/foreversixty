package main

// The /sim page names every action in a result from the build's client tables (spells.json
// and items.json, published in shards by web/scripts/sync-data.mjs) and, for the ids the
// engine reports that the client lacks, data/curated/engine-spell-names.json. These tests
// run every published BiS entry of every spec, as the page sims it, and fail on an action
// the page would show as "An unnamed spell" or "An unnamed item", or in the wrong school.

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// resultIterations is enough for every rotation spell and the common procs to land; the
// names do not depend on the figures.
const resultIterations = 20

// The combat log's school bits (web/src/lib/report/format.ts schoolName), which a summary
// row carries.
const (
	logSchoolFire = 4
	igniteSpellID = 12654
)

var logSchoolOf = map[proto.SpellSchool]int64{
	proto.SpellSchool_SpellSchoolPhysical: 1,
	proto.SpellSchool_SpellSchoolHoly:     2,
	proto.SpellSchool_SpellSchoolFire:     4,
	proto.SpellSchool_SpellSchoolNature:   8,
	proto.SpellSchool_SpellSchoolFrost:    16,
	proto.SpellSchool_SpellSchoolShadow:   32,
	proto.SpellSchool_SpellSchoolArcane:   64,
}

// actionKey is the adapter's name grammar for a spell or an item row
// (sim/adapter/identity.go): "spell:<id>" or "item:<id>", with optional tag and rank.
var actionKey = regexp.MustCompile(`^(spell|item):(\d+)`)

// publishedRun is one published entry's request and its summary.
type publishedRun struct {
	label   string
	request api.SimRequest
	engine  *proto.RaidSimRequest
	summary summary.Summary
}

// publishedRuns holds the runs both tests read, simmed by the first to ask.
var publishedRuns []publishedRun

// runPublishedEntries sims every entry of every published BiS file once per test binary.
func runPublishedEntries(t *testing.T) []publishedRun {
	t.Helper()
	if testing.Short() {
		t.Skip("sims every published entry")
	}
	if publishedRuns == nil {
		publishedRuns = simAllPublished(t)
	}
	return publishedRuns
}

func simAllPublished(t *testing.T) []publishedRun {
	t.Helper()
	inproc.Register()
	dir, err := options{repoRoot: realRoot}.bisDirectory()
	if err != nil {
		t.Fatal(err)
	}
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil || len(files) < 27 {
		t.Fatalf("published BiS files in %s: %d (%v)", dir, len(files), err)
	}
	var runs []publishedRun
	for _, file := range files {
		spec := filepath.Base(file[:len(file)-len(".json")])
		for _, entry := range publishedEntries(t, file) {
			o := options{repoRoot: realRoot, spec: spec, preset: entry.Preset, band: entry.Band, faction: entry.Faction, iterations: resultIterations, seed: 1}
			runs = append(runs, simPublished(t, o))
		}
	}
	return runs
}

func publishedEntries(t *testing.T, file string) []bisEntry {
	t.Helper()
	raw, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	var parsed struct {
		Bands []bisEntry `json:"bands"`
	}
	if err := json.Unmarshal(raw, &parsed); err != nil {
		t.Fatalf("%s: %v", file, err)
	}
	return parsed.Bands
}

func simPublished(t *testing.T, o options) publishedRun {
	t.Helper()
	label := fmt.Sprintf("%s band %d %s %s", o.spec, o.band, o.preset, o.faction)
	req, err := buildRequest(o)
	if err != nil {
		t.Fatalf("%s: %v", label, err)
	}
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		t.Fatalf("%s: %v", label, err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		t.Fatal(err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimConcurrentAsync(engineReq, reporter, label)
	res := simdrain.ToResult(reporter, nil)
	if err := adapter.ResultError(res); err != nil {
		t.Fatalf("%s: %v", label, err)
	}
	sum, err := adapter.Summarize(res, req)
	if err != nil {
		t.Fatalf("%s: %v", label, err)
	}
	return publishedRun{label: label, request: req, engine: engineReq, summary: sum}
}

// resultKeys is every spell or item key a summary names: its damage and healing rows,
// auras and casts, the rows the page resolves.
func resultKeys(s summary.Summary) []string {
	var keys []string
	for _, actors := range [][]summary.Actor{s.DamageDone, s.Healing} {
		for _, actor := range actors {
			for _, ability := range actor.Abilities {
				keys = append(keys, ability.Name)
			}
		}
	}
	for _, aura := range s.Auras {
		keys = append(keys, aura.Name)
	}
	for _, cast := range s.Casts {
		keys = append(keys, cast.SpellName)
	}
	return keys
}

// clientNames is the name source the page publishes: the client's spells and items, and
// the curated names of the engine's own spell ids.
type clientNames struct {
	spell map[string]bool
	item  map[string]bool
}

func loadClientNames(t *testing.T) clientNames {
	t.Helper()
	spells, err := loadSpellNames(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	names := clientNames{spell: map[string]bool{}, item: map[string]bool{}}
	for id := range spells {
		names.spell[strconv.Itoa(int(id))] = true
	}
	var curated struct {
		Spells map[string]struct {
			Name string `json:"name"`
		} `json:"spells"`
	}
	readJSON(t, filepath.Join(realRoot, "data", "curated", "engine-spell-names.json"), &curated)
	for id, row := range curated.Spells {
		if row.Name == "" {
			t.Errorf("engine-spell-names.json: %s has no name", id)
		}
		names.spell[id] = true
	}
	dir, err := buildDirectory(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	var items []struct {
		ID int `json:"id"`
	}
	readJSON(t, filepath.Join(dir, "items.json"), &items)
	for _, row := range items {
		names.item[strconv.Itoa(row.ID)] = true
	}
	return names
}

func readJSON(t *testing.T, path string, into any) {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, into); err != nil {
		t.Fatalf("%s: %v", path, err)
	}
}

// Round 7 (owner ruling 2): /sim's Damage tab showed 36 rows as "An unnamed spell". Every
// spell and item a published spec's result carries has a name the page can read.
func TestEveryPublishedResultActionHasAName(t *testing.T) {
	runs := runPublishedEntries(t)
	names := loadClientNames(t)
	unnamed := map[string][]string{}
	for _, run := range runs {
		for _, key := range resultKeys(run.summary) {
			m := actionKey.FindStringSubmatch(key)
			if m == nil {
				continue
			}
			table := names.spell
			if m[1] == "item" {
				table = names.item
			}
			if !table[m[2]] {
				unnamed[m[1]+":"+m[2]] = append(unnamed[m[1]+":"+m[2]], run.label)
			}
		}
	}
	keys := make([]string, 0, len(unnamed))
	for key := range unnamed {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		t.Errorf("%s has no name (first in %s, %d runs); name it in data/curated/engine-spell-names.json", key, unnamed[key][0], len(unnamed[key]))
	}
}

// Round 7 (owner ruling 3): Ignite and the wand's Shoot read "Physical" on the Damage tab.
// Ignite burns Fire, and Shoot deals the wand's own school.
func TestPublishedResultsCarryIgniteAndShootSchools(t *testing.T) {
	runs := runPublishedEntries(t)
	sawIgnite, sawShoot := false, false
	for _, run := range runs {
		wandSchool, hasWand := wandLogSchool(t, run.engine)
		for _, actor := range run.summary.DamageDone {
			for _, ability := range actor.Abilities {
				m := actionKey.FindStringSubmatch(ability.Name)
				if m == nil || m[1] != "spell" {
					continue
				}
				switch m[2] {
				case strconv.Itoa(igniteSpellID):
					sawIgnite = true
					if ability.School != logSchoolFire {
						t.Errorf("%s: Ignite school %d, want Fire (%d)", run.label, ability.School, logSchoolFire)
					}
				case strconv.Itoa(core.ShootSpellID):
					sawShoot = true
					if !hasWand || ability.School != wandSchool {
						t.Errorf("%s: Shoot school %d, want the wand's %d", run.label, ability.School, wandSchool)
					}
				}
			}
		}
	}
	if !sawIgnite || !sawShoot {
		t.Errorf("no published result carried Ignite (%v) or Shoot (%v)", sawIgnite, sawShoot)
	}
}

// wandLogSchool is the log school of the player's equipped wand, from the item database the
// request carries.
func wandLogSchool(t *testing.T, req *proto.RaidSimRequest) (int64, bool) {
	t.Helper()
	player := req.GetRaid().GetParties()[0].GetPlayers()[0]
	for _, item := range player.GetEquipment().GetItems() {
		for _, row := range player.GetDatabase().GetItems() {
			if row.GetId() == item.GetId() && row.GetRangedWeaponType() == proto.RangedWeaponType_RangedWeaponTypeWand {
				return logSchoolOf[row.GetDamageSchool()], true
			}
		}
	}
	return 0, false
}
