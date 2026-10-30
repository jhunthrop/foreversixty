// Package app serves the routes only the mobile apps call: a character's home answer
// (GET /v1/characters/{region}/{ruleset}/{name}/home) and the apps' remote config
// (GET /v1/app/config). They are mounted only when APP_API is on (config.AppConfig), so they
// land dark. The answer is computed here, from api/internal/upgrades, so neither app
// re-derives a pick, a gain or a score.
package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rankings"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// Characters reads the public character page (rankings.Service.CharacterCard): found is false
// for an unknown region, ruleset or character and for an anonymised one.
type Characters interface {
	CharacterCard(ctx context.Context, region, ruleset, name string) (rankings.Character, bool, error)
}

// Data is the build data the answer reads: the BiS files, the item tables and the races
// (guilds.Service, the same reads the Discord /char panel makes).
type Data interface {
	ActiveBuild() (string, bool)
	BisBuild() string
	BisDataDir() string
	BuildItem(version, classSlug string, itemID int) (trees.Item, bool)
	RaceFaction(raceName string) string
}

// Exports reads a character's stored addon export (Store).
type Exports interface {
	Export(ctx context.Context, key string) (Export, bool, error)
}

// Export is a character's addon_exports row: the FS1 code, where it came from ("addon",
// "blizzard") and when it was true in the game.
type Export struct {
	Code       string
	Source     string
	CapturedAt time.Time
}

// Settings are what GET /v1/app/config answers beside the active build.
type Settings struct {
	WebCheckoutRegions               []string
	MinVersionIOS, MinVersionAndroid string
}

// Service serves the app routes.
type Service struct {
	Characters Characters
	Data       Data
	Exports    Exports
	Settings   Settings
	Log        *slog.Logger
}

// Mount registers the app routes.
func Mount(mux *http.ServeMux, s *Service) {
	mux.HandleFunc("GET /v1/characters/{region}/{ruleset}/{name}/home", s.home)
	mux.HandleFunc("GET /v1/app/config", s.config)
}

func (s *Service) logger() *slog.Logger {
	if s.Log != nil {
		return s.Log
	}
	return slog.Default()
}
