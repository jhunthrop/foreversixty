// api/internal/guilds/raid_night.go
//
// GET /v1/me/raid-nights/{raid}: one raider's raid night (spec 2026-11-08-app-flows F2, M2).
// Before the pull it answers the countdown (starts_at, ends_at), your answer, who is coming by
// role with the late notes, and your ready check; once a log is matched to the raid, the night's
// kills and wipes, the loot you got, your bests on its kills and the coach teaser. Fields a log
// would fill are null until there is one. A verified member of the raid's guild reads it, for
// the character they signed with (else their main, else their first verified one, the D12
// checklist's pick); anyone else, or a raid that does not exist, reads 404.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// The phases of a raid night.
const (
	RaidNightBefore = "before"
	RaidNightLive   = "live"
	RaidNightAfter  = "after"
)

// raidNightLive is how long after the start the night reads as live (spec F2: start to start
// plus 4 hours).
const raidNightLive = 4 * time.Hour

const raidNightFail = "could not load your raid night just now"

// RaidNightCoach reads the coach teaser for one character in a report (coach.Service through an
// adapter); nil notes for an account without the log coach.
type RaidNightCoach interface {
	CoachNotes(ctx context.Context, userID int64, reportID, characterKey string) ([]string, error)
}

// RaidNightSummary is the body of GET /v1/me/raid-nights/{raid}.
type RaidNightSummary struct {
	RaidID     int64              `json:"raid_id"`
	GuildID    int64              `json:"guild_id"`
	GuildName  string             `json:"guild_name"`
	ZoneKey    string             `json:"zone_key"`
	ZoneName   string             `json:"zone_name"`
	ZoneIcon   string             `json:"zone_icon"`
	StartsAt   time.Time          `json:"starts_at"`
	EndsAt     time.Time          `json:"ends_at"`
	Status     string             `json:"status"`
	Phase      string             `json:"phase"`
	Character  RaidNightCharacter `json:"character"`
	You        RaidNightYou       `json:"you"`
	Coming     RaidNightComing    `json:"coming"`
	Ready      *ReadyCheck        `json:"ready"`
	ReportID   *string            `json:"report_id"`
	Kills      *int               `json:"kills"`
	Wipes      *int               `json:"wipes"`
	NewBests   *int               `json:"new_bests"`
	Loot       []RaidNightLoot    `json:"loot"`
	Bests      []RaidNightBest    `json:"bests"`
	CoachNotes []string           `json:"coach_notes"`
}

// RaidNightCharacter is the character the night is read for.
type RaidNightCharacter struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Spec         string `json:"spec"`
	SpecName     string `json:"spec_name"`
	Role         string `json:"role"`
}

// RaidNightYou is your answer for the raid, nulls until you answer.
type RaidNightYou struct {
	State       *string `json:"state"`
	LateMinutes *int16  `json:"late_minutes"`
	Note        string  `json:"note"`
}

// RaidNightComing is the raid's answers: the counts, and who is In or Late by role.
type RaidNightComing struct {
	In      int               `json:"in"`
	Late    int               `json:"late"`
	Bench   int               `json:"bench"`
	Out     int               `json:"out"`
	Tanks   []RaidNightPerson `json:"tanks"`
	Healers []RaidNightPerson `json:"healers"`
	Damage  []RaidNightPerson `json:"damage"`
}

// RaidNightPerson is one raider coming: In, or Late with their minutes and note.
type RaidNightPerson struct {
	Name        string `json:"name"`
	Class       string `json:"class"`
	State       string `json:"state"`
	LateMinutes *int16 `json:"late_minutes"`
	Note        string `json:"note"`
}

// RaidNightLoot is one item you got that night.
type RaidNightLoot struct {
	ItemID    int               `json:"item_id"`
	Name      string            `json:"name"`
	Quality   int               `json:"quality"`
	Icon      string            `json:"icon"`
	Slot      string            `json:"slot"`
	Boss      string            `json:"boss"`
	AwardedBy string            `json:"awarded_by"`
	Replaces  *RaidNightReplace `json:"replaces"`
	// Gain is the upgrade the item is over the worn gear, the Loot tab's gain (bisMatch) in
	// Unit (DPS, HPS or TPS); both null when the character's band cannot say.
	Gain *float64 `json:"gain"`
	Unit *string  `json:"unit"`
}

// RaidNightReplace is the item a drop replaced; the night does not know it yet, so it is null.
type RaidNightReplace struct {
	ItemID    int    `json:"item_id"`
	Name      string `json:"name"`
	Quality   int    `json:"quality"`
	ItemLevel int    `json:"item_level"`
}

// RaidNightBest is your figure on one boss kill of the night: DPS, or HPS for a healer, the
// fight's Forever rating (0 to 100, null when unrated) and the kill's length.
type RaidNightBest struct {
	Boss       string  `json:"boss"`
	DPS        float64 `json:"dps"`
	Percentile *int    `json:"percentile"`
	KillMs     int64   `json:"kill_ms"`
	newBest    bool
}

