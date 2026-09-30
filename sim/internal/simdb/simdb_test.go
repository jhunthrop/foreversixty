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

// A worn item the database lacks -- the quality-1 Ancient Heirloom ring
// (264908), which items/<class>.json leaves out on purpose -- used to reach
// the engine and panic it ("No item with id"). Attach empties that slot and
// leaves the known ones alone.
func TestAttachUnequipsAnItemTheDatabaseLacks(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Skip("embedded database unavailable:", err)
	}
	known := db.Items[0].Id
	player := &proto.Player{Equipment: &proto.EquipmentSpec{Items: []*proto.ItemSpec{
		{Id: known}, {Id: 264908}, {}, nil,
	}}}
	req := &proto.RaidSimRequest{Raid: &proto.Raid{Parties: []*proto.Party{{Players: []*proto.Player{player}}}}}
	if err := Attach(req); err != nil {
		t.Fatal(err)
	}
	if player.Equipment.Items[0].Id != known {
		t.Errorf("the known item %d was removed", known)
	}
	if player.Equipment.Items[1].Id != 0 {
		t.Errorf("the unknown item 264908 stayed equipped: %v", player.Equipment.Items[1])
	}
	removed, err := UnequipUnknown(&proto.Player{Equipment: &proto.EquipmentSpec{Items: []*proto.ItemSpec{{Id: 264908}, {Id: known}}}})
	if err != nil || len(removed) != 1 || removed[0] != 264908 {
		t.Errorf("UnequipUnknown = %v, %v; want [264908]", removed, err)
	}
	if removed, err := UnequipUnknown(nil); err != nil || removed != nil {
		t.Errorf("nil player: %v, %v", removed, err)
	}
}

// TestKnownAgreesWithUnequipUnknown locks Known's own contract to
// UnequipUnknown's real behaviour: an id Known reports true for must
// survive Attach on a character wearing it, and one it reports false
// for must be exactly the id UnequipUnknown strips. A ranker that
// asks Known before trusting a real sim's measurement of an item's
// effect (rank.go's effectVerifiedInSim) is only honest if Known and
// UnequipUnknown can never disagree.
func TestKnownAgreesWithUnequipUnknown(t *testing.T) {
	db, err := load()
	if err != nil {
		t.Skip("embedded database unavailable:", err)
	}
	known := db.Items[0].Id
	if !Known(known) {
		t.Errorf("Known(%d) = false, want true (this id has its own simdb row)", known)
	}

	const unknownID int32 = 264908
	if Known(unknownID) {
		t.Fatalf("Known(%d) = true, want false (test fixture assumes this id has no simdb row)", unknownID)
	}
	removed, err := UnequipUnknown(&proto.Player{Equipment: &proto.EquipmentSpec{Items: []*proto.ItemSpec{{Id: unknownID}}}})
	if err != nil || len(removed) != 1 || removed[0] != unknownID {
		t.Fatalf("UnequipUnknown = %v, %v; want [%d], confirming Known(%d) = false was right", removed, err, unknownID, unknownID)
	}
}

// TestKnownHandOfJusticeRegression documents the defect this lane
// found and fixed (bis-ranker-integrity-8's own follow-on, effect-procs
// lane, 2026-09-30): Hand of Justice (item 11815, sim/common's own
// HandOfJustice constant in the wowsims-forever fork, which DOES
// register a real on-hit proc for it) has no row in this build's
// embedded simdb.bin, because its ItemSparse row is absent from this
// build's client export (data/builds/<build>/raw/ItemSparse.csv) even
// though its Item.csv row and items.json's own classic-db-fallback
// entry both exist - data/pipeline/simdb/items.py's simdb_item_rows
// only keeps a row present in BOTH client tables. Before this lane,
// nothing asked Known before trusting a real sim's measured DPS with
// Hand of Justice equipped, so simdb.Attach's own UnequipUnknown
// silently stripped it from every tournament character and the
// resulting measurement (bit-identical to no trinket at all) read as
// "the engine modelled this effect and it is worth exactly zero" -
// this lane's brief calls that a tenet-8 violation.
//
// This test intentionally names the real item id rather than a
// synthetic one: it is the regression itself, not just a property of
// Known in the abstract, and a future data-pipeline fix that adds
// Hand of Justice's row back to simdb.bin should make this test fail
// loudly (a reminder to update or remove it) rather than pass by
// accident.
func TestKnownHandOfJusticeRegression(t *testing.T) {
	if _, err := load(); err != nil {
		t.Skip("embedded database unavailable:", err)
	}
	// simdb-supplement (2026-09-30) builds the database from the whole
	// catalogue, so Hand of Justice (a classic-db row) must stay Known;
	// its disappearance would silently strip it from every tournament
	// again (the effect-procs finding this test was born from).
	const handOfJustice int32 = 11815
	if !Known(handOfJustice) {
		t.Fatalf("Known(%d) = false; Hand of Justice lost its simdb row - the classic-db supplement (pipeline/simdb) is not feeding the database", handOfJustice)
	}
}
