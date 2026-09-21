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

// An officer-rank member of the guild being claimed (Finding 2: only an officer/leader
// rank in THIS guild may see the claim/confirm controls).
const ME_OFFICER = {
  ...ME,
  guilds: [
    { id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch', rank: 'officer', verified: true },
  ],
};

test('an unclaimed guild shows the claim button to a signed-in officer, and claiming shows "claimed by you"', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: null,
        claim: { state: 'unclaimed', frozen: false },
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
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-21T00:00:00Z', frozen: false },
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
        claim_pending: null,
        claim: { state: 'unclaimed', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-signin')).toBeVisible();
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});

test('a signed-in visitor with no membership, or member rank, in this guild sees the not-eligible line, not a claim button', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  // `ME` above has no membership at all in guild id 501, which is the more common of the
  // two ineligible shapes (a stranger); a plain `rank: 'member'` entry is refused the same
  // way, exercised by the `eligible` check reading a matched membership's own rank.
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: null,
        claim: { state: 'unclaimed', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-not-eligible')).toHaveText(
    'Only an officer or the guild master of this guild can claim it.',
  );
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});

// A pending claim now carries a real `ClaimPendingView` object (`{by, expires_at}`), never
// a boolean -- this exercises the confirm flow off that real shape end to end.
test('a pending claim shows who is confirming it and lets a second officer confirm', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: { by: { battletag: 'OtherOfficer#5678' }, expires_at: '2026-09-28T00:00:00Z' },
        claim: { state: 'pending', since: '2026-09-21T00:00:00Z', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.route('**/v1/guilds/501/claim/confirm', (route) =>
    route.fulfill(envelope({ status: 'confirmed', claimed_by: { battletag: 'OtherOfficer#5678' } })),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-state')).toHaveText(
    'A claim is pending, confirmed by a second officer or the guild master. Expires 2026-09-28.',
  );
  await expect(page.getByTestId('guild-claim-confirm')).toBeVisible();
});

// The claim rules blurb (guildClaimCopy.rules) is new: it renders whenever settings loaded
// successfully, regardless of claim state, so this asserts it independently of any one
// claim-state scenario above.
test('the claiming rules render alongside the claim state', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: null,
        claim_pending: null,
        claim: { state: 'unclaimed', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-rules')).toContainText('A claim can be contested.');
});

// The contested-state notice is a NEW, dedicated testid that renders ADDITIONALLY
// alongside the claimed_by/claim_pending/unclaimed paragraph, not instead of it -- both
// must be visible at once.
test('a contested claim shows the dedicated contested notice alongside the claimed-by state', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'OtherOfficer#5678' },
        claim_pending: null,
        claim: { state: 'contested', since: '2026-09-19T00:00:00Z', frozen: true },
        invite: { rotated_at: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-contested-notice')).toHaveText(
    'This guild’s claim is contested and under review by a moderator.',
  );
  await expect(page.getByTestId('guild-claim-state')).toHaveText('Claimed by OtherOfficer#5678.');
  // Not the claimant, so no release control -- this is a different account's claim.
  await expect(page.getByTestId('guild-claim-release')).toHaveCount(0);
});

// Regression guard for Task 3's own fix: while a claim is BOTH held by the viewer AND
// contested, Release must stay visible and enabled -- contesting must never strand the
// claimant with no way to release the guild.
test('release stays visible and enabled for the claim holder even while their claim is contested', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
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
  await page.goto('/guild/us/hardcore/the-last-watch/claim');
  await expect(page.getByTestId('guild-claim-contested-notice')).toBeVisible();
  await expect(page.getByTestId('guild-claim-state')).toHaveText('You claimed this guild.');
  await expect(page.getByTestId('guild-claim-release')).toBeVisible();
  await expect(page.getByTestId('guild-claim-release')).toBeEnabled();
});
