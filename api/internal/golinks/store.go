package golinks

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
)

// ErrNotFound is the answer for a slug no link has.
var ErrNotFound = errors.New("golinks: no such link")

// ErrExists is Create's answer for a slug already taken.
var ErrExists = errors.New("golinks: that slug is taken")

const uniqueViolation = "23505"

// Store is the links, their daily counters and the accounts' first-touch slug.
type Store struct{ Pool *pgxpool.Pool }

// Target is the path a slug sends a visitor to.
func (s *Store) Target(ctx context.Context, slug string) (string, error) {
	var target string
	err := s.Pool.QueryRow(ctx, `select target from go_links where slug = $1`, slug).Scan(&target)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", ErrNotFound
	}
	if err != nil {
		return "", fmt.Errorf("golinks: target %q: %w", slug, err)
	}
	return target, nil
}

// CountClick adds one to today's (UTC) counter for the slug, the bots' column when bot is
// set. It reports false, counting nothing, for a slug no link has.
func (s *Store) CountClick(ctx context.Context, slug string, bot bool) (bool, error) {
	human, robot := 1, 0
	if bot {
		human, robot = 0, 1
	}
	tag, err := s.Pool.Exec(ctx, `
		insert into go_link_clicks (slug, day, clicks, bot_clicks)
		select slug, (now() at time zone 'utc')::date, $2, $3 from go_links where slug = $1
		on conflict (slug, day) do update set
		  clicks = go_link_clicks.clicks + excluded.clicks,
		  bot_clicks = go_link_clicks.bot_clicks + excluded.bot_clicks`, slug, human, robot)
	if err != nil {
		return false, fmt.Errorf("golinks: count %q: %w", slug, err)
	}
	return tag.RowsAffected() > 0, nil
}

// RecordArrival sets the account's first-touch slug when the account is new (inside
// ArrivalWindow), has none yet and the slug is a link. It never overwrites one, and reports
// whether it recorded anything.
func (s *Store) RecordArrival(ctx context.Context, userID int64, slug string) (bool, error) {
	tag, err := s.Pool.Exec(ctx, `
		update users set first_touch_slug = $2
		where id = $1 and first_touch_slug is null
		  and created_at > now() - make_interval(secs => $3)
		  and exists (select 1 from go_links where slug = $2)`, userID, slug, ArrivalWindow.Seconds())
	if err != nil {
		return false, fmt.Errorf("golinks: arrival %d %q: %w", userID, slug, err)
	}
	return tag.RowsAffected() > 0, nil
}

// Create adds a link, ErrExists when the slug is taken. The caller validates it first.
func (s *Store) Create(ctx context.Context, l Link) (Link, error) {
	err := s.Pool.QueryRow(ctx,
		`insert into go_links (slug, target, label) values ($1, $2, $3) returning created_at`,
		l.Slug, l.Target, l.Label).Scan(&l.CreatedAt)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolation {
		return Link{}, ErrExists
	}
	if err != nil {
		return Link{}, fmt.Errorf("golinks: create %q: %w", l.Slug, err)
	}
	return l, nil
}

// funnelSQL is every link's funnel, oldest link first. Premium is the entitlements package's
// own rule, so the table and the site agree on who holds it.
var funnelSQL = `
	select l.slug, l.target, l.label, l.created_at,
	  coalesce(c.clicks, 0), coalesce(c.bot_clicks, 0),
	  coalesce(a.sign_ins, 0), coalesce(a.linked, 0), coalesce(a.premium, 0)
	from go_links l
	left join (
	  select slug, sum(clicks)::bigint as clicks, sum(bot_clicks)::bigint as bot_clicks
	  from go_link_clicks group by slug
	) c on c.slug = l.slug
	left join (
	  select u.first_touch_slug as slug, count(*) as sign_ins,
	    count(*) filter (where exists (select 1 from characters ch where ch.user_id = u.id)) as linked,
	    count(*) filter (where ` + entitlements.PremiumSQL("u.id") + `) as premium
	  from users u where u.first_touch_slug is not null group by u.first_touch_slug
	) a on a.slug = l.slug
	order by l.created_at, l.slug`

// Funnel is every link with its funnel.
func (s *Store) Funnel(ctx context.Context) ([]FunnelRow, error) {
	rows, err := s.Pool.Query(ctx, funnelSQL)
	if err != nil {
		return nil, fmt.Errorf("golinks: funnel: %w", err)
	}
	defer rows.Close()
	out := []FunnelRow{}
	for rows.Next() {
		var r FunnelRow
		if err := rows.Scan(&r.Slug, &r.Target, &r.Label, &r.CreatedAt,
			&r.Clicks, &r.BotClicks, &r.SignIns, &r.Linked, &r.Premium); err != nil {
			return nil, fmt.Errorf("golinks: funnel: %w", err)
		}
		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("golinks: funnel: %w", err)
	}
	return out, nil
}
