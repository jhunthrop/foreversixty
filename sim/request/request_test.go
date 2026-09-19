package request

import (
	"bytes"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

func fury() api.SimRequest {
	return api.SimRequest{
		EngineVersion: "7779ebb",
		Spec:          "warrior-fury",
		Character: api.CharacterSpec{
			Name:    "Thrall",
			Race:    "orc",
			Class:   "warrior",
			Level:   60,
			Talents: "30305001302-05050005525010051",
			Gear: []api.GearSlot{
				{Slot: "main_hand", ItemID: 19352, Enchant: 2568},
				{Slot: "head", ItemID: 16963},
			},
			Buffs:    []string{"battle_shout"},
			Consumes: []string{"elixir_of_the_mongoose"},
		},
		Encounter:  api.DefaultEncounter(),
		Iterations: 3000,
		RandomSeed: 7,
	}
}

func TestBuildProducesAPlayableRequest(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if got.Raid == nil || len(got.Raid.Parties) != 1 || len(got.Raid.Parties[0].Players) != 1 {
		t.Fatalf("Build did not produce one party with one player: %+v", got.Raid)
	}
	p := got.Raid.Parties[0].Players[0]
	if p.Name != "Thrall" {
		t.Errorf("Name = %q", p.Name)
	}
	if p.Race != proto.Race_RaceOrc {
		t.Errorf("Race = %v, want RaceOrc", p.Race)
	}
	if p.Class != proto.Class_ClassWarrior {
		t.Errorf("Class = %v, want ClassWarrior", p.Class)
	}
	// The brief's draft asserted p.Level here. The engine carries no
	// per-player level: sim/core builds every character at
	// core.CharacterMaxLevel, and proto.Player has no level field, so
	// there is nothing to assert on the request. api.SimRequest.Validate
	// refuses any other level - see
	// TestBuildRejectsALevelTheEngineCannotSimulate in mapping_test.go.
	if p.TalentsString != "30305001302-05050005525010051" {
		t.Errorf("TalentsString = %q", p.TalentsString)
	}
	if p.Spec == nil {
		t.Error("the player has no spec options; the engine cannot build an agent without one")
	}
}

// The engine's EquipmentSpec is positional: slot order is the contract,
// not a name, so a mis-ordered gear list silently equips a helm in the
// weapon slot.
func TestBuildPlacesGearByItsSlot(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	eq := got.Raid.Parties[0].Players[0].Equipment
	if eq == nil {
		t.Fatal("no equipment")
	}
	headIdx, ok := SlotIndex("head")
	if !ok {
		t.Fatal("head is not a known slot")
	}
	mhIdx, ok := SlotIndex("main_hand")
	if !ok {
		t.Fatal("main_hand is not a known slot")
	}
	if eq.Items[headIdx].Id != 16963 {
		t.Errorf("head slot holds item %d, want 16963", eq.Items[headIdx].Id)
	}
	if eq.Items[mhIdx].Id != 19352 {
		t.Errorf("main hand holds item %d, want 19352", eq.Items[mhIdx].Id)
	}
	if eq.Items[mhIdx].Enchant != 2568 {
		t.Errorf("main hand enchant = %d, want 2568", eq.Items[mhIdx].Enchant)
	}
	// Every slot the character does not fill must still exist, empty,
	// because the engine indexes the array rather than searching it.
	if len(eq.Items) != SlotCount {
		t.Errorf("EquipmentSpec has %d slots, want %d", len(eq.Items), SlotCount)
	}
}

func TestBuildCarriesTheEncounter(t *testing.T) {
	req := fury()
	req.Encounter = api.EncounterSpec{DurationSec: 300, Variation: 0.1, Targets: 3, ExecuteRatio: 0.2}
	got, err := Build(req)
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if e.Duration != 300 {
		t.Errorf("Duration = %v, want 300", e.Duration)
	}
	if e.DurationVariation != 30 {
		t.Errorf("DurationVariation = %v, want 30 (0.1 of 300 seconds, in seconds)", e.DurationVariation)
	}
	if len(e.Targets) != 3 {
		t.Errorf("Targets = %d, want 3", len(e.Targets))
	}
	if e.ExecuteProportion_20 != 0.2 {
		t.Errorf("ExecuteProportion_20 = %v, want 0.2", e.ExecuteProportion_20)
	}
}

// Forever's biome-conditional effects read Encounter.Biome (Task 8).
// Our envelope carries no biome, so every request built today is
// BiomeUnknown - but the field is set explicitly, so a request that
// silently left it at a zero value of some future different meaning
// would fail here rather than quietly changing a trinket's damage.
func TestBuildSetsTheEncounterBiome(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if got.Encounter.Biome != proto.Biome_BiomeUnknown {
		t.Errorf("Encounter.Biome = %v, want BiomeUnknown for a request with no encounter profile", got.Encounter.Biome)
	}
}

func TestBuildCarriesTheSimOptions(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	if got.SimOptions.Iterations != 3000 {
		t.Errorf("Iterations = %d, want 3000", got.SimOptions.Iterations)
	}
	if got.SimOptions.RandomSeed != 7 {
		t.Errorf("RandomSeed = %d, want 7; a paired run depends on it", got.SimOptions.RandomSeed)
	}
	// IsTest caps concurrency at three splits and adds per-iteration
	// bookkeeping. It is never right for a real run.
	if got.SimOptions.IsTest {
		t.Error("IsTest is set")
	}
}

func TestBuildRejectsWhatItCannotMap(t *testing.T) {
	cases := []struct {
		name string
		mut  func(*api.SimRequest)
		want string
	}{
		{"unknown race", func(r *api.SimRequest) { r.Character.Race = "vulpera" }, "race"},
		{"unknown class", func(r *api.SimRequest) { r.Character.Class = "demon hunter" }, "class"},
		{"unknown slot", func(r *api.SimRequest) { r.Character.Gear[0].Slot = "tabard_of_doom" }, "slot"},
		{"invalid envelope", func(r *api.SimRequest) { r.Iterations = 17 }, "iterations"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := fury()
			req.Character.Gear = append([]api.GearSlot(nil), req.Character.Gear...)
			tc.mut(&req)
			_, err := Build(req)
			if err == nil {
				t.Fatalf("expected an error mentioning %q", tc.want)
			}
			if !contains(err.Error(), tc.want) {
				t.Errorf("error %q does not mention %q", err, tc.want)
			}
		})
	}
}

