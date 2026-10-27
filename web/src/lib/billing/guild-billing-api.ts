// web/src/lib/billing/guild-billing-api.ts
// The guild billing routes (launch pricing spec section 4): the officer read, the upgrade and
// the plan transfer offer. Same transport and error class as ./api, so a refusal is a
// `BillingApiError` carrying the API's status and `code` (`upgrade_failed`, `not_member`).
import { invalidateGuild } from '../guild/api';
import { API_BASE_URL } from '../planner/config';
import { call, callVoid } from './api';
import { normalizeGuildBilling } from './normalize';
import type { GuildBillingSettings, UpgradePreview } from './contract';

export type ExportKind = 'attendance' | 'loot';

function billingPath(guildId: number): string {
  return `/v1/guilds/${guildId}/billing`;
}

/** Verified officers and leaders only; anyone else gets a 404 `BillingApiError`. */
export async function getGuildBilling(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildBillingSettings> {
  return normalizeGuildBilling(await call<GuildBillingSettings>(billingPath(guildId), apiBase));
}

/** What the Guild to Raid guild move charges now and from the next renewal. A payer's read. */
export function previewUpgrade(guildId: number, apiBase: string = API_BASE_URL): Promise<UpgradePreview> {
  return call<UpgradePreview>(`${billingPath(guildId)}/upgrade-preview`, apiBase);
}

/** Moves the plan to Raid guild at the previewed proration date (409 `preview_expired` after 10 minutes); answers the billing read as it now stands. */
export async function upgrade(
  guildId: number,
  prorationDate: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildBillingSettings> {
  const data = await call<GuildBillingSettings>(`${billingPath(guildId)}/upgrade`, apiBase, {
    method: 'POST',
    body: { proration_date: prorationDate },
  });
  invalidateGuild(guildId, apiBase);
  return normalizeGuildBilling(data);
}

async function transferAction(
  guildId: number,
  action: '' | '/withdraw' | '/decline',
  body: unknown,
  apiBase: string,
): Promise<void> {
  await callVoid(`${billingPath(guildId)}/transfer${action}`, apiBase, { method: 'POST', body });
  invalidateGuild(guildId, apiBase);
}

/** The subscriber offers the plan to another verified officer. */
export function sendTransfer(
  guildId: number,
  toUserId: number,
  apiBase: string = API_BASE_URL,
): Promise<void> {
  return transferAction(guildId, '', { to_user_id: toUserId }, apiBase);
}

export function withdrawTransfer(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  return transferAction(guildId, '/withdraw', {}, apiBase);
}

export function declineTransfer(guildId: number, apiBase: string = API_BASE_URL): Promise<void> {
  return transferAction(guildId, '/decline', {}, apiBase);
}

/**
 * The CSV download link: absolute to the API origin, because the website host does not serve
 * `/v1`. Cross-origin, so a `download` attribute is ignored: the file name comes from the
 * API's Content-Disposition header, and nothing here can set it.
 */
export function exportHref(guildId: number, kind: ExportKind, apiBase: string = API_BASE_URL): string {
  return `${apiBase}/v1/guilds/${guildId}/export/${kind}.csv`;
}
