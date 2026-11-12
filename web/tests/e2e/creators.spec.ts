// web/tests/e2e/creators.spec.ts
// The creator program's pages (design/specs/2026-11-08-creator-program.md): the owner's creators,
// requests and statements (approve asks once, then pays), the creator's onboarding and stats, the
// public page and the application with its states. The API is the fixture JSON served by route.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import type { Page, Route } from '@playwright/test';
import { expect, test } from './support/test';

const FIX = path.join(import.meta.dirname, 'fixtures/creators');
const read = (name: string): unknown => JSON.parse(readFileSync(path.join(FIX, name), 'utf8'));

const envelope = (data: unknown, status = 200, code = 'x', fields: Record<string, string> | null = null) => ({
  status,
  contentType: 'application/json',
  body: JSON.stringify({
    ok: status < 400,
    data: status < 400 ? data : null,
    error: status < 400 ? null : { code, message: 'refused', fields },
    request_id: 'r',
  }),
});

const me = (role: string, extra: Record<string, unknown> = {}) => ({
  user: { id: 78, battletag: 'Thandrel#1894', email: null, role, anonymize: false },
  characters: [],
  guilds: [],
  ...extra,
});

async function signedIn(page: Page, role = 'user', extra: Record<string, unknown> = {}): Promise<void> {
  await page.context().addCookies([{ name: 'fs_csrf', value: 'tok', domain: 'localhost', path: '/' }]);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(me(role, extra))));
}

async function asOwner(page: Page, requests = read('requests.json')): Promise<void> {
  await signedIn(page, 'admin');
  await page.route('**/v1/admin/creators', (route) =>
    route.request().method() === 'GET'
      ? route.fulfill(
          envelope({ creators: read('creators.json'), stripe_enabled: true, payouts_enabled: true }),
        )
      : route.fallback(),
  );
  await page.route('**/v1/admin/creators/requests', (route) => route.fulfill(envelope({ requests })));
}

test.describe('/admin/creators', () => {
  test('lists requests above the creators, with what each figure covers', async ({ page }) => {
    await asOwner(page);
    await page.goto('/admin/creators');
    await expect(page.getByTestId('request-7')).toContainText('twitch.tv/emberfall');
    await expect(page.getByTestId('request-7')).toContainText('Twitch · about 1,200 · 8 Dec 2026');
    const grim = page.getByTestId('creator-row-grimbrew');
    await expect(grim).toContainText('/go/grimbrew · code GRIMBREW · since 8 Nov 2026');
    await expect(grim).toContainText('$184.20');
    await expect(grim).toContainText('December 2026 · of $500 cap');
    await expect(grim).toContainText('all time');
    await expect(page.getByTestId('creator-status-grimbrew')).toHaveText('Payouts ready');
    await expect(page.getByTestId('creator-row-lunawick')).toContainText('December 2026 · $50 cap reached');
  });

  test('approve prefills the form from the request and makes the creator', async ({ page }) => {
    await asOwner(page);
    let sent: Record<string, unknown> | null = null;
    await page.route('**/v1/admin/creators/requests/7/approve', async (route: Route) => {
      sent = route.request().postDataJSON() as Record<string, unknown>;
      await route.fulfill(envelope({ ...(read('creators.json') as object[])[0], name: 'emberfall' }, 201));
    });
    await page.goto('/admin/creators');
    await page.getByTestId('request-approve-7').click();
    await expect(page.getByTestId('creators-form-card')).toContainText('Approve twitch.tv/emberfall');
    await expect(page.getByTestId('creators-slug')).toHaveValue('emberfall');
    await expect(page.getByTestId('creators-code')).toHaveValue('EMBERFALL');
    await page.getByTestId('creators-save').click();
    await expect(page.getByTestId('creators-notice')).toContainText('emberfall is in.');
    expect(sent).toMatchObject({ slug: 'emberfall', code: 'EMBERFALL', cap_cents: 50000 });
  });

  test('edit fixes the link and code, and copies a one-time invite', async ({ page, context }) => {
    await context.grantPermissions(['clipboard-read', 'clipboard-write']);
    await asOwner(page, []);
    await page.route('**/v1/admin/creators/1/invite', (route) =>
      route.fulfill(envelope({ url: 'https://foreversixty.gg/creators/onboarding?invite=tok' })),
    );
    await page.goto('/admin/creators');
    await expect(page.getByTestId('creators-requests')).toHaveCount(0);
    await page.getByTestId('creator-edit-grimbrew').click();
    await expect(page.getByTestId('creators-fixed')).toHaveText('/go/grimbrew · code GRIMBREW');
    await expect(page.getByTestId('creators-slug')).toHaveCount(0);
    await page.getByTestId('creators-invite').click();
    await expect(page.getByTestId('creators-notice')).toHaveText(
      'Invite copied. It works once, for 14 days.',
    );
    expect(await page.evaluate(() => navigator.clipboard.readText())).toContain(
      '/creators/onboarding?invite=',
    );
  });

  test('a field the API refuses shows under its field', async ({ page }) => {
    await asOwner(page, []);
    await page.route('**/v1/admin/creators', (route) =>
      route.request().method() === 'POST'
        ? route.fulfill(envelope(null, 400, 'invalid', { code: 'capitals and digits, 4 to 20' }))
        : route.fallback(),
    );
    await page.goto('/admin/creators');
    await page.getByTestId('creators-save').click();
    await expect(page.getByTestId('creators-form-card')).toContainText('capitals and digits, 4 to 20');
  });

  test('anyone but the owner sees one line', async ({ page }) => {
    await signedIn(page);
    await page.goto('/admin/creators');
    await expect(page.getByTestId('creators-not-owner')).toBeVisible();
    await expect(page.locator('[data-testid^="creator-row-"]')).toHaveCount(0);
  });
});

