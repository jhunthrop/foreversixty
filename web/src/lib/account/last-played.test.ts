import { describe, expect, it } from 'vitest';
import type { MeCharacter } from './api';
import { lastPlayedCharacter } from './last-played';

function character(name: string, capturedAt?: string, level = 60): MeCharacter {
  return {
    key: `us/pvp/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'pvp',
    name,
    class: 'warrior',
    level,
    ...(capturedAt === undefined ? {} : { build: { source: 'addon', captured_at: capturedAt } }),
  };
}

describe('lastPlayedCharacter', () => {
  it('is the character with the newest capture, addon or Battle.net', () => {
    const older = character('Older', '2026-11-08T10:00:00Z');
    const newer = { ...character('Newer'), build: { source: 'blizzard' as const, captured_at: '2026-11-09T10:00:00Z' } };
    expect(lastPlayedCharacter([older, newer], [])?.name).toBe('Newer');
  });

  it('skips a character the site holds no data for', () => {
    const blank = character('Blank', undefined, 60);
    const played = character('Played', '2026-11-01T10:00:00Z', 20);
    expect(lastPlayedCharacter([blank, played], [])?.name).toBe('Played');
  });

  it('is null when no character has any data', () => {
    expect(lastPlayedCharacter([character('A'), character('B')], [])).toBeNull();
  });

  it('breaks a tie by the preferred keys in order, then the higher level', () => {
    const at = '2026-11-09T10:00:00Z';
    const low = character('Low', at, 30);
    const high = character('High', at, 60);
    expect(lastPlayedCharacter([low, high], [])?.name).toBe('High');
    expect(lastPlayedCharacter([low, high], [undefined, low.key])?.name).toBe('Low');
    expect(lastPlayedCharacter([low, high], ['us/pvp/nobody', high.key, low.key])?.name).toBe('High');
  });
});
