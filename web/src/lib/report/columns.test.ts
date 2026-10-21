import { describe, expect, it } from 'vitest';
import { biggestFightIndex, defaultFightIndex } from './fights';
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
  it('is, for the sample, the longest pull', () => {
    expect(biggestFightIndex(night)).toBe(3);
  });
  it('stays the first boss pull for a real report', () => {
    expect(defaultFightIndex([...night, f(4, 'encounter', true, 90_000)])).toBe(4);
    expect(defaultFightIndex(night)).toBe(1);
  });
});
