// logs/engine/mechanics/weights/roles_test.go
package weights

import "testing"

func TestDefaultMatchesTheSpecsTable(t *testing.T) {
	d := Default()
	want := Roles{
		DPS:    RoleWeights{Output: 35, Survival: 15, Mechanics: 20, Utility: 15, Preparation: 10, Activity: 5},
		Healer: RoleWeights{Output: 30, Survival: 10, Mechanics: 20, Utility: 20, Preparation: 10, Activity: 10},
		Tank:   RoleWeights{Output: 10, Survival: 35, Mechanics: 20, Utility: 20, Preparation: 10, Activity: 5},
	}
	if d != want {
		t.Fatalf("Default() = %+v, want %+v", d, want)
	}
}

func TestForSelectsByRoleAndReportsAnUnknownOne(t *testing.T) {
	d := Default()
	if w, ok := d.For("tank"); !ok || w != d.Tank {
		t.Fatalf("For(tank) = %+v, %v", w, ok)
	}
	if w, ok := d.For("healer"); !ok || w != d.Healer {
		t.Fatalf("For(healer) = %+v, %v", w, ok)
	}
	if w, ok := d.For("dps"); !ok || w != d.DPS {
		t.Fatalf("For(dps) = %+v, %v", w, ok)
	}
	if _, ok := d.For("raid-lead"); ok {
		t.Fatal("an unrecognised role must report false")
	}
}

func TestParseRefusesARoleWhoseWeightsDoNotSumToOneHundred(t *testing.T) {
	bad := []byte(`{"dps":{"output":35,"survival":15,"mechanics":20,"utility":15,"preparation":10,"activity":10},
		"healer":{"output":30,"survival":10,"mechanics":20,"utility":20,"preparation":10,"activity":10},
		"tank":{"output":10,"survival":35,"mechanics":20,"utility":20,"preparation":10,"activity":5}}`)
	if _, err := Parse(bad); err == nil {
		t.Fatal("dps weights summing to 105 must be refused")
	}
}

func TestSum(t *testing.T) {
	if got := Default().DPS.Sum(); got != 100 {
		t.Fatalf("DPS.Sum() = %v, want 100", got)
	}
}
