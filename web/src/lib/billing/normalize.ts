// web/src/lib/billing/normalize.ts
// The read boundary for price facts. The API derives `interval` from the price lookup key, so
// a row with no Stripe price (a grant) can send "" where the types say null; normalizing here
// keeps every reader's `interval === null` check honest and stops a price line printing "a year".
import type { Me } from '../account/me-types';
import { GUILD_FACTIONS, type MeGuild } from '../account/me-types';
import {
  CLAIM_STATES,
  COVERED_STATES,
  SUBSCRIBE_STEPS,
  type GuildBillingSettings,
  type GuildClaimView,
  type GuildCoveredView,
  type GuildSubscribeView,
  type MeEndedSeat,
  type MeSeatOffer,
  type TransferView,
} from './contract';
import type { Interval } from './plans';

function intervalOrNull(value: unknown): Interval | null {
  return value === 'monthly' || value === 'yearly' ? value : null;
}

function centsOrNull(value: unknown): number | null {
  return typeof value === 'number' ? value : null;
}

function normalizeTransfer(transfer: TransferView | null): TransferView | null {
  if (transfer === null) return null;
  return {
    ...transfer,
    interval: intervalOrNull(transfer.interval),
    amount_cents: centsOrNull(transfer.amount_cents),
  };
}

export function normalizeGuildBilling(read: GuildBillingSettings): GuildBillingSettings {
  return {
    ...read,
    billing:
      read.billing === null
        ? null
        : {
            ...read.billing,
            interval: intervalOrNull(read.billing.interval),
            amount_cents: centsOrNull(read.billing.amount_cents),
          },
    seats: read.seats.map((seat) => ({ ...seat, offeree: seat.offeree ?? null })),
    transfer: normalizeTransfer(read.transfer),
  };
}

function normalizeOffer(offer: MeSeatOffer | null | undefined): MeSeatOffer | null | undefined {
  if (!offer) return offer;
  return {
    ...offer,
    own_interval: intervalOrNull(offer.own_interval),
    own_amount_cents: centsOrNull(offer.own_amount_cents),
  };
}

function normalizeEnded(ended: MeEndedSeat | null | undefined): MeEndedSeat | null | undefined {
  if (!ended) return ended;
  return {
    ...ended,
    interval: intervalOrNull(ended.interval),
    amount_cents: centsOrNull(ended.amount_cents),
  };
}

/** `GET /v1/me` with its billing price facts, seat offer and ended seat normalized. */
export function normalizeMeBilling(me: Me | null): Me | null {
  if (me === null) return null;
  const billing = me.entitlements?.billing;
  return {
    ...me,
    ...(Array.isArray(me.guilds) ? { guilds: me.guilds.map(normalizeMeGuild) } : {}),
    ...(me.entitlements === undefined || !billing
      ? {}
      : {
          entitlements: {
            ...me.entitlements,
            billing: {
              ...billing,
              interval: intervalOrNull(billing.interval),
              amount_cents: centsOrNull(billing.amount_cents),
            },
          },
        }),
    ...(me.seat_offer === undefined ? {} : { seat_offer: normalizeOffer(me.seat_offer) }),
    ...(me.ended_seat === undefined ? {} : { ended_seat: normalizeEnded(me.ended_seat) }),
  };
}

type Wire = Record<string, unknown>;

function isWire(value: unknown): value is Wire {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

function oneOf<T extends string>(words: readonly T[], value: unknown): T | null {
  return words.find((word) => word === value) ?? null;
}

function stringOrNull(value: unknown): string | null {
  return typeof value === 'string' ? value : null;
}

function normalizeSubscribe(value: unknown): GuildSubscribeView | null {
  if (!isWire(value)) return null;
  const step = oneOf(SUBSCRIBE_STEPS, value.step);
  if (step === null) return null;
  return {
    step,
    trial_used: typeof value.trial_used === 'boolean' ? value.trial_used : null,
    available_at: stringOrNull(value.available_at),
    reason: oneOf(['holds_other', 'rate_limited'] as const, value.reason),
    other_guild_name: stringOrNull(value.other_guild_name),
  };
}

function normalizeClaim(value: unknown): GuildClaimView | null {
  if (!isWire(value)) return null;
  const state = oneOf(CLAIM_STATES, value.state);
  if (state === null) return null;
  return {
    state,
    yours: value.yours === true,
    pending_yours: value.pending_yours === true,
    holder_name: stringOrNull(value.holder_name),
  };
}

function normalizeCovered(value: unknown): GuildCoveredView | null {
  if (!isWire(value)) return null;
  const state = oneOf(COVERED_STATES, value.state);
  const tier = oneOf(['guild', 'raid_guild'] as const, value.tier);
  if (state === null || tier === null) return null;
  return {
    tier,
    state,
    ends_at: stringOrNull(value.ends_at),
    billed_by: typeof value.billed_by === 'string' ? value.billed_by : '',
    you_pay: value.you_pay === true,
    you_hold_seat: value.you_hold_seat === true,
    seats_given: typeof value.seats_given === 'number' ? value.seats_given : 0,
    seats_total: typeof value.seats_total === 'number' ? value.seats_total : 0,
  };
}

/**
 * One `/v1/me` guild entry with its standing read defensively. A field the API did not send
 * (an older response) or a value outside the contract becomes null, which the web draws as
 * "unknown"; an absent `slug` is "" (no hand-off link) and an empty `ruleset_label` is absent.
 */
export function normalizeMeGuild(guild: unknown): MeGuild {
  const { ruleset_label: label, ...wire } = (isWire(guild) ? guild : {}) as Wire & MeGuild;
  const rosterCount = wire.roster_count;
  return {
    ...wire,
    ...(typeof label === 'string' && label !== '' ? { ruleset_label: label } : {}),
    slug: stringOrNull(wire.slug) ?? '',
    crest_url: stringOrNull(wire.crest_url),
    faction: oneOf(GUILD_FACTIONS, wire.faction),
    roster_count:
      typeof rosterCount === 'number' && Number.isInteger(rosterCount) && rosterCount >= 0
        ? rosterCount
        : null,
    claim: normalizeClaim(wire.claim),
    verify: isWire(wire.verify) ? { has_approver: wire.verify.has_approver === true } : null,
    covered: normalizeCovered(wire.covered),
    subscribe: normalizeSubscribe(wire.subscribe),
  };
}
