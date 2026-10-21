// web/tests/e2e/guild-invite.spec.ts
// The fixture-only [...path].astro route only prerenders the exact invite token
// FIXTURE_GUILD_INVITE_TOKEN ('fixtureinvitetoken1', see src/lib/report/shell-paths.ts) --
// navigating to any other /guild/invite/<token> path 404s under the static build this
// suite runs against (`npm run build && npm run preview`), so the token below must match
// that fixture path exactly rather than an arbitrary placeholder.
import { expect, test, type Page } from '@playwright/test';

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

const PEEK = {
  guild: { id: 501, region: 'us', ruleset: 'pvp', name: 'Lifecycle Test', faction: 'alliance' },
};

async function signedOut(page: Page): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
}

test('a signed-in visitor sees the guild that invited them, joins, and is sent to the guild page', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route(`**/v1/guild-invites/${INVITE_TOKEN}`, (route) => route.fulfill(envelope(PEEK)));
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
  await expect(page.getByTestId('guild-join-card')).toBeVisible();
  await expect(page.getByTestId('guild-join-name')).toHaveText('Lifecycle Test');
  await expect(page.getByTestId('guild-join-invited')).toHaveText('Lifecycle Test · PvP US invited you.');
  await expect(page.getByTestId('guild-join-consequence')).toHaveText(
    'Joining puts your characters on its roster and shares your gear with its officers; you can change that on your account.',
  );
  await expect(page.getByTestId('guild-join-crest')).toBeVisible();
  await expect(page.getByTestId('guild-join-button')).toHaveText('Join Lifecycle Test');
  // No horizontal scroll at phone width: the card fits the page.
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - document.documentElement.clientWidth,
  );
  expect(overflow).toBeLessThanOrEqual(0);
  await page.getByTestId('guild-join-button').click();
  await expect(page).toHaveURL(/\/guild\/us\/hardcore\/the-last-watch$/);
});

test('a signed-out visitor sees the same card with a sign-in button, not a join button', async ({ page }) => {
  await signedOut(page);
  await page.route(`**/v1/guild-invites/${INVITE_TOKEN}`, (route) => route.fulfill(envelope(PEEK)));
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join-invited')).toHaveText('Lifecycle Test · PvP US invited you.');
  await expect(page.getByTestId('guild-join-consequence')).toBeVisible();
  await expect(page.getByTestId('guild-join-signin')).toHaveText('Sign in with Battle.net to join');
  // Signing in brings them back to this same invite.
  await expect(page.getByTestId('guild-join-signin')).toHaveAttribute(
    'href',
    new RegExp(`next=${encodeURIComponent(`/guild/invite/${INVITE_TOKEN}`)}`),
  );
  await expect(page.getByTestId('guild-join-button')).toHaveCount(0);
});

test('a bad or rotated token says so plainly and offers no button', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route(`**/v1/guild-invites/${INVITE_TOKEN}`, (route) =>
    route.fulfill({
      status: 404,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'not_found', message: 'that invite link is not valid' },
        request_id: 'r',
      }),
    }),
  );
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join-invalid')).toHaveText(
    'This invite link no longer works. Ask a guild officer for a new one.',
  );
  await expect(page.getByTestId('guild-join-button')).toHaveCount(0);
  await expect(page.getByTestId('guild-join-signin')).toHaveCount(0);
});

test('a member of the guild already is told so, with a way in', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        ...ME,
        guilds: [
          { id: 501, region: 'us', ruleset: 'pvp', name: 'Lifecycle Test', rank: 'member', verified: true },
        ],
      }),
    ),
  );
  await page.route(`**/v1/guild-invites/${INVITE_TOKEN}`, (route) => route.fulfill(envelope(PEEK)));
  await page.goto(`/guild/invite/${INVITE_TOKEN}`);
  await expect(page.getByTestId('guild-join-already')).toHaveText('You are already in Lifecycle Test.');
  await expect(page.getByTestId('guild-join-open')).toHaveAttribute('href', '/guild/us/pvp/lifecycle-test');
  await expect(page.getByTestId('guild-join-button')).toHaveCount(0);
});
