# Forever Sixty — Phase 0 design (foundation)

Date: 2026-09-12. Status: draft for review.

## Goal

Ship foreversixty.gg before the World of Warcraft: Forever beta opens on Sept 17, 2026, with the live-game homepage layout in place, real content where facts exist, and the pipelines that every later phase depends on: a data-ingestion path for the beta client, a Go API, and a deploy that costs nothing at a million visits.

Phase 0 is small on purpose. It exists so that Phase 1 (build planner, from Sept 17) starts on a running site.

## Context

- Research: `research/00-synthesis.md` and the numbered files beside it.
- Look and feel: `design/DESIGN-SYSTEM.md`, mockups in `design/*.dc.html`, canvas at https://claude.ai/code/artifact/7400b706-7d75-44d0-b399-2571afb03023.
- Decisions already made: brand Forever Sixty (foreversixty.gg primary, .com redirects); Astro static site with Svelte islands; Go API; Postgres; user has beta access from Sept 17.

## Phases (for orientation; only Phase 0 is specified here)

| Phase | Window | Ships |
|---|---|---|
| 0 Foundation | now to Sept 17 | this document |
| 1 Build planner | Sept 17 to ~Oct 10 | talents, race/class combos, shareable builds; datamined content pages |
| 2 Launch tools | Oct 10 to Nov 4 | realm picker, queue tracker, new-zone routes, addon v1 + uploader |
| 3 Community and raids | Nov 4 to Dec 9 | Legacy/character tracker, guild board, raid guides |
| 4 Hardcore | winter | deathlog, deathmap |

## Scope of Phase 0

### In

