// web/tests/e2e/checkout-launcher.spec.ts
import { expect, test } from './support/test';
import { accountPlanCopy } from '../../src/lib/billing/account-plan-copy';
import { billingBlockCopy } from '../../src/lib/billing/copy';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import { checkoutGuildCopy } from '../../src/lib/billing/checkout-guild-copy';
import { mockEntitlements } from '../../src/test-support/billing-fixtures';
import {
  OLYMPUS,
  ON_PLAN_YOU_PAY,
  PREMIUM_GRANT_OPEN_ENDED,
  PREMIUM_GRANT_UNTIL,
  PREMIUM_OWN_ENDING,
  PREMIUM_PAST_DUE,
  PREMIUM_SEAT_WITH_PORTAL,
  STEP_FIXTURES,
  buildMe,
} from '../../src/test-support/subscribe-fixtures';
import {
  CHECKOUT_PATH,
  GUILD_ID,
  GUILD_URL,
  checkoutReply,
  refusal,
  routeMeCell,
  routePost,
} from './support/checkout-guild';
import { expectNoAxeViolations } from './support/officer-tools';
import { expectTapHeight } from './support/subscribe-stubs';

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

test('guild already on the plan: the card shows the plan line and checkout is never posted', async ({
  page,
}) => {
  const posts = await routePost(page, CHECKOUT_PATH, () => checkoutReply());
  await routeMeCell(page, ON_PLAN_YOU_PAY);
  await page.goto(`${GUILD_URL}&guild_id=${GUILD_ID}`);
  await expect(page.getByTestId('guild-step-sentence')).toHaveText(
    checkoutGuildCopy.onPlan.youPay(OLYMPUS.name, 'guild'),
  );
  expect(posts).toHaveLength(0);
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

test('forbidden guild checkout: the step re-reads and shows its own mapped line, never the API sentence', async ({
  page,
}) => {
  const cell = await routeMeCell(page, STEP_FIXTURES.ready);
  const posts = await routePost(page, CHECKOUT_PATH, () => {
    cell.set(STEP_FIXTURES.frozen);
    return refusal(403, 'forbidden');
  });
  // A ready guild named by guild_id launches on arrival (spec 5.6); the refusal is the step's to map.
  await page.goto(`${GUILD_URL}&guild_id=${GUILD_ID}`);
  await expect(page.getByTestId('guild-step-sentence')).toHaveText(checkoutGuildCopy.frozen(OLYMPUS.name));
  await expect(page.locator('body')).not.toContainText('the api sentence, never shown');
  // Retrying a permission refusal does not help: no Try again.
  await expect(page.getByRole('button', { name: checkoutGuildCopy.checkoutErrors.tryAgain })).toHaveCount(0);
  expect(posts).toHaveLength(1);
});

test('an unescaped, malformed query string never renders raw into the page', async ({ page }) => {
  await page.route('**/v1/me', (route) =>
    route.fulfill(fulfil({ ok: false, data: null, error: null, request_id: 'r' }, 401)),
  );
  await page.goto('/premium/checkout?plan=%3Cscript%3Ealert(1)%3C%2Fscript%3E&interval=monthly');
  await expect(page.locator('script:has-text("alert(1)")')).toHaveCount(0);
  await expect(page.getByTestId('checkout-message')).toBeVisible();
});

// --- "You already have Premium." (subscribe-and-join spec §4.2, case 14) ------------------------

test.describe('plan=premium for an account that already has it', () => {
  test.use({ timezoneId: 'UTC' });

  const ending = PREMIUM_OWN_ENDING.entitlements?.billing;
  if (ending === null || ending === undefined) throw new Error('PREMIUM_OWN_ENDING carries no billing row');
  const OWN_RENEWING = buildMe({
    billing: { ...ending, cancel_at_period_end: false },
    personalPortal: true,
  });
  const periodEnd = ending.current_period_end ?? '';
  const seatGuild = PREMIUM_SEAT_WITH_PORTAL.seat?.guild_name ?? '';
  const seatGiver = PREMIUM_SEAT_WITH_PORTAL.seat?.given_by ?? '';
  const grantUntilEnd = PREMIUM_GRANT_UNTIL.entitlements?.billing?.current_period_end ?? '';
  const renews = accountPlanCopy.planLine.renews(periodEnd, 600, 'monthly', 'UTC');
  const MANAGE = { kind: 'button', label: premiumSourceCopy.buttons.manage } as const;
  const SEE_PLAN = { kind: 'link', label: premiumSourceCopy.buttons.seePlan } as const;

  const CASES = [
    { name: 'own, renewing', me: OWN_RENEWING, line: renews, line2: null, action: MANAGE },
    {
      name: 'own, ending',
      me: PREMIUM_OWN_ENDING,
      line: accountPlanCopy.planLine.ends(periodEnd, 'UTC'),
      line2: null,
      action: MANAGE,
    },
    {
      name: 'past due',
      me: PREMIUM_PAST_DUE,
      line: renews,
      line2: billingBlockCopy.pastDueBanner,
      action: MANAGE,
    },
    {
      name: 'a guild seat',
      me: PREMIUM_SEAT_WITH_PORTAL,
      line: accountPlanCopy.planLine.seat(seatGuild),
      line2: accountPlanCopy.planLine.givenBy(seatGiver),
      action: SEE_PLAN,
    },
    {
      name: 'a grant with no end',
      me: PREMIUM_GRANT_OPEN_ENDED,
      line: premiumSourceCopy.grant.noEnd,
      line2: premiumSourceCopy.grant.givenNoEnd,
      action: SEE_PLAN,
    },
    {
      name: 'a grant until a date',
      me: PREMIUM_GRANT_UNTIL,
      line: premiumSourceCopy.grant.until(grantUntilEnd, 'UTC'),
      line2: premiumSourceCopy.grant.givenUntil,
      action: SEE_PLAN,
    },
  ] as const;

  for (const { name, me, line, line2, action } of CASES) {
    test(`${name}: the panel reads the source and never calls checkout`, async ({ page }) => {
      await routeMeCell(page, me);
      const posts = await routePost(page, CHECKOUT_PATH, () => refusal(500, 'internal'));
      await page.goto('/premium/checkout?plan=premium&interval=monthly');
      const panel = page.getByTestId('premium-owned');
      await expect(panel).toBeVisible();
      await expect(panel.getByRole('heading', { name: premiumSourceCopy.panelTitle })).toBeFocused();
      await expect(page.getByTestId('premium-owned-line')).toContainText(line);
      if (line2 === billingBlockCopy.pastDueBanner) {
        await expect(page.getByTestId('premium-owned-past-due')).toHaveText(line2);
      } else if (line2 === null) {
        await expect(page.getByTestId('premium-owned-line2')).toHaveCount(0);
      } else {
        await expect(page.getByTestId('premium-owned-line2')).toHaveText(line2);
      }
      const button = page.getByTestId('premium-owned-action');
      await expect(button).toHaveText(action.label);
      if (action.kind === 'link') await expect(button).toHaveAttribute('href', '/account#plan');
      await expectTapHeight(button);
      await expectNoAxeViolations(page, '[data-testid="premium-owned"]');
      expect(posts).toHaveLength(0);
    });
  }
});
