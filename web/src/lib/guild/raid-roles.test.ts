// web/src/lib/guild/raid-roles.test.ts
import { describe, expect, it } from 'vitest';
import { SPECS } from '../sim/specs';
import { SPEC_ROLE, roleForSpec } from './raid-roles';

describe('roleForSpec', () => {
  it('maps tank, healer and damage specs', () => {
    expect(roleForSpec('Protection')).toBe('tank');
    expect(roleForSpec('Feral Bear')).toBe('tank');
    expect(roleForSpec('Holy')).toBe('healer');
    expect(roleForSpec('Restoration')).toBe('healer');
    expect(roleForSpec('Shadow')).toBe('dps');
    expect(roleForSpec('Fury')).toBe('dps');
  });

  it('resolves every spec by name and by key, with the role SPECS holds', () => {
    for (const spec of SPECS) {
      expect(roleForSpec(spec.name), spec.spec).toBe(spec.role);
      expect(roleForSpec(spec.spec), spec.spec).toBe(spec.role);
      expect(SPEC_ROLE[spec.name]).toBe(spec.role);
    }
  });

  it('is case-blind and null for an unknown or empty spec', () => {
    expect(roleForSpec('beast mastery')).toBe('dps');
    expect(roleForSpec('')).toBeNull();
    expect(roleForSpec('Gladiator')).toBeNull();
  });
});
