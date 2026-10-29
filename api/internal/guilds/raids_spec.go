// api/internal/guilds/raids_spec.go
//
// The spec a character brings to this guild's raids (spec 2026-10-28 §P.10): one write path for
// the Discord spec card, its Change spec and the raid page's spec select. The key is re-checked
// as the actor's own verified character and the spec as one of its class's.
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/raidspec"
)

// ErrSpecNotOfClass refuses a spec that is unknown or of another class than the character's.
var ErrSpecNotOfClass = errors.New("guilds: that spec is not one of the character's class")

// RaidSpecCommand saves Spec (a specs.json key) as the character's raid spec. A non-zero RaidID
// also moves the character's In, Late or Bench row for that raid to the spec, never its state.
type RaidSpecCommand struct {
	GuildID, ActorID, RaidID int64
	CharacterKey, Spec       string
}

// RaidSpecResult is the character with its new spec and role, and the state its row for the
// command's raid holds ("" with no row, or no raid).
type RaidSpecResult struct {
	Character SignupCandidate
	State     string
}

// SaveRaidSpec stores the raid spec of the actor's own verified character.
func (s *Service) SaveRaidSpec(ctx context.Context, cmd RaidSpecCommand) (RaidSpecResult, error) {
	char, sp, err := s.raidSpecTarget(ctx, cmd)
	if err != nil {
		return RaidSpecResult{}, err
	}
	if _, err := s.Store.Pool.Exec(ctx, `update guild_characters set raid_spec = $3 where guild_id = $1 and character_key = $2`,
		cmd.GuildID, char.CharacterKey, sp.Key); err != nil {
		return RaidSpecResult{}, fmt.Errorf("guilds: save raid spec: %w", err)
	}
	s.logger().Info("guilds", "op", "raid_spec", "guild_id", cmd.GuildID, "user_id", cmd.ActorID,
		"character_key", char.CharacterKey, "spec", sp.Key)
	char.raidSpec = sp.Key
	char = char.applyRaidSpec()
	res := RaidSpecResult{Character: char.candidate(char.defaultRole(), false)}
	res.Character.Verified = true
	if cmd.RaidID == 0 {
		return res, nil
	}
	if res.State, err = s.Store.respecSignup(ctx, cmd.RaidID, char); err != nil {
		return RaidSpecResult{}, err
	}
	if isRespecState(res.State) {
		s.markRaid(ctx, cmd.RaidID)
	}
	return res, nil
}

// raidSpecTarget is the actor's own verified character and the spec, refused otherwise.
func (s *Service) raidSpecTarget(ctx context.Context, cmd RaidSpecCommand) (raidRosterChar, raidspec.Spec, error) {
	chars, err := s.Store.raidRoster(ctx, cmd.GuildID, cmd.CharacterKey)
	if err != nil {
		return raidRosterChar{}, raidspec.Spec{}, err
	}
	if len(chars) == 0 || cmd.CharacterKey == "" {
		return raidRosterChar{}, raidspec.Spec{}, ErrCharacterNotFound
	}
	char := chars[0]
	switch {
	case char.userID != cmd.ActorID:
		return raidRosterChar{}, raidspec.Spec{}, ErrNotYourCharacter
	case !char.verified:
		return raidRosterChar{}, raidspec.Spec{}, ErrCharacterNotVerified
	}
	sp, ok := raidspec.ByKey(cmd.Spec)
	if !ok || (char.Class != "" && sp.Class != char.Class) {
		return raidRosterChar{}, raidspec.Spec{}, ErrSpecNotOfClass
	}
	return char, sp, nil
}

func isRespecState(state string) bool {
	return state == SignupIn || state == SignupLate || state == SignupBench
}

// respecSignup moves the character's In, Late or Bench row for the raid to its spec and role
// and returns the row's state ("" with no row); an Out row is left as it is.
func (s *Store) respecSignup(ctx context.Context, raidID int64, char raidRosterChar) (string, error) {
	var state string
	err := s.Pool.QueryRow(ctx, `
		update guild_raid_signups set
		  spec = case when state in ('in', 'late', 'bench') then $3 else spec end,
		  role = case when state in ('in', 'late', 'bench') then $4 else role end
		where raid_id = $1 and character_key = $2
		returning state`, raidID, char.CharacterKey, char.Spec, char.defaultRole()).Scan(&state)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("guilds: respec signup: %w", err)
	}
	return state, nil
}
