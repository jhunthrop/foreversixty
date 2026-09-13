package httpx

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func Chain(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func RequestID() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id := r.Header.Get("X-Request-Id")
			if id == "" {
				b := make([]byte, 8)
				_, _ = rand.Read(b)
				id = hex.EncodeToString(b)
			}
			w.Header().Set("X-Request-Id", id)
			next.ServeHTTP(w, r.WithContext(WithRequestID(r.Context(), id)))
		})
	}
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (s *statusWriter) WriteHeader(code int) { s.status = code; s.ResponseWriter.WriteHeader(code) }

func Logger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			sw := &statusWriter{ResponseWriter: w, status: 200}
			next.ServeHTTP(sw, r)
			log.Info("request", "id", RequestIDFrom(r.Context()), "method", r.Method, "path", r.URL.Path, "status", sw.status, "ms", time.Since(start).Milliseconds())
		})
	}
}

func Recover(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Error("panic", "id", RequestIDFrom(r.Context()), "method", r.Method, "path", r.URL.Path, "err", rec)
					WriteError(w, r, http.StatusInternalServerError, "internal", "internal error", nil)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

type limiterEntry struct {
	lim  *rate.Limiter
	seen time.Time
}

func RateLimit(perMinute int) func(http.Handler) http.Handler {
	var mu sync.Mutex
	entries := map[string]*limiterEntry{}
	get := func(ip string) *rate.Limiter {
		mu.Lock()
		defer mu.Unlock()
		now := time.Now()
		for k, e := range entries {
			if now.Sub(e.seen) > 10*time.Minute {
				delete(entries, k)
			}
		}
		e, ok := entries[ip]
		if !ok {
			e = &limiterEntry{lim: rate.NewLimiter(rate.Every(time.Minute/time.Duration(perMinute)), perMinute)}
			entries[ip] = e
		}
		e.seen = now
		return e.lim
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r)
			if !get(ip).Allow() {
				WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// clientIP trusts the first X-Forwarded-For hop, which Cloud Run sets from the
// connecting client; falls back to RemoteAddr for local runs and tests.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if i := strings.IndexByte(xff, ','); i >= 0 {
			return strings.TrimSpace(xff[:i])
		}
		return strings.TrimSpace(xff)
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func CORS(allowedOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Origin") == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Vary", "Origin")
			}
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
