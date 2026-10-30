package upgrades

import (
	"fmt"
	"math"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
)

// The words a gain reads when there is no figure to state, shared by the Discord panel and
// the apps (spec 2026-10-29 copy table; site lib/bis/copy.ts notSimCheckedTag).
const (
	// NotModelledWords: the pick's gain has no figure, or its effect is not in the sim.
	NotModelledWords = "Gain not modelled yet"
	// NotSimCheckedWords: a weapon the pick does not list, whose gain only a full sim measures.
	NotSimCheckedWords = "not sim-checked"
)

// minusSign is U+2212, the SimC convention for a negative figure.
const minusSign = "−"

// SignedFigure is a gain or loss to one decimal with its sign: "+9.2 DPS", "−15.0 DPS".
func SignedFigure(v float64, unit string) string {
	v = math.Round(v*10) / 10
	sign := "+"
	if v < 0 {
		sign, v = minusSign, -v
	}
	return fmt.Sprintf("%s%.1f %s", sign, v, unit)
}

// SourceLine is a pick's short source (web/src/lib/home/source-line.ts): a crafted pick adds
// " · crafted"; every other source reads as it is.
func SourceLine(source, kind string) string {
	if kind == "crafted" {
		return source + " · crafted"
	}
	return source
}

// Unit is what a band's gains count (web/src/lib/bis/tank-view.ts SlotScoreUnit).
type Unit string

const (
	UnitDPS   Unit = "dps"
	UnitHPS   Unit = "hps"
	UnitScore Unit = "tank_score"
)

// UnitOf is a band's unit by its role.
func UnitOf(band bis.Band) Unit {
	switch band.Role {
	case "tank":
		return UnitScore
	case "healer":
		return UnitHPS
	}
	return UnitDPS
}

// Word is the unit as a figure reads it: "DPS", "HPS" or "score".
func (u Unit) Word() string {
	switch u {
	case UnitScore:
		return "score"
	case UnitHPS:
		return "HPS"
	}
	return "DPS"
}
