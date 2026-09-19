package request

import (
	"bytes"
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/jhunthrop/foreversixty/sim/enginever"
	"github.com/jhunthrop/foreversixty/sim/internal/strcase"
	"github.com/wowsims/classic/sim/core/proto"
	googleproto "google.golang.org/protobuf/proto"
)

func fury() api.SimRequest {
	return api.SimRequest{
		EngineVersion: enginever.Version,
		Spec:          "warrior-fury",
		Source:        api.CharacterSource{Kind: api.SourceManual},
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
//
// The old version of this test asserted that biomeFor's stub returned
// its own constant and could not fail. What can fail is the profile
// reaching here at all: an "encounter:<id>" used to be accepted and
// then ignored, producing a sim byte-identical to a blank one under an
// encounter's name.
func TestTheEncounterProfileIsEitherHonouredOrRefused(t *testing.T) {
	for _, profile := range []string{"", api.ProfilePatchwerk} {
		req := fury()
		req.Encounter.Profile = profile
		got, err := Build(req)
		if err != nil {
			t.Fatalf("profile %q: %v", profile, err)
		}
		// A patchwerk IS what the sim builds - a stationary target and
		// nothing else - so it is the same fight as the blank profile,
		// and neither matches a biome-conditional trinket.
		if got.Encounter.Biome != proto.Biome_BiomeUnknown {
			t.Errorf("profile %q: Biome = %v, want BiomeUnknown", profile, got.Encounter.Biome)
		}
	}
	req := fury()
	req.Encounter.Profile = "encounter:onyxia"
	if _, err := Build(req); err == nil {
		t.Error("an encounter profile was accepted; the sim has no encounter table, so the run would be a patchwerk under another name")
	}
}

// The three execute proportions are the share of the fight spent below
// 20%, 25% and 35% health. They are NESTED, so setting all three to one
// number - which is what the builder did - describes a fight no health
// bar can produce, inflates the Execute window by 25% at the default
// ratio and understates the sub-35% one by 29%.
func TestExecuteWindowsAreNestedNotEqual(t *testing.T) {
	// The engine's own reference encounters are {0.2, 0.25, 0.35}: the
	// proportion equals the threshold, which is a target whose health
	// falls at a steady rate. A ratio of 0.2 must reproduce it exactly,
	// or the shape this derivation assumes is not the engine's.
	b20, b25, b35 := executeProportions(0.2)
	for _, tc := range []struct {
		got, want float64
		name      string
	}{{b20, 0.2, "below20"}, {b25, 0.25, "below25"}, {b35, 0.35, "below35"}} {
		if math.Abs(tc.got-tc.want) > 1e-9 {
			t.Errorf("at ratio 0.2, %s = %v, want the engine's own %v", tc.name, tc.got, tc.want)
		}
	}

	req := fury()
	req.Encounter.ExecuteRatio = 0.25
	got, err := Build(req)
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if e.ExecuteProportion_20 != 0.25 {
		t.Errorf("ExecuteProportion_20 = %v, want the requested 0.25: the control names the Execute window", e.ExecuteProportion_20)
	}
	if !(e.ExecuteProportion_20 < e.ExecuteProportion_25 && e.ExecuteProportion_25 < e.ExecuteProportion_35) {
		t.Errorf("the windows are not nested: %v, %v, %v", e.ExecuteProportion_20, e.ExecuteProportion_25, e.ExecuteProportion_35)
	}

	// The whole fight below 20% health means the whole fight below 25%
	// and 35% too: nothing may exceed one.
	b20, b25, b35 = executeProportions(1)
	if b20 != 1 || b25 != 1 || b35 != 1 {
		t.Errorf("at ratio 1 the windows are %v, %v, %v; a proportion over 1 is not a proportion", b20, b25, b35)
	}
	if b20, b25, b35 = executeProportions(0); b20 != 0 || b25 != 0 || b35 != 0 {
		t.Errorf("at ratio 0 the windows are %v, %v, %v", b20, b25, b35)
	}
}

// The engine carries two profession slots and reads them for
// self-only recipes and effects. They were accepted at the boundary and
// never looked at again, so a sim ran without the Engineering trinket
// the player counted on and said nothing.
func TestProfessionsReachThePlayer(t *testing.T) {
	req := fury()
	req.Character.Profession = []string{"engineering", "blacksmithing"}
	got, err := Build(req)
	if err != nil {
		t.Fatal(err)
	}
	p := got.Raid.Parties[0].Players[0]
	if p.Profession1 != proto.Profession_Engineering {
		t.Errorf("Profession1 = %v, want Engineering", p.Profession1)
	}
	if p.Profession2 != proto.Profession_Blacksmithing {
		t.Errorf("Profession2 = %v, want Blacksmithing", p.Profession2)
	}

	none := fury()
	got, err = Build(none)
	if err != nil {
		t.Fatal(err)
	}
	p = got.Raid.Parties[0].Players[0]
	if p.Profession1 != proto.Profession_ProfessionUnknown || p.Profession2 != proto.Profession_ProfessionUnknown {
		t.Errorf("a character with no professions got %v and %v", p.Profession1, p.Profession2)
	}

	for _, tc := range []struct {
		name string
		list []string
		want error
	}{
		{"a profession the engine has no enum for", []string{"cooking"}, ErrUnknownProfession},
		{"a typo", []string{"Engineering"}, ErrUnknownProfession},
		{"three of them", []string{"mining", "tailoring", "alchemy"}, ErrTooManyProfession},
		{"one of them twice", []string{"mining", "mining"}, ErrDuplicateProfess},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bad := fury()
			bad.Character.Profession = tc.list
			if _, err := Build(bad); !errors.Is(err, tc.want) {
				t.Errorf("Build returned %v, want %v", err, tc.want)
			}
		})
	}
}

