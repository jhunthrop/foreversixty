// api/internal/fs1/prep.go
//
// The export's raid-prep sections (owner decision D14), written by the addon's
// Readiness.lua through Codec.encodeFS1:
//
//	school=<fire>:<nature>:<frost>:<shadow>:<arcane>   resistance totals (UnitResistance)
//	supplies=<item_id>:<count>,...                     every Consumable-class stack in the
//	                                                   carried bags, counts summed per item
//	locks=<item_id>,...                                the keys carried: the keyring plus
//	                                                   any Key-class item in the bags
//	quests=<quest_id>:<0|1>,...                        each curated attunement quest asked
//	                                                   about, 1 when the client flags it done
//
// None starts with a letter the client's font strings read as an escape (c, r, n, h, t, a,
// k: see Codec.encodeFS1's "who" note). Each section is written only when the client could
// answer it, and written even when empty (supplies= with nothing in the bags), so presence
// says "read" and absence says "an addon too old to read it": HasSupplies, HasKeyring and
// HasQuests carry that distinction. A malformed section, or a malformed entry, is dropped
// rather than failing the decode (the package rule).
package fs1

import (
	"strconv"
	"strings"
)

// Resistances is the school= section: the character's resistance totals at export time.
type Resistances struct {
	Fire   int `json:"fire"`
	Nature int `json:"nature"`
	Frost  int `json:"frost"`
	Shadow int `json:"shadow"`
	Arcane int `json:"arcane"`
}

// resistanceSchools is how many totals school= carries, in Resistances' field order.
const resistanceSchools = 5

// parseResistances reads school='s five non-negative totals; ok is false for any other shape.
func parseResistances(field string) (Resistances, bool) {
	parts := strings.Split(field, ":")
	if len(parts) != resistanceSchools {
		return Resistances{}, false
	}
	values := make([]int, resistanceSchools)
	for i, part := range parts {
		n, ok := nonNegative(part)
		if !ok {
			return Resistances{}, false
		}
		values[i] = n
	}
	return Resistances{Fire: values[0], Nature: values[1], Frost: values[2], Shadow: values[3], Arcane: values[4]}, true
}

// parseSupplies reads supplies='s `item_id:count` entries into item id -> count, summing a
// repeated id and skipping an entry that is not two non-negative integers.
func parseSupplies(field string) map[int]int {
	out := map[int]int{}
	for _, entry := range nonEmpty(field) {
		idText, countText, found := strings.Cut(entry, ":")
		id, idOK := nonNegative(idText)
		count, countOK := nonNegative(countText)
		if !found || !idOK || !countOK || id == 0 {
			continue
		}
		out[id] += count
	}
	return out
}

// parseIDList reads locks='s comma-joined item ids, skipping a malformed entry.
func parseIDList(field string) []int {
	out := []int{}
	for _, entry := range nonEmpty(field) {
		if id, ok := nonNegative(entry); ok && id > 0 {
			out = append(out, id)
		}
	}
	return out
}

// parseQuests reads quests='s `quest_id:0|1` entries into quest id -> done, skipping a
// malformed entry.
func parseQuests(field string) map[int]bool {
	out := map[int]bool{}
	for _, entry := range nonEmpty(field) {
		idText, flag, found := strings.Cut(entry, ":")
		id, ok := nonNegative(idText)
		if !found || !ok || id == 0 || (flag != "0" && flag != "1") {
			continue
		}
		out[id] = flag == "1"
	}
	return out
}

// nonEmpty is field's comma-separated entries without the empty ones.
func nonEmpty(field string) []string {
	out := []string{}
	for _, entry := range strings.Split(field, ",") {
		if entry != "" {
			out = append(out, entry)
		}
	}
	return out
}

// nonNegative reads a digits-only integer.
func nonNegative(text string) (int, bool) {
	if text == "" {
		return 0, false
	}
	for _, r := range text {
		if r < '0' || r > '9' {
			return 0, false
		}
	}
	n, err := strconv.Atoi(text)
	return n, err == nil
}

// decodePrep records one raid-prep section on d; any other section name is ignored.
func (d *Decoded) decodePrep(name, field string) {
	switch name {
	case "school":
		if r, ok := parseResistances(field); ok {
			d.Resistances = &r
		}
	case "supplies":
		d.Supplies, d.HasSupplies = parseSupplies(field), true
	case "locks":
		d.Keyring, d.HasKeyring = parseIDList(field), true
	case "quests":
		d.Quests, d.HasQuests = parseQuests(field), true
	}
}
