import { expect, test } from './support/test';

for (const path of ['/guides', '/changelog']) {
  test(`${path} carries the updated stamp of its newest entry`, async ({ page }) => {
    await page.goto(path);
    await expect(page.getByText(/^Updated /)).toBeVisible();
  });
}

test('changelog scopes the simulator tools entry to damage specs, not every spec', async ({ page }) => {
  // A pre-launch entry is a title row; its permalink opens it in place (pages spec 2026-11-08 CL1).
  await page.goto('/changelog#2026-09-20-simulator-tools-are-live');
  const entry = page.locator('[id="2026-09-20-simulator-tools-are-live"]');
  await expect(entry).toHaveAttribute('open', '');
  await expect(entry.getByText(/damage specs/).first()).toBeVisible();
});
