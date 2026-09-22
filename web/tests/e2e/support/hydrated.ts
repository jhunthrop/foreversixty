// web/tests/e2e/support/hydrated.ts
import { expect, type Page } from '@playwright/test';

/**
 * Navigate to a page and wait until the island holding `testId` has hydrated.
 *
 * A control driven before its island hydrates fires its event at static HTML and no
 * handler runs: the upload test and the rankings filter tests both failed this way under
 * a loaded machine, each intermittently. Astro marks a not-yet-hydrated island with an
 * `ssr` attribute and removes it once the component has mounted, so waiting for that
 * attribute to go is the one check that means "your click will be heard".
 */
export async function gotoHydrated(page: Page, url: string, testId: string): Promise<void> {
  await page.goto(url);
  await expect(page.getByTestId(testId).first()).toBeVisible();
  await expect(page.locator(`astro-island[ssr]:has([data-testid="${testId}"])`)).toHaveCount(0);
}
