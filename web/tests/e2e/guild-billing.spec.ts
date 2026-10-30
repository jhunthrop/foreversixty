// web/tests/e2e/guild-billing.spec.ts
// Launch pricing spec section 9, guild settings #billing: case 9 (not on a plan), 10 (trial), 11 (Guild
// tier with 28 members and the Raid guild move), 14 (lapsed, CSV) and
// 18 (axe, 44 px at 360, the seat picker open). Seats, transfer and case 15 live in guild-billing-seats.spec.ts. Also
// Manage billing and its Stripe-unavailable retry, which the old #billing section carried. API stubbed.
import { createRequire } from 'node:module';
import { expect, test, type Page } from '@playwright/test';
import type { GuildBillingSettings } from '../../src/lib/billing/contract';
import {
  GUILD_TIER_28_READ,
  LAPSED_READ,
  SETTINGS_GUILD_PATH,
  emptySeat,
  heldSeat,
  TRIAL_READ,
  UPGRADE_PREVIEW,
  guildBillingRead,
  planView,
} from '../../src/test-support/billing-fixtures';
import { guildSettingsCopy } from '../../src/lib/guild/copy';
import { uiCopy } from '../../src/lib/ui/copy';
import { GUILD_ID, envelope, openWith, stubCandidates } from './support/guild-billing';
import { fieldBox } from './support/select';

const PHONE = { width: 360, height: 780 };
const MIN_TAP_PX = 44;
/** Spec 6.1: hand-off links name the site, not the host the page is served from. */
const SITE_ORIGIN = 'https://foreversixty.gg';
const NOON = '2026-12-01T12:00:00Z';

const billing = (page: Page) => page.getByTestId('guild-settings-billing');

test('case 9: not on a plan preselects the fit, and the trial button carries plan, interval and guild', async ({
  page,
}) => {
  await openWith(page, guildBillingRead({ verified_count: 28 }));
  await expect(page.getByTestId('guild-billing-option-raid_guild').locator('input')).toBeChecked();
  await expect(billing(page)).toContainText('28 verified members today.');
  await expect(page.getByTestId('guild-billing-subscribe')).toHaveText('Start the 7-day trial');
  await expect(page.getByTestId('guild-billing-subscribe')).toHaveAttribute(
    'href',
    `/premium/checkout?plan=raid_guild&interval=yearly&guild_id=${GUILD_ID}`,
  );
  await page.getByTestId('guild-billing-interval-monthly').check();
  await page.getByTestId('guild-billing-option-guild').click();
  await expect(page.getByTestId('guild-billing-subscribe')).toHaveAttribute(
    'href',
    `/premium/checkout?plan=guild&interval=monthly&guild_id=${GUILD_ID}`,
  );
  await expect(page.getByTestId('guild-billing-trial-line')).toBeVisible();
});

test('case 9: a used trial shows Subscribe with the price and no trial line', async ({ page }) => {
  await openWith(page, guildBillingRead({ trial_used: true }));
  await expect(page.getByTestId('guild-billing-subscribe')).toContainText('Subscribe · $');
  await expect(page.getByTestId('guild-billing-subscribe')).not.toContainText('trial');
  await expect(page.getByTestId('guild-billing-trial-line')).toHaveCount(0);
});

test('case 10: a trial reads its first charge and the cancel-before date', async ({ page }) => {
  const read: GuildBillingSettings = {
    ...TRIAL_READ,
    billing: { ...TRIAL_READ.billing!, trial_end: '2026-11-04T12:00:00Z' },
  };
  await openWith(page, read);
  await expect(page.getByTestId('guild-billing-plan')).toContainText(
    'Guild · trial · first charge 4 Nov 2026: $15',
  );
  await expect(page.getByTestId('guild-billing-contact')).toHaveText(
    'Cancel before 4 Nov and nothing is charged. Billed to you.',
  );
});

