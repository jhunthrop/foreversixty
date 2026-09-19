// web/tests/e2e/sim-real-engine.spec.ts
// The one end-to-end proof that the real wasm engine web.yml's deploy job publishes
// survives the full browser integration: web/src/lib/sim/engine.ts's loadWasmEngine actually
// fetches, instantiates and runs the artifact under /_sim/<ENGINE_VERSION>/, through the
// exact /sim UI a visitor uses -- not just that the four exports work in Node (sim.yml's
// wasm smoke test already proves that, directly against the artifact, at exactly 500
// iterations).
//
// Every other spec in this directory runs the checked-in fake engine
// (src/fixtures/sim/engine-fake.ts), because PUBLIC_SIM_ENGINE defaults to 'fake' and no
// wasm build or Go toolchain is needed for it. This spec is the opposite by design and
// stays skipped unless both are true: the build was made with PUBLIC_SIM_ENGINE=wasm, and
// the artifact it names was actually published to web/public/_sim/<ENGINE_VERSION>/ before
// that build ran (web/public/_sim/README.md's contract).
//
// Run it with `npm run test:e2e:real-engine` from web/, after publishing the browser pair
// from the repository root: `make simdb && make artifacts && make publish-wasm`. That npm
// script sets PUBLIC_SIM_ENGINE=wasm and E2E_PORT=4325 so its preview server never collides
// with the fixture preview `npm run test:e2e` may already be holding on 4321.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test, type Page } from '@playwright/test';
import { ENGINE_VERSION, engineLabel } from '../../src/lib/sim/version';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const artifactPublished = existsSync(path.join(WEB_ROOT, 'public/_sim', ENGINE_VERSION, 'sim.wasm'));

test.skip(
  process.env.PUBLIC_SIM_ENGINE !== 'wasm' || !artifactPublished,
  `real-engine smoke; publish web/public/_sim/${ENGINE_VERSION}/ first ` +
    '(from the repository root: make simdb && make artifacts && make publish-wasm) and run ' +
    'with npm run test:e2e:real-engine',
);

const activeBuild = JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
  build: string;
};

// A level 60 orc warrior geared with two real Forever item ids from the active build's own
// embedded database (sim/adapter/testdata/warrior-fury.request.json's head and main_hand) --
// loading it here is what proves the wasm's embedded item database (`make simdb`) resolves
// real gear the fake engine only echoes back, rather than proving nothing more than "the
// page renders". Note this is NOT sim-run.spec.ts's fixture string: that one's main_hand
// (11726) is not itemised in every build's simdb.bin, only in the one the fake engine never
// checks against.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=21521`;

async function loadFury(page: Page): Promise<void> {
  await page.goto('/sim');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('a real wasm run against the fixture character returns a positive DPS', async ({ page }) => {
  // Real wasm instantiation plus a full run, not the fake engine's short timers -- this
  // test's whole point is to be slow in a way sim-run.spec.ts's fake-engine coverage is not.
  test.slow();

  // Uncaught exceptions only, not every console.error: a local preview with no API server
  // behind it logs its own unrelated background fetch failures (the history panel's
  // `GET /v1/sims?mine=1`, `GET /v1/me`) as console errors on every /sim load, fake engine or
  // real, and those are not what this test is watching for. A wasm crash is what this
  // guards -- sim/cmd/wasm/main.go's own panic when the page races it (the bug this test's
  // fix for web/src/lib/sim/engine.ts addressed) surfaces as a console.log from the Go
  // runtime, not a pageerror, so the dps/progress assertions below are the load-bearing
  // check; this is a second line of defence for anything that throws in the JS glue itself.
  const pageErrors: string[] = [];
  page.on('pageerror', (error) => pageErrors.push(error.message));

  await loadFury(page);

  // The page names the engine it actually loaded, straight from the wasm module's own
  // simEngineVersion global (engine.ts) -- confirming this run is against the published
  // artifact and not a build that silently fell back to the fake engine.
  await expect(page.getByTestId('sim-engine-version')).toHaveText(engineLabel(ENGINE_VERSION));

  const button = page.getByTestId('sim-run-button');
  await expect(button).toHaveText('Run sim');
  await expect(page.getByTestId('sim-dps')).toHaveText('—');

  await button.click();
  // /sim's run control has no UI for an arbitrary iteration count below the "normal" 3,000
  // it always starts a run at (types.ts's ITERATIONS.live=500 is reserved for the planner's
  // live-DPS preview, not this page) -- sim.yml's wasm smoke test already proves the four
  // exports at exactly 500 iterations. What this test proves instead is the piece that smoke
  // test cannot reach: the browser fetching, instantiating and running THIS published
  // artifact through the real RunControl wiring, for the "normal" run every visitor gets.
  await expect(button).toHaveText('Run again', { timeout: 30_000 });
  await expect(page.getByTestId('sim-progress')).toHaveText('3,000 iterations');

  const dpsText = await page.getByTestId('sim-dps').textContent();
  expect(dpsText).not.toBeNull();
  expect(dpsText).not.toBe('—');
  const dps = Number(dpsText!.replace(/[^0-9.]/g, ''));
  expect(dps).toBeGreaterThan(0);

  expect(pageErrors).toEqual([]);
});
