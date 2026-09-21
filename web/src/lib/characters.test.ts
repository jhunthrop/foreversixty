// web/src/lib/characters.test.ts
import { describe, expect, it } from 'vitest';
import {
  RULESETS,
  characterHref,
  characterKey,
  characterSlug,
  guildClaimHref,
  guildHref,
  guildInviteHref,
  guildSettingsHref,
  isRegion,
  isRuleset,
  parseCharacterKey,
  parseCharacterPath,
  parseGuildClaimPath,
  parseGuildInviteToken,
  parseGuildPath,
  parseGuildSettingsPath,
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

  // The slug is interpolated straight into `/v1/characters/<region>/<ruleset>/<slug>`, and
  // a dot segment there is resolved by the URL parser rather than 404ing: `..` would make
  // the island, and the Worker's own unfurl fetch, call a different endpoint entirely. The
  // percent-encoded spellings resolve the same way, so they are refused the same way.
  it('refuses a slug that could name something other than a character', () => {
    const refused = [
      '..',
      '.',
      '%2e%2e',
      '%2E%2E',
      '%2f%2e%2e',
      'elyra%2fduskvale',
      'elyra duskvale',
      'elyra?x=1',
      'elyra#x',
      'a'.repeat(65),
    ];
    for (const slug of refused) {
      expect(parseCharacterPath(`/character/us/normal/${slug}`)).toBeNull();
      expect(parseGuildPath(`/guild/us/normal/${slug}`)).toBeNull();
    }
  });

  // Not `[a-z0-9-]`, which is what /rankings/<slug> allows: kr, tw and cn are regions this
  // site serves, and a pathname reaches the parser percent-encoded, so a real character
  // page for a Hangul name is nothing but `%`-escapes and has to keep working.
  it('accepts a percent-encoded name from a region that does not write in Latin', () => {
    const encoded = encodeURIComponent('아무개');
    expect(parseCharacterPath(`/character/kr/normal/${encoded}`)).toEqual({
      region: 'kr',
      ruleset: 'normal',
      slug: encoded,
    });
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

describe('parseGuildClaimPath', () => {
  it('parses a claim sub-path', () => {
    expect(parseGuildClaimPath('/guild/us/hardcore/the-last-watch/claim')).toEqual({
      region: 'us',
      ruleset: 'hardcore',
      slug: 'the-last-watch',
    });
  });

  it('refuses the plain guild path (no suffix) and a settings path', () => {
    expect(parseGuildClaimPath('/guild/us/hardcore/the-last-watch')).toBeNull();
    expect(parseGuildClaimPath('/guild/us/hardcore/the-last-watch/settings')).toBeNull();
  });
});

describe('parseGuildSettingsPath', () => {
  it('parses a settings sub-path', () => {
    expect(parseGuildSettingsPath('/guild/eu/normal/iron-vanguard/settings')).toEqual({
      region: 'eu',
      ruleset: 'normal',
      slug: 'iron-vanguard',
    });
  });

  it('refuses an invalid region or ruleset the same way parseGuildPath does', () => {
    expect(parseGuildSettingsPath('/guild/xx/hardcore/name/settings')).toBeNull();
  });
});

describe('parseGuildInviteToken', () => {
  it('reads the token segment', () => {
    expect(parseGuildInviteToken('/guild/invite/abc-123')).toBe('abc-123');
  });

  it('refuses a token with a path separator or empty', () => {
    expect(parseGuildInviteToken('/guild/invite/')).toBeNull();
    expect(parseGuildInviteToken('/guild/invite/a/b')).toBeNull();
    expect(parseGuildInviteToken('/guild/us/hardcore/name')).toBeNull();
  });
});

describe('guild sub-route hrefs', () => {
  it('builds claim, settings and invite links', () => {
    expect(guildClaimHref('us', 'hardcore', 'The Last Watch')).toBe(
      '/guild/us/hardcore/the-last-watch/claim',
    );
    expect(guildSettingsHref('us', 'hardcore', 'The Last Watch')).toBe(
      '/guild/us/hardcore/the-last-watch/settings',
    );
    expect(guildInviteHref('abc 123')).toBe('/guild/invite/abc%20123');
  });
});
