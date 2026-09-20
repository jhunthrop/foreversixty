// web/tests/e2e/sim-gear-real-engine.spec.ts
// The one end-to-end proof that the real wasm's simPlan/simRank survive the full browser
// integration: two candidates, planned, run stage by stage through the pool, and ranked.
// sim.yml's own wasm smoke runs one Top Gear plan of four candidates directly against the
// artifact; what this adds is the browser, the worker pool and the page around it.
//
// Skipped unless the build was made with PUBLIC_SIM_ENGINE=wasm and the published artifact's
// running program actually registers simPlan/simRank/simCount/simWeights -- checked at
// runtime, not assumed from the file's presence. This branch (sim-parity-web-b) has not
// merged the sim module lane that shipped those four exports on main (d464c7c) and must not,
// so a same-branch rebuild of sim/cmd/wasm from this worktree's own, pre-merge Go module
// would publish an artifact at the same path that still lacks them; a gate that only checked
// "does web/public/_sim/<ENGINE_VERSION>/sim.wasm exist" would then load that stale build and
// fail confusingly deep inside the plan/rank loop instead of skipping. Run it with
// `npm run test:e2e:real-engine`, after `make simdb && make artifacts && make publish-wasm`
// from the repository root.
//
// That script also sets FOREVER_DATA=real: playwright.config.ts's webServer otherwise
// defaults every spec's build to the small, fixed fixture item file (which every other
// spec in this directory wants, so its ids never move under it) -- and the fixture's own
// items are exactly where this defect was first reproduced. The fixture's Helm of Wrath
// (16963) has a row in items/warrior.json (a Classic Era build's ItemSparse) but no row in
// the ACTIVE build's own ItemSparse, so `.nth(0)`/`.nth(1)` against the fixture's search
// results could -- and did -- click it, and the real engine refused the whole request:
// `bulk: the build has no such item: 16963`. Real data plus the new simitems.json filter
// is what keeps that from happening again; twoKnownSearchCandidates() below reads both
// from the real build directory, not the fixture, for the same reason.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Browser } from '@playwright/test';
import { engineAssetUrl, ENGINE_VERSION } from '../../src/lib/sim/version';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const artifactPublished = existsSync(path.join(WEB_ROOT, 'public/_sim', ENGINE_VERSION, 'sim.wasm'));
const SKIP_REASON =
  `real-engine Top Gear smoke; publish web/public/_sim/${ENGINE_VERSION}/ first ` +
  '(make simdb && make artifacts && make publish-wasm) and run with npm run test:e2e:real-engine';

/**
 * Instantiates the published artifact in an ephemeral page (never the test's own `page`, so a
 * negative result leaves nothing running behind it) and reports which of the four bulk/weights
 * exports the running Go program actually registered on `window` -- the same handshake
 * `engine.ts`'s own `loadWasmEngine` performs (`Go`, `instantiateStreaming`, race against
 * `wasmready`), duplicated here in miniature because this file cannot import a browser-only
 * module into a Node-side gate check and must not assume what it has not observed.
 */
async function hasNewBulkExports(browser: Browser, baseURL: string, version: string): Promise<boolean> {
  const page = await browser.newPage();
  try {
    await page.goto(baseURL);
    // engineAssetUrl (version.ts) is the one place that knows the `/_sim/<version>/<file>`
    // path shape -- computed here, on the Node side, rather than rebuilt as a template
    // literal inside the evaluated closure, so this probe cannot silently drift from
    // `engine.ts`'s own `loadWasmEngine`, which builds the same two URLs the same way.
    const glueUrl = engineAssetUrl('sim.js', version);
    const wasmUrl = engineAssetUrl('sim.wasm', version);
    await page.addScriptTag({ url: glueUrl });
    return await page.evaluate(async (wasm) => {
      type GoGlue = new () => {
        importObject: WebAssembly.Imports;
        run(instance: WebAssembly.Instance): Promise<void>;
      };
      const w = window as unknown as {
        Go?: GoGlue;
        wasmready?: () => void;
        simPlan?: unknown;
        simRank?: unknown;
        simCount?: unknown;
        simWeights?: unknown;
      };
      try {
        if (typeof w.Go !== 'function') return false;
        const go = new w.Go();
        const { instance } = await WebAssembly.instantiateStreaming(fetch(wasm), go.importObject);
        const ready = new Promise<'ready'>((resolve) => {
          w.wasmready = () => resolve('ready');
        });
        const exited: Promise<'exited'> = go.run(instance).then(
          () => 'exited',
          () => 'exited',
        );
        const outcome = await Promise.race([ready, exited]);
        if (outcome !== 'ready') return false;
        return (
          typeof w.simPlan === 'function' &&
          typeof w.simRank === 'function' &&
          typeof w.simCount === 'function' &&
          typeof w.simWeights === 'function'
        );
      } catch {
        return false;
      }
    }, wasmUrl);
  } catch {
    return false;
  } finally {
    await page.close();
  }
}

