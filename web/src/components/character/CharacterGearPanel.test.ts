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

describe('CharacterGearPanel (L8.6 fix round 1)', () => {
  it('is the section the Gear link opens, landing 96px under the top', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const section = /<section class="([^"]*)" id="gear" data-testid="character-gear"/.exec(body);
    expect(section?.[1].split(' ')).toEqual(expect.arrayContaining(['scroll-mt-24', '-order-2']));
  });

  it('titles it "Gear · +70 pts across 7 slots · 1 not scored", a heading focus can land on (rulings 3, 11)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const title = /<h2[^>]*data-testid="character-gear-title"[\s\S]*?<\/h2>/.exec(body)?.[0] ?? '';
    expect(text(title)).toBe('Gear · +70 pts across 7 slots · 1 not scored');
    expect(title).toMatch(/^<h2[^>]*tabindex="-1"/);
  });

  it('explains the two totals directly under the title: 13px, 4px under it, 16px above row 1 (ruling 3)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const totals = /<p class="([^"]*)" data-testid="character-gear-totals">([\s\S]*?)<\/p>/.exec(body);
    expect(text(totals?.[2] ?? '')).toBe(
      'The card’s 88 counts the band’s best set. These 70 count the picks you can wear.',
    );
    expect(totals?.[1].split(' ')).toEqual(expect.arrayContaining(['text-muted', 'text-[13px]', 'mt-1']));
    expect(body.indexOf('character-gear-totals')).toBeGreaterThan(body.indexOf('character-gear-title'));
    expect(body.indexOf('character-gear-totals')).toBeLessThan(body.indexOf('character-gear-list'));
    expect(classesOf(body, 'character-gear-list')).toContain('mt-4');
  });

  it('leaves the totals line out when the card has no gear gap', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, cardGap: null } });
    expect(body).not.toContain('character-gear-totals');
  });

  it('keeps only the BiS link in the footer: "Full Discipline BiS list at 60 →" (ruling 3)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const footer = /<footer[^>]*data-testid="character-gear-footer"[\s\S]*?<\/footer>/.exec(body)?.[0] ?? '';
    expect(text(footer)).toBe('Full Discipline BiS list at 60 →');
    expect(footer).toContain('href="/bis/priest/discipline?faction=alliance#band-alliance-60"');
    expect(body).not.toContain('Below its top measures');
  });

  it('waits for the items file before drawing the rows; the title and totals draw at once (ruling 2)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: undefined } });
    const list = /<ul class="([^"]*)"[^>]*data-testid="character-gear-list"[^>]*>/.exec(body);
    expect(list?.[1].split(' ')).toContain('invisible');
    expect(list?.[0]).toContain('aria-busy="true"');
    expect(classesOf(body, 'character-gear-totals')).not.toContain('invisible');
    expect(body).not.toContain('text-rarity-common');
  });

  it('draws the 17 rows in the export’s order once the file is in', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(classesOf(body, 'character-gear-list')).not.toContain('invisible');
    const slots = [...body.matchAll(/data-testid="character-gear-row-([a-z_0-9]+)"/g)].map((m) => m[1]);
    expect(slots).toEqual(GEAR_FIXTURE.map((slot) => slot.slot));
  });

  it('marks each known name the way /bis rows are marked, with its quality colour and icon (rulings 1, 10)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    const host = /<span class="([^"]*)" data-bis-item="16921"([^>]*)>([\s\S]*?)<\/span><\/span>/.exec(head);
    expect(host?.[1].split(' ')).toEqual(expect.arrayContaining(['relative', 'hover:underline']));
    expect(host?.[2]).toContain('tabindex="0"');
    expect(host?.[2]).toContain('role="button"');
    expect(host?.[2]).toContain('aria-haspopup="true"');
    expect(host?.[3]).toMatch(
      /<img src="\/data\/[^"]+\/icons\/inv_helmet_24\.webp" alt="" width="20" height="20"[^>]*style="border-color: ?var\(--color-rarity-epic\);?"/,
    );
    expect(head).toMatch(/<span class="text-rarity-epic-text"[^>]*data-testid="character-gear-pick">Halo/);
    expect(head).toMatch(
      /<span class="text-rarity-rare-text"[^>]*data-testid="character-gear-worn">Green Lens/,
    );
    // An item the build does not have is plain muted text: no tooltip host, no icon.
    const legs = rowHtml(body, 'legs');
    expect(legs).not.toContain('data-bis-item="99001"');
    expect(legs).toMatch(/<span class="text-muted"[^>]*data-testid="character-gear-worn">Unknown item/);
  });

  it('a failed items file draws the names in the body colour (ruling 2)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, preloaded: null } });
    expect(rowHtml(body, 'head')).toMatch(/<span class="text-text"[^>]*data-testid="character-gear-pick">/);
    expect(body).not.toContain('text-rarity-');
  });

  it('reads each row as its label first: "Trinket 2, Empty · no pick for your band yet" (rulings 7, 8)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'trinket2'))).toBe('Trinket 2, Empty · no pick for your band yet');
    expect(text(rowHtml(body, 'ranged'))).toBe("Ranged, Skul's Ghastly Touch · no pick for your band yet");
    expect(text(rowHtml(body, 'shoulder'))).toBe('Shoulder, Pauldrons of Transcendence · Best for your band');
    expect(text(rowHtml(body, 'finger1'))).toBe('Ring 1, Blackstone Ring · Best for your band');
  });

  it('an upgrade row: worn → pick, then the source and "17% drop", then the gap (ruling 8)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    expect(text(rowHtml(body, 'head'))).toBe(
      "Head, Green Lens → best for your band: Halo of Transcendence Onyxia's Lair: Onyxia · 17% drop +18.4 pts",
    );
    expect(text(rowHtml(body, 'back'))).toContain('Dire Maul: Tendris Warpwood · <1% drop');
    expect(text(rowHtml(body, 'neck'))).toBe(
      'Neck, Empty → best for your band: Choker of the Fire Lord Molten Core: Ragnaros · 12% drop +12.6 pts',
    );
  });

  it('an unscored gap reads "Not scored", muted and right-aligned (ruling 8)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const legs = rowHtml(body, 'legs');
    expect(text(legs)).toMatch(/Not scored$/);
    expect(classesOf(legs, 'character-gear-gap')).toEqual(
      expect.arrayContaining(['text-muted', 'text-right']),
    );
  });

  it('the off hand under the worn staff: "· comes with the main-hand swap", the gap blank (ruling 9)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const off = rowHtml(body, 'off_hand');
    expect(text(off)).toBe(
      'Off hand, Empty → best for your band: Spirit of Aquementas Blackrock Spire: Gyth · 9% drop · comes with the main-hand swap',
    );
    expect(/data-testid="character-gear-gap">([^<]*)</.exec(off)?.[1]).toBe('');
  });

  it('the five largest gaps in 700, the rest in 400, in the same green (ruling 6)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    for (const slot of ['head', 'neck', 'back', 'chest', 'waist'])
      expect(classesOf(rowHtml(body, slot), 'character-gear-gap'), slot).toEqual(
        expect.arrayContaining(['text-kill', 'font-bold']),
      );
    for (const slot of ['wrist', 'main_hand']) {
      const gap = classesOf(rowHtml(body, slot), 'character-gear-gap');
      expect(gap, slot).toEqual(expect.arrayContaining(['text-kill', 'font-normal']));
      expect(gap, slot).not.toContain('font-bold');
    }
  });

  it('phone: line 1 the label and the gap, line 2 the items, line 3 the source; 44px rows (ruling 4)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    expect(/<li class="([^"]*)"/.exec(head)?.[1].split(' ')).toEqual(
      expect.arrayContaining(['grid', 'min-h-11', 'grid-cols-[minmax(0,1fr)_auto]']),
    );
    expect(classesOf(head, 'character-gear-slot')).toEqual(
      expect.arrayContaining(['col-start-1', 'row-start-1', 'text-muted', 'text-[13px]', 'leading-5']),
    );
    expect(classesOf(head, 'character-gear-gap')).toEqual(
      expect.arrayContaining(['col-start-2', 'row-start-1', 'text-right', 'font-mono', 'leading-5']),
    );
    expect(classesOf(head, 'character-gear-line')).toEqual(
      expect.arrayContaining(['col-span-2', 'row-start-2', 'flex', 'flex-wrap']),
    );
    expect(classesOf(head, 'character-gear-meta')).toEqual(
      expect.arrayContaining(['basis-full', 'text-muted', 'text-[13px]']),
    );
    // A best-for-band row keeps "Best for your band" on line 2, after a dot drawn only there.
    const shoulder = rowHtml(body, 'shoulder');
    expect(classesOf(shoulder, 'character-gear-meta')).not.toContain('basis-full');
    expect(shoulder).toMatch(/<span class="lg:hidden" aria-hidden="true">· <\/span>/);
  });

  it('desktop: slot 104 | worn → pick | source 300 | gap 88, a row tint on hover and focus (ruling 5)', () => {
    const { body } = render(CharacterGearPanel, { props: PROPS });
    const head = rowHtml(body, 'head');
    expect(/<li class="([^"]*)"/.exec(head)?.[1].split(' ')).toEqual(
      expect.arrayContaining([
        'lg:grid-cols-[104px_minmax(0,1fr)_300px_88px]',
        'hover:bg-[rgba(255,255,255,0.04)]',
        'focus-within:bg-[rgba(255,255,255,0.04)]',
      ]),
    );
    expect(classesOf(head, 'character-gear-line')).toContain('lg:contents');
    expect(classesOf(head, 'character-gear-items')).toEqual(
      expect.arrayContaining(['lg:col-start-2', 'lg:row-start-1']),
    );
    expect(classesOf(head, 'character-gear-meta')).toEqual(
      expect.arrayContaining(['lg:col-start-3', 'lg:row-start-1']),
    );
    expect(classesOf(head, 'character-gear-gap')).toEqual(
      expect.arrayContaining(['lg:col-start-4', 'lg:row-start-1']),
    );
  });

  it('`gear: []`: the "Gear" heading over one line that keeps the anchor (ruling 11)', () => {
    const { body } = render(CharacterGearPanel, { props: { ...PROPS, gear: [], total: null } });
    expect(body).not.toContain('character-gear-row-');
    expect(body).not.toContain('id="gear" data-testid="character-gear"');
    const heading = /<h2[^>]*tabindex="-1"[^>]*data-testid="character-gear-title"[^>]*>([\s\S]*?)<\/h2>/.exec(
      body,
    );
    expect(text(heading?.[1] ?? '')).toBe('Gear');
    const line = /<p class="([^"]*)" id="gear" data-testid="character-gear-none">([\s\S]*?)<\/p>/.exec(body);
    expect(text(line?.[2] ?? '')).toBe(gearCopy.noGear);
    expect(body.indexOf('character-gear-title')).toBeLessThan(body.indexOf('character-gear-none'));
  });
});
