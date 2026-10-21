import { afterAll, beforeAll, describe, expect, it } from 'vitest';
import {
  dateOnlyWord,
  dayWord,
  formatDate,
  formatDayMonthYear,
  isPast,
  latest,
  nightDate,
  nightDateOnly,
  relativeDay,
  relativeTime,
  utcDayMonthYear,
} from './dates';

describe('formatDate', () => {
  it('uses the site month style', () => {
    expect(formatDate(new Date('2026-09-17T00:00:00Z'))).toBe('Sept 17');
    expect(formatDate(new Date('2026-11-04T00:00:00Z'))).toBe('Nov 4');
    expect(formatDate(new Date('2027-06-01T00:00:00Z'))).toBe('June 1');
  });
});

describe('isPast', () => {
  it('compares against a supplied now', () => {
    const now = new Date('2026-09-12T12:00:00Z');
    expect(isPast(new Date('2026-09-11T00:00:00Z'), now)).toBe(true);
    expect(isPast(new Date('2026-09-17T00:00:00Z'), now)).toBe(false);
  });
});

describe('latest', () => {
  it('returns the most recent date regardless of order', () => {
    const dates = [
      new Date('2026-09-11T00:00:00Z'),
      new Date('2026-09-17T00:00:00Z'),
      new Date('2026-09-12T00:00:00Z'),
    ];
    expect(latest(dates).toISOString()).toBe('2026-09-17T00:00:00.000Z');
  });

  it('fails loudly on an empty list rather than returning an invalid date', () => {
    expect(() => latest([])).toThrow(/at least one date/);
  });
});

describe('relativeTime', () => {
  const now = new Date('2026-09-21T12:00:00Z');

  it('buckets by elapsed time at each boundary', () => {
    expect(relativeTime(new Date('2026-09-21T11:59:59Z'), now)).toBe('just now');
    expect(relativeTime(new Date('2026-09-21T11:58:59Z'), now)).toBe('1 minute ago');
    expect(relativeTime(new Date('2026-09-21T11:00:01Z'), now)).toBe('59 minutes ago');
    expect(relativeTime(new Date('2026-09-21T11:00:00Z'), now)).toBe('1 hour ago');
    expect(relativeTime(new Date('2026-09-20T13:00:01Z'), now)).toBe('22 hours ago');
    expect(relativeTime(new Date('2026-09-20T12:00:00Z'), now)).toBe('1 day ago');
    expect(relativeTime(new Date('2026-08-23T12:00:00Z'), now)).toBe('29 days ago');
  });

  it('falls back to formatDate past 30 days', () => {
    expect(relativeTime(new Date('2026-08-22T12:00:00Z'), now)).toBe(
      formatDate(new Date('2026-08-22T12:00:00Z')),
    );
  });

  it('never reads a future timestamp as a negative duration', () => {
    expect(relativeTime(new Date('2026-09-21T12:00:05Z'), now)).toBe('just now');
  });
});

describe('relativeDay', () => {
  const now = new Date('2026-10-10T12:00:00Z');

  it('reads whole days inside two weeks', () => {
    expect(relativeDay(new Date('2026-10-10T01:00:00Z'), now)).toBe('today');
    expect(relativeDay(new Date('2026-10-09T12:00:00Z'), now)).toBe('1 day ago');
    expect(relativeDay(new Date('2026-10-07T12:00:00Z'), now)).toBe('3 days ago');
    expect(relativeDay(new Date('2026-09-27T12:00:01Z'), now)).toBe('12 days ago');
    expect(relativeDay(new Date('2026-09-27T12:00:00Z'), now)).toBe('13 days ago');
  });

  it('prints a short date from two weeks on', () => {
    expect(relativeDay(new Date('2026-09-26T12:00:00Z'), now)).toBe('Sep 26');
    expect(relativeDay(new Date('2026-09-15T12:55:13Z'), now)).toBe('Sep 15');
    expect(relativeDay(new Date('2026-06-01T00:00:00Z'), now)).toBe('Jun 1');
  });

  it('never reads a future timestamp as negative, and prints nothing for an invalid one', () => {
    expect(relativeDay(new Date('2026-10-10T12:00:05Z'), now)).toBe('today');
    expect(relativeDay(new Date('not a date'), now)).toBe('');
  });
});

