// web/src/lib/chart/brush.test.ts
import { describe, expect, it } from 'vitest';
import { brushSnap } from './brush';

describe('brushSnap', () => {
  it('snaps a dragged edge to the nearest whole second', () => {
    expect(brushSnap(4_400, 9_879)).toBe(4_000);
    expect(brushSnap(4_600, 9_879)).toBe(5_000);
  });

  it('lands on the fight’s exact end, never a second short of it', () => {
    expect(brushSnap(9_600, 9_879)).toBe(9_879);
    expect(brushSnap(9_879, 9_879)).toBe(9_879);
    expect(brushSnap(12_000, 9_879)).toBe(9_879);
  });

  it('never goes below zero', () => {
    expect(brushSnap(-300, 9_879)).toBe(0);
  });
});
