// web/tests/e2e/premium.spec.ts
// The plans page in the flat system (design/specs/2026-11-08-premium-account.md sections 4, 5, 12,
// 13): founder prices until 6 January 2027 with the founder line said once, the interval bar, the
// Premium column first on a phone, the owned look from the stored hint before /v1/me, the signed-in
// answer, state lines and guild column (G1, G2), the canceled return. Prices come from
// lib/billing/plans, never typed here.
import { createRequire } from 'node:module';
import type { Page } from '@playwright/test';
import { expect, test } from './support/test';
import { priceFor } from '../../src/lib/billing/plan-costs';
import { FOUNDER_WINDOW_CLOSES, PLAN_ORDER, formatPrice } from '../../src/lib/billing/plans';
import { PREMIUM_SOURCE_KEY } from '../../src/lib/account/premium-hint';
import { plansPageCopy } from '../../src/lib/billing/plans-page-copy';
import { premiumSourceCopy } from '../../src/lib/billing/premium-source-copy';
import { billingBlockCopy } from '../../src/lib/billing/copy';
import { checkoutGuildCopy } from '../../src/lib/billing/checkout-guild-copy';
import {
  PREMIUM_FREE,
  PREMIUM_GRANT_OPEN_ENDED,
  PREMIUM_OWN_FOUNDER,
  PREMIUM_PAST_DUE,
  PREMIUM_SEAT_WITH_PORTAL,
} from '../../src/test-support/subscribe-fixtures';
import { plantSessionHint, routeMeCell } from './support/subscribe-stubs';

const LAUNCH_DAY = '2026-10-27T18:00:00Z';
const DECEMBER = '2026-12-08T18:00:00Z';
const JUST_BEFORE_CLOSE = '2027-01-07T07:59:59Z';
const PHONE = { width: 360, height: 780 };
const MIN_TAP_PX = 44;

const premiumPrice = (interval: 'yearly' | 'monthly', founder: boolean): string =>
  formatPrice(priceFor('premium', interval, founder), interval);

async function openAt(page: Page, time: string, options: { path?: string; stored?: string | null } = {}) {
  const { path = '/premium', stored = null } = options;
  if (stored !== null) {
    await page.addInitScript(
      ([key, value]) => localStorage.setItem(key as string, value as string),
      [PREMIUM_SOURCE_KEY, stored],
    );
  }
  await page.clock.install({ time });
  await page.goto(path);
}

/** Below 768 the interval bar is one line with Change (ui/ControlBar): open it to reach its
 *  radios. */
async function openIntervalBar(page: Page): Promise<void> {
  const toggle = page.getByTestId('interval-bar').locator('[data-control-bar-toggle]');
  if (!(await toggle.isVisible())) return;
  // The bar's own script (ControlBar.astro) wires Change once it runs: click until it opens.
  await expect(async () => {
    if ((await toggle.getAttribute('aria-expanded')) !== 'true') await toggle.click();
    await expect(toggle).toHaveAttribute('aria-expanded', 'true', { timeout: 500 });
  }).toPass();
}

const premiumButton = (page: Page) =>
  page.getByTestId('plan-premium').getByRole('link', { name: /^Subscribe · / });
const radio = (page: Page, label: string) =>
  page
    .getByRole('radiogroup', { name: plansPageCopy.interval.groupLabel })
    .getByRole('radio', { name: label });

test('the page itself ships no island: the one astro-island is the layout character selector', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  const islands = page.locator('astro-island');
  await expect(islands).toHaveCount(1);
  await expect(islands.first()).toHaveAttribute('component-url', /CharacterSelector/);
});

test('on launch day the founder line shows once, prices are founder prices, no strike-through', async ({
  page,
}) => {
  await openAt(page, LAUNCH_DAY);
  await expect(page.getByTestId('founder-row')).toContainText(plansPageCopy.founder.open);
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
  await expect(
    page.getByTestId('price-premium').getByTestId('founder-line').filter({ visible: true }),
  ).toHaveText('Founder price for anyone who subscribes by 6 Jan 2027. $50 after that.');
  await expect(
    page.getByTestId('plan-premium').getByText(plansPageCopy.founder.label, { exact: true }),
  ).toHaveCount(0);
  await expect(page.locator('s:visible')).toHaveCount(0);
});

