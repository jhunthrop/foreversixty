package seats

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// SeatFactsForUser reads the seat facts /v1/me carries for userID.
func (s *Store) SeatFactsForUser(ctx context.Context, userID int64, now time.Time) (auth.SeatFacts, error) {
	var facts auth.SeatFacts
	rows, err := s.Pool.Query(ctx, `
		select s.guild_id, g.name, s.state, s.given_at, s.offered_at, s.starts_at, s.given_by
		from premium_seats s join guilds g on g.id = s.guild_id
		where s.holder_user_id = $1 and s.state in ('held', 'accepted', 'offered')
		order by s.guild_id limit 4`, userID)
	if err != nil {
		return facts, fmt.Errorf("seats: facts of %d: %w", userID, err)
	}
	type held struct {
		guildID                      int64
		guild                        string
		state                        State
		givenAt, offeredAt, startsAt *time.Time
		givenBy                      *int64
	}
	held0, err := pgx.CollectRows(rows, func(r pgx.CollectableRow) (held, error) {
		var h held
		err := r.Scan(&h.guildID, &h.guild, &h.state, &h.givenAt, &h.offeredAt, &h.startsAt, &h.givenBy)
		return h, err
	})
	if err != nil {
		return facts, fmt.Errorf("seats: facts of %d: %w", userID, err)
	}
	for _, h := range held0 {
		from, err := s.giverName(ctx, h.guildID, h.givenBy)
		if err != nil {
			return facts, err
		}
		switch {
		case h.state == StateOffered && facts.Offer == nil && h.offeredAt != nil:
			facts.Offer = &auth.SeatOffer{GuildID: h.guildID, GuildName: h.guild, FromName: from, OfferedAt: *h.offeredAt}
		case h.state != StateOffered && facts.Seat == nil && h.givenAt != nil:
			facts.Seat = &auth.HeldSeat{GuildName: h.guild, GivenBy: from, Since: *h.givenAt, StartsAt: h.startsAt}
		}
	}
	facts.Ended, err = s.endedSeat(ctx, userID, now)
	return facts, err
}

func (s *Store) giverName(ctx context.Context, guildID int64, giver *int64) (string, error) {
	if giver == nil {
		return "", nil
	}
	ch, err := characterOf(ctx, s.Pool, guildID, *giver)
	return ch.name, err
}

// endedSeat is the newest comeback still valid and unused.
func (s *Store) endedSeat(ctx context.Context, userID int64, now time.Time) (*auth.EndedSeat, error) {
	var e auth.EndedSeat
	err := s.Pool.QueryRow(ctx, `
		select guild_name, ended_at, id, price_lookup_key from seat_comebacks
		where user_id = $1 and used_at is null and expires_at > $2
		order by ended_at desc limit 1`, userID, now).
		Scan(&e.GuildName, &e.EndedAt, &e.ComebackID, &e.PriceLookupKey)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("seats: ended seat of %d: %w", userID, err)
	}
	return &e, nil
}
