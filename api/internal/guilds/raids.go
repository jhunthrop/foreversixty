// api/internal/guilds/raids.go
//
// GET /v1/guilds/{id}/raids (contract step 3,
// docs/contracts/2026-10-04-guild-centre-api.md): every raid night, cursor-paginated,
// visibility-filtered by the report rules for the viewer (public reports to anyone,
// guild-visibility reports to a verified member, every report to its own owner), and a
// flagged report to no one but a verified member who owns it (visibilityClause) - this
// route is reachable by a signed-out visitor (design spec §4.0's own table: "Raids: Shown
// (public reports only)"), so it is mounted with no auth.RequireSession wrap.
package guilds

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/reportguild"
)

// RaidsDefaultLimit/RaidsMaxLimit bound the ?limit= query parameter (contract's own
// "?limit=20" default).
const (
	RaidsDefaultLimit = 20
	RaidsMaxLimit     = 50
)

// RaidFight is one pull within a raid night.
type RaidFight struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	EncounterID *int64 `json:"encounter_id,omitempty"`
	Kill        bool   `json:"kill"`
	DurationMS  int64  `json:"duration_ms"`
	Deaths      int    `json:"deaths"`
	Players     int    `json:"players"`
}

// PresentCharacter is one roster character who showed up for a raid night.
type PresentCharacter struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
}

// TopParse is a raid row's own best kill-fight parse - fight_metrics' own player_name/class
// (public report data, never gated by the guild's roster consent model - the same figures
// rankings' own leaderboards already show for any report whose visibility permits reading
// it at all).
type TopParse struct {
	Name   string  `json:"name"`
	Class  string  `json:"class"`
	Metric string  `json:"metric"`
	Value  float64 `json:"value"`
}

// RaidRow is one report as the Raids tab shows it.
type RaidRow struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Zone       string    `json:"zone"`
	CreatedAt  time.Time `json:"created_at"`
	DurationMS int64     `json:"duration_ms"`
	FightCount int       `json:"fight_count"`
	KillCount  int       `json:"kill_count"`
	WipeCount  int       `json:"wipe_count"`
	Raiders    int       `json:"raiders"`
	Deaths     int       `json:"deaths"`
	// TopHPS and TopDPS are the night's best healer HPS and best DPS-role DPS (parse_source.go),
	// left out when the night has none.
	TopHPS  *TopParse          `json:"top_hps,omitempty"`
	TopDPS  *TopParse          `json:"top_dps,omitempty"`
	Fights  []RaidFight        `json:"fights"`
	Present []PresentCharacter `json:"present"`
	// RaidID is the guild raid this night is matched to (a verified member's read only).
	RaidID *int64 `json:"raid_id"`

	// evidence is whether the report can show who was there: not private and not flagged
	// (tampered or unverified). Any other report lists nobody.
	evidence bool
}

// RaidsPage is the body of GET /v1/guilds/{id}/raids.
type RaidsPage struct {
	Rows       []RaidRow `json:"rows"`
	NextCursor string    `json:"next_cursor,omitempty"`
	// PlannedSection adds planned, zones, guild_timezone and discord for a verified member only; nil
	// (every key absent) for anyone else (R1).
	*PlannedSection
}

const raidsCursorSep = "|"

type raidsCursor struct {
	CreatedAt time.Time
	ID        string
}

func encodeRaidsCursor(createdAt time.Time, id string) string {
	raw := createdAt.UTC().Format(time.RFC3339Nano) + raidsCursorSep + id
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func decodeRaidsCursor(s string) (raidsCursor, bool) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return raidsCursor{}, false
	}
	createdAt, id, ok := strings.Cut(string(raw), raidsCursorSep)
	if !ok || id == "" {
		return raidsCursor{}, false
	}
	t, err := time.Parse(time.RFC3339Nano, createdAt)
	if err != nil {
		return raidsCursor{}, false
	}
	return raidsCursor{CreatedAt: t, ID: id}, true
}

// visibilityClause is the Raids tab's rule for which reports a viewer sees: the report
// rules for the viewer (contract: "Visibility follows the existing report rules for the
// viewer"), and a flagged report (tampered or unverified) only to a verified member who owns
// it, who manages it. A verified member sees their own reports, and the guild's public and
// guild-visibility ones that are not flagged (presence.NotFlagged). Anyone else (signed out,
// another account, an unverified row) sees the public reports that are evidence
// (presence.EvidenceReport), and their own reports that are not flagged.
func visibilityClause(verified bool) string {
	if verified {
		return `(r.owner_id = $2 or ((r.visibility = 'public' or r.visibility = 'guild') and ` +
			presence.NotFlagged("r") + `))`
	}
	return `((r.owner_id = $2 and ` + presence.NotFlagged("r") + `) or (r.visibility = 'public' and ` +
		presence.EvidenceReport("r") + `))`
}

