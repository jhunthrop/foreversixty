import { describe, expect, it } from 'vitest';
import type { Me, MeGuild } from '../account/api';
import { readFileSync } from 'node:fs';
import * as sessionGuild from './session-guild';
import { currentGuildOf, memberGuildOf, sessionGuildOf } from './session-guild';
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

const meWith = (guilds: MeGuild[]): Me => ({ guilds }) as unknown as Me;

describe('sessionGuildOf', () => {
  it('is null signed out, loading, or with no guild', () => {
    expect(sessionGuildOf(null)).toBeNull();
    expect(sessionGuildOf(undefined)).toBeNull();
    expect(sessionGuildOf(meWith([]))).toBeNull();
  });

  it('names the first verified guild and its page', () => {
    const me = meWith([guild({ id: 2, name: 'Unverified', verified: false }), guild({})]);
    expect(sessionGuildOf(me)).toEqual({ name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' });
  });

  it('is null when no guild is verified', () => {
    expect(sessionGuildOf(meWith([guild({ verified: false })]))).toBeNull();
  });
});

describe('memberGuildOf', () => {
  const olympus = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };
  const other = { name: 'Second Guild', href: '/guild/us/pvp/second-guild' };
  const withChars = (guilds: MeGuild[], main?: string, chars: unknown[] = []): Me =>
    ({ guilds, characters: chars, main_character_key: main }) as unknown as Me;

  it('is null signed out, loading, or with no guild', () => {
    expect(memberGuildOf(null)).toBeNull();
    expect(memberGuildOf(undefined)).toBeNull();
    expect(memberGuildOf(withChars([]))).toBeNull();
  });

  it('names a guild the account belongs to even before an officer verifies it', () => {
    expect(memberGuildOf(withChars([guild({ verified: false })]))).toEqual(olympus);
  });

  it('is the first guild when there is no main character', () => {
    const me = withChars([guild({}), guild({ id: 2, name: 'Second Guild' })]);
    expect(memberGuildOf(me)).toEqual(olympus);
  });

  it("is the main character's guild when the account has several", () => {
    const me = withChars([guild({}), guild({ id: 2, name: 'Second Guild' })], 'us/pvp/alt-main', [
      { key: 'us/pvp/alt-main', guild: { id: 2, name: 'Second Guild', verified: true } },
    ]);
    expect(memberGuildOf(me)).toEqual(other);
  });

  it('falls back to the first guild when the main character has none', () => {
    const me = withChars([guild({}), guild({ id: 2, name: 'Second Guild' })], 'us/pvp/loner', [
      { key: 'us/pvp/loner' },
    ]);
    expect(memberGuildOf(me)).toEqual(olympus);
  });
});

describe('currentGuildOf', () => {
  const olympus = { name: 'OLYMPUS XXVII', href: '/guild/us/pvp/olympus-xxvii' };
  const millennials = { name: 'Employed Millennials', href: '/guild/us/pvp/employed-millennials' };
  const guilds = [guild({}), guild({ id: 2, name: 'Employed Millennials' })];
  const chars = [
    {
      key: 'us/pvp/main',
      name: 'Main',
      class: 'warrior',
      guild: { id: 1, name: 'OLYMPUS XXVII', verified: true },
    },
    {
      key: 'us/pvp/bow',
      name: 'Bow',
      class: 'hunter',
      guild: { id: 2, name: 'Employed Millennials', verified: false },
    },
    { key: 'us/pvp/loner', name: 'Loner', class: 'mage' },
    {
      key: 'us/pvp/stray',
      name: 'Stray',
      class: 'mage',
      guild: { id: 9, name: 'Elsewhere', verified: true },
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

  it('falls back to the main’s guild when the current character has none or a guild me.guilds lacks', () => {
    expect(currentGuildOf(me, at('us/pvp/loner', 'mage'))).toEqual(olympus);
    expect(currentGuildOf(me, at('us/pvp/stray', 'mage'))).toEqual(olympus);
  });

  it('falls back when the pointer names no account character', () => {
    expect(currentGuildOf(me, at('FS1:y'))).toEqual(olympus);
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
