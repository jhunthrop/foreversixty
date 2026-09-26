// web/src/lib/sim/last-upgrade.test.ts
import { describe, expect, it } from 'vitest';
import { readLastUpgrade, writeLastUpgrade } from './last-upgrade';

function fakeStorage(): Storage {
  const map = new Map<string, string>();
  return {
    getItem: (key) => map.get(key) ?? null,
    setItem: (key, value) => void map.set(key, value),
    removeItem: (key) => void map.delete(key),
    clear: () => map.clear(),
    key: () => null,
    get length() {
      return map.size;
    },
  } as Storage;
}

describe('last-upgrade', () => {
  it('round-trips a written upgrade', () => {
    const storage = fakeStorage();
    writeLastUpgrade(
      {
        itemName: 'Bracers of X',
        sourceName: 'Blackfathom Deeps',
        gain: '+14 ± 3',
        savedAt: '2026-01-01T00:00:00Z',
      },
      storage,
    );
    expect(readLastUpgrade(storage)).toEqual({
      itemName: 'Bracers of X',
      sourceName: 'Blackfathom Deeps',
      gain: '+14 ± 3',
      savedAt: '2026-01-01T00:00:00Z',
    });
  });

  it('reads null when nothing is stored', () => {
    expect(readLastUpgrade(fakeStorage())).toBeNull();
  });

  it('reads null for malformed JSON rather than throwing', () => {
    const storage = fakeStorage();
    storage.setItem('fs.lastDroptimizerUpgrade', '{not json');
    expect(readLastUpgrade(storage)).toBeNull();
  });

  it('reads null for a value missing a required field', () => {
    const storage = fakeStorage();
    storage.setItem('fs.lastDroptimizerUpgrade', JSON.stringify({ itemName: 'X' }));
    expect(readLastUpgrade(storage)).toBeNull();
  });
});
