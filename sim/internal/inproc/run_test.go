package inproc

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/leveling"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/score"
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
)

const (
	realRoot   = "../../.."
	testIters  = 3
	bandLevel  = 60
	randomSeed = 7
)

// bandRequest is the committed level-60 bare alliance BiS entry of spec as a
// request: the same gear, race and talents the ranker published.
func bandRequest(t *testing.T, spec string) api.SimRequest {
	t.Helper()
	build, err := leveling.ReadActiveBuild(realRoot)
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(filepath.Join(leveling.BisDir(filepath.Join(realRoot, "data", "builds", build)), spec+".json"))
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Bands []struct {
			Band    int    `json:"band"`
			Preset  string `json:"preset"`
			Faction string `json:"faction"`
			Race    string `json:"race"`
			Talents string `json:"talents"`
			Slots   []struct {
				Slot   string `json:"slot"`
				ItemID int    `json:"item_id"`
			} `json:"slots"`
		} `json:"bands"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	info := specs.ByKey[spec]
	for _, b := range file.Bands {
		if b.Band != bandLevel || b.Preset != request.BarePreset || b.Faction != "alliance" {
			continue
		}
		var gear []api.GearSlot
		for _, s := range b.Slots {
			// Until the nightly ranks a new build the file is the previous build's, and
			// an item the new build dropped would panic the engine instead of unequip.
			if _, known := simdb.Lookup(s.ItemID); s.ItemID != 0 && known {
				gear = append(gear, api.GearSlot{Slot: s.Slot, ItemID: s.ItemID})
			}
		}
		return api.SimRequest{
			EngineVersion: enginever.Version,
			Spec:          spec,
			Source:        api.CharacterSource{Kind: api.SourceBuild},
			Character: api.CharacterSpec{
				Name: "inproc-test", Race: b.Race, Class: info.ClassSlug, Level: bandLevel,
				Talents: b.Talents, Gear: gear,
				Buffs: leveling.KitBuffs(spec, bandLevel), Consumes: leveling.KitConsumes(spec, bandLevel),
			},
			Encounter:  api.DefaultEncounter(),
			Iterations: testIters,
			RandomSeed: randomSeed,
		}
	}
	t.Fatalf("%s has no level-60 bare alliance band", spec)
	return api.SimRequest{}
}

func TestPlainDPSRunsExactlyTheRequestedIterationsDeterministically(t *testing.T) {
	req := bandRequest(t, "rogue-combat")
	a, err := PlainDPS(req)
	if err != nil {
		t.Fatal(err)
	}
	b, err := PlainDPS(req)
	if err != nil {
		t.Fatal(err)
	}
	if a.Mean <= 0 || a != b {
		t.Fatalf("same seed must reproduce: %+v vs %+v", a, b)
	}
}

func TestPlainRunReturnsThePlayersMetricsAndAnEmptyRotationOverrideChangesNothing(t *testing.T) {
	req := bandRequest(t, "rogue-combat")
	est, player, err := PlainRunWithRotation(req, nil)
	if err != nil {
		t.Fatal(err)
	}
	if player == nil || len(player.GetActions()) == 0 {
		t.Fatal("no player metrics came back")
	}
	viaWrapper, err := PlainDPSWithRotation(req, nil)
	if err != nil || viaWrapper != est {
		t.Fatalf("PlainDPSWithRotation(nil) = %+v %v, want %+v", viaWrapper, err, est)
	}
}

func TestPlainRunWithRotationUsesTheGivenPriorityList(t *testing.T) {
	req := bandRequest(t, "rogue-combat")
	idle := &proto.APLRotation{Type: proto.APLRotation_TypeAPL}
	base, err := PlainDPS(req)
	if err != nil {
		t.Fatal(err)
	}
	got, err := PlainDPSWithRotation(req, idle)
	if err != nil {
		t.Fatal(err)
	}
	if got.Mean >= base.Mean {
		t.Fatalf("an empty rotation (%.1f DPS) must do less than the curated one (%.1f)", got.Mean, base.Mean)
	}
}

func TestPlainDPSRejectsARequestTheBuilderRefuses(t *testing.T) {
	req := bandRequest(t, "rogue-combat")
	req.Spec = "no-such-spec"
	if _, err := PlainDPS(req); err == nil {
		t.Fatal("an unknown spec ran")
	}
}

func TestTankRunReportsTheTanksFightAndItsHealth(t *testing.T) {
	req := bandRequest(t, "warrior-protection")
	got, err := TankRun(req)
	if err != nil {
		t.Fatal(err)
	}
	if got.Health <= 0 || got.DTPS.Mean <= 0 || got.ChanceOfDeath < 0 || got.ChanceOfDeath > 1 {
		t.Fatalf("implausible tank result: %+v", got)
	}
	engineReq, err := request.BuildWith(req, buildOptions)
	if err != nil {
		t.Fatal(err)
	}
	max, err := score.MaxHealth(engineReq.GetRaid(), engineReq.GetEncounter())
	if err != nil || max != got.Health {
		t.Fatalf("MaxHealth = %v %v, want the run's %v", max, err, got.Health)
	}
}

func TestTankRunRejectsABadRequest(t *testing.T) {
	req := bandRequest(t, "warrior-protection")
	req.Spec = "no-such-spec"
	if _, err := TankRun(req); err == nil {
		t.Fatal("an unknown spec ran")
	}
}

func TestHealingRunReportsAHealersThroughput(t *testing.T) {
	req := bandRequest(t, "priest-holy")
	profile, err := request.LoadHealProfile(filepath.Join(realRoot, "data", "curated", "heal-profile.json"))
	if err != nil {
		t.Fatal(err)
	}
	got, err := HealingRun(req, profile)
	if err != nil {
		t.Fatal(err)
	}
	if got.Effective.Mean <= 0 || got.Raw.Mean < got.Effective.Mean || got.ManaSpent <= 0 {
		t.Fatalf("implausible healer result: %+v", got)
	}
	if share := got.OverhealShare(); share < 0 || share > 1 {
		t.Fatalf("overheal share %v out of range", share)
	}
	req.Spec = "no-such-spec"
	if _, err := HealingRun(req, profile); err == nil {
		t.Fatal("an unknown spec ran")
	}
}
