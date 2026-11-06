// web/src/lib/guild/segments-keys.test.ts
import { describe, expect, it } from 'vitest';
import { SIGNUP_STATES, tabbableSegment } from './segments-keys';

describe('tabbableSegment', () => {
  it('is the checked segment, else the first', () => {
    expect(tabbableSegment('bench')).toBe(SIGNUP_STATES.indexOf('bench'));
    expect(tabbableSegment(null)).toBe(0);
  });
});
