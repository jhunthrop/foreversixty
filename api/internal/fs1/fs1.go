// Package fs1 reads the handful of facts GET /v1/me shows about a character's build —
// gear, talent split, level, data build and who= realm — out of the FS1 export string
// addon_exports.export stores (addon/ForeverSixty/Codec.lua's encodeFS1,
// web/src/lib/planner/fs1.ts's decodeFS1). It is deliberately not a full port of either:
// addon.ParseFS1Class and addon.ParseFS1Guild already read their own narrower slices of
// the same string, this package reads a third, and none of the three validates the
// sections it does not read.
//
// Unlike the addon's and the planner's own decoders, a malformed *section* here is simply
// not recorded rather than refusing the whole string: GET /v1/me must never fail over an
// old or malformed export, so there is no "refuse the whole code" outcome past the head.
// Only a head that is not FS1 at all, or whose talent field cannot be read, makes Decode
// report ok=false — the caller then shows the rest of the character row with every build
// field this package would have filled simply omitted.
package fs1

import (
	"net/url"
	"strconv"
	"strings"
)

// trees is how many talent trees every class has, in the FS1 head's <t1>/<t2>/<t3> field.
const trees = 3

// Talents is the three-tree rank/point split the head's talent field carries, read in
// tree order (the client's own tab order, position 0 first).
type Talents struct {
	// Trees is each tree's rank string exactly as the export carries it ("503200000"),
	// one base-36 digit per talent in tier/column order with trailing zeros trimmed —
	// never re-encoded or reformatted.
	Trees []string
	// Points is Trees' digit sum per tree: the same number a tree's point total would
	// read as in game, computed straight from the digits rather than by resolving each
	// digit to a talent (which needs client talent data this package does not have).
	Points []int
}

// Decoded is what Decode could read out of one export string. Gear is never nil once ok
// is true; Level and the who= fields are the zero value when the string carries no such
// section, which HasLevel and HasWho distinguish from an honestly-absent fact.
type Decoded struct {
	DataBuild, ClassSlug, RaceSlug string
	Gear                           map[string]int
	Talents                        Talents
	Level                          int
	HasLevel                       bool
	CharacterName, Realm           string
	HasWho                         bool
}

// Decode reads an FS1 export string's head (data build, class, race, talents, gear) and
// its level= and who= sections. ok is false only when the string is not a readable FS1
// head at all: wrong prefix, fewer than six head fields, the talent field not exactly
// three trees, or a talent digit outside 0-9a-z. A bad level= or who= section is read as
// simply absent, the same as a string from an addon old enough never to have written one.
func Decode(export string) (Decoded, bool) {
	pipes := strings.Split(strings.TrimSpace(export), "|")
	parts := strings.Split(pipes[0], ":")
	if len(parts) < 6 || parts[0] != "FS1" {
		return Decoded{}, false
	}

	treeStrings := strings.Split(parts[4], "/")
	if len(treeStrings) != trees {
		return Decoded{}, false
	}
	points := make([]int, trees)
	for i, tree := range treeStrings {
		for _, digit := range tree {
			rank := fromBase36(digit)
			if rank < 0 {
				return Decoded{}, false
			}
			points[i] += rank
		}
	}

	d := Decoded{
		DataBuild: parts[1],
		ClassSlug: parts[2],
		RaceSlug:  parts[3],
		Gear:      parseGear(strings.Join(parts[5:], ":")),
		Talents:   Talents{Trees: treeStrings, Points: points},
	}

	for _, section := range pipes[1:] {
		name, field, _ := strings.Cut(section, "=")
		switch name {
		case "level":
			if level, ok := parseLevel(field); ok {
				d.Level, d.HasLevel = level, true
			}
		case "who":
			name, realm, _ := strings.Cut(field, ":")
			d.CharacterName, d.Realm, d.HasWho = urlDecode(name), urlDecode(realm), true
		}
	}
	return d, true
}

// parseGear reads a comma-joined `<slot>=item_id[:enchant[:suffix]]` list into slot name
// (verbatim, whatever the export carries) -> item id. An entry with no `=`, an empty slot
// name, or a non-numeric item id is skipped rather than failing the whole decode — the
// same "a malformed section is simply not recorded" rule the package comment describes.
func parseGear(field string) map[string]int {
	gear := map[string]int{}
	if field == "" {
		return gear
	}
	for _, entry := range strings.Split(field, ",") {
		slot, value, hasEquals := strings.Cut(entry, "=")
		if !hasEquals || slot == "" {
			continue
		}
		itemField, _, _ := strings.Cut(value, ":")
		itemID, err := strconv.Atoi(itemField)
		if err != nil || itemID < 0 {
			continue
		}
		gear[slot] = itemID
	}
	return gear
}

// parseLevel reads a level= field's digits-only, 1..60 grammar (Codec.lua's own decodeFS1
// check for the same section).
func parseLevel(field string) (int, bool) {
	if field == "" {
		return 0, false
	}
	for _, r := range field {
		if r < '0' || r > '9' {
			return 0, false
		}
	}
	level, err := strconv.Atoi(field)
	if err != nil || level < 1 || level > 60 {
		return 0, false
	}
	return level, true
}

// fromBase36 is a single base-36 digit's value, or -1 when the rune is not one — toBase36's
// inverse, accepting either case the way Codec.lua's fromBase36 does.
func fromBase36(r rune) int {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0')
	case r >= 'a' && r <= 'z':
		return int(r-'a') + 10
	case r >= 'A' && r <= 'Z':
		return int(r-'A') + 10
	default:
		return -1
	}
}

// urlDecode is the who= section's percent-decoding. A malformed escape reads as itself
// rather than failing the decode, the same leniency ParseFS1Guild and the addon's and
// planner's own decoders give a name field no one but a player ever typed.
func urlDecode(s string) string {
	decoded, err := url.PathUnescape(s)
	if err != nil {
		return s
	}
	return decoded
}
