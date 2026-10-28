// Package coach builds the log coach's evidence pack (docs/superpowers/specs/
// 2026-10-27-log-coach-design.md §7): eight rule findings (§6) read from one fight's
// summary, each with its basis, comparison, facts and evidence anchors, plus the
// strengths, links and "Couldn't check" lines. It is pure Go: no database, no network,
// no model call. The pack is the model's only input about the fight and the validator's
// only allowed source.
//
// The package reads the rating engine's windows (logs/engine/rating) and the API
// rating's adherence (api/internal/rating). Neither imports coach (Ruling 1).
package coach
