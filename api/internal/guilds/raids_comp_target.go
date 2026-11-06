// api/internal/guilds/raids_comp_target.go
//
// Comp targets (design/specs/2026-11-05-web-layout.md §7.1 "Comp against the target", D13):
// how many tanks, healers, melee and ranged a raid wants. Every 40-player raid zone has a
// default; an officer may set a raid's own target, and clearing it returns to the default.
// Dungeon runs, guild events and raids with no sourced size have no target (null).
package guilds

import (
	"context"
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// Comp target sources: the zone's default, or the officer's per-raid target.
const (
	CompSourceDefault = "default"
	CompSourceRaid    = "raid"
)

// RoleCounts is a count per roster group (raidspec.Group's four groups).
type RoleCounts struct {
	Tanks   int `json:"tanks"`
	Healers int `json:"healers"`
	Melee   int `json:"melee"`
	Ranged  int `json:"ranged"`
}

// add counts one row of a roster group (raidspec.Group); any group but the first three is ranged,
// as countSignups reads it.
func (c *RoleCounts) add(group string) {
	switch group {
	case raidspec.GroupTank:
		c.Tanks++
	case raidspec.GroupHealer:
		c.Healers++
	case raidspec.GroupMelee:
		c.Melee++
	default:
		c.Ranged++
	}
}

// total is the four groups summed.
func (c RoleCounts) total() int { return c.Tanks + c.Healers + c.Melee + c.Ranged }

// CompTarget is a raid's comp target and where it came from.
type CompTarget struct {
	RoleCounts
	Source string `json:"source"`
}

// compDefaults are the first tier's defaults a raid leader starts from (Forever's raids of 9
// December: Barrow Deeps 10, Hyjal Summit 20, Onyxia's Lair 40). Barrow Deeps: one tank, three
// healers, six damage. Hyjal Summit: two tanks, five healers, thirteen damage. Onyxia: two tanks
// (main tank plus the fear and whelp pickup), eight healers, and a ranged-heavy split for the air
// phase. An officer edits them per raid; a raid zone without a row here has no target.
var compDefaults = map[string]RoleCounts{
	"raid:barrow-deeps": {Tanks: 1, Healers: 3, Melee: 3, Ranged: 3},
	"raid:hyjal-summit": {Tanks: 2, Healers: 5, Melee: 6, Ranged: 7},
	"raid:onyxias-lair": {Tanks: 2, Healers: 8, Melee: 10, Ranged: 20},
}

// zoneCompDefault is the zone's default target; false for a zone with none (not a raid, a raid
// whose size is not sourced, or one with no row in compDefaults).
func zoneCompDefault(z Zone) (RoleCounts, bool) {
	if !z.IsRaid() || z.Size == nil {
		return RoleCounts{}, false
	}
	c, ok := compDefaults[z.Key]
	return c, ok
}

// compTargetOf is the raid's target: its own row when set, else the zone default; nil when the
// zone has no default (an officer cannot set one there either).
func compTargetOf(z Zone, override *RoleCounts) *CompTarget {
	def, ok := zoneCompDefault(z)
	if !ok {
		return nil
	}
	if override != nil {
		return &CompTarget{RoleCounts: *override, Source: CompSourceRaid}
	}
	return &CompTarget{RoleCounts: def, Source: CompSourceDefault}
}

// validateCompTarget checks an officer's target against the zone: every count 0 or more, and
// together no more than the zone's size (a leader may leave seats open, never plan past the
// cap). The fields map is nil when it passes.
func validateCompTarget(z Zone, c RoleCounts) map[string]string {
	fields := map[string]string{}
	for name, n := range map[string]int{"tanks": c.Tanks, "healers": c.Healers, "melee": c.Melee, "ranged": c.Ranged} {
		if n < 0 {
			fields[name] = "Use 0 or more."
		}
	}
	if len(fields) == 0 && z.Size != nil && c.total() > *z.Size {
		fields["total"] = fmt.Sprintf("That adds up to %d. %s takes %d.", c.total(), z.Name, *z.Size)
	}
	if len(fields) == 0 && c.total() == 0 {
		fields["total"] = "Ask for at least one raider."
	}
	if len(fields) == 0 {
		return nil
	}
	return fields
}

// compTargetsOf reads the per-raid targets of raidIDs.
func (s *Store) compTargetsOf(ctx context.Context, raidIDs []int64) (map[int64]RoleCounts, error) {
	out := map[int64]RoleCounts{}
	if len(raidIDs) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx, `select raid_id, tanks, healers, melee, ranged
		from guild_raid_comp_targets where raid_id = any($1)`, raidIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: comp targets: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var c RoleCounts
		if err := rows.Scan(&id, &c.Tanks, &c.Healers, &c.Melee, &c.Ranged); err != nil {
			return nil, fmt.Errorf("guilds: comp targets: %w", err)
		}
		out[id] = c
	}
	return out, rows.Err()
}

// SetCompTarget stores raidID's own target.
func (s *Store) SetCompTarget(ctx context.Context, raidID, by int64, c RoleCounts) error {
	_, err := s.Pool.Exec(ctx, `
		insert into guild_raid_comp_targets (raid_id, tanks, healers, melee, ranged, updated_by)
		values ($1, $2, $3, $4, $5, $6)
		on conflict (raid_id) do update set tanks = excluded.tanks, healers = excluded.healers,
		  melee = excluded.melee, ranged = excluded.ranged, updated_by = excluded.updated_by,
		  updated_at = now()`, raidID, c.Tanks, c.Healers, c.Melee, c.Ranged, by)
	if err != nil {
		return fmt.Errorf("guilds: set comp target: %w", err)
	}
	return nil
}

// ClearCompTarget removes raidID's own target, so it reads the zone default again.
func (s *Store) ClearCompTarget(ctx context.Context, raidID int64) error {
	if _, err := s.Pool.Exec(ctx, `delete from guild_raid_comp_targets where raid_id = $1`, raidID); err != nil {
		return fmt.Errorf("guilds: clear comp target: %w", err)
	}
	return nil
}
