# Lane A3: mobile home and config API

Branch `wt/mobile-home-api`. No migration. Not pushed.

## What landed

- **`api/internal/upgrades`.** A shared Go package that holds the worn-vs-BiS comparison. It contains:
  - the slot frame order and slot words;
  - the level band and preset;
  - the per-slot verdict (`Judge`, which is the site's `upgrades.ts` rule);
  - the stat-weight scorer (`ScoreItem`, which is `score.ts`);
  - the bot's upgrade order (`Order`);
  - the band talent compare (`UnmatchedPoints`, which is `talent-delta.ts`);
  - the signed figure and source line;
  - the `Gain not modelled yet` and `not sim-checked` words.

  `Compare` is the site's `upgradesFor`, ported to Go. `bis.Band` now also carries its band, preset, role, talents, talent points and weights.
- **The Discord `/char` panel reads from the shared package.** The slot list, band choice, verdict, ordering, signed figure and not-modelled words now come from `upgrades`. The panel's own policy stays in `discordbot`:
  - no stat-weight scorer;
  - the 0.05 shown-gain floor;
  - the unmodelled words;
  - the Forever Score points fallback.

  The output is unchanged. Every discordbot test passed untouched with a database, including `TestGearPanelMatchesTheMockPayloads`, along with rating, rankings and guilds.
- **The parity test.** One real-data fixture character (`api/internal/upgrades/testdata/parity/input.json`, build 1.60.1.70658, BM hunter, Alliance, 60 raid) reaches every branch. The fixture has:
  - alternative gaps;
  - a stat-weight loss;
  - an unknown worn item;
  - a weapon with no listed alternative;
  - empty slots with and without a figure;
  - ties.

  `web/src/lib/home/upgrades-parity.test.ts` runs the site's `upgradesFor` and pins `site.json`. The Go test pins the same file.
- **`GET /v1/characters/{region}/{ruleset}/{name}/home`** (`api/internal/app`).
  - It uses the character page's visibility (`rankings.Service.CharacterCard`: anonymised or unknown is a 404) and its cache (`rankings.CacheLive`, now exported).
  - It reads the export (`app.Store`) for the talents and the sync time.
  - It reads the BiS band and item tables through `guilds.Service`, the same reads the bot makes.
- **`GET /v1/app/config`**. The regions and minimum versions come from config. The active build comes from `trees.Data.Active()` through `ActiveBuild()`. The response is cached five minutes.
- **The flag.** `APP_SIGN_IN` is not on main, so this lane adds **`APP_API`** (`config.AppConfig`, off by default) along with:
  - `APP_WEB_CHECKOUT_REGIONS` (`US`, or `none`);
  - `APP_MIN_VERSION_IOS` (`1.0.0`);
  - `APP_MIN_VERSION_ANDROID` (`1.0.0`).

  `server.Deps.App` is mounted only when the flag is set. A router test checks that both routes are a 404 while the flag is off.
- **The contract.** `docs/contracts/app-v1-home-config.md` (`app-v1.md` is not on main) plus the fixtures under `docs/contracts/fixtures/app/`:
  - `characters-home.json`
  - `characters-home-empty.json`
  - `characters-home-no-gear.json`
  - `app-config.json`

  `api/internal/app/contract_test.go` checks each route's shape against its fixture: keys and types at every level, with arrays compared element by element. `-update` regenerates the fixtures.

## Tests

The following ran with `go test -race`, `go vet` and `gofmt`:

- bis
- upgrades
- discordbot (with a database)
- rankings
- rating and guilds (with a database, after the Discord move)
- app (including the store, with a database)
- config
- server
- cmd/api

On the web side, vitest ran the parity test together with `upgrades.test.ts`.

## Concerns for the coordinator

1. **The spec and the dispatch differ on shapes.** Spec §14 C1 names `/v1/me/characters/…/answer`, which is owner-only. C3 names `{plans_link:{ios:["USA"],android:["US"]}, plans_url, min_version}`. I built what the dispatch asked for: the public `/home` path and `{web_checkout_regions, min_app_version, active_build}`. Lane C or D should settle the names before the apps bind to them.
2. **Discord and the site still disagree, by design here.** The `/char` panel shows only the ranker's measured gaps. It drops a worn item the pick does not list. It shows no stat-weight figure and no negative gain, and it falls back to Forever Score points for tanks. The site and the app use the full `upgrades.ts` rule. The shared package holds one verdict, and the bot calls it without a scorer to keep its output unchanged. Making the bot match the site would change Discord's output, which this dispatch forbade. That is an owner call.
3. **Site quirks carried for parity.**
   - A worn item that the weights score above the pick is listed as an "upgrade" with a negative gain, and the gain is counted in the total. The app reads `−0.8 DPS`; the site prints `+-0.8 DPS`.
   - A pick whose effect is not modelled keeps its figure in the site's total, while its row reads `Gain not modelled yet`.
4. **New copy.** `Not in our item data yet` came from the dispatch; the site says `Unknown item`. It needs a copy decision.
5. **Ordering.** The rows follow the bot's order (empty slots first), per spec F2.1, rather than the site's gain-descending order. The parity test compares the sets of rows, not their order.
6. **The band and spec come from the bot's rules.** The band uses the bot's choice: the raid preset at 60, bare below. The spec is the gear panel's `gear_spec`, which falls back to the class's first spec when the export has no talents. The site's "no spec" can therefore be rarer here than on the site.
7. **Not yet done.** The routes are not in `openapi.yaml`, since they are dark. For the web test I linked `web/node_modules` from the main checkout (read-only, gitignored), and removed the link afterwards.
