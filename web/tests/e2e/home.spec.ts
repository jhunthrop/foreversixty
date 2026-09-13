import { expect, test } from '@playwright/test';

test('homepage renders the reference layout without a marketing hero', async ({ page }) => {
  const errors: string[] = [];
  page.on('console', (m) => { if (m.type() === 'error') errors.push(m.text()); });
  await page.goto('/');
  await expect(page.locator('h1')).toHaveCount(0);
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
