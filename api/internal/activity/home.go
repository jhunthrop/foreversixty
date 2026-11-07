package activity

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
	"github.com/jhunthrop/foreversixty/api/internal/fs1"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// ImportSimmer queues a character's import run; sims.Service satisfies it. It answers
// "queued" when it started one.
type ImportSimmer interface {
	RequestFromImport(ctx context.Context, userID int64, characterKey string) (string, error)
}

// importStartTimeout bounds the one import run GET /v1/home may start for a character that
// has gear and was never simmed (one imported before the auto-sim shipped).
const importStartTimeout = 10 * time.Second

// Home answers GET /v1/home.
type Home struct {
	Store   *Store
	Imports ImportSimmer
}

// homeCharacter is what the blocks are read for.
type homeCharacter struct {
	key, faction, spec string
	userID             int64
	level              int
	hasGear            bool
}

func (s *Service) home(w http.ResponseWriter, r *http.Request) {
	key := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("character")))
	if key == "" {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "character is required", map[string]string{"character": "required"})
		return
	}
	actor := auth.ActorFrom(r.Context())
	c, ok, err := s.Home.character(r.Context(), key)
	if err != nil {
		s.fail(w, r, "home character", err)
		return
	}
	if !ok || c.userID != actor.UserID {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such character on your account", nil)
		return
	}
	answer, err := s.Home.answer(r.Context(), c, s.now())
	if err != nil {
		s.fail(w, r, "home", err)
		return
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, answer)
}

func (h *Home) character(ctx context.Context, key string) (homeCharacter, bool, error) {
	c := homeCharacter{key: key}
	var export, class, simSpec *string
	var level *int
	err := h.Store.Pool.QueryRow(ctx,
		`select coalesce(c.user_id, 0), c.level, coalesce(c.faction, 'alliance'), c.class, x.export,
		        (select cs.spec from character_sims cs where cs.character_key = c.key order by cs.created_at desc limit 1)
		 from characters c left join addon_exports x on x.character_key = c.key where c.key = $1`, key).
		Scan(&c.userID, &level, &c.faction, &class, &export, &simSpec)
	if isNoRows(err) {
		return c, false, nil
	}
	if err != nil {
		return c, false, fmt.Errorf("activity: home character %s: %w", key, err)
	}
	if level != nil {
		c.level = *level
	}
	if export != nil {
		if d, ok := fs1.Decode(*export); ok {
			c.hasGear = len(d.Gear) > 0
			c.spec, _ = character.SpecFor(d.ClassSlug, "", d.Talents.Points)
			if d.HasLevel {
				c.level = d.Level
			}
		}
	}
	if simSpec != nil {
		c.spec = *simSpec
	}
	if c.spec == "" && class != nil {
		c.spec, _ = character.SpecFor(strings.ToLower(*class), "", nil)
	}
	return c, true, nil
}

func (h *Home) answer(ctx context.Context, c homeCharacter, now time.Time) (HomeAnswer, error) {
	var out HomeAnswer
	var err error
	if out.ImportSim, err = h.importSim(ctx, c); err != nil {
		return out, err
	}
	if c.spec == "" {
		return out, nil
	}
	band := BandOf(c.level)
	if out.SpecWeek, err = h.Store.specWeek(ctx, c.spec, band, c.faction, now); err != nil {
		return out, err
	}
	if out.Movers, err = h.Store.movers(ctx, c.spec, band, now); err != nil {
		return out, err
	}
	out.GuildSims, err = h.Store.guildSims(ctx, c, band, now)
	return out, err
}

func (h *Home) importSim(ctx context.Context, c homeCharacter) (*ImportSim, error) {
	sim, found, err := h.Store.latestImportSim(ctx, c.key)
	if err != nil || found {
		return sim, err
	}
	if !c.hasGear || h.Imports == nil {
		return nil, nil
	}
	ctx, cancel := context.WithTimeout(ctx, importStartTimeout)
	defer cancel()
	outcome, err := h.Imports.RequestFromImport(ctx, c.userID, c.key)
	if err != nil {
		return nil, err
	}
	if outcome == "queued" {
		return &ImportSim{State: SimPending}, nil
	}
	return nil, nil
}
