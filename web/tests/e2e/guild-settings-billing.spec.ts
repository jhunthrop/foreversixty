// web/tests/e2e/guild-settings-billing.spec.ts
// GuildSettings.svelte's billing section (reconciliation with the landed API lane): a
// guild's billing view is not part of GET /v1/guilds/{id}/settings's own response -- it
// lives on GET /v1/me's guilds[].plan for this same guild (spec 1.4). This file is separate
// from tests/e2e/guild-settings.spec.ts (owned by the guild-web lane) so this lane's new
// coverage never edits a file it does not own.
import { expect, test } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';

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

const SETTINGS = {
  default_visibility: 'guild',
  officer_max_rank_index: 1,
  claimed_by: { battletag: 'Fixture#1234' },
  claim_pending: null,
  claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
  invite: { rotated_at: null },
};

function meWithGuildPlan(plan: unknown) {
  return {
    ok: true,
    data: {
      user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
      characters: [],
      guilds: [
        {
          id: 501,
          region: 'us',
          ruleset: 'hardcore',
          name: 'The Last Watch',
          rank: 'officer',
          verified: true,
          plan,
        },
      ],
      entitlements: mockEntitlements(),
    },
    error: null,
    request_id: 'r',
  };
}

async function routeCommon(page: import('@playwright/test').Page): Promise<void> {
  await page.route('**/v1/guilds/us/hardcore/the-last-watch', (route) => route.fulfill(envelope(GUILD_PAGE)));
  await page.route('**/v1/guilds/501/settings', (route) => route.fulfill(envelope(SETTINGS)));
}

test('not on the plan: shows the honest not-subscribed line and a real subscribe link', async ({ page }) => {
  await routeCommon(page);
  await page.route('**/v1/me', (route) => route.fulfill({ ...envelope(meWithGuildPlan(null)), status: 200 }));
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('not on the guild plan yet');
  await expect(page.getByTestId('guild-subscribe-link')).toHaveAttribute(
    'href',
    '/premium/checkout?plan=guild&interval=monthly&guild_id=501',
  );
});

test('on the plan, billing contact: shows renewal info and a working manage-billing button', async ({
  page,
}) => {
  await routeCommon(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope(
        meWithGuildPlan({
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          billed_by: 'you',
          you_are_billing_contact: true,
        }).data,
      ),
    ),
  );
  let posted: unknown;
  await page.route('**/v1/billing/portal', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill(envelope({ portal_url: 'https://billing.stripe.com/guild' }));
  });
  await page.route('https://billing.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('Renews');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('Billed by you');
  await page.getByTestId('guild-manage-billing').click();
  await page.waitForURL('https://billing.stripe.com/**');
  expect(posted).toEqual({ guild_id: 501 });
});

test('on the plan, not the billing contact: shows renewal info but no manage-billing button', async ({
  page,
}) => {
  await routeCommon(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope(
        meWithGuildPlan({
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          billed_by: 'Otherguy#5678',
          you_are_billing_contact: false,
        }).data,
      ),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('Billed by Otherguy#5678');
  await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
});

test('on the plan via a CLI grant (no Stripe billing contact): shows renewal info, never a blank "Billed by"', async ({
  page,
}) => {
  // auth/handler.go's attachGuildPlans leaves billed_by "" and you_are_billing_contact
  // false when GuildBilling.BillingUserID is nil -- a CLI grant, never billed through
  // Stripe, has no billing contact to name (its own comment: "no billing contact to
  // name"). The settings page must never render "Billed by" with nothing after it.
  await routeCommon(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope(
        meWithGuildPlan({
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          billed_by: '',
          you_are_billing_contact: false,
        }).data,
      ),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('Renews');
  await expect(page.getByTestId('guild-settings-billing')).not.toContainText('Billed by');
  await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
});
