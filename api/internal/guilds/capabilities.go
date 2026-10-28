// api/internal/guilds/capabilities.go
package guilds

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// The capability words GET /v1/me lists on a guild's entry (its "can" array). The web draws
// a control exactly when its word is in the list and never infers one from a rank, so a
// change to who may do what is made here and nowhere else.
const (
	CapApprove  = "approve"
	CapRemove   = "remove"
	CapSettings = "settings"
	CapAward    = "award"
	CapBilling  = "billing"
	CapRelease  = "release"
)

// Capabilities is the one statement of what a viewer may do in guild g, in the order the
// words are listed above. Officers (a verified officer or leader rank) approve, remove,
// change settings and award; a moderator changes settings, removes and releases a claim
// from outside the guild. Billing and releasing the claim belong to the account holding
// the claim, the guild master on the site, and to nobody else: an officer approved by the
// addon's rank alone is not a second guild master, and the claim, not the rank, is what
// makes one. The routes that gate those actions read the same facts (verifiedOfficerOrLeader,
// ReleaseClaim, mayRemoveRow); this list is what the page reads.
func Capabilities(g Guild, actorID int64, moderator, verifiedOfficer bool) []string {
	claimant := g.ClaimedBy != nil && *g.ClaimedBy == actorID
	out := []string{}
	if verifiedOfficer {
		out = append(out, CapApprove)
	}
	if verifiedOfficer || moderator {
		out = append(out, CapRemove, CapSettings)
	}
	if verifiedOfficer {
		out = append(out, CapAward)
	}
	if mayBill(claimant, verifiedOfficer) {
		out = append(out, CapBilling)
	}
	if claimant || moderator {
		out = append(out, CapRelease)
	}
	return out
}

// mayBill is who may bill a guild: the account holding its claim, with a verified officer or
// leader rank there. Capabilities' billing word and SubscribeStep's ready step both read it.
func mayBill(claimant, verifiedOfficer bool) bool {
	return claimant && verifiedOfficer
}

// Capabilities reads userID's standing in guildID and answers Capabilities for it. A
// viewer with no verified membership still gets a (possibly empty) list; a guild that does
// not exist is ErrNotFound.
func (s *Store) Capabilities(ctx context.Context, guildID, userID int64, moderator bool) ([]string, error) {
	g, err := s.getGuild(ctx, guildID)
	if err != nil {
		return nil, err
	}
	var rank string
	err = s.Pool.QueryRow(ctx,
		`select rank from guild_members where guild_id = $1 and user_id = $2 and verified_at is not null`,
		guildID, userID).Scan(&rank)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("guilds: capabilities: %w", err)
	}
	verifiedOfficer := err == nil && (rank == "officer" || rank == "leader")
	return Capabilities(g, userID, moderator, verifiedOfficer), nil
}

// mayRemoveOwnRow is the one extra rule a viewer's own roster row carries: the account that
// holds the claim, and anyone whose own row is an officer or the leader, has no Remove on
// themselves. They leave through Leave on the account page, which says what it releases and
// who loses access, so the roster never offers the guild master (or an officer) a way to strip
// their own standing, and with it the guild's officers, by one stray click. Members keep
// removing their own row, and so do moderators.
func mayRemoveOwnRow(g Guild, actorID int64, rank string, moderator bool) bool {
	if moderator {
		return true
	}
	if g.ClaimedBy != nil && *g.ClaimedBy == actorID {
		return false
	}
	return rank != "officer" && rank != "leader"
}
