// web/tests/e2e/logs-upload.spec.ts
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { routeSampleReportMeta } from './support/logs-sample-fixture';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

// /logs now also mounts the public "Recent public reports" panel (RecentReports.svelte),
// which reads unauthenticated regardless of who is signed in. Every test in this file
// exercises other parts of the page, so it is stubbed once here rather than in each test,
// the way logs-recent-reports.spec.ts exercises the panel itself in full.
test.beforeEach(async ({ page }) => {
  await page.route('**/v1/reports/recent**', (route) => route.fulfill(fulfil({ rows: [] })));
});

const SIGNED_IN = {
  user: { id: 1, battletag: 'F#1', email: null, role: 'user', anonymize: false },
  characters: [],
  guilds: [],
};

/** POST /v1/uploads needs a session, and the form says so up front, so every test that
 *  drives the form is signed in. */
async function signIn(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(SIGNED_IN)));
  // `listDevices()` (lib/account/api.ts) expects the envelope's `data` to be a bare
  // `Device[]`, matching the real API -- not `{ devices: [] }`. Harmless before the pairing
  // poll (logs landing spec §4.D.2) started reading `devices` itself in `onPair`, which is
  // what surfaced the mismatch: a non-array `devices.map` threw and silently swallowed the
  // "Show pairing code" click.
  await page.route('**/v1/devices', (route) => route.fulfill(fulfil([])));
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
}

/** The form is part of the page's one island. A file chosen before it hydrates fires its change
 *  event at nothing, so wait for the island to drop its `ssr` mark. */
async function openLogs(page: Page): Promise<void> {
  await routeSampleReportMeta(page);
  await page.goto('/logs');
  await expect(page.locator('astro-island[ssr]:has([data-testid="logs-grid"])')).toHaveCount(0);
}

test('a log uploads part by part and hands off to the report', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/uploads', (route) =>
    route.fulfill(
      fulfil({
        upload_id: 'up1',
        parts: [{ number: 1, url: 'https://r2.test/p1' }],
        complete_url: 'https://r2.test/c',
      }),
    ),
  );
  await page.route('https://r2.test/p1', (route) =>
    route.fulfill({ status: 200, headers: { etag: '"e1"', 'access-control-expose-headers': 'ETag' } }),
  );
  let completeBody: unknown;
  await page.route('**/v1/uploads/up1/complete', (route) => {
    completeBody = route.request().postDataJSON();
    return route.fulfill(fulfil({ report_id: 'fixture2abcd' }, 202));
  });

  await openLogs(page);
  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n'),
  });
  await page.getByLabel('Title').fill('Tuesday');
  await page.getByText('Unlisted', { exact: true }).click();
  await expect(page.getByRole('radio', { name: 'Unlisted' })).toBeChecked();
  await expect(page.getByTestId('upload-visibility-note')).toContainText('only people with the link');
  await page.getByTestId('upload-start').click();

  await page.waitForURL('**/reports/fixture2abcd');
  expect(completeBody).toEqual({
    etags: [{ number: 1, etag: '"e1"' }],
    title: 'Tuesday',
    visibility: 'unlisted',
  });
});

test('a failed part shows what went wrong instead of a spinner', async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/uploads', (route) =>
    route.fulfill(
      fulfil({
        upload_id: 'up1',
        parts: [{ number: 1, url: 'https://r2.test/p1' }],
        complete_url: 'https://r2.test/c',
      }),
    ),
  );
  await page.route('https://r2.test/p1', (route) => route.fulfill({ status: 500 }));

  await openLogs(page);
  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('x'),
  });
  await page.getByTestId('upload-start').click();

  // Fix round 1 (2026-10-04, build review): the player-facing line is generic now -- no
  // part number, no storage vendor, no doubled "did not finish" wrapper. The technical
  // detail still reaches devtools (Upload.svelte's own console.error), not this assertion.
  await expect(page.getByTestId('upload-error')).toHaveText(
    'The upload did not finish. Your file is untouched; try again.',
    { timeout: 20_000 },
  );
  await expect(page.getByTestId('upload-error')).not.toContainText('R2');
  await expect(page.getByTestId('upload-error')).not.toContainText('part 1');
});

