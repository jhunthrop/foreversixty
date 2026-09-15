package reports

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

// fakeMultipart is the multipart half of the R2 client.
type fakeMultipart struct {
	started   map[string]string
	completed map[string][]r2.CompletedPart
	aborted   []string
	err       error
	// presignErr fails only PresignParts, so a test can exercise a
	// multipart that started successfully in R2 but never got signed.
	presignErr error
}

func newFakeMultipart() *fakeMultipart {
	return &fakeMultipart{started: map[string]string{}, completed: map[string][]r2.CompletedPart{}}
}

func (f *fakeMultipart) StartMultipart(_ context.Context, key, _ string) (string, error) {
	if f.err != nil {
		return "", f.err
	}
	id := "r2-" + key
	f.started[key] = id
	return id, nil
}

func (f *fakeMultipart) PresignParts(_ context.Context, key, uploadID string, parts int, ttl time.Duration) ([]r2.Part, error) {
	if f.err != nil {
		return nil, f.err
	}
	if f.presignErr != nil {
		return nil, f.presignErr
	}
	out := make([]r2.Part, 0, parts)
	for n := 1; n <= parts; n++ {
		out = append(out, r2.Part{Number: n, URL: "https://r2.example/" + key + "?part=" + strconv.Itoa(n)})
	}
	return out, nil
}

func (f *fakeMultipart) CompleteMultipart(_ context.Context, key, uploadID string, parts []r2.CompletedPart) error {
	if f.err != nil {
		return f.err
	}
	f.completed[key] = parts
	return nil
}

func (f *fakeMultipart) AbortMultipart(_ context.Context, key, uploadID string) error {
	f.aborted = append(f.aborted, key)
	return nil
}

func TestAWholeFileUploadIsSignedCompletedAndParsed(t *testing.T) {
	h := newHarness(t)
	h.asSession()

	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":134217728,"filename":"WoWCombatLog.txt"}`)
	var start struct {
		UploadID    string    `json:"upload_id"`
		Parts       []r2.Part `json:"parts"`
		CompleteURL string    `json:"complete_url"`
	}
	h.data(res, &start)
	if len(start.UploadID) != 16 {
		t.Fatalf("upload id = %q", start.UploadID)
	}
	if len(start.Parts) != 2 {
		t.Fatalf("parts = %+v, want two 64 MiB parts", start.Parts)
	}
	if start.CompleteURL != "https://api.foreversixty.gg/v1/uploads/"+start.UploadID+"/complete" {
		t.Fatalf("complete_url = %q", start.CompleteURL)
	}

	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"\"a\""},{"number":2,"etag":"b"}],"title":"Tuesday","visibility":"unlisted"}`)
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("status = %d, want 202", res.StatusCode)
	}
	var complete struct {
		ReportID string `json:"report_id"`
	}
	h.data(res, &complete)
	if complete.ReportID == "" {
		t.Fatal("no report was made")
	}
	if got := len(h.parts.completed[r2.UploadKey(start.UploadID)]); got != 2 {
		t.Fatalf("%d parts completed, want 2", got)
	}
	ran := h.jobs.Ran()
	if len(ran) != 1 || ran[0][0] != ParseJobCommand || ran[0][1] != complete.ReportID {
		t.Fatalf("job ran with %v", ran)
	}
	rep, err := h.store.Get(t.Context(), complete.ReportID)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != StatusProcessing || rep.Visibility != Unlisted || rep.UploadID == nil {
		t.Fatalf("report = %+v", rep)
	}

	// Completing again is harmless and names the same report. It also
	// starts no second parse job: an at-least-once retry of a completion
	// that already succeeded must not launch a second Cloud Run execution.
	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"a"}]}`)
	var again struct {
		ReportID string `json:"report_id"`
	}
	h.data(res, &again)
	if again.ReportID != complete.ReportID {
		t.Fatalf("a second completion made %q, want %q", again.ReportID, complete.ReportID)
	}
	if ran := h.jobs.Ran(); len(ran) != 1 {
		t.Fatalf("job ran %d times after a retry of a completed upload, want 1: %v", len(ran), ran)
	}
}

func TestUploadsRefuseNonsense(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	for name, body := range map[string]string{
		"empty":    `{"size_bytes":0}`,
		"too big":  `{"size_bytes":5000000000}`,
		"not json": `{`,
	} {
		res := h.json(http.MethodPost, "/v1/uploads", body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
	res := h.json(http.MethodPost, "/v1/uploads/nosuchupload/complete", `{"etags":[{"number":1,"etag":"a"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("an unknown upload = %d, want 404", res.StatusCode)
	}
	res = h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	var start struct {
		UploadID string `json:"upload_id"`
	}
	h.data(res, &start)
	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete", `{"etags":[]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("no parts = %d, want 400", res.StatusCode)
	}
	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"a"}],"visibility":"secret"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("a bad visibility = %d, want 400", res.StatusCode)
	}
}

