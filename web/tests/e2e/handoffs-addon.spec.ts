// web/tests/e2e/handoffs-addon.spec.ts
import { expect, test } from '@playwright/test';
import { ACTIVE_BUILD } from './support/active-build';

test('the /addon paste box offers the planner and the simulator once an export decodes', async ({ page }) => {
  await page.goto('/addon');
  await page.getByTestId('addon-paste-code').fill(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`);
  await page.getByTestId('addon-paste-submit').click();

  await expect(page.getByTestId('addon-paste-error')).toHaveCount(0);
  const planner = page.getByTestId('addon-paste-planner');
  const sim = page.getByTestId('addon-paste-sim');
  await expect(planner).toHaveAttribute(
    'href',
    `/planner?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
  await expect(sim).toHaveAttribute(
    'href',
    `/sim?code=${encodeURIComponent(`FS1:${ACTIVE_BUILD}:warrior:human:0/0/0:`)}`,
  );
});

test('a code from another format is refused by name, not silently dropped', async ({ page }) => {
  await page.goto('/addon');
  await page.getByTestId('addon-paste-code').fill('FS2:nope');
  await page.getByTestId('addon-paste-submit').click();
  await expect(page.getByTestId('addon-paste-error')).toHaveText('That code is FS2; this site reads FS1.');
  await expect(page.getByTestId('addon-paste-planner')).toHaveCount(0);
});

test('the page says the in-game UI is in beta testing and ships no screenshot of it', async ({ page }) => {
  await page.goto('/addon');
  await expect(page.getByText('in beta testing in game')).toBeVisible();
  const images = await page.locator('main img').count();
  expect(images).toBe(0);
});
