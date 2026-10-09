package main

import (
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	testSpec    = "rogue-combat"
	testSpellID = 1752
)

func spellID(id int32, tag int32) *proto.ActionID {
	return &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: id}, Tag: tag}
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-6 }

func TestActionName(t *testing.T) {
	names := map[int32]string{testSpellID: "Sinister Strike"}
	tests := []struct {
		name string
		id   *proto.ActionID
		want string
	}{
		{"nil", nil, "?"},
		{"empty raw id", &proto.ActionID{}, "?"},
		{"named spell", spellID(testSpellID, 0), "Sinister Strike (1752)"},
		{"tagged spell", spellID(testSpellID, 3), "Sinister Strike (1752)#3"},
		{"unnamed spell", spellID(99, 0), "spell (99)"},
		{"item", &proto.ActionID{RawId: &proto.ActionID_ItemId{ItemId: 5}, Tag: 2}, "item 5#2"},
		{"other", &proto.ActionID{RawId: &proto.ActionID_OtherId{OtherId: proto.OtherAction_OtherActionWait}}, "OtherActionWait"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := actionName(tc.id, names); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestFoldActionComputesRatesAndOutcomePercentages(t *testing.T) {
	a := &proto.ActionMetrics{
		Id: spellID(testSpellID, 0),
		Targets: []*proto.TargetedActionMetrics{{
			Casts: 12, Hits: 8, Crits: 2, Misses: 1, Dodges: 1, Damage: 1800, GlanceDamage: 0,
		}},
	}
	row := foldAction(a, 2, 10, 180, map[int32]string{testSpellID: "Sinister Strike"})
	checks := map[string][2]float64{
		"DPS":          {row.DPS, 90},
		"Share":        {row.Share, 50},
		"CastsPerIter": {row.CastsPerIter, 6},
		"HitsPerIter":  {row.HitsPerIter, 5},
		"AvgHit":       {row.AvgHit, 180},
		"CritPct":      {row.CritPct, 100.0 * 2 / 12},
		"MissPct":      {row.MissPct, 100.0 / 12},
		"DodgePct":     {row.DodgePct, 100.0 / 12},
	}
	for field, c := range checks {
		if !near(c[0], c[1]) {
			t.Errorf("%s = %v, want %v", field, c[0], c[1])
		}
	}
	if row.Name != "Sinister Strike (1752)" || row.Passive {
		t.Errorf("name/passive = %q/%v", row.Name, row.Passive)
	}
}

func TestFoldActionCountsTicksAndGlancesAndSurvivesNoTargets(t *testing.T) {
	dot := &proto.ActionMetrics{Targets: []*proto.TargetedActionMetrics{
		{Ticks: 4, CritTicks: 1, Glances: 1, Damage: 500, GlanceDamage: 50},
		{Hits: 1, Blocks: 1, Parries: 1},
	}, IsPassive: true}
	row := foldAction(dot, 1, 10, 0, nil)
	if row.Share != 0 {
		t.Errorf("share with zero total = %v, want 0", row.Share)
	}
	if !near(row.GlanceDPS, 5) || !near(row.HitsPerIter, 7) || !row.Passive {
		t.Errorf("glance dps %v hits/iter %v passive %v", row.GlanceDPS, row.HitsPerIter, row.Passive)
	}
	empty := foldAction(&proto.ActionMetrics{}, 1, 10, 100, nil)
	if empty.AvgHit != 0 || empty.CritPct != 0 || empty.Name != "?" {
		t.Errorf("empty action row = %+v", empty)
	}
}

func sampleMetrics() *proto.UnitMetrics {
	return &proto.UnitMetrics{
		Actions: []*proto.ActionMetrics{
			{Id: spellID(1, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 2, Hits: 2, Damage: 200}}},
			{Id: spellID(2, 0), Targets: []*proto.TargetedActionMetrics{{Casts: 2, Hits: 2, Damage: 800}}},
			{Id: spellID(3, 0)},
		},
		Auras: []*proto.AuraMetrics{
			{Id: spellID(10, 0), UptimeSecondsAvg: 2, ProcsAvg: 1},
			{Id: spellID(11, 0), UptimeSecondsAvg: 5, ProcsAvg: 2},
			{Id: spellID(12, 0)},
		},
		Resources: []*proto.ResourceMetrics{
			{Id: spellID(20, 0), Type: proto.ResourceType_ResourceTypeEnergy, Gain: 100, ActualGain: 80},
		},
	}
}

func sampleReport() report {
	req := api.SimRequest{Spec: testSpec, Iterations: 2, Encounter: api.EncounterSpec{DurationSec: 10},
		Character: api.CharacterSpec{Consumes: []string{"a"}, Buffs: []string{"b"}}}
	names := map[int32]string{1: "Low", 2: "High", 10: "Short", 11: "Long", 20: "Combo"}
	return newReport(req, api.Estimate{Mean: 50, Error: 1.5}, sampleMetrics(), names)
}

func TestNewReportSortsAndScalesRows(t *testing.T) {
	r := sampleReport()
	if r.Spec != testSpec || r.DPS != 50 || r.Error != 1.5 {
		t.Errorf("header = %s %v %v", r.Spec, r.DPS, r.Error)
	}
	if !strings.HasPrefix(r.Actions[0].Name, "High") || !strings.HasPrefix(r.Actions[1].Name, "Low") {
		t.Errorf("actions not sorted by DPS: %+v", r.Actions)
	}
	if !near(r.Actions[0].DPS, 40) || !near(r.Actions[0].Share, 80) {
		t.Errorf("top action dps/share = %v/%v", r.Actions[0].DPS, r.Actions[0].Share)
	}
	if !strings.HasPrefix(r.Auras[0].Name, "Long") || !near(r.Auras[0].UptimePct, 50) {
		t.Errorf("auras not sorted by uptime: %+v", r.Auras)
	}
	res := r.Resources[0]
	if res.Type != "ResourceTypeEnergy" || !near(res.GainPerIter, 50) || !near(res.ActualPerIter, 40) {
		t.Errorf("resource row = %+v", res)
	}
}

func TestMarkdownRendersTablesAndSkipsEmptyRows(t *testing.T) {
	out := sampleReport().markdown()
	for _, want := range []string{
		testSpec + "  DPS 50.0 +/- 1.5", "consumes: [a]", "buffs: [b]",
		"| High (2) | 40.0 | 80.0 |", "| Long (11) | 50.0 | 2.00 |", "| Combo (20) | ResourceTypeEnergy | 50.0 | 40.0 |",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("markdown missing %q in:\n%s", want, out)
		}
	}
	for _, absent := range []string{"spell (3)", "spell (12)"} {
		if strings.Contains(out, absent) {
			t.Errorf("markdown kept the empty row %q", absent)
		}
	}
}

func TestLoadRotation(t *testing.T) {
	dir := t.TempDir()
	valid := filepath.Join(dir, "valid.json")
	writeFile(t, valid, `{"rotation":{"priorityList":[{"action":{"castSpell":{"spellId":{"spellId":1752}}}}]}}`)
	missing := filepath.Join(dir, "missing.json")
	writeFile(t, missing, `{"other":{}}`)
	notJSON := filepath.Join(dir, "bad.json")
	writeFile(t, notJSON, `{not json`)
	badProto := filepath.Join(dir, "badproto.json")
	writeFile(t, badProto, `{"rotation":{"noSuchField":1}}`)

	rot, err := loadRotation(valid)
	if err != nil || len(rot.GetPriorityList()) != 1 {
		t.Fatalf("valid rotation: %v, %v", rot, err)
	}
	if rot, err := loadRotation(""); rot != nil || err != nil {
		t.Errorf("empty path = %v, %v; want nil, nil", rot, err)
	}
	tests := []struct{ name, path, wantErr string }{
		{"no such file", filepath.Join(dir, "nope.json"), "no such file"},
		{"not json", notJSON, "decoding"},
		{"no rotation object", missing, `no "rotation" object`},
		{"invalid APL", badProto, "decoding the rotation"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := loadRotation(tc.path); err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("err = %v, want containing %q", err, tc.wantErr)
			}
		})
	}
}

