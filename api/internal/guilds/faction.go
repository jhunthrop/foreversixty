// api/internal/guilds/faction.go
//
// A guild's faction (design/specs/2026-10-04-guild-page.md §12.2): stored on
// guilds.faction (migration 0031) so the guild page header can paint the faction emblem
// and watermark from one read, rather than decoding every roster character's FS1 export
// per request.
package guilds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jhunthrop/foreversixty/api/internal/fs1"
)

// Faction values guilds.faction carries - matches migration 0031's own check constraint.
const (
	FactionAlliance = "alliance"
	FactionHorde    = "horde"
)

// raceFactions maps an FS1 export's race slug to its faction, for exactly the races whose
// faction this feature counts. Mirrors data/builds/1.60.1.70009/races.json's own slug/
// faction fields - the same build this package's one real BiS catalogue already hardcodes
// (bis_lookup.go's bisDataBuild) - faction_test.go cross-checks this table against that
// file directly, so a future race/faction reassignment there fails a test here rather
// than silently drifting.
//
// Skyborne's two faction-specific slugs (high-order-skyborne, windshaper-skyborne) are
// real Alliance/Horde races and are counted here. The generic "skyborne" token the beta
// client itself reports is neutral and never reaches a real export - addon/ForeverSixty/
// Export.lua's own SKYBORNE_BY_FACTION always resolves it to one of the two specific
// slugs first - so "skyborne" alone, like any slug not in this table, simply counts as an
// unknown race.
var raceFactions = map[string]string{
	"human":               FactionAlliance,
	"dwarf":               FactionAlliance,
	"night-elf":           FactionAlliance,
	"gnome":               FactionAlliance,
	"high-order-skyborne": FactionAlliance,
	"orc":                 FactionHorde,
	"undead":              FactionHorde,
	"tauren":              FactionHorde,
	"troll":               FactionHorde,
	"windshaper-skyborne": FactionHorde,
}

// factionFromRaceSlugs is the pure majority rule RecomputeFaction applies: the strict
// majority faction among slugs' known races (alliance or horde; a neutral or
// unrecognised slug - including an empty one, from an export fs1.Decode could not read -
// is simply not counted), or ok=false on a tie, on an all-neutral/unknown roster, or on
// an empty slice (no exports at all).
func factionFromRaceSlugs(slugs []string) (faction string, ok bool) {
	var alliance, horde int
	for _, slug := range slugs {
		switch raceFactions[slug] {
		case FactionAlliance:
			alliance++
		case FactionHorde:
			horde++
		}
	}
	switch {
	case alliance == 0 && horde == 0:
		return "", false
	case alliance > horde:
		return FactionAlliance, true
	case horde > alliance:
		return FactionHorde, true
	default:
		return "", false // tie
	}
}

// queryExecer is the subset of *pgxpool.Pool and pgx.Tx RecomputeFaction needs - both
// satisfy it structurally, so the same function runs inside a membership change's own
// transaction (every trigger point in membership.go/roster.go/invite.go) and standalone
// against the pool (MembershipJob's backfill sweep, jobs.go, which runs with no
// surrounding transaction of its own).
type queryExecer interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

// RecomputeFaction derives guildID's stored faction from the majority faction among its
// guild_characters rows whose addon export decodes to a race this package counts -
// ignoring a row with no addon_exports match, an export that does not decode as FS1
// (fs1.Decode's own ok=false), and a race that is neutral or unrecognised. A tie, or zero
// known rows, clears the column to null rather than guessing - the header then simply
// paints no emblem.
func RecomputeFaction(ctx context.Context, db queryExecer, guildID int64) error {
	rows, err := db.Query(ctx,
		`select ae.export from guild_characters gc
		   join addon_exports ae on ae.character_key = gc.character_key
		 where gc.guild_id = $1`, guildID)
	if err != nil {
		return fmt.Errorf("guilds: recompute faction: read guild %d: %w", guildID, err)
	}
	var exports []string
	for rows.Next() {
		var export string
		if err := rows.Scan(&export); err != nil {
			rows.Close()
			return fmt.Errorf("guilds: recompute faction: scan guild %d: %w", guildID, err)
		}
		exports = append(exports, export)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return fmt.Errorf("guilds: recompute faction: read guild %d: %w", guildID, err)
	}

	slugs := make([]string, 0, len(exports))
	for _, export := range exports {
		if decoded, ok := fs1.Decode(export); ok {
			slugs = append(slugs, decoded.RaceSlug)
		}
	}

	faction, ok := factionFromRaceSlugs(slugs)
	var value *string
	if ok {
		value = &faction
	}
	if _, err := db.Exec(ctx,
		`update guilds set faction = $2, faction_updated_at = now() where id = $1`, guildID, value); err != nil {
		return fmt.Errorf("guilds: recompute faction: update guild %d: %w", guildID, err)
	}
	return nil
}
