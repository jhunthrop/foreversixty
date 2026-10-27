// web/src/test-support/billing-fixtures.ts
// e2e stubs for the billing routes, following account/api.ts's requestEnvelope envelope
// shape and the sim lane's existing sim-api.ts pattern of one small helper per route rather
// than each spec hand-rolling page.route calls.
import type { Page } from '@playwright/test';
import type { EntitlementsView, GuildBillingView } from '../lib/account/api';
import type {
  BillingFacts,
  GuildBillingSettings,
  GuildPlanView,
  MeEndedSeat,
  MeSeat,
  MeSeatOffer,
  SeatCandidate,
  SeatView,
  TransferView,
  UpgradePreview,
} from '../lib/billing/contract';

/** The price facts of a row with no Stripe price behind it (a grant). */
export const NO_PRICE_FACTS: BillingFacts = { founder: false, interval: null, amount_cents: null };

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
    tier: 'guild',
    you_hold_seat: false,
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
 *  Award, Undo and Nudge until 12 Dec 2026, with no billing contact. */
export function scopedGrant(): GuildBillingView {
  return guildBilling({
    scope: 'officer_tools',
    current_period_end: '2026-12-12T00:00:00Z',
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

// ---------------------------------------------------------------------------------------------
// Launch pricing: one fixture per spec section 9 e2e state. Each is the wire shape the API
// sends, so a spec stubs a route with it and a unit test feeds it to a view function.
// ---------------------------------------------------------------------------------------------

/** The guild billing read for a guild with no plan yet (§9 case 9). */
export function guildBillingRead(overrides: Partial<GuildBillingSettings> = {}): GuildBillingSettings {
  return {
    billing: null,
    trial_used: false,
    verified_count: 12,
    seats: [],
    transfer: null,
    transfer_candidates: [],
    can: ['billing', 'seats', 'transfer'],
    ...overrides,
  };
}

/** A paid Guild-tier plan, list price, monthly. */
export function planView(overrides: Partial<GuildPlanView> = {}): GuildPlanView {
  return {
    scope: 'plan',
    tier: 'guild',
    status: 'active',
    current_period_end: '2026-12-01T00:00:00Z',
    cancel_at_period_end: false,
    trial_end: null,
    founder: false,
    interval: 'monthly',
    amount_cents: 1500,
    billed_by: 'Fixture#1234',
    you_are_billing_contact: true,
    ...overrides,
  };
}

export function emptySeat(slot: number, overrides: Partial<SeatView> = {}): SeatView {
  return {
    slot,
    state: 'empty',
    holder: null,
    given_at: null,
    starts_at: null,
    movable_at: null,
    returned: null,
    ...overrides,
  };
}

/** §9 case 12: a seat held, locked for moves until 2 Dec 2026. */
export function heldSeat(slot = 1, overrides: Partial<SeatView> = {}): SeatView {
  return emptySeat(slot, {
    state: 'held',
    holder: { user_id: 31, character_key: 'us-hardcore-kraggor', name: 'Kraggor', class: 'warrior' },
    given_at: '2026-11-02T00:00:00Z',
    starts_at: '2026-11-02T00:00:00Z',
    movable_at: '2026-12-02T00:00:00Z',
    ...overrides,
  });
}

/** §9 case 12: a seat offered to a member who pays for their own Premium (no lock yet). */
export function offeredSeat(slot = 1): SeatView {
  return heldSeat(slot, { state: 'offered', starts_at: null, movable_at: null });
}

/** §9 case 12: a seat whose holder left the guild. */
export function returnedSeat(slot = 1): SeatView {
  return emptySeat(slot, { returned: { name: 'Kraggor', at: '2026-11-20T00:00:00Z' } });
}

export const SEAT_CANDIDATES: readonly SeatCandidate[] = [
  { user_id: 31, name: 'Kraggor', class: 'warrior', pays_premium: false },
  { user_id: 32, name: 'Mirelle', class: 'priest', pays_premium: true },
];

/** §9 case 10: a trial that bills on 4 Nov 2026. */
export const TRIAL_READ: GuildBillingSettings = guildBillingRead({
  billing: planView({ status: 'trialing', trial_end: '2026-11-04T00:00:00Z' }),
});

/** §9 case 11: a Guild-tier guild with 28 verified members, so the size line offers Raid guild. */
export const GUILD_TIER_28_READ: GuildBillingSettings = guildBillingRead({
  billing: planView(),
  verified_count: 28,
  seats: [emptySeat(1)],
});

/** A founder Raid guild on the yearly price. */
export const FOUNDER_RAID_GUILD_READ: GuildBillingSettings = guildBillingRead({
  billing: planView({ tier: 'raid_guild', founder: true, interval: 'yearly', amount_cents: 20000 }),
  seats: [emptySeat(1), emptySeat(2), emptySeat(3)],
});

/** §9 case 11's confirm block: what the upgrade charges now and from the next renewal. */
export const UPGRADE_PREVIEW: UpgradePreview = {
  amount_due_cents: 600,
  next_amount_cents: 2500,
  interval: 'monthly',
  renews_at: '2026-12-01T00:00:00Z',
  proration_date: 1792000000,
};

/** §9 case 13: the sender's view while the offer waits, the receiver's view, and a decline. */
export const TRANSFER_OFFERED: TransferView = {
  to_name: 'Mirelle',
  to_you: false,
  from_name: 'Fixture#1234',
  at: '2026-11-10T00:00:00Z',
  state: 'offered',
  starts_at: '2026-12-01T00:00:00Z',
  amount_cents: 1500,
  interval: 'monthly',
  founder: false,
};
export const TRANSFER_RECEIVED: TransferView = { ...TRANSFER_OFFERED, to_you: true };
export const TRANSFER_DECLINED: TransferView = { ...TRANSFER_OFFERED, state: 'declined' };

/** §9 case 14: a lapsed plan, read-only apart from `Subscribe again` and the CSV downloads. */
export const LAPSED_READ: GuildBillingSettings = guildBillingRead({
  billing: planView({ status: 'canceled', current_period_end: '2026-10-01T00:00:00Z' }),
  can: [],
});

/** §9 case 15: a read-only officer sees the plan, `Billed by`, and seats with no buttons. */
export const READ_ONLY_OFFICER_READ: GuildBillingSettings = guildBillingRead({
  billing: planView({ you_are_billing_contact: false }),
  seats: [heldSeat(1)],
  can: [],
});

/** §9 case 17: a seat offered to an account that pays its own Premium until 20 Dec 2026. */
export const ME_SEAT_OFFER: MeSeatOffer = {
  guild_id: SETTINGS_GUILD.id,
  guild_name: SETTINGS_GUILD.name,
  from_name: 'Fixture#1234',
  offered_at: '2026-11-10T00:00:00Z',
  own_ends_at: '2026-12-20T00:00:00Z',
  own_amount_cents: 600,
  own_interval: 'monthly',
};

/** §9 case 17: an account that holds a seat from 20 Dec 2026. */
export const ME_SEAT: MeSeat = {
  guild_name: SETTINGS_GUILD.name,
  given_by: 'Fixture#1234',
  since: '2026-11-10T00:00:00Z',
  starts_at: '2026-12-20T00:00:00Z',
};

/** §9 case 17: a seat that ended, with its resubscribe link. */
export const ME_ENDED_SEAT: MeEndedSeat = {
  guild_name: SETTINGS_GUILD.name,
  ended_at: '2026-11-25T00:00:00Z',
  resubscribe_url: '/premium/checkout?plan=premium&interval=monthly&comeback=9',
  amount_cents: 600,
  interval: 'monthly',
};

/** §9 case 17: a founder's `entitlements.billing` price facts. */
export const FOUNDER_PREMIUM_FACTS: BillingFacts = { founder: true, interval: 'yearly', amount_cents: 4000 };

/** GET /v1/guilds/{id}/billing, stubbed. */
export async function stubGuildBilling(
  page: Page,
  guildId: number,
  response: { status: number; data?: unknown },
): Promise<void> {
  await page.route(`**/v1/guilds/${guildId}/billing`, (route) =>
    route.fulfill(envelope(response.data ?? null, response.status)),
  );
}
