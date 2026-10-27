// Package seats owns the Premium seats a guild plan includes: giving,
// moving, offering, accepting and releasing them. See
// docs/superpowers/specs/2026-10-27-pricing-design.md §4.
package seats

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/billing/catalog"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// State is a seat's state, as premium_seats.state holds it.
type State string

// The four seat states.
const (
	StateEmpty    State = "empty"
	StateHeld     State = "held"
	StateOffered  State = "offered"
	StateAccepted State = "accepted"
)

// Holder is the account holding a seat, shown by one of its characters: the
// main if it is in the guild, otherwise the first verified character there.
type Holder struct {
	UserID       int64  `json:"user_id"`
	CharacterKey string `json:"character_key"`
	Name         string `json:"name"`
	Class        string `json:"class"`
}

// Returned names the last holder of a seat that was released recently.
type Returned struct {
	Name string    `json:"name"`
	At   time.Time `json:"at"`
}

// Seat is one Premium seat of a guild plan.
type Seat struct {
	Slot      int        `json:"slot"`
	State     State      `json:"state"`
	Holder    *Holder    `json:"holder"`
	GivenAt   *time.Time `json:"given_at"`
	StartsAt  *time.Time `json:"starts_at"`
	MovableAt *time.Time `json:"movable_at"`
	Returned  *Returned  `json:"returned"`
	// Offeree is the member a move has been offered to; the holder keeps the
	// seat until they accept.
	Offeree *Offeree `json:"offeree"`
}

// Offeree is the member a seat move waits on.
type Offeree struct {
	Name      string     `json:"name"`
	Class     string     `json:"class"`
	OfferedAt *time.Time `json:"offered_at"`
}

// Candidate is a verified member who holds no seat here. PaysPremium marks
// one who would be offered the seat rather than given it.
type Candidate struct {
	UserID      int64  `json:"user_id"`
	Name        string `json:"name"`
	Class       string `json:"class"`
	PaysPremium bool   `json:"pays_premium"`
}

// EntitlementReader is the part of entitlements.Store this package reads.
type EntitlementReader interface {
	GuildBilling(ctx context.Context, guildID int64) (*entitlements.GuildBilling, error)
	PersonalBilling(ctx context.Context, userID int64) (*entitlements.Billing, error)
}

// Store is every seats table read and write.
type Store struct {
	Pool         *pgxpool.Pool
	Entitlements EntitlementReader
}

// querier is what pgxpool.Pool and pgx.Tx share.
type querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

// seatRow is one premium_seats row.
type seatRow struct {
	slot         int
	state        State
	holderID     *int64
	givenBy      *int64
	givenAt      *time.Time
	offeredAt    *time.Time
	startsAt     *time.Time
	movableAt    *time.Time
	gaveUpKey    *string
	returnedName *string
	returnedAt   *time.Time
	pendingID    *int64
	pendingBy    *int64
	pendingAt    *time.Time
}

const seatColumns = `slot, state, holder_user_id, given_by, given_at, offered_at, starts_at, movable_at,
	holder_gave_up_lookup_key, returned_name, returned_at,
	pending_user_id, pending_by, pending_offered_at`

// qualifiedSeatColumns is seatColumns for a query that aliases premium_seats as s.
const qualifiedSeatColumns = `s.slot, s.state, s.holder_user_id, s.given_by, s.given_at, s.offered_at, s.starts_at,
	s.movable_at, s.holder_gave_up_lookup_key, s.returned_name, s.returned_at,
	s.pending_user_id, s.pending_by, s.pending_offered_at`

func scanSeatRow(row pgx.Row) (seatRow, error) {
	var r seatRow
	err := row.Scan(&r.slot, &r.state, &r.holderID, &r.givenBy, &r.givenAt, &r.offeredAt, &r.startsAt,
		&r.movableAt, &r.gaveUpKey, &r.returnedName, &r.returnedAt, &r.pendingID, &r.pendingBy, &r.pendingAt)
	return r, err
}

