import { expect, test } from '@playwright/test';

test('search finds a dungeon page and navigates to it', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await expect(box).toBeFocused();
  await box.fill('Thanes');
  const result = page.getByRole('option', { name: /Hall of Thanes/ });
  await expect(result).toBeVisible();
  await result.click();
  await expect(page).toHaveURL(/\/dungeons\/hall-of-thanes$/);
});

test('no keystrokes are dropped when typing immediately after the "/" shortcut', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await expect(box).toBeFocused();
  await page.keyboard.type('Thanes', { delay: 10 });
  await expect(box).toHaveValue('Thanes');
  const result = page.getByRole('option', { name: /Hall of Thanes/ });
  await expect(result).toBeVisible();
});

test('arriving at /search?q= seeds the box from the URL and shows results', async ({ page }) => {
  await page.goto('/search?q=Thanes');
  await expect(page.getByRole('searchbox')).toHaveValue('Thanes');
  await expect(page.getByRole('option', { name: /Hall of Thanes/ })).toBeVisible();
});

test('Enter before the debounce resolves submits the query to the results page', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  await expect(page.getByRole('searchbox')).toBeFocused();
  await page.keyboard.type('Thanes');
  await page.keyboard.press('Enter');
  await expect(page).toHaveURL(/\/search\?q=Thanes$/);
  await expect(page.getByRole('searchbox')).toHaveValue('Thanes');
  await expect(page.getByRole('option', { name: /Hall of Thanes/ })).toBeVisible();
});

test('an unreachable search index falls back to links, once', async ({ page }) => {
  const crashes: string[] = [];
  const indexRequests: string[] = [];
  page.on('pageerror', (e) => crashes.push(e.message));
  page.on('request', (r) => { if (r.url().includes('/pagefind/pagefind.js')) indexRequests.push(r.url()); });
  await page.route('**/pagefind/pagefind.js', (route) => route.fulfill({ status: 404, body: 'not found' }));

  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await box.fill('Thanes');

  const fallback = page.getByText('Search is unavailable. Browse');
  await expect(fallback).toBeVisible();
  for (const [name, href] of [['dungeons', '/dungeons'], ['zones', '/zones'], ['guides', '/guides']]) {
    await expect(fallback.getByRole('link', { name })).toHaveAttribute('href', href);
  }

  await box.fill('Hyjal');
  await expect(fallback).toBeVisible();
  expect(indexRequests).toHaveLength(1);
  expect(crashes).toEqual([]);
});

test('the results dropdown closes on outside click and drops aria-controls', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await expect(box).not.toHaveAttribute('aria-controls');
  await box.fill('Thanes');
  await expect(page.getByRole('listbox')).toBeVisible();
  await expect(box).toHaveAttribute('aria-controls', 'search-results');

  await page.getByRole('heading', { name: 'Tools' }).click();
  await expect(page.getByRole('listbox')).toBeHidden();
  await expect(box).not.toHaveAttribute('aria-controls');
});

test('ArrowUp from the first result hands the query back to the input', async ({ page }) => {
  await page.goto('/');
  await page.keyboard.press('/');
  const box = page.getByRole('searchbox');
  await box.fill('Thanes');
  await expect(box).toHaveAttribute('aria-activedescendant', 'search-opt-0');
  await page.keyboard.press('ArrowUp');
  await expect(box).not.toHaveAttribute('aria-activedescendant');
  await expect(page.getByRole('option').first()).toHaveAttribute('aria-selected', 'false');
  await page.keyboard.press('Enter');
  await expect(page).toHaveURL(/\/search\?q=Thanes$/);
});
