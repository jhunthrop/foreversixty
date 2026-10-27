// web/tests/captures/launch-pricing/capture.spec.ts
// Acceptance screenshots for launch pricing (design/specs/2026-10-27-launch-pricing.md §9, the
// screenshot table, rows 1-12 and 14; row 13, the emails in Gmail web and iOS Mail, needs real
// mail clients). Not a test of behaviour: it drives /premium, the checkout chooser, the guild
// settings #billing section and the account Plan panel into each state with the API stubbed, and
// writes PNGs to design/reviews/2026-10-27-launch-pricing/. Row 8 and 12 print CLS on a warm load.
import { writeFileSync } from 'node:fs';
import path from 'node:path';
import { test, type Locator, type Page } from '@playwright/test';
import { FOUNDER_WINDOW_CLOSES } from '../../../src/lib/billing/plans';
import type { GuildBillingSettings } from '../../../src/lib/billing/contract';
import {
  FOUNDER_PREMIUM_FACTS,
  GUILD_TIER_28_READ,
  LAPSED_READ,
  ME_ENDED_SEAT,
  ME_SEAT,
  ME_SEAT_OFFER,
  READ_ONLY_OFFICER_READ,
  TRANSFER_OFFERED,
  TRANSFER_RECEIVED,
  UPGRADE_PREVIEW,
  emptySeat,
  guildBillingRead,
  heldSeat,
  mockEntitlements,
  pendingSeat,
  planView,
  returnedSeat,
  TRIAL_READ,
} from '../../../src/test-support/billing-fixtures';
import { readLayoutShift, trackLayoutShift } from '../../e2e/support/cls';
import { GUILD_ID, envelope, openLive, openWith, stubCandidates } from '../../e2e/support/guild-billing';
import { gotoHydrated } from '../../e2e/support/hydrated';
import {
  MOCK_DIR,
  atWidths,
  cropOf,
  fresh,
  settle,
  sideBySide,
  sizeTo,
  targetIn,
  two,
} from '../support/capture-kit';

const target = targetIn('2026-10-27-launch-pricing');
const LAUNCH_DAY = '2026-10-27T18:00:00Z';
const DECEMBER = '2026-12-08T18:00:00Z';
const BILLING_DAY = '2026-11-20T12:00:00Z';
const RENEWS = '2026-12-01T12:00:00Z';
const BILLING_ID = 'guild-settings-billing';

const ALL = [360, 390, 1440, 2000];
const THREE = [360, 390, 1440];
const PAIR = [390, 1440];

const file = (row: number, state: string, width: number, kind = ''): string =>
  target(`launch-pricing-${two(row)}-${state}${kind}-${width}.png`);

/** A full-page capture and, with a subject, a crop of it. */
async function shoot(page: Page, row: number, state: string, width: number, subject?: Locator) {
  await page.screenshot({ path: file(row, state, width), fullPage: true });
  if (subject !== undefined) await cropOf(page, subject, file(row, state, width, '-crop'));
}

// ---- /premium -----------------------------------------------------------------------------------

async function openPlans(page: Page, width: number, time: string): Promise<void> {
  await sizeTo(page, width);
  await page.clock.install({ time });
  await page.goto('/premium');
  await settle(page);
}

atWidths('1 plans, founder open, Yearly', ALL, async (page, width) => {
  await openPlans(page, width, DECEMBER);
  await shoot(page, 1, 'founder-yearly', width);
});

atWidths('2 plans, Monthly', PAIR, async (page, width) => {
  await openPlans(page, width, DECEMBER);
  await page.getByLabel('Monthly').check();
  await settle(page);
  await shoot(page, 2, 'monthly', width);
});

atWidths('3 plans, launch day and after the window', PAIR, async (page, width) => {
  // R-P8 ruling: there is no "before" state; launch day is the first day the founder row shows.
  await openPlans(page, width, LAUNCH_DAY);
  await shoot(page, 3, 'launch-day', width);
  const closed = await fresh(page);
  await openPlans(closed, width, FOUNDER_WINDOW_CLOSES);
  await shoot(closed, 3, 'after-window', width);
});

test('4 plans, focus-visible and hover @1440', async ({ page }) => {
  await openPlans(page, 1440, DECEMBER);
  await page.keyboard.press('Shift');
  await page.getByLabel('Yearly').focus();
  await settle(page);
  await page.screenshot({ path: file(4, 'focus-interval', 1440), fullPage: true });
  const button = page.getByRole('link', { name: /^Subscribe · / });
  await button.focus();
  await settle(page);
  await cropOf(page, page.getByTestId('plan-premium'), file(4, 'focus-card-button', 1440, '-crop'));
  await page.mouse.move(0, 0);
  await button.hover();
  await settle(page);
  await cropOf(page, page.getByTestId('plan-premium'), file(4, 'hover-card-button', 1440, '-crop'));
});

// ---- 5: the checkout guild chooser ------------------------------------------------------------

function chooserGuild(id: number, name: string, verifiedCount: number, can: string[]) {
  return { id, region: 'us', ruleset: 'forever', name, verified: true, verified_count: verifiedCount, can, plan: null };
}

