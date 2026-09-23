// web/src/components/character/CharacterRow.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import CharacterRow from './CharacterRow.svelte';

const GUILDED: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'priest',
  guild: { id: 12, name: 'Iron Vanguard', rank: 'officer', verified: true },
};

const PLAIN: MeCharacter = {
  key: 'us/pvp/thoradin',
  region: 'us',
  ruleset: 'pvp',
  name: 'Thoradin',
  class: 'warrior',
  build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
};

describe('CharacterRow', () => {
  it('renders one <li> with the realm descriptor by default', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN } });
    expect(body).toContain('<li');
    expect(body).toContain('PvP · US');
  });

  it('renders the full descriptor when asked', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN, descriptor: 'full' } });
    expect(body).toContain('Warrior');
    expect(body).not.toContain('PvP · US');
  });

  it('renders the build pill with the given test id, addon vs no build', () => {
    const withBuild = render(CharacterRow, { props: { character: PLAIN, pillTestid: 'p' } });
    expect(withBuild.body).toMatch(/data-testid="p">\s*Addon/);

    const withoutBuild = render(CharacterRow, { props: { character: GUILDED, pillTestid: 'p' } });
    expect(withoutBuild.body).toMatch(/data-testid="p">\s*No build yet/);
  });

  it('renders the guild line and verified pill only when guildLine is true', () => {
    const on = render(CharacterRow, { props: { character: GUILDED, guildLine: true } });
    expect(on.body).toContain('data-testid="character-guild-line"');
    expect(on.body).toContain('data-testid="character-guild-verified"');

    const off = render(CharacterRow, { props: { character: GUILDED } });
    expect(off.body).not.toContain('data-testid="character-guild-line"');
  });

  it('omits the guild line for a character with no guild even when guildLine is true', () => {
    const { body } = render(CharacterRow, { props: { character: PLAIN, guildLine: true } });
    expect(body).not.toContain('data-testid="character-guild-line"');
  });

  it('renders the name with the given href and nameTestid', () => {
    const { body } = render(CharacterRow, {
      props: { character: PLAIN, href: '/character/x', nameTestid: 'n' },
    });
    expect(body).toContain('href="/character/x"');
    expect(body).toContain('data-testid="n"');
  });
});
