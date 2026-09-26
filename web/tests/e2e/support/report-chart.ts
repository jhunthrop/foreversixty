// web/tests/e2e/support/report-chart.ts
// The shared chart block collapses to a one-line strip on a phone by default (design review
// 2026-09-26 finding 1/2): TimeChart itself, its sliders and the window select do not exist
// in the DOM until the "Show chart" toggle is pressed. Every spec that reaches into the
// chart on a phone needs this one open call first; on desktop the chart is already open and
// the toggle is not rendered, so `count()` decides rather than a viewport check the calling
// spec would otherwise have to repeat.
import type { Page } from '@playwright/test';

export async function openChartIfCollapsed(page: Page): Promise<void> {
  const toggle = page.getByTestId('chart-toggle');
  const chart = page.getByTestId('time-chart');
  // Waits for whichever the page actually renders -- called right after `goto`, before
  // the island has necessarily finished its first render, a bare `count()` can read 0
  // just because neither has painted yet, not because the chart is already open.
  await toggle.or(chart).first().waitFor();
  if ((await toggle.count()) > 0) await toggle.click();
}
