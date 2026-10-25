// web/tests/e2e/account-billing.spec.ts
import { expect, test } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { ownerMeAnswer } from './support/character-fixture';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

const HERO_CHARACTER = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'Priest',
  level: 60,
};

function meWith(
  entitlements: ReturnType<typeof mockEntitlements>,
  characters: unknown[] = [],
  serverSims?: { runs_this_month: number },
) {
  return {
    ok: true,
    data: {
      user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
      characters,
      guilds: [],
      entitlements,
      server_sims: serverSims,
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

test('the Premium block names each benefit and shows the Supporter mark when it is on', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil(
        meWith(
          mockEntitlements({ server_sims: true, supporter_mark: true, history: true }),
          [HERO_CHARACTER],
          {
            runs_this_month: 4,
          },
        ),
      ),
    ),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.goto('/account');
  const block = page.getByTestId('account-premium-block');
  await expect(page.getByTestId('account-premium-serverSims')).toHaveText(
    /Server sims\s*On · 4 runs this month/,
  );
  await expect(page.getByTestId('account-premium-serverSims-link')).toHaveAttribute('href', '/sim');
  await expect(page.getByTestId('account-premium-supporter-link')).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale',
  );
  await expect(page.getByTestId('account-premium-history-link')).toHaveText('Open on your main character');
  await expect(page.getByTestId('account-premium-history-link')).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale#history',
  );
  await expect(block).toContainText('Discord notifications');
  await expect(page.getByTestId('account-hero').getByTestId('supporter-chip')).toBeVisible();
  // The mark belongs to the person: the character rows never carry it.
  await expect(page.getByTestId('supporter-chip')).toHaveCount(1);
});

test('without the Supporter mark the hero carries no chip', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil(meWith(mockEntitlements({ supporter_mark: false }), [HERO_CHARACTER]))),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.goto('/account');
  await expect(page.getByTestId('account-hero')).toContainText('Elyra Duskvale');
  await expect(page.getByTestId('supporter-chip')).toHaveCount(0);
});

test('the home header keeps the whole name beside the Supporter chip', async ({ page }) => {
  const me = ownerMeAnswer();
  const data = me.data as Record<string, unknown>;
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({ ...me, data: { ...data, entitlements: mockEntitlements({ supporter_mark: true }) } }),
    ),
  );
  await page.route('**/v1/devices', (route) =>
    route.fulfill(fulfil({ ok: true, data: [], error: null, request_id: 'r' })),
  );
  await page.setViewportSize({ width: 1440, height: 1000 });
  await page.goto('/');
  const name = page.getByTestId('character-header-name');
  await expect(name).toHaveText('Elyra Duskvale');
  // The name box is as wide as it is without the chip (about 114px in this column): never squeezed to a letter.
  expect((await name.boundingBox())!.width).toBeGreaterThanOrEqual(100);
  await expect(page.getByTestId('character-header').getByTestId('supporter-chip')).toBeVisible();
});