// bestCharacterLateral picks the character that stands for an account in a
// guild: the main if it is verified there, otherwise the first by name. It
// reads $1 as the guild id and m.user_id as the account.
const bestCharacterLateral = `(
	select gc.character_key as key, coalesce(c.name, ae.name, '') as name, coalesce(c.class, '') as class
	from guild_characters gc
	join users u on u.id = m.user_id
	left join characters c on c.key = gc.character_key
	left join addon_exports ae on ae.character_key = gc.character_key
	where gc.guild_id = $1 and gc.user_id = m.user_id and gc.verified_at is not null
	  and gc.character_key not like 'account:%'
	order by (gc.character_key = u.main_character_key) desc, coalesce(c.name, ae.name, ''), gc.character_key
	limit 1
) ch`

// character is the account's stand-in character, empty when none is known.
type character struct{ key, name, class string }

// characterOf finds the account's character for guildID, falling back to
// the account's main (or first) character anywhere when it has left the
// guild, so a returned seat can still name its last holder.
func characterOf(ctx context.Context, q querier, guildID, userID int64) (character, error) {
	var ch character
	err := q.QueryRow(ctx, `
		select coalesce(ch.key, ''), coalesce(ch.name, ''), coalesce(ch.class, '')
		from (select $2::bigint as user_id) m left join lateral `+bestCharacterLateral+` on true`,
		guildID, userID).Scan(&ch.key, &ch.name, &ch.class)
	if err != nil {
		return character{}, fmt.Errorf("seats: character of %d: %w", userID, err)
	}
	if ch.key != "" {
		return ch, nil
	}
	err = q.QueryRow(ctx, `
		select c.key, c.name, coalesce(c.class, '')
		from characters c join users u on u.id = c.user_id
		where c.user_id = $1
		order by (c.key = u.main_character_key) desc, c.name, c.key limit 1`, userID).
		Scan(&ch.key, &ch.name, &ch.class)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return character{}, fmt.Errorf("seats: fallback character of %d: %w", userID, err)
	}
	return ch, nil
}

// activeTier is the tier of guildID's live whole-plan row, "" when it has
// none (no plan, a lapsed one, or a scoped grant).
func (s *Store) activeTier(ctx context.Context, guildID int64) (string, error) {
	g, err := s.Entitlements.GuildBilling(ctx, guildID)
	if err != nil {
		return "", fmt.Errorf("seats: guild plan %d: %w", guildID, err)
	}
	if g == nil || !g.Active || g.Scoped() {
		return "", nil
	}
	return g.Tier, nil
}

func (s *Store) seatRows(ctx context.Context, q querier, guildID int64) (map[int]seatRow, error) {
	rows, err := q.Query(ctx, `select `+seatColumns+` from premium_seats where guild_id = $1`, guildID)
	if err != nil {
		return nil, fmt.Errorf("seats: read seats of %d: %w", guildID, err)
	}
	defer rows.Close()
	out := map[int]seatRow{}
	for rows.Next() {
		r, err := scanSeatRow(rows)
		if err != nil {
			return nil, fmt.Errorf("seats: scan seat: %w", err)
		}
		out[r.slot] = r
	}
	return out, rows.Err()
}

// view turns a row into the wire seat; a zero row is an empty slot.
func (s *Store) view(ctx context.Context, q querier, guildID int64, slot int, r seatRow, now time.Time) (Seat, error) {
	seat := Seat{Slot: slot, State: StateEmpty}
	if r.state != "" {
		seat.State, seat.GivenAt, seat.StartsAt, seat.MovableAt = r.state, r.givenAt, r.startsAt, r.movableAt
	}
	if r.holderID != nil {
		ch, err := characterOf(ctx, q, guildID, *r.holderID)
		if err != nil {
			return Seat{}, err
		}
		seat.Holder = &Holder{UserID: *r.holderID, CharacterKey: ch.key, Name: ch.name, Class: ch.class}
	}
	if r.pendingID != nil {
		ch, err := characterOf(ctx, q, guildID, *r.pendingID)
		if err != nil {
			return Seat{}, err
		}
		seat.Offeree = &Offeree{Name: ch.name, Class: ch.class, OfferedAt: r.pendingAt}
	}
	if r.returnedName != nil && r.returnedAt != nil && r.returnedAt.After(now.Add(-catalog.SeatReturnedShown)) {
		seat.Returned = &Returned{Name: *r.returnedName, At: *r.returnedAt}
	}
	return seat, nil
}

