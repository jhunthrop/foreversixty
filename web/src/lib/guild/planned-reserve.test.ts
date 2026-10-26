// web/src/lib/guild/planned-reserve.test.ts
import { describe, expect, it } from 'vitest';
import { plannedReserveRows, rememberPlannedRows } from './planned-reserve';

function memoryStorage(initial: Record<string, string> = {}): Storage {
  const data = new Map(Object.entries(initial));
  return {
    getItem: (key) => data.get(key) ?? null,
    setItem: (key, value) => void data.set(key, value),
    removeItem: (key) => void data.delete(key),
    clear: () => data.clear(),
    key: () => null,
    get length() {
      return data.size;
    },
  };
}

describe('the planned reserve', () => {
  it('reads the stored count, 0 to 4, else one row', () => {
    expect(plannedReserveRows(2, memoryStorage({ 'fs.guildPlanned.2': '3' }))).toBe(3);
    expect(plannedReserveRows(2, memoryStorage({ 'fs.guildPlanned.2': '0' }))).toBe(0);
    expect(plannedReserveRows(2, memoryStorage({ 'fs.guildPlanned.2': '9' }))).toBe(1);
    expect(plannedReserveRows(2, memoryStorage())).toBe(1);
    expect(plannedReserveRows(2, null)).toBe(1);
    const throwing = {
      getItem: () => {
        throw new Error('blocked');
      },
    };
    expect(plannedReserveRows(2, throwing)).toBe(1);
  });

  it('writes the rows drawn, capped at four, and nothing for a visitor', () => {
    const storage = memoryStorage();
    rememberPlannedRows(2, { rows: [], next_cursor: null, planned: new Array(7).fill(null) }, storage);
    expect(storage.getItem('fs.guildPlanned.2')).toBe('4');
    rememberPlannedRows(3, { rows: [], next_cursor: null }, storage);
    expect(storage.getItem('fs.guildPlanned.3')).toBeNull();
  });
});
