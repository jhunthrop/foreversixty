// web/tests/e2e/sim-talents-real-engine.spec.ts
// dps-minmaxer round 3, finding 5 (review.md near line 187): "a build added through ADD A
// BUILD gets its own credible DPS number but never joins the ranked table beside the
// current build, so comparing two builds still means subtracting by hand." Talent
// Compare's own fake-engine coverage (sim-talents.spec.ts) cannot catch this: the bug is
// sim/bulk's stage-1 cut (rank.go), which the fake engine stub never runs at all, so a
// two-loadout comparison always looked fine against it. This is the one place that proves
// the real ladder -- fast precision's first cut, the one signed-out browser-lane runs
// actually use -- keeps every ticked talent loadout through to the final result rather
// than narrowing to the survivors the way Top Gear's own candidate search needs to.
//
// Skipped unless the build was made with PUBLIC_SIM_ENGINE=wasm and the published
// artifact's running program registers simPlan/simRank/simCount -- see
// sim-gear-real-engine.spec.ts's own header for why this is checked at runtime rather than
// assumed from the file's presence. Run it with `npm run test:e2e:real-engine`, after
// `make simdb && make artifacts && make publish-wasm` from the repository root.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { apiUnreachable, expect, test } from './support/test';
import {
  artifactPublished,
  E2E_ARTIFACT,
  hasWasmExports,
  publishedTalents,
  realEngineSkipReason,
} from './support/real-engine';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const SKIP_REASON = realEngineSkipReason('talent compare two-loadout ranking');

const activeBuild = JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/active-build.json'), 'utf8')) as {
  build: string;
};

// The exact character and gear the production repro used (dps-minmaxer round 3's own
// r10-r12 scripts, and this fix's talents-fix-01/02 scripts under .superpowers/persona/):
// real item ids from the active build's own embedded database. A near-naked character (as
// sim-gear-real-engine.spec.ts's own FURY is) does not reproduce this reliably -- with so
// little gear, an arms-vs-fury talent swap is too close for the fast ladder's SlackSE
// overlap allowance to ever drop either candidate, which is exactly why this needs the
// same full gear set the production defect was actually found with.
// The talent strings are the published builds' own (support/real-engine.ts), so they stay
// legal when a client build moves a tree.
const FURY =
  `FS1:${activeBuild.build}:warrior:orc:${publishedTalents(WEB_ROOT, activeBuild.build, 'warrior-fury')}:` +
  'head=12640,neck=21809,shoulder=275737,back=21394,chest=19690,wrist=19687,hands=21278,' +
  'waist=19823,legs=16543,feet=16545,ranged=22656,finger1=19325,finger2=21182,' +
  'trinket1=19406,trinket2=249469,main_hand=22384,off_hand=18827';
// A genuinely different, drastically worse spec on the SAME gear -- an arms-heavy talent
// string, no gear after the last colon (valid FS1; the inline planner locks gear to the
// loaded character regardless).
const ALT_TALENTS = `FS1:${activeBuild.build}:warrior:orc:${publishedTalents(WEB_ROOT, activeBuild.build, 'warrior-arms')}:`;

// The saved-builds row (/v1/builds?mine=1) is a signed-in nicety the hermetic site cannot
// reach; answer it unreachable, as sim-gear.spec.ts does.
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/builds?*'));

test('a pasted talent build joins the current build in one ranked table', async ({
  page,
  browser,
  baseURL,
}) => {
  test.slow();

  const hasExports =
    process.env.PUBLIC_SIM_ENGINE === 'wasm' &&
    artifactPublished(WEB_ROOT, E2E_ARTIFACT) &&
    (await hasWasmExports(browser, baseURL ?? 'http://localhost:4341', E2E_ARTIFACT, [
      'simPlan',
      'simRank',
      'simCount',
    ]));
  test.skip(!hasExports, SKIP_REASON);

  const pageErrors: string[] = [];
  page.on('pageerror', (error) => pageErrors.push(error.message));

  await page.goto('/sim/talents');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();

  // ADD A BUILD, paste the alt talent string, submit the import and accept it -- by
  // testid, not an accessible-name regex ("Add a build" labels both the outer toggle and
  // the inner accept button).
  await page.getByTestId('sim-loadout-add').click();
  await page.getByTestId('import-code').fill(ALT_TALENTS);
  await page.getByTestId('import-submit').click();
  await expect(page.getByTestId('import-error')).toHaveCount(0);
  await page.getByTestId('sim-loadout-accept').click();

  const pastedBuild = page.getByTestId('sim-loadout-Build 1');
  await expect(pastedBuild).toBeChecked();
  await page.getByTestId('sim-loadout-current').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(/^2 valid combinations$/, {
    // The count comes back from the engine's own plan pass, and CI loads a 4 MB wasm per worker:
    // give it the same allowance a run gets, not the 5 s default.
    timeout: 20_000,
  });

  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 180_000 });

  // The bug: sim/bulk's fast-precision cut narrowed a two-loadout talents-mode request to
  // its stage-1 survivor, so only one of the two ticked builds ever reached the final
  // result -- the page could rank a build it was never told the answer for. Both must be
  // rows in the one table, with the delta between them, whichever one led.
  const rowTexts = await page.getByTestId('sim-combo-row').allTextContents();
  expect(rowTexts.length, `ranked rows: ${JSON.stringify(rowTexts)}`).toBe(2);
  const joined = rowTexts.join(' | ');
  expect(joined).toContain('Your current build');
  expect(joined).toContain('Build 1');
  expect(pageErrors).toEqual([]);
});