// Every race the data lane publishes must be buildable, and nothing
// else. The list is the slug field of data/builds/<build>/races.json,
// which has ten rows because Skyborne's two faction variants are two
// rows; a slug this map is missing fails a sim at the boundary with a
// clear message instead of producing an orc.
func TestEveryPublishedRaceSlugMaps(t *testing.T) {
	want := []string{
		"human", "orc", "dwarf", "night-elf", "undead",
		"tauren", "gnome", "troll",
		"high-order-skyborne", "windshaper-skyborne",
	}
	if len(races) != len(want) {
		t.Errorf("the race map has %d entries, want %d; compare it against data/builds/<build>/races.json", len(races), len(want))
	}
	for _, slug := range want {
		if _, ok := ParseRace(slug); !ok {
			t.Errorf("ParseRace(%q) failed; the data lane publishes that race", slug)
		}
	}
	// The two Skyborne rows are distinct races to the engine, because
	// their second active racial differs by faction.
	al, _ := ParseRace("high-order-skyborne")
	ho, _ := ParseRace("windshaper-skyborne")
	if al == ho {
		t.Error("both Skyborne slugs map to the same race; their racials differ by faction")
	}
}

// Build must be deterministic: the same request twice must produce the
// same protobuf, or a paired seed buys nothing. The comparison is
// deterministic wire bytes rather than String(), whose whitespace
// protobuf-go randomises on purpose, so a map ranged without sorting
// shows up here instead of hiding behind a per-process seed.
func TestBuildIsDeterministic(t *testing.T) {
	marshal := googleproto.MarshalOptions{Deterministic: true}
	first, err := buildBytes(marshal)
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < 20; i++ {
		got, err := buildBytes(marshal)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(first, got) {
			t.Fatalf("build %d differs from build 0", i)
		}
	}
}

func buildBytes(marshal googleproto.MarshalOptions) ([]byte, error) {
	req, err := Build(fury())
	if err != nil {
		return nil, err
	}
	return marshal.Marshal(req)
}

func contains(h, n string) bool {
	for i := 0; i+len(n) <= len(h); i++ {
		if h[i:i+len(n)] == n {
			return true
		}
	}
	return false
}

// The browser's worker pool builds one engine request per part, and a
// part's iteration count is never one of api.ValidIterations: 3,000
// over four workers is 750. Options.SplitPart is what lets that build,
// and it relaxes nothing else.
func TestBuildWithSplitPartAcceptsAWorkersShare(t *testing.T) {
	part := fury()
	part.Iterations = 750

	if _, err := Build(part); err == nil {
		t.Error("Build accepted 750 iterations without SplitPart")
	}
	got, err := BuildWith(part, Options{SplitPart: true})
	if err != nil {
		t.Fatalf("BuildWith(SplitPart) rejected a worker's share: %v", err)
	}
	if got.SimOptions.Iterations != 750 {
		t.Errorf("Iterations = %d, want 750", got.SimOptions.Iterations)
	}

	// SplitPart is about the iteration count and nothing else: a part
	// with a level the engine cannot sim is still refused.
	bad := part
	bad.Character.Level = 40
	if _, err := BuildWith(bad, Options{SplitPart: true}); err == nil {
		t.Error("BuildWith(SplitPart) accepted a level the engine cannot sim")
	}
}