// Raids lists guildID's reports, newest first, keyset-paginated, each with its fights and
// present-roster list. A report is the guild's as the matcher reads it (reportguild.Is): set
// by an officer, or attributed from its logger's verified standing, so a night the matcher
// linked to a raid is on the tab that links to it. userID is 0 for a signed-out visitor - the visibility clause's own
// r.owner_id = $2 then simply never matches, which is exactly "a signed-out visitor only
// ever sees public, unflagged reports."
func (s *Store) Raids(ctx context.Context, guildID, userID int64, verified bool, limit int, before *raidsCursor) (RaidsPage, error) {
	if limit <= 0 || limit > RaidsMaxLimit {
		limit = RaidsDefaultLimit
	}
	visibility := visibilityClause(verified)
	query := `select r.id, r.title, r.zone, r.created_at,
	       coalesce((select max(f.start_ms + f.duration_ms) - min(f.start_ms) from fights f where f.report_id = r.id), 0),
	       (select count(*) from fights f where f.report_id = r.id),
	       (select count(*) filter (where f.kill) from fights f where f.report_id = r.id),
	       (select count(*) filter (where not f.kill) from fights f where f.report_id = r.id),
	       coalesce((select sum(f.deaths) from fights f where f.report_id = r.id), 0),
	       ` + presence.EvidenceReport("r") + `, gr.id
	     from reports r
	     left join guild_raids gr on gr.report_id = r.id and gr.guild_id = $1
	     where ` + reportguild.Is("r", "$1") + ` and ` + visibility
	args := []any{guildID, userID}
	if before != nil {
		query += ` and (r.created_at, r.id) < ($3, $4)`
		args = append(args, before.CreatedAt, before.ID)
	}
	query += fmt.Sprintf(` order by r.created_at desc, r.id desc limit $%d`, len(args)+1)
	args = append(args, limit)

	rows, err := s.Pool.Query(ctx, query, args...)
	if err != nil {
		return RaidsPage{}, fmt.Errorf("guilds: raids: %w", err)
	}
	defer rows.Close()
	page := RaidsPage{Rows: []RaidRow{}}
	for rows.Next() {
		var row RaidRow
		if err := rows.Scan(&row.ID, &row.Title, &row.Zone, &row.CreatedAt, &row.DurationMS,
			&row.FightCount, &row.KillCount, &row.WipeCount, &row.Deaths, &row.evidence, &row.RaidID); err != nil {
			return RaidsPage{}, fmt.Errorf("guilds: raids: %w", err)
		}
		if !verified {
			row.RaidID = nil
		}
		page.Rows = append(page.Rows, row)
	}
	if err := rows.Err(); err != nil {
		return RaidsPage{}, err
	}

	present, err := s.raidPresence(ctx, page.Rows)
	if err != nil {
		return RaidsPage{}, err
	}
	for i := range page.Rows {
		if err := s.fillRaidDetail(ctx, &page.Rows[i], present[page.Rows[i].ID]); err != nil {
			return RaidsPage{}, err
		}
	}

	if len(page.Rows) == limit {
		last := page.Rows[len(page.Rows)-1]
		page.NextCursor = encodeRaidsCursor(last.CreatedAt, last.ID)
	}
	return page, nil
}

// raidPresence is report id -> the character keys present in it, for every row of the page
// in one presence.In call: a rated or ranked row read by character key, never fights.players
// (the log's GUIDs). A private or flagged report is not asked about, so it lists nobody.
func (s *Store) raidPresence(ctx context.Context, rows []RaidRow) (map[string][]string, error) {
	var ids []string
	for _, row := range rows {
		if row.evidence {
			ids = append(ids, row.ID)
		}
	}
	here, err := presence.In(ctx, s.Pool, ids, nil)
	if err != nil {
		return nil, fmt.Errorf("guilds: raid present: %w", err)
	}
	out := map[string][]string{}
	seen := map[[2]string]bool{}
	for _, p := range here {
		if k := [2]string{p.ReportID, p.Key}; !seen[k] {
			seen[k] = true
			out[p.ReportID] = append(out[p.ReportID], p.Key)
		}
	}
	return out, nil
}

