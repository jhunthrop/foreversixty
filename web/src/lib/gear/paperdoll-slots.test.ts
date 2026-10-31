// web/src/lib/gear/paperdoll-slots.test.ts
import { describe, expect, it } from 'vitest';
import type { Item } from '../planner/types';
import type { GearItemTable } from './qualities';
import type { CharacterGearSlot } from './types';
import { slotsFromCharacterGear, slotsFromWorn } from './paperdoll-slots';

const helm: Item = {
  id: 12640,
  name: 'Lionheart Helm',
  icon: 'inv_helmet_36',
  slot: 'head',
  quality: 4,
  required_level: 56,
  item_level: 61,
  armor: 565,
  stats: {},
  set_id: null,
  unique: false,
};
const champion: Item = { ...helm, id: 18877, name: 'Arcanite Champion', slot: 'main_hand', two_hand: true };
const lookup = (id: number): Item | undefined => [helm, champion].find((item) => item.id === id);

describe('slotsFromWorn (sim, report)', () => {
  it('draws a known item from the items file, an empty slot and an unknown id', () => {
    const slots = slotsFromWorn({ head: 12640, waist: 288887 }, { lookup });
    expect(slots.find((slot) => slot.slot === 'head')).toEqual({
      slot: 'head',
      state: 'item',
      item: {
        id: 12640,
        name: 'Lionheart Helm',
        icon: 'inv_helmet_36',
        quality: 4,
        itemLevel: 61,
        twoHand: false,
      },
    });
    expect(slots.find((slot) => slot.slot === 'waist')).toEqual({
      slot: 'waist',
      state: 'unknown',
      unknownId: 288887,
    });
    expect(slots.find((slot) => slot.slot === 'neck')).toEqual({ slot: 'neck', state: 'empty' });
    expect(slots).toHaveLength(17);
  });

  it('carries two-handedness for the off hand rule', () => {
    const slots = slotsFromWorn({ main_hand: 18877 }, { lookup });
    expect(slots.find((slot) => slot.slot === 'main_hand')?.item?.twoHand).toBe(true);
  });

  it('names an item outside the class file by its name, with no colour, icon or level', () => {
    const slots = slotsFromWorn({ finger1: 5 }, { lookup, outsideName: () => 'Keepsake Ring' });
    expect(slots.find((slot) => slot.slot === 'finger1')?.item).toEqual({
      id: 5,
      name: 'Keepsake Ring',
      icon: null,
      quality: null,
      itemLevel: null,
      twoHand: false,
    });
  });

  it('with a failed items file draws every worn id as a plain item, never as unknown', () => {
    const slots = slotsFromWorn({ head: 12640 }, { lookup: () => undefined, failed: true });
    expect(slots.find((slot) => slot.slot === 'head')).toMatchObject({
      state: 'item',
      item: { id: 12640, name: 'Item 12640', icon: null, quality: null },
    });
  });

  it('takes the log’s item level and detail line, and shirt and tabard when worn', () => {
    const slots = slotsFromWorn(
      { head: 12640, shirt: 6 },
      { lookup, itemLevel: () => 66, detail: (slot) => (slot === 'head' ? 'Item level 66' : undefined) },
    );
    expect(slots.find((slot) => slot.slot === 'head')).toMatchObject({
      item: { itemLevel: 66 },
      detail: 'Item level 66',
    });
    expect(slots.some((slot) => slot.slot === 'shirt')).toBe(true);
    expect(slots.some((slot) => slot.slot === 'tabard')).toBe(false);
  });
});

describe('slotsFromCharacterGear (character page)', () => {
  const table: GearItemTable = new Map([
    [11925, { quality: 3, icon: 'inv_helmet_41', twoHand: false }],
    [12640, { quality: 4, icon: 'inv_helmet_36', twoHand: false }],
  ]);
  const gear: CharacterGearSlot[] = [
    {
      slot: 'head',
      worn: { id: 11925, name: 'Ghostshroud', item_level: 62 },
      pick: { id: 12640, name: 'Lionheart Helm', source: 'Crafted', chance: null },
      gap: 60,
    },
    {
      slot: 'waist',
      worn: { id: 288887, name: null, item_level: null },
      pick: { id: 12640, name: 'Lionheart Helm', source: 'Crafted', chance: null },
      gap: null,
    },
    { slot: 'neck', worn: null, pick: null, gap: null },
  ];

  it('draws the worn item with the file’s colour and icon, the API’s name and level, and its upgrade', () => {
    const [head] = slotsFromCharacterGear(gear, table);
    expect(head).toEqual({
      slot: 'head',
      state: 'item',
      item: {
        id: 11925,
        name: 'Ghostshroud',
        icon: 'inv_helmet_41',
        quality: 3,
        itemLevel: 62,
        twoHand: false,
      },
      upgrade: {
        kind: 'gain',
        points: 60,
        pick: { id: 12640, name: 'Lionheart Helm', icon: 'inv_helmet_36', quality: 4 },
      },
    });
  });

  it('draws an item the data lacks as unknown with a "Not scored" pick, and an empty slot as empty', () => {
    const slots = slotsFromCharacterGear(gear, table);
    expect(slots[1]).toMatchObject({
      slot: 'waist',
      state: 'unknown',
      unknownId: 288887,
      upgrade: { kind: 'unscored' },
    });
    expect(slots[2]).toEqual({ slot: 'neck', state: 'empty' });
  });

  it('leaves colour and icon out while the file is missing', () => {
    const [head] = slotsFromCharacterGear(gear, null);
    expect(head?.item).toMatchObject({ icon: null, quality: null });
  });
});
