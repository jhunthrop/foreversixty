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
