import { expect, test } from '@playwright/test';

// design/DESIGN-SYSTEM.md: 44px minimum hit targets, 18px phone gutter. 360x800 is the
// narrowest phone the site designs for; design/Mobile.dc.html is the reference artboard.
test.describe('planner on a phone', () => {
  test.use({ viewport: { width: 360, height: 800 } });

  test('shows one tree at a time behind a tab switcher', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    await expect(tabs).toHaveCount(2);
    await expect(tabs.nth(0)).toHaveAttribute('aria-selected', 'true');
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeHidden();

    await tabs.nth(1).click();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeHidden();
  });

  // The roving tabindex takes the unselected tabs out of the tab order, so arrow keys are the
  // only way a keyboard can reach the second tree at all.
  test('arrow keys move between the tabs and open the tree they land on', async ({ page }) => {
    await page.goto('/planner');
    const tabs = page.getByRole('tab');
    await tabs.first().focus();

    await page.keyboard.press('ArrowRight');
    await expect(tabs.nth(1)).toBeFocused();
    await expect(page.getByRole('grid', { name: 'Fury talents' })).toBeVisible();

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

    const box = await page.getByTestId('planner-level').boundingBox();
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

  // /planner.html carries a CLS budget of 0.05 (lighthouserc.json), and Lighthouse audits the
  // phone. The island paints a one-line "Loading talent data" placeholder and then replaces it
  // with the whole planner, so unless the loading state reserves the room, everything under it
  // -- the footer above all -- drops down the page when the talent data lands. A shift only
  // counts against CLS where it can be seen, so the rule this pins down is: while the planner
  // is loading, nothing below it is on screen to be pushed.
  test('nothing below the planner is on screen while it loads', async ({ page }) => {
    let land = (): void => {};
    const held = new Promise<void>((resolve) => {
      land = resolve;
    });
    await page.route('**/data/*/talents/warrior.json', async (route) => {
      await held;
      await route.continue();
    });

    await page.goto('/planner');
    await expect(page.getByText('Loading talent data')).toBeVisible();
    const footer = page.locator('footer');
    const viewport = page.viewportSize()?.height ?? 0;
    const loading = await footer.boundingBox();
    expect(loading?.y ?? 0).toBeGreaterThanOrEqual(viewport);

    land();
    await expect(page.getByRole('grid', { name: 'Arms talents' })).toBeVisible();
    // The reserve is never more than the ready planner goes on to need, so the footer only
    // ever settles further down -- it cannot rebound up into the viewport either.
    const ready = await footer.boundingBox();
    expect(ready?.y ?? 0).toBeGreaterThanOrEqual(loading?.y ?? 0);
  });

  test('every talent cell, tab and button clears 44px', async ({ page }) => {
    await page.goto('/planner');
    for (const locator of [
      page.getByTestId('talent-1001'),
      page.getByRole('tab').first(),
      page.getByRole('button', { name: 'Reset' }),
      page.getByRole('button', { name: 'Hide point order' }),
      page.getByLabel('Class'),
    ]) {
      const box = await locator.boundingBox();
      expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
    }
  });
});
