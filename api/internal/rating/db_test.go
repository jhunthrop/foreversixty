// api/internal/rating/db_test.go
package rating

import (
	"os"
	"testing"
)

// testURL is this package's own copy of db.testURL's shape: db.testURL is unexported, so
// every DB-backed package in this repo (compare api/internal/db/db_test.go) declares its
// own. TEST_DATABASE_URL is set by this lane's throwaway-database setup (see the plan's
// Task 1, Step 5) and must point at rating_api_test, not the shared forever_test other
// lanes use, per this lane's dispatch.
func testURL(t *testing.T) string {
	t.Helper()
	u := os.Getenv("TEST_DATABASE_URL")
	if u == "" {
		t.Skip("TEST_DATABASE_URL not set; see docs/superpowers/plans/2026-09-21-rating-api.md Task 1 Step 5")
	}
	return u
}
