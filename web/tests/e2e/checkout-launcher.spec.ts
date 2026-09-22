// web/tests/e2e/checkout-launcher.spec.ts
import { expect, test } from '@playwright/test';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';

function fulfil(body: unknown, status = 200) {
  return { status, contentType: 'application/json', body: JSON.stringify(body) };
}

test('signed out: leads to sign-in and returns to the same checkout url', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/premium/checkout?plan=premium&interval=monthly');
  await expect(page).toHaveURL(/\/login\?next=%2Fpremium%2Fcheckout/);
});

test('signed in: posts checkout and redirects to the returned checkout_url', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
          entitlements: mockEntitlements(),
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  let posted: unknown;
  await page.route('**/v1/billing/checkout', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill(
      fulfil({
        ok: true,
        data: { checkout_url: 'https://checkout.stripe.com/x' },
        error: null,
        request_id: 'r',
      }),
    );
  });
  // The real redirect leaves the page; stop navigation so the assertion can still run.
  await page.route('https://checkout.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.goto('/premium/checkout?plan=premium&interval=monthly');
  await page.waitForURL('https://checkout.stripe.com/**');
  expect(posted).toEqual({ plan: 'premium', interval: 'monthly' });
});

test('billing not configured yet: shows an honest message, no dead redirect', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
          entitlements: mockEntitlements(),
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/billing/checkout', (route) =>
    route.fulfill(
      fulfil({ ok: false, data: null, error: { message: 'billing_unavailable' }, request_id: 'r' }, 503),
    ),
  );
  await page.goto('/premium/checkout?plan=premium&interval=monthly');
  await expect(page.getByTestId('checkout-message')).toContainText('not open yet');
  await expect(page).toHaveURL(/\/premium\/checkout/);
});

test('guild already on the plan: shows the conflict message', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
          entitlements: mockEntitlements(),
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/billing/checkout', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 409)),
  );
  await page.goto('/premium/checkout?plan=guild&interval=monthly&guild_id=42');
  await expect(page.getByTestId('checkout-message')).toContainText('already on the plan');
});

test('the success return re-fetches /v1/me rather than trusting the query string', async ({ page }) => {
  let meCalls = 0;
  await page.route('**/v1/me', (route) => {
    meCalls += 1;
    return route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
          entitlements: mockEntitlements({ server_sims: meCalls > 1 }),
        },
        error: null,
        request_id: 'r',
      }),
    );
  });
  await page.goto('/premium/checkout?plan=premium&interval=monthly&status=success&session_id=cs_test_1');
  await expect(page.getByTestId('checkout-message')).toContainText("You're all set", { timeout: 10000 });
  expect(meCalls).toBeGreaterThan(1);
});

test("forbidden guild checkout: shows the API's own reason verbatim, honest for either refusal", async ({
  page,
}) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      fulfil({
        ok: true,
        data: {
          user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
          characters: [],
          guilds: [],
          entitlements: mockEntitlements(),
        },
        error: null,
        request_id: 'r',
      }),
    ),
  );
  await page.route('**/v1/billing/checkout', (route) =>
    route.fulfill(
      fulfil(
        {
          ok: false,
          data: null,
          error: {
            message:
              "this guild's claim is contested; billing actions are frozen for the disputed claimant until a moderator resolves it",
          },
          request_id: 'r',
        },
        403,
      ),
    ),
  );
  await page.goto('/premium/checkout?plan=guild&interval=monthly&guild_id=42');
  await expect(page.getByTestId('checkout-message')).toContainText('billing actions are frozen');
  // No retry button: retrying a permission refusal does not help either way.
  await expect(page.getByRole('button', { name: 'Try again' })).toHaveCount(0);
});

test('an unescaped, malformed query string never renders raw into the page', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/premium/checkout?plan=%3Cscript%3Ealert(1)%3C%2Fscript%3E&interval=monthly');
  await expect(page.locator('script:has-text("alert(1)")')).toHaveCount(0);
  await expect(page.getByTestId('checkout-message')).toBeVisible();
});
