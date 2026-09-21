// web/tests/e2e/guild-invite.spec.ts
// The fixture-only [...path].astro route only prerenders the exact invite token
// FIXTURE_GUILD_INVITE_TOKEN ('fixtureinvitetoken1', see src/lib/report/shell-paths.ts) --
// navigating to any other /guild/invite/<token> path 404s under the static build this
// suite runs against (`npm run build && npm run preview`), so the token below must match
// that fixture path exactly rather than an arbitrary placeholder.
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const INVITE_TOKEN = 'fixtureinvitetoken1';

const ME = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [],
};

test('a signed-in visitor can join by invite and is sent to the guild page', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route(`**/v1/guilds/invite/${INVITE_TOKEN}/accept`, (route) =>
    route.fulfill(
      envelope({
        guild: { id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
        rank: 'member',
      }),
    ),
  );
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) =>
    route.fulfill(
      envelope({
        guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
        progression: [],
        roster_best: [],
        reports: [],
      }),
    ),
  );

  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join')).toBeVisible();
  await page.getByTestId('guild-join-button').click();
  await expect(page).toHaveURL(/\/guild\/us\/hardcore\/the-last-watch$/);
});

test('a signed-out visitor sees a sign-in prompt, not a join button', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join-signin')).toBeVisible();
  await expect(page.getByTestId('guild-join-button')).toHaveCount(0);
});
