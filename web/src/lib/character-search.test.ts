// web/src/lib/character-search.test.ts
import { describe, expect, it } from 'vitest';
import { DEFAULT_REGION, DEFAULT_RULESET, characterSearchHref, pickLookupMatch } from './character-search';

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

describe('pickLookupMatch', () => {
  const pvp = { key: 'us/pvp/frickin-laserbeams', region: 'us', ruleset: 'pvp', name: 'Frickin Laserbeams' };
  const eu = {
    key: 'eu/normal/frickin-laserbeams',
    region: 'eu',
    ruleset: 'normal',
    name: 'Frickin Laserbeams',
  };

  it('opens the character on another ruleset when the form kept the default', () => {
    expect(pickLookupMatch([pvp], 'Frickin Laserbeams', 'us', 'normal')).toEqual(pvp);
  });

  it('takes a first name alone', () => {
    expect(pickLookupMatch([pvp], 'frickin', 'us', 'normal')).toEqual(pvp);
  });

  it('prefers the chosen region, and gives up when several remain', () => {
    expect(pickLookupMatch([pvp, eu], 'Frickin Laserbeams', 'eu', 'pvp')).toEqual(eu);
    expect(pickLookupMatch([pvp, eu], 'Frickin Laserbeams', 'kr', 'normal')).toBeUndefined();
  });

  it('never matches a different name that only shares letters', () => {
    expect(pickLookupMatch([pvp], 'Frick', 'us', 'pvp')).toBeUndefined();
    expect(pickLookupMatch([], 'Frickin', 'us', 'pvp')).toBeUndefined();
  });
});
