package reports

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

// fakeMultipart is the multipart half of the R2 client.
type fakeMultipart struct {
	started   map[string]string
	completed map[string][]r2.CompletedPart
	aborted   []string
	err       error
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
	out := make([]r2.Part, 0, parts)
	for n := 1; n <= parts; n++ {
		out = append(out, r2.Part{Number: n, URL: "https://r2.example/" + key + "?part=" + itoa(n)})
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

func itoa(n int) string { return string(rune('0' + n)) }

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

	// Completing again is harmless and names the same report.
	res = h.json(http.MethodPost, "/v1/uploads/"+start.UploadID+"/complete",
		`{"etags":[{"number":1,"etag":"a"}]}`)
	var again struct {
		ReportID string `json:"report_id"`
	}
	h.data(res, &again)
	if again.ReportID != complete.ReportID {
		t.Fatalf("a second completion made %q, want %q", again.ReportID, complete.ReportID)
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

func TestStartingAnUploadAnswers500WhenTheDatabaseIsGone(t *testing.T) {
	h := newHarness(t)
	h.asSession()
	h.store.Pool.Close()
	res := h.json(http.MethodPost, "/v1/uploads", `{"size_bytes":10,"filename":"x"}`)
	res.Body.Close()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", res.StatusCode)
	}
}
