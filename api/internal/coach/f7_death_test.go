package coach

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func TestDeath(t *testing.T) {
	cases := []struct {
		name       string
		in         Input
		wantTitle  string
		wantSev    Severity
		wantDetail []string
		strength   string
		none       bool
	}{
		{name: "holy died to Flame Breath at 3:20", in: holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140),
			wantTitle: "Died at 3:20 to Flame Breath", wantSev: SeverityMinor,
			wantDetail: []string{"2,140", "No healing potion or Healthstone used in the 10 s before."}},
		{name: "prot died with Shield Wall not used", in: protDeathShieldWallReady(t, 232_000, 170_000),
			wantTitle: "Died at 2:50, Shield Wall not used", wantSev: SeverityMajor,
			wantDetail: []string{"Onyxia's melee killed you at 2:50. Shield Wall was not used this pull."}},
		{name: "prot died with Shield Wall ready", in: protDeathShieldWallReady(t, 1_200_000, 1_000_000),
			wantTitle: "Died at 16:40 with Shield Wall ready", wantSev: SeverityMinor,
			wantDetail: []string{"Onyxia's melee killed you at 16:40. Shield Wall had been ready since 15:00."}},
		{name: "death in the last 10% of a kill", in: holyDeath(t, 232_000, 220_000, flameBreathID, "Flame Breath", 2140), none: true},
		{name: "wipe tail death is part of the wipe", in: wipeDeath(t, 120_000, 110_000), none: true},
		{name: "no death is a strength", in: furyInput(t), none: true, strength: "Alive the whole fight."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := death(tc.in, CountedTime(tc.in))
			if tc.none != (o.Finding == nil) {
				t.Fatalf("finding %+v", o.Finding)
			}
			if tc.strength != "" && (len(o.Strengths) != 1 || o.Strengths[0].Sentence != tc.strength) {
				t.Errorf("strengths %+v", o.Strengths)
			}
			if tc.none {
				return
			}
			f := o.Finding
			if f.Title != tc.wantTitle || f.Severity != tc.wantSev || o.Anchors[0].Kind != AnchorDeath {
				t.Fatalf("got %+v", f)
			}
			for _, d := range tc.wantDetail {
				if !strings.Contains(f.Detail, d) {
					t.Errorf("detail %q lacks %q", f.Detail, d)
				}
			}
		})
	}
}

// The two §5.3 examples, word for word: title, detail and basis.
func TestDeath_SpecExamplesVerbatim(t *testing.T) {
	v := SpecVerbatim()
	holy := death(holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140), Counted{}).Finding
	if holy == nil {
		t.Fatal("holy: no finding")
	}
	if holy.Title != v["TitleDiedTo"] || holy.Detail != v["DetailDeathHit"]+" "+v["DeathNoPotion"] ||
		holy.Basis.Label != v["BasisAvoidable"] || holy.Basis.Kind != BasisOwnLog || holy.Code != CodeDeath {
		t.Errorf("holy %+v", holy)
	}
	// §5.3's Shield Wall example is a 15-minute cooldown at 2:50: a cast on the previous
	// pull may still have been cooling down, so the title is the §8.1 "not used" form
	// (final review A9) and the cooldown basis is kept for a proven-ready defensive.
	prot := death(protDeathShieldWallReady(t, 232_000, 170_000), Counted{}).Finding
	if prot == nil {
		t.Fatal("prot: no finding")
	}
	if prot.Title != Fill(TitleDiedNotUsed, map[string]string{"t": "2:50", "Spell": "Shield Wall"}) || prot.Basis.Label != BasisFromLog {
		t.Errorf("prot %+v", prot)
	}
	ready := death(protDeathShieldWallReady(t, 1_200_000, 1_000_000), Counted{}).Finding
	if ready == nil || ready.Title != Fill(TitleDiedWithReady, map[string]string{"t": "16:40", "Spell": "Shield Wall"}) ||
		ready.Basis.Label != v["BasisDeathCD"] || ready.Basis.Kind != BasisSpellData {
		t.Errorf("ready %+v", ready)
	}
}

