package parse

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/jhunthrop/foreversixty/api/internal/engine"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/zstdx"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// BackfillOptions is how one backfill runs.
type BackfillOptions struct {
	// Apply writes; false is the dry run.
	Apply bool
	// Layout, when set, names the registered layout the raw stream is parsed with, instead
	// of the one the report's health record names. It is for a report whose live parse
	// inferred its layout (a companion that started mid-file saw no header): the operator
	// names the dialect the night was written in and the engine stops guessing. The CLI
	// validates it against engine.LayoutByName; an unregistered name here infers as before.
	Layout string
}

// BackfillResult is what one backfill run saw, dry or applied.
type BackfillResult struct {
	ReportID string
	Applied  bool
	// Chunks and Bytes are the raw chunks streamed and the decoded bytes they held.
	Chunks int
	Bytes  int64
	// GapAt is the offset the raw stream stopped at because the next chunk did not continue
	// it, nil when every stored chunk was contiguous from zero.
	GapAt *int64
	// Fights is every fight the engine closed; Encounters and Kills are the boss fights among
	// them and the ones won.
	Fights, Encounters, Kills int
	Health                    session.Health
}

// Backfill parses a live report again from the raw chunks its companion uploaded, writing
// the fights the ingest refused at the time through the same writeFight/writeReport path
// the whole-file parse uses. From 10 Nov 2026 the API answered 400 to every fight bundle
// (reports.partBytes), so the week's live reports completed with no fights while their raw
// chunks landed intact; this is how those reports get their fights back.
//
// A dry run (apply false) streams and counts and writes nothing, not even the status. An
// applied run writes every fight, sets the status complete with the engine's health record,
// and runs the completion hook (zone, guild, feed). It never announces: the owner's channel
// was told nothing when the night ended, and a notification days later for a night already
// over would be noise. A report that already holds fights is refused, since the backfill
// would overwrite fights the live path verified.
func Backfill(ctx context.Context, d Deps, reportID string, o BackfillOptions) (BackfillResult, error) {
	apply := o.Apply
	res := BackfillResult{ReportID: reportID, Applied: apply}
	rep, err := d.Reports.Get(ctx, reportID)
	if err != nil {
		return res, err
	}
	existing, err := d.Reports.Fights(ctx, reportID)
	if err != nil {
		return res, err
	}
	if len(existing) > 0 {
		return res, fmt.Errorf("backfill: report %s already holds %d fights; refusing to overwrite them", reportID, len(existing))
	}
	chunks, err := d.Reports.RawChunks(ctx, reportID)
	if err != nil {
		return res, err
	}
	if len(chunks) == 0 {
		return res, fmt.Errorf("backfill: report %s has no raw chunks to parse", reportID)
	}

	readCtx, cancel := context.WithTimeout(ctx, readTimeout)
	defer cancel()
	src := &chunkStream{ctx: readCtx, d: d, reportID: reportID, chunks: chunks, res: &res}
	layoutName := o.Layout
	if layoutName == "" {
		layoutName = layoutOf(rep)
	}
	s := session.New(engine.SessionOptionsFor(reportID, rep.CreatedAt.UTC(), true, layoutName))
	pub := store.Publisher{Keys: reports.Keys(reportID), Put: d.Objects}
	var zone logZone
	onClosed := func(c session.Closed) error {
		zone.see(c.Fight)
		res.Fights++
		if c.Fight.EncounterID != 0 {
			res.Encounters++
			if c.Fight.Kill {
				res.Kills++
			}
		}
		if !apply {
			return nil
		}
		return d.writeFight(readCtx, rep, pub, c)
	}
	if err := stream(readCtx, s, src, onClosed); err != nil {
		return res, fmt.Errorf("backfill: report %s: %w", reportID, err)
	}
	res.Health = s.Health()
	if !apply {
		return res, nil
	}
	return res, d.finishBackfill(ctx, rep, pub, s, zone)
}

