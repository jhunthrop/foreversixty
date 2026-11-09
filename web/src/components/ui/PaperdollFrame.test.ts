import { createRawSnippet } from 'svelte';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { DollCell } from '../../lib/ui/paperdoll';
import { SEAL } from '../../lib/ui/assets';
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

  it('puts the title, its meta and the figure with its unit alone on the one title bar', () => {
    const body = doll({ heading: 'h1' });
    expect(body).toMatch(
      /<h1 class="pdf-title[^"]*"[^>]*>[\s\S]*Beast Mastery Hunter[\s\S]*20 to 29 · Alliance[\s\S]*91\.0 <small[^>]*>DPS<\/small>/,
    );
  });

  it('holds the seal in the medallion with no portrait, the portrait when given', () => {
    expect(doll()).toContain(`src="${SEAL.painted128}"`);
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
});
