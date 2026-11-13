package request

import (
	"errors"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

const curatedRaidHealProfile = "../../data/curated/" + RaidHealProfileFile

func mustRaidHealProfile(t *testing.T) RaidHealProfile {
	t.Helper()
	profile, err := LoadRaidHealProfile(curatedRaidHealProfile)
	if err != nil {
		t.Fatalf("the curated raid heal profile does not load: %v", err)
	}
	return profile
}

func TestCuratedRaidHealProfileIsTheLogSamplesDerivation(t *testing.T) {
	profile := mustRaidHealProfile(t)
	s := profile.LogSample
	if profile.DurationSec < int(s.KillSecondsP90) || profile.DurationSec >= int(s.KillSecondsP90)+raidFightRoundSeconds {
		t.Errorf("fight %d s is not the logged 90th percentile kill (%v s) rounded up", profile.DurationSec, s.KillSecondsP90)
	}
	perSecond := profile.Tank.HitDamage/profile.Tank.SwingSeconds +
		profile.Pulse.Damage*float64(profile.Pulse.Members)/profile.Pulse.IntervalSeconds
	if math.Abs(perSecond-s.HealerDamagePerSecond) > 1 {
		t.Errorf("the fight deals %.1f a second, the logged healer's share is %.1f", perSecond, s.HealerDamagePerSecond)
	}
}

// The logs' spikiness (collect.py's spikiness): one-second bins of the tank's damage, the
// largest five-second sum over the mean five-second sum, the median over many fights.
func TestRaidHealProfileTankSpikeMatchesTheLogs(t *testing.T) {
	profile := mustRaidHealProfile(t)
	rng := rand.New(rand.NewSource(1))
	const fights = 2000
	spikes := make([]float64, 0, fights)
	for range fights {
		spikes = append(spikes, tankSpike5(profile, rng))
	}
	sort.Float64s(spikes)
	median := spikes[fights/2]
	if want := profile.LogSample.TankSpike5; math.Abs(median-want) > 0.1 {
		t.Errorf("the tank's 5 s spike is %.2f, the logs' %.2f", median, want)
	}
}

// tankSpike5 is one fight's tank spikiness, the way scheduleRaidDamage and dealRaidDamage
// land the hits: the first at a random point inside the first swing, each rolled within the
// spread.
func tankSpike5(p RaidHealProfile, rng *rand.Rand) float64 {
	secs := p.DurationSec + 1
	bins := make([]float64, secs)
	for at := rng.Float64() * p.Tank.SwingSeconds; at < float64(p.DurationSec); at += p.Tank.SwingSeconds {
		bins[min(secs-1, int(at))] += p.Tank.HitDamage * (1 + p.DamageSpread*(2*rng.Float64()-1))
	}
	var sum, peak float64
	windows := secs - 4
	for i := range windows {
		w := 0.0
		for _, b := range bins[i : i+5] {
			w += b
		}
		sum += w
		peak = max(peak, w)
	}
	return peak / (sum / float64(windows))
}

func TestRaidHealProfileRefusesFiguresThatAreNotTheLogs(t *testing.T) {
	raw, err := os.ReadFile(curatedRaidHealProfile)
	if err != nil {
		t.Fatal(err)
	}
	for name, edit := range map[string][2]string{
		"tank hit":  {`"hit_damage": 360`, `"hit_damage": 720`},
		"pulse":     {`"damage": 176`, `"damage": 300`},
		"length":    {`"duration_sec": 110`, `"duration_sec": 300`},
		"no reason": {`"reason": "heal-profile.json's raid member: about 5,000 health in Phase 1 gear."`, `"reason": ""`},
	} {
		t.Run(name, func(t *testing.T) {
			if !strings.Contains(string(raw), edit[0]) {
				t.Fatalf("the curated file no longer carries %s", edit[0])
			}
			path := filepath.Join(t.TempDir(), RaidHealProfileFile)
			if err := os.WriteFile(path, []byte(strings.Replace(string(raw), edit[0], edit[1], 1)), 0o600); err != nil {
				t.Fatal(err)
			}
			if _, err := LoadRaidHealProfile(path); err == nil {
				t.Error("loaded")
			} else if !errors.Is(err, ErrBadHealProfile) {
				t.Errorf("err = %v, want ErrBadHealProfile", err)
			}
		})
	}
}

func TestRaidHealProfileAttachLaysOutTheFight(t *testing.T) {
	profile := mustRaidHealProfile(t)
	req := &proto.RaidSimRequest{
		Raid:      &proto.Raid{Parties: []*proto.Party{{Players: []*proto.Player{{}}}}},
		Encounter: &proto.Encounter{Duration: 300, DurationVariation: 36},
	}
	var fight HealFight = profile
	fight.Attach(req)
	model := req.Raid.RaidDamageModel
	if model == nil || model.Profile != profile.ID || model.TankHitDamage != profile.Tank.HitDamage || model.PulseDamage != profile.Pulse.Damage {
		t.Fatalf("model = %+v, want the profile's figures", model)
	}
	if req.Raid.TargetDummies != 5 || len(req.Raid.Parties) != 2 {
		t.Errorf("fake raid = %d members in %d parties, want 5 in 2", req.Raid.TargetDummies, len(req.Raid.Parties))
	}
	if req.Encounter.Duration != float64(profile.DurationSec) || req.Encounter.DurationVariation != 0 {
		t.Errorf("fight = %v +/- %v, want %d s exactly", req.Encounter.Duration, req.Encounter.DurationVariation, profile.DurationSec)
	}
}
