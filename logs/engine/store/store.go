// logs/engine/store/store.go
// Package store writes the report files the edge serves. Putter is the
// whole storage contract: one call that puts bytes at a key with the
// caching headers the spec's storage table specifies, which an S3 client
// for R2 satisfies as directly as the local directory here does.
package store

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/logs/engine/event"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/parquet"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
	"github.com/jhunthrop/foreversixty/logs/engine/units"
)

// PutOptions are the object headers. They map one to one onto S3 PutObject.
type PutOptions struct {
	ContentType     string
	CacheControl    string
	ContentEncoding string
}

// Putter writes one object. An R2 or S3 implementation satisfies this with
// a single PutObject call; the local Dir below writes a file.
type Putter interface {
	Put(ctx context.Context, key string, body []byte, o PutOptions) error
}

// Cache headers from the spec's storage-layout table.
const (
	cacheMutable   = "public, max-age=5"
	cacheImmutable = "public, max-age=31536000, immutable"
	cachePrivate   = "private, max-age=31536000, immutable"
)

// Keys builds the object keys for one report.
type Keys struct{ ReportID string }

// Report is reports/<id>/report.json.
func (k Keys) Report() string { return "reports/" + k.ReportID + "/report.json" }

// FightSummary is reports/<id>/fights/<n>/summary.json.
func (k Keys) FightSummary(n int) string { return k.fightDir(n) + "/summary.json" }

// FightEvents is reports/<id>/fights/<n>/events.parquet.
func (k Keys) FightEvents(n int) string { return k.fightDir(n) + "/events.parquet" }

// FightLive is reports/<id>/fights/<n>/live.json.
func (k Keys) FightLive(n int) string { return k.fightDir(n) + "/live.json" }

// Raw is reports/<id>/raw/<offset>.zst.
func (k Keys) Raw(offset int64) string {
	return "reports/" + k.ReportID + "/raw/" + strconv.FormatInt(offset, 10) + ".zst"
}

func (k Keys) fightDir(n int) string {
	return "reports/" + k.ReportID + "/fights/" + strconv.Itoa(n)
}

// Dir is a Putter backed by a local directory, for the CLI and the tests.
type Dir struct{ Root string }

// NewDir returns a local store rooted at root.
func NewDir(root string) *Dir { return &Dir{Root: root} }

// Put writes one object as a file. Headers are not stored: the local store
// exists to inspect output, and R2 carries the headers in production.
func (d *Dir) Put(ctx context.Context, key string, body []byte, _ PutOptions) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	path, err := safeJoin(d.Root, key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("store: create %s: %w", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, body, 0o644); err != nil {
		return fmt.Errorf("store: write %s: %w", path, err)
	}
	return nil
}

// safeJoin resolves key under root, rejecting any key that would resolve
// outside it. The check is purely lexical (via filepath.Rel against the
// cleaned root) so it gives the same answer regardless of whether root is
// ".", relative, absolute, or "/" itself, and it never touches the
// filesystem.
func safeJoin(root, key string) (string, error) {
	cleanRoot := filepath.Clean(root)
	path := filepath.Join(cleanRoot, filepath.FromSlash(key))
	rel, err := filepath.Rel(cleanRoot, path)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("store: key %q escapes root %s", key, root)
	}
	return path, nil
}

