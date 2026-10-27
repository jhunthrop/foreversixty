// web/tests/e2e/guild-settings-billing.spec.ts
// GuildSettings.svelte's billing section (reconciliation with the landed API lane): a
// guild's billing view is not part of GET /v1/guilds/{id}/settings's own response -- it
// lives on GET /v1/me's guilds[].plan for this same guild (spec 1.4). This file is separate
// from tests/e2e/guild-settings.spec.ts (owned by the guild-web lane) so this lane's new
// coverage never edits a file it does not own.
import { expect, test } from '@playwright/test';
import {
  meWithGuildPlan,
  routeGuildSettingsPage as routeCommon,
} from '../../src/test-support/billing-fixtures';

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-test' }),
  };
}

/** With G7_SHOTS set, saves the billing section's capture there; a plain run writes nothing. */
async function shot(page: import('@playwright/test').Page, name: string): Promise<void> {
  const dir = process.env.G7_SHOTS;
  if (dir)
    await page
      .getByTestId('guild-settings-billing')
      .screenshot({ path: `${dir}/${name}-${page.viewportSize()?.width}.png` });
}

test('not on the plan: shows the honest not-subscribed line and a real subscribe link', async ({ page }) => {
  await routeCommon(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithGuildPlan(null).data)));
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('not on the guild plan yet');
  await expect(page.getByTestId('guild-billing-price')).toHaveText(
    '$15 a month, cancel any time. Officers get the loot council helper and the raid readiness board, with more as they ship.',
  );
  await shot(page, 'g7-billing-not-on-plan');
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
          tier: 'guild',
          you_hold_seat: false,
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          scope: 'plan',
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
  await expect(page.getByTestId('guild-settings-billing')).toContainText('renews 1 Dec 2026');
  await shot(page, 'g7-billing-on-plan');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('Billed to you');
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
          tier: 'guild',
          you_hold_seat: false,
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          scope: 'plan',
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
          tier: 'guild',
          you_hold_seat: false,
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          scope: 'plan',
          billed_by: '',
          you_are_billing_contact: false,
        }).data,
      ),
    ),
  );
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await expect(page.getByTestId('guild-settings-billing')).toContainText('until 1 Dec 2026');
  await expect(page.getByTestId('guild-settings-billing')).not.toContainText('Billed by');
  await expect(page.getByTestId('guild-billing-contact')).toHaveText('Granted by Forever Sixty.');
  await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
});

test('Stripe unavailable: Manage billing says so and offers Try again, which asks again', async ({
  page,
}) => {
  await routeCommon(page);
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope(
        meWithGuildPlan({
          tier: 'guild',
          you_hold_seat: false,
          status: 'active',
          current_period_end: '2026-12-01T00:00:00Z',
          cancel_at_period_end: false,
          scope: 'plan',
          billed_by: 'you',
          you_are_billing_contact: true,
        }).data,
      ),
    ),
  );
  let asked = 0;
  await page.route('**/v1/billing/portal', (route) => {
    asked += 1;
    if (asked === 1) {
      return route.fulfill({
        status: 503,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: { code: 'billing_unavailable', message: 'billing is not available' },
          request_id: 'r',
        }),
      });
    }
    return route.fulfill(envelope({ portal_url: 'https://billing.stripe.com/guild' }));
  });
  await page.route('https://billing.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.goto('/guild/us/hardcore/the-last-watch/settings');
  await page.getByTestId('guild-manage-billing').click();
  await expect(page.getByTestId('guild-billing-unavailable')).toHaveText(
    'Billing is not available right now. Try again in a few minutes.',
  );
  await expect(page.getByTestId('guild-manage-billing')).toHaveCount(0);
  await page.getByTestId('guild-billing-retry').click();
  await page.waitForURL('https://billing.stripe.com/**');
  expect(asked).toBe(2);
});
