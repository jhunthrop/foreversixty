package activity

import "time"

// HomeAnswer is GET /v1/home's body. Every block is null when it has nothing true to say.
type HomeAnswer struct {
	ImportSim *ImportSim `json:"import_sim"`
	SpecWeek  *SpecWeek  `json:"spec_week"`
	Movers    *Movers    `json:"movers"`
	GuildSims *GuildSims `json:"guild_sims"`
}

// The import sim's states.
const (
	SimPending = "pending"
	SimDone    = "done"
	SimFailed  = "failed"
)

// ImportSim is the character's latest import run.
type ImportSim struct {
	State    string     `json:"state"`
	DPS      *float64   `json:"dps"`
	SimmedAt *time.Time `json:"simmed_at"`
	SimID    *string    `json:"sim_id"`
}

// SpecWeek is the character's spec on the tier list's rules this week.
type SpecWeek struct {
	Spec          string  `json:"spec"`
	Band          int     `json:"band"`
	Faction       string  `json:"faction"`
	Role          string  `json:"role"`
	Grade         *string `json:"grade"`
	LastWeekGrade *string `json:"last_week_grade"`
	Place         int     `json:"place"`
	Of            int     `json:"of"`
}

// Movers is the spec board's movement since last week, or its top 3 with no baseline.
type Movers struct {
	Spec     string     `json:"spec"`
	Band     int        `json:"band"`
	Baseline bool       `json:"baseline"`
	Rows     []MoverRow `json:"rows"`
}

// MoverRow is one character on the board.
type MoverRow struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Key   string `json:"key"`
	Delta *int   `json:"delta"`
}

// GuildSims is the viewer's guild ladder: the top rows and the viewer's own.
type GuildSims struct {
	GuildID   int64       `json:"guild_id"`
	GuildName string      `json:"guild_name"`
	Role      string      `json:"role"`
	Band      int         `json:"band"`
	Preset    string      `json:"preset"`
	Unit      string      `json:"unit"`
	Of        int         `json:"of"`
	Rows      []LadderRow `json:"rows"`
	You       *LadderRow  `json:"you"`
}
