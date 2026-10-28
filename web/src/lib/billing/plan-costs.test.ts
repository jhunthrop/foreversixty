// web/src/lib/billing/plan-costs.test.ts
import { describe, expect, it } from 'vitest';
import { FOUNDER_LAST_DAY, type Interval, type PlanKey } from './plans';
import { listPriceAfterFounder, perMemberAmount, perRaiderAmounts, priceFor } from './plan-costs';

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

describe('cost amounts (B3.3)', () => {
  it.each([
    ['yearly', true, ['$5', '$10', '$20']],
    ['yearly', false, ['$6.25', '$12.50', '$25']],
    ['monthly', true, ['$0.50', '$1', '$2']],
    ['monthly', false, ['$0.63', '$1.25', '$2.50']],
  ] as const)('perRaiderAmounts %s founder=%s', (interval, founder, want) => {
    expect(perRaiderAmounts(interval, founder)).toEqual(want);
  });

  it.each([
    ['yearly', true, '$4.80'],
    ['yearly', false, '$6'],
    ['monthly', true, '$0.48'],
    ['monthly', false, '$0.60'],
  ] as const)('perMemberAmount %s founder=%s', (interval, founder, want) => {
    expect(perMemberAmount(interval, founder)).toBe(want);
  });
});

describe('priceFor', () => {
  it.each(TWELVE_AMOUNTS)('%s %s founder=%s is %d', (plan, interval, founder, usd) => {
    expect(priceFor(plan, interval, founder)).toBe(usd);
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
