import { describe, expect, it } from 'vitest';
import { heroCharacter, selectedCharacter } from './hero-character';
import type { CurrentCharacter } from '../current-character';
import type { MeCharacter } from './api';

const CHAR: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
  render_url: 'https://render.worldofwarcraft.com/main-raw.png',
};

const POINTER: CurrentCharacter = {
  source: 'armory',
  ref: 'us/pvp/thoradin',
  label: 'Thoradin',
  classSlug: 'warrior',
  savedAt: '2026-09-22T00:00:00Z',
};

describe('heroCharacter', () => {
  it('matches an armory pointer to its character when a render_url is present', () => {
    expect(heroCharacter(POINTER, [CHAR])).toBe(CHAR);
  });

  it('returns null with no pointer', () => {
    expect(heroCharacter(null, [CHAR])).toBeNull();
  });

  it('matches an addon or code pointer by the character name its label carries', () => {
    const named = { ...POINTER, source: 'addon' as const, ref: 'FS1:x', label: 'Thoradin · Fury Warrior' };
    expect(heroCharacter(named, [CHAR])).toBe(CHAR);
    expect(heroCharacter({ ...named, source: 'code' }, [CHAR])).toBe(CHAR);
  });

  it('does not match a pasted pointer whose label is only the class, or a fight or build pointer', () => {
    const nameless = { ...POINTER, source: 'code' as const, ref: 'FS1:x', label: 'Warrior' };
    expect(heroCharacter(nameless, [CHAR])).toBeNull();
    expect(heroCharacter({ ...POINTER, source: 'fight' }, [CHAR])).toBeNull();
    expect(heroCharacter({ ...POINTER, source: 'build' }, [CHAR])).toBeNull();
  });

  it('returns the character even when it has no render_url (an addon-only import)', () => {
    const noRender = { ...CHAR, render_url: undefined };
    expect(heroCharacter(POINTER, [noRender])).toBe(noRender);
  });

  it('returns null when the pointer names no listed character', () => {
    expect(heroCharacter({ ...POINTER, ref: 'us/pvp/somebody-else' }, [CHAR])).toBeNull();
  });
});

describe('selectedCharacter', () => {
  const SHAMAN: MeCharacter = {
    key: 'us/pvp/totemic',
    region: 'us',
    ruleset: 'pvp',
    name: 'Totemic',
    class: 'shaman',
    spec: 'Enhancement',
  };

  it('follows the current-character pointer over the account main and the list order', () => {
    const me = { characters: [SHAMAN, CHAR], main_character_key: SHAMAN.key };
    expect(selectedCharacter(POINTER, me)).toBe(CHAR);
  });

  it('falls back to the account main when there is no pointer', () => {
    const me = { characters: [SHAMAN, CHAR], main_character_key: CHAR.key };
    expect(selectedCharacter(null, me)).toBe(CHAR);
  });

  it('is the pointer’s own identity, never silently the main, when it names no listed character', () => {
    const me = { characters: [SHAMAN, CHAR], main_character_key: CHAR.key };
    expect(selectedCharacter({ ...POINTER, ref: 'us/pvp/somebody-else' }, me)).toBeNull();
    expect(selectedCharacter({ ...POINTER, source: 'fight', ref: 'abc:1' }, me)).toBeNull();
  });

  it('resolves a pasted export that names an account character to that character', () => {
    const me = { characters: [SHAMAN, CHAR], main_character_key: SHAMAN.key };
    const pasted = { ...POINTER, source: 'addon' as const, ref: 'FS1:x', label: 'Thoradin · Fury Warrior' };
    expect(selectedCharacter(pasted, me)).toBe(CHAR);
  });

  it('is null with no characters at all', () => {
    expect(selectedCharacter(POINTER, { characters: [] })).toBeNull();
  });
});
