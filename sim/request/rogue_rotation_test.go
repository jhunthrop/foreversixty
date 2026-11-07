package request

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	engine "github.com/wowsims/classic/sim"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// Subtlety's Ghostly Strike line says "on cooldown", but Hemorrhage (35
// energy) fired the moment energy reached 35, so energy never pooled to
// Ghostly Strike's 40 and it was cast 1.2 to 1.4 times a fight instead of
// about 9. Hemorrhage now holds back while Ghostly Strike is about to be
// ready, unless energy is high (round 9, physical lane).
func TestSubtletyHemorrhageWaitsForGhostlyStrike(t *testing.T) {
	const hemorrhage = 16511
	curated, err := loadLadderCurated(repoRoot, "rogue-subtlety")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != hemorrhage {
			continue
		}
		cond := protojson.Format(item.GetAction().GetCondition())
		if !strings.Contains(cond, "spellTimeToReady") || !strings.Contains(cond, "14278") {
			t.Errorf("Hemorrhage's condition does not wait for Ghostly Strike (14278): %s", cond)
		}
		return
	}
	t.Fatal("Subtlety casts no Hemorrhage")
}

// Ghostly Strike lets the next Garrote go in without Stealth for 7
// seconds (1329326), and Subtlety's Garrote line sits right below it. But
// Hemorrhage (30 energy) fired whenever Garrote's 40 was not yet there,
// so energy never pooled to 40 inside the window: Garrote went out about
// twice a fight (the opener and one more) instead of after every Ghostly
// Strike. Run on the guide build at 60, it must follow most of them
// (round 16).
func TestSubtletyGarrotesAfterEveryGhostlyStrike(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	const guideBuild = "005302011-320003-035230031321001151"
	req := smokeRequest(t, "rogue-subtlety")
	req.Character.Talents = guideBuild
	engineReq, err := BuildWith(req, Options{OpenIterations: true})
	if err != nil {
		t.Fatalf("building the request: %v", err)
	}
	res := core.RunRaidSim(engineReq)
	if err := adapter.ResultError(res); err != nil {
		t.Fatalf("the sim failed: %v", err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		t.Fatal(err)
	}
	casts := map[string]float64{}
	for _, c := range castSet(player) {
		casts[strings.SplitN(c.name, "/", 2)[0]] += c.casts
	}
	garrote, ghostly := casts["spell:11290"], casts["spell:14278"]
	if ghostly < 5 {
		t.Fatalf("Ghostly Strike cast %.2f times a fight; the check needs it on cooldown (%v)", ghostly, casts)
	}
	if garrote < 0.75*ghostly {
		t.Errorf("Garrote cast %.2f times a fight against %.2f Ghostly Strikes; Hemorrhage is starving the Garrote after each one", garrote, ghostly)
	}
}

// Death Mark comes back every 10 seconds, so holding it for a fifth combo
// point delays every mark after it: marking at four is ahead at every
// level it is cast (bare 40 +2.9%, 50 +0.8%, raid 60 +0.2% at seeds 1
// and 50001; round 16).
func TestSubtletyDeathMarksAtFourComboPoints(t *testing.T) {
	const deathMark = 1325758
	curated, err := loadLadderCurated(repoRoot, "rogue-subtlety")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != deathMark {
			continue
		}
		// protojson.Format randomizes its whitespace on purpose, so compare with it removed.
		cond := strings.Join(strings.Fields(protojson.Format(item.GetAction().GetCondition())), "")
		if !strings.Contains(cond, `"currentComboPoints"`) || !strings.Contains(cond, `"val":"4"`) {
			t.Errorf("Death Mark's condition does not mark at four combo points: %s", cond)
		}
		return
	}
	t.Fatal("Subtlety casts no Death Mark")
}

// subtletyCondition is the whitespace-free condition of Subtlety's first
// curated line casting spellID.
func subtletyCondition(t *testing.T, spellID int32) string {
	t.Helper()
	return specCondition(t, "rogue-subtlety", spellID)
}

// specCondition is spellID's condition on spec's curated card, compacted.
func specCondition(t *testing.T, spec string, spellID int32) string {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() == spellID {
			return strings.Join(strings.Fields(protojson.Format(item.GetAction().GetCondition())), "")
		}
	}
	t.Fatalf("%s casts no %d", spec, spellID)
	return ""
}

