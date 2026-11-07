package activity

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// The feed filters (GET /v1/feed?filter=).
const (
	FilterAll   = "all"
	FilterSpec  = "spec"
	FilterGuild = "guild"
	FilterLogs  = "logs"
)

// Page sizes.
const (
	DefaultLimit = 7
	MaxLimit     = 30
)

// logKinds are the rows the Logs filter keeps: what players did in the game.
var logKinds = []string{KindLogUpload, KindBossKill, KindRankingTop, KindRankingClimb}

// ErrBadCursor is a cursor this API did not write.
var ErrBadCursor = errors.New("activity: bad cursor")

// Query is one feed read.
type Query struct {
	Filter string
	// Spec and Band are the viewer's: they choose the Your spec rows and the tag.
	Spec string
	Band int
	// GuildIDs are the viewer's verified guilds; empty signed out.
	GuildIDs []int64
	Cursor   string
	Limit    int
	Now      time.Time
}

// Row is one feed row as GET /v1/feed answers it.
type Row struct {
	ID       string    `json:"id"`
	Kind     string    `json:"kind"`
	At       time.Time `json:"at"`
	Tag      string    `json:"tag"`
	Icon     *string   `json:"icon"`
	Crest    *string   `json:"crest"`
	Quality  *int      `json:"quality"`
	Href     string    `json:"href"`
	DateOnly bool      `json:"date_only"`
	Segments []Segment `json:"segments"`
}

// Page is the feed answer.
type Page struct {
	Rows       []Row   `json:"rows"`
	NextCursor *string `json:"next_cursor"`
}

// cursor is the last row's sort key: its hour, whether it was the viewer's own, its time
// and id (rows sort newest hour first, the viewer's own rows first within an hour).
type cursor struct {
	hour time.Time
	mine int
	at   time.Time
	id   int64
}

func (c cursor) encode() string {
	raw := fmt.Sprintf("%d|%d|%d|%d", c.hour.UnixMicro(), c.mine, c.at.UnixMicro(), c.id)
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func decodeCursor(s string) (cursor, error) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return cursor{}, ErrBadCursor
	}
	parts := strings.Split(string(raw), "|")
	if len(parts) != 4 {
		return cursor{}, ErrBadCursor
	}
	n := make([]int64, 4)
	for i, p := range parts {
		if n[i], err = strconv.ParseInt(p, 10, 64); err != nil {
			return cursor{}, ErrBadCursor
		}
	}
	return cursor{hour: time.UnixMicro(n[0]).UTC(), mine: int(n[1]), at: time.UnixMicro(n[2]).UTC(), id: n[3]}, nil
}

// mineSQL is 1 for a row of the viewer's own: their guild's, or their spec and band's.
const mineSQL = `(case when e.guild_id is not null or (e.spec = $2 and (e.band = $3 or e.band is null)) then 1 else 0 end)`

func (q Query) where() (string, []any, error) {
	args := []any{q.Now.Add(-FeedWindow), q.Spec, q.Band, q.GuildIDs, q.Now}
	clauses := []string{"e.occurred_at >= $1", "e.occurred_at <= $5", "(e.guild_id is null or e.guild_id = any($4))"}
	switch q.Filter {
	case FilterSpec:
		clauses = append(clauses, "e.spec = $2 and (e.band = $3 or e.band is null)")
	case FilterGuild:
		clauses = append(clauses, "e.guild_id = any($4)")
	case FilterLogs:
		args = append(args, logKinds)
		clauses = append(clauses, fmt.Sprintf("e.kind = any($%d)", len(args)))
	}
	if q.Cursor != "" {
		c, err := decodeCursor(q.Cursor)
		if err != nil {
			return "", nil, err
		}
		args = append(args, c.hour, c.mine, c.at, c.id)
		n := len(args)
		clauses = append(clauses, fmt.Sprintf("(date_trunc('hour', e.occurred_at), %s, e.occurred_at, e.id) < ($%d, $%d, $%d, $%d)",
			mineSQL, n-3, n-2, n-1, n))
	}
	return strings.Join(clauses, " and "), args, nil
}

// Read answers one feed page.
func (s *Store) Read(ctx context.Context, q Query) (Page, error) {
	where, args, err := q.where()
	if err != nil {
		return Page{}, err
	}
	limit := q.Limit
	if limit < 1 || limit > MaxLimit {
		limit = DefaultLimit
	}
	sql := `select e.id, e.kind, e.occurred_at, date_trunc('hour', e.occurred_at), ` + mineSQL + `,
		e.spec, e.band, e.guild_id is not null, e.icon, e.crest, e.quality, e.href, e.date_only, e.segments
		from activity_events e where ` + where + `
		order by date_trunc('hour', e.occurred_at) desc, ` + mineSQL + ` desc, e.occurred_at desc, e.id desc
		limit ` + strconv.Itoa(limit+1)
	rows, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		return Page{}, fmt.Errorf("activity: read feed: %w", err)
	}
	defer rows.Close()
	page := Page{Rows: []Row{}}
	var last cursor
	for rows.Next() {
		if len(page.Rows) == limit {
			next := last.encode()
			page.NextCursor = &next
			break
		}
		row, c, err := scanRow(rows, q)
		if err != nil {
			return Page{}, err
		}
		page.Rows, last = append(page.Rows, row), c
	}
	return page, rows.Err()
}

type scanner interface{ Scan(dest ...any) error }

func scanRow(rows scanner, q Query) (Row, cursor, error) {
	var (
		r        Row
		c        cursor
		spec     *string
		band     *int
		guild    bool
		ic, cr   string
		segments []byte
	)
	if err := rows.Scan(&c.id, &r.Kind, &r.At, &c.hour, &c.mine, &spec, &band, &guild, &ic, &cr,
		&r.Quality, &r.Href, &r.DateOnly, &segments); err != nil {
		return Row{}, cursor{}, fmt.Errorf("activity: scan feed: %w", err)
	}
	r.At = r.At.UTC()
	c.at = r.At
	r.ID = "e_" + strconv.FormatInt(c.id, 10)
	r.Tag = tagFor(KindTag[r.Kind], guild, spec, band, q)
	r.Icon, r.Crest = nonEmpty(ic), nonEmpty(cr)
	if err := json.Unmarshal(segments, &r.Segments); err != nil {
		return Row{}, cursor{}, fmt.Errorf("activity: decode row %d: %w", c.id, err)
	}
	return r, c, nil
}

// tagFor is the row's tag for this viewer: a guild row is always Your guild, a row of the
// viewer's spec and band is Your spec, anything else shows its kind's tag (KindTag).
func tagFor(base string, guild bool, spec *string, band *int, q Query) string {
	switch {
	case guild:
		return TagYourGuild
	case q.Spec != "" && spec != nil && *spec == q.Spec && (band == nil || *band == q.Band):
		return TagYourSpec
	default:
		return base
	}
}

func nonEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
