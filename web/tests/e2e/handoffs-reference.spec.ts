// web/tests/e2e/handoffs-reference.spec.ts
import { expect, test } from '@playwright/test';

test('a class guide links to the planner for that class', async ({ page }) => {
  await page.goto('/guides/warrior');
  const link = page.getByRole('link', { name: 'Open the planner for this class' });
  await expect(link).toHaveAttribute('href', '/planner?class=warrior');
});

test('a dungeon whose zone this site covers links to that zone', async ({ page }) => {
  await page.goto('/dungeons/kroldok-stronghold');
  await expect(page.getByRole('link', { name: 'Riverglades' })).toHaveAttribute('href', '/zones/riverglades');
});

test('a dungeon whose zone this site does not cover names it without a broken link', async ({ page }) => {
  await page.goto('/dungeons/hall-of-thanes');
  await expect(page.getByText('Dun Morogh')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Dun Morogh' })).toHaveCount(0);
});

test('a dungeon with fixture loot data offers drops for your character', async ({ page }) => {
  await page.goto('/dungeons/hall-of-thanes');
  await expect(page.getByTestId('dungeon-drops-link')).toHaveAttribute(
    'href',
    '/sim/drops?instance=hall-of-thanes',
  );
});

test('/zones names which zones are covered and that the rest are coming', async ({ page }) => {
  await page.goto('/zones');
  await expect(page.getByText('The rest of the world', { exact: false })).toBeVisible();
});
