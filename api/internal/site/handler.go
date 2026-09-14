package site

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/jhunthrop/foreversixty/api/internal/builds"
	"github.com/jhunthrop/foreversixty/api/internal/card"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// pageMaxAge is how long a shared build page may be cached. It is short
// because the page is the thing a person lands on from a shared link and
// the view counter behind it should keep moving; the card it points at is
// the expensive artifact, and that one is cached for a week.
const pageMaxAge = 300

const (
	// cardMaxAge is the contract's week-long cache for a preview card.
	cardMaxAge = 604800
	// fallbackCardMaxAge is the short cache used when the real card could
	// not be drawn, so a transient failure is not cached for a week.
	fallbackCardMaxAge = 300
)

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

// Mount registers the HTML build page and its preview card.
func Mount(mux *http.ServeMux, d Deps) {
	mux.HandleFunc("GET /b/{id}", d.page)
	mux.HandleFunc("GET /b/{id}/card.png", d.card)
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

func (d Deps) card(w http.ResponseWriter, r *http.Request) {
	b, err := d.Store.Get(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, builds.ErrNotFound):
		d.writeCard(w, r, http.StatusNotFound, fallbackCardMaxAge, card.Fallback())
		return
	case err != nil:
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "card", "err", err)
		d.writeCard(w, r, http.StatusInternalServerError, fallbackCardMaxAge, card.Fallback())
		return
	}

	desc := builds.Describe(d.Data, b)
	img, err := card.Render(card.Input{
		Title:      desc.Title,
		Race:       desc.RaceName,
		Class:      desc.ClassName,
		ClassColor: desc.ClassColor,
		Split:      desc.Split,
		Level:      desc.Level,
	})
	if err != nil {
		// The spec requires a 200 here, so an unfurl still shows something,
		// with the short cache so a transient failure is not cached for a
		// week. The id is logged.
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "card", "build", b.ID, "err", err)
		d.writeCard(w, r, http.StatusOK, fallbackCardMaxAge, card.Fallback())
		return
	}
	d.writeCard(w, r, http.StatusOK, cardMaxAge, img)
}

// writeCard sends a PNG body, or a plain 500 when even the fallback could
// not be produced (only possible if the embedded fonts are unusable).
func (d Deps) writeCard(w http.ResponseWriter, r *http.Request, status int, maxAge int, body []byte) {
	if len(body) == 0 {
		d.logger().Error("site", "id", httpx.RequestIDFrom(r.Context()), "op", "card", "err", "no fallback card available")
		http.Error(w, "card unavailable", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func (d Deps) logger() *slog.Logger {
	if d.Log != nil {
		return d.Log
	}
	return slog.Default()
}