// atMostFourComboPoints is the compacted `currentComboPoints <= 4` test.
const atMostFourComboPoints = `{"cmp":{"op":"OpLe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"4"}}}}`

// Hemorrhage at five combo points spills its point when no finisher is
// ready (Rupture ticking, Death Mark cooling): gating it at four or fewer
// gains 0.6% at raid 60 on both factions and ties at bare 30 to 60
// (round 22, guide r24 sim request).
func TestSubtletyHemorrhageNeverAtFiveComboPoints(t *testing.T) {
	if cond := subtletyCondition(t, 16511); !strings.Contains(cond, atMostFourComboPoints) {
		t.Errorf("Hemorrhage's condition does not stop at five combo points: %s", cond)
	}
}

// From 17 Assassination puts Slice and Dice back up the moment it falls off,
// at one or more combo points while 6 seconds of fight remain, instead of at
// four with under 3 seconds left or at three once down (MythicSim comparison,
// 2026-11-07, 10,000 iterations, seeds 1 and 50001: raid 60 +0.25%, bare 50
// +0.50%, bare 60 +0.37%, bare 20 to 40 within error). Their Eviscerate at
// three is held: it lost on every bare band.
func TestAssassinationSliceAndDiceWheneverItDrops(t *testing.T) {
	cond := specCondition(t, "rogue-assassination", 6774)
	for _, want := range []string{
		`{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"1"}}}}`,
		`{"cmp":{"op":"OpGe","lhs":{"remainingTime":{}},"rhs":{"const":{"val":"6"}}}}`,
	} {
		if !strings.Contains(cond, want) {
			t.Errorf("Slice and Dice's condition lacks %s: %s", want, cond)
		}
	}
	if evisc := specCondition(t, "rogue-assassination", 11300); strings.Contains(evisc, `"currentComboPoints":{}},"rhs":{"const":{"val":"3"}}`) {
		t.Errorf("Eviscerate spends at three, which lost on every bare band: %s", evisc)
	}
}

// MythicSim comparison (2026-11-07): Combat Shanks only under 80 energy and
// Eviscerates at four from 30 (five below: four lost 0.6% at band 20). With
// the 60-only Slice and Dice rule: raid 60 +1.25%, bare 30 to 60 +0.25 to
// +0.91%, bare 20 unchanged.
func TestCombatShankUnderEightyAndEviscerateAtFourFromThirty(t *testing.T) {
	if cond := specCondition(t, "rogue-combat", 1329549); !strings.Contains(cond, `{"cmp":{"op":"OpLt","lhs":{"currentEnergy":{}},"rhs":{"const":{"val":"80"}}}}`) {
		t.Errorf("Shank is not held at 80 energy: %s", cond)
	}
	cond := specCondition(t, "rogue-combat", 11300)
	for _, want := range []string{
		`{"cmp":{"op":"OpGe","lhs":{"currentLevel":{}},"rhs":{"const":{"val":"30"}}}}`,
		`{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"4"}}}}`,
		`{"cmp":{"op":"OpGe","lhs":{"currentComboPoints":{}},"rhs":{"const":{"val":"5"}}}}`,
	} {
		if !strings.Contains(cond, want) {
			t.Errorf("Eviscerate's condition lacks %s: %s", want, cond)
		}
	}
}

// The 31/20/0 raid build (Coup de Grace, Aggression, Flawless Execution)
// makes Eviscerate cheaper and harder than a Rupture is worth: the card's
// Blood Splatter Rupture line lost 1.5% at raid 60 on both factions, where
// it had gained 0.8% on the old raid build (round 23). The leveling build
// never took Blood Splatter, so the line only ever fired in the raid.
func TestAssassinationCastsNoRupture(t *testing.T) {
	curated, err := loadLadderCurated(repoRoot, "rogue-assassination")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if id := item.GetAction().GetCastSpell().GetSpellId().GetSpellId(); id == 11275 {
			t.Errorf("Assassination's card still casts Rupture: %s", protojson.Format(item))
		}
	}
}

