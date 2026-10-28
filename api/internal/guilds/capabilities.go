// api/internal/guilds/capabilities.go
package guilds

import (
	"context"
	"fmt"
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
	lists, err := s.CapabilitiesFor(ctx, []int64{guildID}, userID, moderator)
	if err != nil {
		return nil, err
	}
	can, ok := lists[guildID]
	if !ok {
		return nil, ErrNotFound
	}
	return can, nil
}

// CapabilitiesFor is Capabilities for many guilds in two queries (the claims, then the
// viewer's verified ranks); a guild that does not exist is absent from the map.
func (s *Store) CapabilitiesFor(ctx context.Context, guildIDs []int64, userID int64, moderator bool) (map[int64][]string, error) {
	claims, err := s.claimHolders(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	officers, err := s.verifiedOfficerGuilds(ctx, guildIDs, userID)
	if err != nil {
		return nil, err
	}
	out := make(map[int64][]string, len(claims))
	for id, claimedBy := range claims {
		out[id] = Capabilities(Guild{ID: id, ClaimedBy: claimedBy}, userID, moderator, officers[id])
	}
	return out, nil
}

// claimHolders maps each existing guild among ids to the account holding its claim.
func (s *Store) claimHolders(ctx context.Context, ids []int64) (map[int64]*int64, error) {
	rows, err := s.Pool.Query(ctx, `select id, claimed_by from guilds where id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("guilds: capabilities: claims: %w", err)
	}
	defer rows.Close()
	out := make(map[int64]*int64, len(ids))
	for rows.Next() {
		var id int64
		var claimedBy *int64
		if err := rows.Scan(&id, &claimedBy); err != nil {
			return nil, fmt.Errorf("guilds: capabilities: claims: %w", err)
		}
		out[id] = claimedBy
	}
	return out, rows.Err()
}

// verifiedOfficerGuilds is the set of guilds among ids where userID is a verified officer or leader.
func (s *Store) verifiedOfficerGuilds(ctx context.Context, ids []int64, userID int64) (map[int64]bool, error) {
	rows, err := s.Pool.Query(ctx,
		`select guild_id, rank from guild_members
		 where guild_id = any($1) and user_id = $2 and verified_at is not null`, ids, userID)
	if err != nil {
		return nil, fmt.Errorf("guilds: capabilities: ranks: %w", err)
	}
	defer rows.Close()
	out := map[int64]bool{}
	for rows.Next() {
		var id int64
		var rank string
		if err := rows.Scan(&id, &rank); err != nil {
			return nil, fmt.Errorf("guilds: capabilities: ranks: %w", err)
		}
		out[id] = rank == "officer" || rank == "leader"
	}
	return out, rows.Err()
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
