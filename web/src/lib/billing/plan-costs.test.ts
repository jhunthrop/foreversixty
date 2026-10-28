// web/src/lib/billing/plan-costs.test.ts
import { describe, expect, it } from 'vitest';
import { perMemberAmount, perRaiderAmounts } from './plan-costs';

describe('cost amounts (B3.3)', () => {
  it.each([
    ['yearly', true, ['$5', '$10', '$20']],
    ['yearly', false, ['$6.25', '$12.50', '$25']],
    ['monthly', true, ['$0.50', '$1', '$2']],
    ['monthly', false, ['$0.63', '$1.25', '$2.50']],
  ] as const)('perRaiderAmounts %s founder=%s', (interval, founder, want) => {
    expect(perRaiderAmounts(interval, founder)).toEqual(want);
  });

  it('perRaiderAmounts monthly list rounds $0.625 to $0.63', () => {
    expect(perRaiderAmounts('monthly', false)[0]).toBe('$0.63');
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
