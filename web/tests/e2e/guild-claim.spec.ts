// web/tests/e2e/guild-claim.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const GUILD_PAGE = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  progression: [],
  roster_best: [],
  reports: [],
};

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [],
};

test('an unclaimed guild shows the claim button to a signed-in officer, and claiming shows "claimed by you"', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: false,
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim')).toBeVisible();
  await expect(page.getByTestId('guild-claim-state')).toHaveText('Nobody has claimed this guild yet.');
  await expect(page.getByTestId('guild-claim-button')).toBeVisible();

  await page.route('**/v1/guilds/501/claim', (route) => route.fulfill(envelope({ status: 'confirmed' })));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'Fixture#1234' },
        claim_pending: false,
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.getByTestId('guild-claim-button').click();
  await expect(page.getByTestId('guild-claim-state')).toHaveText('You claimed this guild.');
  await expect(page.getByTestId('guild-claim-release')).toBeVisible();
});

test('a signed-out visitor sees a sign-in prompt, not a claim button', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: false,
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-signin')).toBeVisible();
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});
