package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestWriteOKEnvelope(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req = req.WithContext(WithRequestID(req.Context(), "abc"))
	WriteOK(rec, req, http.StatusOK, map[string]string{"status": "ok"})
	var body Envelope
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || body.RequestID != "abc" || body.Error != nil {
		t.Fatalf("unexpected envelope %+v", body)
	}
}

func TestWriteErrorEnvelope(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteError(rec, req, http.StatusBadRequest, "invalid", "bad input", map[string]string{"email": "required"})
	if rec.Code != 400 {
		t.Fatalf("status = %d", rec.Code)
	}
	var body Envelope
	_ = json.NewDecoder(rec.Body).Decode(&body)
	if body.OK || body.Error == nil || body.Error.Code != "invalid" || body.Error.Fields["email"] != "required" {
		t.Fatalf("unexpected envelope %+v", body)
	}
}

func TestWriteOKSetsETag(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteOK(rec, req, http.StatusOK, map[string]string{"a": "b"})
	etag := rec.Header().Get("ETag")
	if etag == "" || !strings.HasPrefix(etag, `W/"`) || !strings.HasSuffix(etag, `"`) {
		t.Fatalf("ETag = %q, want a weak tag", etag)
	}
	if len(etag) != len(`W/"`)+16+1 {
		t.Fatalf("ETag = %q, want 16 hex chars inside the quotes", etag)
	}
}

func TestWriteOK304OnMatchingIfNoneMatch(t *testing.T) {
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteOK(rec1, req1, http.StatusOK, map[string]string{"a": "b"})
	etag := rec1.Header().Get("ETag")

	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Set("If-None-Match", etag)
	WriteOK(rec2, req2, http.StatusOK, map[string]string{"a": "b"})
	if rec2.Code != http.StatusNotModified {
		t.Fatalf("status = %d, want 304", rec2.Code)
	}
	if rec2.Body.Len() != 0 {
		t.Fatalf("body = %q, want empty on 304", rec2.Body.String())
	}
	if rec2.Header().Get("ETag") != etag {
		t.Fatalf("304 ETag = %q, want %q", rec2.Header().Get("ETag"), etag)
	}
}

func TestWriteOK304CarriesCacheControlSetBeforehand(t *testing.T) {
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	CachePublic(rec1, 30*time.Second, 300*time.Second)
	WriteOK(rec1, req1, http.StatusOK, map[string]string{"a": "b"})
	etag := rec1.Header().Get("ETag")

	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Set("If-None-Match", etag)
	CachePublic(rec2, 30*time.Second, 300*time.Second)
	WriteOK(rec2, req2, http.StatusOK, map[string]string{"a": "b"})
	if rec2.Code != http.StatusNotModified {
		t.Fatalf("status = %d, want 304", rec2.Code)
	}
	want := "public, max-age=30, stale-while-revalidate=300"
	if got := rec2.Header().Get("Cache-Control"); got != want {
		t.Fatalf("Cache-Control = %q, want %q", got, want)
	}
}

func TestWriteOK200OnMismatchedIfNoneMatch(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", `W/"0000000000000000"`)
	WriteOK(rec, req, http.StatusOK, map[string]string{"a": "b"})
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if rec.Body.Len() == 0 {
		t.Fatalf("body is empty, want the full envelope on a mismatch")
	}
}

func TestWriteOKMatchesStrongTagInIfNoneMatch(t *testing.T) {
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteOK(rec1, req1, http.StatusOK, map[string]string{"a": "b"})
	etag := rec1.Header().Get("ETag")
	strong := strings.TrimPrefix(etag, "W/")

	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Set("If-None-Match", strong)
	WriteOK(rec2, req2, http.StatusOK, map[string]string{"a": "b"})
	if rec2.Code != http.StatusNotModified {
		t.Fatalf("status = %d, want 304 on a strong-form match", rec2.Code)
	}
}

func TestWriteOKMatchesOneOfACommaList(t *testing.T) {
	rec1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteOK(rec1, req1, http.StatusOK, map[string]string{"a": "b"})
	etag := rec1.Header().Get("ETag")

	rec2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.Header.Set("If-None-Match", `W/"deadbeefdeadbeef", `+etag)
	WriteOK(rec2, req2, http.StatusOK, map[string]string{"a": "b"})
	if rec2.Code != http.StatusNotModified {
		t.Fatalf("status = %d, want 304 when the tag is anywhere in the list", rec2.Code)
	}
}

func TestWriteErrorNeverSetsETag(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	WriteError(rec, req, http.StatusBadRequest, "invalid", "bad", nil)
	if rec.Header().Get("ETag") != "" {
		t.Fatalf("ETag = %q, want none on an error response", rec.Header().Get("ETag"))
	}
}

func TestWriteOKPostNeverSetsETag(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	WriteOK(rec, req, http.StatusCreated, map[string]string{"a": "b"})
	if rec.Header().Get("ETag") != "" {
		t.Fatalf("ETag = %q, want none on POST", rec.Header().Get("ETag"))
	}
}
