import { describe, expect, it } from 'vitest';
import { busyLine } from './busy-line';
import { coachCopy } from './copy';
import { utcDayDate } from './dates';

describe('busyLine', () => {
  it("is the site's busy line without an account cap", () => {
    expect(busyLine({ cap: null })).toBe(coachCopy.busy);
  });

  it("names today's budget for the daily cap", () => {
    expect(busyLine({ cap: { period: 'day', until: '2026-10-29' } })).toBe(
      "Your coach has used today's model budget. Reviews are from the log until tomorrow.",
    );
  });

  it('names the first of next month for the monthly cap', () => {
    expect(busyLine({ cap: { period: 'month', until: '2026-11-01' } })).toBe(
      "Your coach has used this month's model budget. Reviews are from the log until 1 Nov.",
    );
  });
});

describe('utcDayDate', () => {
  it('reads the UTC day in every zone', () => {
    expect(utcDayDate('2026-11-01')).toBe('1 Nov');
  });
  it('keeps unparseable text', () => {
    expect(utcDayDate('soon')).toBe('soon');
  });
});
