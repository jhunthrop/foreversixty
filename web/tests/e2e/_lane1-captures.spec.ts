// Throwaway capture spec for the premium lane 1 design review: not committed, deleted after
// the captures land in design/reviews/2026-10-24-premium-lane-1/.
import { test } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import { CHARACTER, CHARACTER_PATH, CHARACTER_URL, ownerMeAnswer, routeGearIcons } from './support/character-fixture';

const OUT = '../design/reviews/2026-10-24-premium-lane-1';
const WIDTHS = [390, 1440, 2000] as const;

function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'req-capture' }),
  };
}

const supporterEntitlements = mockEntitlements({ supporter_mark: true, server_sims: true });

async function shoot(page: import('@playwright/test').Page, name: string): Promise<void> {
  for (const width of WIDTHS) {
    await page.setViewportSize({ width, height: 1100 });
    await page.waitForTimeout(400);
    await page.screenshot({ path: `${OUT}/${name}-${width}.png`, fullPage: true });
  }
}

test('public character page with the Supporter chip', async ({ page }) => {
  await routeGearIcons(page);
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...CHARACTER, data: { ...CHARACTER.data, supporter: true } }),
    }),
  );
  await page.route('**/v1/me', (route) => route.fulfill(envelope(null, 401)));
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto(CHARACTER_PATH);
  await page.getByTestId('supporter-chip').first().waitFor();
  await shoot(page, 'character-public-supporter');
});

test('account page: hero chip and the Premium block', async ({ page }) => {
  const me = ownerMeAnswer();
  const data = me.data as Record<string, unknown>;
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ ...data, entitlements: supporterEntitlements })),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.goto('/account');
  await page.getByTestId('supporter-chip').first().waitFor();
  await shoot(page, 'account-premium-block');
});

test('home page: the shared character header with the chip', async ({ page }) => {
  await routeGearIcons(page);
  const me = ownerMeAnswer();
  const data = me.data as Record<string, unknown>;
  await page.route('**/v1/me', (route) =>
    route.fulfill(envelope({ ...data, entitlements: supporterEntitlements })),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await page.route(CHARACTER_URL, (route) =>
    route.fulfill({
      status: 200,
      contentType: 'application/json',
      body: JSON.stringify({ ...CHARACTER, data: { ...CHARACTER.data, supporter: true } }),
    }),
  );
  await page.goto('/');
  await page.getByTestId('supporter-chip').first().waitFor({ timeout: 15000 }).catch(() => undefined);
  await shoot(page, 'home-header-chip');
});

test('premium page after the copy change', async ({ page }) => {
  await page.route('**/v1/me', (route) => route.fulfill(envelope(null, 401)));
  await page.goto('/premium');
  await shoot(page, 'premium-page');
});
