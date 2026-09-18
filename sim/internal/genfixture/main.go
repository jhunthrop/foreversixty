// Command genfixture writes a checked-in RaidSimResult for sim/adapter's
// tests and for the api lane's, which need a real engine result and no
// engine binary.
//
// It is a bootstrap: sim/cmd/forever-sim (Task 13) does the same job
// from a SimRequest JSON and is what regenerates these fixtures once the
// two Forever specs land. Task 13 deletes this package.
//
// It runs in the site module and writes nothing into the engine
// checkout, which it only reads gear and rotation JSON from:
//
//	go run --tags=with_db ./internal/genfixture -spec warrior-fury -out adapter/testdata/warrior-fury.result.pb
package main

import (
	"flag"
	"log"
	"os"
	"strings"

	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/core/proto"
	"github.com/wowsims/classic/sim/mage"
	dpswarrior "github.com/wowsims/classic/sim/warrior/dps_warrior"
	googleproto "google.golang.org/protobuf/proto"
)

// engineDir is where the engine's own gear-set and APL JSON live.
// core.GetGearSet and core.GetAplRotation read them from disk relative
// to the process's working directory, and those files ship with the
// engine repository rather than with its module, so the path is a flag
// rather than a constant.
var engineDir = flag.String("engine-dir", "/Users/jh/code/wowsims-forever", "the engine checkout, for its gear_sets and apls JSON (read only)")

// fixtureDuration is the fight the fixtures are run at, in seconds. It is
// api.DefaultEncounter's duration, so a golden reads like the sim the
// settings bar opens on.
const fixtureDuration = 180

// spec is one fixture: which engine presets it is built from and what
// player it runs.
type spec struct {
	gearSet string
	apl     string
	player  func() *proto.Player
	options func(*proto.Player)
}

var specs = map[string]spec{
	"warrior-fury": {
		gearSet: "ui/warrior/gear_sets/phase_1",
		apl:     "ui/warrior/apls/dps_reck",
		player: func() *proto.Player {
			return &proto.Player{
				Name: "Fury", Race: proto.Race_RaceOrc, Class: proto.Class_ClassWarrior,
				TalentsString: "30305001302-05050005525010051", Consumes: &proto.Consumes{},
				Buffs: core.FullIndividualBuffs,
			}
		},
		options: func(p *proto.Player) {
			core.WithSpec(p, &proto.Player_Warrior{Warrior: &proto.Warrior{
				Options: &proto.Warrior_Options{StartingRage: 50, Shout: proto.WarriorShout_WarriorShoutBattle},
			}})
		},
	},
	"mage-frost": {
		// The mage's phase-one gear set is named p1.bis, not phase_1;
		// the fixture's job is to be a real engine result, not a
		// particular gear set.
		gearSet: "ui/mage/gear_sets/p1.bis",
		apl:     "ui/mage/apls/p1",
		player: func() *proto.Player {
			return &proto.Player{
				Name: "Frost", Race: proto.Race_RaceGnome, Class: proto.Class_ClassMage,
				// The engine's own frost preset, ui/mage/presets.ts's
				// TalentsP1DPS; the string this task's brief carried is
				// not positional against this fork's mage trees and the
				// engine panics on it.
				TalentsString: "230205021002--05353203102351001", Consumes: &proto.Consumes{},
				Buffs: core.FullIndividualBuffs,
			}
		},
		options: func(p *proto.Player) {
			core.WithSpec(p, &proto.Player_Mage{Mage: &proto.Mage{Options: &proto.Mage_Options{}}})
		},
	},
}

func main() {
	slug := flag.String("spec", "warrior-fury", "spec slug")
	out := flag.String("out", "out.result.pb", "output file")
	iters := flag.Int("iterations", 3000, "iterations")
	flag.Parse()

	dpswarrior.RegisterDpsWarrior()
	mage.RegisterMage()

	sp, ok := specs[*slug]
	if !ok {
		log.Fatalf("unknown spec %q", *slug)
	}

	player := sp.player()
	player.Equipment = core.GetGearSet(split(*engineDir + "/" + sp.gearSet)).GearSet
	player.Rotation = core.GetAplRotation(split(*engineDir + "/" + sp.apl)).Rotation
	sp.options(player)

	enc := core.MakeSingleTargetEncounter(0.2)
	enc.Duration = fixtureDuration
	req := &proto.RaidSimRequest{
		Raid:       core.SinglePlayerRaidProto(player, core.FullPartyBuffs, core.FullRaidBuffs, core.FullDebuffs),
		Encounter:  enc,
		SimOptions: &proto.SimOptions{Iterations: int32(*iters), IsTest: false, RandomSeed: 1},
	}
	res := core.RunRaidSim(req)
	if res.Error != nil {
		log.Fatalf("sim failed: %s", res.Error.Message)
	}
	// The cast log is not part of the fixture: it is one iteration of text
	// and it makes the file large without changing anything the adapter reads.
	res.Logs = ""
	b, err := googleproto.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(*out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s (%d bytes), dps=%.1f", *out, len(b), res.RaidMetrics.Dps.Avg)
}

// split turns a preset path into the directory and base name the
// engine's loaders take as two arguments.
func split(path string) (dir, file string) {
	i := strings.LastIndex(path, "/")
	return path[:i], path[i+1:]
}
