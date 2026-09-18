// Package request turns our JSON SimRequest into the engine's
// RaidSimRequest protobuf.
//
// It is one of exactly two places in the product that touch a protobuf;
// sim/adapter is the other. Both are Go and both run inside the
// browser's wasm as well as on the server, which is what keeps a
// protobuf toolchain out of the front end and stops the mapping being
// written twice in two languages.
package request

import (
	"embed"
	"errors"
	"fmt"

	"github.com/jhunthrop/foreversixty/sim/api"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	ErrUnknownRace  = errors.New("request: unknown race")
	ErrUnknownClass = errors.New("request: unknown class")
	ErrUnknownSlot  = errors.New("request: unknown gear slot")
	ErrUnknownSpec  = errors.New("request: unknown spec")
)

// races and classes map our lower-kebab slugs onto the engine's enums.
// The canonical slug list is data/curated/specs.json and races.json; these
// maps are the engine-side half of that pairing and the test asserts
// every engine enum value is reachable.
// The ten slugs are exactly the `slug` field of every row in
// data/builds/<build>/races.json. Skyborne is one neutral race whose
// faction is chosen at creation and whose second active racial differs
// by faction, so the client's race table carries it as two rows and so
// does this map; Task 4 added the two enum values additively.
var races = map[string]proto.Race{
	"dwarf":               proto.Race_RaceDwarf,
	"gnome":               proto.Race_RaceGnome,
	"human":               proto.Race_RaceHuman,
	"night-elf":           proto.Race_RaceNightElf,
	"orc":                 proto.Race_RaceOrc,
	"tauren":              proto.Race_RaceTauren,
	"troll":               proto.Race_RaceTroll,
	"undead":              proto.Race_RaceUndead,
	"high-order-skyborne": proto.Race_RaceHighOrderSkyborne,
	"windshaper-skyborne": proto.Race_RaceWindshaperSkyborne,
}

var classes = map[string]proto.Class{
	"druid":   proto.Class_ClassDruid,
	"hunter":  proto.Class_ClassHunter,
	"mage":    proto.Class_ClassMage,
	"paladin": proto.Class_ClassPaladin,
	"priest":  proto.Class_ClassPriest,
	"rogue":   proto.Class_ClassRogue,
	"shaman":  proto.Class_ClassShaman,
	"warlock": proto.Class_ClassWarlock,
	"warrior": proto.Class_ClassWarrior,
}

// ParseRace maps a race slug onto the engine's enum.
func ParseRace(slug string) (proto.Race, bool) {
	r, ok := races[slug]
	return r, ok
}

// ParseClass maps a class slug onto the engine's enum.
func ParseClass(slug string) (proto.Class, bool) {
	c, ok := classes[slug]
	return c, ok
}

// Options are the inputs a request needs that the module cannot embed.
//
// Today that is the build's consumable table. It belongs to a build
// rather than to this code, so it is passed in rather than compiled in:
// the api lane loads data/builds/<build>/simconsumes.json once at
// startup and the wasm fetches the same file.
type Options struct {
	// Consumables resolves "item:<id>" consumable ids. Nil means a
	// request may name consumables only by their engine value name; an
	// item id then fails at the boundary rather than silently dropping
	// a consumable the player counted on.
	Consumables *Consumables
}

// Build turns a validated SimRequest into the engine's own request,
// with no build-specific tables. See BuildWith.
func Build(req api.SimRequest) (*proto.RaidSimRequest, error) {
	return BuildWith(req, Options{})
}

// BuildWith turns a validated SimRequest into the engine's own request.
func BuildWith(req api.SimRequest, opt Options) (*proto.RaidSimRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	ch := req.Character

	race, ok := ParseRace(ch.Race)
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownRace, ch.Race)
	}
	class, ok := ParseClass(ch.Class)
	if !ok {
		return nil, fmt.Errorf("%w: %q", ErrUnknownClass, ch.Class)
	}
	equipment, err := equipment(ch.Gear)
	if err != nil {
		return nil, err
	}
	cons, err := consumes(ch.Consumes, opt.Consumables)
	if err != nil {
		return nil, err
	}
	buffs, err := buffsFor(ch.Buffs)
	if err != nil {
		return nil, err
	}

	player := &proto.Player{
		Name:  ch.Name,
		Race:  race,
		Class: class,
		// The engine carries no per-player level: sim/core builds every
		// character at core.CharacterMaxLevel and proto.Player has no
		// level field. api.SimRequest.Validate, which ran above, is
		// what refuses any other level.
		TalentsString: ch.Talents,
		Equipment:     equipment,
		Consumes:      cons,
		Buffs:         buffs.Individual,
	}
	if err := applySpec(player, req.Spec); err != nil {
		return nil, err
	}

	return &proto.RaidSimRequest{
		Raid: &proto.Raid{
			Parties: []*proto.Party{{Players: []*proto.Player{player}, Buffs: buffs.Party}},
			Buffs:   buffs.Raid,
			Debuffs: buffs.Debuffs,
		},
		Encounter: encounter(req.Encounter),
		SimOptions: &proto.SimOptions{
			Iterations: int32(req.Iterations),
			RandomSeed: req.RandomSeed,
			// IsTest caps concurrency at three splits and adds
			// per-iteration bookkeeping; it is never right for a real run.
			IsTest: false,
		},
	}, nil
}

