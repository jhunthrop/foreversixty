// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the four product panels carry Planner, Simulator, Logs and Rankings, in spec order', async ({
  page,
}) => {
  await page.goto('/');
  const grid = page.getByTestId('home-product-panels');
  const labels = await grid.locator('[data-testid^="home-product-"] .label').allTextContents();
  expect(labels).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
  await expect(
    page.getByTestId('home-product-simulator').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
});

test('the addon and companion row links to /addon and /logs#companion, in that order', async ({ page }) => {
  await page.goto('/');
  const row = page.getByTestId('home-companion-row');
  const titles = await row.locator('.font-display').allTextContents();
  expect(titles).toEqual(['The addon', 'The companion']);
  await expect(row.getByRole('link', { name: /The addon/ })).toHaveAttribute('href', '/addon');
  await expect(row.getByRole('link', { name: /The companion/ })).toHaveAttribute('href', '/logs#companion');
});

test('the reference tiles carry Classes, Guides, Zones, Dungeons, in that order', async ({ page }) => {
  await page.goto('/');
  const grid = page.getByTestId('reference-tiles');
  const cards = grid.getByRole('link');
  const titles = await cards.evaluateAll((links) =>
    links.map((a) => a.querySelector('.font-display')?.textContent?.trim() ?? ''),
  );
  expect(titles).toEqual(['Classes', 'Guides', 'Zones', 'Dungeons']);
  await expect(grid.getByRole('link', { name: /Zones/ })).toHaveAttribute('href', '/zones');
  await expect(grid.getByRole('link', { name: /Dungeons/ })).toHaveAttribute('href', '/dungeons');
});
