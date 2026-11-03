// web/src/lib/guides/role-metric.test.ts
import { describe, expect, it } from 'vitest';
import { roleOf } from './role-metric';

describe('roleOf', () => {
  it('defaults a band with no role to dps', () => {
    expect(roleOf({})).toBe('dps');
    expect(roleOf({ role: 'tank' })).toBe('tank');
  });
});