func (s *Service) raidNight(w http.ResponseWriter, r *http.Request) {
	raidID, err := strconv.ParseInt(r.PathValue("raid"), 10, 64)
	if err != nil || raidID <= 0 {
		writeRaidNotFound(w, r)
		return
	}
	v, found, err := s.RaidNightOf(r.Context(), auth.ActorFrom(r.Context()).UserID, raidID, time.Now())
	if err != nil {
		s.fail(w, r, "raid_night", err, raidNightFail)
		return
	}
	if !found {
		writeRaidNotFound(w, r)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, v)
}

// RaidNightOf is viewerID's night for raidID at now; found is false when the raid does not exist
// or is not theirs to read.
func (s *Service) RaidNightOf(ctx context.Context, viewerID, raidID int64, now time.Time) (RaidNightSummary, bool, error) {
	guildID, found, err := s.Store.raidGuildID(ctx, raidID)
	if err != nil || !found {
		return RaidNightSummary{}, false, err
	}
	raid, found, err := s.memberRaidOf(ctx, guildID, viewerID, raidID)
	if err != nil || !found {
		return RaidNightSummary{}, false, err
	}
	char, found, err := s.Store.checklistCharacter(ctx, guildID, viewerID, &raid)
	if err != nil || !found {
		return RaidNightSummary{}, false, err
	}
	g, err := s.Store.getGuild(ctx, guildID)
	if err != nil {
		return RaidNightSummary{}, false, err
	}
	view, err := s.raidViewOf(ctx, guildID, viewerID, raid)
	if err != nil {
		return RaidNightSummary{}, false, err
	}
	out := raidNightHead(g, raid, view, now)
	out.Character = raidNightCharacter(char)
	out.You, out.Coming = raidNightAnswers(view, char.CharacterKey)
	if out.Phase != RaidNightAfter {
		zone := zoneOf(raid.Zone)
		check, err := s.Store.ReadyCheckOf(ctx, char.CharacterKey, &zone)
		if err != nil {
			return RaidNightSummary{}, false, err
		}
		out.Ready = &check
	}
	if raid.ReportID != nil {
		if err := s.fillRaidNightLog(ctx, &out, viewerID, guildID, *raid.ReportID, char); err != nil {
			return RaidNightSummary{}, false, err
		}
	}
	return out, true, nil
}

// raidNightHead is the night's raid line and phase at now.
func raidNightHead(g Guild, raid Raid, view RaidView, now time.Time) RaidNightSummary {
	start := raid.StartsAt.UTC()
	out := RaidNightSummary{
		RaidID: raid.ID, GuildID: g.ID, GuildName: g.Name, ZoneKey: raid.Zone, ZoneName: view.ZoneName, ZoneIcon: view.ZoneIcon,
		StartsAt: start, EndsAt: start.Add(time.Duration(raid.DurationMinutes) * time.Minute), Status: raid.Status,
		Loot: []RaidNightLoot{}, Bests: []RaidNightBest{},
	}
	switch {
	case now.Before(start):
		out.Phase = RaidNightBefore
	case now.Before(start.Add(raidNightLive)) && raid.Status != RaidDone && raid.Status != RaidCancelled:
		out.Phase = RaidNightLive
	default:
		out.Phase = RaidNightAfter
	}
	return out
}

// raidNightCharacter is the checklist's character with its spec resolved.
func raidNightCharacter(c ChecklistCharacter) RaidNightCharacter {
	out := RaidNightCharacter{CharacterKey: c.CharacterKey, Name: c.Name, Class: c.Class, SpecName: c.Spec}
	if sp, ok := raidspec.Resolve(c.Class, c.Spec, ""); ok {
		out.Spec, out.SpecName, out.Role = sp.Key, sp.Name, sp.Role
	}
	return out
}

// raidNightAnswers is your answer (your character's row, else your account's) and who is coming.
func raidNightAnswers(view RaidView, key string) (RaidNightYou, RaidNightComing) {
	you := RaidNightYou{}
	coming := RaidNightComing{In: view.Counts.In, Late: view.Counts.Late, Bench: view.Counts.Bench, Out: view.Counts.Out,
		Tanks: []RaidNightPerson{}, Healers: []RaidNightPerson{}, Damage: []RaidNightPerson{}}
	for _, sg := range view.Signups {
		if sg.CharacterKey == key || (sg.Own && you.State == nil) {
			state := sg.State
			you = RaidNightYou{State: &state, LateMinutes: sg.LateMinutes, Note: sg.Note}
		}
		if !isAttending(sg.State) {
			continue
		}
		p := RaidNightPerson{Name: sg.Name, Class: sg.Class, State: sg.State, LateMinutes: sg.LateMinutes, Note: sg.Note}
		switch sg.Role {
		case RoleTank:
			coming.Tanks = append(coming.Tanks, p)
		case RoleHealer:
			coming.Healers = append(coming.Healers, p)
		default:
			coming.Damage = append(coming.Damage, p)
		}
	}
	return you, coming
}

// raidGuildID is the guild a raid belongs to; found is false for no such raid.
func (s *Store) raidGuildID(ctx context.Context, raidID int64) (int64, bool, error) {
	var id int64
	err := s.Pool.QueryRow(ctx, `select guild_id from guild_raids where id = $1`, raidID).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("guilds: raid guild: %w", err)
	}
	return id, true, nil
}
