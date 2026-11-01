import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import type { CurrentCharacter } from '../current-character';
import { characterFollowHref, guildFollowHref, pointerCharacterHref } from './follow-page';

const millennials = { name: 'Employed Millennials', href: '/guild/us/pvp/employed-millennials' };
const olympus = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };

describe('guildFollowHref', () => {
  const onPage = (pathname: string, hash = '') => ({ pathname, hash });

  it('moves your guild page to the new character’s guild, keeping the tab', () => {
    expect(
      guildFollowHref({
        ...onPage('/guild/us/pvp/employed-millennials', '#roster'),
        previous: millennials,
        next: olympus,
      }),
    ).toBe('/guild/us/pvp/olympus-xxvii#roster');
  });

  it('keeps the settings and claim subpages', () => {
    const base = { hash: '', previous: millennials, next: olympus };
    expect(guildFollowHref({ ...base, pathname: '/guild/us/pvp/employed-millennials/settings' })).toBe(
      '/guild/us/pvp/olympus-xxvii/settings',
    );
    expect(guildFollowHref({ ...base, pathname: '/guild/us/pvp/employed-millennials/claim' })).toBe(
      '/guild/us/pvp/olympus-xxvii/claim',
    );
  });

  it('opens the new guild’s page from a raid, whose id is the old guild’s', () => {
    expect(
      guildFollowHref({
        ...onPage('/guild/us/pvp/employed-millennials/raids/12', '#signups'),
        previous: millennials,
        next: olympus,
      }),
    ).toBe('/guild/us/pvp/olympus-xxvii');
  });

  it('sends a guildless new character to the find-or-claim landing', () => {
    expect(
      guildFollowHref({
        ...onPage('/guild/us/pvp/employed-millennials', '#roster'),
        previous: millennials,
        next: null,
      }),
    ).toBe('/guild');
  });

  it('stays on a guild the visitor looked up', () => {
    expect(
      guildFollowHref({ ...onPage('/guild/eu/normal/the-last-watch'), previous: millennials, next: olympus }),
    ).toBeNull();
    expect(
      guildFollowHref({ ...onPage('/guild/eu/normal/the-last-watch'), previous: null, next: olympus }),
    ).toBeNull();
  });

  it('stays when the new character is in the same guild', () => {
    expect(
      guildFollowHref({
        ...onPage('/guild/us/pvp/employed-millennials'),
        previous: millennials,
        next: millennials,
      }),
    ).toBeNull();
  });

  it('matches the page against the guild however the address spells it', () => {
    const accented = { name: 'Ünd Ër', href: '/guild/us/pvp/ünd-ër' };
    expect(
      guildFollowHref({ ...onPage('/guild/us/pvp/%C3%BCnd-%C3%ABr/'), previous: accented, next: olympus }),
    ).toBe('/guild/us/pvp/olympus-xxvii');
    expect(
      guildFollowHref({
        ...onPage('/guild/us/pvp/Employed-Millennials'),
        previous: millennials,
        next: olympus,
      }),
    ).toBe('/guild/us/pvp/olympus-xxvii');
  });

  it('the landing opens the current character’s guild, and stays for a guildless one', () => {
    expect(guildFollowHref({ ...onPage('/guild'), previous: null, next: olympus })).toBe(olympus.href);
    expect(guildFollowHref({ ...onPage('/guild/', '#claim'), previous: millennials, next: null })).toBeNull();
  });

  it('never moves an invite', () => {
    expect(
      guildFollowHref({ ...onPage('/guild/invite/abc'), previous: millennials, next: olympus }),
    ).toBeNull();
  });
});

describe('characterFollowHref', () => {
  it('moves your character page to the new character, keeping the section', () => {
    expect(
      characterFollowHref({
        pathname: '/character/us/pvp/bow',
        hash: '#gear',
        previous: '/character/us/pvp/bow',
        next: '/character/us/pvp/main',
      }),
    ).toBe('/character/us/pvp/main#gear');
  });

  it('stays on a character the visitor looked up, on the same character and without a page to go to', () => {
    const base = { pathname: '/character/us/pvp/bow', hash: '' };
    expect(
      characterFollowHref({ ...base, previous: '/character/us/pvp/main', next: '/character/us/pvp/x' }),
    ).toBeNull();
    expect(characterFollowHref({ ...base, previous: null, next: '/character/us/pvp/x' })).toBeNull();
    expect(
      characterFollowHref({ ...base, previous: '/character/us/pvp/bow', next: '/character/us/pvp/bow' }),
    ).toBeNull();
    expect(characterFollowHref({ ...base, previous: '/character/us/pvp/bow', next: null })).toBeNull();
  });
});

describe('pointerCharacterHref', () => {
  const me = {
    characters: [
      { key: 'us/pvp/main', region: 'us', ruleset: 'pvp', name: 'Main', class: 'warrior' },
      { key: 'us/pvp/bow', region: 'us', ruleset: 'pvp', name: 'Bow', class: 'hunter' },
    ],
    guilds: [],
    main_character_key: 'us/pvp/main',
  } as unknown as Me;
  const pointer = (source: CurrentCharacter['source'], ref: string, label = ref): CurrentCharacter => ({
    source,
    ref,
    label,
    classSlug: 'hunter',
    savedAt: '2026-11-01T00:00:00Z',
  });

  it('is the armory character’s page, signed in or not', () => {
    expect(pointerCharacterHref(null, pointer('armory', 'eu/normal/far'))).toBe('/character/eu/normal/far');
  });

  it('resolves an addon pointer and the absent pointer through the account', () => {
    expect(pointerCharacterHref(me, pointer('addon', 'FS1:x', 'Bow · Beast Mastery Hunter'))).toBe(
      '/character/us/pvp/bow',
    );
    expect(pointerCharacterHref(me, null)).toBe('/character/us/pvp/main');
  });

  it('is null for a pointer with no character page', () => {
    expect(pointerCharacterHref(null, null)).toBeNull();
    expect(pointerCharacterHref(me, pointer('code', 'FS1:y'))).toBeNull();
  });
});
