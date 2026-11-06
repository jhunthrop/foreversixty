// web/src/lib/ui/spark.test.ts
import { describe, expect, it } from 'vitest';
import { sparkPoints, sparkTickX } from './spark';

describe('spark geometry', () => {
  it('spreads a series across the box, the top value at the top line', () => {
    expect(sparkPoints([0, 50, 100])).toBe('0.0,22.0 500.0,12.0 1000.0,2.0');
    expect(sparkPoints([40, 65], 100)).toBe('0.0,14.0 1000.0,9.0');
  });

  it('places a tick inside the box', () => {
    expect(sparkTickX(0)).toBe(0);
    expect(sparkTickX(1)).toBe(996);
    expect(sparkTickX(2)).toBe(996);
  });
});
