// companion/internal/client/ingest.go
// The ingest routes from the Phase 3 interface contract, one method
// each. Bodies are built here and handed to the queue as bytes, so an
// upload that fails today is replayed byte for byte tomorrow.
package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/character"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

// Character identifies the character a report is logged from. It is
// the shared triple, so the ruleset field is spelled in one place.
type Character = character.Character

// CreateReport is the body of POST /v1/reports.
type CreateReport struct {
	Title            string     `json:"title,omitempty"`
	Visibility       string     `json:"visibility"`
	Zone             string     `json:"zone,omitempty"`
	LoggingCharacter *Character `json:"logging_character,omitempty"`
}

// Report is the API's answer to a report creation.
type Report struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

// MetricsRow is one player's ranking row for one boss fight, with the
// field names the interface contract fixes. The engine's
// summary.MetricRow is one row per metric; the API wants one row per
// player, so MetricsRowsOf pivots the fight's roster instead.
type MetricsRow struct {
	PlayerGUID  string  `json:"player_guid"`
	Name        string  `json:"name"`
	Class       string  `json:"class"`
	Spec        string  `json:"spec"`
	Role        string  `json:"role"`
	ItemLevel   int64   `json:"ilvl"`
	MetricDPS   float64 `json:"metric_dps"`
	MetricHPS   float64 `json:"metric_hps"`
	DamageTaken int64   `json:"damage_taken"`
	ActiveMS    int64   `json:"active_ms"`
	Deaths      int     `json:"deaths"`
	EncounterID int64   `json:"encounter_id"`
	Difficulty  int64   `json:"difficulty"`
	Size        int64   `json:"size"`
	DurationMS  int64   `json:"duration_ms"`
	Kill        bool    `json:"kill"`
}

// MetricsRowsOf derives the fight's metrics rows from its roster. The
// order is the roster's, which the engine already sorts, so two runs
// over the same log post identical bytes.
func MetricsRowsOf(f fight.Fight, s summary.Summary) []MetricsRow {
	rows := make([]MetricsRow, 0, len(s.Roster))
	for _, r := range s.Roster {
		rows = append(rows, MetricsRow{
			PlayerGUID:  r.GUID,
			Name:        r.Name,
			Class:       r.Class,
			Spec:        r.Spec,
			Role:        r.Role,
			ItemLevel:   r.ItemLevel,
			MetricDPS:   r.DPS,
			MetricHPS:   r.HPS,
			DamageTaken: r.DamageTaken,
			ActiveMS:    r.ActiveMS,
			Deaths:      r.Deaths,
			EncounterID: f.EncounterID,
			Difficulty:  f.Difficulty,
			Size:        f.Size,
			DurationMS:  s.DurationMS,
			Kill:        f.Kill,
		})
	}
	return rows
}

// RawRange is the fight's byte range in the original log with the hash
// of those bytes, which the server's raw-sample check verifies later.
type RawRange struct {
	StartOffset int64  `json:"start_offset"`
	EndOffset   int64  `json:"end_offset"`
	SHA256      string `json:"sha256"`
}

// FightBundle is one closed fight, ready to be encoded once and stored
// in the queue until it lands.
type FightBundle struct {
	Summary  summary.Summary
	Events   []byte // Parquet
	Metrics  []MetricsRow
	RawRange RawRange
}

// BundleBoundary is pinned so an encoded bundle is a pure function of
// its content: the queue can compare two encodings byte for byte.
const BundleBoundary = "foreversixty-fight-bundle"

