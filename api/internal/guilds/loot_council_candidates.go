// api/internal/guilds/loot_council_candidates.go
//
// The council's candidates for one drop of one night (design/specs/2026-11-05-web-layout.md
// §7.2): the Loot tab's own two tiers (gain_dps tier 0, ilvl_delta tier 1, lootCandidates), each
// candidate with what it replaces, whether the drop is their BiS pick, attendance over the last
// 4 raids, loot this tier, main or off spec and the officers' votes.
package guilds

import (
	"context"
	"time"
)

// lootAttendanceWindow is the council's attendance window: the last 4 raids (or, before the
// guild has a logged raid, the last 4 reports), the mock's "attended 3 of 4".
const lootAttendanceWindow = 4

// MS/OS sources: the candidate's own (or an officer's) flag, or the derivation rule.
const (
	msOsDeclared = "declared"
	msOsDerived  = "derived"
)

// CouncilCandidate is a Loot tab candidate as the council sees it. Attendance counts the last
// lootAttendanceWindow raids here; ItemsWon is loot this tier (every award
// of a loot table boss).
type CouncilCandidate struct {
	LootCandidate
	// BisForThem is true when the drop is the pick for the slot in the candidate's BiS band, not a
	// listed alternative.
	BisForThem bool `json:"bis_for_them"`
	// MsOs is the candidate's declared flag, else derived: "ms" when their spec's BiS band names the
	// drop (tier 0), "os" otherwise. Null for a quest drop or the bag, which no spec wants.
	MsOs       *string   `json:"ms_os"`
	MsOsSource *string   `json:"ms_os_source"`
	Votes      LootVotes `json:"votes"`
}

// CouncilDropView is the body of GET /v1/guilds/{id}/loot/drops/{item_id}/candidates.
type CouncilDropView struct {
	Night    string `json:"night"`
	ItemID   int    `json:"item_id"`
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Quality  int    `json:"quality"`
	Kind     string `json:"kind"`
	Officers int    `json:"officers"`
	// Closed is true once the drop is awarded on this night: voting and flags are shut.
	Closed     bool               `json:"closed"`
	Award      *LootDecision      `json:"award"`
	Candidates []CouncilCandidate `json:"candidates"`
}

// lootGroupOf is the table's group (items sharing a name) holding itemID.
func lootGroupOf(itemID int) ([]lootItem, bool) {
	for _, b := range lootBosses {
		for _, group := range lootGroups(b.Items) {
			for _, it := range group {
				if it.ItemID == itemID {
					return group, true
				}
			}
		}
	}
	return nil, false
}

// CouncilDrop assembles the council's view of itemID on night for viewer (an officer).
func (s *Store) CouncilDrop(ctx context.Context, guildID int64, night awardNight, itemID int, viewer int64, now time.Time) (CouncilDropView, error) {
	group, ok := lootGroupOf(itemID)
	if !ok {
		return CouncilDropView{}, ErrInvalidAward
	}
	lr, err := s.lootRoster(ctx, guildID)
	if err != nil {
		return CouncilDropView{}, err
	}
	iv := s.lootItemView(lr.roster, group, lr.guild.Faction)
	view := CouncilDropView{Night: night.key(), ItemID: iv.ItemID, Name: iv.Name, Icon: iv.Icon, Quality: iv.Quality,
		Kind: iv.Kind, Candidates: []CouncilCandidate{}}
	cands := iv.Candidates
	if iv.Kind != lootKindGear {
		cands = eligibleCandidates(iv.Eligible)
	}
	if err := s.fillCouncil(ctx, guildID, night, itemID, viewer, now, lr, cands, &view); err != nil {
		return CouncilDropView{}, err
	}
	return view, nil
}

// eligibleCandidates is a quest drop's or the bag's eligible characters as candidate rows.
func eligibleCandidates(eligible []LootEligible) []LootCandidate {
	out := make([]LootCandidate, 0, len(eligible))
	for _, e := range eligible {
		out = append(out, LootCandidate{CharacterKey: e.CharacterKey, Name: e.Name, Class: e.Class, Tier: 1})
	}
	return out
}

