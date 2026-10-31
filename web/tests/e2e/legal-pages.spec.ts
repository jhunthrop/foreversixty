// web/tests/e2e/legal-pages.spec.ts
import { expect, test } from './support/test';

test.describe('legal pages', () => {
  for (const path of ['/terms', '/privacy', '/refunds']) {
    test(`${path} renders final, with no draft banner or OWNER marker`, async ({ page }) => {
      await page.goto(path);
      await expect(page.locator('h1')).toBeVisible();
      await expect(page.getByTestId('legal-draft-banner')).toHaveCount(0);
      await expect(page.getByText('OWNER:', { exact: false })).toHaveCount(0);
      await expect(page.getByText('Commish LLC', { exact: false }).first()).toBeVisible();
      await expect(page.getByRole('link', { name: 'help@foreversixty.gg' }).first()).toBeVisible();
    });
  }

  test('/terms names Commish LLC, Arkansas law, and that Blizzard-sourced data is never sold', async ({
    page,
  }) => {
    await page.goto('/terms');
    await expect(page.getByText('an Arkansas limited liability company', { exact: false })).toBeVisible();
    await expect(page.getByText('laws of the State of Arkansas', { exact: false })).toBeVisible();
    await expect(page.getByText('never sold', { exact: false })).toBeVisible();
  });

  test('/privacy names Stripe as the payment processor and the functional-only cookies', async ({ page }) => {
    await page.goto('/privacy');
    await expect(page.getByText('Stripe', { exact: false }).first()).toBeVisible();
    await expect(page.getByText('fs_session', { exact: false })).toBeVisible();
  });

  test('/refunds says a period is never refunded in part and who issues a refund', async ({ page }) => {
    await page.goto('/refunds');
    await expect(page.getByText('we do not refund part of a period', { exact: false })).toBeVisible();
    await expect(page.getByText('cancels the subscription', { exact: false })).toBeVisible();
  });
});
