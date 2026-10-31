// Package reports owns a report's life: created by a companion or a
// whole-file upload, filled in fight by fight, read back by the report
// page, and governed by its visibility.
package reports

import (
	"encoding/json"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// The visibilities, exactly as the design names them.
const (
	Public   = "public"
	Unlisted = "unlisted"
	Private  = "private"
	GuildTo  = "guild"
)

// The statuses a report moves through.
const (
	StatusLive       = "live"
	StatusProcessing = "processing"
	StatusComplete   = "complete"
	StatusFailed     = "failed"
)

// Visibilities is every valid visibility.
var Visibilities = []string{Public, Unlisted, Private, GuildTo}

// ValidVisibility reports whether v is one of them.
func ValidVisibility(v string) bool {
	for _, s := range Visibilities {
		if s == v {
			return true
		}
	}
	return false
}

// Ranked reports whether a report's fights may appear in rankings. A
// private report never does; the design says so plainly.
func Ranked(visibility string) bool { return visibility != Private }

// Report is one row of the reports table.
type Report struct {
	ID               string
	OwnerID          *int64
	GuildID          *int64
	Title            string
	Visibility       string
	Zone             string
	Status           string
	EngineVersion    string
	UploadID         *string
	LoggingCharacter *string
	Health           json.RawMessage
	Flagged          *string
	CreatedAt        time.Time
	CompletedAt      *time.Time
}

// FightEntry is a fight as the report page reads it: the engine's own
// report.json line plus whether the ingest verified it.
type FightEntry struct {
	store.FightEntry
	Verified bool `json:"verified"`
}

// FightRecord is one fight as the ingest writes it.
type FightRecord struct {
	ReportID    string
	Index       int
	EncounterID *int64
	Name        string
	Difficulty  *int64
	Size        *int64
	Kill        bool
	DurationMS  int64
	StartMS     int64
	Verified    bool
	Players     []string
	Deaths      int
	// The engine 0.6.0 counts (fight.Counter): hostile NPCs that died, the
	// ones the group killed, the group's own deaths, and the group.
	MobDeaths   int
	GroupKills  int
	GroupDeaths int
	Group       []string
	// BossHealthPct is what a wipe got the boss down to; nil on a kill
	// and on trash, -1 when the log never showed the boss.
	BossHealthPct *float64
	RawStart      *int64
	RawEnd        *int64
	RawSHA256     []byte
}

// Owner is the account a report belongs to, as the page shows it. The
// field is named for the contract's `owner: { id, battletag }`, but
// what it carries is auth.User.PublicName: the battletag when there is
// one to show, and a pseudonym otherwise.
type Owner struct {
	ID        int64  `json:"id"`
	Battletag string `json:"battletag,omitempty"`
}

// GuildRef is the guild a report belongs to.
type GuildRef struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
}

// LoggingCharacter is the character a report was logged by: its key (region/ruleset/slug) and
// display name.
type LoggingCharacter struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

// View is the body of GET /v1/reports/{id}.
type View struct {
	ID            string       `json:"id"`
	Title         string       `json:"title"`
	Visibility    string       `json:"visibility"`
	Owner         *Owner       `json:"owner"`
	Guild         *GuildRef    `json:"guild,omitempty"`
	Zone          string       `json:"zone"`
	Status        string       `json:"status"`
	EngineVersion string       `json:"engine_version"`
	Flagged       string       `json:"flagged,omitempty"`
	Fights        []FightEntry `json:"fights"`
	Players       []string     `json:"players"`
	CreatedAt     time.Time    `json:"created_at"`
	// LastUploadAt is the newest of the report's fight uploads and its creation; HeartbeatAt is the
	// companion's last "still tailing" ping, null until the first. Both are what a live report's
	// waiting screen shows instead of "Logging" forever.
	LastUploadAt *time.Time `json:"last_upload_at"`
	HeartbeatAt  *time.Time `json:"heartbeat_at"`
	DataBaseURL  string     `json:"data_base_url"`
	// LoggingCharacter is the character whose client wrote the log; null when the report names
	// none.
	LoggingCharacter *LoggingCharacter `json:"logging_character"`
}
