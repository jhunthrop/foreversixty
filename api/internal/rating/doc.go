// Package rating is the API + storage + job lane of the performance rating feature
// (docs/superpowers/specs/2026-09-21-performance-rating-design.md §8's "API + storage +
// job" row): the percentile-digest adapter that feeds logs/engine/rating.Score, the store
// that writes and reads rating_scores/rating_percentile_digests/kill_duration_digests, the
// two public read endpoints (§5), the async worker api/internal/reports/ingest.go's i.rate
// hook schedules into, and the backfill job (§4.4).
package rating
