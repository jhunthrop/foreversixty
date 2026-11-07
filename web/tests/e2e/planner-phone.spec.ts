import type { Locator, Page } from '@playwright/test';
import { expect, test } from './support/test';
import { fieldBox } from './support/select';
import { openPlannerBar, pointsLeft, treePanel } from './support/planner';

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

  // The planner (In the game, the header selector) makes its own `/v1/me` request,
  // independent of the talent-data load these tests measure. Left unrouted, that request
  // hits the real (unreachable in this suite) API and resolves on its own schedule, which is
  // exactly the kind of unrelated, timing-dependent change the footer-stability tests below
  // exist to rule out. Routing it here as a fast, deterministic signed-out response keeps it
  // out of a window these tests are not about.
  test.beforeEach(async ({ page }) => {
    await page.route('**/v1/me', (route) =>
      route.fulfill({
        status: 401,
        contentType: 'application/json',
        body: JSON.stringify({ ok: false, data: null, error: null, request_id: 'e2e' }),
      }),
    );
  });

  test('shows one panel at a time behind a tab switcher', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    // The phone strip switches between the talent trees alone. Named as well as counted, so a
    // tab that went missing cannot be covered for by one that arrived. Each tab reads its
    // points then its tree's name (spec 2026-11-08 3A, TalentEditor).
    await expect(tabs).toHaveCount(2);
    await expect(tabs.nth(0)).toHaveAccessibleName(/Arms$/);
    await expect(tabs.nth(1)).toHaveAccessibleName(/Fury$/);
    await expect(tabs.nth(0)).toHaveAttribute('aria-selected', 'true');
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeHidden();

    await tabs.nth(1).click();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeHidden();
  });

  // Polish spec 5.3 asked a clipped tab strip to fade its hidden edge and scroll the open tab
  // into view. The strip of spec 2026-11-08 (TalentEditor) cannot clip: the tabs share its width
  // equally, so the narrowest phone shows every tab whole -- which is what the fade was for.
  test('every tree tab sits whole inside the strip, and the open one is marked', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    await expect(tabs).toHaveCount(2);
    const strip = page.getByRole('tablist');
    const edge = (await strip.boundingBox())!;
    for (const tab of await tabs.all()) {
      const box = (await tab.boundingBox())!;
      expect(box.x).toBeGreaterThanOrEqual(edge.x - 1);
      expect(box.x + box.width).toBeLessThanOrEqual(edge.x + edge.width + 1);
    }
    await tabs.nth(1).click();
    await expect(tabs.nth(1)).toHaveAttribute('aria-selected', 'true');
    await expect(tabs.nth(0)).toHaveAttribute('aria-selected', 'false');
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

    // The walk wraps at the end of the (now talent-tree-only, §11 owner ruling) strip.
    await page.keyboard.press('ArrowRight');
    await expect(tabs.nth(0)).toBeFocused();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeHidden();

    await page.keyboard.press('ArrowLeft');
    await expect(tabs.nth(1)).toBeFocused();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();
  });

  // The head is a PageHeader on the page ground (spec 2026-11-08 4.1): points left is on the
  // first screen, and loading the trees never moves the page (the open tab used to pull the
  // page down 95 px).
  test('opens on the first screen with points left, and loading the trees never scrolls the page', async ({
    page,
  }) => {
    await page.goto('/planner');
    await page.getByRole('grid', { name: 'Arms talents' }).waitFor();
    expect(await page.evaluate(() => window.scrollY)).toBe(0);
    const box = await pointsLeft(page).boundingBox();
    const viewport = page.viewportSize();
    expect(box?.y ?? -1).toBeGreaterThanOrEqual(0);
    expect((box?.y ?? 9999) + (box?.height ?? 0)).toBeLessThanOrEqual(viewport?.height ?? 0);
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
  // phone. The island paints a placeholder and then swaps in the trees, so with nothing
  // reserved everything under the placeholder dropped down the page when the talent data
  // landed (0.185 of the page's 0.186 CLS). Spec 2026-11-08 5.2 moved the reserve onto the
  // trees' Plate itself: loading, failed and ready stand the same height, so the Plate and the
  // section under it hold still; the sections that only a ready planner draws (Point order,
  // Share, Import) arrive below the first screen, where no shift is visible. These two tests
  // hold the request open and measure across the swap.
  const plate = (page: Page): Locator => page.getByRole('region', { name: 'Talent trees' });
  const boxOf = async (locator: Locator): Promise<{ y: number; height: number }> => {
    const box = await locator.boundingBox();
    return { y: box?.y ?? -1, height: box?.height ?? -1 };
  };
  const expectBelowFirstScreen = async (page: Page): Promise<void> => {
    expect(await footerTop(page)).toBeGreaterThanOrEqual(page.viewportSize()?.height ?? 0);
  };

  test('the trees and the section under them stay put when the talent data lands', async ({ page }) => {
    const land = await holdTalents(page, 'continue');

    await page.goto('/planner');
    await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
    // The retired character card's loading placeholder was waited out here (fix round 4). Its
    // successor, In the game, mounts with the trees and never during this window: while the
    // talent data is held the rail shows only fixed-size skeletons, signed in or out.
    const band = page.getByRole('heading', { name: 'Level 60 build' });
    const loadingPlate = await boxOf(plate(page));
    const loadingBand = await boxOf(band);
    await expectBelowFirstScreen(page);

    land();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    const readyPlate = await boxOf(plate(page));
    expect(Math.abs(readyPlate.y - loadingPlate.y)).toBeLessThanOrEqual(SETTLED_PX);
    expect(Math.abs(readyPlate.height - loadingPlate.height)).toBeLessThanOrEqual(SETTLED_PX);
    expect(Math.abs((await boxOf(band)).y - loadingBand.y)).toBeLessThanOrEqual(SETTLED_PX);
    await expectBelowFirstScreen(page);
  });

  // The reserve has to cover the failure branch too. It is much the shortest of the three
  // states, so a reserve that applied only while loading would let the region collapse here
  // and haul what follows back up into the viewport -- the same bug, in the visible direction.
  test('the trees stay put when the talent data fails to load', async ({ page }) => {
    const land = await holdTalents(page, 'abort');

    await page.goto('/planner');
    await expect(page.getByTestId('planner-talent-skeleton')).toBeVisible();
    const loadingPlate = await boxOf(plate(page));

    land();
    await expect(page.getByText('Talent data did not load', { exact: true })).toBeVisible();
    const failedPlate = await boxOf(plate(page));
    expect(Math.abs(failedPlate.y - loadingPlate.y)).toBeLessThanOrEqual(SETTLED_PX);
    expect(Math.abs(failedPlate.height - loadingPlate.height)).toBeLessThanOrEqual(SETTLED_PX);
    await expectBelowFirstScreen(page);
  });

  test('every talent cell, tab, class crest, select and button clears 44px', async ({ page }) => {
    await page.goto('/planner');
    const clears44 = async (locator: Locator): Promise<void> => {
      const box = await locator.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
    };

    for (const tab of await page.getByRole('tab').all()) await clears44(tab);

    // Reset shows once a point is spent (spec 2026-11-08 4.4).
    await page.getByTestId('talent-1001').click();
    for (const locator of [
      page.getByTestId('talent-1001'),
      page.getByTestId('planner-reset'),
      page.getByTestId('planner-bar').locator('[data-control-bar-toggle]'),
    ]) {
      await clears44(locator);
    }

    // The ControlBar is one line on a phone; open, the nine crests and the race field.
    await openPlannerBar(page);
    for (const crest of await page.getByTestId('control-class').getByRole('radio').all()) {
      await clears44(crest);
    }
    await clears44(fieldBox(page, 'control-race'));
  });
});

test('a tree and its links fit a phone with no sideways scroll', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 800 });
  await page.goto('/planner');
  await expect(treePanel(page, 'Arms')).toBeVisible();
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
  const grid = page.getByTestId('tree-161');
  const box = await grid.boundingBox();
  expect(box!.width).toBeLessThanOrEqual(390 - 36);
  // Tactical Mastery's link to Deflection, the first of Arms's two (sorted by prerequisite id).
  await expect(grid.locator('path[data-met]').first()).toBeAttached();
});
