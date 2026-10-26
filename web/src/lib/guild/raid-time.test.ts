import { describe, expect, it } from 'vitest';
import {
  clockText,
  dateLabel,
  forYouLine,
  instantToWallClock,
  nextWeekday,
  offsetsDiffer,
  wallClockToInstant,
  zoneOffsetMinutes,
} from './raid-time';

const CHICAGO = 'America/Chicago';
const NEW_YORK = 'America/New_York';

describe('wallClockToInstant', () => {
  it('reads 8 pm Chicago after the spring-forward as CDT', () => {
    expect(wallClockToInstant('2027-03-14', '20:00', CHICAGO).toISOString()).toBe('2027-03-15T01:00:00.000Z');
  });

  it('reads 8 pm Chicago the week before as CST', () => {
    expect(wallClockToInstant('2027-03-07', '20:00', CHICAGO).toISOString()).toBe('2027-03-08T02:00:00.000Z');
  });

  it('keeps 8 pm guild time across the daylight saving change', () => {
    const before = wallClockToInstant('2027-03-07', '20:00', CHICAGO);
    const after = wallClockToInstant('2027-03-14', '20:00', CHICAGO);
    expect(after.getTime() - before.getTime()).toBe(7 * 24 * 3_600_000 - 3_600_000);
  });

  it('lands just after a skipped spring-forward hour', () => {
    const skipped = wallClockToInstant('2027-03-14', '02:30', CHICAGO);
    expect(instantToWallClock(skipped, CHICAGO).time).toBe('03:30');
  });

  it('reads a fall-back evening in CST', () => {
    expect(wallClockToInstant('2027-11-07', '20:00', CHICAGO).toISOString()).toBe('2027-11-08T02:00:00.000Z');
  });

  it('refuses a malformed clock', () => {
    expect(() => wallClockToInstant('14/03/2027', '20:00', CHICAGO)).toThrow('bad wall clock');
  });
});

describe('instantToWallClock', () => {
  it('round-trips the form fields', () => {
    const instant = wallClockToInstant('2026-12-10', '20:00', CHICAGO);
    expect(instantToWallClock(instant, CHICAGO)).toEqual({ date: '2026-12-10', time: '20:00' });
  });

  it('reads midnight as 00, not 24', () => {
    const instant = wallClockToInstant('2026-12-10', '00:00', CHICAGO);
    expect(instantToWallClock(instant, CHICAGO)).toEqual({ date: '2026-12-10', time: '00:00' });
  });
});

describe('zones', () => {
  it('gives the offset in minutes', () => {
    expect(zoneOffsetMinutes(new Date('2026-12-10T12:00:00Z'), CHICAGO)).toBe(-360);
    expect(zoneOffsetMinutes(new Date('2026-07-10T12:00:00Z'), CHICAGO)).toBe(-300);
  });

  it('says whether two zones disagree at an instant', () => {
    const winter = new Date('2026-12-10T12:00:00Z');
    expect(offsetsDiffer(winter, CHICAGO, NEW_YORK)).toBe(true);
    expect(offsetsDiffer(winter, CHICAGO, 'America/Denver')).toBe(true);
    expect(offsetsDiffer(winter, CHICAGO, 'America/Chicago')).toBe(false);
  });
});

describe('forYouLine', () => {
  it('renders the entered wall time in the viewer zone', () => {
    expect(forYouLine('2026-12-10', '20:00', CHICAGO, NEW_YORK, 'en-US')).toBe(
      'For you: Thu 10 Dec, 9:00 pm.',
    );
  });

  it('says same time when the zones agree', () => {
    expect(forYouLine('2026-12-10', '20:00', CHICAGO, CHICAGO, 'en-US')).toBe('For you: same time.');
  });
});

describe('labels', () => {
  const instant = new Date('2026-12-11T02:30:00Z');

  it('writes minutes only off the hour, and always when asked', () => {
    expect(clockText(instant, CHICAGO, 'en-US')).toBe('8:30 pm');
    expect(clockText(new Date('2026-12-11T02:00:00Z'), CHICAGO, 'en-US')).toBe('8 pm');
    expect(clockText(new Date('2026-12-11T02:00:00Z'), CHICAGO, 'en-US', true)).toBe('8:00 pm');
  });

  it('writes a 24-hour clock where the locale does', () => {
    expect(clockText(new Date('2026-12-11T02:00:00Z'), CHICAGO, 'de-DE')).toBe('20:00');
  });

  it('writes the date on the zone calendar', () => {
    expect(dateLabel(instant, CHICAGO)).toBe('Thu 10 Dec');
    expect(dateLabel(instant, 'Europe/Berlin')).toBe('Fri 11 Dec');
  });
});

describe('nextWeekday', () => {
  const last = new Date('2026-12-10T02:00:00Z');

  it('is a week after a raid still ahead', () => {
    expect(nextWeekday(last, new Date('2026-12-01T00:00:00Z')).toISOString()).toBe(
      '2026-12-17T02:00:00.000Z',
    );
  });

  it('skips to the first same weekday after now', () => {
    expect(nextWeekday(last, new Date('2026-12-20T00:00:00Z')).toISOString()).toBe(
      '2026-12-24T02:00:00.000Z',
    );
  });

  it('never returns now itself', () => {
    expect(nextWeekday(last, new Date('2026-12-17T02:00:00Z')).toISOString()).toBe(
      '2026-12-24T02:00:00.000Z',
    );
  });
});
