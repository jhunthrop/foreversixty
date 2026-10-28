package coach

// The web's coach fixtures (web/src/fixtures/coach/*.json): every answer the e2e specs and the
// review captures stub, built here through the real rules (Review over the shipped rule files
// and build tables), the rule-only renderer and the route's own DTO shaping, never typed by
// hand. The two model reviews are the design spec's §5.3 worked reviews, held to the real
// validator against the pack they are written from.
//
//	go test ./internal/coach -run TestWebFixtures_Write -update-coach-fixtures   # rewrite
//	go test ./internal/coach -run TestWebFixtures_NoDrift                        # fail on drift

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

var updateCoachFixtures = flag.Bool("update-coach-fixtures", false, "rewrite web/src/fixtures/coach")

const (
	webFixturesDir = "../../../web/src/fixtures/coach"
	// The e2e fixture report the captures open: its id, the boss pull and the two owned rows.
	webReportID = "fixture2abcd"
	webFight    = 3
	webFuryKey  = "us/hardcore/baelgrim"
	webHolyKey  = "us/hardcore/sunwick"
	webFightMS  = 232_000 // Onyxia (40), kill, 3:52
	webMCZone   = "Molten Core"
	webOnyxZone = "Onyxia's Lair"
	webCooldown = 6_000 // Bloodthirst's
)

var (
	webReviewedAt = time.Date(2026, 10, 25, 2, 10, 0, 0, time.UTC)
	webFoughtAt   = time.Date(2026, 10, 25, 1, 30, 0, 0, time.UTC)
	// The Molten Core night of the account line: eight bosses in kill order.
	webMCBosses = []string{"Lucifron", "Magmadar", "Gehennas", "Garr", "Baron Geddon", "Shazzrah", "Golemagg the Incinerator", "Ragnaros"}
	// The worked example's gaps: nine casts lost, the longest at 1:12–1:25 and 2:44–2:55.
	webFuryGaps = [][2]int64{{24_000, 32_500}, {50_500, 59_000}, {72_000, 85_000}, {103_000, 111_500},
		{129_500, 138_000}, {150_000, 158_000}, {164_000, 175_000}, {193_000, 201_500}, {213_500, 222_000}}
)

// webTables is in with the production tables for its row: the shipped rule file, the active
// build's spell data and trainables, the consumables catalogue and Onyxia's mechanics.
func webTables(t *testing.T, in Input) Input {
	t.Helper()
	fixtureRotation(t, "warrior-fury") // loads the gear data RotationFor reads
	ft := &FileTables{BuildsDir: repoBuilds, CoachDir: ruleFilesDir}
	tables, err := ft.TablesFor(context.Background(), in.Row, onyxiaEncounterID, webFoughtAt)
	if err != nil {
		t.Fatal(err)
	}
	in.Tables = tables
	return in
}

// castsAround is one cast every step or a little more, from first to last inclusive.
func castsAround(first, last, step int64) []int64 {
	if last <= first {
		return []int64{first}
	}
	n := (last - first) / step
	out := make([]int64, 0, n+1)
	for i := int64(0); i <= n; i++ {
		out = append(out, first+(last-first)*i/n)
	}
	return out
}

// bloodthirstWithGaps is Bloodthirst on cooldown but for each [ready, next cast) gap.
func bloodthirstWithGaps(gaps [][2]int64) []int64 {
	var out []int64
	from := int64(0)
	for _, g := range gaps {
		out = append(out, castsAround(from, g[0]-webCooldown, webCooldown)...)
		from = g[1]
	}
	return append(out, castsAround(from, webFightMS-2*webCooldown, webCooldown)...)
}

