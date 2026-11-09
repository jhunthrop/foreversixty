// web/tests/e2e/support/report-frame.ts
// Driving the report's frame (logs spec 2026-11-07 6.1): the Fight select in the control bar and
// the one tab strip, whose tabs past the first seven sit in its More menu below 1760 px.
import type { Page } from '@playwright/test';
import { choose } from './select';

/** Picks a pull (or 'all', the night's boss pulls) on the bar's Fight select. */
export async function pickFight(page: Page, index: number | 'all'): Promise<void> {
  const select = page.locator('select[data-testid="report-fight"]');
  const value = index === 'all' ? '0' : String(index);
  await select.waitFor({ state: 'attached' });
  // A trash pull that killed nothing is behind the list's last row until it is shown.
  if ((await select.locator(`option[value="${value}"]`).count()) === 0) await choose(select, 'idle');
  await choose(select, value);
}

/** Opens a tab: from the strip when it shows, else from the More menu. */
export async function openTab(page: Page, id: string): Promise<void> {
  const shown = page.locator(`[data-testid="report-tabs"] [data-testid="tab-${id}"]:visible`);
  if ((await shown.count()) > 0) {
    await shown.first().click();
    return;
  }
  await page.locator('[data-testid="report-tabs-more"]:visible').first().click();
  await page.locator(`[role="menuitemradio"][data-testid="tab-${id}"]`).click();
}

/** Picks the bar's Source (a scope word or a player's GUID). */
export async function pickSource(page: Page, value: string): Promise<void> {
  await choose(page.locator('select[data-testid="source-scope"]'), value);
}

/** Turns on one of the Filters switches: in the popover from 768 px, in the open bar's stack below. */
export async function checkFilter(page: Page, which: 'boss'): Promise<void> {
  const button = page.getByTestId('report-filters-button');
  await page
    .locator(
      '[data-testid="report-filters-button"]:visible, [data-testid="report-bar"] [data-control-bar-toggle]:visible',
    )
    .first()
    .waitFor();
  if (await button.isVisible()) {
    await button.click();
    await page.getByTestId(`filter-${which}`).check();
    return;
  }
  await page.getByTestId('report-bar').locator('[data-control-bar-toggle]').click();
  await page.getByTestId(`stack-filter-${which}`).check();
}

/** Below 768 the report bar is one line with Change: opens its stack when it is folded. */
export async function openReportBar(page: Page): Promise<void> {
  const toggle = page.getByTestId('report-bar').locator('[data-control-bar-toggle]');
  if (!(await toggle.isVisible())) return;
  if ((await toggle.getAttribute('aria-expanded')) !== 'true') await toggle.click();
}

/** "Sim this fight": the button in the rail's This pull panel (directions b1), at every width. */
export async function openSimThisFight(page: Page): Promise<void> {
  await page.getByTestId('report-sim-fight').click();
}
