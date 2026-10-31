import { describe, expect, it } from 'vitest';
import { defaultFightIndex, sampleFightIndex } from './fights';
import { hasAnySpec, hasRankedKill } from './columns';
import type { FightEntry } from './types';

const f = (index: number, kind: string, kill: boolean, duration_ms: number): FightEntry =>
  ({ index, kind, kill, duration_ms, name: 'x' }) as unknown as FightEntry;

describe('the columns a report shows', () => {
  it('hides Parse on a report with no boss kill, shows it with one', () => {
    expect(hasRankedKill([f(1, 'trash', false, 5), f(2, 'encounter', false, 5)])).toBe(false);
    expect(hasRankedKill([f(1, 'trash', false, 5), f(2, 'encounter', true, 5)])).toBe(true);
    expect(hasRankedKill([])).toBe(false);
  });
  it('hides Spec when no row has one', () => {
    expect(hasAnySpec([{}, { spec: '' }])).toBe(false);
    expect(hasAnySpec([{}, { spec: 'Restoration' }])).toBe(true);
  });
});

describe('the fight a report opens on', () => {
  const night = [f(1, 'trash', false, 206_000), f(2, 'trash', false, 10_000), f(3, 'trash', false, 763_000)];
  it('is, for the sample, its first boss kill', () => {
    const bosses = [
      f(4, 'encounter', false, 300_000),
      f(5, 'encounter', true, 90_000),
      f(6, 'encounter', true, 95_000),
    ];
    expect(sampleFightIndex([...night, ...bosses])).toBe(5);
  });
  it('is, for a sample with no kill, its longest boss attempt', () => {
    const wipes = [
      f(4, 'encounter', false, 60_000),
      f(5, 'encounter', false, 140_000),
      f(6, 'encounter', false, 90_000),
    ];
    expect(sampleFightIndex([...night, ...wipes])).toBe(5);
  });
  it('is, for a sample with no boss, its longest pull', () => {
    expect(sampleFightIndex(night)).toBe(3);
  });
  it('is fight 1 for a sample with no fights', () => {
    expect(sampleFightIndex([])).toBe(1);
  });
  it('stays the first boss pull for a real report', () => {
    expect(defaultFightIndex([...night, f(4, 'encounter', true, 90_000)])).toBe(4);
    expect(defaultFightIndex(night)).toBe(1);
  });
});