// furyNight is the Fury warrior on Onyxia (40) for 3:52 with the rest of the rotation cast:
// Whirlwind and Bloodrage on cooldown, Heroic Strike as the rage dump, three Hamstrings.
func furyNight(t *testing.T, ss seeds, bloodthirst []int64, capped ...summary.Segment) Input {
	t.Helper()
	in := webTables(t, reviewBase(t, webFightMS, ss))
	in = withBloodthirst(in, bloodthirstRank4ID, bloodthirst, capped...)
	in.Fight = withCasts(in.Fight, furyGUID, 1680, "Whirlwind", seqEvery(3_000, 10_000, 23)...)
	in.Fight = withCasts(in.Fight, furyGUID, 2687, "Bloodrage", 0, 60_000, 120_000, 180_000)
	in.Fight = withCasts(in.Fight, furyGUID, 25286, "Heroic Strike", seqEvery(8_000, 12_000, 18)...)
	in.Fight = withCasts(in.Fight, furyGUID, 7373, "Hamstring", 95_000, 140_000, 186_000)
	adherence := 0.93
	in.Refs = poolOf(25, Profile{CastsPerMin: map[string]float64{"Bloodthirst": 9.1, "Whirlwind": 5.6, "Bloodrage": 1},
		Uptime: map[string]float64{SpellKey(25289): 97}, CappedPct: 1.3, Adherence: &adherence})
	return in
}

// furyWorked is the §5.3 Fury example: Bloodthirst gaps, no flask, rage capped 14 s.
func furyWorked(t *testing.T) Input {
	t.Helper()
	return furyNight(t, seededAt(food("Grilled Squid")), bloodthirstWithGaps(webFuryGaps), span(40_000, 45_000), span(165_000, 174_000))
}

// furyOneFinding is the worked example with the flask brought and the bar never full.
func furyOneFinding(t *testing.T) Input {
	t.Helper()
	return furyNight(t, fullKit(), bloodthirstWithGaps(webFuryGaps))
}

// furyClean is everything brought and Bloodthirst on cooldown: nothing to fix.
func furyClean(t *testing.T) Input {
	t.Helper()
	return furyNight(t, fullKit(), bloodthirstWithGaps(nil))
}

// furyRaidBuff is furyClean with a druid in the raid whose Mark of the Wild the warrior lacks.
func furyRaidBuff(t *testing.T) Input {
	t.Helper()
	return rosterOf(furyClean(t), rosterClass("druid", "Druid"))
}

// holyWorked is the §5.3 Holy priest example: died to Flame Breath at 3:20, out of mana from
// 3:05, Greater Heal (Rank 4) at 44% overheal, flask and Inner Fire up.
func holyWorked(t *testing.T) Input {
	t.Helper()
	refs := poolOf(21, Profile{OverhealByRank: map[string]float64{SpellKey(10965): 22}, Uptime: map[string]float64{SpellKey(10952): 98}})
	in := webTables(t, holyOverheal(t, 10965, 23, 44, refs))
	in.Fight = withRoster(in.Fight, in.Row)
	in.Fight = seedPotion(withSeeds(in.Fight, priestGUID, seededAt(flask("Flask of Distilled Wisdom"), food("Nightfin Soup"))), priestGUID)
	in = withGear(in, preparedGear(true))
	in.Fight = withAura(in.Fight, priestGUID, 10952, "Inner Fire", []string{priestGUID}, span(0, 200_000))
	in = f7Died(in, webFightMS, 200_000, flameBreathID, "Flame Breath", 2_140)
	in.Fight = withResource(f4WithoutResources(in.Fight, priestGUID), priestGUID, PowerMana, f4ManaMax,
		f4Series(webFightMS, f4ManaMax*3/5, f4ManaMax/100, span(185_000, webFightMS))...)
	return in
}

func webPersona(t *testing.T, spec string) Persona {
	t.Helper()
	p, err := LoadPersona(ruleFilesDir, spec)
	if err != nil {
		t.Fatal(err)
	}
	return p
}

// packFindingBy is the pack's finding with code; the test fails without one.
func packFindingBy(t *testing.T, p Pack, code Code) Finding {
	t.Helper()
	f, ok := findingOf(p.Findings, code)
	if !ok {
		t.Fatalf("pack has no %s finding", code)
	}
	return f
}

// packStrength is the pack's strength with sentence; the test fails without one.
func packStrength(t *testing.T, p Pack, sentence string) Strength {
	t.Helper()
	for _, s := range p.Strengths {
		if s.Sentence == sentence {
			return s
		}
	}
	t.Fatalf("pack has no strength %q", sentence)
	return Strength{}
}

