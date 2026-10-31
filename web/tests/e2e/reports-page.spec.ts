// web/tests/e2e/reports-page.spec.ts
// /reports is "Your reports", the Logs menu's second row: signed out a sign-in prompt, signed in
// the reports newest first (or "No logs yet"), a failed read an error that retries, and the
// sample report one click away. /reports/<id> stays the report page.
import { expect, test, type Page } from '@playwright/test';
import { SAMPLE_REPORT_ID } from './support/logs-sample-fixture';

const SAMPLE_REPORT_PATH = `/reports/${SAMPLE_REPORT_ID}`;
const FIXTURE_REPORT_ID = 'fixture2abcd';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? body : null,
    error: status < 400 ? null : { message: 'oops' },
    request_id: 'r',
  }),
});

const SIGNED_IN = {
  user: { id: 1, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
};

const HOUR_MS = 60 * 60 * 1000;

function report(id: string, title: string, hoursAgo: number) {
  return {
    id,
    title,
    zone: 'Sanguine Depths',
    status: 'ready',
    visibility: 'public',
    created_at: new Date(Date.now() - hoursAgo * HOUR_MS).toISOString(),
    fight_count: 12,
    kill_count: 2,
    boss_count: 2,
    logging_character: null,
  };
}

async function signedIn(page: Page, reports: unknown[] | 'fail'): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(SIGNED_IN)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      reports === 'fail'
        ? fulfil(null, 500)
        : fulfil({ rows: reports, total: reports.length, page: 1, per_page: 100 }),
    ),
  );
}

test('the Logs menu’s "Your reports" row leads to a page of that name', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/reports');
  await expect(page.getByRole('heading', { level: 1, name: 'Your reports' })).toBeVisible();
  await expect(page.getByTestId('report-shell-title')).toHaveCount(0);
  await expect(page.getByText('No report with that id')).toHaveCount(0);
});

test('signed out, it asks to sign in and offers the sample report', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await page.goto('/reports');
  const signIn = page.getByTestId('your-reports-signin');
  await expect(signIn).toBeVisible();
  await expect(signIn.getByRole('link', { name: 'Sign in with Battle.net' })).toHaveAttribute(
    'href',
    /next=%2Freports/,
  );
  await expect(page.getByTestId('your-reports-sample')).toHaveAttribute('href', SAMPLE_REPORT_PATH);
  await expect(page.getByTestId('my-reports')).toHaveCount(0);
});

test('signed in, it lists the reports newest first, each opening its report', async ({ page }) => {
  await signedIn(page, [
    report('newreport2ab', 'Tuesday raid', 2),
    report('oldreport2ab', 'Monday raid', 30),
  ]);
  await page.goto('/reports');
  const rows = page.getByTestId('your-reports').getByTestId('report-row');
  await expect(rows).toHaveCount(2);
  await expect(rows.nth(0).getByRole('link', { name: 'Tuesday raid' })).toHaveAttribute(
    'href',
    '/reports/newreport2ab',
  );
  await expect(rows.nth(1)).toContainText('Monday raid');
  await expect(page.getByTestId('your-reports-signin')).toHaveCount(0);
});

test('signed in with no reports, it says No logs yet, with the upload and the sample', async ({ page }) => {
  await signedIn(page, []);
  await page.goto('/reports');
  const empty = page.getByTestId('my-reports-empty');
  await expect(empty).toContainText('No logs yet.');
  await expect(empty.getByRole('link', { name: 'Upload a log' })).toHaveAttribute('href', '/logs');
  await expect(page.getByTestId('my-reports-empty-secondary')).toHaveAttribute('href', SAMPLE_REPORT_PATH);
});

test('a list that fails to load says so, and Try again reads it again', async ({ page }) => {
  let failing = true;
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(SIGNED_IN)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(
      failing
        ? fulfil(null, 500)
        : fulfil({ rows: [report('newreport2ab', 'Tuesday raid', 2)], total: 1, page: 1, per_page: 100 }),
    ),
  );
  await page.goto('/reports');
  await expect(page.getByTestId('my-reports-error')).toContainText('Your reports did not load.');
  failing = false;
  await page.getByTestId('my-reports-error').getByRole('button', { name: 'Try again' }).click();
  await expect(page.getByTestId('your-reports').getByTestId('report-row')).toHaveCount(1);
});

test('a session that fails to load says so, and Try again asks again', async ({ page }) => {
  let failing = true;
  await page.route('**/v1/me', (route) => route.fulfill(failing ? fulfil(null, 500) : fulfil(SIGNED_IN)));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.goto('/reports');
  const error = page.getByTestId('your-reports-error');
  await expect(error).toContainText('Your account did not load.');
  failing = false;
  await error.getByRole('button', { name: 'Try again' }).click();
  await expect(page.getByTestId('my-reports-empty')).toBeVisible();
});

test('/reports/<id> still opens the report', async ({ page }) => {
  await page.goto(`/reports/${FIXTURE_REPORT_ID}`);
  await expect(page.getByTestId('report-title')).toHaveText('Sanguine Depths, fixture night');
  await expect(page.getByTestId('fight-selector')).toBeVisible();
  await expect(page.getByTestId('your-reports')).toHaveCount(0);
});
