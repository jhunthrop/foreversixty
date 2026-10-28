package coach

// Shared test fixtures for every coach test (Tasks 5–14). An F task adds helpers only in
// its own fN_fixtures_test.go; a helper two F tasks need belongs here.
//
// The base fight is logs/engine/summary/testdata/v22.summary.json.golden: a 39 s retail
// 12.0.7 log with no encounter, no kill and empty combatants, whose Bloodthirst is the
// retail spell (two casts). It is a SHAPE fixture only: its casts, auras, resources and
// death are real log shapes, and goldenFight sets the fields it lacks itself (Onyxia,
// difficulty 9, a kill, a combatant row per player). Nothing built from it is a
// known-good value. A real Bloodthirst count, a real flask and a real avoidable death
// come only from the owner-provided Forever boss pull (Global Constraints, Fixtures).

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	furyGUID   = "Player-1136-09362954" // Tyrim, the golden's Fury warrior (protInput makes him Protection)
	priestGUID = "Player-11-0E93D4E2"   // Skiimaster, a Discipline priest in the golden; the tests make him Holy

	goldenPath     = "../../../logs/engine/summary/testdata/v22.summary.json.golden"
	buildsDir      = "../../../data/builds"
	activeBuildRef = "../../../web/src/data/active-build.json"

	onyxiaEncounterID = 1084
	onyxiaName        = "Onyxia"
	raidDifficulty40  = 9
	fixtureLevel      = 60
	fixtureVerified   = "1.60.1.70630"
)

var (
	goldenOnce  sync.Once
	goldenBytes []byte
	goldenErr   error
)

// goldenFight is the v22 golden, fresh for each call, with the fields it lacks set: the
// Onyxia encounter, difficulty 9, a kill, and a combatant row with gear per roster player.
func goldenFight(t *testing.T) summary.Summary {
	t.Helper()
	goldenOnce.Do(func() { goldenBytes, goldenErr = os.ReadFile(goldenPath) })
	if goldenErr != nil {
		t.Fatal(goldenErr)
	}
	var s summary.Summary
	if err := json.Unmarshal(goldenBytes, &s); err != nil {
		t.Fatal(err)
	}
	s.EncounterID, s.Difficulty, s.Kill = onyxiaEncounterID, raidDifficulty40, true
	for _, r := range s.Roster {
		s = withCombatant(s, r.GUID, true)
	}
	return s
}

// withDuration is s lasting ms.
func withDuration(s summary.Summary, ms int64) summary.Summary {
	s.DurationMS = ms
	return s
}

// withoutCasts is s with every cast row of guid (and of anything guid owns) removed.
func withoutCasts(s summary.Summary, guid string) summary.Summary {
	s.Casts = slices.DeleteFunc(slices.Clone(s.Casts), func(c summary.CastRow) bool {
		return c.GUID == guid || c.OwnerGUID == guid
	})
	return s
}

// withCasts is s with guid's row for spellID replaced by one that succeeded at each seqMS.
func withCasts(s summary.Summary, guid string, spellID int64, name string, seqMS ...int64) summary.Summary {
	casts := slices.DeleteFunc(slices.Clone(s.Casts), func(c summary.CastRow) bool {
		return c.GUID == guid && c.SpellID == spellID
	})
	n := int64(len(seqMS))
	s.Casts = append(casts, summary.CastRow{GUID: guid, OwnerGUID: guid, SpellID: spellID, SpellName: name,
		Started: n, Succeeded: n, Sequence: slices.Clone(seqMS)})
	return s
}

// withResource is s with guid's track of powerType replaced: one reading per second.
func withResource(s summary.Summary, guid string, powerType, max int64, perSecond ...int64) summary.Summary {
	tracks := slices.DeleteFunc(slices.Clone(s.Resources), func(r summary.ResourceTrack) bool {
		return r.GUID == guid && r.PowerType == powerType
	})
	s.Resources = append(tracks, summary.ResourceTrack{GUID: guid, PowerType: powerType, Max: max, Series: slices.Clone(perSecond)})
	return s
}

