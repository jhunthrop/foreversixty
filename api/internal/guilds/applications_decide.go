// api/internal/guilds/applications_decide.go
//
// Accepting or declining an application (owner decision D16). Accepting joins the applicant's
// account the way a redeemed invite does (joinAsInvited) in the same transaction that marks
// the application. After the commit the bot grants the guild's recruit role to the applicant's
// linked Discord user, and DMs the decision only when the applicant opted in. Neither bot step
// can undo the decision: each one's outcome is recorded and answered instead.
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// The DiscordOutcome words: what the bot did, or why it did nothing.
const (
	OutcomeGranted       = "granted"
	OutcomeSent          = "sent"
	OutcomeNoRoleSet     = "no_role_set"
	OutcomeNoBot         = "no_bot"
	OutcomeNotLinked     = "not_linked"
	OutcomeNotOptedIn    = "not_opted_in"
	OutcomeNotAccepted   = "not_accepted"
	OutcomeDiscordRefuse = "discord_refused"
	OutcomeFailed        = "failed"
)

// ApplicationBot is the Forever Sixty bot as an application's decision needs it; nil on
// Service means the deployment has no bot.
type ApplicationBot interface {
	// GrantGuildRole gives roleID, in guildID's Discord server, to userID's linked Discord user.
	// The outcome is OutcomeGranted, OutcomeNoBot, OutcomeNotLinked or OutcomeDiscordRefuse; an
	// error is a failure to ask at all.
	GrantGuildRole(ctx context.Context, guildID, userID int64, roleID string) (string, error)
	// DirectMessage sends text to userID's linked Discord user: OutcomeSent, OutcomeNotLinked
	// or OutcomeDiscordRefuse.
	DirectMessage(ctx context.Context, userID int64, text string) (string, error)
}

// DiscordOutcome is one bot step's answer: Done when it happened, Reason why not otherwise.
type DiscordOutcome struct {
	Done   bool    `json:"done"`
	Reason *string `json:"reason"`
}

func outcomeOf(word string) DiscordOutcome {
	if word == OutcomeGranted || word == OutcomeSent {
		return DiscordOutcome{Done: true}
	}
	return DiscordOutcome{Reason: &word}
}

// Decision is the answer to an accept or a decline.
type Decision struct {
	Application OfficerApplication `json:"application"`
	DiscordRole DiscordOutcome     `json:"discord_role"`
	DiscordDM   DiscordOutcome     `json:"discord_dm"`
}

// decideApplication marks a pending application of guildID accepted or declined by officerID,
// joining the applicant on accept, in one transaction. ErrApplicationClosed when it is not
// pending, ErrNotFound when it is not this guild's.
func (s *Store) decideApplication(ctx context.Context, guildID, officerID, id int64, accept bool) (Application, error) {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return Application{}, fmt.Errorf("guilds: decide: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	var userID int64
	var status string
	err = tx.QueryRow(ctx, `select user_id, status from guild_applications where id = $1 and guild_id = $2 for update`,
		id, guildID).Scan(&userID, &status)
	if errors.Is(err, pgx.ErrNoRows) {
		return Application{}, ErrNotFound
	}
	if err != nil {
		return Application{}, fmt.Errorf("guilds: decide: read: %w", err)
	}
	if status != ApplicationPending {
		return Application{}, ErrApplicationClosed
	}
	next := ApplicationDeclined
	if accept {
		next = ApplicationAccepted
	}
	if _, err := tx.Exec(ctx, `update guild_applications set status = $2, decided_by = $3, decided_at = now()
		where id = $1`, id, next, officerID); err != nil {
		return Application{}, fmt.Errorf("guilds: decide: %w", err)
	}
	if accept {
		if err := s.joinAsInvited(ctx, tx, guildID, userID); err != nil {
			return Application{}, err
		}
	}
	if err := tx.Commit(ctx); err != nil {
		return Application{}, fmt.Errorf("guilds: decide: commit: %w", err)
	}
	return s.application(ctx, guildID, id)
}

// recordOutcomes stores what the bot did on the decision.
func (s *Store) recordOutcomes(ctx context.Context, id int64, role, dm string) error {
	if _, err := s.Pool.Exec(ctx, `update guild_applications set role_outcome = $2, dm_outcome = $3 where id = $1`,
		id, role, dm); err != nil {
		return fmt.Errorf("guilds: record outcomes: %w", err)
	}
	return nil
}

// Decide accepts or declines, then runs the bot steps and records them.
func (s *Service) Decide(ctx context.Context, guildID, officerID, id int64, accept bool) (Decision, error) {
	a, err := s.Store.decideApplication(ctx, guildID, officerID, id, accept)
	if err != nil {
		return Decision{}, err
	}
	role := s.grantRecruitRole(ctx, guildID, a, accept)
	dm := s.messageDecision(ctx, guildID, a, accept)
	if err := s.Store.recordOutcomes(ctx, id, role, dm); err != nil {
		return Decision{}, err
	}
	view, err := s.Store.OfficerApplication(ctx, guildID, officerID, id)
	if err != nil {
		return Decision{}, err
	}
	return Decision{Application: view, DiscordRole: outcomeOf(role), DiscordDM: outcomeOf(dm)}, nil
}

// grantRecruitRole asks the bot for the guild's recruit role on accept; the word it answers,
// or why nothing was asked.
func (s *Service) grantRecruitRole(ctx context.Context, guildID int64, a Application, accept bool) string {
	if !accept {
		return OutcomeNotAccepted
	}
	roleID, err := s.Store.RecruitRoleID(ctx, guildID)
	if err != nil {
		s.logger().Warn("guilds", "op", "application_role", "guild_id", guildID, "err", err)
		return OutcomeFailed
	}
	if roleID == nil {
		return OutcomeNoRoleSet
	}
	if s.ApplicationBot == nil {
		return OutcomeNoBot
	}
	word, err := s.ApplicationBot.GrantGuildRole(ctx, guildID, a.userID, *roleID)
	if err != nil {
		s.logger().Warn("guilds", "op", "application_role", "guild_id", guildID, "application_id", a.ID, "err", err)
		return OutcomeFailed
	}
	return word
}

// messageDecision DMs the decision to an applicant who opted in.
func (s *Service) messageDecision(ctx context.Context, guildID int64, a Application, accept bool) string {
	if !a.DiscordDM {
		return OutcomeNotOptedIn
	}
	if s.ApplicationBot == nil {
		return OutcomeNoBot
	}
	name, err := s.GuildName(ctx, guildID)
	if err != nil {
		s.logger().Warn("guilds", "op", "application_dm", "guild_id", guildID, "err", err)
		return OutcomeFailed
	}
	word, err := s.ApplicationBot.DirectMessage(ctx, a.userID, decisionText(name, a.Character.Name, accept))
	if err != nil {
		s.logger().Warn("guilds", "op", "application_dm", "guild_id", guildID, "application_id", a.ID, "err", err)
		return OutcomeFailed
	}
	return word
}

// decisionText is the DM: plain, in the guild's voice.
func decisionText(guild, character string, accept bool) string {
	who := ""
	if character != "" {
		who = " with " + character
	}
	if accept {
		return guild + " accepted your application" + who + ". Welcome to the guild. Ask an officer for a guild invite in game."
	}
	return guild + " passed on your application" + who + " this time. Thanks for applying."
}
