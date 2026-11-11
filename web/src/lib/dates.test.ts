import { afterAll, beforeAll, describe, expect, it } from 'vitest';
import {
  compactAgo,
  dateOnlyWord,
  dayWord,
  formatDate,
  formatDayMonthYear,
  isPast,
  latest,
  monthHeading,
  nightDate,
  nightDateOnly,
  relativeDay,
  relativeTime,
  syncedWord,
  utcDayMonth,
  utcDayMonthYear,
  viewerZone,
  zonedDayMonth,
  zonedDayMonthYear,
} from './dates';

describe('utcDayMonth', () => {
  it('drops the year and reads the date in UTC', () => {
    expect(utcDayMonth('2026-12-31')).toBe('31 Dec');
    expect(utcDayMonth('2026-12-31T00:00:00Z')).toBe('31 Dec');
  });
});

describe('formatDate', () => {
  it('uses the site month style', () => {
    expect(formatDate(new Date('2026-09-17T00:00:00Z'))).toBe('17 Sep');
    expect(formatDate(new Date('2026-11-04T00:00:00Z'))).toBe('4 Nov');
    expect(formatDate(new Date('2027-06-01T00:00:00Z'))).toBe('1 Jun');
  });
});

describe('monthHeading', () => {
  it('names the month in full with its year, in UTC', () => {
    expect(monthHeading(new Date('2026-11-05T00:00:00Z'))).toBe('November 2026');
    expect(monthHeading(new Date('2026-09-30T23:30:00Z'))).toBe('September 2026');
    expect(monthHeading(new Date('nope'))).toBe('');
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
    expect(relativeDay(new Date('2026-09-26T12:00:00Z'), now)).toBe('26 Sep');
    expect(relativeDay(new Date('2026-09-15T12:55:13Z'), now)).toBe('15 Sep');
    expect(relativeDay(new Date('2026-06-01T00:00:00Z'), now)).toBe('1 Jun');
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

  it('writes Last synced in one format: words for 7 days, then the date', () => {
    expect(syncedWord(new Date('2026-10-20T15:00:00Z'), eveningTuesday)).toBe('today');
    expect(syncedWord(new Date('2026-10-19T15:00:00Z'), eveningTuesday)).toBe('yesterday');
    expect(syncedWord(new Date('2026-10-16T18:00:00Z'), eveningTuesday)).toBe('4 days ago');
    expect(syncedWord(new Date('2026-10-13T18:00:00Z'), eveningTuesday)).toBe('7 days ago');
    expect(syncedWord(new Date('2026-10-12T18:00:00Z'), eveningTuesday)).toBe('12 Oct');
    expect(syncedWord(new Date('2025-12-31T18:00:00Z'), eveningTuesday)).toBe('31 Dec 2025');
    expect(syncedWord(new Date('2026-12-28T18:00:00Z'), new Date('2027-01-08T18:00:00Z'))).toBe('28 Dec');
    expect(syncedWord(new Date('2026-10-22T18:00:00Z'), eveningTuesday)).toBe('today');
    expect(syncedWord(new Date('nope'), eveningTuesday)).toBe('');
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

describe('zoned day helpers', () => {
  const instant = '2026-12-20T20:00:00Z';

  it('reads an instant on the viewer zone calendar around a day boundary', () => {
    expect(zonedDayMonth(instant, 'America/Los_Angeles')).toBe('20 Dec');
    expect(zonedDayMonth(instant, 'UTC')).toBe('20 Dec');
    expect(zonedDayMonth(instant, 'Pacific/Auckland')).toBe('21 Dec');
    expect(zonedDayMonthYear('2026-12-31T20:00:00Z', 'Pacific/Auckland')).toBe('1 Jan 2027');
    expect(zonedDayMonthYear('2027-01-01T05:00:00Z', 'America/Los_Angeles')).toBe('31 Dec 2026');
  });

  it('reads a date-only value as written in every zone', () => {
    expect(zonedDayMonth('2026-12-20', 'Pacific/Auckland')).toBe('20 Dec');
    expect(zonedDayMonth('2026-12-20', 'America/Los_Angeles')).toBe('20 Dec');
  });

  it('renders an unreadable instant as nothing and defaults to the resolved zone', () => {
    expect(zonedDayMonthYear('nope', 'UTC')).toBe('');
    expect(zonedDayMonthYear(instant)).toBe(zonedDayMonthYear(instant, viewerZone()));
  });
});

describe('compactAgo', () => {
  const now = new Date('2026-11-07T12:00:00Z');
  const ago = (ms: number) => new Date(now.getTime() - ms);
  it('reads the freshness stamp compactly, then the date past 30 days', () => {
    expect(compactAgo(ago(20_000), now)).toBe('just now');
    expect(compactAgo(ago(5 * 60_000), now)).toBe('5 min ago');
    expect(compactAgo(ago(2 * 3_600_000 + 60_000), now)).toBe('2 h ago');
    expect(compactAgo(ago(3 * 86_400_000), now)).toBe('3 d ago');
    expect(compactAgo(new Date('2026-09-25T00:00:00Z'), now)).toBe('25 Sep');
  });
  it('reads a timestamp ahead of the clock as just now, and an invalid one as nothing', () => {
    expect(compactAgo(new Date(now.getTime() + 60_000), now)).toBe('just now');
    expect(compactAgo(new Date('nope'), now)).toBe('');
  });
});
