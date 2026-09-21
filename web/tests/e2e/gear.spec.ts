import { expect, test } from '@playwright/test';
import { openGear, shareBuild } from './support/planner';

// Gear is the optional half of a build and the only part of the planner that reads the item
// data. The set bonus carries the most: it appears only once both pieces are on, so it
// exercises the whole chain from a click through equip to activeSetBonuses.
test('equipping items fills the slot, totals the stats and activates the set bonus', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-head').click();
  await page.getByTestId('item-16963').click();
  await expect(page.getByTestId('slot-head')).toContainText('Helm of Wrath');
  // The button's label is what a screen reader hears, and it is built separately from the
  // text above, so the equipped item has to reach both.
  await expect(page.getByTestId('slot-head')).toHaveAccessibleName('Head: Helm of Wrath');
  await expect(page.getByTestId('item-picker')).toBeHidden();

  await page.getByTestId('slot-shoulder').click();
  await page.getByTestId('item-16966').click();

  // Label and value together rather than as two searches of the same block, so a number that
  // landed on the wrong stat could not pass.
  const totals = page.getByTestId('gear-totals');
  await expect(totals).toContainText(/Armor\s*1110/);
  await expect(totals).toContainText(/Strength\s*45/);
  await expect(totals).toContainText(/Stamina\s*38/);

  const sets = page.getByTestId('gear-sets');
  await expect(sets).toContainText('Battlegear of Wrath');
  await expect(sets).toContainText('(2/2)');
  await expect(sets).toContainText('Increases your chance to parry an attack by 1%.');
});

test('the picker filters by name and can clear the slot again', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-head').click();
  await page.getByLabel('Filter Head items').fill('lionheart');
  await expect(page.getByTestId('item-12640')).toBeVisible();
  await expect(page.getByTestId('item-16963')).toBeHidden();
  await page.getByTestId('item-12640').click();
  await expect(page.getByTestId('slot-head')).toContainText('Lionheart Helm');

  await page.getByTestId('slot-head').click();
  await page.getByRole('button', { name: 'Clear slot' }).click();
  await expect(page.getByTestId('slot-head')).toContainText('Empty');
  await expect(page.getByTestId('slot-head')).toHaveAccessibleName('Head: empty');
  await expect(page.getByTestId('gear-totals')).toContainText('Nothing equipped.');
});

test('a slot no class item fits says so instead of showing an empty list', async ({ page }) => {
  await page.goto('/planner');
  await openGear(page);

  await page.getByTestId('slot-chest').click();
  await expect(page.getByTestId('item-picker')).toContainText('No item in this class list fits Chest.');
});

// tests/e2e/planner-share.spec.ts owns the shape of the save body; this owns the one field
// that only the gear panel can fill.
test('gear travels with the shared build', async ({ page }) => {
  let gear: unknown;
  await page.route('**/v1/builds', async (route) => {
    gear = (route.request().postDataJSON() as { gear: unknown }).gear;
    await route.fulfill({
      status: 201,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
        error: null,
        request_id: 'r',
      }),
    });
  });
  await page.goto('/planner');
  // Spent before the gear panel is opened: below md the talent tree and the gear panel are
  // two tabs of the same strip, so only one of them is on screen at a time.
  await page.getByTestId('talent-1001').click();
  await openGear(page);
  await page.getByTestId('slot-main_hand').click();
  await page.getByTestId('item-12784').click();
  await shareBuild(page);
  await expect(page.getByTestId('share-link')).toBeVisible();
  expect(gear).toEqual({ main_hand: 12784 });
});

test.describe('gear on a phone', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  test('gear is a tab and the page still never scrolls sideways', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    await expect(tabs).toHaveCount(3);
    await tabs.nth(2).click();
    await expect(page.getByTestId('gear-panel')).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeHidden();

    await page.getByTestId('slot-head').click();
    await page.getByTestId('item-16963').click();
    const overflow = await page.evaluate(
      () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
    );
    expect(overflow).toBeLessThanOrEqual(0);
  });
});
