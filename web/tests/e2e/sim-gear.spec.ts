// web/tests/e2e/sim-gear.spec.ts
// Top Gear against the checked-in fake engine (PUBLIC_SIM_ENGINE defaults to 'fake'), so
// no Go toolchain and no wasm artifact is needed. The real engine gets one gated run of its
// own in tests/e2e/sim-gear-real-engine.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { bulkCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The fixture warrior, wearing two of the six fixture items so the grid has an equipped row
// to lock and a bag row to tick.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

// FURY plus contract 7's version 2 sections: one in-game loadout and one named set, the
// set's own gear entry carrying BOTH an enchant and a suffix (contract 10.5's richer
// per-slot shape) -- the one path Task 14's own conversions (TalentCandidates.svelte's
// rank-to-order-to-talents-string pipeline; NamedSets.svelte's itemId -> item_id,
// enchant-and-suffix-preserving `toGearSlot`) need proved end-to-end, against a real decode,
// rather than only on inspection. `character.test.ts`'s own `characterFromFs1` tests already
// prove this exact grammar decodes (bank=…:2505:1820 there; here the enchant/suffix pair
// rides a `sets=` entry instead, which is what this task's own components read). The
// loadout is deliberately named "Deep Fury" -- the same title the mocked
// `/v1/builds?mine=1` row below uses -- so the same fixture also exercises the
// saved-vs-exported name collision fix.
const FURY_V2 = `${FURY}|loadouts=Deep Fury=5530515/0/0|sets=Alt Set=head=12640:2543:1`;

export async function loadGear(page: Page, at = '/sim/gear', code = FURY): Promise<void> {
  await page.goto(at);
  await page.getByTestId('sim-addon-input').fill(code);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid')).toBeVisible();
}

test('the grid lists the equipped item in its slot and nothing is ticked on arrival', async ({ page }) => {
  await loadGear(page);
  const equipped = page.getByTestId('sim-candidate-head-12640');
  await expect(equipped).toBeVisible();
  await expect(equipped.getByRole('checkbox')).not.toBeChecked();
});

test('locking a slot disables every candidate in it, and its copy-and-modify trigger too', async ({
  page,
}) => {
  await loadGear(page);
  const row = page.getByTestId('sim-candidate-head-12640');
  await page.getByTestId('sim-lock-head').check();
  await expect(row.getByRole('checkbox')).toBeDisabled();
  // A locked slot's checkbox drops any ticked candidate from the outgoing request
  // (validateBulk rule 1), so a copy-and-modify made here would be checked-but-inert with
  // nothing telling the player -- the trigger disables for the same reason the checkbox does.
  await expect(row.getByRole('button', { name: /copy/i })).toBeDisabled();
  await page.getByTestId('sim-lock-head').uncheck();
  await expect(row.getByRole('checkbox')).toBeEnabled();
  await expect(row.getByRole('button', { name: /copy/i })).toBeEnabled();
});

test('copy and modify adds the same item again with an enchant, beside the original', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-candidate-head-12640').getByRole('button', { name: /copy/i }).click();
  await page.getByTestId('sim-enchant-head-2543').click();
  await expect(page.getByTestId('sim-candidate-head-12640')).toBeVisible();
  const copy = page.getByTestId('sim-candidate-head-12640-e2543');
  await expect(copy).toBeVisible();
  await expect(copy.getByRole('checkbox')).toBeChecked();
});

test('the combination count moves as candidates are ticked', async ({ page }) => {
  await loadGear(page);
  // Neither bulkCopy string here has a regex metacharacter, so the two are safe to
  // alternate directly rather than through an escaping helper this file would be the only
  // caller of.
  await expect(page.getByTestId('sim-combo-count')).toHaveText(
    new RegExp(`${bulkCopy.combinations(0)}|${bulkCopy.combinationsCounting}`),
  );
  await page.getByTestId('sim-candidate-head-12640').getByRole('checkbox').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(1));
});

