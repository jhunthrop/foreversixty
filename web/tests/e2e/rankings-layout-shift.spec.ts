// web/tests/e2e/rankings-layout-shift.spec.ts
// A board that cannot be read never moves the page (logs spec 2026-11-07 section 13): the failed
// state takes the loading reserve's place at the same height, so the footer and everything under
// the Plate stay put, measured the way Lighthouse's mobile run sees it (360 px, the API out of reach).
import { apiUnreachable, expect, test } from './support/test';
import { readLayoutShift, trackLayoutShift } from './support/cls';
import { signedOut } from './support/selector';

test.use({ viewport: { width: 360, height: 640 } });

test('a board whose read fails shifts nothing', async ({ page }) => {
  await signedOut(page);
  await apiUnreachable(page, '**/v1/rankings/encounters*', '**/v1/rankings?**');
  await trackLayoutShift(page);
  await page.goto('/rankings/warden-kelthas');
  await expect(page.getByTestId('rankings-error')).toBeVisible();
  await page.waitForTimeout(500);
  expect(await readLayoutShift(page)).toBeLessThanOrEqual(0.01);
});

test('a board that loads shifts no more than the budget', async ({ page }) => {
  await signedOut(page);
  const { stubRankingsApi } = await import('./support/rankings-fixture');
  await stubRankingsApi(page);
  await trackLayoutShift(page);
  await page.goto('/rankings/warden-kelthas');
  await expect(page.getByTestId('ranking-1')).toBeVisible();
  await page.waitForTimeout(500);
  expect(await readLayoutShift(page)).toBeLessThanOrEqual(0.05);
});
