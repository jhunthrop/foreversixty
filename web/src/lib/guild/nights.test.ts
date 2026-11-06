import { describe, expect, it } from 'vitest';
import { nightsLabel } from './nights';
import { attendanceLabel } from './loot-view';

describe('nightsLabel', () => {
  it('is singular when the denominator is 1', () => {
    expect(nightsLabel(1, 1)).toBe('1 of 1 raid night');
  });

  it('is plural for a zero numerator over several nights', () => {
    expect(nightsLabel(0, 3)).toBe('0 of 3 raid nights');
  });

  it('reads No raid nights yet with no eligible night, never 0 of 0', () => {
    expect(nightsLabel(0, 0)).toBe('No raid nights yet');
    expect(attendanceLabel(0, 0)).toBe('No raid nights yet');
  });

  it('is plural for 7 of 8', () => {
    expect(nightsLabel(7, 8)).toBe('7 of 8 raid nights');
  });

  it('is the loot line too', () => {
    expect(attendanceLabel(1, 1)).toBe('1 of 1 raid night');
    expect(attendanceLabel(5, 8)).toBe('5 of 8 raid nights');
  });
});
