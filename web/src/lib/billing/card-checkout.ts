// web/src/lib/billing/card-checkout.ts
// The decisions behind the checkout card's pay actions (spec 2026-10-27-subscribe-and-join §5.3,
// §9.1): the ready button's words and fill, which on_plan action a payer gets, what a checkout
// or portal refusal becomes, and how a `guild_id` arrival picks its card. No DOM; the two API
// calls are injected.
import { SECONDARY_BUTTON_FIXED } from '../planner/styles';
import { BillingApiError, type CheckoutRequest, type CheckoutResult, type PortalResult } from './api';
import { checkoutGuildCopy as copy, type GuildTier } from './checkout-guild-copy';
import type { SubscribeStep } from './contract';
import type { StepGuild } from './subscribe-step';

const HTTP_NOT_OPEN = 503;
const HTTP_FORBIDDEN = 403;
const HTTP_CONFLICT = 409;
const HTTP_NOT_FOUND = 404;

/** The settings sub-path and query that open the Guild-to-Raid-guild confirm (launch pricing §3.2.4). */
const UPGRADE_SUFFIX = '/settings?upgrade=1#billing';

const FULL_WIDTH_BUTTON = 'w-full justify-center px-4';
/** The gold-outline button of a secondary action (spec §5.3). */
export const OUTLINE_GOLD_BUTTON = `${SECONDARY_BUTTON_FIXED} border-gold text-gold ${FULL_WIDTH_BUTTON}`;
export const PLAIN_BUTTON = `${SECONDARY_BUTTON_FIXED} ${FULL_WIDTH_BUTTON}`;

export interface ReadyGuild {
  subscribe?: { step: SubscribeStep | null; trial_used?: boolean | null } | null;
  covered?: { state: string } | null;
}

/** Spec §9.1: trial unused names the trial, a lapsed plan says again, else the price. */
export function readyLabel(guild: ReadyGuild, price: string): string {
  if (guild.subscribe?.trial_used === false) return copy.ready.startTrial;
  if (guild.covered?.state === 'lapsed') return copy.ready.subscribeAgain(price);
  return copy.ready.subscribe(price);
}

export function readyCount(guilds: readonly StepGuild[]): number {
  return guilds.filter((guild) => guild.subscribe?.step === 'ready').length;
}

/** Spec §10: the page's one fill; two or more ready cards all draw the gold outline. */
export function isReadyFill(readyCards: number): boolean {
  return readyCards === 1;
}

/** What a card is waiting on; another card's wait only disables this card's buttons. */
export type PayBusy = 'checkout' | 'portal' | null;

export type PayAction = 'manage' | 'move';

export interface PayGuild {
  covered?: { you_pay: boolean; tier: GuildTier } | null;
}

/** Spec §5.3: what the viewer who pays for an `on_plan` guild is offered; null when they do not pay. */
export function payAction(guild: PayGuild, chosen: GuildTier): PayAction | null {
  const covered = guild.covered;
  if (covered?.you_pay !== true) return null;
  return covered.tier === 'guild' && chosen === 'raid_guild' ? 'move' : 'manage';
}

export function upgradeHref(guildPageHref: string): string {
  return `${guildPageHref}${UPGRADE_SUFFIX}`;
}

export type CardOutcome =
  | { kind: 'redirect'; url: string }
  /** `line` null: say nothing (a 403 shows the card's new step). */
  | { kind: 'refused'; line: string | null; rereads: boolean; retry: boolean };

export interface CardIo {
  startCheckout: (input: CheckoutRequest) => Promise<CheckoutResult>;
  openPortal: (guildId: number | undefined) => Promise<PortalResult>;
}

export type RefusalKind = 'not_open' | 'not_found' | 'conflict' | 'forbidden' | 'generic';

/** The one status ladder of a checkout refusal; the card and the launcher each word it their own way. */
export function refusalKind(error: unknown): RefusalKind {
  const status = error instanceof BillingApiError ? error.status : 0;
  if (status === HTTP_NOT_OPEN) return 'not_open';
  if (status === HTTP_NOT_FOUND) return 'not_found';
  if (status === HTTP_CONFLICT) return 'conflict';
  if (status === HTTP_FORBIDDEN) return 'forbidden';
  return 'generic';
}

function refusedCheckout(error: unknown, guildName: string): CardOutcome {
  switch (refusalKind(error)) {
    case 'not_open':
      return { kind: 'refused', line: copy.checkoutErrors.notOpen, rereads: false, retry: false };
    case 'conflict':
      return {
        kind: 'refused',
        line: copy.checkoutErrors.alreadyOnPlan(guildName),
        rereads: true,
        retry: false,
      };
    case 'forbidden':
      return { kind: 'refused', line: null, rereads: true, retry: false };
    default:
      return { kind: 'refused', line: copy.checkoutErrors.generic, rereads: false, retry: true };
  }
}

/** A page restored from the back/forward cache (Stripe's back button) keeps its old busy state. */
export function restoredFromCache(event: { persisted: boolean }): boolean {
  return event.persisted;
}

export async function runCardCheckout(
  io: Pick<CardIo, 'startCheckout'>,
  request: CheckoutRequest,
  guildName: string,
): Promise<CardOutcome> {
  try {
    const result = await io.startCheckout(request);
    return { kind: 'redirect', url: result.checkout_url };
  } catch (error) {
    return refusedCheckout(error, guildName);
  }
}

export async function runCardPortal(io: Pick<CardIo, 'openPortal'>, guildId: number): Promise<CardOutcome> {
  try {
    const result = await io.openPortal(guildId);
    return { kind: 'redirect', url: result.portal_url };
  } catch {
    return { kind: 'refused', line: copy.portalError, rereads: false, retry: false };
  }
}

export interface Arrival {
  /** The guild whose card goes first and scrolls into view. */
  firstId?: number;
  /** The `ready` guild whose checkout starts on arrival (today's direct launch). */
  launchId?: number;
}

/** Spec §5.6: a `guild_id` arrival launches a ready guild, floats any other, ignores an unknown id. */
export function guildArrival(
  guilds: readonly Pick<StepGuild, 'id' | 'subscribe'>[],
  guildId: number | undefined,
): Arrival {
  const guild = guildId === undefined ? undefined : guilds.find((entry) => entry.id === guildId);
  if (guild === undefined) return {};
  return guild.subscribe?.step === 'ready' ? { launchId: guild.id } : { firstId: guild.id };
}
