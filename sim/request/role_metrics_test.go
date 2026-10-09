package request

import (
	"errors"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/enginever"
)

func TestRoleMetricsPutsAHealerInTheHealProfilesFight(t *testing.T) {
	req := fury()
	req.Spec, req.Character.Class, req.Character.Talents = "priest-holy", "priest", ""
	req.EngineVersion = enginever.Version
	req.Iterations = 100
	req.RoleMetrics = true
	req.Encounter.DurationSec, req.Encounter.Variation = 120, 0.2
	got, err := BuildWith(req, Options{OpenIterations: true})
	if err != nil {
		t.Fatal(err)
	}
	profile, err := EmbeddedHealProfile()
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.RaidDamageModel == nil || got.Raid.RaidDamageModel.Profile != profile.ID {
		t.Errorf("raid damage model = %+v, want the profile %q", got.Raid.RaidDamageModel, profile.ID)
	}
	if got.Encounter.Duration != float64(profile.DurationSec) || got.Encounter.DurationVariation != 0 {
		t.Errorf("fight = %v +/- %v, want the profile's %d seconds exactly", got.Encounter.Duration, got.Encounter.DurationVariation, profile.DurationSec)
	}
}

func TestWithoutRoleMetricsAHealerIsAPlainRun(t *testing.T) {
	req := fury()
	req.Spec, req.Character.Class, req.Character.Talents = "priest-holy", "priest", ""
	req.Iterations = 100
	got, err := BuildWith(req, Options{OpenIterations: true})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.RaidDamageModel != nil {
		t.Error("a plain run carries no fake raid")
	}
}

func TestRoleMetricsLeavesATanksFightToTheTankFight(t *testing.T) {
	req := fury()
	req.Spec, req.Character.Class, req.Character.Talents = "warrior-protection", "warrior", ""
	req.Iterations = 100
	req.RoleMetrics = true
	got, err := BuildWith(req, Options{OpenIterations: true})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.RaidDamageModel != nil || len(got.Raid.Tanks) != 1 {
		t.Errorf("a tank fights the boss, not a fake raid: %+v", got.Raid)
	}
}

func TestRoleMetricsRefusesADamageSpec(t *testing.T) {
	req := fury()
	req.RoleMetrics = true
	if _, err := BuildWith(req, Options{OpenIterations: true}); !errors.Is(err, ErrRoleMetricsSpec) {
		t.Fatalf("err = %v, want ErrRoleMetricsSpec", err)
	}
}

func TestIsHealerSpecFollowsTheCanonicalRole(t *testing.T) {
	for spec, want := range map[string]bool{"priest-holy": true, "druid-restoration": true, "warrior-fury": false, "warrior-protection": false, "nope": false} {
		if got := IsHealerSpec(spec); got != want {
			t.Errorf("IsHealerSpec(%q) = %v, want %v", spec, got, want)
		}
	}
}
