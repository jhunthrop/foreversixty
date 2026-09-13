import { expect, test } from '@playwright/test';

test('homepage renders the reference layout without a marketing hero', async ({ page }) => {
  const errors: string[] = [];
  page.on('console', (m) => { if (m.type() === 'error') errors.push(m.text()); });
  await page.goto('/');
  const h1 = page.locator('h1');
  await expect(h1).toHaveCount(1);
  const heading = (await h1.innerText()).trim();
  // A reference heading, not a slogan: short, and it does not end in punctuation.
  expect(heading.length).toBeLessThanOrEqual(90);
  expect(heading).not.toMatch(/[.!?]$/);
  await expect(page.getByRole('searchbox')).toBeVisible();
  await expect(page.getByText('Right now')).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Tools' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'What changed' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Still unknown' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Guides by class' })).toBeVisible();
  await expect(page.getByText('Not affiliated with or endorsed by Blizzard Entertainment')).toBeVisible();
  expect(errors).toEqual([]);
});

test('content pages ship no client JavaScript', async ({ page }) => {
  const scripts: string[] = [];
  page.on('request', (r) => { if (r.resourceType() === 'script') scripts.push(r.url()); });
  await page.goto('/about');
  expect(scripts).toEqual([]);
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
