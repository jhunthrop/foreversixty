# Sim-accuracy round 7: the CI lane (2026-11-01)

Owner ruling: a CI job that builds sim.wasm the way the deploy does and runs the real-engine browser specs, failing (not skipping) when the wasm is missing. This lane did no per-spec audit and changes no DPS.

## The four real-engine specs

Confirmed four, all listed in `web/package.json`'s `test:e2e:real-engine` script (desktop project, port 4341):

1. `web/tests/e2e/sim-real-engine.spec.ts`: a real wasm run on /sim returns a positive DPS.
2. `web/tests/e2e/sim-gear-real-engine.spec.ts`: Top Gear plans, runs and ranks two candidates.
3. `web/tests/e2e/sim-weights-real-engine.spec.ts`: a browser-lane weights run finishes with a table.
4. `web/tests/e2e/sim-talents-real-engine.spec.ts`: a pasted talent build joins the ranked table.

Each calls `test.skip` when `public/_sim/<PUBLIC_SIM_ARTIFACT>/sim.wasm.gz` is absent or the module does not register its exports (`web/tests/e2e/support/real-engine.ts`). No CI job built the wasm before `deploy`, so all four skipped everywhere, and the `e2e` job's fixture dist carries the fake engine.

## Findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **No CI job ran the four real-engine specs.** `build-real` builds the real-data site without `PUBLIC_SIM_ENGINE=wasm` and without a wasm; `e2e` serves the fixture dist; only `deploy` builds the wasm, and it runs no browser test. Round 6 found three of the four stale this way (page.md P3). **Fix:** two jobs in `.github/workflows/web.yml`. `real-engine-scope` answers whether to run: always on push and dispatch, and on a pull request only when the diff touches `sim/`, `logs/`, `Makefile`, the active build's `simdb.bin`/`enchants.json`/`spellranks.json`, `web/src/data/active-build.json`, `web/src/lib/sim/`, the four specs or their support file, `web/playwright.config.ts`, `web/package(-lock).json`, `.github/actions/pin-engine/` or `web.yml` (the engine pin is `sim/enginever/version.go`, inside `sim/`). Not on the nightly schedule. `real-engine` (needs `build`, 30-minute timeout, off the deploy path as build mode keeps it) builds the wasm as `deploy` does: `setup-go` from `sim/go.mod` (go1.25.11), `.github/actions/pin-engine`, `make simdb`, `make publish-wasm` (which runs `make artifacts` and its 4.5 MB gzipped budget and writes `sim.wasm.gz`). Then the real-data build with `PUBLIC_SIM_ENGINE=wasm` on the dead API port, and `npm run test:e2e:real-engine` against it with `E2E_SKIP_BUILD=1`. | `.github/workflows/web.yml` jobs `real-engine-scope`, `real-engine` | CONFIRMED (CI gap) | S cd90cbc22 |
| C2 | **Failing instead of skipping.** Two guards. A step fails if `artifacts/ARTIFACT_ID` does not belong to the pin or `sim.wasm.gz`/`sim.js` are not in `web/public/_sim/<id>/`. After the tests, Playwright's JSON report is read and the step fails when any test skipped or none passed. Playwright itself exits 0 when every test skips, so the report check is what turns a missing wasm into a red job. | Local negative run with `PUBLIC_SIM_ARTIFACT=b57765fb9-000000000000`: Playwright exit 0, "passed 0, failed 0, flaky 0, skipped 4", check exit 1. Presence step on a missing dir: exit 1. | CONFIRMED | S cd90cbc22 |
| C3 | **Caching.** The published pair (`web/public/_sim/` plus `artifacts/ARTIFACT_ID`) is cached under `sim-wasm-v1-<pin>-<16 hex>`, the hex a sha256 of git's ids for the `sim` and `logs` trees, the `Makefile` blob and the active build's three embedded blobs. The key moves only when a compiled-in byte does (the toolchain is `sim/go.mod`'s `go` line, inside the sim tree). A hit skips Go entirely. On a miss, the Go build and module caches are a dedicated `go-wasm-<os>-<hash of go.mod, go.sum, version.go>` cache with a prefix restore key. `setup-go`'s own cache is off: its key is `sim/go.sum` alone, shared with sim.yml's native jobs, so it would restore objects with no js/wasm build in them. Playwright browsers use the same versioned cache as `e2e`. | the job's steps | Aligned | - |
| C4 | **sim.wasm.gz against its budget.** At the pin b57765fb9: raw `sim.wasm` 24,878,724 bytes, published `sim.wasm.gz` 4,470,152 bytes; `make artifacts` reports "4.26 MB (budget 4.50)". The budget is 4.5 MiB = 4,718,592 bytes, so it is at **94.7%, 248,440 bytes (243 KiB) of headroom**. Round 6 read 4,464,151 bytes at 8baffe5fc, so this round's engine added about 6 KB. Not raised (owner decision, round 6 W5). The new job prints the size to its step summary on every run. | local `make publish-wasm` at the pin with go1.25.11; `wc -c` | OPEN (owner) | - |

## Local proof

All in `.claude/worktrees/r7-ci`, with a go.work replacing the engine by a worktree detached at b57765fb9 (CI resolves the same sha through the module proxy).

- `actionlint .github/workflows/web.yml`: clean (shellcheck included).
- Cache key step: `sim-wasm-v1-b57765fb9-0b2d0f9a9d9abeb9` (`shasum -a 256` stands in for the runner's `sha256sum`).
- Scope regex on sample paths: `sim/enginever/version.go`, the gear spec, `support/real-engine.ts` and `simdb.bin` run; `web/src/pages/index.astro`, `sim-gear.spec.ts` and `items.json` do not.
- `make simdb && make publish-wasm`: published `b57765fb9-88f62db2be46/{sim.wasm.gz,sim.js}` in 12 s warm.
- Presence step: id belongs to the pin, both files there.
- `FOREVER_DATA=real PUBLIC_SIM_ENGINE=wasm PUBLIC_API_BASE_URL=http://127.0.0.1:9 npm run build`: exit 0 in 98 s; postbuild reads `check-sim-engine-artifact: dist/_sim/b57765fb9-88f62db2be46/ carries sim.wasm.gz and sim.js`.
- `CI=1 E2E_SKIP_BUILD=1 npm run test:e2e:real-engine -- --reporter=list,json`: **4 passed** (11.8 s), report check "passed 4, failed 0, flaky 0, skipped 0", exit 0.
- No server left running (port 4341 free afterwards).

What only GitHub can prove: the cold-cache timings on a runner, the cache save and restore, and the `pull_request` base-sha fetch in `real-engine-scope`. The integrator should watch the first run after the push.

## Still open

- C4: the gzipped transfer budget is at 94.7%.
- `real-engine` does not gate `deploy` (build mode: deploy needs only `build`). Add it to deploy's `needs` when gating returns.
- A pull request that changes only the /sim page's components (outside `web/src/lib/sim/`) does not run the job; the next main push does.
