// web/src/lib/sim/rotation-card.test.ts
import { describe, expect, it } from 'vitest';
import { obviousNote } from './rotation-card';

describe('obviousNote', () => {
  it('is a note that only says to press it on cooldown', () => {
    expect(obviousNote('On cooldown: the button you lose the most by delaying.')).toBe(true);
    expect(obviousNote('on cooldown, right behind Whirlwind')).toBe(true);
  });

  it('keeps a note with a condition', () => {
    expect(obviousNote('Whenever rage is under 80: 10 rage at once.')).toBe(false);
    expect(obviousNote('From 36, on cooldown under 80 rage')).toBe(false);
    expect(obviousNote('')).toBe(false);
  });
});
