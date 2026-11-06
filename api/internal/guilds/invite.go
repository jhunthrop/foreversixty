// api/internal/guilds/invite.go
package guilds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// InviteTokenPrefix marks an invite token wherever it appears, the same
// way auth.DeviceTokenPrefix does for a device token.
const InviteTokenPrefix = "fsg_"

// inviteTokenChars is the random half of an invite token.
const inviteTokenChars = 32

// syntheticKey is the character_key an invite-joined account gets before
// it ever syncs a real character — a namespace that never collides with a
// real character.Key output (always region/ruleset/name-slug, never
// prefixed "account:").
func syntheticKey(userID int64) string {
	return fmt.Sprintf("account:%d", userID)
}

// RotateInvite issues a fresh invite token for guildID, made by rotatedBy (nil for a tool). The
// SHA-256 hash is the lookup key; with an InviteSealer the token is also kept sealed so a verified
// officer can show it again (InviteURL), and without one it is returned once and never again. The
// old token, its sealed copy and its maker are all replaced, so the old link stops working at once.
func (s *Store) RotateInvite(ctx context.Context, guildID int64, rotatedBy *int64) (string, time.Time, error) {
	token := InviteTokenPrefix + auth.Base32ID(inviteTokenChars)
	box, err := s.sealInvite(token)
	if err != nil {
		return "", time.Time{}, err
	}
	var rotatedAt time.Time
	err = s.Pool.QueryRow(ctx,
		`update guilds set invite_token_hash = $2, invite_token_enc = $3, invite_rotated_by = $4,
		        invite_token_rotated_at = now()
		 where id = $1 returning invite_token_rotated_at`, guildID, auth.TokenHash(token), box, rotatedBy).Scan(&rotatedAt)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("guilds: rotate invite for %d: %w", guildID, err)
	}
	return token, rotatedAt, nil
}

// InviteExports re-syncs, within tx, the guild membership of each of userID's exported
// characters whose export names guildID, the way an addon sync does (addon.Store). A member an
// officer removed keeps his exports, so rejoining by invite brings his synced characters back
// without waiting for the next sync.
type InviteExports interface {
	ResyncGuildExports(ctx context.Context, tx pgx.Tx, userID, guildID int64) error
}

// InviteAccept is what AcceptInvite answers with.
type InviteAccept struct {
	GuildID               int64
	Region, Ruleset, Name string
	Rank                  string
}

// AcceptInvite redeems token for userID: a verified, source="invite" row
// under the synthetic account: key, since an invite-joined account has no
// real character to anchor a normal row by. An officer sharing the link
// is itself the corroboration, so verified_at is set immediately. If the
// account already holds a synthetic row for a different guild, that row
// is transferred (deleted, then recreated here) rather than left behind —
// the unique index on character_key allows only one guild at a time,
// synthetic rows included. The account's exported characters whose export
// names this guild are then re-attached (Exports, the addon sync's path) and
// take the invite's verification (AdoptInviteVerification: member ranks;
// officer and leader ranks still wait for an officer).
func (s *Store) AcceptInvite(ctx context.Context, token string, userID int64) (InviteAccept, error) {
	hash := auth.TokenHash(token)
	var g Guild
	err := s.Pool.QueryRow(ctx,
		`select id, region, ruleset, name from guilds where invite_token_hash = $1`, hash).
		Scan(&g.ID, &g.Region, &g.Ruleset, &g.Name)
	if err != nil {
		// Never distinguish "no such token" from "rotated away" - both a
		// bad token and a hash that no longer matches any row answer the
		// same ErrNotFound here.
		return InviteAccept{}, ErrNotFound
	}

	key := syntheticKey(userID)
	var prevGuildID *int64
	err = s.Pool.QueryRow(ctx,
		`select guild_id from guild_characters where character_key = $1`, key).Scan(&prevGuildID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: read previous: %w", err)
	}

	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if prevGuildID != nil && *prevGuildID != g.ID {
		if _, err := tx.Exec(ctx,
			`delete from guild_characters where character_key = $1 and guild_id = $2`, key, *prevGuildID); err != nil {
			return InviteAccept{}, fmt.Errorf("guilds: accept invite: clear previous: %w", err)
		}
		if err := RecomputeMembership(ctx, tx, *prevGuildID, &userID); err != nil {
			return InviteAccept{}, err
		}
		if err := RecomputeFaction(ctx, tx, *prevGuildID); err != nil {
			return InviteAccept{}, err
		}
		if err := ReleaseClaimIfLost(ctx, tx, *prevGuildID, userID); err != nil {
			return InviteAccept{}, err
		}
	}
	if _, err := tx.Exec(ctx,
		`insert into guild_characters (guild_id, character_key, user_id, rank, source, verified_at, verified_by, refreshed_at)
		 values ($1, $2, $3, 'member', 'invite', now(), 'invite', now())
		 on conflict (guild_id, character_key) do update set
		   user_id = excluded.user_id, refreshed_at = now(),
		   verified_at = coalesce(guild_characters.verified_at, now()),
		   verified_by = coalesce(guild_characters.verified_by, 'invite')`,
		g.ID, key, userID); err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: %w", err)
	}
	if s.Exports != nil {
		if err := s.Exports.ResyncGuildExports(ctx, tx, userID, g.ID); err != nil {
			return InviteAccept{}, err
		}
	}
	if err := AdoptInviteVerification(ctx, tx, g.ID, userID); err != nil {
		return InviteAccept{}, err
	}
	if err := RecomputeMembership(ctx, tx, g.ID, &userID); err != nil {
		return InviteAccept{}, err
	}
	if err := RecomputeFaction(ctx, tx, g.ID); err != nil {
		return InviteAccept{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return InviteAccept{}, fmt.Errorf("guilds: accept invite: commit: %w", err)
	}
	return InviteAccept{GuildID: g.ID, Region: g.Region, Ruleset: g.Ruleset, Name: g.Name, Rank: "member"}, nil
}

