// logs/engine/mechanics/weights/roles.go
// Package weights holds the default per-role component weight table for the
// performance rating engine (docs/superpowers/specs/2026-09-21-performance-
// rating-design.md §1.4), as data rather than code: a guild's own officer
// tooling (out of scope of this spec) edits the same shape this package
// embeds, so that feature edits data, not code.
package weights

import (
	"embed"
	"encoding/json"
	"fmt"
)

// RoleWeights is one role's default weight per component, each 0-100,
// summing to 100.
type RoleWeights struct {
	Output      float64 `json:"output"`
	Survival    float64 `json:"survival"`
	Mechanics   float64 `json:"mechanics"`
	Utility     float64 `json:"utility"`
	Preparation float64 `json:"preparation"`
	Activity    float64 `json:"activity"`
}

// Roles is the whole default weight table, one row per role.
type Roles struct {
	DPS    RoleWeights `json:"dps"`
	Healer RoleWeights `json:"healer"`
	Tank   RoleWeights `json:"tank"`
}

// For selects a role's weights by the same role strings
// logs/engine/summary.RosterRow.Role already produces ("dps", "healer",
// "tank"). An unrecognised role reports false and DPS's weights, which
// never happens in practice since roster.go's role() only ever emits those
// three, but a caller over an untrusted role string should still fail
// explicitly rather than silently mis-weight a card.
func (r Roles) For(role string) (RoleWeights, bool) {
	switch role {
	case "dps":
		return r.DPS, true
	case "healer":
		return r.Healer, true
	case "tank":
		return r.Tank, true
	default:
		return r.DPS, false
	}
}

// Sum is the role's total weight, for validation and for §1.1's
// renormalisation (the weight redistributed across scored components is
// proportional to each one's own share of this total).
func (w RoleWeights) Sum() float64 {
	return w.Output + w.Survival + w.Mechanics + w.Utility + w.Preparation + w.Activity
}

//go:embed roles.json
var rolesFile embed.FS

// A malformed embedded table is a build defect, so it fails the process at
// start rather than the first time a card is scored, matching
// logs/engine/mechanics's own mustParseAll pattern.
var defaultRoles = mustLoad()

func mustLoad() Roles {
	data, err := rolesFile.ReadFile("roles.json")
	if err != nil {
		panic(fmt.Errorf("weights: reading roles.json: %w", err))
	}
	r, err := Parse(data)
	if err != nil {
		panic(fmt.Errorf("weights: roles.json: %w", err))
	}
	return r
}

// weightTolerance absorbs JSON float round-tripping; the curated table's
// own numbers are whole percentages, so any real typo is off by whole
// points, not fractions of one.
const weightTolerance = 0.01

// Parse reads a role weight table and refuses one whose weights do not sum
// to 100 for every role.
func Parse(data []byte) (Roles, error) {
	var r Roles
	if err := json.Unmarshal(data, &r); err != nil {
		return Roles{}, fmt.Errorf("weights: %w", err)
	}
	for role, w := range map[string]RoleWeights{"dps": r.DPS, "healer": r.Healer, "tank": r.Tank} {
		if sum := w.Sum(); sum < 100-weightTolerance || sum > 100+weightTolerance {
			return Roles{}, fmt.Errorf("weights: role %q weights sum to %v, want 100", role, sum)
		}
	}
	return r, nil
}

// Default returns the site-wide default weight table (§1.4).
func Default() Roles {
	return defaultRoles
}
