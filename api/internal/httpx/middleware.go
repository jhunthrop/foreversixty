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

const (
	// limiterSweepEvery controls how often (in calls, not requests-per-IP) the
	// idle-entry sweep runs. Sweeping on every call made the map walk cost
	// O(n) per request; sweeping every Nth call amortizes that cost.
	limiterSweepEvery = 1024
	// limiterIdleTimeout is how long an entry may sit unused before the
	// periodic sweep (every limiterSweepEvery calls) reclaims it.
	limiterIdleTimeout = 10 * time.Minute
	// limiterMaxEntries caps the map size so a flood of distinct IPs cannot
	// grow it without bound.
	limiterMaxEntries = 10000
	// limiterCapIdleTimeout is the shorter idle window used only when the map
	// is at limiterMaxEntries and needs room for a new entry.
	limiterCapIdleTimeout = 1 * time.Minute
)

// ipLimiter tracks a per-IP token bucket. It sweeps idle entries
// periodically (every limiterSweepEvery calls, not every call) and caps its
// size at limiterMaxEntries so an attacker spraying distinct source IPs (or
// spoofed X-Forwarded-For values) cannot grow it without bound.
type ipLimiter struct {
	mu        sync.Mutex
	perMinute int
	entries   map[string]*limiterEntry
	calls     uint64
}

func newIPLimiter(perMinute int) *ipLimiter {
	return &ipLimiter{perMinute: perMinute, entries: map[string]*limiterEntry{}}
}

// sweepLocked removes entries idle for longer than idle. Callers must hold mu.
func (l *ipLimiter) sweepLocked(now time.Time, idle time.Duration) {
	for k, e := range l.entries {
		if now.Sub(e.seen) > idle {
			delete(l.entries, k)
		}
	}
}

// reserve returns the limiter for ip, creating one if needed. The second
// return value is false when the map is at limiterMaxEntries and has no idle
// entries to evict; the caller should treat the request as allowed in that
// case rather than reject legitimate traffic or grow the map unbounded.
func (l *ipLimiter) reserve(ip string, now time.Time) (*rate.Limiter, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calls++
	if l.calls%limiterSweepEvery == 0 {
		l.sweepLocked(now, limiterIdleTimeout)
	}
	if e, ok := l.entries[ip]; ok {
		e.seen = now
		return e.lim, true
	}
	if len(l.entries) >= limiterMaxEntries {
		l.sweepLocked(now, limiterCapIdleTimeout)
		if len(l.entries) >= limiterMaxEntries {
			// At capacity with nothing idle to evict: refuse to add a new
			// entry and let the request through untracked. Deliberately not
			// logged - this path is reachable at high volume during an
			// actual flood, and logging every occurrence would just move the
			// pressure from memory to the log pipeline.
			return nil, false
		}
	}
	e := &limiterEntry{lim: rate.NewLimiter(rate.Every(time.Minute/time.Duration(l.perMinute)), l.perMinute), seen: now}
	l.entries[ip] = e
	return e.lim, true
}

func (l *ipLimiter) allow(ip string, now time.Time) bool {
	lim, ok := l.reserve(ip, now)
	if !ok {
		return true
	}
	return lim.Allow()
}

// RateLimit rate-limits requests per client IP, resolved via clientIP with
// trustedHops trusted reverse proxies in front of this service.
func RateLimit(perMinute, trustedHops int) func(http.Handler) http.Handler {
	l := newIPLimiter(perMinute)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r, trustedHops)
			if !l.allow(ip, time.Now()) {
				WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// clientIP resolves the request's client address, trusting exactly
// trustedHops reverse proxies in front of this service.
//
// X-Forwarded-For is appended to by each proxy the request passes through,
// so it reads [original-client, proxy-1, proxy-2, ...]; a client can put
// anything they like in the leftmost entries. With trustedHops trusted
// proxies, the real client is the trustedHops-th entry counted from the
// right (trustedHops=1 means the rightmost entry). If the header has fewer
// entries than trustedHops - including when it is absent - the header is
// untrustworthy and RemoteAddr (the last hop we ourselves observed) is used
// instead. trustedHops=0 ignores X-Forwarded-For entirely.
func clientIP(r *http.Request, trustedHops int) string {
	if trustedHops > 0 {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			hops := strings.Split(xff, ",")
			if len(hops) >= trustedHops {
				return strings.TrimSpace(hops[len(hops)-trustedHops])
			}
		}
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
