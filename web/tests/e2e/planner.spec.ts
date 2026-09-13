import { expect, test } from '@playwright/test';

test('the planner opens on the default class with an empty build', async ({ page }) => {
  const errors: string[] = [];
  page.on('console', (m) => {
    if (m.type() === 'error') errors.push(m.text());
  });
  await page.goto('/planner');
  await expect(page.getByTestId('planner-level')).toHaveText('9');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(
    page.getByText(
      'Classic Era trees shown until the beta client exports; Forever revamped talents replace them then.',
    ),
  ).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
  // Both trees render, but below the md breakpoint only the selected tab's panel is on screen,
  // so Fury is in the document rather than visible -- tests/e2e/planner-phone.spec.ts owns the
  // one-tree-at-a-time rule, and this project runs at a phone width.
  await expect(page.getByRole('heading', { name: 'Fury', includeHidden: true })).toBeAttached();
  expect(errors).toEqual([]);
});

test('?class and ?race preselect the build', async ({ page }) => {
  await page.goto('/planner?class=warrior&race=dwarf');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page.getByLabel('Race')).toHaveValue('dwarf');
});

test('a failed talent fetch shows the reason and a working retry', async ({ page }) => {
  let attempts = 0;
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    attempts += 1;
    if (attempts === 1) return route.abort('failed');
    return route.continue();
  });
  await page.goto('/planner');
  await expect(page.getByText('Talent data did not load')).toBeVisible();
  await page.getByRole('button', { name: 'Retry' }).click();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
});

test('clicking a talent spends points and the counters follow', async ({ page }) => {
  await page.goto('/planner');
  const improvedHeroicStrike = page.getByTestId('talent-1001');
  await improvedHeroicStrike.click();
  await improvedHeroicStrike.click();
  await improvedHeroicStrike.click();
  await expect(improvedHeroicStrike).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('planner-spent')).toHaveText('3/51');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-level')).toHaveText('12');
});

test('a locked tier refuses the point and says why', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1003').click();
  await expect(page.getByTestId('planner-refusal')).toHaveText('Tier 1 of Arms needs 5 points in Arms first');
  await expect(page.getByTestId('talent-1003')).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
});

test('a full talent refuses a fourth point with its own reason', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 4; i += 1) await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('planner-refusal')).toHaveText(
    'Improved Heroic Strike is already at 3 of 3 points',
  );
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
});

test('right-click removes the last point in a talent', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').click({ button: 'right' });
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '1');
});

test('removal is refused when a later point depends on it', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1002').click();
  await page.getByTestId('talent-1002').click();
  await page.getByTestId('talent-1003').click();
  await page.getByTestId('talent-1001').click({ button: 'right' });
  await expect(page.getByTestId('planner-refusal')).toHaveText('Tier 1 of Arms needs 5 points in Arms first');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
});

test('the tooltip shows the current and the next rank', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByTestId('talent-1001').hover();
  const tip = page.getByRole('tooltip');
  await expect(tip).toContainText('Improved Heroic Strike');
  await expect(tip).toContainText('Rank 1 of 3');
  await expect(tip).toContainText('Reduces the rage cost of Heroic Strike by 1.');
  await expect(tip).toContainText('Reduces the rage cost of Heroic Strike by 2.');
});

test('a keyboard-only visitor can spend and remove points', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').focus();
  await page.keyboard.press('Enter');
  await page.keyboard.press('Enter');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '2');
  await page.keyboard.press('ArrowRight');
  await expect(page.getByTestId('talent-1002')).toBeFocused();
  await page.keyboard.press('Enter');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '1');
  await page.keyboard.press('Backspace');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('planner-spent')).toHaveText('2/51');
});

// Touch-and-hold raises a native contextmenu at roughly the same 500ms threshold as the
// long-press timer, so one gesture can reach both removal paths. Playwright cannot drive
// that native gesture (its touchscreen API has no long press, and synthesised touches do
// not raise contextmenu), so this dispatches both paths at the cell for a single press and
// asserts they remove one point between them -- in either arrival order, and for the
// secondary button, where the same race exists on desktop.
test('a long press that also raises a context menu removes one point, not two', async ({ page }) => {
  await page.goto('/planner');
  const talent = page.getByTestId('talent-1001');
  for (let i = 0; i < 3; i += 1) await talent.click();
  await expect(talent).toHaveAttribute('data-rank', '3');

  // The context menu arrives while the long press is still pending.
  await talent.dispatchEvent('pointerdown', { button: 0 });
  await talent.dispatchEvent('contextmenu');
  await page.waitForTimeout(900);
  await talent.dispatchEvent('pointerup', { button: 0 });
  await talent.dispatchEvent('click', { button: 0 });
  await expect(talent).toHaveAttribute('data-rank', '2');

  // The long press fires first and the context menu follows it.
  await talent.dispatchEvent('pointerdown', { button: 0 });
  await page.waitForTimeout(900);
  await talent.dispatchEvent('contextmenu');
  await talent.dispatchEvent('pointerup', { button: 0 });
  await talent.dispatchEvent('click', { button: 0 });
  await expect(talent).toHaveAttribute('data-rank', '1');

  // A secondary-button press held past the threshold is one gesture too.
  await talent.dispatchEvent('pointerdown', { button: 2 });
  await page.waitForTimeout(900);
  await talent.dispatchEvent('contextmenu');
  await talent.dispatchEvent('pointerup', { button: 2 });
  await expect(talent).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
});

