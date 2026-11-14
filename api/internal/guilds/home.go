// api/internal/guilds/home.go
package guilds

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/reportguild"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
	"github.com/jhunthrop/foreversixty/api/internal/upgrades"
)

// HomeReportsPerPage is the page size of the guild home's "this week's
// reports" list, keyset-paginated the same shape as GET /v1/reports/recent.
const HomeReportsPerPage = 20

// homeReportsWindow is "this week": a trailing 7-day window, not a
// server-specific weekly-reset timestamp, since no reset concept exists
// anywhere in this codebase to anchor to.
const homeReportsWindow = 7 * 24 * time.Hour

const homeCursorSep = "|"

// HomeCursor is the keyset position GET /v1/guilds/{id}/home pages
// reports from.
type HomeCursor struct {
	CreatedAt time.Time
	ID        string
}

func encodeHomeCursor(createdAt time.Time, id string) string {
	raw := createdAt.UTC().Format(time.RFC3339Nano) + homeCursorSep + id
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func decodeHomeCursor(s string) (HomeCursor, bool) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return HomeCursor{}, false
	}
	createdAt, id, ok := strings.Cut(string(raw), homeCursorSep)
	if !ok || id == "" {
		return HomeCursor{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return HomeCursor{}, false
	}
	return HomeCursor{CreatedAt: t, ID: id}, true
}

type GuildIdentity struct {
	ID      int64  `json:"id"`
	Region  string `json:"region"`
	Ruleset string `json:"ruleset"`
	Name    string `json:"name"`
	// Faction is guilds.RecomputeFaction's last computed result - "alliance", "horde",
	// or null when it has never resolved to either, so the header simply paints no
	// emblem (design/specs/2026-10-04-guild-page.md §12.2).
	Faction *string `json:"faction"`
	// CrestURL is this guild's uploaded crest (docs/contracts/2026-10-05-guild-crest-api.md),
	// or null when none is set - the header then falls back to the faction logo.
	CrestURL *string `json:"crest_url"`
}

type HomeReport struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	// Zone lets an untitled report still show something identifying in
	// the list, the same way every other report list in this codebase
	// already does (reports.View, rankings' guild report list) - added
	// here so the guild home matches that convention (item 7, fourth
	// security review response).
	Zone       string    `json:"zone"`
	CreatedAt  time.Time `json:"created_at"`
	FightCount int       `json:"fight_count"`
	KillCount  int       `json:"kill_count"`
	// DurationMS is the night's own elapsed wall-clock span - the last fight's end minus
	// the first fight's start, 0 for a report with no fights yet.
	DurationMS int64 `json:"duration_ms"`
}

// Attendance is a character's presence over the guild's last 8 non-private, unflagged
// reports that someone is present in (attendanceReports; the contract's own window), design
// spec §4.B. Nights is the window's size, the same on every roster row.
type Attendance struct {
	Present int `json:"present"`
	Nights  int `json:"nights"`
}

// BestParse is a roster row's single best fight_metrics row - contract's own shape for
// GET .../home's roster[].best_parse, distinct from progression's ParseRef (progression.go):
// this one identifies the fight, not the character (the roster row it sits on already does
// that).
type BestParse struct {
	Metric     string   `json:"metric"`
	Value      float64  `json:"value"`
	Percentile *float64 `json:"percentile"`
	Encounter  string   `json:"encounter"`
	ReportID   string   `json:"report_id"`
	FightIndex int      `json:"fight_index"`
}

// RatingView is a roster row's rating_scores overall plus its six components - contract's
// own shape, null when the character has no rated fight.
type RatingView struct {
	Overall     float64 `json:"overall"`
	Output      float64 `json:"output"`
	Survival    float64 `json:"survival"`
	Mechanics   float64 `json:"mechanics"`
	Utility     float64 `json:"utility"`
	Preparation float64 `json:"preparation"`
	Activity    float64 `json:"activity"`
}

