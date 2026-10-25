import { describe, expect, it } from 'vitest';
import {
  axisTicks,
  linearScale,
  parseDomain,
  parseTicks,
  plotRect,
  scoreDomain,
  scoreTicks,
  yScale,
} from './history-scales';

describe('plotRect', () => {
  it('leaves 40 left, 14 right, 12 top and 26 bottom, and 36 left on a phone', () => {
    expect(plotRect({ width: 588, height: 294, phone: false })).toEqual({
      left: 40,
      top: 12,
      width: 534,
      height: 256,
    });
    expect(plotRect({ width: 320, height: 213, phone: true }).left).toBe(36);
  });
});

describe('scoreDomain', () => {
  it('pads 60 and rounds out to whole hundreds', () => {
    expect(scoreDomain([740, 1830])).toEqual({ min: 600, max: 1900 });
  });

  it('is at least 400 wide, centred on the midpoint', () => {
    expect(scoreDomain([1500, 1520])).toEqual({ min: 1300, max: 1700 });
  });

  it('clamps to 0..2100 by shifting, keeping 400 wide', () => {
    expect(scoreDomain([100, 150])).toEqual({ min: 0, max: 400 });
    expect(scoreDomain([2050, 2090])).toEqual({ min: 1700, max: 2100 });
  });

  it('clamps a window already 400 or wider bound by bound', () => {
    expect(scoreDomain([1500, 2090])).toEqual({ min: 1400, max: 2100 });
    expect(scoreDomain([30, 400])).toEqual({ min: 0, max: 500 });
  });

  it('fails fast with no values', () => {
    expect(() => scoreDomain([])).toThrow();
  });
});

describe('scoreTicks', () => {
  it('labels the tier boundaries inside a domain wider than 1200', () => {
    expect(scoreTicks({ min: 400, max: 1900 })).toEqual([500, 1000, 1500]);
    expect(scoreTicks({ min: 0, max: 2100 })).toEqual([0, 500, 1000, 1500, 2000]);
  });

  it('labels every 200 up to a 1200 span and every 100 up to 600', () => {
    expect(scoreTicks({ min: 600, max: 1800 })).toEqual([600, 800, 1000, 1200, 1400, 1600, 1800]);
    expect(scoreTicks({ min: 1300, max: 1700 })).toEqual([1300, 1400, 1500, 1600, 1700]);
  });
});

describe('parseDomain', () => {
  it.each([
    [[12, 90], 0],
    [[29, 90], 0],
    [[30, 90], 25],
    [[54, 90], 25],
    [[55, 90], 50],
    [[79, 99], 50],
    [[80, 99], 75],
  ])('%j starts at %i and ends at 100', (values, min) => {
    expect(parseDomain(values)).toEqual({ min, max: 100 });
  });

  it('depends only on the lowest parse, never on the others', () => {
    expect(parseDomain([60, 70])).toEqual(parseDomain([60, 99, 100]));
    expect(parseDomain([60, 70])).toEqual(parseDomain([60]));
    expect(parseDomain([40, 99])).not.toEqual(parseDomain([60, 99]));
  });
});

describe('parseTicks', () => {
  it('writes 25, 50, 75 and 100 inside the domain, and 0 when it starts there', () => {
    expect(parseTicks({ min: 0, max: 100 })).toEqual([0, 25, 50, 75, 100]);
    expect(parseTicks({ min: 25, max: 100 })).toEqual([25, 50, 75, 100]);
    expect(parseTicks({ min: 75, max: 100 })).toEqual([75, 100]);
  });
});

describe('scales', () => {
  it('puts the domain max at the plot top', () => {
    const plot = { left: 40, top: 12, width: 100, height: 200 };
    const y = yScale({ min: 0, max: 100 }, plot);
    expect(y(100)).toBe(12);
    expect(y(0)).toBe(212);
  });

  it('maps a zero-width domain to its start', () => {
    expect(linearScale({ min: 5, max: 5 }, 10, 20)(5)).toBe(10);
  });
});

describe('axisTicks', () => {
  const first = Date.UTC(2026, 9, 12);
  const last = Date.UTC(2026, 11, 11);

  it('writes up to six labels at least 96 px apart on a desktop', () => {
    const ticks = axisTicks(first, last, 450, false);
    expect(ticks).toHaveLength(5);
    expect(450 / (ticks.length - 1)).toBeGreaterThanOrEqual(96);
    expect(axisTicks(first, last, 534, false)).toHaveLength(6);
  });

  it('thins to three on a phone: first, middle, last', () => {
    const ticks = axisTicks(first, last, 270, true);
    expect(ticks.map((tick) => tick.anchor)).toEqual(['start', 'middle', 'end']);
    expect(ticks[1].time).toBe((first + last) / 2);
  });

  it('keeps two labels when even those are closer than 96 px', () => {
    expect(axisTicks(first, last, 80, true)).toHaveLength(2);
  });

  it('anchors the first at the left edge and the last at the right edge', () => {
    const ticks = axisTicks(first, last, 534, false);
    expect(ticks[0]).toMatchObject({ x: 0, anchor: 'start', time: first });
    expect(ticks[ticks.length - 1]).toMatchObject({ x: 534, anchor: 'end', time: last });
    expect(ticks[1].anchor).toBe('middle');
  });

  it('is one label for a single instant', () => {
    expect(axisTicks(first, first, 500, false)).toHaveLength(1);
  });
});