test('the order strip lists every point with the level it was spent at', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  const points = page.getByTestId('order-strip').getByRole('listitem');
  await expect(points).toHaveCount(3);
  await expect(points.nth(0)).toContainText('10');
  await expect(points.nth(1)).toContainText('11');
  await expect(points.nth(2)).toContainText('12');
  await expect(points.nth(0)).toContainText('Improved Heroic Strike');
});

test('the order strip collapses and reopens', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  const toggle = page.getByRole('button', { name: 'Hide point order' });
  await toggle.click();
  await expect(page.getByTestId('order-strip').getByRole('list')).toBeHidden();
  await page.getByRole('button', { name: 'Show point order' }).click();
  await expect(page.getByTestId('order-strip').getByRole('list')).toBeVisible();
});

test('reset asks before it clears the build', async ({ page }) => {
  await page.goto('/planner');
  for (let i = 0; i < 2; i += 1) await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Reset' }).click();
  await expect(page.getByTestId('planner-spent')).toHaveText('2/51');
  // The safe answer takes the focus, so a second Enter out of habit keeps the build.
  await expect(page.getByRole('button', { name: 'Keep the build' })).toBeFocused();
  await page.getByRole('button', { name: 'Keep the build' }).click();
  await expect(page.getByTestId('planner-spent')).toHaveText('2/51');

  await page.getByRole('button', { name: 'Reset' }).click();
  await page.getByRole('button', { name: 'Clear all points' }).click();
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '0');
  await expect(page.getByTestId('order-strip').getByRole('listitem')).toHaveCount(0);
});

// The strip is the one part of the planner that grows without bound, and the phone layout
// budget (tests/e2e/layout.spec.ts) allows no sideways page scroll at 360px.
test('a long point order scrolls inside the strip, not across the page', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 800 });
  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  for (let i = 0; i < 5; i += 1) await page.getByTestId('talent-1002').click();
  await expect(page.getByTestId('order-strip').getByRole('listitem')).toHaveCount(8);
  const widths = await page.evaluate(() => {
    const list = document.querySelector('[data-testid="order-strip"] ol');
    return {
      pageScroll: document.documentElement.scrollWidth,
      pageClient: document.documentElement.clientWidth,
      listScroll: list?.scrollWidth ?? 0,
      listClient: list?.clientWidth ?? 0,
    };
  });
  expect(widths.listScroll).toBeGreaterThan(widths.listClient);
  expect(widths.pageScroll).toBeLessThanOrEqual(widths.pageClient);
});

// Gear is the optional half of a build and the only part of the planner that reads the item
// data. The set bonus is the assertion that carries the most: it appears only once both
// pieces of the set are on, so it exercises the whole chain from a click to activeSetBonuses.
test('equipping gear fills the slot, sums the stats and unlocks the set bonus', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('slot-head').click();
  await page.getByTestId('item-16963').click();
  await expect(page.getByTestId('slot-head')).toHaveAccessibleName('Head: Helm of Wrath');
  await expect(page.getByTestId('item-picker')).toBeHidden();

  await page.getByTestId('slot-shoulder').click();
  await page.getByTestId('item-16966').click();

  const totals = page.getByTestId('gear-totals');
  await expect(totals).toContainText(/Armor\s*1110/);
  await expect(totals).toContainText(/Strength\s*45/);
  await expect(totals).toContainText(/Stamina\s*38/);

  const sets = page.getByTestId('gear-sets');
  await expect(sets).toContainText('Battlegear of Wrath');
  await expect(sets).toContainText('(2/2)');
  await expect(sets).toContainText('Increases your chance to parry an attack by 1%.');
});

test('the item picker filters by name, and the slot can be cleared again', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('slot-head').click();
  const rows = page.getByTestId('item-picker').getByRole('listitem');
  await expect(rows).toHaveCount(2);
  await page.getByLabel('Filter Head items').fill('lionheart');
  await expect(rows).toHaveCount(1);
  await page.getByTestId('item-12640').click();
  await expect(page.getByTestId('slot-head')).toHaveAccessibleName('Head: Lionheart Helm');

  await page.getByTestId('slot-head').click();
  await page.getByRole('button', { name: 'Clear slot' }).click();
  await expect(page.getByTestId('slot-head')).toHaveAccessibleName('Head: empty');
  await expect(page.getByTestId('gear-totals')).toContainText('Nothing equipped.');
});

// Switching class empties the build by itself, and the class selector stays reachable while the
// confirm is open, so a confirm that survived the switch would ask about points already gone.
test('a class switch drops a reset confirm that is still open', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Reset' }).click();
  await expect(page.getByRole('button', { name: 'Clear all points' })).toBeVisible();

  // Warrior is the only class with talent data on this build, so the way back to a ready
  // planner under a different class is out and back again.
  await page.getByLabel('Class').selectOption('paladin');
  await expect(page.getByText('Talent data did not load')).toBeVisible();
  await page.getByLabel('Class').selectOption('warrior');
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();

  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByRole('button', { name: 'Clear all points' })).toBeHidden();
  await expect(page.getByRole('button', { name: 'Reset' })).toBeVisible();
});