test('case 11: 28 members shows the size line and one Move to Raid guild; Escape closes the confirm', async ({
  page,
}) => {
  await openWith(page, GUILD_TIER_28_READ);
  await expect(page.getByTestId('guild-size-line')).toContainText('28 verified members.');
  await expect(page.getByRole('button', { name: 'Move to Raid guild' })).toHaveCount(1);

  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, (route) =>
    route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: NOON })),
  );
  await page.getByTestId('guild-move-raid-guild').click();
  await expect(page.getByTestId('guild-upgrade-keep')).toBeFocused();
  await expect(page.getByTestId('guild-upgrade-confirm')).toContainText('Today you pay $6');
  await page.keyboard.press('Escape');
  await expect(page.getByTestId('guild-upgrade-confirm')).toHaveCount(0);
});

test('case 11: the commit updates the plan line and says it is done', async ({ page }) => {
  await openWith(page, GUILD_TIER_28_READ);
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, (route) =>
    route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: NOON })),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade`, (route) =>
    route.fulfill(
      envelope({
        ...GUILD_TIER_28_READ,
        billing: planView({ tier: 'raid_guild', amount_cents: 2500, current_period_end: NOON }),
      }),
    ),
  );
  await page.getByTestId('guild-move-raid-guild').click();
  await expect(page.getByTestId('guild-upgrade-commit')).toBeEnabled();
  await page.getByTestId('guild-upgrade-commit').click();
  await expect(page.getByTestId('guild-billing-plan')).toContainText('Raid guild · renews 1 Dec 2026');
  await expect(page.getByTestId('guild-upgrade-done')).toBeVisible();
  await expect(page.getByTestId('guild-size-line')).toHaveCount(0);
});

test('a 409 upgrade_failed shows its line and nothing moves', async ({ page }) => {
  await openWith(page, GUILD_TIER_28_READ);
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, (route) =>
    route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: NOON })),
  );
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade`, (route) =>
    route.fulfill(envelope({ code: 'upgrade_failed', message: 'declined' }, 409)),
  );
  await page.getByTestId('guild-move-raid-guild').click();
  await page.getByTestId('guild-upgrade-commit').click();
  await expect(page.getByTestId('guild-upgrade-error')).toHaveText(
    'That did not go through. Nothing was charged. Try again.',
  );
});

