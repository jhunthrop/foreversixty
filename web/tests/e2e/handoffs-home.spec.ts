// web/tests/e2e/handoffs-home.spec.ts
import { expect, test } from '@playwright/test';

test('the four next-steps cards carry Planner, Simulator, Logs and Rankings, in spec order', async ({
  page,
}) => {
  await page.goto('/');
  const grid = page.getByTestId('home-next-steps-signed-out');
  const labels = await grid.locator('[data-testid^="home-next-"] .label').allTextContents();
  expect(labels).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
  await expect(
    page.getByTestId('home-next-simulator-signed-out').getByRole('link', { name: 'Open the simulator' }),
  ).toHaveAttribute('href', '/sim');
});

test('the "Get set up" card links to /setup', async ({ page }) => {
  await page.goto('/');
  const row = page.getByTestId('home-companion-row');
  await expect(row.getByRole('link', { name: /Get set up/ })).toHaveAttribute('href', '/setup');
});