// AdoptInviteVerification moves an invite's verification from userID's placeholder row
// (syntheticKey) in guildID onto the account's own character rows there, then drops the
// placeholder, within tx. It runs whenever either side lands: an invited account whose
// character syncs into the guild, and an account with characters already on the roster that
// joins by invite. The join page promises "No officer has to approve you", so a member-rank
// character of an invite-verified account never reads "Waiting for approval". An officer or
// leader rank is the character's own export's word, and an invite link can be passed on, so such
// a character still waits for an officer (Approve all names it: "including 1 officer"); the
// placeholder then stays, keeping the account verified as a member. With no member-rank
// character in the guild, or no verified placeholder, nothing changes. Callers recompute
// membership after.
func AdoptInviteVerification(ctx context.Context, tx pgx.Tx, guildID, userID int64) error {
	if _, err := tx.Exec(ctx, `
		with placeholder as (
		  delete from guild_characters p
		  where p.guild_id = $1 and p.character_key = $2 and p.verified_at is not null
		    and exists (select 1 from guild_characters c
		                where c.guild_id = $1 and c.user_id = $3 and c.character_key not like 'account:%'
		                  and c.rank not in ('officer', 'leader'))
		  returning p.verified_at
		)
		update guild_characters gc set verified_at = placeholder.verified_at, verified_by = 'invite'
		from placeholder
		where gc.guild_id = $1 and gc.user_id = $3 and gc.character_key not like 'account:%'
		  and gc.verified_at is null and gc.rank not in ('officer', 'leader')`,
		guildID, syntheticKey(userID), userID); err != nil {
		return fmt.Errorf("guilds: adopt invite verification for user %d in guild %d: %w", userID, guildID, err)
	}
	return nil
}

// InvitePeek is what PeekInvite answers with: the guild a token opens, how many verified
// accounts it has, and the officer who made the link by their character on the guild's roster
// (Inviter, nil when the maker is unknown or has no character there), and nothing about who else
// holds the token or who is in the guild.
type InvitePeek struct {
	GuildID               int64
	Region, Ruleset, Name string
	Faction, CrestKey     *string
	Members               int
	Inviter               *ClaimHolderRef
}

// PeekInvite looks a token up without redeeming it, so the invite page can name the guild and
// the inviter before anyone joins. Members is the billing count, verified accounts. Unknown and
// rotated-away tokens both answer ErrNotFound.
func (s *Store) PeekInvite(ctx context.Context, token string) (InvitePeek, error) {
	var p InvitePeek
	var rotatedBy *int64
	err := s.Pool.QueryRow(ctx,
		`select g.id, g.region, g.ruleset, g.name, g.faction, g.crest_key, g.invite_rotated_by,
		        (select count(distinct m.user_id) from guild_members m
		         where m.guild_id = g.id and m.verified_at is not null)
		 from guilds g where g.invite_token_hash = $1`,
		auth.TokenHash(token)).Scan(&p.GuildID, &p.Region, &p.Ruleset, &p.Name, &p.Faction, &p.CrestKey,
		&rotatedBy, &p.Members)
	if errors.Is(err, pgx.ErrNoRows) {
		return InvitePeek{}, ErrNotFound
	}
	if err != nil {
		return InvitePeek{}, fmt.Errorf("guilds: peek invite: %w", err)
	}
	if rotatedBy != nil {
		if p.Inviter, err = s.claimHolderRef(ctx, p.GuildID, *rotatedBy); err != nil {
			return InvitePeek{}, err
		}
	}
	return p, nil
}

// JoinsAs is the character an invited account would join a guild as (spec 2026-11-05 §3.4): the
// name, class and level as the site knows them. OnRoster is false for the stand-in officers see
// until one of the account's characters shows on the guild's roster.
type JoinsAs struct {
	Name     string  `json:"name"`
	Class    *string `json:"class"`
	Race     *string `json:"race"`
	Level    *int    `json:"level"`
	OnRoster bool    `json:"on_roster"`
}