// RosterRow is one guild_characters row as the home shows it. Class,
// Spec and ItemLevel are nil whenever the account's consent withholds
// them — the query itself never selects those columns for a row below
// the consent level that would show them (see HomeRoster's SQL), so
// there is no Go-side filtering step to forget.
type RosterRow struct {
	CharacterKey   string  `json:"character_key"`
	Region         string  `json:"region"`
	Ruleset        string  `json:"ruleset"`
	Name           string  `json:"name"`
	Rank           string  `json:"rank"`
	Verified       bool    `json:"verified"`
	LoggedRecently bool    `json:"logged_recently"`
	Consent        string  `json:"consent"`
	Class          *string `json:"class,omitempty"`
	Spec           *string `json:"spec,omitempty"`
	Role           *string `json:"role,omitempty"`
	// SpecName ("Holy Paladin") and SpecIcon are the catalogue's words for Spec
	// (describeRosterRows), omitted with it.
	SpecName *string `json:"spec_name,omitempty"`
	SpecIcon *string `json:"spec_icon,omitempty"`
	// Group is the raid group the row fights in (tank, healer, melee or ranged; raidspec.Group,
	// the rule the raid page's sign-ups use), omitted with the spec (describeRosterRows).
	Group     *string `json:"group,omitempty"`
	ItemLevel *int    `json:"item_level,omitempty"`
	// Level is the character's level: its export's level= section, else the guild's newest
	// addon roster snapshot's line (homeAddonSnapshot); null when neither carries one. Both are
	// read only at gear consent.
	Level *int `json:"level"`
	// LoggedAt is addon_exports.updated_at - the contract's own replacement for
	// logged_recently (kept below too, unchanged, per the contract's own instruction).
	LoggedAt   time.Time  `json:"logged_at"`
	Attendance Attendance `json:"attendance"`
	// Attendance30 is the same figure over the guild's last 30 logged raids; raid mode only.
	Attendance30 *Attendance `json:"attendance_30,omitempty"`
	BestParse    *BestParse  `json:"best_parse,omitempty"`
	Rating       *RatingView `json:"rating,omitempty"`
	Professions  []string    `json:"professions,omitempty"`
	AccountKey   string      `json:"account_key"`
	// Supporter is true when the row's account shows the supporter mark (Service.markSupporters).
	Supporter bool `json:"supporter"`
	// DiscordPickedBy is the Discord name of the Discord-only account that picked this character
	// from the guild's directory roster on Discord, until a Battle.net sign-in confirms it.
	DiscordPickedBy string     `json:"discord_picked_by,omitempty"`
	LastReportAt    *time.Time `json:"last_report_at"`
	// MayRemove is computed server-side from the exact same
	// mayRemoveRow rule the DELETE .../characters/{key} route enforces
	// (item 7, fourth security review response), from the viewpoint of
	// whoever is asking for this home - so the web can show the remove
	// control exactly where it would actually succeed, never a control
	// that then answers 403.
	MayRemove bool `json:"may_remove"`
	// MayApprove is true exactly when the viewer is a verified officer/leader and this
	// row is not already verified - the Roster tab's own Approve control (contract's
	// own field, mirroring MayRemove's convention).
	MayApprove bool `json:"may_approve"`
	// IsMain is true for the character its account marked as main (users.main_character_key):
	// the Roster tab's "alt of" label names that one.
	IsMain bool `json:"is_main"`
	// SimExport is true when the character's sim input is its FS1 export (sims.ExportIsInputSQL),
	// so the Roster tab's Open in simulator / planner slot has links to offer; false below gear
	// consent. The web decides from it at first paint whether the row holds that slot at all.
	SimExport bool `json:"sim_export"`
	// Addon is this character's line in the guild's newest addon roster snapshot
	// (AddonMemberView), matched by name; null when the guild has no snapshot or the
	// snapshot does not carry this character. Never omitted.
	Addon *AddonMemberView `json:"addon"`
	// ownerUserID is gc.user_id: read to compute MayRemove/MayApprove/AccountKey against
	// the viewing actor, never marshaled into the response itself.
	ownerUserID int64
	// awardKey is the guild_characters key a loot award has to name when it differs from
	// CharacterKey: an invited member's picked character is ranked under its own key but awarded
	// under the account's synthetic row (invitedRosterRows). Empty everywhere else.
	awardKey string
	// faction/className/specName/export/gear/enchants/bags/talentPointsSpent/hasTalents
	// back this row's own standing/readiness computation (standing.go) without a second
	// query - never marshaled.
	faction            string
	className, specStr string
	gear, enchants     map[string]int
	bags               []int
	hasBagsSection     bool
	// prep is the export's raid-prep sections (raid_prep.go).
	prep              exportPrep
	talentPointsSpent int
	hasTalents        bool
}

