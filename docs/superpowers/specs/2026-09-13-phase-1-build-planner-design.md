# Forever Sixty — Phase 1 design (build planner)

Date: 2026-09-13. Status: approved in conversation; sections 3 onward decided by the assistant at the user's request.

## Goal

Ship a shareable build planner for World of Warcraft: Forever before the beta opens on Sept 17, 2026, on Classic Era talent data labeled as a placeholder, and swap to the Forever trees the day the beta client exports. A build is a class, a race, the order talent points are spent in, and optionally an item per slot. Every saved build gets a short link on foreversixty.gg with a preview card.

Phase 1 is the site's first tool. The research says every independent site that broke through in this ecosystem did it with one narrow, free, shareable tool, and that Wowhead will own the phrase "talent calculator" but not the builds themselves.

## Decisions made with the user

| Question | Decision |
|---|---|
| Primary product | A shareable build planner. Pages are generated around it, not the other way round. |
| Where a shared build lives | Stored by the Go API under a short id. |
| What a build contains | Class, race, talents with point order, and gear per slot. Talents and order ship first; gear is a second release inside Phase 1. |
| Editing | Builds are immutable. Editing a shared build forks it into a new id. |
| Data timing | Build on Era data now, swap to the Forever build when it exports. |
| Addon | Phase 2, unchanged. |
| Generated pages | One races-and-classes reference page. No page per class or per combination. |
| Gear depth | Slot picker with stat totals and set bonuses. No item pages. |
| Launch target | Planner live on Era data before Sept 17. |
| Rendering of `/b/:id` | Approach A: the API server-renders the build page and its preview card; the site stays static. |

## Scope

### In

1. **Planner page** at `/planner`, with `?class=<slug>` to preselect a class and `?race=<slug>` to preselect a race. Three trees, point counter, level readout, tooltips from spell text, point order strip, share button, reset, fork from a shared build.
2. **Gear panel** on the same page: one item per slot from the datamined item table, summed stats, armor, and set bonuses. Ships second; the planner is complete without it.
3. **Shared build pages** at `/b/:id`, server-rendered by the API with Open Graph tags and a preview card at `/b/:id/card.png`.
4. **Races and classes page** at `/classes`: the nine classes, the nine races including Skyborne, the legal combinations, what changed in Forever per row with a source pill, and a link into the planner per class and per new combination.
5. **Data pipeline additions**: per-class talent files with resolved spell text, per-class item files, sets, and the Forever race and class facts. Deterministic, keyed by client build id, checked in.
6. **API additions**: builds table, save and fetch endpoints, the build page and card renderers, validation against the tree data, rate limits.
7. **Site routing**: `/b/*` proxied to the API through a Worker fetch handler in front of the static assets, so links stay on foreversixty.gg.
8. **Homepage**: the Build planner tool card flips from "Sept 17" to live; the community panel gains an email subscribe box wired to the existing subscribe endpoint.

### Out

Item, spell, quest pages and a browsable database (Phase 2). Realm picker, addon, uploader (Phase 2). Accounts, "your builds", claiming builds (Phase 3). Gear stat weights, simulation, best-in-slot lists. Build comments, votes, or a public build browser. Light mode.

## Architecture

```
foreversixty.gg/planner        static page + Planner island; loads /data/<build>/talents/<class>.json
foreversixty.gg/classes        static page from classes.json + races.json
foreversixty.gg/b/<id>   ──▶   Worker fetch handler ──▶ api.foreversixty.gg/b/<id>   (HTML + island, OG tags)
foreversixty.gg/b/<id>/card.png ─────────────────────▶ api.foreversixty.gg/b/<id>/card.png (satori PNG)
Planner island ──POST /v1/builds──▶ API ──▶ Postgres builds table
data/ pipeline ──▶ data/builds/<id>/{talents/<class>.json, items/<class>.json, sets.json, classes.json, races.json}
```

- The site remains static. The only new client JavaScript is the Planner island, hydrated with `client:load` on `/planner` and on the API-rendered build page, because the planner is the page's purpose.
- The API reads tree data from a directory named by `TREE_DATA_DIR` (default `/data`), populated in the Docker image from `data/builds/`. The Docker build context becomes the repository root with `-f api/Dockerfile`; the CI workflow changes accordingly.
- Short ids are the first eight characters of a base32 SHA-256 over the canonical JSON of `{class, race, tree_version, order, gear}`. Saving the same build twice returns the same id. Title is not part of the hash, so a retitled build is the same build; the first title wins.

