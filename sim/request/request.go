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

	// NoSampleIteration switches off the engine's sample cast log.
	//
	// One caller sets it: sim/bulk's stage requests. A bulk stage is
	// dozens of runs and the page shows one cast log, so recording
	// one per combination is paid for on every sim and read on none.
	// A plain run always records it, because the report's sample card
	// is not optional.
	NoSampleIteration bool
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
	cooldowns, err := cooldownsFor(ch.Cooldowns, opt.Consumables)
	if err != nil {
		return nil, err
	}

	player := &proto.Player{
		Name:  ch.Name,
		Race:  race,
		Class: class,
		// The engine builds the character at this level (sim/core's
		// EffectiveCharacterLevel): base stats, health, mana and which
		// spell ranks register all follow it, as the rotation's spell
		// ranks (rewriteRotationRanks) and the target's level already do.
		// The envelope bounds it to 1..api.MaxLevel.
		Level:         engineCharacterLevel(ch),
		TalentsString: ch.Talents,
		Equipment:     equipment,
		Consumes:      cons,
		Cooldowns:     cooldowns,
		Buffs:         buffs.Individual,
		Profession1:   first,
		Profession2:   second,
		// This lane's brief (bis-ranker-integrity-5, item 7): passed
		// straight through - api.CharacterSpec.DistanceFromTarget's own
		// doc explains what zero (the default, every request before this
		// field existed) does and why a caller sets it.
		DistanceFromTarget: ch.DistanceFromTarget,
	}
	if err := applySpec(player, req.Spec, ch.Class, ch.Level); err != nil {
		return nil, err
	}

	return &proto.RaidSimRequest{
		Raid: &proto.Raid{
			Parties: []*proto.Party{{Players: []*proto.Player{player}, Buffs: buffs.Party}},
			Buffs:   buffs.Raid,
			Debuffs: buffs.Debuffs,
		},
		Encounter: encounter(req.Encounter, ch.Level),
		SimOptions: &proto.SimOptions{
			Iterations: int32(req.Iterations),
			RandomSeed: req.RandomSeed,
			// IsTest caps concurrency at three splits and adds
			// per-iteration bookkeeping; it is never right for a real run.
			IsTest: false,
			// The median-DPS iteration's cast log, for the report's
			// sample card. See Options.NoSampleIteration.
			SampleIteration: !opt.NoSampleIteration,
		},
	}, nil
}

