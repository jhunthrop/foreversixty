// web/tests/e2e/checkout-guild.spec.ts
// Launch pricing spec section 9, case 8 (the chooser) and case 18 (axe, 44px controls).
import { expect, test, type Page } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { expectNoAxeViolations } from './support/officer-tools';

const CHOOSER_URL = '/premium/checkout?plan=guild&interval=yearly';

function guild(id: number, name: string, verifiedCount: number, can: string[]) {
  return {
    id,
    region: 'us',
    ruleset: 'forever',
    name,
    verified: true,
    verified_count: verifiedCount,
    can,
    plan: null,
  };
}

async function stubMe(page: Page, guilds: unknown[]): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds,
          entitlements: mockEntitlements(),
        },
        error: null,
        request_id: 'r',
      }),
    }),
  );
}

test('chooser lists billable guilds and names the fitting plan', async ({ page }) => {
  await stubMe(page, [
    guild(1, 'Olympus', 12, ['billing']),
    guild(2, 'Titans', 28, ['billing']),
    guild(3, 'Other', 9, []),
  ]);
  await page.goto(CHOOSER_URL);
  const rows = page.getByTestId('checkout-guild-row');
  await expect(rows).toHaveCount(2);
  await expect(page.getByText('Raid guild fits 28 members')).toBeVisible();
  expect((await rows.first().boundingBox())?.height).toBeCloseTo(56, 0);
});

test('choosing a row posts checkout with that guild id and the url plan', async ({ page }) => {
  await stubMe(page, [guild(2, 'Titans', 28, ['billing'])]);
  let posted: unknown;
  await page.route('**/v1/billing/checkout', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: true,
        data: { checkout_url: 'https://checkout.stripe.com/x' },
        error: null,
        request_id: 'r',
      }),
    });
  });
  await page.route('https://checkout.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.goto('/premium/checkout?plan=raid_guild&interval=monthly');
  await page.getByTestId('checkout-guild-row').click();
  await page.waitForURL('https://checkout.stripe.com/**');
  expect(posted).toEqual({ plan: 'raid_guild', interval: 'monthly', guild_id: 2 });
});

test('no billable guild shows the none line and the Find your guild link', async ({ page }) => {
  await stubMe(page, [guild(3, 'Other', 9, [])]);
  await page.goto(CHOOSER_URL);
  await expect(page.getByText('You can subscribe a guild you lead or have claimed.')).toBeVisible();
  await expect(page.getByRole('link', { name: 'Find your guild' })).toHaveAttribute('href', '/guild');
});

test('the chooser is axe-clean and its controls are 44px at 360', async ({ page }) => {
  await page.setViewportSize({ width: 360, height: 740 });
  await stubMe(page, [guild(1, 'Olympus', 12, ['billing'])]);
  await page.goto(CHOOSER_URL);
  await expect(page.getByTestId('checkout-guild-row')).toBeVisible();
  const box = await page.getByTestId('checkout-guild-row').boundingBox();
  expect(box?.height).toBeGreaterThanOrEqual(44);
  await expectNoAxeViolations(page, 'main');
});

test('a comeback checkout posts comeback_id; a 404 shows the no-retry line', async ({ page }) => {
  await stubMe(page, []);
  let posted: unknown;
  await page.route('**/v1/billing/checkout', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill({
      status: 404,
      contentType: 'application/json',
      body: JSON.stringify({
        ok: false,
        data: null,
        error: { code: 'not_found', message: 'gone' },
        request_id: 'r',
      }),
    });
  });
  await page.goto('/premium/checkout?plan=premium&interval=yearly&comeback=12');
  await expect(page.getByTestId('checkout-message')).toContainText('no longer available');
  await expect(page.getByRole('button', { name: 'Try again' })).toHaveCount(0);
  expect(posted).toEqual({ plan: 'premium', interval: 'yearly', comeback_id: 12 });
});