// Last Stand is a talent: never cast that fight, it is not claimed ready. Shield Wall cast
// 70 s before on its 15-minute cooldown is not ready either.
func TestDeath_TalentDefensiveNeverCastIsNotReady(t *testing.T) {
	in := f7Cast(protDeathShieldWallReady(t, 232_000, 170_000), shieldWallID, "Shield Wall", 100_000)
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 2:50 to Onyxia's melee" {
		t.Fatalf("got %+v", f)
	}
	if strings.Contains(textOf(outcome{Finding: f}), "Last Stand") {
		t.Errorf("Last Stand claimed: %q", f.Detail)
	}
	if f.Basis.Label != FromYourLog+"." {
		t.Errorf("basis %q", f.Basis.Label)
	}
}

// A talent cast that fight and back off cooldown before the death is ready, but not
// "since the pull": the detail says only what the log shows.
func TestDeath_TalentCastAndBackOffCooldownIsReady(t *testing.T) {
	in := protDeathShieldWallReady(t, 400_000, 200_000)
	in = f7Cast(in, shieldWallID, "Shield Wall", 150_000)
	in = f7Cast(in, lastStandID, "Last Stand", 10_000)
	o := death(in, CountedTime(in))
	f := o.Finding
	if f == nil || f.Title != "Died at 3:20 with Last Stand ready" {
		t.Fatalf("got %+v", f)
	}
	if f.Detail != "Onyxia's melee killed you at 3:20. Last Stand had been ready since 3:10." {
		t.Errorf("detail %q", f.Detail)
	}
	if len(o.Anchors) != 2 || o.Anchors[1].Kind != AnchorCast || o.Anchors[1].SpellID != lastStandID {
		t.Errorf("anchors %+v", o.Anchors)
	}
}

// A baseline defensive counts only from its trainable level, and only when the level is
// known.
func TestDeath_BaselineDefensiveNeedsItsTrainableLevel(t *testing.T) {
	for _, level := range []int64{20, 0} {
		in := f7Level(protDeathShieldWallReady(t, 232_000, 170_000), level)
		f := death(in, CountedTime(in)).Finding
		if f == nil || f.Title != "Died at 2:50 to Onyxia's melee" {
			t.Errorf("level %d: got %+v", level, f)
		}
	}
	in := f7Level(protDeathShieldWallReady(t, 232_000, 170_000), 28)
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 2:50, Shield Wall not used" {
		t.Errorf("level 28: got %+v", f)
	}
}

// A defensive whose cooldown the data does not state is never claimed ready once cast.
func TestDeath_UnknownCooldownIsNotReady(t *testing.T) {
	in := protDeathShieldWallReady(t, 400_000, 300_000)
	in = f7Defensives(in, RuleEntry{Family: "Shield Wall", SpellIDs: []int64{99999}, Verified: fixtureVerified})
	in = f7Cast(in, 99999, "Shield Wall", 10_000)
	if f := death(in, CountedTime(in)).Finding; f == nil || strings.Contains(f.Title, "ready") {
		t.Errorf("got %+v", f)
	}
}

func TestDeath_PotionWindow(t *testing.T) {
	cases := []struct {
		name     string
		spellID  int64
		cast     string
		atMS     int64
		wantLine bool
	}{
		{name: "healing potion 5 s before", spellID: majorHealingPotID, cast: "Healing Potion", atMS: 195_000},
		{name: "healthstone at the 10 s edge", spellID: majorHealthstone, cast: "Major Healthstone", atMS: 190_000},
		{name: "healthstone 12 s before is outside", spellID: majorHealthstone, cast: "Major Healthstone", atMS: 188_000, wantLine: true},
		{name: "a healthstone conjured is not one used", spellID: 11730, cast: "Create Healthstone", atMS: 195_000, wantLine: true},
		{name: "rejuvenation potion 3 s before", spellID: majorRejuvPotionID, cast: "Major Rejuvenation Potion", atMS: 197_000},
		{name: "the warlock talent is not a use", spellID: 18693, cast: "Improved Healthstone", atMS: 195_000, wantLine: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := f7Cast(holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140), tc.spellID, tc.cast, tc.atMS)
			f := death(in, CountedTime(in)).Finding
			if f == nil {
				t.Fatal("no finding")
			}
			if got := strings.Contains(f.Detail, DeathNoPotion); got != tc.wantLine {
				t.Errorf("detail %q: no-potion line %v, want %v", f.Detail, got, tc.wantLine)
			}
		})
	}
}