// AddonMemberView is one roster row's member line from the guild's newest roster
// snapshot (guild_roster_snapshots, written by POST /v1/addon/exports' outbox, contract 4 of
// docs/contracts/addon-v2.md): what a guildmate's addon saw of this character on its guild
// tab. Level, Spec (a slug), GearGap and Build are as the member's addon announced them, and
// all four are null unless the row's member gave gear consent (the home withholds spec and
// gear for a roster-consent member, and the addon's copy follows);
// Build is null when it announced none. SeenAt is the sender's local "YYYY-MM-DD HH:MM" as
// stored; SnapshotAt is when the site received the snapshot, RFC 3339 UTC.
type AddonMemberView struct {
	Level      *int      `json:"level"`
	Spec       *string   `json:"spec"`
	GearGap    *int      `json:"gear_gap"`
	Build      *string   `json:"build"`
	SeenAt     string    `json:"seen_at"`
	SnapshotAt time.Time `json:"snapshot_at"`
}

// AddonSnapshotView is the home's top-level account of the guild's newest roster snapshot:
// when it was received (RFC 3339 UTC), the character key whose export carried it, and how
// many members it listed (unmatched ones included; they become the roster's no-account entries, HomeView.Members).
type AddonSnapshotView struct {
	ReceivedAt  time.Time `json:"received_at"`
	ByCharacter string    `json:"by_character"`
	Members     int       `json:"members"`
}

// ViewerView is the home endpoint's own "who is asking" object.
type ViewerView struct {
	Role         string  `json:"role"`
	CharacterKey *string `json:"character_key,omitempty"`
	Verified     bool    `json:"verified"`
	// OfficerTools is true when the viewer is a verified officer/leader of a guild
	// whose officer views are unlocked (a plan or a scoped grant).
	OfficerTools bool `json:"officer_tools"`
	// RaidGuildTools is true when the guild also holds the Raid guild tier, which
	// unlocks the loot council and the readiness board.
	RaidGuildTools bool `json:"raid_guild_tools"`
	// Verify is whether anyone in the guild can approve a character; the web
	// tells an unverified member whom to ask from it.
	Verify auth.GuildVerifyView `json:"verify"`
}

// SummaryView is the Overview tab's own glance row (design spec §4.A).
type SummaryView struct {
	Raiders            int `json:"raiders"`
	WaitingForApproval int `json:"waiting_for_approval"`
	BelowRatingFloor   int `json:"below_rating_floor"`
	// Tier is the Progression tab's tier bar (Store.Progression for a member): the band and the
	// Overview read their counts from it. NamedEncountersDown and PullsThisTier restate its
	// down and pulls_total for the apps.
	Tier                TierView  `json:"tier"`
	NamedEncountersDown int       `json:"named_encounters_down"`
	PullsThisTier       int       `json:"pulls_this_tier"`
	UpdatedAt           time.Time `json:"updated_at"`
}

// StandingView is the member/officer-only "where do I stand" line (design spec §4.A.1),
// nil whenever it cannot be derived (no known spec/class for the viewer, no gear consent,
// or an unverified viewer character).
type StandingView struct {
	Spec                string   `json:"spec"`
	Class               string   `json:"class"`
	SameSpecCount       int      `json:"same_spec_count"`
	RankByItemLevel     int      `json:"rank_by_item_level"`
	ItemLevel           int      `json:"item_level"`
	NeedsBeforeNextRaid []string `json:"needs_before_next_raid"`
	// NeedsRaid is the raid needs_before_next_raid is before: the next planned raid of kind
	// raid (Store.NextRaid, the one Readiness's chip, the nudge and /raid name first), never a
	// dungeon run or guild event the band may show; null when none is planned.
	NeedsRaid *NeedsRaidView `json:"needs_raid"`
}

// NeedsRaidView names the raid the standing line's needs are before. CharacterKey is the
// character the viewer's account signed up for it, its best place (In, then Late, then Bench;
// an Out answer names none): the needs line is that character's, null when the account has no
// such sign-up and the line keeps the header character (guild round 3: the line named the
// header character, not the one the account signed up).
type NeedsRaidView struct {
	Title        string    `json:"title"`
	StartsAt     time.Time `json:"starts_at"`
	CharacterKey *string   `json:"character_key"`
}

// needsRaid is the standing line's raid at now for viewerUserID (NeedsRaidView), nil when none
// is planned.
func (s *Store) needsRaid(ctx context.Context, guildID, viewerUserID int64, now time.Time) (*NeedsRaidView, error) {
	raid, ok, err := s.NextRaid(ctx, guildID, now)
	if err != nil || !ok {
		return nil, err
	}
	signups, err := s.RaidSignups(ctx, raid.ID)
	if err != nil {
		return nil, err
	}
	return &NeedsRaidView{Title: raid.Title, StartsAt: raid.StartsAt.UTC(),
		CharacterKey: signedUpCharacter(signups, viewerUserID)}, nil
}

