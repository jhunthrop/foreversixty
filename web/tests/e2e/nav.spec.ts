// web/tests/e2e/nav.spec.ts
// The primary nav's own behaviours: the Reference disclosure (mouse and keyboard), aria-
// current, the phone-only right-edge fade and tool-first order, "Sign in" on every page,
// and the "/" shortcut. web/tests/e2e/layout.spec.ts keeps the general phone-fit sweep and
// the pre-existing Discord/search hit-target check; this file is the nav's own behaviour,
// which is cohesive enough to want its own file rather than growing that one further.
import { expect, test } from '@playwright/test';

test.describe('Reference disclosure', () => {
  test('opens and closes by mouse click, with no layout shift to the header', async ({ page }) => {
    await page.goto('/');
    const header = page.locator('header');
    const before = await header.boundingBox();
    const summary = page.locator('summary');
    await summary.click();
    await expect(page.getByRole('link', { name: 'Classes' })).toBeVisible();
    const afterOpen = await header.boundingBox();
    expect(afterOpen?.height).toBe(before?.height);
    await summary.click();
    await expect(page.getByRole('link', { name: 'Classes' })).toBeHidden();
  });

  test('opens by keyboard: Tab to the summary, Enter toggles it', async ({ page }) => {
    await page.goto('/');
    // Scoped to the primary nav: the homepage also has a "Live Rankings" content card whose
    // accessible name contains "Rankings", so an unscoped role query is ambiguous here.
    await page.getByTestId('primary-nav').getByRole('link', { name: 'Rankings' }).focus();
    await page.keyboard.press('Tab');
    await expect(page.locator('summary')).toBeFocused();
    await page.keyboard.press('Enter');
    // Scoped for the same reason as the Rankings focus above: the homepage's own "Live
    // Class guides" content card also has an accessible name containing "Guides".
    await expect(page.getByTestId('primary-nav').getByRole('link', { name: 'Guides' })).toBeVisible();
  });
});

test.describe('aria-current', () => {
  test('the Reference summary carries aria-current=page on a reference page', async ({ page }) => {
    await page.goto('/classes');
    await expect(page.locator('summary')).toHaveAttribute('aria-current', 'page');
  });

  test('a reference child link carries aria-current=page when it is the current page', async ({ page }) => {
    await page.goto('/classes');
    await page.locator('summary').click();
    await expect(page.getByRole('link', { name: 'Classes' })).toHaveAttribute('aria-current', 'page');
    await expect(page.getByRole('link', { name: 'Guides' })).not.toHaveAttribute('aria-current', 'page');
  });

  test('the Simulator tool link carries aria-current=page on /sim', async ({ page }) => {
    await page.goto('/sim');
    await expect(page.getByRole('link', { name: 'Simulator' })).toHaveAttribute('aria-current', 'page');
  });
});

test.describe('phone nav', () => {
  test.use({ viewport: { width: 360, height: 800 } });
  test.skip(() => test.info().project.name !== 'mobile', 'phone layout only');

  test('the four tools come before Reference and The addon in scroll order', async ({ page }) => {
    await page.goto('/');
    const nav = page.getByTestId('primary-nav');
    // Scoped to the nav's own top-level items (direct-child anchors, plus the Reference
    // <summary>) rather than `a, summary` unscoped: the Reference panel's own Classes /
    // Guides / Zones / Dungeons links are `<a>` elements too, nested inside the closed
    // `<details>`, and an unscoped selector picks them up in DOM order right after
    // "Reference" -- a false failure, not a real one, since they are never part of the
    // top-level scroll order this test is about.
    const labels = await nav.locator(':scope > a, :scope > details > summary').allTextContents();
    const trimmed = labels.map((label) => label.trim());
    expect(trimmed.slice(0, 4)).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
    expect(trimmed.slice(4)).toEqual(['Reference', 'The addon']);
  });

  test('the nav row scrolls and carries the right-edge fade background while more is reachable', async ({
    page,
  }) => {
    await page.goto('/');
    const nav = page.getByTestId('primary-nav');
    const { scrollWidth, clientWidth, backgroundImage } = await nav.evaluate((element) => ({
      scrollWidth: element.scrollWidth,
      clientWidth: element.clientWidth,
      backgroundImage: getComputedStyle(element).backgroundImage,
    }));
    expect(scrollWidth).toBeGreaterThan(clientWidth);
    expect(backgroundImage).not.toBe('none');
  });
});

test.describe('sign in on every page', () => {
  for (const path of ['/', '/planner', '/sim', '/classes', '/guides', '/zones']) {
    test(`${path} shows Sign in in the header`, async ({ page }) => {
      await page.route('**/v1/me', (route) =>
        route.fulfill({
          status: 401,
          contentType: 'application/json',
          body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
        }),
      );
      await page.goto(path);
      await expect(page.getByTestId('session-nav').getByRole('link', { name: 'Sign in' })).toBeVisible();
    });
  }
});

test.describe('"/" shortcut', () => {
  test('on a page with a search box, "/" focuses it', async ({ page }) => {
    await page.goto('/');
    await page.keyboard.press('/');
    await expect(page.getByLabel('Search the site')).toBeFocused();
  });

  test('on a page without a search box, "/" goes to /search', async ({ page }) => {
    await page.goto('/classes');
    await page.keyboard.press('/');
    await expect(page).toHaveURL(/\/search$/);
  });

  test('ignored while typing in a field', async ({ page }) => {
    await page.goto('/search');
    const box = page.getByLabel('Search the site');
    await box.click();
    await box.pressSequentially('a/b');
    await expect(box).toHaveValue('a/b');
  });
});

// The panel was `position: fixed` at a measured 137px. It looked right at the top of a page
// and floated over the content, detached from the header, as soon as the page scrolled.
test('the Reference panel sits under the header and stays attached to it when the page scrolls', async ({
  page,
}) => {
  await page.goto('/classes');
  const summary = page.getByTestId('primary-nav').locator('summary');
  await summary.scrollIntoViewIfNeeded();
  await summary.click();
  const panel = page.locator('.reference-panel');
  await expect(panel).toBeVisible();

  const gap = (): Promise<number> =>
    page.evaluate(() => {
      const header = document.querySelector('header')!.getBoundingClientRect();
      const box = document.querySelector('.reference-panel')!.getBoundingClientRect();
      return Math.round(box.top - header.bottom);
    });
  const before = await gap();
  // Under the header (a few pixels of margin on desktop), never over it and never far below.
  expect(before).toBeGreaterThanOrEqual(-60);
  expect(before).toBeLessThanOrEqual(12);

  await page.evaluate(() => window.scrollTo(0, 80));
  expect(await gap()).toBe(before);
  // All four links are reachable, not clipped by the scrolling nav row.
  for (const name of ['Classes', 'Guides', 'Zones', 'Dungeons']) {
    await expect(panel.getByRole('link', { name })).toBeVisible();
  }
});
