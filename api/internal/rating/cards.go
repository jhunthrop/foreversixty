package rating

import (
	"encoding/json"
	"time"

	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// momentDTO is one Moment, spec §5.2's exact JSON field names.
type momentDTO struct {
	Kind      string `json:"kind"`
	AtMS      int64  `json:"at_ms"`
	SpellID   int64  `json:"spell_id,omitempty"`
	SpellName string `json:"spell_name,omitempty"`
	Avoidable bool   `json:"avoidable,omitempty"`
	Anchor    string `json:"anchor,omitempty"`
}

// componentDTO is one Component, spec §5.2's exact JSON field names. Score is a pointer so
// an excluded component serialises "score": null (spec §5.2: "an excluded entry carries
// ... score: null").
type componentDTO struct {
	Name       string      `json:"name"`
	Score      *float64    `json:"score"`
	Weight     float64     `json:"weight"`
	Basis      string      `json:"basis"`
	Percentile *float64    `json:"percentile,omitempty"`
	BracketN   int64       `json:"bracket_n,omitempty"`
	Excluded   bool        `json:"excluded"`
	Reason     string      `json:"reason,omitempty"`
	Moments    []momentDTO `json:"moments"`
}

func toMomentDTOs(ms []ratingengine.Moment) []momentDTO {
	out := make([]momentDTO, 0, len(ms))
	for _, m := range ms {
		out = append(out, momentDTO{
			Kind: m.Kind, AtMS: m.AtMS, SpellID: m.SpellID, SpellName: m.SpellName,
			Avoidable: m.Avoidable, Anchor: m.Anchor,
		})
	}
	return out
}

// toComponentDTOs converts the engine's six components into the site's response shape.
// Used both to build the stored components jsonb column and to build the two HTTP
// responses, so the two never drift out of sync with each other.
func toComponentDTOs(cs [6]ratingengine.Component) []componentDTO {
	out := make([]componentDTO, 0, len(cs))
	for _, c := range cs {
		d := componentDTO{
			Name: c.Name, Weight: c.Weight, Basis: c.Basis, BracketN: c.BracketN,
			Excluded: c.Excluded, Reason: c.Reason, Moments: toMomentDTOs(c.Moments),
		}
		if !c.Excluded {
			score := c.Score
			d.Score = &score
		}
		if c.Percentile != nil {
			pct := *c.Percentile
			d.Percentile = &pct
		}
		out = append(out, d)
	}
	return out
}

// fightMeta is the per-fight fields every roster player's stored row shares - read once
// per fight, not recomputed per player.
type fightMeta struct {
	ReportID     string
	FightIndex   int
	EncounterID  int64
	Difficulty   int64
	Size         int
	DurationMS   int64
	Kill         bool
	KillTimeBand string
	ModelVersion string
	FoughtAt     time.Time
}

// CardRow is one player's rating_scores row, ready to bind to an insert statement.
// PlayerKey is left blank here: the caller (Store.RateFight) is the one place that knows
// the fight's region/ruleset, so it fills PlayerKey via character.KeyFromUnit after
// newCardRow returns, rather than this package importing character.KeyFromUnit's inputs
// just to thread two more strings through this function's signature.
type CardRow struct {
	ReportID, PlayerKey, PlayerName, Class, Spec, Role, KillTimeBand, ModelVersion string
	FightIndex                                                                     int
	EncounterID, Difficulty                                                        int64
	Size                                                                           int
	DurationMS                                                                     int64
	Kill, OverallCapped                                                            bool
	Overall, OverallUncapped                                                       float64
	Components                                                                     json.RawMessage
	FoughtAt                                                                       time.Time
}

// newCardRow builds a CardRow from one player's already-scored Card. PlayerKey is left
// empty; the caller sets it.
func newCardRow(f fightMeta, row summary.RosterRow, card ratingengine.Card) CardRow {
	components, err := json.Marshal(toComponentDTOs(card.Components))
	if err != nil {
		// toComponentDTOs never produces a value json.Marshal can refuse (no channels,
		// funcs, or cycles in componentDTO), so this is unreachable outside a future,
		// accidentally-unmarshalable field addition - panicking here turns that mistake
		// into a test failure immediately rather than a silently empty components column.
		panic("rating: components must always marshal: " + err.Error())
	}
	return CardRow{
		ReportID: f.ReportID, FightIndex: f.FightIndex, PlayerName: row.Name,
		Class: row.Class, Spec: row.Spec, Role: row.Role,
		EncounterID: f.EncounterID, Difficulty: f.Difficulty, Size: f.Size,
		DurationMS: f.DurationMS, Kill: f.Kill, KillTimeBand: f.KillTimeBand,
		Overall: card.Overall, OverallUncapped: card.OverallUncapped, OverallCapped: card.OverallCapped,
		Components: components, ModelVersion: card.ModelVersion, FoughtAt: f.FoughtAt,
	}
}
