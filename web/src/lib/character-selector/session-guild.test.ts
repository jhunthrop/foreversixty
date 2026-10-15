import { describe, expect, it } from 'vitest';
import type { Me, MeGuild } from '../account/api';
import { sessionGuildOf } from './session-guild';

const guild = (overrides: Partial<MeGuild>): MeGuild => ({
  id: 1,
  region: 'us',
  ruleset: 'pvp',
  name: 'OLYMPUS XXVII',
  verified: true,
  ...overrides,
});

const meWith = (guilds: MeGuild[]): Me => ({ guilds }) as unknown as Me;

describe('sessionGuildOf', () => {
  it('is null signed out, loading, or with no guild', () => {
    expect(sessionGuildOf(null)).toBeNull();
    expect(sessionGuildOf(undefined)).toBeNull();
    expect(sessionGuildOf(meWith([]))).toBeNull();
  });

  it('names the first verified guild and its page', () => {
    const me = meWith([guild({ id: 2, name: 'Unverified', verified: false }), guild({})]);
    expect(sessionGuildOf(me)).toEqual({ name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' });
  });

  it('is null when no guild is verified', () => {
    expect(sessionGuildOf(meWith([guild({ verified: false })]))).toBeNull();
  });
});
