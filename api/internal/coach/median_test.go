package coach

import (
	"reflect"
	"testing"
)

func cpmProfile(v float64) Profile {
	return Profile{CastsPerMin: map[string]float64{"Bloodthirst": v}}
}

func TestMedianProfile_OddAndEvenCounts(t *testing.T) {
	odd := MedianProfile([]Profile{cpmProfile(9), cpmProfile(1), cpmProfile(5)})
	if v, ok := odd.CastsPerMinOf("Bloodthirst"); !ok || v != 5 {
		t.Fatalf("odd median = %v, %v; want 5", v, ok)
	}
	even := MedianProfile([]Profile{cpmProfile(1), cpmProfile(2), cpmProfile(4), cpmProfile(100)})
	if v, _ := even.CastsPerMinOf("Bloodthirst"); v != 3 {
		t.Fatalf("even median = %v; want 3", v)
	}
}

func TestMedianProfile_KeyInSomeProfilesTakesTheirMedian(t *testing.T) {
	a := Profile{CastsPerMin: map[string]float64{"Bloodthirst": 10, "Whirlwind": 4}}
	b := Profile{CastsPerMin: map[string]float64{"Bloodthirst": 8}}
	c := Profile{CastsPerMin: map[string]float64{"Bloodthirst": 6, "Whirlwind": 8}}
	m := MedianProfile([]Profile{a, b, c})
	if v, _ := m.CastsPerMinOf("Whirlwind"); v != 6 {
		t.Fatalf("Whirlwind median over the two that carry it = %v; want 6", v)
	}
	if v, _ := m.CastsPerMinOf("Bloodthirst"); v != 8 {
		t.Fatalf("Bloodthirst median = %v; want 8", v)
	}
	if _, ok := m.CastsPerMinOf("Slam"); ok {
		t.Fatal("a key no profile carries must stay absent")
	}
}

func TestMedianProfile_ScalarsPointersAndFlags(t *testing.T) {
	adh1, adh3 := 0.5, 0.9
	first := int64(100_000)
	ps := []Profile{
		{CountedMS: 100, ActivePct: 90, CappedPct: 1, ManaLowMS: 10, Adherence: &adh1,
			Uptime: map[string]float64{"6673": 100}, Consumables: map[string]bool{"flask": true}},
		{CountedMS: 300, ActivePct: 70, CappedPct: 3, ManaLowMS: 30, Adherence: &adh3, ManaLowFirstMS: &first,
			OverhealByRank: map[string]float64{"10965": 20}, DefensiveUptime: map[string]float64{"2565": 30},
			Consumables: map[string]bool{"flask": true}},
		{CountedMS: 200, ActivePct: 80, CappedPct: 2, ManaLowMS: 20, Consumables: map[string]bool{"flask": false}},
	}
	m := MedianProfile(ps)
	if m.CountedMS != 200 || m.ActivePct != 80 || m.CappedPct != 2 || m.ManaLowMS != 20 {
		t.Fatalf("scalar medians wrong: %+v", m)
	}
	if v, ok := m.AdherenceOf(); !ok || v != 0.7 {
		t.Fatalf("adherence = %v, %v; want 0.7 over the two that carry it", v, ok)
	}
	if m.ManaLowFirstMS == nil || *m.ManaLowFirstMS != first {
		t.Fatalf("mana low first = %v", m.ManaLowFirstMS)
	}
	if v, _ := m.UptimeOf(6673); v != 100 {
		t.Fatalf("uptime = %v", v)
	}
	if !m.Consumables["flask"] {
		t.Fatal("a flask on two of three is the median player's")
	}
}

func TestMedianProfile_NoProfilesIsZero(t *testing.T) {
	if got := MedianProfile(nil); !reflect.DeepEqual(got, Profile{}) {
		t.Fatalf("empty median = %+v", got)
	}
}
