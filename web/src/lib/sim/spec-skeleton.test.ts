// web/src/lib/sim/spec-skeleton.test.ts
// spec-skeleton.ts is the one source both sim/specs.astro's static shell and
// SpecGrid.svelte read for the specs grid's shape (see that module's own header comment),
// so its two slot functions are the load-bearing contract for the "27 cards, 20 simulated
// and 7 not" split task-2-brief.md calls for: a drift here is a CLS regression there.
import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { dpsSpecs, nonDpsSpecs } from './spec-label';
import { SPECS } from './specs';
import { specSkeletonSlots, unsimulatedSpecSkeletonSlots } from './spec-skeleton';

describe('specSkeletonSlots', () => {
  it('has one slot per dps spec, as a plain 0..n-1 index sequence', () => {
    const slots = specSkeletonSlots();
    expect(slots).toHaveLength(dpsSpecs().length);
    expect(slots).toEqual(Array.from({ length: dpsSpecs().length }, (_, index) => index));
  });
});

describe('unsimulatedSpecSkeletonSlots', () => {
  it('has one slot per non-dps (healer/tank) spec, as a plain 0..n-1 index sequence', () => {
    const slots = unsimulatedSpecSkeletonSlots();
    expect(slots).toHaveLength(nonDpsSpecs().length);
    expect(slots).toEqual(Array.from({ length: nonDpsSpecs().length }, (_, index) => index));
  });
});

describe('the /sim/specs grid split', () => {
  // The reviewers' repro: 27 specs total, 20 rendered as real cards, 7 rendered as a
  // distinct "not simulated yet" card -- never fewer (a spec silently missing) and never
  // more (a spec counted twice).
  it('is 27 cards total: 20 simulated, 7 not, and every spec appears exactly once', () => {
    expect(dpsSpecs().length).toBe(20);
    expect(nonDpsSpecs().length).toBe(7);
    expect(dpsSpecs().length + nonDpsSpecs().length).toBe(SPECS.length);
  });

  it('is exactly the copy task-2-brief.md specifies for the unsimulated card', () => {
    expect(simCopy.specsUnsimulatedBody).toBe(
      'Not simulated yet — damage specs first; healers and tanks come later',
    );
  });
});
