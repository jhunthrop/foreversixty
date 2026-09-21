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
	ItemLevel      *int    `json:"item_level,omitempty"`
	// MayRemove is computed server-side from the exact same
	// mayRemoveRow rule the DELETE .../characters/{key} route enforces
	// (item 7, fourth security review response), from the viewpoint of
	// whoever is asking for this home - so the web can show the remove
	// control exactly where it would actually succeed, never a control
	// that then answers 403.
	MayRemove bool `json:"may_remove"`
	// ownerUserID is gc.user_id: read to compute MayRemove against the
	// viewing actor, never marshaled into the response itself.
	ownerUserID int64
}

type HomeView struct {
	Guild      GuildIdentity  `json:"guild"`
	Claim      ClaimStateView `json:"claim"`
	Reports    []HomeReport   `json:"reports"`
	NextCursor string         `json:"next_cursor,omitempty"`
	Roster     []RosterRow    `json:"roster"`
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
	       (select count(*) from fights f where f.report_id = r.id and f.kill)`
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
		if err := rows.Scan(&h.ID, &h.Title, &h.Zone, &h.CreatedAt, &h.FightCount, &h.KillCount); err != nil {
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
		       ae.updated_at >= now() - interval '24 hours', gm.consent,
		       case when gm.consent in ('gear', 'gear_bags') then fm.class end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.spec end,
		       case when gm.consent in ('gear', 'gear_bags') then fm.ilvl end
		from guild_characters gc
		join addon_exports ae on ae.character_key = gc.character_key
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		left join lateral (
		  select class, spec, ilvl from fight_metrics
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
		if err := rows.Scan(&row.CharacterKey, &row.ownerUserID, &row.Region, &row.Ruleset, &row.Name, &row.Rank,
			&row.Verified, &row.LoggedRecently, &row.Consent, &row.Class, &row.Spec, &row.ItemLevel); err != nil {
			return nil, fmt.Errorf("guilds: home roster: %w", err)
		}
		row.MayRemove = mayRemoveRow(g, actorID, row.ownerUserID, row.Rank, moderator, verifiedOfficer)
		out = append(out, row)
	}
	return out, rows.Err()
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
	view := HomeView{
		Guild:   GuildIdentity{ID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name},
		Claim:   claimState(g, time.Now()),
		Reports: reports, Roster: roster,
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
	httpx.WriteOK(w, r, http.StatusOK, view)
}