// withAura is s with target's track of spellID replaced by segs, applied by appliers.
func withAura(s summary.Summary, target string, spellID int64, name string, appliers []string, segs ...summary.Segment) summary.Summary {
	tracks := slices.DeleteFunc(slices.Clone(s.Auras), func(a summary.AuraTrack) bool {
		return a.TargetGUID == target && a.SpellID == spellID
	})
	var up int64
	for _, sg := range segs {
		up += sg.EndMS - sg.StartMS
	}
	s.Auras = append(tracks, summary.AuraTrack{TargetGUID: target, SpellID: spellID, Name: name, Type: "BUFF",
		Applications: int64(len(segs)), UptimeMS: up, Segments: slices.Clone(segs), Appliers: slices.Clone(appliers)})
	return s
}

// withDeath is s with guid dying at atMS to killingSpellID (0 is a melee swing).
func withDeath(s summary.Summary, guid string, atMS, killingSpellID int64, killingName string) summary.Summary {
	blow := &summary.DamageRef{AtMS: atMS, SourceName: onyxiaName, SpellID: killingSpellID, SpellName: killingName}
	s.Deaths = append(slices.Clone(s.Deaths), summary.Death{GUID: guid, AtMS: atMS, KillingBlow: blow})
	s.Roster = slices.Clone(s.Roster)
	for i := range s.Roster {
		if s.Roster[i].GUID == guid {
			s.Roster[i].Deaths++
		}
	}
	return s
}

// withCombatant is s with guid's combatant row replaced: gear true is a character sheet
// (one enchanted main hand), false a row with no gear, which reads as no sheet.
func withCombatant(s summary.Summary, guid string, gear bool) summary.Summary {
	rows := slices.DeleteFunc(slices.Clone(s.Combatants), func(c summary.CombatantRow) bool { return c.GUID == guid })
	row := summary.CombatantRow{GUID: guid}
	if gear {
		row.Gear = []event.Item{{ID: 19364, ItemLevel: 83, Enchants: []int64{1900}}}
	}
	s.Combatants = append(rows, row)
	return s
}

// withRole is s with guid's roster row given class, spec and role, added if absent.
func withRole(s summary.Summary, guid, class, spec, role string) summary.Summary {
	return mapRosterRow(s, guid, func(r summary.RosterRow) summary.RosterRow {
		r.Class, r.Spec, r.Role = class, spec, role
		return r
	})
}

// withLevel is s with guid's roster row at level (0: the log has no advanced logging).
func withLevel(s summary.Summary, guid string, level int64) summary.Summary {
	return mapRosterRow(s, guid, func(r summary.RosterRow) summary.RosterRow {
		r.Level = level
		return r
	})
}

// withRoster is s with its roster replaced by rows.
func withRoster(s summary.Summary, rows ...summary.RosterRow) summary.Summary {
	s.Roster = slices.Clone(rows)
	return s
}

func mapRosterRow(s summary.Summary, guid string, f func(summary.RosterRow) summary.RosterRow) summary.Summary {
	rows := slices.Clone(s.Roster)
	for i := range rows {
		if rows[i].GUID == guid {
			rows[i] = f(rows[i])
			s.Roster = rows
			return s
		}
	}
	s.Roster = append(rows, f(summary.RosterRow{GUID: guid}))
	return s
}

// seeds is what a player brought into the pull: catalogue entries seeded at 0 ms.
type seeds []consumables.Entry

func seededAt(entries ...consumables.Entry) seeds { return entries }

// withSeeds is s with one aura track on guid per entry, one segment from the pull to the end.
func withSeeds(s summary.Summary, guid string, ss seeds) summary.Summary {
	for _, e := range ss {
		s = withAura(s, guid, e.SpellID, e.Name, []string{guid}, span(0, s.DurationMS))
	}
	return s
}

