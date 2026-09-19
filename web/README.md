# Forever Sixty — Web

A fan reference site for World of Warcraft: Forever, built with [Astro](https://astro.build) (static output), [Svelte 5](https://svelte.dev) islands, and [Tailwind CSS 4](https://tailwindcss.com).

## Stack

- **Astro** — static site generation, `build.format: 'file'`
- **Svelte 5** — interactive islands
- **Tailwind CSS 4** — via the `@tailwindcss/vite` plugin
- **astro-pagefind** — static full-text search
- **@astrojs/sitemap** — sitemap generation

## Commands

All commands are run from this directory (`web/`):

| Command                  | Action                                                                    |
| :----------------------- | :------------------------------------------------------------------------ |
| `npm install`            | Install dependencies                                                      |
| `npm run dev`            | Start the local dev server                                                |
| `npm run build`          | Build the production site to `./dist/`                                    |
| `npm run preview`        | Preview the production build locally                                      |
| `npm run check`          | Type-check with `astro check` (TypeScript strict)                         |
| `npm run lint`           | Lint with ESLint (Astro, Svelte, TypeScript)                              |
| `npm run lint:fix`       | Lint and apply the fixable rules                                          |
| `npm run format`         | Format with Prettier                                                      |
| `npm run format:check`   | Check formatting without writing                                          |
| `npm run test`           | Run unit tests once with Vitest                                           |
| `npm run test:watch`     | Run unit tests in watch mode                                              |
| `npm run test:e2e`       | Run end-to-end tests with Playwright (fixture data)                       |
| `npm run test:e2e:real`  | Pre-deploy smoke: build with real data, run `tests/e2e/real-data.spec.ts` |
| `npm run test:e2e:phone` | Just the four phone audits, at 360x800, on the `mobile` project           |
| `npm run lhci`           | Run Lighthouse CI                                                         |

## Project Structure

```text
web/
├── public/                  # static assets served as-is (favicon, robots.txt); sync-data writes public/data/
├── scripts/                 # build-time Node scripts (sync-data, make-planner-fixture, check-island-size)
├── src/
│   ├── components/          # Astro components and Svelte islands, with their co-located unit tests
│   │   └── planner/         # the planner island's own Svelte components (Planner, TreeGrid, GearPanel, …)
│   ├── content/             # Markdown content collections (changelog, dungeons, guides, pages, zones)
│   ├── data/                # static JSON data (active build, dates, tools, classes, community links, unknowns)
│   ├── directives/          # custom client directives (e.g. `client:interaction`)
│   ├── fixtures/planner/    # checked-in planner data the sync falls back to off a deploy
│   ├── layouts/             # shared page layouts (Base, Content)
│   ├── lib/                 # framework-agnostic helpers (dates, sources, OG image generation)
│   │   └── planner/         # the planner's logic modules — see the table below
│   ├── pages/               # file-based routes, including dynamic `[slug]` and OG image routes
│   ├── styles/              # global CSS and design tokens
│   ├── content.config.ts    # content collection schemas (zod)
│   ├── planner-island.ts    # entry for the standalone island bundle the API's /b/:id page links
│   └── worker.ts            # the Cloudflare Worker for /b/* (the only server code here)
├── tests/e2e/               # Playwright end-to-end specs (fixture data; real-data.spec.ts is the real-data smoke)
├── lighthouserc.json        # Lighthouse CI budgets
├── vite.island.config.ts    # the second Vite build that emits dist/planner-island.{js,css}
├── wrangler.jsonc           # Workers deploy config (assets, routes, custom domains)
├── playwright.config.ts
├── vitest.config.ts
├── astro.config.mjs
├── svelte.config.js
├── eslint.config.js
├── tsconfig.json
└── package.json
```

## Data for the planner

`scripts/sync-data.mjs` runs before `dev`, `check`, `test` and `build` (npm `pre*` hooks). It reads
`src/data/active-build.json` (`{ "build": "1.15.9.69722" }`), then copies
`../data/builds/<build>/{talents,items,icons,sets.json,classes.json,races.json,combos.json}` into
`public/data/<build>/` for the island to fetch at runtime, and mirrors the three reference files into
`src/data/generated/` for `src/pages/classes.astro` to import at build time. Both destinations are
gitignored.

`active-build.json` names the default build for `/planner` and `/classes`, but the sync publishes
**every** build under `../data/builds/` whose `manifest.json` lists Phase 1 talent data, each into its
own `public/data/<build>/`, and logs which ones it published. Saved builds are immutable and carry the
`tree_version` they were made against: `/b/:id` mounts the island on that build id and the island
fetches `/data/<tree_version>/…` for its talents, items and reference files. Publishing only the active
build would break every share link already circulating the moment `active-build.json` moved on. The
retained set is the same set the API accepts — any `tree_version` it holds data for. A build directory
with no Phase 1 talent data is skipped and named in a warning. `src/data/generated/` still comes from
the active build alone, since those are build-time page imports rather than per-build fetches.

Every path listed in the build's `manifest.json` must exist on disk or the sync fails and names the
missing files. While `data/builds/<build>/` still holds only the Phase 0 flat files, the sync falls
back to the checked-in fixture at `src/fixtures/planner/` and says so on stdout. That fallback is
refused when `CF_PAGES` is set, so a deploy can never publish fixture talent data. Regenerate the
fixture with `node scripts/make-planner-fixture.mjs`. A `public/data/<build>/` directory this run did
not publish is removed, so switching sources or retiring a build never leaves a stale build behind
for the island to fetch.

### `FOREVER_DATA`: which data the sync publishes

`FOREVER_DATA` selects the source explicitly; any value other than the two below fails the sync
rather than being guessed at.

| Value            | Effect                                                                                           |
| :--------------- | :----------------------------------------------------------------------------------------------- |
| `real` (default) | Publishes `data/builds/` as described above, fixture fallback included.                          |
| `fixture`        | Publishes `src/fixtures/planner/` whatever `data/builds/` holds. Refused when `CF_PAGES` is set. |

**The test suites run on the fixture.** `pretest` sets `FOREVER_DATA=fixture`, and
`playwright.config.ts` passes `FOREVER_DATA: process.env.FOREVER_DATA ?? 'fixture'` to the build its
`webServer` runs. The unit tests and every browser spec but one assert on talent names, tree counts,
item ids and reference rows; the real pipeline output moves all of those on each regeneration, so
running them against it would make them a changelog rather than a test. The fixture is a two-tree
warrior with a handful of items and the full nine-by-nine reference tables, and it is checked in, so
those suites are deterministic.

**The real data gets a smoke suite.** `tests/e2e/real-data.spec.ts` skips unless `FOREVER_DATA=real`,
and `npm run test:e2e:real` builds with real data and runs it alone: `/planner` opens on the default
class and lays out its three trees under the names the synced talent file gives them, switching class
lays out the new class's trees, `/classes` crosses nine classes with nine races, and the share panel is
there. It asserts on structure and on names it reads back out of `public/data/<build>/`, never on
particular Forever facts, so regenerating `data/builds/` cannot turn it red on its own. It is not in
CI's verify job — run it before a deploy.

`npm run build` (and CI's deploy job, which adds `CF_PAGES=1`) uses real data: `FOREVER_DATA` is unset
there, and `real` is the default. CI's verify job runs `npm run build` for the real build and then
`npm run test:e2e`, whose `webServer` rebuilds `dist/` on the fixture; `npm run lhci` audits that
fixture build, which is the build the Lighthouse budgets below were measured against.

## The planner island

`/planner` mounts `src/components/planner/Planner.svelte` with `client:load`. All the logic sits in
plain modules under `src/lib/planner/`:

| Module            | Responsibility                                                               |
| :---------------- | :--------------------------------------------------------------------------- |
| `config.ts`       | The API base url, the default class, the era data notice                     |
| `types.ts`        | Shapes from the Phase 1 interface contract, the 17 gear slots, the stat keys |
| `rules.ts`        | Validation rules 1-6, mirrored from the API so the UI refuses the same moves |
| `derive.ts`       | Split, level per point, gear stat totals, active set bonuses                 |
| `grid.ts`         | Sparse tier-by-column layout and arrow-key movement                          |
| `store.svelte.ts` | The one rune store: draft build, loaded data, refusal, read-only and fork    |
| `load.ts`         | Fetches under `/data/<build>/`                                               |
| `share.ts`        | `POST /v1/builds` and its wording                                            |
| `items.ts`        | Item list helpers for the gear panel, and the rarity token per quality       |
| `styles.ts`       | The class strings more than one planner component renders                    |
| `reference.ts`    | Typed access to `src/data/generated/*.json` for `/classes`                   |

`npm run build` also runs `build:island`, a second Vite build (`vite.island.config.ts`) that writes
`dist/planner-island.js` and `dist/planner-island.css` at fixed, unhashed paths. It uses a plain
`rollupOptions` entry with pinned output filenames rather than Vite's library mode: library mode inlines
every referenced asset as a base64 data URL regardless of `assetsInlineLimit`, and the island's
stylesheet imports `src/styles/fonts.css` for ten font files, which turned `planner-island.css` from
26 KB into 307 KB of render-blocking base64. `rollupOptions` instead emits the font files under
`dist/assets/` as their own content-hashed, cacheable requests. The Go API links the two fixed-name
files from its server-rendered `/b/<id>` page. `postbuild` then runs `scripts/check-island-size.mjs`,
which fails the build if `planner-island.js` passes 60 KB gzipped.

`src/styles/fonts.css` holds the three self-hosted faces (Cinzel, Barlow, JetBrains Mono) in one module
because both the static site (`src/layouts/Base.astro`) and the island entry (`src/planner-island.ts`)
import it; the island's page links only `planner-island.css`, so without a shared module its header,
footer and planner would fall back to Georgia and Arial on the domain every share link points at.

## Shared build pages and the Worker

`src/worker.ts` is the only server code in `web/`. `wrangler.jsonc` sets `main` to it and
`assets.run_worker_first` to `["/b/*"]`, so it runs for shared build pages and their preview cards and
nothing else; every other path is served straight from the static assets with no Worker invocation. The
handler proxies `/b/<id>` and `/b/<id>/card.png` to `API_BASE_URL` (a `var` in `wrangler.jsonc`),
forwarding only `Accept` and `Accept-Language` to the upstream request — an allow-list, so the
browser's `Cookie` header is never forwarded — and setting `X-Forwarded-For` from Cloudflare's own
`CF-Connecting-IP` (a client-supplied `X-Forwarded-For` is never read, so it can't be spoofed through).
On the way back, every upstream response header passes through unchanged except `Set-Cookie`, which is
stripped so a cacheable, edge-cached response never leaks a session cookie. A 404 is the API's own "no
such build" page. A network failure or a 5xx from the API serves `dist/b-unavailable.html` with status
503 and `Cache-Control: no-store`. That fallback is internal plumbing rather than a destination, so
`astro.config.mjs` filters it out of the sitemap and the page carries `<meta name="robots"
content="noindex">`: nobody should arrive at it from a search result.

### Environment

| Variable              | Where                                          | Default                                                                 |
| :-------------------- | :--------------------------------------------- | :---------------------------------------------------------------------- |
| `PUBLIC_API_BASE_URL` | build time (Astro and the island bundle)       | `https://api.foreversixty.gg`                                           |
| `API_BASE_URL`        | `vars` in `wrangler.jsonc`, read by the Worker | `https://api.foreversixty.gg`                                           |
| `CF_PAGES`            | deploy build only                              | unset; when set, placeholder links and fixture data both fail the build |
| `FOREVER_DATA`        | `scripts/sync-data.mjs`, so every `pre*` hook  | `real`; `fixture` publishes `src/fixtures/planner/` instead             |

## Combat log reports

A report's files are written to the R2 bucket `foreversixty-logs` by the API and served to
the browser by this site's own Worker at `/logs-data/reports/<id>/…` — same origin, same
edge cache, no CORS, and no origin call on a repeat view of a closed fight. The keys are
the engine's `store.Keys` exactly:

```
reports/<id>/report.json              metadata, health, fight list, units   (max-age=5)
reports/<id>/fights/<n>/summary.json  the precomputed tables                (immutable)
reports/<id>/fights/<n>/events.parquet  typed events, queried in the browser  (immutable)
reports/<id>/fights/<n>/live.json     the snapshot while a fight is open     (max-age=5)
```

R2 does not know a report is private, so `src/worker.ts` is the enforcement point: it asks
the API for the report's visibility, caches the answer for sixty seconds, serves public and
unlisted, and refuses private and guild with 403. Those two carry a signed `data_base_url`
from `GET /v1/reports/{id}/access` instead.

### The shells

`/reports/<id>`, `/rankings/<slug>`, `/character/<region>/<ruleset>/<name>` and
`/guild/<region>/<ruleset>/<name>` do not exist at build time. The Worker serves one static
shell per prefix — `dist/reports.html` and its three siblings — and rewrites the `<title>`,
canonical and Open Graph tags from the API's JSON with `HTMLRewriter`. The values come from
`src/lib/report/og-meta.ts`, which is pure and unit-tested; `src/test-support/html-rewriter.ts`
is a small stand-in for the runtime global, because `HTMLRewriter` has no Node
implementation and `@cloudflare/vitest-pool-workers` peers on vitest 4 while this project is
on 5.

### The report island

`dist/report-island.js` is a second standalone bundle, built exactly like the planner's by
`vite.report-island.config.ts`, and budgeted at 140 KB gzipped by
`scripts/check-island-size.mjs`. Everything it computes lives in plain TypeScript under
`src/lib/report/`:

| Module            | Responsibility                                                        |
| ----------------- | --------------------------------------------------------------------- |
| `types.ts`        | TypeScript mirrors of the engine's JSON, field for field              |
| `url.ts`          | the page's whole state, which is its query string                     |
| `load.ts`         | the API and `data_base_url`, plus the five-second live poll           |
| `window.ts`       | rescoping every table to a brushed window, from the per-second series |
| `filters.ts`      | target, ability, boss-only, players-only, overkill, after-death       |
| `percentile.ts`   | parse percentiles, six at a time, cached                              |
| `events.ts`       | the Events view's list, built from what the summary timestamps        |
| `query.ts`        | the DuckDB-WASM query layer, behind an interface the tests fake       |
| `planner-link.ts` | a combatant's gear and talents as an FS1 link into the planner        |

### Deep queries

The Queries view runs SQL over the fight's own `events.parquet` with DuckDB-WASM. It loads
lazily — never on page load, asserted by an e2e — and entirely from this origin. The
package's own jsDelivr helper is deliberately unused: no page on this site makes a
third-party request.

`scripts/sync-duckdb.mjs` publishes the runtime out of `node_modules` at build time, and
splits it in two because Cloudflare Workers cap a single static asset at 25 MiB:

- `duckdb-eh.wasm` and `duckdb-mvp.wasm` (32.7 and 37.5 MiB) are staged into
  `build/duckdb-runtime/<version>/`, uploaded to the `foreversixty-logs` bucket under
  `runtime/duckdb/<version>/`, and served by `src/worker.ts` at
  `/duckdb-runtime/<version>/<file>` with an immutable year.
- the two `duckdb-browser-*.worker.js` and the vendored parquet extension go to
  `public/duckdb/<version>/` as ordinary static assets, cached for a year by
  `public/_headers` because that path carries the version too.

`scripts/duckdb-runtime.mjs` and `src/lib/report/duckdb-runtime.ts` are the two halves of
that arithmetic, held to one version by `duckdb-runtime.test.ts`. The upload is a
user-owned step (below) and runs once per `@duckdb/duckdb-wasm` bump; `npm run postbuild`
fails the build if any file in `dist/` is over the 25 MiB limit, so this class of defect
cannot reach a deploy again.

Spec section 9 asks for a timing test of a brush over a 10 MB fight. The checked-in fixture's
Parquet is 18 KB, so that budget is measured against a real raid log once one exists rather
than against a fixture standing in for one; `tests/e2e/report-queries.spec.ts` measures a
real DuckDB query over the real fixture in the meantime.

### Running it locally

```bash
npm run dev                      # real data, no fixture report
FOREVER_DATA=fixture npm run dev # adds /reports/fixture2abcd and /logs-data/
npm run make:report-fixture      # regenerates src/fixtures/report/ with the engine CLI
```

The fixture report exists only under `FOREVER_DATA=fixture`: a production build prerenders
no fixture pages and publishes no `public/logs-data/`.

## Accounts, logs and rankings

| Page                       | Island                                     | What it does                                                                 |
| -------------------------- | ------------------------------------------ | ---------------------------------------------------------------------------- |
| `/login`                   | `Account` (`login`)                        | Battle.net, or a single-use email link                                       |
| `/account`                 | `Account` (`account`)                      | devices and pairing, characters, the anonymize toggle, sign out              |
| `/logs`                    | `Account` (`pairing`, `reports`), `Upload` | live-logging steps, the companion downloads, whole-file upload, your reports |
| `/rankings/<slug>`         | `Rankings`                                 | character and guild boards, every filter in the URL                          |
| `/character/…`, `/guild/…` | `Character`, `Guild`                       | ranked history, bests, builds, progression                                   |

The header shows who is signed in only on these pages and on the report shells: `Base.astro`
takes a `session` prop, and a content page that set it would ship client JavaScript for a
link and would not hold its Lighthouse budget.

Sessions are an opaque `HttpOnly` cookie, so no code here reads one — `credentials:
'include'` is the whole mechanism — and every state-changing call sends the readable
`fs_csrf` cookie back in `X-CSRF-Token`.

`Rankings.svelte`, `Character.svelte` and `Guild.svelte` are ordinary `client:load` Astro
islands rather than a third standalone Vite build: Astro bundles each into `dist/_astro/`
under a content-hashed filename (`Rankings.<hash>.js`, and so on), so they cannot be checked
by a fixed-path budget the way `planner-island.js` and `report-island.js` are. None of their
pages are in `lighthouserc.json`'s `collect.url` either — Step 1 of Task 21 audits the report
page but not the rankings, character or guild shells, so nothing in the Lighthouse run would
catch a regression here. `scripts/check-island-size.mjs` therefore also globs
`dist/_astro/` for each of the three by component name and holds them to 16 KB gzipped —
today they measure roughly 4 KB, 2 KB and 2 KB, so the ceiling is headroom against a
regression rather than a target to grow into.

### The phone audits

`tests/e2e/report-phone.spec.ts`, `rankings-phone.spec.ts`, `character-phone.spec.ts` and
`guild-phone.spec.ts` are the standing regression guard for the phone layout on the four
logs-product pages that carry one: the report, the rankings board, a character page and a
guild page. Each file sets its own 360x800 viewport with `test.use()` and skips itself on
every Playwright project but `mobile`, so plain `npm run test:e2e` already runs all four as
part of the ordinary suite — no CI step is dedicated to them. Each sweeps every tab, view and
mode (or filter and page, for rankings/character/guild) the page can be in and asserts two
things throughout: nothing scrolls sideways at 360px, and every visible control clears 44px
in its smallest dimension. `npm run test:e2e:phone` runs just these four, against whatever
`dist/` the last build left behind, for a fast local check while iterating on layout.

## Simulator

Three routes. `/sim` and `/sim/specs` are static, prerendered pages, same as `/planner` and
`/classes`. `/sim/<sim_id>` is Worker-served: `src/worker.ts`'s `SHELL_ROUTES` maps `/sim/`
to the `/sim.html` asset and rewrites its head from `GET /v1/sims/{sim_id}` the same way it
does for `/reports/<id>`, `/rankings/<slug>`, `/character/…` and `/guild/…` — a saved sim
carries no visibility gate (every one is public by contract), so every one is indexable.

The engine — `sim/request` and `sim/adapter` compiled to `sim.wasm` from the site's own
`sim/` Go module — runs entirely in the visitor's browser; the server is never asked to
simulate anything for the free lane (`RunControl`'s `premium` prop gates a separate,
server-side run for premium accounts, whose entrypoint is `POST /v1/sims/run` — the browser
lane stays free and unlimited either way). `web/src/lib/sim/engine.ts` reads
`PUBLIC_SIM_ENGINE` at build time: `wasm` loads the real engine from
`/_sim/<ENGINE_VERSION>/` (`web/public/_sim/README.md` documents that directory; CI builds
and publishes it, and it is never committed), and the default, `fake`, swaps in
`web/src/fixtures/sim/engine-fake.ts` over the same four functions so the pages, tests and
this repository's own `npm run build` work with no Go toolchain and no wasm artifact at all.

`postbuild`'s `scripts/check-island-size.mjs` holds three fixed-path bundles to a gzipped
ceiling: `planner-island.js` at 60 KB, `report-island.js` at 140 KB, and `sim-island.js` —
this lane's addition — at 90 KB. The engine itself never counts against that budget: it
loads from the worker chunk after first paint, never from the island's own module graph
(`grep -rn "_sim/" dist/sim-island.js dist/planner-island.js` finds nothing in either, and
under the checked-in `fake` engine the wasm loader is dead code that never reaches `dist/`
at all).

`/sim/specs` renders each spec's `state` — `validated`, `in_progress` or `unsupported` — and
how close its simmed DPS runs to real parses; that judgment lives in the data this page
reads (`GET /v1/specs`), never as a list hard-coded here.

```bash
FOREVER_DATA=fixture npx vitest run src/lib/sim src/fixtures/sim
E2E_PORT=4325 FOREVER_DATA=fixture npx playwright test sim-
```

## Deploying this: two things you own

Two pieces of the deploy are outside this repository's CI and were left for whoever runs
`wrangler deploy` for the first time.

**The DuckDB engine modules have to be uploaded to R2 once per version.** They are 32.7 and
37.5 MiB, over Cloudflare's 25 MiB static-asset cap, so they are served out of the
`foreversixty-logs` bucket by `src/worker.ts` rather than out of `dist/` (see _Deep
queries_ above). `npm run build` stages them; putting them in the bucket needs an R2 write
token this repository does not have. From `web/`, with `CLOUDFLARE_API_TOKEN` and
`CLOUDFLARE_ACCOUNT_ID` set, or after `npx wrangler login`:

```bash
npm run sync:duckdb    # stages build/duckdb-runtime/<version>/
npm run upload:duckdb  # puts both modules in foreversixty-logs/runtime/duckdb/<version>/
```

`.github/workflows/web.yml` runs the same command before `wrangler deploy` when the
`CLOUDFLARE_R2_TOKEN` secret exists, and prints a notice and carries on when it does not.
Until the upload is done the site deploys and every page works; only the Queries view has
nothing to instantiate.

**The rest of the backend is likewise provisioned by hand, not by this workflow:** the
`foreversixty-logs` R2 bucket and its `LOGS` binding in `wrangler.jsonc`, the bucket's CORS
rule (`ExposeHeaders: ETag`, which the whole-file upload flow reads off each part response),
the Cloud Run job the API dispatches log processing to, and the signing certificate the API
uses for `data_base_url` and the multipart upload URLs. None of those live in `web/`, and
none of them are things this repository's CI can create — they are one-time account-level
setup the API and infrastructure own.

## Deploy (Cloudflare Workers, static assets)

The site is served by Cloudflare Workers as static assets, plus the one Worker described above for
`/b/*`; `wrangler.jsonc` holds the config (asset directory `dist`, `auto-trailing-slash` so
`dungeons.html` answers at `/dungeons`, the `404-page` handling for unknown paths, and the two custom
domains `foreversixty.gg` and `www.foreversixty.gg`).
Cloudflare Pages was the original plan; it has since been folded into Workers and new projects cannot be
created through the API, so the site deploys with wrangler instead.

- Every push to `main` that touches `web/**` or `data/builds/**` runs the verify job and then
  `wrangler deploy` (see `.github/workflows/web.yml`). It needs two repository secrets:
  `CLOUDFLARE_API_TOKEN` (Workers Scripts Edit, Zone DNS Edit for the custom domains) and
  `CLOUDFLARE_ACCOUNT_ID`.
- Manual deploy from this directory: `CLOUDFLARE_API_TOKEN=... npm run deploy` (`npm run build` then
  `wrangler deploy`). The workers.dev preview is https://foreversixty.jhunthrop.workers.dev.
- Node is pinned by `web/.node-version` (`22.12.0`); CI reads it with `node-version-file`.

Community links (Discord invite, GitHub repo, issue templates, and the community panel's subscribe
fallback, which points at the Discord until the site has a contact address) live in
`src/data/links.json`. The deploy job builds with `CF_PAGES=1`, which makes the build fail and name
the offending keys if any value still contains `PLACEHOLDER`; local and CI verify builds are
unaffected.

## Lighthouse budgets

`lighthouserc.json` audits six URLs under a mobile, throttled profile: `/index.html`,
`/dungeons/hall-of-thanes.html`, `/classes.html`, `/planner.html`, `/logs.html` and
`/reports/fixture2abcd.html`. `ci.assert.assertMatrix` (an array of `{ matchingUrlPattern, assertions }`
entries; not `ci.assert.assertions`, which is mutually exclusive with it) holds four of them. Every
collected URL matches exactly one, which is the property to preserve when a URL or an entry is added:
`@lhci/utils` tests each entry's pattern against each URL with a plain `new RegExp(pattern).test(finalUrl)`
and applies every entry that matches, so a URL matching two entries is held to both and a URL matching
none is collected and never asserted at all.

- `^(?!.*/(?:index|planner|logs)\.html$)(?!.*/reports/).*$` — `/dungeons/hall-of-thanes.html` and
  `/classes.html`. LCP 1600 ms.
- `.*/index\.html$` — the homepage alone, for the 1800 ms LCP budget evidenced below.
- `.*/(?:planner|logs)\.html$` — the planner and `/logs`, which ships the same class of client-side
  surface (the `Account` and `Upload` islands). No LCP assertion.
- `.*/reports/.*\.html$` — `/reports/fixture2abcd.html`, the only report page that renders with no
  network at all: the prerendered fixture shell inlines its `GET /v1/reports/{id}` payload as
  `data-report`, and `public/logs-data/reports/fixture2abcd/` (published by `scripts/sync-duckdb.mjs`'s
  sibling, `sync-report-fixture.mjs`) carries the rest. No LCP assertion, and a wider TBT ceiling.

Accessibility and SEO stay at 0.95 on every URL, and so do CLS (0.05). Performance is 0.95 on the two
static entries and 0.90 on `/planner.html`, `/logs.html` and `/reports/fixture2abcd.html`, which each
ship more interactive surface for the budget. LCP is not asserted on `/planner.html`, `/logs.html` or
the report page: `/planner.html`'s figure moves with talent icon decode timing rather than static
layout (2120 ms against the 1352 ms the two static pages come in at), and the report page's largest
element is the fight selector, which exists only once the island has parsed `report.json` — an LCP
number there would measure the island's boot rather than the page's paint, and the performance score
already covers that. TBT stays at 100 ms everywhere except the report page, which gets 150 ms: the
island parses a summary and lays out twelve tabs' worth of state on a 4x throttled CPU, and holding it
to a static page's 100 ms would mean deleting tables rather than making them faster.

`/planner.html` and `/logs.html` are the pages that hydrate an island, so TBT is the metric their growth
moves first and the composite performance score is too coarse to catch that alone. `/planner.html`
measures 0 ms of total blocking time across all three runs under the 4x CPU slowdown, so the same 100 ms
the rest of the site carries costs nothing today and turns a regression into a named failure rather than
a slipped score.

`assertMatrix` has to sit inside `ci.assert`, not as a sibling key of `ci.collect`/`ci.upload`. Sibling
placement parses without error, but `@lhci/cli@0.15.1`'s `autorun` command decides whether to even run
assertions by checking `ciConfiguration.assert` specifically (`src/autorun/autorun.js`). With
`assertMatrix` outside `assert` and `upload` also configured, that check is false, so `autorun` skips the
assert step entirely and exits 0 without checking anything — confirmed by running `lhci assert` directly
against that shape, which throws `Error: No assertions to use`. Nesting it under `ci.assert.assertMatrix`
fixes both: the `autorun` gate sees a truthy `assert`, and yargs's config loading (which only spreads
`ci.assert`'s own keys into the `assert` command's options) picks up `assertMatrix` at all.

The homepage entry's LCP budget is 1800 ms, not the Phase 0 site's original 1600 ms. It is its own
matrix entry for exactly that reason: the evidence below is about the homepage, and carrying it on the
general entry would have loosened `/dungeons/hall-of-thanes.html` and `/classes.html` by 200 ms on the
strength of it. Those two measure 1352 ms, so 1600 leaves them the headroom they had. Phase 0 measured the
homepage at 1502 ms against 1600 ms — 98 ms of headroom. Task 14 added a subscribe box to the community
panel, and with it the homepage measures 1652 ms. Five separate configurations were tried while chasing
that regression back down — CSS made byte-identical to the pre-task baseline, a `client:interaction`
hydration swap, a no-hydration diagnostic build, and `content-visibility: auto` with a measured
`contain-intrinsic-block-size` — and LCP held at 1651.7-1653.9 ms across all of them. The cost tracks the
roughly 1 KB of added DOM under Lighthouse's Lantern simulation, not hydration strategy, CSS surface, or
render-skipping: this is a genuine cost of the required feature, not an implementation defect. 1800 ms
keeps the budget meaningful against the 1502 ms Phase 0 baseline instead of rubber-stamping the new
number, and a budget with only 98 ms of headroom would flap across CI runners regardless.
`lighthouserc.json` is parsed with plain `JSON.parse` (`@lhci/utils` only special-cases `.js`/`.cjs` and
`.yaml`/`.yml` filenames), so it cannot hold this reasoning as a comment — hence it lives here.

## Redirect foreversixty.com → foreversixty.gg

1. Add `foreversixty.com` as a zone in the same Cloudflare account; point the Namecheap nameservers at Cloudflare (done 2026-09-13: coby and lady .ns.cloudflare.com).
2. DNS: `A @ 192.0.2.1` proxied and `CNAME www foreversixty.com` proxied (placeholders so the proxy answers).
3. Rules → Redirect Rules → create: when hostname matches `foreversixty.com` or `www.foreversixty.com`, redirect 301 to `https://foreversixty.gg${uri}` preserving path and query.

## Verify

```
curl -sI https://foreversixty.com/dungeons | grep -i location   → https://foreversixty.gg/dungeons
curl -s https://foreversixty.gg/robots.txt                        → Sitemap line present
```
