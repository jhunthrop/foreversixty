package activity

import "time"

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

// The tags a row carries for its viewer (rows[].tag). A row's stored base tag is what it
// shows to anyone whose spec it does not match; TagYourSpec replaces it on a match.
const (
	TagYourSpec  = "your_spec"
	TagLogs      = "logs"
	TagGuides    = "guides"
	TagYourGuild = "your_guild"
	TagPatch     = "patch"
)

// The segment styles (rows[].segments[].style). Plain text has none.
const (
	StyleStrong = "strong"
	StyleName   = "name"
	StyleItem   = "item"
)

// Segment is one run of a row's sentence.
type Segment struct {
	Text    string `json:"text"`
	Style   string `json:"style,omitempty"`
	Class   string `json:"class,omitempty"`
	Quality *int   `json:"quality,omitempty"`
	ItemID  *int   `json:"item_id,omitempty"`
}

// Event is one feed row as it is written.
type Event struct {
	Kind         string
	OccurredAt   time.Time
	Spec         string // "warrior-fury", or "" for a row of no spec
	Band         int    // 20..60, or 0
	GuildID      int64  // set: the guild's members only
	CharacterKey string
	TagBase      string
	Icon         string
	Crest        string
	Quality      *int
	Href         string
	DateOnly     bool
	Segments     []Segment
	DedupeKey    string
}

// Text is a plain segment.
func Text(s string) Segment { return Segment{Text: s} }

// Strong is a bold segment.
func Strong(s string) Segment { return Segment{Text: s, Style: StyleStrong} }

// Name is a character's name in its class colour.
func Name(s, class string) Segment { return Segment{Text: s, Style: StyleName, Class: class} }

// Item is an item's name in its quality colour.
func Item(name string, id, quality int) Segment {
	return Segment{Text: name, Style: StyleItem, ItemID: &id, Quality: &quality}
}
