import { expect, test } from '@playwright/test';
import { collectPageErrors } from './support/console';

import { treeSourceNotice } from '../../src/lib/planner/tree-source';
import { ACTIVE_BUILD } from './support/active-build';

// tree-source.ts carries no import.meta.env dependency (unlike its neighbour config.ts,
// which reads PUBLIC_API_BASE_URL at module scope and cannot be imported under
// Playwright's plain Node loader), so this is a real import rather than a reproduction.
const ACTIVE_BUILD_NOTICE = treeSourceNotice(ACTIVE_BUILD);

test('the planner opens on the default class with an empty build', async ({ page }) => {
  const errors = collectPageErrors(page);
  await page.goto('/planner');
  // A bare /planner has no current-character pointer, so Level -- which reads as a real
  // character's level -- is not shown for a build nobody has loaded (spec 2026-09-25 §6).
  await expect(page.getByTestId('planner-level')).toHaveCount(0);
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page.getByText(ACTIVE_BUILD_NOTICE)).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
  // Both trees render, but below the md breakpoint only the selected tab's panel is on screen,
  // so Fury is in the document rather than visible -- tests/e2e/planner-phone.spec.ts owns the
  // one-tree-at-a-time rule, and this project runs at a phone width.
  await expect(page.getByRole('heading', { name: 'Fury', includeHidden: true })).toBeAttached();
  expect(errors).toEqual([]);
});

test('Level appears once a build opens from an addon code (a real character)', async ({ page }) => {
  await page.goto('/planner?code=FS1%3A1.15.9.69722%3Awarrior%3Ahuman%3A3%2F0%2F0%3A');
  await expect(page.getByTestId('planner-level')).toBeVisible();
  await expect(page.getByTestId('planner-level')).toHaveText('12');
});

test('a prerequisite link is drawn and turns gold when the rank is met', async ({ page }) => {
  await page.goto('/planner');
  // Tactical Mastery (1004) needs Deflection (1002) at rank 2.
  const link = page.getByTestId('connector-1002-1004');
  await expect(link).toHaveAttribute('data-met', 'false');
  await expect(page.getByTestId('talent-1004')).toHaveAttribute('data-state', 'locked');

  await page.getByTestId('talent-1002').click();
  await expect(link).toHaveAttribute('data-met', 'false');
  await page.getByTestId('talent-1002').click();
  await expect(link).toHaveAttribute('data-met', 'true');
  await expect(page.getByTestId('talent-1002')).toHaveAttribute('data-state', 'filled');
  // Meeting the prerequisite is not enough on its own: tier 1 still needs
  // POINTS_PER_TIER (5) points somewhere in Arms before it opens at all. Three more in
  // Improved Heroic Strike (1001) get there without disturbing the prerequisite above.
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('talent-1004')).toHaveAttribute('data-state', 'available');
});

test('reset says it will ask before it clears the build', async ({ page }) => {
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Reset…' }).click();
  await expect(page.getByText('Clear every point in this build?')).toBeVisible();
  await page.getByRole('button', { name: 'Keep the build' }).click();
  await expect(page.getByTestId('planner-spent')).toHaveText('1/51');
});

test('?class and ?race preselect the build', async ({ page }) => {
  await page.goto('/planner?class=warrior&race=dwarf');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page.getByLabel('Race')).toHaveValue('dwarf');
});

test('switching class and race rewrites the address, so a refresh keeps the choice', async ({ page }) => {
  await page.goto('/planner?class=warrior&race=dwarf');
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await page.getByLabel('Class').selectOption('mage');
  await expect(page).toHaveURL(/\/planner\?class=mage(&race=[a-z]+)?$/);
  await page.getByLabel('Race').selectOption('gnome');
  await expect(page).toHaveURL(/\/planner\?class=mage&race=gnome$/);
  await page.reload();
  await expect(page.getByLabel('Class')).toHaveValue('mage');
  await expect(page.getByLabel('Race')).toHaveValue('gnome');
});

test('a code stays in the address until its class or race is left', async ({ page }) => {
  const code = 'FS1%3A1.15.9.69722%3Awarrior%3Ahuman%3A3%2F0%2F0%3A';
  await page.goto(`/planner?code=${code}`);
  await expect(page.getByLabel('Class')).toHaveValue('warrior');
  await expect(page).toHaveURL(new RegExp(`/planner\\?code=${code}$`));
  await page.getByLabel('Race').selectOption('dwarf');
  await expect(page).toHaveURL(/\/planner\?class=warrior&race=dwarf$/);
});