// signedUpCharacter is the key of the account's best-placed sign-up among signups (In, Late,
// Bench; the first of a tie), nil when the account has none of those.
func signedUpCharacter(signups []Signup, accountID int64) *string {
	var best *Signup
	for i := range signups {
		sg := &signups[i]
		if sg.SignedBy != accountID || sg.State == SignupOut {
			continue
		}
		if best == nil || placeRank[sg.State] < placeRank[best.State] {
			best = sg
		}
	}
	if best == nil {
		return nil
	}
	key := best.CharacterKey
	return &key
}

// signedKeyOf is needs' signed-up character key, "" when there is none.
func signedKeyOf(needs *NeedsRaidView) string {
	if needs == nil || needs.CharacterKey == nil {
		return ""
	}
	return *needs.CharacterKey
}

type HomeView struct {
	Guild      GuildIdentity  `json:"guild"`
	Viewer     ViewerView     `json:"viewer"`
	Claim      ClaimStateView `json:"claim"`
	Summary    SummaryView    `json:"summary"`
	Standing   *StandingView  `json:"standing,omitempty"`
	Reports    []HomeReport   `json:"reports"`
	NextCursor string         `json:"next_cursor,omitempty"`
	Roster     []RosterRow    `json:"roster"`
	// Pending is every unverified roster row, officer view only - empty, never omitted,
	// for everyone else (contract: "others: []").
	Pending []RosterRow `json:"pending"`
	// AddonSnapshot is the guild's newest addon roster snapshot, null when it has none.
	AddonSnapshot *AddonSnapshotView `json:"addon_snapshot"`
	// Members is the roster as a person counts it: everyone the newest addon snapshot names plus
	// every verified character, merged by name, each with a state (RosterEntry), and Counts the
	// numbers above it. Roster and Pending above stay the account rows the readiness and loot
	// views are built from.
	Members []RosterEntry `json:"members"`
	Counts  RosterCounts  `json:"counts"`
	// NextRaid is the guild band's raid line: the next planned or locked raid, absent when
	// there is none (and for anyone but a verified member).
	NextRaid *NextRaidView `json:"next_raid,omitempty"`
	// AttendanceSource is "raids" once the guild has a logged raid, else "reports".
	AttendanceSource string `json:"attendance_source"`
}

// HomeReports lists this guild's reports from the trailing week, newest
// first, keyset-paginated. A report is the guild's as the matcher reads it
// (reportguild.Is): set by an officer, or attributed from its logger's
// verified standing. A member sees their own report regardless of
// visibility, every public report regardless of their own verification,
// and a guild-visible report only once they are verified - never a
// private or unlisted row that is not their own, even once verified: an
// unlisted report is discoverable by every member on this page, which
// defeats its whole point (2026-09-21 security review response, spec
// §3.3's amendment).
func (s *Store) HomeReports(ctx context.Context, guildID, userID int64, verified bool, before *HomeCursor) ([]HomeReport, error) {
	since := time.Now().Add(-homeReportsWindow)
	const columns = `r.id, r.title, r.zone, r.created_at,
	       (select count(*) from fights f where f.report_id = r.id),
	       (select count(*) from fights f where f.report_id = r.id and f.kill),
	       coalesce((select max(f.start_ms + f.duration_ms) - min(f.start_ms)
	                 from fights f where f.report_id = r.id), 0)`
	visibility := `(r.owner_id = $3 or r.visibility = 'public')`
	if verified {
		visibility = `(r.owner_id = $3 or r.visibility = 'public' or r.visibility = 'guild')`
	}
	query := `select ` + columns + ` from reports r
		where ` + reportguild.Is("r", "$1") + ` and r.created_at >= $2 and ` + visibility + `
		order by r.created_at desc, r.id desc limit $4`
	args := []any{guildID, since, userID, HomeReportsPerPage}
	if before != nil {
		query = `select ` + columns + ` from reports r
			where ` + reportguild.Is("r", "$1") + ` and r.created_at >= $2 and ` + visibility + `
			  and (r.created_at, r.id) < ($4, $5)
			order by r.created_at desc, r.id desc limit $6`
		args = []any{guildID, since, userID, before.CreatedAt, before.ID, HomeReportsPerPage}
	}
	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("guilds: home reports: %w", err)
	}
	defer rows.Close()
	out := []HomeReport{}
	for rows.Next() {
		var h HomeReport
		if err := rows.Scan(&h.ID, &h.Title, &h.Zone, &h.CreatedAt, &h.FightCount, &h.KillCount, &h.DurationMS); err != nil {
			return nil, fmt.Errorf("guilds: home reports: %w", err)
		}
		out = append(out, h)
	}
	return out, rows.Err()
}