test.describe('/admin/creators/payouts', () => {
  async function statements(page: Page, enabled = true, availableCents = 50000): Promise<void> {
    await signedIn(page, 'admin');
    await page.route('**/v1/admin/creators/statements**', (route) =>
      route.fulfill(envelope({ ...(read('statements.json') as object), payouts_enabled: enabled })),
    );
    await page.route('**/v1/admin/creators/balance', (route) =>
      route.fulfill(envelope({ available_cents: availableCents })),
    );
  }

  test('approve asks once, names the month, then shows the transfer', async ({ page }) => {
    await statements(page);
    let calls = 0;
    await page.route('**/v1/admin/creators/1/statements/2026-11/approve', async (route) => {
      calls++;
      const paid = { ...(read('statements.json') as { statements: object[] }).statements[0] };
      await route.fulfill(
        envelope({
          ...paid,
          status: 'paid',
          paid_at: '2026-12-09T12:00:00Z',
          transfer_id: 'tr_creator-1-2026-11',
        }),
      );
    });
    await page.goto('/admin/creators/payouts');
    await expect(page.getByTestId('payout-row-1')).toContainText('Ready to approve');
    await expect(page.getByTestId('payout-blocked-2')).toHaveText('Stripe not ready');
    await expect(page.getByTestId('payout-approve-3')).toHaveCount(0);
    await page.getByTestId('payout-approve-1').click();
    await expect(page.getByTestId('pay-dialog')).toContainText('Pay Grimbrew $168.60 for November 2026?');
    await expect(page.getByTestId('pay-balance')).toHaveText(
      'Available in Stripe: $500.00. This payout: $168.60.',
    );
    await page.getByTestId('pay-cancel').click();
    expect(calls).toBe(0);
    await page.getByTestId('payout-approve-1').click();
    await page.getByTestId('pay-confirm').click();
    await expect(page.getByTestId('payout-row-1')).toContainText(
      'Paid 9 Dec 2026 · transfer tr_creator-1-2026-11',
    );
    expect(calls).toBe(1);
  });

  test('a balance that cannot cover the payout turns Pay off and says what to do', async ({ page }) => {
    await statements(page, true, 12000);
    let calls = 0;
    await page.route('**/v1/admin/creators/1/statements/2026-11/approve', async (route) => {
      calls++;
      await route.fulfill(envelope(null, 409, 'balance_insufficient'));
    });
    await page.goto('/admin/creators/payouts');
    await page.getByTestId('payout-approve-1').click();
    await expect(page.getByTestId('pay-balance')).toHaveText(
      'Available in Stripe: $120.00. This payout: $168.60.',
    );
    await expect(page.getByTestId('pay-problem')).toContainText(
      'Not enough in your Stripe balance. Add funds in Stripe',
    );
    await expect(page.getByTestId('pay-confirm')).toBeDisabled();
    expect(calls).toBe(0);
  });

  test('a transfer Stripe refuses for the balance pays nothing and says so', async ({ page }) => {
    await statements(page);
    await page.route('**/v1/admin/creators/1/statements/2026-11/approve', (route) =>
      route.fulfill(envelope(null, 409, 'balance_insufficient')),
    );
    await page.goto('/admin/creators/payouts');
    await page.getByTestId('payout-approve-1').click();
    await page.getByTestId('pay-confirm').click();
    await expect(page.getByTestId('pay-problem')).toContainText('Nothing was paid');
    await expect(page.getByTestId('pay-dialog')).toBeVisible();
    await page.getByTestId('pay-cancel').click();
    await expect(page.getByTestId('payout-row-1')).toContainText('Ready to approve');
  });

  test('with payouts off the notice shows and nothing can be approved', async ({ page }) => {
    await statements(page, false);
    await page.goto('/admin/creators/payouts');
    await expect(page.getByTestId('payouts-off')).toHaveText(
      'Payouts are switched off on the server. Statements can be read but not paid.',
    );
    await expect(page.getByTestId('payout-approve-1')).toBeDisabled();
  });

  test('a key without Connect permissions says so and blocks Approve', async ({ page }) => {
    await signedIn(page, 'admin');
    await page.route('**/v1/admin/creators/statements**', (route) =>
      route.fulfill(envelope({ ...(read('statements.json') as object), connect_denied: true })),
    );
    await page.goto('/admin/creators/payouts');
    await expect(page.getByTestId('payouts-connect-denied')).toContainText(
      'Stripe payouts key missing or lacks Connect permissions',
    );
    await expect(page.getByTestId('payout-approve-1')).toBeDisabled();
  });

  test('Lines shows a creator’s lines, a refund signed', async ({ page }) => {
    await statements(page);
    await page.goto('/admin/creators/payouts');
    await expect(page.getByTestId('payout-lines')).toContainText('Player #10498 · refund 21 Nov 2026');
    await expect(page.getByTestId('payout-lines')).toContainText('−$2.50');
    await page.getByTestId('payout-lines-2').click();
    await expect(page.getByTestId('payout-lines')).toContainText('No lines this month.');
  });
});

