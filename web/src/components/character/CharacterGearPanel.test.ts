// web/src/components/character/CharacterGearPanel.test.ts
// Task L8.6: the profile's Gear section, rendered with svelte/server (no effects run, so the
// item qualities have not loaded and every name takes the common colour; the e2e spec
// character-gear.spec.ts checks the colours in a browser).
//   Ruling 1. 17 rows in the export's order; the title line; absent for `[]`, one line instead.
//   Ruling 2. The section is the Gear link's anchor; the BiS page link is in its footer.
//   Ruling 3. A muted footer sentence says how the slot gaps differ from "below its top".
//   Ruling 4. Two lines a row on a phone, 44px rows, nothing wider than the page.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GEAR_FIXTURE } from '../../fixtures/character-gear';
import { gearCopy } from '../../lib/gear/copy';
import CharacterGearPanel from './CharacterGearPanel.svelte';

const PROPS = {
  gear: GEAR_FIXTURE,
  total: 70,
  classSlug: 'priest',
  bisHref: '/bis/priest/discipline?faction=alliance#band-alliance-60',
  spec: 'Discipline',
  band: 60,
};

function text(html: string): string {
  return html
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, '')
    .replace(/\s+/g, ' ')
    .trim();
}

function rowHtml(body: string, slot: string): string {
  return new RegExp(`<li[^>]*data-testid="character-gear-row-${slot}"[\\s\\S]*?</li>`).exec(body)?.[0] ?? '';
}

