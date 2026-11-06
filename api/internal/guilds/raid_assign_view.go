// api/internal/guilds/raid_assign_view.go
//
// The raid view's assignments block (spec 2026-11-05-web-layout §7.2): each encounter of the
// raid's zone with its template's slots and the names the raid leader picked, the raid warning
// lines a leader pastes into /rw, and the raids the client has no bosses for yet. A picked
// character whose sign-up is no longer In, Late or Bench drops out of every reading (the row
// stays, so a change of heart puts it back).
package guilds

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// maxRaidWarningBytes is the client's chat line limit, which /rw shares.
const maxRaidWarningBytes = 255

// RaidAssignments is RaidView.assignments: null for a run that is not a raid.
type RaidAssignments struct {
	Encounters []EncounterAssignments `json:"encounters"`
	// NotInClient names the raid when the game client carries none of its bosses yet.
	NotInClient []string `json:"not_in_client"`
	// Candidates is every character signed In, Late or Bench, the pickers' list.
	Candidates []AssignmentRaider `json:"candidates"`
	// PostedAt is when the assignments last went to Discord; null before the first post.
	PostedAt *time.Time `json:"posted_at"`
}

// EncounterAssignments is one boss's template with its names.
type EncounterAssignments struct {
	EncounterID int64            `json:"encounter_id"`
	Name        string           `json:"name"`
	RaidZone    string           `json:"raid_zone"`
	Slots       []SlotAssignment `json:"slots"`
	// RaidWarning is one short line per filled slot, ready for /rw.
	RaidWarning []string `json:"raid_warning"`
}

// SlotAssignment is one template slot and who holds it.
type SlotAssignment struct {
	Key     string         `json:"key"`
	Label   string         `json:"label"`
	Purpose string         `json:"purpose"`
	Count   int            `json:"count"`
	Role    string         `json:"role"`
	Names   []AssignedName `json:"names"`
}

// AssignmentRaider is a signed-up character as the assignments name it. Own is true for the
// viewer's own characters.
type AssignmentRaider struct {
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
	Group        string `json:"group"`
	State        string `json:"state"`
	Own          bool   `json:"own"`
}

// AssignedName is a slot's pick. OffRole is true when the slot is for a role the character
// does not fill tonight (a healer in a tank slot): a warning, never a refusal.
type AssignedName struct {
	AssignmentRaider
	OffRole bool `json:"off_role"`
}

// assignable reports whether a sign-up state may hold an assignment.
func assignable(state string) bool {
	return state == SignupIn || state == SignupLate || state == SignupBench
}

// groupRank orders the pickers: tanks, healers, melee, ranged.
var groupRank = map[string]int{raidspec.GroupTank: 0, raidspec.GroupHealer: 1, raidspec.GroupMelee: 2, raidspec.GroupRanged: 3}

// assignmentCandidates is every assignable sign-up, tanks first, then by name.
func assignmentCandidates(signups []Signup, vc raidViewContext) []AssignmentRaider {
	out := []AssignmentRaider{}
	for _, sg := range signups {
		if !assignable(sg.State) {
			continue
		}
		out = append(out, AssignmentRaider{
			CharacterKey: sg.CharacterKey, Name: sg.Name, Class: sg.Class,
			Group: raidspec.Group(sg.Class, sg.Spec, sg.Role), State: sg.State, Own: vc.owns(sg),
		})
	}
	sort.SliceStable(out, func(i, j int) bool {
		if a, b := groupRank[out[i].Group], groupRank[out[j].Group]; a != b {
			return a < b
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// buildAssignments is the block for a raid zone; nil for any other run.
func buildAssignments(zone Zone, stored map[slotRef][]string, candidates []AssignmentRaider, posted *time.Time) *RaidAssignments {
	if !zone.IsRaid() {
		return nil
	}
	out := &RaidAssignments{Encounters: []EncounterAssignments{}, NotInClient: []string{}, Candidates: candidates, PostedAt: posted}
	raid, ok := tierRaidOf(zone.Key)
	if !ok || len(raid.Bosses) == 0 {
		out.NotInClient = append(out.NotInClient, zone.Name)
		return out
	}
	byKey := make(map[string]AssignmentRaider, len(candidates))
	for _, c := range candidates {
		byKey[c.CharacterKey] = c
	}
	for _, boss := range raid.Bosses {
		if slots, has := templateOf(boss.EncounterID); has {
			out.Encounters = append(out.Encounters, encounterAssignments(raid.Name, boss, slots, stored, byKey))
		}
	}
	return out
}

func encounterAssignments(raidName string, boss tierBoss, slots []slotTemplate, stored map[slotRef][]string, byKey map[string]AssignmentRaider) EncounterAssignments {
	e := EncounterAssignments{EncounterID: boss.EncounterID, Name: boss.Name, RaidZone: raidName,
		Slots: make([]SlotAssignment, len(slots)), RaidWarning: []string{}}
	for i, t := range slots {
		names := []AssignedName{}
		for _, key := range stored[slotRef{EncounterID: boss.EncounterID, SlotKey: t.Key}] {
			if c, ok := byKey[key]; ok {
				names = append(names, AssignedName{AssignmentRaider: c, OffRole: offRole(t.Role, c.Group)})
			}
		}
		e.Slots[i] = SlotAssignment{Key: t.Key, Label: t.Label, Purpose: t.Purpose, Count: t.Count, Role: t.Role, Names: names}
		if line := raidWarningLine(t, names); line != "" {
			e.RaidWarning = append(e.RaidWarning, line)
		}
	}
	return e
}

// offRole reports whether a raider of group is outside a slot for role.
func offRole(role, group string) bool {
	return role != slotRoleAny && role != group
}

// assignedNames is the slot's names as one comma list.
func assignedNames(names []AssignedName) string {
	parts := make([]string, len(names))
	for i, n := range names {
		parts[i] = n.Name
	}
	return strings.Join(parts, ", ")
}

// raidWarningLine is "MT: Thornhide", cut to the chat limit on a whole name; empty for an empty slot.
func raidWarningLine(t slotTemplate, names []AssignedName) string {
	if len(names) == 0 {
		return ""
	}
	line := t.Warning + ": " + names[0].Name
	for _, n := range names[1:] {
		next := line + ", " + n.Name
		if len(next) > maxRaidWarningBytes {
			break
		}
		line = next
	}
	return line
}

// fillAssignments sets every raid view's assignments from the stored picks and its sign-ups.
func (s *Store) fillAssignments(ctx context.Context, vc raidViewContext, views []RaidView, byRaid map[int64][]Signup) error {
	ids := raidIDsOf(views)
	stored, err := s.assignmentsByRaid(ctx, ids)
	if err != nil {
		return err
	}
	posts, err := s.lastAssignmentPosts(ctx, ids)
	if err != nil {
		return err
	}
	for i := range views {
		var posted *time.Time
		if at, ok := posts[views[i].ID]; ok {
			utc := at.UTC()
			posted = &utc
		}
		candidates := assignmentCandidates(byRaid[views[i].ID], vc)
		views[i].Assignments = buildAssignments(zoneOf(views[i].Zone), stored[views[i].ID], candidates, posted)
	}
	return nil
}
