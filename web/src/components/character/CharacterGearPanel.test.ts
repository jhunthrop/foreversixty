// web/src/components/character/CharacterGearPanel.test.ts
// Task L8.6 and its fix round 1: the profile's Gear section, rendered with svelte/server (no
// effects run). `preloaded` stands in for the build's items file having loaded (or failed);
// without it the section is still waiting for that file. The e2e spec character-gear.spec.ts
// checks the same in a browser.
import { readFileSync } from 'node:fs';
import { compile } from 'svelte/compiler';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_FIXTURE_ITEMS, GEAR_FIXTURE_LONG_NAME } from '../../fixtures/character-gear';
import { gearCopy, notReadCopy } from '../../lib/gear/copy';
import type { GearItemTable } from '../../lib/gear/qualities';
import CharacterGearPanel from './CharacterGearPanel.svelte';

const FACTS: GearItemTable = new Map(
  GEAR_FIXTURE_ITEMS.map((item) => [
    item.id,
    { quality: item.quality, icon: item.icon, twoHand: item.two_hand === true },
  ]),
);

const PROPS = {
  gear: GEAR_FIXTURE,
  total: 70,
  cardGap: 88,
  classSlug: 'priest',
  bisHref: '/bis/priest/discipline?faction=alliance#band-alliance-60',
  spec: 'Discipline',
  band: 60,
  preloaded: FACTS,
};

function text(html: string): string {
  return html
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/&#39;/g, "'")
    .replace(/&lt;/g, '<')
    .replace(/\s+/g, ' ')
    .trim();
}

/** The markup without Svelte's hydration comments. */
function bare(html: string): string {
  return html.replace(/<!--[\s\S]*?-->/g, '');
}

function rowHtml(body: string, slot: string): string {
  return new RegExp(`<li[^>]*data-testid="character-gear-row-${slot}"[\\s\\S]*?</li>`).exec(body)?.[0] ?? '';
}

function classesOf(html: string, testid: string): string[] {
  return new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(html)?.[1].split(' ') ?? [];
}

/** Hover styling only where there is a fine pointer that hovers (fix 4 item C). */
const POINTER_HOVER = '[@media(hover:hover)_and_(pointer:fine)]:hover:';
/** A separator drawn before its text, never a text node of its own (fix 4 item B), from the
 *  component's own CSS so that no screen reader reads it (fix 5 item 2). */
const DOT_BEFORE = 'gear-dot-before';

/** The separator a dot-trailing element draws after itself: a no-break space, then the dot. */
const TRAILING_DOT = 'gear-dot-after';