// flask is the catalogue's flask of that name, in any role; it panics naming an absent one.
func flask(name string) consumables.Entry {
	return catalogueEntry(consumables.Load().AllRoles().Flask, "flask", name)
}

// food is the catalogue's food of that name, in any role; it panics naming an absent one.
func food(name string) consumables.Entry {
	return catalogueEntry(consumables.Load().AllRoles().Food, "food", name)
}

func catalogueEntry(entries []consumables.Entry, category, name string) consumables.Entry {
	for _, e := range entries {
		if e.Name == name {
			return e
		}
	}
	panic("fixture: no " + category + " named " + name + " in the consumables catalogue")
}

// repeat is n copies of v: a flat resource series, one reading per second.
func repeat(v int64, n int) []int64 {
	out := make([]int64, n)
	for i := range out {
		out[i] = v
	}
	return out
}

// seqEvery is n instants from fromMS, everyMS apart.
func seqEvery(fromMS, everyMS int64, n int) []int64 {
	out := make([]int64, n)
	for i := range out {
		out[i] = fromMS + int64(i)*everyMS
	}
	return out
}

// span is one aura segment [fromMS, toMS) with one stack.
func span(fromMS, toMS int64) summary.Segment {
	return summary.Segment{StartMS: fromMS, EndMS: toMS, Stacks: 1}
}

// noPool is no top-logs reference, with count distinct players found.
func noPool(count int) Refs { return Refs{PoolCount: count} }

// poolOf is a pool of n players (nil below PoolMinPlayers, as production builds it)
// whose median is m, compared over at most PoolTopN of them.
func poolOf(n int, m Profile) Refs {
	if n < PoolMinPlayers {
		return noPool(n)
	}
	return Refs{Pool: &PoolRef{N: min(n, PoolTopN), Median: m}, PoolCount: n}
}

func poolCPM(family string, cpm float64, n int) Refs {
	return poolOf(n, Profile{CastsPerMin: map[string]float64{family: cpm}})
}

func poolUptime(spellID int64, pct float64, n int) Refs {
	return poolOf(n, Profile{Uptime: map[string]float64{SpellKey(spellID): pct}})
}

func poolDefensive(spellID int64, pct float64, n int) Refs {
	return poolOf(n, Profile{DefensiveUptime: map[string]float64{SpellKey(spellID): pct}})
}

func poolOverheal(spellID int64, pct float64, n int) Refs {
	return poolOf(n, Profile{OverhealByRank: map[string]float64{SpellKey(spellID): pct}})
}

func poolAdherence(adherence float64, n int) Refs {
	return poolOf(n, Profile{Adherence: &adherence})
}

// withPool is a copy of in with its Refs replaced.
func withPool(in Input, r Refs) Input {
	in.Refs = r
	return in
}

// withCombatantInput is withCombatant on in.Fight.
func withCombatantInput(in Input, guid string, gear bool) Input {
	in.Fight = withCombatant(in.Fight, guid, gear)
	return in
}

// textOf is every player-facing string of o, newline-joined: the finding's title, detail
// and basis, every fact's text and name, every strength sentence and couldn't-check line.
func textOf(o outcome) string {
	var parts []string
	if f := o.Finding; f != nil {
		parts = append(parts, f.Title, f.Detail, f.Basis.Label)
		if f.Comparison != nil {
			parts = append(parts, f.Comparison.Label)
		}
		parts = append(parts, factsText(f.Facts)...)
	}
	for _, s := range o.Strengths {
		parts = append(parts, s.Sentence)
		parts = append(parts, factsText(s.Facts)...)
	}
	parts = append(parts, o.CouldNotCheck...)
	return strings.Join(parts, "\n")
}

