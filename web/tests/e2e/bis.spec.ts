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

  // Alliance is the default panel; its own first band table is visible, Horde's is not.
  // The band list is the nightly's (20..60 step 10), so the test reads the first band off
  // the page rather than pinning its level.
  await expect(page.getByTestId('bis-faction-panel-alliance')).toBeVisible();
  const firstBand = page.locator('[data-testid^="bis-band-alliance-"]:visible').first();
  await expect(firstBand).toBeVisible();
  const firstLevel = (await firstBand.getAttribute('data-testid'))!.replace('bis-band-alliance-', '');
  expect(Number(firstLevel)).toBeGreaterThanOrEqual(20);

  // A filled slot shows the real item (ItemHover's pill), and an empty one says why rather
  // than a bare dash (tenet 4, this lane's own brief item 1) -- every slot row is one or
  // the other, never a bare dash.
  await expect(firstBand.locator('[data-testid^="item-hover-"]').first()).toBeVisible();
  const slotRows = firstBand.locator(`[data-testid^="bis-slot-alliance-${firstLevel}-"]`);
  const slotCount = await slotRows.count();
  expect(slotCount).toBeGreaterThan(0);
  for (let i = 0; i < slotCount; i += 1) {
    const row = slotRows.nth(i);
    const filled = (await row.locator('[data-testid^="item-hover-"]').count()) > 0;
    if (!filled) await expect(row).toContainText('No sourced item at this level yet');
  }

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

  // "What changed since level N" (a wall of diff rows above the list, owner screenshot
  // review 2026-09-29) is now a single disclosure line under the list -- "N upgrades since
  // level 20" -- that expands to the same before/after diff.
  await expect(band30).toContainText(/upgrades since level \d+/);
});

test('Leveling BiS: a spec with no ranked list yet shows the empty state, not a 404', async ({ page }) => {
  // Healers have no written rotation, so the nightly ranks nothing for them (a dps spec
  // gained a real file the night the nightly first ran, which is what this test once used).
  await page.goto('/bis/priest/holy');
  await expect(page.getByTestId('bis-empty-state')).toBeVisible();
  await expect(page.getByTestId('bis-empty-state')).toContainText('No leveling BiS list yet');
});
