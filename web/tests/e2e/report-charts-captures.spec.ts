// web/tests/e2e/report-charts-captures.spec.ts
// The chart lane's review captures (controls and charts spec §D), written to
// design/reviews/2026-10-30-controls-charts/build-charts-<surface>-<width>.png. Skipped unless
// CHART_CAPTURES=1: they are pictures for the reviewers, not assertions.
//   CHART_CAPTURES=1 npx playwright test report-charts-captures.spec.ts --project=desktop
import { join } from 'node:path';
import { expect, test, type Page } from '@playwright/test';
import { beLightningbolltz, RAID40_URL, serveRaid40 } from './support/raid40';

const OUT = join(import.meta.dirname, '../../../design/reviews/2026-10-30-controls-charts');
const WIDTHS = [1440, 2000, 390] as const;

test.skip(process.env.CHART_CAPTURES !== '1', 'review captures run on request');

async function settle(page: Page): Promise<void> {
  await expect(page.locator('[data-state="ready"]').first()).toBeVisible();
  await page.waitForTimeout(400);
}

async function hoverPlot(page: Page, testid: string, fraction: number): Promise<void> {
  const plot = page.getByTestId(testid).getByTestId('chart-plot');
  await plot.scrollIntoViewIfNeeded();
  const box = (await plot.boundingBox())!;
  await page.mouse.move(box.x + 52 + (box.width - 64) * fraction, box.y + box.height / 3);
  await page.waitForTimeout(200);
}

for (const width of WIDTHS) {
  const phone = width < 768;
  test.describe(`${width}`, () => {
    test.use({
      viewport: { width, height: phone ? 640 : 1100 },
      hasTouch: phone,
      isMobile: phone,
    });

    test(`threat ${width}`, async ({ page }) => {
      await serveRaid40(page);
      await page.goto(`${RAID40_URL}&tab=threat&start=90000&end=160000`);
      await settle(page);
      if (phone) {
        await page.goto(`${RAID40_URL}&tab=damage-done`);
        await page.locator('#report-tab-select').selectOption('threat', { force: true });
        await settle(page);
        await page.waitForTimeout(800);
      } else {
        await page.getByTestId('threat-chart').scrollIntoViewIfNeeded();
        await hoverPlot(page, 'threat-chart', 0.53);
      }
      await page.screenshot({ path: join(OUT, `build-charts-threat-${width}.png`), fullPage: !phone });
    });

    test(`damage ${width}`, async ({ page }) => {
      await serveRaid40(page);
      await beLightningbolltz(page);
      await page.goto(`${RAID40_URL}&tab=damage-done`);
      if (phone) await page.getByTestId('chart-toggle').click();
      await settle(page);
      await page.getByTestId('time-chart').scrollIntoViewIfNeeded();
      if (!phone) {
        await page.getByTestId('time-chart').getByTestId('legend-chip').nth(4).hover();
        await page.waitForTimeout(200);
      }
      await page.screenshot({ path: join(OUT, `build-charts-damage-${width}.png`) });
      if (!phone) {
        await hoverPlot(page, 'time-chart', 0.37);
        await page.screenshot({ path: join(OUT, `build-charts-damage-tooltip-${width}.png`) });
      }
    });

    test(`players ${width}`, async ({ page }) => {
      await serveRaid40(page);
      await beLightningbolltz(page);
      await page.goto(`${RAID40_URL}&tab=damage-done`);
      if (phone) await page.getByTestId('chart-toggle').click();
      await settle(page);
      await page.getByTestId('time-chart').scrollIntoViewIfNeeded();
      if (phone) await page.getByTestId('time-chart').getByTestId('legend-more').click();
      else await page.getByTestId('players-trigger').click();
      await page.waitForTimeout(300);
      await page.screenshot({ path: join(OUT, `build-charts-players-${width}.png`) });
    });

    test(`timelines and resources ${width}`, async ({ page }) => {
      // Resources draws no damage chart above its own lines (fix round 2).
      await page.goto('/reports/fixture2abcd?fight=3&tab=resources');
      await expect(page.getByTestId('resource-head')).toBeVisible();
      await page.screenshot({ path: join(OUT, `build-charts-resources-${width}.png`), fullPage: true });
      await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
      await expect(page.getByTestId('timelines')).toBeVisible();
      await page.screenshot({ path: join(OUT, `build-charts-timelines-${width}.png`), fullPage: true });
    });
  });
}

test.describe('death link', () => {
  test.use({ viewport: { width: 1440, height: 1100 } });

  test('death marker to the Deaths tab, before and after (1440)', async ({ page }) => {
    await page.goto('/reports/fixture2abcd?fight=3&tab=damage-done');
    await settle(page);
    const marker = page.getByTestId('time-chart').getByTestId('death-marker');
    await marker.hover();
    await page.waitForTimeout(300);
    await page.screenshot({ path: join(OUT, 'build-charts-death-link-before-1440.png') });
    await marker.click();
    await expect(page).toHaveURL(/tab=deaths/);
    await page.waitForTimeout(500);
    await page.screenshot({ path: join(OUT, 'build-charts-death-link-after-1440.png') });
  });
});
