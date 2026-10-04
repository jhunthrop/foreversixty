// web/src/lib/guides/signed-in-callout.test.ts
import { describe, expect, it } from 'vitest';
import { signedInCalloutView } from './signed-in-callout';
import type { MeCharacter } from '../account/api';

function character(overrides: Partial<MeCharacter> = {}): MeCharacter {
  return { key: 'k', region: 'us', ruleset: 'era', name: 'Obnoxious', ...overrides };
}

describe('signedInCalloutView', () => {
  it('builds the guide link and callout text for a character with a known class and spec', () => {
    const view = signedInCalloutView(character({ class: 'warrior', spec: 'Fury' }));
    expect(view).toEqual({
      href: '/guides/warrior/fury',
      crestSrc: '/icons/hd/crests/warrior.webp',
      crestColorVar: 'var(--color-class-warrior)',
      text: 'Your guide: Fury Warrior →',
    });
  });

  it('kebab-cases a multi-word spec name into its guide slug', () => {
    const view = signedInCalloutView(character({ class: 'hunter', spec: 'Beast Mastery' }));
    expect(view?.href).toBe('/guides/hunter/beast-mastery');
  });

  it('is undefined with no character at all', () => {
    expect(signedInCalloutView(undefined)).toBeUndefined();
  });

  it('is undefined when class or spec has not been learned yet', () => {
    expect(signedInCalloutView(character({ class: 'warrior' }))).toBeUndefined();
    expect(signedInCalloutView(character({ spec: 'Fury' }))).toBeUndefined();
  });
});
