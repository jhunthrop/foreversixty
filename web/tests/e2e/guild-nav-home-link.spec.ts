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
  guilds: [{ id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'member' }],
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

test('a signed-in member sees "My guild" on the homepage', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_WITH_GUILD)));
  await page.goto('/');
  await expect(page.getByTestId('home-my-guild')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch',
    // client:visible on an already-visible hero element hydrates promptly, but this still
    // waits for the assertion's own retry loop rather than assuming synchronous readiness.
  );
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
