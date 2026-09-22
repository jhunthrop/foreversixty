// api/internal/rating/dto.go
package rating

import (
	"encoding/json"
	"time"
)

// playerRatingDTO is one roster player's rating, as either read endpoint returns it.
type playerRatingDTO struct {
	PlayerKey       string          `json:"player_key"`
	PlayerName      string          `json:"player_name"`
	Class           string          `json:"class"`
	Spec            string          `json:"spec"`
	Role            string          `json:"role"`
	Overall         float64         `json:"overall"`
	OverallUncapped float64         `json:"overall_uncapped"`
	OverallCapped   bool            `json:"overall_capped"`
	Components      json.RawMessage `json:"components"`
}

func toPlayerRatingDTO(cr CardRow) playerRatingDTO {
	return playerRatingDTO{
		PlayerKey: cr.PlayerKey, PlayerName: cr.PlayerName, Class: cr.Class, Spec: cr.Spec, Role: cr.Role,
		Overall: cr.Overall, OverallUncapped: cr.OverallUncapped, OverallCapped: cr.OverallCapped,
		Components: cr.Components,
	}
}

// fightRatingsDTO is the body of GET /v1/reports/{id}/fights/{n}/ratings.
type fightRatingsDTO struct {
	FightIndex   int               `json:"fight_index"`
	Kill         bool              `json:"kill"`
	KillTimeBand string            `json:"kill_time_band"`
	ModelVersion string            `json:"model_version"`
	Players      []playerRatingDTO `json:"players"`
}

// trendPointDTO is one point of the character endpoint's trend[].
type trendPointDTO struct {
	FoughtAt   time.Time `json:"fought_at"`
	Overall    float64   `json:"overall"`
	ReportID   string    `json:"report_id"`
	FightIndex int       `json:"fight_index"`
}

// characterRatingDTO is the body of GET /v1/characters/{region}/{ruleset}/{name}/rating.
type characterRatingDTO struct {
	PlayerKey      string           `json:"player_key"`
	SampleSize     int              `json:"sample_size"`
	Trend          []trendPointDTO  `json:"trend"`
	NextCursor     string           `json:"next_cursor,omitempty"`
	Latest         *playerRatingDTO `json:"latest,omitempty"`
	BestComponent  string           `json:"best_component,omitempty"`
	WorstComponent string           `json:"worst_component,omitempty"`
}

func buildCharacterRatingDTO(playerKey string, rows []CardRow, hasMore bool) characterRatingDTO {
	out := characterRatingDTO{PlayerKey: playerKey, SampleSize: len(rows)}
	for _, cr := range rows {
		out.Trend = append(out.Trend, trendPointDTO{
			FoughtAt: cr.FoughtAt, Overall: cr.Overall, ReportID: cr.ReportID, FightIndex: cr.FightIndex,
		})
	}
	if len(rows) > 0 {
		latest := toPlayerRatingDTO(rows[0])
		out.Latest = &latest
		out.BestComponent, out.WorstComponent = bestWorstComponent(latest.Components)
	}
	if hasMore && len(rows) > 0 {
		last := rows[len(rows)-1]
		out.NextCursor = encodeTrendCursor(cursorPos{FoughtAt: last.FoughtAt, ReportID: last.ReportID, FightIndex: last.FightIndex})
	}
	return out
}

// bestWorstComponent decodes latest's stored components (the most recent card only - not
// the whole trend) to name the highest- and lowest-scoring non-excluded component, for the
// spec §5.2 example payload's best_component/worst_component fields.
func bestWorstComponent(latest json.RawMessage) (best, worst string) {
	var cs []componentDTO
	if err := json.Unmarshal(latest, &cs); err != nil {
		return "", ""
	}
	var bestScore, worstScore *float64
	for _, c := range cs {
		if c.Excluded || c.Score == nil {
			continue
		}
		if bestScore == nil || *c.Score > *bestScore {
			bestScore, best = c.Score, c.Name
		}
		if worstScore == nil || *c.Score < *worstScore {
			worstScore, worst = c.Score, c.Name
		}
	}
	return best, worst
}
