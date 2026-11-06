// web/tests/e2e/guild-settings.spec.ts
import { apiUnreachable, expect, test, GUILD_OVERVIEW_READS } from './support/test';

test.beforeEach(({ page }) => apiUnreachable(page, ...GUILD_OVERVIEW_READS));

// Reads this spec does not answer go unanswered, as the API out of reach (support/test.ts).
test.beforeEach(({ page }) => apiUnreachable(page, '**/v1/guilds/*/billing'));

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

test('a verified officer sees settings and can rotate the invite link', async ({ page, context }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
        invite: { rotated_at: null, url: null, made_by: null },
      }),
    ),
  );
  await page.route('**/v1/guilds/501/invite/rotate', (route) =>
    route.fulfill(
      envelope({
        token: 'newtoken1',
        url: '/guild/invite/newtoken1',
        rotated_at: '2026-09-21T00:00:00Z',
        invite: { rotated_at: '2026-09-21T00:00:00Z', url: null, made_by: null },
      }),
    ),
  );

  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings')).toBeVisible();
  await expect(page.getByText('as a verified member: no officer has to approve them')).toBeVisible();

  // A first link needs no confirm: nothing breaks.
  await expect(page.getByTestId('guild-invite-none')).toHaveText(
    'Make a link to share in guild chat or Discord.',
  );
  await page.getByTestId('guild-invite-make').click();
  // The full URL on this site, in a field, with a Copy button that reads "Copied" after a click.
  await expect(page.getByTestId('guild-invite-url')).toHaveValue(
    new RegExp('^https?://[^/]+/guild/invite/newtoken1$'),
  );
  await expect(page.getByTestId('guild-invite-copy')).toBeFocused();
  await expect(page.getByTestId('guild-invite-shown-once')).toContainText('Copy it now.');
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await page.getByTestId('guild-invite-copy').click();
  await expect(page.getByTestId('guild-invite-copy')).toHaveText('Copied');
  expect(await page.evaluate(() => navigator.clipboard.readText())).toContain('/guild/invite/newtoken1');
  await expect(page.getByTestId('guild-invite-copy')).toHaveText('Copy link', { timeout: 6000 });
});

