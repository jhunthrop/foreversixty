// api/internal/guilds/raids_activity.go
package guilds

import "context"

// RaidActivity keeps a raid's row in the site feed (design/specs/2026-11-07-home-b.md, P3).
type RaidActivity interface {
	RaidChanged(ctx context.Context, raidID int64)
}

// raidActivity tells the feed a raid was posted or its sign-ups moved.
func (s *Service) raidActivity(ctx context.Context, raidID int64) {
	if s.Activity != nil {
		s.Activity.RaidChanged(ctx, raidID)
	}
}
