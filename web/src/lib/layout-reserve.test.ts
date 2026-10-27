import { afterEach, describe, expect, it, vi } from 'vitest';
import { readReserve, writeReserve } from './layout-reserve';

function memoryStorage(): Storage {
  const values = new Map<string, string>();
  return {
    get length() {
      return values.size;
    },
    clear: () => values.clear(),
    getItem: (key) => values.get(key) ?? null,
    key: (index) => [...values.keys()][index] ?? null,
    removeItem: (key) => void values.delete(key),
    setItem: (key, value) => void values.set(key, value),
  };
}

afterEach(() => vi.unstubAllGlobals());

describe('readReserve', () => {
  it('reads back what writeReserve stored', () => {
    vi.stubGlobal('localStorage', memoryStorage());
    writeReserve('k', 612.4);
    expect(readReserve('k')).toBe(612);
  });

  it('is null when missing, malformed or out of bounds', () => {
    const store = memoryStorage();
    vi.stubGlobal('localStorage', store);
    expect(readReserve('missing')).toBeNull();
    store.setItem('bad', 'tall');
    expect(readReserve('bad')).toBeNull();
    store.setItem('huge', '9000');
    expect(readReserve('huge')).toBeNull();
    store.setItem('zero', '0');
    expect(readReserve('zero')).toBeNull();
  });

  it('is null when storage throws', () => {
    vi.stubGlobal('localStorage', {
      getItem: () => {
        throw new Error('blocked');
      },
    });
    expect(readReserve('k')).toBeNull();
  });
});
