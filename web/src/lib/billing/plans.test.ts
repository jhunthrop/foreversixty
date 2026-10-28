// web/src/lib/billing/plans.test.ts
import { describe, expect, it } from 'vitest';
import {
  FOUNDER_LAST_DAY,
  FOUNDER_WINDOW_CLOSES,
  PLANS,
  RAID_SIZES,
  TRIAL_REMINDER_DAYS,
  founderOpen,
  formatAmount,
  formatCents,
  formatPrice,
  listPriceAfterFounder,
  perMonthLine,
  priceFor,
  type Interval,
  type PlanKey,
} from './plans';

const TWELVE_AMOUNTS: ReadonlyArray<[PlanKey, Interval, boolean, number]> = [
  ['premium', 'monthly', false, 6],
  ['premium', 'yearly', false, 50],
  ['premium', 'monthly', true, 5],
  ['premium', 'yearly', true, 40],
  ['guild', 'monthly', false, 15],
  ['guild', 'yearly', false, 150],
  ['guild', 'monthly', true, 12],
  ['guild', 'yearly', true, 120],
  ['raid_guild', 'monthly', false, 25],
  ['raid_guild', 'yearly', false, 250],
  ['raid_guild', 'monthly', true, 20],
  ['raid_guild', 'yearly', true, 200],
];

describe('priceFor', () => {
  it.each(TWELVE_AMOUNTS)('%s %s founder=%s is %d', (plan, interval, founder, usd) => {
    expect(priceFor(plan, interval, founder)).toBe(usd);
  });
});

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

describe('formatAmount', () => {
  it('shows cents only when the amount is not whole', () => {
    expect(formatAmount(50 / 12)).toBe('$4.17');
    expect(formatAmount(150 / 12)).toBe('$12.50');
    expect(formatAmount(40 / 12)).toBe('$3.33');
    expect(formatAmount(120 / 12)).toBe('$10');
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

describe('listPriceAfterFounder', () => {
  it('names the list price and the last founder day in one date style', () => {
    expect(FOUNDER_LAST_DAY).toBe('6 Jan 2027');
    expect(listPriceAfterFounder('premium', 'yearly')).toBe('$50 after 6 Jan 2027');
    expect(listPriceAfterFounder('premium', 'monthly')).toBe('$6 a month after 6 Jan 2027');
    expect(listPriceAfterFounder('guild', 'yearly')).toBe('$150 after 6 Jan 2027');
    expect(listPriceAfterFounder('raid_guild', 'monthly')).toBe('$25 a month after 6 Jan 2027');
  });
});

describe('raid sizes and trial reminder (B3.3)', () => {
  it('RAID_SIZES leads with 40', () => expect(RAID_SIZES).toEqual([40, 20, 10]));
  it('TRIAL_REMINDER_DAYS is 3', () => expect(TRIAL_REMINDER_DAYS).toBe(3));
});
