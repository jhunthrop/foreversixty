package request

import (
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/wowsims/classic/sim/core/proto"
	gproto "google.golang.org/protobuf/proto"
)

// Round 11 (regression lane C9): below 60 the curated raid preset, resolved engine-side, is
// not cut to the character's level (a band-20 resolution carries Elemental Sharpening Stone
// and Brilliant Wizard Oil). No server path resolves it: the api never imports this package
// (api.yml), and a request's `preset` only picks spell ranks. This holds that line: a
// level-20 request naming the raid preset gets no buff and no consumable it did not list.
func TestARequestPresetAddsNoBuffOrConsumable(t *testing.T) {
	for _, preset := range []string{RaidPreset, BarePreset} {
		req := api.SimRequest{
			EngineVersion: enginever.Version, Spec: "warrior-fury",
			Source:    api.CharacterSource{Kind: api.SourceManual},
			Character: api.CharacterSpec{Name: "x", Race: "orc", Class: "warrior", Level: 20},
			Encounter: api.DefaultEncounter(), Iterations: 3000, RandomSeed: 1,
			Preset: preset,
		}
		built, err := Build(req)
		if err != nil {
			t.Fatalf("%s: %v", preset, err)
		}
		player := built.GetRaid().GetParties()[0].GetPlayers()[0]
		if c := player.GetConsumes(); c != nil && !gproto.Equal(c, &proto.Consumes{}) {
			t.Errorf("%s: the request listed no consumables, the build carries %v", preset, c)
		}
		if b := built.GetRaid().GetBuffs(); b != nil && !gproto.Equal(b, &proto.RaidBuffs{}) {
			t.Errorf("%s: the request listed no buffs, the build carries raid buffs %v", preset, b)
		}
		if b := player.GetBuffs(); b != nil && !gproto.Equal(b, &proto.IndividualBuffs{}) {
			t.Errorf("%s: the request listed no buffs, the build carries %v", preset, b)
		}
		if b := built.GetRaid().GetParties()[0].GetBuffs(); b != nil && !gproto.Equal(b, &proto.PartyBuffs{}) {
			t.Errorf("%s: the request listed no buffs, the build carries party buffs %v", preset, b)
		}
		if d := built.GetRaid().GetDebuffs(); d != nil && !gproto.Equal(d, &proto.Debuffs{}) {
			t.Errorf("%s: the request listed no debuffs, the build carries %v", preset, d)
		}
	}
}
