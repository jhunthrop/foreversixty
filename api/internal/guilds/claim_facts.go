// api/internal/guilds/claim_facts.go
//
// The claim and verify facts GET /v1/me adds to each guilds[] entry (spec
// design/specs/2026-10-27-subscribe-and-join.md §8), read for many guilds at once from the
// same columns Claim, ConfirmClaim and the contest code read.
package guilds

import (
	"context"
	"fmt"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// approverSQL is the predicate for a verified officer or leader on a guild_members row: the one
// Capabilities, verifiedOfficerOrLeader and every fact here read.
const approverSQL = `verified_at is not null and rank in ('officer', 'leader')`

// claimHolderState is a guild's claim view and the account a holder name would come from.
type claimHolderState struct {
	view   auth.GuildClaimView
	holder *int64 // the claim holder, else the active pending claimant
}

// claimAndVerifyViews answers, for each of the guilds identities describe, the claim view the
// viewer sees and whether anyone in the guild can approve a character, from approverFacts'
// answer (which GuildStanding reads once and shares with the subscribe step). It costs at most
// one query whatever the guild count, the holder names, skipped when none is to be shown. The
// claim columns themselves arrive with the identity rows.
func (s *Store) claimAndVerifyViews(ctx context.Context, viewerID int64, identities []identityRow,
	approvers map[int64]approverFact) (map[int64]auth.GuildClaimView, map[int64]auth.GuildVerifyView, error) {
	now := time.Now()
	states := make(map[int64]claimHolderState, len(identities))
	named := map[int64]int64{} // guild -> the account whose character name this viewer may see
	for _, id := range identities {
		st := claimHolderStateOf(id.claim, viewerID, now)
		states[id.id] = st
		if approvers[id.id].viewerApproves && st.holder != nil {
			named[id.id] = *st.holder
		}
	}
	names, err := s.guildCharacterNames(ctx, named)
	if err != nil {
		return nil, nil, err
	}
	claims := make(map[int64]auth.GuildClaimView, len(identities))
	verify := make(map[int64]auth.GuildVerifyView, len(identities))
	for id, st := range states {
		if name, ok := names[id]; ok {
			st.view.HolderName = &name
		}
		claims[id] = st.view
		verify[id] = auth.GuildVerifyView{HasApprover: approvers[id].hasApprover}
	}
	return claims, verify, nil
}

// VerifyViews is the verify view of each of guildIDs, for a caller that needs no claim facts.
func (s *Store) VerifyViews(ctx context.Context, guildIDs []int64) (map[int64]auth.GuildVerifyView, error) {
	approvers, err := s.approverFacts(ctx, 0, guildIDs)
	if err != nil {
		return nil, err
	}
	verify := make(map[int64]auth.GuildVerifyView, len(guildIDs))
	for _, id := range guildIDs {
		verify[id] = auth.GuildVerifyView{HasApprover: approvers[id].hasApprover}
	}
	return verify, nil
}

// claimHolderStateOf is g's claim phase and who holds it, from viewerID's side.
func claimHolderStateOf(g Guild, viewerID int64, now time.Time) claimHolderState {
	view := auth.GuildClaimView{State: claimState(g, now).State}
	var holder *int64
	switch {
	case g.ClaimedBy != nil:
		holder = g.ClaimedBy
		view.Yours = *g.ClaimedBy == viewerID
	case g.pendingActive(now):
		holder = g.ClaimPendingBy
		view.PendingYours = *g.ClaimPendingBy == viewerID
	}
	return claimHolderState{view: view, holder: holder}
}

// approverFact is whether a guild has a verified officer or leader, and whether the viewer is one.
type approverFact struct {
	hasApprover    bool
	viewerApproves bool
}

// approverFacts is approverFact for many guilds in one query; a guild with no approver is absent.
func (s *Store) approverFacts(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]approverFact, error) {
	rows, err := s.Pool.Query(ctx, `
		select guild_id, bool_or(user_id = $1)
		from guild_members
		where guild_id = any($2) and `+approverSQL+`
		group by guild_id`, viewerID, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: approver facts: %w", err)
	}
	defer rows.Close()
	out := map[int64]approverFact{}
	for rows.Next() {
		var id int64
		var viewer bool
		if err := rows.Scan(&id, &viewer); err != nil {
			return nil, fmt.Errorf("guilds: approver facts: %w", err)
		}
		out[id] = approverFact{hasApprover: true, viewerApproves: viewer}
	}
	return out, rows.Err()
}

