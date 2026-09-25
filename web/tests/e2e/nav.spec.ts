// web/tests/e2e/nav.spec.ts
// The primary nav's behaviours after the cut (spec 2026-09-25): five doors plus Get set up,
// no Reference disclosure, a fixed two-row phone grid instead of a horizontally scrolling
// row, and "Sign in" on every page.
import { expect, test } from '@playwright/test';

test.describe('aria-current', () => {
  test('the Simulator tool link carries aria-current=page on /sim', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('primary-nav').getByRole('link', { name: 'Simulator' })).toHaveAttribute(
      'aria-current',
      'page',
    );
  });

  test('the Guides link carries aria-current=page on a guide sub-path', async ({ page }) => {
    await page.goto('/guides/warrior');
    await expect(page.getByTestId('primary-nav').getByRole('link', { name: 'Guides' })).toHaveAttribute(
      'aria-current',
      'page',
    );
  });
});

test.describe('phone nav', () => {
  test.use({ viewport: { width: 360, height: 800 } });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('all six items are visible without horizontal scroll', async ({ page }) => {
    await page.goto('/');
    const nav = page.getByTestId('primary-nav');
    const { scrollWidth, clientWidth } = await nav.evaluate((element) => ({
      scrollWidth: element.scrollWidth,
      clientWidth: element.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth + 1);
    for (const label of ['Planner', 'Simulator', 'Logs', 'Rankings', 'Guides', 'Get set up']) {
      await expect(nav.getByRole('link', { name: label })).toBeVisible();
    }
  });

  test('the page does not scroll sideways with the six-item nav', async ({ page }) => {
    await page.goto('/');
    const { scrollWidth, clientWidth } = await page.evaluate(() => ({
      scrollWidth: document.documentElement.scrollWidth,
      clientWidth: document.documentElement.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
  });
});

test.describe('sign in on every page', () => {
  for (const path of ['/', '/planner', '/sim', '/guides', '/setup']) {
    test(`${path} shows Sign in in the header`, async ({ page }) => {
      await page.route('**/v1/me', (route) =>
        route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        }),
      );
      await page.goto(path);
      await expect(page.getByRole('banner').getByRole('link', { name: 'Sign in' })).toBeVisible();
    });
  }

  test('a content page gets its Sign in link without shipping a script for it', async ({ page }) => {
    const scripts: string[] = [];
    page.on('request', (request) => {
      if (request.resourceType() === 'script') scripts.push(request.url());
    });
    await page.goto('/guides');
    await expect(page.getByTestId('session-nav-static')).toHaveAttribute('href', '/login');
    expect(scripts).toEqual([]);
  });
});