func factsText(f Facts) []string {
	var out []string
	for _, n := range f.Numbers {
		out = append(out, n.Text, n.Of)
	}
	for _, s := range f.Spells {
		out = append(out, s.Name)
	}
	out = append(out, f.Items...)
	return append(out, f.Times...)
}

// assertNumbers fails unless every want is the text of one of facts' numbers.
func assertNumbers(t *testing.T, facts Facts, want ...string) {
	t.Helper()
	var got []string
	for _, n := range facts.Numbers {
		got = append(got, n.Text)
	}
	assertContainsAll(t, "numbers", got, want)
}

// assertTimes fails unless every want is one of facts' times.
func assertTimes(t *testing.T, facts Facts, want ...string) {
	t.Helper()
	assertContainsAll(t, "times", facts.Times, want)
}

// assertSpells fails unless every want names one of facts' spells.
func assertSpells(t *testing.T, facts Facts, want ...string) {
	t.Helper()
	var got []string
	for _, s := range facts.Spells {
		got = append(got, s.Name)
	}
	assertContainsAll(t, "spells", got, want)
}

func assertContainsAll(t *testing.T, what string, got, want []string) {
	t.Helper()
	for _, w := range want {
		if !slices.Contains(got, w) {
			t.Errorf("facts' %s %q lack %q", what, got, w)
		}
	}
}

// furyInput is the golden's Fury warrior at level 60 on Onyxia (40), with the warrior's
// fixture tables and no pool.
func furyInput(t *testing.T) Input {
	t.Helper()
	s := withLevel(withRole(goldenFight(t), furyGUID, "Warrior", "Fury", RoleDPS), furyGUID, fixtureLevel)
	return fixtureInput(t, s, furyGUID, furyRules, warriorSpells, warriorTrainable)
}

// holyInput is the golden's priest made a level 60 Holy priest.
func holyInput(t *testing.T) Input {
	t.Helper()
	s := withLevel(withRole(goldenFight(t), priestGUID, "Priest", "Holy", RoleHealer), priestGUID, fixtureLevel)
	return fixtureInput(t, s, priestGUID, holyRules, priestSpells, priestTrainable)
}

// protInput is the golden's warrior made a level 60 Protection warrior.
func protInput(t *testing.T) Input {
	t.Helper()
	s := withLevel(withRole(goldenFight(t), furyGUID, "Warrior", "Protection", RoleTank), furyGUID, fixtureLevel)
	return fixtureInput(t, s, furyGUID, protRules, warriorSpells, warriorTrainable)
}

func fixtureInput(t *testing.T, s summary.Summary, guid string, rules SpecRules, spells map[int64]SpellConst, trainable map[string][]TrainableRank) Input {
	t.Helper()
	row := rosterRowOf(t, s, guid)
	rot, rotOK := fixtureRotation(t, rating.DataSpecSlug(row.Class, row.Spec))
	cat := consumables.Load()
	role, ok := cat.For(row.Role)
	if !ok {
		t.Fatalf("no consumables catalogue for role %q", row.Role)
	}
	table, ok := mechanics.Load(onyxiaEncounterID)
	if !ok {
		t.Fatal("no Onyxia mechanics table")
	}
	return Input{
		Fight: s, Row: row, Size: 40, BossName: onyxiaName, Refs: noPool(0),
		Tables: Tables{
			Rules: rules, Spells: spells, Trainable: trainable, Rotation: rot, RotationOK: rotOK,
			Consumables: role, RaidBuffs: cat.RaidBuffs, Mechanics: &table, WorstActions: map[string]bool{},
		},
	}
}

func rosterRowOf(t *testing.T, s summary.Summary, guid string) summary.RosterRow {
	t.Helper()
	for _, r := range s.Roster {
		if r.GUID == guid {
			return r
		}
	}
	t.Fatalf("no roster row for %s", guid)
	return summary.RosterRow{}
}

var (
	gearOnce sync.Once
	gearErr  error
)

