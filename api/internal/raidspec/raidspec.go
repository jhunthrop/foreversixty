// Package raidspec is the raid roster's view of the spec catalogue (data/curated/specs.json,
// through its generated copy sim/specs): each spec's key, class, name, role, the roster group it
// fights in, its application emoji's name and the icon the emoji is drawn from
// (spec 2026-10-28 §P.9.1, §P.9.4).
package raidspec

import (
	"strings"

	"github.com/jhunthrop/foreversixty/sim/specs"
)

// The roles a spec plays, as specs.json and guild_raid_signups spell them.
const (
	RoleTank   = "tank"
	RoleHealer = "healer"
	RoleDPS    = "dps"
)

// emojiPrefix starts every roster emoji's name: fs_{class}_{spec}.
const emojiPrefix = "fs_"

// meleeSpecs are the DPS specs that fight in melee (§P.9.4; Survival is melee in Forever).
var meleeSpecs = map[string]bool{
	"druid-feral": true, "hunter-survival": true, "paladin-retribution": true,
	"rogue-assassination": true, "rogue-combat": true, "rogue-subtlety": true,
	"shaman-enhancement": true, "warrior-arms": true, "warrior-fury": true,
}

// meleeClasses place a DPS line with no spec (§P.9.4).
var meleeClasses = map[string]bool{"paladin": true, "rogue": true, "warrior": true}

// Spec is one playable spec as the roster uses it.
type Spec struct {
	// Key is the specs.json key ("paladin-holy"), the value guild_characters.raid_spec holds.
	Key   string
	Class string // the class slug, "paladin"
	Name  string // the full catalogue name, "Holy"
	Role  string
	Icon  string // the icon file's base name, no extension
}

// Emoji is the spec's application emoji name: fs_{class}_{spec}, hyphens as underscores.
func (s Spec) Emoji() string { return emojiPrefix + strings.ReplaceAll(s.Key, "-", "_") }

// Melee reports whether a DPS player of this spec stands in melee.
func (s Spec) Melee() bool { return meleeSpecs[s.Key] }

// All is every spec in catalogue order (class, then talent tree).
func All() []Spec {
	out := make([]Spec, 0, len(specs.All))
	for _, s := range specs.All {
		out = append(out, of(s))
	}
	return out
}

func of(s specs.Spec) Spec {
	return Spec{Key: s.Spec, Class: s.ClassSlug, Name: s.Name, Role: s.Role, Icon: s.Icon}
}

// ByKey is the spec with this specs.json key.
func ByKey(key string) (Spec, bool) {
	for _, s := range specs.All {
		if s.Spec == key {
			return of(s), true
		}
	}
	return Spec{}, false
}

// ForClass is the class's specs in catalogue order; nil for an unknown class.
func ForClass(class string) []Spec {
	class = strings.ToLower(strings.TrimSpace(class))
	var out []Spec
	for _, s := range specs.All {
		if s.ClassSlug == class {
			out = append(out, of(s))
		}
	}
	return out
}

// Resolve is the catalogue spec a sign-up's class, spec name and role name: the name matched
// case-insensitively within the class, the role deciding between two specs of one name (a logged
// "Feral" tank is Feral Bear). ok is false for a spec the catalogue does not name.
func Resolve(class, name, role string) (Spec, bool) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Spec{}, false
	}
	if strings.EqualFold(class, "druid") && (strings.EqualFold(name, "feral") || strings.EqualFold(name, "feral combat")) {
		if role == RoleTank {
			return ByKey("druid-feral-bear")
		}
		return ByKey("druid-feral")
	}
	for _, s := range ForClass(class) {
		if SameName(s.Name, name) {
			return s, true
		}
	}
	return Spec{}, false
}

// SameName reports whether two spec names are one spec however they are written: case aside and
// with the spaces gone, so a log's "BeastMastery" is the catalogue's "Beast Mastery"
// (rankings/query.go compares board filters the same way).
func SameName(a, b string) bool {
	return strings.EqualFold(strings.ReplaceAll(a, " ", ""), strings.ReplaceAll(b, " ", ""))
}

// MeleeClass reports whether a DPS line with no spec goes to Melee, by class.
func MeleeClass(class string) bool { return meleeClasses[strings.ToLower(class)] }

// ClassOrder is the game's class order (classChoices); the roster sorts by it.
var ClassOrder = []string{"druid", "hunter", "mage", "paladin", "priest", "rogue", "shaman", "warlock", "warrior"}

// ClassRank is the class's place in ClassOrder; an unknown or empty class sorts after every class.
func ClassRank(class string) int {
	for i, c := range ClassOrder {
		if c == strings.ToLower(class) {
			return i
		}
	}
	return len(ClassOrder)
}
