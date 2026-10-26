// web/src/lib/guild/segments-keys.test.ts
import { describe, expect, it } from 'vitest';
import { SIGNUP_STATES, segmentKeyTarget, tabbableSegment } from './segments-keys';

describe('segmentKeyTarget', () => {
  it('moves forward on ArrowRight and ArrowDown, wrapping at the end', () => {
    expect(segmentKeyTarget('ArrowRight', 0, 4)).toBe(1);
    expect(segmentKeyTarget('ArrowDown', 2, 4)).toBe(3);
    expect(segmentKeyTarget('ArrowRight', 3, 4)).toBe(0);
  });

  it('moves back on ArrowLeft and ArrowUp, wrapping at the start', () => {
    expect(segmentKeyTarget('ArrowLeft', 2, 4)).toBe(1);
    expect(segmentKeyTarget('ArrowUp', 1, 4)).toBe(0);
    expect(segmentKeyTarget('ArrowLeft', 0, 4)).toBe(3);
  });

  it('selects the focused segment on Space and Enter', () => {
    expect(segmentKeyTarget(' ', 2, 4)).toBe(2);
    expect(segmentKeyTarget('Enter', 1, 4)).toBe(1);
  });

  it('ignores every other key', () => {
    expect(segmentKeyTarget('Tab', 1, 4)).toBeNull();
    expect(segmentKeyTarget('a', 1, 4)).toBeNull();
    expect(segmentKeyTarget('Escape', 1, 4)).toBeNull();
  });
});

describe('tabbableSegment', () => {
  it('is the checked segment, else the first', () => {
    expect(tabbableSegment('bench')).toBe(SIGNUP_STATES.indexOf('bench'));
    expect(tabbableSegment(null)).toBe(0);
  });
});
