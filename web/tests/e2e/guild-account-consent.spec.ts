// web/tests/e2e/guild-account-consent.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const GUILD = {
  id: 501,
  region: 'us',
  ruleset: 'hardcore',
  name: 'The Last Watch',
  rank: 'member',
  consent: 'gear',
  verified: true,
};

const CHARACTER = {
  key: 'us/hardcore/quillfeather',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Quillfeather',
  class: 'priest',
  guild: { id: 501, name: 'The Last Watch', rank: 'member', verified: true },
};

function meFor(consent: string, inGuild: boolean) {
  return {
    user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
    characters: [inGuild ? CHARACTER : { ...CHARACTER, guild: undefined }],
    guilds: inGuild ? [{ ...GUILD, consent }] : [],
  };
}

test('changing consent calls the API and leaving removes the row and the cards guild line', async ({
  page,
}) => {
  // The server's own state: the account page re-reads /v1/me after each mutation.
  let consent = 'gear';
  let inGuild = true;
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meFor(consent, inGuild))));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.route('**/v1/guilds/501/members/me', (route) => {
    if (route.request().method() === 'PATCH') {
      consent = 'roster';
      return route.fulfill(envelope({ consent: 'roster' }));
    }
    inGuild = false;
    return route.fulfill(envelope({ status: 'left' }));
  });

  await page.goto('/account');
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');
  await expect(page.getByTestId('character-guild-line')).toContainText('The Last Watch');
  await page.getByTestId('account-guild-consent').selectOption('roster');
  await expect(page.getByTestId('account-guild-consent')).toHaveValue('roster');

  await page.getByTestId('account-guild-leave').click();
  await expect(page.getByTestId('account-guilds')).toHaveCount(0);
  // The character card drops its guild line in the same view, no reload.
  await expect(page.getByTestId('character-guild-line')).toHaveCount(0);
  await expect(page.getByTestId('account-characters')).toContainText('Quillfeather');
});
