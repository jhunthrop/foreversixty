# Phase 1 interface contract

Binding for the three Phase 1 plans (data, api, web). The spec is
`docs/superpowers/specs/2026-09-13-phase-1-build-planner-design.md`; this file pins the exact
shapes the plans share so they can be executed in parallel worktrees. Values here are verbatim.

## Repository facts

- Monorepo: `web/` (Astro 7.3, Svelte 5, Tailwind 4, TypeScript strict, Vitest, Playwright,
  Lighthouse CI, ESLint flat config, Prettier), `api/` (Go 1.25, stdlib mux, pgx, golang-migrate,
  slog; module `github.com/jhunthrop/foreversixty/api`), `data/` (Python 3.12, uv, pydantic,
  pytest with `--cov-fail-under=80`), `design/DESIGN-SYSTEM.md` (binding tokens and principles).
- Node 22.12.0 at `/Users/jh/.nvm/versions/node/v22.12.0/bin`; npm from `web/`. Go from
  `/usr/local/go/bin`. Local test Postgres: `postgres://forever:forever@localhost:5434/forever_test?sslmode=disable`
  (docker compose in `api/docker-compose.test.yml`). Playwright needs `ASTRO_PREVIEW_BACKGROUND=1`.
- Commits: conventional subject, body, and these two trailer lines together in ONE final `-m`
  argument (git parses trailers only from the last paragraph, so separate `-m` flags lose the first):
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>` and
  `Claude-Session: https://claude.ai/code/session_01EkkERdonhxS2PMgXD7qZ6k`. Never bypass git hooks.
- Hosting: site on Cloudflare Workers static assets via `web/wrangler.jsonc` and `wrangler deploy`
  from `.github/workflows/web.yml`; API on Cloud Run `api` in us-east1 from
  `.github/workflows/api.yml` (Docker build, Workload Identity). API health is `GET /health`.
- Existing pipeline (`data/pipeline/`): `wago.py` fetches DB2 CSVs from wago.tools for `TABLES`,
  `normalize/{classes,races,talents,spells,items,zones,dungeons}.py`, `manifest.py`, `diff.py`,
  `__main__.py` orchestrator; output under `data/builds/<build>/` as flat JSON lists. Current
  shapes: talents `{id, tab_id, tab_name, class_id, tier, column, spell_ids[], prereq_talent_id}`;
  spells `{id, name}`; items `{id, name, quality, item_level, required_level, class_id,
  subclass_id, inventory_type}`; classes `{id, name, slug, color}`; races `{id, name, slug, faction}`.
  Current Era build id: `1.15.9.69722`, product `wow_classic_era`.

## Client build id

`TREE_VERSION` is the client build string, e.g. `1.15.9.69722`. The site reads
`web/src/data/active-build.json` → `{ "build": "1.15.9.69722" }`. The API accepts any
`tree_version` for which `TREE_DATA_DIR/<build>/` exists.

## Pipeline outputs (data plan produces; web and api consume)

Directory `data/builds/<build>/` keeps the Phase 0 files and adds:

`talents/<class-slug>.json`
```json
{
  "build": "1.15.9.69722",
  "class_id": 2,
  "class_slug": "paladin",
  "trees": [
    {
      "id": 382, "name": "Holy", "position": 0,
      "talents": [
        {
          "id": 1451, "name": "Divine Strength", "icon": "ability_golemthunderclap",
          "max_rank": 5, "tier": 0, "column": 0,
          "prereq_talent_id": null, "prereq_rank": null,
          "ranks": [ { "spell_id": 20262, "description": "Increases your Strength by 2%." } ]
        }
      ]
    }
  ]
}
```
- `tier` and `column` are 0-based. `position` orders trees left to right as in the client.
- `ranks` has exactly `max_rank` entries. `description` is the spell description with `$s1`,
  `$o1`, `$d`, `$h`-style tokens substituted when the normalizer can resolve them from
  SpellEffect and SpellDuration; unresolved tokens stay verbatim (never invented).
- `icon` is the lowercase icon file name without extension. Icon images are written to
  `data/builds/<build>/icons/<icon>.webp` (64×64) for every talent icon; the web build copies
  them to `public/data/<build>/icons/`.

