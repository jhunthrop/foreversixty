// web/tests/e2e/guild-account-consent.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const GUILD = { id: 501, name: 'The Last Watch', rank: 'member', verified: true };

function character(name: string) {
  return {
    key: `us/hardcore/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'hardcore',
    name,
    class: 'warrior',
    guild: GUILD,
  };
}

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [character('Aldric'), character('Brenna'), character('Corvin')],
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

test('changing consent says what it covers, and leaving asks first then drops the guild lines', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  let leaves = 0;
  await page.route('**/v1/guilds/501/members/me', (route) => {
    if (route.request().method() === 'PATCH') return route.fulfill(envelope({ consent: 'roster' }));
    leaves += 1;
    return route.fulfill(envelope({ status: 'left' }));
  });

  await page.goto('/account');
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');
  await expect(page.getByTestId('character-guild-line').first()).toBeVisible();

  await page.getByTestId('account-guild-consent').selectOption('roster');
  await expect(page.getByTestId('account-guild-consent')).toHaveValue('roster');
  await expect(page.getByTestId('account-guild-consent-saved')).toHaveText(
    'Saved. Applies to all 3 of your characters.',
  );

  // Leave asks first, naming what it removes; Cancel changes nothing.
  await page.getByTestId('account-guild-leave').click();
  await expect(page.getByTestId('account-guild-leave-confirm')).toContainText(
    'Leave The Last Watch? This removes all 3 of your characters from its roster.',
  );
  await page.getByTestId('account-guild-leave-confirm').getByRole('button', { name: 'Cancel' }).click();
  await expect(page.getByTestId('account-guild-leave-confirm')).toHaveCount(0);
  expect(leaves).toBe(0);
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');

  await page.getByTestId('account-guild-leave').click();
  await page.getByTestId('account-guild-leave-confirm-button').click();
  await expect(page.getByTestId('account-guilds')).toHaveCount(0);
  expect(leaves).toBe(1);
  // The character cards read the same source: their guild line goes with the guild.
  await expect(page.getByTestId('character-guild-line')).toHaveCount(0);
});
