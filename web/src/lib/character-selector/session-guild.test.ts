import { describe, expect, it } from 'vitest';
import type { Me, MeGuild } from '../account/api';
import type { CurrentCharacter } from '../current-character';
import { guildPagePointer, memberGuildOf, sessionGuildOf } from './session-guild';

const guild = (overrides: Partial<MeGuild>): MeGuild => ({
  id: 1,
  region: 'us',
  ruleset: 'pvp',
  name: 'OLYMPUS XXVII',
  verified: true,
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

describe('guildPagePointer', () => {
  const LOADED = Date.parse('2026-10-20T20:00:00Z');
  const GUILD_PAGE = '/guild/us/pvp/lifecycle-test';
  const lifecycle = guild({ id: 3, name: 'Lifecycle Test' });
  const pipsqueak = {
    key: 'us/pvp/pipsqueak',
    name: 'Pipsqueak',
    class: 'Priest',
    spec: 'Holy',
    region: 'us',
    ruleset: 'pvp',
  };
  const bow = {
    key: 'us/pvp/bow-jackzon',
    name: 'Bow Jackzon',
    class: 'Hunter',
    spec: 'BeastMastery',
    region: 'us',
    ruleset: 'pvp',
    guild: { id: 3, name: 'Lifecycle Test', verified: true },
  };
  const account = (main?: string): Me =>
    ({ guilds: [lifecycle], characters: [pipsqueak, bow], main_character_key: main }) as unknown as Me;
  const pointerOf = (character: { key: string; name: string }, savedAt: string): CurrentCharacter => ({
    source: 'armory',
    ref: character.key,
    label: `${character.name} · Holy Priest`,
    classSlug: 'priest',
    savedAt,
  });
  const before = '2026-10-19T09:00:00Z';

  it('shows the account character that is in the guild, not the current non-guild one', () => {
    const shown = guildPagePointer(account(pipsqueak.key), GUILD_PAGE, pointerOf(pipsqueak, before), LOADED);
    expect(shown?.ref).toBe('us/pvp/bow-jackzon');
    expect(shown?.label).toContain('Bow Jackzon');
  });

  it('leaves the pointer alone on any other page', () => {
    const pointer = pointerOf(pipsqueak, before);
    expect(guildPagePointer(account(), '/planner', pointer, LOADED)).toBe(pointer);
    expect(guildPagePointer(account(), '/guild/us/pvp/another-guild', pointer, LOADED)).toBe(pointer);
  });

  it('leaves the pointer alone signed out and for an account with no character in the guild', () => {
    const pointer = pointerOf(pipsqueak, before);
    expect(guildPagePointer(null, GUILD_PAGE, pointer, LOADED)).toBe(pointer);
    const stranger = { guilds: [lifecycle], characters: [pipsqueak] } as unknown as Me;
    expect(guildPagePointer(stranger, GUILD_PAGE, pointer, LOADED)).toBe(pointer);
  });

  it('keeps a pointer that already is a character of the guild', () => {
    const pointer = pointerOf(bow, before);
    expect(guildPagePointer(account(), GUILD_PAGE, pointer, LOADED)).toBe(pointer);
  });

  it('keeps a character the player picked on this page', () => {
    const picked = pointerOf(pipsqueak, '2026-10-20T20:05:00Z');
    expect(guildPagePointer(account(), GUILD_PAGE, picked, LOADED)).toBe(picked);
  });

  it('falls back to the guild character when there is no pointer at all', () => {
    expect(guildPagePointer(account(), GUILD_PAGE, null, LOADED)?.ref).toBe('us/pvp/bow-jackzon');
  });
});
