import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { readReserve } from '../layout-reserve';
import {
  HOME_UPGRADES_MAX_RESERVE_PX,
  HOME_UPGRADES_RESERVE_KEY,
  storeUpgradesReserve,
} from './upgrades-layout';

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

beforeEach(() => vi.stubGlobal('localStorage', memoryStorage()));
afterEach(() => vi.unstubAllGlobals());

describe('storeUpgradesReserve', () => {
  it('keeps the settled height for the next visit', () => {
    storeUpgradesReserve(412.4);
    expect(readReserve(HOME_UPGRADES_RESERVE_KEY, HOME_UPGRADES_MAX_RESERVE_PX)).toBe(412);
  });

  it('drops an empty or oversized height', () => {
    storeUpgradesReserve(0);
    storeUpgradesReserve(HOME_UPGRADES_MAX_RESERVE_PX + 1);
    expect(localStorage.getItem(HOME_UPGRADES_RESERVE_KEY)).toBeNull();
  });
});
