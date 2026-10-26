// web/src/lib/guild/reserve-keys.test.ts
import { describe, expect, it } from 'vitest';
import { bandHoldKey } from './band-hold';
import { forgetGuildReserves, RAID_ROWS_KEY } from './reserve-keys';

function memoryStorage(entries: Record<string, string>): Storage {
  const map = new Map(Object.entries(entries));
  return {
    get length() {
      return map.size;
    },
    key: (index: number) => [...map.keys()][index] ?? null,
    getItem: (key: string) => map.get(key) ?? null,
    setItem: (key: string, value: string) => void map.set(key, value),
    removeItem: (key: string) => void map.delete(key),
    clear: () => map.clear(),
  };
}

describe('forgetGuildReserves', () => {
  it("drops every guild reserve on sign-out and keeps the rest of the site's keys", () => {
    const storage = memoryStorage({
      [bandHoldKey('us', 'pvp', 'olympus-xxvii')]: '{"desktop":132}',
      [bandHoldKey('eu', 'pve', 'other')]: '{"phone":200}',
      'fs.guildPlanned.2': '3',
      [RAID_ROWS_KEY]: '2',
      'fs.theme': 'dark',
    });
    forgetGuildReserves(storage);
    expect(storage.length).toBe(1);
    expect(storage.getItem('fs.theme')).toBe('dark');
  });

  it('survives a missing or blocked store', () => {
    expect(() => forgetGuildReserves(null)).not.toThrow();
    const blocked = {
      get length(): number {
        throw new Error('blocked');
      },
    } as unknown as Storage;
    expect(() => forgetGuildReserves(blocked)).not.toThrow();
  });
});
