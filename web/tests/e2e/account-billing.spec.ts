// web/tests/e2e/account-billing.spec.ts
import { expect, test } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

function meWith(entitlements: ReturnType<typeof mockEntitlements>) {
  return {
    ok: true,
    data: {
      user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
      characters: [],
      guilds: [],
      entitlements,
    },
    error: null,
    request_id: 'r',
  };
}

test('not subscribed: plain text with a link to the plans, no upsell tone', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(fulfil(meWith(mockEntitlements()))));
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.goto('/account');
  await expect(page.getByTestId('account')).toContainText('Not on Premium');
  await expect(page.getByRole('link', { name: 'See plans' })).toHaveAttribute('href', '/premium');
});

test('active subscription: shows the renewal date and a working manage-billing button', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil(
        meWith(
          mockEntitlements({
            server_sims: true,
            billing: {
              plan: 'premium',
              status: 'active',
              current_period_end: '2026-12-01T00:00:00Z',
              cancel_at_period_end: false,
            },
          }),
        ),
      ),
    ),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  let posted: unknown;
  await page.route('**/v1/billing/portal', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill(
      fulfil({
        ok: true,
        data: { portal_url: 'https://billing.stripe.com/p' },
        error: null,
        request_id: 'r',
      }),
    );
  });
  await page.route('https://billing.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.goto('/account');
  await expect(page.getByTestId('account')).toContainText('Renews');
  await page.getByRole('button', { name: 'Manage billing' }).click();
  await page.waitForURL('https://billing.stripe.com/**');
  expect(posted).toEqual({});
});

test('past due: shows the payment-failed banner without losing access', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil(
        meWith(
          mockEntitlements({
            server_sims: true,
            billing: {
              plan: 'premium',
              status: 'past_due',
              current_period_end: '2026-12-01T00:00:00Z',
              cancel_at_period_end: false,
            },
          }),
        ),
      ),
    ),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.goto('/account');
  await expect(page.getByTestId('account')).toContainText('payment failed');
});
