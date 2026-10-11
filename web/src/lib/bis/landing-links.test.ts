// web/src/lib/bis/landing-links.test.ts
import { describe, expect, it } from 'vitest';
import { specLinkHref } from './landing-links';

describe('specLinkHref', () => {
  it('carries no faction signed out, and opens the alliance band the spec page defaults to', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 20 })).toBe(
      '/bis/warrior/fury#band-alliance-20',
    );
  });

  it('carries the character’s faction in the query and the band hash', () => {
    expect(specLinkHref({ classSlug: 'warrior', specSlug: 'fury', band: 40, faction: 'horde' })).toBe(
      '/bis/warrior/fury?faction=horde#band-horde-40',
    );
  });

  it('writes the top band as 60', () => {
    expect(specLinkHref({ classSlug: 'druid', specSlug: 'feral-bear', band: 60, faction: 'alliance' })).toBe(
      '/bis/druid/feral-bear?faction=alliance#band-alliance-60',
    );
  });
});
