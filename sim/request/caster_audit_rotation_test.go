package request

import (
	"slices"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/wowsims/classic/sim/core/proto"
)

// The 2026-10-31 caster theorycraft audit
// (design/reviews/2026-10-31-theorycraft-audit/casters.md) found rotation
// lines that disagreed with the build or with Forever theorycraft. These
// pin the corrected lines so a later rotation search cannot quietly undo
// them without a decision.

// castOrder lists the spell ids of a curated rotation's top-level castSpell
// lines, in priority order.
func castOrder(t *testing.T, spec string) []int32 {
	t.Helper()
	curated, err := loadLadderCurated(repoRoot, spec)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatalf("%s: %v", spec, err)
	}
	var ids []int32
	for _, item := range rot.PriorityList {
		if cast := item.GetAction().GetCastSpell(); cast != nil {
			ids = append(ids, cast.GetSpellId().GetSpellId())
		}
		if channel := item.GetAction().GetChannelSpell(); channel != nil {
			ids = append(ids, channel.GetSpellId().GetSpellId())
		}
	}
	return ids
}

func requireBefore(t *testing.T, spec string, ids []int32, first, second int32) {
	t.Helper()
	i, j := slices.Index(ids, first), slices.Index(ids, second)
	if i < 0 || j < 0 {
		t.Fatalf("%s: want both %d and %d cast, got %v", spec, first, second, ids)
	}
	if i > j {
		t.Errorf("%s: %d must sit above %d, got %v", spec, first, second, ids)
	}
}

// W1, W2: the Affliction build takes Siphon Life, so the rotation keeps it
// up. Since the 10 Nov hotfix (sim-accuracy 2026-11-11), Wrack speeds up
// Corruption, Siphon Life and Bane of Agony and Soul Siphon counts only those
// three, so Agony is the bane all fight and Bane of Doom is never cast.
func TestAfflictionKeepsSiphonLifeAndAgony(t *testing.T) {
	const (
		siphonLife = 18881
		agony      = 11713
		doom       = 603
	)
	ids := castOrder(t, "warlock-affliction")
	for _, id := range []int32{siphonLife, agony} {
		if !slices.Contains(ids, id) {
			t.Errorf("Affliction never casts %d: %v", id, ids)
		}
	}
	if slices.Contains(ids, doom) {
		t.Errorf("Affliction casts Bane of Doom (%d), which Wrack and Soul Siphon leave behind: %v", doom, ids)
	}
}

// D2: Demonic Pact is real and modeled. The Demonology raid build takes it,
// and the prepull summons the Imp, sacrifices it, then summons the Succubus
// before Soul Link; Shadow Bolt is the filler under both Shadow bonuses.
func TestDemonologySacrificesTheImpUnderDemonicPact(t *testing.T) {
	const (
		summonImp         = 688
		demonicSacrifice  = 18788
		summonSuccubus    = 712
		soulLink          = 19028
		shadowBolt        = 11661
		searingPain       = 17923
		demonologySpecKey = "warlock-demonology"
	)
	curated, err := loadLadderCurated(repoRoot, demonologySpecKey)
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	var prepull []int32
	for _, p := range rot.PrepullActions {
		prepull = append(prepull, p.GetAction().GetCastSpell().GetSpellId().GetSpellId())
	}
	want := []int32{summonImp, demonicSacrifice, summonSuccubus, soulLink}
	if !slices.Equal(prepull, want) {
		t.Errorf("Demonology prepull = %v, want %v", prepull, want)
	}
	ids := castOrder(t, demonologySpecKey)
	if !slices.Contains(ids, shadowBolt) || slices.Contains(ids, searingPain) {
		t.Errorf("Demonology's filler must be Shadow Bolt (%d), not Searing Pain (%d): %v", shadowBolt, searingPain, ids)
	}
}

// P1, P2: Mind Blast sits above Mind Flay, and Devouring Plague is kept up
// by its own line rather than only behind Inner Focus.
func TestShadowCastsMindBlastBeforeMindFlayAndKeepsDevouringPlague(t *testing.T) {
	const (
		mindBlast       = 10947
		mindFlay        = 18807
		devouringPlague = 19280
		shadowSpecKey   = "priest-shadow"
	)
	ids := castOrder(t, shadowSpecKey)
	requireBefore(t, shadowSpecKey, ids, mindBlast, mindFlay)
	if !slices.Contains(ids, devouringPlague) {
		t.Errorf("Shadow has no standalone Devouring Plague line: %v", ids)
	}
}