// testBuildDir is the build directory a test root's active-build.json names.
const testBuildDir = "data/builds/9.9.9.9"

// activeRoot is a temp repo root whose active build is testBuildDir's.
func activeRoot(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "web", "src", "data", "active-build.json"), `{"build":"9.9.9.9"}`)
	return root
}

func TestLoadSpellNames(t *testing.T) {
	if _, err := loadSpellNames(t.TempDir()); err == nil {
		t.Error("a root with no active-build.json should fail")
	}
	root := activeRoot(t)
	if _, err := loadSpellNames(root); err == nil {
		t.Error("missing spells.json should fail")
	}
	path := filepath.Join(root, testBuildDir, "spells.json")
	writeFile(t, path, `[{"id":1752,"name":"Sinister Strike"},{"id":2,"name":"Other"}]`)
	names, err := loadSpellNames(root)
	if err != nil || len(names) != 2 || names[1752] != "Sinister Strike" {
		t.Fatalf("names = %v, err = %v", names, err)
	}
	writeFile(t, path, `{"not":"a list"}`)
	if _, err := loadSpellNames(root); err == nil {
		t.Error("malformed spells.json should fail")
	}
}

func bisFixture(t *testing.T) string {
	t.Helper()
	root := activeRoot(t)
	writeFile(t, filepath.Join(root, testBuildDir, "bis", testSpec+".json"), `{"bands":[
	  {"band":60,"preset":"bare","faction":"horde","race":"orc","talents":"other","slots":[]},
	  {"band":60,"preset":"bare","faction":"alliance","race":"human","talents":"0023",
	   "slots":[{"slot":"head","item_id":12},{"slot":"neck","item_id":34}]}]}`)
	return root
}

