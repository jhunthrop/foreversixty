// web/tests/e2e/sim-unsimulated-spec.spec.ts
// Task 3 (healer review): a Restoration Druid used to load into the full damage UI, let the
// player set up and start a run, and then print sim/request's raw
// `combine: part 0 failed: request: … unsupported spec: "druid-restoration"` -- the only
// honest sentence on the page. This proves the fix from the browser, on the two pages the
// brief names: `/sim` (RunControl.svelte) and `/sim/weights` (BulkRunBar.svelte, the
// BLOCKER -- the run used to accept and print nothing for 64s).
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { simCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The fixture data build (FOREVER_DATA=fixture) ships only a warrior talent tree
// (src/fixtures/planner/talents/warrior.json) -- tests/e2e/sim-specs.spec.ts's own
// ROGUE_TALENTS already established the pattern this follows: stub
// `/data/<build>/talents/druid.json` with the smallest legal file that decodes to a chosen
// spec, one point in one talent, no prereq. Three trees (a druid always has three), all
// empty but the last, so the character's every point lands in Restoration
// (`data/curated/specs.json`'s druid-restoration is tree_index 2, the third tree).
//
// The healer reviewer's own paste export
// (FS1:1.60.1.69893:druid:tauren:0/0/5553335153113251:head=19999,…) names real item ids and
// a full 51-point Restoration build; it does not decode against this synthetic file (the
// talent count and ranks would not match), and no druid item file exists under
// FOREVER_DATA=fixture for real gear ids to resolve against either (a fetch failure there
// is caught and shown as an empty strip, not a load failure, but there is no need to rely on
// that when the rogue precedent already proves the minimal-file approach loads cleanly). The
// gear list is empty for the same reason ROGUE_FS1 carries none.
const DRUID_TALENTS = {
  build: activeBuild.build,
  class_id: 11,
  class_slug: 'druid',
  trees: [
    { id: 1101, name: 'Balance', position: 0, background: 'fixture_balance', talents: [] },
    { id: 1102, name: 'Feral Combat', position: 1, background: 'fixture_feral', talents: [] },
    {
      id: 1103,
      name: 'Restoration',
      position: 2,
      background: 'fixture_restoration',
      talents: [
        {
          id: 11031,
          name: 'Improved Mark of the Wild',
          icon: 'fixture_improved_mark_of_the_wild',
          max_rank: 5,
          tier: 0,
          column: 0,
          prereq_talent_id: null,
          prereq_rank: null,
          spell_id: 17055,
          ranks: Array.from({ length: 5 }, (_, i) => ({
            spell_id: 17055 + i,
            description: `Rank ${i + 1}.`,
          })),
        },
      ],
    },
  ],
};

const RESTO_DRUID_FS1 = `FS1:${activeBuild.build}:druid:tauren:0/0/1:`;

async function stubDruidTalents(page: Page): Promise<void> {
  await page.route('**/data/*/talents/druid.json', (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(DRUID_TALENTS) }),
  );
}

async function loadRestoDruid(page: Page, route: string): Promise<void> {
  await stubDruidTalents(page);
  await page.goto(route);
  await page.getByTestId('sim-addon-input').fill(RESTO_DRUID_FS1);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

/** Fix round 1: signs the visitor in as premium, the same shape sim-drops.spec.ts's own
 *  `stubPremiumRun` uses, so `RunControl`'s "Run on servers" button renders at all --
 *  `{#if premium}` -- and its own `!simulated` gate can be asserted rather than the button
 *  simply being absent. */
async function stubPremium(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ data: { user: { premium: true }, characters: [] }, error: null }),
    }),
  );
}

test('/sim: a Restoration Druid loads the strip, but the run control is disabled and honest', async ({
  page,
}) => {
  await stubPremium(page);
  await loadRestoDruid(page, '/sim');

  // The BLOCKER/MAJOR findings' first demand: the strip still loads and still shows the
  // character -- gear, talents, header -- whatever the spec.
  await expect(page.getByTestId('sim-character')).toContainText('Restoration Druid');

  const button = page.getByTestId('sim-run-button');
  await expect(button).toBeDisabled();

  const note = page.getByTestId('sim-run-not-simulated');
  await expect(note).toBeVisible();
  await expect(note).toHaveText(simCopy.runNotSimulated('Restoration'));
  // The honest line is a <p>, never a title= tooltip (Task 7 removes every title= in this
  // lane; nothing here should add one back).
  await expect(note).toHaveJSProperty('tagName', 'P');
  await expect(button).not.toHaveAttribute('title', /.+/);

  // Fix round 1: the premium "Run on servers" button is a second run path beside the
  // primary one, and used to stay enabled for an unsimulated spec.
  const serverButton = page.getByTestId('sim-server-run');
  await expect(serverButton).toBeVisible();
  await expect(serverButton).toBeDisabled();

  await button.click({ force: true });
  await expect(button).toBeDisabled();
  await expect(page.getByTestId('sim-dps')).toHaveText('—');

  await serverButton.click({ force: true });
  await expect(page.getByTestId('sim-dps')).toHaveText('—');

  await expect(page.locator('body')).not.toContainText('combine: part 0 failed');
  await expect(page.locator('body')).not.toContainText('unsupported spec');
});

test('/sim/weights: the same Restoration Druid never accepts a run and returns nothing (BLOCKER)', async ({
  page,
}) => {
  await loadRestoDruid(page, '/sim/weights');

  await expect(page.getByTestId('sim-character')).toContainText('Restoration Druid');

  const button = page.getByTestId('sim-run-bulk');
  await expect(button).toBeDisabled();

  const note = page.getByTestId('sim-run-not-simulated');
  await expect(note).toBeVisible();
  await expect(note).toHaveText(simCopy.runNotSimulated('Restoration'));
  await expect(note).toHaveJSProperty('tagName', 'P');

  // The BLOCKER itself: clicking (or trying to) must not start a run that prints nothing
  // for 64 seconds -- the button refuses the click, and the page settles immediately with
  // no stage progress and no weights.
  await button.click({ force: true });
  await expect(page.getByTestId('sim-stage-progress')).toHaveCount(0);
  await expect(page.getByTestId('sim-weights')).toHaveCount(0);

  await expect(page.locator('body')).not.toContainText('combine: part 0 failed');
  await expect(page.locator('body')).not.toContainText('unsupported spec');
});
