// web/src/components/account/CharacterList.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterList from './CharacterList.svelte';
import { characterListCopy } from '../../lib/account/character-list-copy';
import type { MeCharacter } from '../../lib/account/api';
import { ctaCopy } from '../../lib/addon/cta-copy';

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

const WITH_BLIZZARD_BUILD: MeCharacter = {
  ...UNGUILDED,
  key: 'us/pvp/blizzardbuilt',
  name: 'Blizzardbuilt',
  build: { source: 'blizzard', captured_at: '2026-09-20T00:00:00Z' },
};

const WITH_ADDON_BUILD: MeCharacter = {
  ...UNGUILDED,
  key: 'us/pvp/addonbuilt',
  name: 'Addonbuilt',
  build: { source: 'addon', captured_at: '2026-09-20T00:00:00Z' },
};

/** The rendered text, tags stripped and spaces collapsed: a descriptor's phrases are spans. */
function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\s+/g, ' ');
}

describe('CharacterList', () => {
  it('marks the chosen main with a pill and offers Set as main on every alt', () => {
    const { body } = render(CharacterList, {
      props: { characters: [GUILDED, UNGUILDED], mainKey: GUILDED.key, onSetMain: async () => {} },
    });
    expect((body.match(/character-main-pill/g) ?? []).length).toBe(1);
    expect((body.match(/character-set-main"/g) ?? []).length).toBe(1);
    expect(body).toContain(characterListCopy.main);
    expect(body).toContain(characterListCopy.setAsMain);
  });

  it('draws the circular class crest for every row, never an avatar or a letter square', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED, WITH_AVATAR] } });
    expect(body).toContain('/icons/hd/crests/priest.webp');
    expect(body).toContain('/icons/hd/crests/warrior.webp');
    expect(body).not.toContain('/fixtures/avatar-placeholder.jpg');
    expect(body).not.toContain('data-testid="character-avatar-fallback"');
    expect(text(body)).toContain('Night Elf Priest · Level 60 · Whitemane (Hardcore US)');
  });

  it("puts a guilded character's guild on the second sub line", () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain('data-testid="character-guild">Iron Vanguard');
  });

  it('renders an unguilded character with no guild line, and omits absent fields from the descriptor', () => {
    const { body } = render(CharacterList, { props: { characters: [UNGUILDED] } });
    expect(body).toContain('Thoradin');
    expect(body).not.toContain('data-testid="character-guild"');
    expect(text(body)).toContain('Warrior · PvP US');
  });

  it('shows the export explanation once, as the panel footer line', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).toContain(characterListCopy.introBattlenetLine);
    expect(body).toContain(characterListCopy.introInstallAddonLink);
    expect(body).toContain('href="/setup"');
    expect(body).toContain(characterListCopy.introPasteLink);
    expect(body).toContain('href="/setup#paste"');
  });

  /** The build-source Tag's word on the only row. */
  const source = (body: string): string =>
    /data-testid="character-source"[^>]*>([^<]*)/.exec(body.replace(/<!--.*?-->/g, ''))?.[1]?.trim() ?? '';

  it('tags the build source: Battle.net, Addon, or No gear yet', () => {
    expect(source(render(CharacterList, { props: { characters: [WITH_BLIZZARD_BUILD] } }).body)).toBe(
      characterListCopy.battlenetSource,
    );
    expect(source(render(CharacterList, { props: { characters: [WITH_ADDON_BUILD] } }).body)).toBe(
      characterListCopy.addonSource,
    );
    expect(source(render(CharacterList, { props: { characters: [UNGUILDED] } }).body)).toBe(
      characterListCopy.noBuildYet,
    );
  });

  it("links each row's Open to the character page; the row itself is not a link", () => {
    const { body } = render(CharacterList, { props: { characters: [WITH_BLIZZARD_BUILD] } });
    expect(body).toContain('href="/character/us/pvp/blizzardbuilt" data-testid="character-open"');
    expect(text(body)).toContain(`${characterListCopy.open} ›`);
    expect(body).not.toContain('<a href="/character/us/pvp/blizzardbuilt" class="ui-row');
  });

  it('is one framed panel, Characters, with the reports link inside it', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body.match(/data-region="plate"/g)?.length).toBe(1);
    expect(body).toContain('data-region="section-characters"');
    expect(body).toMatch(/class="[^"]*ui-framed[^"]*" data-region="plate"/);
    expect(body.indexOf('account-reports-link')).toBeLessThan(body.lastIndexOf('</section>'));
    expect(body).toContain('href="/reports" data-testid="account-reports-link"');
  });

  it('draws no last-raid line before a raid is read (spec 2026-11-08 G5)', () => {
    const { body } = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(body).not.toContain('coach-last-raid');
  });

  it('says what a Battle.net refresh did, and offers Try again when it did not answer', () => {
    const changed = render(CharacterList, {
      props: { characters: [GUILDED], refreshResult: { kind: 'changed', count: 5, added: 0 } },
    });
    expect(changed.body).toContain('Updated just now · 5 characters');
    const unchanged = render(CharacterList, {
      props: { characters: [GUILDED], refreshResult: { kind: 'unchanged', count: 1 } },
    });
    expect(unchanged.body).toContain('No changes · 1 character');
    const failed = render(CharacterList, {
      props: { characters: [GUILDED], refreshResult: { kind: 'error' } },
    });
    expect(failed.body).toContain(characterListCopy.refreshError);
    expect(failed.body).toContain(characterListCopy.tryAgain);
  });

  it('fills the Plate with the empty sentence, Get set up first and Paste an export, when there are no characters', () => {
    const { body } = render(CharacterList, { props: { characters: [] } });
    expect(body).toContain(characterListCopy.emptyLead);
    expect(body).toContain(characterListCopy.emptyKorea);
    expect(body).toContain('data-testid="account-characters-empty"');
    expect(body).toContain(characterListCopy.getSetUp);
    expect(body).toContain(characterListCopy.getSetUpHref);
    expect(body.indexOf(characterListCopy.getSetUp)).toBeLessThan(
      body.indexOf(characterListCopy.pasteAnExport),
    );
    expect(body).toContain(characterListCopy.pasteHref);
    expect(characterListCopy.empty).toBe(
      `No Forever characters on this Battle.net account yet. Make one in game, then refresh from Battle.net. Playing in Korea? ${ctaCopy.installAddonAndCompanion}, or paste an export.`,
    );
  });

  it('says the setup story once on an empty account: the intro line is left out under the empty state', () => {
    const { body } = render(CharacterList, { props: { characters: [] } });
    expect(body).not.toContain(characterListCopy.introBattlenetLine);
    expect(body.split('href="/setup"').length - 1).toBe(1);
  });

  it('shows the updated-from-Battle.net line and the refresh link only when bnetImportedAt is set', () => {
    const withImport = render(CharacterList, {
      props: { characters: [GUILDED], bnetImportedAt: '2026-09-20T00:00:00Z' },
    });
    expect(withImport.body).toMatch(/Updated [^<]* from Battle\.net/);
    expect(withImport.body).toContain(characterListCopy.refreshFromBattlenet);

    const without = render(CharacterList, { props: { characters: [GUILDED] } });
    expect(without.body).not.toMatch(/Updated [^<]* from Battle\.net/);
    expect(without.body).not.toContain(characterListCopy.refreshFromBattlenet);
  });

  it('never draws the Supporter chip on a row: the mark belongs to the person, not the character', () => {
    const body = render(CharacterList, { props: { characters: [GUILDED, UNGUILDED] } }).body;
    expect(body).not.toContain('supporter-chip');
  });

  it('folds characters below level 10 behind the toggle but keeps the current one', () => {
    const at = (name: string, level: number): MeCharacter => ({
      ...UNGUILDED,
      key: `us/normal/${name.toLowerCase()}`,
      name,
      level,
    });
    const body = render(CharacterList, {
      props: {
        characters: [at('Raider', 60), at('Bankalt', 1), at('Fresh', 5)],
        currentKey: 'us/normal/fresh',
      },
    }).body;
    expect(body).toContain('Raider');
    expect(body).toContain('Fresh');
    expect(body).not.toContain('Bankalt');
    expect(body).toContain('Show low-level characters (1)');
  });
});
