// api/internal/guilds/loot_council_votes.go
//
// Loot council votes (decision D13): an officer votes yes or no on a candidate for a drop of a
// night, one vote per officer per candidate, changeable and removable. Votes are officer-only to
// read and write, and awarding the drop closes its voting.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Vote and MS/OS words, as stored and as the API carries them.
const (
	voteYes = "yes"
	voteNo  = "no"
	msOsMS  = "ms"
	msOsOS  = "os"
)

// ErrVotingClosed is a vote or a flag on a drop that is already awarded on that night.
var ErrVotingClosed = errors.New("guilds: loot voting closed")

// LootVotes is one candidate's tally: yes and no votes cast, the guild's officer count (who can
// vote), and the viewing officer's own vote (null when they have not voted).
type LootVotes struct {
	Yes    int     `json:"yes"`
	No     int     `json:"no"`
	Voters int     `json:"voters"`
	Mine   *string `json:"mine"`
}

// officerCount is how many verified officers and leaders the guild has: the votes' "of N".
func (s *Store) officerCount(ctx context.Context, guildID int64) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `
		select count(*) from guild_members
		where guild_id = $1 and verified_at is not null and rank in ('officer', 'leader')`, guildID).Scan(&n); err != nil {
		return 0, fmt.Errorf("guilds: officer count: %w", err)
	}
	return n, nil
}

// dropVotes is each candidate's tally for itemID on night, keyed by character key, with viewer's
// own vote. Voters is left for the caller (officerCount).
func (s *Store) dropVotes(ctx context.Context, guildID int64, night awardNight, itemID int, viewer int64, now time.Time) (map[string]LootVotes, error) {
	rows, err := s.Pool.Query(ctx, `
		select v.character_key, v.voter_id, v.vote from loot_votes v
		where v.guild_id = $1 and v.encounter_id = $5 and v.item_id = $6 and `+nightMatchSQL("v", "$2", "$3", "$4", "voted_at"),
		append([]any{guildID}, append(night.args(now), onyxiaEncounterID, itemID)...)...)
	if err != nil {
		return nil, fmt.Errorf("guilds: loot votes: %w", err)
	}
	defer rows.Close()
	out := map[string]LootVotes{}
	for rows.Next() {
		var key, vote string
		var voter int64
		if err := rows.Scan(&key, &voter, &vote); err != nil {
			return nil, fmt.Errorf("guilds: loot votes: %w", err)
		}
		out[key] = tallyVote(out[key], vote, voter == viewer)
	}
	return out, rows.Err()
}

func tallyVote(t LootVotes, vote string, mine bool) LootVotes {
	if vote == voteYes {
		t.Yes++
	} else {
		t.No++
	}
	if mine {
		v := vote
		t.Mine = &v
	}
	return t
}

// checkDropTarget is ErrInvalidAward unless itemID is one of Onyxia's drops and characterKey is
// on the guild's roster.
func (s *Store) checkDropTarget(ctx context.Context, guildID int64, itemID int, characterKey string) error {
	if !validOnyxiaItem(itemID) {
		return ErrInvalidAward
	}
	var exists bool
	if err := s.Pool.QueryRow(ctx,
		`select exists(select 1 from guild_characters where guild_id = $1 and character_key = $2)`,
		guildID, characterKey).Scan(&exists); err != nil {
		return fmt.Errorf("guilds: loot council: %w", err)
	}
	if !exists {
		return ErrInvalidAward
	}
	return nil
}

// dropAwarded reports whether itemID already has an award filed under night.
func (s *Store) dropAwarded(ctx context.Context, guildID int64, night awardNight, itemID int, now time.Time) (bool, error) {
	var awarded bool
	err := s.Pool.QueryRow(ctx, `
		select exists(select 1 from loot_awards la
		  where la.guild_id = $1 and la.encounter_id = $5 and la.item_id = $6 and `+nightMatchSQL("la", "$2", "$3", "$4", "awarded_at")+`)`,
		append([]any{guildID}, append(night.args(now), onyxiaEncounterID, itemID)...)...).Scan(&awarded)
	if err != nil {
		return false, fmt.Errorf("guilds: loot council: %w", err)
	}
	return awarded, nil
}

// openDrop is checkDropTarget plus ErrVotingClosed once the drop is awarded on night.
func (s *Store) openDrop(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey string, now time.Time) error {
	if err := s.checkDropTarget(ctx, guildID, itemID, characterKey); err != nil {
		return err
	}
	awarded, err := s.dropAwarded(ctx, guildID, night, itemID, now)
	if err != nil {
		return err
	}
	if awarded {
		return ErrVotingClosed
	}
	return nil
}

// SetLootVote records voter's yes or no on characterKey for itemID on night, replacing their
// earlier vote; vote "" removes it.
func (s *Store) SetLootVote(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey string, voter int64, vote string, now time.Time) error {
	if err := s.openDrop(ctx, guildID, night, itemID, characterKey, now); err != nil {
		return err
	}
	if vote == "" {
		_, err := s.Pool.Exec(ctx, `
			delete from loot_votes
			where guild_id = $1 and coalesce(raid_id, 0) = coalesce($2::bigint, 0) and coalesce(report_id, '') = coalesce($3::text, '')
			  and encounter_id = $4 and item_id = $5 and character_key = $6 and voter_id = $7`,
			guildID, night.raidID, night.reportID, onyxiaEncounterID, itemID, characterKey, voter)
		if err != nil {
			return fmt.Errorf("guilds: loot vote: %w", err)
		}
		return nil
	}
	_, err := s.Pool.Exec(ctx, `
		insert into loot_votes (guild_id, raid_id, report_id, encounter_id, item_id, character_key, voter_id, vote)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		on conflict (guild_id, (coalesce(raid_id, 0)), (coalesce(report_id, '')), encounter_id, item_id, character_key, voter_id)
		do update set vote = excluded.vote, voted_at = now()`,
		guildID, night.raidID, night.reportID, onyxiaEncounterID, itemID, characterKey, voter, vote)
	if err != nil {
		return fmt.Errorf("guilds: loot vote: %w", err)
	}
	return nil
}

// candidateVotes is one candidate's tally with the officer count filled in.
func (s *Store) candidateVotes(ctx context.Context, guildID int64, night awardNight, itemID int, characterKey string, viewer int64, now time.Time) (LootVotes, error) {
	all, err := s.dropVotes(ctx, guildID, night, itemID, viewer, now)
	if err != nil {
		return LootVotes{}, err
	}
	officers, err := s.officerCount(ctx, guildID)
	if err != nil {
		return LootVotes{}, err
	}
	t := all[characterKey]
	t.Voters = officers
	return t, nil
}
