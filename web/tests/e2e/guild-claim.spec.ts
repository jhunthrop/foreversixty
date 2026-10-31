// web/tests/e2e/guild-claim.spec.ts
import type { Page } from '@playwright/test';
import { apiUnreachable, expect, test } from './support/test';
import { priceFor } from '../../src/lib/billing/plan-costs';
import { checkoutGuildCopy } from '../../src/lib/billing/checkout-guild-copy';
import { PLANS, formatPrice, founderOpen } from '../../src/lib/billing/plans';
import { guildClaimCopy } from '../../src/lib/guild/copy';
import { SETUP_NAV_ITEM } from '../../src/lib/nav';
import { uiCopy } from '../../src/lib/ui/copy';
import type { Me } from '../../src/lib/account/me-types';
import { buildMe, stepGuild, subscribeView } from '../../src/test-support/subscribe-fixtures';
import { gotoHydrated } from './support/hydrated';
import { expectNoAxeViolations } from './support/officer-tools';
import { expectTapHeight, routeMeCell } from './support/subscribe-stubs';

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/devices', '**/v1/guilds/*/home'));

// The claim page is an island: wait for it to hydrate before asserting on it, or a slow CI shard
// asserts against the server shell (the mobile shard did, once).
const openClaim = (page: Page, path = '/guild/us/hardcore/the-last-watch/claim'): Promise<void> =>
  gotoHydrated(page, path, 'guild-claim');

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

test('an unclaimed guild asks an officer first, then says "You hold" and where to go next', async ({
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
  await openClaim(page);
  await expect(page.getByTestId('guild-claim')).toBeVisible();
  await expect(page.getByTestId('guild-claim-state')).toHaveText(
    'Nobody has claimed this guild. Claiming lets you approve members, set report visibility and manage billing.',
  );
  await expect(page.getByTestId('guild-claim-button')).toBeVisible();
  // Nothing to contest: the claim does not exist yet.
  await expect(page.getByTestId('guild-claim-contest-button')).toHaveCount(0);

  // Claiming asks first, with the rules in one line; Cancel backs out without a request.
  let claimPosts = 0;
  await page.route('**/v1/guilds/501/claim', (route) => {
    claimPosts += 1;
    return route.fulfill(envelope({ status: 'confirmed' }));
  });
  await page.getByTestId('guild-claim-button').click();
  await expect(page.getByTestId('guild-claim-confirm-box')).toContainText('Claim The Last Watch?');
  await expect(page.getByTestId('guild-claim-confirm-box')).toContainText('one claim every 30 days');
  await page.getByTestId('guild-claim-confirm-box').getByRole('button', { name: 'Cancel' }).click();
  await expect(page.getByTestId('guild-claim-confirm-box')).toHaveCount(0);
  expect(claimPosts).toBe(0);
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
  await page.getByTestId('guild-claim-confirm-claim').click();
  await expect(page.getByTestId('guild-claim-state')).toHaveText('You hold The Last Watch.');
  expect(claimPosts).toBe(1);
  // Once the claim is held the page is about the claim.
  await expect(page.getByTestId('guild-claim-heading')).toHaveText('The Last Watch · your claim');
  await expect(page.getByTestId('guild-claim-next-roster')).toHaveText('Approve your officers and members');
  await expect(page.getByTestId('guild-claim-next-roster')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch#roster',
  );
  await expect(page.getByTestId('guild-claim-next-invite')).toHaveText('Make an invite link');
  // The next steps are two 44 px buttons: the roster is the gold primary, the invite link a ghost.
  for (const id of ['guild-claim-next-roster', 'guild-claim-next-invite']) {
    const box = await page.getByTestId(id).boundingBox();
    expect(box?.height).toBe(44);
  }
  // The pointer still rests where the claim button was, which is where the primary now sits;
  // move it off so the resting colour is read, not the hover one.
  await page.mouse.move(0, 0);
  await expect(page.getByTestId('guild-claim-next-roster')).toHaveCSS(
    'background-color',
    'rgb(229, 185, 85)',
  );
  await expect(page.getByTestId('guild-claim-next-invite')).toHaveCSS('background-color', 'rgba(0, 0, 0, 0)');
  await expect(page.getByTestId('guild-claim-next-invite')).toHaveAttribute(
    'href',
    '/guild/us/hardcore/the-last-watch/settings',
  );
  // Your own claim is not something to contest.
  await expect(page.getByTestId('guild-claim-contest-button')).toHaveCount(0);

  // Release stays, behind a confirm.
  let releasePosts = 0;
  await page.route('**/v1/guilds/501/claim/release', (route) => {
    releasePosts += 1;
    return route.fulfill(envelope({ status: 'released' }));
  });
  await page.getByTestId('guild-claim-release').click();
  await expect(page.getByTestId('guild-claim-release-confirm')).toContainText(
    'Release your claim on The Last Watch? Officers will have to claim it again.',
  );
  await page.getByTestId('guild-claim-release-confirm').getByRole('button', { name: 'Cancel' }).click();
  expect(releasePosts).toBe(0);
  await page.getByTestId('guild-claim-release').click();
  await page.getByTestId('guild-claim-release-confirm-button').click();
  await expect.poll(() => releasePosts).toBe(1);
});