// linkBetween is the ids of the pack's links joining a and b.
func linkBetween(p Pack, a, b string) []string {
	out := []string{}
	for _, l := range p.Links {
		if (l.Findings[0] == a && l.Findings[1] == b) || (l.Findings[0] == b && l.Findings[1] == a) {
			out = append(out, l.ID)
		}
	}
	return out
}

func fixFrom(f Finding, why string, persona Persona, anchors, links []string) *FixItem {
	drill := drillIDPrefix + string(f.Code)
	return &FixItem{Finding: f.ID, Title: f.Title, Why: why, Drill: drill, DrillText: persona.Drills[drill], Anchors: anchors, Links: links}
}

func thenFrom(f Finding, why string, links []string) ThenItem {
	return ThenItem{Finding: f.ID, Title: f.Title, Why: why, Anchors: slices.Clone(f.Anchors), Links: links}
}

// furyModelDoc is the §5.3 Fury review, written from the worked pack.
func furyModelDoc(t *testing.T, p Pack, persona Persona) ReviewDoc {
	t.Helper()
	core, prep, res := packFindingBy(t, p, CodeCoreCasts), packFindingBy(t, p, CodePreparation), packFindingBy(t, p, CodeResource)
	alive, shout := packStrength(t, p, "Alive the whole fight."), packStrength(t, p, "Battle Shout up 100% of the fight. Top 25: 97%.")
	return ReviewDoc{
		Opening: "Solid kill: you were alive the whole fight and Battle Shout was up 100% of it. Almost everything you left behind sits in one button.",
		FixFirst: fixFrom(core, "Bloodthirst is the button the Fury rotation is built around, and 9 casts went unused while you had the rage, "+
			"longest at 1:12–1:25. At 2:44–2:55 you sat on a full bar instead, so that one gap cost you twice.",
			persona, slices.Concat(core.Anchors, res.Anchors), linkBetween(p, core.ID, res.ID)),
		ThenFix: []ThenItem{
			thenFrom(prep, "A flask is the cheapest damage you will ever buy; bring one to every 40-player boss.", []string{}),
			thenFrom(res, "When the bar is full, Heroic Strike turns rage into damage instead of losing it.", linkBetween(p, core.ID, res.ID)),
		},
		DoingWell: []WellItem{
			{Strength: alive.ID, Text: "You were alive the whole fight.", Anchors: slices.Clone(alive.Anchors)},
			{Strength: shout.ID, Text: "Battle Shout was up 100% of the fight, against 97% for the top 25.", Anchors: slices.Clone(shout.Anchors)},
		},
		CouldntCheck: slices.Clone(p.CouldNotCheck),
	}
}

// holyModelDoc is the §5.3 Holy priest review, written from the worked pack.
func holyModelDoc(t *testing.T, p Pack, persona Persona) ReviewDoc {
	t.Helper()
	dead, res, rank := packFindingBy(t, p, CodeDeath), packFindingBy(t, p, CodeResource), packFindingBy(t, p, CodeRanks)
	flask := packStrength(t, p, "Flask of Distilled Wisdom at the pull.")
	fire := packStrength(t, p, "Inner Fire up 100% of the fight. Top 21: 98%.")
	return ReviewDoc{
		Opening: "Your preparation was right: flask at the pull and Inner Fire up 100% of the fight. This one is about lasting to the end.",
		FixFirst: fixFrom(dead, "A dead healer heals nothing for the rest of the pull, and Flame Breath is damage you can step out of. "+
			"You were also out of mana when it landed, so there was no answer left.",
			persona, slices.Concat(dead.Anchors, res.Anchors), linkBetween(p, dead.ID, res.ID)),
		ThenFix: []ThenItem{
			thenFrom(res, "The mana ran out 15 s before the death; the fix is upstream, in how you spend it.", []string{}),
			thenFrom(rank, "Top Holy priests here overheal 22% with this rank. A lower rank costs less mana, and that is where your last 15 s went.",
				linkBetween(p, rank.ID, res.ID)),
		},
		DoingWell: []WellItem{
			{Strength: flask.ID, Text: flask.Sentence, Anchors: slices.Clone(flask.Anchors)},
			{Strength: fire.ID, Text: "Inner Fire was up 100% of the fight, against 98% for the top 21.", Anchors: slices.Clone(fire.Anchors)},
		},
		CouldntCheck: slices.Clone(p.CouldNotCheck),
	}
}