// engineCharacterLevel is the level BuildWith hands the engine's
// proto.Player. It is ch.Level unchanged - the envelope already bounds it
// to 1..api.MaxLevel - kept as its own named function so every place a
// level belongs reads the same one.
func engineCharacterLevel(ch api.CharacterSpec) int32 {
	return int32(ch.Level)
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

// secondaryProfessions are the client's secondary skills as the addon
// slugs them (Export.slugify of GetProfessionInfo's name): every
// character can hold all of them beside two primaries, and the engine
// models nothing for any of them. They are skipped, not refused -- an
// export listing leatherworking, enchanting and cooking is a character
// with two professions, not three.
var secondaryProfessions = map[string]bool{
	"cooking":   true,
	"fishing":   true,
	"first-aid": true,
}

// professionsFor maps the character's primary professions onto the
// engine's two slots. The engine carries exactly two, so a third primary
// is an error rather than a silently dropped profession.
func professionsFor(slugs []string) (proto.Profession, proto.Profession, error) {
	primaries := make([]string, 0, len(slugs))
	for _, slug := range slugs {
		if !secondaryProfessions[slug] {
			primaries = append(primaries, slug)
		}
	}
	if len(primaries) > 2 {
		return 0, 0, fmt.Errorf("%w, got %d: %v", ErrTooManyProfession, len(primaries), primaries)
	}
	var out [2]proto.Profession
	for i, slug := range primaries {
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
	// The feral UI preset's assumed player reaction time, in
	// milliseconds; the spec's rotation code shifts its clip windows by
	// it and zero would model a player who never misses a tick.
	defaultLatencyMS = 100
	// A hunter pet that is alive for the whole fight. The UI preset's
	// default, and the only honest one for a patchwerk.
	fullPetUptime = 1.0
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

// mobTypes maps our target-type ids onto the engine's enum. The ids are
// the enum names in lower snake case with the prefix stripped, and
// TestTargetTypesMatchTheEngineEnum holds the map to the enum in both
// directions: a creature type the engine models and we cannot name is
// a Hunter's Slaying bonus nobody can ask for.
var mobTypes = map[string]proto.MobType{
	"beast":               proto.MobType_MobTypeBeast,
	"demon":               proto.MobType_MobTypeDemon,
	"dragonkin":           proto.MobType_MobTypeDragonkin,
	"elemental":           proto.MobType_MobTypeElemental,
	"giant":               proto.MobType_MobTypeGiant,
	"humanoid":            proto.MobType_MobTypeHumanoid,
	"mechanical":          proto.MobType_MobTypeMechanical,
	"undead":              proto.MobType_MobTypeUndead,
	api.TargetTypeUnknown: proto.MobType_MobTypeUnknown,
}

// targetCount is how many targets the encounter needs BUILT.
//
// A target-count timeline overrides the fixed count, and the SITE
// SENDS ONE TARGET for it: the engine pads the list by repeating the
// last target up to the timeline's maximum (contract 10.3), so
// building the pool here would be the same work done twice and would
// disagree with the engine the day the padding rule changes.
func targetCount(e api.EncounterSpec) int {
	if len(e.TargetsOverTime) > 0 {
		return 1
	}
	return e.Targets
}

// targetStats is one target's stat array, with the encounter's armor in
// it. The engine indexes the array by proto.Stat, so it is built to the
// enum's length rather than to the highest index we happen to set.
//
// level is the RESOLVED target level - encounter's own computation of
// e.TargetLevel or its level-dependent default, never the raw field -
// so this always asks TargetArmorFor for the level the target was
// actually built at. Passing e.TargetLevel straight through here used
// to work only because TargetArmorFor's map-miss fallback happened to
// equal the one default (BossLevel) every request used; now that the
// default is the character's own level plus three, a target level of
// 0 and an unresolved default no longer mean the same armor.
func targetStats(level int, override *int) []float64 {
	out := make([]float64, len(proto.Stat_name))
	out[proto.Stat_StatArmor] = float64(api.TargetArmorFor(level, override))
	return out
}

// movement turns our window into the engine's pattern. Our two kinds
// are the engine's one boolean: "casting" interrupts spells without
// moving, "away" leaves melee range too.
func movement(m *api.Movement) *proto.MovementPattern {
	if m == nil {
		return nil
	}
	return &proto.MovementPattern{
		IntervalSeconds: float64(m.IntervalSec),
		DurationSeconds: float64(m.DurationSec),
		CastingOnly:     m.Kind == api.MovementCasting,
	}
}

// targetsOverTime turns our timeline into the engine's.
func targetsOverTime(steps []api.TargetCount) []*proto.TargetCountAt {
	if len(steps) == 0 {
		return nil
	}
	out := make([]*proto.TargetCountAt, len(steps))
	for i, s := range steps {
		out[i] = &proto.TargetCountAt{AtSeconds: float64(s.AtSec), Count: int32(s.Count)}
	}
	return out
}

// encounter builds the target(s) a character of characterLevel fights.
// A request naming no target_level fights api.DefaultTargetLevel(characterLevel)
// - three above the character, the same offset a level-MaxLevel
// character always fought (api.BossLevel) before a request could name
// any other character level.
func encounter(e api.EncounterSpec, characterLevel int) *proto.Encounter {
	below20, below25, below35 := executeProportions(e.ExecuteRatio)
	// A target dummy has no execute window: nothing kills it, so its
	// health never falls. The envelope's Dummy flag is the one place
	// that says so, rather than the page being asked to zero the ratio
	// as well as tick the box.
	if e.Dummy {
		below20, below25, below35 = 0, 0, 0
	}
	level := e.TargetLevel
	if level == 0 {
		level = api.DefaultTargetLevel(characterLevel)
	}
	mob, ok := mobTypes[e.TargetType]
	if !ok {
		// "" is not a type the page offers; it is the shape of every
		// request written before the field existed, and those fought a
		// humanoid. Keeping that is what stops the pin bump changing
		// every stored spec's number.
		mob = proto.MobType_MobTypeHumanoid
	}
	targets := make([]*proto.Target, targetCount(e))
	for i := range targets {
		targets[i] = &proto.Target{
			Id:        targetDummyID,
			Name:      targetDummyName,
			Level:     int32(level),
			MobType:   mob,
			Stats:     targetStats(level, e.TargetArmor),
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
		Movement:             movement(e.Movement),
		TargetsOverTime:      targetsOverTime(e.TargetsOverTime),
		TargetDummy:          e.Dummy,
	}
}

// aplFS carries every written spec's default rotation, so the wasm needs
// no fetch to attach one.
//
// These files are GENERATED, by `make apl-sync`: each is the `rotation`
// block of data/curated/apl/<spec>.json, which is the one place a
// rotation is edited. The engine fork's ui/<class>/apls/forever_<spec>.apl.json
// is the other copy of the same source, and `make apl-check` proves both
// against it. Editing one here would make the artifacts measure a
// rotation nobody wrote down.
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
// One class the engine models as a single package serves every one of
// its specs from one set of options - the fork has one hunter, mage,
// rogue and warlock package, and one DPS warrior package for Arms and
// Fury - so those specs share a function here rather than repeating
// it. A spec the fork models on its own gets its own.
var specOptions = map[string]func(*proto.Player){
	"druid-balance":        balanceDruidOptions,
	"druid-feral":          feralDruidOptions,
	"druid-restoration":    restorationDruidOptions,
	"hunter-beast-mastery": hunterOptions,
	"hunter-marksmanship":  hunterOptions,
	"hunter-survival":      hunterOptions,
	"mage-arcane":          mageOptions,
	"mage-fire":            mageOptions,
	"mage-frost":           mageOptions,
	"paladin-holy":         holyPaladinOptions,
	"paladin-retribution":  retributionPaladinOptions,
	"priest-discipline":    healingPriestOptions,
	"priest-holy":          healingPriestOptions,
	"priest-shadow":        shadowPriestOptions,
	"rogue-assassination":  rogueOptions,
	"rogue-combat":         rogueOptions,
	"rogue-subtlety":       rogueOptions,
	"shaman-elemental":     elementalShamanOptions,
	"shaman-enhancement":   enhancementShamanOptions,
	"shaman-restoration":   restorationShamanOptions,
	"warlock-affliction":   warlockOptions,
	"warlock-demonology":   warlockOptions,
	"warlock-destruction":  warlockOptions,
	"warrior-arms":         warriorOptions,
	"warrior-fury":         warriorOptions,
}

func balanceDruidOptions(p *proto.Player) {
	p.Spec = &proto.Player_BalanceDruid{BalanceDruid: &proto.BalanceDruid{
		// The engine reads the innervate target through this reference
		// and a nil one is a nil dereference, not a druid who innervates
		// nobody. The empty reference is the UI's own default: self.
		Options: &proto.BalanceDruid_Options{InnervateTarget: &proto.UnitReference{}},
	}}
}

func feralDruidOptions(p *proto.Player) {
	p.Spec = &proto.Player_FeralDruid{FeralDruid: &proto.FeralDruid{
		Options: &proto.FeralDruid_Options{
			InnervateTarget: &proto.UnitReference{},
			LatencyMs:       defaultLatencyMS,
		},
	}}
}

func restorationDruidOptions(p *proto.Player) {
	p.Spec = &proto.Player_RestorationDruid{RestorationDruid: &proto.RestorationDruid{
		// As for the balance druid: the empty reference is self.
		Options: &proto.RestorationDruid_Options{InnervateTarget: &proto.UnitReference{}},
	}}
}

func hunterOptions(p *proto.Player) {
	p.Spec = &proto.Player_Hunter{Hunter: &proto.Hunter{
		Options: &proto.Hunter_Options{
			Ammo:           proto.Hunter_Options_ThoriumHeadedArrow,
			QuiverBonus:    proto.Hunter_Options_Speed15,
			PetType:        proto.Hunter_Options_Cat,
			PetAttackSpeed: proto.Hunter_Options_OneTwo,
			PetUptime:      fullPetUptime,
		},
	}}
}

func mageOptions(p *proto.Player) {
	p.Spec = &proto.Player_Mage{Mage: &proto.Mage{
		Options: &proto.Mage_Options{Armor: proto.Mage_Options_MoltenArmor},
	}}
}

func holyPaladinOptions(p *proto.Player) {
	p.Spec = &proto.Player_HolyPaladin{HolyPaladin: &proto.HolyPaladin{
		Options: &proto.PaladinOptions{Aura: proto.PaladinAura_ConcentrationAura},
	}}
}

func healingPriestOptions(p *proto.Player) {
	p.Spec = &proto.Player_HealingPriest{HealingPriest: &proto.HealingPriest{
		// The reference is self: a healing priest's Power Infusion has no
		// one else in a one-healer sim to go to.
		Options: &proto.HealingPriest_Options{UseInnerFire: true, PowerInfusionTarget: &proto.UnitReference{}},
	}}
}

func retributionPaladinOptions(p *proto.Player) {
	p.Spec = &proto.Player_RetributionPaladin{RetributionPaladin: &proto.RetributionPaladin{
		Options: &proto.PaladinOptions{
			PrimarySeal: proto.PaladinSeal_Righteousness,
			Aura:        proto.PaladinAura_SanctityAura,
		},
	}}
}

func shadowPriestOptions(p *proto.Player) {
	p.Spec = &proto.Player_ShadowPriest{ShadowPriest: &proto.ShadowPriest{
		Options: &proto.ShadowPriest_Options{},
	}}
}

func rogueOptions(p *proto.Player) {
	p.Spec = &proto.Player_Rogue{Rogue: &proto.Rogue{Options: &proto.RogueOptions{}}}
}

func elementalShamanOptions(p *proto.Player) {
	p.Spec = &proto.Player_ElementalShaman{ElementalShaman: &proto.ElementalShaman{
		Options: &proto.ElementalShaman_Options{},
	}}
}

func enhancementShamanOptions(p *proto.Player) {
	p.Spec = &proto.Player_EnhancementShaman{EnhancementShaman: &proto.EnhancementShaman{
		Options: &proto.EnhancementShaman_Options{SyncType: proto.ShamanSyncType_Auto},
	}}
}

func restorationShamanOptions(p *proto.Player) {
	p.Spec = &proto.Player_RestorationShaman{RestorationShaman: &proto.RestorationShaman{
		Options: &proto.RestorationShaman_Options{},
	}}
}

func warlockOptions(p *proto.Player) {
	p.Spec = &proto.Player_Warlock{Warlock: &proto.Warlock{
		Options: &proto.WarlockOptions{
			Armor:  proto.WarlockOptions_DemonArmor,
			Summon: proto.WarlockOptions_Succubus,
		},
	}}
}

func warriorOptions(p *proto.Player) {
	p.Spec = &proto.Player_Warrior{Warrior: &proto.Warrior{
		Options: &proto.Warrior_Options{
			StartingRage: defaultStartingRage,
			Shout:        proto.WarriorShout_WarriorShoutBattle,
		},
	}}
}

// applySpec attaches the spec's options and its default rotation,
// rewritten to the character's own rank of every ranked spell it
// casts. The engine cannot build an agent for a player carrying
// neither.
func applySpec(player *proto.Player, slug, class string, level int) error {
	apply, ok := specOptions[slug]
	if !ok {
		return fmt.Errorf("%w: %q; the specs this build carries are %v", ErrUnsupportedSpec, slug, supportedSpecs())
	}
	rot, err := rotation(slug, class, level)
	if err != nil {
		return err
	}
	apply(player)
	player.Rotation = rot
	return nil
}

// rotation parses one embedded APL, rewritten to level's ranks (see
// rewriteRotationRanks). It is parsed per build rather than cached, so
// no two requests ever share a mutable rotation.
func rotation(name, class string, level int) (*proto.APLRotation, error) {
	b, err := aplFS.ReadFile("apl/" + name + ".apl.json")
	if err != nil {
		return nil, fmt.Errorf("request: no embedded APL for %q: %w", name, err)
	}
	// At MaxLevel the rotation was authored for exactly the ranks it
	// already carries, so skipping the rewrite is both an optimization
	// and the guarantee that a level-MaxLevel rotation is byte-for-byte
	// what it always was.
	if level < api.MaxLevel {
		b, err = rewriteRotationRanks(b, class, level)
		if err != nil {
			return nil, fmt.Errorf("request: rewriting %q's rotation for level %d: %w", name, level, err)
		}
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
