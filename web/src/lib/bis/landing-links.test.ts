// web/src/lib/bis/landing-links.test.ts
import { describe, expect, it } from 'vitest';
import { specLinkHref } from './landing-links';

describe('specLinkHref', () => {
  it('opens the first band on the spec’s own page, with no faction signed out', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 20 })).toBe('/bis/warrior/fury');
  });

  it('opens a later band at its own page, the character’s Horde faction in the query', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 40, faction: 'horde' })).toBe(
      '/bis/warrior/fury/40-49?faction=horde',
    );
  });

  it('writes the top band as 60 and leaves the default faction out', () => {
    expect(specLinkHref({ classSlug: 'druid', specSlug: 'feral-bear', band: 60, faction: 'alliance' })).toBe(
      '/bis/druid/feral-bear/60',
    );
  });
});
