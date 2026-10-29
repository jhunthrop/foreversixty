package guilds

import "context"

// IsOfficer reports whether viewerID is a verified officer or leader of the guild: the rank
// rule behind every officer read, without loading any of them.
func (s *Service) IsOfficer(ctx context.Context, guildID, viewerID int64) (bool, error) {
	return s.seesEveryAttendanceRow(ctx, guildID, viewerID)
}
