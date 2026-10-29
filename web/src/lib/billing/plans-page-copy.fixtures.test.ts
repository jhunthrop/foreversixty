// web/src/lib/billing/plans-page-copy.fixtures.test.ts
// B8 check 2: swap every source figure and the copy must move with it. A figure typed into a
// string instead of imported fails here.
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

type Copy = typeof import('./plans-page-copy').plansPageCopy;
type Model = typeof import('../../components/premium/model');

function ruleSentence(c: Copy, term: string): string {
  const rule = c.rules.find((r) => r.term === term);
  if (!rule) throw new Error(`no rule named ${term}`);
  return rule.sentence;
}

function expectGuildFiguresFollow(c: Copy, perMemberLine: Copy['price']['perMember']): void {
  expect(c.guild.audience).toContain('30 verified members');
  expect(c.guild.features[3]).toMatch(/^2 Premium seat/);
  expect(c.guild.features[4]).toContain('Up to 30 verified members');
  expect(c.guild.features[4]).toContain('5- or 15-person');
  expect(perMemberLine('yearly', true)).toBe('$3 per member a year at 30 members.');
  expect(c.trial).toContain('First 9 days free');
  expect(c.trial).toContain('Cancel before day 10');
  expect(ruleSentence(c, 'Trial')).toContain('first 9 days');
  expect(ruleSentence(c, 'Trial')).toContain('Cancel before day 10');
  expect(ruleSentence(c, 'Size')).toContain('Above 30');
  expect(ruleSentence(c, 'Moving a seat')).toContain('once every 45 days');
}

function expectModelPricesFollow(model: Model): void {
  expect(model.priceCell('guild', 'yearly', true).amount).toBe('$90');
  expect(model.priceCell('raid_guild', 'yearly', true).amount).toBe('$300');
  expect(model.buttonLabel('raid_guild', 'yearly', true)).toBe('Subscribe a guild · $300 a year');
  expect(model.buttonLabel('guild', 'monthly', true)).toBe('Subscribe a guild · $9 a month');
  expect(model.trialParts().lead).toBe('First 9 days free');
}

describe('plans-page figures follow their sources', () => {
  beforeEach(() => vi.resetModules());
  afterEach(() => {
    vi.doUnmock('../sim/precision');
    vi.doUnmock('./plans');
    vi.resetModules();
  });

  it('every string follows swapped figures', async () => {
    vi.doMock('../sim/precision', async (importOriginal) => ({
      ...(await importOriginal<typeof import('../sim/precision')>()),
      CAPS: { browser: 333, server: 7777 },
      LANE_ITERATION_CEILING: { browser: 11111, server: 222222 },
    }));
    // The cost helpers must read the imported PLANS / RAID_SIZES values rather than plans.ts's own
    // helpers bound to the real table, so swapping the values here moves every string.
    vi.doMock('./plans', async (importOriginal) => {
      const real = await importOriginal<typeof import('./plans')>();
      return {
        ...real,
        RAID_SIZES: [30, 15, 5] as const,
        TRIAL_REMINDER_DAYS: 4,
        GUILD_TIER_MEMBER_LIMIT: 30,
        GUILD_TRIAL_DAYS: 9,
        SEAT_MOVE_LOCK_DAYS: 45,
        PLANS: {
          ...real.PLANS,
          guild: { ...real.PLANS.guild, founderYearly: 90, founderMonthly: 9, seats: 2 },
          raid_guild: { ...real.PLANS.raid_guild, founderYearly: 300, seats: 4 },
        },
      };
    });
    const { plansPageCopy: c, perRaiderLine, perMemberLine } = await import('./plans-page-copy');
    expect(c.premium.features[1]?.text).toContain('7,777');
    expect(c.premium.features[1]?.text).toContain('333');
    expect(c.premium.features[3]?.text).toContain('222,222');
    expect(c.premium.features[3]?.text).toContain('11,111');
    expect(c.trialEmail).toContain('4 days');
    expect(perRaiderLine('yearly', true)).toBe(
      '$10 per raider a year in a 30-person raid, $20 in a 15, $60 in a 5.',
    );
    expectGuildFiguresFollow(c, perMemberLine);
    const model = await import('../../components/premium/model');
    expect(model.costLine('raid_guild', 'yearly', true)).toContain('30-person');
    expect(model.costLine('guild', 'yearly', true)).toBe('$3 per member a year at 30 members.');
    expectModelPricesFollow(model);
  });
});
