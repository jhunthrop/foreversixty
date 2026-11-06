// api/internal/guilds/me_checklist_handler.go
//
// GET /v1/guilds/{id}/me/checklist?raid={raidId}: the signed-in member's checklist
// (me_checklist.go) for the raid (default: the next raid, else the next run), for the character
// they signed with, else their main, else their first verified character by name.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const checklistFailMessage = "could not load your checklist just now"

func (s *Service) meChecklist(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	viewer := auth.ActorFrom(r.Context()).UserID
	member, err := s.Store.IsVerifiedMember(r.Context(), guildID, viewer)
	if err != nil {
		s.fail(w, r, "me_checklist", err, checklistFailMessage)
		return
	}
	if !member {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you are not a member of that guild", nil)
		return
	}
	raid, ok := s.checklistRaid(w, r, guildID)
	if !ok {
		return
	}
	view, found, err := s.Store.Checklist(r.Context(), guildID, viewer, raid)
	if err != nil {
		s.fail(w, r, "me_checklist", err, checklistFailMessage)
		return
	}
	if !found {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "you have no verified character in that guild", nil)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// checklistRaid is the ?raid= raid of the guild (404 when it has none such), else the next raid,
// else the next run; nil when nothing is planned. ok is false once it has answered.
func (s *Service) checklistRaid(w http.ResponseWriter, r *http.Request, guildID int64) (*Raid, bool) {
	ctx := r.Context()
	if q := r.URL.Query().Get("raid"); q != "" {
		id, err := strconv.ParseInt(q, 10, 64)
		if err != nil || id <= 0 {
			writeRaidNotFound(w, r)
			return nil, false
		}
		raid, found, err := s.Store.GetRaid(ctx, guildID, id)
		if !s.answeredMemberRaid(w, r, "me_checklist", found, err) {
			return nil, false
		}
		return &raid, true
	}
	now := time.Now()
	raid, found, err := s.Store.NextRaid(ctx, guildID, now)
	if err == nil && !found {
		raid, found, err = s.Store.NextPlanned(ctx, guildID, now)
	}
	if err != nil {
		s.fail(w, r, "me_checklist", err, checklistFailMessage)
		return nil, false
	}
	if !found {
		return nil, true
	}
	return &raid, true
}

// Checklist is viewerID's checklist in the guild for raid (nil: none planned); found is false
// when the viewer has no verified character there.
func (s *Store) Checklist(ctx context.Context, guildID, viewerID int64, raid *Raid) (ChecklistView, bool, error) {
	char, found, err := s.checklistCharacter(ctx, guildID, viewerID, raid)
	if err != nil || !found {
		return ChecklistView{}, false, err
	}
	roster, err := s.readinessRoster(ctx, guildID)
	if err != nil {
		return ChecklistView{}, false, err
	}
	in := checklistInput{}
	for _, row := range roster {
		if row.CharacterKey == char.CharacterKey {
			in.row, in.onRoster = row, true
			in.band, in.hasBand = s.loadBandFor(row.className, row.specStr, row.faction)
			in.readiness = computeReadiness(readinessInputOf(row, in.band, in.hasBand, medianItemLevel(roster)))
			break
		}
	}
	view := ChecklistView{Character: char}
	if raid != nil {
		zone := zoneOf(raid.Zone)
		in.zone = &zone
		view.Raid = &ChecklistRaid{ID: raid.ID, ZoneKey: zone.Key, ZoneName: zone.Name}
	}
	view.Character.Level = in.readiness.Level
	view.Items, view.Verdict = buildChecklist(in)
	return view, true, nil
}

// checklistCharacter is the viewer's character for the raid: the one signed with (its best
// place), else the account's main, else the first verified character by name.
func (s *Store) checklistCharacter(ctx context.Context, guildID, viewerID int64, raid *Raid) (ChecklistCharacter, bool, error) {
	roster, err := s.raidRoster(ctx, guildID, "")
	if err != nil {
		return ChecklistCharacter{}, false, err
	}
	own := map[string]raidRosterChar{}
	for _, c := range roster {
		if c.userID == viewerID && c.verified {
			own[c.CharacterKey] = c
		}
	}
	if len(own) == 0 {
		return ChecklistCharacter{}, false, nil
	}
	key, err := s.checklistKey(ctx, viewerID, raid, own)
	if err != nil {
		return ChecklistCharacter{}, false, err
	}
	c := own[key]
	return ChecklistCharacter{CharacterKey: c.CharacterKey, Name: c.Name, Class: c.Class, Spec: c.Spec}, true, nil
}

// checklistKey picks among own (non-empty) by checklistCharacter's rule.
func (s *Store) checklistKey(ctx context.Context, viewerID int64, raid *Raid, own map[string]raidRosterChar) (string, error) {
	if raid != nil {
		signups, err := s.RaidSignups(ctx, raid.ID)
		if err != nil {
			return "", err
		}
		if key := bestOwnSignup(signups, own); key != "" {
			return key, nil
		}
	}
	var main *string
	if err := s.Pool.QueryRow(ctx, `select main_character_key from users where id = $1`, viewerID).Scan(&main); err != nil {
		return "", fmt.Errorf("guilds: checklist main: %w", err)
	}
	if main != nil {
		if _, ok := own[*main]; ok {
			return *main, nil
		}
	}
	keys := make([]string, 0, len(own))
	for k := range own {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return own[keys[i]].Name < own[keys[j]].Name })
	return keys[0], nil
}

// bestOwnSignup is the key of the best place (In, Late, Bench, Out) among own's sign-ups; ""
// when none of own signed.
func bestOwnSignup(signups []Signup, own map[string]raidRosterChar) string {
	best, rankOf := "", len(placeRank)
	for _, sg := range signups {
		if _, ok := own[sg.CharacterKey]; !ok {
			continue
		}
		if r := placeRank[sg.State]; r < rankOf {
			best, rankOf = sg.CharacterKey, r
		}
	}
	return best
}
