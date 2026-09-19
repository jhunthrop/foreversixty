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
	"github.com/jhunthrop/foreversixty/sim/specs"
	"github.com/wowsims/classic/sim/core/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	ErrUnknownRace  = errors.New("request: unknown race")
	ErrUnknownClass = errors.New("request: unknown class")
	ErrUnknownSlot  = errors.New("request: unknown gear slot")
	ErrUnknownSpec  = errors.New("request: unknown spec")
	// ErrUnsupportedSpec is a spec on sim/specs' canonical list that
	// this engine build has no agent for yet.
	ErrUnsupportedSpec   = errors.New("request: unsupported spec")
	ErrDuplicateSlot     = errors.New("request: two items in one gear slot")
	ErrUnknownProfession = errors.New("request: unknown profession")
	ErrTooManyProfession = errors.New("request: a character has at most two professions")
	ErrDuplicateProfess  = errors.New("request: one profession listed twice")
	// ErrSpecClassMismatch is returned when the spec and the class
	// disagree. The engine would build the player from the class and the
	// agent from the spec, and a warrior would run a mage's rotation.
	ErrSpecClassMismatch = errors.New("request: the spec does not belong to the character's class")
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

// professions maps our lower-kebab slugs onto the engine's enum. The
// engine models a profession as a source of self-only recipes and
// effects (Engineering's grenades and trinkets, Blacksmithing's socket),
// so an unrecognised one is refused rather than dropped: a sim that
// quietly ran without the profession the player counted on would report
// a wrong number and say nothing about why.
var professions = map[string]proto.Profession{
	"alchemy":        proto.Profession_Alchemy,
	"blacksmithing":  proto.Profession_Blacksmithing,
	"enchanting":     proto.Profession_Enchanting,
	"engineering":    proto.Profession_Engineering,
	"herbalism":      proto.Profession_Herbalism,
	"leatherworking": proto.Profession_Leatherworking,
	"mining":         proto.Profession_Mining,
	"skinning":       proto.Profession_Skinning,
	"tailoring":      proto.Profession_Tailoring,
}

// ParseProfession maps a profession slug onto the engine's enum.
func ParseProfession(slug string) (proto.Profession, bool) {
	p, ok := professions[slug]
	return p, ok
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

	// OpenIterations validates the request with
	// api.SimRequest.ValidatePart instead of Validate: everything except
	// the closed set of iteration counts the settings bar offers.
	//
	// Two callers need it and both are the same shape - a run whose
	// iteration count nobody chose from the UI. The browser's worker
	// pool splits 3,000 four ways and each part asks for 750; the
	// operator running forever-sim passes -iterations 100 to reproduce
	// something quickly. A whole request from a client never sets it.
	OpenIterations bool
}

// Build turns a validated SimRequest into the engine's own request,
// with no build-specific tables. See BuildWith.
func Build(req api.SimRequest) (*proto.RaidSimRequest, error) {
	return BuildWith(req, Options{})
}