// fixtureRotation is the newest build's shipped rotation for spec at level 60.
func fixtureRotation(t *testing.T, spec string) (sims.Rotation, bool) {
	t.Helper()
	gearOnce.Do(func() {
		data, err := trees.Load(buildsDir)
		if err != nil {
			gearErr = err
			return
		}
		sims.SetGearData(data, buildsDir)
	})
	if gearErr != nil {
		t.Fatal(gearErr)
	}
	return sims.RotationFor(activeBuild(t), spec, fixtureLevel)
}

// activeBuild is web/src/data/active-build.json's build, the newest build (Global
// Constraints).
func activeBuild(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile(activeBuildRef)
	if err != nil {
		t.Fatal(err)
	}
	var v struct {
		Build string `json:"build"`
	}
	if err := json.Unmarshal(b, &v); err != nil || v.Build == "" {
		t.Fatalf("active-build.json: %v %q", err, v.Build)
	}
	return v.Build
}

// buildFile is a file of the newest build's data, by path under data/builds/<build>/.
func buildFile(t *testing.T, parts ...string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(append([]string{buildsDir, activeBuild(t)}, parts...)...))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// The fixture spell constants, by hand; TestFixtureSpellConstMatchesData proves each one
// equals the newest build's spellconst. Rage costs are in tenths (300 is 30 rage).
var warriorSpells = map[int64]SpellConst{
	23881: {Name: "Bloodthirst", Rank: 1, CooldownMS: 6000, GCDMS: 1500, Cost: 300, CostType: 1, CostKnown: true},
	23894: {Name: "Bloodthirst", Rank: 4, CooldownMS: 6000, GCDMS: 1500, Cost: 300, CostType: 1, CostKnown: true},
	6673:  {Name: "Battle Shout", Rank: 1, GCDMS: 1500, Cost: 100, CostType: 1, CostKnown: true},
	25289: {Name: "Battle Shout", Rank: 7, GCDMS: 1500, Cost: 100, CostType: 1, CostKnown: true},
	11596: {Name: "Sunder Armor", Rank: 4, GCDMS: 1500, Cost: 150, CostType: 1, CostKnown: true},
	11597: {Name: "Sunder Armor", Rank: 5, GCDMS: 1500, Cost: 150, CostType: 1, CostKnown: true},
	11601: {Name: "Revenge", Rank: 5, CooldownMS: 5000, GCDMS: 1500, Cost: 50, CostType: 1, CostKnown: true},
	25288: {Name: "Revenge", Rank: 6, CooldownMS: 5000, GCDMS: 1500, Cost: 50, CostType: 1, CostKnown: true},
	23925: {Name: "Shield Slam", Rank: 4, CooldownMS: 6000, GCDMS: 1500, Cost: 200, CostType: 1, CostKnown: true},
	2565:  {Name: "Shield Block", CooldownMS: 5000, Cost: 100, CostType: 1, CostKnown: true},
	871:   {Name: "Shield Wall", CooldownMS: 900000, GCDMS: 1500, CostType: 1, CostKnown: true},
	12975: {Name: "Last Stand", CooldownMS: 180000, CostKnown: true},
}

var priestSpells = map[int64]SpellConst{
	10964: {Name: "Greater Heal", Rank: 3, GCDMS: 1500, Cost: 545, CostKnown: true},
	10965: {Name: "Greater Heal", Rank: 4, GCDMS: 1500, Cost: 655, CostKnown: true},
	25314: {Name: "Greater Heal", Rank: 5, GCDMS: 1500, Cost: 710, CostKnown: true},
	14751: {Name: "Inner Focus", CooldownMS: 180000, CostKnown: true},
	10952: {Name: "Inner Fire", Rank: 6, GCDMS: 1500, Cost: 315, CostKnown: true},
}

