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
  rememberDone,
  rememberStripRows,
  reservedDoneRows,
  stripReserveRows,
} from './raid-page';
import { RAID_ROWS_KEY } from './reserve-keys';
import { attendanceRow } from '../../test-support/report-fixtures';

function memoryStorage(
  seed: Record<string, string> = {},
): Pick<Storage, 'getItem' | 'setItem' | 'removeItem'> {
  const data = new Map(Object.entries(seed));
  return {
    getItem: (key) => data.get(key) ?? null,
    setItem: (key, value) => void data.set(key, value),
    removeItem: (key) => void data.delete(key),
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
  it('reads the date, times and zone in the guild clock, with no your-time line when the offsets agree', () => {
    expect(raidWhen(plannedRaid(), NOW, CHICAGO, 'en-US')).toEqual({
      date: 'Thursday 10 December',
      start: '8:00 pm',
      end: '11:00 pm\u00a0CST',
      yourTime: null,
      countdown: 'in 2 days 8 h',
    });
  });

  it('keeps the guild clock first and adds the viewer’s own start when the offsets differ (spec 3.1)', () => {
    const when = raidWhen(plannedRaid(), NOW, NEW_YORK, 'en-US');
    expect(when.date).toBe('Thursday 10 December');
    expect(when.start).toBe('8:00 pm');
    expect(when.end).toBe('11:00 pm\u00a0CST');
    expect(when.yourTime).toBe('9:00 pm EST your time');
  });

  it('reads a raid that runs to midnight as `to midnight`, never 12:00 am (guild loop round 3)', () => {
    const late = plannedRaid({ starts_at: '2026-12-11T03:00:00Z', duration_minutes: 180 });
    expect(raidWhen(late, NOW, CHICAGO, 'en-US').start).toBe('9:00 pm');
    expect(raidWhen(late, NOW, CHICAGO, 'en-US').end).toBe('midnight\u00a0CST');
    expect(raidWhen(late, NOW, CHICAGO, 'en-US-u-hc-h23').end).toBe('midnight\u00a0CST');
    const past = plannedRaid({ starts_at: '2026-12-11T03:00:00Z', duration_minutes: 210 });
    expect(raidWhen(past, NOW, CHICAGO, 'en-US').end).toBe('12:30 am\u00a0CST');
  });

  it('names the viewer’s weekday when their date is another', () => {
    expect(raidWhen(plannedRaid(), NOW, 'Europe/London', 'en-US').yourTime).toBe('Fri 2:00 am GMT your time');
  });
});

describe('raidCount', () => {
  it('counts in plus late against the size, with the bar fill', () => {
    expect(raidCount(plannedRaid())).toEqual({
      line: '23 of 40 signed up',
      ratio: '23/40',
      // Each clause glued with no-break spaces: the line wraps only between clauses.
      breakdown: '3\u00a0tanks · 7\u00a0healers · 7\u00a0melee · 6\u00a0ranged · 2\u00a0late',
      fill: 57,
    });
  });

  it('drops the bar while the size is unknown', () => {
    expect(raidCount(plannedRaid({ size: null }))).toMatchObject({ line: '23 signed up', fill: null });
  });

  it('says nobody signed up, with no role breakdown, at zero', () => {
    const counts = { in: 0, late: 0, bench: 1, out: 2, tanks: 0, healers: 0, dps: 0 };
    expect(raidCount(plannedRaid({ counts, size: 5 }))).toEqual({
      line: 'Nobody signed up yet',
      ratio: '0/5',
      breakdown: '',
      fill: 0,
    });
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

describe('done reserve', () => {
  it('remembers a done raid by id and forgets it once the raid is not done', () => {
    const storage = memoryStorage();
    expect(reservedDoneRows('7', storage)).toBeNull();
    rememberDone(plannedRaid({ status: 'done' }), storage);
    expect(reservedDoneRows('7', storage)).toBe(0);
    expect(reservedDoneRows('8', storage)).toBeNull();
    rememberDone(plannedRaid(), storage);
    expect(reservedDoneRows('7', storage)).toBeNull();
  });

  it('counts the own attendance rows Your night draws', () => {
    const storage = memoryStorage();
    rememberDone(
      plannedRaid({
        status: 'done',
        attendance: [attendanceRow({ own: true }), attendanceRow({ character_key: 'x', own: false })],
      }),
      storage,
    );
    expect(reservedDoneRows('7', storage)).toBe(1);
  });

  it('survives a store that throws', () => {
    const throwing = {
      getItem: (): string | null => {
        throw new Error('blocked');
      },
      setItem: (): void => {
        throw new Error('quota');
      },
      removeItem: (): void => {
        throw new Error('blocked');
      },
    };
    expect(reservedDoneRows('7', throwing)).toBeNull();
    expect(() => rememberDone(plannedRaid({ status: 'done' }), throwing)).not.toThrow();
  });
});
