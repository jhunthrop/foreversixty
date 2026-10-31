// web/tests/e2e/guild-account-consent.spec.ts
import { expect, test } from './support/test';
import { choose } from './support/select';

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
    characters: ['Quillfeather', 'Brenna', 'Corvin'].map((name) => ({
      ...CHARACTER,
      key: `us/hardcore/${name.toLowerCase()}`,
      name,
      ...(inGuild ? {} : { guild: undefined }),
    })),
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
  let leaves = 0;
  await page.route('**/v1/guilds/501/members/me', (route) => {
    if (route.request().method() === 'PATCH') {
      consent = 'roster';
      return route.fulfill(envelope({ consent: 'roster' }));
    }
    inGuild = false;
    leaves += 1;
    return route.fulfill(envelope({ status: 'left' }));
  });

  await page.goto('/account');
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');
  await expect(page.getByTestId('character-guild-line').first()).toContainText('The Last Watch');
  await choose(page.getByTestId('account-guild-consent'), 'roster');
  await expect(page.getByTestId('account-guild-consent')).toHaveValue('roster');
  await expect(page.getByTestId('account-guild-consent-saved')).toHaveText(
    'Saved. Applies to all 3 of your characters.',
  );

  // Leave asks first, naming what it removes; Cancel changes nothing.
  await page.getByTestId('account-guild-leave').click();
  await expect(page.getByTestId('account-guild-leave-confirm')).toContainText(
    'Leave The Last Watch? This removes all 3 of your characters from its roster.',
  );
  // A member holds no claim, so the confirm says nothing about releasing one.
  await expect(page.getByTestId('account-guild-leave-releases')).toHaveCount(0);
  await page.getByTestId('account-guild-leave-confirm').getByRole('button', { name: 'Cancel' }).click();
  await expect(page.getByTestId('account-guild-leave-confirm')).toHaveCount(0);
  expect(leaves).toBe(0);
  await expect(page.getByTestId('account-guilds')).toContainText('The Last Watch');

  await page.getByTestId('account-guild-leave').click();
  await page.getByTestId('account-guild-leave-confirm-button').click();
  await expect(page.getByTestId('account-guilds')).toHaveCount(0);
  expect(leaves).toBe(1);
  // The character cards drop its guild line in the same view, no reload.
  await expect(page.getByTestId('character-guild-line')).toHaveCount(0);
  await expect(page.getByTestId('account-characters')).toContainText('Quillfeather');
});

test('the guild master is warned that leaving releases the claim and what the officers lose', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        ...meFor('gear', true),
        guilds: [
          {
            ...GUILD,
            name: 'Lifecycle Test',
            rank: 'leader',
            can: ['approve', 'remove', 'settings', 'award', 'billing', 'release'],
          },
        ],
      }),
    ),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto('/account');
  await page.getByTestId('account-guild-leave').click();
  const confirm = page.getByTestId('account-guild-leave-confirm');
  await expect(confirm).toContainText(
    'Leave Lifecycle Test? This removes all 3 of your characters from its roster.',
  );
  await expect(page.getByTestId('account-guild-leave-releases')).toHaveText(
    'Leaving Lifecycle Test means you stop running it here. Your officers lose the roster and loot tools and can no longer see your gear.',
  );
  await expect(page.getByTestId('account-guild-leave-no-handover')).toHaveText(
    'There is no hand-over yet; if you want an officer to run the guild after you, ask us on Discord before you leave.',
  );
  // Leave is the red primary and Cancel the plain one, with focus on Cancel; the consent select is labelled.
  const leave = page.getByTestId('account-guild-leave-confirm-button');
  await expect(leave).toHaveText('Leave guild');
  expect(await leave.evaluate((node) => getComputedStyle(node).backgroundColor)).not.toBe('rgba(0, 0, 0, 0)');
  await expect(page.getByTestId('account-guild-leave-cancel')).toBeFocused();
  await expect(page.getByRole('combobox', { name: 'Share with this guild' })).toBeVisible();
  // P.1: nothing under the open confirm -- no hairline on the row, no empty line after the card.
  const row = page.getByTestId('account-guild-row');
  await expect(row).toHaveCSS('border-bottom-width', '0px');
  expect(await row.evaluate((node) => node.lastElementChild?.getAttribute('data-testid'))).toBe(
    'account-guild-leave-confirm',
  );
  const dir = process.env.G7_SHOTS;
  if (dir) await page.screenshot({ path: `${dir}/g7-leave-claimant-1280.png`, fullPage: true });
  // The confirm clears the plan card below it by a gutter, at a phone's width too.
  await page.setViewportSize({ width: 390, height: 900 });
  const [box, plan] = await Promise.all([
    confirm.boundingBox(),
    page.getByTestId('account-guilds-plan').boundingBox(),
  ]);
  expect(box && plan).toBeTruthy();
  if (box && plan) expect(plan.y - (box.y + box.height)).toBeGreaterThanOrEqual(12);
  if (dir) await page.screenshot({ path: `${dir}/g7-leave-claimant-390.png`, fullPage: true });
});

test('a plain officer is told they lose their officer tools, and nothing about a claim', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        ...meFor('gear', true),
        guilds: [
          {
            ...GUILD,
            name: 'Lifecycle Test',
            rank: 'officer',
            can: ['approve', 'remove', 'settings', 'award'],
          },
        ],
      }),
    ),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto('/account');
  await page.getByTestId('account-guild-leave').click();
  await expect(page.getByTestId('account-guild-leave-officer')).toHaveText(
    'You lose your officer tools in Lifecycle Test.',
  );
  await expect(page.getByTestId('account-guild-leave-releases')).toHaveCount(0);
  await expect(page.getByTestId('account-guild-leave-no-handover')).toHaveCount(0);
});