// webAnswer is one answer as the route shapes it for (player, access) from a served review.
func webAnswer(in Input, player string, access Access, sv served) fullDTO {
	sv.at = webReviewedAt
	req := fightRequest{reportID: webReportID, fight: webFight, player: player}
	return fullOf(req, loaded{pack: Review(in), tables: in.Tables}, access, sv)
}

// rulesOf is the rule-only rendering, served with status (ready unless given).
func rulesOf(p Pack, persona Persona) served {
	return served{doc: RenderRules(p, persona.Drills), source: SourceRules, status: string(StatusReady), persona: persona, hasPersona: true}
}

func modelOf(doc ReviewDoc, persona Persona) served {
	return served{doc: doc, source: SourceModel, status: string(StatusReady), persona: persona, hasPersona: true}
}

// withStatus is the rule-only rendering of in served under status, as serve answers a pack
// that is not ready.
func withStatus(in Input, persona Persona, status Status) fullDTO {
	sv := rulesOf(Review(in), persona)
	sv.status = string(status)
	return webAnswer(in, webFuryKey, AccessFull, sv)
}

// assertValid holds a model review to the generator's own validator.
func assertValid(t *testing.T, name string, doc ReviewDoc, p Pack, persona Persona) {
	t.Helper()
	v := Validator{Names: buildNames(t), Roster: []string{"Baelgrim", "Sunwick"}}
	if probs := v.Validate(doc, p, persona); len(probs) > 0 {
		t.Fatalf("%s model review fails the validator: %+v", name, probs)
	}
}

// webNight is the last-raid answer for a night of packs, as lastRaid shapes it.
func webNight(zone string, fights []NightFight) lastRaidDTO {
	s := &Service{}
	out := lastRaidDTO{Zone: zone, FoughtAt: webFoughtAt, ReportID: webReportID, URL: s.ReportURL(webReportID, 0)}
	if line, ok := Night(fights); ok {
		out.FindingCount = line.FindingCount
		out.First = &firstDTO{Stem: line.Stem, On: line.On, Of: line.Of}
		out.URL = s.ReportURL(webReportID, line.FightIndex)
	}
	return out
}

// moltenCoreNightPacks is eight bosses: Bloodthirst lost on five (fights 3 to 7), the rest clean.
func moltenCoreNightPacks(t *testing.T) []NightFight {
	t.Helper()
	out := make([]NightFight, 0, len(webMCBosses))
	for i, boss := range webMCBosses {
		in := furyClean(t)
		if i >= 2 && i <= 6 {
			in = furyWorked(t)
		}
		in.BossName = boss
		out = append(out, NightFight{FightIndex: i + 1, Pack: Review(in)})
	}
	return out
}

