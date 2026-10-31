# App contract v1: character home and app config

Two routes only the iOS and Android apps call. They fold into `docs/contracts/app-v1.md` when lane C's file lands. Both are **dark**: they are mounted only when the API runs with `APP_API=true` (`config.AppConfig`, off by default). The fixtures are in `docs/contracts/fixtures/app/`. `api/internal/app/contract_test.go` checks each route's real response against its fixture's shape. A fixture pins the shape (keys and JSON types), never counts. `go test ./internal/app -run Contract -update` rewrites the fixtures from the handlers.

Both answer the API's usual envelope `{ ok, data, error, request_id }` and carry a weak `ETag` on `data`.

## `GET /v1/characters/{region}/{ruleset}/{name}/home`

The home answer (app spec §5.2 F2, §14 C1). It is computed in Go by `api/internal/upgrades`, the package the Discord `/char` panel reads. The apps show the rows in the server's order and never re-rank them.

- **Who may see it.** The rules of the character page, `GET /v1/characters/{region}/{ruleset}/{name}`, since the answer is read from it. A character that page answers 404 for (an unknown region, ruleset or name, or an anonymised account) is a 404 `not_found` here. The worn gear, talents and sync time are already public through that page and `…/sim-input`.
- **Caching.** The same as the character page: `public, max-age=30, stale-while-revalidate=300`.
- **The gain rule.** The site's `web/src/lib/home/upgrades.ts`, pinned to it by a parity test on one real-data character (`api/internal/upgrades/testdata/parity`, read by Go and by `web/src/lib/home/upgrades-parity.test.ts`). The rule:
  - The pick worn, or a listed alternative with a 0 delta, is already best in slot.
  - A listed alternative gains `−dps_delta`.
  - A weapon with no listed alternative has no figure (`not sim-checked`).
  - Otherwise the gain is `(pick.score − stat-weight score of the worn item) × reference_dps_per_point`. An empty slot scores 0. A worn item the item data lacks has no figure.
- **The band.** The character's level band. At 60 it is the Phase 1 raid preset (`raid`), and below 60 it is `bare`, the `/char` panel's choice. The faction is the race's, else the export's race's.

| Field | Type | Meaning |
|---|---|---|
| `character` | string | The character key. |
| `state` | `ok`, `no_gear`, `under_20`, `no_spec`, `no_list` | Checked in that order of precedence: `under_20` (the level is known and under 20), then `no_gear` (no gear shared), then `no_spec`, then `no_list` (no list published for the spec, band and faction). Any state but `ok` has `upgrades` and `already_best` empty and `total_gain` null. "Every slot matches" is `ok` with `upgrades: []`. |
| `spec` | string or null | The spec key (`hunter-beast-mastery`). Null when the data knows no spec. |
| `band`, `band_label`, `preset` | int, string, string, or null | `60`, `"60"` (or `"20 to 29"`), `"raid"` or `"bare"`. Null when no band applies. |
| `unit`, `unit_word` | string or null | `dps`, `hps` or `tank_score`, and how a figure reads it: `DPS`, `HPS` or `score`. |
| `data_build` | string | The build the item and icon ids belong to. |
| `upgrade_count` | int | `len(upgrades)`. |
| `total_gain` | number or null | The sum of every row's gain that has a figure, to one decimal. Null when it is not above 0. |
| `not_sim_checked_count` | int | Rows whose words are `not sim-checked`. |
| `not_modelled_count` | int | Rows whose words are not a figure: the note's `{n} gains not modelled yet`. |
| `synced_at`, `synced_from` | RFC 3339 and `addon` or `blizzard`, or null | When the stored export was true in the game, and its source. Null with no export. |
| `upgrades[]` | array | Ordered: empty slots in frame order, then measured gains largest first, then the rest in frame order. |
| `already_best[]` | `{ slot, slot_label, item }` | The slots already wearing the pick, in frame order. |
| `talents` | `{ state, points_differ, band_points }` | `state` is `optimized`, `differs`, `no_export` or `no_build`. `points_differ` is Σ max(band − worn, 0) over every talent cell, against the band's `band_points`. Both are null unless the state is `optimized` or `differs`. |

An **item** is `{ id, name, quality, icon, item_level }`, read from the build's table. Every field but `id` is null when the table lacks the item. `icon` is the icon id: draw `/data/{data_build}/icons/{icon}.webp` on the site. `quality` is the game's number, where 4 is Epic.

An **upgrade** is:

- `slot` and `slot_label`, for example `finger1` and `Finger 1`.
- `pick`: an item plus `source_line`, which is the source with ` · crafted` added for a crafted pick.
- `worn`: null for an empty slot. Otherwise an item plus `known` and `label`. When `known` is false, `label` is `Not in our item data yet`, and the name, quality, icon and item level are null.
- `gain`: the figure to one decimal, or null.
- `gain_words`: what the row reads. This is the signed figure (`+11.1 DPS`, `−0.8 DPS`), or `Gain not modelled yet` (no figure, or a pick whose effect the sim leaves out), or `not sim-checked`.
- `modelled`: true exactly when `gain_words` is the figure. A pick whose effect is not modelled keeps its `gain`, and the site counts it in the total, but it reads `Gain not modelled yet`.

Fixtures:

- `character.home.json`: every branch.
- `character.home.empty.json`: every slot matches.
- `character.home.no-gear.json`: no gear and no export.

## `GET /v1/app/config`

Public, with no credential. Cached `private, max-age=300` with `Vary: X-App` (the CDN does not key on `X-App`, so no shared cache keeps it), so a change reaches every app within five minutes. A 503 `unavailable` means no client build is loaded.

| Field | Type | Source |
|---|---|---|
| `web_checkout_regions` | string[] (ISO 3166 alpha-2) | Per platform, from the `X-App` header's platform (decision L): `ios` reads `APP_WEB_CHECKOUT_REGIONS_IOS` (default `US`), `android` reads `APP_WEB_CHECKOUT_REGIONS_ANDROID` (default `none`). `none` gives `[]`, and so does a missing, malformed or unknown platform. These are the storefront regions where the app may show See plans and link to web checkout. The list is never null. |
| `min_app_version` | `{ ios, android }` | `APP_MIN_VERSION_IOS` and `APP_MIN_VERSION_ANDROID`, default `1.0.0`, in the form `x.y.z`. An older build shows the blocking update screen. |
| `active_build` | string | The API's active client build (`trees.Data.Active()`), the one the static data serves. |
| `features` | `{ key: bool }` | `APP_FEATURES`, a comma list of `apps/shared/features.json` keys, each sent as `true`. Unset gives `{}`, and each app keeps its bundled defaults. |

Fixture: `app.config.json`.
