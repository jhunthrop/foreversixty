// web/src/lib/billing/normalize.test.ts
import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import {
  emptySeat,
  guildBillingRead,
  ME_ENDED_SEAT,
  ME_SEAT_OFFER,
  mockEntitlements,
  planView,
  TRANSFER_OFFERED,
} from '../../test-support/billing-fixtures';
import type { GuildBillingSettings, SeatView } from './contract';
import { normalizeGuildBilling, normalizeMeBilling } from './normalize';

/** What the API sends for a row with no price: empty strings where the types say null. */
const blank = { interval: '', amount_cents: '' } as unknown as { interval: null; amount_cents: null };

describe('normalizeGuildBilling', () => {
  it('turns an empty interval and amount into null on the plan and the transfer', () => {
    const read = guildBillingRead({
      billing: planView(blank),
      transfer: { ...TRANSFER_OFFERED, ...blank },
    });
    const out = normalizeGuildBilling(read);
    expect(out.billing).toMatchObject({ interval: null, amount_cents: null });
    expect(out.transfer).toMatchObject({ interval: null, amount_cents: null });
  });

  it('keeps real price facts and a null plan or transfer', () => {
    const priced = guildBillingRead({ billing: planView({ interval: 'yearly', amount_cents: 12000 }) });
    expect(normalizeGuildBilling(priced).billing).toMatchObject({ interval: 'yearly', amount_cents: 12000 });
    const bare: GuildBillingSettings = guildBillingRead();
    expect(normalizeGuildBilling(bare)).toEqual(bare);
  });

  it('does not mutate its input', () => {
    const read = guildBillingRead({ billing: planView(blank) });
    normalizeGuildBilling(read);
    expect(read.billing?.interval).toBe('');
  });
});

describe('normalizeMeBilling', () => {
  const me = (extra: Partial<Me>): Me => ({
    user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
    characters: [],
    guilds: [],
    ...extra,
  });

  it('normalizes the premium billing, the offer and the ended seat', () => {
    const out = normalizeMeBilling(
      me({
        entitlements: mockEntitlements({
          billing: {
            plan: 'premium',
            status: 'active',
            current_period_end: null,
            cancel_at_period_end: false,
            founder: false,
            ...blank,
          },
        }),
        seat_offer: { ...ME_SEAT_OFFER, own_interval: '' as never, own_amount_cents: '' as never },
        ended_seat: { ...ME_ENDED_SEAT, ...blank },
      }),
    );
    expect(out?.entitlements?.billing).toMatchObject({ interval: null, amount_cents: null });
    expect(out?.seat_offer).toMatchObject({ own_interval: null, own_amount_cents: null });
    expect(out?.ended_seat).toMatchObject({ interval: null, amount_cents: null });
  });

  it('leaves absent and null facts as they were', () => {
    const out = normalizeMeBilling(me({ seat_offer: null }));
    expect(out?.seat_offer).toBeNull();
    expect(out !== null && 'ended_seat' in out).toBe(false);
    expect(out?.entitlements).toBeUndefined();
  });
});

describe('normalizeGuildBilling seats', () => {
  it('reads an absent offeree as null and keeps a present one', () => {
    const offeree = { name: 'Mirelle', class: 'priest', offered_at: '2026-11-10T12:00:00Z' };
    const absent = { ...emptySeat(1), offeree: undefined } as unknown as SeatView;
    const read = guildBillingRead({ seats: [absent, emptySeat(2, { offeree })] });
    expect(normalizeGuildBilling(read).seats.map((seat) => seat.offeree)).toEqual([null, offeree]);
  });
});
