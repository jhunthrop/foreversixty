// web/src/lib/bis/landing-config.test.ts
import { describe, expect, it } from 'vitest';
import { classFor, specKeyFor, type LandingConfig } from './landing-config';

const CONFIG: LandingConfig = {
  build: 'b1',
  classes: {
    warrior: { slug: 'warrior', name: 'Warrior', color: '#c69b6d', defaultSpecKey: 'warrior-arms' },
    hunter: { slug: 'hunter', name: 'Hunter', color: '#aad372', defaultSpecKey: 'hunter-beast-mastery' },
  },
  specs: {
    'warrior-arms': { classSlug: 'warrior', specSlug: 'arms', name: 'Arms' },
    'warrior-fury': { classSlug: 'warrior', specSlug: 'fury', name: 'Fury' },
    'hunter-beast-mastery': { classSlug: 'hunter', specSlug: 'beast-mastery', name: 'Beast Mastery' },
  },
};

describe('classFor', () => {
  it('reads the class name /v1/me reports, in any case', () => {
    expect(classFor(CONFIG, 'Warrior')?.slug).toBe('warrior');
    expect(classFor(CONFIG, 'hunter')?.slug).toBe('hunter');
  });

  it('is undefined for a class the site does not know or none', () => {
    expect(classFor(CONFIG, 'Monk')).toBeUndefined();
    expect(classFor(CONFIG, undefined)).toBeUndefined();
  });
});

describe('specKeyFor', () => {
  it('finds the spec by class and spec name', () => {
    expect(specKeyFor(CONFIG, 'Warrior', 'Fury')).toBe('warrior-fury');
    expect(specKeyFor(CONFIG, 'Hunter', 'Beast Mastery')).toBe('hunter-beast-mastery');
  });

  it('never matches a spec name under another class', () => {
    expect(specKeyFor(CONFIG, 'Hunter', 'Fury')).toBeUndefined();
  });

  it('is undefined when the class or the spec is not known yet', () => {
    expect(specKeyFor(CONFIG, 'Warrior', undefined)).toBeUndefined();
    expect(specKeyFor(CONFIG, undefined, 'Fury')).toBeUndefined();
  });
});
