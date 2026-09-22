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
  realm: 'Whitemane',
  level: 60,
  faction: 'alliance',
  source: 'bnet',
  guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', rank_index: 1, verified: true },
};

const WITH_AVATAR: MeCharacter = {
  ...GUILDED,
  key: 'us/pvp/thoradin',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
  guild: undefined,
  avatar_url: '/fixtures/avatar-placeholder.jpg',
};

const UNGUILDED: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
};

describe('CharacterList', () => {
  it('shows a class-coloured letter square and the race/class/level/realm descriptor', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain('data-testid="character-avatar-fallback"');
    expect(body).toContain('>P<'); // Priest's initial
    expect(body).toContain('Night Elf Priest · Level 60 · Whitemane (Hardcore US)');
    expect(body).not.toContain('data-testid="character-avatar"');
  });

  it('shows the avatar image, not the letter square, when avatar_url is set', () => {
    const { body } = render(CharacterList, { props: { characters: [WITH_AVATAR] } });
    expect(body).toContain('data-testid="character-avatar"');
    expect(body).toContain('src="/fixtures/avatar-placeholder.jpg"');
    expect(body).toContain('alt=""');
    expect(body).toContain('loading="lazy"');
    expect(body).not.toContain('data-testid="character-avatar-fallback"');
  });

  it('shows a guilded, verified character with the guild line and a verified pill', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain('Iron Vanguard');
    expect(body).toContain('Officer');
    expect(body).toContain('data-testid="character-guild-verified"');
  });

  it('renders an unguilded character with no guild line, and omits absent fields from the descriptor', () => {
    const { body } = render(CharacterList, { props: { characters: [UNGUILDED] } });
    expect(body).toContain('Thoradin');
    expect(body).not.toContain('data-testid="character-guild-line"');
    expect(body).toContain('Warrior · PvP US');
  });

  it('shows the export explanation once, as the panel footer line', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain(characterListCopy.introLead);
    expect(body).toContain('href="/addon#paste"');
  });

  it('shows EmptyState with one action (Refresh from Battle.net) when there are no characters', () => {
    const { body } = render(CharacterList, { props: { characters: [] } });
    expect(body).toContain(characterListCopy.empty);
    expect(body).toContain('data-testid="account-characters-empty"');
    expect(body).toContain(characterListCopy.refreshFromBattlenet);
  });

  it('shows the imported-from line and the refresh link only when bnetImportedAt is set', () => {
    const withImport = render(CharacterList, {
      props: { characters: [GUILDED], bnetImportedAt: '2026-09-20T00:00:00Z' },
    });
    expect(withImport.body).toContain('data-testid="bnet-imported"');
    expect(withImport.body).toContain('Imported from Battle.net');
    expect(withImport.body).toContain(characterListCopy.refreshFromBattlenet);

    const without = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(without.body).not.toContain('data-testid="bnet-imported"');
  });
});
