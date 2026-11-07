// web/tests/e2e/support/report-chart.ts
// The report's chart is an open Section (logs spec 2026-11-07 6.1); below 768 its body folds into
// a Disclosure (`chart-collapsed`) whose plot and controls exist only once it is opened. Every spec
// that reaches into the chart on a phone opens it first; on desktop the chart is already open.
import type { Page } from '@playwright/test';

export async function openChartIfCollapsed(page: Page): Promise<void> {
  const fold = page.getByTestId('chart-collapsed');
  const chart = page.getByTestId('time-chart');
  await chart.first().waitFor();
  if ((await fold.count()) > 0 && (await fold.getAttribute('open')) === null)
    await fold.locator('summary').click();
}