// Flame Breath is the tank's to take (Onyxia's table, role tank): a tank who dies to it is
// never told it was avoidable.
func TestDeath_RoleTaggedAbilityIsNotAvoidableForThatRole(t *testing.T) {
	in := f7Defensives(f7Died(protInput(t), 232_000, 200_000, flameBreathID, "Flame Breath", 2140))
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 3:20 to Flame Breath" {
		t.Fatalf("got %+v", f)
	}
	if f.Basis.Label == Fill(BasisAvoidable, map[string]string{"Spell": "Flame Breath"}) {
		t.Errorf("tank told Flame Breath was avoidable")
	}
}

// The death anchor is the death's time only (the page adds the GUID); the window is the
// death instant.
func TestDeath_AnchorsAndWindow(t *testing.T) {
	notUsed := protDeathShieldWallReady(t, 232_000, 170_000)
	if o := death(notUsed, CountedTime(notUsed)); o.Finding == nil || len(o.Anchors) != 1 || len(o.Finding.Anchors) != 1 {
		t.Fatalf("a defensive never cast has no Casts row to anchor: %+v", o)
	}
	in := protDeathShieldWallReady(t, 1_200_000, 1_000_000)
	o := death(in, CountedTime(in))
	if o.Finding == nil || len(o.Anchors) != 2 {
		t.Fatalf("got %+v", o)
	}
	a := o.Anchors[0]
	if a != (Anchor{Kind: AnchorDeath, AtMS: a.AtMS}) || a.AtMS == nil || *a.AtMS != 1_000_000 {
		t.Errorf("death anchor %+v", a)
	}
	if b := o.Anchors[1]; b.Kind != AnchorCast || b.SpellID != shieldWallID {
		t.Errorf("defensive anchor %+v", b)
	}
	if got := o.Finding.Anchors; len(got) != 2 || got[0] != "#0" || got[1] != "#1" {
		t.Errorf("finding anchors %v", got)
	}
	if len(o.Windows) != 1 || o.Windows[0].FromMS != 1_000_000 || o.Windows[0].ToMS != 1_000_000 {
		t.Errorf("windows %+v", o.Windows)
	}
	if len(o.Strengths) != 0 {
		t.Errorf("a dead player has strength %+v", o.Strengths)
	}
}

func TestDeath_FactsCarryEveryQuotedValue(t *testing.T) {
	f := death(holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140), Counted{}).Finding
	if f == nil {
		t.Fatal("no finding")
	}
	assertTimes(t, f.Facts, "3:20")
	assertSpells(t, f.Facts, "Flame Breath")
	assertNumbers(t, f.Facts, "2,140")
	p := death(protDeathShieldWallReady(t, 232_000, 170_000), Counted{}).Finding
	if p == nil {
		t.Fatal("prot: no finding")
	}
	assertSpells(t, p.Facts, "Shield Wall", "Onyxia's melee")
	assertTimes(t, p.Facts, "2:50")
}

// On a wipe any reviewable death fires, the earliest first, and before 75% it is major.
func TestDeath_WipeEarlyDeathIsTheFirstAndMajor(t *testing.T) {
	in := wipeDeath(t, 120_000, 60_000)
	in.Fight = withDeath(in.Fight, furyGUID, 30_000, flameBreathID, "Flame Breath")
	f := death(in, CountedTime(in)).Finding
	if f == nil || f.Title != "Died at 0:30 to Flame Breath" || f.Severity != SeverityMajor {
		t.Fatalf("got %+v", f)
	}
	if f.Basis.Label != Fill(BasisAvoidable, map[string]string{"Spell": "Flame Breath"}) {
		t.Errorf("basis %q", f.Basis.Label)
	}
}