test.describe('the creator', () => {
  test('stats: weeks without names, payouts and the dashboard link', async ({ page }) => {
    await signedIn(page, 'user', { creator: true });
    await page.route('**/v1/creators/me', (route) => route.fulfill(envelope(read('me.json'))));
    await page.route('**/v1/creators/me/dashboard', (route) =>
      route.fulfill(envelope({ url: '/creators/stats?dash=1' })),
    );
    await page.goto('/creators/stats');
    await expect(page.getByRole('heading', { level: 1, name: 'Grimbrew' })).toBeVisible();
    await expect(page.getByTestId('stats-figures')).toContainText(
      'December 2026, an estimate until the month closes · $500 monthly cap',
    );
    await expect(page.getByTestId('stats-week-2026-12-07')).toHaveText('Week of 7 Dec: 18 signed up so far');
    await expect(page.getByTestId('stats-week-2026-11-23')).toHaveText(
      'Week of 23 Nov: 54 signed up · 36 returning · 10 paying',
    );
    await expect(page.getByTestId('stats-paid-so-far')).toHaveText('$0.00 paid so far');
    await page.getByTestId('stats-dashboard').click();
    await expect(page).toHaveURL(/dash=1/);
  });

  test('onboarding: accepting an invite goes on to Stripe', async ({ page }) => {
    await signedIn(page);
    await page.route('**/v1/creators/invite/preview', (route) =>
      route.fulfill(envelope({ name: 'Grimbrew' })),
    );
    await page.route('**/v1/creators/invite', (route) => route.fulfill(envelope({ name: 'Grimbrew' })));
    await page.route('**/v1/creators/me/onboarding', (route) =>
      route.fulfill(envelope({ url: '/creators/onboarding?done=1' })),
    );
    await page.route('**/v1/creators/me**', (route) =>
      route.request().method() === 'GET' ? route.fulfill(envelope(read('me.json'))) : route.fallback(),
    );
    await page.goto('/creators/onboarding?invite=tok');
    await expect(page.getByTestId('onboarding-join')).toHaveText(
      'You’re joining the creator program as Grimbrew.',
    );
    await page.getByTestId('onboarding-accept').click();
    await expect(page.getByTestId('onboarding-ready')).toHaveText(
      'You’re set up. Payouts arrive after each month is approved.',
    );
  });

  test('onboarding: a used invite says so', async ({ page }) => {
    await signedIn(page);
    await page.route('**/v1/creators/invite/preview', (route) =>
      route.fulfill(envelope(null, 410, 'invite_used')),
    );
    await page.goto('/creators/onboarding?invite=old');
    await expect(page.getByTestId('onboarding-bad-invite')).toBeVisible();
  });
});