// List returns slots 1..SeatsForTier of the guild's live whole plan; a slot
// with no row reads as empty. No plan, or a scoped grant, gives no slots.
func (s *Store) List(ctx context.Context, guildID int64, now time.Time) ([]Seat, error) {
	tier, err := s.activeTier(ctx, guildID)
	if err != nil {
		return nil, err
	}
	return s.listForTier(ctx, s.Pool, guildID, tier, now)
}

func (s *Store) listForTier(ctx context.Context, q querier, guildID int64, tier string, now time.Time) ([]Seat, error) {
	count := catalog.SeatsForTier(tier)
	seats := []Seat{}
	if count == 0 {
		return seats, nil
	}
	rows, err := s.seatRows(ctx, q, guildID)
	if err != nil {
		return nil, err
	}
	for slot := 1; slot <= count; slot++ {
		seat, err := s.view(ctx, q, guildID, slot, rows[slot], now)
		if err != nil {
			return nil, err
		}
		seats = append(seats, seat)
	}
	return seats, nil
}

// IsPayer reports whether userID may manage guildID's seats: a verified
// member who is the plan's billing user (the subscriber, of any rank), or a
// verified officer or leader who holds the guild's claim.
func (s *Store) IsPayer(ctx context.Context, guildID, userID int64) (bool, error) {
	var ok bool
	err := s.Pool.QueryRow(ctx, `
		select exists (
		  select 1 from guild_members m join guilds g on g.id = m.guild_id
		  where m.guild_id = $1 and m.user_id = $2 and m.verified_at is not null
		    and (exists (
		          select 1 from entitlements e
		          where e.guild_id = $1 and e.plan = 'guild' and e.billing_user_id = $2)
		      or (g.claimed_by = $2 and m.rank in ('officer', 'leader')))
		)`, guildID, userID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("seats: payer check %d/%d: %w", guildID, userID, err)
	}
	return ok, nil
}

// IsSubscriber reports whether userID is the billing user of guildID's plan.
func (s *Store) IsSubscriber(ctx context.Context, guildID, userID int64) (bool, error) {
	var ok bool
	err := s.Pool.QueryRow(ctx, `
		select exists (select 1 from entitlements e
		  where e.guild_id = $1 and e.plan = 'guild' and e.billing_user_id = $2)`, guildID, userID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("seats: subscriber check %d/%d: %w", guildID, userID, err)
	}
	return ok, nil
}

