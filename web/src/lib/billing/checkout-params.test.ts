// web/src/lib/billing/checkout-params.test.ts
import { describe, expect, it } from 'vitest';
import type { MeGuild } from '../account/api';
import { billableGuilds, guildFitsRaidTier, parseCheckoutParams } from './checkout-params';

describe('parseCheckoutParams', () => {
  it('accepts a premium request, with or without a comeback id', () => {
    expect(parseCheckoutParams('?plan=premium&interval=yearly')).toEqual({
      plan: 'premium',
      interval: 'yearly',
      guildId: undefined,
      comebackId: undefined,
      status: null,
    });
    expect(parseCheckoutParams('?plan=premium&interval=monthly&comeback=12')?.comebackId).toBe(12);
  });

  it('accepts both guild plans, with or without a guild id', () => {
    expect(parseCheckoutParams('?plan=guild&interval=yearly&guild_id=9')?.guildId).toBe(9);
    expect(parseCheckoutParams('?plan=raid_guild&interval=monthly')?.guildId).toBeUndefined();
  });

  it.each([
    '?plan=gold&interval=yearly',
    '?plan=premium',
    '?plan=guild&interval=weekly',
    '?plan=guild&interval=yearly&guild_id=-1',
    '?plan=guild&interval=yearly&guild_id=0',
    '?plan=guild&interval=yearly&guild_id=1.5',
    '?plan=guild&interval=yearly&comeback=3',
    '?plan=premium&interval=yearly&comeback=0',
    '?plan=premium&interval=yearly&guild_id=4',
  ])('rejects %s', (search) => {
    expect(parseCheckoutParams(search)).toBeNull();
  });

  it('reads the success return', () => {
    expect(parseCheckoutParams('?plan=premium&interval=yearly&status=success')?.status).toBe('success');
  });
});

describe('billableGuilds and the tier fit', () => {
  const base = { region: 'us', ruleset: 'x', verified: true, verified_count: 3 };
  const guilds: MeGuild[] = [
    { ...base, id: 1, name: 'A', can: ['billing'], plan: null },
    { ...base, id: 2, name: 'B', can: ['approve'] },
    { ...base, id: 3, name: 'C', can: ['billing'], plan: { tier: 'guild' } as MeGuild['plan'] },
  ];

  it('keeps only guilds that can bill and have no plan', () => {
    expect(billableGuilds(guilds).map((guild) => guild.id)).toEqual([1]);
  });

  it('switches to the raid tier above 25 members', () => {
    expect(guildFitsRaidTier(25)).toBe(false);
    expect(guildFitsRaidTier(28)).toBe(true);
  });
});
