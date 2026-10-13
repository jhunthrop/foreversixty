// api/internal/guilds/progression.go
//
// GET /v1/guilds/{id}/progression (contract step 4,
// docs/contracts/2026-10-04-guild-centre-api.md). The tab is public (design spec §4.0: shown
// to every role), and its counts (pulls, kills, first kill, best kill, deaths per pull, pulls
// by night with their report ids, the unnamed zones' pulls and nights, the tier's bosses
// down) depend on the viewer. A verified member of the guild reads the guild's own numbers,
// which count its private nights too (presence.NotFlagged), in a private response. Everyone
// else (signed out, another account, an unverified row) reads the public numbers, the ones the
// public guild page prints (rankings.Store.Guild): reports that are evidence only
// (presence.EvidenceReport: not private, not flagged), in a shareable response. A flagged
// report (tampered or unverified) counts for no viewer. The best parse per role, a parse a
// player sees and opens, is read from EvidenceReport for every viewer, as every other parse
// reader is (task L6.12).
//
// The tier (Barrow Deeps/Hyjal Summit/Onyxia's Lair, Onyxia the one published encounter)
// is hardcoded rather than read from data/curated/loot/forever-raid-phases.json, matching
// api/cmd/seedguild/raid.go's own choice: that file's own "replace" list carries open
// dates, not display names or a raid roster, and data/curated/ is not copied into this
// service's Docker image at all (api/Dockerfile only copies data/builds/) - so a runtime
// read of it would 404 in production exactly like data/curated/specs.json would (see
// api/internal/bis/catalog.go's own reasoning).
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
)

const (
	tierName        = "First tier"
	zoneBarrowDeeps = "Barrow Deeps"
	zoneHyjal       = "Hyjal Summit"
	zoneOnyxia      = "Onyxia's Lair"
	encounterOnyxia = int64(1084)
	encounterName   = "Onyxia"
)

// tierRaids is this tier's three raids, in the order the design doc names them.
var tierRaids = []string{zoneBarrowDeeps, zoneHyjal, zoneOnyxia}

// unnamedZones are this tier's raids with no published encounter list at all - every pull
// in them counts toward the tier total but cannot be named or ranked per boss (design spec
// §12.1's own ruling).
var unnamedZones = []string{zoneBarrowDeeps, zoneHyjal}

// TierView is the Progression tab's own tier bar.
type TierView struct {
	Name            string   `json:"name"`
	Raids           []string `json:"raids"`
	NamedEncounters int      `json:"named_encounters"`
	Down            int      `json:"down"`
}

// ParseRef identifies one fight_metrics row's own player - distinct from home.go's
// BestParse, which identifies the fight instead (the roster row it sits on already names
// the player).
type ParseRef struct {
	Name       string  `json:"name"`
	Class      string  `json:"class"`
	Spec       string  `json:"spec"`
	Metric     string  `json:"metric"`
	Value      float64 `json:"value"`
	ReportID   string  `json:"report_id"`
	FightIndex int     `json:"fight_index"`
}

// NightPulls is one report's own pull count against a named encounter.
type NightPulls struct {
	ReportID string `json:"report_id"`
	Date     string `json:"date"`
	Pulls    int    `json:"pulls"`
	Killed   bool   `json:"killed"`
}

// EncounterProgress is one named encounter's own standing for the guild.
type EncounterProgress struct {
	EncounterID   int64                `json:"encounter_id"`
	Name          string               `json:"name"`
	Zone          string               `json:"zone"`
	FirstKillAt   *time.Time           `json:"first_kill_at,omitempty"`
	Pulls         int                  `json:"pulls"`
	Kills         int                  `json:"kills"`
	BestKillMS    *int64               `json:"best_kill_ms,omitempty"`
	PullsByNight  []NightPulls         `json:"pulls_by_night"`
	DeathsPerPull float64              `json:"deaths_per_pull"`
	BestByRole    map[string]*ParseRef `json:"best_by_role"`
}

// UnnamedZoneProgress is one of this tier's raids with no published encounter list.
type UnnamedZoneProgress struct {
	Zone   string `json:"zone"`
	Pulls  int    `json:"pulls"`
	Nights int    `json:"nights"`
}