func TestAnotherPersonsUploadCannotBeCompleted(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	var start struct {
		UploadID string `json:"upload_id"`
	}
	h.data(res, &start)

	h.actor = auth.Actor{UserID: h.owner + 3, Role: "user", Method: "session"}
	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"a"}]}`)
	res.Body.Close()
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("status = %d, want 403", res.StatusCode)
	}
}

func TestUploadsNeedABrowserSession(t *testing.T) {
	h := newHarness(t)
	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401 for a device token", res.StatusCode)
	}
}

func TestTrimFilenameKeepsTheBaseName(t *testing.T) {
	for in, want := range map[string]string{
		`C:\Games\WoW\Logs\WoWCombatLog.txt`: "WoWCombatLog.txt",
		"/home/me/WoWCombatLog.txt":          "WoWCombatLog.txt",
		"  spaced.txt  ":                     "spaced.txt",
	} {
		if got := trimFilename(in); got != want {
			t.Errorf("trimFilename(%q) = %q, want %q", in, got, want)
		}
	}
	if got := trimFilename(string(make([]byte, 300))); len(got) != 120 {
		t.Errorf("a long name was not bounded: %d", len(got))
	}
}

// A filename over the bound that is not ASCII used to be byte-sliced,
// which can split a multi-byte rune; uploads.filename is a text column
// and Postgres refuses the result with "invalid byte sequence for
// encoding UTF8", so the upload 500s on every attempt.
func TestALongNonASCIIFilenameStartsAnUpload(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	// Each "日" is three bytes, so 50 of them is 150 bytes and a cut at
	// 120 lands inside the forty-first rune.
	name := strings.Repeat("日", 50) + ".txt"

	body, err := json.Marshal(map[string]any{"size_bytes": 10, "filename": name})
	if err != nil {
		t.Fatal(err)
	}
	res := h.json(http.MethodPost, "/v1/uploads", string(body))
	var start struct {
		UploadID string `json:"upload_id"`
	}
	h.data(res, &start)

	var stored string
	if err := h.store.Pool.QueryRow(t.Context(),
		`select filename from uploads where id = $1`, start.UploadID).Scan(&stored); err != nil {
		t.Fatal(err)
	}
	if !utf8.ValidString(stored) {
		t.Fatalf("stored filename is not valid UTF-8: %q", stored)
	}
	if len(stored) > 120 || stored != strings.Repeat("日", 40) {
		t.Fatalf("stored filename = %q (%d bytes), want the whole-rune prefix", stored, len(stored))
	}
}

func TestStartingAnUploadAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	h.store.Pool.Close()
	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", res.StatusCode)
	}
	// CreateUpload failed after StartMultipart already opened the
	// multipart in R2: that upload must be aborted, not orphaned.
	var key string
	for k := range h.parts.started {
		key = k
	}
	if len(h.parts.aborted) != 1 || h.parts.aborted[0] != key {
		t.Fatalf("aborted = %v, want [%q]", h.parts.aborted, key)
	}
}

// TestStartAbortsTheMultipartWhenPresigningFails proves the start path
// does not orphan a multipart upload in R2 when a step after
// StartMultipart fails: the upload it opened must be aborted before the
// 500 is answered.
func TestStartAbortsTheMultipartWhenPresigningFails(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	h.parts.presignErr = errors.New("presign failed")

	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", res.StatusCode)
	}
	if len(h.parts.started) != 1 {
		t.Fatalf("started = %v, want exactly one multipart start", h.parts.started)
	}
	var key string
	for k := range h.parts.started {
		key = k
	}
	if len(h.parts.aborted) != 1 || h.parts.aborted[0] != key {
		t.Fatalf("aborted = %v, want [%q]", h.parts.aborted, key)
	}
}

// TestARetryAfterTheReportRowWasNeverCreatedConverges reproduces the
// state a partial failure between R2.CompleteMultipart succeeding and
// Store.Create running would leave: the upload row already names a
// reserved report id, but that report does not exist yet. A retried
// completion must create the report under that same id, start the job,
// and never call CompleteMultipart again on an upload R2 already
// finalized.
func TestARetryAfterTheReportRowWasNeverCreatedConverges(t *testing.T) {
	h := newHarness(t)
	h.asSession()

	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	var start struct {
		UploadID string `json:"upload_id"`
	}
	h.data(res, &start)

	reserved := auth.NewReportID()
	if err := h.store.FinishUpload(t.Context(), start.UploadID, reserved); err != nil {
		t.Fatal(err)
	}

	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"a"}]}`)
	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("status = %d, want 202", res.StatusCode)
	}
	var out struct {
		ReportID string `json:"report_id"`
	}
	h.data(res, &out)
	if out.ReportID != reserved {
		t.Fatalf("report id = %q, want the reserved id %q", out.ReportID, reserved)
	}
	if len(h.parts.completed) != 0 {
		t.Fatalf("CompleteMultipart was called on an upload R2 had already finalized: %v", h.parts.completed)
	}
	rep, err := h.store.Get(t.Context(), reserved)
	if err != nil {
		t.Fatal(err)
	}
	if rep.Status != StatusProcessing {
		t.Fatalf("report = %+v", rep)
	}
	var count int
	if err := h.store.Pool.QueryRow(t.Context(),
		`select count(*) from reports where upload_id = $1`, start.UploadID).Scan(&count); err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("reports for upload = %d, want exactly 1", count)
	}
	ran := h.jobs.Ran()
	if len(ran) != 1 || ran[0][0] != ParseJobCommand || ran[0][1] != reserved {
		t.Fatalf("job ran with %v, want exactly one run for %q", ran, reserved)
	}
}
