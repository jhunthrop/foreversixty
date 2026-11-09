// web/tests/e2e/logs.spec.ts
// The logs landing (design/specs/2026-11-07-logs-reports.md section 4): signed in with a report,
// the head's answer is the latest report and Your reports is the Plate, the upload open below
// it; with no report the upload is the Plate and Your reports says No reports yet; the Upload
// button appears only after a file is chosen; a paired device's status line sits in the
// companion rail. Signed out, the upload Plate carries the page's one sign-in link and the sample
// report, and the shared character band is gone (decision L2). On a phone the reports still come
// before the upload.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { routeSampleReportMeta } from './support/logs-sample-fixture';
import { OBNOXIOUS, signInWith, signedOut } from './support/selector';

const DESKTOP = { width: 1440, height: 900 };
const PHONE = { width: 390, height: 844 };
const DAY_MS = 86_400_000;
const NEWEST_REPORT_ID = 'newestreport1';
const envelope = (data: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'e2e' }),
});

const daysAgo = (days: number): string => new Date(Date.now() - days * DAY_MS - 60_000).toISOString();

async function routeReports(page: Page, createdDaysAgo: number | null): Promise<void> {
  const rows =
    createdDaysAgo === null
      ? []
      : [
          {
            id: NEWEST_REPORT_ID,
            title: 'Onyxia, wipes',
            zone: "Onyxia's Lair",
            status: 'complete',
            visibility: 'public',
            created_at: daysAgo(createdDaysAgo),
            fight_count: 6,
            kill_count: 0,
          },
        ];
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(envelope({ rows, total: rows.length, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(envelope({ rows: [] })));
  await routeSampleReportMeta(page);
}

async function routeDevices(page: Page, paired: boolean): Promise<void> {
  const devices = paired
    ? [
        {
          id: 'dev1',
          name: 'MacBook Pro',
          platform: 'macOS',
          created_at: daysAgo(30),
          last_seen_at: new Date(Date.now() - 4 * 60_000).toISOString(),
        },
      ]
    : [];
  await page.route('**/v1/devices', (route) => route.fulfill(envelope(devices)));
}

async function openSignedIn(
  page: Page,
  options: { paired?: boolean; reportDaysAgo?: number | null } = {},
): Promise<void> {
  await signInWith(page, [OBNOXIOUS]);
  await routeDevices(page, options.paired ?? false);
  await routeReports(page, options.reportDaysAgo === undefined ? 1 : options.reportDaysAgo);
  await page.goto('/logs');
  await expect(page.getByTestId('logs-companion-pair')).toBeVisible();
}

const top = async (page: Page, selector: string): Promise<number> =>
  ((await page.locator(selector).first().boundingBox())?.y ?? Number.NaN) as number;

test.describe('signed in, desktop', () => {
  test.use({ viewport: DESKTOP });

  test('the head is the title alone (directions b1); Your reports opens the newest report', async ({
    page,
  }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('logs-answer')).toHaveCount(0);
    await expect(page.locator('#reports').getByTestId('report-row').first()).toHaveAttribute(
      'href',
      `/reports/${NEWEST_REPORT_ID}`,
    );
    await expect(page.locator('#reports').getByRole('link', { name: 'All ›' })).toHaveAttribute(
      'href',
      '/reports',
    );
    await expect(page.getByTestId('character-header')).toHaveCount(0);
    await expect(page.getByTestId('current-character-bar')).toHaveCount(0);
  });

  test('no device paired, the rail has no status line and offers a pairing code', async ({ page }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('companion-status')).toHaveCount(0);
    await expect(page.getByTestId('logs-companion').getByText('Pair it with your account')).toBeVisible();
  });

  test('with no logs the upload is the Plate, Your reports says so, and there is no answer', async ({
    page,
  }) => {
    await openSignedIn(page, { reportDaysAgo: null });
    await expect(page.getByTestId('my-reports-empty')).toHaveText(
      'No reports yet. Upload a log, or let the companion send your next raid.',
    );
    await expect(page.getByTestId('logs-answer')).toHaveCount(0);
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '#reports'));
    await expect(page.locator('#upload [data-region="plate"]')).toHaveCount(1);
  });

  test('report rows read zone, pulls and bosses with a relative age; finished rows carry no Tag', async ({
    page,
  }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('report-row-summary').first()).toHaveText(
      "Onyxia's Lair · 6 pulls · Trash only",
    );
    await expect(page.getByTestId('report-row-when').first()).toHaveText('1 day ago');
    await expect(page.getByTestId('report-status')).toHaveCount(0);
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
  });

  test('with reports, the upload still leads and Your reports is the panel under it', async ({ page }) => {
    await openSignedIn(page);
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '#reports'));
    await expect(page.locator('#reports [data-region="plate"]')).toHaveCount(1);
    await expect(page.locator('#upload [data-region="plate"]')).toHaveCount(1);
  });

  test('a paired device shows its status line in the companion rail', async ({ page }) => {
    await openSignedIn(page, { paired: true });
    await expect(page.getByTestId('companion-status-line')).toContainText('MacBook Pro');
    await expect(page.getByTestId('logs-companion').getByText('Pair another device')).toBeVisible();
  });

  test('the Upload button appears only after a file is chosen', async ({ page }) => {
    await openSignedIn(page);
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
    await page.getByTestId('upload-file').setInputFiles({
      name: 'WoWCombatLog.txt',
      mimeType: 'text/plain',
      buffer: Buffer.from('9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n'),
    });
    await expect(page.getByTestId('upload-start')).toBeVisible();
    await expect(page.getByTestId('upload-start')).toHaveText('Upload');
  });
});

