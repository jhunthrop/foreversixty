package sheet

import (
	"regexp"
	"strings"
)

// The sheet's words, verbatim from the spec's copy table (rows sheet.*, tile.alt and the gain
// words the panel shares). Tests read each one back from the spec.
const (
	CopyBisKicker   = "Best in slot"                         // sheet.bis.kicker
	CopyItemLevel   = "Item level"                           // sheet.char.ilvl, drawn uppercase
	CopyEmpty       = "Empty"                                // sheet.empty
	CopyUpgrade     = "Upgrade available"                    // sheet.upgrade
	CopyNone        = "No pick"                              // sheet.none
	CopyMark        = "foreversixty.gg"                      // sheet.mark
	CopyNotModelled = "Gain not modelled yet"                // item.notModelled
	CopyAltOffHand  = "Off hand: none, two-handed main hand" // sheet.alt.offhand
	CopyTwoHand     = "Two-hand in main hand"                // paperdoll: the off hand under a two-hander

	altBisLead  = "Best in slot sheet, "
	altCharLead = "Character sheet, "
	altEmpty    = "empty"
	altNoPick   = "no pick"
	altGainTail = " available"
	metaILvl    = "item level "
)

// SlotOff is the off hand's slot key: under a two-handed main hand it is drawn empty.
const SlotOff = "off_hand"

// slotLabels are the 17 slot keys (the BiS data's bis.Slot.Slot) and their words, in the
// character frame's order: the /bis slot choices' names, used by the panel and the alt texts.
var slotLabels = []struct{ key, label string }{
	{"head", "Head"}, {"neck", "Neck"}, {"shoulder", "Shoulder"}, {"back", "Back"},
	{"chest", "Chest"}, {"wrist", "Wrist"}, {"hands", "Hands"}, {"waist", "Waist"},
	{"legs", "Legs"}, {"feet", "Feet"}, {"finger1", "Finger 1"}, {"finger2", "Finger 2"},
	{"trinket1", "Trinket 1"}, {"trinket2", "Trinket 2"}, {"main_hand", "Main hand"},
	{SlotOff, "Off hand"}, {"ranged", "Ranged"},
}

// SlotKeys is the 17 slot keys in the character frame's order.
func SlotKeys() []string {
	out := make([]string, len(slotLabels))
	for i, s := range slotLabels {
		out[i] = s.key
	}
	return out
}

// SlotLabel is "Finger 1" for "finger1", and "" for a key that is not a gear slot.
func SlotLabel(key string) string {
	for _, s := range slotLabels {
		if s.key == key {
			return s.label
		}
	}
	return ""
}

// instanceShort is the fit rule's step 2 (spec §B.3): the abbreviations players use for the
// instances, and nothing else. Dire Maul and Uldaman stay whole ("DM" is also the Deadmines).
// Longer names come first so "Lower Blackrock Spire" never loses only its tail.
var instanceShort = []struct{ long, short string }{
	{"Lower Blackrock Spire", "LBRS"}, {"Upper Blackrock Spire", "UBRS"}, {"Blackrock Depths", "BRD"},
	{"Temple of Ahn'Qiraj", "AQ40"}, {"Ruins of Ahn'Qiraj", "AQ20"}, {"Stratholme", "Strat"},
	{"Scholomance", "Scholo"}, {"Molten Core", "MC"}, {"Blackwing Lair", "BWL"}, {"Onyxia's Lair", "Ony"},
	{"Zul'Gurub", "ZG"}, {"Naxxramas", "Naxx"}, {"Sunken Temple", "ST"}, {"Zul'Farrak", "ZF"},
	{"Maraudon", "Mara"}, {"Scarlet Monastery", "SM"}, {"Razorfen Downs", "RFD"}, {"Razorfen Kraul", "RFK"},
	{"Shadowfang Keep", "SFK"}, {"Blackfathom Deeps", "BFD"}, {"Wailing Caverns", "WC"},
	{"Ragefire Chasm", "RFC"}, {"Gnomeregan", "Gnomer"},
}

var (
	questLine   = regexp.MustCompile(`^Quest: .* · (requires level \d+)$`)
	pvpLine     = regexp.MustCompile(`^PvP rank (\d+) · .* · (Alliance|Horde)$`)
	craftedLine = regexp.MustCompile(`^(Crafted: [^·]+?) · .*$`)
)

// ShortSource is a source line's short form (spec §B.3 fit rule, step 2, copy rows
// source.*.short): a quest's requirement without its title, a PvP rank without its title, a
// recipe by its profession, an instance by its abbreviation. A line with nothing to shorten
// comes back unchanged.
func ShortSource(full string) string {
	if m := questLine.FindStringSubmatch(full); m != nil {
		return "Quest · " + m[1]
	}
	if strings.HasPrefix(full, "Quest: ") {
		return "Quest"
	}
	if m := pvpLine.FindStringSubmatch(full); m != nil {
		return "PvP rank " + m[1] + " · " + m[2]
	}
	if m := craftedLine.FindStringSubmatch(full); m != nil {
		return m[1]
	}
	out := full
	for _, in := range instanceShort {
		out = strings.ReplaceAll(out, in.long, in.short)
	}
	return out
}

// metaJoin joins a meta line's slot label to what follows it ("Head · Vendor: Mokvar").
const metaJoin = " · "

// labelled leads each meta candidate with the slot's label, then offers the label alone, so
// every cell names its slot whatever fits (the persona round's ruling).
func labelled(label string, candidates ...string) []string {
	out := make([]string, 0, len(candidates)+1)
	for _, c := range candidates {
		if c != "" {
			out = append(out, label+metaJoin+c)
		}
	}
	return append(out, label)
}

// metaCandidates is a /bis cell's meta lines, longest first and without repeats: the full
// source line, then its short form.
func metaCandidates(full string) []string {
	if full == "" {
		return nil
	}
	short := ShortSource(full)
	if short == full || short == "" {
		return []string{full}
	}
	return []string{full, short}
}
