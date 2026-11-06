// web/tests/e2e/support/logs-parity.ts
// The logs, reports list, rankings and ratings states of the parity gate (design/parity.md,
// design/specs/2026-11-07-logs-reports.md section 14), served from the files the mocks read:
// web/src/fixtures/reports (production's z2opt2x2gs5u as the signed-in list and the public
// feed), the rankings stubs (support/rankings-fixture.ts) and the bundled sample report. The
// clock is the fixture's own day, so ages read as the mock's.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import type { Page } from '@playwright/test';
import { expect } from './test';
import { stubRankingsApi } from './rankings-fixture';
import { OBNOXIOUS, armoryPointer, signInWith, signedOut, storePointer } from './selector';

const FIXTURES = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../src/fixtures/reports');

interface ReportsFixture {
  now: string;
  me: Record<string, unknown>;
  page: { rows: unknown[]; total: number; page: number; per_page: number };
}

function readFixture<T>(file: string): T {
  return JSON.parse(readFileSync(path.join(FIXTURES, file), 'utf8')) as T;
}

export const MINE = readFixture<ReportsFixture>('mine.json');
export const RECENT = readFixture<{ page: { rows: unknown[] } }>('recent.json');

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { message: 'failed' },
    request_id: 'r',
  }),
});

export type LogsParityState =
  | 'logs'
  | 'logs-empty'
  | 'logs-you'
  | 'reports'
  | 'reports-empty'
  | 'reports-list'
  | 'reports-failed'
  | 'rankings'
  | 'rankings-empty'
  | 'rankings-board'
  | 'rankings-you'
  | 'rankings-off'
  | 'rankings-unranked'
  | 'ratings';

/** The fixture's signed-in account (no characters), its list and the public feed. */
async function stubReports(page: Page, list: 'none' | 'some' | 'fail' | 'out'): Promise<void> {
  if (list === 'out') await signedOut(page);
  else await page.route('**/v1/me', (route) => route.fulfill(envelope(MINE.me)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      list === 'fail'
        ? envelope(null, 500)
        : envelope(list === 'some' ? MINE.page : { ...MINE.page, rows: [], total: 0 }),
    ),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(envelope(RECENT.page)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
}

async function openLogs(page: Page, id: LogsParityState): Promise<void> {
  const list = { logs: 'out', 'logs-empty': 'none', 'logs-you': 'some' } as const;
  await stubReports(page, list[id as keyof typeof list]);
  await page.goto('/logs');
  await expect(page.getByTestId('recent-reports-sample')).toBeVisible();
  await expect(page.getByTestId('recent-reports-skeleton')).toHaveCount(0);
  if (id === 'logs-you') await expect(page.getByTestId('report-row').first()).toBeVisible();
  if (id === 'logs-empty') await expect(page.getByTestId('my-reports-empty')).toBeVisible();
  if (id !== 'logs') await expect(page.getByTestId('logs-companion-pair')).toBeVisible();
}

async function openReports(page: Page, id: LogsParityState): Promise<void> {
  const list = {
    reports: 'out',
    'reports-empty': 'none',
    'reports-list': 'some',
    'reports-failed': 'fail',
  } as const;
  await stubReports(page, list[id as keyof typeof list]);
  await page.goto('/reports');
  const ready = {
    reports: 'your-reports-signin',
    'reports-empty': 'my-reports-empty',
    'reports-list': 'report-row',
    'reports-failed': 'my-reports-error',
  } as const;
  await expect(page.getByTestId(ready[id as keyof typeof ready]).first()).toBeVisible();
}

async function openRankings(page: Page, id: LogsParityState): Promise<void> {
  if (id === 'rankings' || id === 'rankings-empty' || id === 'rankings-board') {
    await signedOut(page);
    if (id === 'rankings-empty') {
      await page.route('**/v1/encounters', (route) => route.fulfill(envelope({ rows: [] })));
    } else {
      await stubRankingsApi(page);
    }
    await page.goto(id === 'rankings-board' ? '/rankings/warden-kelthas' : '/rankings');
    await expect(
      page.getByTestId(
        {
          rankings: 'rankings-encounters',
          'rankings-empty': 'rankings-no-encounters',
          'rankings-board': 'ranking-1',
        }[id],
      ),
    ).toBeVisible();
    return;
  }
  await signInWith(page, [OBNOXIOUS]);
  await storePointer(page, armoryPointer(OBNOXIOUS));
  await stubRankingsApi(page, {
    ownOffPage: id === 'rankings-off',
    unranked: id === 'rankings-unranked',
  });
  // Signed in, /rankings opens the encounter the character last killed (Onyxia).
  await page.goto('/rankings');
  await expect(page.getByTestId('ranking-1')).toBeVisible();
  await expect(page.getByTestId('rankings-answer')).toBeVisible();
}

/** Opens one state, every read answered from the fixtures, the clock on the fixture's day. */
export async function openLogsState(page: Page, id: LogsParityState): Promise<void> {
  await page.clock.setFixedTime(new Date(MINE.now));
  if (id.startsWith('logs')) return openLogs(page, id);
  if (id.startsWith('reports')) return openReports(page, id);
  if (id.startsWith('rankings')) return openRankings(page, id);
  await page.goto('/ratings');
  await expect(page.getByRole('heading', { level: 1 })).toBeVisible();
}
