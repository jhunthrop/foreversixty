// api/internal/guilds/raid_assign_addon.go
//
// A character's own assignments for its next raid, for the addon inbox's raid_assignment
// messages (api/internal/addon, docs/contracts/addon-v2.md contract 3): the earliest open raid
// (planned or locked, not yet over) in which the character holds a slot and is still signed
// In, Late or Bench.
package guilds

import (
	"context"
	"fmt"
	"time"
)

// CharacterAssignment is one slot a character holds in its next raid. With is the slot's
// other names, in the raid leader's order.
type CharacterAssignment struct {
	CharacterKey string
	RaidID       int64
	RaidTitle    string
	ZoneName     string
	StartsAt     time.Time
	Encounter    string
	SlotKey      string
	Label        string
	Warning      string
	Purpose      string
	With         []string
}

// heldSlot is a character's next raid with an assignment in it.
type heldSlot struct {
	key, title, zone string
	raidID           int64
	startsAt         time.Time
}

// NextRaidAssignments is every slot each of keys holds in its next raid, by character then the
// template's order.
func (s *Store) NextRaidAssignments(ctx context.Context, keys []string, now time.Time) ([]CharacterAssignment, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	held, err := s.nextHeldSlots(ctx, keys, now)
	if err != nil || len(held) == 0 {
		return nil, err
	}
	ids := make([]int64, 0, len(held))
	for _, h := range held {
		ids = append(ids, h.raidID)
	}
	stored, err := s.assignmentsByRaid(ctx, ids)
	if err != nil {
		return nil, err
	}
	signups, err := s.SignupsByRaid(ctx, ids)
	if err != nil {
		return nil, err
	}
	out := []CharacterAssignment{}
	for _, h := range held {
		out = append(out, heldAssignments(h, stored[h.raidID], signups[h.raidID])...)
	}
	return out, nil
}

// nextHeldSlots reads the characters' held slots in each one's earliest open raid.
func (s *Store) nextHeldSlots(ctx context.Context, keys []string, now time.Time) ([]heldSlot, error) {
	rows, err := s.Pool.Query(ctx, `
		with held as (
		  select a.character_key, a.raid_id, a.encounter_id, a.slot_key, r.title, r.zone, r.starts_at,
		    min(r.starts_at) over (partition by a.character_key) as first_start
		  from raid_assignments a
		  join guild_raids r on r.id = a.raid_id
		  join guild_raid_signups sg on sg.raid_id = a.raid_id and sg.character_key = a.character_key
		  where a.character_key = any($1) and r.status in ('planned', 'locked')
		    and r.starts_at + make_interval(mins => r.duration_minutes) > $2
		    and sg.state in ('in', 'late', 'bench'))
		select distinct on (character_key) character_key, raid_id, title, zone, starts_at
		from held where starts_at = first_start order by character_key, raid_id`, keys, now)
	if err != nil {
		return nil, fmt.Errorf("guilds: next held slots: %w", err)
	}
	defer rows.Close()
	var out []heldSlot
	for rows.Next() {
		var h heldSlot
		if err := rows.Scan(&h.key, &h.raidID, &h.title, &h.zone, &h.startsAt); err != nil {
			return nil, fmt.Errorf("guilds: next held slots: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}

// heldAssignments is every slot of the raid's encounters that h's character holds, in the
// template's order, each with the slot's other assignable names.
func heldAssignments(h heldSlot, stored map[slotRef][]string, signups []Signup) []CharacterAssignment {
	names := make(map[string]string, len(signups))
	for _, sg := range signups {
		if assignable(sg.State) {
			names[sg.CharacterKey] = sg.Name
		}
	}
	zone := zoneOf(h.zone)
	raid, ok := tierRaidOf(zone.Key)
	if !ok {
		return nil
	}
	var out []CharacterAssignment
	for _, boss := range raid.Bosses {
		slots, _ := templateOf(boss.EncounterID)
		for _, t := range slots {
			keys := stored[slotRef{EncounterID: boss.EncounterID, SlotKey: t.Key}]
			if !contains(keys, h.key) {
				continue
			}
			out = append(out, CharacterAssignment{
				CharacterKey: h.key, RaidID: h.raidID, RaidTitle: h.title, ZoneName: zone.Name,
				StartsAt: h.startsAt.UTC(), Encounter: boss.Name, SlotKey: t.Key, Label: t.Label,
				Warning: t.Warning, Purpose: t.Purpose, With: otherNames(keys, h.key, names),
			})
		}
	}
	return out
}

// otherNames is the slot's assignable names other than self's, in order.
func otherNames(keys []string, self string, names map[string]string) []string {
	out := []string{}
	for _, k := range keys {
		if n, ok := names[k]; ok && k != self {
			out = append(out, n)
		}
	}
	return out
}

func contains(keys []string, key string) bool {
	for _, k := range keys {
		if k == key {
			return true
		}
	}
	return false
}