describe('CharacterGearPanel (L8.6)', () => {
  it('is a section with the anchor the Gear link opens, landing 96px under the top (ruling 2)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const section = /<section class="([^"]*)" id="gear" data-testid="character-gear"/.exec(body);
    expect(section).not.toBeNull();
    expect(section?.[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-24', '-order-2']));
  });

  it('titles the section "Gear · +70 pts across 7 slots" (ruling 1)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const title = /<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '';
    expect(text(title)).toBe('Gear · +70 pts across 7 slots');
  });

  it('sets the title’s total in the body face, after "Gear" in the section-title face', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const title = /<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '';
    expect(title).toMatch(/<span class="section-title[^"]*">Gear<\/span>/);
    expect(title).toMatch(/<span class="[^"]*font-body[^"]*"> · \+70 pts across 7 slots<\/span>/);
  });

  it('glues the chance to the source’s last word, so a wrap never strands "<1%" or the dot', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(rowHtml(body, 'back')).toMatch(/Dire Maul: Tendris Warpwood(<!--[^>]*-->)*\u00a0·\u00a0&lt;1%/);
  });

  it('draws the 17 rows in the export’s order (ruling 1)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const slots = [...body.matchAll(/data-testid="character-gear-row-([a-z_0-9]+)"/g)].map((m) => m[1]);
    expect(slots).toEqual(GEAR_FIXTURE.map((slot) => slot.slot));
  });

  it('an upgrade row: Head · worn → pick · source · chance, and the gap right-aligned in mono', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    expect(text(head)).toBe(
      'Head Green Lens → best for your band: Halo of Transcendence · Onyxia’s Lair: Onyxia · 17% +18.4 pts',
    );
    const gap = /<span class="([^"]*)"[^>]*data-testid="character-gear-gap"/.exec(head)?.[1].split(' ') ?? [];
    expect(gap).toEqual(expect.arrayContaining(['text-right', 'font-mono', 'tabular']));
    // The arrow is drawn, not read: the screen reader hears "best for your band:".
    expect(head).toMatch(/<span[^>]*aria-hidden="true"[^>]*> → <\/span><span class="sr-only">/);
  });

  it('the names take their item’s quality colour (common until the item table loads)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    expect(head).toMatch(/<span class="text-rarity-common[^"]*"[^>]*data-testid="character-gear-worn"/);
    expect(head).toMatch(/<span class="text-rarity-common[^"]*"[^>]*data-testid="character-gear-pick"/);
  });

  it('the worn item that is the pick reads "Best for your band", muted, with no arrow and no gap', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const shoulder = rowHtml(body, 'shoulder');
    expect(text(shoulder)).toBe('Shoulder Pauldrons of Transcendence · Best for your band');
    expect(shoulder).not.toContain('→');
    expect(shoulder).toMatch(/<span class="[^"]*text-muted[^"]*"[^>]*data-testid="character-gear-meta"/);
  });

  it('an empty slot reads "Empty", muted, with the pick', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const neck = rowHtml(body, 'neck');
    expect(text(neck)).toBe(
      'Neck Empty → best for your band: Choker of the Fire Lord · Molten Core: Ragnaros · 12% +12.6 pts',
    );
    expect(neck).toMatch(/<span class="text-muted"[^>]*data-testid="character-gear-empty"/);
  });

  it('a slot with no pick reads only the worn item; nothing at all reads "Empty"', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'ranged'))).toBe('Ranged Skul’s Ghastly Touch');
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket Empty');
  });

  it('a source with no rate shows no chance; an unscored gap is a dash named "Not scored"', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'chest'))).toBe(
      'Chest Robes of the Exalted → best for your band: Robe of the Archmage · Tailoring +11.3 pts',
    );
    const legs = rowHtml(body, 'legs');
    expect(text(legs)).toBe(
      'Legs Unknown item → best for your band: Pants of Prophecy · Molten Core: Magmadar · 18% –',
    );
    expect(legs).toMatch(/aria-label="Not scored"[^>]*data-testid="character-gear-gap"/);
  });

  it('the footer: the muted sentence apart from "below its top", then the BiS page link (rulings 2, 3)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const footer = /<footer[^>]*data-testid="character-gear-footer"[\s\S]*?<\/footer>/.exec(body)?.[0] ?? '';
    expect(footer).toMatch(
      new RegExp(
        `<p class="[^"]*text-muted[^"]*"[^>]*data-testid="character-gear-difference">${gearCopy.difference}</p>`,
      ),
    );
    const link = /<a[^>]*href="([^"]*)"[^>]*data-testid="character-gear-bis"[\s\S]*?<\/a>/.exec(footer);
    expect(link?.[1]).toBe('/bis/priest/discipline?faction=alliance#band-alliance-60');
    expect(text(link?.[0] ?? '')).toBe('Every pick for Discipline at 60 →');
  });

  it('no BiS page known: the footer keeps its sentence and has no link', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, bisHref: null, spec: null } });
    expect(body).toContain('data-testid="character-gear-difference"');
    expect(body).not.toContain('character-gear-bis');
  });

  it('`gear: []`: no section, one line in its place that the Gear link still lands on (ruling 1)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, gear: [], total: null } });
    expect(body).not.toContain('<section');
    expect(body).not.toContain('character-gear-row-');
    const line = /<p class="([^"]*)" id="gear" data-testid="character-gear-none">([\s\S]*?)<\/p>/.exec(body);
    expect(line).not.toBeNull();
    expect(text(line?.[2] ?? '')).toBe(gearCopy.noGear);
    expect(line?.[1].split(' ')).toEqual(expect.arrayContaining(['-order-2', 'scroll-mt-24']));
  });

  it('a phone row is two lines, worn → pick then source · chance with the gap; 44px rows (ruling 4)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    const li = /<li class="([^"]*)"/.exec(head)?.[1].split(' ') ?? [];
    expect(li).toEqual(
      expect.arrayContaining([
        'grid',
        'min-h-11',
        'grid-cols-[72px_minmax(0,1fr)_auto]',
        'md:grid-cols-[112px_fit-content(60%)_minmax(0,1fr)_88px]',
      ]),
    );
    const cell = (testid: string) =>
      new RegExp(`<span class="([^"]*)"[^>]*data-testid="${testid}"`).exec(head)?.[1].split(' ') ?? [];
    // Line 1: the items across the two right columns; line 2: the source, then the gap.
    expect(cell('character-gear-items')).toEqual(
      expect.arrayContaining(['col-span-2', 'row-start-1', 'md:col-span-1']),
    );
    expect(cell('character-gear-meta')).toEqual(
      expect.arrayContaining(['col-start-2', 'row-start-2', 'md:col-start-3', 'md:row-start-1']),
    );
    expect(cell('character-gear-gap')).toEqual(
      expect.arrayContaining(['col-start-3', 'row-start-2', 'md:col-start-4', 'md:row-start-1']),
    );
    // Long names wrap rather than widen the page.
    expect(cell('character-gear-items')).toEqual(expect.arrayContaining(['min-w-0', 'break-words']));
  });

  it('the meta line starts with its dot only from md, where it follows the pick on one line', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    for (const slot of ['head', 'shoulder']) {
      const meta = /data-testid="character-gear-meta">(<span[^>]*>[^<]*<\/span>)/.exec(rowHtml(body, slot));
      expect(meta?.[1], slot).toBe('<span class="max-md:hidden" aria-hidden="true">· </span>');
    }
  });
});
