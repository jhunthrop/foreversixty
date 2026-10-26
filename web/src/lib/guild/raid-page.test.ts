// web/src/lib/guild/raid-page.test.ts
// The raid page's pure rules: the poll window, the page flags, the band's when and count lines and
// the strip's loading reserve kept in localStorage.
import { describe, expect, it } from 'vitest';
import { CHICAGO, NEW_YORK, NOW, plannedRaid, viewerCharacter } from '../../test-support/raid-fixtures';
import {
  pollOpen,
  raidCount,
  raidPageFlags,
  raidWhen,
  rememberStripRows,
  stripReserveRows,
} from './raid-page';
import { RAID_ROWS_KEY } from './reserve-keys';

function memoryStorage(seed: Record<string, string> = {}): Pick<Storage, 'getItem' | 'setItem'> {
  const data = new Map(Object.entries(seed));
  return {
    getItem: (key) => data.get(key) ?? null,
    setItem: (key, value) => void data.set(key, value),
  };
}

const HOUR = 3_600_000;
/** The fixture raid's end: 8 pm Chicago plus 3 h. */
const END = Date.parse('2026-12-11T05:00:00Z');

describe('pollOpen', () => {
  it('polls planned and locked raids until 6 h after the end', () => {
    expect(pollOpen(plannedRaid(), NOW)).toBe(true);
    expect(pollOpen(plannedRaid({ status: 'locked' }), new Date(END + 6 * HOUR - 1))).toBe(true);
    expect(pollOpen(plannedRaid(), new Date(END + 6 * HOUR))).toBe(false);
  });

  it('never polls a done or cancelled raid', () => {
    expect(pollOpen(plannedRaid({ status: 'done' }), NOW)).toBe(false);
    expect(pollOpen(plannedRaid({ status: 'cancelled' }), NOW)).toBe(false);
  });
});

describe('raidPageFlags', () => {
  it('reads locked, ended, cancelled and done', () => {
    expect(raidPageFlags(plannedRaid(), NOW)).toEqual({
      status: 'planned',
      locked: false,
      ended: false,
      cancelled: false,
      done: false,
    });
    expect(raidPageFlags(plannedRaid({ status: 'locked' }), NOW).locked).toBe(true);
    expect(raidPageFlags(plannedRaid({ status: 'locked' }), new Date(END + HOUR))).toMatchObject({
      status: 'ended',
      locked: false,
      ended: true,
    });
    expect(raidPageFlags(plannedRaid({ status: 'cancelled' }), NOW).cancelled).toBe(true);
    expect(raidPageFlags(plannedRaid({ status: 'done' }), NOW).done).toBe(true);
  });
});

describe('raidWhen', () => {
  it('reads the date and times in the viewer zone, with no guild time when the offsets agree', () => {
    expect(raidWhen(plannedRaid(), NOW, CHICAGO, 'en-US')).toEqual({
      date: 'Thursday 10 December',
      start: '8:00 pm',
      end: '11:00 pm',
      guildTime: null,
      countdown: 'in 2 days 8 h',
    });
  });

  it('adds the guild time when the offsets differ', () => {
    const when = raidWhen(plannedRaid(), NOW, NEW_YORK, 'en-US');
    expect(when.start).toBe('9:00 pm');
    expect(when.end).toBe('12:00 am');
    expect(when.guildTime).toBe('8:00 pm CST guild time');
  });
});

describe('raidCount', () => {
  it('counts in plus late against the size, with the bar fill', () => {
    expect(raidCount(plannedRaid())).toEqual({
      line: '23 of 40 signed up',
      ratio: '23/40',
      breakdown: '3 tanks · 7 healers · 13 DPS · 2 late',
      fill: 57,
    });
  });

  it('drops the bar while the size is unknown', () => {
    expect(raidCount(plannedRaid({ size: null }))).toMatchObject({ line: '23 signed up', fill: null });
  });
});

describe('strip reserve rows', () => {
  it('is one row when nothing usable is stored, or the store throws', () => {
    expect(stripReserveRows(memoryStorage())).toBe(1);
    expect(stripReserveRows(null)).toBe(1);
    expect(stripReserveRows(memoryStorage({ [RAID_ROWS_KEY]: 'lots' }))).toBe(1);
    const throwing = {
      getItem: (): string | null => {
        throw new Error('blocked');
      },
    };
    expect(stripReserveRows(throwing)).toBe(1);
  });

  it('reads back the own-character count written after a load', () => {
    const storage = memoryStorage();
    const second = viewerCharacter({ character_key: 'us-pvp-menda', name: 'Menda' });
    rememberStripRows(plannedRaid({ viewer: { characters: [viewerCharacter(), second] } }), storage);
    expect(storage.getItem(RAID_ROWS_KEY)).toBe('2');
    expect(stripReserveRows(storage)).toBe(2);
  });

  it('survives a store that throws on write', () => {
    const throwing = {
      setItem: (): void => {
        throw new Error('quota');
      },
    };
    expect(() => rememberStripRows(plannedRaid(), throwing)).not.toThrow();
  });
});
