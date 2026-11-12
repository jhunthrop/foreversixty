import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { DollCell } from '../../lib/ui/paperdoll';
import { SEAL } from '../../lib/ui/seal';
import PaperdollFrame from './PaperdollFrame.svelte';

const BUILD = '1.60.1.70790';
const HARPOON = { id: 7, name: 'Impaling Harpoon', quality: 3, icon: 'inv_spear_04' };
const CELLS: DollCell[] = [
  { slot: 'head', label: 'Head', item: { id: 1, name: "Brawler's Leather Hood", quality: 3, icon: 'a' } },
  { slot: 'back', label: 'Back', item: { id: 2, name: 'Glowing Lizardscale Cloak', quality: 2, icon: 'b' } },
  { slot: 'trinket1', label: 'Trinket', item: null },
  { slot: 'main_hand', label: 'Two-Hand', item: HARPOON },
  { slot: 'off_hand', label: 'Off Hand', item: null, coveredBy: HARPOON },
];

function doll(props: Record<string, unknown> = {}): string {
  return render(PaperdollFrame, {
    props: {
      cells: CELLS,
      build: BUILD,
      title: 'Beast Mastery Hunter',
      meta: '20 to 29 · Alliance',
      figure: { value: '91.0', unit: 'DPS' },
      ...props,
    },
  }).body.replace(/<!--[^>]*-->/g, '');
}

describe('PaperdollFrame', () => {
  it('positions a button per filled slot from the measured boxes, with its accessible name', () => {
    const body = doll();
    expect(body).toMatch(/data-testid="paperdoll-frame-slot-head"[^>]*|style="left:[\d.]+%;top:[\d.]+%/);
    expect(body).toContain('aria-label="Head: Brawler\'s Leather Hood"');
    expect(body).toContain('aria-label="Back: Glowing Lizardscale Cloak"');
  });

  it('draws nothing over an empty slot, so the frame glyph shows', () => {
    expect(doll()).not.toContain('paperdoll-frame-slot-trinket1');
  });

  it('shows a two-hander greyed in the off hand', () => {
    const body = doll();
    expect(body).toContain('aria-label="Off Hand: Impaling Harpoon, two-handed"');
    expect(body).toMatch(/data-testid="paperdoll-frame-slot-off_hand"[\s\S]*?pdf-covered/);
  });

  it('borders each icon in the client quality colour', () => {
    expect(doll()).toContain('--qb: var(--color-rarity-rare)');
  });

  it('puts the title (the heading alone), its meta and the figure with its unit on the one title bar', () => {
    const body = doll({ heading: 'h1' });
    expect(body).toMatch(
      /class="pdf-title[^"]*"[^>]*>\s*<h1 class="pdf-title-text[^"]*">Beast Mastery Hunter<\/h1>[\s\S]*20 to 29 · Alliance[\s\S]*91\.0 <small[^>]*>DPS<\/small>/,
    );
  });

  it('stands the character render in the centre only when given', () => {
    expect(doll()).not.toContain('pdf-model');
    const body = doll({ render: '/fixtures/render-placeholder.png' });
    expect(body).toMatch(/class="pdf-centre[^"]*pdf-modelled/);
    expect(body).toContain('src="/fixtures/render-placeholder.png"');
  });

  it('marks itself as showing its tooltips in place only when it has a centre to show them in', () => {
    const centre = createRawSnippet(() => ({ render: () => '<p>centre</p>' }));
    expect(doll({ centre })).toContain('data-inline-tip');
    expect(doll()).not.toContain('data-inline-tip');
  });

  it('holds the seal in the medallion with no portrait, the portrait when given', () => {
    expect(doll()).toContain(`src="${SEAL.disc128}"`);
    expect(doll({ avatar: 'https://example.test/a.jpg' })).toContain('src="https://example.test/a.jpg"');
  });

  it('renders the caller centre for the selected slot', () => {
    const centre = createRawSnippet((cell: () => DollCell | undefined) => ({
      render: () => `<p>centre:${cell()?.item?.name ?? 'none'}</p>`,
    }));
    const body = doll({ selected: 'back', centre });
    expect(body).toContain('centre:Glowing Lizardscale Cloak');
    expect(body).toContain('data-region="paperdoll-centre"');
    expect(body).toContain('data-region="paperdoll-medallion"');
  });

  it('carries the list form for narrow widths, without the shirt and tabard', () => {
    const body = doll();
    expect(body).toMatch(/class="pdf-list[ "]/);
    expect(body).toContain('>Neck<');
    expect(body).toContain('data-testid="paperdoll-list-slot-head"');
    expect(body).not.toContain('paperdoll-list-slot-shirt');
  });

  it('draws only the list form, never the frame image, with form="list"', () => {
    const body = doll({ form: 'list' });
    expect(body).not.toContain('class="pdf-frame');
    expect(body).not.toContain('paperdoll-frame-slot-head');
    expect(body).toContain('pdf-listonly');
    expect(body).toContain('data-testid="paperdoll-list-slot-head"');
  });

  it('drops the strip with head={false}', () => {
    expect(doll()).toContain('class="pdf-strip');
    const body = doll({ form: 'list', head: false });
    expect(body).not.toContain('class="pdf-strip');
    expect(body).toMatch(/class="pdf-items[^"]*pdf-headless/);
  });

  it('gives each row its line 2 and gain with facts: the enchant in green, the pick icon and figure', () => {
    const cells: DollCell[] = [
      {
        slot: 'head',
        label: 'Head',
        item: { id: 1, name: 'Lionheart Helm', quality: 4, icon: 'a' },
        source: 'Item level 61',
        enchant: 'Stamina +10',
        gain: { value: '+1.8', unit: 'DPS', tone: 'gold', basis: { prefix: 'vs', item: HARPOON } },
      },
      { slot: 'neck', label: 'Neck', item: { id: 2, name: 'Mark of Fordring', quality: 3, icon: 'b' } },
    ];
    const body = doll({ cells, form: 'list', facts: true });
    expect(body).toMatch(
      /Head · <\/span><span[^>]*>Item level 61<\/span><\/span>\s*<span class="pdf-sub pdf-enchant[^"]*"[^>]*>Stamina \+10<\/span>/,
    );
    expect(body).toMatch(
      /data-testid="paperdoll-list-gain-head"[\s\S]*data-icon="inv_spear_04"[\s\S]*\+1\.8 DPS/,
    );
    expect(body).not.toContain('paperdoll-list-gain-neck');
    expect(doll({ cells, form: 'list' })).not.toContain('Stamina +10');
  });

  it('keeps the centre to the frame with listCentre={false}: the list opens the page tooltip', () => {
    const centre = createRawSnippet(() => ({ render: () => '<p>centre</p>' }));
    const body = doll({ selected: 'head', centre, listCentre: false });
    expect(body).toMatch(/class="pdf-frame[^"]*"[^>]*data-inline-tip/);
    expect(body).not.toMatch(/class="pdf[^-][^>]*data-inline-tip/);
    expect(body).not.toContain('pdf-list-centre');
  });

  it('draws a headless list straight on the page panel', () => {
    expect(doll({ form: 'list', head: false })).toMatch(/class="pdf-items[^"]*pdf-headless/);
  });
});