func baseOptions(root string) options {
	return options{repoRoot: root, spec: testSpec, preset: "bare", band: 60, faction: "alliance", iterations: 7, seed: 3}
}

func TestLoadEntrySelectsByBandPresetAndFaction(t *testing.T) {
	root := bisFixture(t)
	e, err := loadEntry(baseOptions(root))
	if err != nil || e.Race != "human" || len(e.Slots) != 2 {
		t.Fatalf("entry = %+v, err = %v", e, err)
	}
	o := baseOptions(root)
	o.band = 50
	if _, err := loadEntry(o); err == nil || !strings.Contains(err.Error(), "no rogue-combat band 50") {
		t.Errorf("missing band err = %v", err)
	}
	o = baseOptions(t.TempDir())
	if _, err := loadEntry(o); err == nil {
		t.Error("missing bis file should fail")
	}
}

func TestBuildRequestLayersEntryAndOverrides(t *testing.T) {
	root := bisFixture(t)
	req, err := buildRequest(baseOptions(root))
	if err != nil {
		t.Fatal(err)
	}
	c := req.Character
	if c.Race != "human" || c.Talents != "0023" || c.Class != "rogue" || c.Level != 60 || len(c.Gear) != 2 {
		t.Errorf("character = %+v", c)
	}
	if c.Gear[1] != (api.GearSlot{Slot: "neck", ItemID: 34}) || req.Iterations != 7 || req.RandomSeed != 3 {
		t.Errorf("gear/iterations/seed = %+v %d %d", c.Gear, req.Iterations, req.RandomSeed)
	}

	o := baseOptions(root)
	o.talents, o.consumes, o.buffs = "9999", "x,y", "extra"
	over, err := buildRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	if over.Character.Talents != "9999" || strings.Join(over.Character.Consumes, ",") != "x,y" {
		t.Errorf("overrides not applied: %+v", over.Character)
	}
	if got := over.Character.Buffs; got[len(got)-1] != "extra" || len(got) != len(c.Buffs)+1 {
		t.Errorf("buffs = %v, want kit buffs plus extra", got)
	}
}

func TestBuildRequestRejectsUnknownSpecAndMissingEntry(t *testing.T) {
	root := activeRoot(t)
	writeFile(t, filepath.Join(root, testBuildDir, "bis", "nope.json"),
		`{"bands":[{"band":60,"preset":"bare","faction":"alliance"}]}`)
	o := baseOptions(root)
	o.spec = "nope"
	if _, err := buildRequest(o); err == nil || !strings.Contains(err.Error(), `unknown spec "nope"`) {
		t.Errorf("unknown spec err = %v", err)
	}
	if _, err := buildRequest(baseOptions(root)); err == nil {
		t.Error("missing entry should fail")
	}
}

func TestRunRequiresSpec(t *testing.T) {
	if err := run(options{}); err == nil || !strings.Contains(err.Error(), "-spec is required") {
		t.Errorf("err = %v", err)
	}
}
