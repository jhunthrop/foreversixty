// web/src/lib/account/refresh-snapshot.test.ts
import { describe, expect, it } from 'vitest';
import {
  compareSnapshot,
  saveSnapshot,
  snapshotOf,
  takeSnapshot,
  type SnapshotStore,
} from './refresh-snapshot';
import type { MeCharacter } from './me-types';

const char = (key: string, level: number): MeCharacter => ({
  key,
  region: 'us',
  ruleset: 'pvp',
  name: key,
  level,
});
const me = (at: string | undefined, ...characters: MeCharacter[]) => ({ bnet_imported_at: at, characters });

describe('compareSnapshot', () => {
  const before = snapshotOf(me('2026-11-04T10:00:00Z', char('a', 1), char('b', 1)));

  it('reads a newer import with a new character or a new level as changed', () => {
    expect(
      compareSnapshot(before, me('2026-11-04T11:00:00Z', char('a', 1), char('b', 1), char('c', 1))),
    ).toEqual({
      kind: 'changed',
      count: 3,
      added: 1,
    });
    expect(compareSnapshot(before, me('2026-11-04T11:00:00Z', char('a', 2), char('b', 1)))).toEqual({
      kind: 'changed',
      count: 2,
      added: 0,
    });
  });

  it('reads a newer import with the same characters and levels as no changes', () => {
    expect(compareSnapshot(before, me('2026-11-04T11:00:00Z', char('a', 1), char('b', 1)))).toEqual({
      kind: 'unchanged',
      count: 2,
    });
  });

  it('reads an import time that did not move as an error', () => {
    expect(compareSnapshot(before, me('2026-11-04T10:00:00Z', char('a', 1)))).toEqual({ kind: 'error' });
    expect(compareSnapshot(before, me(undefined, char('a', 1)))).toEqual({ kind: 'error' });
  });

  it('reads a return with no snapshot as updated', () => {
    expect(compareSnapshot(null, me('2026-11-04T11:00:00Z', char('a', 1)))).toEqual({
      kind: 'changed',
      count: 1,
      added: 0,
    });
  });
});

describe('saveSnapshot and takeSnapshot', () => {
  function memory(): SnapshotStore {
    const map = new Map<string, string>();
    return {
      getItem: (k) => map.get(k) ?? null,
      setItem: (k, v) => void map.set(k, v),
      removeItem: (k) => void map.delete(k),
    };
  }

  it('round-trips once, then is gone', () => {
    const store = memory();
    const snapshot = snapshotOf(me('2026-11-04T10:00:00Z', char('a', 1)));
    saveSnapshot(snapshot, store);
    expect(takeSnapshot(store)).toEqual(snapshot);
    expect(takeSnapshot(store)).toBeNull();
  });

  it('survives a store that throws', () => {
    const broken: SnapshotStore = {
      getItem: () => {
        throw new Error('blocked');
      },
      setItem: () => {
        throw new Error('blocked');
      },
      removeItem: () => {
        throw new Error('blocked');
      },
    };
    expect(() => saveSnapshot({ at: null, chars: {} }, broken)).not.toThrow();
    expect(takeSnapshot(broken)).toBeNull();
  });
});
