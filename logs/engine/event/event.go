// logs/engine/event/event.go
// Package event turns lexed parameters into one flat typed struct. Every
// optional field is explicitly nullable, so the Parquet schema is a direct
// projection of Event and a missing field is never confused with a zero.
// Events the layout does not know keep their raw text; events that do not
// match their layout row become ParseError rather than being mis-read.
package event

import "time"

// Kind classifies an event once decoded.
type Kind uint8

// The kinds. Their string forms are stable: they appear in the Parquet file
// and in summary JSON.
const (
	Unknown Kind = iota
	ParseError
	Header
	Damage
	// DamageLanded is a SWING_DAMAGE_LANDED line: the same swing SWING_DAMAGE
	// already reported, written again at impact with the target's advanced
	// block. It is never added to a total, only read for the target's health.
	DamageLanded
	Heal
	Missed
	Absorbed
	HealAbsorbed
	Energize
	AuraApplied
	AuraRemoved
	AuraRefresh
	AuraBroken
	AuraDose
	CastStart
	CastSuccess
	CastFailed
	Interrupt
	Dispel
	Summon
	Create
	Resurrect
	Instakill
	ExtraAttacks
	Death
	PartyKill
	EncounterStart
	EncounterEnd
	ZoneChange
	MapChange
	CombatantInfo
	Enchant
	Emote
	ChallengeModeStart
	ChallengeModeEnd
	Durability
)

var kindNames = map[Kind]string{
	Unknown: "unknown", ParseError: "parse_error", Header: "header",
	Damage: "damage", DamageLanded: "damage_landed", Heal: "heal", Missed: "missed", Absorbed: "absorbed",
	HealAbsorbed: "heal_absorbed", Energize: "energize",
	AuraApplied: "aura_applied", AuraRemoved: "aura_removed",
	AuraRefresh: "aura_refresh", AuraBroken: "aura_broken", AuraDose: "aura_dose",
	CastStart: "cast_start", CastSuccess: "cast_success", CastFailed: "cast_failed",
	Interrupt: "interrupt", Dispel: "dispel", Summon: "summon", Create: "create",
	Resurrect: "resurrect", Instakill: "instakill", ExtraAttacks: "extra_attacks",
	Death: "death", PartyKill: "party_kill",
	EncounterStart: "encounter_start", EncounterEnd: "encounter_end",
	ZoneChange: "zone_change", MapChange: "map_change",
	CombatantInfo: "combatant_info", Enchant: "enchant", Emote: "emote",
	ChallengeModeStart: "challenge_mode_start", ChallengeModeEnd: "challenge_mode_end",
	Durability: "durability",
}

func (k Kind) String() string {
	if s, ok := kindNames[k]; ok {
		return s
	}
	return "unknown"
}

// OptInt is a nullable integer. OK is false when the log did not carry the
// field at all; a field written as "nil" is also not OK.
type OptInt struct {
	V  int64
	OK bool
}

// OptBool is a nullable boolean. A field written as "nil" is present and
// false; an absent field is not OK.
type OptBool struct {
	V  bool
	OK bool
}

// Unit is one side of an event.
type Unit struct {
	GUID  string
	Name  string
	Flags uint32
	Raid  uint32
}

// Spell is a spell reference: the prefix spell, or an extra spell such as
// the one an interrupt stopped.
type Spell struct {
	ID     int64
	Name   string
	School int64
}

// Advanced is the 17-field advanced-logging block. Level is the creature's
// level for an NPC and the player's item level for a player: one field,
// two meanings, exactly as the game writes it.
type Advanced struct {
	OK           bool
	InfoGUID     string
	OwnerGUID    string
	CurrentHP    int64
	MaxHP        int64
	AttackPower  int64
	SpellPower   int64
	Armor        int64
	Absorb       int64
	PowerType    int64
	CurrentPower int64
	MaxPower     int64
	PowerCost    int64
	PositionX    float64
	PositionY    float64
	UIMapID      int64
	Facing       float64
	Level        int64
}

// Encounter carries ENCOUNTER_START and ENCOUNTER_END.
type Encounter struct {
	ID         int64
	Name       string
	Difficulty int64
	Size       int64
	InstanceID int64
	Kill       bool
}

// Zone carries ZONE_CHANGE and MAP_CHANGE.
type Zone struct {
	ID                     int64
	Name                   string
	Difficulty             int64
	MinX, MaxX, MinY, MaxY float64
}

// Item is one equipped item from COMBATANT_INFO.
type Item struct {
	ID        int64
	ItemLevel int64
	Enchants  []int64
	BonusIDs  []int64
	Gems      []int64
}

// Aura is one entry of the auras-at-pull list in COMBATANT_INFO.
type Aura struct {
	SourceGUID string
	SpellID    int64
}

// Combatant is a decoded COMBATANT_INFO line.
type Combatant struct {
	GUID       string
	Faction    int64
	Stats      map[string]int64
	SpecID     int64
	Talents    []int64
	PvPTalents []int64
	Borrowed   string // kept raw: its shape is expansion-specific
	Gear       []Item
	Auras      []Aura
	ItemLevel  int64 // mean of the equipped items that have one
}

// Event is one decoded line. Every field a given event does not use is left
// at its zero value with OK false.
type Event struct {
	Time   time.Time
	Offset int64
	Line   int64
	Name   string
	Kind   Kind

	Source Unit
	Dest   Unit
	Spell  Spell

	ExtraUnit  Unit
	ExtraSpell Spell

	Adv Advanced

	Amount       OptInt
	BaseAmount   OptInt
	Overkill     OptInt
	School       OptInt
	Resisted     OptInt
	Blocked      OptInt
	Absorbed     OptInt
	Overheal     OptInt
	Total        OptInt
	Stacks       OptInt
	PowerType    OptInt
	MaxPower     OptInt
	OverEnergize OptInt
	ItemID       OptInt

	Critical OptBool
	Glancing OptBool
	Crushing OptBool
	OffHand  OptBool

	MissType   string
	AuraType   string
	FailedType string
	EnvType    string
	ItemName   string

	Encounter *Encounter
	Zone      *Zone
	Combatant *Combatant

	Raw   string // set for Unknown and ParseError
	Error string // set for ParseError
}

// Effective is the damage or healing that actually landed: damage minus
// nothing (amount is already post-mitigation), healing minus overheal.
func (e Event) Effective() int64 {
	switch e.Kind {
	case Damage:
		// Amount includes overkill, as the log writes it; the damage that
		// counts is what the target could still take. Overkill is -1 on a
		// hit that did not kill, hence the floor.
		return e.Amount.V - max(e.Overkill.V, 0)
	case Heal:
		return e.Amount.V - e.Overheal.V
	default:
		return 0
	}
}
