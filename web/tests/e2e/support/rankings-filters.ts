// web/tests/e2e/support/rankings-filters.ts
// Driving the Rankings ControlBar (design/specs/2026-11-07-logs-reports.md section 7). From 768 px
// the bar's controls are always on screen and Region, Phase, Faction and Today only sit behind the
// Filters popover; below 768 px the bar is one summary line whose toggle opens the stack, the same
// four fields inline in it, and a choice in the bar's own groups collapses it again.
import type { Locator, Page } from '@playwright/test';

/** The extra fields behind Filters (`rankings-pop-*` in the popover, `rankings-stack-*` on a phone). */
export type MoreFilter = 'region' | 'phase' | 'faction' | 'metric' | 'today';

/** The phone bar's summary toggle: hidden from 768 px. */
export function barToggle(page: Page): Locator {
  return page.getByTestId('rankings-filters').locator('[data-control-bar-toggle]');
}

/** Opens the phone bar's stack when it is collapsed; from 768 px the controls are already shown.
 *  Call it before driving any `control-*` on `/rankings`. */
export async function openFilters(page: Page): Promise<void> {
  const toggle = barToggle(page);
  if ((await toggle.isVisible()) && (await toggle.getAttribute('aria-expanded')) !== 'true') {
    await toggle.click();
  }
}

/** One of the fields behind Filters, opened wherever it lives at this width. */
export async function moreFilter(page: Page, field: MoreFilter): Promise<Locator> {
  await openFilters(page);
  const button = page.getByTestId('rankings-filters-button');
  if (!(await button.isVisible())) return page.getByTestId(`rankings-stack-${field}`);
  if ((await button.getAttribute('aria-expanded')) !== 'true') await button.click();
  return page.getByTestId(`rankings-pop-${field}`);
}
