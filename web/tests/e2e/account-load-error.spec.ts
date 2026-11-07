import { addSessionHint } from './support/account-reserve';
import { gotoHydrated } from './support/hydrated';
import { expect, test } from './support/test';

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

test('a failed /v1/me shows Try again, which re-fires the same request', async ({ page, baseURL }) => {
  // The error Plate is for a signed-in visitor: without the session hint a failed read leaves the
  // sign-in page (the test below).
  await addSessionHint(page, baseURL);
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

  await expect(page.getByTestId('account-character-row').filter({ hasText: 'Elyra Duskvale' })).toBeVisible();
  await expect(page.getByTestId('account-load-error')).toHaveCount(0);
});

test('without a session hint a failed /v1/me leaves the sign-in page, not the error', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: { message: 'oops' }, request_id: 'r' }, 500)),
  );
  const answered = page.waitForResponse('**/v1/me');
  await gotoHydrated(page, '/account', 'account');
  await answered;
  // The island has hydrated and its read has failed: give it the frame it would paint the error in.
  await page.evaluate(() => new Promise((done) => requestAnimationFrame(() => requestAnimationFrame(done))));
  await expect(page.getByTestId('account-load-error')).toHaveCount(0);
  await expect(page.getByTestId('battlenet')).toBeVisible();
});
