// logs/engine/parquet/schema.go
// Package parquet writes a fight's events as one Parquet file and reads it
// back. The schema is a direct projection of event.Event: one nullable
// column per suffix field, the seventeen advanced fields as their own
// nullable columns, and raw for anything the decoder could not type.
package parquet

import (
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// CreatedBy is pinned so two runs over the same input produce byte-identical
// files. Never put a timestamp or a host name in it.
const (
	CreatedBy      = "forever-logs"
	CreatedVersion = "1"
	CreatedBuild   = "schema-1"
)

// Row is one event. Field order here is the column order in the file.
type Row struct {
	TimeUnixNano int64  `parquet:"time_unix_nano,delta"`
	Line         int64  `parquet:"line,delta"`
	Offset       int64  `parquet:"offset,delta"`
	Event        string `parquet:"event,dict"`
	Kind         string `parquet:"kind,dict"`

	SourceGUID  string `parquet:"source_guid,dict"`
	SourceName  string `parquet:"source_name,dict"`
	SourceFlags int64  `parquet:"source_flags"`
	SourceRaid  int64  `parquet:"source_raid"`
	DestGUID    string `parquet:"dest_guid,dict"`
	DestName    string `parquet:"dest_name,dict"`
	DestFlags   int64  `parquet:"dest_flags"`
	DestRaid    int64  `parquet:"dest_raid"`

	SpellID     int64  `parquet:"spell_id"`
	SpellName   string `parquet:"spell_name,dict"`
	SpellSchool int64  `parquet:"spell_school"`

	ExtraGUID        string `parquet:"extra_guid,dict"`
	ExtraName        string `parquet:"extra_name,dict"`
	ExtraSpellID     int64  `parquet:"extra_spell_id"`
	ExtraSpellName   string `parquet:"extra_spell_name,dict"`
	ExtraSpellSchool int64  `parquet:"extra_spell_school"`

	Amount       *int64 `parquet:"amount,optional"`
	BaseAmount   *int64 `parquet:"base_amount,optional"`
	Overkill     *int64 `parquet:"overkill,optional"`
	School       *int64 `parquet:"school,optional"`
	Resisted     *int64 `parquet:"resisted,optional"`
	Blocked      *int64 `parquet:"blocked,optional"`
	Absorbed     *int64 `parquet:"absorbed,optional"`
	Overheal     *int64 `parquet:"overheal,optional"`
	Total        *int64 `parquet:"total,optional"`
	Stacks       *int64 `parquet:"stacks,optional"`
	PowerType    *int64 `parquet:"power_type,optional"`
	MaxPower     *int64 `parquet:"max_power,optional"`
	OverEnergize *int64 `parquet:"over_energize,optional"`
	ItemID       *int64 `parquet:"item_id,optional"`

	Critical *bool `parquet:"critical,optional"`
	Glancing *bool `parquet:"glancing,optional"`
	Crushing *bool `parquet:"crushing,optional"`
	OffHand  *bool `parquet:"off_hand,optional"`

	MissType   string `parquet:"miss_type,dict"`
	AuraType   string `parquet:"aura_type,dict"`
	FailedType string `parquet:"failed_type,dict"`
	EnvType    string `parquet:"env_type,dict"`
	ItemName   string `parquet:"item_name,dict"`

	// The seventeen advanced-logging fields.
	AdvPresent      bool     `parquet:"adv_present"`
	AdvInfoGUID     string   `parquet:"adv_info_guid,dict"`
	AdvOwnerGUID    string   `parquet:"adv_owner_guid,dict"`
	AdvCurrentHP    *int64   `parquet:"adv_current_hp,optional"`
	AdvMaxHP        *int64   `parquet:"adv_max_hp,optional"`
	AdvAttackPower  *int64   `parquet:"adv_attack_power,optional"`
	AdvSpellPower   *int64   `parquet:"adv_spell_power,optional"`
	AdvArmor        *int64   `parquet:"adv_armor,optional"`
	AdvAbsorb       *int64   `parquet:"adv_absorb,optional"`
	AdvPowerType    *int64   `parquet:"adv_power_type,optional"`
	AdvCurrentPower *int64   `parquet:"adv_current_power,optional"`
	AdvMaxPower     *int64   `parquet:"adv_max_power,optional"`
	AdvPowerCost    *int64   `parquet:"adv_power_cost,optional"`
	AdvPositionX    *float64 `parquet:"adv_position_x,optional"`
	AdvPositionY    *float64 `parquet:"adv_position_y,optional"`
	AdvUIMapID      *int64   `parquet:"adv_ui_map_id,optional"`
	AdvFacing       *float64 `parquet:"adv_facing,optional"`
	AdvLevel        *int64   `parquet:"adv_level,optional"`

	Raw   string `parquet:"raw"`
	Error string `parquet:"error"`
}

func optI(v event.OptInt) *int64 {
	if !v.OK {
		return nil
	}
	n := v.V
	return &n
}

func optB(v event.OptBool) *bool {
	if !v.OK {
		return nil
	}
	b := v.V
	return &b
}

func fromI(p *int64) event.OptInt {
	if p == nil {
		return event.OptInt{}
	}
	return event.OptInt{V: *p, OK: true}
}

func fromB(p *bool) event.OptBool {
	if p == nil {
		return event.OptBool{}
	}
	return event.OptBool{V: *p, OK: true}
}

func i64(v int64) *int64     { return &v }
func f64(v float64) *float64 { return &v }
func deref(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}
func derefF(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

// RowOf projects an event onto the schema.
func RowOf(e event.Event) Row {
	r := Row{
		TimeUnixNano: e.Time.UnixNano(),
		Line:         e.Line,
		Offset:       e.Offset,
		Event:        e.Name,
		Kind:         e.Kind.String(),

		SourceGUID: e.Source.GUID, SourceName: e.Source.Name,
		SourceFlags: int64(e.Source.Flags), SourceRaid: int64(e.Source.Raid),
		DestGUID: e.Dest.GUID, DestName: e.Dest.Name,
		DestFlags: int64(e.Dest.Flags), DestRaid: int64(e.Dest.Raid),

		SpellID: e.Spell.ID, SpellName: e.Spell.Name, SpellSchool: e.Spell.School,

		ExtraGUID: e.ExtraUnit.GUID, ExtraName: e.ExtraUnit.Name,
		ExtraSpellID: e.ExtraSpell.ID, ExtraSpellName: e.ExtraSpell.Name,
		ExtraSpellSchool: e.ExtraSpell.School,

		Amount: optI(e.Amount), BaseAmount: optI(e.BaseAmount), Overkill: optI(e.Overkill),
		School: optI(e.School), Resisted: optI(e.Resisted), Blocked: optI(e.Blocked),
		Absorbed: optI(e.Absorbed), Overheal: optI(e.Overheal), Total: optI(e.Total),
		Stacks: optI(e.Stacks), PowerType: optI(e.PowerType), MaxPower: optI(e.MaxPower),
		OverEnergize: optI(e.OverEnergize), ItemID: optI(e.ItemID),

		Critical: optB(e.Critical), Glancing: optB(e.Glancing),
		Crushing: optB(e.Crushing), OffHand: optB(e.OffHand),

		MissType: e.MissType, AuraType: e.AuraType, FailedType: e.FailedType,
		EnvType: e.EnvType, ItemName: e.ItemName,

		Raw: e.Raw, Error: e.Error,
	}
	if e.Adv.OK {
		r.AdvPresent = true
		r.AdvInfoGUID, r.AdvOwnerGUID = e.Adv.InfoGUID, e.Adv.OwnerGUID
		r.AdvCurrentHP, r.AdvMaxHP = i64(e.Adv.CurrentHP), i64(e.Adv.MaxHP)
		r.AdvAttackPower, r.AdvSpellPower = i64(e.Adv.AttackPower), i64(e.Adv.SpellPower)
		r.AdvArmor, r.AdvAbsorb = i64(e.Adv.Armor), i64(e.Adv.Absorb)
		r.AdvPowerType, r.AdvCurrentPower = i64(e.Adv.PowerType), i64(e.Adv.CurrentPower)
		r.AdvMaxPower, r.AdvPowerCost = i64(e.Adv.MaxPower), i64(e.Adv.PowerCost)
		r.AdvPositionX, r.AdvPositionY = f64(e.Adv.PositionX), f64(e.Adv.PositionY)
		r.AdvUIMapID, r.AdvFacing = i64(e.Adv.UIMapID), f64(e.Adv.Facing)
		r.AdvLevel = i64(e.Adv.Level)
	}
	return r
}

// EventOf reads a row back. The specials that carry their own structs
// (Encounter, Zone, Combatant) are not in the events file: they live in
// report.json and summary.json, which is where the report page reads them.
func EventOf(r Row) event.Event {
	e := event.Event{
		Time:   time.Unix(0, r.TimeUnixNano).UTC(),
		Line:   r.Line,
		Offset: r.Offset,
		Name:   r.Event,
		Kind:   kindByName(r.Kind),

		Source: event.Unit{GUID: r.SourceGUID, Name: r.SourceName,
			Flags: uint32(r.SourceFlags), Raid: uint32(r.SourceRaid)},
		Dest: event.Unit{GUID: r.DestGUID, Name: r.DestName,
			Flags: uint32(r.DestFlags), Raid: uint32(r.DestRaid)},
		Spell: event.Spell{ID: r.SpellID, Name: r.SpellName, School: r.SpellSchool},

		ExtraUnit:  event.Unit{GUID: r.ExtraGUID, Name: r.ExtraName},
		ExtraSpell: event.Spell{ID: r.ExtraSpellID, Name: r.ExtraSpellName, School: r.ExtraSpellSchool},

		Amount: fromI(r.Amount), BaseAmount: fromI(r.BaseAmount), Overkill: fromI(r.Overkill),
		School: fromI(r.School), Resisted: fromI(r.Resisted), Blocked: fromI(r.Blocked),
		Absorbed: fromI(r.Absorbed), Overheal: fromI(r.Overheal), Total: fromI(r.Total),
		Stacks: fromI(r.Stacks), PowerType: fromI(r.PowerType), MaxPower: fromI(r.MaxPower),
		OverEnergize: fromI(r.OverEnergize), ItemID: fromI(r.ItemID),

		Critical: fromB(r.Critical), Glancing: fromB(r.Glancing),
		Crushing: fromB(r.Crushing), OffHand: fromB(r.OffHand),

		MissType: r.MissType, AuraType: r.AuraType, FailedType: r.FailedType,
		EnvType: r.EnvType, ItemName: r.ItemName,

		Raw: r.Raw, Error: r.Error,
	}
	if r.AdvPresent {
		e.Adv = event.Advanced{
			OK: true, InfoGUID: r.AdvInfoGUID, OwnerGUID: r.AdvOwnerGUID,
			CurrentHP: deref(r.AdvCurrentHP), MaxHP: deref(r.AdvMaxHP),
			AttackPower: deref(r.AdvAttackPower), SpellPower: deref(r.AdvSpellPower),
			Armor: deref(r.AdvArmor), Absorb: deref(r.AdvAbsorb),
			PowerType: deref(r.AdvPowerType), CurrentPower: deref(r.AdvCurrentPower),
			MaxPower: deref(r.AdvMaxPower), PowerCost: deref(r.AdvPowerCost),
			PositionX: derefF(r.AdvPositionX), PositionY: derefF(r.AdvPositionY),
			UIMapID: deref(r.AdvUIMapID), Facing: derefF(r.AdvFacing),
			Level: deref(r.AdvLevel),
		}
	}
	return e
}

// kinds is the reverse of event.Kind.String, built once.
var kinds = func() map[string]event.Kind {
	m := map[string]event.Kind{}
	for k := event.Unknown; k <= event.Durability; k++ {
		m[k.String()] = k
	}
	return m
}()

func kindByName(s string) event.Kind {
	if k, ok := kinds[s]; ok {
		return k
	}
	return event.Unknown
}
