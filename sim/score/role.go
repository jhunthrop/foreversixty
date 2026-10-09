package score

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core/proto"
)

// RoleBlocks reads the figures a role_metrics request asked for from a
// finished engine run: a healer's Healing block or a tank's Tank block.
// engineReq is the request the run was built from (a healer's fixed fight
// length and a tank's maximum health come from it). A request that did not
// set role_metrics gets neither block.
func RoleBlocks(res *proto.RaidSimResult, engineReq *proto.RaidSimRequest, req api.SimRequest) (*api.HealingResult, *api.TankResult, error) {
	if !req.RoleMetrics {
		return nil, nil, nil
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		return nil, nil, fmt.Errorf("reading the player's metrics: %w", err)
	}
	switch {
	case request.IsHealerSpec(req.Spec):
		return HealingResultOf(player, res.IterationsDone, engineReq.GetEncounter().GetDuration()).Block(), nil, nil
	case request.IsTankSpec(req.Spec):
		health, err := MaxHealth(engineReq.GetRaid(), engineReq.GetEncounter())
		if err != nil {
			return nil, nil, err
		}
		block, err := TankBlockOf(TankRunOf(player, res.IterationsDone, health), req.Character.Level)
		return nil, block, err
	}
	return nil, nil, fmt.Errorf("%w: %q", request.ErrRoleMetricsSpec, req.Spec)
}
