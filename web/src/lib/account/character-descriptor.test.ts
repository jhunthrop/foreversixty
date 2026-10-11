import { describe, expect, it } from 'vitest';
import { characterDescriptor } from './character-descriptor';
import type { MeCharacter } from './api';

const FULL: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'hunter',
  race: 'Night Elf',
  realm: 'Living Flame',
  level: 25,
};

const BARE: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
};

describe('characterDescriptor', () => {
  it('builds the full "Race Class · Level N · Realm (Ruleset REGION)" line', () => {
    expect(characterDescriptor(FULL)).toBe('Night Elf Hunter · Level 25 · Living Flame (Hardcore US)');
  });

  it('falls back to the bare ruleset/region location when race, class, level and realm are all unknown', () => {
    expect(characterDescriptor(BARE)).toBe('PvP US');
  });
});
