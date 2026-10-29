package sheet

import (
	"strconv"
	"strings"
)

// AltMax is Discord's limit on an attachment's and a gallery item's description (spec §0).
const AltMax = 1024

// Alt is the sheet's alt text (spec §B.7), at most AltMax characters and never "…":
//
//	/bis:  Best in slot sheet, {Spec} {Class}, {band label}, {Faction}. {Slot}: {Item}; …
//	/char: Character sheet, {Name}, level {n} {Spec} {Class}, {Guild}. {Slot}: {Item}, {gain} available; …
//
// Over the limit, the gain clauses go first, then the slots from the end, whole, and the text
// ends with ".".
func Alt(s Sheet) string {
	lead := altLead(s)
	full := lead + strings.Join(altSlots(s, true), "; ") + "."
	if len([]rune(full)) <= AltMax {
		return full
	}
	parts := altSlots(s, false)
	for n := len(parts); n > 0; n-- {
		out := lead + strings.Join(parts[:n], "; ") + "."
		if len([]rune(out)) <= AltMax {
			return out
		}
	}
	return string([]rune(strings.TrimSuffix(lead, " "))[:min(len([]rune(lead)), AltMax)])
}

func altLead(s Sheet) string {
	class := className(s.Class)
	faction := "Alliance"
	if s.Faction == "horde" {
		faction = "Horde"
	}
	if s.Kind == KindBis {
		return altBisLead + joinKnown(", ", strings.TrimSpace(s.Spec+" "+class), s.Band, faction) + ". "
	}
	who := ""
	if s.Level > 0 {
		who = "level " + strconv.Itoa(s.Level) + " "
	}
	return altCharLead + joinKnown(", ", s.Title, who+strings.TrimSpace(s.Spec+" "+class), s.Guild) + ". "
}

// altSlots is one clause per slot in frame order; gains only on /char, only figures (an
// unmodelled gain is left out), and only when withGains.
func altSlots(s Sheet, withGains bool) []string {
	out := make([]string, 0, len(slotLabels))
	for _, sl := range slotLabels {
		if sl.key == SlotOff && s.TwoHand {
			out = append(out, CopyAltOffHand)
			continue
		}
		c := s.slot(sl.key)
		switch {
		case c.Empty && s.Kind == KindChar:
			out = append(out, sl.label+": "+altEmpty)
		case c.Empty:
			out = append(out, sl.label+": "+altNoPick)
		default:
			clause := sl.label + ": " + c.Name
			if withGains && s.Kind == KindChar && isFigure(c.Gain) {
				clause += ", " + c.Gain + altGainTail
			}
			out = append(out, clause)
		}
	}
	return out
}

// isFigure reports whether a gold line is a measured gain ("+9.3 DPS"), not a word.
func isFigure(gain string) bool {
	return gain != "" && gain != CopyNotModelled && gain != CopyUpgrade
}

// className is "Hunter" for "hunter".
func className(slug string) string {
	if slug == "" {
		return ""
	}
	return strings.ToUpper(slug[:1]) + slug[1:]
}

func joinKnown(sep string, parts ...string) string {
	kept := parts[:0:0]
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			kept = append(kept, p)
		}
	}
	return strings.Join(kept, sep)
}
