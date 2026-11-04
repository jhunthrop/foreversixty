package request

import (
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
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
	const guideBuild = "005302011-320003-035131031321001151"
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
