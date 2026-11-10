// web/src/lib/sim/rotation-card.test.ts
import { describe, expect, it } from 'vitest';
import { obviousNote } from './rotation-card';

describe('obviousNote', () => {
  it('is a note that says to press it on cooldown, or restates a whenever tooltip', () => {
    expect(obviousNote('On cooldown: the button you lose the most by delaying.')).toBe(true);
    expect(obviousNote('From 36, on cooldown under 80 rage')).toBe(true);
    expect(obviousNote('Whenever rage is under 80: 10 rage at once.')).toBe(true);
    expect(obviousNote('Whenever it drops: 10 rage for the melee attack power.')).toBe(true);
  });

  it('keeps a note that changes when to press', () => {
    expect(obviousNote('In the Execute phase, or in the last 15 seconds when that phase is short.')).toBe(
      false,
    );
    expect(obviousNote('Ice Barrier whenever the shield is down.')).toBe(false);
    expect(obviousNote('')).toBe(false);
  });
});
