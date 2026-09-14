package reports

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
	"github.com/jhunthrop/foreversixty/api/internal/zstdx"
	logparquet "github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

const (
	// maxBundleBytes is the ceiling on one fight bundle: a fight's
	// Parquet is two to ten megabytes compressed, so sixty-four is
	// generous and still bounded.
	maxBundleBytes = 64 << 20
	// maxRawChunkBytes is the contract's 8 MiB cap on a raw chunk.
	maxRawChunkBytes = 8 << 20
	// maxLiveBytes is the ceiling on a live snapshot.
	maxLiveBytes = 8 << 20
	// rawSHAHeader carries the hash of a raw chunk's decoded bytes.
	rawSHAHeader = "X-Raw-SHA256"
	// msgFightIndex is what both fight routes say about a bad {n}.
	msgFightIndex = "the fight index must be a number, counted from one"
	// rawCacheControl is the header the engine's publisher gives a raw
	// chunk. Its own constant is unexported, and a chunk is stored here
	// exactly as it arrived rather than through Publisher.WriteRaw,
	// which would decompress and recompress what the companion sent.
	rawCacheControl = "private, max-age=31536000, immutable"
)

// RankedFight is a verified fight handed to the rankings store.
type RankedFight struct {
	ReportID   string
	FightIndex int
	FoughtAt   time.Time
	Region     string
	Ruleset    string
	GuildID    *int64
	Rows       []metrics.Row
	Combatants []summary.CombatantRow
	// Factions is each player GUID's faction, where the parse could
	// read it. Only a parse of the original log text can: the Parquet
	// schema drops the COMBATANT_INFO payload, so live bundles leave
	// this empty.
	Factions map[string]string
}

// Ranker writes and withdraws ranking rows. The rankings store
// satisfies it; the ingest and the report handlers hold it as an
// interface so a test can watch what they send without a rankings
// store behind it.
type Ranker interface {
	WriteFight(ctx context.Context, f RankedFight) error
	RemoveReport(ctx context.Context, reportID, reason string) error
}

// The reasons this package withdraws a report's ranking rows.
const (
	// ReasonUnverified is a fight that had ranked being demoted by a
	// re-sent bundle whose metrics no longer match its events.
	ReasonUnverified = "a re-sent bundle's metrics do not match its events"
	// ReasonNotRankable is an owner moving a report out of the
	// visibilities that may rank.
	ReasonNotRankable = "the report is no longer visible to the rankings"
)

// Sampler schedules the after-the-fact raw-sample verification of a
// report. The job runs out of band, so the companion's complete call
// returns at once.
type Sampler interface {
	Schedule(reportID string)
}

// Ingest serves the companion's routes: one fight at a time, a live
// snapshot while a fight is open, raw chunks in the background, and a
// completion call at the end of the night.
type Ingest struct {
	Store *Store
	Put   store.Putter
	Rank  Ranker
	Samp  Sampler
	Log   *slog.Logger
}

// MountIngest registers the companion's routes.
func MountIngest(mux *http.ServeMux, i *Ingest) {
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}", auth.RequireDevice(i.putFight))
	mux.HandleFunc("PUT /v1/reports/{id}/fights/{n}/live", auth.RequireDevice(i.putLive))
	mux.HandleFunc("PUT /v1/reports/{id}/raw", auth.RequireDevice(i.putRaw))
	mux.HandleFunc("POST /v1/reports/{id}/complete", auth.RequireDevice(i.complete))
}

func (i *Ingest) logger() *slog.Logger {
	if i.Log != nil {
		return i.Log
	}
	return slog.Default()
}

func (i *Ingest) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	i.logger().Error("ingest", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// owned reads the report and checks the caller may write to it. It
// writes the response itself when the answer is no.
func (i *Ingest) owned(w http.ResponseWriter, r *http.Request) (Report, bool) {
	rep, err := i.Store.Get(r.Context(), r.PathValue("id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such report", nil)
		return Report{}, false
	}
	if err != nil {
		i.fail(w, r, "report", err, "could not read that report just now")
		return Report{}, false
	}
	a := auth.ActorFrom(r.Context())
	if rep.OwnerID == nil || *rep.OwnerID != a.UserID {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "that report is not yours to write to", nil)
		return Report{}, false
	}
	return rep, true
}

// fightIndex reads the {n} path value. The engine numbers fights from
// one, and every route and column keyed on a fight index follows it, so
// zero is not a fight.
func fightIndex(r *http.Request) (int, bool) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		return 0, false
	}
	return n, true
}

