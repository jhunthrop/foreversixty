// web/src/components/character/CharacterGearPanel.test.ts
// Task L8.6 and its fix round 1: the profile's Gear section, rendered with svelte/server (no
// effects run). `preloaded` stands in for the build's items file having loaded (or failed);
// without it the section is still waiting for that file. The e2e spec character-gear.spec.ts
// checks the same in a browser.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE, GEAR_FIXTURE_ITEMS } from '../../fixtures/character-gear';
import { gearCopy } from '../../lib/gear/copy';
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

function rowHtml(body: string, slot: string): string {
  return new RegExp(`<li[^>]*data-testid="character-gear-row-${slot}"[\\s\\S]*?</li>`).exec(body)?.[0] ?? '';
}

function classesOf(html: string, testid: string): string[] {
  return new RegExp(`<[a-z]+ class="([^"]*)"[^>]*data-testid="${testid}"`).exec(html)?.[1].split(' ') ?? [];
}

/** The separator a dot-trailing element draws after itself: a no-break space, then the dot. */
const TRAILING_DOT = "after:content-['\\00a0·']";

describe('CharacterGearPanel (L8.6 fix rounds 1 and 2)', () => {
  it('is the section the Gear link opens, landing 96px under the top', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const section = /<section class="([^"]*)" id="gear" data-testid="character-gear"/.exec(body);
    expect(section?.[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-24', '-order-2']));
  });

  it('titles it "Gear" then "+70 pts across 7 slots · 1 not scored", a heading focus can land on', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const title = /<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '';
    expect(text(title)).toBe('Gear +70 pts across 7 slots · 1 not scored');
    expect(title).toMatch(/^<h2[^>]*tabindex="-1"/);
  });

  it('the title’s dot trails "Gear"; at 480px and under the summary is its own block, no dot (fix 2 ruling 4)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(classesOf(body, 'character-gear-heading')).toEqual(
      expect.arrayContaining([TRAILING_DOT, 'max-[480px]:after:content-none']),
    );
    expect(classesOf(body, 'character-gear-summary')).toContain('max-[480px]:block');
    // The summary's own dot is glued to the word before it.
    expect(body).toMatch(/data-testid="character-gear-summary">\+70 pts across 7 slots\u00a0· 1 not scored</);
  });

  it('explains the two totals directly under the title: 13px, 4px under it, 16px above row 1', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const totals = /<p class="([^"]*)" data-testid="character-gear-totals">([\s\S]*?)<\/p>/.exec(body);
    expect(text(totals?.[2] ?? '')).toBe(
      'The card’s 88 counts the band’s best set. These 70 count the picks you can wear.',
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
    // a screen reader, its text unpainted on the bar, so it is exactly as wide as the name.
    const held = [
      ...rowHtml(body, 'head').matchAll(
        /<span class="([^"]*)" aria-hidden="true"><span class="([^"]*)"><\/span><span class="([^"]*)" data-testid="character-gear-skeleton">([^<]*)<\/span><\/span>/g,
      ),
    ];
    expect(held.map((m) => m[4])).toEqual(['Green Lens', 'Halo of Transcendence']);
    for (const [, box, icon, bar] of held) {
      expect(box.split(' ')).toEqual(
        expect.arrayContaining(['inline-flex', 'min-h-6', 'max-w-full', 'items-center', 'gap-1.5']),
      );
      expect(icon.split(' ')).toEqual(expect.arrayContaining(['h-5', 'w-5', 'bg-[#1a1d24]']));
      expect(bar.split(' ')).toEqual(
        expect.arrayContaining(['bg-[#1a1d24]', 'leading-[14px]', 'text-transparent']),
      );
    }
    // The loaded name's box has the same shape classes, so the two wrap alike.
    const loaded = render(CharacterGearPanel, { props: PROPS }).body;
    const host = /<span class="([^"]*)" data-bis-item="16921"/.exec(loaded)?.[1].split(' ') ?? [];
    expect(host).toEqual(expect.arrayContaining(held[1][1].split(' ')));
    // Nothing opens a tooltip yet, and no name is drawn in its colour.
    expect(body).not.toContain('data-bis-item');
    expect(body).not.toContain('character-gear-pick');
    // "Empty" and "Unknown item" are not names: they draw at once.
    expect(text(rowHtml(body, 'legs'))).toContain('Unknown item');
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket 2 Empty No pick for your band yet');
    // The off hand's swap note is laid out but held, so nothing moves when the file answers.
    expect(rowHtml(body, 'off_hand')).toMatch(
      /<span class="invisible">[^<]*comes with the main-hand swap<\/span>/,
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
    expect(host?.[1].split(' ')).toEqual(expect.arrayContaining(['relative', 'hover:underline', 'min-h-6']));
    expect(host?.[2]).toContain('tabindex="0"');
    expect(host?.[2]).toContain('role="button"');
    expect(host?.[2]).toContain('aria-haspopup="true"');
    expect(head).toMatch(/<span class="text-rarity-epic-text"[^>]*data-testid="character-gear-pick">Halo/);
    expect(head).toMatch(
      /<span class="text-rarity-rare-text"[^>]*data-testid="character-gear-worn">Green Lens/,
    );
    const legs = rowHtml(body, 'legs');
    expect(legs).not.toContain('data-bis-item="99001"');
    expect(legs).toMatch(/<span class="text-muted"[^>]*data-testid="character-gear-worn">Unknown item/);
  });

  it('every icon sits in a 20px bordered box filled #12151c; no icon name, the box alone (fix 2 ruling 1)', () => {
    const noIcon: GearItemTable = new Map(
      [...FACTS].map(([id, facts]) => [id, id === 16921 ? { ...facts, icon: null } : facts]),
    );
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: noIcon } });
    const head = rowHtml(body, 'head');
    const boxes = [
      ...head.matchAll(
        /<span class="([^"]*)" style="border-color: ?([^;"]+);?" aria-hidden="true" data-testid="character-gear-icon">([\s\S]*?)<\/span>/g,
      ),
    ];
    expect(boxes).toHaveLength(2);
    for (const box of boxes)
      expect(box[1].split(' ')).toEqual(
        expect.arrayContaining(['h-5', 'w-5', 'border', 'bg-[#12151c]', 'overflow-hidden']),
      );
    // Green Lens has its icon, 20 by 20; Halo of Transcendence (no icon name) keeps the empty box.
    expect(boxes[0][3]).toMatch(
      /<img src="\/data\/[^"]+\/icons\/inv_helmet_44\.webp" alt="" width="20" height="20"/,
    );
    expect(boxes[1][2]).toBe('var(--color-rarity-epic)');
    expect(boxes[1][3]).not.toContain('<img');
  });

  it('a failed items file draws the names in the body colour, and the empty off hand no figure (fix 2 ruling 7)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } });
    expect(rowHtml(body, 'head')).toMatch(/<span class="text-text"[^>]*data-testid="character-gear-pick">/);
    expect(body).not.toContain('text-rarity-');
    const off = rowHtml(body, 'off_hand');
    expect(/data-testid="character-gear-gap">([^<]*)</.exec(off)?.[1]).toBe('');
    expect(text(off)).not.toContain('swap');
  });

  it('copy: "No pick for your band yet", "Best for your band", "17% drop" (fix 2 ruling 6)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket 2 Empty No pick for your band yet');
    expect(text(rowHtml(body, 'ranged'))).toBe("Ranged Skul's Ghastly Touch No pick for your band yet");
    expect(text(rowHtml(body, 'shoulder'))).toBe('Shoulder Pauldrons of Transcendence Best for your band');
    expect(text(rowHtml(body, 'head'))).toBe(
      "Head Green Lens → best for your band: Halo of Transcendence Onyxia's Lair: Onyxia · 17% drop +18.4 pts",
    );
  });

  it('no separator starts a text: the dots are glued to what precedes them (fix 2 ruling 4)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const texts = [...body.replace(/<!--[\s\S]*?-->/g, '').matchAll(/>([^<]+)</g)].map((m) => m[1]);
    for (const piece of texts) expect(piece, JSON.stringify(piece)).not.toMatch(/^[ \t\r\n]*·/);
    // Below 1024px a note on the items' line has its dot drawn after the items.
    expect(classesOf(rowHtml(body, 'shoulder'), 'character-gear-items')).toEqual(
      expect.arrayContaining(["max-lg:after:content-['\\00a0·']"]),
    );
    expect(classesOf(rowHtml(body, 'head'), 'character-gear-items')).not.toContain(
      "max-lg:after:content-['\\00a0·']",
    );
    expect(text(rowHtml(body, 'off_hand'))).toContain('9% drop · comes with the main-hand swap');
    expect(rowHtml(body, 'off_hand')).toMatch(/\u00a0· comes with the main-hand swap/);
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
      expect.arrayContaining(['col-start-2', 'row-start-1', 'text-right', 'font-mono', 'leading-5']),
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
        'hover:bg-[rgba(255,255,255,0.04)]',
        'focus-within:bg-[rgba(255,255,255,0.04)]',
      ]),
    );
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
    expect(text(line?.[2] ?? '')).toBe(gearCopy.noGear);
  });
});
