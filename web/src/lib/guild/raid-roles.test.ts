// web/src/lib/guild/raid-roles.test.ts
import { describe, expect, it } from 'vitest';
import { roleForSpec } from './raid-roles';

describe('roleForSpec', () => {
  it('maps tank, healer and damage specs', () => {
    expect(roleForSpec('Protection')).toBe('tank');
    expect(roleForSpec('Holy')).toBe('healer');
    expect(roleForSpec('Restoration')).toBe('healer');
    expect(roleForSpec('Shadow')).toBe('dps');
    expect(roleForSpec('Fury')).toBe('dps');
  });

  it('is case-blind and null for an unknown or empty spec', () => {
    expect(roleForSpec('beast mastery')).toBe('dps');
    expect(roleForSpec('')).toBeNull();
    expect(roleForSpec('Gladiator')).toBeNull();
  });
});