/** The rows with no source line, whose note is the meta line's whole text. */
const NOTE_ROWS = ['shoulder', 'hands', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2', 'ranged'];

/** The component's compiled CSS (its scoped <style>). */
function componentCss(): string {
  const source = readFileSync(new URL('./CharacterGearPanel.svelte', import.meta.url), 'utf8');
  return compile(source, { filename: 'CharacterGearPanel.svelte', css: 'external' }).css?.code ?? '';
}

/** The declarations of the first rule whose selector names `className` and `pseudo`. */
function pseudoRule(css: string, className: string, pseudo: '::before' | '::after'): string {
  const rule = new RegExp(`\\.${className}[^{},]*${pseudo}\\s*\\{([^}]*)\\}`).exec(css);
  return rule?.[1].replace(/\s+/g, ' ').trim() ?? '';
}

describe('CharacterGearPanel (L8.6 fix rounds 1 and 2)', () => {
  it('is the section the Gear link opens, landing 96px under the top', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const section = /<section class="([^"]*)" id="gear" data-testid="character-gear"/.exec(body);
    expect(section?.[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-24', '-order-2']));
  });

  it('titles it "Gear" then the best-in-slot count and "+70 points across 7 slots · 1 not scored", a heading focus can land on', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const title = /<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '';
    expect(text(title)).toBe('Gear 6 of 15 slots best in slot · +70 points across 7 slots · 1 not scored');
    expect(title).toMatch(/^<h2[^>]*tabindex="-1"/);
  });

  it('the title’s dot trails "Gear"; at 480px and under the summary is its own block, no dot (fix 2 ruling 4)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(classesOf(body, 'character-gear-heading')).toEqual(expect.arrayContaining([TRAILING_DOT]));
    expect(componentCss()).toMatch(
      /@media \(max-width: 480px\)\s*\{\s*\.gear-dot-after[^{]*::after\s*\{\s*content:\s*none;?\s*\}/,
    );
    expect(classesOf(body, 'character-gear-summary')).toContain('max-[480px]:block');
    // The summary's own dot is glued to the word before it.
    expect(body).toMatch(
      /data-testid="character-gear-summary">6 of 15 slots best in slot · \+70 points across 7 slots\u00a0· 1 not scored</,
    );
  });

  it('explains the two totals directly under the title: 13px, 4px under it, 16px above row 1', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const totals = /<p class="([^"]*)" data-testid="character-gear-totals">([\s\S]*?)<\/p>/.exec(body);
    expect(text(totals?.[2] ?? '')).toBe(
      'The card’s 88 counts the level range’s best set. These 70 count the picks you can wear.',
    );
    expect(totals?.[1].split(' ')).toEqual(expect.arrayContaining(['text-muted', 'text-[13px]', 'mt-1']));
    expect(classesOf(body, 'character-gear-list')).toContain('mt-4');
  });

  it('leaves the totals line out when the card has no gear gap', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, cardGap: null } });
    expect(body).not.toContain('character-gear-totals');
  });

  it('keeps only the BiS link in the footer, its target 12px of padding above and below (fix 2 ruling 9)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const footer = /<footer[^>]*data-testid="character-gear-footer"[\s\S]*?<\/footer>/.exec(body)?.[0] ?? '';
    expect(text(footer)).toBe('Full Discipline BiS list at 60 →');
    expect(footer).toContain('href="/bis/priest/discipline?faction=alliance#band-alliance-60"');
    expect(classesOf(footer, 'character-gear-bis')).toContain('py-3');
    expect(classesOf(footer, 'character-gear-bis')).not.toContain('invisible');
  });

  it('mid-load: labels, gaps and sources draw at once; each name is held, unpainted, on a 14px bar; the footer keeps its height (fix 2 ruling 8, fix 3 item 2)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: undefined } });
    const list = /<ul class="([^"]*)"[^>]*data-testid="character-gear-list"[^>]*>/.exec(body);
    expect(list?.[1].split(' ')).not.toContain('invisible');
    expect(list?.[0]).toContain('aria-busy="true"');
    // Each held name is the name's own box -- a 20px icon place, then the real name -- hidden from
    // a screen reader: the bar is the wrapper's background, the name inside it visibility-hidden
    // (its width kept, never found by a page search) (fix 3 item 2, fix 4 item 7).
    const held = [
      ...rowHtml(body, 'head').matchAll(
        /<span class="([^"]*)" aria-hidden="true"><span class="([^"]*)"><\/span><span class="([^"]*)" data-testid="character-gear-skeleton"><span class="invisible">([^<]*)<\/span><\/span><\/span>/g,
      ),
    ];
    expect(held.map((m) => m[4])).toEqual(['Green Lens', 'Halo of Transcendence']);
    for (const [, box, icon, bar] of held) {
      expect(box.split(' ')).toEqual(
        expect.arrayContaining(['inline-flex', 'min-h-6', 'max-w-full', 'items-center', 'gap-1.5']),
      );
      expect(icon.split(' ')).toEqual(expect.arrayContaining(['h-5', 'w-5', 'bg-[#1a1d24]']));
      expect(bar.split(' ')).toEqual(expect.arrayContaining(['bg-[#1a1d24]', 'leading-[14px]']));
      expect(bar.split(' ')).not.toContain('text-transparent');
    }
    // The loaded name's box has the same shape classes, so the two wrap alike.
    const loaded = render(CharacterGearPanel, { props: PROPS }).body;
    const host = /<span class="([^"]*)" data-bis-item="16921"/.exec(loaded)?.[1].split(' ') ?? [];
    expect(host).toEqual(expect.arrayContaining(held[1][1].split(' ')));
    // Nothing opens a tooltip yet, and no name is drawn in its colour.
    expect(body).not.toContain('data-bis-item');
    expect(body).not.toContain('character-gear-pick');
    // "Empty" and "Not in our item data yet" are not names: they draw at once.
    expect(text(rowHtml(body, 'legs'))).toContain('Not in our item data yet');
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket 2 Empty No pick for your level range yet');
    // The off hand's swap note's line is reserved as a bar of its width (fix 4 item 4).
    expect(bare(rowHtml(body, 'off_hand'))).toMatch(
      /<span class="[^"]*bg-\[#1a1d24\][^"]*" data-testid="character-gear-note-bar"><span class="[^"]*invisible[^"]*"[^>]*>comes with the <span class="whitespace-nowrap">main-hand swap<\/span><\/span><\/span>/,
    );
    expect(/data-testid="character-gear-gap">([^<]*)</.exec(rowHtml(body, 'off_hand'))?.[1]).toBe('');
    // The footer link holds its place (visibility, not display).
    expect(classesOf(body, 'character-gear-bis')).toContain('invisible');
  });

  it('draws the 17 rows in the export’s order once the file is in', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(body).not.toContain('character-gear-skeleton');
    const slots = [...body.matchAll(/data-testid="character-gear-row-([a-z_0-9]+)"/g)].map((m) => m[1]);
    expect(slots).toEqual(GEAR_FIXTURE.map((slot) => slot.slot));
  });

  it('marks each known name the way /bis rows are marked: colour, icon box, a 24px target (fix 2 ruling 9)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    const host = /<span class="([^"]*)" data-bis-item="16921"([^>]*)>([\s\S]*?)<\/span><\/span>/.exec(head);
    expect(host?.[1].split(' ')).toEqual(
      expect.arrayContaining(['relative', `${POINTER_HOVER}underline`, 'min-h-6']),
    );
    expect(host?.[1].split(' ')).not.toContain('hover:underline');
    expect(host?.[2]).toContain('tabindex="0"');
    expect(host?.[2]).toContain('role="button"');
    expect(host?.[2]).toContain('aria-haspopup="true"');
    expect(head).toMatch(/<span class="text-rarity-epic-text"[^>]*data-testid="character-gear-pick">Halo/);
    expect(head).toMatch(
      /<span class="text-rarity-rare-text"[^>]*data-testid="character-gear-worn">Green Lens/,
    );
    const legs = rowHtml(body, 'legs');
    expect(legs).not.toContain('data-bis-item="99001"');
    expect(legs).toMatch(
      /<span class="text-muted"[^>]*data-testid="character-gear-worn">Not in our item data yet/,
    );
  });

  it('every icon sits in a 20px box filled #14171f; with no image it takes the muted line border (fix 4 item A)', () => {
    const noIcon: GearItemTable = new Map(
      [...FACTS].map(([id, facts]) => [id, id === 16921 ? { ...facts, icon: null } : facts]),
    );
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: noIcon } });
    const head = rowHtml(body, 'head');
    const boxes = [
      ...head.matchAll(
        /<span class="([^"]*)"(?: style="border-color: ?([^;"]+);?")? aria-hidden="true" data-testid="character-gear-icon">([\s\S]*?)<\/span>/g,
      ),
    ];
    expect(boxes).toHaveLength(2);
    for (const box of boxes)
      expect(box[1].split(' ')).toEqual(
        expect.arrayContaining(['h-5', 'w-5', 'border', 'bg-[#14171f]', 'overflow-hidden']),
      );
    // Green Lens has its icon, 20 by 20, in its quality's border.
    expect(boxes[0][3]).toMatch(
      /<img src="\/data\/[^"]+\/icons\/inv_helmet_44\.webp" alt="" width="20" height="20"/,
    );
    expect(boxes[0][2]).toBe('var(--color-rarity-rare)');
    // Halo of Transcendence (no icon name): the skeleton's treatment, the muted line border, no
    // quality colour, never a white checkbox.
    expect(boxes[1][1].split(' ')).toContain('border-line');
    expect(boxes[1][2]).toBeUndefined();
    expect(boxes[1][3]).not.toContain('<img');
  });

  it('a failed items file draws the names in the body colour, and the empty off hand no figure (fix 2 ruling 7)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } });
    expect(rowHtml(body, 'head')).toMatch(/<span class="text-text"[^>]*data-testid="character-gear-pick">/);
    expect(body).not.toContain('text-rarity-');
    const off = rowHtml(body, 'off_hand');
    expect(/data-testid="character-gear-gap">([^<]*)</.exec(off)?.[1]).toBe('');
    // The note's line is kept, blank: the same rows as mid-load (fix 4 item 4).
    expect(bare(off)).toMatch(
      /<span class="[^"]*invisible[^"]*" data-testid="character-gear-note">comes with the <span class="whitespace-nowrap">main-hand swap<\/span><\/span>/,
    );
    expect(off).not.toContain('character-gear-note-bar');
    // Every icon box takes the muted line border: no quality is known.
    expect(body).not.toContain('style="border-color');
    // One muted line says so, in the footer (fix 6 item 2).
    expect(text(/<p[^>]*data-testid="character-gear-failed">[^<]*<\/p>/.exec(body)?.[0] ?? '')).toBe(
      'Item details didn’t load.',
    );
    expect(render(CharacterGearPanel, { props: PROPS }).body).not.toContain('character-gear-failed');
  });

  it('a failed items file: every name is drawn, none held, and each note shows (fix 6 item 1)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } });
    expect(body).not.toContain('character-gear-skeleton');
    expect(body).not.toContain('aria-busy');
    // The only thing kept unpainted is the off hand's swap note (fix 4 item 4).
    for (const slot of GEAR_FIXTURE.map((s) => s.slot)) {
      const row = rowHtml(body, slot);
      const unpainted = [...row.matchAll(/<span class="([^"]*)"[^>]*data-testid="([^"]*)"/g)]
        .filter((m) => m[1].split(' ').includes('invisible'))
        .map((m) => m[2]);
      expect(unpainted, slot).toEqual(slot === 'off_hand' ? ['character-gear-note'] : []);
    }
    // A worn item with no pick: the name in the body colour, its note beside it.
    const ranged = rowHtml(body, 'ranged');
    expect(ranged).toMatch(
      /<span class="text-text"[^>]*data-testid="character-gear-worn">Skul's Ghastly Touch<\/span>/,
    );
    expect(text(ranged)).toBe("Ranged Skul's Ghastly Touch No pick for your level range yet");
  });

  it('the failed line sits in the footer under the link, so nothing above it moves (fix 6 item 2)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } });
    const footer = /<footer[^>]*data-testid="character-gear-footer"[\s\S]*?<\/footer>/.exec(body)?.[0] ?? '';
    expect(footer).toMatch(/data-testid="character-gear-bis"[\s\S]*data-testid="character-gear-failed"/);
    expect(text(footer)).toBe('Full Discipline BiS list at 60 → Item details didn’t load.');
    expect(classesOf(footer, 'character-gear-failed')).toEqual(
      expect.arrayContaining(['text-muted', 'text-[13px]', 'leading-5']),
    );
    // Nothing between the totals line and the list.
    expect(body).toMatch(/data-testid="character-gear-totals">[^<]*<\/p>\s*(<!--[^>]*-->\s*)*<ul/);
    // With no BiS link the footer holds the line alone.
    const bare = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null, bisHref: null } }).body;
    const alone = /<footer[^>]*data-testid="character-gear-footer"[\s\S]*?<\/footer>/.exec(bare)?.[0] ?? '';
    expect(text(alone)).toBe('Item details didn’t load.');
    expect(alone).not.toContain('character-gear-bis');
    // Loaded or loading, no footer line.
    expect(render(CharacterGearPanel, { props: { ...PROPS, preloaded: undefined } }).body).not.toContain(
      'character-gear-failed',
    );
  });

  it('"main-hand swap" never breaks: kept on one line in every state (fix 6 item 3)', () => {
    const kept = '>comes with the <span class="whitespace-nowrap">main-hand swap</span></span>';
    for (const preloaded of [FACTS, null, undefined]) {
      const off = rowHtml(render(CharacterGearPanel, { props: { ...PROPS, preloaded } }).body, 'off_hand');
      expect(bare(off), String(preloaded)).toContain(kept);
      expect(text(off), String(preloaded)).toContain('comes with the main-hand swap');
    }
    // The words are unchanged.
    expect(gearCopy.swap).toBe('comes with the main-hand swap');
  });

  it('copy: "No pick for your level range yet", "Already your best at this level", "17% drop" (fix 2 ruling 6)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket 2 Empty No pick for your level range yet');
    expect(text(rowHtml(body, 'ranged'))).toBe(
      "Ranged Skul's Ghastly Touch No pick for your level range yet",
    );
    expect(text(rowHtml(body, 'shoulder'))).toBe(
      'Shoulder Pauldrons of Transcendence Already your best at this level',
    );
    expect(text(rowHtml(body, 'head'))).toBe(
      "Head Green Lens → best at this level: Halo of Transcendence Onyxia's Lair: Onyxia · 17% drop +18.4 points",
    );
    expect(text(rowHtml(body, 'legs'))).toContain('Not in our item data yet');
  });

  it('no text begins or ends with a separator: each dot is drawn before its text (fix 4 item B)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const texts = [...body.replace(/<!--[\s\S]*?-->/g, '').matchAll(/>([^<]+)</g)].map((m) => m[1].trim());
    for (const piece of texts) {
      // The one separator between a source and its drop chance is a text node of its own, aria-hidden.
      if (piece === '·') continue;
      expect(piece, JSON.stringify(piece)).not.toMatch(/^·/);
      expect(piece, JSON.stringify(piece)).not.toMatch(/·$/);
    }
    expect(classesOf(rowHtml(body, 'shoulder'), 'character-gear-items').join(' ')).not.toContain('after:');
    // The swap note draws its dot before itself; the chance is separated from its source by a real
    // " · " in the text (aria-hidden), so "Skum · 50% drop" reads that way wherever the text is read.
    const off = rowHtml(body, 'off_hand');
    expect(classesOf(off, 'character-gear-note')).toContain(DOT_BEFORE);
    expect(classesOf(off, 'character-gear-chance')).not.toContain(DOT_BEFORE);
    expect(off).toContain('<span aria-hidden="true"> · </span>');
    expect(text(off).split('·').length - 1).toBe(1);
  });

  it('below 1024px a row’s note is a line of its own under the items, with no dot, at every width (fix 5 item 1)', () => {
    for (const gear of [GEAR_FIXTURE, GEAR_FIXTURE_LONG_NAME]) {
      const { body } = render(CharacterGearPanel, { props: { ...PROPS, gear } });
      for (const slot of NOTE_ROWS) {
        const row = rowHtml(body, slot);
        const note = classesOf(row, 'character-gear-meta');
        // The same full-width line as a source line: no width puts it beside the items.
        expect(note, slot).toContain('basis-full');
        expect(
          note.filter((c) => /before:|after:|gear-dot|min-\[|max-\[/.test(c)),
          slot,
        ).toEqual([]);
        expect(row, slot).not.toContain('·');
      }
    }
    const long = render(CharacterGearPanel, { props: { ...PROPS, gear: GEAR_FIXTURE_LONG_NAME } }).body;
    expect(text(rowHtml(long, 'shoulder'))).toBe(
      'Shoulder Mantle of the Unforgiving Majesty Already your best at this level',
    );
  });

  it('every drawn separator is read by no screen reader: content "·" / "", after the plain dot for a browser without that form (fix 5 item 2)', () => {
    const css = componentCss();
    expect(pseudoRule(css, DOT_BEFORE, '::before')).toBe(
      "content: '\\00a0·\\00a0'; content: '\\00a0·\\00a0' / '';",
    );
    expect(pseudoRule(css, TRAILING_DOT, '::after')).toBe("content: '\\00a0·'; content: '\\00a0·' / '';");
    // No drawn dot comes from a content utility: one declaration, so no fallback.
    for (const preloaded of [FACTS, null, undefined]) {
      const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded } });
      expect(body).not.toMatch(/content-\[/);
    }
  });

  it('an unscored gap reads "Not scored", muted and right-aligned', () => {
    const legs = rowHtml(render(CharacterGearPanel, { props: PROPS }).body, 'legs');
    expect(text(legs)).toMatch(/Not scored$/);
    expect(classesOf(legs, 'character-gear-gap')).toEqual(
      expect.arrayContaining(['text-muted', 'text-right']),
    );
  });

  it('the five largest gaps in 700, the rest in 400, in the same green', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    for (const slot of ['head', 'neck', 'back', 'chest', 'waist'])
      expect(classesOf(rowHtml(body, slot), 'character-gear-gap'), slot).toEqual(
        expect.arrayContaining(['text-kill', 'font-bold']),
      );
    for (const slot of ['wrist', 'main_hand'])
      expect(classesOf(rowHtml(body, slot), 'character-gear-gap'), slot).toEqual(
        expect.arrayContaining(['text-kill', 'font-normal']),
      );
  });

  it('names each row by its slot label and the list "Gear by slot" (fix 2 ruling 9)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(body).toMatch(/<ul [^>]*aria-label="Gear by slot"[^>]*data-testid="character-gear-list"/);
    for (const slot of GEAR_FIXTURE.map((s) => s.slot)) {
      const row = rowHtml(body, slot);
      expect(row, slot).toMatch(new RegExp(`^<li [^>]*aria-labelledby="gear-slot-${slot}"`));
      expect(row, slot).toMatch(
        new RegExp(`<span class="[^"]*" id="gear-slot-${slot}" data-testid="character-gear-slot">`),
      );
    }
  });

  it('phone: line 1 the label and the gap, line 2 the items, line 3 the source; 44px rows', () => {
    const head = rowHtml(render(CharacterGearPanel, { props: PROPS }).body, 'head');
    expect(/<li [^>]*class="([^"]*)"/.exec(head)?.[1].split(' ')).toEqual(
      expect.arrayContaining(['grid', 'min-h-11', 'grid-cols-[minmax(0,1fr)_auto]']),
    );
    expect(classesOf(head, 'character-gear-slot')).toEqual(
      expect.arrayContaining(['col-start-1', 'row-start-1', 'leading-5']),
    );
    expect(classesOf(head, 'character-gear-gap')).toEqual(
      expect.arrayContaining(['col-start-2', 'row-start-1', 'text-right', 'numeric', 'leading-5']),
    );
    expect(classesOf(head, 'character-gear-line')).toEqual(
      expect.arrayContaining(['col-span-2', 'row-start-2']),
    );
    expect(classesOf(head, 'character-gear-meta')).toContain('basis-full');
  });

  it('desktop: slot 104 | worn → pick 5fr | source 4fr | gap 88 (fix 2 ruling 3)', () => {
    const head = rowHtml(render(CharacterGearPanel, { props: PROPS }).body, 'head');
    expect(/<li [^>]*class="([^"]*)"/.exec(head)?.[1].split(' ')).toEqual(
      expect.arrayContaining([
        'lg:grid-cols-[104px_minmax(0,5fr)_minmax(0,4fr)_88px]',
        `${POINTER_HOVER}bg-[rgba(255,255,255,0.04)]`,
        'has-[:focus-visible]:bg-[rgba(255,255,255,0.04)]',
      ]),
    );
    expect(/<li [^>]*class="([^"]*)"/.exec(head)?.[1]).not.toMatch(/(^| )(hover|focus-within):/);
    expect(classesOf(head, 'character-gear-line')).toContain('lg:contents');
    expect(classesOf(head, 'character-gear-meta')).toEqual(
      expect.arrayContaining(['lg:col-start-3', 'lg:row-start-1']),
    );
  });

  it('`gear: []`: the "Gear" heading over one line that keeps the anchor', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, gear: [], total: null } });
    expect(body).not.toContain('character-gear-row-');
    const heading = /<h2[^>]*tabindex="-1"[^>]*data-testid="character-gear-title"[^>]*>([\s\S]*?)<\/h2>/.exec(
      body,
    );
    expect(text(heading?.[1] ?? '')).toBe('Gear');
    expect(classesOf(body, 'character-gear-heading')).not.toContain(TRAILING_DOT);
    const line = /<p class="([^"]*)" id="gear" data-testid="character-gear-none">([\s\S]*?)<\/p>/.exec(body);
    // The default viewer owns the character: sync, and nothing else (W.10).
    expect(text(line?.[2] ?? '')).toBe('No gear on file yet. Sync the addon.');
    expect(body).not.toContain('character-gear-planner');
  });

  it('`gear: []`, a visitor: the one quiet line says sync or sign in; with a score the planner is the one button (W.10)', () => {
    const visitor = render(CharacterGearPanel, {
      props: { ...PROPS, gear: [], total: null, owned: false },
    }).body;
    const line = /<p class="[^"]*" id="gear" data-testid="character-gear-none">([^]*?)[<][/]p>/.exec(visitor);
    expect(text(line?.[1] ?? '')).toBe(gearCopy.noGear);
    expect(line?.[1]).toContain('href="/setup"');
    expect(line?.[1]).toContain('href="/login"');
    expect(visitor).not.toContain('character-gear-planner');
    const scored = render(CharacterGearPanel, {
      props: { ...PROPS, gear: [], total: null, owned: false, planner: true },
    }).body;
    expect(scored).toMatch(/href="[/]planner"[^>]*data-testid="character-gear-planner">Open the planner</);
  });
});