// finishBackfill is Report's ending without the announcement: status, the final report.json
// with the unit list, and the completion hook.
func (d Deps) finishBackfill(ctx context.Context, rep reports.Report, pub store.Publisher, s *session.Session, zone logZone) error {
	health, err := json.Marshal(s.Health())
	if err != nil {
		return fmt.Errorf("backfill: marshal health of %s: %w", rep.ID, err)
	}
	if err := d.Reports.SetStatus(ctx, rep.ID, reports.StatusComplete, engine.Version, health); err != nil {
		return err
	}
	rep.Status, rep.EngineVersion, rep.Health = reports.StatusComplete, engine.Version, health
	if _, err := d.writeReport(ctx, pub, rep, s); err != nil {
		return err
	}
	reports.FinishCompletion(ctx, d.Reports, d.Completed, d.logger(), rep.ID, zone.name(rep.Zone))
	return nil
}

// chunkStream reads a report's raw chunks back out of the bucket in offset order as one
// decoded stream, stopping (with GapAt set) at the first chunk that does not continue the
// one before it. Each chunk is fetched when the stream reaches it, so memory holds one
// decoded chunk at a time, never the night.
type chunkStream struct {
	ctx      context.Context
	d        Deps
	reportID string
	chunks   []reports.RawChunk
	res      *BackfillResult

	next   int
	cursor int64
	buf    []byte
	done   bool
}

func (c *chunkStream) Read(p []byte) (int, error) {
	for len(c.buf) == 0 {
		if c.done {
			return 0, io.EOF
		}
		if err := c.fill(); err != nil {
			return 0, err
		}
	}
	n := copy(p, c.buf)
	c.buf = c.buf[n:]
	return n, nil
}

// fill decodes the next chunk into buf, or marks the stream done at its end or its first gap.
func (c *chunkStream) fill() error {
	if c.next >= len(c.chunks) {
		c.done = true
		return nil
	}
	ch := c.chunks[c.next]
	if ch.Start != c.cursor {
		gap := c.cursor
		c.res.GapAt = &gap
		c.done = true
		return nil
	}
	body, err := c.d.Objects.Get(c.ctx, reports.Keys(c.reportID).Raw(ch.Start))
	if err != nil {
		return fmt.Errorf("read raw chunk %s@%d: %w", c.reportID, ch.Start, err)
	}
	packed, err := io.ReadAll(io.LimitReader(body, zstdx.MaxChunk))
	body.Close()
	if err != nil {
		return fmt.Errorf("read raw chunk %s@%d: %w", c.reportID, ch.Start, err)
	}
	decoded, err := zstdx.DecodeAll(packed, zstdx.MaxChunk)
	if err != nil {
		return fmt.Errorf("decode raw chunk %s@%d: %w", c.reportID, ch.Start, err)
	}
	if int64(len(decoded)) != ch.End-ch.Start {
		return fmt.Errorf("raw chunk %s@%d decodes to %d bytes, the record says %d",
			c.reportID, ch.Start, len(decoded), ch.End-ch.Start)
	}
	c.buf, c.cursor, c.next = decoded, ch.End, c.next+1
	c.res.Chunks++
	c.res.Bytes += int64(len(decoded))
	return nil
}

// LogBackfill writes one result as the CLI reports it: counts an operator reads before an
// apply, and the health figures a parse is judged by.
func LogBackfill(log *slog.Logger, r BackfillResult) {
	mode := "dry-run"
	if r.Applied {
		mode = "applied"
	}
	attrs := []any{"report", r.ReportID, "mode", mode, "chunks", r.Chunks, "bytes", r.Bytes,
		"fights", r.Fights, "encounters", r.Encounters, "kills", r.Kills,
		"lines", r.Health.Lines, "parse_errors", r.Health.ParseErrors, "layout", r.Health.Layout,
		"missing_header", r.Health.MissingHeader}
	if r.GapAt != nil {
		attrs = append(attrs, "gap_at", *r.GapAt)
	}
	log.Info("backfill-report", attrs...)
}
