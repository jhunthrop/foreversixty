// web/src/lib/bis/landing-links.test.ts
import { describe, expect, it } from 'vitest';
import { landingSpecHref, specLinkHref } from './landing-links';

describe('specLinkHref', () => {
  it('opens the first band by its hash on the spec’s own page, with no faction signed out', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 20 })).toBe(
      '/bis/warrior/fury#levels-20-29',
    );
  });

  it('opens a later band at its own page, the character’s Horde faction in the query', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 40, faction: 'horde' })).toBe(
      '/bis/warrior/fury?faction=horde#levels-40-49',
    );
  });

  it('writes the top band as 60 and leaves the default faction out', () => {
    expect(specLinkHref({ classSlug: 'druid', specSlug: 'feral-bear', band: 60, faction: 'alliance' })).toBe(
      '/bis/druid/feral-bear/60',
    );
  });
});

describe('landingSpecHref', () => {
  it('opens level 60 on the spec’s own page, which shows 60 first', () => {
    expect(landingSpecHref({ classSlug: 'warrior', specSlug: 'fury', band: 60 })).toBe('/bis/warrior/fury');
    expect(landingSpecHref({ classSlug: 'warrior', specSlug: 'fury', band: 60, faction: 'horde' })).toBe(
      '/bis/warrior/fury?faction=horde',
    );
  });

  it('links a leveling band like specLinkHref', () => {
    expect(landingSpecHref({ classSlug: 'warrior', specSlug: 'fury', band: 40 })).toBe(
      '/bis/warrior/fury#levels-40-49',
    );
  });
});
