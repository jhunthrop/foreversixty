// web/tests/e2e/logs-upload.spec.ts
import { expect, test } from '@playwright/test';

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({ ok: status < 400, data: body, error: null, request_id: 'r' }),
});

test('a log uploads part by part and hands off to the report', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        user: { id: 1, battletag: 'F#1', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
      }),
    ),
  );
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
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

  await page.goto('/logs');
  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('9/26 20:10:00.000  COMBAT_LOG_VERSION,16\n'),
  });
  await page.getByLabel('Title').fill('Tuesday');
  await page.getByRole('radio', { name: 'Unlisted' }).check();
  await page.getByTestId('upload-start').click();

  await page.waitForURL('**/reports/fixture2abcd');
  expect(completeBody).toEqual({
    etags: [{ number: 1, etag: '"e1"' }],
    title: 'Tuesday',
    visibility: 'unlisted',
  });
});

test('a failed part shows what went wrong instead of a spinner', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
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

  await page.goto('/logs');
  await page.getByTestId('upload-file').setInputFiles({
    name: 'WoWCombatLog.txt',
    mimeType: 'text/plain',
    buffer: Buffer.from('x'),
  });
  await page.getByTestId('upload-start').click();

  await expect(page.getByTestId('upload-error')).toContainText('part 1', { timeout: 20_000 });
});

// The signal was only read between parts and nothing ever called xhr.abort(), so there was
// no way to stop an upload at all: no control on the page, and no effect if there had been.
test('cancel stops the part in flight and leaves the form usable', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(null, 401)));
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

  await page.goto('/logs');
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

test('the pairing code and the companion downloads are on the page', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        user: { id: 1, battletag: 'F#1', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
      }),
    ),
  );
  await page.route('**/v1/reports?mine=1**', (route) =>
    route.fulfill(fulfil({ rows: [], total: 0, page: 1, per_page: 100 })),
  );
  await page.route('**/v1/devices/pair', (route) =>
    route.fulfill(fulfil({ code: '4821-9930', expires_in: 600 })),
  );

  await page.goto('/logs');
  await expect(page.getByTestId('companion-downloads').getByRole('link')).toHaveCount(4);
  await page.getByRole('button', { name: 'Show pairing code' }).click();
  await expect(page.getByTestId('pairing-code')).toHaveText('4821-9930');
});
