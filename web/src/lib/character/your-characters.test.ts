// /characters' "Your characters" rail (directions mock b2-characters).
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { yourCharacterRows } from './your-characters';

const thandrel: MeCharacter = {
  key: 'us/normal/thandrel',
  region: 'us',
  ruleset: 'normal',
  name: 'Thandrel',
  class: 'Warrior',
  spec: 'Fury',
  level: 60,
} as MeCharacter;

describe('yourCharacterRows', () => {
  it('names the level, spec and class, in the class colour, linked to the page', () => {
    expect(yourCharacterRows([thandrel])).toEqual([
      {
        key: 'us/normal/thandrel',
        name: 'Thandrel',
        tone: { classSlug: 'warrior' },
        classSlug: 'warrior',
        line: 'Level 60 Fury Warrior',
        href: '/character/us/normal/thandrel',
      },
    ]);
  });

  it('adds the guild and the item level the account already has', () => {
    const geared = {
      ...thandrel,
      item_level: 61,
      guild: { id: 2, name: 'Cx66 Olympus', verified: true },
    } as MeCharacter;
    expect(yourCharacterRows([geared])[0]?.line).toBe(
      'Level 60 Fury Warrior · <Cx66 Olympus> · Item level 61',
    );
  });

  it('leaves out what the Battle.net import has not learned', () => {
    const bare = {
      key: 'eu/hardcore/felsnap',
      region: 'eu',
      ruleset: 'hardcore',
      name: 'Felsnap',
    } as MeCharacter;
    expect(yourCharacterRows([bare])[0]).toMatchObject({ tone: 'strong', classSlug: '', line: '' });
  });
});
