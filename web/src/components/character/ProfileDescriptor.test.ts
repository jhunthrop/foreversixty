// web/src/components/character/ProfileDescriptor.test.ts
// L8.4 fix round 2 items 3, 4 and 5: the profile's lines under the character's name.
//   Line 1: "Holy Priest · 31/20/0 · <Olympus XXVII> · Doomhowl", "Holy Priest" in
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
    realm: 'Doomhowl',
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
/** The words a reader sees: no tags, the no-break space read as a space and the chip's
 *  zero-width word joiner (fix round 3 item 2) dropped. */
const text = (html: string): string =>
  html
    .replace(/<[^>]+>/g, '')
    .replace(/<!--[^>]*-->/g, '')
    .replace(/\u00a0|&nbsp;/g, ' ')
    .replace(/\u2060/g, '');
const tag = (html: string, testid: string): string =>
  new RegExp(`<[a-z]+[^>]*data-testid="${testid}"[^>]*>`).exec(html)?.[0] ?? '';
const classes = (element: string): string[] => (/class="([^"]*)"/.exec(element)?.[1] ?? '').split(' ');
/** A `<p>` by its testid, open tag to close tag. */
const para = (html: string, testid: string): string =>
  new RegExp(`<p[^>]*data-testid="${testid}"[\\s\\S]*?</p>`).exec(html)?.[0] ?? '';