test("someone else's claim names the holder and offers Contest, with its confirm", async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME_OFFICER)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { battletag: 'OtherOfficer#5678' },
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-20T00:00:00Z', frozen: false },
        invite: { rotated_at: null },
      }),
    ),
  );
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-state')).toHaveText(
    'Claimed by OtherOfficer#5678 since 20 Sep.',
  );
  await expect(page.getByTestId('guild-claim-release')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-next')).toHaveCount(0);
  await page.getByTestId('guild-claim-contest-button').click();
  await expect(page.getByTestId('guild-claim-contest-confirm')).toBeVisible();
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
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-signin')).toBeVisible();
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});

test('a signed-in visitor with no membership in this guild is told so, and sees no claim button', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  // `ME` above has no membership at all in guild id 501: a stranger. Since 2026-09-28 the
  // page names that case (the site knows no character of theirs here) rather than the
  // generic not-eligible line; a plain `rank: 'member'` entry gets the member line instead
  // (the "refused settings" case further down).
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
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-not-member')).toContainText(
    'does not know a character of yours',
  );
  await expect(page.getByTestId('guild-claim-not-eligible')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});

// --- a claim link made for a plan (subscribe-and-join spec §6.2, case 17) ---------------------------

const CLAIM_PATH = '/guild/us/hardcore/the-last-watch/claim';
const PLAN_QUERY = '?plan=guild&interval=yearly';
const LAST_WATCH = {
  id: 501,
  name: 'The Last Watch',
  region: 'us',
  ruleset: 'hardcore',
  slug: 'the-last-watch',
} as const;
const OWNER_BATTLETAG = 'hunthrop#1894';

const lastWatchMe = (guild: Parameters<typeof stepGuild>[0]): Me =>
  buildMe({ guilds: [stepGuild({ ...LAST_WATCH, ...guild })] });

const unclaimedSettings = {
  default_visibility: 'guild',
  officer_max_rank_index: 1,
  claimed_by: null,
  claim_pending: null,
  claim: { state: 'unclaimed', frozen: false },
  invite: { rotated_at: null },
};
const heldSettings = {
  ...unclaimedSettings,
  claimed_by: { battletag: OWNER_BATTLETAG },
  claim: { state: 'claimed', since: '2026-10-20T00:00:00Z', frozen: false },
};
const FORBIDDEN = {
  status: 403,
  contentType: 'application/json',
  body: '{"ok":false,"data":null,"error":{"code":"forbidden","message":"officers only"},"request_id":"r"}',
};

async function stubGuildPage(page: Page): Promise<void> {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
}

const planPriceText = (): string =>
  formatPrice(priceFor('guild', 'yearly', founderOpen(new Date())), 'yearly');

test('signed out, the sign-in keeps next on the claim page and its plan', async ({ page }) => {
  await stubGuildPage(page);
  await routeMeCell(page, () => null);
  await page.route('**/v1/guilds/501/settings', (route) => route.fulfill(FORBIDDEN));
  await openClaim(page, `${CLAIM_PATH}${PLAN_QUERY}`);
  const signIn = page
    .getByTestId('guild-claim-signin')
    .getByRole('link', { name: 'Sign in with Battle.net' });
  await expect(signIn).toHaveAttribute(
    'href',
    new RegExp(`next=${encodeURIComponent(`${CLAIM_PATH}${PLAN_QUERY}`)}$`),
  );
  await expect(page.getByTestId('guild-claim-plan-line')).toHaveText(
    guildClaimCopy.planLine(LAST_WATCH.name, PLANS.guild.name, planPriceText()),
  );
});

