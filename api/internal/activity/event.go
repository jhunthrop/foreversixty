package activity

import (
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/events"
)

// The feed's event kinds (GET /v1/feed rows[].kind).
const (
	KindLogUpload    = "log_upload"
	KindBossKill     = "boss_kill"
	KindGuildRaid    = "guild_raid"
	KindRankingTop   = "ranking_top"
	KindRankingClimb = "ranking_climb"
	KindBisChange    = "bis_change"
	KindGuildSim     = "guild_sim"
)

// The tags a row carries for its viewer (rows[].tag). A row's kind decides its tag (KindTag);
// a guild row is always TagYourGuild, and TagYourSpec replaces the kind's tag on a row of the
// viewer's spec and band.
const (
	TagYourSpec   = "your_spec"
	TagLogs       = "logs"
	TagRankings   = "rankings"
	TagBestInSlot = "best_in_slot"
	TagYourGuild  = "your_guild"
	TagPatch      = "patch"
)

// KindTag is the one map from an event kind to the tag it shows anyone it is not personal to.
var KindTag = map[string]string{
	KindLogUpload:    TagLogs,
	KindBossKill:     TagLogs,
	KindGuildRaid:    TagYourGuild,
	KindRankingTop:   TagRankings,
	KindRankingClimb: TagRankings,
	KindBisChange:    TagBestInSlot,
	KindGuildSim:     TagYourGuild,
}

// The segment styles (rows[].segments[].style), the events package's: the feed and the
// Activity inbox share one segment builder.
const (
	StyleStrong = events.StyleStrong
	StyleName   = events.StyleName
	StyleItem   = events.StyleItem
)

// Segment is one run of a row's sentence (events.Segment).
type Segment = events.Segment

// The segment builders, the events package's.
var (
	Text     = events.Text
	Segmentf = events.Segmentf
	Strong   = events.Strong
	Name     = events.Name
	Item     = events.Item
)

// Event is one feed row as it is written.
type Event struct {
	Kind         string
	OccurredAt   time.Time
	Spec         string // "warrior-fury", or "" for a row of no spec
	Band         int    // 20..60, or 0
	GuildID      int64  // set: the guild's members only
	CharacterKey string
	Icon         string
	Crest        string
	Quality      *int
	Href         string
	DateOnly     bool
	Segments     []Segment
	DedupeKey    string
	// ReportID is the report that wrote the row, if one did: the feed shows the row only while
	// that report is public evidence (feed.go's reportVisible).
	ReportID string
}
