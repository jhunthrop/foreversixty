package subscribe

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct{ Pool *pgxpool.Pool }

// UpsertResult describes the row Upsert acted on. Existing is false only for
// a brand-new row, in which case Token/Confirmed/Unsubscribed/
// ConfirmationSentAt are the row's zero values (the caller has no use for
// them: it already knows the token it just inserted, and a fresh row is
// never confirmed, unsubscribed, or previously sent-to).
type UpsertResult struct {
	Existing           bool
	Token              string
	Confirmed          bool
	Unsubscribed       bool
	ConfirmationSentAt *time.Time
}

// Upsert inserts a new subscriber row when normalized is not already present.
// When the address already exists, nothing is inserted; the existing row's
// token, confirmation state, and unsubscribed state are returned so the
// caller can decide whether to resend the confirmation (never inventing a
// new token for a row that already has one, and never re-mailing an
// already-confirmed-and-subscribed address). Whether a resend is currently
// allowed by cooldown is decided separately and atomically by
// ClaimConfirmationSend, not from the ConfirmationSentAt snapshot returned
// here.
func (s *Store) Upsert(ctx context.Context, email, normalized, token, unsubscribeToken string) (UpsertResult, error) {
	var returnedToken string
	err := s.Pool.QueryRow(ctx,
		`insert into subscribers (email, email_normalized, token, unsubscribe_token) values ($1, $2, $3, $4)
		 on conflict (email_normalized) do nothing
		 returning token`, email, normalized, token, unsubscribeToken).Scan(&returnedToken)
	if err == nil {
		return UpsertResult{Existing: false}, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return UpsertResult{}, err
	}
	var res UpsertResult
	res.Existing = true
	err = s.Pool.QueryRow(ctx,
		`select token, confirmed_at is not null, unsubscribed_at is not null, confirmation_sent_at
		 from subscribers where email_normalized = $1`, normalized).
		Scan(&res.Token, &res.Confirmed, &res.Unsubscribed, &res.ConfirmationSentAt)
	if err != nil {
		return UpsertResult{}, err
	}
	return res, nil
}

// ClaimConfirmationSend atomically claims the right to send a confirmation
// email to normalized. It succeeds (claimed=true) and stamps
// confirmation_sent_at = now() in the same statement only when no
// confirmation was sent within cooldown; on failure (claimed=false) nothing
// is written. Doing the check and the write in one statement closes a race
// where concurrent callers on the same address would each read a stale
// confirmation_sent_at and each decide to send. The caller must send only
// after a true result, and must not separately record the send afterward -
// the timestamp this call writes is the record.
func (s *Store) ClaimConfirmationSend(ctx context.Context, normalized string, cooldown time.Duration) (bool, error) {
	tag, err := s.Pool.Exec(ctx,
		`update subscribers set confirmation_sent_at = now()
		 where email_normalized = $1
		   and (confirmation_sent_at is null or confirmation_sent_at < now() - ($2 * interval '1 second'))`,
		normalized, cooldown.Seconds())
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}

func (s *Store) Confirm(ctx context.Context, token string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set confirmed_at = coalesce(confirmed_at, now()), unsubscribed_at = null where token = $1`, token)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}

func (s *Store) Unsubscribe(ctx context.Context, unsubscribeToken string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `update subscribers set unsubscribed_at = now() where unsubscribe_token = $1`, unsubscribeToken)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() == 1, nil
}
