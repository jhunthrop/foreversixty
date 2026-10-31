package httpx

import (
	"crypto/rand"
	"crypto/sha256"
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
			log.Info("request", "id", RequestIDFrom(r.Context()), "method", r.Method, "path", LogPath(r.URL.Path), "status", sw.status, "ms", time.Since(start).Milliseconds())
		})
	}
}

func Recover(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Error("panic", "id", RequestIDFrom(r.Context()), "method", r.Method, "path", LogPath(r.URL.Path), "err", rec)
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
	mu      sync.Mutex
	n       int
	window  time.Duration
	entries map[string]*limiterEntry
	calls   uint64
}

func newIPLimiter(n int, window time.Duration) *ipLimiter {
	return &ipLimiter{n: n, window: window, entries: map[string]*limiterEntry{}}
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
	e := &limiterEntry{lim: rate.NewLimiter(rate.Every(l.window/time.Duration(l.n)), l.n), seen: now}
	l.entries[ip] = e
	return e.lim, true
}

func (l *ipLimiter) allow(ip string, now time.Time) bool {
	lim, ok := l.reserve(ip, now)
	if !ok {
		return true
	}
	return lim.AllowN(now, 1)
}

// RateLimit rate-limits requests per client IP to perMinute per minute,
// with trustedHops trusted reverse proxies in front of this service.
func RateLimit(perMinute, trustedHops int) func(http.Handler) http.Handler {
	return RateLimitPer(perMinute, time.Minute, trustedHops)
}

// RateLimitExcept rate-limits per client IP to perMinute per minute exactly
// as RateLimit does, except for requests where exempt reports true: those are
// passed straight to next, neither counted against the bucket nor rejected by
// it. A nil exempt exempts nothing. The caller owns the predicate, and with it
// the reason a route may skip the limiter.
func RateLimitExcept(perMinute, trustedHops int, exempt func(*http.Request) bool) func(http.Handler) http.Handler {
	limit := RateLimit(perMinute, trustedHops)
	return func(next http.Handler) http.Handler {
		limited := limit(next)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if exempt != nil && exempt(r) {
				next.ServeHTTP(w, r)
				return
			}
			limited.ServeHTTP(w, r)
		})
	}
}

// deviceTokenPrefix marks a device bearer. It mirrors auth.DeviceTokenPrefix;
// httpx cannot import auth (auth depends on httpx).
const deviceTokenPrefix = "fsd_"

// deadTokens remembers device tokens the auth layer has answered 401, keyed by
// the hex SHA-256 of the token. An entry lives limiterIdleTimeout from its
// last mark, and the set is capped at limiterMaxEntries.
var deadTokens = struct {
	mu sync.Mutex
	m  map[string]time.Time
}{m: map[string]time.Time{}}

func sha256Sum(token string) []byte {
	sum := sha256.Sum256([]byte(token))
	return sum[:]
}

// MarkDeadToken records that the device token whose SHA-256 is hash was
// refused with a 401, so RateLimitByCredential cuts it to the dead budget.
func MarkDeadToken(hash []byte) {
	key := hex.EncodeToString(hash)
	now := Now()
	deadTokens.mu.Lock()
	defer deadTokens.mu.Unlock()
	if _, ok := deadTokens.m[key]; !ok && len(deadTokens.m) >= limiterMaxEntries {
		for k, t := range deadTokens.m {
			if now.Sub(t) > limiterIdleTimeout {
				delete(deadTokens.m, k)
			}
		}
		if len(deadTokens.m) >= limiterMaxEntries {
			return
		}
	}
	deadTokens.m[key] = now
}

func isDeadToken(key string, now time.Time) bool {
	deadTokens.mu.Lock()
	defer deadTokens.mu.Unlock()
	t, ok := deadTokens.m[key]
	if ok && now.Sub(t) > limiterIdleTimeout {
		delete(deadTokens.m, key)
		return false
	}
	return ok
}

// deviceKey returns the hex SHA-256 of the request's device bearer, if it
// carries one (the fsd_ prefix, malformed or not).
func deviceKey(r *http.Request) (string, bool) {
	h := r.Header.Get("Authorization")
	if !strings.HasPrefix(h, "Bearer ") {
		return "", false
	}
	token := strings.TrimSpace(strings.TrimPrefix(h, "Bearer "))
	if !strings.HasPrefix(token, deviceTokenPrefix) {
		return "", false
	}
	return hex.EncodeToString(sha256Sum(token)), true
}

