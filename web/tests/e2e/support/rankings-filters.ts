// web/tests/e2e/support/rankings-filters.ts
import type { Page } from '@playwright/test';

/** On a phone the Rankings filters sit behind the summary line's button; elsewhere they are
 *  always open. Call it before driving any `filter-*` control on `/rankings`. */
export async function openFilters(page: Page): Promise<void> {
  const toggle = page.getByTestId('rankings-filters-toggle');
  if (await toggle.isVisible()) await toggle.click();
}
