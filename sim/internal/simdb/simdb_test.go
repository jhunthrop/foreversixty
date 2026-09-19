package simdb

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/request"
	"github.com/wowsims/classic/sim/core/proto"
)

func TestDatabaseLoads(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	if len(db.Items) == 0 || len(db.Enchants) == 0 {
		t.Fatalf("items=%d enchants=%d; the active build's simdb.bin is empty or stale",
			len(db.Items), len(db.Enchants))
	}
	t.Logf("items=%d enchants=%d random_suffixes=%d", len(db.Items), len(db.Enchants), len(db.RandomSuffixes))
}

// Every item and enchant the checked-in fixture requests name must exist
// in the build the artifacts embed. Forever re-itemises, so a gear id
// copied from a vanilla preset resolves to nothing and the sim dies with
// "No item with id" - at the fixture refresh if we are lucky and in a
// visitor's browser if we are not.
func TestFixtureGearResolves(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Fatal(err)
	}
	items := map[int32]string{}
	for _, it := range db.Items {
		items[it.Id] = it.Name
	}
	enchants := map[int32]bool{}
	for _, e := range db.Enchants {
		enchants[e.EffectId] = true
	}

	for _, spec := range []string{"warrior-fury", "mage-frost"} {
		t.Run(spec, func(t *testing.T) {
			b, err := os.ReadFile(filepath.Join("..", "..", "adapter", "testdata", spec+".request.json"))
			if err != nil {
				t.Fatal(err)
			}
			// Strictly: the fixture requests are plain SimRequests and
			// carry no key the envelope does not define. The envelope
			// crosses lane boundaries - the web mirrors it tag for tag
			// and the api lane decodes it - so provenance that means
			// nothing to the product lives in the .request.notes.md
			// beside each file, not in a field here.
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.DisallowUnknownFields()
			var req api.SimRequest
			if err := dec.Decode(&req); err != nil {
				t.Fatalf("%v; if this is an unknown field, it belongs in %s.request.notes.md", err, spec)
			}
			// The fixture carries no engine_version: it is a request
			// about a character and an encounter, and the engine that
			// runs it is whichever binary is asked to. forever-sim
			// fills the field in the same way.
			if req.EngineVersion != "" {
				t.Errorf("engine_version = %q; the fixture must not name an engine, or it expires on every pin", req.EngineVersion)
			}
			req.EngineVersion = enginever.Version
			if err := req.Validate(); err != nil {
				t.Errorf("the fixture request is not a valid SimRequest: %v", err)
			}
			notes := filepath.Join("..", "..", "adapter", "testdata", spec+".request.notes.md")
			if _, err := os.Stat(notes); err != nil {
				t.Errorf("%s is missing; every re-pointed slot needs its reason written down", notes)
			}
			if len(req.Character.Gear) == 0 {
				t.Fatal("the fixture request equips nothing; it is meant to be a real geared fight")
			}
			for _, g := range req.Character.Gear {
				if _, ok := items[int32(g.ItemID)]; !ok {
					t.Errorf("%s: item %d has no row in the embedded database", g.Slot, g.ItemID)
				}
				if g.Enchant != 0 && !enchants[int32(g.Enchant)] {
					t.Errorf("%s: enchant %d has no row in the embedded database", g.Slot, g.Enchant)
				}
			}
		})
	}
}

// Attach is the only way the database reaches the engine, so a request
// that went through it must carry it on the player.
func TestAttachPutsTheDatabaseOnEveryPlayer(t *testing.T) {
	built, err := request.Build(api.SimRequest{
		EngineVersion: enginever.Version, Spec: "warrior-fury",
		Source:     api.CharacterSource{Kind: api.SourceManual},
		Character:  api.CharacterSpec{Name: "T", Race: "orc", Class: "warrior", Level: 60},
		Encounter:  api.DefaultEncounter(),
		Iterations: 500,
	})
	if err != nil {
		t.Fatal(err)
	}
	if built.Raid.Parties[0].Players[0].Database != nil {
		t.Fatal("request.Build attached a database of its own; simdb is meant to be the only source")
	}
	if err := Attach(built); err != nil {
		t.Fatal(err)
	}
	got := built.Raid.Parties[0].Players[0].Database
	if got == nil || len(got.Items) == 0 {
		t.Fatal("Attach left the player without a database")
	}

	// A nil request is not an error: the callers check the build error
	// first, and Attach must not be a second place that can panic.
	if err := Attach(nil); err != nil {
		t.Errorf("Attach(nil) = %v, want nil", err)
	}
	if err := Attach(&proto.RaidSimRequest{}); err != nil {
		t.Errorf("Attach(empty) = %v, want nil", err)
	}
}

// AttachWeights is the sibling of Attach for a stat weights request: the
// database reaches the engine through the same proto.Player.Database
// door, and a weights run equips the character the same way a DPS run
// does.
func TestAttachWeightsPutsTheDatabaseOnThePlayer(t *testing.T) {
	built, err := request.BuildWeights(api.SimRequest{
		EngineVersion: enginever.Version, Spec: "warrior-fury",
		Source:    api.CharacterSource{Kind: api.SourceManual},
		Character: api.CharacterSpec{Name: "T", Race: "orc", Class: "warrior", Level: 60},
		Encounter: api.DefaultEncounter(),
		Weights: &api.WeightsSpec{
			Stats:     []string{"agility", "attack_power"},
			Reference: "attack_power",
		},
		Iterations: 500,
	}, request.Options{})
	if err != nil {
		t.Fatal(err)
	}
	if built.Player.Database != nil {
		t.Fatal("request.BuildWeights attached a database of its own; simdb is meant to be the only source")
	}
	if err := AttachWeights(built); err != nil {
		t.Fatal(err)
	}
	got := built.Player.Database
	if got == nil || len(got.Items) == 0 {
		t.Fatal("AttachWeights left the player without a database")
	}

	// A nil request, and a request with no player, are not errors: the
	// callers check the build error first, and AttachWeights must not
	// be a second place that can panic.
	if err := AttachWeights(nil); err != nil {
		t.Errorf("AttachWeights(nil) = %v, want nil", err)
	}
	if err := AttachWeights(&proto.StatWeightsRequest{}); err != nil {
		t.Errorf("AttachWeights(empty) = %v, want nil", err)
	}
}
