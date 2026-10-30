// web/tests/e2e/select-captures.spec.ts
// The dropdown lane's review captures (controls and charts spec §D), written to
// design/reviews/2026-10-30-controls-charts/build-dropdown-<surface>-<width>.png. Skipped unless
// DROPDOWN_CAPTURES=1: they are pictures for the reviewers, not assertions. On a fine pointer each
// surface shows its list open; at 390 (a coarse pointer) the native field, or the legend sheet.
//   DROPDOWN_CAPTURES=1 npx playwright test select-captures.spec.ts --project=desktop
import { readFileSync } from 'node:fs';
import { join } from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { beLightningbolltz, RAID40_URL, serveRaid40 } from './support/raid40';
import { enhanced, listOf } from './support/select';

const OUT = join(import.meta.dirname, '../../../design/reviews/2026-10-30-controls-charts');
const { build } = JSON.parse(
  readFileSync(join(import.meta.dirname, '..', '..', 'src', 'data', 'active-build.json'), 'utf8'),
) as { build: string };
/** sim-settings.spec.ts's Fury export: the sim draws its settings once a character runs. */
const FURY = `FS1:${build}:warrior:orc:0/5530515/0:head=12640,main_hand=11726`;

test.skip(process.env.DROPDOWN_CAPTURES !== '1', 'review captures run on request');

async function open(page: Page, testid: string, phone: boolean): Promise<void> {
  if (phone) return;
  const trigger = await enhanced(page, testid);
  await trigger.scrollIntoViewIfNeeded();
  await page.evaluate(() => window.scrollBy(0, -140));
  await trigger.click();
  await expect(listOf(page, testid)).toBeVisible();
}

async function loadFury(page: Page): Promise<void> {
  await page.getByTestId('sim-cold-paste-input').fill(FURY);
  await page.getByTestId('sim-cold-paste-run').click();
  await expect(page.getByTestId('sim-run-button')).toHaveText('Run again', { timeout: 20_000 });
}

async function shot(page: Page, surface: string, width: number, fullPage = false): Promise<void> {
  await page.waitForTimeout(300);
  await page.screenshot({ path: join(OUT, `build-dropdown-${surface}-${width}.png`), fullPage });
}

for (const width of [1024, 1440, 2000, 390] as const) {
  const phone = width < 768;
  test.describe(`${width}`, () => {
    test.use({
      viewport: { width, height: phone ? 844 : 1000 },
      hasTouch: phone,
      isMobile: phone,
      reducedMotion: 'reduce',
    });

    test(`rankings ${width}`, async ({ page }) => {
      await page.goto('/rankings');
      await page.getByTestId('rankings-filters').waitFor();
      if (phone) await page.getByTestId('rankings-filters-toggle').click();
      else await open(page, 'filter-class', phone);
      await shot(page, 'rankings', width);
    });

    test(`report source with 40 players ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the report captures are 1440, 2000 and 390');
      await serveRaid40(page);
      await beLightningbolltz(page);
      await page.goto(`${RAID40_URL}&tab=damage-done`);
      await page.getByTestId('source-scope').waitFor({ state: 'attached' });
      await open(page, 'source-scope', phone);
      await shot(page, 'report-source-40', width);
    });

    test(`report source, the fixture's five ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the report captures are 1440, 2000 and 390');
      await page.goto('/reports/fixture2abcd?fight=3&tab=threat');
      await page.getByTestId('source-scope').waitFor({ state: 'attached' });
      await open(page, 'source-scope', phone);
      await shot(page, 'report-source', width);
    });

    test(`players ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the report captures are 1440, 2000 and 390');
      await serveRaid40(page);
      await beLightningbolltz(page);
      await page.goto(`${RAID40_URL}&tab=damage-done`);
      if (phone) await page.getByTestId('chart-toggle').click();
      await expect(page.locator('[data-state="ready"]').first()).toBeVisible();
      if (phone) await page.getByTestId('time-chart').getByTestId('legend-more').click();
      else await open(page, 'players', phone);
      await shot(page, 'players', width);
    });

    test(`report more and chart fields ${width}`, async ({ page }) => {
      test.skip(width === 1024 || phone, 'the desktop tab row');
      await page.goto('/reports/fixture2abcd?fight=3&tab=threat');
      await page.getByTestId('tab-more').click();
      await page.getByTestId('tab-more-menu').waitFor();
      await shot(page, 'report-more', width);
    });

    test(`sim item search ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the sim captures are 1440, 2000 and 390');
      await page.goto('/sim/gear');
      await page.getByTestId('sim-addon-input').fill(FURY);
      await page.getByTestId('sim-addon-load').click();
      await page.getByTestId('sim-search-source').waitFor({ state: 'attached' });
      await open(page, 'sim-search-source', phone);
      if (!phone) await page.keyboard.type('black');
      await shot(page, 'sim-search-black', width);
      if (phone) return;
      await page.keyboard.press('Control+A');
      await page.keyboard.type('naxx');
      await shot(page, 'no-match', width);
    });

    test(`sim fight style ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the sim captures are 1440, 2000 and 390');
      await page.goto('/sim');
      await loadFury(page);
      await open(page, 'sim-style', phone);
      await shot(page, 'sim-style', width);
    });

    test(`flip above ${width}`, async ({ page }) => {
      test.skip(width === 1024 || phone, 'the list is the fine pointer');
      await page.goto('/design-sheet');
      const trigger = await enhanced(page, 'sheet-enemy');
      await trigger.evaluate((node) =>
        window.scrollBy(0, node.getBoundingClientRect().bottom - window.innerHeight + 60),
      );
      await trigger.click();
      await expect(listOf(page, 'sheet-enemy')).toHaveAttribute('data-side', 'top');
      await shot(page, 'flip', width);
    });

    test(`sheet ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the sheet captures are 1440, 2000 and 390');
      await page.goto('/design-sheet');
      await page.getByTestId('sheet-dialog-open').waitFor();
      await shot(page, 'sheet', width, true);
    });

    test(`home search ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the home captures are 1440, 2000 and 390');
      await page.goto('/');
      await page.locator('#home-character-search-region').waitFor({ state: 'attached' });
      if (!phone) {
        const trigger = page.locator('#home-character-search-region-trigger');
        await trigger.waitFor();
        await trigger.scrollIntoViewIfNeeded();
        await trigger.click();
      }
      await shot(page, 'home-search', width);
    });

    test(`planner ${width}`, async ({ page }) => {
      test.skip(width === 1024, 'the planner captures are 1440, 2000 and 390');
      await page.goto('/planner');
      await page.getByTestId('planner-race-select').waitFor({ state: 'attached' });
      await open(page, 'planner-race-select', phone);
      await shot(page, 'planner', width);
    });
  });
}
