package coach

// F8 fixtures: a player whose casts are a fixed count, a share of them the written
// rotation's and the rest spread over three class spells it does not list.

import (
	"math"
	"testing"
)

const (
	adherenceCasts      = 100
	adherenceRotationID = 23894 // Bloodthirst rank 4, listed by the Fury rotation
	adherenceHamstring  = 7373
	adherenceRend       = 11574
	adherenceOverpower  = 11585
	adherenceFirstShare = 0.5 // of the off-rotation casts: the first spell, then
	adherenceNextShare  = 0.3 // the second; the third takes the rest
)

type offSpell struct {
	id   int64
	name string
}

var offSpells = []offSpell{
	{adherenceHamstring, "Hamstring"},
	{adherenceRend, "Rend"},
	{adherenceOverpower, "Overpower"},
}

// adherenceOf is in's player casting adherenceCasts abilities, share of them spellID.
func adherenceOf(t *testing.T, in Input, spellID int64, name string, share float64) Input {
	t.Helper()
	guid := in.Row.GUID
	listed := int(math.Round(share * adherenceCasts))
	off := adherenceCasts - listed
	first := int(float64(off) * adherenceFirstShare)
	next := int(float64(off) * adherenceNextShare)
	counts := []int{first, next, off - first - next}
	f := withoutCasts(in.Fight, guid)
	f = withCasts(f, guid, spellID, name, repeat(0, listed)...)
	for i, s := range offSpells {
		if counts[i] > 0 {
			f = withCasts(f, guid, s.id, s.name, repeat(0, counts[i])...)
		}
	}
	in.Fight = f
	return in
}

func furyAdherence(t *testing.T, share float64, refs Refs) Input {
	t.Helper()
	return withPool(adherenceOf(t, furyInput(t), adherenceRotationID, "Bloodthirst", share), refs)
}

func protAdherence(t *testing.T, share float64, refs Refs) Input {
	t.Helper()
	return withPool(adherenceOf(t, protInput(t), 23925, "Shield Slam", share), refs)
}

func holyAdherence(t *testing.T, share float64) Input {
	t.Helper()
	return adherenceOf(t, holyInput(t), 25314, "Greater Heal", share)
}