test('Yearly is the default; Monthly changes every price and button, writes the URL and survives a reload', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  await openIntervalBar(page);
  await expect(radio(page, 'Yearly')).toHaveAttribute('aria-checked', 'true');
  await expect(premiumButton(page)).toHaveAttribute('href', /plan=premium&interval=yearly/);

  await radio(page, 'Monthly').click();
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('monthly', true)}`);
  await expect(premiumButton(page)).toHaveAttribute('href', /plan=premium&interval=monthly/);
  for (const plan of PLAN_ORDER) {
    await expect(page.getByTestId(`price-${plan}`)).toContainText(`$${priceFor(plan, 'monthly', true)}`);
  }
  await expect(page).toHaveURL(/interval=monthly/);
  await page.reload();
  await openIntervalBar(page);
  await expect(radio(page, 'Monthly')).toHaveAttribute('aria-checked', 'true');
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('monthly', true)}`);

  await openIntervalBar(page);
  await radio(page, 'Monthly').focus();
  await page.keyboard.press('ArrowLeft');
  // A choice closes the phone's open bar (ui/ControlBar): open it again to read the radios.
  await openIntervalBar(page);
  await expect(radio(page, 'Yearly')).toHaveAttribute('aria-checked', 'true');
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
});

test('the founder window ends at the instant, to the second', async ({ page }) => {
  await openAt(page, JUST_BEFORE_CLOSE);
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', true)}`);
  await expect(page.getByTestId('founder-row')).toBeVisible();

  await openAt(page, FOUNDER_WINDOW_CLOSES);
  await expect(page.getByTestId('founder-row')).toBeHidden();
  await expect(premiumButton(page)).toHaveText(`Subscribe · ${premiumPrice('yearly', false)}`);
  await expect(
    page.getByTestId('price-premium').getByTestId('founder-line').filter({ visible: true }),
  ).toHaveCount(0);
});

test('at 390 Premium stacks first and its Subscribe is in the first screen (PR3)', async ({ page }) => {
  await page.setViewportSize({ width: 390, height: 844 });
  await openAt(page, DECEMBER);
  const premium = await page.getByTestId('plan-premium').boundingBox();
  const free = await page.getByTestId('plan-free').boundingBox();
  expect(premium!.y).toBeLessThan(free!.y);
  const button = await premiumButton(page).boundingBox();
  expect(button!.y + button!.height).toBeLessThanOrEqual(844);
});

test('forbidden words stay off the page, the coach line leads Premium, rules and questions are open', async ({
  page,
}) => {
  await openAt(page, DECEMBER);
  const text = (await page.locator('main').textContent()) ?? '';
  expect(text).not.toMatch(/chip in|retention|compare more than/i);
  expect(text).not.toContain('!');
  expect(text).toContain(plansPageCopy.premium.features[0]?.text ?? 'missing');
  expect(text).not.toContain(plansPageCopy.trialEmail);
  for (const entry of plansPageCopy.faq) await expect(page.getByText(entry.answer)).toBeVisible();
  await expect(page.locator('main details')).toHaveCount(0);
});

test('the sample coach review is labelled Sample in its head and its caption', async ({ page }) => {
  await openAt(page, DECEMBER);
  const sample = page.getByTestId('coach-sample');
  await expect(sample.getByTestId('coach-sample-tag')).toHaveText(plansPageCopy.coachSample.tag);
  await expect(sample.getByTestId('coach-sample-caption')).toHaveText(plansPageCopy.coachSample.caption);
});

interface AxeResults {
  violations: { id: string }[];
}
const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

for (const stored of [null, 'own'] as const) {
  test(`axe finds no violations and every button is 44px tall at 360 (${stored ?? 'free visitor'})`, async ({
    page,
  }) => {
    await page.setViewportSize(PHONE);
    await openAt(page, DECEMBER, { stored });
    await page.addScriptTag({ path: AXE_SOURCE });
    const results = await page.evaluate(async (): Promise<AxeResults> => {
      const axe = (window as unknown as { axe: { run: () => Promise<AxeResults> } }).axe;
      return axe.run();
    });
    expect(results.violations.map((violation) => violation.id)).toEqual([]);
    const heights = await page
      .locator('main .ui-plh ~ * a:visible, main [role=radio]:visible')
      .evaluateAll((nodes) => nodes.map((node) => node.getBoundingClientRect().height));
    for (const height of heights) expect(height).toBeGreaterThanOrEqual(MIN_TAP_PX);
  });
}

const SOURCE_STATES = [
  { source: 'own', me: PREMIUM_OWN_FOUNDER },
  { source: 'seat', me: PREMIUM_SEAT_WITH_PORTAL },
  { source: 'grant', me: PREMIUM_GRANT_OPEN_ENDED },
] as const;

for (const { source, me } of SOURCE_STATES) {
  test(`stored source ${source}: the Premium column reads its line and See your plan on first paint`, async ({
    page,
  }) => {
    await plantSessionHint(page);
    await routeMeCell(page, () => me);
    await openAt(page, DECEMBER, { stored: source });
    const card = page.getByTestId('plan-premium');
    await expect(card.locator(`.owned-line[data-source="${source}"]`)).toHaveText(
      premiumSourceCopy.premiumCard[source],
    );
    await expect(card.getByRole('link', { name: premiumSourceCopy.buttons.seePlan })).toHaveAttribute(
      'href',
      '/account#plan',
    );
    await expect(card.getByRole('link', { name: /^Subscribe · / })).toBeHidden();
    await expect(card.getByText(plansPageCopy.yourPlanTag, { exact: true })).toBeVisible();
    await expect(page.getByTestId('premium-answer')).toContainText(premiumSourceCopy.buttons.seePlan);
  });
}

