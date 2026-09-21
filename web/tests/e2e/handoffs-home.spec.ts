// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the homepage tools grid carries Simulator and The addon, in spec order', async ({ page }) => {
  await page.goto('/');
  const grid = page.getByTestId('tools-grid');
  const cards = grid.getByRole('link');
  const titles = await cards.evaluateAll((links) =>
    links.map((a) => a.querySelector('.font-display')?.textContent?.trim() ?? ''),
  );
  expect(titles).toEqual([
    'Build planner',
    'Simulator',
    'Combat logs',
    'Rankings',
    'The addon',
    'Dungeons',
    'Zone atlas',
    'Class guides',
  ]);
  await expect(grid.getByRole('link', { name: /Simulator/ })).toHaveAttribute('href', '/sim');
  await expect(grid.getByRole('link', { name: /The addon/ })).toHaveAttribute('href', '/addon');
});
