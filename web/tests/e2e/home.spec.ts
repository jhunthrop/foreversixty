import { expect, test } from '@playwright/test';
import { collectPageErrors } from './support/console';

test('homepage states the product and stops, not a marketing slogan', async ({ page }) => {
  const errors = collectPageErrors(page);
  await page.goto('/');
  const h1 = page.locator('h1');
  await expect(h1).toHaveCount(1);
  const heading = (await h1.innerText()).trim();
  // A reference heading, not a slogan: short, and it never ends with "!" or "?" -- a
  // trailing "." is fine (spec 2026-09-23 §2's own headline sentence, reference voice
  // "state the thing and stop").
  expect(heading.length).toBeLessThanOrEqual(90);
  expect(heading).not.toMatch(/[!?]$/);
  expect(heading).toBe('Your character, planned, simmed, logged and ranked.');
  await expect(page.getByRole('combobox', { name: 'Search the site' })).toBeVisible();
  await expect(page.getByText('Right now')).toBeVisible();
  await expect(page.getByTestId('home-product-planner')).toBeVisible();
  await expect(page.getByTestId('home-product-simulator')).toBeVisible();
  await expect(page.getByTestId('home-product-logs')).toBeVisible();
  await expect(page.getByTestId('home-product-rankings')).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Your guild' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'What changed' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Still unknown' })).toBeVisible();
  await expect(page.getByText('Not affiliated with or endorsed by Blizzard Entertainment')).toBeVisible();
  expect(errors).toEqual([]);
});

test('the four product panels link to their tools, with the planner and simulator live elements', async ({
  page,
}) => {
  await page.goto('/');
  await expect(
    page.getByTestId('home-product-planner').getByRole('link', { name: 'Open the planner' }),
  ).toHaveAttribute('href', '/planner');
  await expect(
    page.getByTestId('home-product-simulator').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
  await expect(
    page.getByTestId('home-product-logs').getByRole('link', { name: 'Open the logs' }),
  ).toHaveAttribute('href', '/logs');
  await expect(
    page.getByTestId('home-product-rankings').getByRole('link', { name: 'Open the rankings' }),
  ).toHaveAttribute('href', '/rankings');
  // The planner panel's live element: nine class tiles linking into the planner.
  await expect(
    page.getByTestId('home-product-planner').getByRole('link', { name: 'Warrior' }),
  ).toHaveAttribute('href', '/planner?class=warrior');
  // The simulator panel's live element: damage-spec pills, real coverage, linking into the
  // simulator itself -- plain /sim, since /sim has no per-spec URL state to preselect from
  // (a per-spec query string would look distinct while landing on an identical page).
  await expect(
    page
      .getByTestId('home-product-simulator')
      .getByRole('link', { name: /Warrior/ })
      .first(),
  ).toHaveAttribute('href', '/sim');
});

// The reference tiles' and addon/companion row's own ordering and hrefs are
// handoffs-home.spec.ts's job; the assertions above already cover this page's presence.

test('content pages ship no client JavaScript', async ({ page }) => {
  for (const path of ['/about', '/premium']) {
    const scripts: string[] = [];
    page.on('request', (r) => {
      if (r.resourceType() === 'script') scripts.push(r.url());
    });
    await page.goto(path);
    expect(scripts).toEqual([]);
    page.removeAllListeners('request');
  }
});

test('a keyboard user can skip the header straight to the content', async ({ page }) => {
  await page.goto('/');
  const skip = page.getByRole('link', { name: 'Skip to content' });
  const offScreen = await skip.boundingBox();
  expect(offScreen?.height ?? 0).toBeLessThanOrEqual(1);

  await page.keyboard.press('Tab');
  await expect(skip).toBeFocused();
  const onScreen = await skip.boundingBox();
  expect(onScreen?.height ?? 0).toBeGreaterThanOrEqual(44);

  await page.keyboard.press('Enter');
  await expect(page).toHaveURL(/#main$/);
  await expect(page.locator('main')).toBeFocused();
});

test('the header and footer navigations are distinguishable landmarks', async ({ page }) => {
  await page.goto('/');
  await expect(page.getByRole('navigation', { name: 'Primary' })).toBeVisible();
  await expect(page.getByRole('navigation', { name: 'Footer' })).toBeVisible();
});
