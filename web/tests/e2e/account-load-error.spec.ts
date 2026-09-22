import { expect, test } from '@playwright/test';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

const ME_OK = {
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

test('a failed /v1/me shows Try again, which re-fires the same request', async ({ page }) => {
  let call = 0;
  await page.route('**/v1/me', (route) => {
    call += 1;
    // /account also mounts the header's SessionNav (Base.astro's `session` slot), which
    // shares `fetchMeOnce` with this island (src/lib/account/api.ts). SessionNav's bundle
    // is far smaller than Account.svelte's, so it hydrates and completes its own
    // `/v1/me` round trip first; the account island's own initial load is therefore the
    // *second* request, not the first. Failing both covers SessionNav's request and the
    // account island's own initial load, so the island actually reaches its failed state.
    if (call <= 2)
      return route.fulfill(
        fulfil({ ok: false, data: null, error: { message: 'oops' }, request_id: 'r' }, 500),
      );
    return route.fulfill(fulfil(ME_OK));
  });
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );

  await page.goto('/account');

  await expect(page.getByTestId('account-load-error')).toBeVisible();
  await page.getByTestId('account-load-error-retry').click();

  await expect(page.getByRole('link', { name: 'Elyra Duskvale' })).toBeVisible();
  await expect(page.getByTestId('account-load-error')).toHaveCount(0);
});
