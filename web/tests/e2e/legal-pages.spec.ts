// web/tests/e2e/legal-pages.spec.ts
import { expect, test } from '@playwright/test';

test.describe('legal pages', () => {
  for (const path of ['/terms', '/privacy', '/refunds']) {
    test(`${path} renders and shows the unfinished-draft banner while OWNER markers remain`, async ({
      page,
    }) => {
      await page.goto(path);
      await expect(page.locator('h1')).toBeVisible();
      await expect(page.getByText('OWNER:', { exact: false }).first()).toBeVisible();
      await expect(page.getByTestId('legal-draft-banner')).toBeVisible();
    });
  }

  test('/terms names COMMISH LLC and that Blizzard-sourced data is never sold', async ({ page }) => {
    await page.goto('/terms');
    await expect(page.getByText('COMMISH LLC', { exact: false })).toBeVisible();
    await expect(page.getByText('never sold', { exact: false })).toBeVisible();
  });

  test('/privacy names Stripe as the payment processor and the functional-only cookies', async ({ page }) => {
    await page.goto('/privacy');
    await expect(page.getByText('Stripe', { exact: false }).first()).toBeVisible();
    await expect(page.getByText('fs_session', { exact: false })).toBeVisible();
  });

  test('/refunds explains the manual cancel-with-refund process honestly', async ({ page }) => {
    await page.goto('/refunds');
    await expect(page.getByText('cancels the subscription', { exact: false })).toBeVisible();
  });
});