test('a failed talent fetch shows the reason and a working retry', async ({ page }) => {
  let attempts = 0;
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    attempts += 1;
    if (attempts === 1) return route.abort('failed');
    return route.continue();
  });
  await page.goto('/planner');
  await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
  await page.getByRole('button', { name: 'Try again' }).click();
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
});

test('a slow class switch cannot leave one class holding another class trees', async ({ page }) => {
  // Warrior is the class the planner opens on, so its talent request is the one already in
  // flight when the switch happens. Held open until paladin has finished loading, it is the
  // stale response `load` in Planner.svelte has to discard: without the generation guard it
  // lands last and puts warrior trees under `paladin`, which then posts warrior talent ids
  // with paladin's class_id and is refused by the API.
  let release = (): void => {};
  const held = new Promise<void>((resolve) => {
    release = resolve;
  });
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    await held;
    await route.continue();
  });
  // This build ships talent data for warrior alone, so paladin's file is supplied here. One
  // tier-0 talent is enough: what the assertions read is the tree name, not the grid.
  await page.route('**/data/*/talents/paladin.json', (route) =>
    route.fulfill({
      json: {
        build: ACTIVE_BUILD,
        class_slug: 'paladin',
        class_id: 2,
        trees: [
          {
            id: 382,
            name: 'Holy',
            position: 0,
            talents: [
              {
                id: 3001,
                name: 'Divine Strength',
                icon: 'fixture_divine_strength',
                max_rank: 1,
                tier: 0,
                column: 0,
                prereq_talent_id: null,
                prereq_rank: null,
                ranks: [{ spell_id: 30011, description: 'Increases Strength by 2%.' }],
              },
            ],
          },
        ],
      },
    }),
  );
  // Gear is optional, and a 404 is how the planner is told a class ships no item file.
  await page.route('**/data/*/items/paladin.json', (route) =>
    route.fulfill({ status: 404, contentType: 'text/plain', body: 'not found' }),
  );

  await page.goto('/planner');
  await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
  await page.getByLabel('Class').selectOption('paladin');
  await expect(page.getByRole('heading', { name: 'Holy' })).toBeVisible();

  const warriorResponse = page.waitForResponse('**/data/*/talents/warrior.json');
  release();
  await warriorResponse;
  // The response event fires when the headers arrive; the store write the unguarded code
  // made followed a microtask later, once the body parsed. A discarded response changes
  // nothing and so raises no event to wait on, which is what this settle stands in for.
  await page.waitForTimeout(500);

  await expect(page.getByLabel('Class')).toHaveValue('paladin');
  await expect(page.getByRole('heading', { name: 'Holy' })).toBeVisible();
  await expect(page.getByRole('heading', { name: 'Arms', includeHidden: true })).toHaveCount(0);
  await expect(page.getByRole('heading', { name: 'Fury', includeHidden: true })).toHaveCount(0);
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
  // This is a bare build (no current-character pointer), so Level is not shown here at all
  // (spec 2026-09-25 §6) -- the Level-appears/tracks-a-real-character case is covered by
  // 'Level appears once a build opens from an addon code (a real character)', above.
  await expect(page.getByTestId('planner-level')).toHaveCount(0);
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
  await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
  await page.getByLabel('Class').selectOption('warrior');
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();

  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');
  await expect(page.getByRole('button', { name: 'Clear all points' })).toBeHidden();
  await expect(page.getByRole('button', { name: 'Reset' })).toBeVisible();
});

test('a build opens from an addon code in the URL, with the order noted as reconstructed', async ({
  page,
}) => {
  // The fixture warrior's first tier-0 talent is 1001 with max rank 3; see
  // src/fixtures/planner/talents/warrior.json.
  await page.goto('/planner?code=FS1%3A1.15.9.69722%3Awarrior%3Ahuman%3A3%2F0%2F0%3A');

  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-code-note')).toContainText('not recorded in game');
});

test('a code from another format is refused by name rather than ignored', async ({ page }) => {
  await page.goto('/planner?code=FS2%3A1%3Awarrior%3Ahuman%3A3%2F0%2F0%3A');
  await expect(page.getByTestId('planner-code-note')).toHaveText('That code is FS2; this site reads FS1.');
});