test('Manage billing opens the portal for the guild', async ({ page }) => {
  await openWith(page, guildBillingRead({ billing: planView({ current_period_end: NOON }) }));
  let posted: unknown;
  await page.route('**/v1/billing/portal', (route) => {
    posted = route.request().postDataJSON();
    return route.fulfill(envelope({ portal_url: 'https://billing.stripe.com/guild' }));
  });
  await page.route('https://billing.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.getByTestId('guild-manage-billing').click();
  await page.waitForURL('https://billing.stripe.com/**');
  expect(posted).toEqual({ guild_id: GUILD_ID });
});

test('case 14: lapsed reads the read-only line, Subscribe again, and both CSVs download', async ({
  page,
}) => {
  const lapsed: GuildBillingSettings = {
    ...LAPSED_READ,
    can: ['billing'],
    billing: { ...LAPSED_READ.billing!, current_period_end: '2026-10-01T12:00:00Z' },
  };
  await openWith(page, lapsed);
  await expect(page.getByTestId('guild-billing-plan')).toHaveText('Guild plan ended 1 Oct 2026.');
  await expect(page.getByTestId('guild-billing-contact')).toContainText('Officer tools are read-only.');
  await expect(page.getByTestId('guild-billing-subscribe')).toHaveText('Subscribe again');

  for (const kind of ['attendance', 'loot'] as const) {
    let contentType = '';
    await page.route(`**/v1/guilds/${GUILD_ID}/export/${kind}.csv`, (route) => {
      contentType = 'text/csv';
      return route.fulfill({
        status: 200,
        headers: {
          'content-type': 'text/csv',
          'content-disposition': `attachment; filename="the-last-watch-${kind}.csv"`,
          'access-control-allow-origin': '*',
        },
        body: 'name\n',
      });
    });
    const download = page.waitForEvent('download');
    await page.getByTestId(`guild-export-${kind}`).click();
    expect((await download).suggestedFilename()).toBe(`the-last-watch-${kind}.csv`);
    expect(contentType).toBe('text/csv');
  }
});

// --- the hand-off links (subscribe-and-join spec §6.1, §6.4) ---------------------------------------

test('a settings link made for a plan preselects that card and interval', async ({ page }) => {
  await openWith(page, guildBillingRead(), `${SETTINGS_GUILD_PATH}?plan=raid_guild&interval=monthly#billing`);
  await expect(page.getByTestId('guild-billing-option-raid_guild').locator('input')).toBeChecked();
  await expect(page.getByTestId('guild-billing-interval-monthly')).toBeChecked();
  await expect(page.getByTestId('guild-billing-subscribe')).toHaveAttribute(
    'href',
    `/premium/checkout?plan=raid_guild&interval=monthly&guild_id=${GUILD_ID}`,
  );
});

test('upgrade=1 opens the Move to Raid guild confirm on arrival, with focus on keeping Guild', async ({
  page,
}) => {
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, (route) =>
    route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: NOON })),
  );
  await openWith(
    page,
    GUILD_TIER_28_READ,
    `${SETTINGS_GUILD_PATH}?plan=raid_guild&interval=monthly&upgrade=1#billing`,
  );
  await expect(page.getByTestId('guild-upgrade-confirm')).toBeVisible();
  await expect(page.getByTestId('guild-upgrade-keep')).toBeFocused();
});

test('an officer who does not hold the claim gets the link to send the claim holder', async ({ page }) => {
  await openWith(page, guildBillingRead({ can: [], verified_count: 12 }));
  const block = page.getByTestId('guild-billing-handoff-block');
  await expect(block).toContainText(guildSettingsCopy.sendThemThisLink);
  const origin = SITE_ORIGIN;
  await expect(page.getByTestId('guild-billing-handoff-url')).toHaveValue(
    `${origin}${SETTINGS_GUILD_PATH}?plan=guild&interval=yearly#billing`,
  );
  await expect(page.getByTestId('guild-billing-handoff-copy')).toHaveText(uiCopy.copyLink);
  expect(
    (await page.getByTestId('guild-billing-handoff-copy').boundingBox())?.height ?? 0,
  ).toBeGreaterThanOrEqual(MIN_TAP_PX);
});

test('the hand-off link names Raid guild above 25 verified members', async ({ page }) => {
  await openWith(page, guildBillingRead({ can: [], verified_count: 28 }));
  const origin = SITE_ORIGIN;
  await expect(page.getByTestId('guild-billing-handoff-url')).toHaveValue(
    `${origin}${SETTINGS_GUILD_PATH}?plan=raid_guild&interval=yearly#billing`,
  );
});

interface AxeResults {
  violations: { id: string }[];
}
const AXE_SOURCE = createRequire(import.meta.url).resolve('axe-core/axe.min.js');

test('case 18: axe finds no violations and every control is at least 44px tall at 360', async ({ page }) => {
  await page.setViewportSize(PHONE);
  await stubCandidates(page);
  await openWith(page, {
    ...GUILD_TIER_28_READ,
    seats: [emptySeat(1), heldSeat(2)],
    transfer_candidates: [{ user_id: 40, name: 'Mirelle' }],
  });
  await page.getByTestId('guild-move-raid-guild').click();
  await page.getByTestId('guild-seat-1-give').click();
  await expect(fieldBox(page, 'guild-seat-1-select')).toBeVisible();
  await page.addScriptTag({ path: AXE_SOURCE });
  const results = await page.evaluate(async (): Promise<AxeResults> => {
    const axe = (window as unknown as { axe: { run: (context: string) => Promise<AxeResults> } }).axe;
    return axe.run('#billing');
  });
  expect(results.violations.map((violation) => violation.id)).toEqual([]);

  const heights = await billing(page)
    .locator('a:visible, button:visible, input:visible')
    .evaluateAll((nodes) => nodes.map((node) => node.getBoundingClientRect().height));
  for (const height of heights) expect(height).toBeGreaterThanOrEqual(MIN_TAP_PX);
});