test.describe('signed out, desktop', () => {
  test.use({ viewport: DESKTOP });

  test.beforeEach(async ({ page }) => {
    await signedOut(page);
    await routeReports(page, null);
    await page.goto('/logs');
  });

  test('the one sign-in link is in the upload Plate, beside the sample report', async ({ page }) => {
    await expect(page.getByTestId('character-header')).toHaveCount(0);
    const main = page.locator('#main');
    await expect(main.getByRole('link', { name: /sign in/i })).toHaveCount(1);
    await expect(page.locator('#upload').getByTestId('reports-sign-in')).toBeVisible();
    await expect(main.getByRole('button', { name: /sign in/i })).toHaveCount(0);
    await expect(page.getByTestId('upload-signed-out')).toBeVisible();
    await expect(page.getByTestId('upload-zone')).toHaveAttribute('aria-disabled', 'true');
    await expect(page.getByTestId('upload-sample')).toHaveAttribute('href', '/reports/sample');
    await expect(page.getByTestId('reports-signed-out')).toHaveText('Sign in to see your reports.');
    // The companion is a sentence and a button, with no pairing block.
    await expect(page.getByTestId('logs-companion-pair')).toHaveCount(0);
    await expect(page.getByTestId('logs-companion-link')).toBeVisible();
  });

  test('the upload comes first, the companion is in the rail, and the sample is one row', async ({
    page,
  }) => {
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '[data-testid="recent-reports"]'));
    const upload = await page.locator('#upload').boundingBox();
    const companion = await page.locator('#companion').boundingBox();
    expect(companion?.x ?? 0).toBeGreaterThan((upload?.x ?? 0) + (upload?.width ?? 0));
    await expect(page.getByTestId('recent-reports-sample').getByTestId('report-row')).toHaveCount(1);
  });
});

test.describe('signed in, phone', () => {
  test.use({ viewport: PHONE });

  test('the upload leads, then Your reports with the newest report', async ({ page }) => {
    await openSignedIn(page);
    await expect(page.locator('#reports').getByTestId('report-row').first()).toHaveAttribute(
      'href',
      `/reports/${NEWEST_REPORT_ID}`,
    );
    await expect(page.getByTestId('upload-start')).toHaveCount(0);
    expect(await top(page, '#upload')).toBeLessThan(await top(page, '#reports'));
  });
});
