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

// Save inserts b and returns the stored record with created true. When
// the id already exists nothing is written and the existing record is
// returned with created false, provided it really is the same build: an
// id is a content hash, so the title that was saved first wins. An
// existing row whose content differs is an id collision and returns
// ErrIDCollision.
//
// userID may be nil: an anonymous save is saved and shareable, it simply
// has no owner and never appears in anyone's list. A signed-in save of a
// build that already exists claims the row when nobody owns it yet —
// otherwise a player could save a build somebody had already shared and
// never find it in their own list — and leaves an owned row alone, the
// way the first title wins.
func (s *Store) Save(ctx context.Context, b Build, userID *int64) (Build, bool, error) {
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
		`insert into builds (id, class_id, race_id, tree_version, point_order, gear, title, user_id)
		 values ($1, $2, $3, $4, $5, $6, $7, $8)
		 on conflict (id) do nothing
		 returning created_at, views`,
		b.ID, classID, raceID, b.TreeVersion, order, gear, title, userID).
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
	if userID != nil {
		if _, err := s.Pool.Exec(ctx,
			`update builds set user_id = $2 where id = $1 and user_id is null`,
			b.ID, *userID); err != nil {
			return Build{}, false, fmt.Errorf("builds: claim %s: %w", b.ID, err)
		}
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

// buildRow is what both a single-row query and a multi-row one satisfy,
// so one function reads a build in one column order.
type buildRow interface{ Scan(dest ...any) error }

// buildColumns is that column order. Every query below selects exactly
// these, in this order, and scanBuild reads them.
const buildColumns = `id, class_id, race_id, tree_version, point_order, gear, title,
	created_at, views`

// scanBuild reads one row. It exists because Get, GetMany and Mine had
// three copies of the same conversions between them, and a fourth was
// one too many.
func scanBuild(row buildRow) (Build, error) {
	var (
		b               Build
		classID, raceID int16
		order           []int32
		title           *string
	)
	if err := row.Scan(&b.ID, &classID, &raceID, &b.TreeVersion, &order, &b.Gear, &title,
		&b.CreatedAt, &b.Views); err != nil {
		return Build{}, err
	}
	b.ClassID, b.RaceID = int(classID), int(raceID)
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

func (s *Store) Get(ctx context.Context, id string) (Build, error) {
	b, err := scanBuild(s.Pool.QueryRow(ctx, `select `+buildColumns+` from builds where id = $1`, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return Build{}, ErrNotFound
	}
	if err != nil {
		return Build{}, fmt.Errorf("builds: get %s: %w", id, err)
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
		`select `+buildColumns+` from builds where id = any($1)`, ids)
	if err != nil {
		return nil, fmt.Errorf("builds: get many: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		b, err := scanBuild(rows)
		if err != nil {
			return nil, fmt.Errorf("builds: get many: %w", err)
		}
		out[b.ID] = b
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("builds: get many: %w", err)
	}
	return out, nil
}

// PerPage is the page size of a player's own build list, the same
// hundred the sim history and the rankings use.
const PerPage = 100

// Page is one page of a player's own builds.
type Page struct {
	Rows    []Build `json:"rows"`
	Total   int     `json:"total"`
	Page    int     `json:"page"`
	PerPage int     `json:"per_page"`
}

// Mine answers one page of a player's own builds, newest first.
func (s *Store) Mine(ctx context.Context, userID int64, page int) (Page, error) {
	if page < 1 {
		page = 1
	}
	out := Page{Rows: []Build{}, Page: page, PerPage: PerPage}
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from builds where user_id = $1`, userID).Scan(&out.Total); err != nil {
		return Page{}, fmt.Errorf("builds: count: %w", err)
	}
	rows, err := s.Pool.Query(ctx,
		`select `+buildColumns+` from builds where user_id = $1
		 order by created_at desc, id limit $2 offset $3`,
		userID, PerPage, (page-1)*PerPage)
	if err != nil {
		return Page{}, fmt.Errorf("builds: list: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		b, err := scanBuild(rows)
		if err != nil {
			return Page{}, fmt.Errorf("builds: scan: %w", err)
		}
		out.Rows = append(out.Rows, b)
	}
	return out, rows.Err()
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
