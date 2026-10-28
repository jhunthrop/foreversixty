// web/src/lib/billing/plans-page-copy.fixtures.test.ts
// B8 check 2: swap every source figure and the copy must move with it. A figure typed into a
// string instead of imported fails here.
import { beforeEach, describe, expect, it, vi } from 'vitest';

describe('plans-page figures follow their sources', () => {
  beforeEach(() => vi.resetModules());

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
        PLANS: { ...real.PLANS, raid_guild: { ...real.PLANS.raid_guild, founderYearly: 300 } },
      };
    });
    vi.doMock('../../data/premium-examples.json', () => ({
      default: {
        serverRun: {
          simId: 'swap1',
          characterName: 'Swapname',
          level: 58,
          specLabel: 'Arms Warrior',
          preset: true,
          placeholder: false,
          combinations: 1234,
          ranAt: '2027-02-03T12:00:00Z',
        },
      },
    }));
    const { plansPageCopy: c, serverRunCaption, perRaiderLine } = await import('./plans-page-copy');
    expect(c.premium.features[0]).toContain('7,777');
    expect(c.premium.features[0]).toContain('333');
    expect(c.premium.features[2]).toContain('222,222');
    expect(c.premium.features[2]).toContain('11,111');
    expect(c.trialEmail).toContain('4 days');
    expect(perRaiderLine('yearly', true)).toBe(
      '$10 per raider a year in a 30-person raid, $20 in a 15, $60 in a 5.',
    );
    expect(serverRunCaption()).toContain('level-58 Arms Warrior');
    expect(serverRunCaption()).toContain('1,234');
    expect(serverRunCaption()).toContain('3 Feb 2027');
    const { costLine } = await import('../../components/premium/model');
    expect(costLine('raid_guild', 'yearly', true)).toContain('30-person');
  });
});