// ProgressionView is the body of GET /v1/guilds/{id}/progression.
type ProgressionView struct {
	Tier       TierView              `json:"tier"`
	Encounters []EncounterProgress   `json:"encounters"`
	Unnamed    []UnnamedZoneProgress `json:"unnamed"`
}

// roleMetric is rankings.Store's own metricOf (api/internal/rankings/store.go), duplicated
// narrowly rather than imported: healer reads hps, every other role (including tank) reads
// dps - the existing site-wide convention this endpoint must not invent a second one for.
func roleMetric(role string) string {
	if role == "healer" {
		return "hps"
	}
	return "dps"
}

// countedReports is the rule progression's counts select reports by for a viewer: a verified
// member of the guild counts its private nights (presence.NotFlagged), anyone else reads the
// public numbers (presence.EvidenceReport). Neither counts a flagged report.
func countedReports(member bool) string {
	if member {
		return presence.NotFlagged("r")
	}
	return presence.EvidenceReport("r")
}

// Progression assembles the guild's own tier bar, Onyxia's depth and the two unnamed
// zones' pull counts, as member (a verified member of the guild) or anyone else sees them
// (countedReports).
func (s *Store) Progression(ctx context.Context, guildID int64, member bool) (ProgressionView, error) {
	view := ProgressionView{
		Tier:       TierView{Name: tierName, Raids: tierRaids, NamedEncounters: 1},
		Encounters: []EncounterProgress{},
		Unnamed:    []UnnamedZoneProgress{},
	}
	counted := countedReports(member)

	onyxia, killed, err := s.encounterProgress(ctx, guildID, encounterOnyxia, encounterName, zoneOnyxia, counted)
	if err != nil {
		return ProgressionView{}, err
	}
	view.Encounters = append(view.Encounters, onyxia)
	if killed {
		view.Tier.Down = 1
	}

	for _, zone := range unnamedZones {
		uz, err := s.unnamedZoneProgress(ctx, guildID, zone, counted)
		if err != nil {
			return ProgressionView{}, err
		}
		view.Unnamed = append(view.Unnamed, uz)
	}
	return view, nil
}

// encounterProgress reads one named encounter's standing from the reports counted selects
// (countedReports), and whether it is down.
func (s *Store) encounterProgress(ctx context.Context, guildID, encounterID int64, name, zone, counted string) (EncounterProgress, bool, error) {
	ep := EncounterProgress{EncounterID: encounterID, Name: name, Zone: zone, PullsByNight: []NightPulls{}, BestByRole: map[string]*ParseRef{}}

	var firstKillMS *int64
	var bestKillMS *int64
	var avgDeaths *float64
	err := s.Pool.QueryRow(ctx, `
		select count(*), count(*) filter (where f.kill),
		       min(f.start_ms) filter (where f.kill), min(f.duration_ms) filter (where f.kill),
		       avg(f.deaths)
		from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.encounter_id = $2 and `+counted,
		guildID, encounterID).Scan(&ep.Pulls, &ep.Kills, &firstKillMS, &bestKillMS, &avgDeaths)
	if err != nil {
		return EncounterProgress{}, false, fmt.Errorf("guilds: encounter progress: %w", err)
	}
	if firstKillMS != nil {
		t := time.UnixMilli(*firstKillMS).UTC()
		ep.FirstKillAt = &t
	}
	ep.BestKillMS = bestKillMS
	if avgDeaths != nil {
		ep.DeathsPerPull = *avgDeaths
	}

	nightRows, err := s.Pool.Query(ctx, `
		select r.id, r.created_at::date, count(*), count(*) filter (where f.kill) > 0
		from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.encounter_id = $2 and `+counted+`
		group by r.id, r.created_at order by r.created_at`,
		guildID, encounterID)
	if err != nil {
		return EncounterProgress{}, false, fmt.Errorf("guilds: encounter nights: %w", err)
	}
	for nightRows.Next() {
		var n NightPulls
		var date time.Time
		if err := nightRows.Scan(&n.ReportID, &date, &n.Pulls, &n.Killed); err != nil {
			nightRows.Close()
			return EncounterProgress{}, false, fmt.Errorf("guilds: encounter nights: %w", err)
		}
		n.Date = date.Format("2006-01-02")
		ep.PullsByNight = append(ep.PullsByNight, n)
	}
	nightRows.Close()
	if err := nightRows.Err(); err != nil {
		return EncounterProgress{}, false, err
	}

	for _, role := range []string{"tank", "healer", "dps"} {
		ref, ok, err := s.bestByRole(ctx, guildID, encounterID, role)
		if err != nil {
			return EncounterProgress{}, false, err
		}
		if ok {
			ep.BestByRole[role] = &ref
		}
	}
	return ep, ep.Kills > 0, nil
}