// With no killing blow the last hit names the death; with no hit at all nothing is claimed.
func TestDeath_KillingBlowFallsBackToTheLastHit(t *testing.T) {
	in := holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140)
	in.Fight.Deaths[len(in.Fight.Deaths)-1].KillingBlow = nil
	in.Fight.Deaths[len(in.Fight.Deaths)-1].Last = []summary.DamageRef{
		{AtMS: 198_000, SourceName: onyxiaName, SpellID: 15847, SpellName: "Tail Sweep", Amount: 900},
		{AtMS: 199_500, SourceName: onyxiaName, SpellID: flameBreathID, SpellName: "Flame Breath", Amount: 1_500},
	}
	if f := death(in, CountedTime(in)).Finding; f == nil || f.Title != "Died at 3:20 to Flame Breath" || !strings.Contains(f.Detail, "1,500 at 3:19") {
		t.Errorf("got %+v", f)
	}
	in.Fight.Deaths[len(in.Fight.Deaths)-1].Last = nil
	o := death(in, CountedTime(in))
	if o.Finding != nil || len(o.Strengths) != 0 {
		t.Errorf("unproven death claimed: %+v", o)
	}
	want := "Your death at 3:20: the log doesn't say what killed you."
	if len(o.CouldNotCheck) != 1 || o.CouldNotCheck[0] != want {
		t.Errorf("couldn't check %q, want %q", o.CouldNotCheck, want)
	}
}

// A narrowed defensive check says so (§20.4): no defensives list, or a level the log
// doesn't record; a defensive found ready, or a full check, says nothing.
func TestDeath_NarrowedDefensiveCheckSaysSo(t *testing.T) {
	cases := []struct {
		name string
		in   Input
		want []string
	}{
		{name: "no defensives list", in: holyDeath(t, 232_000, 200_000, flameBreathID, "Flame Breath", 2140),
			want: []string{"Defensive cooldowns: the coach doesn't know your spec's list yet."}},
		{name: "unknown level", in: f7Level(protDeathShieldWallReady(t, 232_000, 170_000), 0),
			want: []string{"Defensive cooldowns you never cast: this log doesn't tell us your level."}},
		{name: "known level, Shield Wall ready", in: protDeathShieldWallReady(t, 232_000, 170_000)},
		{name: "known level, nothing ready", in: f7Cast(protDeathShieldWallReady(t, 232_000, 170_000), shieldWallID, "Shield Wall", 100_000)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			o := death(tc.in, CountedTime(tc.in))
			if o.Finding == nil {
				t.Fatal("no finding")
			}
			if strings.Join(o.CouldNotCheck, "|") != strings.Join(tc.want, "|") {
				t.Errorf("couldn't check %q, want %q", o.CouldNotCheck, tc.want)
			}
		})
	}
}

// With a defensive ready, an avoidable killer is still said, in §6.2's order: killing
// blow, avoidable, defensive.
func TestDeath_ReadyKeepsTheAvoidableFact(t *testing.T) {
	notUsed := f7Died(protInput(t), 232_000, 170_000, tailSweepID, "Tail Sweep", 1_200)
	f := death(notUsed, CountedTime(notUsed)).Finding
	want := "Tail Sweep killed you at 2:50, and this boss lets you avoid it. Shield Wall was not used this pull."
	if f == nil || f.Title != "Died at 2:50, Shield Wall not used" || f.Detail != want {
		t.Fatalf("got %+v, want detail %q", f, want)
	}
	back := f7Died(protInput(t), 400_000, 200_000, tailSweepID, "Tail Sweep", 1_200)
	back = f7Cast(f7Cast(back, shieldWallID, "Shield Wall", 150_000), lastStandID, "Last Stand", 10_000)
	f = death(back, CountedTime(back)).Finding
	want = "Tail Sweep killed you at 3:20, and this boss lets you avoid it. Last Stand had been ready since 3:10."
	if f == nil || f.Title != "Died at 3:20 with Last Stand ready" || f.Detail != want {
		t.Fatalf("got %+v, want detail %q", f, want)
	}
}

// Task 1's golden (logs/engine/summary/testdata/forever-1.60.summary.json.golden) is the
// one real Forever log on disk. When its pull holds a death to an ability on its
// encounter's avoidable list, F7's title names that ability; it holds none today (no
// encounter, no death: Task 1's owner-provided boss pull is not on disk).
func TestDeath_KnownGoodAvoidableOnTheRealPull(t *testing.T) {
	raw, err := os.ReadFile("../../../logs/engine/summary/testdata/forever-1.60.summary.json.golden")
	if err != nil {
		t.Fatal(err)
	}
	var s summary.Summary
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatal(err)
	}
	table, ok := mechanics.Load(s.EncounterID)
	if !ok {
		t.Skip("Task 1's pull has no avoidable death (no curated encounter)")
	}
	seen := map[string]bool{}
	for _, d := range s.Deaths {
		first := !seen[d.GUID]
		seen[d.GUID] = true
		if !first || d.KillingBlow == nil {
			continue
		}
		m, ok := table.Lookup(d.KillingBlow.SpellID)
		if !ok || m.Kind != mechanics.Avoidable || m.Role != "" {
			continue
		}
		in := Input{Fight: s, Row: rosterRowOf(t, s, d.GUID), BossName: table.Name, Tables: Tables{Mechanics: &table}}
		f := death(in, CountedTime(in)).Finding
		if f == nil || !strings.Contains(f.Title, m.Name) {
			t.Fatalf("death %s at %d to %s: got %+v", d.GUID, d.AtMS, m.Name, f)
		}
		return
	}
	t.Skip("Task 1's pull has no avoidable death")
}