test('a later visit has no URL in the page, and a new link asks first', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-21T00:00:00Z', frozen: false, yours: true },
        invite: { rotated_at: '2026-09-12T00:00:00Z', url: null, made_by: null },
      }),
    ),
  );
  await page.route('**/v1/guilds/501/invite/rotate', (route) =>
    route.fulfill(
      envelope({
        token: 'newtoken2',
        url: '/guild/invite/newtoken2',
        rotated_at: '2026-09-21T00:00:00Z',
        invite: { rotated_at: '2026-09-21T00:00:00Z', url: null, made_by: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-invite-made')).toHaveText(
    'Made 12 Sep. It still works, but it cannot be shown again. Make a new one to copy it.',
  );
  await expect(page.getByTestId('guild-invite-url')).toHaveCount(0);
  expect(await page.content()).not.toContain('/guild/invite/');

  await page.getByTestId('guild-invite-rotate').click();
  await expect(page.getByTestId('guild-invite-rotate-confirm')).toContainText(
    'Make a new link? The old one stops working, so anyone you sent it to needs the new one.',
  );
  await expect(page.getByTestId('guild-invite-rotate-cancel')).toBeFocused();
  await page.keyboard.press('Escape');
  await expect(page.getByTestId('guild-invite-rotate-confirm')).toHaveCount(0);
  await expect(page.getByTestId('guild-invite-rotate')).toBeFocused();

  await page.getByTestId('guild-invite-rotate').click();
  await page.getByTestId('guild-invite-rotate-confirm-button').click();
  await expect(page.getByTestId('guild-invite-url')).toHaveValue(/\/guild\/invite\/newtoken2$/);
});

test('a kept link shows again on a later visit, with who made it, and a new one replaces it', async ({
  page,
}) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  const thornhide = { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' };
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: thornhide,
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-21T00:00:00Z', frozen: false, yours: true },
        invite: { rotated_at: '2026-09-12T00:00:00Z', url: '/guild/invite/keptold', made_by: thornhide },
      }),
    ),
  );
  await page.route('**/v1/guilds/501/invite/rotate', (route) =>
    route.fulfill(
      envelope({
        token: 'keptnew',
        url: '/guild/invite/keptnew',
        rotated_at: '2026-09-21T00:00:00Z',
        invite: { rotated_at: '2026-09-21T00:00:00Z', url: '/guild/invite/keptnew', made_by: thornhide },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-invite-url')).toHaveValue(/\/guild\/invite\/keptold$/);
  await expect(page.getByTestId('guild-invite-made')).toHaveText(/^Made 12 Sep by\s+Thornhide\.$/);
  await expect(page.getByTestId('guild-invite-shown-once')).toHaveCount(0);

  await page.getByTestId('guild-invite-rotate').click();
  await page.getByTestId('guild-invite-rotate-confirm-button').click();
  await expect(page.getByTestId('guild-invite-url')).toHaveValue(/\/guild\/invite\/keptnew$/);
  await expect(page.getByTestId('guild-invite-made')).toHaveText(/^Made 21 Sep by\s+Thornhide\.$/);
});

test('the claim line names the holder by character, never an account tag', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) =>
    route.fulfill(
      envelope({
        default_visibility: 'guild',
        officer_max_rank_index: 1,
        claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
        claim_pending: null,
        claim: { state: 'claimed', since: '2026-09-21T00:00:00Z', frozen: false },
        invite: { rotated_at: null, url: null, made_by: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-claimed-by')).toHaveText(/Claimed by\s+Thornhide/);
  await expect(page.locator('body')).not.toContainText(/#[0-9]{4}/);
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
        claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
        claim_pending: null,
        claim: { state: 'contested', since: '2026-09-19T00:00:00Z', frozen: true },
        invite: { rotated_at: null, url: null, made_by: null },
      }),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-frozen')).toBeVisible();
  await expect(page.getByTestId('guild-visibility-select')).toBeDisabled();
  await expect(page.getByTestId('guild-officer-threshold-input')).toBeDisabled();
  // No link yet, so the row offers its first link; the freeze disables that press.
  await expect(page.getByTestId('guild-invite-make')).toBeDisabled();
  // The plain claim-state line is suppressed while frozen -- the frozen alert replaces it,
  // never both at once.
  await expect(page.getByTestId('guild-settings-claim-state')).toHaveCount(0);
});

test('the officer rank has a Save control and says "Saved." once it took', async ({ page }) => {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  const settings = {
    default_visibility: 'guild',
    officer_max_rank_index: 1,
    claimed_by: { name: 'Thornhide', class: 'druid', character_key: 'us/hardcore/thornhide' },
    claim_pending: null,
    claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
    invite: { rotated_at: null, url: null, made_by: null },
  };
  const patches: unknown[] = [];
  await page.route('**/v1/guilds/501/settings', async (route) => {
    if (route.request().method() === 'PATCH') {
      const body = route.request().postDataJSON() as { officer_max_rank_index: number };
      patches.push(body);
      settings.officer_max_rank_index = body.officer_max_rank_index;
    }
    await route.fulfill(envelope(settings));
  });
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  // The label has no trailing colon, and Save stays off until the number changes.
  await expect(page.getByText('Officer ranks', { exact: true })).toBeVisible();
  const save = page.getByTestId('guild-officer-threshold-save');
  await expect(save).toBeDisabled();
  await page.getByTestId('guild-officer-threshold-input').fill('2');
  await expect(save).toBeEnabled();
  await save.click();
  await expect(page.getByTestId('guild-officer-threshold-saved')).toHaveText('Saved.');
  expect(patches).toEqual([{ officer_max_rank_index: 2 }]);
});