// FightEntry is one fight's line in report.json.
type FightEntry struct {
	Index       int       `json:"index"`
	Kind        string    `json:"kind"`
	Name        string    `json:"name"`
	EncounterID int64     `json:"encounter_id,omitempty"`
	Difficulty  int64     `json:"difficulty,omitempty"`
	Size        int64     `json:"size,omitempty"`
	Kill        bool      `json:"kill"`
	InProgress  bool      `json:"in_progress"`
	Zone        string    `json:"zone,omitempty"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	DurationMS  int64     `json:"duration_ms"`
	Players     []string  `json:"players"`
	Deaths      int       `json:"deaths"`
	NPCKills    int       `json:"npc_kills"`
}

// EntryOf projects a fight onto its report.json line.
func EntryOf(f fight.Fight) FightEntry {
	return FightEntry{
		Index: f.Index, Kind: string(f.Kind), Name: f.Name,
		EncounterID: f.EncounterID, Difficulty: f.Difficulty, Size: f.Size,
		Kill: f.Kill, InProgress: f.InProgress, Zone: f.Zone,
		Start: f.Start, End: f.End, DurationMS: f.Duration().Milliseconds(),
		Players: f.Players, Deaths: f.Deaths, NPCKills: f.NPCKills,
	}
}

// Report is reports/<id>/report.json.
type Report struct {
	ReportID      string         `json:"report_id"`
	EngineVersion string         `json:"engine_version"`
	Health        session.Health `json:"health"`
	Fights        []FightEntry   `json:"fights"`
	Units         []units.Unit   `json:"units"`
}

// Publisher writes a report's files through a Putter.
type Publisher struct {
	Keys Keys
	Put  Putter
}

// WriteReport writes report.json. It is mutable, so it gets the short cache.
func (p Publisher) WriteReport(ctx context.Context, r Report) error {
	sorted := append([]FightEntry(nil), r.Fights...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Index < sorted[j].Index })
	r.Fights = sorted
	b, err := marshal(r)
	if err != nil {
		return err
	}
	return p.Put.Put(ctx, p.Keys.Report(), b, PutOptions{
		ContentType: "application/json", CacheControl: cacheMutable,
	})
}

// WriteFight writes a closed fight's summary and events. Both are
// immutable once written.
func (p Publisher) WriteFight(ctx context.Context, n int, s summary.Summary, events []event.Event) error {
	b, err := marshal(s)
	if err != nil {
		return err
	}
	if err := p.Put.Put(ctx, p.Keys.FightSummary(n), b, PutOptions{
		ContentType: "application/json", CacheControl: cacheImmutable,
	}); err != nil {
		return err
	}
	pq, err := parquet.Marshal(events)
	if err != nil {
		return err
	}
	return p.Put.Put(ctx, p.Keys.FightEvents(n), pq, PutOptions{
		ContentType: "application/vnd.apache.parquet", CacheControl: cacheImmutable,
	})
}

// WriteLive writes the snapshot of a fight in progress.
func (p Publisher) WriteLive(ctx context.Context, n int, s summary.Summary) error {
	b, err := marshal(s)
	if err != nil {
		return err
	}
	return p.Put.Put(ctx, p.Keys.FightLive(n), b, PutOptions{
		ContentType: "application/json", CacheControl: cacheMutable,
	})
}

// WriteRaw writes one compressed chunk of the original log, addressed by
// its byte offset so a re-sent chunk overwrites itself harmlessly.
func (p Publisher) WriteRaw(ctx context.Context, offset int64, chunk []byte) error {
	packed, err := compress(chunk)
	if err != nil {
		return err
	}
	return p.Put.Put(ctx, p.Keys.Raw(offset), packed, PutOptions{
		ContentType: "application/zstd", CacheControl: cachePrivate,
		ContentEncoding: "zstd",
	})
}

func marshal(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("store: marshal: %w", err)
	}
	return b, nil
}

// compress packs a raw chunk. The encoder is created per call with fixed
// settings so the output depends only on the input.
func compress(chunk []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := zstd.NewWriter(&buf,
		zstd.WithEncoderLevel(zstd.SpeedDefault),
		zstd.WithEncoderConcurrency(1),
	)
	if err != nil {
		return nil, fmt.Errorf("store: zstd writer: %w", err)
	}
	if _, err := w.Write(chunk); err != nil {
		return nil, fmt.Errorf("store: zstd write: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("store: zstd close: %w", err)
	}
	return buf.Bytes(), nil
}

// Decompress unpacks a raw chunk, for the reprocess job and the tests.
//
// TODO: cap the decoded size. DecodeAll's default ceiling is 64 GiB, which
// is only safe while every object this reads is one this engine wrote.
// The trigger is the Phase 3 reprocess job, which will read R2 objects
// whose provenance is not guaranteed: pass zstd.WithDecoderMaxMemory for a
// bound derived from the chunk size the publisher writes.
//
// The reader argument is nil because DecodeAll is stateless and ignores
// it; passing a real one would only spawn a decoder goroutine per call.
func Decompress(packed []byte) ([]byte, error) {
	r, err := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
	if err != nil {
		return nil, fmt.Errorf("store: zstd reader: %w", err)
	}
	defer r.Close()
	out, err := r.DecodeAll(packed, nil)
	if err != nil {
		return nil, fmt.Errorf("store: zstd decode: %w", err)
	}
	return out, nil
}
