package livesurface

import (
	"context"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/app"
)

// HomeAnswers reads a character's Home answer (app.Service, lane A3).
type HomeAnswers interface {
	HomeOf(ctx context.Context, region, ruleset, name string) (app.Home, bool, error)
}

// HomeUpgrades is Upgrades over the Home answer: the snapshot's rows, count and total are the
// Home answer's own, so the widget and Home never disagree.
type HomeUpgrades struct{ Homes HomeAnswers }

// The sentences the widget shows for a Home answer with no rows (spec §5.6 states; the
// no-gear line is widget.noGear, the others are the Home answer's own words).
const (
	sentenceNoGear  = "Add the addon to see your upgrades."
	sentenceUnder20 = "Best in slot starts at 20."
	sentenceNoSpec  = "Set a spec in the planner to see this."
	sentenceNoList  = "No best in slot list published yet for this spec and level range."
	sentenceAllBest = "Every slot matches the list."
)

// WidgetUpgrades answers nil for a character whose page is not found.
func (h HomeUpgrades) WidgetUpgrades(ctx context.Context, characterKey string) (*WidgetUpgrades, error) {
	parts := strings.SplitN(characterKey, "/", 3)
	if len(parts) != 3 {
		return nil, nil
	}
	home, found, err := h.Homes.HomeOf(ctx, parts[0], parts[1], parts[2])
	if err != nil || !found {
		return nil, err
	}
	return upgradesOf(home), nil
}

// upgradesOf is the Home answer as the snapshot carries it: every row (WidgetFor keeps the
// first three), the Home answer's count and total, and a sentence for any answer with no rows.
func upgradesOf(h app.Home) *WidgetUpgrades {
	out := &WidgetUpgrades{State: h.State, Count: h.UpgradeCount, TotalGain: h.TotalGain, Rows: []UpgradeRow{}}
	if h.Unit != nil {
		out.Unit = *h.Unit
	}
	for _, u := range h.Upgrades {
		out.Rows = append(out.Rows, UpgradeRow{Slot: u.Slot, SlotLabel: u.SlotLabel, Gain: u.Gain, Pick: pickOf(u.Pick)})
	}
	if len(out.Rows) == 0 {
		out.Sentence = sentenceFor(h.State)
	}
	return out
}

func sentenceFor(state string) *string {
	words := map[string]string{
		app.StateNoGear: sentenceNoGear, app.StateUnder20: sentenceUnder20, app.StateNoSpec: sentenceNoSpec,
		app.StateNoList: sentenceNoList, app.StateOK: sentenceAllBest,
	}[state]
	if words == "" {
		return nil
	}
	return &words
}

func pickOf(p app.Pick) UpgradePick {
	out := UpgradePick{ID: p.ID}
	if p.Name != nil {
		out.Name = *p.Name
	}
	if p.Quality != nil {
		out.Quality = *p.Quality
	}
	if p.Icon != nil {
		out.Icon = *p.Icon
	}
	if p.ItemLevel != nil {
		out.ItemLevel = *p.ItemLevel
	}
	return out
}
