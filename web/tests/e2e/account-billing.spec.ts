// web/tests/e2e/account-billing.spec.ts
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { mockEntitlements, NO_PRICE_FACTS } from '../../src/test-support/billing-fixtures';
import { ownerMeAnswer } from './support/character-fixture';
import type { Me } from '../../src/lib/account/me-types';
import { accountPlanCopy } from '../../src/lib/billing/account-plan-copy';
import { fittingGuildPlan } from '../../src/lib/billing/checkout-params';
import { billingBlockCopy } from '../../src/lib/billing/copy';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import {
  MILLENNIALS,
  OLYMPUS,
  ON_PLAN_YOU_PAY,
  OWNER_MEMBER,
  PREMIUM_FALLBACK,
  PREMIUM_GRANT_OPEN_ENDED,
  PREMIUM_GRANT_UNTIL,
  PREMIUM_OWN_ENDING,
  PREMIUM_OWN_FOUNDER,
  PREMIUM_OWN_THEN_SEAT,
  PREMIUM_PAST_DUE,
  PREMIUM_SEAT_NO_PORTAL,
  PREMIUM_SEAT_WITH_PORTAL,
} from '../../src/test-support/subscribe-fixtures';
import { expectNoAxeViolations } from './support/officer-tools';
import { expectTapHeight, routeDevices, routeMeCell } from './support/subscribe-stubs';

// A Stripe date reads on the viewer's calendar (launch pricing spec §6, lib/dates.ts zonedDayMonthYear);
// the fixtures' instants are midnight UTC, so the viewer is in UTC for them to read as written.
test.use({ timezoneId: 'UTC' });

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
      // The deployment has notifications on, so the Premium block carries its Discord row.
      notifications: { discord: null },
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
  await expect(page.getByTestId('account-see-plans')).toHaveAttribute('href', '/premium#premium');
});

test('active subscription: shows the renewal date and a working manage-billing button', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil(
        meWith(
          mockEntitlements({
            server_sims: true,
            billing: {
              ...NO_PRICE_FACTS,
              // A source-less, priceless row reads as a grant (plan-line isGrant); this is Stripe.
              source: 'stripe',
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
  await expect(page.getByTestId('account')).toContainText('renews 1 Dec 2026');
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
              ...NO_PRICE_FACTS,
              // A source-less, priceless row reads as a grant (plan-line isGrant); this is Stripe.
              source: 'stripe',
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
  await expect(page.getByTestId('account-premium-serverSims')).toHaveText(
    /Server sims\s*On · 4 runs this month/,
  );
  await expect(page.getByTestId('account-premium-serverSims-link')).toHaveAttribute('href', '/sim');
  await expect(page.getByTestId('account-premium-supporter-link')).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale',
  );
  await expect(page.getByTestId('account-premium-history-link')).toContainText('Open on your main character');
  await expect(page.getByTestId('account-premium-history-link')).toHaveAttribute(
    'href',
    '/character/us/hardcore/elyra-duskvale#history',
  );
  await expect(page.getByTestId('account-premium-notifications')).toContainText(
    // Not entitled to notifications: the free wording (launch pricing spec §6 amendment).
    /Discord notifications\s*Premium\s*only/,
  );
  await expect(page.getByTestId('account-premium-notifications-link')).toHaveCount(0);
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

// --- the Plan panel by source, and the guild rows (subscribe-and-join spec §4.3, §6.5; cases 15) ---

/** The same account with its Premium rows On, as the API resolves them for a paying or given account. */
function withPremiumRowsOn(me: Me): Me {
  return {
    ...me,
    entitlements: { ...mockEntitlements(), ...me.entitlements, server_sims: true, history: true },
  };
}

async function openFixture(page: Page, me: Me): Promise<void> {
  await routeMeCell(page, () => me);
  await routeDevices(page);
  await page.goto('/account');
  await expect(page.getByTestId('account-plan-headline')).toBeVisible();
}

test("the owner's grant reads granted with no end date and offers no billing button", async ({ page }) => {
  await openFixture(page, withPremiumRowsOn(PREMIUM_GRANT_OPEN_ENDED));
  await expect(page.getByTestId('account-billing-row')).toContainText(premiumSourceCopy.grant.noEnd);
  await expect(page.getByTestId('account-plan-line2')).toHaveText(premiumSourceCopy.grant.givenNoEnd);
  await expect(page.getByTestId('account-billing-button')).toHaveCount(0);
  await expectNoAxeViolations(page, '#plan');
});

test('every fixture whose Premium rows are On reads a plan line, never Not on Premium', async ({ page }) => {
  const sources = [
    PREMIUM_OWN_FOUNDER,
    PREMIUM_OWN_ENDING,
    PREMIUM_PAST_DUE,
    PREMIUM_OWN_THEN_SEAT,
    PREMIUM_SEAT_WITH_PORTAL,
    PREMIUM_SEAT_NO_PORTAL,
    PREMIUM_GRANT_OPEN_ENDED,
    PREMIUM_GRANT_UNTIL,
    PREMIUM_FALLBACK,
  ];
  let current: Me = withPremiumRowsOn(sources[0]!);
  await routeMeCell(page, () => current);
  await routeDevices(page);
  for (const source of sources) {
    current = withPremiumRowsOn(source);
    await page.goto('/account');
    await expect(page.getByTestId('account-premium-serverSims-link')).toBeVisible();
    await expect(page.getByTestId('account-billing-row')).not.toContainText(billingBlockCopy.notSubscribed);
  }
});

test('an own subscription gets Manage billing, tall enough to tap', async ({ page }) => {
  await openFixture(page, withPremiumRowsOn(PREMIUM_OWN_ENDING));
  const button = page.getByTestId('account-billing-button');
  await expect(button).toHaveText(premiumSourceCopy.buttons.manage);
  await expectTapHeight(button);
});

test('the guild rows say Not verified, that the guild has no plan, and link the fitting next step', async ({
  page,
}) => {
  await openFixture(page, OWNER_MEMBER);
  const rows = page.getByTestId('account-guild-row');
  await expect(rows).toHaveCount(2);
  for (const guild of [OLYMPUS, MILLENNIALS]) {
    const row = rows.filter({ hasText: guild.name });
    await expect(row.getByTestId('account-guild-not-verified')).toHaveText(
      accountPlanCopy.guildRow.notVerified,
    );
    await expect(row.getByTestId('account-guild-no-plan')).toContainText(
      accountPlanCopy.guildRow.noPlan(guild.name),
    );
    await expect(row.getByTestId('account-guild-plan-tag')).toHaveText(accountPlanCopy.guildRow.planTag.none);
    const next = row.getByRole('link', { name: new RegExp(accountPlanCopy.guildRow.nextStep) });
    await expect(next).toContainText(accountPlanCopy.guildRow.nextStep);
    await expect(next).toHaveAttribute(
      'href',
      `/premium/checkout?plan=${fittingGuildPlan(0)}&interval=yearly#guild-${guild.id}`,
    );
  }
  await expectNoAxeViolations(page, '[data-testid="account-guilds-panel"]');
});

test('a guild on a plan, verified, carries neither the mark nor the no-plan line', async ({ page }) => {
  await openFixture(page, ON_PLAN_YOU_PAY);
  const row = page.getByTestId('account-guild-row');
  await expect(row).toHaveCount(1);
  await expect(row.getByTestId('account-guild-not-verified')).toHaveCount(0);
  await expect(row.getByTestId('account-guild-no-plan')).toHaveCount(0);
});