test('a code that spends a locked tier drops it, and the dropped count renders as a wrapped number', async ({
  page,
}) => {
  // Warrior's talent 1005 sits at tier 2 of Arms; asking for a point there with nothing spent
  // in the lower tiers is not reachable by any legal order, so orderFromRanks drops it. "00001"
  // is talent 1005's tab position (the fifth of Arms's seven talents) encoded as base-36 digits
  // with the trailing zeros trimmed; see src/fixtures/planner/talents/warrior.json.
  await page.goto('/planner?code=FS1%3A1.15.9.69722%3Awarrior%3Ahuman%3A00001%2F0%2F0%3A');

  const note = page.getByTestId('planner-code-note');
  await expect(note).toContainText('minus 1 that no legal order reaches');
  await expect(note.locator('span.font-mono')).toHaveText('1');
});

test('a code naming a class with no talent data fails clean, and switching class does not replay it', async ({
  page,
}) => {
  // Paladin is a real class in this build's classes.json but ships no talents/paladin.json
  // fixture -- the same gap the 'a failed talent fetch...' test above uses to force a load
  // failure without a route mock.
  await page.goto('/planner?code=FS1%3A1.15.9.69722%3Apaladin%3Ahuman%3A5%2F0%2F0%3A');

  await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
  await expect(page.getByTestId('planner-code-note')).toHaveText(
    'That code names a class this planner does not have: paladin.',
  );

  // Recovering by picking a class that does have data must not replay the paladin code's tree
  // ranks onto it: codeApplied is armed the moment the first load began, not only once that
  // load succeeds, so warrior comes up with an empty build rather than paladin's digits
  // reinterpreted against warrior's own talent tab positions.
  await page.getByLabel('Class').selectOption('warrior');
  await expect(page.getByRole('heading', { name: 'Arms' })).toBeVisible();
  await expect(page.getByTestId('planner-spent')).toHaveText('0/51');

  // The paladin message is now stale and, worse, no longer true -- warrior's build is working
  // fine -- so switching class must clear it rather than leave it sitting under the new build.
  await expect(page.getByTestId('planner-code-note')).toHaveCount(0);
});

test('a good code survives a transient failure on its own class, and still applies on a same-class Retry', async ({
  page,
}) => {
  // Warrior does have talent data on this build, so this is a network blip on the way to it --
  // nothing to do with the code itself -- unlike the paladin case above, which fails because
  // the class named by the code has no data at all. The code must get a second chance here.
  let attempts = 0;
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    attempts += 1;
    if (attempts === 1) return route.abort('failed');
    return route.continue();
  });

  await page.goto('/planner?code=FS1%3A1.15.9.69722%3Awarrior%3Ahuman%3A3%2F0%2F0%3A');
  await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();

  await page.getByRole('button', { name: 'Try again' }).click();
  await expect(page.getByTestId('talent-1001')).toHaveAttribute('data-rank', '3');
  await expect(page.getByTestId('planner-split')).toHaveText('3/0');
  await expect(page.getByTestId('planner-code-note')).toContainText('not recorded in game');
});

test('each tree draws its own art and counts its own points', async ({ page }) => {
  await page.goto('/planner');
  const art = page.getByTestId('tree-art-161');
  await expect(art).toHaveAttribute('src', `/data/${ACTIVE_BUILD}/trees/fixture_arms.webp`);
  // The art is a real image, not a broken one.
  await expect.poll(() => art.evaluate((node: HTMLImageElement) => node.naturalWidth)).toBeGreaterThan(0);

  await expect(page.getByTestId('tree-points-161')).toHaveText('0');
  await expect(page.getByTestId('planner-remaining')).toHaveText('51');
  await page.getByTestId('talent-1001').click();
  await expect(page.getByTestId('tree-points-161')).toHaveText('1');
  await expect(page.getByTestId('planner-remaining')).toHaveText('50');
});

test('opening /planner bare with a stored, non-restorable (armory) pointer still loads cleanly', async ({
  page,
}) => {
  await page.addInitScript(() => {
    window.localStorage.setItem(
      'fs.currentCharacter',
      JSON.stringify({
        source: 'armory',
        ref: 'us/normal/simfury',
        label: 'Simfury · Fury Warrior',
        classSlug: 'warrior',
        savedAt: '2026-09-25T00:00:00.000Z',
      }),
    );
  });
  await page.goto('/planner');
  await expect(page.getByTestId('current-character-bar')).toBeVisible();
  await expect(page.getByTestId('current-character-restored')).toHaveCount(0);
  await expect(page.getByTestId('planner')).toBeVisible();
});
