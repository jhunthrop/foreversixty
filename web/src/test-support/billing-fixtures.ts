// web/src/test-support/billing-fixtures.ts
// e2e stubs for the billing routes, following account/api.ts's requestEnvelope envelope
// shape and the sim lane's existing sim-api.ts pattern of one small helper per route rather
// than each spec hand-rolling page.route calls.
import type { Page } from '@playwright/test';
import type { EntitlementsView, GuildBillingView } from '../lib/account/api';

export function mockEntitlements(overrides: Partial<EntitlementsView> = {}): EntitlementsView {
  return {
    server_sims: false,
    retention: false,
    multi_compare: false,
    history: false,
    notifications: false,
    officer_views: false,
    roster_check: false,
    supporter_mark: false,
    billing: null,
    ...overrides,
  };
}

function envelope(data: unknown, status: number): { status: number; contentType: string; body: string } {
  return {
    status,
    contentType: 'application/json',
    body: JSON.stringify({ ok: status < 400, data, error: null, request_id: 'r' }),
  };
}

export async function stubCheckout(page: Page, response: { status: number; data?: unknown }): Promise<void> {
  await page.route('**/v1/billing/checkout', (route) =>
    route.fulfill(envelope(response.data ?? null, response.status)),
  );
}

export async function stubPortal(page: Page, response: { status: number; data?: unknown }): Promise<void> {
  await page.route('**/v1/billing/portal', (route) =>
    route.fulfill(envelope(response.data ?? null, response.status)),
  );
}

/** The guild the settings-billing specs open at SETTINGS_GUILD_PATH. */
export const SETTINGS_GUILD = { id: 501, name: 'The Last Watch', region: 'us', ruleset: 'hardcore' } as const;
export const SETTINGS_GUILD_PATH = '/guild/us/hardcore/the-last-watch/settings';

/** Every capability the claim holder has; an officer without the claim lacks billing and release. */
export const CLAIM_HOLDER_CAN: readonly string[] = [
  'approve',
  'remove',
  'settings',
  'award',
  'billing',
  'release',
];
export const OFFICER_CAN: readonly string[] = ['approve', 'remove', 'settings', 'award'];

/** A `/v1/me` guilds[].plan: the paid guild plan billed to the viewer unless overridden. */
export function guildBilling(overrides: Partial<GuildBillingView> = {}): GuildBillingView {
  return {
    scope: 'plan',
    status: 'active',
    current_period_end: '2026-12-01T00:00:00Z',
    cancel_at_period_end: false,
    billed_by: 'you',
    you_are_billing_contact: true,
    ...overrides,
  };
}

/** The scoped CLI grant (officer tools gate spec §5.1) the spec's copy is written against:
 *  Award, Undo and Nudge until 11 Dec 2026, with no billing contact. */
export function scopedGrant(): GuildBillingView {
  return guildBilling({
    scope: 'officer_tools',
    current_period_end: '2026-12-11T00:00:00Z',
    billed_by: '',
    you_are_billing_contact: false,
  });
}

/** The `/v1/me` envelope for an officer of SETTINGS_GUILD with the given plan and capabilities. */
export function meWithGuildPlan(plan: GuildBillingView | null, can: readonly string[] = CLAIM_HOLDER_CAN) {
  return {
    ok: true,
    data: {
      user: { id: 7, battletag: 'Fixture#1234', email: null, role: 'user', anonymize: false },
      characters: [],
      guilds: [{ ...SETTINGS_GUILD, rank: 'officer', verified: true, can, plan }],
      entitlements: mockEntitlements(),
    },
    error: null,
    request_id: 'r',
  };
}

/** GET /v1/guilds/{id}/settings for a claimed guild. */
const GUILD_SETTINGS = {
  default_visibility: 'guild',
  officer_max_rank_index: 1,
  claimed_by: { battletag: 'Fixture#1234' },
  claim_pending: null,
  claim: { state: 'claimed', since: '2026-09-01T00:00:00Z', frozen: false },
  invite: { rotated_at: null },
};

/** Routes the public guild page and the settings read that SETTINGS_GUILD_PATH makes. */
export async function routeGuildSettingsPage(page: Page): Promise<void> {
  const { region, ruleset } = SETTINGS_GUILD;
  await page.route(`**/v1/guilds/${region}/${ruleset}/the-last-watch`, (route) =>
    route.fulfill(envelope({ guild: SETTINGS_GUILD, progression: [], roster_best: [], reports: [] }, 200)),
  );
  await page.route(`**/v1/guilds/${SETTINGS_GUILD.id}/settings`, (route) =>
    route.fulfill(envelope(GUILD_SETTINGS, 200)),
  );
}
