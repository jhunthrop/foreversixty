// web/tests/e2e/support/logs-hero-fixture.ts
// The canonical sample report's public `GET /v1/reports/{id}`, which /logs reads on every visit
// to draw the one sample row under "Recent public reports" (lib/logs/sample-row.ts). A spec that
// needs the row to resolve to its real title, date and kills -- rather than the link-only fallback
// an unmocked fetch gives -- routes through here, the same "one fixture, several callers"
// convention active-build.ts and character-fixture.ts already use.
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
  data_base_url: `/sample-report/${SAMPLE_REPORT_ID}`,
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

/** Routes the bundled sample's meta, a plain JSON file beside its data (lib/report/sample.ts). */
export async function routeSampleReportMeta(page: Page): Promise<void> {
  await page.route(`**/sample-report/${SAMPLE_REPORT_ID}/meta.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(SAMPLE_REPORT_META) }),
  );
}
