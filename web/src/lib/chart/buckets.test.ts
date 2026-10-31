// web/src/lib/chart/buckets.test.ts
import { describe, expect, it } from 'vitest';
import { bucketMeans, chartBucket, MAX_POINTS, valueAt } from './buckets';
import { timeStepSeconds } from './time-axis';

const S = 1000;
const MIN = 60 * S;

describe('chartBucket', () => {
  it('draws a short pull at one point per second, labelled every 2 s', () => {
    for (const durationMs of [5 * S, 10 * S]) {
      const bucket = chartBucket(durationMs);
      expect(bucket.seconds).toBe(1);
      expect(bucket.labelFloorS).toBe(2);
    }
    // The 10 s pull across a desktop plot (1180 px) and a phone strip (330 px): a label every 2 s.
    expect(timeStepSeconds(10, 1180, chartBucket(10 * S).labelFloorS)).toBe(2);
    expect(timeStepSeconds(10, 330, chartBucket(10 * S).labelFloorS)).toBe(2);
  });

  it('widens to 2 s on a minute, so the line has 30 points and not 12', () => {
    expect(chartBucket(60 * S)).toEqual({ seconds: 2, labelFloorS: 2 });
  });

  it('keeps the 5 s average on a long fight', () => {
    expect(chartBucket(4 * MIN)).toEqual({ seconds: 5, labelFloorS: 5 });
    expect(chartBucket(15 * MIN)).toEqual({ seconds: 5, labelFloorS: 5 });
  });

  it('never draws more than 300 points a series', () => {
    for (const durationMs of [5 * S, 10 * S, 60 * S, 4 * MIN, 15 * MIN, 30 * MIN, 60 * MIN, 3 * 60 * MIN]) {
      const { seconds } = chartBucket(durationMs);
      expect(Math.ceil(durationMs / S / seconds)).toBeLessThanOrEqual(MAX_POINTS);
    }
    expect(chartBucket(30 * MIN)).toEqual({ seconds: 10, labelFloorS: 10 });
  });

  it('is one second for an empty or broken duration', () => {
    expect(chartBucket(0).seconds).toBe(1);
    expect(chartBucket(Number.NaN).seconds).toBe(1);
  });
});

describe('bucketMeans', () => {
  it('is each bucket’s mean per second, the last bucket over its own seconds', () => {
    expect(bucketMeans([10, 20, 30, 40, 50], 2, 5)).toEqual([15, 35, 50]);
  });

  it('counts the seconds past a series’ last busy one as zero', () => {
    // The summary's series stops at its last busy second; the fight runs 6 s.
    expect(bucketMeans([10, 20, 30], 2, 6)).toEqual([15, 15, 0]);
  });

  it('is the series itself at one second', () => {
    expect(bucketMeans([1, 2, 3], 1, 3)).toEqual([1, 2, 3]);
  });

  it('falls to zero and stays flat through a phase where the boss takes no damage', () => {
    const series = [100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 50];
    expect(bucketMeans(series, 5, 15)).toEqual([100, 0, 10]);
  });
});

describe('valueAt', () => {
  it('reads the bucket a moment falls in', () => {
    const line = { values: [1, 2, 3], stepMs: 5000 };
    expect(valueAt(line, 0)).toBe(1);
    expect(valueAt(line, 4999)).toBe(1);
    expect(valueAt(line, 5000)).toBe(2);
    expect(valueAt(line, 20_000)).toBeUndefined();
  });

  it('reads one value per second when the line names no step', () => {
    expect(valueAt({ values: [1, 2, 3] }, 2500)).toBe(3);
  });
});
