package builds

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"math"
	"slices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound is returned by Get for an id that has never been saved.
var ErrNotFound = errors.New("builds: not found")

// ErrIDCollision is returned by Save when the row already under a build's
// id is a different build. An id is eight base32 characters of a SHA-256,
// so forty bits: at a million stored builds the chance that some pair
// collides is roughly even. Returning the other build would hand the
// planner a link to someone else's work and silently drop their own, so
// Save refuses instead and the handler answers 500.
var ErrIDCollision = errors.New("builds: id collision")

type Store struct {
	Pool *pgxpool.Pool
	Log  *slog.Logger
}

func (s *Store) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}

// Save inserts b and returns the stored record with created true. When the
// id already exists nothing is written and the existing record is returned
// with created false, provided it really is the same build: an id is a
// content hash, so the title that was saved first wins. An existing row
// whose content differs is an id collision and returns ErrIDCollision.
func (s *Store) Save(ctx context.Context, b Build) (Build, bool, error) {
	classID, err := smallint(b.ClassID, "class_id")
	if err != nil {
		return Build{}, false, err
	}
	raceID, err := smallint(b.RaceID, "race_id")
	if err != nil {
		return Build{}, false, err
	}
	order, err := integers(b.PointOrder)
	if err != nil {
		return Build{}, false, err
	}
	gear := b.Gear
	if gear == nil {
		gear = map[string]int{}
	}
	var title *string
	if b.Title != "" {
		t := b.Title
		title = &t
	}

	err = s.Pool.QueryRow(ctx,
		`insert into builds (id, class_id, race_id, tree_version, point_order, gear, title)
		 values ($1, $2, $3, $4, $5, $6, $7)
		 on conflict (id) do nothing
		 returning created_at, views`,
		b.ID, classID, raceID, b.TreeVersion, order, gear, title).
		Scan(&b.CreatedAt, &b.Views)
	if err == nil {
		b.Gear = gear
		return b, true, nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return Build{}, false, fmt.Errorf("builds: insert %s: %w", b.ID, err)
	}
	existing, err := s.Get(ctx, b.ID)
	if err != nil {
		return Build{}, false, err
	}
	b.Gear = gear
	if !sameContent(existing, b) {
		s.logger().Error("builds", "op", "save", "err", "id collision", "id", b.ID,
			"stored", contentOf(existing), "incoming", contentOf(b))
		return Build{}, false, fmt.Errorf("%w on %s", ErrIDCollision, b.ID)
	}
	return existing, false, nil
}

// sameContent reports whether two records hash to the same id for the same
// reason: every field ID covers. The title is deliberately excluded, since
// it is excluded from the hash and the first one saved wins.
func sameContent(a, b Build) bool {
	return a.ClassID == b.ClassID &&
		a.RaceID == b.RaceID &&
		a.TreeVersion == b.TreeVersion &&
		slices.Equal(a.PointOrder, b.PointOrder) &&
		maps.Equal(a.Gear, b.Gear)
}

// contentOf is the collision log's view of a record: the hashed fields, so
// the two builds that landed on one id can be told apart in the logs.
func contentOf(b Build) string {
	return fmt.Sprintf("class_id=%d race_id=%d tree_version=%s point_order=%v gear=%v",
		b.ClassID, b.RaceID, b.TreeVersion, b.PointOrder, b.Gear)
}

func (s *Store) Get(ctx context.Context, id string) (Build, error) {
	b := Build{ID: id}
	var (
		classID, raceID int16
		order           []int32
		title           *string
	)
	err := s.Pool.QueryRow(ctx,
		`select class_id, race_id, tree_version, point_order, gear, title, created_at, views
		 from builds where id = $1`, id).
		Scan(&classID, &raceID, &b.TreeVersion, &order, &b.Gear, &title, &b.CreatedAt, &b.Views)
	if errors.Is(err, pgx.ErrNoRows) {
		return Build{}, ErrNotFound
	}
	if err != nil {
		return Build{}, fmt.Errorf("builds: get %s: %w", id, err)
	}
	b.ClassID = int(classID)
	b.RaceID = int(raceID)
	b.PointOrder = make([]int, len(order))
	for i, v := range order {
		b.PointOrder[i] = int(v)
	}
	if b.Gear == nil {
		b.Gear = map[string]int{}
	}
	if title != nil {
		b.Title = *title
	}
	return b, nil
}

// GetMany reads every row named by ids in one query, keyed by id. It is
// Get's batched counterpart, for a caller that would otherwise fetch one
// row per id in a loop - the addon inbox, rendering several queued builds
// at once. An id with no matching row is simply absent from the result:
// GetMany is not a validator, and the caller decides what a missing build
// means.
func (s *Store) GetMany(ctx context.Context, ids []string) (map[string]Build, error) {
	out := make(map[string]Build, len(ids))
	if len(ids) == 0 {
		return out, nil
	}
	rows, err := s.Pool.Query(ctx,
		`select id, class_id, race_id, tree_version, point_order, gear, title, created_at, views
		 from builds where id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("builds: get many: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			b               Build
			classID, raceID int16
			order           []int32
			title           *string
		)
		if err := rows.Scan(&b.ID, &classID, &raceID, &b.TreeVersion, &order, &b.Gear, &title,
			&b.CreatedAt, &b.Views); err != nil {
			return nil, fmt.Errorf("builds: get many: %w", err)
		}
		b.ClassID = int(classID)
		b.RaceID = int(raceID)
		b.PointOrder = make([]int, len(order))
		for i, v := range order {
			b.PointOrder[i] = int(v)
		}
		if b.Gear == nil {
			b.Gear = map[string]int{}
		}
		if title != nil {
			b.Title = *title
		}
		out[b.ID] = b
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("builds: get many: %w", err)
	}
	return out, nil
}

// AddViews adds each count to the matching row's view counter in one round
// trip. Ids that no longer exist are simply not updated.
func (s *Store) AddViews(ctx context.Context, counts map[string]int64) error {
	if len(counts) == 0 {
		return nil
	}
	batch := &pgx.Batch{}
	for id, n := range counts {
		batch.Queue(`update builds set views = views + $2 where id = $1`, id, n)
	}
	res := s.Pool.SendBatch(ctx, batch)
	for range counts {
		if _, err := res.Exec(); err != nil {
			_ = res.Close()
			return fmt.Errorf("builds: add views: %w", err)
		}
	}
	if err := res.Close(); err != nil {
		return fmt.Errorf("builds: add views: %w", err)
	}
	return nil
}

// smallint converts a validated id to the column's element type. The
// validator rejects unknown classes, races, and talents before a build
// reaches the store, so an out-of-range value here is a programming error
// rather than user input - but it is checked instead of silently truncated.
func smallint(v int, name string) (int16, error) {
	if v < math.MinInt16 || v > math.MaxInt16 {
		return 0, fmt.Errorf("builds: %s %d does not fit a smallint column", name, v)
	}
	return int16(v), nil
}

// integer converts a validated talent id to point_order's element type. The
// column is integer, not smallint, because the beta client's trait node ids
// run into six digits, well past a smallint's range.
func integer(v int, name string) (int32, error) {
	if v < math.MinInt32 || v > math.MaxInt32 {
		return 0, fmt.Errorf("builds: %s %d does not fit an integer column", name, v)
	}
	return int32(v), nil
}

func integers(in []int) ([]int32, error) {
	out := make([]int32, len(in))
	for i, v := range in {
		s, err := integer(v, "talent id")
		if err != nil {
			return nil, err
		}
		out[i] = s
	}
	return out, nil
}
