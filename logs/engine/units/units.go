// logs/engine/units/units.go
// Package units parses GUIDs and keeps the registry of everyone seen in a
// log: who is a player, which pet belongs to whom, and what class a player
// is. Class comes from COMBATANT_INFO when the log has it and is otherwise
// inferred from the spells cast, always labelled with its source.
package units

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
)

// Unit flag bits, as the game writes them in sourceFlags and destFlags.
const (
	FlagAffiliationMine     uint32 = 0x00000001
	FlagAffiliationParty    uint32 = 0x00000002
	FlagAffiliationRaid     uint32 = 0x00000004
	FlagAffiliationOutsider uint32 = 0x00000008
	FlagReactionFriendly    uint32 = 0x00000010
	FlagReactionNeutral     uint32 = 0x00000020
	FlagReactionHostile     uint32 = 0x00000040
	FlagControlPlayer       uint32 = 0x00000100
	FlagControlNPC          uint32 = 0x00000200
	FlagTypePlayer          uint32 = 0x00000400
	FlagTypeNPC             uint32 = 0x00000800
	FlagTypePet             uint32 = 0x00001000
	FlagTypeGuardian        uint32 = 0x00002000
	FlagTypeObject          uint32 = 0x00004000
)

// Hostile reports whether the flags mark a hostile unit.
func Hostile(f uint32) bool { return f&FlagReactionHostile != 0 }

// Friendly reports whether the flags mark a friendly unit.
func Friendly(f uint32) bool { return f&FlagReactionFriendly != 0 }

// Enemy reports whether the flags describe a unit on the other side of the
// fight: one the log gave a reaction (neutral or hostile) that is not friendly.
// Flags with no reaction bit at all -- the environment, a unit the log never
// described -- are nobody's enemy, so a fall does not become a boss.
func Enemy(f uint32) bool {
	const reaction = FlagReactionFriendly | FlagReactionNeutral | FlagReactionHostile
	return f&reaction != 0 && !Friendly(f)
}

// SameSide reports whether two flag sets put their units on one side of the
// fight: both friendly or both hostile. Flags with no reaction bit at all
// (a unit the log never described) are nobody's side, so they never match.
func SameSide(a, b uint32) bool {
	const reaction = FlagReactionFriendly | FlagReactionHostile
	if a&reaction == 0 || b&reaction == 0 {
		return false
	}
	return Friendly(a) == Friendly(b) && Hostile(a) == Hostile(b)
}

// Kind is what a GUID refers to.
type Kind string

// The kinds. The strings are stable: they appear in report.json.
const (
	KindNone       Kind = "none"
	KindPlayer     Kind = "player"
	KindCreature   Kind = "creature"
	KindPet        Kind = "pet"
	KindVehicle    Kind = "vehicle"
	KindGameObject Kind = "object"
	KindUnknown    Kind = "unknown"
)

// NoGUID is the placeholder the game writes when there is no unit.
const NoGUID = "0000000000000000"

// GUID is a parsed unit identifier.
type GUID struct {
	Raw      string
	Kind     Kind
	ServerID int64  // realm id for a player, server id for an NPC
	NPCID    int64  // creature entry for an NPC, pet, vehicle, or object
	SpawnUID string // the per-spawn suffix
}

// Parse reads a GUID. The two shapes are Player-<realm>-<hex uid> and
// <Kind>-<subtype>-<server>-<instance>-<zone>-<npcID>-<spawnUID>.
func Parse(s string) GUID {
	g := GUID{Raw: s}
	switch {
	case s == "" || strings.Trim(s, "0") == "":
		g.Kind = KindNone
		return g
	case strings.HasPrefix(s, "Player-"):
		g.Kind = KindPlayer
	case strings.HasPrefix(s, "Creature-"):
		g.Kind = KindCreature
	case strings.HasPrefix(s, "Pet-"):
		g.Kind = KindPet
	case strings.HasPrefix(s, "Vehicle-"):
		g.Kind = KindVehicle
	case strings.HasPrefix(s, "GameObject-"):
		g.Kind = KindGameObject
	default:
		g.Kind = KindUnknown
		return g
	}
	p := strings.Split(s, "-")
	if g.Kind == KindPlayer {
		if len(p) >= 3 {
			g.ServerID, _ = strconv.ParseInt(p[1], 10, 64)
			g.SpawnUID = p[2]
		}
		return g
	}
	if len(p) >= 7 {
		g.ServerID, _ = strconv.ParseInt(p[2], 10, 64)
		g.NPCID, _ = strconv.ParseInt(p[5], 10, 64)
		g.SpawnUID = p[6]
	}
	return g
}

