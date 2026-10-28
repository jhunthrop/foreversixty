package coach

import (
	"math"
	"slices"
	"strconv"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// F8's fact labels, and adherenceEntry, one spell in DetailAdherence's list as §8.1's
// example shows it ("Heroic Strike (31), Slam (12) and Rend (9)").
const (
	adherenceEntry         = "{Spell} ({n})"
	adherenceFactYou       = "casts from the rotation"
	adherenceFactMedian    = "casts from the rotation, top logs"
	adherenceCastsPerSpell = "{Spell} casts outside the rotation"
	adherenceWholeRotation = 100.0
	adherenceUnitCasts     = "casts"
)

// adherence is F8 (§6.2): the share of a DPS's or tank's casts that are the written
// rotation's, with the spells cast most outside it. A healer, a spec whose rule file holds
// no rotation and a log with too few casts get nothing.
func adherence(in Input, _ Counted) outcome {
	if (in.Row.Role != RoleDPS && in.Row.Role != RoleTank) || !in.Tables.RotationOK {
		return outcome{}
	}
	share := rating.FightAdherence(in.Fight, in.Row, in.Tables.Rotation)
	if share == nil {
		return outcome{}
	}
	you := math.Round(*share * percentOf)
	ref, basis := adherenceReference(in)
	switch {
	case adherenceFires(you, ref, basis):
		return adherenceOutcome(in, you, ref, basis)
	case basis.Kind == BasisTopLogs && you >= ref:
		return adherenceStrength(in, you, ref, basis)
	}
	return outcome{}
}

// adherenceReference is the top logs' median adherence when the pool has one, else the
// written rotation alone (every cast from it).
func adherenceReference(in Input) (float64, Basis) {
	pool := func(p Profile) (float64, bool) {
		v, ok := p.AdherenceOf()
		return math.Round(v * percentOf), ok
	}
	if ref, b, _ := comparison(in, pool, nil, false); b != nil {
		return ref, *b
	}
	label := Fill(BasisAdherenceWritten, map[string]string{"Spec": in.Row.Spec, "floor": PctNumber(AdherenceFiresUnderPct)})
	return adherenceWholeRotation, Basis{Label: label, Kind: BasisOwnLog}
}

// adherenceFires: with a pool, 10 points under its median; against the rotation alone,
// under 85%.
func adherenceFires(you, ref float64, b Basis) bool {
	if b.Kind == BasisTopLogs {
		return ref-you >= AdherenceFiresPoints
	}
	return you < AdherenceFiresUnderPct
}

func adherenceOutcome(in Input, you, ref float64, b Basis) outcome {
	off := rating.OffRotation(in.Fight, in.Row, in.Tables.Rotation, OffRotationShown)
	anchors := make([]Anchor, len(off))
	refs := make([]string, len(off))
	for i, row := range off {
		anchors[i] = Anchor{Kind: AnchorCast, SpellID: row.SpellID}
		refs[i] = "#" + strconv.Itoa(i)
	}
	severity := SeverityMinor
	if ref-you >= AdherenceMajorPoints {
		severity = SeverityMajor
	}
	pct := PctNumber(you)
	return outcome{
		Anchors: anchors,
		Finding: &Finding{
			Code:     CodeAdherence,
			Severity: severity,
			Title:    Fill(TitleAdherence, map[string]string{"pct": pct}),
			Detail:   Fill(DetailAdherence, map[string]string{"spells": offRotationNames(off)}),
			Basis:    adherenceBasis(in, b, ref, you),
			Anchors:  refs,
			Facts:    adherenceFacts(in, off, you, ref, b),
		},
	}
}

func adherenceStrength(in Input, you, ref float64, b Basis) outcome {
	pct := PctNumber(you)
	top := rotationTopCasts(in, OffRotationShown)
	anchors := make([]Anchor, len(top))
	refs := make([]string, len(top))
	for i, id := range top {
		anchors[i] = Anchor{Kind: AnchorCast, SpellID: id}
		refs[i] = "#" + strconv.Itoa(i)
	}
	return outcome{Anchors: anchors, Strengths: []Strength{{
		Sentence: Fill(StrengthAdherence, map[string]string{"pct": pct, "n": strconv.Itoa(b.N), "ref": PctNumber(ref)}),
		Facts: Facts{Numbers: []NumberFact{
			{V: you, Text: pct, Unit: "%", Of: adherenceFactYou},
			{V: ref, Text: PctNumber(ref), Unit: "%", Of: adherenceFactMedian},
			{V: float64(b.N), Text: strconv.Itoa(b.N), Unit: "players", Of: "top logs"},
		}},
		Anchors: refs,
	}}}
}

// rotationTopCasts is the n rotation spells the player cast most (successful casts summed
// per spell id), most first, ties by spell id.
func rotationTopCasts(in Input, n int) []int64 {
	bySpell := map[int64]int64{}
	for _, c := range in.Fight.Casts {
		if c.GUID == in.Row.GUID && c.Succeeded > 0 && in.Tables.Rotation.Lists(c.SpellID) {
			bySpell[c.SpellID] += c.Succeeded
		}
	}
	ids := make([]int64, 0, len(bySpell))
	for id := range bySpell {
		ids = append(ids, id)
	}
	slices.SortFunc(ids, func(a, b int64) int {
		if bySpell[a] != bySpell[b] {
			return int(bySpell[b] - bySpell[a])
		}
		return int(a - b)
	})
	return ids[:min(n, len(ids))]
}

// adherenceBasis words the pool's basis with both values; the rotation-alone basis is
// its own sentence.
func adherenceBasis(in Input, b Basis, ref, you float64) Basis {
	if b.Kind != BasisTopLogs {
		return b
	}
	fields := topFields(in, b.N)
	fields["ref"], fields["you"] = PctNumber(ref), PctNumber(you)
	b.Label = Fill(BasisAdherenceTop, fields)
	return b
}

func offRotationNames(off []summary.CastRow) string {
	names := make([]string, len(off))
	for i, row := range off {
		names[i] = Fill(adherenceEntry, map[string]string{"Spell": row.SpellName, "n": strconv.FormatInt(row.Succeeded, 10)})
	}
	return joinAnd(names)
}

func adherenceFacts(in Input, off []summary.CastRow, you, ref float64, b Basis) Facts {
	f := Facts{Numbers: []NumberFact{{V: you, Text: PctNumber(you), Unit: "%", Of: adherenceFactYou}}}
	if b.Kind == BasisTopLogs {
		f.Numbers = append(f.Numbers,
			NumberFact{V: ref, Text: PctNumber(ref), Unit: "%", Of: adherenceFactMedian},
			NumberFact{V: float64(b.N), Text: strconv.Itoa(b.N), Unit: "players", Of: "top logs"})
	}
	for _, row := range off {
		f.Numbers = append(f.Numbers, NumberFact{
			V: float64(row.Succeeded), Text: strconv.FormatInt(row.Succeeded, 10), Unit: adherenceUnitCasts,
			Of: Fill(adherenceCastsPerSpell, map[string]string{"Spell": row.SpellName}),
		})
		f.Spells = append(f.Spells, SpellFact{ID: row.SpellID, Name: row.SpellName, Rank: in.Tables.Spells[row.SpellID].Rank})
	}
	return f
}
