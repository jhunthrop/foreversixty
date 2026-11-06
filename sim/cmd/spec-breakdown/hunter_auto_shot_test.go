package main

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/adapter"
	"github.com/jhunthrop/foreversixty/sim/internal/inproc"
	"github.com/jhunthrop/foreversixty/sim/internal/simdb"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
)

const arcaneShotsAt50Floor = 15.0

// arcaneShotRank1 is Arcane Shot rank 1; the rank rows pick the rank a
// level casts, so the count takes every rank.
const arcaneShotRank1 = 3044

// autoShots is the hunter's Auto Shot casts per iteration under rotation.
func autoShots(t *testing.T, rotation *proto.APLRotation) float64 {
	t.Helper()
	inproc.Register()
	o := options{repoRoot: "../../..", spec: "hunter-marksmanship", preset: "raid", band: 60, faction: "horde", iterations: 200, seed: 1}
	req, err := legalRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		t.Fatal(err)
	}
	for _, party := range engineReq.GetRaid().GetParties() {
		for _, player := range party.GetPlayers() {
			player.Rotation = rotation
			player.DistanceFromTarget = 25
		}
	}
	if err := simdb.Attach(engineReq); err != nil {
		t.Fatal(err)
	}
	res := core.RunRaidSim(engineReq)
	if err := adapter.ResultError(res); err != nil {
		t.Fatal(err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		t.Fatal(err)
	}
	var casts float64
	for _, action := range player.GetActions() {
		if action.GetId().GetOtherId() != proto.OtherAction_OtherActionShoot {
			continue
		}
		for _, target := range action.GetTargets() {
			casts += float64(target.GetCasts())
		}
	}
	return casts / float64(o.iterations)
}

// In Forever an Auto Shot fires while Aimed Shot is being cast: in the
// beta logs (foreverlogs 2695-2764, nine hunters, 2-10 October 2026) about
// 60% of Aimed Shot casts have an Auto Shot logged inside the cast, about
// what a timer that ignores the cast gives (a 1.8 s cast against a 2.2 s
// swing). The engine cancelled the auto timer for the cast and fired the
// missed shot only when the cast ended, so every Aimed Shot cost a share
// of an Auto Shot (sim-accuracy round 21, guide-accuracy round 21 lead).
// A hunter casting Aimed Shot on every cooldown shoots as many Auto Shots
// as one who never casts it.
func TestAutoShotKeepsFiringThroughAimedShot(t *testing.T) {
	aimed := &proto.APLListItem{Action: &proto.APLAction{Action: &proto.APLAction_CastSpell{CastSpell: &proto.APLActionCastSpell{
		SpellId: &proto.ActionID{RawId: &proto.ActionID_SpellId{SpellId: 20904}},
	}}}}
	idle := autoShots(t, &proto.APLRotation{Type: proto.APLRotation_TypeAPL})
	casting := autoShots(t, &proto.APLRotation{Type: proto.APLRotation_TypeAPL, PriorityList: []*proto.APLListItem{aimed}})
	t.Logf("Auto Shots per fight: %.1f with no casts, %.1f with Aimed Shot on cooldown", idle, casting)
	if casting < 0.97*idle {
		t.Errorf("Auto Shots with Aimed Shot on cooldown = %.1f, want within 3%% of %.1f with no casts", casting, idle)
	}
}

// Marksmanship at 50 and 51 casts the top Aimed Shot rank, which costs more
// than a hunter under a quarter of his mana often has. The shared cooldown
// then sat ready with nothing to spend it, and Serpent Sting and Arcane
// Shot, which wait on it, stood idle too (sim-accuracy round 21, rendered
// card review). Under a quarter mana they keep firing at every level.
func TestMarksmanshipAt50KeepsShootingWhenAimedShotIsOutOfReach(t *testing.T) {
	inproc.Register()
	o := options{repoRoot: "../../..", spec: "hunter-marksmanship", preset: "bare", band: 50, faction: "horde", iterations: 200, seed: 1}
	req, err := legalRequest(o)
	if err != nil {
		t.Fatal(err)
	}
	engineReq, err := request.BuildWith(req, request.Options{OpenIterations: true, NoSampleIteration: true})
	if err != nil {
		t.Fatal(err)
	}
	if err := simdb.Attach(engineReq); err != nil {
		t.Fatal(err)
	}
	res := core.RunRaidSim(engineReq)
	if err := adapter.ResultError(res); err != nil {
		t.Fatal(err)
	}
	player, err := adapter.PlayerMetrics(res)
	if err != nil {
		t.Fatal(err)
	}
	var arcane float64
	for _, action := range player.GetActions() {
		switch action.GetId().GetSpellId() {
		case arcaneShotRank1, 14281, 14282, 14283, 14284, 14285, 14286, 14287:
			for _, target := range action.GetTargets() {
				arcane += float64(target.GetCasts())
			}
		}
	}
	arcane /= float64(o.iterations)
	t.Logf("Arcane Shot casts at 50: %.1f", arcane)
	if arcane < arcaneShotsAt50Floor {
		t.Errorf("Arcane Shot cast %.1f times at 50, want %v or more", arcane, arcaneShotsAt50Floor)
	}
}
