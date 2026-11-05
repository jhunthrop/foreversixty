// web/src/lib/guild/raid-roles.test.ts
import { describe, expect, it } from 'vitest';
import { SPECS } from '../sim/specs';
import { SPEC_ROLE, roleForSpec, rolesForClass } from './raid-roles';

describe('rolesForClass', () => {
  it('offers only the roles the class plays in Classic', () => {
    expect(rolesForClass('warrior')).toEqual(['tank', 'dps']);
    expect(rolesForClass('Druid')).toEqual(['tank', 'healer', 'dps']);
    expect(rolesForClass('paladin')).toEqual(['tank', 'healer', 'dps']);
    expect(rolesForClass('priest')).toEqual(['healer', 'dps']);
    expect(rolesForClass('shaman')).toEqual(['healer', 'dps']);
    expect(rolesForClass('mage')).toEqual(['dps']);
    expect(rolesForClass('')).toEqual(['tank', 'healer', 'dps']);
  });
});

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