async function stubMe(page: Page, guilds: unknown[]): Promise<void> {
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds,
        entitlements: mockEntitlements(),
      }),
    ),
  );
}

atWidths('5 checkout chooser', PAIR, async (page, width) => {
  await sizeTo(page, width);
  await stubMe(page, [
    chooserGuild(1, 'Olympus', 12, ['billing']),
    chooserGuild(2, 'Titans', 28, ['billing']),
    chooserGuild(3, 'Other', 9, []),
  ]);
  await page.goto('/premium/checkout?plan=guild&interval=yearly');
  await page.getByTestId('checkout-guild-row').first().waitFor();
  await settle(page);
  await shoot(page, 5, 'two-guilds', width);
  const none = await fresh(page);
  await sizeTo(none, width);
  await stubMe(none, [chooserGuild(3, 'Other', 9, [])]);
  await none.goto('/premium/checkout?plan=guild&interval=yearly');
  await none.getByText('Find your guild').waitFor();
  await settle(none);
  await shoot(none, 5, 'no-billable-guild', width);
});

// ---- 6-11: guild settings #billing ------------------------------------------------------------

async function billingPage(page: Page, width: number, read: GuildBillingSettings, live = false) {
  await sizeTo(page, width);
  await page.clock.install({ time: BILLING_DAY });
  await stubCandidates(page);
  const cell = live ? await openLive(page, read) : (await openWith(page, read), undefined);
  await settle(page);
  return cell;
}

const billing = (page: Page): Locator => page.getByTestId(BILLING_ID);

atWidths('6 billing, not on a plan', THREE, async (page, width) => {
  await billingPage(page, width, guildBillingRead({ verified_count: 28 }));
  await shoot(page, 6, 'trial-available', width, billing(page));
  const used = await fresh(page);
  await billingPage(used, width, guildBillingRead({ trial_used: true }));
  await shoot(used, 6, 'trial-used', width, billing(used));
});

atWidths('7 billing, trial', PAIR, async (page, width) => {
  await billingPage(page, width, {
    ...TRIAL_READ,
    billing: { ...TRIAL_READ.billing!, trial_end: '2026-11-27T12:00:00Z' },
  });
  await shoot(page, 7, 'trial', width, billing(page));
});

atWidths('8 billing, Guild with 28 members', THREE, async (page, width) => {
  await billingPage(page, width, GUILD_TIER_28_READ);
  await shoot(page, 8, 'size-line', width, billing(page));
  let release: () => void = () => undefined;
  const held = new Promise<void>((resolve) => (release = resolve));
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, async (route) => {
    await held;
    await route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: RENEWS }));
  });
  await page.getByTestId('guild-move-raid-guild').click();
  await page.getByTestId('guild-upgrade-confirm').waitFor();
  await settle(page);
  await shoot(page, 8, 'confirm-loading', width, billing(page));
  release();
  await page.getByTestId('guild-upgrade-commit').waitFor();
  await page.waitForFunction(() => !document.querySelector('[data-testid="guild-upgrade-commit"][disabled]'));
  await settle(page);
  await shoot(page, 8, 'confirm-ready', width, billing(page));
});

const RAID_PLAN = planView({ tier: 'raid_guild', interval: 'yearly', amount_cents: 20000, current_period_end: RENEWS });
const SEAT_STATES = [
  heldSeat(1, { movable_at: '2026-12-02T12:00:00Z' }),
  pendingSeat(2),
  heldSeat(3, { state: 'accepted', movable_at: '2026-11-01T12:00:00Z' }),
  emptySeat(4),
  returnedSeat(5),
];

atWidths('9 billing, Raid guild seats', ALL, async (page, width) => {
  const read = guildBillingRead({ billing: RAID_PLAN, verified_count: 31, seats: SEAT_STATES });
  await billingPage(page, width, read);
  await shoot(page, 9, 'seats', width, page.getByTestId('guild-seats'));
  await page.getByTestId('guild-seat-4-give').click();
  await page
    .getByTestId('guild-seat-4-select')
    .selectOption({ label: 'Mirelle · priest · pays for Premium' });
  await page.getByTestId('guild-seat-4-pays-note').waitFor();
  await settle(page);
  await shoot(page, 9, 'picker-payer', width, page.getByTestId('guild-seats'));
});

