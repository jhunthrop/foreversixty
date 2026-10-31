// web/tests/e2e/report-processing.spec.ts
// A report being parsed is a real screen, polled, and filled in without a reload; and the
// 404 right after an upload is "not visible yet" rather than "No report with that id".
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { expect, test } from './support/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');
const ID = 'fixture2live'; // the prerendered shell that carries no inlined meta

const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
});

async function meta(): Promise<Record<string, unknown>> {
  return JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8')) as Record<string, unknown>;
}

test('a processing report shows the parsing screen, then fills in without navigation', async ({ page }) => {
  const complete = { ...(await meta()), id: ID, status: 'complete' };
  const processing = { ...complete, status: 'processing', title: '', zone: '', fights: [], players: [] };

  let calls = 0;
  await page.route(`**/v1/reports/${ID}`, (route) => {
    calls += 1;
    return route.fulfill(envelope(calls <= 2 ? processing : complete));
  });

  await page.goto(`/reports/${ID}`);
  await page.evaluate(() => ((window as unknown as { __kept: boolean }).__kept = true));

  await expect(page.getByTestId('report-processing')).toBeVisible();
  await expect(page.getByTestId('report-title')).toHaveText('Your upload');
  await expect(page.getByTestId('report-parsing')).toContainText('Parsing your log');
  await expect(page.getByTestId('report-spinner')).toBeVisible();
  const tab = page.getByTestId('report-tabs-disabled').getByRole('button').first();
  await expect(tab).toBeDisabled();
  await expect(tab).toHaveAttribute('title', /Parsing your log/);

  // Two polls at five seconds each, plus slack.
  await expect(page.getByTestId('fight-3-outcome')).toBeVisible({ timeout: 25_000 });
  await expect(page.getByTestId('report-processing')).toHaveCount(0);
  await expect(page.getByTestId('report-title')).not.toHaveText('Your upload');
  // Same document: nothing navigated or reloaded.
  expect(await page.evaluate(() => (window as unknown as { __kept?: boolean }).__kept)).toBe(true);
  expect(calls).toBeGreaterThanOrEqual(3);
});

test('a failed parse says so and points at the logging setup', async ({ page }) => {
  await page.route(`**/v1/reports/${ID}`, async (route) =>
    route.fulfill(envelope({ ...(await meta()), id: ID, status: 'failed', fights: [] })),
  );
  await page.goto(`/reports/${ID}`);
  await expect(page.getByTestId('report-parse-failed')).toContainText('This log could not be parsed.');
  await expect(page.getByTestId('report-parse-failed-link')).toHaveAttribute('href', '/setup#logging');
});

test('a 404 right after the upload is retried, not shown as a missing report', async ({ page }) => {
  const complete = { ...(await meta()), id: ID, status: 'complete' };
  await page.addInitScript((id) => {
    sessionStorage.setItem(`fs.justUploaded.${id}`, String(Date.now()));
  }, ID);

  let calls = 0;
  await page.route(`**/v1/reports/${ID}`, (route) => {
    calls += 1;
    return route.fulfill(calls <= 2 ? envelope(null, 404) : envelope(complete));
  });

  await page.goto(`/reports/${ID}`);
  await expect(page.getByTestId('report-processing')).toBeVisible();
  await expect(page.getByTestId('report-error')).toHaveCount(0);
  await expect(page.getByTestId('fight-3-outcome')).toBeVisible({ timeout: 15_000 });
  expect(calls).toBe(3);
});

test('without the upload mark a 404 is a missing report straight away', async ({ page }) => {
  let calls = 0;
  await page.route(`**/v1/reports/${ID}`, (route) => {
    calls += 1;
    return route.fulfill(envelope(null, 404));
  });
  await page.goto(`/reports/${ID}`);
  await expect(page.getByTestId('report-error')).toContainText('No report with that id');
  expect(calls).toBe(1);
});
