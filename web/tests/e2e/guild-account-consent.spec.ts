// web/tests/e2e/guild-account-consent.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [
    {
      id: 501,
      region: 'us',
      ruleset: 'hardcore',
      name: 'The Last Watch',
      rank: 'member',
      consent: 'gear',
      verified: true,
    },
  ],
};

test('changing consent calls the API and leaving removes the row', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.route('**/v1/guilds/501/members/me', (route) => {
    if (route.request().method() === 'PATCH') return route.fulfill(envelope({ consent: 'roster' }));
    return route.fulfill(envelope({ status: 'left' }));
  });

  await page.goto('/account');
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');
  await page.getByTestId('account-guild-consent').selectOption('roster');
  await expect(page.getByTestId('account-guild-consent')).toHaveValue('roster');

  await page.getByTestId('account-guild-leave').click();
  await expect(page.getByTestId('account-guilds')).toHaveCount(0);
});
