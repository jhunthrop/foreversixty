// web/tests/e2e/account-seat.spec.ts
// Launch pricing spec section 9 cases 17 and 18: the account Plan panel for a seat offer, a seat
// holder, a founder and an ended seat. The viewer's zone is UTC so stored instants read as written.
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import {
  FOUNDER_PREMIUM_FACTS,
  ME_ENDED_SEAT,
  ME_SEAT,
  ME_SEAT_OFFER,
  mockEntitlements,
} from '../../src/test-support/billing-fixtures';
import { accountPlanCopy } from '../../src/lib/billing/account-plan-copy';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import { PREMIUM_SEAT_NO_PORTAL, PREMIUM_SEAT_WITH_PORTAL } from '../../src/test-support/subscribe-fixtures';
import { expectNoAxeViolations } from './support/officer-tools';
import { expectTapHeight } from './support/subscribe-stubs';

test.use({ timezoneId: 'UTC' });

const OWN_PREMIUM = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-20T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 600,
};

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
  };
}

async function openAccount(
  page: Page,
  extra: Record<string, unknown>,
  billing: unknown = null,
): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ billing: billing as never }),
        notifications: { discord: null },
        ...extra,
      }),
    ),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto('/account');
}

const expectClean = (page: Page): Promise<void> => expectNoAxeViolations(page, '#plan');

test('a payer sees the offer, accepts, and reads the new status', async ({ page }) => {
  await openAccount(page, { seat_offer: ME_SEAT_OFFER }, OWN_PREMIUM);
  const offer = page.getByTestId('account-seat-offer');
  await expect(offer).toContainText('The Last Watch offers you a Premium seat.');
  await expect(offer).toContainText('ends 20 Dec 2026, and the seat starts then');
  await expect(offer).toContainText('$6 a month');
  await expectClean(page);

  let hit = '';
  await page.route('**/v1/guilds/501/seats/offer/accept', (route) => {
    hit = route.request().url();
    return route.fulfill({ status: 204, body: '' });
  });
  // After the accept the API sends the seat and no offer.
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ billing: OWN_PREMIUM as never }),
        notifications: { discord: null },
        seat: ME_SEAT,
      }),
    ),
  );
  await page.getByTestId('account-seat-offer-accept').click();
  await expect(page.getByTestId('account-seat-offer-status')).toHaveText(
    'Accepted. Your Premium ends 20 Dec 2026; the seat starts then.',
  );
  expect(hit).toContain('/seats/offer/accept');
  await expect(page.getByTestId('account-seat-offer-accept')).toHaveCount(0);
  await expect(page.getByTestId('account-billing-row')).toContainText(
    'Premium · ends 20 Dec 2026 · then a seat from The Last Watch',
  );
});

test('a refused accept shows the error line and keeps the offer open', async ({ page }) => {
  await openAccount(page, { seat_offer: ME_SEAT_OFFER }, OWN_PREMIUM);
  await page.route('**/v1/guilds/501/seats/offer/accept', (route) => route.fulfill(envelope(null, 500)));
  await page.getByTestId('account-seat-offer-accept').click();
  await expect(page.getByTestId('account-seat-offer-error')).toHaveText('That did not save. Try again.');
  await expect(page.getByTestId('account-seat-offer-accept')).toBeEnabled();
});

test('declining gives its status', async ({ page }) => {
  await openAccount(page, { seat_offer: ME_SEAT_OFFER }, OWN_PREMIUM);
  await page.route('**/v1/guilds/501/seats/offer/decline', (route) =>
    route.fulfill({ status: 204, body: '' }),
  );
  await page.getByTestId('account-seat-offer-decline').click();
  await expect(page.getByTestId('account-seat-offer-status')).toHaveText(
    'Declined. Your Premium carries on as before.',
  );
});

test('own Premium ending into a seat reads the hand-over', async ({ page }) => {
  await openAccount(page, { seat: ME_SEAT }, OWN_PREMIUM);
  await expect(page.getByTestId('account-billing-row')).toContainText(
    'Premium · ends 20 Dec 2026 · then a seat from The Last Watch',
  );
  await expect(page.getByRole('button', { name: 'Manage billing' })).toBeVisible();
});

test('a seat holder sees a seat from the guild and no Manage billing', async ({ page }) => {
  await openAccount(page, { seat: { ...ME_SEAT, starts_at: null } });
  await expect(page.getByTestId('account-billing-row')).toContainText('Premium · a seat from The Last Watch');
  await expect(page.getByTestId('account-plan-line2')).toContainText('Given by Thornwick');
  await expect(page.getByRole('button', { name: 'Manage billing' })).toHaveCount(0);
  await expectClean(page);
});

test('a founder sees the Founder mark with its title', async ({ page }) => {
  await openAccount(page, {}, { ...OWN_PREMIUM, ...FOUNDER_PREMIUM_FACTS });
  const mark = page.getByTestId('account-founder');
  await expect(mark).toContainText('Founder');
  await expect(mark).toHaveAttribute('title', 'Founder price, kept while your plan stays active.');
  await expectClean(page);
});

test('the ended seat line links to its resubscribe URL', async ({ page }) => {
  await openAccount(page, { ended_seat: ME_ENDED_SEAT });
  const line = page.getByTestId('account-seat-ended');
  await expect(line).toContainText('Your seat from The Last Watch ended 18 Nov.');
  await expect(line.getByRole('link', { name: 'Come back at $6 a month' })).toHaveAttribute(
    'href',
    ME_ENDED_SEAT.resubscribe_url,
  );
  await expectClean(page);
});

test('a seat holder with a personal portal sees Billing history, not Manage billing', async ({ page }) => {
  await openAccount(page, {
    seat: PREMIUM_SEAT_WITH_PORTAL.seat,
    entitlements: { ...PREMIUM_SEAT_WITH_PORTAL.entitlements, server_sims: true },
  });
  await expect(page.getByTestId('account-billing-row')).toContainText(
    accountPlanCopy.planLine.seat(PREMIUM_SEAT_WITH_PORTAL.seat?.guild_name ?? ''),
  );
  const button = page.getByTestId('account-billing-button');
  await expect(button).toHaveText(premiumSourceCopy.buttons.history);
  await expectTapHeight(button);
  await expectClean(page);
});

test('a seat holder without a personal portal sees no billing button', async ({ page }) => {
  await openAccount(page, {
    seat: PREMIUM_SEAT_NO_PORTAL.seat,
    entitlements: { ...PREMIUM_SEAT_NO_PORTAL.entitlements, server_sims: true },
  });
  await expect(page.getByTestId('account-billing-row')).toBeVisible();
  await expect(page.getByTestId('account-billing-button')).toHaveCount(0);
});

test('the offer controls are at least 44px tall at 360 wide', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 800 });
  await openAccount(page, { seat_offer: ME_SEAT_OFFER }, OWN_PREMIUM);
  for (const id of ['account-seat-offer-accept', 'account-seat-offer-decline']) {
    const box = await page.getByTestId(id).boundingBox();
    expect(box?.height ?? 0).toBeGreaterThanOrEqual(44);
  }
});
