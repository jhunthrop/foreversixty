# GET /v1/guilds/{id}/reports — implementation report

Spec: `design/specs/2026-10-04-logs-landing.md` §4.C.1 (finding 3: "No guild-scoped report
list despite `me.guilds` already existing").

## What shipped

A new route, `GET /v1/guilds/{id}/reports`, registered in `reports.Mount`
(`api/internal/reports/handler.go`) alongside the package's other two list routes (`mine`,
`recent`), since it needs the same `*reports.Store` and `Accounts` (for `GuildRank`) those
already use.

- **Shape:** reuses `MinePage` verbatim — `{ rows: Summary[], total, page, per_page }`, the
  exact body `GET /v1/reports?mine=1` already returns and the frontend's `MyReportPage`
  already types. Each row is a `Summary` (`id, title, zone, status, visibility, created_at,
  fight_count, kill_count`). Page size is `MinePerPage` (100), same as `mine`; `?page=` works
  the same way. This is what the design spec's own §4.C.1 text says ("the same `MyReportPage`
  shape `listMyReports` already returns") — it is *not* cursor-paginated and carries no
  `guild_name` field, unlike `GET /v1/reports/recent`'s `RecentPage`/`RecentSummary`. The task
  brief that kicked this off described the target shape as `MyReportPage` but then listed
  `guild_name` and `next_cursor` among its fields; those two actually belong to `RecentSummary`/
  `RecentPage`, not `MinePage`/`Summary` — I went with the spec document and the real
  `MyReportPage` TypeScript type (`web/src/lib/account/api.ts`), both of which agree with each
  other and not with that aside. Flagging this now in case the intent really was the cursor
  shape; it would be a different, not-yet-built response type.
- **Rows:** a guild's own `guild`-visibility reports plus its `public` ones — the exact set
  `mayView`'s guild branch already grants a verified member on any one report of that guild.
  `private` and `unlisted` guild reports are excluded even for members, matching `mayView`'s own
  rule (only the owner/moderator sees those, membership alone does not).
- **Authorization:** `auth.RequireSession` (401 for anyone not signed in via session, same wrap
  `mine` uses) followed by the same `Accounts.GuildRank(ctx, guildID, userID)` check `mayView`
  already runs for a guild report — any verified rank counts, matching "a member... sees that
  guild's `guild`-visibility reports." No standing (or an unknown guild id — the two are
  indistinguishable on purpose) answers 404, never 403, mirroring every other report route's
  "whether it exists is itself private" rule. A moderator actor bypasses the membership check,
  matching `mayView`'s own moderator bypass for a single report.
- **Query builder reuse:** `Store.OwnedBy` was refactored to share a new
  `reportSummaryColumns` constant and `scanSummaries` helper with a new `Store.ForGuild` method
  (`api/internal/reports/store.go`) — same `Summary` projection and scan on both, differing only
  in the `where` clause (`owner_id = $1` vs. `guild_id = $1 and visibility in (guild, public)`).
  No second list implementation; the duplication that existed only inside `OwnedBy` is gone.
- **OpenAPI:** documented at `/v1/guilds/{id}/reports` in `api/openapi.yaml`, next to the other
  `/v1/guilds/{region}/{ruleset}/{name}` route, and added to `requiredPaths` in
  `api/internal/server/openapi_test.go` so a future removal is caught.

## Tests added

`api/internal/reports/guild_reports_test.go`, table-driven against the real harness/Postgres:

- `TestGuildReportsMemberSeesGuildAndPublicRowsNewestFirst` — a verified member sees exactly the
  guild's `guild` and `public` rows (not another guild's, not its own `private`/`unlisted`
  rows), newest first, with correct `total`/`page`/`per_page`.
- `TestGuildReportsNonMemberGets404` — a signed-in stranger gets 404.
- `TestGuildReportsAnonymousGets401` — no session at all gets 401.
- `TestGuildReportsEmptyGuildReturnsEmptyRows` — a member of a guild with nothing logged gets
  `rows: []`, not `null`, with `total: 0`.
- `TestGuildReportsRejectsAnInvalidGuildID` — a non-numeric `{id}` answers 404 rather than 500.

## Gate results

- `cd api && go vet ./...` — **exit 0**, clean.
- `cd api && go test ./... -count=1` — **exit 1** as literally invoked, but the only failures
  are `db: migrate up: Dirty database version 23` cascades, a pre-existing race from running
  every package's migration against one shared Postgres test instance concurrently (default
  `-p` matches `GOMAXPROCS`). This is why `.github/workflows/api.yml` runs CI as
  `go test -race -cover -p 1 ./...`. Re-run with `-p 1 -count=1`: **exit 0, every package
  passes**, including `internal/reports` (13.6s) and `internal/server` (0.5s, the OpenAPI
  path/schema check). Not caused by this change — confirmed by restoring the dirty flag and
  re-running clean. Left the test database's `schema_migrations` table un-dirtied afterward.

## Files touched

- `api/internal/reports/handler.go` — route registration + `guildReports` handler.
- `api/internal/reports/store.go` — `reportSummaryColumns`, `scanSummaries`, `ForGuild`;
  `OwnedBy` refactored onto the same two.
- `api/internal/reports/guild_reports_test.go` — new, table-driven tests above.
- `api/openapi.yaml` — new path.
- `api/internal/server/openapi_test.go` — new path added to `requiredPaths`.