// Encode writes the multipart body the fight route takes.
func (b FightBundle) Encode() (contentType string, body []byte, err error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if err := w.SetBoundary(BundleBoundary); err != nil {
		return "", nil, err
	}
	writeJSON := func(field string, v any) error {
		p, err := w.CreateFormField(field)
		if err != nil {
			return err
		}
		enc := json.NewEncoder(p)
		return enc.Encode(v)
	}
	if err := writeJSON("summary", b.Summary); err != nil {
		return "", nil, err
	}
	p, err := w.CreateFormFile("events", "events.parquet")
	if err != nil {
		return "", nil, err
	}
	if _, err := p.Write(b.Events); err != nil {
		return "", nil, err
	}
	if err := writeJSON("metrics", b.Metrics); err != nil {
		return "", nil, err
	}
	if err := writeJSON("raw_range", b.RawRange); err != nil {
		return "", nil, err
	}
	if err := w.Close(); err != nil {
		return "", nil, err
	}
	return w.FormDataContentType(), buf.Bytes(), nil
}

// Live is the body of the live-snapshot route.
type Live struct {
	Summary   summary.Summary `json:"summary"`
	ElapsedMS int64           `json:"elapsed_ms"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// Complete is the body of the completion route.
type Complete struct {
	FinalOffset   int64          `json:"final_offset"`
	EngineVersion string         `json:"engine_version"`
	Health        session.Health `json:"health"`
}

// Stored says how the server took a write: Created for the first time,
// Duplicate when the same bytes were already stored. Both are success;
// the contract makes at-least-once delivery safe.
type Stored int

// The two outcomes of an idempotent write.
const (
	Created Stored = iota
	Duplicate
)

// CreateReport creates a report and returns its server-assigned id.
func (c *Client) CreateReport(ctx context.Context, in CreateReport) (Report, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return Report{}, err
	}
	var out Report
	if _, err := c.do(ctx, request{
		Method: http.MethodPost, Path: "/v1/reports",
		Body: body, Type: "application/json", Out: &out,
	}); err != nil {
		return Report{}, err
	}
	if out.ID == "" {
		return Report{}, fmt.Errorf("POST /v1/reports: the response carried no report id")
	}
	return out, nil
}

// PutFight uploads one closed fight. The body is pre-encoded so the
// queue can replay it unchanged.
func (c *Client) PutFight(ctx context.Context, reportID string, index int, contentType string, body []byte) (Stored, error) {
	status, err := c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/fights/" + strconv.Itoa(index),
		Body:   body, Type: contentType,
	})
	return storedOf(status), err
}

// PutLive uploads the running summary of the fight in progress. It is
// never queued: a snapshot that fails is replaced five seconds later.
func (c *Client) PutLive(ctx context.Context, reportID string, index int, in Live) error {
	body, err := json.Marshal(in)
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/fights/" + strconv.Itoa(index) + "/live",
		Body:   body, Type: "application/json",
	})
	return err
}

// PutRaw uploads one compressed chunk of the original log, addressed
// by the uncompressed byte offset of its first byte.
//
// X-Raw-SHA256 is the hash of the DECODED bytes, not of the frame in
// body: the server decodes the chunk, hashes what comes out and
// derives the end offset from its length. The caller passes that hash
// because it is the only party that still has the plaintext — the
// queue stores the compressed frame and nothing else.
func (c *Client) PutRaw(ctx context.Context, reportID string, offset int64, decodedSHA256 string, chunk []byte) (Stored, error) {
	status, err := c.do(ctx, request{
		Method: http.MethodPut,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/raw",
		Query:  url.Values{"offset": []string{strconv.FormatInt(offset, 10)}},
		Body:   chunk, Type: "application/zstd",
		Header: map[string]string{"X-Raw-SHA256": decodedSHA256},
	})
	return storedOf(status), err
}

// SHA256 is the hex digest the raw route and the fight bundle's
// raw_range both carry, over plaintext log bytes.
func SHA256(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

// Complete closes the report and schedules the server's raw-sample
// verification.
func (c *Client) Complete(ctx context.Context, reportID string, in Complete) error {
	body, err := json.Marshal(in)
	if err != nil {
		return err
	}
	_, err = c.do(ctx, request{
		Method: http.MethodPost,
		Path:   "/v1/reports/" + url.PathEscape(reportID) + "/complete",
		Body:   body, Type: "application/json",
	})
	return err
}

func storedOf(status int) Stored {
	if status == http.StatusOK {
		return Duplicate
	}
	return Created
}