// fillCouncil reads the night's awards, flags, votes and attendance and decorates cands.
func (s *Store) fillCouncil(ctx context.Context, guildID int64, night awardNight, itemID int, viewer int64, now time.Time,
	lr lootRosterSet, cands []LootCandidate, view *CouncilDropView) error {
	all, err := s.lootAwards(ctx, guildID, allLoot.encounters)
	if err != nil {
		return err
	}
	lr.nameInvited(all)
	view.Award = nightAward(all, night, itemID, now)
	view.Closed = view.Award != nil
	flags, err := s.dropInterest(ctx, guildID, night, itemID, now)
	if err != nil {
		return err
	}
	votes, err := s.dropVotes(ctx, guildID, night, itemID, viewer, now)
	if err != nil {
		return err
	}
	if view.Officers, err = s.officerCount(ctx, guildID); err != nil {
		return err
	}
	attendance, err := s.councilAttendance(ctx, guildID, lr.roster)
	if err != nil {
		return err
	}
	wonBy := awardCounts(all)
	for _, c := range cands {
		c.ItemsWon = wonBy[c.CharacterKey]
		c.Attendance = attendance[c.CharacterKey]
		cc := CouncilCandidate{LootCandidate: c, BisForThem: c.Tier == 0 && c.bisPick, Votes: votes[c.CharacterKey]}
		cc.Votes.Voters = view.Officers
		cc.MsOs, cc.MsOsSource = candidateMsOs(c, view.Kind, flags[c.CharacterKey])
		view.Candidates = append(view.Candidates, cc)
	}
	return nil
}

// candidateMsOs is the declared flag, else the rule: a gear drop their spec's BiS band names (tier
// 0) is main spec, any other gear drop off spec; a quest drop or the bag has none.
func candidateMsOs(c LootCandidate, kind, declared string) (*string, *string) {
	if declared != "" {
		return stringRef(declared), stringRef(msOsDeclared)
	}
	if kind != lootKindGear {
		return nil, nil
	}
	if c.Tier == 0 {
		return stringRef(msOsMS), stringRef(msOsDerived)
	}
	return stringRef(msOsOS), stringRef(msOsDerived)
}

func stringRef(s string) *string { return &s }

// councilAttendance is attendance over the last lootAttendanceWindow raids by award key: logged
// raids once the guild has one, else its last reports (the roster's own two sources).
func (s *Store) councilAttendance(ctx context.Context, guildID int64, roster []RosterRow) (map[string]Attendance, error) {
	keys := make([]string, 0, len(roster))
	byAward := make(map[string]string, len(roster))
	for _, row := range roster {
		keys = append(keys, row.CharacterKey)
		byAward[awardKeyOf(row)] = row.CharacterKey
	}
	byKey, err := s.attendanceOver(ctx, guildID, keys, lootAttendanceWindow)
	if err != nil {
		return nil, err
	}
	out := make(map[string]Attendance, len(byAward))
	for award, key := range byAward {
		out[award] = byKey[key]
	}
	return out, nil
}

// attendanceOver is attendance per key over the guild's last window nights, by its source.
func (s *Store) attendanceOver(ctx context.Context, guildID int64, keys []string, window int) (map[string]Attendance, error) {
	source, err := s.attendanceSource(ctx, guildID)
	if err != nil {
		return nil, err
	}
	if source == AttendanceSourceRaids {
		return s.raidAttendance(ctx, guildID, keys, window)
	}
	nights, present, err := s.attendance(ctx, guildID, keys, window)
	if err != nil {
		return nil, err
	}
	out := make(map[string]Attendance, len(keys))
	for _, key := range keys {
		out[key] = Attendance{Present: present[key], Nights: nights}
	}
	return out, nil
}
