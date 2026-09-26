// web/tests/e2e/logs-recent-reports.spec.ts
// Spec section 4: a "Recent public reports" panel above "Your reports" on /logs, an honest
// empty state, and the one-line framing for a visitor who is not raiding yet.
import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

/** /logs also mounts "Your reports" (MyReports via Account), which needs a session read. */
async function signedOut(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
}

test('the framing line tells a non-raider what counts as a log', async ({ page }) => {
  await signedOut(page);
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));

  await page.goto('/logs');

  await expect(page.getByTestId('logs-framing')).toHaveText(
    'Logs are for group content at any level: a dungeon run logs the same way a raid does.',
  );
});

test('lists a public report by title, date, counts and guild', async ({ page }) => {
  await signedOut(page);
  await page.route('**/v1/reports/recent**', (route) =>
    route.fulfill(
      fulfil({
        rows: [
          {
            id: 'abc123def456',
            title: 'Progress night',
            created_at: '2026-12-09T22:10:00Z',
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
  // The panel hydrates only once it is on screen (client:visible), the same reason
  // logs-upload.spec.ts scrolls to reports-signin before looking for it.
  await panel.scrollIntoViewIfNeeded();
  const link = panel.getByRole('link', { name: 'Progress night' });
  await expect(link).toHaveAttribute('href', '/reports/abc123def456');
  await expect(panel).toContainText('2026-12-09');
  await expect(panel).toContainText('8 fights');
  await expect(panel).toContainText('3 kills');
  await expect(panel).toContainText('The Last Watch');
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

test('the spine bar mounts on /logs and "Your reports" is the first panel for a signed-in visitor', async ({
  page,
}) => {
  // The chip-slot pre-paint rule (global.css) hides the spine bar before hydration unless
  // `fs_csrf` is present, the same signed-in-visitor cookie sim-landing.spec.ts and
  // home-panel.spec.ts set; without it the bar's `.chip-slot` is `display: none` at first
  // paint and this test's own visibility assertion fails regardless of the /v1/me mock.
  await page.context().addCookies([{ name: 'fs_csrf', value: 'token', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        data: { user: { battletag: 'Fixture#1234' }, characters: [] },
        error: null,
      }),
    }),
  );
  await page.goto('/logs');
  await expect(page.getByTestId('current-character-bar')).toBeVisible();
  const myReports = page.getByTestId('my-reports');
  const companion = page.locator('#companion');
  const myReportsBox = await myReports.boundingBox();
  const companionBox = await companion.boundingBox();
  expect(myReportsBox).not.toBeNull();
  expect(companionBox).not.toBeNull();
  expect((myReportsBox as { y: number }).y).toBeLessThan((companionBox as { y: number }).y);
});
