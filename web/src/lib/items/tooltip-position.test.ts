// web/src/lib/items/tooltip-position.test.ts
import { describe, expect, it } from 'vitest';
import { tooltipShiftPx, tooltipVerticalPlacementFor } from './tooltip-position';

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

describe('tooltipVerticalPlacementFor', () => {
  it('opens below, unclamped, when the panel fits between the anchor and the bottom edge', () => {
    // Anchor bottom at 200 on an 800-tall viewport: 800 - 200 - 8 (gap) - 8 (margin) = 584px
    // of room below, plenty for a 300px panel.
    expect(tooltipVerticalPlacementFor(150, 200, 300, 800)).toEqual({ side: 'below', maxHeightPx: null });
  });

  it('flips above, unclamped, when below cannot fit the panel but above can', () => {
    // Anchor near the bottom of a short viewport: 400 - 780 - 16 is negative below, but
    // 780 - 16 = 764px of room above easily fits a 300px panel.
    expect(tooltipVerticalPlacementFor(780, 800, 300, 800)).toEqual({ side: 'above', maxHeightPx: null });
  });

  it('clamps to the larger side and reports a max height when neither side fits', () => {
    // A 700px-tall panel on an anchor roughly in the middle of a 400px-tall viewport: neither
    // 400 - 220 - 16 = 164px below nor 200 - 16 = 184px above fits it, so this picks the
    // larger (above, 184px) and clamps to it.
    expect(tooltipVerticalPlacementFor(200, 220, 700, 400)).toEqual({ side: 'above', maxHeightPx: 184 });
  });

  it('prefers below on an exact tie between the two clamped sides', () => {
    // A zero-height anchor dead center of a 400px-tall viewport: 184px of room on both
    // sides (200 - 8 - 8), neither enough for a 900px panel -- an exact tie prefers below.
    expect(tooltipVerticalPlacementFor(200, 200, 900, 400)).toEqual({ side: 'below', maxHeightPx: 184 });
  });

  it('never reports a negative clamped height even when both sides are fully out of room', () => {
    const result = tooltipVerticalPlacementFor(0, 0, 500, 4);
    expect(result.maxHeightPx).not.toBeNull();
    expect(result.maxHeightPx!).toBeGreaterThanOrEqual(0);
  });
});
