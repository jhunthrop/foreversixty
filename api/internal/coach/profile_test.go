package coach

import (
	"encoding/json"
	"testing"
)

func TestProfileAccessors(t *testing.T) {
	adh := 0.9
	p := Profile{
		CastsPerMin:     map[string]float64{"Bloodthirst": 9.1},
		Uptime:          map[string]float64{"6673": 97},
		OverhealByRank:  map[string]float64{"10965": 22},
		DefensiveUptime: map[string]float64{"2565": 34},
		Adherence:       &adh,
	}
	present := []struct {
		name string
		v    float64
		ok   bool
		want float64
	}{
		{"CastsPerMinOf", first(p.CastsPerMinOf("Bloodthirst")), second(p.CastsPerMinOf("Bloodthirst")), 9.1},
		{"UptimeOf", first(p.UptimeOf(6673)), second(p.UptimeOf(6673)), 97},
		{"OverhealOf", first(p.OverhealOf(10965)), second(p.OverhealOf(10965)), 22},
		{"DefensiveOf", first(p.DefensiveOf(2565)), second(p.DefensiveOf(2565)), 34},
		{"AdherenceOf", first(p.AdherenceOf()), second(p.AdherenceOf()), 0.9},
	}
	for _, c := range present {
		if !c.ok || c.v != c.want {
			t.Errorf("%s = %v, %v; want %v, true", c.name, c.v, c.ok, c.want)
		}
	}
	var empty Profile
	absent := map[string]bool{
		"CastsPerMinOf": second(p.CastsPerMinOf("Whirlwind")), "UptimeOf": second(p.UptimeOf(1)),
		"OverhealOf": second(p.OverhealOf(1)), "DefensiveOf": second(empty.DefensiveOf(2565)),
		"AdherenceOf": second(empty.AdherenceOf()),
	}
	for name, ok := range absent {
		if ok {
			t.Errorf("%s of an absent key answered ok=true", name)
		}
	}
}

// The jsonb keys are §11.1's.
func TestProfileJSONNames(t *testing.T) {
	b, err := json.Marshal(Profile{CastsPerMin: map[string]float64{"Bloodthirst": 9.1}})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"alive_ms", "counted_ms", "active_pct", "cpm", "uptime", "capped_pct", "mana_low_ms", "overheal", "defensive", "consumables"} {
		if _, ok := m[k]; !ok {
			t.Errorf("profile json lacks %q: %s", k, b)
		}
	}
}

func first(v float64, _ bool) float64 { return v }
func second(_ float64, ok bool) bool  { return ok }