// HomeRoster lists this guild's characters (synthetic account:-prefixed
// invite rows excluded — they have no class/spec/item level to show),
// verified and unverified both, gear columns withheld in the query
// itself for any row whose account has not consented to show them.
// Class, spec, role, item level and faction are the character's newest
// ranked row in a report that is evidence (presence.EvidenceReport: not
// private, not flagged).
// "Logged recently" (the literal `interval '24 hours'` below) is a
// 24-hour proxy for "ran the companion recently", in the absence of a
// real raid-night concept to anchor to. MayRemove is computed per row
// against the viewing actor (item 7, fourth security review response):
// actorID, moderator and verifiedOfficer describe who is asking, and g
// carries the guild's claimed_by the same rule reads for an
// officer-rank row - both already in the caller's hands from Home, so
// this costs no extra query.
func (s *Store) HomeRoster(ctx context.Context, guildID int64, g Guild, actorID int64, moderator, verifiedOfficer bool) ([]RosterRow, error) {
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, err
	}
	return s.homeRoster(ctx, guildID, source, g, actorID, moderator, verifiedOfficer)
}

// homeRoster is HomeRoster for a caller that has already read the guild's attendance source,
// so one request reads it once.
func (s *Store) homeRoster(ctx context.Context, guildID int64, source string, g Guild, actorID int64, moderator, verifiedOfficer bool) ([]RosterRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, gc.user_id, coalesce(ae.region, c.region), coalesce(ae.ruleset, c.ruleset),
		       coalesce(ae.name, c.name), gc.rank, gc.verified_at is not null,
		       coalesce(ae.updated_at >= now() - interval '24 hours', false), coalesce(ae.updated_at, gc.refreshed_at), gm.consent,
		       case when gm.consent in ('gear', 'gear_bags') then fm.class end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.spec end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.role end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.ilvl end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.faction end,
		       case when gm.consent in ('gear', 'gear_bags') then ae.export end,
		       `+rosterSimExportSQL+`,
		       case when gm.consent in ('gear', 'gear_bags') then gc.raid_spec end,
		       coalesce(u.main_character_key = gc.character_key, false),
		       coalesce(case when gc.verified_by = 'discord' then (select l.display_name from discord_links l where l.user_id = gc.user_id) end, '')
		from guild_characters gc
		left join addon_exports ae on ae.character_key = gc.character_key
		left join characters c on c.key = gc.character_key
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		left join users u on u.id = gc.user_id
		left join lateral (
		  select m.class, m.spec, m.role, m.ilvl, m.faction
		  from fight_metrics m join reports r on r.id = m.report_id
		  where m.player_key = gc.character_key and `+presence.EvidenceReport("r")+`
		  order by m.fought_at desc limit 1
		) fm on true
		where gc.guild_id = $1 and gc.character_key not like 'account:%'
		  -- A character with no export is listed only when it was picked on Discord.
		  and (ae.character_key is not null or (gc.verified_by = 'discord' and c.key is not null))
		order by case gc.rank when 'leader' then 0 when 'officer' then 1 else 2 end, coalesce(ae.name, c.name)
	`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: home roster: %w", err)
	}
	defer rows.Close()
	out := []RosterRow{}
	for rows.Next() {
		var raidSpec *string
		var isMain bool
		var pickedBy string
		row, err := s.scanRosterRow(rows, g, actorID, moderator, verifiedOfficer, &raidSpec, &isMain, &pickedBy)
		if err != nil {
			return nil, err
		}
		row.IsMain = isMain
		if moderator || verifiedOfficer {
			// The picker's Discord name is for officers, who free a wrong pick.
			row.DiscordPickedBy = pickedBy
		}
		out = append(out, row.withRaidSpec(raidSpec))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := s.enrichRoster(ctx, guildID, source, out); err != nil {
		return nil, err
	}
	return out, nil
}

// rosterSimExportSQL is RosterRow.SimExport's column in both roster queries (addon_exports
// aliased ae, guild_members gm): the export is the sim input, read only at gear consent.
var rosterSimExportSQL = `case when gm.consent in ('gear', 'gear_bags') then ` +
	sims.ExportIsInputSQL("ae") + ` else false end`

// scanRosterRow reads one row of a roster query (the columns HomeRoster and
// invitedRosterRows both select, in this order) and finishes it: the remove and approve flags
// against the viewing actor, and the class, spec and gear the export decodes to. extra are a
// query's own trailing columns, scanned after the standard ones.
func (s *Store) scanRosterRow(rows pgx.Rows, g Guild, actorID int64, moderator, verifiedOfficer bool, extra ...any) (RosterRow, error) {
	var row RosterRow
	var export, faction *string
	dest := append([]any{&row.CharacterKey, &row.ownerUserID, &row.Region, &row.Ruleset, &row.Name, &row.Rank,
		&row.Verified, &row.LoggedRecently, &row.LoggedAt, &row.Consent, &row.Class, &row.Spec, &row.Role,
		&row.ItemLevel, &faction, &export, &row.SimExport}, extra...)
	if err := rows.Scan(dest...); err != nil {
		return RosterRow{}, fmt.Errorf("guilds: home roster: %w", err)
	}
	if faction != nil {
		row.faction = *faction
	}
	row.MayRemove = mayRemoveRow(g, actorID, row.ownerUserID, row.Rank, moderator, verifiedOfficer)
	row.MayApprove = verifiedOfficer && !row.Verified
	row.AccountKey = accountKeyFor(row.ownerUserID)
	if row.Class != nil {
		row.className = *row.Class
	}
	if row.Spec != nil {
		row.specStr = *row.Spec
	}
	if export != nil {
		s.applyExport(&row, *export)
	}
	return row, nil
}

// applyExport fills row from its addon export: gear, enchants, bags and the prep sections,
// professions, talents, level, and the class, faction and item level the export carries. An
// export that does not decode leaves the row as it was.
func (s *Store) applyExport(row *RosterRow, export string) {
	decoded, ok := fs1.Decode(export)
	if !ok {
		return
	}
	row.gear, row.enchants = decoded.Gear, decoded.Enchants
	row.bags, row.hasBagsSection, row.prep = decoded.Bags, decoded.HasBags, prepOf(decoded)
	row.Professions = decoded.Professions
	row.talentPointsSpent, row.hasTalents = sumPoints(decoded.Talents.Points), true
	if decoded.HasLevel {
		level := decoded.Level
		row.Level = &level
	}
	// Class comes from fight_metrics first (a logged fight is live evidence of
	// what was actually played), the FS1 export second - the export always
	// carries a class (fs1.Decoded.ClassSlug), so a character with gear consent
	// but no fight_metrics row yet (readiness/standing's own "no fight data"
	// case) still gets a class instead of reading empty and drawing a broken
	// crest web-side. The export carries no spec field at all, so specStr has
	// no equivalent second source.
	if row.className == "" && decoded.ClassSlug != "" {
		row.className = decoded.ClassSlug
		class := decoded.ClassSlug
		row.Class = &class
	}
	// faction has no fight_metrics source at all (that table carries no
	// faction column - see factionForRace's own doc comment), so the export's
	// race is the only place this ever comes from.
	if row.faction == "" {
		if f, ok := s.factionForRace(decoded.RaceSlug); ok {
			row.faction = f
		}
	}
	if level, ok := equippedItemLevel(decoded, sims.GearItem); ok {
		row.ItemLevel = &level
	}
}

// itemLookup is an item in a build's per-class table: sims.GearItem, the table the character
// card reads its worn item levels from.
type itemLookup func(build, classSlug string, itemID int) (trees.Item, bool)

// equippedItemLevel is an export's one item-level figure (upgrades.EquippedItemLevel, the number
// /char shows too) over its equipped gear slots, each item read from the export's own data build
// as the card reads it; ok is false with no gear, or an item that build lacks, and the row keeps
// its latest log's item level.
func equippedItemLevel(decoded fs1.Decoded, item itemLookup) (int, bool) {
	var levels []*int
	for _, slot := range upgrades.Slots {
		id, worn := decoded.Gear[slot.Key]
		if !worn || id <= 0 {
			continue
		}
		var level *int
		if it, found := item(decoded.DataBuild, decoded.ClassSlug, id); found {
			level = &it.ItemLevel
		}
		levels = append(levels, level)
	}
	return upgrades.EquippedItemLevel(levels)
}

// snapshotMember is one stored member line of guild_roster_snapshots.members, as
// addon.RosterMember writes it (that package imports this one, so the shape is restated).
type snapshotMember struct {
	Name    string `json:"name"`
	Level   int    `json:"level"`
	Spec    string `json:"spec"`
	GearGap int    `json:"gear_gap"`
	Build   string `json:"build"`
	SeenAt  string `json:"seen_at"`
	Class   string `json:"class"`
}

// homeAddonSnapshot reads the guild's newest roster snapshot in one query and sets each
// roster row's Addon from it. The snapshot's names carry no realm, or the sender's
// Name-Realm form: the realm half is dropped and the name matched to a row's name, any
// case, the guild's realm being the row's. The snapshot's member lines come back too, for the
// merged roster. The first line of a name wins; a line that
// matches no row is counted in Members and nothing else. The newest snapshot wins even when
// it lists nobody (an empty roster is a fact), so every row then reads null. Nil, with every
// row's Addon nil, when the guild has no snapshot.
func (s *Store) homeAddonSnapshot(ctx context.Context, guildID int64, roster []RosterRow) (*AddonSnapshotView, []snapshotMember, error) {
	var snap AddonSnapshotView
	var members []snapshotMember
	err := s.Pool.QueryRow(ctx, `
		select character_key, members, received_at from guild_roster_snapshots
		where guild_id = $1 order by received_at desc limit 1`, guildID).
		Scan(&snap.ByCharacter, &members, &snap.ReceivedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("guilds: home addon snapshot: %w", err)
	}
	snap.ReceivedAt = snap.ReceivedAt.UTC()
	snap.Members = len(members)
	byName := make(map[string]snapshotMember, len(members))
	for _, m := range members {
		name, _ := character.SplitUnit(m.Name)
		key := strings.ToLower(name)
		if _, seen := byName[key]; !seen && key != "" {
			byName[key] = m
		}
	}
	for i := range roster {
		m, ok := byName[strings.ToLower(roster[i].Name)]
		if !ok {
			continue
		}
		line := &AddonMemberView{SeenAt: m.SeenAt, SnapshotAt: snap.ReceivedAt}
		if hasGearConsent(roster[i].Consent) {
			level, spec := m.Level, m.Spec
			line.Level, line.Spec = &level, &spec
			if roster[i].Level == nil && m.Level > 0 {
				roster[i].Level = &level
			}
			line.GearGap = s.rosterGearGap(roster[i])
			if m.Build != "" {
				build := m.Build
				line.Build = &build
			}
		}
		roster[i].Addon = line
	}
	return &snap, members, nil
}

// rosterGearGap is the roster's "Gear gap N DPS": the readiness board's own gap
// (computeReadiness over the same band and export gear; /char lists the same band's slots) in
// whole DPS, rounded as the board rounds it. The addon's own announced figure is never shown:
// it measured against the addon's data, not the site's, and read 38 DPS where the board said
// no upgrades (guild round 3). Nil when the board has no DPS figure for the row: no export gear
// or band, a healer's or tank's band, or a leveling character.
func (s *Store) rosterGearGap(row RosterRow) *int {
	cr := s.readinessOf(row, nil)
	if !cr.GearChecked || cr.GearUnit != gearUnitDPS || len(row.gear) == 0 {
		return nil
	}
	gap := int(math.Round(cr.GearGap.GainDps))
	return &gap
}

// accountKeyFor is the contract's own "same value for a main and its alts" grouping key.
func accountKeyFor(userID int64) string { return fmt.Sprintf("%s%d", accountKeyPrefix, userID) }

// sumPoints totals a decoded talent split's per-tree points into the single spent-points
// figure the readiness board's unspent-points check needs.
func sumPoints(points []int) int {
	total := 0
	for _, p := range points {
		total += p
	}
	return total
}

// Home assembles the signed-in guild home shell: identity, claim state,
// this week's reports, and the roster. moderator and verifiedOfficer
// describe userID's own standing, read once here rather than
// re-derived per roster row.
func (s *Store) Home(ctx context.Context, guildID, userID int64, verified, moderator, verifiedOfficer bool, before *HomeCursor) (HomeView, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return HomeView{}, err
	}
	reports, err := s.HomeReports(ctx, guildID, userID, verified, before)
	if err != nil {
		return HomeView{}, err
	}
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return HomeView{}, err
	}
	roster, err := s.homeRoster(ctx, guildID, source, g, userID, moderator, verifiedOfficer)
	if err != nil {
		return HomeView{}, err
	}
	// Before pending is copied out of roster, so a pending row carries its line too.
	addonSnapshot, snapshotLines, err := s.homeAddonSnapshot(ctx, guildID, roster)
	if err != nil {
		return HomeView{}, err
	}
	invited, err := s.invitedMembers(ctx, guildID)
	if err != nil {
		return HomeView{}, err
	}
	markRemovable(invited, g, userID, moderator, verifiedOfficer)
	describeRosterRows(roster)
	snapshotLines, err = s.withExportClasses(ctx, g, snapshotLines)
	if err != nil {
		return HomeView{}, err
	}
	members, counts := mergeRoster(roster, invited, snapshotLines)
	describeMembers(members, roster)
	claim, err := s.withClaimHolder(ctx, claimState(g, time.Now()), g, userID)
	if err != nil {
		return HomeView{}, err
	}

	summary, err := s.homeSummary(ctx, guildID, roster)
	if err != nil {
		return HomeView{}, err
	}

	pending := []RosterRow{}
	if verifiedOfficer {
		for _, row := range roster {
			if !row.Verified {
				pending = append(pending, row)
			}
		}
	}

	role := "member"
	switch {
	case moderator:
		role = "moderator"
	case verifiedOfficer:
		role = "officer"
	}
	var viewerKey *string
	for i := range roster {
		if roster[i].ownerUserID == userID {
			key := roster[i].CharacterKey
			viewerKey = &key
			if roster[i].Verified {
				break // a verified character is the more representative "who is this" pick
			}
		}
	}

	verifyViews, err := s.VerifyViews(ctx, []int64{guildID})
	if err != nil {
		return HomeView{}, err
	}

	needs, err := s.needsRaid(ctx, guildID, userID, time.Now())
	if err != nil {
		return HomeView{}, err
	}
	view := HomeView{
		Guild: GuildIdentity{
			ID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name, Faction: g.Faction,
			CrestURL: CrestURL(s.APIBaseURL, g.ID, g.CrestKey),
		},
		Viewer:   ViewerView{Role: role, CharacterKey: viewerKey, Verified: verified, Verify: verifyViews[guildID]},
		Claim:    claim,
		Summary:  summary,
		Standing: standingFor(s, roster, userID, signedKeyOf(needs)),
		Reports:  reports, Roster: roster, Pending: pending,
		AddonSnapshot: addonSnapshot, Members: members, Counts: counts,
		AttendanceSource: source,
	}
	if view.Standing != nil {
		view.Standing.NeedsRaid = needs
	}
	if len(reports) == HomeReportsPerPage {
		last := reports[len(reports)-1]
		view.NextCursor = encodeHomeCursor(last.CreatedAt, last.ID)
	}
	return view, nil
}

// home is gated by IsVerifiedMember: the summary counts the guild's
// private nights (task L6.12), so an unverified row (a bare export,
// possibly forged, or an account an officer has not approved) gets 403
// like any other non-member. Report *content* access stays governed by
// mayView server-side wherever a report link is actually followed.
func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	member, err := s.Store.IsVerifiedMember(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return
	}
	rank, verified, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	verifiedOfficer := verified && (rank == "officer" || rank == "leader")
	moderator := actor.IsModerator()
	var before *HomeCursor
	if v := r.URL.Query().Get("cursor"); v != "" {
		c, ok := decodeHomeCursor(v)
		if !ok {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a cursor",
				map[string]string{"cursor": "the next_cursor a previous page returned"})
			return
		}
		before = &c
	}
	view, err := s.Store.Home(r.Context(), guildID, actor.UserID, verified, moderator, verifiedOfficer, before)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "home", err, "could not load that guild's home just now")
		return
	}
	if verified {
		if view.NextRaid, err = s.Store.NextRaidFor(r.Context(), guildID, actor.UserID, view.Viewer.CharacterKey, time.Now()); err != nil {
			s.fail(w, r, "home", err, "could not load that guild's home just now")
			return
		}
	}
	s.markSupporters(r.Context(), view.Roster, view.Pending)
	if verifiedOfficer {
		view.Viewer.OfficerTools = s.officerTools(r.Context(), guildID, actor.UserID)
		view.Viewer.RaidGuildTools = s.raidGuildTools(r.Context(), guildID, actor.UserID)
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}