// Two rings in finger1 used to equip one and lose the other with no
// word, and the character the sim reported on was not the one the
// planner sent.
func TestADuplicateGearSlotIsRefused(t *testing.T) {
	req := fury()
	req.Character.Gear = append(req.Character.Gear, api.GearSlot{Slot: "main_hand", ItemID: 12345})
	_, err := Build(req)
	if !errors.Is(err, ErrDuplicateSlot) {
		t.Fatalf("Build returned %v, want ErrDuplicateSlot", err)
	}
	if !strings.Contains(err.Error(), "12345") || !strings.Contains(err.Error(), "19352") {
		t.Errorf("the error names neither item: %v", err)
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
			if !strings.Contains(err.Error(), tc.want) {
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

// The browser's worker pool builds one engine request per part, and a
// part's iteration count is never one of api.ValidIterations: 3,000
// over four workers is 750. Options.OpenIterations is what lets that build,
// and it relaxes nothing else.
func TestBuildWithOpenIterationsAcceptsAWorkersShare(t *testing.T) {
	part := fury()
	part.Iterations = 750

	if _, err := Build(part); err == nil {
		t.Error("Build accepted 750 iterations without OpenIterations")
	}
	got, err := BuildWith(part, Options{OpenIterations: true})
	if err != nil {
		t.Fatalf("BuildWith(OpenIterations) rejected a worker's share: %v", err)
	}
	if got.SimOptions.Iterations != 750 {
		t.Errorf("Iterations = %d, want 750", got.SimOptions.Iterations)
	}

	// OpenIterations is about the iteration count and nothing else: a part
	// with a level the engine cannot sim is still refused.
	bad := part
	bad.Character.Level = 40
	if _, err := BuildWith(bad, Options{OpenIterations: true}); err == nil {
		t.Error("BuildWith(OpenIterations) accepted a level the engine cannot sim")
	}
}

// The five encounter fields the parity contract added all reach the
// engine, and the target the sim has always built is unchanged when
// none of them is set.
func TestEncounterCarriesTheParityFields(t *testing.T) {
	req := fury()
	req.Encounter.Movement = &api.Movement{IntervalSec: 20, DurationSec: 5, Kind: api.MovementCasting}
	req.Encounter.TargetsOverTime = []api.TargetCount{{AtSec: 0, Count: 1}, {AtSec: 40, Count: 3}}
	req.Encounter.TargetLevel = 61
	req.Encounter.TargetArmor = 2500
	req.Encounter.TargetType = "undead"
	req.Encounter.Dummy = true

	got, err := Build(req)
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if m := e.GetMovement(); m == nil || m.IntervalSeconds != 20 || m.DurationSeconds != 5 || !m.CastingOnly {
		t.Errorf("movement = %+v", e.GetMovement())
	}
	if len(e.GetTargetsOverTime()) != 2 || e.GetTargetsOverTime()[1].AtSeconds != 40 || e.GetTargetsOverTime()[1].Count != 3 {
		t.Errorf("targets_over_time = %+v", e.GetTargetsOverTime())
	}
	if !e.GetTargetDummy() {
		t.Error("target_dummy is not set")
	}
	// A timeline overrides the fixed count and the site sends ONE
	// target: the engine pads the list up to the timeline's maximum
	// itself (contract 10.3).
	if len(e.Targets) != 1 {
		t.Fatalf("the request carries %d targets; a timeline sends one and the engine pads", len(e.Targets))
	}
	for i, target := range e.Targets {
		if target.Level != 61 {
			t.Errorf("target %d is level %d", i, target.Level)
		}
		if target.MobType != proto.MobType_MobTypeUndead {
			t.Errorf("target %d is %v", i, target.MobType)
		}
		if got := target.Stats[proto.Stat_StatArmor]; got != 2500 {
			t.Errorf("target %d armor = %v, want the override 2500", i, got)
		}
	}
}

// Nothing set is the fight the sim has always built, plus the armor
// preset the contract now says every target carries.
func TestEncounterDefaultsAreUnchanged(t *testing.T) {
	got, err := Build(fury())
	if err != nil {
		t.Fatal(err)
	}
	e := got.Encounter
	if e.GetMovement() != nil || len(e.GetTargetsOverTime()) != 0 || e.GetTargetDummy() {
		t.Errorf("a plain request set a parity field: %+v", e)
	}
	if len(e.Targets) != 1 {
		t.Fatalf("targets = %d", len(e.Targets))
	}
	target := e.Targets[0]
	if target.Level != api.BossLevel {
		t.Errorf("level = %d, want %d", target.Level, api.BossLevel)
	}
	if target.MobType != proto.MobType_MobTypeHumanoid {
		t.Errorf("mob_type = %v, want humanoid, which is what the sim has always fought", target.MobType)
	}
	if got := target.Stats[proto.Stat_StatArmor]; got != float64(api.TargetArmorByLevel[api.BossLevel]) {
		t.Errorf("armor = %v, want the boss preset %d", got, api.TargetArmorByLevel[api.BossLevel])
	}
}

// The target-type vocabulary is the engine's enum, and a name that
// resolved to nothing would fight a creature with no type and change
// what Hunter and Warlock abilities do without a word.
func TestTargetTypesMatchTheEngineEnum(t *testing.T) {
	for _, id := range api.TargetTypes {
		if _, ok := mobTypes[id]; !ok {
			t.Errorf("api.TargetTypes lists %q, which this package cannot map", id)
		}
	}
	for value, name := range proto.MobType_name {
		id := strcase.Snake(strings.TrimPrefix(name, "MobType"))
		if _, ok := mobTypes[id]; !ok {
			t.Errorf("the engine has MobType %s (%d) and no id maps to it", name, value)
		}
	}
	if len(mobTypes) != len(proto.MobType_name) {
		t.Errorf("mobTypes has %d entries, the enum has %d", len(mobTypes), len(proto.MobType_name))
	}
}
