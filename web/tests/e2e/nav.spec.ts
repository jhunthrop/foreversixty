// web/tests/e2e/nav.spec.ts
// The primary nav's behaviours after the cut (spec 2026-09-25): five doors plus Get set up,
// no Reference disclosure, a fixed two-row phone grid instead of a horizontally scrolling
// row, and "Sign in" on every page.
import { expect, test } from '@playwright/test';

test.describe('aria-current', () => {
  test('the Simulator tool link carries aria-current=page on /sim', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByTestId('primary-nav').locator('a', { hasText: 'Simulator' })).toHaveAttribute(
      'aria-current',
      'page',
    );
  });

  test('the Guides link carries aria-current=page on a guide sub-path', async ({ page }) => {
    await page.goto('/guides/warrior');
    await expect(page.getByTestId('primary-nav').locator('a', { hasText: 'Guides' })).toHaveAttribute(
      'aria-current',
      'page',
    );
  });
});

test.describe('desktop header', () => {
  test.use({ viewport: { width: 1440, height: 900 } });
  test.skip(() => test.info().project.name !== 'desktop', 'desktop layout only');

  test('the wordmark, the doors and the session sit left to right, with no Menu button', async ({ page }) => {
    await page.goto('/');
    const mark = await page.getByRole('banner').getByRole('link', { name: 'Forever Sixty' }).boundingBox();
    const nav = await page.getByTestId('primary-nav').boundingBox();
    const discord = await page.getByRole('banner').getByRole('link', { name: 'Discord' }).boundingBox();
    expect(mark!.x + mark!.width).toBeLessThanOrEqual(nav!.x);
    expect(nav!.x + nav!.width).toBeLessThanOrEqual(discord!.x);
    await expect(page.getByTestId('menu-button')).toBeHidden();
  });
});

test.describe('phone nav', () => {
  test.use({ viewport: { width: 360, height: 800 } });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the closed header is one 56px bar and the menu opens the doors beneath it', async ({ page }) => {
    await page.goto('/');
    const header = page.getByRole('banner');
    const nav = page.getByTestId('primary-nav');
    const button = page.getByTestId('menu-button');
    await expect(nav).toBeHidden();
    await expect(button).toHaveAttribute('aria-expanded', 'false');
    expect((await header.boundingBox())!.height).toBeLessThanOrEqual(60);
    await button.click();
    await expect(nav).toBeVisible();
    await expect(button).toHaveAttribute('aria-expanded', 'true');
    await page.keyboard.press('Escape');
    await expect(nav).toBeHidden();
    await expect(button).toBeFocused();
  });

  test('all eight items are visible without horizontal scroll', async ({ page }) => {
    await page.goto('/');
    await page.getByTestId('menu-button').click();
    const nav = page.getByTestId('primary-nav');
    const { scrollWidth, clientWidth } = await nav.evaluate((element) => ({
      scrollWidth: element.scrollWidth,
      clientWidth: element.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth + 1);
    for (const label of [
      'Planner',
      'Simulator',
      'Logs',
      'Rankings',
      'Tier List',
      'Guides',
      'Leveling BiS',
      'Get set up',
    ]) {
      await expect(nav.getByRole('link', { name: label })).toBeVisible();
    }
  });

  test('the page does not scroll sideways with the eight-item nav', async ({ page }) => {
    await page.goto('/');
    const { scrollWidth, clientWidth } = await page.evaluate(() => ({
      scrollWidth: document.documentElement.scrollWidth,
      clientWidth: document.documentElement.clientWidth,
    }));
    expect(scrollWidth).toBeLessThanOrEqual(clientWidth);
  });

  // Fix round (night-site-ux, 2026-09-28): "Leveling BiS" (98.75px of tracked, uppercase
  // text) did not fit its 80px column in this four-up grid and, forced onto one line,
  // spilled ~10px into "Get set up" next door -- the two links visually ran together as
  // "LEVELING BISGET SET UP" with no gap between them, even though each link's own tap
  // target box was still correctly positioned and separately clickable. A visibility check
  // alone can't catch this (both links are still individually "visible"), so this compares
  // actual rendered boxes directly.
  test('no two nav links visually overlap or touch', async ({ page }) => {
    await page.goto('/');
    await page.getByTestId('menu-button').click();
    const nav = page.getByTestId('primary-nav');
    const boxes = await nav.locator('a').evaluateAll((links) =>
      links.map((link) => {
        const r = link.getBoundingClientRect();
        return { left: r.left, right: r.right, top: r.top, bottom: r.bottom };
      }),
    );
    for (let i = 0; i < boxes.length; i++) {
      for (let j = i + 1; j < boxes.length; j++) {
        const a = boxes[i];
        const b = boxes[j];
        const overlapsHorizontally = a.left < b.right && b.left < a.right;
        const overlapsVertically = a.top < b.bottom && b.top < a.bottom;
        expect(overlapsHorizontally && overlapsVertically, `nav links ${i} and ${j} overlap`).toBe(false);
      }
    }
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

  test('a content page gets the same account menu island as the tool pages', async ({ page }) => {
    await page.goto('/guides');
    await expect(page.locator('astro-island[component-url*="AccountMenu"]')).toHaveCount(1);
    await expect(page.getByRole('banner').getByRole('link', { name: 'Sign in' })).toHaveAttribute(
      'href',
      /\/login/,
    );
  });
});