// fillRaidDetail reads one report's own fights, present roster (the keys raidPresence found
// in it) and top parse.
func (s *Store) fillRaidDetail(ctx context.Context, row *RaidRow, keys []string) error {
	fightRows, err := s.Pool.Query(ctx,
		`select fight_index, name, encounter_id, kill, duration_ms, deaths, coalesce(array_length(players, 1), 0)
		 from fights where report_id = $1 order by fight_index`, row.ID)
	if err != nil {
		return fmt.Errorf("guilds: raid fights: %w", err)
	}
	row.Fights = []RaidFight{}
	for fightRows.Next() {
		var f RaidFight
		var encounterID *int64
		if err := fightRows.Scan(&f.Index, &f.Name, &encounterID, &f.Kill, &f.DurationMS, &f.Deaths, &f.Players); err != nil {
			fightRows.Close()
			return fmt.Errorf("guilds: raid fights: %w", err)
		}
		if encounterID != nil && *encounterID != 0 {
			f.EncounterID = encounterID
		}
		row.Fights = append(row.Fights, f)
	}
	fightRows.Close()
	if err := fightRows.Err(); err != nil {
		return err
	}

	row.Present = []PresentCharacter{}
	if len(keys) == 0 {
		// A private, flagged or trash-only night, or one not ranked yet: nobody to list.
		return s.fillTopParses(ctx, row)
	}
	presentRows, err := s.Pool.Query(ctx, `
		select distinct on (p.character_key) p.character_key, coalesce(ae.name, fm.player_name, ''), coalesce(fm.class, '')
		from unnest($2::text[]) as p(character_key)
		left join addon_exports ae on ae.character_key = p.character_key
		left join fight_metrics fm on fm.report_id = $1 and fm.player_key = p.character_key
		order by p.character_key`, row.ID, keys)
	if err != nil {
		return fmt.Errorf("guilds: raid present: %w", err)
	}
	for presentRows.Next() {
		var pc PresentCharacter
		if err := presentRows.Scan(&pc.CharacterKey, &pc.Name, &pc.Class); err != nil {
			presentRows.Close()
			return fmt.Errorf("guilds: raid present: %w", err)
		}
		row.Present = append(row.Present, pc)
	}
	presentRows.Close()
	if err := presentRows.Err(); err != nil {
		return err
	}
	row.Raiders = len(row.Present)
	return s.fillTopParses(ctx, row)
}

// fillTopParses reads the night's Top HPS (healers) and Top DPS (the DPS role), each from boss
// kills only and the role's own metric, the rule Progression's bests read (parse_source.go).
// Each is nil when the night has none yet, and both are nil for a report that is not evidence
// (presence.EvidenceReport: private or flagged), which lists nobody present either.
func (s *Store) fillTopParses(ctx context.Context, row *RaidRow) error {
	var err error
	if row.TopHPS, err = s.topByRole(ctx, row.ID, RoleHealer); err != nil {
		return err
	}
	row.TopDPS, err = s.topByRole(ctx, row.ID, RoleDPS)
	return err
}

// topByRole is a report's best boss-kill figure among role's rows, nil when it has none.
func (s *Store) topByRole(ctx context.Context, reportID, role string) (*TopParse, error) {
	col := roleMetricColumn(role)
	tp := TopParse{Metric: roleMetric(role)}
	err := s.Pool.QueryRow(ctx, `
		select coalesce(m.player_name, ''), coalesce(m.class, ''), coalesce(m.`+col+`, 0)
		from fight_metrics m join reports r on r.id = m.report_id
		where m.report_id = $1 and m.role = $2 and `+bossKillMetrics+` and `+presence.EvidenceReport("r")+`
		order by coalesce(m.`+col+`, 0) desc
		limit 1`, reportID, role).Scan(&tp.Name, &tp.Class, &tp.Value)
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("guilds: raid top %s: %w", tp.Metric, err)
	}
	return &tp, nil
}

func (s *Service) raids(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
	if err != nil {
		s.fail(w, r, "raids", err, "could not load that guild's raids just now")
		return
	}

	limit := RaidsDefaultLimit
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			limit = n
		}
	}
	var before *raidsCursor
	if v := r.URL.Query().Get("cursor"); v != "" {
		c, ok := decodeRaidsCursor(v)
		if !ok {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a cursor",
				map[string]string{"cursor": "the next_cursor a previous page returned"})
			return
		}
		before = &c
	}

	page, err := s.Store.Raids(r.Context(), guildID, actor.UserID, verified, limit, before)
	if err != nil {
		s.fail(w, r, "raids", err, "could not load that guild's raids just now")
		return
	}
	if verified {
		if page.PlannedSection, err = s.plannedSection(r, guildID, actor.UserID); err != nil {
			s.fail(w, r, "raids", err, "could not load that guild's raids just now")
			return
		}
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, page)
}