describe('formatDayMonthYear', () => {
  it('reads "9 Dec 2026" in UTC, the Forever Score tables’ date (fix round 2 item 6)', () => {
    expect(formatDayMonthYear(new Date('2026-12-09T22:10:00Z'))).toBe('9 Dec 2026');
    expect(formatDayMonthYear(new Date('2026-09-30T23:59:00Z'))).toBe('30 Sep 2026');
    expect(formatDayMonthYear(new Date('not a date'))).toBe('');
  });
});

describe("the guild pages' day words, in a zone west of Greenwich", () => {
  const zone = process.env.TZ;
  beforeAll(() => {
    process.env.TZ = 'America/Chicago';
  });
  afterAll(() => {
    if (zone === undefined) delete process.env.TZ;
    else process.env.TZ = zone;
  });

  // Tuesday 20 Oct 2026, 20:00 CDT, is already Wednesday 21 Oct in UTC.
  const eveningTuesday = new Date('2026-10-21T01:00:00Z');

  it('reads a moment in the viewer zone, so a Tuesday evening is today and not tomorrow', () => {
    expect(dayWord(eveningTuesday, eveningTuesday)).toBe('today');
    expect(dayWord(new Date('2026-10-20T15:00:00Z'), eveningTuesday)).toBe('today');
    expect(dayWord(new Date('2026-10-19T15:00:00Z'), eveningTuesday)).toBe('yesterday');
  });

  it('names a weekday inside six days and the date beyond', () => {
    expect(dayWord(new Date('2026-10-16T18:00:00Z'), eveningTuesday)).toBe('Friday');
    expect(dayWord(new Date('2026-10-04T18:00:00Z'), eveningTuesday)).toBe('4 Oct');
    expect(dayWord(new Date('2025-12-31T18:00:00Z'), eveningTuesday)).toBe('31 Dec 2025');
  });

  it('never reads a future moment as a negative age', () => {
    expect(dayWord(new Date('2026-10-22T18:00:00Z'), eveningTuesday)).toBe('today');
  });

  it('reads a date-only value as written, not through the local zone', () => {
    expect(dateOnlyWord('2026-10-20', eveningTuesday)).toBe('today');
    expect(dateOnlyWord('2026-10-19', eveningTuesday)).toBe('yesterday');
    expect(dateOnlyWord('2026-10-21', eveningTuesday)).toBe('today');
  });

  it('prints the renewal date of a grant until 2026-12-31 as the 31st, not the 30th', () => {
    // The trap this guards: midnight UTC on the 31st is the evening of the 30th here.
    expect(new Date('2026-12-31T00:00:00Z').getDate()).toBe(30);
    expect(utcDayMonthYear('2026-12-31')).toBe('31 Dec 2026');
    expect(utcDayMonthYear('2026-12-31T00:00:00Z')).toBe('31 Dec 2026');
    expect(utcDayMonthYear('2026-12-31T00:00:00.000Z')).toBe('31 Dec 2026');
  });

  it('renders nothing for a value it cannot read', () => {
    expect(dayWord(new Date('nope'))).toBe('');
    expect(dateOnlyWord('soon')).toBe('');
  });
});

describe('nightDate (P.1: one night, one date)', () => {
  const now = new Date(2026, 9, 21, 9, 0);

  it('reads the viewer-zone day as a weekday, day and month, never today or yesterday', () => {
    expect(nightDate(new Date(2026, 9, 20, 23, 30), now)).toBe('Tue 20 Oct');
    expect(nightDate(new Date(2026, 9, 21, 1, 0), now)).toBe('Wed 21 Oct');
  });

  it('is the same words whichever day it is read on', () => {
    const night = new Date(2026, 9, 20, 20, 0);
    expect(nightDate(night, new Date(2026, 9, 20, 22, 0))).toBe(
      nightDate(night, new Date(2026, 9, 21, 22, 0)),
    );
  });

  it('adds the year only in another year, and reads a date-only value as written', () => {
    expect(nightDate(new Date(2025, 11, 31, 20, 0), now)).toBe('Wed 31 Dec 2025');
    expect(nightDateOnly('2026-10-20', now)).toBe('Tue 20 Oct');
    expect(nightDateOnly('nope', now)).toBe('');
    expect(nightDate(new Date('nope'), now)).toBe('');
  });
});