// viewerApproves is the guilds in approvers in which the viewer is a verified officer or
// leader: Capabilities' verified-rank read.
func viewerApproves(approvers map[int64]approverFact) map[int64]bool {
	out := map[int64]bool{}
	for id, f := range approvers {
		if f.viewerApproves {
			out[id] = true
		}
	}
	return out
}

// ClaimHolderRef names the account holding (or asking for) a guild's claim the way the game
// does: its character on this guild's roster, never an account tag. Class is nil when the site
// has not learned it.
type ClaimHolderRef struct {
	Name         string  `json:"name"`
	Class        *string `json:"class"`
	CharacterKey string  `json:"character_key"`
}

// guildCharacterRefs is, for each guild in holders (guild -> account), that account's character
// on the guild's roster, one query for all: verified before unverified (an unapproved export names
// nobody), then its lowest rank number, ties to its main,
// then by name. A guild whose account has no named character there is absent.
func (s *Store) guildCharacterRefs(ctx context.Context, holders map[int64]int64) (map[int64]ClaimHolderRef, error) {
	out := map[int64]ClaimHolderRef{}
	if len(holders) == 0 {
		return out, nil
	}
	guildIDs, userIDs := make([]int64, 0, len(holders)), make([]int64, 0, len(holders))
	for guildID, userID := range holders {
		guildIDs, userIDs = append(guildIDs, guildID), append(userIDs, userID)
	}
	rows, err := s.Pool.Query(ctx, `
		select distinct on (gc.guild_id) gc.guild_id, coalesce(ae.name, c.name), nullif(c.class, ''), gc.character_key
		from guild_characters gc
		join unnest($1::bigint[], $2::bigint[]) as h(guild_id, user_id)
		  on h.guild_id = gc.guild_id and h.user_id = gc.user_id
		join users u on u.id = gc.user_id
		left join addon_exports ae on ae.character_key = gc.character_key
		left join characters c on c.key = gc.character_key
		where gc.character_key not like 'account:%' and coalesce(ae.name, c.name) is not null
		order by gc.guild_id, gc.verified_at is not null desc, gc.rank_index nulls last, `+rankOrderSQL("gc")+`,
		         (gc.character_key = u.main_character_key) desc nulls last, coalesce(ae.name, c.name)`,
		guildIDs, userIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: claim holder characters: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var guildID int64
		var ref ClaimHolderRef
		if err := rows.Scan(&guildID, &ref.Name, &ref.Class, &ref.CharacterKey); err != nil {
			return nil, fmt.Errorf("guilds: claim holder characters: %w", err)
		}
		out[guildID] = ref
	}
	return out, rows.Err()
}

// guildCharacterNames is guildCharacterRefs' names: the one rule that names a claim holder or a
// payer in a guild.
func (s *Store) guildCharacterNames(ctx context.Context, holders map[int64]int64) (map[int64]string, error) {
	refs, err := s.guildCharacterRefs(ctx, holders)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]string, len(refs))
	for guildID, ref := range refs {
		out[guildID] = ref.Name
	}
	return out, nil
}

// claimHolderRef is guildCharacterRefs for one guild; nil when the account has no named
// character on its roster (the page then says "the guild master").
func (s *Store) claimHolderRef(ctx context.Context, guildID, userID int64) (*ClaimHolderRef, error) {
	refs, err := s.guildCharacterRefs(ctx, map[int64]int64{guildID: userID})
	if err != nil {
		return nil, err
	}
	ref, ok := refs[guildID]
	if !ok {
		return nil, nil
	}
	return &ref, nil
}

// withClaimHolder is a claim view as viewerID reads it: whether the claim is theirs and which
// character holds it. claimState itself stays pure.
func (s *Store) withClaimHolder(ctx context.Context, view ClaimStateView, g Guild, viewerID int64) (ClaimStateView, error) {
	if g.ClaimedBy == nil {
		return view, nil
	}
	ref, err := s.claimHolderRef(ctx, g.ID, *g.ClaimedBy)
	if err != nil {
		return view, err
	}
	view.Yours = *g.ClaimedBy == viewerID
	view.ClaimedBy = ref
	return view, nil
}