test.describe('/creators and /creators/apply', () => {
  test('the public page says what creators get and links Apply', async ({ page }) => {
    await page.goto('/creators');
    await expect(page.getByTestId('creators-section-what-you-get')).toContainText(
      '$2 for each player you send who comes back',
    );
    await expect(page.getByTestId('creators-section-disclosure')).toContainText(
      'the first time it is missing we warn you, and a second time ends your place in the program.',
    );
    await expect(page.getByTestId('creators-section-payouts')).toContainText(
      'A month under $10 carries over',
    );
    await expect(page.getByTestId('creators-apply-top')).toHaveAttribute('href', '/creators/apply');
    await expect(page.locator('footer a[href="/creators"]').first()).toBeVisible();
  });

  test('signed out, apply asks for the Battle.net sign-in and comes back', async ({ page }) => {
    await page.route('**/v1/me', (route) => route.fulfill(envelope(null, 401, 'unauthorized')));
    await page.goto('/creators/apply');
    await expect(page.getByTestId('apply-signin')).toBeVisible();
  });

  test('an application is sent, then shows Pending', async ({ page }) => {
    await signedIn(page);
    let sent: Record<string, unknown> | null = null;
    let status: string | null = null;
    await page.route('**/v1/creators/application', async (route) => {
      if (route.request().method() === 'POST') {
        sent = route.request().postDataJSON() as Record<string, unknown>;
        status = 'pending';
        await route.fulfill(envelope({ status }, 201));
        return;
      }
      await route.fulfill(
        envelope({ status, created_at: status ? '2026-12-08T16:20:00Z' : null, creator: false }),
      );
    });
    await page.goto('/creators/apply');
    await page.getByTestId('apply-channel').fill('https://twitch.tv/emberfall');
    await page.getByRole('radio', { name: 'YouTube' }).check();
    await page.getByTestId('apply-audience').fill('1,200');
    await page.getByTestId('apply-about').fill('Hardcore runs three nights a week.');
    await expect(page.getByTestId('apply-form')).toContainText('34 of 500 characters');
    await page.getByTestId('apply-submit').click();
    await expect(page.getByTestId('apply-error')).toHaveText(
      'Tick this to apply: the FTC and Twitch require it.',
    );
    expect(sent).toBeNull();
    await page.getByTestId('apply-disclose').locator('input').check();
    await page.getByTestId('apply-submit').click();
    await expect(page.getByTestId('apply-status-pending')).toContainText(
      'Sent 8 Dec 2026. We’ll reply here.',
    );
    expect(sent).toEqual({
      channel_url: 'https://twitch.tv/emberfall',
      platform: 'youtube',
      audience: 1200,
      about: 'Hardcore runs three nights a week.',
      contact_email: '',
    });
  });

  for (const [status, words] of [
    ['approved', 'You’re in. Set up payouts so we can pay you.'],
    ['declined', 'Not this time. Thank you for applying.'],
  ] as const) {
    test(`shows ${status}`, async ({ page }) => {
      await signedIn(page);
      await page.route('**/v1/creators/application', (route) =>
        route.fulfill(
          envelope({ status, created_at: '2026-12-08T16:20:00Z', creator: false, reapply_on: null }),
        ),
      );
      await page.goto('/creators/apply');
      await expect(page.getByTestId(`apply-status-${status}`)).toContainText(words);
      await expect(page.getByTestId('apply-form')).toHaveCount(status === 'declined' ? 1 : 0);
    });
  }

  test('a recent decline says when to apply again and holds the form', async ({ page }) => {
    await signedIn(page);
    await page.route('**/v1/creators/application', (route) =>
      route.fulfill(
        envelope({
          status: 'declined',
          created_at: '2026-12-08T16:20:00Z',
          creator: false,
          reapply_on: '2027-01-08',
        }),
      ),
    );
    await page.goto('/creators/apply');
    await expect(page.getByTestId('apply-reapply')).toHaveText('You can apply again from 8 Jan 2027.');
    await expect(page.getByTestId('apply-form')).toHaveCount(0);
  });
});
