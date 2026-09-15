// web/tests/e2e/auth.spec.ts
import { expect, test } from '@playwright/test';

const ME = {
  ok: true,
  data: {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
    characters: [
      {
        key: 'us/hardcore/elyra-duskvale',
        region: 'us',
        ruleset: 'hardcore',
        name: 'Elyra Duskvale',
        class: 'Priest',
      },
    ],
    guilds: [],
  },
  error: null,
  request_id: 'r',
};

const fulfil = (body: unknown, status = 200) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify(body),
});

test('a signed-out visitor is offered both sign-in routes', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/login');

  await expect(page.getByTestId('session-nav').getByRole('link', { name: 'Sign in' })).toBeVisible();
  await expect(page.getByTestId('battlenet')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Flogs$/,
  );
  await expect(page.getByTestId('email-form')).toBeVisible();
});

test('asking for an email link shows what happens next', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  let body: unknown;
  await page.route('**/v1/auth/email', (route) => {
    body = route.request().postDataJSON();
    return route.fulfill(fulfil({ ok: true, data: { sent: true }, error: null, request_id: 'r' }));
  });

  await page.goto('/login');
  await page.getByLabel('Or a sign-in link by email').fill('raider@example.com');
  await page.getByRole('button', { name: 'Send link' }).click();

  await expect(page.getByTestId('account-notice')).toContainText('works once, for twenty minutes');
  expect(body).toEqual({ email: 'raider@example.com' });
});

test('the account page lists devices, pairs one, and shows the character', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(ME)));
  await page.route('**/v1/devices', (route) =>
    route.request().method() === 'POST'
      ? route.fulfill(
          fulfil({ ok: true, data: { code: '4821-9930', expires_in: 600 }, error: null, request_id: 'r' }),
        )
      : route.fulfill(
          fulfil({
            ok: true,
            data: [
              {
                id: 'dev1',
                name: 'Raid PC',
                platform: 'windows',
                created_at: '2026-11-05T10:00:00Z',
                last_seen_at: null,
              },
            ],
            error: null,
            request_id: 'r',
          }),
        ),
  );
  await page.route('**/v1/devices/pair', (route) =>
    route.fulfill(
      fulfil({ ok: true, data: { code: '4821-9930', expires_in: 600 }, error: null, request_id: 'r' }),
    ),
  );

  await page.goto('/account');

  await expect(page.getByTestId('session-nav')).toContainText('Fixture#1234');
  await expect(page.getByText('Raid PC')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Elyra Duskvale' })).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale',
  );

  await page.getByRole('button', { name: 'Pair a device' }).click();
  await expect(page.getByTestId('pairing-code')).toHaveText('4821-9930');
});

test('every control on the account page clears 44px on phone', async ({ page }, testInfo) => {
  test.skip(testInfo.project.name !== 'mobile', 'phone hit targets');
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(ME)));
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.goto('/account');

  for (const control of await page.getByRole('button').all()) {
    const box = await control.boundingBox();
    expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
  }
});