func equipment(gear []api.GearSlot) (*proto.EquipmentSpec, error) {
	items := make([]*proto.ItemSpec, SlotCount)
	for i := range items {
		// Every slot exists, filled or not: the engine indexes this
		// array rather than searching it.
		items[i] = &proto.ItemSpec{}
	}
	for _, g := range gear {
		idx, ok := SlotIndex(g.Slot)
		if !ok {
			return nil, fmt.Errorf("%w: %q", ErrUnknownSlot, g.Slot)
		}
		items[idx] = &proto.ItemSpec{
			Id:           int32(g.ItemID),
			Enchant:      int32(g.Enchant),
			RandomSuffix: int32(g.Suffix),
		}
	}
	return &proto.EquipmentSpec{Items: items}, nil
}

// biomeFor maps an encounter profile onto a biome. The profile is
// ignored until there is a table to read: today every profile
// is BiomeUnknown: the data lane's zones.json has no biome column, so
// there is nothing to map from, and inventing one would put a damage
// multiplier on a guess. The function exists so the mapping has one
// home when that table lands.
func biomeFor(_ string) proto.Biome {
	return proto.Biome_BiomeUnknown
}

// The target every sim fights. A boss is three levels above the player,
// which is what the attack table's suppression terms are derived
// against; the id and name are the engine's own target dummy.
const (
	targetDummyID   = 31146
	targetDummyName = "Target Dummy"
	targetBossLevel = 63
	targetNotTanked = -1
	// The engine's own UI preset opens a fury pull at no rage.
	defaultStartingRage = 0
)

func encounter(e api.EncounterSpec) *proto.Encounter {
	targets := make([]*proto.Target, e.Targets)
	for i := range targets {
		targets[i] = &proto.Target{
			Id:        targetDummyID,
			Name:      targetDummyName,
			Level:     targetBossLevel,
			MobType:   proto.MobType_MobTypeHumanoid,
			TankIndex: targetNotTanked,
		}
	}
	return &proto.Encounter{
		Duration: float64(e.DurationSec),
		// Forever's biome-conditional trinkets read the encounter's
		// biome (Task 8). Our envelope has no biome field - the contract
		// gives EncounterSpec a Profile and nothing else - so a plain
		// sim is BiomeUnknown, which matches nothing, which is exactly a
		// vanilla fight. Naming it here rather than leaving the field at
		// its zero value is what gives the encounter-profile feature one
		// place to fill in when the data lane publishes a zone-to-biome
		// table; today zones.json carries no biome column.
		Biome: biomeFor(e.Profile),
		// The engine's variation is in seconds; ours is a fraction of
		// the duration, because that is what the settings bar offers.
		DurationVariation:    float64(e.DurationSec) * e.Variation,
		ExecuteProportion_20: e.ExecuteRatio,
		ExecuteProportion_25: e.ExecuteRatio,
		ExecuteProportion_35: e.ExecuteRatio,
		Targets:              targets,
	}
}

// aplFS carries the launch specs' default rotations, so the wasm needs no
// fetch to attach one. They are the engine's own presets, copied from
// ui/<class>/apls; Tasks 11 and 12 replace them with the Forever specs'
// own, and the request builder's job is only to attach *a* rotation.
//
//go:embed apl/*.apl.json
var aplFS embed.FS

// specOptions is the per-spec half of the player: it attaches the
// engine's spec oneof, whose concrete type the engine's own package
// owns. The spec's default rotation is the embedded APL of the same
// name, so apl/<slug>.apl.json and this table stay in step.
//
// data/curated/specs.json does not exist yet, so this table is the only
// spec list in the module and it fails closed: an unsupported spec is an
// error at the boundary rather than a player the engine cannot build an
// agent for. The option values are the engine's own UI presets
// (ui/<class>/presets.ts); Tasks 11 and 12 own what Forever's specs
// default to.
var specOptions = map[string]func(*proto.Player){
	"warrior-fury": func(p *proto.Player) {
		p.Spec = &proto.Player_Warrior{Warrior: &proto.Warrior{
			Options: &proto.Warrior_Options{
				StartingRage: defaultStartingRage,
				Shout:        proto.WarriorShout_WarriorShoutBattle,
			},
		}}
	},
	"mage-frost": func(p *proto.Player) {
		p.Spec = &proto.Player_Mage{Mage: &proto.Mage{
			Options: &proto.Mage_Options{Armor: proto.Mage_Options_MoltenArmor},
		}}
	},
}

// applySpec attaches the spec's options and its default rotation. The
// engine cannot build an agent for a player carrying neither.
func applySpec(player *proto.Player, slug string) error {
	apply, ok := specOptions[slug]
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownSpec, slug)
	}
	rot, err := rotation(slug)
	if err != nil {
		return err
	}
	apply(player)
	player.Rotation = rot
	return nil
}

// rotation parses one embedded APL. It is parsed per build rather than
// cached, so no two requests ever share a mutable rotation.
func rotation(name string) (*proto.APLRotation, error) {
	b, err := aplFS.ReadFile("apl/" + name + ".apl.json")
	if err != nil {
		return nil, fmt.Errorf("request: no embedded APL for %q: %w", name, err)
	}
	apl := &proto.APLRotation{}
	if err := protojson.Unmarshal(b, apl); err != nil {
		return nil, fmt.Errorf("request: the embedded APL for %q is corrupt: %w", name, err)
	}
	return apl, nil
}
