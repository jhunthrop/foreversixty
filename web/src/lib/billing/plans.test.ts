// web/src/lib/billing/plans.test.ts
import { describe, expect, it } from 'vitest';
import {
  FOUNDER_LAST_DAY,
  FOUNDER_PRICE_UNTIL,
  FOUNDER_WINDOW_CLOSES,
  PLANS,
  RAID_SIZES,
  TRIAL_REMINDER_DAYS,
  founderOpen,
  formatUsd,
  formatCents,
  formatPrice,
  perMonthLine,
} from './plans';

describe('PLANS', () => {
  it('names the plans and their seats', () => {
    expect(PLANS.premium).toMatchObject({ name: 'Premium', guild: false, seats: 0 });
    expect(PLANS.guild).toMatchObject({ name: 'Guild', guild: true, seats: 1 });
    expect(PLANS.raid_guild).toMatchObject({ name: 'Raid guild', guild: true, seats: 3 });
  });
});

describe('founderOpen', () => {
  it('is open at the last second and closed at 08:00:00Z on 7 January', () => {
    expect(FOUNDER_WINDOW_CLOSES).toBe('2027-01-07T08:00:00Z');
    expect(founderOpen(new Date('2027-01-07T07:59:59Z'))).toBe(true);
    expect(founderOpen(new Date('2027-01-07T08:00:00Z'))).toBe(false);
  });

  it('has no opening instant', () => {
    expect(founderOpen(new Date('2020-01-01T00:00:00Z'))).toBe(true);
  });
});

describe('the founder dates', () => {
  // The literals stand in for these two time-zoned formats (built at import they cost the
  // hydration task ~5 ms); this keeps them the window's own last Pacific day.
  const lastInstant = new Date(Date.parse(FOUNDER_WINDOW_CLOSES) - 1);
  const pacific = (month: 'short' | 'long'): string =>
    lastInstant.toLocaleDateString('en-GB', {
      day: 'numeric',
      month,
      year: 'numeric',
      timeZone: 'America/Los_Angeles',
    });

  it('are the last Pacific day of the founder window', () => {
    expect(FOUNDER_LAST_DAY).toBe(pacific('short'));
    expect(FOUNDER_PRICE_UNTIL).toBe(pacific('long'));
  });
});

describe('formatUsd', () => {
  it('shows cents only when the amount is not whole', () => {
    expect(formatUsd(50 / 12)).toBe('$4.17');
    expect(formatUsd(150 / 12)).toBe('$12.50');
    expect(formatUsd(40 / 12)).toBe('$3.33');
    expect(formatUsd(120 / 12)).toBe('$10');
  });
});

describe('formatPrice', () => {
  it('reads as running text', () => {
    expect(formatPrice(40, 'yearly')).toBe('$40 a year');
    expect(formatPrice(5, 'monthly')).toBe('$5 a month');
  });
});

describe('perMonthLine', () => {
  it('divides the yearly price over twelve months', () => {
    expect(perMonthLine(40)).toBe('$3.33 a month, billed yearly');
    expect(perMonthLine(120)).toBe('$10 a month, billed yearly');
  });
});

describe('formatCents', () => {
  it('matches formatPrice', () => {
    expect(formatCents(4000, 'yearly')).toBe('$40 a year');
    expect(formatCents(500, 'monthly')).toBe('$5 a month');
  });
});

describe('raid sizes and trial reminder (B3.3)', () => {
  it('RAID_SIZES leads with 40', () => expect(RAID_SIZES).toEqual([40, 20, 10]));
  it('TRIAL_REMINDER_DAYS is 3', () => expect(TRIAL_REMINDER_DAYS).toBe(3));
});
