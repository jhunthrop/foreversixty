// web/tests/e2e/report-live-waiting.spec.ts
// The companion opens a report when the game starts a combat log, long before the first fight
// is uploaded. The meta answers and report.json does not exist yet: the page waits for the
// first fight and fills in by itself, and never says "No report with that id".
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');
const ID = 'fixture2live'; // the prerendered shell that carries no inlined meta
const CAPTURES = process.env.CAPTURE_DIR;

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

async function fixtureMeta(): Promise<Record<string, unknown>> {
  return JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8')) as Record<string, unknown>;
}

/** What GET /v1/reports/{id} answers for a log that has been started and has no fight yet. */
async function emptyLiveMeta(credit: { key: string; name: string } | null): Promise<Record<string, unknown>> {
  return {
    ...(await fixtureMeta()),
    id: ID,
    status: 'live',
    title: '',
    zone: '',
    fights: [],
    players: [],
    logging_character: credit,
    created_at: '2026-10-20T09:00:00Z',
  };
}

test('a live report with no fight yet waits for the first fight, then fills in on its own', async ({
  page,
}) => {
  await page.setViewportSize({ width: 1280, height: 800 });
  const live = await emptyLiveMeta({ key: 'classic/era/baelgrim', name: 'Baelgrim' });
  await page.route(`**/v1/reports/${ID}`, (route) => route.fulfill(envelope(live)));

  let reads = 0;
  await page.route('**/logs-data/reports/fixture2abcd/report.json', (route) => {
    reads += 1;
    return reads <= 2 ? route.fulfill({ status: 404, body: '' }) : route.fallback();
  });

  await page.goto(`/reports/${ID}`);
  await page.evaluate(() => ((window as unknown as { __kept: boolean }).__kept = true));

  await expect(page.getByTestId('report-waiting')).toContainText('waiting for the first fight');
  await expect(page.getByTestId('report-waiting')).toContainText('this page updates by itself');
  await expect(page.getByTestId('report-live')).toBeVisible();
  await expect(page.getByTestId('report-title')).toContainText('20 Oct 2026 · 0 fights');
  await expect(page.getByTestId('report-credit')).toContainText('Credited to Baelgrim');
  await expect(page.getByTestId('report-error')).toHaveCount(0);
  await expect(page.getByText('No report with that id')).toHaveCount(0);
  if (CAPTURES) await page.screenshot({ path: path.join(CAPTURES, 'l4-live-waiting-1280.png') });

  // Two polls at five seconds each, plus slack: the first fight arrives and the page switches.
  await expect(page.getByTestId('fight-3-outcome')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('report-processing')).toHaveCount(0);
  expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
});

test('the waiting screen on a phone, with no one credited yet', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  const live = await emptyLiveMeta(null);
  await page.route(`**/v1/reports/${ID}`, (route) => route.fulfill(envelope(live)));
  await page.route('**/logs-data/reports/fixture2abcd/report.json', (route) =>
    route.fulfill({ status: 404, body: '' }),
  );

  await page.goto(`/reports/${ID}`);
  await expect(page.getByTestId('report-waiting')).toBeVisible();
  await expect(page.getByTestId('report-credit')).toContainText('Not credited yet');
  if (CAPTURES) await page.screenshot({ path: path.join(CAPTURES, 'l4-live-waiting-390.png') });
});

test('a finished report whose data is missing is a data error, not a missing report', async ({ page }) => {
  const complete = { ...(await fixtureMeta()), id: ID, status: 'complete' };
  await page.route(`**/v1/reports/${ID}`, (route) => route.fulfill(envelope(complete)));
  await page.route('**/logs-data/reports/fixture2abcd/report.json', (route) =>
    route.fulfill({ status: 404, body: '' }),
  );
  await page.goto(`/reports/${ID}`);
  await expect(page.getByTestId('report-error')).toContainText("This report's data is missing");
  await expect(page.getByText('No report with that id')).toHaveCount(0);
});
