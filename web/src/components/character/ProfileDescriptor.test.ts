// web/src/components/character/ProfileDescriptor.test.ts
// L8.4 fix round 2 items 3, 4 and 5: the profile's lines under the character's name.
//   Line 1: "Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP", "Holy Priest" in
//           --text and the rest muted, the guild a link with a 44px tap target on a phone.
//   Line 2: "Level 60 Night Elf".
//   The "HARDCORE US" chip ends the last of those two lines, 8px after its text.
//   Line 3 (the owner's): the Character header's sync line, "Addon · 4 minutes ago", the stale
//           words in ember.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { SyncView } from '../../lib/character-header/sync';
import { profileHeader } from '../../lib/character-profile';
import type { CharacterPage } from '../../lib/rankings/api';
import ProfileDescriptor from './ProfileDescriptor.svelte';

const PATH = { region: 'us', ruleset: 'hardcore' };
const FULL: CharacterPage = {
  character: {
    name: 'Elyra Duskvale',
    region: 'us',
    ruleset: 'hardcore',
    class: 'Priest',
    guild: { name: 'Olympus XXVII', rank: 'member' },
    level: 60,
    race: 'Night Elf',
    realm: 'Classic Beta PvP',
  },
  best: [],
  history: [
    {
      encounter: 'Warden Kelthas',
      encounter_id: 9001,
      difficulty: 8,
      metric: 'hps',
      value: 1840,
      spec: 'Holy',
      fought_at: '2026-12-09T22:10:00Z',
      report_id: 'r1',
      fight_index: 3,
      execution_score: null,
    },
  ],
  builds_seen: [{ talent_split: '31/20/0', spec: 'Holy', first_seen: '2026-12-09T22:10:00Z' }],
};
const BARE: CharacterPage = {
  ...FULL,
  character: { ...FULL.character, guild: null, level: null, race: null, realm: null },
  builds_seen: [],
};

/** The server render without Svelte's hydration comments. */
function view(props: { header: ReturnType<typeof profileHeader>; sync: SyncView | null }): { body: string } {
  return { body: render(ProfileDescriptor, { props }).body.replace(/<!--[^>]*-->/g, '') };
}
const text = (html: string): string => html.replace(/<[^>]+>/g, '').replace(/<!--[^>]*-->/g, '');
const tag = (html: string, testid: string): string =>
  new RegExp(`<[a-z]+[^>]*data-testid="${testid}"[^>]*>`).exec(html)?.[0] ?? '';
const classes = (element: string): string[] => (/class="([^"]*)"/.exec(element)?.[1] ?? '').split(' ');
/** A `<p>` by its testid, open tag to close tag. */
const para = (html: string, testid: string): string =>
  new RegExp(`<p[^>]*data-testid="${testid}"[\\s\\S]*?</p>`).exec(html)?.[0] ?? '';

describe('ProfileDescriptor', () => {
  it('reads "Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP", the spec in --text and the rest muted', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const line = para(body, 'character-descriptor');
    expect(text(line).replace(/&lt;/g, '<').replace(/&gt;/g, '>')).toBe(
      'Holy Priest · 31/20/0 · <Olympus XXVII> · Classic Beta PvP',
    );
    expect(classes(tag(body, 'character-descriptor'))).toContain('text-muted');
    expect(tag(body, 'character-spec')).toMatch(/class="text-text"/);
    expect(classes(tag(body, 'character-split'))).not.toContain('text-text');
    // The brackets are escaped text, still inside the link.
    expect(body).toMatch(/data-testid="character-guild-link"[^>]*>&lt;Olympus XXVII(&gt;|>)<\/a>/);
  });

  it('gives the guild link a 44px tap target on a phone with no layout shift', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const link = tag(body, 'character-guild-link');
    expect(link).toContain('href="/guild/us/hardcore/olympus-xxvii"');
    // 22px of line plus 11px above and below is 44px; the negative margins take the padding
    // back out of the layout.
    expect(classes(link)).toEqual(
      expect.arrayContaining(['max-md:inline-block', 'max-md:py-[11px]', 'max-md:-my-[11px]', 'underline']),
    );
    expect(classes(link)).not.toContain('min-h-11');
  });

  it('ends the level line with the chip, 8px after "Level 60 Night Elf"', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const levelLine = para(body, 'character-level-line');
    expect(levelLine).toMatch(
      /data-testid="character-level-race">Level 60 Night Elf<\/span><span class="[^"]*\bml-2\b[^"]*" data-testid="character-ruleset-chip">Hardcore US<\/span><\/p>/,
    );
    expect(classes(tag(body, 'character-level-line'))).toContain('whitespace-nowrap');
    expect(para(body, 'character-descriptor')).not.toContain('character-ruleset-chip');
  });

  it('ends the first line with the chip when nothing else is known: "Holy Priest" then the chip', () => {
    const { body } = view({ header: profileHeader(BARE, PATH), sync: null });
    expect(body).not.toContain('character-level-line');
    const line = para(body, 'character-descriptor');
    // The last piece and the chip never part: they wrap as one.
    expect(line).toMatch(
      /<span class="whitespace-nowrap"><span class="text-text" data-testid="character-spec">Holy Priest<\/span><span class="[^"]*\bml-2\b[^"]*" data-testid="character-ruleset-chip">Hardcore US<\/span><\/span><\/p>/,
    );
  });

  it('draws the chip with muted text and one border token, whichever line it ends', () => {
    for (const page of [FULL, BARE]) {
      const { body } = view({ header: profileHeader(page, PATH), sync: null });
      const chip = classes(tag(body, 'character-ruleset-chip'));
      expect(chip).toEqual(expect.arrayContaining(['pill', 'text-muted', 'border-line', 'ml-2']));
      expect(chip.filter((token) => token.startsWith('border-'))).toEqual(['border-line']);
      expect(chip.filter((token) => token.startsWith('text-'))).toEqual(['text-muted']);
    }
  });

  it('adds the owner’s sync line as the third line: "Addon · 4 minutes ago", 12px muted', () => {
    const { body } = view({
      header: profileHeader(FULL, PATH),
      sync: { source: 'Addon', age: '4 minutes ago', staleTail: null },
    });
    const sync = para(body, 'character-sync');
    expect(text(sync)).toBe('Addon · 4 minutes ago');
    expect(classes(tag(body, 'character-sync'))).toEqual(
      expect.arrayContaining(['text-muted', 'text-[12px]']),
    );
    expect(body.indexOf('character-level-line')).toBeLessThan(body.indexOf('character-sync'));
    expect(body).not.toContain('Last seen in game');
    expect(body).not.toContain('character-sync-stale');
  });

  it('puts only the stale words in ember, after the age', () => {
    const { body } = view({
      header: profileHeader(FULL, PATH),
      sync: { source: 'Addon', age: '19 days ago', staleTail: 'Log out in the game to sync' },
    });
    const sync = para(body, 'character-sync');
    expect(sync).toMatch(
      /Addon · 19 days ago · <span class="text-ember" data-testid="character-sync-stale">Log out in the game to sync<\/span>/,
    );
  });

  it('has no sync line without one (a visitor, or no build on file)', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    expect(body).not.toContain('character-sync');
  });
});
