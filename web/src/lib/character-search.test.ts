// web/src/lib/character-search.test.ts
import { describe, expect, it } from 'vitest';
import { DEFAULT_REGION, DEFAULT_RULESET, characterSearchHref } from './character-search';

describe('characterSearchHref', () => {
  it('routes a name to its profile at the default region and ruleset', () => {
    expect(DEFAULT_REGION).toBe('us');
    expect(DEFAULT_RULESET).toBe('normal');
    expect(characterSearchHref('Elyra Duskvale')).toBe('/character/us/normal/elyra-duskvale');
  });

  it('takes the ruleset from a -Realm suffix, the way the addon writes Name-Realm', () => {
    expect(characterSearchHref('Elyra Duskvale-Hardcore')).toBe('/character/us/hardcore/elyra-duskvale');
    expect(characterSearchHref('  Elyra Duskvale-PvP ', 'eu')).toBe('/character/eu/pvp/elyra-duskvale');
  });

  it('keeps a hyphen that names no ruleset as part of the name', () => {
    expect(characterSearchHref('Jean-Luc')).toBe('/character/us/normal/jean-luc');
  });

  it('uses the region and ruleset picked beside the box', () => {
    expect(characterSearchHref('Elyra Duskvale', 'eu', 'rp')).toBe('/character/eu/rp/elyra-duskvale');
  });

  it('is null for nothing typed, or for a name no character could have', () => {
    expect(characterSearchHref('   ')).toBeNull();
    expect(characterSearchHref('../account')).toBeNull();
    expect(characterSearchHref('a/b')).toBeNull();
  });
});
