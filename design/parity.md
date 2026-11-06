# Mock parity gate

Tenet 15: a page region lands only when the parity spec matches the mock. The mock generator is the source of truth for structure: it decides every block and where it sits. A build that grows a block the mock does not have, or moves a region, fails CI.

The specs render their mocks **at test time**, from the same build they serve (see "Test-time mocks"), so the two sides always show the same data: a data refresh (new picks, a longer trust line, new numbers) cannot fail the gate and needs no hand regen to land. The committed mocks in `design/mocks/web` stay the human reference and the side-by-side source (see "The reference mocks").

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

`design/mocks/bis_facts.mjs` reads names, numbers, sources and group membership from a running build; `design/mocks/ws_bis.py` renders the mocks from those facts. The facts carry no structure, so a block the build grows never reaches the mock.

## The simulator states

`design/mocks/web/sim-parity.json`: `landing-out`, `landing-in`, `result`, `saved`. `design/mocks/ws_sim.py` draws them from `web/src/fixtures/sim/result-full.json` and `data/builds/<active build>`, the files `web/tests/e2e/support/sim-parity.ts` serves to the fixture build.

## Test-time mocks

Before its tests each spec renders its mocks into `web/test-results/parity/mocks/<page>/` (`web/tests/e2e/support/parity-mocks.ts`) and compares against those:

- **BiS:** `node design/mocks/bis_facts.mjs <the spec's own baseURL> <out>/bis-facts.json`, then `ws_bis.py` with `FS_BIS_FACTS` and `FS_MOCK_OUT` pointing there. The URL is the spec's preview (`E2E_PORT`), never whatever holds another port; `bis_facts.mjs` without a URL stops with the exact command.
- **Simulator:** `ws_sim.py` with `FS_DATA_ROOT` at the checkout and `FS_MOCK_OUT` there.

The generators need `python3` (standard library only; Pillow is only for `ws_sim.py compare` and `ws_compose.py`). Rendering takes under a second for the simulator and a few seconds of page reads for BiS.

So a content-only change passes and a structural one fails. Proved on 2026-11-06 (land/parity-data): a trust line with one long fact added, which wraps and grows the head 18 px, passes all four BiS states (against the committed mocks it failed every region below the head, the land/bis-final2 failure); a `details`/`summary` in the head answer and an extra card above the summary fail (`[head-answer] inventory: disclosure:details: mock 0, build 1`, `[outside-regions] inventory: outside:surface: mock 0, build 1`, and every region below moved).

## The reference mocks

The committed `design/mocks/web/*.html` are what a reviewer reads and what the side-by-side compares by hand. After the specs, `node scripts/check-parity-mocks.mjs` (in `web/`) compares them with the test-time mocks (icon files aside: a checkout without an icon draws its empty socket, same box) and, when they differ, names each stale file, the first text that differs and the commands that refresh them:

```
# in web/, in its own terminal
npm run build && npm run preview -- --port 4500
# from the repo root
node design/mocks/bis_facts.mjs http://localhost:4500 && python3 design/mocks/ws_bis.py
python3 design/mocks/ws_sim.py
```

Then commit `design/mocks/web` (and `design/mocks/data/bis-facts.json`).

**It warns, it does not block.** In CI the `parity` job runs it after the specs with `continue-on-error`: a stale reference writes a warning annotation and the commands into the job summary, and the job stays green. The gate already compared against the test-time mocks, so a stale reference cannot hide a structural regression; it only makes the hand side-by-side read older numbers. Blocking on it would bring back exactly what this design removes: a nightly data refresh (bis.yml, data.yml) that cannot land until someone regenerates mocks by hand. Refresh the reference when a page's design changes, or whenever the warning shows up on `main`.

## When it fails

Read the side-by-side (`web/test-results/parity/<page>-<state>.png`, the test-time mock against the build). A data refresh never lands here: both sides read the same data. If the build grew something, remove it or take a mock change to review (`ux-designer`, `wow-player`): change the generator (`ws_bis.py`, `ws_sim.py`, `ws_prims.py`), rerun the gate, and regenerate the reference mocks with the commands above. Never widen the tolerance to pass. A known difference another lane is fixing is listed in the spec's `EXPECTED_DRIFT` with a date and the lane; the regions below it are still compared, and the test fails once the drift is gone so the entry is deleted.

## CI

`web.yml` job `parity` (needs `build`, runs on every push to `main` and `land/**` and on pull requests, launch mode included) serves the fixture dist and runs `tests/e2e/parity-*.spec.ts`, then the reference report (`scripts/check-parity-mocks.mjs`, non-blocking). Its sparse checkout carries the generators (`design/mocks/*.py`, `*.mjs`) and, from `data/builds`, only the files `web/tests/e2e/support/parity-data-paths.mjs` lists (what the simulator spec serves and `ws_sim.py` reads); the runner's own `python3` renders. `tools/land.sh` waits for every job of the branch's runs, so a red `parity` stops a landing. The four-shard `e2e` job also runs the specs when launch mode is off.

Locally, with a port no other checkout holds: `E2E_PORT=4500 npx playwright test tests/e2e/parity-*.spec.ts --project=desktop && node scripts/check-parity-mocks.mjs` in `web/`.
