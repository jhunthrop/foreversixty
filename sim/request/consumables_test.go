package request

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/wowsims/classic/sim/core/proto"
)

// testConsumables loads the ten-row excerpt of the real
// data/builds/<build>/simconsumes.json, copied from the data lane's
// 1.60.1.69893 build. The full file is 1,579 rows and belongs to a
// build, so it is loaded at run time rather than committed here.
func testConsumables(t *testing.T) *Consumables {
	t.Helper()
	f, err := os.Open(filepath.Join("testdata", "simconsumes.excerpt.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	c, err := LoadConsumables(f)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestLoadConsumables(t *testing.T) {
	c := testConsumables(t)
	if c.Len() != 10 {
		t.Errorf("Len = %d, want the excerpt's 10 rows", c.Len())
	}
	if _, err := LoadConsumables(strings.NewReader("[]")); err == nil {
		t.Error("an empty consumables file loaded without error")
	}
	if _, err := LoadConsumables(strings.NewReader("not json")); err == nil {
		t.Error("a malformed consumables file loaded without error")
	}
	if _, err := LoadConsumables(strings.NewReader(`[{"id":0,"name":""}]`)); err == nil {
		t.Error("a row with no id loaded without error")
	}
}

// An item id names the same engine value its name does. The join is by
// name, so nothing hand-written pairs the client's items with the
// engine's enums.
func TestItemIDsResolveThroughTheTable(t *testing.T) {
	table := testConsumables(t)
	got, err := consumes([]string{"item:13452", "item:13510", "item:13928"}, table)
	if err != nil {
		t.Fatal(err)
	}
	if got.AgilityElixir != proto.AgilityElixir_ElixirOfTheMongoose {
		t.Errorf("AgilityElixir = %v, want ElixirOfTheMongoose from item 13452", got.AgilityElixir)
	}
	if got.Flask != proto.Flask_FlaskOfTheTitans {
		t.Errorf("Flask = %v, want FlaskOfTheTitans from item 13510", got.Flask)
	}
	// Food's values repeat the enum's own name, and the item does not,
	// so the id has to answer to both spellings.
	if got.Food != proto.Food_FoodGrilledSquid {
		t.Errorf("Food = %v, want FoodGrilledSquid from item 13928", got.Food)
	}
}

func TestItemIDsNeedATable(t *testing.T) {
	_, err := consumes([]string{"item:13452"}, nil)
	if !errors.Is(err, ErrUnknownConsume) {
		t.Fatalf("an item id with no table returned %v, want ErrUnknownConsume", err)
	}
	if !strings.Contains(err.Error(), "simconsumes.json") {
		t.Errorf("the error does not say where the table comes from: %v", err)
	}
}

func TestUnknownAndUnmodelledItems(t *testing.T) {
	table := testConsumables(t)
	// An item the build does not carry.
	if _, err := consumes([]string{"item:999999"}, table); !errors.Is(err, ErrUnknownConsume) {
		t.Errorf("an item the build has no row for returned %v", err)
	}
	// An item the build carries and the engine does not model: most of
	// the file is food no sim has a value for.
	if _, err := consumes([]string{"item:117"}, table); !errors.Is(err, ErrUnknownConsume) {
		t.Errorf("Tough Jerky, which the engine has no enum value for, returned %v", err)
	}
	// Not a number.
	if _, err := consumes([]string{"item:frog"}, table); !errors.Is(err, ErrUnknownConsume) {
		t.Errorf("a non-numeric item id returned %v", err)
	}
}

// An imbue item is as ambiguous as an imbue name: both hands hold the
// same enum, so the request has to say which.
func TestItemIDsCanBeSlotQualified(t *testing.T) {
	table := testConsumables(t)
	if _, err := consumes([]string{"item:3824"}, table); !errors.Is(err, ErrAmbiguousConsume) {
		t.Fatalf("a bare imbue item returned %v, want ErrAmbiguousConsume", err)
	}
	got, err := consumes([]string{"off_hand_imbue:item:3824"}, table)
	if err != nil {
		t.Fatal(err)
	}
	if got.OffHandImbue != proto.WeaponImbue_ShadowOil {
		t.Errorf("OffHandImbue = %v, want ShadowOil from item 3824", got.OffHandImbue)
	}
	if got.MainHandImbue != proto.WeaponImbue_WeaponImbueUnknown {
		t.Errorf("MainHandImbue = %v, want it untouched", got.MainHandImbue)
	}
}

// The table's own vocabulary lists only the items the engine can act
// on, so a settings bar built from it cannot offer an unsimulable one.
func TestConsumableTableIDs(t *testing.T) {
	ids := testConsumables(t).IDs()
	if len(ids) == 0 {
		t.Fatal("the table lists no ids")
	}
	for i := 1; i < len(ids); i++ {
		if ids[i-1] >= ids[i] {
			t.Fatalf("IDs() is not sorted: %q before %q", ids[i-1], ids[i])
		}
	}
	for _, id := range ids {
		if _, err := consumes([]string{id}, testConsumables(t)); err != nil && !errors.Is(err, ErrAmbiguousConsume) {
			t.Errorf("%s is listed but does not resolve: %v", id, err)
		}
	}
	if slices.Contains(ids, "item:117") {
		t.Error("Tough Jerky is listed, but the engine has no value for it")
	}
}

// BuildWith is how the api lane passes the table in.
func TestBuildWithResolvesItemConsumables(t *testing.T) {
	req := fury()
	req.Character.Consumes = []string{"item:13452"}
	if _, err := Build(req); !errors.Is(err, ErrUnknownConsume) {
		t.Fatalf("Build without a table accepted an item id: %v", err)
	}
	got, err := BuildWith(req, Options{Consumables: testConsumables(t)})
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid.Parties[0].Players[0].Consumes.AgilityElixir != proto.AgilityElixir_ElixirOfTheMongoose {
		t.Error("the item id did not reach the player's consumes")
	}
}
