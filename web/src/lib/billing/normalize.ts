// web/src/lib/billing/normalize.ts
// The read boundary for price facts. The API derives `interval` from the price lookup key, so
// a row with no Stripe price (a grant) can send "" where the types say null; normalizing here
// keeps every reader's `interval === null` check honest and stops a price line printing "a year".
import type { Me } from '../account/api';
import type { GuildBillingSettings, MeEndedSeat, MeSeatOffer, TransferView } from './contract';
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
