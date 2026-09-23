package httpx

import (
	"net/http"
	"strconv"
)

// Page is one page of a caller's own rows, the shape every "mine=1"
// list route answers with: sims, builds, and any future one. T is the
// row type, so sims and builds each get their own JSON shape without
// each declaring the envelope around it.
type Page[T any] struct {
	Rows    []T `json:"rows"`
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

// RequireMine enforces the one shape a "my own rows" list route takes:
// mine=1 is required, because there is no "everyone's rows" list and
// inventing one by omission would be a surprise. It writes the 400
// itself and reports whether the request may continue.
func RequireMine(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Query().Get("mine") != "1" {
		WriteError(w, r, http.StatusBadRequest, "invalid", "this list is mine=1 only",
			map[string]string{"mine": "1"})
		return false
	}
	return true
}

// ParsePage reads the page query parameter, defaulting to 1. It writes
// the 400 itself for anything but a positive integer and reports
// whether the request may continue.
func ParsePage(w http.ResponseWriter, r *http.Request) (int, bool) {
	page := 1
	if v := r.URL.Query().Get("page"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			WriteError(w, r, http.StatusBadRequest, "invalid", "page must be 1 or more",
				map[string]string{"page": "a page number from 1"})
			return 0, false
		}
		page = n
	}
	return page, true
}

// SetPrivateListCache marks a per-account list response as never
// cached at a shared edge: the response depends on who is asking, and
// a shared cache does not know that. This is CachePrivate under its
// list-specific name, kept as its own function so every "mine=1" call
// site reads as what it is - httpx/cache.go is still the one place
// the actual header value lives.
func SetPrivateListCache(w http.ResponseWriter) {
	CachePrivate(w)
}
