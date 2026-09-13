package httpx

import (
	"context"
	"encoding/json"
	"net/http"
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

func write(w http.ResponseWriter, status int, e Envelope) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(e)
}

func WriteOK(w http.ResponseWriter, r *http.Request, status int, data any) {
	write(w, status, Envelope{OK: true, Data: data, RequestID: RequestIDFrom(r.Context())})
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, code, message string, fields map[string]string) {
	write(w, status, Envelope{OK: false, Error: &ErrorBody{Code: code, Message: message, Fields: fields}, RequestID: RequestIDFrom(r.Context())})
}
