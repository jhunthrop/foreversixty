// web/tests/e2e/support/guild-billing.ts
// The guild settings #billing page, API stubbed: the envelope the API answers with, and the openers
// the billing specs share. `openLive` keeps the read in a cell so an action can change the reload.
import { expect, type Page } from '@playwright/test';
import type { GuildBillingSettings } from '../../../src/lib/billing/contract';
import {
  SEAT_CANDIDATES,
  SETTINGS_GUILD,
  SETTINGS_GUILD_PATH,
  meWithGuildPlan,
  routeGuildSettingsPage,
} from '../../../src/test-support/billing-fixtures';
import { gotoHydrated } from './hydrated';

export const GUILD_ID = SETTINGS_GUILD.id;

export function envelope(data: unknown, status = 200) {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({
      ok: status < 400,
      data: status < 400 ? data : null,
      error: status < 400 ? null : data,
      request_id: 'req-test',
    }),
  };
}

export async function openWith(page: Page, read: GuildBillingSettings): Promise<void> {
  await routeGuildSettingsPage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithGuildPlan(null).data)));
  await page.route(`**/v1/guilds/${GUILD_ID}/billing`, (route) => route.fulfill(envelope(read)));
  await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings-billing');
  await expect(page.getByTestId('guild-billing-skeleton')).toHaveCount(0);
}

/** The billing read is a mutable cell so a seat or transfer action can change what the reload sees. */
export async function openLive(
  page: Page,
  read: GuildBillingSettings,
): Promise<{ current: GuildBillingSettings }> {
  const cell = { current: read };
  await routeGuildSettingsPage(page);
  await page.route('**/v1/me', (route) => route.fulfill(envelope(meWithGuildPlan(null).data)));
  await page.route(`**/v1/guilds/${GUILD_ID}/billing`, (route) => route.fulfill(envelope(cell.current)));
  await gotoHydrated(page, SETTINGS_GUILD_PATH, 'guild-settings-billing');
  await expect(page.getByTestId('guild-billing-skeleton')).toHaveCount(0);
  return cell;
}

export async function stubCandidates(page: Page): Promise<void> {
  await page.route(`**/v1/guilds/${GUILD_ID}/seats/candidates`, (route) =>
    route.fulfill(envelope({ candidates: SEAT_CANDIDATES })),
  );
}
