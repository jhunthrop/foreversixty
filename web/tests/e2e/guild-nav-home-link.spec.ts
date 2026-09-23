// web/tests/e2e/guild-nav-home-link.spec.ts
import { expect, test } from '@playwright/test';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

const ME_WITH_GUILD = {
  user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false, premium: false },
  characters: [],
  guilds: [
    { id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'member', verified: true },
  ],
};

// The home page's "Your guild" panel (spec 2026-09-23 §2 item 5) reads the guild's
// progression from GET /v1/guilds/<region>/<ruleset>/<slug> once a guild is known, so the
// two homepage tests below stub it: one boss down of two.
const GUILD_PAGE = {
  guild: { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' },
  progression: [
    { encounter: 'Skolex', encounter_id: 1, difficulty: 1, kills: 2, pull_count: 5 },
    { encounter: 'Warden Kelthas', encounter_id: 2, difficulty: 1, kills: 0, pull_count: 3 },
  ],
  roster_best: [],
  reports: [],
};

test('a signed-in member sees "My guild" in the header on a tool page', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_WITH_GUILD)));
  await page.goto('/planner');
  await expect(page.getByTestId('session-my-guild')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch',
  );
});

test('a signed-out visitor never sees "My guild" in the header', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  await page.goto('/planner');
  await expect(page.getByTestId('session-my-guild')).toHaveCount(0);
});

test('a signed-in member sees "My guild" and its progression on the homepage', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_WITH_GUILD)));
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.goto('/');
  await expect(page.getByTestId('home-my-guild')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch',
    // client:visible on an already-visible hero element hydrates promptly, but this still
    // waits for the assertion's own retry loop rather than assuming synchronous readiness.
  );
  await expect(page.getByTestId('home-guild-ready')).toContainText('1 of 2 bosses down');
});

test('a signed-out visitor sees no "My guild" on the homepage and no extra script for it', async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 401,
      contentType: 'application/json',
      body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
    }),
  );
  await page.goto('/');
  await expect(page.getByTestId('home-my-guild')).toHaveCount(0);
});
