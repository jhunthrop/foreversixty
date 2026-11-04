// web/src/lib/bis/weight-reading.test.ts
import { describe, expect, it } from 'vitest';
import { readingOf as railReadingOf } from '../guides/rail-stats';
import { bandReadingOf, isRanked, perPercentOf, readingOf } from './weight-reading';

describe('readingOf', () => {
  it('reads a weight inside its error, or flagged insignificant, as not significant', () => {
    expect(readingOf({ stat: 'spirit', weight: 0.05, error: 0.1 })).toBe('not-significant');
    expect(readingOf({ stat: 'spirit', weight: -0.05, error: 0.1 })).toBe('not-significant');
    expect(readingOf({ stat: 'spirit', weight: 3, error: 0.1, insignificant: true })).toBe('not-significant');
    expect(readingOf(undefined)).toBe('not-significant');
  });

  it('reads a weight below zero beyond its error as no gain, never a number', () => {
    expect(readingOf({ stat: 'spell_haste', weight: -6.62, error: 0.74 })).toBe('no-gain');
  });

  it('reads a positive weight beyond its error as a value', () => {
    expect(readingOf({ stat: 'agility', weight: 2, error: 0.1 })).toBe('value');
  });

  it('is the one rule the guide stat table uses too', () => {
    expect(railReadingOf).toBe(readingOf);
  });
});

describe('perPercentOf', () => {
  it('names haste and expertise per 1%, every other stat per point', () => {
    expect(perPercentOf('melee_haste')).toBe('haste');
    expect(perPercentOf('spell_haste')).toBe('haste');
    expect(perPercentOf('expertise')).toBe('expertise');
    expect(perPercentOf('strength')).toBeUndefined();
  });
});

describe('bandReadingOf', () => {
  it('reads haste no item of the band carries as not on any item, and nothing else', () => {
    const haste = { stat: 'spell_haste', weight: 4.9, error: 0.4 };
    expect(bandReadingOf(haste, false)).toBe('not-on-items');
    expect(bandReadingOf(haste, true)).toBe('value');
    expect(bandReadingOf({ stat: 'spell_power', weight: 1, error: 0.01 }, false)).toBe('value');
    // Round 20: haste no item carries reads "Not on any item" whatever its weight says (Arcane
    // read "Not significant" beside Fire and Frost's "Not on any item" for the same empty pool).
    expect(bandReadingOf({ stat: 'spell_haste', weight: 0.4, error: 0.9 }, false)).toBe('not-on-items');
    expect(bandReadingOf({ stat: 'spell_haste', weight: 0.4, error: 0.9 }, true)).toBe('not-significant');
    expect(bandReadingOf(undefined, false)).toBe('not-significant');
  });

  it('ranks only a value', () => {
    expect(isRanked('value')).toBe(true);
    expect(isRanked('not-on-items')).toBe(false);
    expect(isRanked('not-significant')).toBe(false);
    expect(isRanked('no-gain')).toBe(false);
  });
});