// The fixture trainable ranks; TestFixtureTrainablesMatchData proves them.
var warriorTrainable = map[string][]TrainableRank{
	"Bloodthirst":  {{23881, 1, 40}, {23892, 2, 48}, {23893, 3, 54}, {23894, 4, 60}},
	"Sunder Armor": {{7386, 1, 10}, {7405, 2, 22}, {8380, 3, 34}, {11596, 4, 46}, {11597, 5, 58}},
	"Revenge":      {{6572, 1, 14}, {6574, 2, 24}, {7379, 3, 34}, {11600, 4, 44}, {11601, 5, 54}, {25288, 6, 60}},
	"Shield Slam":  {{23922, 1, 40}, {23923, 2, 48}, {23924, 3, 54}, {23925, 4, 60}},
	"Shield Wall":  {{871, 0, 28}},
	"Last Stand":   {{12975, 0, 1}},
}

var priestTrainable = map[string][]TrainableRank{
	"Greater Heal": {{2060, 1, 40}, {10963, 2, 46}, {10964, 3, 52}, {10965, 4, 58}, {25314, 5, 60}},
	"Inner Fire":   {{588, 1, 12}, {7128, 2, 20}, {602, 3, 30}, {1006, 4, 40}, {10951, 5, 50}, {10952, 6, 60}},
}

// The fixture rule files, inline; Task 15 ships the real ones.
var furyRules = SpecRules{
	Spec: "warrior-fury", Build: fixtureVerified,
	Core: []RuleEntry{{Family: "Bloodthirst", SpellIDs: []int64{23881, 23892, 23893, 23894}, Talent: true, Verified: fixtureVerified}},
	Auras: []AuraRule{
		{SpellID: 6673, Name: "Battle Shout", Kind: AuraSelfBuff, Verified: fixtureVerified},
		{SpellID: 11597, Name: "Sunder Armor", Kind: AuraDebuff, Verified: fixtureVerified},
	},
}

var protRules = SpecRules{
	Spec: "warrior-protection", Build: fixtureVerified,
	Core: []RuleEntry{
		{Family: "Revenge", SpellIDs: []int64{6572, 6574, 7379, 11600, 11601, 25288}, Verified: fixtureVerified},
		{Family: "Shield Slam", SpellIDs: []int64{23922, 23923, 23924, 23925}, Talent: true, Verified: fixtureVerified},
	},
	Auras: []AuraRule{{SpellID: 2565, Name: "Shield Block", Kind: AuraDefensive, Verified: fixtureVerified}},
	Defensives: []RuleEntry{
		{Family: "Shield Wall", SpellIDs: []int64{871}, Verified: fixtureVerified},
		{Family: "Last Stand", SpellIDs: []int64{12975}, Talent: true, Verified: fixtureVerified},
	},
}

var holyRules = SpecRules{
	Spec: "priest-holy", Build: fixtureVerified,
	Auras:      []AuraRule{{SpellID: 10952, Name: "Inner Fire", Kind: AuraSelfBuff, Verified: fixtureVerified}},
	Efficiency: []RuleEntry{{Family: "Inner Focus", SpellIDs: []int64{14751}, Talent: true, Verified: fixtureVerified}},
}

