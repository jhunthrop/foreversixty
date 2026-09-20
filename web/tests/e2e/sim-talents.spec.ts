// web/tests/e2e/sim-talents.spec.ts
// Design 3.4: /sim/talents is /sim/gear's own island, TopGear.svelte, with `store.tool`
// set to 'talents' instead of 'gear'. There is no separate component to test here -- this
// file proves the mode difference: gear locked, no gear-shaped sections, the talent list is
// the only thing that can be picked.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { bulkCopy } from '../../src/lib/sim/copy';

const activeBuild = JSON.parse(
  readFileSync(path.join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };

const FURY = `FS1:${activeBuild.build}:warrior:orc:0/5530515/0:head=12640,main_hand=12784`;

async function loadTalents(page: Page): Promise<void> {
  await page.goto('/sim/talents');
  await page.getByTestId('sim-addon-input').fill(FURY);
  await page.getByTestId('sim-addon-load').click();
  await expect(page.getByTestId('sim-character')).toBeVisible();
}

test('there is no slot grid, no item search and no named sets — only the talent list', async ({ page }) => {
  await loadTalents(page);
  await expect(page.getByTestId('sim-talent-candidates')).toBeVisible();
  await expect(page.getByTestId('sim-slot-grid')).toHaveCount(0);
  await expect(page.getByTestId('sim-item-search')).toHaveCount(0);
  await expect(page.getByTestId('sim-named-sets')).toHaveCount(0);
});

test('ticking the character’s own build makes the run button live and ranks it', async ({ page }) => {
  await loadTalents(page);
  await page.getByTestId('sim-loadout-current').check();
  await expect(page.getByTestId('sim-combo-count')).toHaveText(bulkCopy.combinations(1));
  await page.getByTestId('sim-run-bulk').click();
  await expect(page.getByTestId('sim-combos')).toBeVisible({ timeout: 20_000 });
  await expect(page.getByTestId('sim-combo-row').first()).toContainText(bulkCopy.talentsOwn);
});

test('the page has its own heading and its own canonical', async ({ page }) => {
  await page.goto('/sim/talents');
  await expect(page.getByRole('heading', { level: 1 })).toHaveText(bulkCopy.talentsTitle);
  await expect(page.locator('link[rel="canonical"]')).toHaveAttribute('href', /\/sim\/talents$/);
});
