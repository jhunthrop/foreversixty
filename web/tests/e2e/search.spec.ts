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
