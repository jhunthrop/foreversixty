import { describe, expect, it } from 'vitest';
import { coachRulesReview } from '../../test-support/coach-fixtures';
import { SLOT_RESERVE, mergeReserve, reserveFor, slotClass } from './reserves';

describe('slot reserves', () => {
  it('pins the spec table: lines x 22 px, phone then desktop', () => {
    expect(SLOT_RESERVE.opening).toBe('min-h-[154px] md:min-h-[66px]');
    expect(SLOT_RESERVE.why).toBe('min-h-[176px] md:min-h-[88px]');
    expect(SLOT_RESERVE.drill).toBe('min-h-[132px] md:min-h-[66px]');
    expect(SLOT_RESERVE.thenWhy).toBe('min-h-[110px] md:min-h-[66px]');
    expect(SLOT_RESERVE.well).toBe('min-h-[88px] md:min-h-[44px]');
    expect(SLOT_RESERVE.chips).toBe('min-h-[94px] md:min-h-[36px]');
    expect(SLOT_RESERVE.links).toBe('min-h-[80px] md:min-h-[40px]');
  });
  it('reserves only while held', () => {
    expect(slotClass('opening', true)).toBe(SLOT_RESERVE.opening);
    expect(slotClass('opening', false)).toBe('');
  });
});

describe('reserveFor and mergeReserve', () => {
  it('counts the rule-only rendering items', () => {
    const review = coachRulesReview().review;
    expect(reserveFor(review)).toEqual({ thenCount: 1, wellCount: 1, links: [false, false] });
  });
  it('never shrinks', () => {
    const big = { thenCount: 3, wellCount: 2, links: [true, false, true] };
    const small = { thenCount: 1, wellCount: 1, links: [false] };
    expect(mergeReserve(big, small)).toEqual(big);
    expect(mergeReserve(null, small)).toBe(small);
  });
});