`items/<class-slug>.json`
```json
{
  "build": "1.15.9.69722",
  "class_slug": "paladin",
  "items": [
    {
      "id": 12640, "name": "Lionheart Helm", "icon": "inv_helmet_29",
      "slot": "head", "quality": 4, "required_level": 60, "item_level": 65,
      "armor": 565, "stats": { "strength": 18, "crit": 2, "hit": 2 },
      "set_id": null, "unique": false
    }
  ]
}
```
- `slot` ∈ `head neck shoulder back chest wrist hands waist legs feet finger1 finger2 trinket1 trinket2 main_hand off_hand ranged`. Rings and trinkets are emitted with slot `finger` and `trinket`; the planner maps them to the two slots.
- `stats` keys ∈ `strength agility stamina intellect spirit armor crit hit spell_power healing attack_power defense dodge parry block mp5 fire_res frost_res nature_res shadow_res arcane_res`.
- `quality` is the client value (0 poor … 5 legendary). Items are filtered to `required_level` ≤ 60, class-equippable, and `inventory_type` in the slot map; the file is emitted only if the item table normalizes without error.
- Item icons go to the same `icons/` directory.

`sets.json`: `[ { "id": 209, "name": "Lawbringer Armor", "item_ids": [...], "bonuses": [ { "pieces": 2, "description": "..." } ] } ]`

`classes.json` rows become `{ id, name, slug, color, forever_changes: [ { text, sources: [ { label, url, kind } ] } ] }`
plus a top-level companion `combos.json`: `[ { "race_id": 5, "class_id": 2, "new_in_forever": true } ]`.
`races.json` rows become `{ id, name, slug, faction, placeholder, forever_changes: [...] }` (`placeholder` is a boolean on every row). Skyborne is present
with `slug: "skyborne"`, `faction: "neutral"` once the beta data has it; until then it is added
from `data/curated/races.json` with `"placeholder": true`. Curated files live in
`data/curated/{classes,races,combos}.json` and are merged by the pipeline; every `forever_changes`
entry must carry at least one source or the pipeline fails. Source `kind` is one of `blizzard | datamined | community | site`, the vocabulary of `web/src/lib/sources.ts`.

`manifest.json` lists every emitted file path with a sha256; `files` keys are paths relative to
the build directory.

## Build record (api stores; web and api render)

```json
{
  "id": "k7x2qm4a",
  "class_id": 2,
  "race_id": 5,
  "tree_version": "1.15.9.69722",
  "point_order": [1451, 1451, 1451, 1451, 1451, 1452, ...],
  "gear": { "head": 12640, "finger1": 19325 },
  "title": "Holy leveling, 10 to 30",
  "created_at": "2026-09-14T03:12:44Z",
  "views": 12
}
```
- `id`: first 8 characters of lowercase base32 (RFC 4648, no padding) of SHA-256 over the
  canonical JSON `{"class_id":2,"gear":{...sorted keys...},"point_order":[...],"race_id":5,"tree_version":"..."}`
  with sorted keys and no whitespace. Title excluded.
- `point_order` length 0..51; level at index i is `10 + i` (first point at level 10).
- `gear` keys are the 17 planner slots above; values are item ids. Absent key means empty slot.
- `title`: optional, trimmed, 1..60 characters, plain text.

Validation rules (api implements; web mirrors the same rules client-side in
`web/src/lib/planner/rules.ts` so the UI refuses the same moves the API refuses):
1. `class_id` and `race_id` exist and the pair is in `combos.json`.
2. every id in `point_order` is a talent of that class in `tree_version`.
3. for point i with talent t: points already spent in t's tree before i ≥ `5 * t.tier`.
4. if t has `prereq_talent_id`, points already spent in that talent before i ≥ `prereq_rank`.
5. points in t after i ≤ `t.max_rank`; total ≤ 51.
6. every gear key is a planner slot; the item exists in `items/<class>.json` and its `slot`
   matches (finger/trinket map to both numbered slots); `finger1`≠`finger2` and
   `trinket1`≠`trinket2` unless `unique` is false.

## API (api plan produces; web consumes)

All JSON responses use the Phase 0 envelope `{ ok, data, error, request_id }`.