describe('ProfileDescriptor', () => {
  it('reads "Holy Priest · 31/20/0 · <Olympus XXVII> · Doomhowl", the spec in --text and the rest muted', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const line = para(body, 'character-descriptor');
    expect(text(line).replace(/&lt;/g, '<').replace(/&gt;/g, '>')).toBe(
      'Holy Priest · 31/20/0 · <Olympus XXVII> · Doomhowl',
    );
    expect(classes(tag(body, 'character-descriptor'))).toContain('text-muted');
    // The class is in its colour: Priest has a token, so the spec piece takes it.
    expect(tag(body, 'character-spec')).toMatch(/style="color: var\(--color-class-priest\)"/);
    expect(classes(tag(body, 'character-split'))).not.toContain('text-text');
    // The brackets are escaped text, still inside the link.
    expect(body).toMatch(/data-testid="character-guild-link"[^>]*>&lt;Olympus XXVII(&gt;|>)<\/a>/);
  });

  // Fix round 3 item 2 (D2): the link stays inline, so a long guild name can wrap; its phone
  // tap target is vertical padding on the inline box, which an inline box never adds to the
  // line's height. Fix round 4 items 6 and 8: 14px above and below the 16px content box make
  // 44px, not the 46 that 15px made (Chrome measures 11px at 38).
  // Fix round 5 item 6: inside its nowrap segment the link turns wrapping back on, so a long
  // guild still wraps (D2) while its trailing dot stays glued to its last word.
  it('gives the guild link a 44px tap target on a phone as an inline box that can wrap', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const link = tag(body, 'character-guild-link');
    expect(link).toContain('href="/guild/us/hardcore/olympus-xxvii"');
    expect(classes(link)).toEqual(
      expect.arrayContaining(['max-md:py-[14px]', 'underline', 'whitespace-normal']),
    );
    expect(classes(link).filter((token) => /inline-block|nowrap|-my-/.test(token))).toEqual([]);
    expect(classes(link)).not.toContain('min-h-11');
  });

  // Fix round 4 item 8: a pressed link takes --gold-hover, the design system's link :active.
  it('gives the guild link a pressed state in --gold-hover', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    expect(classes(tag(body, 'character-guild-link'))).toContain('active:text-gold-hover');
  });

  // L8.5 ruling 3 (replacing fix round 3 item 11's offset 0): the focus ring stands 2px off the
  // guild's text like every other ring, and the guild takes 4px on each side a dot sits -- the
  // ring's 2px offset plus its 2px width -- so the ring ends where the text did and the dots
  // keep their space and stay visible. A guild that ends the line has the chip after it, not a
  // dot, and keeps the chip's 8px.
  it('draws the guild link’s focus ring at a 2px offset, with room beside each dot for it (L8.5 ruling 3)', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const link = classes(tag(body, 'character-guild-link'));
    expect(link).toEqual(expect.arrayContaining(['focus-visible:outline-offset-2', 'ml-1', 'mr-1']));
    expect(link).not.toContain('focus-visible:outline-offset-0');
    expect(link.filter((token) => /outline-\d|outline-solid|outline-none/.test(token))).toEqual([]);
    const last = view({
      header: profileHeader({ ...FULL, character: { ...FULL.character, realm: null } }, PATH),
      sync: null,
    }).body;
    const lastLink = classes(tag(last, 'character-guild-link'));
    expect(lastLink).toContain('ml-1');
    expect(lastLink).not.toContain('mr-1');
  });

  // Fix round 5 item 6 (replacing fix round 3 item 10's leading dot): every segment is one
  // nowrap unit that ends in its "·" (a no-break space, then the dot), and plain spaces outside
  // the units part them, so a wrap falls only after a dot and no line starts with one. The
  // realm, last, is one unit with no dot.
  it('ends every segment but the last with its dot, inside one nowrap unit, and breaks only between them', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const line = para(body, 'character-descriptor').replace(/&nbsp;/g, '\u00a0');
    const unit =
      /<span class="(?:inline-block max-w-full )?whitespace-nowrap" data-testid="character-segment">([\s\S]*?)<\/span>(?= |<\/p>)/g;
    const segments = [...line.matchAll(unit)].map((match) =>
      match[1]
        .replace(/<[^>]+>/g, '')
        .replace(/&lt;/g, '<')
        .replace(/&gt;/g, '>'),
    );
    expect(segments).toEqual(['Holy Priest\u00a0·', '31/20/0\u00a0·', '<Olympus XXVII>\u00a0·', 'Doomhowl']);
    // Between the units only a plain, breakable space, and no dot outside them.
    expect(line.replace(unit, '|').replace(/<[^>]+>/g, '')).toBe('| | | |');
    // Only the guild's unit is an inline-block, so a guild longer than the line can still wrap.
    expect(
      line.match(/class="inline-block max-w-full whitespace-nowrap" data-testid="character-segment"><a /g),
    ).toHaveLength(1);
    // The realm is one unit, holding no dot.
    expect(line).toMatch(
      /<span class="whitespace-nowrap" data-testid="character-segment"><span data-testid="character-realm">Doomhowl<\/span><\/span><\/p>/,
    );
  });

  // Fix round 2 item 3, fix round 3 item 14: the chip sits on the line box, centred with its
  // text (vertical-align: top on a 22px pill in a 22px line), not 1.7px low.
  it('aligns the chip to the line box, not the baseline', () => {
    for (const page of [FULL, BARE]) {
      const { body } = view({ header: profileHeader(page, PATH), sync: null });
      const chip = classes(tag(body, 'character-ruleset-chip'));
      expect(chip).toContain('align-top');
      expect(chip).not.toContain('align-middle');
    }
  });

  // Fix round 3 item 2 (D2): with no realm the guild ends the line; only the chip is glued to
  // its last word (a word joiner in a nowrap span), so the guild text itself can still wrap
  // (fix round 5 item 6: its segment is nowrap, the link inside it whitespace-normal).
  it('glues only the chip to a guild that ends the first line, never the whole guild', () => {
    const page: CharacterPage = {
      ...FULL,
      character: {
        ...FULL.character,
        guild: { name: 'Olympus XXVII Veterans of the Long March', rank: 'member' },
        realm: null,
        level: null,
        race: null,
      },
    };
    const line = para(view({ header: profileHeader(page, PATH), sync: null }).body, 'character-descriptor');
    expect(line).toMatch(
      /<a [^>]*data-testid="character-guild-link">&lt;Olympus XXVII Veterans of the Long March(&gt;|>)<\/a><span class="whitespace-nowrap">\u2060<span [^>]*><span [^>]*data-testid="character-ruleset-chip">Hardcore US<\/span><\/span><\/span><\/span><\/p>/,
    );
    // The guild’s segment moves whole when it fits a line and wraps inside itself when it
    // does not (an inline-block at most the line’s width), the link turning wrapping back on.
    expect(line).toMatch(
      /<span class="inline-block max-w-full whitespace-nowrap" data-testid="character-segment"><a class="[^"]*\bwhitespace-normal\b/,
    );
  });

  it('ends the level line with the chip, 12px after "Level 60 Night Elf"', () => {
    const { body } = view({ header: profileHeader(FULL, PATH), sync: null });
    const levelLine = para(body, 'character-level-line');
    expect(levelLine).toMatch(
      /data-testid="character-level-race">Level 60 Night Elf<\/span><span class="[^"]*\bml-3\b[^"]*"><span class="[^"]*" data-testid="character-ruleset-chip">Hardcore US<\/span><\/span><\/p>/,
    );
    expect(classes(tag(body, 'character-level-line'))).toContain('whitespace-nowrap');
    expect(para(body, 'character-descriptor')).not.toContain('character-ruleset-chip');
  });

  it('ends the first line with the chip when nothing else is known: "Holy Priest" then the chip', () => {
    const { body } = view({ header: profileHeader(BARE, PATH), sync: null });
    expect(body).not.toContain('character-level-line');
    const line = para(body, 'character-descriptor');
    // The last word and the chip never part (fix round 3 item 2: a word joiner in a nowrap span
    // glues the chip to it), inside the spec’s own nowrap segment (fix round 5 item 6).
    expect(line).toMatch(
      /<span style="color: var\(--color-class-priest\)" data-testid="character-spec">Holy Priest<\/span><span class="whitespace-nowrap">\u2060<span class="[^"]*\bml-3\b[^"]*"><span class="[^"]*" data-testid="character-ruleset-chip">Hardcore US<\/span><\/span><\/span><\/span><\/p>/,
    );
  });

  it('draws the chip with muted text and one border token, whichever line it ends', () => {
    for (const page of [FULL, BARE]) {
      const { body } = view({ header: profileHeader(page, PATH), sync: null });
      const chip = classes(tag(body, 'character-ruleset-chip'));
      expect(chip).toEqual(expect.arrayContaining(['pill', 'text-muted', 'border-line']));
      // 12px from the text before it (W.9); after the race, on a phone, it drops to a line of its own.
      const box =
        /<span class="([^"]*)"><span class="[^"]*" data-testid="character-ruleset-chip"/.exec(body)?.[1] ??
        '';
      const tokens = box.split(' ');
      expect(tokens).toContain('ml-3');
      if (page === FULL) {
        expect(tokens).toEqual(expect.arrayContaining(['max-sm:block', 'max-sm:ml-0', 'max-sm:w-fit']));
      } else {
        expect(tokens).not.toContain('max-sm:block');
      }
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