// R1: Frost casts Arcane Missiles on Missile Barrage procs; A4: Arcane
// stacks Arcane Blast to 3 before its Frostbolt filler, and since client
// 1.60.1.70707 (+100% cost a stack, 18 s) keeps blasting past 3 while the
// mana bar is ahead of the fight.
func TestMageMissileBarrageAndArcaneBlastStacks(t *testing.T) {
	// Rank 7: rank 8 (25345) is an Ahn'Qiraj book (sim-accuracy round 4).
	const arcaneMissiles = 10212
	if ids := castOrder(t, "mage-frost"); !slices.Contains(ids, arcaneMissiles) {
		t.Errorf("Frost never casts Arcane Missiles on Missile Barrage: %v", ids)
	}
	curated, err := loadLadderCurated(repoRoot, "mage-arcane")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		either := item.GetAction().GetCondition().GetOr().GetVals()
		if len(either) != 2 || either[0].GetCmp().GetLhs().GetAuraNumStacks() == nil {
			continue
		}
		if got := either[0].GetCmp().GetRhs().GetConst().GetVal(); got != "3" {
			t.Errorf("Arcane Blast stacks to %s before Frostbolt, want 3", got)
		}
		if !conditionPacesMana(either[1]) {
			t.Error("Arcane Blast's past-3 branch does not compare mana against the fight's remaining time")
		}
		return
	}
	t.Error("the Arcane rotation has no Arcane Blast stack condition")
}

// conditionPacesMana reports whether v (or one of the values it ands)
// compares currentManaPercent against remainingTimePercent.
func conditionPacesMana(v *proto.APLValue) bool {
	if cmp := v.GetCmp(); cmp != nil {
		return cmp.GetLhs().GetCurrentManaPercent() != nil && cmp.GetRhs().GetRemainingTimePercent() != nil
	}
	for _, inner := range v.GetAnd().GetVals() {
		if conditionPacesMana(inner) {
			return true
		}
	}
	return false
}

// Round 11 (design/reviews/2026-11-03-sim-audit-round11/casters.md S1): the
// published horde Shadow priest is Undead, whose Dark Sacrifice (ranks
// 1277324-1277328, 20 to 60, SkillLineAbility race mask 16) trades health
// for mana. The bare bands run dry, so the rotation casts it when mana is
// low and behind the fight's pace, above the damage lines.
func TestShadowCastsDarkSacrificeWhenManaIsBehind(t *testing.T) {
	const (
		darkSacrifice = 1277328
		mindBlast     = 10947
	)
	ids := castOrder(t, "priest-shadow")
	requireBefore(t, "priest-shadow", ids, darkSacrifice, mindBlast)
	curated, err := loadLadderCurated(repoRoot, "priest-shadow")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	for _, item := range rot.PriorityList {
		if item.GetAction().GetCastSpell().GetSpellId().GetSpellId() != darkSacrifice {
			continue
		}
		if !conditionPacesMana(item.GetAction().GetCondition()) {
			t.Error("Dark Sacrifice does not compare mana against the fight's remaining time")
		}
	}
}

// Sim audit round 13: Balance opens like the other casters, with a Wrath
// timed to finish on the pull (1.5 s with Improved Wrath 5/5, then its
// flight), instead of starting its first cast at the pull.
func TestBalanceOpensWithAPrePullWrath(t *testing.T) {
	const wrathRank8 = 9912
	curated, err := loadLadderCurated(repoRoot, "druid-balance")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	if len(rot.PrepullActions) != 1 {
		t.Fatalf("Balance prepull has %d lines, want one Wrath", len(rot.PrepullActions))
	}
	p := rot.PrepullActions[0]
	if got := p.GetAction().GetCastSpell().GetSpellId().GetSpellId(); got != wrathRank8 {
		t.Errorf("Balance prepull casts %d, want Wrath rank 8 (%d)", got, wrathRank8)
	}
	if got := p.GetDoAtValue().GetConst().GetVal(); got != "-1.5s" {
		t.Errorf("Balance prepull Wrath at %q, want -1.5s (its cast ends on the pull)", got)
	}
}

