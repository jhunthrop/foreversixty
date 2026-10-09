// web/src/lib/tiers/tier-list.test.ts
import { describe, expect, it } from 'vitest';
import { band, catalogEntry, dpsInput, tankInput } from './tier-test-support';
import {
  RULER_PERCENTS,
  listItemsFor,
  rankRole,
  raceLabel,
  sortMetricOf,
  withRulers,
  type TierRow,
} from './tier-list';

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

describe('race, confidence and rulers', () => {
  it('names the band race and carries the low-confidence flag', () => {
    const entry = catalogEntry('warlock', 'affliction', 'Affliction');
    const rows = rankRole(
      [{ entry, band: band(entry, { race: 'undead', weights_low_confidence: true }) }],
      'dps',
    );
    expect(rows[0]).toMatchObject({ race: 'Undead', lowConfidence: true });
  });

  it('puts a ruler before the first row at or past each percent and none after the last row', () => {
    const rows = rankRole(
      [100, 95, 88, 79].map((v, i) => dpsInput('c', `s${i}`, `S${i}`, v)),
      'dps',
    );
    const items = withRulers(rows);
    expect(items.map((i) => (i.kind === 'ruler' ? `r${i.percent}` : `s${i.row.rank}`))).toEqual([
      's1',
      's2',
      'r10',
      's3',
      'r20',
      's4',
    ]);
    expect(RULER_PERCENTS).toEqual([10, 20, 30]);
  });

  it('draws rulers for the DPS list only', () => {
    const rows = rankRole([dpsInput('a', 'a', 'A', 100), dpsInput('b', 'b', 'B', 50)], 'dps');
    expect(listItemsFor(rows, 'dps').some((i) => i.kind === 'ruler')).toBe(true);
    expect(listItemsFor(rows, 'healer').some((i) => i.kind === 'ruler')).toBe(false);
  });
});
