// web/tests/e2e/guild-settings.spec.ts
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

test('a verified officer sees settings and can rotate the invite link with the leak warning shown', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'Fixture#1234' },
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.route('**/v1/guilds/501/invite/rotate', (route) =>
    route.fulfill(
      envelope({
        token: 'newtoken1',
        url: 'https://foreversixty.gg/guild/invite/newtoken1',
        rotated_at: '2026-09-21T00:00:00Z',
      }),
    ),
  );

  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings')).toBeVisible();
  await expect(
    page.getByText('Anyone with this link can join as a member. Rotate it if it leaks.'),
  ).toBeVisible();

  await page.getByTestId('guild-invite-rotate').click();
  await expect(page.getByTestId('guild-invite-token')).toContainText('newtoken1');
});

test('a non-officer sees the honest forbidden line, never a bare 403', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill({
      status: 403,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'forbidden' }, request_id: 'r' }),
    }),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-forbidden')).toHaveText(
    'You need to be a verified officer of this guild to see its settings.',
  );
});

// A contest always freezes officer tools now (a later security-review response simplified
// the freeze rule to exactly `state === 'contested'`) -- there is no longer a
// contested-but-not-frozen case.
test('a contested claim disables officer controls and shows the frozen notice', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'Fixture#1234' },
        claim_pending: null,
        claim: { state: 'contested', since: '2026-09-19T00:00:00Z', frozen: true },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-frozen')).toBeVisible();
  await expect(page.getByTestId('guild-visibility-select')).toBeDisabled();
  await expect(page.getByTestId('guild-officer-threshold-input')).toBeDisabled();
  await expect(page.getByTestId('guild-invite-rotate')).toBeDisabled();
  // The plain claim-state line is suppressed while frozen -- the frozen alert replaces it,
  // never both at once.
  await expect(page.getByTestId('guild-settings-claim-state')).toHaveCount(0);
});
