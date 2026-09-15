// web/src/lib/report/fights.test.ts
// parseReportState accepts any run of digits for ?fight=, because a total parser is what
// makes a pasted link render something rather than throw. It cannot know which fights a
// report actually has -- report.json has not loaded when the url is read -- so the island
// resolves the parsed number against the real list, and this covers the three ways that
// number can miss: 0 (fight_index is 1-based, so no report ever has one), a number past
// the end, and a gap.
import { describe, expect, it } from 'vitest';
import { defaultFightIndex, pullNumbers, resolveFightIndex } from './fights';
import type { FightEntry } from './types';
import { ALL_FIGHTS } from './url';

function fight(index: number): FightEntry {
  return {
    index,
    kind: 'encounter',
    name: `Fight ${index}`,
    kill: true,
    in_progress: false,
    start: '2026-09-26T20:10:00Z',
    end: '2026-09-26T20:11:00Z',
    duration_ms: 60_000,
    players: [],
    deaths: 0,
    npc_kills: 1,
  };
}

describe('resolveFightIndex', () => {
  const fights = [fight(1), fight(2), fight(4)];

  it('keeps a fight the report actually has', () => {
    expect(resolveFightIndex(fights, 2, 1)).toBe(2);
    expect(resolveFightIndex(fights, 4, 1)).toBe(4);
  });

  it('reads ?fight=0 as the whole night, which needs a boss pull to mean anything', () => {
    expect(resolveFightIndex(fights, 0, 1)).toBe(ALL_FIGHTS);
    expect(resolveFightIndex([{ ...fight(1), kind: 'trash' }], 0, 1)).toBe(1);
  });

  it('falls back for a fight past the end and for a gap in the numbering', () => {
    expect(resolveFightIndex(fights, 999, 1)).toBe(1);
    expect(resolveFightIndex(fights, 3, 1)).toBe(1);
  });

  it('falls back to the first fight the report has, not to 1', () => {
    expect(resolveFightIndex([fight(7), fight(8)], 3, 7)).toBe(7);
  });

  it('returns the fallback when the report has no fights at all', () => {
    expect(resolveFightIndex([], 3, 1)).toBe(1);
  });
});

describe('pullNumbers', () => {
  it('numbers each boss pull in order and leaves trash unnumbered', () => {
    const fights: FightEntry[] = [
      { ...fight(1), name: 'General Kaal', kill: false },
      { ...fight(2), kind: 'trash', name: 'Trash' },
      { ...fight(3), name: 'General Kaal', kill: true },
      { ...fight(4), name: 'Kryxis the Voracious' },
    ];
    const numbers = pullNumbers(fights);
    expect(numbers.get(1)).toEqual({ pull: 1, of: 2 });
    expect(numbers.get(2)).toBeUndefined();
    expect(numbers.get(3)).toEqual({ pull: 2, of: 2 });
    expect(numbers.get(4)).toEqual({ pull: 1, of: 1 });
  });
});

describe('defaultFightIndex', () => {
  it('opens on the first boss pull, or the first fight when there is no boss', () => {
    const trash = { ...fight(1), kind: 'trash' };
    expect(defaultFightIndex([trash, fight(2), fight(3)])).toBe(2);
    expect(defaultFightIndex([trash, { ...fight(5), kind: 'trash' }])).toBe(1);
    expect(defaultFightIndex([])).toBe(1);
  });

  it('accepts the whole night only when the report has a boss pull', () => {
    expect(resolveFightIndex([fight(1), fight(2)], ALL_FIGHTS, 1)).toBe(ALL_FIGHTS);
    expect(resolveFightIndex([{ ...fight(1), kind: 'trash' }], ALL_FIGHTS, 1)).toBe(1);
  });
});
