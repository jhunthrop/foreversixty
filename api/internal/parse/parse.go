// Package parse runs the engine over a whole uploaded log, and re-parses
// samples of a live report's raw chunks to check the companion told the
// truth. Both are the same engine the ingest verifies with.
package parse

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/metrics"
	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/zstdx"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// chunkSize is how much of the upload is fed to the engine at a time.
// Memory stays bounded by the open fight, never the file.
const chunkSize = 1 << 20

// maxDecoded is the ceiling on the log one upload may yield. The upload
// routes cap what a browser can store at reports.MaxUploadBytes, but a
// zstd frame inside that cap can decode to far more, so the decoded
// stream needs a bound of its own. Four times the upload cap is the
// same ratio zstdx.MaxChunk uses: ample headroom for a log that
// compresses well, and still a number.
const maxDecoded int64 = 4 * reports.MaxUploadBytes

// readTimeout bounds the part of a parse that waits on the network. The
// R2 client sets no HTTP timeout, so a body that stalls mid-stream would
// otherwise hold a Cloud Run job open forever; a parse that has not read
// its upload in two hours is stuck rather than slow. The report's own
// status is written outside this deadline, so a parse that runs out of
// time still records that it failed.
const readTimeout = 2 * time.Hour

// Objects is the object store the job reads the upload from and writes
// the report's files to.
type Objects interface {
	store.Putter
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}

// Deps is everything a job needs. The same set serves the whole-file
// parse and the raw-sample check.
type Deps struct {
	Reports *reports.Store
	Objects Objects
	Rank    reports.Ranker
	Log     *slog.Logger
}

func (d Deps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}

// Report parses a whole uploaded log into a report: fights are written
// as they close, so the page fills in progressively, and the status
// ends as complete or failed with the engine's health record.
func Report(ctx context.Context, d Deps, reportID string) error {
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		return err
	}
	if rep.UploadID == nil {
		return fmt.Errorf("parse: report %s has no upload to parse", reportID)
	}
	up, err := d.Reports.Upload(ctx, *rep.UploadID)
	if err != nil {
		return err
	}

	// Everything that reads the upload runs under its own deadline; the
	// status writes below stay on ctx so a parse that times out can
	// still say so.
	readCtx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()

	body, err := d.Objects.Get(readCtx, up.ObjectKey)
	if err != nil {
		return d.failed(ctx, rep, session.Health{}, err)
	}
	defer body.Close()
	src, err := maybeDecompress(body)
	if err != nil {
		return d.failed(ctx, rep, session.Health{}, err)
	}
	defer src.Close()

	s := session.New(engine.SessionOptions(reportID, time.Now().UTC(), true))
	pub := store.Publisher{Keys: reports.Keys(reportID), Put: d.Objects}
	onClosed := func(c session.Closed) error { return d.writeFight(readCtx, rep, pub, c) }
	if err := stream(readCtx, s, src, onClosed); err != nil {
		return d.failed(ctx, rep, s.Health(), err)
	}

	health, err := json.Marshal(s.Health())
	if err != nil {
		return fmt.Errorf("parse: marshal health of %s: %w", reportID, err)
	}
	if err := d.Reports.SetStatus(ctx, reportID, reports.StatusComplete, engine.Version, health); err != nil {
		return err
	}
	rep.Status, rep.EngineVersion, rep.Health = reports.StatusComplete, engine.Version, health
	return d.writeReport(ctx, pub, rep, s)
}

// failed records a parse that could not finish, keeping the health the
// engine had reached so the report page can say why. The fights already
// published stay published: a log that broke halfway is still worth the
// half that parsed.
func (d Deps) failed(ctx context.Context, rep reports.Report, h session.Health, cause error) error {
	d.logger().Error("parse", "report", rep.ID, "err", cause)
	health, err := json.Marshal(h)
	if err != nil {
		return fmt.Errorf("parse: marshal health of %s: %w", rep.ID, err)
	}
	if err := d.Reports.SetStatus(ctx, rep.ID, reports.StatusFailed, engine.Version, health); err != nil {
		return err
	}
	return fmt.Errorf("parse: report %s failed: %w", rep.ID, cause)
}

