// Command genfixture writes a checked-in RaidSimResult for sim/adapter's
// tests and for the api lane's, which need a real engine result and no
// engine binary.
//
// The player it runs is the one sim/request.Build produces, so a fixture
// is a real sim of a real request rather than a second opinion about
// what a fury warrior is. Only the gear comes from elsewhere: the site
// has no gear sets of its own yet, so the engine checkout's presets are
// read (and nothing there is written).
//
// It is a bootstrap: sim/cmd/forever-sim (Task 13) does the same job
// from a SimRequest JSON and is what regenerates these fixtures once the
// two Forever specs land. Task 13 deletes this package.
//
//	go run --tags=with_db ./internal/genfixture -spec warrior-fury -out adapter/testdata/warrior-fury.result.pb
package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/request"
	_ "github.com/wowsims/classic/sim/common"
	"github.com/wowsims/classic/sim/core"
	"github.com/wowsims/classic/sim/mage"
	dpswarrior "github.com/wowsims/classic/sim/warrior/dps_warrior"
	googleproto "google.golang.org/protobuf/proto"
)

// engineDir is where the engine's own gear-set JSON lives.
// core.GetGearSet reads it from disk relative to the process's working
// directory, and those files ship with the engine repository rather
// than with its module, so the path is a flag rather than a constant.
var engineDir = flag.String("engine-dir", "/Users/jh/code/wowsims-forever", "the engine checkout, for its gear_sets JSON (read only)")

// fixture is one checked-in result: the character sim/request builds it
// from, and the engine preset its gear comes from.
type fixture struct {
	character api.CharacterSpec
	gearSet   string
}

var fixtures = map[string]fixture{
	"warrior-fury": {
		character: api.CharacterSpec{
			Name: "Fury", Race: "orc", Class: "warrior", Level: 60,
			Talents: "30305001302-05050005525010051",
		},
		gearSet: "ui/warrior/gear_sets/phase_1",
	},
	"mage-frost": {
		character: api.CharacterSpec{
			Name: "Frost", Race: "gnome", Class: "mage", Level: 60,
			// The engine's own frost preset, ui/mage/presets.ts's
			// TalentsP1DPS. The string this task's brief carried is not
			// positional against this fork's mage trees and the engine
			// panics on it.
			Talents: "230205021002--05353203102351001",
		},
		// The mage's phase-one gear set is named p1.bis, not phase_1;
		// the fixture's job is to be a real engine result, not a
		// particular gear set.
		gearSet: "ui/mage/gear_sets/p1.bis",
	},
}

func main() {
	slug := flag.String("spec", "warrior-fury", "spec slug")
	out := flag.String("out", "out.result.pb", "output file")
	iters := flag.Int("iterations", 3000, "iterations; one of api.ValidIterations")
	flag.Parse()

	dpswarrior.RegisterDpsWarrior()
	mage.RegisterMage()

	fx, ok := fixtures[*slug]
	if !ok {
		log.Fatalf("unknown spec %q", *slug)
	}

	req, err := request.Build(api.SimRequest{
		EngineVersion: "genfixture",
		Spec:          *slug,
		Character:     fx.character,
		Encounter:     api.DefaultEncounter(),
		Iterations:    *iters,
		RandomSeed:    1,
	})
	if err != nil {
		log.Fatal(err)
	}
	player := req.Raid.Parties[0].Players[0]
	player.Equipment = core.GetGearSet(split(*engineDir + "/" + fx.gearSet)).GearSet
	// A fixture is a buffed fight, because that is the one whose damage
	// table has something in every column. The envelope's buff ids
	// cannot say "everything", so the engine's own full-buff sets stand
	// in here; sim/request maps the settings bar's ids in production.
	player.Buffs = core.FullIndividualBuffs
	req.Raid.Parties[0].Buffs = core.FullPartyBuffs
	req.Raid.Buffs = core.FullRaidBuffs
	req.Raid.Debuffs = core.FullDebuffs

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
// engine's loader takes as two arguments.
func split(path string) (dir, file string) {
	i := strings.LastIndex(path, "/")
	return path[:i], path[i+1:]
}
