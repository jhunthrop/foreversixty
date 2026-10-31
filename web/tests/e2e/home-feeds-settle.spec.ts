import type { Page, Route } from '@playwright/test';
import { expect, test } from './support/test';

// Tenet 9: loading must end in a ready, empty or error state. Recent reports and Top guilds
// are below-the-fold islands (client:visible); a read that answers with no rows or fails must
// replace the skeleton within the budget, and a visitor at the fold must never be left on one.
const SETTLE_BUDGET_MS = 1500;

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});
const failure = (route: Route) => route.fulfill({ status: 500, contentType: 'application/json', body: '{}' });

async function stubReads(page: Page, mode: 'empty' | 'failing'): Promise<void> {
  const handler = mode === 'empty' ? (route: Route) => route.fulfill(envelope({ rows: [] })) : failure;
  await page.route('**/v1/reports/recent**', handler);
  await page.route('**/v1/rankings/guilds?**', handler);
}

async function openFeeds(page: Page): Promise<void> {
  await page.goto('/', { waitUntil: 'networkidle' });
  await page.getByTestId('home-around-the-site').scrollIntoViewIfNeeded();
}

test.describe('home feeds settle', () => {
  test('empty reads end in the empty states, no skeleton left', async ({ page }) => {
    await stubReads(page, 'empty');
    await openFeeds(page);
    await expect(page.getByTestId('home-top-guilds-empty')).toBeVisible({ timeout: SETTLE_BUDGET_MS });
    await expect(page.getByTestId('recent-reports-sample')).toBeVisible({ timeout: SETTLE_BUDGET_MS });
    await expect(page.getByTestId('recent-reports-skeleton')).toHaveCount(0);
    await expect(page.getByTestId('home-top-guilds-skeleton')).toHaveCount(0);
  });

  test('failing reads end in the error states with a retry, no skeleton left', async ({ page }) => {
    await stubReads(page, 'failing');
    await openFeeds(page);
    await expect(page.getByTestId('recent-reports-error')).toBeVisible({ timeout: SETTLE_BUDGET_MS });
    await expect(page.getByTestId('home-top-guilds-error')).toBeVisible({ timeout: SETTLE_BUDGET_MS });
    await expect(page.locator('[data-testid$="-skeleton"]')).toHaveCount(0);
  });

  test('a visitor who has not scrolled still sees no skeleton once the page is idle', async ({ page }) => {
    await stubReads(page, 'empty');
    await page.goto('/', { waitUntil: 'networkidle' });
    await page.waitForTimeout(SETTLE_BUDGET_MS);
    await expect(page.locator('[data-testid$="-skeleton"]')).toHaveCount(0);
  });
});
