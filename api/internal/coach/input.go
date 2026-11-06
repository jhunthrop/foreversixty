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
	Rules  SpecRules            // data/curated/coach/<spec>.json; zero value = no file
	Spells map[int64]SpellConst // the class's spellconst for the fight's build
	// BaseMana is the class's base mana at the row's level (the build's levels.json), what a
	// cost given as a percent of base mana is read against; 0 is unknown.
	BaseMana  int64
	Trainable map[string][]TrainableRank // family -> ranks with the level each is learned
	// RankChoices is family -> the ranks data/curated/rank-choices.json has the spec's
	// rotation cast at the player's level (a mana-bound band's downrank); F5 never calls
	// those low. Nil where no row holds the level.
	RankChoices map[string][]int
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
// uses. CostKnown is false when the data states the cost only as a share of the bar:
// CostPct, the percent of base mana it costs (Judgement's 6), which CostAt reads.
type SpellConst struct {
	Name       string
	Rank       int
	CooldownMS int64
	GCDMS      int64
	Cost       int64
	CostType   int64
	CostKnown  bool
	CostPct    float64
	// MeleeOnly is true when every effect of the spell raises dodge or parry (Evasion,
	// Deterrence): it can stop a melee swing or a physical attack, never a spell.
	MeleeOnly bool
}

// CostAt is the spell's cost for a player with baseMana: the flat cost when the data
// states one, else CostPct of baseMana for a mana spell (rounded down, as the client
// does). false when neither is known: no flat cost and no base mana (0) to read a percent
// against.
func (sc SpellConst) CostAt(baseMana int64) (int64, bool) {
	if sc.CostKnown {
		return sc.Cost, true
	}
	if sc.CostPct > 0 && sc.CostType == PowerMana && baseMana > 0 {
		return int64(sc.CostPct * float64(baseMana) / 100), true
	}
	return 0, false
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
