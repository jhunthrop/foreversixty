// api/internal/guilds/recruiting.go
//
// GET /v1/guilds/recruiting?spec=&faction=&region=&ruleset=: the guilds recruiting a spec, for
// the apps' Raid tab with no guild (spec 2026-11-08-app-flows, Raid, no guild). A guild is listed
// while it recruits (recruit_open), on the faction asked (and the realm, when given), and its
// wanted list names the spec, its class with no spec, or nothing at all (open to everyone).
// Schedule is the officers' raid nights, else the nights and time of its planned raids over the
// next two weeks in its raid time zone ("Wed and Sat, 8:00 pm CST"); needs is the spec's role.
// Public, like the recruiting card.
package guilds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

const (
	// recruitingLimit is the most guilds the list answers.
	recruitingLimit = 20
	// recruitingCacheSeconds is the edge cache.
	recruitingCacheSeconds = 60
)

// recruitingNeeds is the words a spec's role reads as.
var recruitingNeeds = map[string]string{RoleDPS: "damage", RoleHealer: "healer", RoleTank: "tank"}

// RecruitingGuild is one row of the list.
type RecruitingGuild struct {
	GuildID  int64   `json:"guild_id"`
	Name     string  `json:"name"`
	Faction  string  `json:"faction"`
	Schedule string  `json:"schedule"`
	Needs    string  `json:"needs"`
	CrestURL *string `json:"crest_url"`
}

// RecruitingQuery is a validated request.
type RecruitingQuery struct {
	Spec            raidspec.Spec
	Faction         string
	Region, Ruleset string
}

const recruitingSQL = `
select g.id, g.name, g.region, g.ruleset, g.faction, g.crest_key, g.recruit_nights, g.recruit_wanted,
       coalesce(nullif(g.raid_timezone, ''), 'UTC')
from guilds g
where g.recruit_open and g.faction = $1
  and ($2 = '' or g.region = $2) and ($3 = '' or g.ruleset = $3)
order by g.recruit_updated_at desc nulls last, g.id`

// Recruiting is the guilds recruiting q's spec; never nil.
func (s *Store) Recruiting(ctx context.Context, q RecruitingQuery) ([]RecruitingGuild, error) {
	rows, err := s.Pool.Query(ctx, recruitingSQL, q.Faction, q.Region, q.Ruleset)
	if err != nil {
		return nil, fmt.Errorf("guilds: recruiting: %w", err)
	}
	type candidate struct {
		g        Guild
		faction  string
		nights   string
		wanted   []byte
		timezone string
	}
	var cands []candidate
	for rows.Next() {
		var c candidate
		if err := rows.Scan(&c.g.ID, &c.g.Name, &c.g.Region, &c.g.Ruleset, &c.faction, &c.g.CrestKey,
			&c.nights, &c.wanted, &c.timezone); err != nil {
			rows.Close()
			return nil, fmt.Errorf("guilds: recruiting: %w", err)
		}
		cands = append(cands, c)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: recruiting: %w", err)
	}
	out := []RecruitingGuild{}
	for _, c := range cands {
		var wanted []WantedEntry
		if err := json.Unmarshal(c.wanted, &wanted); err != nil {
			return nil, fmt.Errorf("guilds: recruiting wanted %d: %w", c.g.ID, err)
		}
		if !wantsSpec(wanted, q.Spec) {
			continue
		}
		schedule := c.nights
		if schedule == "" {
			if schedule, err = s.plannedSchedule(ctx, c.g.ID, c.timezone); err != nil {
				return nil, err
			}
		}
		out = append(out, RecruitingGuild{GuildID: c.g.ID, Name: c.g.Name, Faction: c.faction, Schedule: schedule,
			Needs: recruitingNeeds[q.Spec.Role], CrestURL: CrestURL(s.APIBaseURL, c.g.ID, c.g.CrestKey)})
		if len(out) == recruitingLimit {
			break
		}
	}
	return out, nil
}

// wantsSpec reports whether a wanted list takes sp: empty, its class with no spec, or sp itself.
func wantsSpec(wanted []WantedEntry, sp raidspec.Spec) bool {
	if len(wanted) == 0 {
		return true
	}
	for _, w := range wanted {
		if w.Class == sp.Class && (w.Spec == nil || *w.Spec == sp.Key) {
			return true
		}
	}
	return false
}

// plannedSchedule words the guild's planned raids over the next two weeks: their weekdays in
// week order, then the earliest raid's time, in tz; "" with none planned.
func (s *Store) plannedSchedule(ctx context.Context, guildID int64, tz string) (string, error) {
	rows, err := s.Pool.Query(ctx, `select starts_at from guild_raids
		where guild_id = $1 and status in ('planned', 'locked')
		  and starts_at > now() and starts_at < now() + make_interval(days => $2)
		order by starts_at limit $3`, guildID, recruitUpcomingDays, recruitUpcomingMax)
	if err != nil {
		return "", fmt.Errorf("guilds: recruiting schedule: %w", err)
	}
	defer rows.Close()
	loc := raidLocation(tz)
	var starts []time.Time
	for rows.Next() {
		var t time.Time
		if err := rows.Scan(&t); err != nil {
			return "", fmt.Errorf("guilds: recruiting schedule: %w", err)
		}
		starts = append(starts, t.In(loc))
	}
	if err := rows.Err(); err != nil || len(starts) == 0 {
		return "", err
	}
	return scheduleWords(starts), nil
}

// scheduleWords is "Wed and Sat, 8:00 pm CST": the distinct weekdays Monday first, then the
// first start's time.
func scheduleWords(starts []time.Time) string {
	seen := map[time.Weekday]bool{}
	var days []time.Weekday
	for _, t := range starts {
		if !seen[t.Weekday()] {
			seen[t.Weekday()] = true
			days = append(days, t.Weekday())
		}
	}
	monday := func(d time.Weekday) int { return (int(d) + 6) % 7 }
	sort.Slice(days, func(i, j int) bool { return monday(days[i]) < monday(days[j]) })
	words := make([]string, len(days))
	for i, d := range days {
		words[i] = d.String()[:3]
	}
	list := words[0]
	if n := len(words); n > 1 {
		list = strings.Join(words[:n-1], ", ") + " and " + words[n-1]
	}
	return list + ", " + starts[0].Format(confirmationTimeLayout)
}

// recruitingQuery validates the request; it writes the 400 itself when a parameter is bad.
func recruitingQuery(w http.ResponseWriter, r *http.Request) (RecruitingQuery, bool) {
	p := r.URL.Query()
	sp, ok := raidspec.ByKey(p.Get("spec"))
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "spec must be a spec key such as warrior-fury",
			map[string]string{"spec": "a specs.json key"})
		return RecruitingQuery{}, false
	}
	q := RecruitingQuery{Spec: sp, Faction: strings.ToLower(p.Get("faction"))}
	if q.Faction != "alliance" && q.Faction != "horde" {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "faction must be alliance or horde",
			map[string]string{"faction": "alliance, horde"})
		return q, false
	}
	q.Region, q.Ruleset, ok = realmFilter(w, r)
	return q, ok
}

// recruiting is GET /v1/guilds/recruiting.
func (s *Service) recruiting(w http.ResponseWriter, r *http.Request) {
	q, ok := recruitingQuery(w, r)
	if !ok {
		return
	}
	rows, err := s.Store.Recruiting(r.Context(), q)
	if err != nil {
		s.fail(w, r, "recruiting", err, "could not list the recruiting guilds just now")
		return
	}
	httpx.CachePublic(w, recruitingCacheSeconds*time.Second, 0)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"rows": rows})
}