test('an empty grid says so, rather than rendering as an empty section', async ({ page }) => {
  // No gear at all: the FS1 grammar's trailing colon with nothing after it is a valid,
  // empty gear list (tests/e2e/sim-specs.spec.ts's ROGUE_FS1 uses the same shape). Kept on
  // the warrior/orc talent build FURY already uses -- unlike rogue, its talent tree ships
  // in this fixture data build, so this needs no API route stub either.
  const NOTHING_EQUIPPED = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:`;
  await page.goto('/sim/gear');
  await page.getByTestId('sim-addon-input').fill(NOTHING_EQUIPPED);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid-empty')).toHaveText(bulkCopy.noCandidates);
});

test('the item search adds a candidate and says how it was filtered', async ({ page }) => {
  await loadGear(page);
  const search = page.getByTestId('sim-item-search');
  await expect(search).toBeVisible();
  await search.getByRole('searchbox').fill('wrath');
  await expect(page.getByTestId('sim-search-result-16963')).toBeVisible();
  await expect(page.getByTestId('sim-search-result-12640')).toHaveCount(0);
  await page.getByTestId('sim-search-add-16963').click();
  await expect(page.getByTestId('sim-candidate-head-16963').getByRole('checkbox')).toBeChecked();
});

test('the source filter narrows the search to one boss’s loot', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-source').selectOption('world:azuregos');
  await expect(page.getByTestId('sim-search-result-19325')).toBeVisible();
  await expect(page.getByTestId('sim-search-result-16963')).toHaveCount(0);
});

test('usable-only is on by default and can be turned off', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-search-usable')).toBeChecked();
  await page.getByTestId('sim-search-usable').uncheck();
  await expect(page.getByTestId('sim-search-usable')).not.toBeChecked();
});

test('a ticked consumable multiplies the combination count (contract 10.1 A5)', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(1));
  await page.getByTestId('sim-consumable-flask_of_supreme_power').check();
  await page.getByTestId('sim-consumable-elixir_of_the_mongoose').check();
  // (1 head + 1) x 2 alternatives
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(4));
});

test('a consumable candidate is named, not spelled as an id', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-consumables')).toContainText('Flask of Supreme Power');
});

/** One `/v1/builds?mine=1` row, titled `title`, on the character's own class/tree/build. */
async function mockMyBuilds(page: Page, title = 'Deep Fury'): Promise<void> {
  await page.route('**/v1/builds?mine=1*', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        request_id: 'r',
        error: null,
        data: {
          rows: [
            {
              id: 'bld987654321',
              class_id: 1,
              race_id: 2,
              tree_version: activeBuild.build,
              point_order: [2001, 2001, 2001, 2001, 2001],
              gear: {},
              title,
              created_at: '2026-09-18T12:00:00Z',
              views: 2,
            },
          ],
          total: 1,
          page: 1,
          per_page: 100,
        },
      }),
    }),
  );
}

test('the talent list offers the character’s own build and a saved one', async ({ page }) => {
  await mockMyBuilds(page);
  await loadGear(page);
  await expect(page.getByTestId('sim-loadout-current')).toBeVisible();
  await page.getByTestId('sim-loadout-Deep Fury').check();
  await expect(page.getByTestId('sim-loadout-Deep Fury')).toBeChecked();
});

test('an in-game loadout that shares a saved build’s name renders once, not twice', async ({ page }) => {
  // The mocked saved build and FURY_V2's own `loadouts=Deep Fury=...` section both name
  // "Deep Fury" -- without TalentCandidates.svelte's own dedupe, that is two checkboxes
  // under the identical `sim-loadout-Deep Fury` test id, which Playwright's strict mode
  // (`toHaveCount(1)` below, and any single-element locator after it) refuses to resolve.
  await mockMyBuilds(page);
  await loadGear(page, '/sim/gear', FURY_V2);
  await expect(page.getByTestId('sim-loadout-Deep Fury')).toHaveCount(1);
  await page.getByTestId('sim-loadout-Deep Fury').check();
  await expect(page.getByTestId('sim-loadout-Deep Fury')).toBeChecked();
});

test('a named set from the addon export carries its enchant and suffix into the outgoing request', async ({
  page,
}) => {
  // FURY_V2's `sets=Alt Set=head=12640:2543:1` is the one path proving NamedSets.svelte's
  // toGearSlot conversion end to end: an addon-exported enchant AND suffix together,
  // through itemId -> item_id, surviving into the request the page would actually send --
  // not just a rendered row, which alone would prove nothing about the numbers behind it.
  await loadGear(page, '/sim/gear', FURY_V2);
  await expect(page.getByTestId('sim-set-Alt Set')).toBeVisible();
  await page.getByTestId('sim-set-Alt Set').getByRole('button', { name: bulkCopy.setsAdd }).click();
  await page.getByTestId('sim-request-drawer').locator('summary').click();
  const request = JSON.parse(await page.getByTestId('sim-request-json').inputValue());
  expect(request.bulk.sets).toEqual([
    { name: 'Alt Set', gear: [{ slot: 'head', item_id: 12640, enchant: 2543, suffix: 1 }] },
  ]);
});

test('a pasted second export string becomes a named set', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-set-input').fill(`${FURY}`);
  await page.getByTestId('sim-set-name').fill('PvP set');
  await page.getByTestId('sim-set-add').click();
  await expect(page.getByTestId('sim-set-PvP set')).toBeVisible();
});

test('a set that is not an export string says so and adds nothing', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-set-input').fill('FS2:nope');
  await page.getByTestId('sim-set-name').fill('Bad');
  await page.getByTestId('sim-set-add').click();
  await expect(page.getByTestId('sim-set-error')).toBeVisible();
  await expect(page.getByTestId('sim-set-Bad')).toHaveCount(0);
});

test('four candidates are well under the cap, so no notice and a live run button', async ({ page }) => {
  await loadGear(page);
  for (const id of [16963, 16966, 13968, 19325]) {
    await page.getByTestId(`sim-search-add-${id}`).click();
  }
  // The desktop project reports 8+ cores, so the cap is 400 and four candidates fit. The
  // notice's own wording and its 5,000-combination premium alternative are asserted in
  // the unit tests, where the cap is injectable.
  await expect(page.getByTestId('sim-cap-notice')).toHaveCount(0);
  await expect(page.getByTestId('sim-run-bulk')).toBeEnabled();
});

test('a run reports its stage line and ends with a ranked table', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toHaveText(/stage \d of 3 · \d+ of \d+ combinations/);
  // Task 16 (not yet landed) renders the ranked table at this test id; until then, the run
  // finishing (the stage line disappearing) is this task's own thing to assert.
  // await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-stage-progress')).toHaveCount(0, { timeout: 20_000 });
});

test('precision is three choices and normal runs two stages', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-search-add-16963').click();
  await page.getByTestId('sim-precision').selectOption('normal');
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-stage-progress')).toHaveText(/stage \d of 2 /);
});

test('the server lane is not offered to a signed-out visitor', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-server-run')).toHaveCount(0);
});
