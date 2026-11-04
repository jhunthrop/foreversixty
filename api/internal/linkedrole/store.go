package linkedrole

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// Link is an account's role connection: the Discord account and its sealed refresh token.
type Link struct {
	UserID        int64
	DiscordUserID string
	RefreshBox    []byte
}

// Profile is what a push tells Discord about the account.
type Profile struct {
	Premium   bool
	Battletag string
	Anonymize bool
}

// PGStore is discord_role_links on Postgres. Writes that follow a token refresh compare the
// box they read (compare-and-set), so two syncs racing on one account never drop the link.
type PGStore struct{ Pool *pgxpool.Pool }

// Get returns the account's link; ok is false when it has none.
func (s *PGStore) Get(ctx context.Context, userID int64) (Link, bool, error) {
	l := Link{UserID: userID}
	err := s.Pool.QueryRow(ctx, `select discord_user_id, refresh_token_box from discord_role_links
		where user_id = $1`, userID).Scan(&l.DiscordUserID, &l.RefreshBox)
	if errors.Is(err, pgx.ErrNoRows) {
		return Link{}, false, nil
	}
	if err != nil {
		return Link{}, false, fmt.Errorf("linkedrole: get %d: %w", userID, err)
	}
	return l, true, nil
}

// Save links the account to the Discord account. A Discord account linked to another account
// moves to this one (the newer authorization wins).
func (s *PGStore) Save(ctx context.Context, l Link) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("linkedrole: save: %w", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	if _, err := tx.Exec(ctx, `delete from discord_role_links where discord_user_id = $1 and user_id <> $2`,
		l.DiscordUserID, l.UserID); err != nil {
		return fmt.Errorf("linkedrole: save: release: %w", err)
	}
	if _, err := tx.Exec(ctx, `insert into discord_role_links (user_id, discord_user_id, refresh_token_box)
		values ($1, $2, $3)
		on conflict (user_id) do update set discord_user_id = excluded.discord_user_id,
		  refresh_token_box = excluded.refresh_token_box, pushed_premium = null, pushed_at = null,
		  linked_at = now()`, l.UserID, l.DiscordUserID, l.RefreshBox); err != nil {
		return fmt.Errorf("linkedrole: save: write: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("linkedrole: save: commit: %w", err)
	}
	return nil
}

// Rotate replaces the refresh token if the row still holds oldBox; false when another sync
// rotated it first.
func (s *PGStore) Rotate(ctx context.Context, userID int64, oldBox, newBox []byte) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update discord_role_links set refresh_token_box = $3
		where user_id = $1 and refresh_token_box = $2`, userID, oldBox, newBox)
	if err != nil {
		return false, fmt.Errorf("linkedrole: rotate %d: %w", userID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// DeleteIf removes the link if it still holds box: a revoked token, unless another sync
// rotated it meanwhile. It reports whether a row went.
func (s *PGStore) DeleteIf(ctx context.Context, userID int64, box []byte) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `delete from discord_role_links where user_id = $1 and refresh_token_box = $2`, userID, box)
	if err != nil {
		return false, fmt.Errorf("linkedrole: delete %d: %w", userID, err)
	}
	return tag.RowsAffected() == 1, nil
}

// MarkPushed records what Discord was told.
func (s *PGStore) MarkPushed(ctx context.Context, userID int64, premium bool) error {
	if _, err := s.Pool.Exec(ctx, `update discord_role_links set pushed_premium = $2, pushed_at = now()
		where user_id = $1`, userID, premium); err != nil {
		return fmt.Errorf("linkedrole: mark pushed %d: %w", userID, err)
	}
	return nil
}

// Profile reads the account's Premium (the one rule in package entitlements) and battletag.
func (s *PGStore) Profile(ctx context.Context, userID int64) (Profile, error) {
	var p Profile
	var tag *string
	err := s.Pool.QueryRow(ctx, `select `+entitlements.PremiumSQL("u.id")+`, u.battletag, u.anonymize
		from users u where u.id = $1`, userID).Scan(&p.Premium, &tag, &p.Anonymize)
	if err != nil {
		return Profile{}, fmt.Errorf("linkedrole: profile %d: %w", userID, err)
	}
	if tag != nil {
		p.Battletag = *tag
	}
	return p, nil
}

// LinkedUsers lists every linked account, or with guildID > 0, the linked members of that guild
// (whose Premium may come from its seats).
func (s *PGStore) LinkedUsers(ctx context.Context, guildID int64) ([]int64, error) {
	rows, err := s.Pool.Query(ctx, `select l.user_id from discord_role_links l
		where $1 = 0 or exists (select 1 from guild_members m where m.user_id = l.user_id and m.guild_id = $1)
		order by l.user_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("linkedrole: linked users: %w", err)
	}
	ids, err := pgx.CollectRows(rows, pgx.RowTo[int64])
	if err != nil {
		return nil, fmt.Errorf("linkedrole: linked users: %w", err)
	}
	return ids, nil
}
