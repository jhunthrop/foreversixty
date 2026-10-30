// web/src/lib/chart/choice.test.ts
import { describe, expect, it } from 'vitest';
import { matchesFilter, togglePlayer } from './choice';
import type { PlayerEntry } from './types';

const entry = (id: string, drawn: boolean, pinned = false): PlayerEntry => ({
  id,
  label: id,
  colour: 'var(--color-text)',
  role: 'dps',
  figure: 1,
  drawn,
  pinned,
  you: false,
});

describe('togglePlayer', () => {
  it('starts from the drawn players (less the pinned ones) when the default holds', () => {
    const entries = [entry('a', true), entry('b', true), entry('me', true, true), entry('c', false)];
    expect([...togglePlayer(null, entries, 'c')].sort()).toEqual(['a', 'b', 'c']);
    expect([...togglePlayer(null, entries, 'a')]).toEqual(['b']);
  });

  it('toggles inside an existing choice', () => {
    expect([...togglePlayer(new Set(['a']), [], 'a')]).toEqual([]);
  });
});

describe('matchesFilter', () => {
  it('matches case- and accent-insensitively', () => {
    expect(matchesFilter('Élyra', 'ely')).toBe(true);
    expect(matchesFilter('Garrosk', 'ROSK')).toBe(true);
    expect(matchesFilter('Garrosk', 'z')).toBe(false);
  });
});