test('a malformed stored source renders the selling column and the key is gone', async ({ page }) => {
  await plantSessionHint(page);
  await routeMeCell(page, () => null);
  await openAt(page, DECEMBER, { stored: 'premium' });
  await expect(page.getByTestId('plan-premium').locator('.owned-line:visible')).toHaveCount(0);
  await expect(premiumButton(page)).toBeVisible();
  await expect(page.locator('html')).not.toHaveAttribute('data-premium', /.*/);
  await expect.poll(() => page.evaluate((key) => localStorage.getItem(key), PREMIUM_SOURCE_KEY)).toBeNull();
});

test('signed in on a free account: the answer says Free and Free carries Your plan (G1)', async ({
  page,
}) => {
  await plantSessionHint(page);
  await routeMeCell(page, () => PREMIUM_FREE);
  await openAt(page, DECEMBER);
  await expect(page.getByTestId('premium-answer')).toContainText('Free');
  await expect(
    page.getByTestId('plan-free').getByText(plansPageCopy.yourPlanTag, { exact: true }),
  ).toBeVisible();
  await expect(premiumButton(page)).toBeVisible();
});

test('a failed payment says so under the bar, as an alert with Manage billing', async ({ page }) => {
  await plantSessionHint(page);
  await routeMeCell(page, () => PREMIUM_PAST_DUE);
  await openAt(page, DECEMBER, { stored: 'own' });
  const line = page.getByTestId('premium-state-pastdue');
  await expect(line.getByRole('alert')).toHaveText(billingBlockCopy.pastDueBanner);
  await expect(line.getByRole('button', { name: /Manage billing/ })).toBeVisible();
});

test('an ended seat links back at its own price (G2)', async ({ page }) => {
  const ended = {
    ...PREMIUM_FREE,
    ended_seat: {
      guild_name: 'Olympus XXVII',
      ended_at: '2026-12-02T10:00:00Z',
      resubscribe_url: '/premium/checkout?plan=premium&interval=yearly&comeback=9',
      amount_cents: 4000,
      interval: 'yearly',
    },
  };
  await plantSessionHint(page);
  await routeMeCell(page, () => ended as typeof PREMIUM_FREE);
  await openAt(page, DECEMBER);
  const line = page.getByTestId('premium-state-ended');
  await expect(line).toContainText('Olympus XXVII');
  await expect(line.getByRole('link')).toHaveAttribute('href', ended.ended_seat.resubscribe_url);
});

test('an officer paying for a guild on Guild sees it on that column with Manage billing', async ({
  page,
}) => {
  const officer = {
    ...PREMIUM_OWN_FOUNDER,
    guilds: [
      {
        id: 4,
        region: 'us',
        ruleset: 'normal',
        name: 'Olympus XXVII',
        rank: 'Officer',
        can: ['billing'],
        plan: {
          tier: 'guild',
          you_hold_seat: false,
          scope: 'plan',
          status: 'active',
          current_period_end: '2027-10-04T10:00:00Z',
          cancel_at_period_end: false,
          billed_by: 'Thandrel',
          you_are_billing_contact: true,
        },
      },
    ],
  };
  await plantSessionHint(page);
  await routeMeCell(page, () => officer as unknown as typeof PREMIUM_OWN_FOUNDER);
  await openAt(page, DECEMBER, { stored: 'own' });
  const column = page.getByTestId('plan-guild');
  await expect(column.getByTestId('guild-on-plan')).toHaveText(
    checkoutGuildCopy.onPlan.youPay('Olympus XXVII', 'guild'),
  );
  await expect(column.getByTestId('guild-on-plan-manage')).toBeVisible();
  await expect(column.getByRole('link', { name: /^Subscribe your guild/ })).toBeHidden();
  await expect(
    page.getByTestId('plan-raid_guild').getByRole('link', { name: /^Subscribe your guild/ }),
  ).toBeVisible();
});

test('?canceled=1 acknowledges a canceled Stripe Checkout and loads no third-party script', async ({
  page,
}) => {
  const scripts: string[] = [];
  page.on('request', (r) => {
    if (r.resourceType() === 'script') scripts.push(r.url());
  });
  await page.goto('/premium?canceled=1');
  await expect(page.getByTestId('premium-canceled-notice')).toContainText(
    'Checkout was canceled. Nothing was charged.',
  );
  const origin = new URL(page.url()).origin;
  expect(scripts.filter((url) => !url.startsWith(`${origin}/`))).toEqual([]);

  await page.goto('/premium');
  await expect(page.getByTestId('premium-canceled-notice')).toBeHidden();
});