// examplePack is the §5.3 Fury pack, value for value (invented values: tests only).
func examplePack() Pack {
	at := func(id string) Anchor { return Anchor{ID: id, Kind: AnchorCast} }
	return Pack{
		Version: "pack-v1.example", Status: StatusReady,
		Context: Context{SpecLabel: "Fury", Class: "Warrior", Role: "dps", Boss: "Onyxia", Size: 40, Kill: true, Duration: "3:52"},
		Findings: []Finding{
			{ID: "f1", Code: CodeCoreCasts, Severity: SeverityMajor, Rank: 1, Title: "Bloodthirst: 9 casts lost",
				Detail: "31 of a possible 40 with rage in your bar. Longest gaps 1:12–1:25 and 2:40–2:51.",
				Facts:  Facts{Spells: []SpellFact{{Name: "Bloodthirst"}}}, Anchors: []string{"a1", "a2", "a3"}},
			{ID: "f6", Code: CodePreparation, Severity: SeverityMajor, Rank: 2, Title: "No flask at the pull",
				Detail: "You pulled with Grilled Squid and an Elemental Sharpening Stone, and no flask.",
				Facts:  Facts{Items: []string{"Grilled Squid", "Elemental Sharpening Stone"}}, Anchors: []string{"a4"}},
			{ID: "f4", Code: CodeResource, Severity: SeverityMinor, Rank: 3, Title: "Rage full for 14 seconds",
				Detail: "Your rage sat at its cap for 14 s outside the air phase, longest at 2:41–2:50.", Anchors: []string{"a5"}},
		},
		Strengths: []Strength{
			{ID: "s1", Sentence: "Alive the whole fight."},
			{ID: "s2", Sentence: "Battle Shout up 100% of the fight. Top 25: 97%."},
		},
		Links: []Link{{ID: "l1", Findings: [2]string{"f1", "f4"}, Kind: LinkGapOverlapsCap,
			Sentence: "The Bloodthirst gap at 2:40–2:51 overlaps the rage cap at 2:41–2:50."}},
		CouldNotCheck: []string{"Talents and race: Forever logs don't record them.", "Range and movement: the log doesn't say where you stood."},
		Anchors:       []Anchor{at("a1"), at("a2"), at("a3"), at("a4"), at("a5")},
		Allowed: Allowed{
			Numbers: []string{"31", "40", "9", "8", "9.1", "14", "3", "25", "100%", "97%", "3:52", "0:21", "1:12", "2:40", "2:41", "2:50", "2:51", "1:12–1:25", "2:40–2:51", "2:41–2:50"},
			Names:   []string{"Bloodthirst", "Heroic Strike", "Battle Shout", "Onyxia", "Grilled Squid", "Elemental Sharpening Stone"},
			IDs:     []string{"a1", "a2", "a3", "a4", "a5", "f1", "f4", "f6", "l1", "s1", "s2"},
		},
	}
}

// exampleReview is the §5.3 Fury review, word for word.
func exampleReview() ReviewDoc {
	return ReviewDoc{
		Opening: "Solid kill: you were alive the whole fight and Battle Shout was up 100% of it. Almost everything you left behind sits in one button.",
		FixFirst: &FixItem{
			Finding: "f1", Title: "Bloodthirst: 9 casts lost",
			Why: "Bloodthirst is the button the Fury rotation is built around, and 9 of a possible 40 went unused while you had the rage, " +
				"longest at 1:12–1:25. At 2:40–2:51 you sat on a full bar instead, so that one gap cost you twice.",
			Drill:     "drill.core_casts",
			DrillText: "On a target dummy, press Bloodthirst the moment it comes off cooldown, before anything else; Whirlwind and Heroic Strike wait their turn.",
			Anchors:   []string{"a1", "a2", "a3", "a5"}, Links: []string{"l1"},
		},
		ThenFix: []ThenItem{
			{Finding: "f6", Title: "No flask at the pull", Why: "A flask is the cheapest damage you will ever buy; bring one to every 40-player boss.", Anchors: []string{"a4"}, Links: []string{}},
			{Finding: "f4", Title: "Rage full for 14 seconds", Why: "When the bar is full, Heroic Strike turns rage into damage instead of losing it.", Anchors: []string{"a5"}, Links: []string{"l1"}},
		},
		DoingWell: []WellItem{
			{Strength: "s1", Text: "You were alive the whole fight.", Anchors: []string{}},
			{Strength: "s2", Text: "Battle Shout was up 100% of the fight, against 97% for the top 25.", Anchors: []string{}},
		},
		CouldntCheck: []string{"Talents and race: Forever logs don't record them.", "Range and movement: the log doesn't say where you stood."},
	}
}
