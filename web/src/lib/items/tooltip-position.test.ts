// web/src/lib/items/tooltip-position.test.ts
import { describe, expect, it } from 'vitest';
import { tooltipShiftPx } from './tooltip-position';

describe('tooltipShiftPx', () => {
  it('is 0 when the panel already fits between the anchor and the right edge', () => {
    expect(tooltipShiftPx(20, 1000, 288, 8)).toBe(0);
  });

  it('shifts left by exactly enough to clear the viewport margin near the right edge', () => {
    // anchor at 320 on a 360-wide viewport: a 288-wide panel would run off by
    // 320 + 288 + 8 - 360 = 256px, so the shift is -256.
    expect(tooltipShiftPx(320, 360, 288, 8)).toBe(-256);
  });

  it('never shifts right (positive)', () => {
    expect(tooltipShiftPx(-50, 1000, 288, 8)).toBe(0);
  });

  it('defaults to the panel width and margin the real ItemTooltip.svelte uses', () => {
    // Same math as the explicit-args case above, using the exported defaults.
    expect(tooltipShiftPx(320, 360)).toBe(-256);
  });
});