// writeFight publishes one closed fight and indexes it, then rewrites
// report.json so the open page sees the new fight within a poll.
func (d Deps) writeFight(ctx context.Context, rep reports.Report, pub store.Publisher, c session.Closed) error {
	if err := pub.WriteFight(ctx, c.Fight.Index, c.Summary, c.Events); err != nil {
		return err
	}
	rec := reports.FightRecord{
		ReportID: rep.ID, Index: c.Fight.Index, Name: c.Fight.Name, Kill: c.Fight.Kill,
		DurationMS: c.Fight.Duration().Milliseconds(), StartMS: c.Fight.Start.UnixMilli(),
		// Parsed here from the original bytes, so there is nothing to
		// verify against: this is the verification.
		Verified: true, Players: c.Fight.Players, Deaths: c.Fight.Deaths, NPCKills: c.Fight.NPCKills,
		RawStart: &c.Fight.StartOffset, RawEnd: &c.Fight.EndOffset,
	}
	if c.Fight.EncounterID != 0 {
		rec.EncounterID, rec.Difficulty, rec.Size = &c.Fight.EncounterID, &c.Fight.Difficulty, &c.Fight.Size
		if !c.Fight.Kill {
			health := c.Fight.BossHealthPct
			rec.BossHealthPct = &health
		}
	}
	if _, err := d.Reports.UpsertFight(ctx, rec); err != nil {
		return err
	}
	if d.Rank != nil && c.Fight.EncounterID != 0 && reports.Ranked(rep.Visibility) {
		region, ruleset := reports.ReportRealm(rep)
		if err := d.Rank.WriteFight(ctx, reports.RankedFight{
			ReportID: rep.ID, FightIndex: c.Fight.Index, FoughtAt: c.Fight.Start.UTC(),
			Region: region, Ruleset: ruleset, GuildID: rep.GuildID,
			Rows: metrics.Derive(c.Fight, c.Summary), Combatants: c.Summary.Combatants,
			Factions: rankings.FactionsFromEvents(c.Events),
		}); err != nil {
			return err
		}
	}
	fights, err := d.Reports.Fights(ctx, rep.ID)
	if err != nil {
		return err
	}
	return pub.WriteReport(ctx, store.Report{
		ReportID: rep.ID, EngineVersion: engine.Version, Fights: entriesOf(fights),
	})
}

// writeReport writes the final report.json, with the health record and
// the unit list the page names players from.
func (d Deps) writeReport(ctx context.Context, pub store.Publisher, rep reports.Report, s *session.Session) error {
	fights, err := d.Reports.Fights(ctx, rep.ID)
	if err != nil {
		return err
	}
	return pub.WriteReport(ctx, store.Report{
		ReportID: rep.ID, EngineVersion: engine.Version, Health: s.Health(),
		Fights: entriesOf(fights), Units: s.Units().All(),
	})
}

func entriesOf(fights []reports.FightEntry) []store.FightEntry {
	out := make([]store.FightEntry, 0, len(fights))
	for _, f := range fights {
		out = append(out, f.FightEntry)
	}
	return out
}

// stream feeds a reader through a session in bounded chunks, calling
// onClosed for each fight as it closes. It reads at most maxDecoded
// bytes: the upload is a body of whatever length R2 hands back, and an
// unbounded read of it is an unbounded parse.
func stream(ctx context.Context, s *session.Session, src io.Reader, onClosed func(session.Closed) error) error {
	buf := make([]byte, chunkSize)
	offset := s.Offset()
	read := int64(0)
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		n, rerr := src.Read(buf)
		if n > 0 {
			if read += int64(n); read > maxDecoded {
				return fmt.Errorf("parse: the upload decodes to more than %d bytes", maxDecoded)
			}
			res, err := s.Feed(buf[:n], offset)
			if err != nil {
				return err
			}
			offset += int64(n)
			for _, c := range res.Closed {
				if err := onClosed(c); err != nil {
					return err
				}
			}
		}
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			return fmt.Errorf("parse: read upload: %w", rerr)
		}
	}
	res, err := s.Close()
	if err != nil {
		return err
	}
	for _, c := range res.Closed {
		if err := onClosed(c); err != nil {
			return err
		}
	}
	return nil
}

// zstdMagic is the first four bytes of a zstd frame.
var zstdMagic = []byte{0x28, 0xb5, 0x2f, 0xfd}

// maybeDecompress wraps r in a zstd decoder when the stream starts with
// a zstd frame, and hands it back as it is otherwise. The upload key
// ends in .zst, but a browser that could not compress the file still
// uploads a readable log, and a report is worth more than a naming rule.
func maybeDecompress(r io.ReadCloser) (io.ReadCloser, error) {
	br := bufio.NewReaderSize(r, chunkSize)
	head, err := br.Peek(len(zstdMagic))
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("parse: read upload: %w", err)
	}
	if !bytes.Equal(head, zstdMagic) {
		return io.NopCloser(br), nil
	}
	return zstdx.Reader(io.NopCloser(br))
}
