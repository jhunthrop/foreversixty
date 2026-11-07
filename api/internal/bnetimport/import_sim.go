// api/internal/bnetimport/import_sim.go
package bnetimport

import (
	"context"
	"time"
)

// ImportSimmer queues a character's import run when its gear changed (sims.Service
// satisfies it; design/specs/2026-11-07-home-b.md, "Auto-sim on import"). It dedupes per
// gear itself, so it is told after every enrichment.
type ImportSimmer interface {
	RequestFromImport(ctx context.Context, userID int64, characterKey string) (string, error)
}

// importSimTimeout bounds one import-sim request.
const importSimTimeout = 20 * time.Second

// simImport asks for the character's import run after its enrichment committed. The sign-in
// import does it detached, so a slow job start never holds the import; the nightly refresh
// waits for it, since a job process ends when its loop does. A failure is logged only.
func (s *Service) simImport(ctx context.Context, ref characterRef, detached bool) {
	if s.ImportSims == nil {
		return
	}
	run := func() {
		ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), importSimTimeout)
		defer cancel()
		if _, err := s.ImportSims.RequestFromImport(ctx, ref.UserID, ref.Key); err != nil {
			s.logger().Warn("bnetimport", "op", "import_sim", "key", ref.Key, "err", err)
		}
	}
	if detached {
		go run()
		return
	}
	run()
}