// JoinsAsFor picks userID's character for p's guild: the character on the guild's roster (a
// guild_characters row or its newest addon snapshot names it) at the highest level; else the main
// with an export, else the newest export, else the main, else the first by name. Nil for an
// account with no characters.
func (s *Store) JoinsAsFor(ctx context.Context, p InvitePeek, userID int64) (*JoinsAs, error) {
	rows, err := s.Pool.Query(ctx, `
		with mine as (
		  select c.key, c.name, nullif(c.class, '') as class, nullif(c.race, '') as race, c.level,
		         c.region = $3 and c.ruleset = $4 and (
		           exists (select 1 from guild_characters gc where gc.guild_id = $2 and gc.character_key = c.key)
		           or exists (select 1 from guild_roster_snapshots snap, jsonb_array_elements(snap.members) m
		                      where snap.guild_id = $2
		                        and lower(split_part(m->>'name', '-', 1)) = lower(c.name))
		         ) as on_roster,
		         c.key is not distinct from u.main_character_key as main,
		         ae.updated_at as exported_at
		  from characters c
		  join users u on u.id = c.user_id
		  left join addon_exports ae on ae.character_key = c.key
		  where c.user_id = $1
		)
		select name, class, race, level, on_roster from mine
		order by on_roster desc,
		         case when on_roster then level end desc nulls last,
		         (main and exported_at is not null) desc,
		         exported_at desc nulls last,
		         main desc, name
		limit 1`, userID, p.GuildID, p.Region, p.Ruleset)
	if err != nil {
		return nil, fmt.Errorf("guilds: joins as: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}
	var j JoinsAs
	if err := rows.Scan(&j.Name, &j.Class, &j.Race, &j.Level, &j.OnRoster); err != nil {
		return nil, fmt.Errorf("guilds: joins as: %w", err)
	}
	return &j, rows.Err()
}

func (s *Service) rotateInvite(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to rotate its invite link", nil)
		return
	}
	if !s.officerNotFrozen(w, r, guildID, "rotate_invite", "could not rotate that invite just now") {
		return
	}
	token, rotatedAt, err := s.Store.RotateInvite(r.Context(), guildID, &actor.UserID)
	if err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	}
	invite, err := s.rotatedInviteView(r.Context(), guildID, actor.UserID, token, rotatedAt)
	if err != nil {
		s.fail(w, r, "rotate_invite", err, "could not rotate that invite just now")
		return
	}
	s.logger().Info("guilds", "op", "invite_rotate", "guild_id", guildID, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"token": token, "url": invitePath(token), "rotated_at": rotatedAt, "invite": invite,
	})
}

// peekInvite answers GET /v1/guild-invites/{token}: the guild's name, realm, faction, crest and
// verified count for a valid token, signed in or not, and for a signed-in viewer the character
// they would join as. The token is the secret (a 32-char random value only an officer can mint),
// so holding it is already the right to know which guild it opens; the route is rate-limited like
// the redemption it precedes.
func (s *Service) peekInvite(w http.ResponseWriter, r *http.Request) {
	p, err := s.Store.PeekInvite(r.Context(), r.PathValue("token"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that invite link is not valid", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "peek_invite", err, "could not read that invite just now")
		return
	}
	var joinsAs *JoinsAs
	if actor := auth.ActorFrom(r.Context()); actor.UserID != 0 {
		if joinsAs, err = s.Store.JoinsAsFor(r.Context(), p, actor.UserID); err != nil {
			s.fail(w, r, "peek_invite", err, "could not read that invite just now")
			return
		}
	}
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{
		"guild": map[string]any{
			"id": p.GuildID, "region": p.Region, "ruleset": p.Ruleset, "name": p.Name,
			"faction":   p.Faction,
			"crest_url": CrestURL(s.Store.APIBaseURL, p.GuildID, p.CrestKey),
		},
		"members":  p.Members,
		"joins_as": joinsAs,
		"inviter":  p.Inviter,
	})
}

func (s *Service) acceptInvite(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	actor := auth.ActorFrom(r.Context())
	result, err := s.Store.AcceptInvite(r.Context(), token, actor.UserID)
	switch {
	case errors.Is(err, ErrNotFound):
		// Never distinguish "no such token" from "rotated away" — both
		// fall out of the same zero-row lookup in AcceptInvite.
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "that invite link is not valid", nil)
	case err != nil:
		s.fail(w, r, "accept_invite", err, "could not accept that invite just now")
	default:
		s.logger().Info("guilds", "op", "invite_accept", "guild_id", result.GuildID, "user_id", actor.UserID)
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{
			"guild": map[string]any{
				"id": result.GuildID, "region": result.Region, "ruleset": result.Ruleset, "name": result.Name,
			},
			"rank": result.Rank,
		})
	}
}