// Every further ready defensive after the first is named (A12, DetailAlsoReady): Shield
// Wall never cast on a pull longer than its cooldown, Last Stand cast and back.
func TestDeath_FurtherReadyDefensivesAreNamed(t *testing.T) {
	in := f7Cast(protDeathShieldWallReady(t, 1_200_000, 1_000_000), lastStandID, "Last Stand", 10_000)
	f := death(in, CountedTime(in)).Finding
	want := "Onyxia's melee killed you at 16:40. Shield Wall had been ready since 15:00. Also ready: Last Stand."
	if f == nil || f.Title != "Died at 16:40 with Shield Wall ready" || f.Detail != want {
		t.Fatalf("got %+v, want detail %q", f, want)
	}
	assertSpells(t, f.Facts, "Shield Wall", "Last Stand")
}

// A melee killing blow from a player's or a pet's GUID never carries its name (B1).
func TestDeath_APlayerOrPetKillerIsNeverNamed(t *testing.T) {
	cases := []struct{ guid, name, title string }{
		{"Player-1136-0A1B2C3D", "Zoë-Aegwynn-US", "Died at 3:20 to another player's melee"},
		{"Pet-0-4372-409-3-299-0103B0DC1E", "Fluffy", "Died at 3:20 to a pet"},
		{"Creature-0-4372-249-6-10184-000045E1A2", "Onyxia", "Died at 3:20 to Onyxia's melee"},
	}
	for _, c := range cases {
		in := withKiller(holyDeath(t, 232_000, 200_000, meleeSpellID, "Melee", 1_500), c.guid, c.name)
		f := death(in, CountedTime(in)).Finding
		if f == nil || f.Title != c.title {
			t.Fatalf("%s: got %+v, want %q", c.guid, f, c.title)
		}
	}
	in := withKiller(protDeathShieldWallReady(t, 232_000, 170_000), "Player-1136-0A1B2C3D", "Zoë-Aegwynn-US")
	if f := death(in, CountedTime(in)).Finding; f == nil || !strings.HasPrefix(f.Detail, "Another player's melee killed you at 2:50.") {
		t.Fatalf("the killer leads the sentence upper case: %+v", f)
	}
}

// withKiller is in with its player's last death's killing blow coming from guid, named name.
func withKiller(in Input, guid, name string) Input {
	in.Fight.Deaths = slices.Clone(in.Fight.Deaths)
	last := len(in.Fight.Deaths) - 1
	blow := *in.Fight.Deaths[last].KillingBlow
	blow.SourceGUID, blow.SourceName = guid, name
	in.Fight.Deaths[last].KillingBlow = &blow
	return in
}

// The pack built for a fight where a mind-controlled player kills the reviewed player
// carries no roster name or realm anywhere in its JSON (B1, §19).
func TestReview_AMindControlledKillerLeavesNoNameInThePack(t *testing.T) {
	in := holyDeath(t, 232_000, 100_000, meleeSpellID, "Melee", 1_500)
	in.BossName = onyxiaName
	killer := rosterRowOf(t, in.Fight, furyGUID)
	in = withKiller(in, killer.GUID, killer.Name)
	raw, err := json.Marshal(Review(in))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), KillerPlayerMelee) {
		t.Fatalf("the death finding is missing: %s", raw)
	}
	body := strings.ToLower(string(raw))
	for _, r := range in.Fight.Roster {
		for _, part := range append([]string{r.Name}, strings.Split(r.Name, "-")...) {
			if len(part) > 2 && strings.Contains(body, strings.ToLower(part)) {
				t.Errorf("%q (from a roster name) is in the pack", part)
			}
		}
	}
}
