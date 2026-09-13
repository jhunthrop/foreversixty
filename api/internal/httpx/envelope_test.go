package httpx

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