test('after a claim made for a plan, subscribe leads Next as the one gold fill', async ({ page }) => {
  let claimed = false;
  await stubGuildPage(page);
  await routeMeCell(page, () =>
    claimed
      ? lastWatchMe({
          rank: 'leader',
          verified: true,
          verified_count: 12,
          claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Obnoxious Yell' },
          subscribe: subscribeView('ready', { trial_used: false }),
        })
      : lastWatchMe({ step: 'claim', rank: 'leader' }),
  );
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(envelope(claimed ? heldSettings : unclaimedSettings)),
  );
  await page.route('**/v1/guilds/501/claim', (route) => {
    claimed = true;
    return route.fulfill(envelope({ status: 'confirmed' }));
  });
  await openClaim(page, `${CLAIM_PATH}${PLAN_QUERY}`);
  await expect(page.getByTestId('guild-claim-plan-line')).toHaveText(
    guildClaimCopy.planLine(LAST_WATCH.name, PLANS.guild.name, planPriceText()),
  );
  await page.getByTestId('guild-claim-button').click();
  await page.getByTestId('guild-claim-confirm-claim').click();

  const next = page.getByTestId('guild-claim-next');
  await expect(next).toBeVisible();
  const order = await next
    .locator('[data-testid^="guild-claim-next-"]')
    .evaluateAll((nodes) => nodes.map((node) => node.getAttribute('data-testid')));
  expect(order).toEqual(['guild-claim-next-subscribe', 'guild-claim-next-roster', 'guild-claim-next-invite']);
  // The viewer now holds the claim: "Claim the guild first" is no longer theirs to read (row 16).
  await expect(page.getByTestId('guild-claim-plan-line')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-next-subscribe')).toHaveText(checkoutGuildCopy.ready.startTrial);
  await expect(page.getByTestId('guild-claim-next-roster')).toHaveText(guildClaimCopy.nextRoster);

  // One fill: the subscribe button is gold, the roster link and the invite link are ghosts.
  await page.mouse.move(0, 0);
  await expect(page.getByTestId('guild-claim-next-subscribe')).toHaveCSS(
    'background-color',
    'rgb(229, 185, 85)',
  );
  for (const ghost of ['guild-claim-next-roster', 'guild-claim-next-invite']) {
    await expect(page.getByTestId(ghost)).toHaveCSS('background-color', 'rgba(0, 0, 0, 0)');
  }
  for (const control of ['subscribe', 'roster', 'invite']) {
    await expectTapHeight(page.getByTestId(`guild-claim-next-${control}`));
  }
  await expectNoAxeViolations(page, '[data-testid="guild-claim"]');

  const posted: unknown[] = [];
  await page.route('**/v1/billing/checkout', (route) => {
    posted.push(route.request().postDataJSON());
    return route.fulfill(envelope({ checkout_url: 'https://checkout.stripe.com/x' }));
  });
  await page.route('https://checkout.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.getByTestId('guild-claim-next-subscribe').click();
  await page.waitForURL('https://checkout.stripe.com/**');
  expect(posted).toEqual([{ plan: 'guild', interval: 'yearly', guild_id: LAST_WATCH.id }]);
});

test('a member sees the plan line and a copy field holding the claim link with the same plan', async ({
  page,
}) => {
  await stubGuildPage(page);
  await routeMeCell(page, () => lastWatchMe({ rank: 'member' }));
  await page.route('**/v1/guilds/501/settings', (route) => route.fulfill(FORBIDDEN));
  await page.route('**/v1/guilds/501/home', (route) =>
    route.fulfill(
      envelope({
        guild: LAST_WATCH,
        claim: { state: 'unclaimed', frozen: false },
        reports: [],
        roster: [],
      }),
    ),
  );
  await openClaim(page, `${CLAIM_PATH}${PLAN_QUERY}`);
  await expect(page.getByTestId('guild-claim-plan-line')).toBeVisible();
  await expect(page.getByTestId('guild-claim-not-eligible')).toContainText(
    guildClaimCopy.memberHandoff(LAST_WATCH.name),
  );
  const origin = new URL(page.url()).origin;
  await expect(page.getByTestId('guild-claim-handoff-url')).toHaveValue(
    `${origin}${CLAIM_PATH}?plan=guild&interval=yearly`,
  );
  await expect(page.getByTestId('guild-claim-handoff-copy')).toHaveText(uiCopy.copyLink);
  await expectTapHeight(page.getByTestId('guild-claim-handoff-copy'));
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
  await expectNoAxeViolations(page, '[data-testid="guild-claim"]');
});

test('a stranger is offered Get set up, which carries this claim page as next', async ({ page }) => {
  await stubGuildPage(page);
  await routeMeCell(page, () => buildMe({ guilds: [] }));
  await page.route('**/v1/guilds/501/settings', (route) => route.fulfill(FORBIDDEN));
  await openClaim(page, `${CLAIM_PATH}${PLAN_QUERY}`);
  const link = page.getByTestId('guild-claim-not-member').getByRole('link', { name: SETUP_NAV_ITEM.label });
  await expect(link).toHaveAttribute(
    'href',
    `/setup?next=${encodeURIComponent(`${CLAIM_PATH}${PLAN_QUERY}`)}`,
  );
});

test('/setup offers Back to the claim once a save makes the account know the guild, and on focus', async ({
  page,
}) => {
  let saved = false;
  await routeMeCell(page, () => (saved ? lastWatchMe({ rank: 'leader' }) : buildMe({ guilds: [] })));
  const claimLink = `${CLAIM_PATH}${PLAN_QUERY}`;
  await page.goto(`/setup?next=${encodeURIComponent(claimLink)}`);
  const slot = page.getByTestId('setup-back-to-claim');
  await expect(slot).toBeAttached();
  const back = slot.getByRole('link', { name: guildClaimCopy.backToClaim(LAST_WATCH.name) });
  await expect(back).toHaveCount(0);

  saved = true;
  await page.evaluate(() => window.dispatchEvent(new Event('focus')));
  await expect(back).toHaveAttribute('href', claimLink);
  await expectTapHeight(back);
  await expectNoAxeViolations(page, '[data-testid="setup-back-to-claim"]');
});

test('/setup ignores a next that is not a claim page', async ({ page }) => {
  await routeMeCell(page, () => lastWatchMe({ rank: 'member' }));
  await page.goto(`/setup?next=${encodeURIComponent('https://evil.example/guild/us/hardcore/x/claim')}`);
  await expect(page.getByTestId('setup-back-to-claim')).toHaveCount(0);
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
  await openClaim(page);
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
  await openClaim(page);
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
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-contested-notice')).toHaveText(
    'This guild’s claim is contested and under review by a moderator.',
  );
  await expect(page.getByTestId('guild-claim-state')).toHaveText(
    'Claimed by OtherOfficer#5678 since 19 Sep.',
  );
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
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-contested-notice')).toBeVisible();
  await expect(page.getByTestId('guild-claim-state')).toHaveText('You hold The Last Watch.');
  await expect(page.getByTestId('guild-claim-release')).toBeVisible();
  await expect(page.getByTestId('guild-claim-release')).toBeEnabled();
  // While frozen the next steps are not offered, and a held claim is never contestable by its holder.
  await expect(page.getByTestId('guild-claim-next')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-contest-button')).toHaveCount(0);
});

// Found live (2026-09-28, /guild/us/pvp/olympus-xxvii/claim): GET .../settings is
// officer-only, so for a plain member it answers 403 -- and the page rendered nothing but
// its heading. The claim state now comes from the member home for a member, with the
// honest line about who can act.
test('a plain member, refused settings, still sees the claim state and who can act', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        ...ME,
        guilds: [
          {
            id: 501,
            region: 'us',
            ruleset: 'hardcore',
            name: 'The Last Watch',
            rank: 'member',
            verified: false,
          },
        ],
      }),
    ),
  );
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill({
      status: 403,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'forbidden', message: 'officers only' },
        request_id: 'r',
      }),
    }),
  );
  await page.route('**/v1/guilds/501/home', (route) =>
    route.fulfill(
      envelope({
        guild: { id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
        claim: { state: 'unclaimed', frozen: false },
        reports: [],
        roster: [],
      }),
    ),
  );
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-state')).toHaveText('Nobody has claimed this guild yet.');
  await expect(page.getByTestId('guild-claim-not-eligible')).toContainText(
    'You are a member of The Last Watch',
  );
  // Nothing is claimed, so there is nothing to contest.
  await expect(page.getByTestId('guild-claim-contest-button')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-rules')).toBeVisible();
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
});

test('a signed-in stranger, refused both reads, is told the site knows no character of theirs here', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/me', (route) => route.fulfill(envelope(ME)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill({
      status: 403,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'forbidden', message: 'officers only' },
        request_id: 'r',
      }),
    }),
  );
  await openClaim(page);
  await expect(page.getByTestId('guild-claim-not-member')).toContainText(
    'does not know a character of yours',
  );
  await expect(page.getByTestId('guild-claim-button')).toHaveCount(0);
  await expect(page.getByTestId('guild-claim-contest-button')).toHaveCount(0);
});
