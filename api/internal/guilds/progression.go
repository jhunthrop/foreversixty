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
// The tier's raids are first-tier.json's (raidZones, raids_zones.go); its named encounters are the
// generated boss table's (loot_table.go). A raid the game client does not carry yet has no named
// encounter, so its pulls count toward the tier total as an unnamed zone (design spec §12.1).
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

const tierName = "First tier"

// tierRaids is this tier's raids by name, in first-tier.json's order.
var tierRaids = func() []string {
	out := make([]string, 0, len(raidZones))
	for _, z := range raidZones {
		out = append(out, z.Name)
	}
	return out
}()

// unnamedZones are this tier's raids the boss table holds no boss for (the game client does not
// carry them yet): every pull in them counts toward the tier total but cannot be named or ranked
// per boss (design spec §12.1's own ruling).
var unnamedZones = func() []string {
	var out []string
	for _, z := range raidZones {
		if !raidHasBosses(z.Key) {
			out = append(out, z.Name)
		}
	}
	return out
}()

// raidHasBosses reports whether the boss table holds a boss of zoneKey.
func raidHasBosses(zoneKey string) bool {
	for _, b := range lootBosses {
		if b.RaidKey == zoneKey {
			return true
		}
	}
	return false
}

// TierView is the Progression tab's own tier bar.
type TierView struct {
	Name            string   `json:"name"`
	Raids           []string `json:"raids"`
	NamedEncounters int      `json:"named_encounters"`
	Down            int      `json:"down"`
	// PullsTotal is every pull this tier: the named encounters' and the unnamed zones'. Zero
	// means the guild has logged nothing in the tier yet (week-one spec §4.1).
	PullsTotal int `json:"pulls_total"`
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

// NightPulls is one report's own pull count against a named encounter, with that night's
// fastest kill (null on a night without one) and every death across its pulls.
type NightPulls struct {
	ReportID string `json:"report_id"`
	Date     string `json:"date"`
	Pulls    int    `json:"pulls"`
	Killed   bool   `json:"killed"`
	KillMS   *int64 `json:"kill_ms"`
	Deaths   int    `json:"deaths"`
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

// Progression assembles the guild's own tier bar, each table boss's depth and the unnamed
// zones' pull counts, as member (a verified member of the guild) or anyone else sees them
// (countedReports).
func (s *Store) Progression(ctx context.Context, guildID int64, member bool) (ProgressionView, error) {
	view := ProgressionView{
		Tier:       TierView{Name: tierName, Raids: tierRaids, NamedEncounters: len(lootBosses)},
		Encounters: []EncounterProgress{},
		Unnamed:    []UnnamedZoneProgress{},
	}
	counted := countedReports(member)

	for _, b := range lootBosses {
		ep, err := s.encounterProgress(ctx, guildID, b.EncounterID, b.Name, b.RaidName, counted)
		if err != nil {
			return ProgressionView{}, err
		}
		view.Encounters = append(view.Encounters, ep)
		view.Tier.PullsTotal += ep.Pulls
	}
	down, err := s.tierDown(ctx, guildID, counted)
	if err != nil {
		return ProgressionView{}, err
	}
	view.Tier.Down = down

	for _, zone := range unnamedZones {
		uz, err := s.unnamedZoneProgress(ctx, guildID, zone, counted)
		if err != nil {
			return ProgressionView{}, err
		}
		view.Unnamed = append(view.Unnamed, uz)
		view.Tier.PullsTotal += uz.Pulls
	}
	return view, nil
}

// tierDown counts this tier's named encounters the guild has killed in the reports counted
// selects (countedReports): the one figure the Progression tab's tier bar, the Overview and the
// page header all read, so they never disagree.
func (s *Store) tierDown(ctx context.Context, guildID int64, counted string) (int, error) {
	var down int
	if err := s.Pool.QueryRow(ctx, `
		select count(distinct f.encounter_id) from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.kill and f.encounter_id = any($2) and `+counted,
		guildID, allLoot.encounters).Scan(&down); err != nil {
		return 0, fmt.Errorf("guilds: tier down: %w", err)
	}
	return down, nil
}

// encounterProgress reads one named encounter's standing from the reports counted selects
// (countedReports).
func (s *Store) encounterProgress(ctx context.Context, guildID, encounterID int64, name, zone, counted string) (EncounterProgress, error) {
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
		return EncounterProgress{}, fmt.Errorf("guilds: encounter progress: %w", err)
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
		select r.id, r.created_at::date, count(*), count(*) filter (where f.kill) > 0,
		       min(f.duration_ms) filter (where f.kill), coalesce(sum(f.deaths), 0)::int
		from fights f join reports r on r.id = f.report_id
		where r.guild_id = $1 and f.encounter_id = $2 and `+counted+`
		group by r.id, r.created_at order by r.created_at`,
		guildID, encounterID)
	if err != nil {
		return EncounterProgress{}, fmt.Errorf("guilds: encounter nights: %w", err)
	}
	for nightRows.Next() {
		var n NightPulls
		var date time.Time
		if err := nightRows.Scan(&n.ReportID, &date, &n.Pulls, &n.Killed, &n.KillMS, &n.Deaths); err != nil {
			nightRows.Close()
			return EncounterProgress{}, fmt.Errorf("guilds: encounter nights: %w", err)
		}
		n.Date = date.Format("2006-01-02")
		ep.PullsByNight = append(ep.PullsByNight, n)
	}
	nightRows.Close()
	if err := nightRows.Err(); err != nil {
		return EncounterProgress{}, err
	}

	for _, role := range []string{"tank", "healer", "dps"} {
		ref, ok, err := s.bestByRole(ctx, guildID, encounterID, role)
		if err != nil {
			return EncounterProgress{}, err
		}
		if ok {
			ep.BestByRole[role] = &ref
		}
	}
	return ep, nil
}

// bestByRole is the guild's best kill parse on encounterID for role, from its reports that are
// evidence (presence.EvidenceReport: not private, not flagged); false when it has none.
func (s *Store) bestByRole(ctx context.Context, guildID, encounterID int64, role string) (ParseRef, bool, error) {
	orderCol := roleMetricColumn(role)
	var ref ParseRef
	ref.Metric = roleMetric(role)
	err := s.Pool.QueryRow(ctx, `
		select coalesce(m.player_name, ''), coalesce(m.class, ''), coalesce(m.spec, ''),
		       coalesce(m.`+orderCol+`, 0), m.report_id, m.fight_index
		from fight_metrics m join reports r on r.id = m.report_id
		where r.guild_id = $1 and m.encounter_id = $2 and m.role = $3 and `+bossKillMetrics+`
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
