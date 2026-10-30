# Mobile lane C: contracts and fixtures

STATUS: done, not pushed.

Delivered:
- `docs/contracts/app-v1.md`: every endpoint and static file the v1 apps read or write.
- `docs/contracts/fixtures/app/*.json`: 73 fixtures, one per response, with empty and error variants.
- `api/internal/contract`: the shape comparer (field names, kinds, nullability; never values or counts), its tests, the CDN tests and the envelope test over every fixture.
- `app_contract_test.go` in `auth`, `builds`, `guilds`, `rankings`, `reports` and `sims`: each serves the real handler with test data and checks the answer and its Cache-Control against the fixture.

Run: `TEST_DATABASE_URL=... go test -race -p 1 ./internal/contract ./internal/{auth,builds,guilds,rankings,reports,sims} -run 'TestAppContract|TestCDN|TestEvery|TestFixtures|TestCompare'`. Use `-p 1`: the packages share one database and deadlock when run in parallel.

Open items for the coordinator:
1. `active-build.json` and `dates.json` are not on the CDN. `web/scripts/sync-data.mjs` does not copy them into `public/data/`. The apps need the active build id. Publish both (a two-line addition to the sync, which needs the full web suite) or serve the build from `GET /v1/app/config`.
2. `GET /v1/app/config` has no handler anywhere; its shape is the plan's (`web_checkout_regions`, `min_app_version`). A1 and A2 shapes (sign-in, logout, `/v1/me/notifications`, `/v1/me/push-tokens`) were read from the lane branches and are held only to the envelope rule until they merge; reconcile at merge.
3. Fixtures for `app.*` paths were written by hand; once A1/A2 land, add `app_contract_test.go` beside their handlers.
