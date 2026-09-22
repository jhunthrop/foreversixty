// api/internal/dataaddon/aggregate.go
package dataaddon

import (
	"encoding/json"
	"math"
)

// componentNames is the addon's six component keys, in Ratings.lua's own
// COMPONENTS order -- the order this package writes them in Data.lua, so a
// diff between two nights' files stays minimal.
var componentNames = []string{"output", "survival", "mechanics", "utility", "preparation", "activity"}

// componentScore is one fight's contribution to one component: nil when
// the engine excluded it for that fight (a stored componentDTO's score is
// JSON null exactly when excluded is true -- see api/internal/rating/cards.go).
type componentScore struct {
	Name  string
	Score *float64
}

// fightScore is one public, in-window, non-anonymized rated fight read for
// one character.
type fightScore struct {
	Overall    float64
	Components []componentScore
}

// characterRow is one character's line in Data.lua: the rounded rating,
// whichever components had at least one non-excluded fight to average
// (absent from the map for the rest -- "may omit what it cannot compute"),
// and the fight count the rating rests on.
type characterRow struct {
	Rating     int
	Components map[string]int
	Fights     int
}

// decodeComponents reads one fight's stored components jsonb
// (api/internal/rating/cards.go's componentDTO array) into the scores this
// package aggregates over. A row whose components column does not decode
// as the expected shape is not a reason to drop the whole fight -- its
// Overall still counts toward the rating -- so this returns nil rather
// than an error, and every component is simply excluded for that one
// fight: "never write a wrong number" outranks "use every fight for every
// field."
func decodeComponents(raw json.RawMessage) []componentScore {
	var decoded []struct {
		Name     string   `json:"name"`
		Score    *float64 `json:"score"`
		Excluded bool     `json:"excluded"`
	}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil
	}
	out := make([]componentScore, 0, len(decoded))
	for _, d := range decoded {
		if d.Excluded {
			out = append(out, componentScore{Name: d.Name, Score: nil})
			continue
		}
		out = append(out, componentScore{Name: d.Name, Score: d.Score})
	}
	return out
}

// aggregateCharacter folds one character's fights into its Data.lua row.
// ok is false for zero fights: a character with no public rated fights in
// the window has nothing to say, and gets no row at all rather than a
// zeroed one.
func aggregateCharacter(fights []fightScore) (row characterRow, ok bool) {
	if len(fights) == 0 {
		return characterRow{}, false
	}
	var overallSum float64
	sums := make(map[string]float64, len(componentNames))
	counts := make(map[string]int, len(componentNames))
	for _, f := range fights {
		overallSum += f.Overall
		for _, c := range f.Components {
			if c.Score == nil {
				continue
			}
			sums[c.Name] += *c.Score
			counts[c.Name]++
		}
	}
	components := make(map[string]int, len(componentNames))
	for _, name := range componentNames {
		if n := counts[name]; n > 0 {
			components[name] = roundHalfUp(sums[name] / float64(n))
		}
	}
	return characterRow{
		Rating:     roundHalfUp(overallSum / float64(len(fights))),
		Components: components,
		Fights:     len(fights),
	}, true
}

// roundHalfUp rounds to the nearest integer, ties away from zero -- Go's
// math.Round's own rule. Every number this package writes (rating and
// every component) takes this same rounding, so Data.lua's formatting
// stays uniform and its golden test has no fractional tie to adjudicate.
func roundHalfUp(v float64) int {
	return int(math.Round(v))
}