// Holding Shank at four or five combo points cost its energy on every
// leveling band (bare 30 to 59: +0.1 to +1.4% with the hold lifted) and
// only paid at 60, where Death Mark's top rank makes every five-point
// finisher count. Holding it only with 50 energy or more is ahead or level
// everywhere, raid 60 included (round 21).
func TestSubtletyShanksWhenShortOnEnergy(t *testing.T) {
	cond := subtletyCondition(t, 1329549)
	if !strings.Contains(cond, `"currentEnergy"`) || !strings.Contains(cond, `"val":"50"`) {
		t.Errorf("Shank's condition does not let it go below 50 energy: %s", cond)
	}
}

// Rupture trains at 20, so a finisher gated on Rupture ticking named a
// condition no rogue below 20 can meet. Both finishers carry a
// before-Rupture branch (SIM-R11-1, carried since round 11).
func TestSubtletyFinishersBeforeRupture(t *testing.T) {
	for _, id := range []int32{11300, 6774} {
		if cond := subtletyCondition(t, id); !strings.Contains(cond, `"spellIsKnown":{"spellId":{"spellId":11275`) {
			t.Errorf("spell %d has no branch for a rogue without Rupture: %s", id, cond)
		}
	}
}

// Before Mutilate (40) and Hemorrhage (30), Assassination and Subtlety
// built with Sinister Strike even with a dagger in the main hand, behind
// a boss, where Backstab is the bigger hit (and Opportunity, which
// Subtlety's leveling build takes at 13, raises Backstab, not Sinister
// Strike). Combat already backstabs at every level. With a dagger, both
// must build with Backstab (round 17).
func TestDaggerLevelingBuildsWithBackstab(t *testing.T) {
	registerEngine.Do(engine.RegisterAll)
	const level = 25
	for _, spec := range []string{"rogue-assassination", "rogue-subtlety"} {
		req := smokeRequest(t, spec)
		req.Character.Level = level
		req.Character.Talents = ""
		// Blackfang (dagger) in the main hand, Diamond Hammer in the off hand.
		req.Character.Gear = []api.GearSlot{{Slot: "main_hand", ItemID: 2236}, {Slot: "off_hand", ItemID: 2194}}
		engineReq, err := BuildWith(req, Options{OpenIterations: true})
		if err != nil {
			t.Fatalf("%s: building the request: %v", spec, err)
		}
		if err := simdb.Attach(engineReq); err != nil {
			t.Fatalf("%s: attaching the item database: %v", spec, err)
		}
		res := core.RunRaidSim(engineReq)
		if err := adapter.ResultError(res); err != nil {
			t.Fatalf("%s: the sim failed: %v", spec, err)
		}
		player, err := adapter.PlayerMetrics(res)
		if err != nil {
			t.Fatal(err)
		}
		casts := map[string]float64{}
		for _, c := range castSet(player) {
			casts[strings.SplitN(c.name, "/", 2)[0]] += c.casts
		}
		backstab, sinister := casts["spell:2591"]+casts["spell:2590"], casts["spell:1759"]+casts["spell:1758"]
		if backstab < 15 || sinister > backstab/5 {
			t.Errorf("%s at %d with a dagger: Backstab %.2f casts, Sinister Strike %.2f (%v)", spec, level, backstab, sinister, casts)
		}
	}
}

// Subtlety's build put its last tier-two point in Improved Ambush, which
// only raises Ambush's crit: the rotation never casts Ambush (a Garrote
// opener, a Garrote after every Ghostly Strike), and the published level
// 60 weapons are not daggers, so Ambush cannot even be pressed. Dirty
// Deeds' second point sits in the same tier, keeps the points total that
// opens Boot Knife and Death Mark, and takes 10 more energy off every
// Garrote, about nine a fight (round 17).
func TestSubtletyBuildSpendsNoPointOnAmbush(t *testing.T) {
	const (
		dirtyDeeds     = 3 // Subtlety tree positions, talents/rogue.json 1.60.1.70755
		improvedAmbush = 5
	)
	_, trees, err := leveling.GuideBuildTalents(repoRoot, "rogue", "subtlety")
	if err != nil {
		t.Fatal(err)
	}
	sub := trees[2]
	if got := sub[improvedAmbush]; got != '0' {
		t.Errorf("the Subtlety build spends %c points in Improved Ambush, which the rotation never uses (%s)", got, sub)
	}
	if got := sub[dirtyDeeds]; got != '2' {
		t.Errorf("the Subtlety build has Dirty Deeds at %c of 2 (%s)", got, sub)
	}
}
