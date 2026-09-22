// web/src/lib/billing/plans.test.ts
import { describe, expect, it } from 'vitest';
import { PLANS, formatPrice } from './plans';

describe('PLANS', () => {
  it('has the two plans at the decided prices', () => {
    expect(PLANS.premium.monthly).toBe(4);
    expect(PLANS.premium.yearly).toBe(40);
    expect(PLANS.guild.monthly).toBe(15);
    expect(PLANS.guild.yearly).toBe(150);
    expect(PLANS.guild.coversGuild).toBe(true);
    expect(PLANS.premium.coversGuild).toBe(false);
  });
});

describe('formatPrice', () => {
  it('formats a monthly and a yearly amount', () => {
    expect(formatPrice(4, 'monthly')).toBe('$4/month');
    expect(formatPrice(40, 'yearly')).toBe('$40/year');
  });
});
