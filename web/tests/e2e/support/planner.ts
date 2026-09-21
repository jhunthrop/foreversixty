// web/tests/e2e/support/planner.ts
// Moves shared by the planner specs, so the phone layout is spelled out in one place.
import { expect, type Page } from '@playwright/test';

/** Tailwind's `md`. Below it the planner shows one panel at a time behind the tab strip. */
const MD_BREAKPOINT = 768;

/**
 * Brings the gear panel on screen. From md up it is always in the column under the order
 * strip; below md it is the last tab, so it has to be selected first. Both Playwright
 * projects run the gear specs and the mobile one is narrower than md, so every gear move
 * goes through here rather than assuming the panel is already on screen.
 */
export async function openGear(page: Page): Promise<void> {
  if ((page.viewportSize()?.width ?? MD_BREAKPOINT) >= MD_BREAKPOINT) {
    await expect(page.getByTestId('gear-panel')).toBeVisible();
    return;
  }
  await page.getByRole('tab', { name: 'Gear' }).click();
}

/**
 * The fixture warrior one point short of a whole build: Arms full (24), Fury at 26 with
 * Flurry (talent 2007) empty. The live estimate waits for all 51 points, so a spec that
 * wants a figure starts here and spends the last point with `finishBuild`.
 */
export const NEARLY_FINISHED_BUILD = `/planner?code=${encodeURIComponent(
  'FS1:1.15.9.69722:warrior:human:3535125/5555510/0:',
)}`;
export const LAST_TALENT = 'talent-2007';

/** Below md the planner shows one tree at a time, so a talent's tree is brought up first. */
export async function showTree(page: Page, tree: 'Arms' | 'Fury'): Promise<void> {
  if ((page.viewportSize()?.width ?? MD_BREAKPOINT) < MD_BREAKPOINT) {
    await page.getByRole('tab', { name: tree }).click();
  }
}

/** `LAST_TALENT` is in the second tree. */
export async function showLastTalent(page: Page): Promise<void> {
  await showTree(page, 'Fury');
  await expect(page.getByTestId(LAST_TALENT)).toBeVisible();
}

/**
 * Spends the last point and brings the estimate on. A touch-first device (the mobile
 * project) is asked before anything runs, so it presses "Show DPS"; a desktop is not.
 */
export async function finishBuild(page: Page): Promise<void> {
  await showLastTalent(page);
  await page.getByTestId(LAST_TALENT).click();
  await expect(page.getByTestId('planner-remaining')).toHaveText('0');
  const show = page.getByTestId('planner-dps-show');
  if (await show.isVisible()) await show.click();
}
