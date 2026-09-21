package reports

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// ErrNotFound is returned for a report, fight, or upload that is not there.
var ErrNotFound = errors.New("reports: not found")

// ErrOverlap is returned when a raw chunk's range overlaps a stored range
// that is not the same chunk.
var ErrOverlap = errors.New("reports: the offset overlaps a different stored range")

// Store is every report read and write.
type Store struct{ Pool *pgxpool.Pool }

const reportColumns = `id, owner_id, guild_id, title, visibility, zone, status, engine_version,
	upload_id, logging_character, health, flagged, created_at, completed_at`

func scanReport(row pgx.Row) (Report, error) {
	var r Report
	err := row.Scan(&r.ID, &r.OwnerID, &r.GuildID, &r.Title, &r.Visibility, &r.Zone, &r.Status,
		&r.EngineVersion, &r.UploadID, &r.LoggingCharacter, &r.Health, &r.Flagged, &r.CreatedAt, &r.CompletedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Report{}, ErrNotFound
	}
	if err != nil {
		return Report{}, fmt.Errorf("reports: read: %w", err)
	}
	return r, nil
}

// Create inserts a report.
func (s *Store) Create(ctx context.Context, r Report) (Report, error) {
	err := s.Pool.QueryRow(ctx,
		`insert into reports (id, owner_id, guild_id, title, visibility, zone, status, upload_id, logging_character)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning `+reportColumns,
		r.ID, r.OwnerID, r.GuildID, r.Title, r.Visibility, r.Zone, r.Status, r.UploadID, r.LoggingCharacter).
		Scan(&r.ID, &r.OwnerID, &r.GuildID, &r.Title, &r.Visibility, &r.Zone, &r.Status,
			&r.EngineVersion, &r.UploadID, &r.LoggingCharacter, &r.Health, &r.Flagged, &r.CreatedAt, &r.CompletedAt)
	if err != nil {
		return Report{}, fmt.Errorf("reports: create %s: %w", r.ID, err)
	}
	return r, nil
}

// Get reads one report.
func (s *Store) Get(ctx context.Context, id string) (Report, error) {
	return scanReport(s.Pool.QueryRow(ctx, `select `+reportColumns+` from reports where id = $1`, id))
}

// OwnedBy lists a user's own reports, newest first, with the count of
// fights each one holds.
func (s *Store) OwnedBy(ctx context.Context, userID int64, page, perPage int) ([]Summary, int, error) {
	var total int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from reports where owner_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("reports: count of %d: %w", userID, err)
	}
	rows, err := s.Pool.Query(ctx,
		`select r.id, r.title, r.zone, r.status, r.visibility, r.created_at,
		        (select count(*) from fights f where f.report_id = r.id),
		        (select count(*) from fights f where f.report_id = r.id and f.kill
		           and f.encounter_id is not null)
		 from reports r where r.owner_id = $1
		 order by r.created_at desc limit $2 offset $3`, userID, perPage, (page-1)*perPage)
	if err != nil {
		return nil, 0, fmt.Errorf("reports: list of %d: %w", userID, err)
	}
	defer rows.Close()
	out := []Summary{}
	for rows.Next() {
		var s Summary
		if err := rows.Scan(&s.ID, &s.Title, &s.Zone, &s.Status, &s.Visibility,
			&s.CreatedAt, &s.FightCount, &s.KillCount); err != nil {
			return nil, 0, fmt.Errorf("reports: list of %d: %w", userID, err)
		}
		out = append(out, s)
	}
	return out, total, rows.Err()
}

// Recent lists the newest complete public reports, keyset-paged on
// (created_at, id) so a page inserted between two reads of the feed is
// never duplicated or skipped the way an offset page would be. before
// nil reads the first page.
func (s *Store) Recent(ctx context.Context, before *recentCursor, limit int) ([]RecentSummary, error) {
	var beforeCreated *time.Time
	var beforeID *string
	if before != nil {
		beforeCreated, beforeID = &before.CreatedAt, &before.ID
	}
	rows, err := s.Pool.Query(ctx,
		`select r.id, coalesce(nullif(r.title, ''), r.zone), r.created_at,
		        (select count(*) from fights f where f.report_id = r.id),
		        (select count(*) from fights f where f.report_id = r.id and f.kill
		           and f.encounter_id is not null),
		        coalesce(g.name, '')
		 from reports r
		 left join guilds g on g.id = r.guild_id
		 where r.visibility = $1 and r.status = $2
		   and ($3::timestamptz is null or (r.created_at, r.id) < ($3, $4))
		 order by r.created_at desc, r.id desc
		 limit $5`,
		Public, StatusComplete, beforeCreated, beforeID, limit)
	if err != nil {
		return nil, fmt.Errorf("reports: recent: %w", err)
	}
	defer rows.Close()
	out := []RecentSummary{}
	for rows.Next() {
		var row RecentSummary
		if err := rows.Scan(&row.ID, &row.Title, &row.CreatedAt, &row.FightCount,
			&row.KillCount, &row.GuildName); err != nil {
			return nil, fmt.Errorf("reports: recent: %w", err)
		}
		out = append(out, row)
	}
	return out, rows.Err()
}

