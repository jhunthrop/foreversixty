// web/tests/e2e/report-live.spec.ts
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { expect, test } from '@playwright/test';

const FIXTURES = path.resolve(process.cwd(), 'src/fixtures/report');
const DATA = '/logs-data/reports/fixture2live';

const envelope = (data: unknown) => ({
  status: 200,
  contentType: 'application/json',
  body: JSON.stringify({ ok: true, data, error: null, request_id: 'r' }),
});

test('a live fight polls, then settles when it closes', async ({ page }) => {
  const report = JSON.parse(await readFile(path.join(FIXTURES, 'report.json'), 'utf8'));
  const summary = JSON.parse(await readFile(path.join(FIXTURES, 'fights/3/summary.json'), 'utf8'));
  const meta = JSON.parse(await readFile(path.join(FIXTURES, 'meta.json'), 'utf8'));

  const open = {
    ...report,
    report_id: 'fixture2live',
    fights: [{ ...report.fights[2], in_progress: true, kill: false, duration_ms: 12_000 }],
  };
  const closed = { ...open, fights: [{ ...report.fights[2], in_progress: false }] };

  let reportCalls = 0;
  await page.route(`**${DATA}/report.json`, (route) => {
    reportCalls += 1;
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify(reportCalls < 3 ? open : closed),
    });
  });
  await page.route(`**${DATA}/fights/3/live.json`, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...summary, duration_ms: 12_000 }),
    }),
  );
  await page.route(`**${DATA}/fights/3/summary.json`, (route) =>
    route.fulfill({ status: 200, contentType: 'application/json', body: JSON.stringify(summary) }),
  );
  await page.route('**/v1/reports/fixture2live', (route) =>
    route.fulfill(
      envelope({
        ...meta,
        id: 'fixture2live',
        status: 'live',
        data_base_url: DATA,
        fights: open.fights,
      }),
    ),
  );

  await page.goto('/reports/fixture2live');

  await expect(page.getByTestId('report-live')).toBeVisible();
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Live');

  // Two polls at five seconds each, plus slack.
  await expect(page.getByTestId('fight-3-outcome')).toHaveText('Kill', { timeout: 20_000 });
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  expect(reportCalls).toBeGreaterThanOrEqual(3);
});

test('a closed report never polls', async ({ page }) => {
  let calls = 0;
  await page.route('**/logs-data/reports/fixture2abcd/report.json', (route) => {
    calls += 1;
    return route.continue();
  });
  await page.goto('/reports/fixture2abcd');
  await expect(page.getByTestId('report-title')).toBeVisible();
  await page.waitForTimeout(12_000);
  expect(calls).toBe(1);
});

test('the timelines and events views render from the summary', async ({ page }) => {
  await page.goto('/reports/fixture2abcd?fight=3&view=timelines');
  await expect(page.getByTestId('timelines')).toContainText('Baelgrim');
  await expect(page.getByTestId('lane-Player-4184-000000A4')).toBeVisible();

  await page.goto('/reports/fixture2abcd?fight=3&view=events');
  await expect(page.getByTestId('event-list')).toContainText('Thalgrit died');
  await page.getByTestId('events-search').fill('frostbolt');
  await expect(page.getByTestId('event-list')).toContainText('Frostbolt');
  await expect(page.getByTestId('event-list')).not.toContainText('died');
});
