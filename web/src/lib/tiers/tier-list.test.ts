// web/src/lib/tiers/tier-list.test.ts
import { describe, expect, it } from 'vitest';
import { band, catalogEntry, dpsInput, tankInput } from './tier-test-support';
import {
  hrefsFor,
  groupByTier,
  rankRole,
  raceLabel,
  sortMetricOf,
  roleHasTiers,
  tierOf,
  type TierRow,
} from './tier-list';
import { TIER_BANDS } from './tier-rules';

describe('raceLabel', () => {
  it('capitalises each word of a hyphenated race', () => {
    expect(raceLabel('night-elf')).toBe('Night Elf');
    expect(raceLabel('troll')).toBe('Troll');
  });
});

describe('sortMetricOf', () => {
  it('reads damage taken for a tank and set DPS for everything else', () => {
    const tank = tankInput('x', 'y', 'Z', { dtps: 5, effective_health: 1, tps: 1 });
    expect(sortMetricOf(tank.band)).toBe(5);
    expect(sortMetricOf(dpsInput('a', 'b', 'C', 77).band)).toBe(77);
  });
});

describe('rankRole sort keys', () => {
  it('sorts DPS by set DPS, highest first, ties by name', () => {
    const rows = rankRole(
      [
        dpsInput('mage', 'fire', 'Fire', 90),
        dpsInput('warrior', 'fury', 'Fury', 100),
        dpsInput('rogue', 'combat', 'Combat', 90),
      ],
      'dps',
    );
    expect(rows.map((r) => r.name)).toEqual(['Fury', 'Combat', 'Fire']);
    expect(rows.map((r) => r.rank)).toEqual([1, 2, 3]);
  });

  it('sorts tanks by damage taken per second, lowest first, not by effective health', () => {
    const rows = rankRole(
      [
        tankInput('druid', 'feral-bear', 'Feral Bear', { dtps: 330, effective_health: 40000, tps: 700 }),
        tankInput('paladin', 'protection', 'Protection', { dtps: 300, effective_health: 30000, tps: 600 }),
        tankInput('warrior', 'protection', 'Protection', { dtps: 360, effective_health: 50000, tps: 800 }),
      ],
      'tank',
    );
    expect(rows.map((r) => r.classSlug)).toEqual(['paladin', 'druid', 'warrior']);
  });

  it('keeps each role in its own list', () => {
    const inputs = [
      dpsInput('mage', 'fire', 'Fire', 90),
      tankInput('paladin', 'protection', 'Protection', { dtps: 1, effective_health: 1, tps: 1 }),
    ];
    expect(rankRole(inputs, 'dps')).toHaveLength(1);
    expect(rankRole(inputs, 'healer')).toEqual([]);
  });

  it('treats a band without a role as DPS', () => {
    const entry = catalogEntry('mage', 'fire', 'Fire');
    const rows = rankRole([{ entry, band: band(entry, { role: undefined }) }], 'dps');
    expect(rows).toHaveLength(1);
  });

  it('refuses a tank band without metrics', () => {
    const entry = catalogEntry('paladin', 'protection', 'Protection', 'tank');
    expect(() => rankRole([{ entry, band: band(entry, { metrics: null }) }], 'tank')).toThrow(
      /no tank metrics/,
    );
  });
});

describe('gap to the top and bar length', () => {
  it('gaps a DPS row as the share behind the top and sizes its bar to value over top', () => {
    const rows = rankRole([dpsInput('a', 'a', 'A', 200), dpsInput('b', 'b', 'B', 150)], 'dps');
    expect(rows[0]).toMatchObject({ gapPercent: 0, fraction: 1 });
    expect(rows[1]!.gapPercent).toBeCloseTo(25, 10);
    expect(rows[1]!.fraction).toBeCloseTo(0.75, 10);
  });

  it('gaps a tank row as the extra damage taken and sizes its bar to best over value', () => {
    const rows = rankRole(
      [
        tankInput('a', 'a', 'A', { dtps: 300, effective_health: 1, tps: 1 }),
        tankInput('b', 'b', 'B', { dtps: 345, effective_health: 1, tps: 1 }),
      ],
      'tank',
    );
    expect(rows[1]!.gapPercent).toBeCloseTo(15, 10);
    expect(rows[1]!.fraction).toBeCloseTo(300 / 345, 10);
  });

  it('marks the best of the effective health and threat columns independently of the sort', () => {
    const rows = rankRole(
      [
        tankInput('a', 'a', 'A', { dtps: 300, effective_health: 100, tps: 900 }),
        tankInput('b', 'b', 'B', { dtps: 345, effective_health: 200, tps: 500 }),
      ],
      'tank',
    );
    expect(rows.map((r) => r.isBestEffectiveHealth)).toEqual([false, true]);
    expect(rows.map((r) => r.isBestThreat)).toEqual([true, false]);
  });
});

describe('ties', () => {
  const tied = (values: number[]): TierRow['tie'][] =>
    rankRole(
      values.map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    ).map((r) => r.tie);

  it('marks both rows of a tied pair', () => {
    expect(tied([100, 99.5, 80])).toEqual(['below', 'above', null]);
  });

  it('marks a row tied on both sides "both"', () => {
    expect(tied([100, 99.5, 99])).toEqual(['below', 'both', 'above']);
  });

  it('does not mark rows just past the 1% margin', () => {
    expect(tied([100, 98.9])).toEqual([null, null]);
  });

  it('never marks a tank list', () => {
    const rows = rankRole(
      [
        tankInput('a', 'a', 'A', { dtps: 300, effective_health: 1, tps: 1 }),
        tankInput('b', 'b', 'B', { dtps: 300.5, effective_health: 1, tps: 1 }),
      ],
      'tank',
    );
    expect(rows.map((r) => r.tie)).toEqual([null, null]);
  });
});