// Round 13: the Shadow guide keeps Vampiric Embrace up on the boss ("On the
// boss, keep it up: it costs a global every 30 seconds"), so the list casts
// it, above Mind Flay, its filler.
func TestShadowKeepsVampiricEmbraceUp(t *testing.T) {
	const (
		vampiricEmbrace = 15286
		mindFlay        = 18807
		shadowSpecKey   = "priest-shadow"
	)
	ids := castOrder(t, shadowSpecKey)
	requireBefore(t, shadowSpecKey, ids, vampiricEmbrace, mindFlay)
}

// Round 15 (priest): Shadow opens with a Mind Blast cast before the pull, timed
// to land on it (1.5 s cast), after Shadowform at -20 s. Both Forever Shadow
// guides open this way (Warcraft Tavern Forever Shadow guide; Mobalytics Forever
// Shadow guide, 2026-11-02); raid 60 gains about 0.8% and no bare band loses.
func TestShadowOpensWithAPrePullMindBlast(t *testing.T) {
	const (
		shadowform    = 15473
		mindBlastRank = 10947
	)
	curated, err := loadLadderCurated(repoRoot, "priest-shadow")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, p := range rot.PrepullActions {
		got = append(got, protojson.Format(p.GetAction().GetCastSpell().GetSpellId())+"@"+p.GetDoAtValue().GetConst().GetVal())
	}
	if len(rot.PrepullActions) != 2 ||
		rot.PrepullActions[0].GetAction().GetCastSpell().GetSpellId().GetSpellId() != shadowform ||
		rot.PrepullActions[1].GetAction().GetCastSpell().GetSpellId().GetSpellId() != mindBlastRank ||
		rot.PrepullActions[1].GetDoAtValue().GetConst().GetVal() != "-1.5s" {
		t.Errorf("Shadow prepull %v, want Shadowform then Mind Blast rank 9 at -1.5s", got)
	}
}

// Round 17 (priest): Discipline puts Renew on the tank before the pull,
// after Lightwell, at the rank its tank Renew line casts. Raid 60 gains 0.2
// to 0.4% and bare 60 about 1% (seeds 1 and 50001, both factions); band 20 is
// unchanged and no band loses. Round 15 held it only because band 20 lost.
// Round 19: Lightwell at -3 s and Renew at -1.5 s, so Renew's global is over
// at the pull (at -1 s it ran to +0.5 s and held the first heal back).
func TestDisciplineRenewsTheTankBeforeThePull(t *testing.T) {
	const (
		lightwell  = 27871
		renewRank7 = 10927
		tankIndex  = 5
	)
	curated, err := loadLadderCurated(repoRoot, "priest-discipline")
	if err != nil {
		t.Fatal(err)
	}
	rot := &proto.APLRotation{}
	if err := protojson.Unmarshal(curated.Rotation, rot); err != nil {
		t.Fatal(err)
	}
	pre := rot.PrepullActions
	if len(pre) != 2 ||
		pre[0].GetAction().GetCastSpell().GetSpellId().GetSpellId() != lightwell {
		t.Fatalf("Discipline prepull has %d lines, want Lightwell then Renew", len(pre))
	}
	renew := pre[1].GetAction().GetCastSpell()
	if renew.GetSpellId().GetSpellId() != renewRank7 || renew.GetSpellId().GetRank() != 7 ||
		renew.GetTarget().GetType() != proto.UnitReference_Player || renew.GetTarget().GetIndex() != tankIndex ||
		pre[0].GetDoAtValue().GetConst().GetVal() != "-3s" ||
		pre[1].GetDoAtValue().GetConst().GetVal() != "-1.5s" {
		t.Errorf("Discipline second prepull line %s at %s, want Renew rank 7 on the tank (player %d) at -1.5s after Lightwell at -3s",
			protojson.Format(renew), pre[1].GetDoAtValue().GetConst().GetVal(), tankIndex)
	}
}
