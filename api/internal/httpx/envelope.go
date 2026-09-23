package httpx

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
)

type ctxKey int

const requestIDKey ctxKey = 1

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFrom(ctx context.Context) string {
	id, _ := ctx.Value(requestIDKey).(string)
	return id
}

type ErrorBody struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

type Envelope struct {
	OK        bool       `json:"ok"`
	Data      any        `json:"data"`
	Error     *ErrorBody `json:"error"`
	RequestID string     `json:"request_id"`
}

// etagHexLen is how many hex characters of the data field's sha256 form
// the weak ETag: 16 hex characters (8 bytes) makes two different
// answers to the same route collide only astronomically often, while
// keeping the header short (spec §2.1).
const etagHexLen = 16

// dataETag computes the spec §2.1 weak ETag over data alone, never the
// whole envelope: request_id differs on every request and would make
// the tag change even when nothing the caller reads has.
func dataETag(data any) (string, bool) {
	body, err := json.Marshal(data)
	if err != nil {
		return "", false
	}
	sum := sha256.Sum256(body)
	return `W/"` + hex.EncodeToString(sum[:])[:etagHexLen] + `"`, true
}

// ifNoneMatch reports whether etag (already in its W/"..." form)
// appears anywhere in header, a comma-separated If-None-Match value
// that may list a mix of weak and strong tags. RFC 9110 §13.1.2 says a
// weak comparison is always used for If-None-Match, so a strong-form
// copy of the same tag matches too.
func ifNoneMatch(header, etag string) bool {
	want := strings.TrimPrefix(etag, "W/")
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimPrefix(strings.TrimSpace(part), "W/")
		if part == want {
			return true
		}
	}
	return false
}

// write is every response's one exit door. A successful (2xx) answer
// to a GET request gets a data ETag and, on a matching If-None-Match,
// a bodyless 304 carrying that same ETag plus whatever Cache-Control
// (and Vary) the handler already set on w before calling WriteOK -
// CachePublic/CachePrivate always run first, so those headers are
// already staged on the ResponseWriter by the time write sees them.
func write(w http.ResponseWriter, r *http.Request, status int, e Envelope) {
	if r != nil && r.Method == http.MethodGet && status >= 200 && status < 300 {
		if etag, ok := dataETag(e.Data); ok {
			w.Header().Set("ETag", etag)
			if inm := r.Header.Get("If-None-Match"); inm != "" && ifNoneMatch(inm, etag) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(e)
}

func WriteOK(w http.ResponseWriter, r *http.Request, status int, data any) {
	write(w, r, status, Envelope{OK: true, Data: data, RequestID: RequestIDFrom(r.Context())})
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, code, message string, fields map[string]string) {
	write(w, r, status, Envelope{OK: false, Error: &ErrorBody{Code: code, Message: message, Fields: fields}, RequestID: RequestIDFrom(r.Context())})
}
