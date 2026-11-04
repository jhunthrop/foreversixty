package leveling

// Hand-set leveling orders (guide-accuracy round 1, owner 2026-11-04).
//
// data/curated/ladder-order.json holds ranks a spec takes ahead of the walk for the
// in-game order between checkpoints, where the sim cannot decide: the rule is that a row
// changes NO checkpoint build (the levels the sims run, ladderCheckpointLevels), so every
// band, ladder golden and published number stays the plain walk's (or ladder-first's), and
// only the levels between checkpoints move. Real leveling play decides those, researched
// on the web with Forever sources first (owner rule, 2026-11-04), and each row's reason
// says why in one line and cites them. A row that would move a checkpoint belongs in
// ladder-first.json, which needs a measured gain instead; TestLadderOrderMovesNoCheckpoint
// holds every row to that. The file shares ladder-first.json's schema and ForSpec spends
// it after the spec's ladder-first row.

import (
	_ "embed"
	"sync"
)

// ladderCheckpointLevels are the levels any sim runs a band build at: the rotation ladder
// (sim/request/ladder.go ladderLevels: 10, 20, 30, 38, 40, 50, 60), the leveling-bis and
// guide-strip bands (20 to 60 by ten) and dps-per-point's mid-band levels (bandLevel:
// 25, 35, 45, 55).
var ladderCheckpointLevels = []int{10, 20, 25, 30, 35, 38, 40, 45, 50, 55, 60}

// ladderOrderJSON is data/curated/ladder-order.json (`make ladder-first-sync`);
// TestEmbeddedLadderOrderMatchesCurated holds it to the curated file.
//
//go:embed ladder-order.json
var ladderOrderJSON []byte

var embeddedLadderOrder = sync.OnceValue(func() ladderFirstFile {
	f, err := parseLadderFirst(ladderOrderJSON)
	if err != nil {
		panic(err) // embedded data; TestEmbeddedLadderOrderMatchesCurated fails first
	}
	return f
})