// RawRange is the raw_range part of a bundle: where in the log the
// fight came from, and the hash of those bytes.
type RawRange struct {
	StartOffset int64  `json:"start_offset"`
	EndOffset   int64  `json:"end_offset"`
	SHA256      string `json:"sha256"`
}

func (i *Ingest) putFight(w http.ResponseWriter, r *http.Request) {
	rep, ok := i.owned(w, r)
	if !ok {
		return
	}
	n, ok := fightIndex(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgFightIndex, nil)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBundleBytes)
	if err := r.ParseMultipartForm(maxBundleBytes); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the bundle must be multipart with summary, events, metrics, and raw_range parts", nil)
		return
	}
	defer func() { _ = r.MultipartForm.RemoveAll() }()

	var (
		posted   summary.Summary
		rows     []metrics.Row
		rawRange RawRange
	)
	if err := partJSON(r, "summary", &posted); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	if err := partJSON(r, "metrics", &rows); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	if err := partJSON(r, "raw_range", &rawRange); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	eventsBytes, err := partBytes(r, "events")
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}

	sha, err := hex.DecodeString(strings.TrimSpace(rawRange.SHA256))
	if err != nil || len(sha) != sha256.Size {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"raw_range.sha256 must be a hex SHA-256", nil)
		return
	}

	// Idempotency: the same fight with the same raw hash is already
	// stored and verified, so at-least-once delivery costs nothing. A
	// stored fight that did *not* verify falls through and is checked
	// again: a re-send must be answered the way the first send was, and
	// a corrected bundle over the same bytes must still be able to pass.
	switch stored, verified, err := i.Store.FightSHA(r.Context(), rep.ID, n); {
	case err == nil && verified && string(stored) == string(sha):
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{"fight_index": n, "verified": true})
		return
	case err != nil && !errors.Is(err, ErrNotFound):
		i.fail(w, r, "fight", err, "could not store that fight just now")
		return
	}

	events, err := logparquet.Unmarshal(eventsBytes)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the events part is not a Parquet file this engine can read", nil)
		return
	}
	if len(rows) == 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the metrics part is empty", nil)
		return
	}
	head := engine.Header{
		EncounterID: rows[0].EncounterID, Difficulty: rows[0].Difficulty,
		Size: rows[0].Size, Kill: rows[0].Kill, Zone: rep.Zone,
	}
	f, rebuilt, err := engine.Rebuild(n, head, events)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"those events do not make a fight: "+err.Error(), nil)
		return
	}
	derived := metrics.Derive(f, rebuilt)

	record := FightRecord{
		ReportID: rep.ID, Index: n, Name: engine.NameFromEvents(head, rebuilt),
		Kill: f.Kill, DurationMS: rebuilt.DurationMS, StartMS: f.Start.UnixMilli(),
		Players: f.Players, Deaths: f.Deaths, NPCKills: f.NPCKills,
		RawStart: &rawRange.StartOffset, RawEnd: &rawRange.EndOffset, RawSHA256: sha,
	}
	if f.EncounterID != 0 {
		record.EncounterID, record.Difficulty, record.Size = &f.EncounterID, &f.Difficulty, &f.Size
	}

	if mismatch := metrics.Compare(rows, derived); mismatch != nil {
		record.Verified = false
		if _, err := i.Store.UpsertFight(r.Context(), record); err != nil {
			i.fail(w, r, "fight", err, "could not store that fight just now")
			return
		}
		if err := i.Store.Flag(r.Context(), rep.ID, "metrics_mismatch"); err != nil {
			i.fail(w, r, "fight", err, "could not store that fight just now")
			return
		}
		// The fight may have verified on an earlier send and ranked
		// then; this send demoted it. Storing the demotion without
		// withdrawing the rows would leave the report's numbers on the
		// leaderboards with nothing behind them.
		if i.Rank != nil {
			if err := i.Rank.RemoveReport(r.Context(), rep.ID, ReasonUnverified); err != nil {
				i.fail(w, r, "fight", err, "could not store that fight just now")
				return
			}
		}
		i.logger().Warn("ingest", "id", httpx.RequestIDFrom(r.Context()), "op", "verify",
			"report", rep.ID, "fight", n, "mismatch", mismatch.Error())
		// The response names the field and nothing more. The mismatch's
		// own message carries the value the events give, and a companion
		// that reads that on rejection has been told what to post next
		// time; it stays in the line logged just above.
		httpx.WriteError(w, r, http.StatusConflict, "unverified",
			"the metrics do not match the events", map[string]string{"metrics": mismatchField(mismatch)})
		return
	}
	record.Verified = true

	pub := store.Publisher{Keys: Keys(rep.ID), Put: i.Put}
	if err := pub.WriteFight(r.Context(), n, posted, events); err != nil {
		i.fail(w, r, "fight", err, "could not store that fight just now")
		return
	}
	if _, err := i.Store.UpsertFight(r.Context(), record); err != nil {
		i.fail(w, r, "fight", err, "could not store that fight just now")
		return
	}
	if err := i.rank(r.Context(), rep, f.EncounterID, n, f.Start, derived, rebuilt.Combatants); err != nil {
		i.fail(w, r, "fight", err, "could not store that fight just now")
		return
	}
	if err := i.WriteReportJSON(r.Context(), rep); err != nil {
		i.fail(w, r, "fight", err, "could not store that fight just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]any{"fight_index": n, "verified": true})
}

// mismatchField is the name of the field a verification rejected, which
// is the only part of a mismatch a rejection may carry back.
func mismatchField(err error) string {
	var m *metrics.Mismatch
	if errors.As(err, &m) {
		return m.Field
	}
	return "rows"
}

// rank hands a verified encounter to the rankings store. Trash is never
// ranked, and neither is a private report.
func (i *Ingest) rank(ctx context.Context, rep Report, encounterID int64, n int, at time.Time,
	rows []metrics.Row, combatants []summary.CombatantRow) error {
	if i.Rank == nil || encounterID == 0 || !Ranked(rep.Visibility) {
		return nil
	}
	region, ruleset := ReportRealm(rep)
	return i.Rank.WriteFight(ctx, RankedFight{
		ReportID: rep.ID, FightIndex: n, FoughtAt: at.UTC(),
		Region: region, Ruleset: ruleset, GuildID: rep.GuildID,
		Rows: rows, Combatants: combatants,
	})
}

// ReportRealm is the region and ruleset a report's players are keyed
// under: the logging character's, and the defaults when there is none.
func ReportRealm(rep Report) (region, ruleset string) {
	region, ruleset = "us", character.RulesetNormal
	if rep.LoggingCharacter == nil {
		return region, ruleset
	}
	parts := strings.Split(*rep.LoggingCharacter, "/")
	if len(parts) == 3 && character.ValidRegion(parts[0]) && character.ValidRuleset(parts[1]) {
		return parts[0], parts[1]
	}
	return region, ruleset
}

// LiveInput is the body of the live snapshot route.
type LiveInput struct {
	Summary   summary.Summary `json:"summary"`
	ElapsedMS int64           `json:"elapsed_ms"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (i *Ingest) putLive(w http.ResponseWriter, r *http.Request) {
	rep, ok := i.owned(w, r)
	if !ok {
		return
	}
	n, ok := fightIndex(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", msgFightIndex, nil)
		return
	}
	var in LiveInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxLiveBytes)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with summary, elapsed_ms, and updated_at", nil)
		return
	}
	pub := store.Publisher{Keys: Keys(rep.ID), Put: i.Put}
	if err := pub.WriteLive(r.Context(), n, in.Summary); err != nil {
		i.fail(w, r, "live", err, "could not store that snapshot just now")
		return
	}
	if err := i.WriteReportJSON(r.Context(), rep); err != nil {
		i.fail(w, r, "live", err, "could not store that snapshot just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (i *Ingest) putRaw(w http.ResponseWriter, r *http.Request) {
	rep, ok := i.owned(w, r)
	if !ok {
		return
	}
	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	if err != nil || offset < 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "offset must be a byte offset", nil)
		return
	}
	packed, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxRawChunkBytes))
	if err != nil {
		httpx.WriteError(w, r, http.StatusRequestEntityTooLarge, "too_large",
			"a raw chunk may be at most 8 MiB compressed", nil)
		return
	}
	claimed, err := hex.DecodeString(strings.TrimSpace(r.Header.Get(rawSHAHeader)))
	if err != nil || len(claimed) != sha256.Size {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			rawSHAHeader+" must be the hex SHA-256 of the chunk's decoded bytes", nil)
		return
	}
	decoded, err := zstdx.DecodeAll(packed, zstdx.MaxChunk)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "the body must be zstd", nil)
		return
	}
	sum := sha256.Sum256(decoded)
	if string(sum[:]) != string(claimed) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"the chunk does not hash to "+rawSHAHeader, nil)
		return
	}

	stored, err := i.Store.PutRawChunk(r.Context(), rep.ID, RawChunk{
		Start: offset, End: offset + int64(len(decoded)), SHA256: sum[:],
	})
	switch {
	case errors.Is(err, ErrOverlap):
		httpx.WriteError(w, r, http.StatusConflict, "overlap",
			"that offset overlaps a different stored range", nil)
		return
	case err != nil:
		i.fail(w, r, "raw", err, "could not store that chunk just now")
		return
	case !stored:
		// The same bytes at the same offset: already have them.
		w.WriteHeader(http.StatusOK)
		return
	}
	// The body is already zstd, so it is stored as it arrived rather
	// than decoded and recompressed.
	if err := i.Put.Put(r.Context(), Keys(rep.ID).Raw(offset), packed, store.PutOptions{
		ContentType: "application/zstd", CacheControl: rawCacheControl,
		ContentEncoding: "zstd",
	}); err != nil {
		// The row is recorded but the object is not. Left alone it would
		// make the companion's retry look like a chunk already held, and
		// the hole would survive until the raw-sample job re-read the
		// range. Undo the row so the retry re-records and re-uploads.
		// The row is recorded first, and not last, so that a chunk that
		// conflicts with a stored range is refused before it can
		// overwrite that range's object.
		if undo := i.Store.DeleteRawChunk(r.Context(), rep.ID, offset); undo != nil {
			i.logger().Error("ingest", "id", httpx.RequestIDFrom(r.Context()), "op", "raw",
				"report", rep.ID, "offset", offset, "err", undo,
				"msg", "the chunk was recorded but not stored and the record could not be undone")
		}
		i.fail(w, r, "raw", err, "could not store that chunk just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// CompleteInput is the body of the completion call.
type CompleteInput struct {
	FinalOffset   int64          `json:"final_offset"`
	EngineVersion string         `json:"engine_version"`
	Health        session.Health `json:"health"`
}

func (i *Ingest) complete(w http.ResponseWriter, r *http.Request) {
	rep, ok := i.owned(w, r)
	if !ok {
		return
	}
	var in CompleteInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxLiveBytes)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with final_offset, engine_version, and health", nil)
		return
	}
	health, err := json.Marshal(in.Health)
	if err != nil {
		i.fail(w, r, "complete", err, "could not finish that report just now")
		return
	}
	if err := i.Store.SetStatus(r.Context(), rep.ID, StatusComplete, in.EngineVersion, health); err != nil {
		i.fail(w, r, "complete", err, "could not finish that report just now")
		return
	}
	rep.Status, rep.EngineVersion, rep.Health = StatusComplete, in.EngineVersion, health
	if err := i.WriteReportJSON(r.Context(), rep); err != nil {
		i.fail(w, r, "complete", err, "could not finish that report just now")
		return
	}
	if i.Samp != nil {
		i.Samp.Schedule(rep.ID)
	}
	w.WriteHeader(http.StatusNoContent)
}

// WriteReportJSON rewrites reports/<id>/report.json from the database,
// which is the index the report page polls while a raid is live.
func (i *Ingest) WriteReportJSON(ctx context.Context, rep Report) error {
	fights, err := i.Store.Fights(ctx, rep.ID)
	if err != nil {
		return err
	}
	entries := make([]store.FightEntry, 0, len(fights))
	for _, f := range fights {
		entries = append(entries, f.FightEntry)
	}
	out := store.Report{
		ReportID: rep.ID, EngineVersion: rep.EngineVersion, Fights: entries,
	}
	if len(rep.Health) > 0 {
		if err := json.Unmarshal(rep.Health, &out.Health); err != nil {
			return fmt.Errorf("reports: read health of %s: %w", rep.ID, err)
		}
	}
	pub := store.Publisher{Keys: Keys(rep.ID), Put: i.Put}
	return pub.WriteReport(ctx, out)
}

// partJSON decodes one JSON part of a bundle.
func partJSON(r *http.Request, name string, into any) error {
	b, err := partBytes(r, name)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, into); err != nil {
		return fmt.Errorf("the %s part is not the JSON this route expects", name)
	}
	return nil
}

// partBytes reads one part of a bundle whole.
func partBytes(r *http.Request, name string) ([]byte, error) {
	f, _, err := r.FormFile(name)
	if err != nil {
		return nil, fmt.Errorf("the bundle has no %s part", name)
	}
	defer f.Close()
	b, err := io.ReadAll(io.LimitReader(f, maxBundleBytes))
	if err != nil {
		return nil, fmt.Errorf("the %s part could not be read", name)
	}
	return b, nil
}
