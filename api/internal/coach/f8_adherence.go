package coach

import (
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// F8 pack strings the spec does not show (drafts until Task 33 pins them).
const (
	basisAdherenceValue    = "{ref}% of casts from the rotation. You: {you}%."
	basisAdherenceRotation = "Measured against the written rotation alone."
	adherenceFactYou       = "casts from the rotation"
	adherenceFactMedian    = "casts from the rotation, top logs"
	adherenceCastsPerSpell = "{Spell} casts outside the rotation"
	adherenceWholeRotation = 100.0
	adherenceUnitCasts     = "casts"
	adherenceSeparator     = ", "
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
	you := *share * percentOf
	ref, basis := adherenceReference(in)
	switch {
	case adherenceFires(you, ref, basis):
		return adherenceOutcome(in, you, ref, basis)
	case basis.Kind == BasisTopLogs && you >= ref:
		return adherenceStrength(you, ref, basis)
	}
	return outcome{}
}

// adherenceReference is the top logs' median adherence when the pool has one, else the
// written rotation alone (every cast from it).
func adherenceReference(in Input) (float64, Basis) {
	pool := func(p Profile) (float64, bool) {
		v, ok := p.AdherenceOf()
		return v * percentOf, ok
	}
	if ref, b, _ := comparison(in, pool, nil, false); b != nil {
		return ref, *b
	}
	return adherenceWholeRotation, Basis{Label: basisAdherenceRotation, Kind: BasisOwnLog}
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
			Title:    Fill(Drafts["title_adherence"], map[string]string{"pct": pct}),
			Detail:   Fill(Drafts["detail_adherence"], map[string]string{"pct": pct, "spells": offRotationNames(off)}),
			Basis:    adherenceBasis(b, ref, you),
			Anchors:  refs,
			Facts:    adherenceFacts(in, off, you, ref, b),
		},
	}
}

func adherenceStrength(you, ref float64, b Basis) outcome {
	pct := PctNumber(you)
	return outcome{Strengths: []Strength{{
		Sentence: Fill(Drafts["strength_adherence"], map[string]string{"pct": pct}),
		Facts: Facts{Numbers: []NumberFact{
			{V: you, Text: pct, Unit: "%", Of: adherenceFactYou},
			{V: ref, Text: PctNumber(ref), Unit: "%", Of: adherenceFactMedian},
			{V: float64(b.N), Text: strconv.Itoa(b.N), Unit: "players", Of: "top logs"},
		}},
		Anchors: []string{},
	}}}
}

// adherenceBasis appends the values to the pool's label; the rotation-alone basis is its
// own sentence.
func adherenceBasis(b Basis, ref, you float64) Basis {
	if b.Kind != BasisTopLogs {
		return b
	}
	b.Label += Fill(basisAdherenceValue, map[string]string{"ref": PctNumber(ref), "you": PctNumber(you)})
	return b
}

func offRotationNames(off []summary.CastRow) string {
	names := make([]string, len(off))
	for i, row := range off {
		names[i] = row.SpellName
	}
	return strings.Join(names, adherenceSeparator)
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
