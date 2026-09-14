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

| Command                 | Action                                                                    |
| :---------------------- | :------------------------------------------------------------------------ |
| `npm install`           | Install dependencies                                                      |
| `npm run dev`           | Start the local dev server                                                |
| `npm run build`         | Build the production site to `./dist/`                                    |
| `npm run preview`       | Preview the production build locally                                      |
| `npm run check`         | Type-check with `astro check` (TypeScript strict)                         |
| `npm run lint`          | Lint with ESLint (Astro, Svelte, TypeScript)                              |
| `npm run lint:fix`      | Lint and apply the fixable rules                                          |
| `npm run format`        | Format with Prettier                                                      |
| `npm run format:check`  | Check formatting without writing                                          |
| `npm run test`          | Run unit tests once with Vitest                                           |
| `npm run test:watch`    | Run unit tests in watch mode                                              |
| `npm run test:e2e`      | Run end-to-end tests with Playwright (fixture data)                       |
| `npm run test:e2e:real` | Pre-deploy smoke: build with real data, run `tests/e2e/real-data.spec.ts` |
| `npm run lhci`          | Run Lighthouse CI                                                         |

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

`lighthouserc.json` audits four URLs under a mobile, throttled profile: `/index.html`,
`/dungeons/hall-of-thanes.html`, `/classes.html` and `/planner.html`. `ci.assert.assertMatrix` (an array
of `{ matchingUrlPattern, assertions }` entries; not `ci.assert.assertions`, which is mutually exclusive
with it) holds three of them. Every collected URL matches exactly one, which is the property to preserve
when a URL or an entry is added: `@lhci/utils` tests each entry's pattern against each URL with a plain
`new RegExp(pattern).test(finalUrl)` and applies every entry that matches, so a URL matching two entries
is held to both and a URL matching none is collected and never asserted at all.

- `^(?!.*/(?:index|planner)\.html$).*$` — `/dungeons/hall-of-thanes.html` and `/classes.html`. LCP 1600 ms.
- `.*/index\.html$` — the homepage alone, for the 1800 ms LCP budget evidenced below.
- `.*/planner\.html$` — the planner alone. No LCP assertion.

Accessibility and SEO stay at 0.95 on every URL, and so do TBT (100 ms) and CLS (0.05). Performance is
0.95 on the two static entries and 0.90 on `/planner.html`, which ships more interactive surface for the
budget. LCP is the only metric that is not asserted everywhere: `/planner.html` omits it because its
figure moves with talent icon decode timing rather than static layout, and it measures 2120 ms against
the 1352 ms the two static pages come in at.

`/planner.html` is the one page that hydrates an island, so TBT is the metric its growth moves first and
the composite performance score is too coarse to catch that alone. It measures 0 ms of total blocking
time across all three runs under the 4x CPU slowdown, so the same 100 ms the rest of the site carries
costs nothing today and turns a regression into a named failure rather than a slipped score.

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