const activeBuild = JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
  build: string;
};

// Real Forever item ids from the active build's own embedded database, the same two
// sim-real-engine.spec.ts uses -- which is what proves the wasm's simdb resolves real gear
// rather than echoing back whatever it was handed.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=21521`;

const DATA_ROOT = path.resolve(WEB_ROOT, '..', 'data', 'builds', activeBuild.build);

/**
 * Two warrior item ids the item search will actually offer AND the embedded engine
 * database actually carries -- read from the same two files the fix this test covers reads
 * from (`items/warrior.json`, the planner's per-class list; `simitems.json`, the engine's
 * kept set), rather than trusting `.nth(0)`/`.nth(1)` to land on usable rows by luck. Sorted
 * the same way `item-search.ts`'s `searchItems` sorts an untyped, unfiltered query (item
 * level desc, then name), so these are the first two rows the page itself would show.
 */
function twoKnownSearchCandidates(): [number, number] {
  const items = (
    JSON.parse(readFileSync(path.join(DATA_ROOT, 'items', 'warrior.json'), 'utf8')) as {
      items: { id: number; name: string; item_level: number; required_level: number }[];
    }
  ).items;
  const known = new Set(
    (JSON.parse(readFileSync(path.join(DATA_ROOT, 'simitems.json'), 'utf8')) as { items: number[] }).items,
  );
  const candidates = items
    .filter((item) => item.required_level <= 60 && known.has(item.id))
    .sort((a, b) => b.item_level - a.item_level || a.name.localeCompare(b.name));
  if (candidates.length < 2) {
    throw new Error(
      `build ${activeBuild.build} offers fewer than two usable, engine-known warrior items to search for`,
    );
  }
  return [candidates[0].id, candidates[1].id];
}

test('two candidates plan, run and rank against the equipped set', async ({ page, browser, baseURL }) => {
  test.slow();

  const hasExports =
    process.env.PUBLIC_SIM_ENGINE === 'wasm' &&
    artifactPublished &&
    (await hasNewBulkExports(browser, baseURL ?? 'http://localhost:4340', ENGINE_VERSION));
  test.skip(!hasExports, SKIP_REASON);

  const pageErrors: string[] = [];
  page.on('pageerror', (error) => pageErrors.push(error.message));

  await page.goto('/sim/gear');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-slot-grid')).toBeVisible();

  // Two search candidates rather than two bag rows: this export carries no bags, and the
  // search is the path that exercises the item database on both sides. Both ids are known
  // to be usable AND known to the engine (twoKnownSearchCandidates), so clicking each by
  // its own id proves the item search actually offered it, rather than clicking whatever
  // happened to render first.
  await page.getByTestId('sim-item-search').getByRole('searchbox').fill('');
  const [firstId, secondId] = twoKnownSearchCandidates();
  await page.getByTestId(`sim-search-add-${firstId}`).click();
  await page.getByTestId(`sim-search-add-${secondId}`).click();

  await expect(page.getByTestId('sim-combo-count')).toHaveText(/[1-9]\d* valid combinations?/);

  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toBeVisible();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 180_000 });

  const rows = page.getByTestId('sim-combo-row');
  expect(await rows.count()).toBeGreaterThan(0);
  await expect(page.getByTestId('sim-equipped-line')).toContainText(/\d/);
  expect(pageErrors).toEqual([]);
});
