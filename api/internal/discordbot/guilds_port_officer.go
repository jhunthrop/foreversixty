package discordbot

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/guilds"
)

// OfficerReads is what /report, /attendance and /readiness ask of the guilds package: the
// service functions the raid page, the roster and the readiness board answer through.
// GuildService embeds it.
type OfficerReads interface {
	// ReportFor is a done raid's report as the viewer may see it (0: a signed-out visitor).
	ReportFor(ctx context.Context, guildID, viewerID int64, raidID *int64) (guilds.ReportAnswer, bool, error)
	// DoneRaids is the guild's logged raids, newest first; publicOnly keeps public logs.
	DoneRaids(ctx context.Context, guildID int64, publicOnly bool, limit int) ([]guilds.Raid, error)
	// AttendanceFor is the viewer's own records, or the named character's.
	AttendanceFor(ctx context.Context, guildID, viewerID int64, name string) ([]guilds.AttendanceRecord, error)
	// LowestAttendance is the officer's list of members who came least, lowest first; others get ErrOfficerOnly.
	LowestAttendance(ctx context.Context, guildID, viewerID int64, limit int) ([]guilds.AttendanceRecord, error)
	// AttendanceNames is the roster's names for an officer; empty for anyone else.
	AttendanceNames(ctx context.Context, guildID, viewerID int64, typed string, limit int) ([]string, error)
	// NextRaidReadiness is the next raid's readiness board for an officer.
	NextRaidReadiness(ctx context.Context, guildID, viewerID int64, now time.Time) (guilds.ReadinessAnswer, error)
	// GuildPagePathOf is the guild page's site path.
	GuildPagePathOf(ctx context.Context, guildID int64) (string, error)
}