test('Manage billing: Stripe 503 shows the unavailable line, and Try again succeeds', async ({ page }) => {
  await openWith(page, guildBillingRead({ billing: planView({ current_period_end: NOON }) }));
  let asked = 0;
  await page.route('**/v1/billing/portal', (route) => {
    asked += 1;
    if (asked === 1) {
      return route.fulfill(
        envelope({ code: 'billing_unavailable', message: 'billing is not available' }, 503),
      );
    }
    return route.fulfill(envelope({ portal_url: 'https://billing.stripe.com/guild' }));
  });
  await page.route('https://billing.stripe.com/**', (route) => route.fulfill({ status: 200, body: 'ok' }));
  await page.getByTestId('guild-manage-billing').click();
  await expect(page.getByTestId('guild-billing-unavailable')).toBeVisible();
  await page.getByTestId('guild-billing-retry').click();
  await page.waitForURL('https://billing.stripe.com/**');
  expect(asked).toBe(2);
});

test('Manage billing: a 404 refusal shows the unavailable line and no Try again', async ({ page }) => {
  await openWith(page, guildBillingRead({ billing: planView({ current_period_end: NOON }) }));
  await page.route('**/v1/billing/portal', (route) =>
    route.fulfill(envelope({ code: 'not_found', message: 'no customer' }, 404)),
  );
  await page.getByTestId('guild-manage-billing').click();
  await expect(page.getByTestId('guild-billing-unavailable')).toBeVisible();
  await expect(page.getByTestId('guild-billing-retry')).toHaveCount(0);
  await expect(page.getByTestId('guild-manage-billing')).toBeVisible();
});

test('a 409 preview_expired shows the prices-changed line and reloads the preview', async ({ page }) => {
  await openWith(page, GUILD_TIER_28_READ);
  let previews = 0;
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade-preview`, (route) => {
    previews += 1;
    return route.fulfill(envelope({ ...UPGRADE_PREVIEW, renews_at: NOON }));
  });
  await page.route(`**/v1/guilds/${GUILD_ID}/billing/upgrade`, (route) => {
    expect(route.request().postDataJSON()).toEqual({ proration_date: UPGRADE_PREVIEW.proration_date });
    return route.fulfill(envelope({ code: 'preview_expired', message: 'old' }, 409));
  });
  await page.getByTestId('guild-move-raid-guild').click();
  await page.getByTestId('guild-upgrade-commit').click();
  await expect(page.getByTestId('guild-upgrade-error')).toHaveText(
    "Today's amount was updated. Check it, then move.",
  );
  expect(previews).toBe(2);
});

test.describe('a viewer in Los Angeles', () => {
  test.use({ timezoneId: 'America/Los_Angeles' });

  test('a Stripe date just after midnight UTC reads as the day before; an operator grant stays on its day', async ({
    page,
  }) => {
    await openWith(
      page,
      guildBillingRead({ billing: planView({ current_period_end: '2026-12-01T03:00:00Z' }) }),
    );
    await expect(page.getByTestId('guild-billing-plan')).toContainText('renews 30 Nov 2026');
  });

  test('an operator grant at midnight UTC stays on its day', async ({ page }) => {
    await openWith(
      page,
      guildBillingRead({
        billing: planView({ current_period_end: '2026-12-01T00:00:00Z', billed_by: '' }),
      }),
    );
    await expect(page.getByTestId('guild-billing-plan')).toContainText('until 1 Dec 2026');
  });
});
