// web/tests/e2e/auth.spec.ts
import { expect, test } from './support/test';
import { meBnetFixture } from '../../src/fixtures/me-bnet';

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

test('a signed-out visitor is offered Battle.net as the one sign-in route', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/login');

  await expect(page.getByRole('button', { name: 'Sign in or choose a character' })).toBeVisible();
  await expect(page.getByTestId('battlenet')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Fsigned_in%3D1$/,
  );
  await expect(page.getByTestId('email-form')).toHaveCount(0);
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

  await expect(page.getByTestId('character-selector')).toBeVisible();
  await expect(page.getByText('Raid PC')).toBeVisible();
  await expect(
    page
      .getByTestId('account-character-row')
      .filter({ hasText: 'Elyra Duskvale' })
      .getByTestId('character-open'),
  ).toHaveAttribute('href', '/character/us/hardcore/elyra-duskvale');

  // One device is paired, so the entry is the quiet "Pair another computer", and the code
  // comes with the line that says it is only the fallback.
  await expect(page.getByRole('button', { name: 'Get a pairing code' })).toHaveCount(0);
  await page.getByRole('button', { name: 'Pair another computer' }).click();
  await expect(page.getByTestId('pairing-code')).toHaveText('4821-9930');
  await expect(page.getByTestId('pairing-fallback-line')).toContainText('cannot open a browser');

  // Revoke asks first, and Cancel leaves the device alone.
  await page.getByTestId('revoke').click();
  await expect(page.getByTestId('revoke-question')).toHaveText('Revoke Raid PC? The companion signs out.');
  await page.getByTestId('revoke-cancel').click();
  await expect(page.getByTestId('revoke')).toBeVisible();
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

test('a signed-in visitor sees imported characters, refreshes, and is told what the refresh did', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: true, data: meBnetFixture, error: null, request_id: 'r' })),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );

  await page.goto('/account?refreshed=1');

  // No snapshot from before the redirect (a fresh test page): the return reads as updated.
  // The result line takes the Characters head's meta for a while.
  await expect(page.getByTestId('account-characters')).toContainText('Updated just now');
  await expect.poll(() => new URL(page.url()).search).toBe('');

  // The guilded character carries its guild on the row's second line.
  const rows = page.getByTestId('account-character-row');
  await expect(rows.filter({ hasText: 'Thoradin' }).getByTestId('character-guild')).toHaveText(
    'Iron Vanguard',
  );

  // The unguilded character carries no guild line at all.
  await expect(rows.filter({ hasText: 'Elyra Duskvale' }).getByTestId('character-guild')).toHaveCount(0);

  // The result line gives way, after a few seconds, to the import time and the refresh link.
  await expect(page.getByTestId('account-characters')).toContainText(/Updated .* from Battle\.net/, {
    timeout: 12_000,
  });
  await expect(page.getByTestId('bnet-refresh')).toHaveAttribute(
    'href',
    /\/v1\/auth\/battlenet\/start\?next=%2Faccount%3Frefreshed%3D1$/,
  );
});

test('choosing a main on the account page records it and makes it the head answer', async ({ page }) => {
  let mainKey: string | undefined;
  const me = () => ({
    ok: true,
    data: {
      user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
      characters: [
        { key: 'us/pvp/reloadd', region: 'us', ruleset: 'pvp', name: 'Reloadd', class: 'hunter', level: 25 },
        { key: 'us/pvp/dottzz', region: 'us', ruleset: 'pvp', name: 'Dottzz', class: 'priest', level: 12 },
      ],
      guilds: [],
      ...(mainKey === undefined ? {} : { main_character_key: mainKey }),
    },
    error: null,
    request_id: 'r',
  });
  await page.route('**/v1/me', async (route) => {
    if (route.request().method() === 'PATCH') {
      mainKey = (route.request().postDataJSON() as { main_character_key: string }).main_character_key;
    }
    await route.fulfill({ json: me() });
  });
  await page.route('**/v1/devices', (route) => route.fulfill({ json: { ok: true, data: [], error: null } }));
  await page.route('**/v1/reports**', (route) =>
    route.fulfill({ json: { ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null } }),
  );
  await page.route('**/v1/characters/**', (route) =>
    route.fulfill({ status: 404, json: { ok: false, data: null, error: { message: 'none' } } }),
  );
  await page.goto('/account');
  const list = page.getByTestId('account-characters');
  await expect(list).toBeVisible();
  // Reloadd is the site's guess (highest level) until the player chooses; Dottzz offers Set as main.
  await expect(list.getByTestId('character-main-pill')).toHaveCount(1);
  const dottzzRow = list.getByTestId('account-character-row').filter({ hasText: 'Dottzz' });
  await dottzzRow.getByTestId('character-set-main').click();
  await expect(dottzzRow.getByTestId('character-main-pill')).toBeVisible();
  expect(mainKey).toBe('us/pvp/dottzz');
});
