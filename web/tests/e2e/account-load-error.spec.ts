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
  // Every /v1/me answers 500 until the retry is clicked. Counting requests is not
  // deterministic: the header's SessionNav and the account island share one deduped
  // fetchMeOnce promise when their calls overlap (one request) and make two when they do
  // not, so "fail the first N" sometimes failed the retry itself under CPU contention.
  let failing = true;
  await page.route('**/v1/me', (route) => {
    if (failing)
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
  failing = false;
  await page.getByTestId('account-load-error-retry').click();

  await expect(
    page.getByTestId('account-characters').getByRole('link', { name: 'Elyra Duskvale' }),
  ).toBeVisible();
  await expect(page.getByTestId('account-load-error')).toHaveCount(0);
});