// Unit is everything the registry knows about one GUID.
type Unit struct {
	GUID        string    `json:"guid"`
	Name        string    `json:"name"`
	Kind        Kind      `json:"kind"`
	NPCID       int64     `json:"npc_id,omitempty"`
	Flags       uint32    `json:"flags"`
	OwnerGUID   string    `json:"owner_guid,omitempty"`
	Class       string    `json:"class,omitempty"`
	ClassSource string    `json:"class_source,omitempty"` // "combatant_info" or "inferred"
	SpecID      int64     `json:"spec_id,omitempty"`
	ItemLevel   int64     `json:"item_level,omitempty"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
}

// IsPlayer reports whether the unit is a player character.
func (u Unit) IsPlayer() bool { return u.Kind == KindPlayer }

// Options injects the lookup tables. Both are optional: with neither, class
// stays empty rather than being guessed.
type Options struct {
	// ClassBySpell maps a spell id to a class name. Build it from the
	// pipeline's talent data, which lists every talent rank's spell id per
	// class, so no spell id is ever hard-coded here.
	ClassBySpell map[int64]string
	// ClassBySpec maps a COMBATANT_INFO spec id to a class name.
	ClassBySpec map[int64]string
}

// Registry accumulates units as events arrive.
type Registry struct {
	opt   Options
	units map[string]*Unit
}

// State is the registry in a form that survives a restart.
type State struct {
	Units []Unit `json:"units"`
}

// NewRegistry returns an empty registry.
func NewRegistry(o Options) *Registry {
	return &Registry{opt: o, units: map[string]*Unit{}}
}

// RestoreRegistry rebuilds a registry from saved state.
func RestoreRegistry(o Options, s State) *Registry {
	r := NewRegistry(o)
	for i := range s.Units {
		u := s.Units[i]
		r.units[u.GUID] = &u
	}
	return r
}

// State captures the registry for serialisation, sorted so the bytes are
// reproducible.
func (r *Registry) State() State {
	s := State{Units: r.All()}
	return s
}

// Observe folds one event into the registry.
func (r *Registry) Observe(e event.Event) {
	r.see(e.Source, e.Time)
	r.see(e.Dest, e.Time)
	r.see(e.ExtraUnit, e.Time)

	// The advanced block names a pet's owner directly. It names an owner for a
	// boss mechanic's hostile add too (the player the mechanic targeted), so
	// the same test as a summon's applies: a pet or a guardian, or a unit on
	// the owner's own side.
	if e.Adv.OK && e.Adv.InfoGUID != "" && e.Adv.OwnerGUID != "" &&
		Parse(e.Adv.OwnerGUID).Kind != KindNone {
		if u := r.units[e.Adv.InfoGUID]; u != nil && r.ownable(u, e.Adv.OwnerGUID) {
			u.OwnerGUID = e.Adv.OwnerGUID
		}
	}
	// A summon is the authoritative ownership signal, for a pet or a guardian, or
	// for anything on the summoner's own side. A boss mechanic writes SPELL_SUMMON
	// with the player it targets as the source of a hostile add; that add is not
	// theirs, and its damage to the raid is not their damage.
	if e.Kind == event.Summon && e.Dest.GUID != "" && e.Source.GUID != "" {
		if u := r.units[e.Dest.GUID]; u != nil && r.ownable(u, e.Source.GUID) {
			u.OwnerGUID = e.Source.GUID
		}
	}
	// COMBATANT_INFO wins over inference.
	if e.Kind == event.CombatantInfo && e.Combatant != nil {
		u := r.units[e.Combatant.GUID]
		if u == nil {
			u = r.touch(e.Combatant.GUID, "", 0, e.Time)
		}
		u.SpecID = e.Combatant.SpecID
		u.ItemLevel = e.Combatant.ItemLevel
		if class, ok := r.opt.ClassBySpec[e.Combatant.SpecID]; ok {
			u.Class, u.ClassSource = class, "combatant_info"
		}
		return
	}
	// Otherwise infer from the spell cast, but never overwrite a class that
	// came from COMBATANT_INFO.
	if e.Spell.ID != 0 && e.Source.GUID != "" {
		if class, ok := r.opt.ClassBySpell[e.Spell.ID]; ok {
			if u := r.units[e.Source.GUID]; u != nil && u.IsPlayer() && u.ClassSource != "combatant_info" {
				u.Class, u.ClassSource = class, "inferred"
			}
		}
	}
}

// ownable reports whether a unit can belong to the named owner: a pet or a
// guardian always can; anything else only when it is on the owner's side,
// which a hostile add spawned on a player is not.
func (r *Registry) ownable(u *Unit, ownerGUID string) bool {
	if u.Flags&(FlagTypePet|FlagTypeGuardian) != 0 {
		return true
	}
	owner, ok := r.units[ownerGUID]
	if !ok {
		return false
	}
	return SameSide(u.Flags, owner.Flags)
}

func (r *Registry) see(u event.Unit, at time.Time) {
	if u.GUID == "" || u.GUID == NoGUID {
		return
	}
	r.touch(u.GUID, u.Name, u.Flags, at)
}

func (r *Registry) touch(guid, name string, flags uint32, at time.Time) *Unit {
	u, ok := r.units[guid]
	if !ok {
		g := Parse(guid)
		u = &Unit{GUID: guid, Kind: g.Kind, NPCID: g.NPCID, FirstSeen: at}
		r.units[guid] = u
	}
	if name != "" {
		u.Name = name
	}
	if flags != 0 {
		u.Flags = flags
	}
	if at.After(u.LastSeen) {
		u.LastSeen = at
	}
	if u.FirstSeen.IsZero() || (!at.IsZero() && at.Before(u.FirstSeen)) {
		u.FirstSeen = at
	}
	return u
}

// Get returns one unit.
func (r *Registry) Get(guid string) (Unit, bool) {
	u, ok := r.units[guid]
	if !ok {
		return Unit{}, false
	}
	return *u, true
}

// Name returns a unit's display name, falling back to the GUID.
func (r *Registry) Name(guid string) string {
	if u, ok := r.units[guid]; ok && u.Name != "" {
		return u.Name
	}
	return guid
}

// Owner walks the pet chain to the player who owns a unit, and returns the
// GUID itself when it owns nothing. A cycle stops the walk rather than
// hanging.
func (r *Registry) Owner(guid string) string {
	seen := map[string]bool{}
	for range 8 {
		if seen[guid] {
			return guid
		}
		seen[guid] = true
		u, ok := r.units[guid]
		if !ok || u.OwnerGUID == "" || u.OwnerGUID == guid {
			return guid
		}
		guid = u.OwnerGUID
	}
	return guid
}

// Players returns every player seen, sorted by GUID.
func (r *Registry) Players() []Unit {
	var out []Unit
	for _, u := range r.units {
		if u.IsPlayer() {
			out = append(out, *u)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GUID < out[j].GUID })
	return out
}

// All returns every unit seen, sorted by GUID.
func (r *Registry) All() []Unit {
	out := make([]Unit, 0, len(r.units))
	for _, u := range r.units {
		out = append(out, *u)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].GUID < out[j].GUID })
	return out
}