describe('CharacterGearPanel on an export that carried no gear and no talents (W.3)', () => {
  const UNREAD = GEAR_FIXTURE.map((slot) => ({ ...slot, worn: null, gap: slot.pick === null ? null : 5 }));
  const props = { ...PROPS, gear: UNREAD, unread: true, capturedAt: '2026-10-19T21:40:00', band: 20 };

  it('says the gear has not been read, in place of the slot list, the Empty rows and the totals', () => {
    const body = bare(render(CharacterGearPanel, { props }).body);
    expect(body).toContain('data-testid="export-not-read"');
    expect(text(body)).toContain(notReadCopy.lead);
    expect(text(body)).toContain('carried no gear or talents.');
    expect(text(body)).toContain(notReadCopy.steps);
    expect(body).not.toContain('character-gear-empty');
    expect(body).not.toContain('character-gear-totals');
    expect(body).not.toContain('character-gear-gap');
    expect(text(body)).not.toContain(gearCopy.empty);
    expect(text(body)).not.toContain(gearCopy.noPick);
    expect(text(body)).not.toContain('pts');
  });

  it('keeps the band’s best in slot below it, titled for the band, picks only', () => {
    const body = bare(render(CharacterGearPanel, { props }).body);
    expect(text(/data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '')).toContain(
      'Best in slot for level 20 to 29',
    );
    expect(body.indexOf('export-not-read')).toBeLessThan(body.indexOf('character-gear-title'));
    expect(body).toContain('character-gear-row-head');
    expect(body).toContain('Halo of Transcendence');
    expect(body).not.toContain('character-gear-worn');
  });

  it('speaks to a visitor without the owner’s steps', () => {
    const body = text(bare(render(CharacterGearPanel, { props: { ...props, owned: false } }).body));
    expect(body).toContain(notReadCopy.visitorLead);
    expect(body).not.toContain(notReadCopy.steps);
  });

  it('titles the list by band: 60 alone, a range below it, plain with none', () => {
    expect(notReadCopy.bestInSlot(60)).toBe('Best in slot for level 60');
    expect(notReadCopy.bestInSlot(20)).toBe('Best in slot for level 20 to 29');
    expect(notReadCopy.bestInSlot(null)).toBe('Best in slot');
    expect(notReadCopy.synced('19 Oct, 21:40')).toBe(
      'The last sync from the game (19 Oct, 21:40) carried no gear or talents.',
    );
  });
});

describe('the spec the picks are for (A.2)', () => {
  it('names the spec and its source after the heading when the talents named it', () => {
    const body = render(CharacterGearPanel, {
      props: { ...PROPS, specLine: 'Fury · from your talents' },
    }).body;
    expect(bare(body)).toContain('data-testid="character-gear-spec"');
    expect(text(body)).toContain('Fury · from your talents');
  });

  it('draws nothing extra when a fight named the spec', () => {
    const body = render(CharacterGearPanel, { props: PROPS }).body;
    expect(bare(body)).not.toContain('character-gear-spec');
  });
});
