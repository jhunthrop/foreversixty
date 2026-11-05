import { describe, expect, it } from 'vitest';
import { guildSuggestion, type GuildCandidate } from './guild-suggest';

const LIFECYCLE: GuildCandidate = { name: 'Lifecycle Test', region: 'us', ruleset: 'pvp' };
const at = (slug: string, ruleset: 'pvp' | 'normal' = 'pvp') => ({ region: 'us' as const, ruleset, slug });

describe('guildSuggestion', () => {
  it('offers the guild whose name starts with what was typed, with where it lives', () => {
    expect(guildSuggestion(at('lifecycle'), [LIFECYCLE])).toEqual({
      name: 'Lifecycle Test',
      href: '/guild/us/pvp/lifecycle-test',
      where: 'US PvP',
    });
  });

  it('offers a name that only contains the typed words, after any that start with them', () => {
    const lazy = { name: 'The Lifecycle', region: 'us', ruleset: 'pvp' };
    expect(guildSuggestion(at('cycle'), [lazy])?.name).toBe('The Lifecycle');
    expect(guildSuggestion(at('lifecycle'), [lazy, LIFECYCLE])?.name).toBe('Lifecycle Test');
  });

  it('prefers the region and ruleset typed, else the order given', () => {
    const normal = { ...LIFECYCLE, ruleset: 'normal' };
    expect(guildSuggestion(at('lifecycle', 'normal'), [LIFECYCLE, normal])?.where).toBe('US Normal');
    expect(guildSuggestion(at('lifecycle'), [normal, LIFECYCLE])?.where).toBe('US PvP');
  });

  it('offers the full name typed on the wrong ruleset, never the address that just failed', () => {
    expect(guildSuggestion(at('lifecycle-test', 'normal'), [LIFECYCLE])?.where).toBe('US PvP');
    expect(guildSuggestion(at('lifecycle-test'), [LIFECYCLE])).toBeNull();
  });

  it('reads the slug as words, any case, and guesses nothing from under three letters', () => {
    expect(guildSuggestion(at('LIFECYCLE_TE'), [LIFECYCLE])?.name).toBe('Lifecycle Test');
    expect(guildSuggestion(at('li'), [LIFECYCLE])).toBeNull();
    expect(guildSuggestion(at('lifecyle-tset'), [LIFECYCLE])).toBeNull();
    expect(guildSuggestion(at('lifecycle'), [])).toBeNull();
  });
});