atWidths('10 billing, Who pays', PAIR, async (page, width) => {
  const read = guildBillingRead({
    billing: planView(),
    transfer_candidates: [{ user_id: 40, name: 'Mirelle' }],
  });
  const cell = await billingPage(page, width, read, true);
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/transfer`, (route) => {
    if (cell !== undefined) cell.current = { ...read, transfer: TRANSFER_OFFERED };
    return route.fulfill({ status: 204 });
  });
  await page.getByTestId('guild-transfer-open').click();
  await page.getByTestId('guild-transfer-consequence').waitFor();
  await settle(page);
  await shoot(page, 10, 'send-confirm', width, billing(page));
  await page.getByTestId('guild-transfer-send').click();
  await page.getByTestId('guild-transfer-waiting').waitFor();
  await settle(page);
  await shoot(page, 10, 'waiting', width, billing(page));
  const receiver = await fresh(page);
  await billingPage(
    receiver,
    width,
    guildBillingRead({
      billing: planView({ you_are_billing_contact: false }),
      transfer: TRANSFER_RECEIVED,
      can: [],
    }),
  );
  await shoot(receiver, 10, 'receiver-offer', width, billing(receiver));
});

atWidths('11 billing, lapsed and read-only', PAIR, async (page, width) => {
  await billingPage(page, width, {
    ...LAPSED_READ,
    can: ['billing'],
    billing: { ...LAPSED_READ.billing!, current_period_end: '2026-10-01T12:00:00Z' },
  });
  await shoot(page, 11, 'lapsed-records', width, billing(page));
  const officer = await fresh(page);
  await billingPage(officer, width, READ_ONLY_OFFICER_READ);
  await shoot(officer, 11, 'read-only-officer', width, billing(officer));
});

// ---- 12: the account Plan panel ---------------------------------------------------------------

const OWN_PREMIUM = {
  plan: 'premium',
  status: 'active',
  current_period_end: '2026-12-20T00:00:00Z',
  cancel_at_period_end: false,
  founder: false,
  interval: 'monthly',
  amount_cents: 600,
};

async function openAccount(page: Page, width: number, extra: Record<string, unknown>, own: unknown = null) {
  await sizeTo(page, width);
  await page.clock.install({ time: BILLING_DAY });
  await page.route('**/v1/me', (route) =>
    route.fulfill(
      envelope({
        user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
        characters: [],
        guilds: [],
        entitlements: mockEntitlements({ billing: own as never }),
        notifications: { discord: null },
        ...extra,
      }),
    ),
  );
  await page.route('**/v1/devices', (route) => route.fulfill(envelope([])));
  await gotoHydrated(page, '/account#plan', 'account');
  await settle(page);
}

atWidths('12 account plan', THREE, async (page, width) => {
  const plan = (on: Page): Locator => on.locator('#plan');
  await openAccount(page, width, { seat_offer: ME_SEAT_OFFER }, OWN_PREMIUM);
  await shoot(page, 12, 'offer', width, plan(page));
  await page.route('**/v1/guilds/501/seats/offer/accept', (route) => route.fulfill({ status: 204, body: '' }));
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
  await page.getByTestId('account-seat-offer-status').waitFor();
  await settle(page);
  await shoot(page, 12, 'accepted', width, plan(page));
  const states: [string, Record<string, unknown>, unknown][] = [
    ['seat-holder', { seat: { ...ME_SEAT, starts_at: null } }, null],
    ['founder', {}, { ...OWN_PREMIUM, ...FOUNDER_PREMIUM_FACTS }],
    ['seat-ended', { ended_seat: ME_ENDED_SEAT }, null],
  ];
  for (const [state, extra, own] of states) {
    const on = await fresh(page);
    await openAccount(on, width, extra, own);
    await shoot(on, 12, state, width, plan(on));
  }
});

// ---- CLS: #billing warm load -------------------------------------------------------------------

atWidths('CLS billing warm', PAIR, async (page, width) => {
  await sizeTo(page, width);
  await page.clock.install({ time: BILLING_DAY });
  await trackLayoutShift(page);
  await stubCandidates(page);
  const read = guildBillingRead({ billing: RAID_PLAN, verified_count: 31, seats: SEAT_STATES });
  await openWith(page, read);
  await page.reload();
  await billing(page).waitFor();
  const cls = Math.round((await readLayoutShift(page)) * 10_000) / 10_000;
  writeFileSync(target(`cls-${width}.json`), `${JSON.stringify({ width, billing: cls, cap: 0.02 })}\n`);
});

// ---- 14: side by side with the mock ------------------------------------------------------------

for (const width of [1024, 1440, 1920, 2000]) {
  test(`14 plans side by side @${width}`, async ({ browser }) => {
    const context = await browser.newContext({ reducedMotion: 'reduce', timezoneId: 'America/Chicago' });
    const mockPage = await context.newPage();
    await sizeTo(mockPage, width);
    const board = width === 2000 ? 'launch-pricing-2000' : 'launch-pricing';
    await mockPage.goto(`file://${path.join(MOCK_DIR, `${board}.html`)}`);
    await mockPage.waitForLoadState('networkidle');
    const mockPng = file(14, 'mock', width);
    await mockPage.screenshot({ path: mockPng, fullPage: true });
    const built = await context.newPage();
    await built.goto(`http://localhost:${process.env.CAPTURE_PORT ?? '4481'}/premium`);
    await sizeTo(built, width);
    await built.clock.install({ time: DECEMBER });
    await built.reload();
    await settle(built);
    const builtPng = file(14, 'built', width);
    await built.screenshot({ path: builtPng, fullPage: true });
    await sideBySide(await context.newPage(), mockPng, builtPng, file(14, 'side-by-side', width), width);
    await context.close();
  });
}
