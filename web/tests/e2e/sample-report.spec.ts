// web/tests/e2e/sample-report.spec.ts
// The sample report is bundled with the site (lib/report/sample.ts): /logs offers it, the
// click lands on fights, and no API is involved at any step.
import { expect, test } from '@playwright/test';
import { SAMPLE_REPORT_META } from './support/logs-sample-fixture';

const SHOT_DIR = process.env.SAMPLE_SHOT_DIR;

test('the sample report opens from /logs and shows fights, with no API', async ({ page }, info) => {
  const api: string[] = [];
  page.on('request', (request) => {
    if (request.url().includes('/v1/reports/sample')) api.push(request.url());
  });
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'no' }, request_id: 'r' }),
    }),
  );
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ok: true, data: { rows: [] }, error: null, request_id: 'r' }),
    }),
  );

  await page.goto('/logs');
  const sample = page.getByTestId('recent-reports-sample');
  await sample.scrollIntoViewIfNeeded();
  await expect(sample.getByTestId('report-row-summary')).toHaveText(
    `${SAMPLE_REPORT_META.zone} · ${SAMPLE_REPORT_META.fights.length} fights`,
  );
  await expect(sample.getByTestId('report-row-when')).toHaveText('17 Sept 2026');
  await sample.getByRole('link').first().click();

  await expect(page).toHaveURL(/\/reports\/sample$/);
  await expect(page.getByTestId('report-title')).toHaveText('Beta night, 17 Sept 2026');
  await expect(page.getByText('17 Sept 2026').first()).toBeVisible();
  await expect(page.getByTestId('report-sample-badge')).toHaveText('Sample report');
  await expect(page.getByTestId('report-live')).toHaveCount(0);
  await expect(page.getByTestId('report-owner')).toHaveCount(0);
  await expect(page.getByTestId('report-credit')).toHaveCount(0);
  await expect(page.getByTestId('report-error')).toHaveCount(0);
  await expect(page.getByText(SAMPLE_REPORT_META.zone).first()).toBeVisible();
  expect(api).toEqual([]);

  // L.8: a first look that convinces. Nothing the log cannot know is drawn; the class shows.
  const roster = page.getByTestId('summary-tab');
  await expect(roster).toBeVisible();
  await expect(roster.getByTestId(/^roster-Player-/).first()).toBeVisible();
  const body = (await page.locator('main').innerText()) ?? '';
  expect(body).not.toContain('REPLAY');
  expect(body).not.toContain('could not be reached');
  await expect(page.getByTestId('mode-replay')).toHaveCount(0);
  await expect(page.getByTestId('parse-unavailable')).toHaveCount(0);
  await expect(page.getByTestId('roster-percentile')).toHaveCount(0);
  await expect(page.getByTestId('roster-spec')).toHaveCount(0);
  await expect(page.getByTestId('report-sample-note')).toContainText(
    'A real beta night with the names changed. Level 10 to 30 players, so the numbers are small. Your own reports look like this.',
  );
  await expect(page).not.toHaveURL(/fight=/);
  // Opens on its first boss kill, else its longest boss attempt, else its longest pull. The
  // sample has no boss, so that is its longest pull (fight 6, 12:43), not the first.
  await expect(page.getByTestId('fight-6')).toHaveAttribute('aria-current', 'true');
  // Class colours: no name keeps the plain text colour.
  const colours = await roster
    .getByTestId(/^roster-Player-/)
    .evaluateAll((rows) =>
      rows.slice(0, 40).map((row) => getComputedStyle(row.querySelector('span') as Element).color),
    );
  expect(new Set(colours).size).toBeGreaterThan(2);
  // A trash row prints its length once.
  const row1 = await page.getByTestId('fight-1').innerText();
  expect((row1.match(/3:26/g) ?? []).length).toBe(1);

  if (SHOT_DIR !== undefined) {
    const width = info.project.name === 'mobile' ? 390 : 1280;
    await page.screenshot({ path: `${SHOT_DIR}/l8-sample-${width}.png`, fullPage: false });
  }
});
