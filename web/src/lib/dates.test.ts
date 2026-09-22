import { describe, expect, it } from 'vitest';
import { formatDate, isPast, latest, relativeTime } from './dates';

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
