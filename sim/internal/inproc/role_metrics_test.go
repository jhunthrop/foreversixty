package inproc

import (
	"math"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/combine"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/internal/simdrain"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/jhunthrop/foreversixty/sim/score"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

// wasmPart is what the browser's simRun does for one worker's share of a
// role_metrics request: build, attach the item rows, run serially, read the
// summary and the role's blocks. The wasm itself only builds for js.
func wasmPart(t *testing.T, req api.SimRequest) api.SimResult {
	t.Helper()
	Register()
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: req.NoSample})
	if err != nil {
		t.Fatal(err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		t.Fatal(err)
	}
	reporter := make(chan *proto.ProgressMetrics, 32)
	core.RunRaidSimAsync(engineReq, reporter, nextRunID())
	res := simdrain.ToResult(reporter, nil)
	if err := adapter.ResultError(res); err != nil {
		t.Fatal(err)
	}
	sum, err := adapter.Summarize(res, req)
	if err != nil {
		t.Fatalf("the summary of a role fight: %v", err)
	}
	healing, tank, err := score.RoleBlocks(res, engineReq, req)
	if err != nil {
		t.Fatal(err)
	}
	return api.SimResult{
		EngineVersion: req.EngineVersion, Request: req, DPS: adapter.DPS(res),
		IterationsRun: int(res.IterationsDone), Summary: sum, Healing: healing, Tank: tank,
	}
}

func runShards(t *testing.T, req api.SimRequest, shards int) api.SimResult {
	t.Helper()
	parts, err := combine.Split(req, shards)
	if err != nil {
		t.Fatal(err)
	}
	results := make([]api.SimResult, len(parts))
	for i, part := range parts {
		results[i] = wasmPart(t, part)
	}
	out, err := combine.Results(results)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func near(a, b float64) bool { return math.Abs(a-b) <= 1e-6*math.Max(1, math.Abs(b)) }

func TestAHealersRoleMetricsPoolAcrossShardsAsOneRun(t *testing.T) {
	req := bandRequest(t, "priest-holy")
	req.Iterations, req.RoleMetrics = 24, true
	whole, pooled := runShards(t, req, 1), runShards(t, req, 3)
	if whole.Healing == nil || whole.Healing.EffectiveHPS.Mean <= 0 || whole.Tank != nil {
		t.Fatalf("a healer's run carries healing and no tank block: %+v", whole.Healing)
	}
	if got := whole.Healing.EffectiveHPS.Mean; got <= 0 || whole.Healing.HPM <= 0 || whole.Healing.ManaLastsSec <= 0 {
		t.Fatalf("implausible healing block: %+v", whole.Healing)
	}
	direct, err := HealingRun(req, mustProfile(t))
	if err != nil {
		t.Fatal(err)
	}
	if !near(whole.Healing.EffectiveHPS.Mean, direct.Effective.Mean) || !near(whole.Healing.HPM, direct.HealingPerMana()) {
		t.Errorf("browser path %+v differs from the ranker's %+v", whole.Healing, direct)
	}
	for name, pair := range map[string][2]float64{
		"effective hps": {pooled.Healing.EffectiveHPS.Mean, whole.Healing.EffectiveHPS.Mean},
		"hps":           {pooled.Healing.HPS.Mean, whole.Healing.HPS.Mean},
		"mana lasts":    {pooled.Healing.ManaLastsSec, whole.Healing.ManaLastsSec},
		"hpm":           {pooled.Healing.HPM, whole.Healing.HPM},
		"error":         {pooled.Healing.EffectiveHPS.Error, whole.Healing.EffectiveHPS.Error},
	} {
		if !near(pair[0], pair[1]) {
			t.Errorf("%s: three shards %v, one run %v", name, pair[0], pair[1])
		}
	}
}

func TestATanksRoleMetricsScoreAfterPoolingAsOneRun(t *testing.T) {
	req := bandRequest(t, "warrior-protection")
	req.Iterations, req.RoleMetrics = 24, true
	whole, pooled := runShards(t, req, 1), runShards(t, req, 3)
	if whole.Tank == nil || whole.Tank.Score.Mean <= 0 || whole.Healing != nil {
		t.Fatalf("a tank's run carries a tank block and no healing block: %+v", whole.Tank)
	}
	direct, err := TankRun(req)
	if err != nil {
		t.Fatal(err)
	}
	rawDPS, err := score.BossRawDPS(req.Character.Level)
	if err != nil {
		t.Fatal(err)
	}
	figures, err := score.NewTankFigures(direct, rawDPS)
	if err != nil {
		t.Fatal(err)
	}
	if !near(whole.Tank.Score.Mean, figures.Score) || !near(whole.Tank.Score.Error, figures.ScoreError) {
		t.Errorf("browser path score %+v differs from the ranker's %v +/- %v", whole.Tank.Score, figures.Score, figures.ScoreError)
	}
	for name, pair := range map[string][2]float64{
		"score":            {pooled.Tank.Score.Mean, whole.Tank.Score.Mean},
		"score error":      {pooled.Tank.Score.Error, whole.Tank.Score.Error},
		"effective health": {pooled.Tank.EffectiveHealth, whole.Tank.EffectiveHealth},
		"dtps":             {pooled.Tank.DTPS.Mean, whole.Tank.DTPS.Mean},
		"tmi":              {pooled.Tank.TMI.Mean, whole.Tank.TMI.Mean},
		"chance of death":  {pooled.Tank.ChanceOfDeath, whole.Tank.ChanceOfDeath},
	} {
		if !near(pair[0], pair[1]) {
			t.Errorf("%s: three shards %v, one run %v", name, pair[0], pair[1])
		}
	}
}

func TestRoleMetricsIsAbsentFromAPlainRun(t *testing.T) {
	req := bandRequest(t, "priest-holy")
	req.Iterations = 6
	if got := wasmPart(t, req); got.Healing != nil || got.Tank != nil {
		t.Errorf("a plain run carries role blocks: %+v %+v", got.Healing, got.Tank)
	}
}

func mustProfile(t *testing.T) request.HealProfile {
	t.Helper()
	profile, err := request.EmbeddedHealProfile()
	if err != nil {
		t.Fatal(err)
	}
	return profile
}
