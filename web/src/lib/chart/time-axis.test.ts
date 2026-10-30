// web/src/lib/chart/time-axis.test.ts
import { describe, expect, it } from 'vitest';
import { clockLabel, timeStepSeconds, timeTicks } from './time-axis';

describe('clockLabel', () => {
  it('reads m:ss under an hour and h:mm:ss from one hour', () => {
    expect(clockLabel(0)).toBe('0:00');
    expect(clockLabel(5_000)).toBe('0:05');
    expect(clockLabel(127_400)).toBe('2:07');
    expect(clockLabel(340_000)).toBe('5:40');
    expect(clockLabel(3_600_000)).toBe('1:00:00');
    expect(clockLabel(3_725_000)).toBe('1:02:05');
  });

  it('never reads a negative or broken time', () => {
    expect(clockLabel(-5)).toBe('0:00');
    expect(clockLabel(Number.NaN)).toBe('0:00');
  });
});

describe('timeStepSeconds', () => {
  it('is the smallest listed step that keeps labels 64 px apart', () => {
    // 340 s across 1180 px: 5 s is 17 px, 30 s is 104 px; 15 s is 52 px (too close).
    expect(timeStepSeconds(340, 1180)).toBe(30);
    // A zoomed 70 s window at the same width: 5 s is 84 px.
    expect(timeStepSeconds(70, 1180)).toBe(5);
    // A phone: 70 s across 280 px.
    expect(timeStepSeconds(70, 280)).toBe(30);
  });

  it('falls back to the largest step for a very long span', () => {
    expect(timeStepSeconds(10_000, 300)).toBe(300);
  });
});

describe('timeTicks', () => {
  it('puts ticks on whole steps inside the range', () => {
    expect(timeTicks(90, 160, 1180)).toEqual([
      90, 95, 100, 105, 110, 115, 120, 125, 130, 135, 140, 145, 150, 155, 160,
    ]);
    expect(timeTicks(0, 340, 1180)).toEqual([0, 30, 60, 90, 120, 150, 180, 210, 240, 270, 300, 330]);
  });

  it('is empty for an empty range', () => {
    expect(timeTicks(10, 10, 600)).toEqual([]);
  });
});