## Data model

### Pipeline outputs (per client build)

- `talents/<class-slug>.json`: `{ build, class_id, trees: [ { id, name, background, talents: [ { id, name, icon, max_rank, tier, column, prereq_talent_id, prereq_rank, ranks: [ { spell_id, description } ] } ] } ] }`. Description text comes from the spell table with `$` variables resolved where the normalizer can, otherwise left as the raw token; the planner never invents numbers.
- `items/<class-slug>.json`: items the class can equip: `{ id, name, icon, slot, quality, required_level, armor, stats: { [stat]: value }, set_id, source? }`. Emitted only when the item table normalizes cleanly for that build.
- `sets.json`: `{ id, name, item_ids, bonuses: [ { pieces, description } ] }`.
- `classes.json` and `races.json` gain `combos: [[race_id, class_id], ...]` on the classes file and a `forever_changes: [ { text, sources } ]` array per row, hand-maintained under `data/curated/` and merged by the pipeline, because Blizzard states these facts in posts, not in client tables.
- `manifest.json` lists every emitted file with its hash; the site build fails if a referenced file is missing.

### Build record (API, Postgres)

```
builds(
  id            text primary key,          -- 8 chars base32
  class_id      smallint not null,
  race_id       smallint not null,
  tree_version  text not null,             -- client build id, e.g. 1.15.9.69722
  point_order   smallint[] not null,       -- talent id per point spent, in order
  gear          jsonb not null default '{}',  -- { "head": 12345, ... }
  title         text,                      -- up to 60 chars, plain text
  created_at    timestamptz not null default now(),
  views         bigint not null default 0
)
```

Validation before insert, against the tree data for `tree_version`: class and race exist and the pair is legal; every talent id belongs to the class; each point's prerequisite is met by the points before it; each point's tier is unlocked by points already in that tree (five per tier); no talent exceeds its max rank; total points at most 51; gear keys are known slots and each item id exists, is equippable by the class, and fits the slot. Violations return 400 with per-field messages. `tree_version` must be one the API has data for, else 400.

### Derived, never stored

Final tree (counts of each id), level at each point (9 plus the index), points per tree, stat totals, set bonuses active. The island and the card renderer compute these from the record and the tree data.

## API

| Method and path | Purpose |
|---|---|
| `POST /v1/builds` | Save. Body is the build record fields. 201 with `{ id, url }`, or 200 with the existing id if the hash already exists. Rate limit 20 per IP per hour, body limit 8 KB. |
| `GET /v1/builds/{id}` | The record as JSON, cache-control public for a day. 404 in the envelope if unknown. |
| `GET /b/{id}` | HTML. Title `"<title or class/race> · <a/b/c> · Forever Sixty"`, description with the tree split and level, Open Graph and Twitter card tags, canonical `https://foreversixty.gg/b/<id>`, the site's header and footer markup, and the Planner island mounted with the record inlined. Increments `views` asynchronously. 404 renders a plain page with a link to the planner. |
| `GET /b/{id}/card.png` | 1200 by 630 PNG rendered with satori and resvg: class color band, title, race and class, the three-number split, level reached, the site wordmark. Cache-control public for a week; Cloudflare caches it. |

The build page and card are rendered by Go from templates and the same tree data the validator uses. Header and footer markup are copied into a Go template once, with a test that compares them to the site's rendered output so drift fails CI.

## Planner island

