// web/tests/e2e/support/logs-hero-fixture.ts
// The header hero (LogsHero.svelte, logs landing spec 2026-10-04 §4.A) now fetches on every
// /logs visit regardless of what a given test cares about: `/v1/me`, then either
// `listMyReports(1)` (signed in) or the canonical sample's `GET /v1/reports/{id}` (signed
// out, or signed in with none of their own). A spec that needs the hero to actually render
// -- rather than just tolerate its graceful LoadError on an unmocked fetch -- routes through
// here, the same "one fixture, several callers" convention active-build.ts and
// character-fixture.ts already use.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import type { ReportMeta } from '../../../src/lib/report/types';

const WEB_ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../..');

export const SAMPLE_REPORT_ID: string = (
  JSON.parse(readFileSync(path.join(WEB_ROOT, 'src/data/sample-report.json'), 'utf8')) as { id: string }
).id;

export const SAMPLE_REPORT_META: ReportMeta = {
  id: SAMPLE_REPORT_ID,
  title: 'Sanguine Depths, sample log',
  visibility: 'public',
  owner: null,
  zone: 'Sanguine Depths',
  status: 'complete',
  engine_version: '0.1.0',
  created_at: '2026-09-15T12:55:13Z',
  data_base_url: `/logs-data/reports/${SAMPLE_REPORT_ID}`,
  players: ['Player-1', 'Player-2'],
  fights: [
    {
      index: 1,
      kind: 'trash',
      name: 'Trash',
      kill: false,
      in_progress: false,
      start: '2026-09-15T12:00:00Z',
      end: '2026-09-15T12:01:00Z',
      duration_ms: 60_000,
      players: ['Player-1'],
      deaths: 0,
      npc_kills: 2,
    },
    {
      index: 62,
      kind: 'encounter',
      name: 'Nalthor the Rimebinder',
      encounter_id: 9100,
      difficulty: 8,
      size: 5,
      kill: true,
      in_progress: false,
      start: '2026-09-15T13:40:00Z',
      end: '2026-09-15T13:42:52Z',
      duration_ms: 172_000,
      players: ['Player-1', 'Player-2'],
      deaths: 0,
      npc_kills: 1,
    },
  ],
};

function envelope(data: unknown, status = 200): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'e2e' }),
  };
}

/** Routes the canonical sample's own `GET /v1/reports/{id}` -- the hero's fallback for a
 *  signed-out visitor, or a signed-in one with no reports of their own. */
export async function routeSampleReportMeta(page: Page): Promise<void> {
  await page.route(`**/v1/reports/${SAMPLE_REPORT_ID}`, (route) =>
    route.fulfill(envelope(SAMPLE_REPORT_META)),
  );
}
