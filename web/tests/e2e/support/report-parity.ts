// web/tests/e2e/support/report-parity.ts
// The report states of the parity gate (design/parity.md, design/specs/2026-11-07-logs-reports.md
// section 14) on the bundled real log: /reports/sample on the state's tab, fight and source,
// signed out, every read served from web/public/sample-report/sample. The loading state holds the
// default pull's summary so the island stays on its skeleton.
import type { Page } from '@playwright/test';
import { expect } from './test';
import { signedOut } from './selector';

export type ReportParityState =
  | 'report-summary'
  | 'report-damage-done'
  | 'report-healing'
  | 'report-damage-taken'
  | 'report-deaths'
  | 'report-empty'
  | 'report-loading'
  | 'report-sample';

/** The element each state's ready page shows last. */
const READY: Readonly<Record<Exclude<ReportParityState, 'report-loading'>, string>> = {
  'report-summary': 'summary-tab',
  'report-damage-done': 'report-table-card',
  'report-healing': 'report-table-card',
  'report-damage-taken': 'report-table-card',
  'report-deaths': 'deaths-tab',
  'report-empty': 'report-table-card',
  'report-sample': 'summary-tab',
};

/** Opens one state at `path` (report-parity.json), the report settled or held on its skeleton. */
export async function openReportState(page: Page, id: ReportParityState, path: string): Promise<void> {
  await signedOut(page);
  if (id === 'report-loading') {
    // Never answered: the island keeps its loading skeleton for the capture.
    await page.route('**/sample-report/sample/fights/6/summary.json**', () => undefined);
    await page.goto(path);
    await expect(page.getByTestId('report-skeleton')).toBeVisible();
    return;
  }
  await page.goto(path);
  await expect(page.getByTestId(READY[id])).toBeVisible();
  await expect(page.getByTestId('report-skeleton')).toHaveCount(0);
  // The chart paints after its chunk: wait for its plot so its box is final.
  await expect(page.locator('[data-region^="section-"] [role="img"]').first())
    .toBeVisible()
    .catch(() => undefined);
}
