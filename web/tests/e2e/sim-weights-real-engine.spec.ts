// web/tests/e2e/sim-weights-real-engine.spec.ts
// The one end-to-end proof that a browser-lane /sim/weights run against the real,
// published wasm artifact actually finishes with something on screen: the exact
// production repro for defect A (2026-09-21) -- click Run at the default ("fast")
// precision, wait, and the page used to land on "Run again" with no table, no message and
// no error of any kind. Every other sim-weights spec in this directory runs the checked-in
// fake engine (src/fixtures/sim/engine-fake.ts), which always succeeds and so never
// exercised the real engine's own iteration accounting or its failure path.
//
// Two things were wrong and both are fixed by the time this spec is expected to pass:
//   * sim/api's own closed iteration set (500/3000/10000) refused a weights request's
//     Iterations even though a weights request's own count is the engine's per-direction
//     base, not a plain run's choice from the settings bar -- so the browser lane's own
//     guarded "fast" default (weights.ts's WEIGHTS_BROWSER_DEFAULT_ITERATIONS, 60) was
//     refused on every single run at the page's own default precision.
//   * that refusal came back as an ordinary-looking, resolved SimResult (weightsJSON's own
//     failJSON, sim/cmd/wasm/exports.go) which engine.ts's `simWeights` binding never
//     unwrapped the way `simValidate`/`simRank` already do, so `runWeightsRun` read it as a
//     finished run and the page never learned anything had gone wrong.
//
// This spec proves the fix at the one layer neither of the two unit-test layers can reach:
// the real, published wasm artifact, run through the actual browser worker pool. Skipped
// unless the build was made with PUBLIC_SIM_ENGINE=wasm, the artifact was published first,
// and the running program actually registers simWeights. Run it with
// `npm run test:e2e:real-engine`, after `make simdb && make artifacts && make publish-wasm`
// from the repository root.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { expect, test } from '@playwright/test';
import { ENGINE_VERSION } from '../../src/lib/sim/version';
import { artifactPublished, hasWasmExports, realEngineSkipReason } from './support/real-engine';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');

const activeBuild = JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
  build: string;
};

// Defect A's own repro character (production, engine 7a45995ae): a Fury Warrior with real
// gear ids from the active build's embedded database, talents included -- the field report
// pasted this exact FS1 export against https://foreversixty.gg/sim/weights.
const FURY_REPRO =
  `FS1:${activeBuild.build}:warrior:orc:33305013002/150531000051310051/0:` +
  'head=12640,neck=21809,shoulder=275737,back=21394,chest=19690,wrist=19687,hands=21278,' +
  'waist=19823,legs=16543,feet=16545,ranged=22656,finger1=19325,finger2=21182,' +
  'trinket1=19406,trinket2=249469,main_hand=22384,off_hand=18827';

test('a browser-lane weights run at the default precision finishes with a table, not silence', async ({
  page,
  browser,
  baseURL,
}) => {
  test.slow();

  const published = artifactPublished(WEB_ROOT, ENGINE_VERSION);
  const hasExports =
    process.env.PUBLIC_SIM_ENGINE === 'wasm' &&
    published &&
    (await hasWasmExports(browser, baseURL ?? 'http://localhost:4340', ENGINE_VERSION, ['simWeights']));
  test.skip(!hasExports, realEngineSkipReason('weights smoke'));

  const pageErrors: string[] = [];
  page.on('pageerror', (error) => pageErrors.push(error.message));

  await page.goto('/sim/weights');
  await page.getByTestId('sim-addon-input').fill(FURY_REPRO);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // The default precision is "fast" (bulk-store.svelte.ts's own `precision = $state('fast')`)
  // -- left untouched, because that is exactly the button the field report pressed.
  await page.getByTestId('sim-run-bulk').click();

  // The page must never end this run silently: either the table appears, or a message
  // does. Racing the two, rather than asserting the table alone first, is what makes this
  // assertion fail loudly (a real Playwright timeout naming which one never showed) instead
  // of hanging for the suite's own default timeout with no clue which branch was expected.
  await Promise.race([
    page.getByTestId('sim-weights').waitFor({ state: 'visible', timeout: 60_000 }),
    page.getByTestId('sim-message').waitFor({ state: 'visible', timeout: 60_000 }),
  ]);

  const message = page.getByTestId('sim-message');
  if (await message.isVisible()) {
    // A message is an honest, acceptable outcome for a genuinely failed run -- but not for
    // this character: warrior-fury is a modelled dps spec, its default weight_stats are all
    // known ids, and this exact request already runs clean end to end against this engine
    // build natively (sim/api's own TestWeightsValidation, the browser-lane 60-iteration
    // case). A message here means the fix regressed, not that this spec's assumption was
    // ever meant to allow it.
    const text = await message.textContent();
    throw new Error(`weights run ended with a message instead of a table: ${text}`);
  }

  await expect(page.getByTestId('sim-weights')).toBeVisible();
  // Contract 10.8: the reference stat is always exactly 1.
  await expect(page.getByTestId('sim-weight-attack_power')).toContainText('1.00');
  await expect(page.getByTestId('sim-pawn')).toContainText('( Pawn: v1:');

  expect(pageErrors).toEqual([]);
});
