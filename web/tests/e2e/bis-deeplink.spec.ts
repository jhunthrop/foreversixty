// web/tests/e2e/bis-deeplink.spec.ts
// R1: a shared link's state is the page's state. ?band=60 opens on 60 (and the old
// #band-<faction>-<band> hash still does), with no visible swap from the default band: until the
// island applies it, the changing regions are hidden behind html[data-bis-pending].
import { test, expect } from './support/test';
import { openBis } from './support/bis-page';

test('?band=60 selects 60 in the control bar and shows the 60 doll', async ({ page }) => {
  await openBis(page, '/bis/hunter/marksmanship?band=60');
  await expect(page.getByTestId('control-band-60')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('bis-page')).toHaveAttribute('data-bis-band', '60');
  await expect(page.getByTestId('bis-band-alliance-60')).toBeVisible();
  await expect(page).toHaveURL(/[?&]band=60/);
});

test('?faction=horde&band=40 opens on Horde 40', async ({ page }) => {
  await openBis(page, '/bis/mage/fire?faction=horde&band=40');
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
  await expect(page.getByTestId('bis-band-horde-40')).toBeVisible();
});

test('the legacy #band hash still opens on its band and faction', async ({ page }) => {
  await openBis(page, '/bis/hunter/marksmanship#band-horde-60');
  await expect(page.getByTestId('bis-band-horde-60')).toBeVisible();
  await expect(page.getByTestId('control-faction-horde')).toHaveAttribute('aria-checked', 'true');
});

test('the default band never shows under a deep link: the doll is hidden until the state applies', async ({ page }) => {
  await page.route('**/_astro/BisSpecPage.*.js', async (route) => {
    await new Promise((resolve) => setTimeout(resolve, 1500));
    await route.continue();
  });
  await page.goto('/bis/hunter/marksmanship?band=60');
  await expect(page.locator('html[data-bis-pending]')).toHaveCount(1);
  await expect(page.locator('[data-testid^="bis-band-alliance-"]')).toBeHidden();
  await expect(page.getByTestId('bis-band-alliance-60')).toBeVisible({ timeout: 10_000 });
});
