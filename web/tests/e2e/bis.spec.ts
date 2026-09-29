// web/tests/e2e/bis.spec.ts
// Lane bis-page-ux (leveling-bis-design.md, the 2026-09-28 UX pass): index -> spec page ->
// toggle faction -> pick a band -> see that band's own "new" row marker and a real quest
// source line, not a generic pill. The page's data comes straight off disk at build time
// (src/lib/bis/load.ts), independent of FOREVER_DATA, so this runs in the default fixture
// suite same as every other content spec here.
//
// Character-panel redesign (lane bis-character-panel, 2026-09-29): the band panel is now
// the paperdoll's three-column grid plus a weapon row (`.paperdoll-left`/`-right`/
// `-bottom`, testids below), a pick's alternatives are compact rows under it rather than a
// runner-up `<details>`, and an empty off-hand under a two-handed main hand says so
// specifically rather than the generic no-known-source line. `hunter/marksmanship`'s own
// real published data already has a two-hander band (main hand a staff/axe, off-hand
// empty, at levels 20 and 40) -- no fixture edit needed to exercise that path. The
// alternatives row and `reference_dps_per_point`-driven weight line are not in any real
// published file yet (the nightly has not regenerated with them), so those two states are
// exercised against `warrior/protection`'s own committed fixture instead
// (`src/data/fixtures/bis/warrior-protection.json`), a spec with no real file for the
// active build to shadow it.
import { test, expect } from '@playwright/test';

/** hunter-marksmanship's real data always empties off-hand together with a two-handed
 *  main hand (bands 20 and 40) -- never any other slot -- so a filled-vs-empty sweep over a
 *  band's rows can assert the right copy for either shape without hardcoding which band. */
async function expectEmptyRowCopy(
  row: import('@playwright/test').Locator,
  slotTestId: string,
): Promise<void> {
  if (slotTestId.endsWith('-off_hand')) {
    await expect(row).toContainText(/Two-hander equipped|No sourced item at this level yet/);
  } else {
    await expect(row).toContainText('No sourced item at this level yet');
  }
}

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
  expect(Number(firstLevel)).toBeGreaterThan(0);

  // The paperdoll's three regions are all present: the two flanking gear columns (6 + 8
  // slots) and the weapon row underneath (main hand, off hand, ranged).
  await expect(page.getByTestId(`bis-band-alliance-${firstLevel}-left`)).toBeVisible();
  await expect(page.getByTestId(`bis-band-alliance-${firstLevel}-right`)).toBeVisible();
  const weaponsRow = page.getByTestId(`bis-band-alliance-${firstLevel}-weapons`);
  await expect(weaponsRow).toBeVisible();
  await expect(weaponsRow.locator(`[data-testid="bis-slot-alliance-${firstLevel}-main_hand"]`)).toBeVisible();
  await expect(weaponsRow.locator(`[data-testid="bis-slot-alliance-${firstLevel}-ranged"]`)).toBeVisible();

  // A filled slot shows the real item (a GearRow with its own icon and tooltip host), and
  // an empty one says why rather than a bare dash (tenet 4) -- every slot row is one or the
  // other. An empty off-hand under a two-handed main hand gets its own copy (ruling 4),
  // never the generic no-known-source line.
  await expect(firstBand.locator('[data-testid^="item-hover-"]').first()).toBeVisible();
  const slotRows = firstBand.locator(`[data-testid^="bis-slot-alliance-${firstLevel}-"]`);
  const slotCount = await slotRows.count();
  expect(slotCount).toBe(17);
  for (let i = 0; i < slotCount; i += 1) {
    const row = slotRows.nth(i);
    const slotTestId = (await row.getAttribute('data-testid'))!;
    const filled = (await row.locator('[data-testid^="item-hover-"]').count()) > 0;
    if (!filled) await expectEmptyRowCopy(row, slotTestId);
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

test('Leveling BiS: a pick’s alternatives render as rows under it, every one openable in the shared tooltip', async ({
  page,
}) => {
  await page.goto('/bis/warrior/protection#band-alliance-60');
  const band60 = page.getByTestId('bis-band-alliance-60');
  await expect(band60).toBeVisible();

  // The head slot's pick (fixture: Bonescythe Helmet) carries two alternatives.
  const headRow = band60.getByTestId('bis-slot-alliance-60-head');
  await expect(headRow.getByText('Also:')).toBeVisible();
  const altHosts = headRow.locator('[data-testid^="item-hover-"]');
  // The pick itself plus its two alternatives -- three tooltip hosts on this one row.
  await expect(altHosts).toHaveCount(3);

  // One alternative is a tie (dps_delta 0, fixture: Circlet of Faith) and reads "same DPS",
  // never a signed "+0.0 DPS" (ruling 2).
  await expect(headRow).toContainText('same DPS');
  // Another is behind by a measured amount, in muted text (ruling 2's own example shape).
  await expect(headRow).toContainText(/−\d+\.\d DPS/);

  // An alternative's icon opens the exact same shared tooltip a main pick's does.
  const altHost = headRow.locator('[data-testid^="item-hover-"]').nth(1);
  await altHost.hover();
  await expect(page.getByTestId('item-tooltip')).toBeVisible();
  await page.mouse.move(0, 0);
  await expect(page.getByTestId('item-tooltip')).toBeHidden();

  // A pick with no alternatives (fixture: every other slot) renders no "Also:" line at all
  // -- absent, not an empty placeholder (spec §4).
  const neckRow = band60.getByTestId('bis-slot-alliance-60-neck');
  await expect(neckRow.getByText('Also:')).toHaveCount(0);
});

test('Leveling BiS: the weight rail shows a DPS-per-point line when the band carries reference_dps_per_point, and "No effect" for an insignificant stat', async ({
  page,
}) => {
  await page.goto('/bis/warrior/protection#band-alliance-60');
  const band60 = page.getByTestId('bis-band-alliance-60');
  const weights60 = band60.getByTestId('bis-weights-alliance-60');
  // The fixture's own reference_dps_per_point (3.42): "1 Attack power = 3.42 DPS".
  await expect(band60).toContainText('1 Attack power = 3.42 DPS');
  // A significant, non-reference row reads its own weight-times-reference DPS clause.
  await expect(weights60).toContainText(/DPS per point/);
  // melee_haste is flagged insignificant in the fixture -- "No effect", never a number or
  // a bar implying a real measurement.
  await expect(weights60).toContainText('No effect');

  // Band 20 carries no reference_dps_per_point at all -- the rail falls back to the plain
  // reference sentence and today's bare ratios, never a fabricated DPS number.
  await page.getByTestId('bis-band-pill-alliance-20').click();
  const band20 = page.getByTestId('bis-band-alliance-20');
  const weights20 = band20.getByTestId('bis-weights-alliance-20');
  await expect(weights20).not.toContainText('DPS per point');
  await expect(band20).toContainText('reference stat');
});

test('Leveling BiS: an ordinary empty slot (not a two-hander gap) reads the plain no-source copy', async ({
  page,
}) => {
  await page.goto('/bis/warrior/protection#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  const finger2 = band20.getByTestId('bis-slot-alliance-20-finger2');
  await expect(finger2).toContainText('No sourced item at this level yet');

  // The two-hander case, on the very same band, reads its own copy instead.
  const offHand = band20.getByTestId('bis-slot-alliance-20-off_hand');
  await expect(offHand).toContainText('Two-hander equipped');
});
