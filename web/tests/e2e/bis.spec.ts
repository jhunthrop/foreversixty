// web/tests/e2e/bis.spec.ts
// Lane bis-page-ux (leveling-bis-design.md, the 2026-09-28 UX pass): index -> spec page ->
// toggle faction -> pick a band -> see that band's own "new" row marker and a real quest
// source line, not a generic pill. The page's data comes straight off disk at build time
// (src/lib/bis/load.ts), independent of FOREVER_DATA, so this runs in the default fixture
// suite same as every other content spec here.
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
  const band10 = page.getByTestId('bis-band-alliance-10');
  await expect(band10).toBeVisible();

  // A filled slot shows the real item (ItemHover's pill), and an empty one says why rather
  // than a bare dash (tenet 4, this lane's own brief item 1).
  await expect(band10.locator('[data-testid^="item-hover-"]').first()).toBeVisible();
  await expect(band10.getByTestId('bis-slot-alliance-10-neck')).toContainText('No sourced item at this level yet');

  // Toggle to Horde -- a label click on a hidden radio, no navigation.
  await page.getByTestId('bis-faction-toggle-horde').click();
  await expect(page.getByTestId('bis-faction-panel-horde')).toBeVisible();

  // Pick band 30 within the Horde panel.
  await page.getByTestId('bis-band-pill-horde-30').click();
  const band30 = page.getByTestId('bis-band-horde-30');
  await expect(band30).toBeVisible();

  // A row this band marks new (hunter-marksmanship/horde/30 has several -- the neck slot's
  // own quest pick among them), and a real quest source line, not the generic "Quests" pill
  // the picker uses elsewhere -- tenet 2's "an item is never just a name" applied to sources.
  await expect(band30.getByTestId('bis-row-new').first()).toBeVisible();
  await expect(band30).toContainText('Quest:');
  await expect(band30).toContainText('What changed since level 25');
});

test('Leveling BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  // Healers have no written rotation, so the nightly ranks nothing for them (a dps spec
  // gained a real file the night the nightly first ran, which is what this test once used).
  await page.goto('/bis/priest/holy');
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});