describe('ties on the published Alliance numbers (rows 6 to 14)', () => {
  const ROWS_6_TO_14 = [799.4, 781.1, 778.2, 757.7, 750.6, 746.6, 700.6, 669.3, 667.9];

  it('marks exactly the rows that have a partner within 1%, never one that has none', () => {
    const rows = rankRole(
      ROWS_6_TO_14.map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    // 781.1 and 778.2 pair; 757.7, 750.6 and 746.6 pair in turn; 669.3 and 667.9 pair.
    // 799.4 (2.3% above 781.1) and 700.6 (4.5% below 746.6 and above 669.3) have no partner.
    expect(rows.map((r) => r.tie)).toEqual([
      null,
      'below',
      'above',
      'below',
      'both',
      'above',
      null,
      'below',
      'above',
    ]);
  });

  it('gives every mark a partner within 1% on the side it names', () => {
    const rows = rankRole(
      ROWS_6_TO_14.map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    const within = (a: number, b: number): boolean => Math.abs(1 - a / b) * 100 <= 1;
    rows.forEach((row, i) => {
      if (row.tie === 'above' || row.tie === 'both')
        expect(within(row.metric, rows[i - 1]!.metric)).toBe(true);
      if (row.tie === 'below' || row.tie === 'both')
        expect(within(rows[i + 1]!.metric, row.metric)).toBe(true);
    });
  });
});

describe('race, confidence and rulers', () => {
  it('names the band race and carries the low-confidence flag', () => {
    const entry = catalogEntry('warlock', 'affliction', 'Affliction');
    const rows = rankRole(
      [{ entry, band: band(entry, { race: 'undead', weights_low_confidence: true }) }],
      'dps',
    );
    expect(rows[0]).toMatchObject({ race: 'Undead', lowConfidence: true });
  });

  it('cuts the DPS list into tiers and draws none that is empty', () => {
    const rows = rankRole(
      [100, 95, 88, 79].map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    expect(rows.map((r) => r.tier)).toEqual(['S', 'A', 'B', 'C']);
    const groups = groupByTier(rows);
    expect(groups.map((g) => [g.letter, g.from, g.to, g.rows.length])).toEqual([
      ['S', 0, 5, 1],
      ['A', 5, 10, 1],
      ['B', 10, 20, 1],
      ['C', 20, 30, 1],
    ]);
    expect(groupByTier(rows.slice(0, 1)).map((g) => g.letter)).toEqual(['S']);
    expect(TIER_BANDS.map((b) => [b.letter, b.from])).toEqual([
      ['S', 0],
      ['A', 5],
      ['B', 10],
      ['C', 20],
      ['D', 30],
    ]);
  });

  it('gives letters to the DPS list only', () => {
    const inputs = [dpsInput('a', 'a', 'A', 100), dpsInput('b', 'b', 'B', 50)];
    expect(rankRole(inputs, 'dps').map((r) => r.tier)).toEqual(['S', 'D']);
    expect(
      rankRole(
        inputs.map((i) => ({ ...i, band: { ...i.band, role: 'healer' as const } })),
        'healer',
      ),
    ).toSatisfy((rows: TierRow[]) => rows.every((r) => r.tier === null));
    expect(roleHasTiers('dps')).toBe(true);
    expect(roleHasTiers('tank')).toBe(false);
    expect(roleHasTiers('healer')).toBe(false);
  });
});

describe('tierOf at the exact cut-offs', () => {
  it.each([
    [0, 'S'],
    [4.99, 'S'],
    [5.0, 'A'],
    [9.99, 'A'],
    [10.0, 'B'],
    [19.99, 'B'],
    [20.0, 'C'],
    [29.99, 'C'],
    [30.0, 'D'],
    [55, 'D'],
  ])('puts a %s%% gap in tier %s', (gap, letter) => {
    expect(tierOf(gap)).toBe(letter);
  });

  it('is not thrown off by float noise on a spec exactly on a line', () => {
    const rows = rankRole(
      [100, 95, 90, 80, 70].map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    expect(rows.map((r) => r.tier)).toEqual(['S', 'A', 'B', 'C', 'D']);
  });

  it('lets a tie pair straddle a line, keeping the mark on both rows', () => {
    const rows = rankRole(
      [100, 95.3, 94.8].map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    expect(rows.map((r) => r.tier)).toEqual(['S', 'S', 'A']);
    expect(rows.map((r) => r.tie)).toEqual([null, 'below', 'above']);
  });
});

describe('hrefsFor', () => {
  it('links the BiS page, the guide and the level 60 build in the planner', () => {
    const [row] = rankRole([dpsInput('warrior', 'fury', 'Fury', 100)], 'dps');
    expect(hrefsFor({ ...row!, talents: '1-2-3' })).toEqual({
      bis: '/bis/warrior/fury',
      guide: '/guides/warrior/fury',
      planner: '/planner?spec=warrior-fury&talents=1-2-3',
    });
  });
});
