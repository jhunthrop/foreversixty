import { expect, test, type Locator, type Page } from '@playwright/test';
import { openGear } from './support/planner';

/**
 * Holds the talent request open so the planner's loading state can be measured, and returns
 * the release. `outcome` decides which state it lands in: 'continue' serves the data, 'abort'
 * drops it and sends the island down its failure branch.
 */
async function holdTalents(page: Page, outcome: 'continue' | 'abort'): Promise<() => void> {
  let land = (): void => {};
  const held = new Promise<void>((resolve) => {
    land = resolve;
  });
  await page.route('**/data/*/talents/warrior.json', async (route) => {
    await held;
    return outcome === 'abort' ? route.abort('failed') : route.continue();
  });
  return () => land();
}

const footerTop = async (page: Page): Promise<number> => (await page.locator('footer').boundingBox())?.y ?? 0;

// The reserve is measured from the loaded layout rather than derived from it, so it sits a
// fraction under the real height and a pixel or two of travel is expected. A jump is not.
const SETTLED_PX = 4;

// design/DESIGN-SYSTEM.md: 44px minimum hit targets, 18px phone gutter. 360x800 is the
// narrowest phone the site designs for; design/Mobile.dc.html is the reference artboard.
test.describe('planner on a phone', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  test('shows one panel at a time behind a tab switcher', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    // The two warrior trees and the gear panel. Named as well as counted, so a tab that went
    // missing cannot be covered for by one that arrived, and so the order is pinned: the gear
    // tab is last because Planner.svelte derives its index from the tree count.
    await expect(tabs).toHaveCount(3);
    await expect(tabs.nth(0)).toHaveAccessibleName(/^Arms/);
    await expect(tabs.nth(1)).toHaveAccessibleName(/^Fury/);
    await expect(tabs.nth(2)).toHaveAccessibleName('Gear');
    await expect(tabs.nth(0)).toHaveAttribute('aria-selected', 'true');
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeHidden();

    await tabs.nth(1).click();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeHidden();
  });

  // The roving tabindex takes the unselected tabs out of the tab order, so arrow keys are the
  // only way a keyboard can reach anything past the first tab at all.
  test('arrow keys move between the tabs and open the panel they land on', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    await tabs.first().focus();

    await page.keyboard.press('ArrowRight');
    await expect(tabs.nth(1)).toBeFocused();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();

    // The walk carries on to the end of the strip: the gear tab is behind the same roving
    // tabindex as the trees, so arrow keys are the only way a keyboard reaches it at all.
    await page.keyboard.press('ArrowRight');
    await expect(tabs.nth(2)).toBeFocused();
    await expect(page.getByTestId('gear-panel')).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeHidden();

    await page.keyboard.press('ArrowLeft');
    await expect(tabs.nth(1)).toBeFocused();
    await page.keyboard.press('ArrowLeft');
    await expect(tabs.nth(0)).toBeFocused();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
  });

  test('keeps the summary bar in view while the trees scroll', async ({ page }) => {
    await page.goto('/planner');
    await page.getByRole('grid', { name: 'Arms talents' }).waitFor();
    // `mouse.wheel` starts a scroll and returns before it lands, so the position has to be
    // waited for -- reading it straight after measures the unscrolled page. A phone shows one
    // tree, so 600px may be further than the page goes; wait for as far as it does go.
    const target = await page.evaluate(() =>
      Math.min(600, document.documentElement.scrollHeight - window.innerHeight),
    );
    await page.mouse.wheel(0, 600);
    await expect.poll(() => page.evaluate(() => window.scrollY)).toBeGreaterThanOrEqual(target - 1);

    // planner-remaining, not planner-level: a bare /planner (this test's own load) has no
    // current-character pointer, so Level does not render at all (spec 2026-09-25 §6) --
    // Points left sits in the same summary bar row and is always present.
    const box = await page.getByTestId('planner-remaining').boundingBox();
    expect(box?.y ?? -1).toBeGreaterThanOrEqual(0);
    expect(box?.y ?? 9999).toBeLessThan(200);
  });

  test('never scrolls sideways, with or without points spent', async ({ page }) => {
    await page.goto('/planner');
    const overflow = async () =>
      page.evaluate(() => document.documentElement.scrollWidth - document.documentElement.clientWidth);
    expect(await overflow()).toBeLessThanOrEqual(0);
    for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
    await page.getByTestId('talent-1002').click();
    expect(await overflow()).toBeLessThanOrEqual(0);
  });

  // /planner.html carries a CLS budget of 0.05 (lighthouserc.json) and Lighthouse audits the
  // phone. The island paints a placeholder and then swaps in a planner the better part of a
  // thousand pixels tall, so with nothing reserved everything under it -- the footer above all
  // -- drops down the page when the talent data lands. That swap measured 0.185 of the page's
  // 0.186 CLS. These two tests hold the request open and measure the footer across the swap.
  test('the footer stays put when the talent data lands', async ({ page }) => {
    const land = await holdTalents(page, 'continue');

    await page.goto('/planner');
    await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
    const loading = await footerTop(page);

    land();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    expect(Math.abs((await footerTop(page)) - loading)).toBeLessThanOrEqual(SETTLED_PX);
  });

  // The reserve has to cover the failure branch too. It is much the shortest of the three
  // states, so a reserve that applied only while loading would let the region collapse here
  // and haul the footer back up into the viewport -- the same bug, in the visible direction.
  test('the footer stays put when the talent data fails to load', async ({ page }) => {
    const land = await holdTalents(page, 'abort');

    await page.goto('/planner');
    await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
    const loading = await footerTop(page);

    land();
    await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
    expect(Math.abs((await footerTop(page)) - loading)).toBeLessThanOrEqual(SETTLED_PX);
  });

  test('every talent cell, tab, button and gear target clears 44px', async ({ page }) => {
    await page.goto('/planner');
    const clears44 = async (locator: Locator): Promise<void> => {
      const box = await locator.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
    };

    // Every tab rather than the first: the gear tab is built from the same recipe as the
    // tree tabs but not from the same loop, so it is free to drift on its own.
    for (const tab of await page.getByRole('tab').all()) await clears44(tab);

    for (const locator of [
      page.getByTestId('talent-1001'),
      page.getByRole('button', { name: 'Reset' }),
      page.getByLabel('Class'),
    ]) {
      await clears44(locator);
    }

    // The point-order toggle exists only once a point is spent (an empty strip has nothing to
    // hide), so spend one before measuring it.
    await page.getByTestId('talent-1001').click();
    await clears44(page.getByRole('button', { name: 'Hide point order' }));

    // A gear slot and an item row are hit with a finger like everything else here, and the
    // rows are the narrowest thing the planner asks anyone to tap, so the picker is opened
    // for the measurement rather than left out of it. It comes after the talent cell because
    // the gear tab hides the tree: a locator measured from the wrong tab has no box at all.
    await openGear(page);
    await page.getByTestId('slot-head').click();
    await clears44(page.getByTestId('slot-head'));
    await clears44(page.getByTestId('item-16963'));
  });
});

test('a tree and its links fit a phone with no sideways scroll', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 800 });
  await page.goto('/planner');
  await expect(page.getByTestId('tree-panel-161')).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
  const grid = page.getByTestId('tree-161');
  const box = await grid.boundingBox();
  expect(box!.width).toBeLessThanOrEqual(390 - 36);
  await expect(page.getByTestId('connector-1002-1004')).toBeAttached();
});
