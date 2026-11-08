// web/tests/e2e/app-beta.spec.ts
// foreversixty.gg/app, the page the apps' QR codes open: the tester signup (received, whether or
// not already listed; a refused address; a failure), the QR on a desktop and none on a phone, the phone's
// own platform chosen for it; and the owner's /admin/testers list with its copy buttons.
import type { Page, Request } from '@playwright/test';
import { appBetaCopy as copy } from '../../src/lib/app-beta/copy';
import { adminTestersCopy as adminCopy } from '../../src/lib/app-beta/admin';
import { expect, test } from './support/test';

const envelope = (status: number, data: unknown, fields: Record<string, string> | null = null) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code: 'invalid', message: 'refused', fields },
    request_id: 'r',
  }),
});

/** Answers the signup with one status and keeps every request it saw. */
async function answerSignup(
  page: Page,
  status: number,
  fields: Record<string, string> | null = null,
): Promise<Request[]> {
  const seen: Request[] = [];
  await page.route('**/v1/app-testers', (route) => {
    seen.push(route.request());
    return route.fulfill(envelope(status, { status: 'received' }, fields));
  });
  return seen;
}

async function fillAndSend(page: Page, email: string): Promise<void> {
  await page.getByTestId('app-beta-email').fill(email);
  await page.getByTestId('app-beta-submit').click();
}

test.describe('/app', () => {
  test('says the apps are coming, what they do, and offers the signup', async ({ page }) => {
    await page.goto('/app');
    await expect(page.getByRole('heading', { level: 1, name: copy.heading })).toBeVisible();
    await expect(page.getByTestId('app-beta-soon')).toHaveText(copy.soon);
    for (const line of copy.features) await expect(page.getByText(line)).toBeVisible();
    await expect(page.getByTestId('app-beta-consent')).toHaveText(copy.consent);
  });

  test('signs an address up for the chosen phone, with no cookie sent', async ({ page, isMobile }) => {
    test.skip(isMobile, 'the desktop default and the switch');
    const seen = await answerSignup(page, 201);
    await page.goto('/app');
    await page.getByTestId('app-beta-platform-android').click();
    await fillAndSend(page, '  Tester@Example.com ');
    await expect(page.getByTestId('app-beta-done')).toHaveText(copy.received('android'));
    await expect(page.getByTestId('app-beta-form')).toHaveCount(0);
    expect(seen).toHaveLength(1);
    expect(seen[0].postDataJSON()).toEqual({ email: 'Tester@Example.com', platform: 'android' });
    expect(await seen[0].headerValue('cookie')).toBeNull();
  });

  test('reads the same whether or not the address was already on the list', async ({ page }) => {
    await answerSignup(page, 201);
    await page.goto('/app');
    await page.getByTestId('app-beta-platform-ios').click();
    await fillAndSend(page, 'tester@example.com');
    await expect(page.getByTestId('app-beta-done')).toHaveText(copy.received('ios'));
    await expect(page.getByText(/already/)).toHaveCount(0);
  });

  test('stops a malformed address before sending, and shows the API refusal on the field', async ({
    page,
  }) => {
    const seen = await answerSignup(page, 422, { email: 'enter one email address' });
    await page.goto('/app');
    await fillAndSend(page, 'not-an-email');
    await expect(page.getByText(copy.invalidEmail)).toBeVisible();
    await expect(page.getByTestId('app-beta-email')).toHaveAttribute('aria-invalid', 'true');
    expect(seen).toHaveLength(0);
    await fillAndSend(page, 'tester@example.org');
    await expect(page.getByText(copy.invalidEmail)).toBeVisible();
    expect(seen).toHaveLength(1);
  });

  test('keeps the form and says to try again when the API fails', async ({ page }) => {
    await answerSignup(page, 500);
    await page.goto('/app');
    await fillAndSend(page, 'tester@example.com');
    await expect(page.getByTestId('app-beta-failed')).toHaveText(copy.failed);
    await expect(page.getByTestId('app-beta-submit')).toBeEnabled();
  });

  test('a desktop gets the QR for this page, iPhone chosen by default', async ({ page, isMobile }) => {
    test.skip(isMobile, 'desktop only');
    await page.goto('/app');
    await expect(page.getByRole('img', { name: copy.qrAlt })).toBeVisible();
    await expect(page.getByText(copy.qrLine)).toBeVisible();
    await expect(page.getByTestId('app-beta-platform-ios')).toHaveAttribute('aria-checked', 'true');
  });

  test('a phone gets no QR and its own platform chosen', async ({ page, isMobile }) => {
    test.skip(!isMobile, 'the phone project (Pixel 7)');
    await page.goto('/app');
    await expect(page.getByTestId('app-beta-qr')).toBeHidden();
    await expect(page.getByTestId('app-beta-platform-android')).toHaveAttribute('aria-checked', 'true');
  });
});

test.describe('/admin/testers', () => {
  const TESTERS = [
    { email: 'new@example.com', platform: 'ios', created_at: '2026-11-08T12:00:00Z' },
    { email: 'droid@example.com', platform: 'android', created_at: '2026-11-07T09:00:00Z' },
    { email: 'old@example.com', platform: 'ios', created_at: '2026-11-01T23:30:00Z' },
  ];

  const me = (role: string) => ({
    user: { id: 78, battletag: 'Owner#1894', email: null, role, anonymize: false },
    characters: [],
    guilds: [],
  });

  test('the owner sees every signup and copies one platform’s addresses', async ({ page }) => {
    await page.addInitScript(() => {
      const copied: string[] = [];
      Object.defineProperty(window, '__copied', { value: copied });
      Object.defineProperty(navigator, 'clipboard', {
        value: { writeText: async (text: string) => void copied.push(text) },
      });
    });
    await page.route('**/v1/me', (route) => route.fulfill(envelope(200, me('admin'))));
    await page.route('**/v1/admin/app-testers', (route) =>
      route.fulfill(envelope(200, { testers: TESTERS })),
    );
    await page.goto('/admin/testers');
    await expect(page.getByRole('heading', { level: 1, name: adminCopy.pageTitle })).toBeVisible();
    await expect(page.getByTestId('testers-grid-row')).toHaveCount(3);
    await expect(page.getByTestId('testers-grid-row').first()).toContainText('new@example.com');
    await expect(page.getByTestId('testers-grid-row').first()).toContainText('8 Nov 2026');
    await page.getByTestId('testers-copy-ios').click();
    await expect(page.getByTestId('testers-copy-ios')).toHaveText(adminCopy.copied);
    expect(await page.evaluate(() => (window as unknown as { __copied: string[] }).__copied)).toEqual([
      'new@example.com, old@example.com',
    ]);
    await expect(page.getByTestId('testers-copy-android')).toHaveText(adminCopy.copyEmails('android', 1));
  });

  test('anyone else sees one line and never the list', async ({ page }) => {
    let listed = false;
    await page.route('**/v1/me', (route) => route.fulfill(envelope(200, me('user'))));
    await page.route('**/v1/admin/app-testers', (route) => {
      listed = true;
      return route.fulfill(envelope(403, null));
    });
    await page.goto('/admin/testers');
    await expect(page.getByTestId('testers-not-owner')).toHaveText(adminCopy.notOwner);
    expect(listed).toBe(false);
  });
});
