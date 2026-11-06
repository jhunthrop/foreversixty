// api/internal/fs1/drops.go
//
// The export's drops= section (guild slice 7), written by the addon's RaidDrops.lua through
// Codec.encodeDropCorpse: every boss corpse the character's loot window read in a raid, `;`-joined:
//
//	<unix>:<instance_id>:<encounter_id>:<method>:<source>:<item>,<item>,...
//
// <encounter_id> is 0 when the addon saw no kill, <method> the loot method's name ("" unknown),
// <source> the corpse GUID ("" where the client has none) and each <item> `<item_id>x<count>`,
// then `@` and the url-encoded winners `+`-joined when the loot chat named any. Read leniently, the
// package rule: a malformed corpse or item is dropped, never the export.
package fs1

import "strings"

// Drop bounds: past these the rest of the section is not read (the addon writes far fewer).
const (
	maxDropCorpses = 64
	maxDropItems   = 16
	maxDropWinners = 40
	maxDropCount   = 40
	maxWinnerLen   = 48
)

// DropItem is one item a corpse held: its id, how many and who the loot chat says received them.
type DropItem struct {
	ItemID  int
	Count   int
	Winners []string
}

// DropCorpse is one boss corpse as the addon recorded it.
type DropCorpse struct {
	// At is the wall clock (unix seconds) the loot window opened.
	At          int64
	InstanceID  int
	EncounterID int64
	Method      string
	Source      string
	Items       []DropItem
}

// parseDrops reads drops='s corpses, skipping a malformed one.
func parseDrops(field string) []DropCorpse {
	out := []DropCorpse{}
	for _, entry := range strings.Split(field, ";") {
		if len(out) == maxDropCorpses {
			break
		}
		if c, ok := parseDropCorpse(entry); ok {
			out = append(out, c)
		}
	}
	return out
}

func parseDropCorpse(entry string) (DropCorpse, bool) {
	parts := strings.SplitN(entry, ":", 6)
	if len(parts) != 6 || !lowerLetters(parts[3]) || !guidChars(parts[4]) {
		return DropCorpse{}, false
	}
	at, atOK := nonNegative(parts[0])
	instance, instanceOK := nonNegative(parts[1])
	encounter, encounterOK := nonNegative(parts[2])
	if !atOK || !instanceOK || !encounterOK || at == 0 {
		return DropCorpse{}, false
	}
	c := DropCorpse{At: int64(at), InstanceID: instance, EncounterID: int64(encounter), Method: parts[3],
		Source: parts[4], Items: []DropItem{}}
	for _, text := range nonEmpty(parts[5]) {
		if len(c.Items) == maxDropItems {
			break
		}
		if it, ok := parseDropItem(text); ok {
			c.Items = append(c.Items, it)
		}
	}
	return c, len(c.Items) > 0
}

func parseDropItem(text string) (DropItem, bool) {
	head, winners, _ := strings.Cut(text, "@")
	idText, countText, found := strings.Cut(head, "x")
	id, idOK := nonNegative(idText)
	count, countOK := nonNegative(countText)
	if !found || !idOK || !countOK || id == 0 || count == 0 || count > maxDropCount {
		return DropItem{}, false
	}
	it := DropItem{ItemID: id, Count: count, Winners: []string{}}
	for _, name := range strings.Split(winners, "+") {
		name = strings.TrimSpace(urlDecode(name))
		if name != "" && len(name) <= maxWinnerLen && len(it.Winners) < maxDropWinners {
			it.Winners = append(it.Winners, name)
		}
	}
	return it, true
}

func lowerLetters(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}

func guidChars(s string) bool {
	for _, r := range s {
		if !(r == '-' || (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return len(s) <= 96
}
