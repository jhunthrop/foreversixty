// api/internal/guilds/readiness_raid.go
//
// Lane G2: the readiness board cut to one raid (design/specs/2026-10-27-guild-attendance.md
// §3.7). GET …/readiness?raid={id} keeps the raid's In and Late sign-ups, judges each export's
// freshness against the raid's start, and lists a signed character with no export as a
// no_export row built from its sign-up snapshot. The unfiltered read adds next_raid.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// readinessFreshDays is READINESS_FRESH_DAYS: an export synced longer than this before the
// raid's start is stale and counts as one failing check.
const readinessFreshDays = 7

const hoursPerDay = 24

// noExportScore sorts a no_export row below every row that has an export.
const noExportScore = -1

// ReadinessRaid is the raid a filtered board is cut to. CutoffAt is the raid's start: the
// moment each export is judged against.
type ReadinessRaid struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	StartsAt time.Time `json:"starts_at"`
	CutoffAt time.Time `json:"cutoff_at"`
}

// ReadinessNextRaid is the unfiltered board's pointer to the guild's next raid.
type ReadinessNextRaid struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	ZoneName string    `json:"zone_name"`
	StartsAt time.Time `json:"starts_at"`
}

// readinessWithNext is the unfiltered board: next_raid is always present, null with no raid.
type readinessWithNext struct {
	ReadinessView
	NextRaid *ReadinessNextRaid `json:"next_raid"`
}

// readinessScope is the raid a filtered read is cut to and who signed In or Late. A nil scope
// is the unfiltered board: every method then leaves the board as it was.
type readinessScope struct {
	raid   Raid
	signed map[string]Signup
}

func (s *Store) newReadinessScope(ctx context.Context, raid *Raid) (*readinessScope, error) {
	if raid == nil {
		return nil, nil
	}
	signups, err := s.RaidSignups(ctx, raid.ID)
	if err != nil {
		return nil, err
	}
	sc := &readinessScope{raid: *raid, signed: map[string]Signup{}}
	for _, sg := range signups {
		if sg.State == SignupIn || sg.State == SignupLate {
			sc.signed[sg.CharacterKey] = sg
		}
	}
	return sc, nil
}

func (sc *readinessScope) includes(key string) bool {
	if sc == nil {
		return true
	}
	_, ok := sc.signed[key]
	return ok
}

func (sc *readinessScope) view() *ReadinessRaid {
	if sc == nil {
		return nil
	}
	start := sc.raid.StartsAt.UTC()
	return &ReadinessRaid{ID: sc.raid.ID, Title: sc.raid.Title, StartsAt: start, CutoffAt: start}
}

// fresh reports whether an export synced at loggedAt is within the fresh window of the raid.
func (sc *readinessScope) fresh(loggedAt time.Time) bool {
	return !loggedAt.Before(sc.raid.StartsAt.AddDate(0, 0, -readinessFreshDays))
}

// apply judges the row's export against the raid: a stale one adds a failing check.
func (sc *readinessScope) apply(rr *ReadinessRow) {
	if sc == nil {
		return
	}
	fresh := sc.fresh(rr.LoggedAt)
	rr.SyncedBeforeCutoff = &fresh
	if !fresh {
		rr.Failing++
	}
}

// staleNudge is the sentence an officer's nudge gains for a stale export.
func (sc *readinessScope) staleNudge(rr ReadinessRow) string {
	if sc == nil || rr.SyncedBeforeCutoff == nil || *rr.SyncedBeforeCutoff {
		return ""
	}
	days := int(sc.raid.StartsAt.Sub(rr.LoggedAt).Hours() / hoursPerDay)
	return fmt.Sprintf(" Last sync was %d days before the raid.", days)
}

// noExportRows lists the signed characters the board had no export row for, last.
func (sc *readinessScope) noExportRows(seen map[string]bool) []ReadinessRow {
	if sc == nil {
		return nil
	}
	var rows []ReadinessRow
	for key, sg := range sc.signed {
		if seen[key] {
			continue
		}
		rows = append(rows, ReadinessRow{CharacterKey: key, Name: sg.Name, Class: sg.Class, Spec: sg.Spec,
			Consent: "none", Enchants: EnchantsView{MissingSlots: []string{}}, Consumables: ConsumablesView{State: "unknown"},
			NoExport: true, readinessScore: noExportScore})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Name < rows[j].Name })
	return rows
}

func (s *Store) readinessNextRaid(ctx context.Context, guildID int64, now time.Time) (*ReadinessNextRaid, error) {
	raid, ok, err := s.NextRaid(ctx, guildID, now)
	if err != nil || !ok {
		return nil, err
	}
	return &ReadinessNextRaid{ID: raid.ID, Title: raid.Title, ZoneName: zoneOf(raid.Zone).Name, StartsAt: raid.StartsAt.UTC()}, nil
}

// readinessRaid reads the optional raid query: filtered is false with no raid param; ok is
// false once the request is answered (404 for a raid that is not this guild's).
func (s *Service) readinessRaid(w http.ResponseWriter, r *http.Request, guildID int64) (raid *Raid, filtered, ok bool) {
	raw := r.URL.Query().Get("raid")
	if raw == "" {
		return nil, false, true
	}
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id <= 0 {
		writeRaidNotFound(w, r)
		return nil, true, false
	}
	found, exists, err := s.Store.GetRaid(r.Context(), guildID, id)
	if err != nil {
		s.fail(w, r, "readiness", err, "could not load that guild's readiness just now")
		return nil, true, false
	}
	if !exists {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such raid", nil)
		return nil, true, false
	}
	return &found, true, true
}
