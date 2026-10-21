import { describe, expect, it } from 'vitest';
import { nightsLabel } from './nights';
import { attendanceLabel } from './loot-view';

describe('nightsLabel', () => {
  it('is singular when the denominator is 1', () => {
    expect(nightsLabel(1, 1)).toBe('1 of 1 night');
  });

  it('is plural for a zero numerator over several nights', () => {
    expect(nightsLabel(0, 3)).toBe('0 of 3 nights');
  });

  it('is plural for 5 of 8', () => {
    expect(nightsLabel(5, 8)).toBe('5 of 8 nights');
  });

  it('keeps the loot lines longer noun through the same rule', () => {
    expect(nightsLabel(1, 1, 'raid night')).toBe('1 of 1 raid night');
    expect(attendanceLabel(1, 1)).toBe('1 of 1 raid night');
    expect(attendanceLabel(5, 8)).toBe('5 of 8 raid nights');
  });
});
