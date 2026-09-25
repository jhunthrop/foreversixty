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
  await expect(page.getByTestId('home-timeline')).toBeVisible();
  await expect(page.getByTestId('home-next-planner-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-simulator-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-logs-signed-out')).toBeVisible();
  await expect(page.getByTestId('home-next-rankings-signed-out')).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Your guild' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'What changed' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Still unknown' })).toBeVisible();
  await expect(page.getByText('Not affiliated with or endorsed by Blizzard Entertainment')).toBeVisible();
  expect(errors).toEqual([]);
});

test('the four next-steps cards link to their tools, with the planner and simulator live elements gone from this page', async ({
  page,
}) => {
  await page.goto('/');
  await expect(
    page.getByTestId('home-next-planner-signed-out').getByRole('link', { name: 'Open the planner' }),
  ).toHaveAttribute('href', '/planner');
  await expect(
    page.getByTestId('home-next-simulator-signed-out').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
  await expect(
    page.getByTestId('home-next-logs-signed-out').getByRole('link', { name: 'Open the logs' }),
  ).toHaveAttribute('href', '/logs');
  await expect(
    page.getByTestId('home-next-rankings-signed-out').getByRole('link', { name: 'Open the rankings' }),
  ).toHaveAttribute('href', '/rankings');
  // The class-tile row and the spec-pill wall left the home page (spec 2026-09-24 §3).
  await expect(page.getByRole('link', { name: 'Warrior', exact: true })).toHaveCount(0);
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
