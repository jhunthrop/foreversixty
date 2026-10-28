// web/tests/captures/launch-pricing/capture-premium-value.spec.ts
// Capture rows 15 to 22 (design/specs/2026-10-27-launch-pricing.md Amendment 2026-10-27b, B9): the
// plans page's Region A and Region B, the band link, the failed example image, the /sim landing
// card, and the full page beside the premium-value board. PNGs go to
// design/reviews/2026-10-27-launch-pricing/. Row 18 also prints the founder-line wrap check and
// the CLS of the example card and the 54 px cost-line box into cls-premium-value-<width>.json.
// Same run as capture.spec.ts (see playwright.capture.config.ts).
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { expect, test } from '@playwright/test';
import { FOUNDER_WINDOW_CLOSES } from '../../../src/lib/billing/plans';
import { plansPageCopy } from '../../../src/lib/billing/plans-page-copy';
import { readLayoutShift, trackLayoutShift } from '../../e2e/support/cls';
import { COST_BOX_PX, founderLineBoxes, guildPriceBoxes } from '../../e2e/support/premium-layout';
import { MOCK_DIR, atWidths, cropOf, fresh, settle, sideBySide, sizeTo } from '../support/capture-kit';
import {
  DECEMBER,
  clearStorageOnLoad,
  file,
  openPlans,
  overlayDiff,
  shoot,
  target,
} from './plans-capture-kit';

const ALL = [360, 390, 1024, 1440, 2000];
const PAIR = [390, 1440];
const REGION_B_WIDTHS = [360, 390, 1024, 1440];
const SIM_WIDTHS = [360, 390, 1440];
const BOARD_WIDTHS = [1024, 1440, 2000];
const BOARD_FILE: Record<number, string> = {
  1024: 'premium-value-1024.html',
  1440: 'premium-value.html',
  2000: 'premium-value-2000.html',
};
const FOUNDER_LINE_WIDTHS = [360, 1024];
const CLS_CAP = 0.05;

atWidths('15 region A, founder open, Yearly', ALL, async (page, width) => {
  await openPlans(page, width, DECEMBER);
  await shoot(page, 15, 'region-a', width, page.locator('[data-region="a"]'));
});

atWidths('16 region A, Monthly, diffed against Yearly', PAIR, async (page, width) => {
  await openPlans(page, width, DECEMBER);
  const yearlyPng = file(16, 'yearly', width);
  await page.screenshot({ path: yearlyPng, fullPage: true });
  await page.getByLabel('Monthly').check();
  await settle(page);
  const monthlyPng = file(16, 'monthly', width);
  await page.screenshot({ path: monthlyPng, fullPage: true });
  await overlayDiff(await fresh(page), yearlyPng, monthlyPng, file(16, 'diff', width), width);
});

atWidths('17 region B, founder open, Yearly', ALL, async (page, width) => {
  await openPlans(page, width, DECEMBER);
  await shoot(page, 17, 'region-b', width, page.locator('#guild-leaders'));
});

atWidths('18 region B, Monthly and after the founder window', REGION_B_WIDTHS, async (page, width) => {
  await trackLayoutShift(page);
  await openPlans(page, width, DECEMBER);
  await page.getByLabel('Monthly').check();
  await settle(page);
  await shoot(page, 18, 'monthly', width, page.locator('#guild-leaders'));
  const closed = await fresh(page);
  await openPlans(closed, width, FOUNDER_WINDOW_CLOSES);
  await shoot(closed, 18, 'after-window-yearly', width, closed.locator('#guild-leaders'));
  await closed.getByLabel('Monthly').check();
  await settle(closed);
  await shoot(closed, 18, 'after-window-monthly', width, closed.locator('#guild-leaders'));

  // The cost line fits its 54 px box in every form (B8 check 4).
  for (const on of [page, closed]) {
    for (const box of await guildPriceBoxes(on)) {
      expect(box.costBox, `${box.plan} cost box`).toBe(COST_BOX_PX);
      expect(box.costOverflows, `${box.plan} cost line overflows`).toBe(false);
    }
  }
  // The founder line must stay on one line, or PriceBlock's fixed 144 px block breaks (PV review).
  const founderLines = FOUNDER_LINE_WIDTHS.includes(width) ? await founderLineBoxes(page) : [];
  for (const line of founderLines) expect(line.lines, `founder line wraps at ${width}: ${line.text}`).toBe(1);
  const cls = Math.round((await readLayoutShift(page)) * 10_000) / 10_000;
  writeFileSync(
    target(`cls-premium-value-${width}.json`),
    `${JSON.stringify({ width, cls, cap: CLS_CAP, founderLines }, null, 2)}\n`,
  );
  expect(cls).toBeLessThanOrEqual(CLS_CAP);
});

atWidths('19 /sim landing with the real example card', SIM_WIDTHS, async (page, width) => {
  await sizeTo(page, width);
  await clearStorageOnLoad(page);
  await page.goto('/sim');
  await settle(page);
  await shoot(page, 19, 'sim-landing', width, page.getByTestId('sim-example-result').first());
});

test('20 band link: default, hover, focus-visible, activated @1440', async ({ page }) => {
  await openPlans(page, 1440, DECEMBER);
  const link = page.getByTestId('band-guild-link');
  const band = page.locator('.band');
  await cropOf(page, band, file(20, 'default', 1440, '-crop'));
  await link.hover();
  await settle(page);
  await cropOf(page, band, file(20, 'hover', 1440, '-crop'));
  await page.mouse.move(0, 0);
  await page.keyboard.press('Shift');
  await link.focus();
  await settle(page);
  await cropOf(page, band, file(20, 'focus-visible', 1440, '-crop'));
  await link.click();
  await settle(page);
  await expect(
    page.getByRole('heading', { level: 2, name: plansPageCopy.guildRegion.heading }),
  ).toBeInViewport();
  await page.screenshot({ path: file(20, 'activated', 1440) });
});

atWidths('21 example image blocked', PAIR, async (page, width) => {
  await page.route('**/premium/example-server-run.*', (route) => route.abort());
  await openPlans(page, width, DECEMBER);
  await shoot(page, 21, 'image-failed', width, page.getByTestId('premium-example'));
});

for (const width of BOARD_WIDTHS) {
  test(`22 plans side by side with the premium-value board @${width}`, async ({ browser }) => {
    const context = await browser.newContext({ reducedMotion: 'reduce', timezoneId: 'America/Chicago' });
    const mockPage = await context.newPage();
    await sizeTo(mockPage, width);
    await mockPage.goto(`file://${path.join(MOCK_DIR, BOARD_FILE[width]!)}`);
    await mockPage.waitForLoadState('networkidle');
    const mockPng = file(22, 'mock', width);
    await mockPage.screenshot({ path: mockPng, fullPage: true });
    const built = await context.newPage();
    await openPlans(built, width, DECEMBER);
    const builtPng = file(22, 'built', width);
    await built.screenshot({ path: builtPng, fullPage: true });
    await sideBySide(await context.newPage(), mockPng, builtPng, file(22, 'side-by-side', width), width);
    await context.close();
  });
}
