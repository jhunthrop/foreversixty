// api/internal/addon/raid_assignments.go
//
// The raid_assignment inbox message (docs/contracts/addon-v2.md contract 3): one per slot a
// character holds in its next raid (guilds.Store.NextRaidAssignments), so the addon tells the
// player their own job. The message is flat (strings and numbers only), so the companion
// carries it through as a type it does not know; With is the slot's other names, comma joined.
package addon

import (
	"context"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

const messageRaidAssignment = "raid_assignment"

// AssignmentSource is enough of guilds.Store to build the raid_assignment messages.
type AssignmentSource interface {
	NextRaidAssignments(ctx context.Context, keys []string, now time.Time) ([]guilds.CharacterAssignment, error)
}

// raidAssignmentMessages is one message per slot each of keys holds in its next raid; none on
// a failed read, which is logged (the sync never fails over it).
func (s *Service) raidAssignmentMessages(ctx context.Context, keys []string, now time.Time) []MessageEntry {
	if s.Assignments == nil || len(keys) == 0 {
		return nil
	}
	held, err := s.Assignments.NextRaidAssignments(ctx, keys, now)
	if err != nil {
		s.logger().Warn("addon", "op", "messages_raid_assignment", "err", err)
		return nil
	}
	out := make([]MessageEntry, 0, len(held))
	for _, a := range held {
		out = append(out, MessageEntry{
			Type: messageRaidAssignment, Character: a.CharacterKey, RaidID: a.RaidID, Raid: a.ZoneName,
			Title: a.RaidTitle, StartsAt: a.StartsAt.UTC().Format(time.RFC3339), Encounter: a.Encounter,
			Slot: a.SlotKey, Label: a.Label, Warning: a.Warning, Purpose: a.Purpose, With: strings.Join(a.With, ", "),
		})
	}
	return out
}
