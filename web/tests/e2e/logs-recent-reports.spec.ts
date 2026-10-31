// web/tests/e2e/logs-recent-reports.spec.ts
// Logs spec 2026-10-11 sections 2.5 and 9: the "Recent public reports" panel under "Your reports"
// on /logs, an honest empty state, and the one sample report as a plain labelled row, shown once.
import { apiUnreachable, expect, test } from './support/test';
import { routeSampleReportMeta, SAMPLE_REPORT_ID, SAMPLE_REPORT_META } from './support/logs-sample-fixture';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/devices'));

const THREE_DAYS_MS = 3 * 24 * 60 * 60 * 1000;

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

/** /logs also mounts "Your reports" (MyReports via Account), which needs a session read. */
async function signedOut(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
}

test('the sample report is one plain labelled row, with a date and its kills, and no Sample pill', async ({
  page,
}) => {
  await signedOut(page);
  await routeSampleReportMeta(page);
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));

  await page.goto('/logs');
  const sample = page.getByTestId('recent-reports-sample');
  await sample.scrollIntoViewIfNeeded();
  await expect(sample).toContainText('A sample report');
  const row = sample.getByTestId('report-row');
  await expect(row.getByRole('link', { name: SAMPLE_REPORT_META.title })).toHaveAttribute(
    'href',
    `/reports/${SAMPLE_REPORT_ID}`,
  );
  // The date the report page says (its first fight), not a second format (W.9); the zone and
  // count come from the sample's meta, the same source the page reads (L.7).
  await expect(row.getByTestId('report-row-when')).toHaveText('17 Sept 2026');
  await expect(row.getByTestId('report-row-summary')).toHaveText(
    `${SAMPLE_REPORT_META.zone} · ${SAMPLE_REPORT_META.fights.length} fights`,
  );
  await expect(page.locator('.pill-sample')).toHaveCount(0);
});

test('the sample appears once: a feed that carries it is filtered by id', async ({ page }) => {
  await signedOut(page);
  await routeSampleReportMeta(page);
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: SAMPLE_REPORT_ID,
            title: SAMPLE_REPORT_META.title,
            created_at: SAMPLE_REPORT_META.created_at,
            fight_count: 2,
            kill_count: 1,
          },
        ],
      }),
    ),
  );

  await page.goto('/logs');
  const panel = page.getByTestId('recent-reports');
  await panel.scrollIntoViewIfNeeded();
  await expect(page.getByTestId('recent-reports-empty')).toBeVisible();
  await expect(panel.getByRole('link', { name: SAMPLE_REPORT_META.title })).toHaveCount(1);
});

test('lists a public report by title, relative date, kills and guild', async ({ page }) => {
  await signedOut(page);
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'abc123def456',
            title: 'Progress night',
            created_at: new Date(Date.now() - THREE_DAYS_MS - 60_000).toISOString(),
            fight_count: 8,
            kill_count: 3,
            guild_name: 'The Last Watch',
          },
        ],
      }),
    ),
  );

  await page.goto('/logs');
  const panel = page.getByTestId('recent-reports');
  await panel.scrollIntoViewIfNeeded();
  const link = panel.getByRole('link', { name: 'Progress night' });
  await expect(link).toHaveAttribute('href', '/reports/abc123def456');
  await expect(panel.getByTestId('report-row-when').first()).toHaveText('3 days ago');
  await expect(panel.getByTestId('report-row-summary').first()).toHaveText(
    '3 bosses · 8 fights · The Last Watch',
  );
});

test('an honest empty state before the first public report lands', async ({ page }) => {
  await signedOut(page);
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));

  await page.goto('/logs');
  await page.getByTestId('recent-reports').scrollIntoViewIfNeeded();
  await expect(page.getByTestId('recent-reports-empty')).toHaveText(
    'No public reports yet. The first raid logs land in December; dungeon logs are welcome now.',
  );
});

test('an "Older reports" control appears only when the API says a further page exists', async ({ page }) => {
  await signedOut(page);
  let calls = 0;
  await page.route('**/v1/reports/recent**', (route) => {
    calls += 1;
    const cursor = new URL(route.request().url()).searchParams.get('cursor');
    if (cursor === null) {
      return route.fulfill(
        fulfil({
          rows: [
            {
              id: 'page1report1',
              title: 'Night one',
              created_at: '2026-12-09T22:10:00Z',
              fight_count: 4,
              kill_count: 1,
            },
          ],
          next_cursor: 'opaque-cursor',
        }),
      );
    }
    return route.fulfill(
      fulfil({
        rows: [
          {
            id: 'page2report1',
            title: 'Night two',
            created_at: '2026-12-08T22:10:00Z',
            fight_count: 6,
            kill_count: 2,
          },
        ],
      }),
    );
  });

  await page.goto('/logs');
  const panel = page.getByTestId('recent-reports');
  await panel.scrollIntoViewIfNeeded();
  await expect(panel.getByRole('link', { name: 'Night one' })).toBeVisible();
  const older = page.getByTestId('recent-reports-older');
  await expect(older).toBeVisible();

  await older.click();
  await expect(panel.getByRole('link', { name: 'Night two' })).toBeVisible();
  await expect(page.getByTestId('recent-reports-older')).toHaveCount(0);
  expect(calls).toBe(2);
});

test('"Your reports" is the first panel for a signed-in visitor, and the page has no spine bar', async ({
  page,
}) => {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ user: { battletag: 'Fixture#1234' }, characters: [], guilds: [] })),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
  await routeSampleReportMeta(page);
  await page.goto('/logs');
  await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  await page.getByTestId('my-reports').waitFor();
  const myReportsBox = await page.getByTestId('my-reports').boundingBox();
  const recentBox = await page.getByTestId('recent-reports').boundingBox();
  expect(myReportsBox).not.toBeNull();
  expect(recentBox).not.toBeNull();
  expect((myReportsBox as { y: number }).y).toBeLessThan((recentBox as { y: number }).y);
});
