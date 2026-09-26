// web/tests/e2e/support/planner.ts
// Moves shared by the planner specs, so the phone layout is spelled out in one place.
import { expect, type Page } from '@playwright/test';

/** Tailwind's `md`. Below it the planner shows one panel at a time behind the tab strip. */
const MD_BREAKPOINT = 768;

/**
 * Brings the gear panel on screen. From md up it is always directly under the tree row (in
 * the rail layout's own left column from lg, or last in the stack between md and lg); below
 * md it is the last tab, so it has to be selected first. Both Playwright projects run the
 * gear specs and the mobile one is narrower than md, so every gear move goes through here
 * rather than assuming the panel is already on screen.
 */
export async function openGear(page: Page): Promise<void> {
  if ((page.viewportSize()?.width ?? MD_BREAKPOINT) >= MD_BREAKPOINT) {
    await expect(page.getByTestId('gear-panel')).toBeVisible();
    return;
  }
  await page.getByRole('tab', { name: 'Gear' }).click();
}

/**
 * Opens the point order panel's content. From md up it is always open (design loop, planner
 * round); below md it is a native, closed-by-default `<details>`, so its own summary has to
 * be clicked first. Every spec that reads the strip's list goes through here rather than
 * assuming the disclosure is already open on a phone-width viewport.
 */
export async function openOrderStrip(page: Page): Promise<void> {
  await openDisclosure(page, 'order-strip');
}

/**
 * Opens the import box's content. From md up it is always open; below md it is a native,
 * closed-by-default `<details>` (design loop, planner round), so a spec that fills or clicks
 * inside it on a phone-width viewport has to open the disclosure first -- Playwright refuses
 * to act on an element a closed `<details>` hides.
 */
export async function openImportBox(page: Page): Promise<void> {
  await openDisclosure(page, 'import-box');
}

async function openDisclosure(page: Page, testId: string): Promise<void> {
  if ((page.viewportSize()?.width ?? MD_BREAKPOINT) >= MD_BREAKPOINT) return;
  const el = page.getByTestId(testId);
  if (await el.evaluate((node) => node.tagName === 'DETAILS' && !(node as HTMLDetailsElement).open)) {
    await el.locator('summary').click();
  }
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

/**
 * Share, all the way to a saved build. Task 11: the Share button (`share-open`) no longer
 * saves on its own click -- it opens a confirm step naming what becomes public, and only
 * "Share anyway" (`share-confirm-proceed`) actually calls `saveBuild`. Every planner spec
 * that used to do `getByRole('button', { name: 'Share' }).click()` and land straight on a
 * saved link goes through here now, so the confirm step is exercised (not bypassed) by
 * every one of them rather than each spec growing its own two-click copy.
 */
export async function shareBuild(page: Page): Promise<void> {
  await page.getByTestId('share-open').click();
  await expect(page.getByTestId('share-confirm')).toBeVisible();
  await page.getByTestId('share-confirm-proceed').click();
}
