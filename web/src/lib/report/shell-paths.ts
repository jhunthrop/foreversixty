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

// Task 20's fixture character and guild pages -- one of each, the same pairing
// fixtureReportPaths() and fixtureRankingsPaths() use above: a path here is what
// Playwright navigates to and what the fixture-only [...path].astro route prerenders,
// and it must never ship in a real build.
export const FIXTURE_CHARACTER_PATH = 'us/hardcore/elyra-duskvale';
export const FIXTURE_GUILD_PATH = 'us/hardcore/the-last-watch';

export function fixtureCharacterPaths(): { params: { path: string } }[] {
  return process.env.FOREVER_DATA === 'fixture' ? [{ params: { path: FIXTURE_CHARACTER_PATH } }] : [];
}

export function fixtureGuildPaths(): { params: { path: string } }[] {
  return process.env.FOREVER_DATA === 'fixture' ? [{ params: { path: FIXTURE_GUILD_PATH } }] : [];
}
