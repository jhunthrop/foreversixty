package rating

import (
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
)

// historyDTO is GET /v1/characters/{region}/{ruleset}/{name}/history's body
// (design/specs/2026-10-24-character-history.md section 10). Every array is present, never null.
type historyDTO struct {
	AsOf       string                `json:"as_of"`
	Score      []historyPointDTO     `json:"score"`
	Encounters []historyEncounterDTO `json:"encounters"`
	Gaps       []historyGapDTO       `json:"gaps"`
	Fights     []rankings.OwnerFight `json:"fights"`
}

type historyPointDTO struct {
	At              string   `json:"at"`
	Score           int      `json:"score"`
	Tier            string   `json:"tier"`
	ReportID        string   `json:"report_id"`
	FightIndex      int      `json:"fight_index"`
	Encounter       string   `json:"encounter"`
	EncounterID     int64    `json:"encounter_id"`
	Difficulty      int64    `json:"difficulty"`
	ParsePercentile *float64 `json:"parse_percentile,omitempty"`
	Kill            bool     `json:"kill"`
	Visibility      string   `json:"visibility"`
}

type historyParseDTO struct {
	At              string  `json:"at"`
	ParsePercentile float64 `json:"parse_percentile"`
	ReportID        string  `json:"report_id"`
	FightIndex      int     `json:"fight_index"`
}

type historyEncounterDTO struct {
	Encounter   string            `json:"encounter"`
	EncounterID int64             `json:"encounter_id"`
	Difficulty  int64             `json:"difficulty"`
	Points      []historyParseDTO `json:"points"`
}

type historyGapDTO struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Reason string `json:"reason"`
}

func utc(t time.Time) string { return t.UTC().Format(time.RFC3339) }

func toHistoryDTO(h History) historyDTO {
	out := historyDTO{
		AsOf:       h.AsOf,
		Score:      make([]historyPointDTO, len(h.Score)),
		Encounters: make([]historyEncounterDTO, len(h.Encounters)),
		Gaps:       make([]historyGapDTO, len(h.Gaps)),
		Fights:     h.Fights,
	}
	if out.Fights == nil {
		out.Fights = []rankings.OwnerFight{}
	}
	for i, p := range h.Score {
		out.Score[i] = historyPointDTO{
			At: utc(p.At), Score: p.Score, Tier: p.Tier, ReportID: p.ReportID,
			FightIndex: p.FightIndex, Encounter: p.Encounter, EncounterID: p.EncounterID,
			Difficulty: p.Difficulty, ParsePercentile: p.ParsePercentile, Kill: p.Kill, Visibility: p.Visibility,
		}
	}
	for i, e := range h.Encounters {
		pts := make([]historyParseDTO, len(e.Points))
		for j, p := range e.Points {
			pts[j] = historyParseDTO{At: utc(p.At), ParsePercentile: p.ParsePercentile,
				ReportID: p.ReportID, FightIndex: p.FightIndex}
		}
		out.Encounters[i] = historyEncounterDTO{Encounter: e.Encounter, EncounterID: e.EncounterID,
			Difficulty: e.Difficulty, Points: pts}
	}
	for i, g := range h.Gaps {
		out.Gaps[i] = historyGapDTO{From: g.From, To: g.To, Reason: g.Reason}
	}
	return out
}
