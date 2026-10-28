package coach

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// sampleItem is one sampled review as the moderator reads it.
type sampleItem struct {
	ReportID      string          `json:"report_id"`
	FightIndex    int             `json:"fight_index"`
	PlayerKey     string          `json:"player_key"`
	CharacterName string          `json:"character_name"`
	Spec          string          `json:"spec"`
	Source        string          `json:"source"`
	Validation    json.RawMessage `json:"validation"`
	Pack          json.RawMessage `json:"pack"`
	Review        json.RawMessage `json:"review"`
	SampledAt     time.Time       `json:"sampled_at"`
}

// SamplesHandler serves GET /v1/admin/coach/samples?day=YYYY-MM-DD (today UTC when absent):
// moderators only, private, no-store. Register it on the router in api/cmd/api.
func SamplesHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		actor := auth.ActorFrom(r.Context())
		switch {
		case !actor.Signed():
			httpx.WriteError(w, r, http.StatusUnauthorized, "unauthorized", "sign in first", nil)
			return
		case !actor.IsModerator():
			httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "moderators only", nil)
			return
		}
		day, ok := requestedDay(w, r)
		if !ok {
			return
		}
		items, err := daySamples(r, pool, day)
		if err != nil {
			slog.Error("coach samples", "err", err)
			httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not read the samples", nil)
			return
		}
		w.Header().Set("Cache-Control", "private, no-store")
		httpx.WriteOK(w, r, http.StatusOK, map[string]any{"day": day.Format(sampleDayFmt), "samples": items})
	}
}

func requestedDay(w http.ResponseWriter, r *http.Request) (time.Time, bool) {
	raw := r.URL.Query().Get("day")
	if raw == "" {
		return time.Now().UTC().Truncate(hoursPerDay * time.Hour), true
	}
	day, err := ParseSampleDay(raw)
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "bad_request", "day must be YYYY-MM-DD", map[string]string{"day": raw})
		return time.Time{}, false
	}
	return day, true
}

func daySamples(r *http.Request, pool *pgxpool.Pool, day time.Time) ([]sampleItem, error) {
	rows, err := pool.Query(r.Context(), `
		select report_id, fight_index, player_key, character_name, spec, source, validation, pack, review, sampled_at
		from coach_review_samples where day = $1 order by spec, report_id, fight_index, player_key`, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []sampleItem{}
	for rows.Next() {
		var s sampleItem
		if err := rows.Scan(&s.ReportID, &s.FightIndex, &s.PlayerKey, &s.CharacterName, &s.Spec, &s.Source,
			&s.Validation, &s.Pack, &s.Review, &s.SampledAt); err != nil {
			return nil, err
		}
		items = append(items, s)
	}
	return items, rows.Err()
}
