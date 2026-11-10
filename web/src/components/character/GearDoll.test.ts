// web/src/components/character/GearDoll.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { PaperdollSlot } from '../../lib/gear/paperdoll';
import GearDoll from './GearDoll.svelte';

const HELM = {
  id: 12640,
  name: 'Lionheart Helm',
  icon: 'inv_helmet_36',
  quality: 4,
  itemLevel: 61,
  twoHand: false,
};
const SLOTS: PaperdollSlot[] = [
  { slot: 'head', state: 'item', item: HELM, enchant: 'Stamina +10', detail: 'Item level 66' },
  { slot: 'neck', state: 'empty' },
];

function doll(props: { mode: 'worn' | 'report' } & Record<string, unknown>): string {
  return render(GearDoll, {
    props: {
      slots: SLOTS,
      build: '1.60.1.70790',
      classSlug: 'warrior',
      title: 'Thandrel',
      label: 'Gear',
      testid: 'doll',
      ...props,
    },
  }).body.replace(/<!--[^>]*-->/g, '');
}

describe('GearDoll', () => {
  it('draws a report set on the frame, opening on the first item with its log item level and enchant', () => {
    const body = doll({ mode: 'report', meta: 'Warrior', figure: { value: '182.4', unit: 'DPS' } });
    expect(body).toContain('data-testid="paperdoll-frame-slot-head"');
    expect(body).toMatch(/class="pdf-centre[\s\S]*Lionheart Helm[\s\S]*Item level 66 · Stamina \+10/);
    expect(body).toContain('182.4 <small');
  });

  it('lists a worn set with its enchant in green and never draws the frame', () => {
    const body = doll({ mode: 'worn', form: 'list', head: false });
    expect(body).not.toContain('paperdoll-frame-slot-head');
    expect(body).not.toContain('class="pdf-strip');
    expect(body).toMatch(
      /Head · <\/span><span[^>]*>Item level 66<\/span><\/span>\s*<span class="pdf-sub pdf-enchant[^"]*"[^>]*>Stamina \+10/,
    );
    expect(body).not.toContain('data-inline-tip');
  });

  it('waits empty while the items file loads', () => {
    expect(doll({ mode: 'worn', form: 'list', loading: true })).not.toContain('Lionheart Helm');
  });
});