// Patch is the set of fields PATCH /v1/reports/{id} may change. A nil
// field is left alone.
type Patch struct {
	Title      *string `json:"title"`
	Visibility *string `json:"visibility"`
	GuildID    *int64  `json:"guild_id"`
}

// Update applies a patch and returns the report as it now stands.
func (s *Store) Update(ctx context.Context, id string, p Patch) (Report, error) {
	return scanReport(s.Pool.QueryRow(ctx,
		`update reports set
		   title = coalesce($2, title),
		   visibility = coalesce($3, visibility),
		   guild_id = coalesce($4, guild_id)
		 where id = $1 returning `+reportColumns, id, p.Title, p.Visibility, p.GuildID))
}

// SetStatus moves a report to a new status, recording the engine version
// and health the parse ended with. completedAt is stamped for the
// terminal statuses.
func (s *Store) SetStatus(ctx context.Context, id, status, engineVersion string, health json.RawMessage) error {
	var completed *time.Time
	if status == StatusComplete || status == StatusFailed {
		now := time.Now().UTC()
		completed = &now
	}
	tag, err := s.Pool.Exec(ctx,
		`update reports set status = $2,
		   engine_version = coalesce(nullif($3, ''), engine_version),
		   health = coalesce($4, health),
		   completed_at = coalesce($5, completed_at)
		 where id = $1`, id, status, engineVersion, health, completed)
	if err != nil {
		return fmt.Errorf("reports: set status %s: %w", id, err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

// Flag marks a report as suspect, or clears the flag with an empty
// reason. A flagged report keeps its files; its rankings rows are
// removed by the caller.
func (s *Store) Flag(ctx context.Context, id, reason string) error {
	if _, err := s.Pool.Exec(ctx,
		`update reports set flagged = nullif($2, '') where id = $1`, id, reason); err != nil {
		return fmt.Errorf("reports: flag %s: %w", id, err)
	}
	return nil
}

// UpsertFight writes one fight. It reports created false when the row
// was already there, which is how a re-sent bundle is recognised.
func (s *Store) UpsertFight(ctx context.Context, f FightRecord) (bool, error) {
	if f.Players == nil {
		// The column is not null: a fight with nobody in it is an empty
		// list, not a missing one.
		f.Players = []string{}
	}
	tag, err := s.Pool.Exec(ctx,
		`insert into fights (report_id, fight_index, encounter_id, name, difficulty, size, kill,
		   duration_ms, start_ms, verified, players, deaths, npc_kills,
		   raw_start_offset, raw_end_offset, raw_sha256, boss_health_pct)
		 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		 on conflict (report_id, fight_index) do update set
		   encounter_id = excluded.encounter_id, name = excluded.name,
		   difficulty = excluded.difficulty, size = excluded.size, kill = excluded.kill,
		   duration_ms = excluded.duration_ms, start_ms = excluded.start_ms,
		   verified = excluded.verified, players = excluded.players,
		   deaths = excluded.deaths, npc_kills = excluded.npc_kills,
		   raw_start_offset = excluded.raw_start_offset, raw_end_offset = excluded.raw_end_offset,
		   raw_sha256 = excluded.raw_sha256, boss_health_pct = excluded.boss_health_pct`,
		f.ReportID, f.Index, f.EncounterID, f.Name, f.Difficulty, f.Size, f.Kill,
		f.DurationMS, f.StartMS, f.Verified, f.Players, f.Deaths, f.NPCKills,
		f.RawStart, f.RawEnd, f.RawSHA256, f.BossHealthPct)
	if err != nil {
		return false, fmt.Errorf("reports: write fight %s/%d: %w", f.ReportID, f.Index, err)
	}
	return tag.RowsAffected() == 1, nil
}

// FightSHA reads the raw hash a fight was stored with, and whether that
// fight verified. A re-sent bundle matching both is answered 200 rather
// than written twice; a re-send matching the hash of a fight that did
// not verify is checked again rather than accepted, or an uploader
// would be told its rejected fight was fine.
func (s *Store) FightSHA(ctx context.Context, reportID string, index int) (sha []byte, verified bool, err error) {
	err = s.Pool.QueryRow(ctx,
		`select raw_sha256, verified from fights where report_id = $1 and fight_index = $2`,
		reportID, index).Scan(&sha, &verified)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, false, ErrNotFound
	}
	if err != nil {
		return nil, false, fmt.Errorf("reports: read fight %s/%d: %w", reportID, index, err)
	}
	return sha, verified, nil
}

// FightRawRange reads the byte range a fight was parsed from, which is
// what the raw-sample check re-parses. Both returns are nil for a fight
// whose bundle carried no range.
func (s *Store) FightRawRange(ctx context.Context, reportID string, index int) (*int64, *int64, error) {
	var start, end *int64
	err := s.Pool.QueryRow(ctx,
		`select raw_start_offset, raw_end_offset from fights where report_id = $1 and fight_index = $2`,
		reportID, index).Scan(&start, &end)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil, ErrNotFound
	}
	if err != nil {
		return nil, nil, fmt.Errorf("reports: read fight range %s/%d: %w", reportID, index, err)
	}
	return start, end, nil
}

// Fights lists a report's fights in order.
func (s *Store) Fights(ctx context.Context, reportID string) ([]FightEntry, error) {
	rows, err := s.Pool.Query(ctx,
		`select fight_index, encounter_id, name, difficulty, size, kill, duration_ms, start_ms,
		        verified, players, deaths, npc_kills, boss_health_pct
		 from fights where report_id = $1 order by fight_index`, reportID)
	if err != nil {
		return nil, fmt.Errorf("reports: list fights %s: %w", reportID, err)
	}
	defer rows.Close()
	out := []FightEntry{}
	for rows.Next() {
		var (
			e                            FightEntry
			encounterID, difficulty, siz *int64
			startMS                      int64
			bossHealth                   *float64
		)
		if err := rows.Scan(&e.Index, &encounterID, &e.Name, &difficulty, &siz, &e.Kill,
			&e.DurationMS, &startMS, &e.Verified, &e.Players, &e.Deaths, &e.NPCKills, &bossHealth); err != nil {
			return nil, fmt.Errorf("reports: list fights %s: %w", reportID, err)
		}
		if bossHealth != nil {
			e.BossHealthPct = *bossHealth
		}
		e.Kind = string(fight.Trash)
		if encounterID != nil {
			e.Kind, e.EncounterID = string(fight.Encounter), *encounterID
		}
		if difficulty != nil {
			e.Difficulty = *difficulty
		}
		if siz != nil {
			e.Size = *siz
		}
		e.Start = time.UnixMilli(startMS).UTC()
		e.End = e.Start.Add(time.Duration(e.DurationMS) * time.Millisecond)
		out = append(out, e)
	}
	return out, rows.Err()
}

// Players is every player seen in a report, sorted, for the page's
// source picker.
func (s *Store) Players(ctx context.Context, reportID string) ([]string, error) {
	rows, err := s.Pool.Query(ctx,
		`select distinct unnest(players) as p from fights where report_id = $1 order by p`, reportID)
	if err != nil {
		return nil, fmt.Errorf("reports: list players %s: %w", reportID, err)
	}
	defer rows.Close()
	out := []string{}
	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			return nil, fmt.Errorf("reports: list players %s: %w", reportID, err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// RawChunk is one stored raw range.
type RawChunk struct {
	Start  int64
	End    int64
	SHA256 []byte
}

// PutRawChunk records a stored raw chunk. It returns stored false when
// the same offset and hash are already recorded, so a re-sent chunk is
// recognised, and ErrOverlap when the range conflicts with a different
// stored one.
func (s *Store) PutRawChunk(ctx context.Context, reportID string, c RawChunk) (bool, error) {
	var (
		existingEnd int64
		existingSHA []byte
	)
	err := s.Pool.QueryRow(ctx,
		`select end_offset, sha256 from raw_chunks where report_id = $1 and start_offset = $2`,
		reportID, c.Start).Scan(&existingEnd, &existingSHA)
	switch {
	case err == nil:
		if existingEnd == c.End && string(existingSHA) == string(c.SHA256) {
			return false, nil
		}
		return false, ErrOverlap
	case !errors.Is(err, pgx.ErrNoRows):
		return false, fmt.Errorf("reports: read raw chunk %s@%d: %w", reportID, c.Start, err)
	}
	var overlaps int
	if err := s.Pool.QueryRow(ctx,
		`select count(*) from raw_chunks
		 where report_id = $1 and start_offset < $3 and end_offset > $2`,
		reportID, c.Start, c.End).Scan(&overlaps); err != nil {
		return false, fmt.Errorf("reports: check raw overlap %s: %w", reportID, err)
	}
	if overlaps > 0 {
		return false, ErrOverlap
	}
	if _, err := s.Pool.Exec(ctx,
		`insert into raw_chunks (report_id, start_offset, end_offset, sha256) values ($1, $2, $3, $4)`,
		reportID, c.Start, c.End, c.SHA256); err != nil {
		return false, fmt.Errorf("reports: write raw chunk %s@%d: %w", reportID, c.Start, err)
	}
	return true, nil
}

// DeleteRawChunk removes the record of one raw chunk. It is how the
// ingest undoes a row whose object never reached the bucket: a chunk
// recorded but not stored would be recognised as already held and the
// retry answered without ever uploading it. Removing a chunk that is
// not there is not an error.
func (s *Store) DeleteRawChunk(ctx context.Context, reportID string, start int64) error {
	if _, err := s.Pool.Exec(ctx,
		`delete from raw_chunks where report_id = $1 and start_offset = $2`,
		reportID, start); err != nil {
		return fmt.Errorf("reports: delete raw chunk %s@%d: %w", reportID, start, err)
	}
	return nil
}

// RawChunks lists a report's stored raw ranges in offset order, for the
// raw-sample verification job.
func (s *Store) RawChunks(ctx context.Context, reportID string) ([]RawChunk, error) {
	rows, err := s.Pool.Query(ctx,
		`select start_offset, end_offset, sha256 from raw_chunks where report_id = $1 order by start_offset`,
		reportID)
	if err != nil {
		return nil, fmt.Errorf("reports: list raw chunks %s: %w", reportID, err)
	}
	defer rows.Close()
	out := []RawChunk{}
	for rows.Next() {
		var c RawChunk
		if err := rows.Scan(&c.Start, &c.End, &c.SHA256); err != nil {
			return nil, fmt.Errorf("reports: list raw chunks %s: %w", reportID, err)
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

// Upload is one whole-file upload in progress.
type Upload struct {
	ID         string
	UserID     *int64
	ObjectKey  string
	R2UploadID string
	SizeBytes  int64
	Filename   string
	ReportID   *string
}

// CreateUpload records a started multipart upload.
func (s *Store) CreateUpload(ctx context.Context, u Upload) error {
	if _, err := s.Pool.Exec(ctx,
		`insert into uploads (id, user_id, object_key, r2_upload_id, size_bytes, filename)
		 values ($1, $2, $3, $4, $5, $6)`,
		u.ID, u.UserID, u.ObjectKey, u.R2UploadID, u.SizeBytes, u.Filename); err != nil {
		return fmt.Errorf("reports: create upload %s: %w", u.ID, err)
	}
	return nil
}

// Upload reads one upload.
func (s *Store) Upload(ctx context.Context, id string) (Upload, error) {
	var u Upload
	err := s.Pool.QueryRow(ctx,
		`select id, user_id, object_key, r2_upload_id, size_bytes, filename, report_id
		 from uploads where id = $1`, id).
		Scan(&u.ID, &u.UserID, &u.ObjectKey, &u.R2UploadID, &u.SizeBytes, &u.Filename, &u.ReportID)
	if errors.Is(err, pgx.ErrNoRows) {
		return Upload{}, ErrNotFound
	}
	if err != nil {
		return Upload{}, fmt.Errorf("reports: read upload %s: %w", id, err)
	}
	return u, nil
}

// FinishUpload ties an upload to the report its parse will fill in.
func (s *Store) FinishUpload(ctx context.Context, id, reportID string) error {
	if _, err := s.Pool.Exec(ctx,
		`update uploads set completed_at = now(), report_id = $2 where id = $1`, id, reportID); err != nil {
		return fmt.Errorf("reports: finish upload %s: %w", id, err)
	}
	return nil
}

// Keys is the object-key builder for a report, so callers do not have to
// import the engine's store package for one type.
func Keys(reportID string) store.Keys { return store.Keys{ReportID: reportID} }
