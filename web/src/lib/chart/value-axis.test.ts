// web/src/lib/chart/value-axis.test.ts
import { describe, expect, it } from 'vitest';
import { LOG_FLOOR, logTicks, niceTop, valueLabel, valueTicks } from './value-axis';

describe('niceTop', () => {
  it('is the smallest of 1, 1.5, 2, 2.5, 3, 4, 5 or 10 times a power of ten that holds the peak', () => {
    expect(niceTop(1150)).toBe(1500);
    expect(niceTop(1200)).toBe(1500);
    expect(niceTop(160_000)).toBe(200_000);
    expect(niceTop(140_000)).toBe(150_000);
    expect(niceTop(380)).toBe(400);
    expect(niceTop(1000)).toBe(1000);
    expect(niceTop(9)).toBe(10);
  });

  it('is 1 for an empty or zero chart', () => {
    expect(niceTop(0)).toBe(1);
    expect(niceTop(-4)).toBe(1);
  });
});

describe('valueTicks', () => {
  it('draws 4 or 5 gridlines from zero to the top', () => {
    expect(valueTicks(1200)).toEqual([0, 300, 600, 900, 1200]);
    expect(valueTicks(1500)).toEqual([0, 500, 1000, 1500]);
    expect(valueTicks(160_000)).toEqual([0, 40_000, 80_000, 120_000, 160_000]);
    expect(valueTicks(400)).toEqual([0, 100, 200, 300, 400]);
  });
});

describe('valueLabel', () => {
  it('reads like the spec: 0, 400, 1.2k, 60k, 1.2M', () => {
    expect(valueLabel(0)).toBe('0');
    expect(valueLabel(400)).toBe('400');
    expect(valueLabel(1200)).toBe('1.2k');
    expect(valueLabel(60_000)).toBe('60k');
    expect(valueLabel(1_200_000)).toBe('1.2M');
    expect(valueLabel(1000)).toBe('1k');
  });
});

describe('logTicks', () => {
  it('ticks at powers of ten from the floor', () => {
    expect(LOG_FLOOR).toBe(1);
    expect(logTicks(1200)).toEqual([1, 10, 100, 1000]);
    expect(logTicks(5)).toEqual([1]);
  });
});
