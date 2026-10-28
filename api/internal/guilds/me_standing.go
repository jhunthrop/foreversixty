// api/internal/guilds/me_standing.go
//
// GuildStanding answers the per-guild facts GET /v1/me adds to each guilds[] entry (spec
// design/specs/2026-10-27-subscribe-and-join.md §8), for every guild of one account in one
// read per table, never one per guild.
package guilds

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// rulesetLabel is a ruleset as the site writes it (the web's rulesetLabel; the character
// package carries only the ids): on Forever a realm is its ruleset. A ruleset the site has
// not heard of is echoed.
func rulesetLabel(ruleset string) string {
	switch ruleset {
	case character.RulesetNormal:
		return "Normal"
	case character.RulesetPvP:
		return "PvP"
	case character.RulesetRP:
		return "Roleplay"
	case character.RulesetHardcore:
		return "Hardcore"
	}
	return ruleset
}

// identityRow is one guild's identity columns and its newest roster snapshot's member lines
// (nil members: the guild has no snapshot).
type identityRow struct {
	id       int64
	name     string
	ruleset  string
	faction  *string
	crestKey *string
	members  []snapshotMember
	snapshot bool
	// claim carries only the columns claimState and claimHolderStateOf read.
	claim Guild
}

// GuildStanding returns each of guildIDs' identity facts: ruleset label, slug, faction, crest
// url, the roster count, claim and verify. A guild id that does not exist is absent from the
// map. These cost four fixed queries per call whatever the guild count: identity (with the
// newest snapshot and the claim columns), roster rows, invited accounts, and who approves; a
// fifth, the claim holders' names, runs only when a verified officer may see one. Cache the
// roster count by snapshot received_at if /v1/me latency ever shows it.
func (s *Store) GuildStanding(ctx context.Context, viewerID int64, guildIDs []int64) (map[int64]auth.GuildStanding, error) {
	out := make(map[int64]auth.GuildStanding, len(guildIDs))
	if len(guildIDs) == 0 {
		return out, nil
	}
	identities, err := s.identityRows(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	rows, err := s.rosterRowsByGuild(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	invited, err := s.invitedByGuild(ctx, guildIDs)
	if err != nil {
		return nil, err
	}
	claims, verify, err := s.ClaimAndVerifyViews(ctx, viewerID, identities)
	if err != nil {
		return nil, err
	}
	subscribe, err := s.subscribeViews(ctx, viewerID, guildIDs, claims)
	if err != nil {
		return nil, err
	}
	covered, err := s.coveredViews(ctx, viewerID, guildIDs)
	if err != nil {
		return nil, err
	}
	for _, id := range identities {
		claim, verified := claims[id.id], verify[id.id]
		st := auth.GuildStanding{
			Claim:        &claim,
			Verify:       &verified,
			RulesetLabel: rulesetLabel(id.ruleset),
			Slug:         character.Slug(id.name),
			CrestURL:     CrestURL(s.APIBaseURL, id.id, id.crestKey),
			Faction:      id.faction,
			Subscribe:    subscribe[id.id],
			Covered:      covered[id.id],
		}
		if id.snapshot {
			_, counts := mergeRoster(rows[id.id], invited[id.id], id.members)
			n := counts.OnRoster + counts.InGameNoAccount
			st.RosterCount = &n
		}
		out[id.id] = st
	}
	return out, nil
}

// identityRows reads every guild's identity and newest snapshot in one query.
func (s *Store) identityRows(ctx context.Context, guildIDs []int64) ([]identityRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select g.id, g.name, g.ruleset, g.faction, g.crest_key, snap.members,
		       g.claimed_by, g.claimed_at, g.claim_pending_by, g.claim_requested_at, g.claim_contested_at
		from guilds g
		left join guild_roster_snapshots snap on snap.guild_id = g.id
		where g.id = any($1)`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: standing identity: %w", err)
	}
	defer rows.Close()
	var out []identityRow
	for rows.Next() {
		var r identityRow
		var raw []byte
		g := &r.claim
		if err := rows.Scan(&r.id, &r.name, &r.ruleset, &r.faction, &r.crestKey, &raw,
			&g.ClaimedBy, &g.ClaimedAt, &g.ClaimPendingBy, &g.ClaimRequestedAt, &g.ClaimContestedAt); err != nil {
			return nil, fmt.Errorf("guilds: standing identity: %w", err)
		}
		g.ID = r.id
		if raw != nil {
			r.snapshot = true
			if err := json.Unmarshal(raw, &r.members); err != nil {
				return nil, fmt.Errorf("guilds: standing snapshot members of %d: %w", r.id, err)
			}
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

// rosterRowsByGuild reads the characters that make up each guild's roster (the rows homeRoster
// lists), reduced to the name and state mergeRoster counts by.
func (s *Store) rosterRowsByGuild(ctx context.Context, guildIDs []int64) (map[int64][]RosterRow, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.guild_id, gc.character_key, ae.name, gc.rank, gc.verified_at is not null
		from guild_characters gc
		join addon_exports ae on ae.character_key = gc.character_key
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		where gc.guild_id = any($1) and gc.character_key not like 'account:%'`, guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: standing roster: %w", err)
	}
	defer rows.Close()
	out := map[int64][]RosterRow{}
	for rows.Next() {
		var guildID int64
		var r RosterRow
		if err := rows.Scan(&guildID, &r.CharacterKey, &r.Name, &r.Rank, &r.Verified); err != nil {
			return nil, fmt.Errorf("guilds: standing roster: %w", err)
		}
		out[guildID] = append(out[guildID], r)
	}
	return out, rows.Err()
}

// invitedByGuild is invitedMembers for many guilds at once, naming each account by the same
// rule (inviteeName), so the roster count here matches the home's.
func (s *Store) invitedByGuild(ctx context.Context, guildIDs []int64) (map[int64][]invitedMember, error) {
	rows, err := s.Pool.Query(ctx, `
		select gc.guild_id, gc.user_id, c.name
		from guild_characters gc
		join guild_members gm on gm.guild_id = gc.guild_id and gm.user_id = gc.user_id
		join users u on u.id = gc.user_id
		left join lateral (
		  select c.key, c.name from characters c where c.user_id = gc.user_id
		  order by (c.key = u.main_character_key) desc, c.name limit 1
		) c on true
		where gc.guild_id = any($1) and gc.character_key like 'account:%' and gc.verified_at is not null
		  and not exists (select 1 from guild_characters x where x.guild_id = gc.guild_id and x.character_key = c.key)`,
		guildIDs)
	if err != nil {
		return nil, fmt.Errorf("guilds: standing invited: %w", err)
	}
	return s.collectInvited(ctx, rows)
}

// collectInvited reads the rows invitedByGuild selects.
func (s *Store) collectInvited(ctx context.Context, rows pgx.Rows) (map[int64][]invitedMember, error) {
	defer rows.Close()
	out := map[int64][]invitedMember{}
	for rows.Next() {
		var guildID int64
		var m invitedMember
		var name *string
		if err := rows.Scan(&guildID, &m.UserID, &name); err != nil {
			return nil, fmt.Errorf("guilds: standing invited: %w", err)
		}
		m.Name = s.inviteeName(ctx, m.UserID, name)
		out[guildID] = append(out[guildID], m)
	}
	return out, rows.Err()
}