| Route | Request | Response |
|---|---|---|
| `POST /v1/builds` | body `{ class_id, race_id, tree_version, point_order, gear?, title? }`, max 8 KB | 201 `{ id, url }` where `url` = `PUBLIC_BASE_URL + "/b/" + id`; 200 with the same body if the id already exists; 400 `{ error: { message, fields: { "point_order[7]": "Tier 2 of Holy needs 10 points in Holy first" } } }`; 429 when over 20 saves per IP per hour |
| `GET /v1/builds/{id}` | | 200 build record; 404 envelope |
| `GET /b/{id}` | `Accept: text/html` | 200 HTML page (see below); 404 HTML page |
| `GET /b/{id}/card.png` | | 200 `image/png` 1200×630; `Cache-Control: public, max-age=604800` |

`GET /b/{id}` HTML: `<title>` = `"{title or "{Race} {Class}"} · {a}/{b}/{c} · Forever Sixty"`; meta description =
`"{Race} {Class} build, {a}/{b}/{c} at level {10 + len(point_order) - 1}. Forever Sixty build planner."`;
`<link rel="canonical" href="https://foreversixty.gg/b/{id}">`; `og:title`, `og:description`,
`og:image` = `https://foreversixty.gg/b/{id}/card.png`, `og:url`, `twitter:card = summary_large_image`.
Body: the site header and footer markup from `api/internal/site/chrome.html` (kept in parity with
the site by a test that renders the site's `Base.astro` chrome and compares), a `<main id="main">`
containing `<div id="planner" data-build='{record JSON}' data-tree-version="…"></div>` and
`<script type="module" src="https://foreversixty.gg/planner-island.js"></script>` where the
site publishes the island bundle at that fixed path (web plan builds it as a separate entry).

Environment: `TREE_DATA_DIR` (default `/data`) contains `<build>/talents/*.json`,
`<build>/items/*.json`, `<build>/sets.json`, `<build>/classes.json`, `<build>/races.json`,
`<build>/combos.json`, `<build>/icons/*.webp` copied from `data/builds/` by the Dockerfile
(`COPY data/builds /data`; Docker context is the repo root, `-f api/Dockerfile`).

Migration `0004_builds.up.sql` creates the `builds` table exactly as in the spec.

## Site (web plan produces)

- `web/src/pages/planner.astro`: Base layout, `<Planner client:load classSlug={…} raceSlug={…} treeVersion={…} />`; class and race from the query string are read client-side (the page is static), default class `warrior`, default race the first legal one.
- `web/src/components/planner/Planner.svelte` and children; pure logic in `web/src/lib/planner/{rules,derive,store}.ts` with unit tests.
- `web/src/pages/classes.astro`: from `classes.json`, `races.json`, `combos.json`; links `/planner?class=<slug>` and `/planner?class=<slug>&race=<slug>` for `new_in_forever` combos; every `forever_changes` row rendered with `SourcePill`.
- Island bundle: `web/src/planner-island.ts` mounts `Planner` onto `#planner` reading `data-build`; built by Vite as a separate library entry to `dist/planner-island.js` (and its CSS inlined or emitted as `planner-island.css` and linked by the API page).
- Worker: `web/src/worker.ts` fetch handler; `/b/*` → `${API_BASE_URL}${pathname}${search}`; `wrangler.jsonc` gains `"main": "src/worker.ts"`, `"assets": { ..., "binding": "ASSETS", "run_worker_first": ["/b/*"] }`, and `"vars": { "API_BASE_URL": "https://api.foreversixty.gg" }`.
- Data at build time: `web/scripts/sync-data.mjs` copies `data/builds/<active>/{talents,items,icons,sets.json,classes.json,races.json,combos.json}` to `web/public/data/<build>/` before `astro build` (npm `prebuild`). The classes page imports the JSON directly.
- Homepage: `tools.json` Build planner card status becomes `{ label: "Live", kind: "site" }` with href `/planner`; community panel gains `SubscribeBox.svelte` (`client:visible`) posting to `${PUBLIC_API_BASE_URL}/v1/subscribe`.
- `web/lighthouserc.json` adds `/planner.html` and `/classes.html` URLs; `/planner` asserts performance ≥ 0.90, others unchanged.
