package coach

import (
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics"
	"github.com/jhunthrop/foreversixty/logs/engine/mechanics/consumables"
	ratingengine "github.com/jhunthrop/foreversixty/logs/engine/rating"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Input is everything one player's pack is built from: the fight, the player's roster
// row, the raid size and boss name, the comparison references and the data tables.
type Input struct {
	Fight    summary.Summary
	Row      summary.RosterRow
	Size     int
	BossName string
	Refs     Refs
	Tables   Tables
}

// Tables is the data the checks read beside the fight.
type Tables struct {
	Rules       SpecRules                  // data/curated/coach/<spec>.json; zero value = no file
	Spells      map[int64]SpellConst       // the class's spellconst for the fight's build
	Trainable   map[string][]TrainableRank // family -> ranks with the level each is learned
	Rotation    sims.Rotation
	RotationOK  bool
	Consumables consumables.RoleCatalogue
	RaidBuffs   []consumables.RaidBuff
	Mechanics   *mechanics.Table
	// Assignments are the officer marks ActivityIdleMS subtracts; rating/store.go passes
	// nil today, and so does Task 29b.
	Assignments  []ratingengine.Assignment
	WorstActions map[string]bool // families on the spec's sims WorstActions list
}

// SpellConst is one spell's constants. CooldownMS is the larger of spellconst's
// cooldown_ms and category_cooldown_ms (Bloodthirst's 6 s is a category cooldown in
// 1.60.1.70630). Cost is spellconst's cost, in the same tenths the log's rage series
// uses. CostKnown is false when the data states the cost only as a share of the bar.
type SpellConst struct {
	Name       string
	Rank       int
	CooldownMS int64
	GCDMS      int64
	Cost       int64
	CostType   int64
	CostKnown  bool
}

// TrainableRank is one rank of a trainable family and the level it is learned at.
type TrainableRank struct {
	ID    int64
	Rank  int
	Level int
}

// Refs is the comparison references for one fight and player (§11).
type Refs struct {
	Pool         *PoolRef // nil below PoolMinPlayers
	PoolCount    int      // distinct players found, for the "not enough logs yet" line
	Sim          *SimRef  // nil unless the spec is validated DPS and the reference is in
	SimPreparing bool
	SimNoLevel   bool // a validated DPS spec whose sim band needs the level the log lacks
}

// PoolRef is the top logs' median profile over N players.
type PoolRef struct {
	N      int
	Median Profile
}

// SimRef is the sim reference at a level band: casts per minute by family and uptime by
// aura spell id.
type SimRef struct {
	Band        int
	CastsPerMin map[string]float64
	Uptime      map[int64]float64
}
