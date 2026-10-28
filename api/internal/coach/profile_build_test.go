package coach

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const profileSizeCap = 2048

func TestBuildProfile_MeasuresWhatThePackCompares(t *testing.T) {
	p := BuildProfile(furyFight(t, 60_000, 23881, []int64{0, 6000, 12000, 18000, 24000}, repeat(500, 60)))
	if got, ok := p.CastsPerMinOf("Bloodthirst"); !ok || math.Abs(got-5.0) > 0.01 {
		t.Fatalf("cpm %v %v", got, ok)
	}
	if p.Consumables == nil {
		t.Fatal("consumables at the pull belong in the profile")
	}
	if p.CountedMS != 60_000 || p.CappedPct != 0 {
		t.Fatalf("counted %d capped %v, want 60000 and 0", p.CountedMS, p.CappedPct)
	}
}

func TestBuildProfile_CappedShareIsOfCountedTime(t *testing.T) {
	rage := append(repeat(fixtureRageMax, 12), repeat(500, 48)...)
	p := BuildProfile(furyFight(t, 60_000, 23881, []int64{0}, rage))
	if math.Abs(p.CappedPct-20) > 0.01 {
		t.Fatalf("capped %v, want 20", p.CappedPct)
	}
}

func TestBuildProfile_SharedDebuffIsLeftOut(t *testing.T) {
	in := furyFight(t, 60_000, 23881, []int64{0}, repeat(500, 60))
	in.Fight = withAura(in.Fight, "Creature-boss", 11597, "Sunder Armor",
		[]string{furyGUID, "other-1", "other-2"}, span(0, 50_000))
	in.Fight = withAura(in.Fight, furyGUID, 6673, "Battle Shout", []string{furyGUID}, span(0, 30_000))
	p := BuildProfile(in)
	if _, ok := p.UptimeOf(11597); ok {
		t.Fatal("a debuff three warriors applied has no own uptime")
	}
	if got, ok := p.UptimeOf(6673); !ok || math.Abs(got-50) > 0.5 {
		t.Fatalf("battle shout uptime %v %v, want 50", got, ok)
	}
}

func TestBuildProfile_IsAboutOneKilobyte(t *testing.T) {
	p := BuildProfile(furyFight(t, 60_000, 23881, []int64{0, 6000, 12000}, repeat(500, 60)))
	raw, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) > profileSizeCap {
		t.Fatalf("profile is %d bytes, cap %d", len(raw), profileSizeCap)
	}
}

func TestBuildProfile_HealerOverhealByRankNeedsTenCasts(t *testing.T) {
	in := holyInput(t)
	in.Fight = withCasts(in.Fight, priestGUID, 10965, "Greater Heal", seqEvery(0, 3000, 12)...)
	in.Fight = withCasts(in.Fight, priestGUID, 10964, "Greater Heal", 0, 3000)
	in.Fight.Healing = []summary.Actor{{GUID: priestGUID, Abilities: []summary.Ability{
		{SpellID: 10965, Total: 1000, Overheal: 400}, {SpellID: 10964, Total: 1000, Overheal: 900},
	}}}
	p := BuildProfile(in)
	if got, ok := p.OverhealOf(10965); !ok || math.Abs(got-40) > 0.01 {
		t.Fatalf("overheal %v %v, want 40", got, ok)
	}
	if _, ok := p.OverhealOf(10964); ok {
		t.Fatal("a rank cast twice is too thin to compare")
	}
}
