// web/src/lib/guild/mark.test.ts
import { describe, expect, it } from 'vitest';
import { guildMarkSrc } from './mark';

describe('guildMarkSrc', () => {
  it('prefers an uploaded crest over the faction logo', () => {
    expect(guildMarkSrc({ crest_url: '/crest/2024.webp', faction: 'horde' })).toBe('/crest/2024.webp');
  });

  it('falls back to the faction logo when there is no crest', () => {
    expect(guildMarkSrc({ crest_url: null, faction: 'alliance' })).toBe(
      '/icons/hd/faction/alliance-logo-512.webp',
    );
    expect(guildMarkSrc({ faction: 'horde' })).toBe('/icons/hd/faction/horde-logo-512.webp');
  });

  it('is null when there is neither a crest nor a known faction', () => {
    expect(guildMarkSrc({ crest_url: null, faction: null })).toBeNull();
    expect(guildMarkSrc({})).toBeNull();
  });

  it('is null for a missing guild', () => {
    expect(guildMarkSrc(null)).toBeNull();
    expect(guildMarkSrc(undefined)).toBeNull();
  });
});
