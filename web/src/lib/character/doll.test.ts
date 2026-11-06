// The character page's paperdoll cells (design/specs/2026-11-06-character.md C4).
import { describe, expect, it } from 'vitest';
import type { PaperdollSlot } from '../gear/paperdoll';
import type { GearItemTable } from '../gear/qualities';
import type { CharacterGearSlot } from '../gear/types';
import { pickCells, wornCells } from './doll';

const slots: PaperdollSlot[] = [
  {
    slot: 'head',
    state: 'item',
    item: { id: 5, name: 'Circlet', icon: 'inv_helm', quality: 3, itemLevel: 62, twoHand: false },
    upgrade: {
      kind: 'gain',
      value: 19.2,
      unit: 'hps',
      pick: { id: 6, name: 'Crown', icon: null, quality: 4 },
    },
  },
];

describe('wornCells', () => {
  it('keeps the gain for the owner and drops it for a visitor', () => {
    expect(wornCells(slots, 'priest', true)[0].gain?.value).toBe('+19.2');
    const visitor = wornCells(slots, 'priest', false)[0];
    expect(visitor.gain).toBeUndefined();
    expect(visitor.badge).toBeUndefined();
    expect(visitor.source).toBe('Item level 62');
  });
});

describe('pickCells', () => {
  it('draws the band pick with its source for every slot that has one', () => {
    const gear: CharacterGearSlot[] = [
      {
        slot: 'head',
        worn: null,
        pick: { id: 6, name: 'Crown', source: 'Scholomance', chance: null },
        gap: 1,
      },
      { slot: 'neck', worn: null, pick: null, gap: null },
    ];
    const table: GearItemTable = new Map([[6, { quality: 4, icon: 'inv_crown', twoHand: false }]]);
    expect(pickCells(gear, table, 'priest')).toEqual([
      {
        slot: 'head',
        label: 'Head',
        item: { id: 6, name: 'Crown', quality: 4, icon: 'inv_crown' },
        source: 'Scholomance',
      },
    ]);
  });
});