// bestByRole is the guild's best kill parse on encounterID for role, from its reports that are
// evidence (presence.EvidenceReport: not private, not flagged); false when it has none.
func (s *Store) bestByRole(ctx context.Context, guildID, encounterID int64, role string) (ParseRef, bool, error) {
	metric := roleMetric(role)
	orderCol := "metric_dps"
	if metric == "hps" {
		orderCol = "metric_hps"
	}
	var ref ParseRef
	ref.Metric = metric
	err := s.Pool.QueryRow(ctx, `
		select coalesce(m.player_name, ''), coalesce(m.class, ''), coalesce(m.spec, ''),
		       coalesce(m.`+orderCol+`, 0), m.report_id, m.fight_index
		from fight_metrics m join reports r on r.id = m.report_id
		where r.guild_id = $1 and m.encounter_id = $2 and m.role = $3 and m.kill and m.state <> 'removed'
		  and `+presence.EvidenceReport("r")+`
		order by coalesce(m.`+orderCol+`, 0) desc limit 1`,
		guildID, encounterID, role).Scan(&ref.Name, &ref.Class, &ref.Spec, &ref.Value, &ref.ReportID, &ref.FightIndex)
	if errors.Is(err, pgx.ErrNoRows) {
		return ParseRef{}, false, nil
	}
	if err != nil {
		return ParseRef{}, false, fmt.Errorf("guilds: best by role %s: %w", role, err)
	}
	return ref, true, nil
}

// unnamedZoneProgress counts one unnamed zone's pulls and nights in the reports counted selects
// (countedReports).
func (s *Store) unnamedZoneProgress(ctx context.Context, guildID int64, zone, counted string) (UnnamedZoneProgress, error) {
	uz := UnnamedZoneProgress{Zone: zone}
	if err := s.Pool.QueryRow(ctx, `
		select count(*), count(distinct r.id)
		from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and r.zone = $2 and (f.encounter_id is null or f.encounter_id = 0)
		  and `+counted,
		guildID, zone).Scan(&uz.Pulls, &uz.Nights); err != nil {
		return UnnamedZoneProgress{}, fmt.Errorf("guilds: unnamed zone %s: %w", zone, err)
	}
	return uz, nil
}

// progression answers anyone (no session needed), and the counts depend on who asks: the
// handler reads whatever actor the auth middleware put in context and, for a signed-in one,
// whether it is a verified member of the guild (GuildRank, the rule report access reads). A
// member's answer counts the guild's private nights, so it is private and never shared; any
// other answer is the public numbers and may be shared. Vary keeps a shared cache from
// handing the public answer to a member's request, or the other way round.
func (s *Service) progression(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	member := false
	if actor := auth.ActorFrom(r.Context()); actor.Signed() {
		_, verified, err := s.Accounts.GuildRank(r.Context(), guildID, actor.UserID)
		if err != nil {
			s.fail(w, r, "progression", err, "could not load that guild's progression just now")
			return
		}
		member = verified
	}
	view, err := s.Store.Progression(r.Context(), guildID, member)
	if err != nil {
		s.fail(w, r, "progression", err, "could not load that guild's progression just now")
		return
	}
	if member {
		httpx.CachePrivate(w)
	} else {
		httpx.CachePublic(w, progressionMaxAge, progressionStale)
		w.Header().Set("Vary", "Cookie, Authorization")
	}
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// progressionMaxAge/progressionStale match rankings' own public-board cache window, for the
// public numbers (the Live public boards class).
const (
	progressionMaxAge = 30 * time.Second
	progressionStale  = 60 * time.Second
)
