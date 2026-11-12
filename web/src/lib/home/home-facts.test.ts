// Home's density facts (lib/home/home-facts.ts, density pass 2026-11-09).
import { describe, expect, it } from 'vitest';
import {
  guildSummaryFacts,
  moverDelta,
  moversTitle,
  simGap,
  simShare,
  standingFacts,
  upgradesFacts,
} from './home-facts';
import type { GuildSimRow } from './home-answer';
import type { UpgradesResult } from './upgrades';

const STANDING = {
  spec: 'Fury',
  class: 'Warrior',
  same_spec_count: 4,
  rank_by_item_level: 2,
  item_level: 61,
  needs_before_next_raid: [],
};

const sim = (value: number, you = false): GuildSimRow => ({
  rank: 1,
  name: 'Kaelra',
  class: 'mage',
  key: 'us/normal/kaelra',
  value,
  delta: null,
  simmed_at: '2026-11-06T08:00:00Z',
  you,
});

describe('standingFacts', () => {
  it('reads your item level place among your spec and your raids attended', () => {
    expect(
      standingFacts(
        STANDING,
        { id: 2, name: 'G', verified: true, attendance: { present: 3, nights: 4 } },
        'raids',
      ),
    ).toEqual([
      { n: '#2', text: 'of 4 guild Fury Warriors by item level' },
      { n: '3', text: 'of 4 raids attended' },
    ]);
  });

  it('counts nights before the guild logs raids, and leaves out a place with no peers', () => {
    const guild = { id: 2, name: 'G', verified: true, attendance: { present: 1, nights: 2 } };
    expect(standingFacts({ ...STANDING, same_spec_count: 1 }, guild, undefined)).toEqual([
      { n: '1', text: 'of 2 raid nights attended' },
    ]);
    expect(standingFacts(null, undefined, undefined)).toEqual([]);
  });
});

describe('guildSummaryFacts', () => {
  const summary = {
    raiders: 14,
    waiting_for_approval: 0,
    below_rating_floor: 0,
    tier: { name: 'T1', raids: [], named_encounters: 9, down: 3, pulls_total: 40 },
    named_encounters_down: 3,
    pulls_this_tier: 40,
    updated_at: '2026-11-06T08:00:00Z',
  };

  it('reads the raiders and, once the tier is pulled, the bosses down', () => {
    expect(guildSummaryFacts(summary)).toEqual([
      { n: '14', text: 'raiders' },
      { n: '3', text: 'of 9 bosses down' },
    ]);
    expect(guildSummaryFacts({ ...summary, raiders: 1, tier: { ...summary.tier, pulls_total: 0 } })).toEqual([
      { n: '1', text: 'raider' },
    ]);
    expect(guildSummaryFacts(undefined)).toEqual([]);
  });
});

describe('upgradesFacts', () => {
  const result = (
    total: number,
    count: number,
    bis: number,
    scoreUnit: UpgradesResult['scoreUnit'] = 'dps',
  ): UpgradesResult =>
    ({
      upgrades: Array.from({ length: count }, () => ({})),
      alreadyBis: Array.from({ length: bis }, () => ({})),
      totalGainDps: total,
      scoreUnit,
      notSimCheckedCount: 0,
    }) as unknown as UpgradesResult;

  it('reads the whole gap and the slots already best in slot', () => {
    expect(upgradesFacts(result(52.08, 12, 5))).toEqual([
      { n: '+52.1', text: 'DPS' },
      { n: '5', text: 'best in slot' },
    ]);
    expect(upgradesFacts(result(0, 1, 16))).toEqual([{ n: '16', text: 'best in slot' }]);
    expect(upgradesFacts(result(7.8, 14, 0))).toEqual([{ n: '+7.8', text: 'DPS' }]);
    expect(upgradesFacts(result(0, 0, 17))).toEqual([]);
    expect(upgradesFacts(undefined)).toEqual([]);
  });

  it('counts the gap in the band unit: a tank score, a healer HPS', () => {
    expect(upgradesFacts(result(3895.7, 5, 1, 'tank_score'))[0]).toEqual({ n: '+3,895.7', text: 'score' });
    expect(upgradesFacts(result(41.2, 5, 1, 'hps'))[0]).toEqual({ n: '+41.2', text: 'HPS' });
  });
});

describe('movers', () => {
  it('titles the board climbers once a baseline exists, else the top of the spec', () => {
    expect(moversTitle({ spec: 'warrior-fury', band: 60, baseline: true, rows: [] })).toBe(
      'Climbing this week',
    );
    expect(moversTitle({ spec: 'warrior-fury', band: 60, baseline: false, rows: [] })).toBe('Top this week');
  });

  it('shows a climb only', () => {
    const row = { rank: 3, name: 'A', class: 'hunter', key: 'k', delta: 6 };
    expect(moverDelta(row)).toBe('+6');
    expect(moverDelta({ ...row, delta: 0 })).toBeNull();
    expect(moverDelta({ ...row, delta: null })).toBeNull();
  });
});

describe('guild sims', () => {
  it('reads each row against your sim and the top of the ladder', () => {
    expect(simGap(sim(1012.3), sim(948.6, true), 'DPS')).toBe('63.7 DPS ahead of you');
    expect(simGap(sim(940.0), sim(948.6, true), 'DPS')).toBe('8.6 DPS behind you');
    expect(simGap(sim(948.6, true), sim(948.6, true), 'DPS')).toBeNull();
    expect(simGap(sim(1012.3), null, 'DPS')).toBeNull();
    expect(simShare(sim(506.15), 1012.3)).toBeCloseTo(0.5);
    expect(simShare(sim(1), 0)).toBe(0);
  });
});
