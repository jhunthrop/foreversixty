// web/src/components/account/CharacterList.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterList from './CharacterList.svelte';
import { characterListCopy } from '../../lib/account/character-list-copy';
import type { MeCharacter } from '../../lib/account/api';

const GUILDED: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'priest',
  race: 'Night Elf',
  item_level: 63,
  realm: 'Whitemane',
  level: 60,
  faction: 'alliance',
  source: 'bnet',
  guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', rank_index: 1, verified: true },
};

const UNGUILDED: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
};

describe('CharacterList', () => {
  it('prints the race and the class display name from the class slug, and the item level', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain('Night Elf');
    expect(body).toContain('>Priest<');
    expect(body).not.toContain('>priest<');
    expect(body).toContain(characterListCopy.itemLevelPrefix(63));
  });

  it('shows a guilded, verified character with the guild line and a verified pill', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain('Iron Vanguard');
    expect(body).toContain('Officer');
    expect(body).toContain('data-testid="character-guild-verified"');
    expect(body).toContain('Whitemane');
    expect(body).toContain('Level 60');
  });

  it('renders an unguilded character with no guild line at all', () => {
    const { body } = render(CharacterList, { props: { characters: [UNGUILDED] } });
    expect(body).toContain('Thoradin');
    expect(body).not.toContain('data-testid="character-guild-line"');
  });

  it('shows the intro line once, with the paste link, above the list', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain(characterListCopy.introLead);
    expect(body).toContain('href="/addon#paste"');
  });

  it('shows EmptyState with one action (Refresh from Battle.net) when there are no characters', () => {
    const { body } = render(CharacterList, { props: { characters: [] } });
    expect(body).toContain(characterListCopy.empty);
    expect(body).toContain('data-testid="account-characters-empty"');
    expect(body).toContain(characterListCopy.refreshFromBattlenet);
    expect(body).not.toContain('data-testid="characters-paste"');
  });

  it('shows the imported-from line only when bnetImportedAt is set', () => {
    const withImport = render(CharacterList, {
      props: { characters: [GUILDED], bnetImportedAt: '2026-09-20T00:00:00Z' },
    });
    expect(withImport.body).toContain('data-testid="bnet-imported"');
    expect(withImport.body).toContain('Imported from Battle.net');

    const without = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(without.body).not.toContain('data-testid="bnet-imported"');
  });
});
