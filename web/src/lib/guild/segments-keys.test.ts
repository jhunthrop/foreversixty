// web/src/lib/guild/segments-keys.test.ts
import { describe, expect, it } from 'vitest';
import { SIGNUP_STATES, segmentKeyAction, tabbableSegment } from './segments-keys';

const move = (index: number): { kind: 'move'; index: number } => ({ kind: 'move', index });
const select = (index: number): { kind: 'select'; index: number } => ({ kind: 'select', index });

describe('segmentKeyAction', () => {
  it('moves forward on ArrowRight and ArrowDown, wrapping at the end, without selecting', () => {
    expect(segmentKeyAction('ArrowRight', 0, 4)).toEqual(move(1));
    expect(segmentKeyAction('ArrowDown', 2, 4)).toEqual(move(3));
    expect(segmentKeyAction('ArrowRight', 3, 4)).toEqual(move(0));
  });

  it('moves back on ArrowLeft and ArrowUp, wrapping at the start, without selecting', () => {
    expect(segmentKeyAction('ArrowLeft', 2, 4)).toEqual(move(1));
    expect(segmentKeyAction('ArrowUp', 1, 4)).toEqual(move(0));
    expect(segmentKeyAction('ArrowLeft', 0, 4)).toEqual(move(3));
  });

  it('selects the focused segment on Space and Enter', () => {
    expect(segmentKeyAction(' ', 2, 4)).toEqual(select(2));
    expect(segmentKeyAction('Enter', 1, 4)).toEqual(select(1));
  });

  it('ignores every other key', () => {
    expect(segmentKeyAction('Tab', 1, 4)).toBeNull();
    expect(segmentKeyAction('a', 1, 4)).toBeNull();
    expect(segmentKeyAction('Escape', 1, 4)).toBeNull();
  });
});

describe('tabbableSegment', () => {
  it('is the checked segment, else the first', () => {
    expect(tabbableSegment('bench')).toBe(SIGNUP_STATES.indexOf('bench'));
    expect(tabbableSegment(null)).toBe(0);
  });
});