1. **Repository and tooling**: monorepo with `web/` (Astro + Svelte + Tailwind, TypeScript strict), `api/` (Go), `data/` (ingestion scripts and exported game data), `design/`, `docs/`, `research/`. Lint, format, type-check, and test commands at the root. CI runs them on every push.
2. **Homepage in pre-launch state**, using the live-game layout with pre-launch contents in each slot:
   - Search box (client-side search over the site's own pages; item/quest search arrives with Phase 1 data).
   - State panel showing the key dates (Sept 13 panels, Sept 17 beta, Oct 27 names, Nov 4 launch, Dec 9 raids) with an "updated" stamp.
   - Tools grid with availability pills (Sept 17, Oct 27, Live).
   - "What changed" feed driven by a content collection; every entry has a date and a source tag.
   - "Still unknown" list in the side slot; it becomes "Your characters" in Phase 3.
   - Guides-by-class tiles linking to stub pages that state plainly what is and isn't known.
   - Community panel: Discord link, report-an-error and write-a-guide links (GitHub issues to start).
   - Footer with the Blizzard disclaimer.
3. **Content pages from research**: "Everything we know," a dungeons index with the nine names and single-source level ranges labeled as such, a zones index for the four new zones, a Skyborne page, an editions and pricing page, a roadmap page. All Markdown in an Astro content collection with a schema requiring `updated` and `sources`.
4. **Source and freshness UI**: the source-pill component, the "updated" stamp component, and a convention that every fact page lists its sources at the bottom.
5. **Data ingestion skeleton** in `data/`: scripts that fetch DB2 exports for a given build from wago.tools, store them under `data/builds/<build>/`, and emit normalized JSON for zones, dungeons, items, spells, and talents. Runs against the current Classic Era build now to prove the pipeline; runs against the Forever beta build on Sept 17. Output is checked into the repo (or an object store if it grows past a few hundred MB).
6. **Go API skeleton** in `api/`: HTTP service with health, version, and a single `POST /v1/subscribe` endpoint (email capture into Postgres, double opt-in). Structured logging, request IDs, rate limiting, migrations. Deployed so that Phase 1 endpoints have a home.
7. **Hosting and DNS**: Astro output on Cloudflare Workers static assets at foreversixty.gg (Pages was the plan; it was folded into Workers before the first deploy on 2026-09-13); .com redirects with 301; Go container on Google Cloud Run (scale to zero, one region, Postgres on Neon) at api.foreversixty.gg; Cloudflare DNS in front of both. Privacy-respecting analytics (Cloudflare Web Analytics) with no cookie banner. Decided 2026-09-12 after considering Fly.io and Vercel: Cloud Run is the natural home for a Go container and Cloudflare does not meter static bandwidth.
8. **SEO and sharing basics**: titles, descriptions, canonical URLs, sitemap, Open Graph images generated per page, `robots.txt`.
9. **Performance budget** enforced in CI with Lighthouse CI: mobile performance 95+, LCP under 1.6s on a Lighthouse mobile-throttled run (the three self-hosted typefaces cost about 0.3–0.6 s of simulated LCP and were kept by decision), zero JavaScript on content pages, islands only where declared.
10. **Legal and trust**: disclaimer footer, an About page explaining who runs the site and how facts are sourced, a Sources page, a Changelog page.

### Out (deferred to later phases)

Item, quest, and NPC database pages (Phase 1). Talent calculator (Phase 1). Realm list and queues (Phase 2). Addon and desktop uploader (Phase 2). Accounts, sign-in, character tracking (Phase 3). Guild board (Phase 3). Deathlog (Phase 4). Ads and any monetization. Light mode.

## Architecture

```
foreversixty.gg  ──▶  Cloudflare Pages (static HTML/CSS + Svelte islands)
                          │
                          │ islands fetch JSON
                          ▼
api.foreversixty.gg ──▶  Go container on Cloud Run ──▶ Postgres (Neon)

data/  (ingestion scripts, run locally or in CI)
   wago.tools DB2 exports ──▶ normalized JSON ──▶ committed ──▶ Astro build reads it
```

- **Astro** renders every page to static HTML at build time. Content pages read Markdown collections and the normalized JSON in `data/`. No Astro server runtime in production.
- **Svelte 5 islands** are used only for interactive components. In Phase 0 that is the search box. Islands are hydrated with `client:idle` or `client:visible` unless above the fold.
- **Go API** is the only server. Standard library router, `pgx` for Postgres, `golang-migrate` for migrations, `slog` for logging. OpenAPI document committed and tested against the routes. Runs as a container on Cloud Run; large uploads in later phases go direct to object storage via signed URLs rather than through the API.
- **Data pipeline** is Python or Go scripts (decide in the plan; Python has the mature DB2 tooling via `pywowlib`). Deterministic output; a build ID in every generated file.
- **Rebuilds**: a GitHub Action rebuilds and deploys the site on push to `main` and on a manual "content updated" dispatch. Data updates land as commits.

## Data model (Phase 0 only)

Postgres, one table for now:

- `subscribers(id, email, confirmed_at, created_at, unsubscribed_at, token)`; unique on email; token for double opt-in and unsubscribe.

Content collections in Astro: `pages`, `changelog`, `dungeons`, `zones`, `guides`. Shared schema fields: `title`, `updated` (date, required), `sources` (array of `{label, url, kind}` where kind is `blizzard | datamined | community | site`, required, min 1), `confidence` (`confirmed | single-source | inferred`).

## Error handling

- API returns a consistent envelope `{ok, data, error}` with request ID; validation errors are 400 with field messages; rate limit is 429; everything else 500 with the ID logged.
- Site build fails loudly if any content entry lacks `updated` or `sources`, or if a data file's build ID doesn't match the manifest.
- Islands render a static fallback when the API is unreachable (search falls back to a link to the sitemap).

## Testing

- `web/`: Vitest for Svelte components and utilities; Playwright smoke test (homepage loads, search island hydrates, no console errors); Lighthouse CI budget.
- `api/`: Go table-driven tests for handlers and the subscribe flow against a test Postgres (dockerized); coverage target 80%.
- `data/`: unit tests on the normalizers with fixture DB2 rows; a golden-file test per output type.
- Every pull request runs all three.

## Security and privacy

No secrets in the repo; Google Secret Manager and Cloudflare hold them. The subscribe endpoint is rate-limited, validates email format, and never returns whether an address exists. Analytics has no cookies and no personal data. No third-party scripts on the page.

## Timeline

| By | Milestone |
|---|---|
| Sept 13 (Sat) | Repo, tooling, CI, hosting wired; homepage skeleton live at foreversixty.gg behind a "coming soon" robots rule |
| Sept 14 | Content pages from research; source/freshness components; Deep Dive and Hardcore panel notes folded in |
| Sept 15 | Data pipeline proven on a Classic Era build; Go API deployed with subscribe |
| Sept 16 | Lighthouse budget green; OG images; site opened to crawlers; announcement on Discord only |
| Sept 17 | Beta opens: run the pipeline on the Forever build; test whether addons load; Phase 1 begins |

## Open questions

1. Python or Go for the data pipeline. Recommendation: Python for Phase 0 because the DB2 tooling exists; revisit if it becomes the slowest part.
2. Email provider for double opt-in mail (Resend, Postmark, or SES). Recommendation: Resend for its free tier and simple API.
3. Whether to open the site to crawlers on Sept 16 or wait for Phase 1 content. Recommendation: open on Sept 16; the research pages are already better sourced than most coverage.

## Success criteria

- foreversixty.gg serves the homepage and research pages with mobile Lighthouse 95+ before Sept 17.
- The data pipeline produces normalized JSON from a real build with tests passing.
- The Go API is deployed, health-checked, and accepting subscriptions.
- Everything the homepage shows carries a date and a source.
