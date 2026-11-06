# Mock parity gate

Tenet 15: a page region lands only when the parity spec matches the mock. The committed mock is the source of truth; a build that grows a block the mock does not have, or moves a region, fails CI.

## Regions

The mock library (`design/mocks/ws_lib.py` `REGION`, emitted by `ws_lib` and `ws_prims`) and the built primitives (`web/src/lib/ui/regions.ts`, emitted by the components below) use the same `data-region` names. Add a region in both places and here, never as a literal in a page.

| Region | Built by | What it holds |
|---|---|---|
| `head` | `PageHeader` | The whole page head |
| `head-title` | `PageHeader` | Crest, title, meta line |
| `head-meta` | `PageHeader` | The one meta line under the title |
| `head-answer` | `PageHeader` (its `answer` slot) | BiS: one figure with its unit, the trust line, the links. Nothing else |
| `controls` | `ControlBar` | The control bar |
| `summary` | the page (BiS: `BisSpecPage`) | The summary line and the view toggle |
| `notes` | the page (BiS) | One-line standing notes over the Plate (source rule, hit cap, weapon) |
| `plate` | `Plate` | The page's one primary object |
| `slots` | `Paperdoll` | The slot cells |
| `inspector` / `inspector-inline` | `Inspector` | The slot detail at the side (from 1760) / in place |
| `section-<slug>` | `Section` | A section; the slug is its title (`section-what-to-stack`) |
| `footer-note` | the page (BiS) | The closing one-liner (the guide line) |

## What is compared

`web/tests/e2e/support/parity.ts`, per region present on either side:

- **Present on both.** A region on one side only fails.
- **Inventory.** The region's accessibility tree as role chains, names and text ignored (`listbox>option` x17, `paragraph>link`, `heading[1]`), plus every visible `details`, `summary` and `[aria-expanded]` (a select's own popup excepted). Any count that differs fails, so a new disclosure, note, list or link fails whatever it says.
- **Box.** x and y (y from the top of `<main>`, so the site header is not every region's error) within **8 px**; width and height within **8 px or 5%**, whichever is larger. 8 px absorbs sub-pixel and hinting drift (Linux CI against macOS) and a padding nudge, never a text line (18 to 24 px). 5% of the BiS plate is about 30 px: one row wrapping to a third line passes, a new row (58 px) or a second bar line (46 px) fails.
- **Deterministic.** Both sides at 1440 x 1000 use the site's own font files (the mock's Google Fonts request is answered with them), motion off, animations frozen.

Each run writes `web/test-results/parity/<page>-<state>.png`: mock | build at the same scale, every region outlined and named. CI uploads them as `parity-side-by-side`.

## The BiS states

`design/mocks/web/bis-parity.json`: `dps-60` (marksmanship hunter, raid), `healer-60` (holy priest), `tank-60` (protection warrior), `low-20` (beast mastery hunter, 20 to 29, Where to go). Mocks `web-bis-<state>.html`, rev 4 of `design/mocks/ws_bis.py` (spec 5 and the 5.13 build decisions).

The data on both sides is the same: `node design/mocks/bis_facts.mjs http://localhost:4321` reads names, numbers, sources and group membership from a running build into `design/mocks/data/bis-facts.json`, and `python3 design/mocks/ws_bis.py` renders the mocks from it. The facts carry no structure, so a block the build grows never reaches the mock. After a data refresh that moves the numbers, rerun both and commit the mocks with the change.

## The home states

`design/mocks/web/home-parity.json`: `out` (signed out), `34` (a level 34 leveler, no guild), `60` (a level 60 raider with a raid planned), `new` (signed in, no character), `loading` (the session hinted, `/v1/me` in flight), `error` (the upgrades failed). Mocks `web-home-<state>.html` from `design/mocks/ws_home.py` (spec `design/specs/2026-11-06-home.md`). The signed-in data is committed in `web/tests/e2e/fixtures/home/` (`python3 design/mocks/home_fixtures.py` writes it); `tests/e2e/support/home-page.ts` routes the same files, so the parity job needs no `data/builds`.

## When it fails

Read the side-by-side. If the build grew something, remove it or take a mock change to review (`ux-designer`, `wow-player`): regenerate the mock, then the gate. Never widen the tolerance to pass. A known difference another lane is fixing is listed in the spec's `EXPECTED_DRIFT` with a date and the lane; the regions below it are still compared, and the test fails once the drift is gone so the entry is deleted.

## CI

`web.yml` job `parity` (needs `build`, runs on every push to `main` and `land/**` and on pull requests, launch mode included) serves the fixture dist and runs `tests/e2e/parity-*.spec.ts`. `tools/land.sh` waits for every job of the branch's runs, so a red `parity` stops a landing. The four-shard `e2e` job also runs the specs when launch mode is off.