// Candidates returns the verified members holding no seat here, by name.
func (s *Store) Candidates(ctx context.Context, guildID int64) ([]Candidate, error) {
	rows, err := s.Pool.Query(ctx, `
		select m.user_id, coalesce(ch.name, ''), coalesce(ch.class, ''), pe.status is not null, coalesce(pe.status, '')
		from guild_members m
		left join lateral `+bestCharacterLateral+` on true
		left join entitlements pe on pe.user_id = m.user_id and pe.plan = 'premium' and pe.source = 'stripe'
		where m.guild_id = $1 and m.verified_at is not null
		  and not exists (select 1 from premium_seats s
		    where s.guild_id = m.guild_id and m.user_id in (s.holder_user_id, s.pending_user_id))
		order by lower(coalesce(ch.name, '')), m.user_id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("seats: candidates of %d: %w", guildID, err)
	}
	defer rows.Close()
	out := []Candidate{}
	for rows.Next() {
		var c Candidate
		var hasRow bool
		var status string
		if err := rows.Scan(&c.UserID, &c.Name, &c.Class, &hasRow, &status); err != nil {
			return nil, fmt.Errorf("seats: scan candidate: %w", err)
		}
		c.PaysPremium = hasRow && paysPremium(sourceStripe, status)
		out = append(out, c)
	}
	return out, rows.Err()
}

// sourceStripe is entitlements.source for a paid subscription.
const sourceStripe = "stripe"

// paysPremium is the one rule for "this account pays for its own Premium":
// a Stripe row whose status still grants access. Candidates and Give share it.
func paysPremium(source, status string) bool {
	return source == sourceStripe && entitlements.StatusGrantsAccess(status)
}

// isVerifiedMember reports whether userID is a verified member of guildID.
func isVerifiedMember(ctx context.Context, q querier, guildID, userID int64) (bool, error) {
	var ok bool
	err := q.QueryRow(ctx, `select exists (select 1 from guild_members
		where guild_id = $1 and user_id = $2 and verified_at is not null)`, guildID, userID).Scan(&ok)
	if err != nil {
		return false, fmt.Errorf("seats: membership %d/%d: %w", guildID, userID, err)
	}
	return ok, nil
}

// emailParties is the guild's name and the account's email address.
func emailParties(ctx context.Context, q querier, guildID, userID int64) (guildName, email string, err error) {
	err = q.QueryRow(ctx, `select g.name, u.email from guilds g, users u where g.id = $1 and u.id = $2`,
		guildID, userID).Scan(&guildName, &email)
	if err != nil {
		return "", "", fmt.Errorf("seats: parties %d/%d: %w", guildID, userID, err)
	}
	return guildName, email, nil
}

// externalStep is a call to another system made as the last step before
// commit. undo runs if the commit then fails.
type externalStep struct {
	do   func(ctx context.Context) error
	undo func(ctx context.Context)
}

// withGuildLock runs fn in a transaction holding the guild's seat lock, so
// two actions on one guild's seats never interleave.
func (s *Store) withGuildLock(ctx context.Context, guildID int64, fn func(tx pgx.Tx) error) error {
	return s.withGuildLockStep(ctx, guildID, func(tx pgx.Tx) (*externalStep, error) { return nil, fn(tx) })
}

// withGuildLockStep is withGuildLock whose fn may return one external step.
// The step runs after fn and immediately before commit; if the commit fails
// its undo runs, so the other system is never left ahead of the database.
func (s *Store) withGuildLockStep(ctx context.Context, guildID int64, fn func(tx pgx.Tx) (*externalStep, error)) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("seats: begin: %w", err)
	}
	defer func() { _ = tx.Rollback(context.WithoutCancel(ctx)) }()
	if _, err := tx.Exec(ctx, `select pg_advisory_xact_lock(hashtextextended('premium-seats:' || $1::bigint::text, 0))`,
		guildID); err != nil {
		return fmt.Errorf("seats: lock guild %d: %w", guildID, err)
	}
	step, err := fn(tx)
	if err != nil {
		return err
	}
	if step == nil {
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("seats: commit: %w", err)
		}
		return nil
	}
	return commitAfterStep(ctx, tx, step)
}

// externalStepTimeout bounds the external step, which runs detached from the
// request so a client that goes away cannot cut it off halfway.
const externalStepTimeout = 20 * time.Second

// commitAfterStep runs step then commits, both detached from the request: once
// the other system may have applied the change, the database must follow it or
// the change must be undone. Any failure of either runs undo; a step that fails
// may still have landed on the other side.
func commitAfterStep(ctx context.Context, tx pgx.Tx, step *externalStep) error {
	detached := context.WithoutCancel(ctx)
	stepCtx, cancel := context.WithTimeout(detached, externalStepTimeout)
	defer cancel()
	if err := step.do(stepCtx); err != nil {
		step.undo(detached)
		return err
	}
	if err := tx.Commit(detached); err != nil {
		step.undo(detached)
		return fmt.Errorf("seats: commit: %w", err)
	}
	return nil
}
