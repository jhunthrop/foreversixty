// web/tests/captures/raider-ranks/capture.spec.ts
// Review captures for the Settings tab's raider ranks row (addon hub L9). Not a test of behaviour:
// it drives the guild settings page into each state and writes PNGs to
// design/reviews/2026-11-13-raider-ranks/: the whole page at 1440, the Members and access panel at
// every width, and the unset, saving, saved, ranks-failed and frozen states.
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { test, type Page } from '@playwright/test';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  RAIDER_RANKS_URL,
  stubRaiderRanks,
  type RaiderRanksStub,
} from '../../e2e/support/guild-raider-ranks';

const OUT_DIR = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../../design/reviews/2026-11-13-raider-ranks',
);
const VIEWPORT_HEIGHT = 1000;
const SETTLE_MS = 400;
const WIDTHS = [390, 1024, 1440, 2000];

const target = (name: string): string => path.join(OUT_DIR, name);

async function open(page: Page, width: number, stub: RaiderRanksStub, testId = 'guild-raider-ranks-list') {
  await page.setViewportSize({ width, height: VIEWPORT_HEIGHT });
  await page.route('**/v1/guilds/*/billing', (route) => route.abort());
  await page.route('**/v1/me', (route) => route.abort());
  const bodies = await stubRaiderRanks(page, stub);
  await gotoHydrated(page, RAIDER_RANKS_URL, testId);
  await page.waitForTimeout(SETTLE_MS);
  return bodies;
}

const panel = (page: Page) => page.getByTestId('guild-settings-members');

for (const width of WIDTHS) {
  test(`set at ${width}`, async ({ page }) => {
    await open(page, width, { saved: [2, 3] });
    await panel(page).screenshot({ path: target(`set-panel-${width}.png`) });
    if (width === 1440 || width === 390)
      await page.screenshot({ path: target(`set-page-${width}.png`), fullPage: true });
  });
}

for (const width of [390, 1440]) {
  test(`unset at ${width}`, async ({ page }) => {
    await open(page, width, { saved: null });
    await panel(page).screenshot({ path: target(`unset-panel-${width}.png`) });
  });

  test(`saving and saved at ${width}`, async ({ page }) => {
    let release = (): void => {};
    const hold = new Promise<void>((resolve) => (release = resolve));
    await open(page, width, { saved: [2], hold });
    await page.getByTestId('guild-raider-rank-3-check').check();
    await page.getByTestId('guild-raider-ranks-save').click();
    await page.waitForTimeout(SETTLE_MS);
    await panel(page).screenshot({ path: target(`saving-panel-${width}.png`) });
    release();
    await page.getByTestId('guild-raider-ranks-saved').waitFor();
    await panel(page).screenshot({ path: target(`saved-panel-${width}.png`) });
  });
}

test('cleared at 1440', async ({ page }) => {
  await open(page, 1440, { saved: [2] });
  await page.getByTestId('guild-raider-rank-2-check').uncheck();
  await panel(page).screenshot({ path: target('clearing-panel-1440.png') });
});

test('ranks failed at 1440', async ({ page }) => {
  await open(page, 1440, { saved: [2], setupFails: true }, 'guild-raider-ranks');
  await panel(page).screenshot({ path: target('failed-panel-1440.png') });
});

test('frozen at 1440', async ({ page }) => {
  await open(page, 1440, { saved: [2, 3], frozen: true });
  await panel(page).screenshot({ path: target('frozen-panel-1440.png') });
});
