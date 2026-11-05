import { describe, expect, it } from 'vitest';
import { guildSuggestions, searchTerm, type GuildCandidate } from './guild-suggest';

const LIFECYCLE: GuildCandidate = { name: 'Lifecycle Test', region: 'us', ruleset: 'pvp' };
const at = (slug: string, ruleset: 'pvp' | 'normal' = 'pvp') => ({ region: 'us' as const, ruleset, slug });
const first = (...args: Parameters<typeof guildSuggestions>) => guildSuggestions(...args)[0] ?? null;

describe('guildSuggestions from the guilds the page reads', () => {
  it('offers the guild whose name starts with what was typed, with where it lives', () => {
    expect(guildSuggestions(at('lifecycle'), [LIFECYCLE])).toEqual([
      { name: 'Lifecycle Test', href: '/guild/us/pvp/lifecycle-test', where: 'US PvP' },
    ]);
  });

  it('offers a name that only contains the typed words, after any that start with them', () => {
    const lazy = { name: 'The Lifecycle', region: 'us', ruleset: 'pvp' };
    expect(first(at('cycle'), [lazy])?.name).toBe('The Lifecycle');
    expect(guildSuggestions(at('lifecycle'), [lazy, LIFECYCLE]).map((s) => s.name)).toEqual([
      'Lifecycle Test',
      'The Lifecycle',
    ]);
  });

  it('prefers the region and ruleset typed, else the order given', () => {
    const normal = { ...LIFECYCLE, ruleset: 'normal' };
    expect(first(at('lifecycle', 'normal'), [LIFECYCLE, normal])?.where).toBe('US Normal');
    expect(first(at('lifecycle'), [normal, LIFECYCLE])?.where).toBe('US PvP');
  });

  it('offers the full name typed on the wrong ruleset, never the address that just failed', () => {
    expect(first(at('lifecycle-test', 'normal'), [LIFECYCLE])?.where).toBe('US PvP');
    expect(guildSuggestions(at('lifecycle-test'), [LIFECYCLE])).toEqual([]);
  });

  it('reads the slug as words, any case, and guesses nothing from under three letters', () => {
    expect(first(at('LIFECYCLE_TE'), [LIFECYCLE])?.name).toBe('Lifecycle Test');
    expect(guildSuggestions(at('li'), [LIFECYCLE])).toEqual([]);
    expect(guildSuggestions(at('lifecyle-tset'), [LIFECYCLE])).toEqual([]);
    expect(guildSuggestions(at('lifecycle'), [])).toEqual([]);
  });
});

describe('guildSuggestions with the guild search', () => {
  const search = (name: string, ruleset = 'pvp'): GuildCandidate => ({ name, region: 'us', ruleset });

  it('adds the search rows after the known guilds, as the search ordered them, a typo included', () => {
    const rows = [search('Lifecycle Testers'), search('Lifecycle Test', 'normal')];
    expect(guildSuggestions(at('lifecyle-tset'), [], rows).map((s) => `${s.name} (${s.where})`)).toEqual([
      'Lifecycle Testers (US PvP)',
      'Lifecycle Test (US Normal)',
    ]);
  });

  it('names a guild once, never the failed address, and at most three', () => {
    const rows = [LIFECYCLE, search('Lifecycle'), search('Lifecycle A'), search('Lifecycle B'), search('Lifecycle C')];
    const names = guildSuggestions(at('lifecycle'), [LIFECYCLE], rows).map((s) => s.name);
    expect(names).toEqual(['Lifecycle Test', 'Lifecycle A', 'Lifecycle B']);
  });

  it('asks nothing under three letters', () => {
    expect(searchTerm(at('li'))).toBeNull();
    expect(searchTerm(at('Lifecyle_Tset'))).toBe('lifecyle tset');
    expect(guildSuggestions(at('li'), [], [LIFECYCLE])).toEqual([]);
  });
});
