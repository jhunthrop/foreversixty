import { describe, expect, it } from 'vitest';
import type { Me, MeGuild } from '../account/api';
import { readFileSync } from 'node:fs';
import * as sessionGuild from './session-guild';
import { memberGuildOf, sessionGuildOf } from './session-guild';

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

describe('the header chip on a guild page (P.1)', () => {
  const source = readFileSync(new URL('./selector-state.svelte.ts', import.meta.url), 'utf8');

  it('never swaps in the guild’s character: it shows the account’s current character on every page', () => {
    expect(Object.keys(sessionGuild)).not.toContain('guildPagePointer');
    expect(source).not.toMatch(/guildPagePointer|window\.location\.pathname/);
    expect(source).toMatch(/buildSelectorModel\(\{[\s\S]*?pointer,/);
  });
});
