package app

import (
	"math"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/bis"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// homeInput is everything one answer reads, already loaded.
type homeInput struct {
	card  rankings.Character
	level *int
	spec  string
	// band is the character's band, found false when none applies or none is published.
	band  bis.Band
	found bool
	// talents are the export's tree rank strings; hasExport is false with no decodable export.
	talents   []string
	hasExport bool
	synced    *time.Time
	source    string
	build     string
	items     upgrades.Items
}

// answer is the home answer over in.
func answer(in homeInput) Home {
	h := Home{Character: in.card.Character.Key, DataBuild: in.build, Upgrades: []Upgrade{}, AlreadyBest: []Best{},
		Spec: optional(in.spec), SyncedAt: in.synced, SyncedFrom: optional(in.source), Talents: talentsOf(in)}
	if in.found {
		unit := upgrades.UnitOf(in.band)
		h.Band, h.BandLabel, h.Preset = &in.band.Band, ptrTo(bandLabel(in.band.Band)), &in.band.Preset
		h.Unit, h.UnitWord = ptrTo(string(unit)), ptrTo(unit.Word())
	}
	if h.State = stateOf(in); h.State != StateOK {
		return h
	}
	cmp := upgrades.Compare(in.band, wornGear(in.card.Gear), in.items)
	for _, r := range cmp.Upgrades {
		u := in.upgrade(r, cmp.Unit)
		if !u.Modelled {
			h.NotModelledCount++
		}
		h.Upgrades = append(h.Upgrades, u)
	}
	for _, b := range cmp.AlreadyBest {
		h.AlreadyBest = append(h.AlreadyBest, Best{Slot: b.Slot, SlotLabel: upgrades.SlotLabel(b.Slot), Item: in.item(b.ItemID, b.ItemName)})
	}
	h.UpgradeCount, h.NotSimCheckedCount = len(h.Upgrades), cmp.NotSimChecked
	if cmp.TotalGain > 0 {
		h.TotalGain = ptrTo(oneDecimal(cmp.TotalGain))
	}
	return h
}

// stateOf is why there is no comparison, or StateOK.
func stateOf(in homeInput) string {
	switch {
	case in.level != nil && *in.level < upgrades.LevelBands[0]:
		return StateUnder20
	case len(in.card.Gear) == 0:
		return StateNoGear
	case in.spec == "":
		return StateNoSpec
	case !in.found:
		return StateNoList
	}
	return StateOK
}

// talentsOf is the worn talents against the band's build.
func talentsOf(in homeInput) Talents {
	switch {
	case !in.hasExport:
		return Talents{State: TalentsNoExport}
	case !in.found || in.band.Talents == "":
		return Talents{State: TalentsNoBuild}
	}
	differ := upgrades.UnmatchedPoints(in.talents, in.band.Talents)
	state := TalentsDiffer
	if differ == 0 {
		state = TalentsOptimized
	}
	return Talents{State: state, PointsDiffer: &differ, BandPoints: &in.band.TalentPoints}
}

// upgrade is one comparison row as the apps draw it.
func (in homeInput) upgrade(r upgrades.Row, unit upgrades.Unit) Upgrade {
	u := Upgrade{Slot: r.Slot, SlotLabel: upgrades.SlotLabel(r.Slot),
		Pick: Pick{Item: in.item(r.Pick.ItemID, r.Pick.ItemName), SourceLine: upgrades.SourceLine(r.Pick.Source, r.Pick.SourceKind)}}
	if r.WornID > 0 {
		u.Worn = in.worn(r.WornID, r.WornKnown)
	}
	switch {
	case r.Gain != nil && !r.Pick.EffectUnmodelled:
		u.Gain, u.GainWords, u.Modelled = ptrTo(oneDecimal(*r.Gain)), upgrades.SignedFigure(*r.Gain, unit.Word()), true
	case r.NotSimChecked:
		u.GainWords = upgrades.NotSimCheckedWords
	default:
		u.GainWords = upgrades.NotModelledWords
		if r.Gain != nil {
			u.Gain = ptrTo(oneDecimal(*r.Gain))
		}
	}
	return u
}

// item is id as the build's table knows it; name is the BiS data's, used when the table
// lacks the item or names it nothing.
func (in homeInput) item(id int, name string) Item {
	out := Item{ID: id, Name: optional(name)}
	it, ok := in.items(id)
	if !ok {
		return out
	}
	if it.Name != "" {
		out.Name = &it.Name
	}
	out.Quality, out.ItemLevel, out.Icon = &it.Quality, &it.ItemLevel, optional(it.Icon)
	return out
}

func (in homeInput) worn(id int, known bool) *Worn {
	if !known {
		return &Worn{Item: Item{ID: id}, Label: WornNotInData}
	}
	it := in.item(id, "")
	return &Worn{Item: it, Known: true, Label: strOf(it.Name)}
}

// wornGear is the card's gear panel as slot -> worn item id.
func wornGear(gear []rankings.GearSlot) map[string]int {
	out := map[string]int{}
	for _, g := range gear {
		if g.Worn != nil {
			out[g.Slot] = g.Worn.ID
		}
	}
	return out
}

// bandLabel is a band's range as the site writes it (web/src/lib/planner/band-title.ts):
// "20 to 29", or "60" for the last band.
func bandLabel(band int) string {
	if band == upgrades.TopBand {
		return strconv.Itoa(band)
	}
	return strconv.Itoa(band) + " to " + strconv.Itoa(band+9)
}

func oneDecimal(v float64) float64 { return math.Round(v*10) / 10 }

func ptrTo[T any](v T) *T { return &v }

// optional is s, or nil when it is empty.
func optional(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func strOf(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
