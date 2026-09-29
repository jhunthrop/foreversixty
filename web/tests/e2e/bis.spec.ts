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
// specifically rather than the generic no-known-source line. Every real published file now
// carries `alternatives` and `reference_dps_per_point` (fix round 1, 2026-09-29 -- the
// ranker's own follow-up), so every state below is exercised against real published data:
// `hunter/marksmanship` band 20 for the two-hander empty off-hand and for alternatives
// (including a tie), `hunter/marksmanship` band 60 for the weight rail's DPS-per-point line
// and its "No effect" row, and `druid/balance` (which never equips a ranged weapon) for the
// ordinary empty-slot case. No fixture is committed for this page any more.
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

test('Leveling BiS: a pick’s alternatives render as rows beside it, every one its own hover target in the shared tooltip', async ({
  page,
}) => {
  await page.goto('/bis/hunter/marksmanship#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  await expect(band20).toBeVisible();

  // The head slot's real pick at band 20 carries three alternatives (Flying Tiger Goggles,
  // Shadow Goggles, Lucky Fishing Hat), each behind by a measured amount.
  const headRow = band20.getByTestId('bis-slot-alliance-20-head');
  await expect(headRow.getByText('Also:')).toBeVisible();

  // The pick's own hover target never contains a nested [role="button"] (fix round 1: a
  // button nested in a button broke keyboard/AT tab order when alternatives lived inside
  // it) -- every alternative is a sibling now, so the pick's own host has none.
  const pickHostForA11y = headRow.locator('[data-testid^="item-hover-"]').first();
  await expect(pickHostForA11y.locator('[role="button"]')).toHaveCount(0);
  const altHosts = headRow.locator('[data-testid^="item-hover-"]');
  // The pick itself plus its three alternatives -- four independent tooltip hosts on this
  // one row (fix round 1: alternatives are a sibling block, never nested inside the pick's
  // own host).
  await expect(altHosts).toHaveCount(4);
  // Fix round 1 (wow-player): each alternative names its item level.
  await expect(headRow).toContainText(/ilvl \d+/);
  await expect(headRow).toContainText(/−\d+\.\d DPS/);

  // legs, the same band, has a genuine tie -- "same DPS", never a signed "+0.0 DPS".
  const legsRow = band20.getByTestId('bis-slot-alliance-20-legs');
  await expect(legsRow).toContainText('same DPS');

  // Hovering the pick's own NAME (fix round 1: the regression this test guards against was
  // hovering the pick's visual centre landing on a nested alternative instead) shows the
  // pick's own tooltip, not an alternative's.
  const pickHost = headRow.locator('[data-testid^="item-hover-"]').first();
  const pickName = (await pickHost.locator('.gear-row-name').innerText()).trim();
  await pickHost.locator('.gear-row-name').hover();
  const tooltip = page.getByTestId('item-tooltip');
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(pickName);
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();

  // An alternative's own icon opens the exact same shared tooltip, for that alternative.
  const altHost = headRow.locator('[data-testid^="item-hover-"]').nth(1);
  const altName = (await altHost.locator('.gear-row-alt-name').innerText()).trim();
  await altHost.hover();
  await expect(tooltip).toBeVisible();
  await expect(tooltip).toContainText(altName);
  await page.mouse.move(0, 0);
  await expect(tooltip).toBeHidden();

  // trinket1 at this band has no alternatives at all -- "Also:" is absent entirely, not an
  // empty placeholder (spec §4).
  const trinket1Row = band20.getByTestId('bis-slot-alliance-20-trinket1');
  await expect(trinket1Row.getByText('Also:')).toHaveCount(0);
});

test('Leveling BiS: the weight rail shows a DPS-per-point line when the band carries reference_dps_per_point, and "No effect" for an insignificant stat', async ({
  page,
}) => {
  await page.goto('/bis/hunter/marksmanship#band-alliance-60');
  const band60 = page.getByTestId('bis-band-alliance-60');
  const weights60 = band60.getByTestId('bis-weights-alliance-60');
  // Band 60's own reference_dps_per_point: "1 Ranged attack power = <n> DPS".
  await expect(band60).toContainText(/1 Ranged attack power = \d+\.\d\d DPS/);
  // A significant, non-reference row reads its own weight-times-reference DPS clause.
  await expect(weights60).toContainText(/DPS per point/);
  // hit is flagged insignificant at band 60 -- "No effect", never a number or a bar
  // implying a real measurement.
  await expect(weights60).toContainText('No effect');
});

test('Leveling BiS: an ordinary empty slot (not a two-hander gap) reads the plain no-source copy', async ({
  page,
}) => {
  // Druids never equip a ranged weapon -- every band's own ranged slot is a genuine "no
  // known source" gap, never a two-hander side effect.
  await page.goto('/bis/druid/balance#band-alliance-20');
  const band20 = page.getByTestId('bis-band-alliance-20');
  const ranged = band20.getByTestId('bis-slot-alliance-20-ranged');
  await expect(ranged).toContainText('No sourced item at this level yet');
  await expect(ranged).not.toContainText('Two-hander equipped');
});
