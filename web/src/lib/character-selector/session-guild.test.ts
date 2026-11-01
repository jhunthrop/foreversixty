import { describe, expect, it } from 'vitest';
import type { Me, MeGuild } from '../account/api';
import { readFileSync } from 'node:fs';
import * as sessionGuild from './session-guild';
import { currentGuildOf } from './session-guild';
import type { CurrentCharacter } from '../current-character';

const guild = (overrides: Partial<MeGuild>): MeGuild => ({
  id: 1,
  region: 'us',
  ruleset: 'pvp',
  name: 'OLYMPUS XXVII',
  verified: true,
  verified_count: 1,
  ...overrides,
});

describe('currentGuildOf', () => {
  const olympus = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };
  const millennials = { name: 'Employed Millennials', href: '/guild/us/pvp/employed-millennials' };
  const guilds = [guild({}), guild({ id: 2, name: 'Employed Millennials' })];
  const usPvp = { region: 'us', ruleset: 'pvp' };
  const chars = [
    {
      key: 'us/pvp/main',
      ...usPvp,
      name: 'Main',
      class: 'warrior',
      guild: { id: 1, name: 'OLYMPUS XXVII', verified: true },
    },
    {
      key: 'us/pvp/bow',
      ...usPvp,
      name: 'Bow',
      class: 'hunter',
      guild: { id: 2, name: 'Employed Millennials', verified: false },
    },
    { key: 'us/pvp/loner', ...usPvp, name: 'Loner', class: 'mage' },
    {
      key: 'us/pvp/stray',
      ...usPvp,
      name: 'Stray',
      class: 'mage',
      guild: { id: 9, name: 'Elsewhere', verified: true },
    },
    {
      key: 'eu/hardcore/far',
      region: 'eu',
      ruleset: 'hardcore',
      name: 'Far',
      class: 'mage',
      guild: { id: 4, name: 'The Last Watch', verified: true },
    },
  ];
  const me = { guilds, characters: chars, main_character_key: 'us/pvp/main' } as unknown as Me;
  const at = (ref: string, classSlug = 'hunter'): CurrentCharacter => ({
    source: 'armory',
    ref,
    label: ref,
    classSlug,
    savedAt: '2026-10-21T00:00:00Z',
  });

  it('is null signed out', () => {
    expect(currentGuildOf(null, at('us/pvp/bow'))).toBeNull();
    expect(currentGuildOf(undefined, null)).toBeNull();
  });

  it("is the current character's guild, an unverified one included", () => {
    expect(currentGuildOf(me, at('us/pvp/bow'))).toEqual(millennials);
    expect(currentGuildOf(me, at('us/pvp/main'))).toEqual(olympus);
  });

  it('resolves an addon pointer by the name its label carries', () => {
    const addon: CurrentCharacter = { ...at('FS1:x'), source: 'addon', label: 'Bow · Beast Mastery Hunter' };
    expect(currentGuildOf(me, addon)).toEqual(millennials);
  });

  it('is the main character’s guild with no pointer', () => {
    expect(currentGuildOf(me, null)).toEqual(olympus);
  });

  it('is null for a guildless current character, never the main’s guild', () => {
    expect(currentGuildOf(me, at('us/pvp/loner', 'mage'))).toBeNull();
  });

  it('names the character’s own guild even when me.guilds does not list it', () => {
    expect(currentGuildOf(me, at('us/pvp/stray', 'mage'))).toEqual({
      name: 'Elsewhere',
      href: '/guild/us/pvp/elsewhere',
    });
  });

  it('builds the page from the character’s own region and ruleset', () => {
    expect(currentGuildOf(me, at('eu/hardcore/far', 'mage'))).toEqual({
      name: 'The Last Watch',
      href: '/guild/eu/hardcore/the-last-watch',
    });
  });

  it('is null when the pointer names no account character', () => {
    expect(currentGuildOf(me, at('FS1:y'))).toBeNull();
  });
});

describe('the header chip on a guild page (P.1)', () => {
  const source = readFileSync(new URL('./selector-state.svelte.ts', import.meta.url), 'utf8');

  it('never swaps in the guild’s character: it shows the account’s current character on every page', () => {
    expect(Object.keys(sessionGuild)).not.toContain('guildPagePointer');
    expect(source).not.toMatch(/guildPagePointer|window\.location\.pathname/);
    expect(source).toMatch(/buildSelectorModel\(\{[\s\S]*?pointer,/);
  });
});
