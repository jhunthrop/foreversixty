// web/src/lib/report/shell-paths.ts
// Which report pages Astro prerenders. In production: none -- every /reports/<id> is the
// Worker serving dist/reports.html. Under FOREVER_DATA=fixture, two: one complete report
// backed by the checked-in fixture, and one live report the e2e stubs entirely. Those two
// are what Playwright navigates to and what Lighthouse audits, and they must never ship
// in a real build, which is why this reads the same selector scripts/sync-data.mjs does.
import process from 'node:process';

export const FIXTURE_REPORT_ID = 'fixture2abcd';
export const FIXTURE_LIVE_REPORT_ID = 'fixture2live';

export function fixtureReportPaths(): { params: { id: string } }[] {
  if (process.env.FOREVER_DATA !== 'fixture') return [];
  return [{ params: { id: FIXTURE_REPORT_ID } }, { params: { id: FIXTURE_LIVE_REPORT_ID } }];
}
