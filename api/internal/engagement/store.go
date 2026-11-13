package engagement

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store reads and writes user_active_days.
type Store struct {
	Pool *pgxpool.Pool
}

// MarkActive adds surface to the account's row for day (a UTC date). A row the backfill wrote
// becomes a live one; a surface already recorded writes nothing.
func (s *Store) MarkActive(ctx context.Context, userID int64, day time.Time, surface Surface) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into user_active_days (user_id, day, surfaces) values ($1, $2::date, $3)
		 on conflict (user_id, day) do update
		   set surfaces = user_active_days.surfaces | excluded.surfaces, backfilled = false
		 where user_active_days.surfaces & excluded.surfaces = 0 or user_active_days.backfilled`,
		userID, day.Format(time.DateOnly), int16(surface)); err != nil {
		return fmt.Errorf("engagement: mark %d: %w", userID, err)
	}
	return nil
}

// internalUsers is every account kept out of the figures: the owner (role admin) and every
// account holding a granted Premium (the Caffeine members and other comped accounts).
const internalUsers = `internal_users as (
  select id as user_id from users where role = 'admin'
  union
  select user_id from entitlements where source = 'grant' and plan = 'premium' and user_id is not null
)`

// periodCounts is the aggregate one daily or weekly row carries over p (start, surfaces): every
// account, the ones whose signup period (the %[1]s expression) is this one, and each surface.
const periodCounts = `count(*),
  count(*) filter (where %[1]s = p.start),
  count(*) filter (where p.surfaces & 1 <> 0),
  count(*) filter (where p.surfaces & 2 <> 0),
  count(*) filter (where p.surfaces & 4 <> 0)`

var dailySQL = `with ` + internalUsers + `,
p as (select user_id, day as start, surfaces from user_active_days where day >= $1::date)
select p.start, ` + fmt.Sprintf(periodCounts, `(u.created_at at time zone 'UTC')::date`) + `
from p join users u on u.id = p.user_id
where p.user_id not in (select user_id from internal_users)
group by p.start`

var weeklySQL = `with ` + internalUsers + `,
p as (
  select user_id, date_trunc('week', day::timestamp)::date as start, bit_or(surfaces) as surfaces
  from user_active_days where day >= $1::date group by 1, 2
)
select p.start, ` + fmt.Sprintf(periodCounts, `date_trunc('week', u.created_at at time zone 'UTC')::date`) + `
from p join users u on u.id = p.user_id
where p.user_id not in (select user_id from internal_users)
group by p.start`

// cohortSQL answers (signup week, weeks later, accounts): offset 0 is the cohort's size, offset
// k how many of its accounts were active in week N+k.
const cohortSQL = `with ` + internalUsers + `,
cohort as (
  select id as user_id, date_trunc('week', created_at at time zone 'UTC')::date as week
  from users
  where (created_at at time zone 'UTC') >= $1::date and id not in (select user_id from internal_users)
),
weekly as (
  select distinct user_id, date_trunc('week', day::timestamp)::date as week
  from user_active_days where day >= $1::date
)
select week, 0, count(*) from cohort group by week
union all
select c.week, (w.week - c.week) / 7, count(*)
from cohort c join weekly w on w.user_id = c.user_id and w.week > c.week
group by c.week, (w.week - c.week) / 7`

// Report reads the owner's figures as of now.
func (s *Store) Report(ctx context.Context, now time.Time) (Report, error) {
	today, thisWeek := dayOf(now), weekOf(now)
	liveSince, err := s.liveSince(ctx)
	if err != nil {
		return Report{}, err
	}
	internal, err := s.internalCount(ctx)
	if err != nil {
		return Report{}, err
	}
	daily, err := s.periodRows(ctx, dailySQL, stepDays(today, 1-DaysShown))
	if err != nil {
		return Report{}, err
	}
	weekly, err := s.periodRows(ctx, weeklySQL, stepWeeks(thisWeek, 1-WeeksShown))
	if err != nil {
		return Report{}, err
	}
	firstCohort := stepWeeks(thisWeek, 1-CohortsShown)
	cells, err := s.cohortCells(ctx, firstCohort)
	if err != nil {
		return Report{}, err
	}
	r := Report{
		Today: today.Format(time.DateOnly), InternalAccounts: internal,
		Days:    periods(daily, today, DaysShown, stepDays, liveSince),
		Weeks:   periods(weekly, thisWeek, WeeksShown, stepWeeks, liveSince),
		Cohorts: cohorts(cells, firstCohort, thisWeek, liveSince),
	}
	if liveSince != nil {
		stamp := liveSince.Format(time.DateOnly)
		r.LiveSince = &stamp
	}
	return r, nil
}

func (s *Store) liveSince(ctx context.Context) (*time.Time, error) {
	var since *time.Time
	if err := s.Pool.QueryRow(ctx, `select min(day) from user_active_days where not backfilled`).Scan(&since); err != nil {
		return nil, fmt.Errorf("engagement: live since: %w", err)
	}
	if since != nil {
		d := dayOf(*since)
		since = &d
	}
	return since, nil
}

func (s *Store) internalCount(ctx context.Context) (int, error) {
	var n int
	if err := s.Pool.QueryRow(ctx, `with `+internalUsers+` select count(*) from internal_users`).Scan(&n); err != nil {
		return 0, fmt.Errorf("engagement: internal accounts: %w", err)
	}
	return n, nil
}

func (s *Store) periodRows(ctx context.Context, query string, from time.Time) ([]periodRow, error) {
	rows, err := s.Pool.Query(ctx, query, from.Format(time.DateOnly))
	if err != nil {
		return nil, fmt.Errorf("engagement: periods: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (periodRow, error) {
		var p periodRow
		err := row.Scan(&p.start, &p.active, &p.fresh, &p.site, &p.addon, &p.app)
		return p, err
	})
	if err != nil {
		return nil, fmt.Errorf("engagement: periods: %w", err)
	}
	return out, nil
}

func (s *Store) cohortCells(ctx context.Context, from time.Time) ([]cohortCell, error) {
	rows, err := s.Pool.Query(ctx, cohortSQL, from.Format(time.DateOnly))
	if err != nil {
		return nil, fmt.Errorf("engagement: cohorts: %w", err)
	}
	out, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (cohortCell, error) {
		var c cohortCell
		err := row.Scan(&c.week, &c.offset, &c.count)
		return c, err
	})
	if err != nil {
		return nil, fmt.Errorf("engagement: cohorts: %w", err)
	}
	return out, nil
}
