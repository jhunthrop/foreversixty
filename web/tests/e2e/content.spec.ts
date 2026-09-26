import { expect, test } from '@playwright/test';

for (const path of ['/guides', '/changelog']) {
  test(`${path} carries the updated stamp of its newest entry`, async ({ page }) => {
    await page.goto(path);
    await expect(page.getByText(/^Updated /)).toBeVisible();
  });
}

test('changelog scopes the simulator tools entry to damage specs, not every spec', async ({ page }) => {
  await page.goto('/changelog');
  const entry = page
    .getByText(/Top Gear, Droptimizer, talent compare and stat weights/)
    .locator('xpath=ancestor::article');
  await expect(entry).toContainText('damage specs');
});
