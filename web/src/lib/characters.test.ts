// web/src/lib/characters.test.ts
import { describe, expect, it } from 'vitest';
import {
  RULESETS,
  characterHref,
  characterKey,
  characterSlug,
  guildHref,
  isRegion,
  isRuleset,
  parseCharacterKey,
  parseCharacterPath,
  parseGuildPath,
  rulesetLabel,
  splitUnitName,
} from './characters';

describe('rulesets and regions', () => {
  it('offers Forever’s four rulesets, which replaced realms', () => {
    expect(RULESETS.map((r) => r.id)).toEqual(['normal', 'pvp', 'rp', 'hardcore']);
    expect(RULESETS.map((r) => r.label)).toEqual(['Normal', 'PvP', 'Roleplay', 'Hardcore']);
  });

  it('labels a ruleset, and echoes anything it does not know', () => {
    expect(rulesetLabel('rp')).toBe('Roleplay');
    expect(rulesetLabel('seasonal')).toBe('seasonal');
  });

  it('guards both enumerations', () => {
    expect(isRuleset('hardcore')).toBe(true);
    expect(isRuleset('nightslayer')).toBe(false);
    expect(isRegion('eu')).toBe(true);
    expect(isRegion('uk')).toBe(false);
  });
});

describe('logged unit names', () => {
  it('splits the trailing segment off a logged name', () => {
    expect(splitUnitName('Elyra Duskvale-Hardcore')).toEqual({ name: 'Elyra Duskvale', segment: 'Hardcore' });
    expect(splitUnitName('Baelgrim-Nightslayer')).toEqual({ name: 'Baelgrim', segment: 'Nightslayer' });
  });

  it('leaves a name with no segment alone', () => {
    expect(splitUnitName('Hollow Sentinel')).toEqual({ name: 'Hollow Sentinel', segment: '' });
    expect(splitUnitName('')).toEqual({ name: '', segment: '' });
  });

  it('splits on the last hyphen, which is the game’s own rule', () => {
    expect(splitUnitName('Anne-Marie Vale-Normal')).toEqual({ name: 'Anne-Marie Vale', segment: 'Normal' });
  });
});

describe('character and guild links', () => {
  it('slugs a two-part name with spaces as hyphens', () => {
    expect(characterSlug('Elyra Duskvale')).toBe('elyra-duskvale');
    expect(characterSlug('  Elyra   Duskvale ')).toBe('elyra-duskvale');
  });

  it('builds the contract’s character key', () => {
    expect(characterKey('us', 'hardcore', 'Elyra Duskvale')).toBe('us/hardcore/elyra-duskvale');
  });

  it('builds the shell routes', () => {
    expect(characterHref('us', 'hardcore', 'Elyra Duskvale')).toBe('/character/us/hardcore/elyra-duskvale');
    expect(guildHref('eu', 'normal', 'The Last Watch')).toBe('/guild/eu/normal/the-last-watch');
  });

  it('reads a character path back, and refuses one that is not a real region and ruleset', () => {
    expect(parseCharacterPath('/character/us/hardcore/elyra-duskvale')).toEqual({
      region: 'us',
      ruleset: 'hardcore',
      slug: 'elyra-duskvale',
    });
    expect(parseCharacterPath('/character/us/hardcore/elyra-duskvale/')).toEqual({
      region: 'us',
      ruleset: 'hardcore',
      slug: 'elyra-duskvale',
    });
    expect(parseCharacterPath('/character/us/nightslayer/elyra-duskvale')).toBeNull();
    expect(parseCharacterPath('/character/us/hardcore')).toBeNull();
    expect(parseCharacterPath('/guild/us/hardcore/the-last-watch')).toBeNull();
  });

  it('reads a guild path back', () => {
    expect(parseGuildPath('/guild/eu/normal/the-last-watch')).toEqual({
      region: 'eu',
      ruleset: 'normal',
      slug: 'the-last-watch',
    });
  });

  it('reads a ranking row’s player.key back into its own region and ruleset', () => {
    expect(parseCharacterKey('eu/normal/other-raider')).toEqual({ region: 'eu', ruleset: 'normal' });
    // Not us/normal, not any default: a key with a real region and ruleset that simply are
    // not the defaults a careless fallback would guess.
    expect(parseCharacterKey('kr/hardcore/other-raider')).toEqual({ region: 'kr', ruleset: 'hardcore' });
  });

  it('refuses a malformed player.key rather than guessing', () => {
    expect(parseCharacterKey('nightslayer/elyra-duskvale')).toBeNull();
    expect(parseCharacterKey('us/nightslayer/elyra-duskvale')).toBeNull();
    expect(parseCharacterKey('')).toBeNull();
  });
});
