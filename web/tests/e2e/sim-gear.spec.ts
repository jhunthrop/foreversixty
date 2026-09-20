// web/tests/e2e/sim-gear.spec.ts
// Top Gear against the checked-in fake engine (PUBLIC_SIM_ENGINE defaults to 'fake'), so
// no Go toolchain and no wasm artifact is needed. The real engine gets one gated run of its
// own in tests/e2e/sim-gear-real-engine.spec.ts.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

// The fixture warrior, wearing two of the six fixture items so the grid has an equipped row
// to lock and a bag row to tick.
const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

export async function loadGear(page: Page, at = '/sim/gear'): Promise<void> {
  await page.goto(at);
  await page.getByTestId('sim-addon-input').fill(FURY);
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

test('locking a slot disables every candidate in it', async ({ page }) => {
  await loadGear(page);
  await page.getByTestId('sim-lock-head').check();
  await expect(page.getByTestId('sim-candidate-head-12640').getByRole('checkbox')).toBeDisabled();
  await page.getByTestId('sim-lock-head').uncheck();
  await expect(page.getByTestId('sim-candidate-head-12640').getByRole('checkbox')).toBeEnabled();
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
  await expect(page.getByTestId('sim-combo-count')).toHaveText(/0 valid combinations|Counting/);
  await page.getByTestId('sim-candidate-head-12640').getByRole('checkbox').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText('1 valid combination');
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
  await expect(page.getByTestId('sim-combo-count')).toHaveText('1 valid combination');
  await page.getByTestId('sim-consumable-flask_of_supreme_power').check();
  await page.getByTestId('sim-consumable-elixir_of_the_mongoose').check();
  // (1 head + 1) x 2 alternatives
  await expect(page.getByTestId('sim-combo-count')).toHaveText('4 valid combinations');
});

test('a consumable candidate is named, not spelled as an id', async ({ page }) => {
  await loadGear(page);
  await expect(page.getByTestId('sim-consumables')).toContainText('Flask of Supreme Power');
});
