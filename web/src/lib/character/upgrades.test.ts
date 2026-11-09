// The character paperdoll's Upgrades list (directions mock b2-character, the frame's centre).
import { describe, expect, it } from 'vitest';
import type { PaperdollSlot } from '../gear/paperdoll';
import { UPGRADES_SHOWN, upgradeRows } from './upgrades';

function gain(slot: PaperdollSlot['slot'], value: number | null, pickId: number): PaperdollSlot {
  return {
    slot,
    state: 'item',
    item: { id: pickId - 100, name: 'Worn', icon: 'inv_worn', quality: 2, itemLevel: 50, twoHand: false },
    upgrade: {
      kind: 'gain',
      value,
      unit: 'hps',
      pick: { id: pickId, name: `Pick ${pickId}`, icon: pickId === 201 ? null : 'inv_pick', quality: 4 },
    },
  };
}

describe('upgradeRows', () => {
  it('lists the band picks by their priced gain, largest first, the slot named', () => {
    const rows = upgradeRows(
      [gain('head', 18.4, 200), gain('legs', 25.7, 202), gain('feet', 16.1, 203)],
      'priest',
    );
    expect(rows.map((row) => [row.slot, row.label, row.item.name, row.value])).toEqual([
      ['legs', 'Legs', 'Pick 202', 25.7],
      ['head', 'Head', 'Pick 200', 18.4],
      ['feet', 'Feet', 'Pick 203', 16.1],
    ]);
    expect(rows[0].item).toEqual({ id: 202, name: 'Pick 202', quality: 4, icon: 'inv_pick' });
  });

  it('leaves out unpriced gains, ties, best-in-slot and empty slots, and a missing icon', () => {
    const slots: PaperdollSlot[] = [
      gain('head', null, 200),
      gain('neck', 0.01, 201),
      { slot: 'back', state: 'item', upgrade: { kind: 'best' } },
      { slot: 'chest', state: 'empty' },
      gain('wrist', 3, 201),
    ];
    const rows = upgradeRows(slots, 'priest');
    expect(rows.map((row) => row.slot)).toEqual(['wrist']);
    expect(rows[0].item.icon).toBeUndefined();
  });

  it(`shows ${UPGRADES_SHOWN} at most`, () => {
    const many = (['head', 'neck', 'legs', 'feet', 'wrist', 'waist'] as const).map((slot, index) =>
      gain(slot, index + 1, 300 + index),
    );
    expect(upgradeRows(many, 'priest')).toHaveLength(UPGRADES_SHOWN);
  });
});