// BuildWith turns a validated SimRequest into the engine's own request.
func BuildWith(req api.SimRequest, opt Options) (*proto.RaidSimRequest, error) {
	validate := req.Validate
	if opt.OpenIterations {
		validate = req.ValidatePart
	}
	if err := validate(); err != nil {
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
	if err := checkSpecClass(req.Spec, ch.Class); err != nil {
		return nil, err
	}
	equipment, err := equipment(ch.Gear)
	if err != nil {
		return nil, err
	}
	first, second, err := professionsFor(ch.Profession)
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
		Profession1:   first,
		Profession2:   second,
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

// checkSpecClass refuses a spec that belongs to another class. Without
// it the player is built from Character.Class and the agent from Spec,
// so an orc warrior carrying spec "mage-frost" reaches the engine as a
// warrior running a frost mage's rotation and the result looks like a
// number rather than a mistake. sim/specs is the authoritative pairing.
func checkSpecClass(spec, class string) error {
	known, ok := specs.ByKey[spec]
	if !ok {
		return fmt.Errorf("%w: %q", ErrUnknownSpec, spec)
	}
	if known.ClassSlug != class {
		return fmt.Errorf("%w: %q is a %s spec, but character.class is %q", ErrSpecClassMismatch, spec, known.ClassSlug, class)
	}
	return nil
}

// professionsFor maps the character's professions onto the engine's two
// slots. The engine carries exactly two, so a third is an error rather
// than a silently dropped profession.
func professionsFor(slugs []string) (proto.Profession, proto.Profession, error) {
	if len(slugs) > 2 {
		return 0, 0, fmt.Errorf("%w, got %d: %v", ErrTooManyProfession, len(slugs), slugs)
	}
	var out [2]proto.Profession
	for i, slug := range slugs {
		p, ok := ParseProfession(slug)
		if !ok {
			return 0, 0, fmt.Errorf("%w: %q", ErrUnknownProfession, slug)
		}
		// Two slots holding one profession is a client that meant to
		// send two, and taking it would silently halve what the
		// character has.
		if i == 1 && p == out[0] {
			return 0, 0, fmt.Errorf("%w: %q", ErrDuplicateProfess, slug)
		}
		out[i] = p
	}
	return out[0], out[1], nil
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
		// Last-wins would equip one of two rings and lose the other
		// without a word, and the character the sim reports on would
		// not be the one the planner sent.
		if items[idx].Id != 0 {
			return nil, fmt.Errorf("%w: %q holds both item %d and item %d", ErrDuplicateSlot, g.Slot, items[idx].Id, g.ItemID)
		}
		items[idx] = &proto.ItemSpec{
			Id:           int32(g.ItemID),
			Enchant:      int32(g.Enchant),
			RandomSuffix: int32(g.Suffix),
		}
	}
	return &proto.EquipmentSpec{Items: items}, nil
}

// biomeFor maps an encounter profile onto a biome.
//
// Only the two profiles that mean "a stationary target and nothing else"
// reach it: api.SimRequest.Validate refuses an "encounter:<id>" outright,
// because the data lane's zones.json has no biome column and a sim that
// answered one would be a patchwerk wearing an encounter's name. So
// every profile here is BiomeUnknown, which matches no biome-conditional
// trinket, which is exactly a vanilla fight. The function is where the
// mapping goes when that table lands, and it is total by construction:
// an unknown profile is already an error.
func biomeFor(_ string) proto.Biome {
	return proto.Biome_BiomeUnknown
}

// The target every sim fights. The id and name are the engine's own
// target dummy; the level is api.BossLevel, which sim/measure reads too.
const (
	targetDummyID   = 31146
	targetDummyName = "Target Dummy"
	targetNotTanked = -1
	// The engine's own UI preset opens a fury pull at no rage.
	defaultStartingRage = 0
)

// The three health thresholds the engine's Encounter carries one
// proportion for each of. They are percentages of the target's health,
// and they are what makes the three fields three different questions:
// ExecuteProportion_20 is the share of the fight spent below 20% health,
// not the share spent in "the execute window".
const (
	executeThreshold20 = 20.0
	executeThreshold25 = 25.0
	executeThreshold35 = 35.0
)

// executeProportions turns the settings bar's one execute_ratio into the
// engine's three nested windows.
//
// The ratio is the sub-20% share, because that is the window the control
// is named for: Execute, Hammer of Wrath and Improved Expose Weakness
// all start at 20%. The other two follow from one assumption about the
// fight's shape - that the target's health falls at a steady rate, so
// the time spent below X% is proportional to X. That assumption is the
// engine's own: its reference encounters are {0.2, 0.25, 0.35}, which is
// exactly what this returns for a ratio of 0.2.
//
// Setting all three to one number, as an earlier draft did, inflates the
// Execute window by 25% at the default ratio and understates the sub-35%
// one by 29%, and it describes a fight no health bar can produce: the
// three are nested, so they can only be equal at 0 and at 1.
func executeProportions(ratio float64) (below20, below25, below35 float64) {
	scale := func(threshold float64) float64 {
		return min(ratio*threshold/executeThreshold20, 1)
	}
	return scale(executeThreshold20), scale(executeThreshold25), scale(executeThreshold35)
}

func encounter(e api.EncounterSpec) *proto.Encounter {
	below20, below25, below35 := executeProportions(e.ExecuteRatio)
	targets := make([]*proto.Target, e.Targets)
	for i := range targets {
		targets[i] = &proto.Target{
			Id:        targetDummyID,
			Name:      targetDummyName,
			Level:     api.BossLevel,
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
		ExecuteProportion_20: below20,
		ExecuteProportion_25: below25,
		ExecuteProportion_35: below35,
		Targets:              targets,
	}
}

// aplFS carries the launch specs' default rotations, so the wasm needs
// no fetch to attach one.
//
// These two files are still the engine's Era presets, copied from
// ui/<class>/apls, NOT the canonical rotations in data/curated/apl. The
// canonical source is the curated file and `make apl-check` compares
// the engine fork's own forever_<spec>.apl.json against it; bringing
// these two into line moves both adapter fixtures, so it is done in the
// same pass as the fixture regeneration rather than on its own.
//
//go:embed apl/*.apl.json
var aplFS embed.FS

// specOptions is the per-spec half of the player: it attaches the
// engine's spec oneof, whose concrete type the engine's own package
// owns. The spec's default rotation is the embedded APL of the same
// name, so apl/<slug>.apl.json and this table stay in step.
//
// The canonical spec list is sim/specs, generated from
// data/curated/specs.json, and checkSpecClass above is what refuses a
// spec that is not on it or does not match the class. This table is the
// narrower question of which of those specs the engine can build an
// agent for today, and it fails closed: an unsupported spec is an error
// at the boundary rather than a player with no rotation. The option
// values are the engine's own UI presets (ui/<class>/presets.ts).
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
		return fmt.Errorf("%w: %q; the launch specs are %v", ErrUnsupportedSpec, slug, supportedSpecs())
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

// supportedSpecs lists the specs specOptions can build an agent for, in
// sim/specs' canonical order so the message is stable.
func supportedSpecs() []string {
	out := make([]string, 0, len(specOptions))
	for _, s := range specs.All {
		if _, ok := specOptions[s.Spec]; ok {
			out = append(out, s.Spec)
		}
	}
	return out
}