- **Layout.** Desktop: the three trees side by side, each a grid of tiers by columns, with the tree's point count and name above it; a summary bar with class, race, level, and the split; the point order strip beneath as a horizontal list of icons with level labels. Phone: one tree at a time with a three-tab switcher, the summary bar pinned at the top, the order strip collapsible. Every talent cell is at least 44 px.
- **Interaction.** Click adds a point, right-click or long-press removes the last point in that talent, keyboard: arrow keys move, Enter adds, Backspace removes. Adding is refused with a short inline reason when the tier is locked or a prerequisite is missing. Removing is refused when later points depend on the one being removed, with the same inline reason. Hover or focus shows the tooltip with the current rank's text and the next rank's text.
- **State.** The build lives in one Svelte 5 rune store; the URL is not updated while editing. Reset clears with a confirm step. Share saves through the API and shows the link with a copy button and the preview card. A build opened from `/b/:id` is read-only until the visitor presses Fork, which copies it into the editable store.
- **Gear panel.** A slot grid in the game's layout; picking a slot opens a searchable list filtered to that slot and the class, sorted by required level then name, with rarity colors; totals update live. Sets show their active bonus count.
- **Copy voice.** Reference style, no exclamation marks, no calls to action. The placeholder label on Era data reads "Classic Era trees shown until the beta client exports; Forever revamped talents replace them then."
- **Performance.** The island's script stays under 60 KB gzipped. Tree data for the selected class is fetched once and cached by the browser under its build-id path. `/planner` has its own Lighthouse assertion: performance at least 0.90, accessibility and SEO at least 0.95. Content pages keep the Phase 0 budgets.

## Site routing

`web/src/worker.ts` becomes the Worker entry with a fetch handler that proxies `/b/*` to `API_BASE_URL` with the original path, forwards `Accept` and `Accept-Language`, strips cookies, and passes through cache headers. `wrangler.jsonc` sets `main` to it and `assets.run_worker_first` to `["/b/*"]`, so every other path is still served as static assets with no Worker execution.

## Error handling

- Island cannot load tree data: render the summary bar with "Talent data did not load" and a retry, never a blank page.
- Save fails: the build stays in the store; the share panel shows the API's message and a retry. Rate limited: "Too many saves from this connection; try again in an hour."
- `/b/:id` when the API is down: the Worker returns a static fallback page from assets with a link to `/planner`, status 503, no cache.
- Card render fails: the API returns a static fallback card PNG shipped in the image, status 200, short cache, and logs the id.
- Build validation: every rule has a specific message naming the talent or slot.

## Testing

- **data/**: normalizer unit tests with fixture rows for the per-class split, prerequisite resolution, spell text substitution, item slot mapping, and set bonuses; golden files for one class; a test that the curated Forever facts merge and that every fact has a source.
- **api/**: table-driven validator tests covering each rule with both a passing and a failing build; handler tests for save, dedupe, fetch, 404, rate limit; a render test that the build page contains the OG tags and the inlined record; a card test that produces a PNG of the right dimensions; the header and footer parity test.
- **web/**: unit tests on the pure planner functions (can add, can remove, level for index, split, stat totals); Container tests for the classes page; Playwright: build a tree by clicking, hit a locked tier and see the reason, share and follow the link, fork from a shared build, phone layout has no horizontal overflow, keyboard-only point spending; Lighthouse on `/planner` and `/classes`.
- **worker**: a unit test with the Workers test runner that `/b/x` proxies and `/about` does not.

## Timeline

| By | Milestone |
|---|---|
| Sept 14 | Pipeline emits per-class talents with spell text from the Era build; API builds table, validator, save and fetch; Planner island spends points on Era trees locally. |
| Sept 15 | Share works end to end; `/b/:id` renders with OG tags and card; Worker proxy live; classes page live. |
| Sept 16 | Lighthouse and e2e green; planner announced on Discord as Era-data placeholder. |
| Sept 17 to 18 | Beta client exports; pipeline run on the Forever build; normalizer fixes as needed; data swap; Forever facts curated. |
| Sept 24 | Gear panel live on Forever item data. |
| Oct 10 | Phase 1 closes; Phase 2 (realm picker, routes, addon) begins. |

## Open questions

1. Whether the beta client keeps talents in the same DB2 tables. If not, the normalizer is the only code that changes; if the trees are not in the client at all, the fallback is hand-curated JSON under `data/curated/` with the same shape.
2. Whether Skyborne's racials and class list are in the client on day one. The classes page marks unknown rows as such rather than guessing.

## Success criteria

- A build made on foreversixty.gg/planner can be shared as a short link that unfurls with a card on Discord and Twitter.
- The planner runs on the real Forever trees within two days of the beta client exporting.
- `/planner` scores at least 0.90 performance and 0.95 accessibility on mobile Lighthouse; content pages keep their Phase 0 budgets.
- No build that violates the tree rules can be saved.
