package reports

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/card"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
)

const (
	// cardMaxAge is the contract's five-minute cache for a report card.
	// A report is still filling in while people are sharing it.
	cardMaxAge = 300
	// fallbackCardMaxAge is the short cache for the stand-in card, so a
	// transient failure is not cached at all long.
	fallbackCardMaxAge = 60
)

// The two Cache-Control scopes a card is served at. See cacheScope.
const (
	sharedCache  = "public"
	privateCache = "private"
)

// card draws the 1200x630 image an unfurled report link shows. A
// report nobody may read gets the site's stand-in card rather than a
// 404: the unfurl should not say whether the link is real.
//
// Every response here varies on the session cookie, because whether a
// card is the report's or the stand-in depends on who asked. The
// stand-in is safe to share; a real card for a private or guild report
// is not, and is marked private below.
func (s *Service) card(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Vary", "Cookie")
	rep, err := s.Store.Get(r.Context(), r.PathValue("id"))
	switch {
	case errors.Is(err, ErrNotFound):
		writeCard(w, http.StatusNotFound, sharedCache, fallbackCardMaxAge, card.Fallback())
		return
	case err != nil:
		s.logger().Error("reports", "id", httpx.RequestIDFrom(r.Context()), "op", "card", "err", err)
		writeCard(w, http.StatusInternalServerError, sharedCache, fallbackCardMaxAge, card.Fallback())
		return
	}
	if !s.mayView(r, rep) {
		writeCard(w, http.StatusNotFound, sharedCache, fallbackCardMaxAge, card.Fallback())
		return
	}
	scope := cacheScope(rep)
	fights, err := s.Store.Fights(r.Context(), rep.ID)
	if err != nil {
		s.logger().Error("reports", "id", httpx.RequestIDFrom(r.Context()), "op", "card", "err", err)
		writeCard(w, http.StatusOK, scope, fallbackCardMaxAge, card.Fallback())
		return
	}
	img, err := card.RenderReport(cardInput(rep, fights))
	if err != nil {
		s.logger().Error("reports", "id", httpx.RequestIDFrom(r.Context()), "op", "card",
			"report", rep.ID, "err", err)
		writeCard(w, http.StatusOK, scope, fallbackCardMaxAge, card.Fallback())
		return
	}
	writeCard(w, http.StatusOK, scope, cardMaxAge, img)
}

// cacheScope is the Cache-Control scope a report's own card may be
// held at. A card for a private or guild report was drawn because
// mayView passed for this one caller; api.foreversixty.gg sits behind
// Cloudflare, which caches .png by extension, so marking it public
// would let an owner's own fetch fill a shared entry that is then
// served to strangers.
func cacheScope(rep Report) string {
	if rep.Visibility == Public || rep.Visibility == Unlisted {
		return sharedCache
	}
	return privateCache
}

// cardInput projects a report onto what the card draws: the bosses in
// order, and the span the night ran over.
func cardInput(rep Report, fights []FightEntry) card.ReportInput {
	in := card.ReportInput{Title: rep.Title, Zone: rep.Zone}
	if in.Title == "" && rep.Zone != "" {
		in.Title = rep.Zone
	}
	var first, last time.Time
	for _, f := range fights {
		if f.Kind == string(fight.Encounter) {
			in.Bosses = append(in.Bosses, card.Boss{Name: f.Name, Kill: f.Kill})
		}
		if first.IsZero() || f.Start.Before(first) {
			first = f.Start
		}
		if f.End.After(last) {
			last = f.End
		}
	}
	if !first.IsZero() && last.After(first) {
		in.DurationMS = last.Sub(first).Milliseconds()
	}
	return in
}

// writeCard sends a PNG, or a plain error when even the stand-in could
// not be produced, which only happens if the embedded fonts are broken.
func writeCard(w http.ResponseWriter, status int, scope string, maxAge int, body []byte) {
	if len(body) == 0 {
		http.Error(w, "card unavailable", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", fmt.Sprintf("%s, max-age=%d", scope, maxAge))
	w.WriteHeader(status)
	_, _ = w.Write(body)
}
