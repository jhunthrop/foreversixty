// api/internal/addon/import_sim.go
package addon

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// ImportSimmer queues a character's import run when its gear changed (design/specs/
// 2026-11-07-home-b.md, "Auto-sim on import").
type ImportSimmer interface {
	RequestFromImport(ctx context.Context, userID int64, characterKey string) (string, error)
}

// importSimTimeout bounds the detached import-sim requests a sync starts.
const importSimTimeout = 20 * time.Second

// simImports asks for each synced character's import run, detached from the sync: the
// companion's post never waits on a job start, and a failure is logged only.
func (s *Service) simImports(ctx context.Context, userID int64, exports []Export) {
	if s.ImportSims == nil || len(exports) == 0 {
		return
	}
	keys := make([]string, len(exports))
	for i, e := range exports {
		keys[i] = character.Key(e.Region, character.RulesetFromRealm(e.Ruleset, ""), e.Name)
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), importSimTimeout)
		defer cancel()
		for _, key := range keys {
			if _, err := s.ImportSims.RequestFromImport(ctx, userID, key); err != nil {
				s.logger().Warn("addon", "op", "import_sim", "character_key", key, "err", err)
			}
		}
	}()
}
