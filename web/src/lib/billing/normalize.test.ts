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
import { normalizeGuildBilling, normalizeMeBilling, normalizeMeGuild } from './normalize';

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

describe('normalizeMeGuild', () => {
  const OLDER = {
    id: 1,
    region: 'us',
    ruleset: 'forever',
    name: 'X',
    verified: false,
    verified_count: 0,
    can: [],
  };
  const SUBSCRIBE = {
    step: 'ready',
    trial_used: false,
    available_at: null,
    reason: null,
    other_guild_name: null,
  };

  it('an older guild entry normalizes to unknown', () => {
    const out = normalizeMeGuild(OLDER);
    expect(out).toMatchObject({
      id: 1,
      name: 'X',
      subscribe: null,
      claim: null,
      verify: null,
      covered: null,
    });
    expect(out.ruleset_label).toBeUndefined();
  });

  it('an unknown step becomes null', () => {
    expect(normalizeMeGuild({ ...OLDER, subscribe: { step: 'pay_now' } }).subscribe).toBeNull();
  });

  it('keeps a known step and its values', () => {
    expect(normalizeMeGuild({ ...OLDER, subscribe: SUBSCRIBE }).subscribe).toEqual(SUBSCRIBE);
    const blocked = {
      ...SUBSCRIBE,
      step: 'claim_blocked',
      reason: 'rate_limited',
      available_at: '2026-11-01T00:00:00Z',
    };
    expect(normalizeMeGuild({ ...OLDER, subscribe: blocked }).subscribe).toEqual(blocked);
  });

  it('treats an empty ruleset_label as absent and keeps a real one', () => {
    expect('ruleset_label' in normalizeMeGuild({ ...OLDER, ruleset_label: '' })).toBe(false);
    expect(normalizeMeGuild({ ...OLDER, ruleset_label: 'Forever' }).ruleset_label).toBe('Forever');
  });

  it('keeps the slug and normalizes an absent one to empty', () => {
    expect(normalizeMeGuild({ ...OLDER, slug: 'olympus' }).slug).toBe('olympus');
    expect(normalizeMeGuild(OLDER).slug).toBe('');
  });

  it('validates faction, crest_url and roster_count', () => {
    expect(
      normalizeMeGuild({ ...OLDER, faction: 'horde', crest_url: '/c.png', roster_count: 12 }),
    ).toMatchObject({
      faction: 'horde',
      crest_url: '/c.png',
      roster_count: 12,
    });
    expect(normalizeMeGuild({ ...OLDER, faction: 'orc', crest_url: 5, roster_count: -1 })).toMatchObject({
      faction: null,
      crest_url: null,
      roster_count: null,
    });
    expect(normalizeMeGuild({ ...OLDER, roster_count: 1.5 }).roster_count).toBeNull();
  });

  it('drops a claim with an unknown state and keeps a known one', () => {
    expect(normalizeMeGuild({ ...OLDER, claim: { state: 'bogus' } }).claim).toBeNull();
    const claim = { state: 'claimed', yours: true, pending_yours: false, holder_name: 'Aria' };
    expect(normalizeMeGuild({ ...OLDER, claim }).claim).toEqual(claim);
  });

  it('keeps a covered view and drops one with an unknown state', () => {
    const covered = {
      tier: 'guild',
      state: 'trial',
      ends_at: null,
      billed_by: '',
      you_pay: false,
      you_hold_seat: false,
      seats_given: 0,
      seats_total: 3,
    };
    expect(normalizeMeGuild({ ...OLDER, covered }).covered).toEqual(covered);
    expect(normalizeMeGuild({ ...OLDER, covered: { ...covered, state: 'x' } }).covered).toBeNull();
  });

  it('is applied to every guild by normalizeMeBilling', () => {
    const me = { user: { id: 1 }, characters: [], guilds: [OLDER] } as unknown as Me;
    expect(normalizeMeBilling(me)?.guilds[0]?.subscribe).toBeNull();
  });
});