// The signal was only read between parts and nothing ever called xhr.abort(), so there was
// no way to stop an upload at all: no control on the page, and no effect if there had been.
test('cancel stops the part in flight and leaves the form usable', async ({ page }) => {
  await signIn(page);
  let completed = 0;
  await page.route('**/v1/uploads', (route) =>
    route.fulfill(
      fulfil({
        upload_id: 'up1',
        parts: [{ number: 1, url: 'https://r2.test/p1' }],
        complete_url: 'https://r2.test/c',
      }),
    ),
  );
  await page.route('**/v1/uploads/up1/complete', (route) => {
    completed += 1;
    return route.fulfill(fulfil({ report_id: 'fixture2abcd' }, 202));
  });
  // Never answered: the part is in flight for as long as the test wants it to be, which is
  // the state Cancel has to be able to interrupt.
  let releasePart = (): void => {};
  const held = new Promise<void>((resolve) => (releasePart = resolve));
  await page.route('https://r2.test/p1', async (route) => {
    await held;
    await route.fulfill({ status: 200, headers: { etag: '"e1"' } });
  });

  await openLogs(page);
  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('x'),
  });
  await page.getByTestId('upload-start').click();
  await expect(page.getByTestId('upload-progress')).toBeVisible();

  await page.getByTestId('upload-cancel').click();

  // Back to where it started, with no error: the visitor asked for this and knows.
  await expect(page.getByTestId('upload-progress')).toHaveCount(0);
  await expect(page.getByTestId('upload-cancel')).toHaveCount(0);
  await expect(page.getByTestId('upload-error')).toHaveCount(0);
  await expect(page.getByTestId('upload-start')).toBeEnabled();

  // And the upload is really over: releasing the part does not complete it behind the
  // visitor's back or navigate away from the page.
  releasePart();
  await page.waitForTimeout(500);
  expect(completed).toBe(0);
  await expect(page).toHaveURL(/\/logs$/);
});

test("the pairing code is on the page, and the companion's downloads point to /setup", async ({ page }) => {
  await signIn(page);
  await page.route('**/v1/devices/pair', (route) =>
    route.fulfill(fulfil({ code: '4821-9930', expires_in: 600 })),
  );

  await openLogs(page);
  await expect(page.getByTestId('logs-companion-link')).toHaveAttribute('href', '/setup#companion');
  await expect(page.getByTestId('companion-downloads')).toHaveCount(0);
  await page.getByRole('button', { name: 'Show pairing code' }).click();
  await expect(page.getByTestId('pairing-code')).toHaveText('4821-9930');
});

test('before a file is chosen the panel is the explanation and the zone, with no Upload button', async ({
  page,
}) => {
  await signIn(page);
  await openLogs(page);
  const panel = page.getByTestId('upload');
  await expect(panel.getByTestId('upload-zone')).toContainText('WoWCombatLog.txt');
  await expect(panel.getByText('choose a file', { exact: true })).toBeVisible();
  await expect(page.getByTestId('upload-start')).toHaveCount(0);
  await expect(page.getByLabel('Title')).toHaveCount(0);

  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('x'),
  });
  await expect(page.getByTestId('upload-start')).toBeVisible();
  await expect(page.getByLabel('Title')).toBeVisible();
  await expect(page.getByText('Who can open it')).toBeVisible();
});

test('a signed-out visitor is told in plain words, with no sign-in control anywhere on the page', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
  await openLogs(page);

  await expect(page.getByTestId('reports-signed-out')).toHaveText('Sign in to see your reports.');
  await expect(page.getByTestId('upload-signed-out')).toHaveText(
    'Uploading needs a signed-in account; the report is filed under it.',
  );
  // Signed out the companion card is a sentence and a button, and the pairing block is not drawn.
  await expect(page.getByTestId('pairing')).toHaveCount(0);
  // Signed out, the form is not drawn at all.
  await expect(page.getByTestId('upload-file')).toHaveCount(0);
  await expect(page.getByTestId('upload-start')).toHaveCount(0);
  // One sign-in button on the page, in Your reports (W.9); the other panels only say what they need.
  const main = page.locator('#main');
  await expect(main.getByRole('link', { name: /sign in/i })).toHaveCount(1);
  await expect(main.getByRole('button', { name: /sign in/i })).toHaveCount(0);
});
