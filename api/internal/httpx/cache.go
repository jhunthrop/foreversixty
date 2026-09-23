// Package httpx: cache.go holds the two functions every GET route in
// this API uses to set its Cache-Control header (spec §2.2). Nothing
// else in this codebase should call w.Header().Set("Cache-Control", …)
// directly for a JSON-envelope route - the two functions here are the
// one place the four classes' exact values live.
package httpx

import (
	"fmt"
	"net/http"
	"time"
)

// CachePublic marks a read cacheable by a shared cache and the
// browser alike. staleWhileRevalidate of zero omits that directive
// entirely, which is the Reference class's shape (spec §2.2's table):
// those routes keep today's plain max-age with no stale-while-revalidate
// window.
func CachePublic(w http.ResponseWriter, maxAge, staleWhileRevalidate time.Duration) {
	if staleWhileRevalidate <= 0 {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int(maxAge.Seconds())))
		return
	}
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d, stale-while-revalidate=%d",
		int(maxAge.Seconds()), int(staleWhileRevalidate.Seconds())))
}

// CachePrivate marks a read as never stored by a shared cache: a
// browser may keep it only for the current page, and Vary keeps a
// shared cache that ignores Cache-Control anyway (a misconfigured
// corporate proxy, a browser extension) from ever keying one caller's
// answer for another (spec §2.1, §2.2).
func CachePrivate(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "private, no-cache")
	w.Header().Set("Vary", "Cookie, Authorization")
}
