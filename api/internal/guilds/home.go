// api/internal/guilds/home.go
package guilds

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
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

// Attendance is a character's presence over the guild's last 8 reports (contract's own
// window), design spec §4.B.
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
	ItemLevel      *int    `json:"item_level,omitempty"`
	// LoggedAt is addon_exports.updated_at - the contract's own replacement for
	// logged_recently (kept below too, unchanged, per the contract's own instruction).
	LoggedAt     time.Time   `json:"logged_at"`
	Attendance   Attendance  `json:"attendance"`
	BestParse    *BestParse  `json:"best_parse,omitempty"`
	Rating       *RatingView `json:"rating,omitempty"`
	Professions  []string    `json:"professions,omitempty"`
	AccountKey   string      `json:"account_key"`
	LastReportAt *time.Time  `json:"last_report_at"`
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
	// ownerUserID is gc.user_id: read to compute MayRemove/MayApprove/AccountKey against
	// the viewing actor, never marshaled into the response itself.
	ownerUserID int64
	// faction/className/specName/export/gear/enchants/bags/talentPointsSpent/hasTalents
	// back this row's own standing/readiness computation (standing.go) without a second
	// query - never marshaled.
	faction            string
	className, specStr string
	gear, enchants     map[string]int
	bags               []int
	hasBagsSection     bool
	talentPointsSpent  int
	hasTalents         bool
	// level/hasLevel are the export's own level= section (fs1.Decoded.Level/HasLevel) -
	// the readiness board's talent-points-unspent check needs the character's real level,
	// never a level-60 assumption, when the export carries one.
	level    int
	hasLevel bool
}

// ViewerView is the home endpoint's own "who is asking" object.
type ViewerView struct {
	Role         string  `json:"role"`
	CharacterKey *string `json:"character_key,omitempty"`
	Verified     bool    `json:"verified"`
}

// SummaryView is the Overview tab's own glance row (design spec §4.A).
type SummaryView struct {
	Raiders             int       `json:"raiders"`
	WaitingForApproval  int       `json:"waiting_for_approval"`
	BelowRatingFloor    int       `json:"below_rating_floor"`
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
}

// HomeReports lists this guild's reports from the trailing week, newest
// first, keyset-paginated. A member sees their own report regardless of
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
		where r.guild_id = $1 and r.created_at >= $2 and ` + visibility + `
		order by r.created_at desc, r.id desc limit $4`
	args := []any{guildID, since, userID, HomeReportsPerPage}
	if before != nil {
		query = `select ` + columns + ` from reports r
			where r.guild_id = $1 and r.created_at >= $2 and ` + visibility + `
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
// "Logged recently" (the literal `interval '24 hours'` below) is a
// 24-hour proxy for "ran the companion recently", in the absence of a
// real raid-night concept to anchor to. MayRemove is computed per row
// against the viewing actor (item 7, fourth security review response):
// actorID, moderator and verifiedOfficer describe who is asking, and g
// carries the guild's claimed_by the same rule reads for an
// officer-rank row - both already in the caller's hands from Home, so
// this costs no extra query.
func (s *Store) HomeRoster(ctx context.Context, guildID int64, g Guild, actorID int64, moderator, verifiedOfficer bool) ([]RosterRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.character_key, gc.user_id, ae.region, ae.ruleset, ae.name, gc.rank, gc.verified_at is not null,
		       ae.updated_at >= now() - interval '24 hours', ae.updated_at, gm.consent,
		       case when gm.consent in ('gear', 'gear_bags') then fm.class end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.spec end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.role end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.ilvl end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.faction end,
		       case when gm.consent in ('gear', 'gear_bags') then ae.export end
		from guild_characters gc
		join addon_exports ae on ae.character_key = gc.character_key
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		left join lateral (
		  select class, spec, role, ilvl, faction from fight_metrics
		  where player_key = gc.character_key order by fought_at desc limit 1
		) fm on true
		where gc.guild_id = $1 and gc.character_key not like 'account:%'
		order by case gc.rank when 'leader' then 0 when 'officer' then 1 else 2 end, ae.name
	`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: home roster: %w", err)
	}
	defer rows.Close()
	out := []RosterRow{}
	for rows.Next() {
		var row RosterRow
		var export, faction *string
		if err := rows.Scan(&row.CharacterKey, &row.ownerUserID, &row.Region, &row.Ruleset, &row.Name, &row.Rank,
			&row.Verified, &row.LoggedRecently, &row.LoggedAt, &row.Consent, &row.Class, &row.Spec, &row.Role,
			&row.ItemLevel, &faction, &export); err != nil {
			return nil, fmt.Errorf("guilds: home roster: %w", err)
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
			decoded, ok := fs1.Decode(*export)
			if ok {
				row.gear, row.enchants = decoded.Gear, decoded.Enchants
				row.bags, row.hasBagsSection = decoded.Bags, hasBagsSection(*export)
				row.Professions = decoded.Professions
				row.talentPointsSpent, row.hasTalents = sumPoints(decoded.Talents.Points), true
				row.level, row.hasLevel = decoded.Level, decoded.HasLevel
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
			}
		}
		out = append(out, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if err := s.enrichRoster(ctx, guildID, out); err != nil {
		return nil, err
	}
	return out, nil
}

// accountKeyFor is the contract's own "same value for a main and its alts" grouping key.
func accountKeyFor(userID int64) string { return fmt.Sprintf("u:%d", userID) }

// sumPoints totals a decoded talent split's per-tree points into the single spent-points
// figure the readiness board's unspent-points check needs.
func sumPoints(points []int) int {
	total := 0
	for _, p := range points {
		total += p
	}
	return total
}

// hasBagsSection reports whether export's pipe sections include a bags= entry at all
// (distinguishing "no bags= section" from "an empty one," both of which fs1.Decode reads as
// an empty slice) - the consumables check's own "unknown" vs "short" distinction needs this.
func hasBagsSection(export string) bool {
	return strings.Contains(export, "|bags=") || strings.HasPrefix(export, "bags=")
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
	roster, err := s.HomeRoster(ctx, guildID, g, userID, moderator, verifiedOfficer)
	if err != nil {
		return HomeView{}, err
	}
	claim := claimState(g, time.Now())
	if claim.State == "claimed" && g.ClaimedBy != nil && s.Accounts != nil {
		if u, uerr := s.Accounts.User(ctx, *g.ClaimedBy); uerr == nil {
			name := u.PublicName()
			claim.ClaimedByName = &name
		}
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

	view := HomeView{
		Guild:    GuildIdentity{ID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name, Faction: g.Faction},
		Viewer:   ViewerView{Role: role, CharacterKey: viewerKey, Verified: verified},
		Claim:    claim,
		Summary:  summary,
		Standing: standingFor(s, roster, userID),
		Reports:  reports, Roster: roster, Pending: pending,
	}
	if len(reports) == HomeReportsPerPage {
		last := reports[len(reports)-1]
		view.NextCursor = encodeHomeCursor(last.CreatedAt, last.ID)
	}
	return view, nil
}

// home is gated by IsMember — verified or not — never by GuildRank: an
// unverified member still sees the home shell, just as they still
// appear on the roster. Report *content* access stays governed by
// mayView server-side wherever a report link is actually followed.
func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	member, err := s.Store.IsMember(r.Context(), guildID, actor.UserID)
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
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}
