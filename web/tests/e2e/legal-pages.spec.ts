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
    // The phone's On this page list (closed) names the Stripe section too: read the document's own text.
    await expect(page.getByText('Stripe', { exact: false }).filter({ visible: true }).first()).toBeVisible();
    await expect(page.getByText('fs_session', { exact: false })).toBeVisible();
  });

  test('the three legal pages link each other in tabs, the current one marked', async ({ page }) => {
    await page.goto('/refunds');
    const tabs = page.getByRole('navigation', { name: 'Legal pages' });
    await expect(tabs.getByRole('link')).toHaveText(['Terms', 'Privacy', 'Refunds']);
    await expect(tabs.getByRole('link', { name: 'Refunds' })).toHaveAttribute('aria-current', 'page');
    await tabs.getByRole('link', { name: 'Privacy' }).click();
    await expect(page).toHaveURL(/\/privacy$/);
    await expect(page.getByRole('heading', { level: 1 })).toHaveText('Privacy policy');
  });

  test('/privacy lists every section in its contents rail, the Log coach section anchored', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 1440, height: 1000 });
    await page.goto('/privacy#log-coach');
    const rail = page.getByTestId('content-toc');
    await expect(rail.getByRole('link')).toHaveCount(8);
    await expect(rail.getByRole('link', { name: 'Log coach' })).toHaveAttribute('href', '#log-coach');
    await expect(page.locator('#log-coach')).toBeInViewport();
    await expect(rail.getByRole('link', { name: 'Log coach' })).toHaveAttribute('aria-current', 'location');
  });

  test('on a phone the contents are one closed line that opens onto the sections and closes on a choice', async ({
    page,
  }) => {
    await page.setViewportSize({ width: 390, height: 844 });
    await page.goto('/privacy');
    const contents = page.getByTestId('content-toc-phone').locator('details');
    await expect(contents).not.toHaveAttribute('open', '');
    await contents.locator('summary').click();
    await contents.getByRole('link', { name: 'Cookies' }).click();
    await expect(contents).not.toHaveAttribute('open', '');
    await expect(page).toHaveURL(/#cookies$/);
  });

  test('/refunds says a period is never refunded in part and who issues a refund', async ({ page }) => {
    await page.goto('/refunds');
    await expect(page.getByText('we do not refund part of a period', { exact: false })).toBeVisible();
    await expect(page.getByText('cancels the subscription', { exact: false })).toBeVisible();
  });
});
