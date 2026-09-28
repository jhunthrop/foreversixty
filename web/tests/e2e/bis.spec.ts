// web/tests/e2e/bis.spec.ts
// Lane bis-web (leveling-bis-design.md): index -> spec page -> toggle faction -> pick a
// band -> see that band's "New at" strip. The page's data comes straight off disk at build
// time (src/lib/bis/load.ts), independent of FOREVER_DATA, so this runs in the default
// fixture suite same as every other content spec here.
import { test, expect } from '@playwright/test';

test('Leveling BiS: index links to a spec, faction and band pills switch panels with no reload', async ({
  page,
}) => {
  await page.goto('/bis');
  await expect(page.getByRole('heading', { name: 'Leveling BiS', level: 1 })).toBeVisible();

  const marksmanshipLink = page.locator('a[href="/bis/hunter/marksmanship"]');
  await expect(marksmanshipLink).toBeVisible();
  await marksmanshipLink.click();
  await expect(page).toHaveURL(/\/bis\/hunter\/marksmanship$/);

  // Alliance is the default panel; its own band 10 table is visible, Horde's is not.
  await expect(page.getByTestId('bis-faction-panel-alliance')).toBeVisible();
  await expect(page.getByTestId('bis-band-alliance-10')).toBeVisible();

  // Toggle to Horde -- a label click on a hidden radio, no navigation.
  await page.getByTestId('bis-faction-toggle-horde').click();
  await expect(page.getByTestId('bis-faction-panel-horde')).toBeVisible();

  // Pick band 30 within the Horde panel.
  await page.getByTestId('bis-band-pill-horde-30').click();
  await expect(page.getByTestId('bis-band-horde-30')).toBeVisible();
  await expect(page.getByTestId('bis-new-at-horde-30')).toContainText('New at 30');
});

test('Leveling BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  await page.goto('/bis/druid/balance');
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});