// webFixtures is every file under web/src/fixtures/coach, by name.
func webFixtures(t *testing.T) map[string]any {
	t.Helper()
	fury, holy := webPersona(t, "warrior-fury"), webPersona(t, "priest-holy")
	starter := fury
	starter.Status = PersonaStarter

	worked := furyWorked(t)
	workedPack := Review(worked)
	furyDoc := furyModelDoc(t, workedPack, fury)
	assertValid(t, "fury", furyDoc, workedPack, fury)
	holyIn := holyWorked(t)
	holyPack := Review(holyIn)
	holyDoc := holyModelDoc(t, holyPack, holy)
	assertValid(t, "holy", holyDoc, holyPack, holy)

	writing, busy := rulesOf(workedPack, fury), rulesOf(workedPack, fury)
	writing.status, busy.busy = statusWriting, true
	trash, short, noSpec, preparing := furyWorked(t), furyWorked(t), furyWorked(t), furyWorked(t)
	trash.Fight.EncounterID = 0
	noSpec.Row.Spec = ""
	short.Fight = withDuration(short.Fight, 25_000)
	preparing.Refs = Refs{SimPreparing: true}
	clean, oneFinding, raidBuff := furyClean(t), furyOneFinding(t), furyRaidBuff(t)
	historyFirst, historyOlder := webHistoryPages(webHistoryItems(t, workedPack, furyDoc, fury))

	return map[string]any{
		"fury-model.json":                webAnswer(worked, webFuryKey, AccessFull, modelOf(furyDoc, fury)),
		"fury-rules.json":                webAnswer(worked, webFuryKey, AccessFull, rulesOf(workedPack, fury)),
		"fury-writing.json":              webAnswer(worked, webFuryKey, AccessFull, writing),
		"fury-busy.json":                 webAnswer(worked, webFuryKey, AccessFull, busy),
		"fury-starter.json":              webAnswer(worked, webFuryKey, AccessFull, modelOf(furyDoc, starter)),
		"fury-officer.json":              webAnswer(worked, webFuryKey, AccessOfficer, modelOf(furyDoc, fury)),
		"fury-no-findings.json":          webAnswer(clean, webFuryKey, AccessFull, rulesOf(Review(clean), fury)),
		"fury-raid-buffs-only.json":      webAnswer(raidBuff, webFuryKey, AccessFull, rulesOf(Review(raidBuff), fury)),
		"fury-preparing.json":            withStatus(preparing, fury, StatusPreparing),
		"fury-not-reviewable-trash.json": withStatus(trash, fury, StatusNotReviewable),
		"fury-not-reviewable-short.json": withStatus(short, fury, StatusNotReviewable),
		"not-reviewable-spec.json":       withStatus(noSpec, fury, StatusNotReviewable),
		"fury-teaser-several.json":       teaserOf(workedPack),
		"fury-teaser-one.json":           teaserOf(Review(oneFinding)),
		"fury-teaser-empty.json":         teaserOf(Review(clean)),
		"fury-teaser-no-first.json":      teaserOf(Review(raidBuff)),
		"holy-model.json":                webAnswer(holyIn, webHolyKey, AccessFull, modelOf(holyDoc, holy)),
		"holy-rules.json":                webAnswer(holyIn, webHolyKey, AccessFull, rulesOf(holyPack, holy)),
		"last-raid-molten-core.json":     webNight(webMCZone, moltenCoreNightPacks(t)),
		"last-raid-onyxia.json":          webNight(webOnyxZone, []NightFight{{FightIndex: webFight, Pack: workedPack}}),
		"last-raid-clean.json":           webNight(webOnyxZone, []NightFight{{FightIndex: webFight, Pack: Review(clean)}}),
		"history-first.json":             historyFirst,
		"history-older.json":             historyOlder,
	}
}

func marshalFixture(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return append(b, '\n')
}

// TestWebFixtures_Write rewrites the web's fixtures under -update-coach-fixtures.
func TestWebFixtures_Write(t *testing.T) {
	if !*updateCoachFixtures {
		t.Skip("-update-coach-fixtures is not set")
	}
	if err := os.MkdirAll(webFixturesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for name, v := range webFixtures(t) {
		if err := os.WriteFile(filepath.Join(webFixturesDir, name), marshalFixture(t, v), 0o600); err != nil {
			t.Fatal(err)
		}
	}
}

// TestWebFixtures_NoDrift fails when a checked-in fixture is not what the rules produce today.
func TestWebFixtures_NoDrift(t *testing.T) {
	if *updateCoachFixtures {
		t.Skip("rewriting")
	}
	want := webFixtures(t)
	for name, v := range want {
		got, err := os.ReadFile(filepath.Join(webFixturesDir, name))
		if err != nil {
			t.Fatalf("%s: %v (run TestWebFixtures_Write with -update-coach-fixtures)", name, err)
		}
		if !bytes.Equal(got, marshalFixture(t, v)) {
			t.Errorf("%s drifted from the coach's output; run TestWebFixtures_Write with -update-coach-fixtures", name)
		}
	}
	files, err := filepath.Glob(filepath.Join(webFixturesDir, "*.json"))
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if _, ok := want[filepath.Base(f)]; !ok {
			t.Errorf("%s is not produced by the coach; delete it", filepath.Base(f))
		}
	}
}
