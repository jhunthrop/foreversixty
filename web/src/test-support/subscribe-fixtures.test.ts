// The fixtures resolve to the premium source and guild shapes their names promise, and cannot be
// mutated by a test (the shared objects would leak into every other fixture).
import { describe, expect, it } from 'vitest';
import type { Me } from '../lib/account/api';
import { premiumSource } from '../lib/account/plan-line';
import type { PremiumSource } from '../lib/billing/contract';
import * as fx from './subscribe-fixtures';

const NOW = new Date('2026-11-15T00:00:00Z');

describe('premium source per fixture', () => {
  it.each<[string, Me, PremiumSource]>([
    ['OWNER_MEMBER', fx.OWNER_MEMBER, 'grant'],
    ['OWNER_GUILD_MASTER', fx.OWNER_GUILD_MASTER, 'grant'],
    ['OWNER_GUILD_MASTER_CLAIMED', fx.OWNER_GUILD_MASTER_CLAIMED, 'grant'],
    ['PREMIUM_OWN_FOUNDER', fx.PREMIUM_OWN_FOUNDER, 'own'],
    ['PREMIUM_OWN_ENDING', fx.PREMIUM_OWN_ENDING, 'own'],
    ['PREMIUM_PAST_DUE', fx.PREMIUM_PAST_DUE, 'own'],
    ['PREMIUM_OWN_THEN_SEAT', fx.PREMIUM_OWN_THEN_SEAT, 'own_then_seat'],
    ['PREMIUM_SEAT_WITH_PORTAL', fx.PREMIUM_SEAT_WITH_PORTAL, 'seat'],
    ['PREMIUM_SEAT_NO_PORTAL', fx.PREMIUM_SEAT_NO_PORTAL, 'seat'],
    ['PREMIUM_GRANT_OPEN_ENDED', fx.PREMIUM_GRANT_OPEN_ENDED, 'grant'],
    ['PREMIUM_GRANT_UNTIL', fx.PREMIUM_GRANT_UNTIL, 'grant'],
    ['PREMIUM_FALLBACK', fx.PREMIUM_FALLBACK, 'fallback'],
    ['PREMIUM_FREE', fx.PREMIUM_FREE, 'none'],
  ])('%s', (_name, me, source) => {
    expect(premiumSource(me, NOW)).toBe(source);
  });
});

describe('guild shapes', () => {
  it('only a claimed guild reaches checkout_open, and only a verified officer sees the holder', () => {
    expect(fx.STEP_FIXTURES.checkout_open.guilds[0]?.claim?.state).toBe('claimed');
    for (const me of Object.values(fx.STEP_FIXTURES)) {
      for (const guild of me.guilds) {
        if (!guild.verified) expect(guild.claim?.holder_name ?? null).toBeNull();
      }
    }
  });

  it('a plain ready guild starts the trial, a lapsed one has spent it', () => {
    const plain = fx.STEP_FIXTURES.ready.guilds[0];
    expect(plain?.covered).toBeNull();
    expect(plain?.subscribe?.trial_used).toBe(false);
    const lapsed = fx.READY_LAPSED.guilds[0];
    expect(lapsed?.covered?.state).toBe('lapsed');
    expect(lapsed?.subscribe?.trial_used).toBe(true);
  });
});

describe('immutability', () => {
  it('refuses mutation of a fixture and of what it shares', () => {
    expect(() => {
      (fx.OWNER_MEMBER.characters[0] as { name: string }).name = 'X';
    }).toThrow(TypeError);
    expect(() => {
      (fx.GRANT_OPEN_ENDED as { plan: string }).plan = 'x';
    }).toThrow(TypeError);
  });

  it('builders hand out fresh guilds that do not leak into each other', () => {
    const first = fx.stepGuild();
    const second = fx.stepGuild();
    expect(first.claim).not.toBe(second.claim);
  });
});
