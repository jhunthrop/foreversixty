package site

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// pageMaxAge is how long a shared build page may be cached. It is short
// because the page is the thing a person lands on from a shared link and
// the view counter behind it should keep moving; the card it points at is
// the expensive artifact, and that one is cached for a week.
const pageMaxAge = 300

// BuildGetter is the part of builds.Store the page needs.
type BuildGetter interface {
	Get(ctx context.Context, id string) (builds.Build, error)
}

// ViewRecorder counts a page view without blocking the response.
type ViewRecorder interface{ Record(id string) }

type Deps struct {
	Store         BuildGetter
	Data          *trees.Data
	PublicBaseURL string
	Views         ViewRecorder
	Log           *slog.Logger
}

// Mount registers the HTML build page.
func Mount(mux *http.ServeMux, d Deps) {
	mux.HandleFunc("GET /b/{id}", d.page)
}

func (d Deps) page(w http.ResponseWriter, r *http.Request) {
	b, err := d.Store.Get(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, builds.ErrNotFound):
		d.writeMessage(w, r, http.StatusNotFound, "Build not found",
			"That link does not point at a saved build. It may have been mistyped.")
		return
	case err != nil:
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "page", "err", err)
		d.writeMessage(w, r, http.StatusInternalServerError, "Build unavailable",
			"That build could not be loaded just now.")
		return
	}

	page, err := d.buildPageHTML(b)
	if err != nil {
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "page", "build", b.ID, "err", err)
		d.writeMessage(w, r, http.StatusInternalServerError, "Build unavailable",
			"That build could not be rendered just now.")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", pageMaxAge))
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, page)

	if d.Views != nil {
		d.Views.Record(b.ID)
	}
}

func (d Deps) writeMessage(w http.ResponseWriter, r *http.Request, status int, heading, body string) {
	page, err := d.messageHTML(heading, body)
	if err != nil {
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "message", "err", err)
		http.Error(w, heading, status)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_, _ = io.WriteString(w, page)
}

func (d Deps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}