// RateLimitByCredential budgets a request by the kind of credential it
// carries. A request with a device bearer is counted in a bucket keyed by the
// token's hash (devicePerMinute), never in the IP bucket, so one companion
// cannot spend the budget of the browser on the same address. A token the
// auth layer has refused (MarkDeadToken) drops to deadPerMinute. Everything
// else is counted per client IP (ipPerMinute). Requests where except reports
// true skip the limiter entirely.
func RateLimitByCredential(ipPerMinute, devicePerMinute, deadPerMinute, trustedHops int, except func(*http.Request) bool) func(http.Handler) http.Handler {
	ipL := newIPLimiter(ipPerMinute, time.Minute)
	devL := newIPLimiter(devicePerMinute, time.Minute)
	deadL := newIPLimiter(deadPerMinute, time.Minute)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if except != nil && except(r) {
				next.ServeHTTP(w, r)
				return
			}
			now := Now()
			var ok bool
			if key, isDev := deviceKey(r); isDev {
				if isDeadToken(key, now) {
					ok = deadL.allow(key, now)
				} else {
					ok = devL.allow(key, now)
				}
			} else {
				ok = ipL.allow(clientIP(r, trustedHops), now)
			}
			if !ok {
				WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitPer rate-limits each client IP to n requests per window, with a
// burst of n: a fresh IP may spend the whole allowance at once, and the
// bucket then refills at n per window. Used with an hour-long window for
// POST /v1/builds, which the contract caps at 20 saves per IP per hour.
// Now is the limiter's clock. Tests replace it with a frozen time so a bucket
// spent in one loop cannot refill between the last allowed request and the
// first refused one on a slow runner.
var Now = time.Now

func RateLimitPer(n int, window time.Duration, trustedHops int) func(http.Handler) http.Handler {
	l := newIPLimiter(n, window)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := clientIP(r, trustedHops)
			if !l.allow(ip, Now()) {
				WriteError(w, r, http.StatusTooManyRequests, "rate_limited", "too many requests", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimitByKey rate-limits requests to n per window, with a burst of n,
// keyed by whatever keyOf derives from the request rather than by client
// IP - the same token-bucket-per-string machinery RateLimitPer uses,
// generalized for routes that must budget per account rather than per
// address (a compromised session dispatching many premium runs, or many
// uploads, from a rotating set of addresses must still hit one budget).
// keyOf reports ok=false to exempt a request from this limiter entirely
// - an unsigned request on a route that also requires a session, say,
// which is refused downstream anyway and should not spend or contend for
// a shared "no key" bucket.
func RateLimitByKey(n int, window time.Duration, keyOf func(*http.Request) (key string, ok bool)) func(http.Handler) http.Handler {
	return RateLimitWith(NewKeyLimiter(n, window), keyOf)
}

// KeyLimiter is a token bucket per string key: n allowed per window, burst n.
// RateLimitWith puts one behind a route; a caller outside HTTP (the Discord bot's
// sign-up) shares the same budget by calling Allow on the same instance.
type KeyLimiter struct{ l *ipLimiter }

// NewKeyLimiter builds a limiter allowing n per window per key.
func NewKeyLimiter(n int, window time.Duration) *KeyLimiter {
	return &KeyLimiter{l: newIPLimiter(n, window)}
}

// Allow spends one token of key's bucket, false when it is empty.
func (k *KeyLimiter) Allow(key string) bool { return k.l.allow(key, Now()) }

// RateLimitWith is RateLimitByKey over a caller-owned limiter, so the budget can be shared
// with code that is not an HTTP route.
func RateLimitWith(l *KeyLimiter, keyOf func(*http.Request) (key string, ok bool)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key, ok := keyOf(r)
			if !ok {
				next.ServeHTTP(w, r)
				return
			}
			if !l.Allow(key) {
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
// anything they like in the leftmost entries - including sending their own
// entire X-Forwarded-For header line ahead of the one the trusted proxy
// adds, since Go's net/http preserves repeated header lines in order and
// only Values (not Get) sees all of them. So every X-Forwarded-For line is
// joined into one list before counting from the right. With trustedHops
// trusted proxies, the real client is the trustedHops-th entry counted from
// the right (trustedHops=1 means the rightmost entry). If the combined list
// has fewer entries than trustedHops - including when the header is absent
// - it is untrustworthy and RemoteAddr (the last hop we ourselves observed)
// is used instead. trustedHops=0 ignores X-Forwarded-For entirely.
func clientIP(r *http.Request, trustedHops int) string {
	if trustedHops > 0 {
		if xff := strings.Join(r.Header.Values("X-Forwarded-For"), ","); xff != "" {
			hops := strings.Split(xff, ",")
			if len(hops) >= trustedHops {
				return strings.TrimSpace(hops[len(hops)-trustedHops])
			}
		}
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// CORSMethods and CORSHeaders are what the site's islands need from
// this service: the report page reads and patches reports, the upload
// page PUTs and POSTs, the account page DELETEs a device, and every
// state-changing call carries the CSRF header. Credentials are allowed
// because the session cookie is on .foreversixty.gg and the site is a
// different origin from the API.
const (
	CORSMethods = "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	CORSHeaders = "Content-Type, X-CSRF-Token, Authorization"
)

func CORS(allowedOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Origin") == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
				w.Header().Set("Access-Control-Allow-Methods", CORSMethods)
				w.Header().Set("Access-Control-Allow-Headers", CORSHeaders)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
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

// ClientIP is the address RateLimitPer keys by, for a handler that budgets some outcomes of a
// request (a miss, say) rather than every request.
func ClientIP(r *http.Request, trustedHops int) string { return clientIP(r, trustedHops) }
