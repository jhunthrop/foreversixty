// web/src/lib/report/shell-paths.ts
// Which report pages Astro prerenders. In production: none -- every /reports/<id> is the
// Worker serving dist/reports.html. Under FOREVER_DATA=fixture, two: one complete report
// backed by the checked-in fixture, and one live report the e2e stubs entirely. Those two
// are what Playwright navigates to and what Lighthouse audits, and they must never ship
// in a real build, which is why this reads the same selector scripts/sync-data.mjs does.
import process from 'node:process';
import { encounterSlug } from '../rankings/api';

export const FIXTURE_REPORT_ID = 'fixture2abcd';
export const FIXTURE_LIVE_REPORT_ID = 'fixture2live';

export function fixtureReportPaths(): { params: { id: string } }[] {
  if (process.env.FOREVER_DATA !== 'fixture') return [];
  return [{ params: { id: FIXTURE_REPORT_ID } }, { params: { id: FIXTURE_LIVE_REPORT_ID } }];
}

// Task 19's fixture rankings page. Derived through the same encounterSlug() the API and
// the report island's own Rankings mode use, rather than a hand-typed literal that could
// drift from it: "Warden Kelthas" becomes "warden-kelthas".
export const FIXTURE_ENCOUNTER_SLUG = encounterSlug('Warden Kelthas');

export function fixtureRankingsPaths(): { params: { slug: string } }[] {
  if (process.env.FOREVER_DATA !== 'fixture') return [];
  return [{ params: { slug: FIXTURE_ENCOUNTER_SLUG } }];
}
