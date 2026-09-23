import { describe, expect, it } from 'vitest';
import { heroCharacter } from './hero-character';
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

  it('returns null when the pointer is not an armory source', () => {
    expect(heroCharacter({ ...POINTER, source: 'addon' }, [CHAR])).toBeNull();
  });

  it('returns the character even when it has no render_url (an addon-only import)', () => {
    const noRender = { ...CHAR, render_url: undefined };
    expect(heroCharacter(POINTER, [noRender])).toBe(noRender);
  });

  it('returns null when the pointer names no listed character', () => {
    expect(heroCharacter({ ...POINTER, ref: 'us/pvp/somebody-else' }, [CHAR])).toBeNull();
  });
});
