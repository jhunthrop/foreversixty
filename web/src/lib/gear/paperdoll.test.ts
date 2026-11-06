// web/src/lib/gear/paperdoll.test.ts
import { describe, expect, it } from 'vitest';
import type { CharacterGearSlot } from './types';
import {
  averageItemLevel,
  paperdollNote,
  twoHandCover,
  upgradeFor,
  type GainPricer,
  type PaperdollSlot,
} from './paperdoll';

function known(slot: PaperdollSlot['slot'], id: number, itemLevel = 60, twoHand = false): PaperdollSlot {
  return {
    slot,
    state: 'item',
    item: { id, name: `Item ${id}`, icon: 'inv_x', quality: 3, itemLevel, twoHand },
  };
}

const FULL_KEYS = [
  'head',
  'neck',
  'shoulder',
  'back',
  'chest',
  'wrist',
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
  'main_hand',
  'off_hand',
  'ranged',
] as const;

function fullSet(): PaperdollSlot[] {
  return FULL_KEYS.map((slot, index) => known(slot, 100 + index, 60 + (index % 4)));
}

describe('twoHandCover', () => {
  it('is the two-hander over an empty off hand, and only then', () => {
    const twoHander = fullSet().map((slot) =>
      slot.slot === 'main_hand'
        ? known('main_hand', 7, 63, true)
        : slot.slot === 'off_hand'
          ? ({ slot: 'off_hand', state: 'empty' } as const)
          : slot,
    );
    // The game's character pane shows the two-hander greyed in the off-hand slot.
    expect(twoHandCover(twoHander)).toMatchObject({ id: 7, icon: 'inv_x' });
    const oneHander = twoHander.map((slot) => (slot.slot === 'main_hand' ? known('main_hand', 7) : slot));
    expect(twoHandCover(oneHander)).toBeNull();
    expect(twoHandCover(fullSet())).toBeNull();
  });
});

describe('averageItemLevel (§3.5)', () => {
  it('is the mean of every equipped slot, one decimal', () => {
    expect(averageItemLevel(fullSet())).toBe('61.4');
  });

  it('is absent when any equipped slot is not in our item data', () => {
    const slots = fullSet().map((slot) =>
      slot.slot === 'waist' ? ({ slot: 'waist', state: 'unknown', unknownId: 288887 } as const) : slot,
    );
    expect(averageItemLevel(slots)).toBeNull();
  });

  it('does not count an empty slot, nor the empty off hand under a two-hander', () => {
    const slots: PaperdollSlot[] = [
      known('head', 1, 60),
      known('main_hand', 2, 63, true),
      { slot: 'off_hand', state: 'empty' },
      { slot: 'ranged', state: 'empty' },
    ];
    expect(averageItemLevel(slots)).toBe('61.5');
  });

  it('does not count shirt or tabard', () => {
    expect(averageItemLevel([known('head', 1, 60), known('shirt', 2, 1), known('tabard', 3, 20)])).toBe(
      '60.0',
    );
  });

  it('is absent with nothing equipped, or while an item level is unknown', () => {
    expect(averageItemLevel([{ slot: 'head', state: 'empty' }])).toBeNull();
    const noLevel: PaperdollSlot = {
      slot: 'head',
      state: 'item',
      item: { id: 1, name: 'A', icon: null, quality: null, itemLevel: null, twoHand: false },
    };
    expect(averageItemLevel([noLevel])).toBeNull();
  });
});

describe('upgradeFor (§3.4)', () => {
  const pick = { id: 12640, name: 'Lionheart Helm', source: 'Crafted', chance: null };
  const gear = (overrides: Partial<CharacterGearSlot>): CharacterGearSlot => ({
    slot: 'head',
    worn: { id: 11925, name: 'Ghostshroud', item_level: 62 },
    pick,
    gap: 60,
    ...overrides,
  });

  const ref = (id: number, name: string) => ({ id, name, icon: null, quality: null });
  const lionheart = ref(12640, 'Lionheart Helm');
  const priced: GainPricer = { unit: 'dps', price: (_slot, _pick, worn) => (worn === undefined ? 4.4 : 1.8) };
  const unpriced: GainPricer = { unit: 'hps', price: () => null };

  it('is a gain for a different, higher-scoring pick, priced in the band’s unit', () => {
    expect(upgradeFor(gear({}), {}, ref, priced)).toEqual({
      kind: 'gain',
      value: 1.8,
      unit: 'dps',
      pick: lionheart,
    });
  });

  it('keeps a gain the band cannot price as a gain without a number', () => {
    expect(upgradeFor(gear({}), {}, ref, unpriced)).toEqual({
      kind: 'gain',
      value: null,
      unit: 'hps',
      pick: lionheart,
    });
  });

  it('is "Best in slot" when the pick is worn, never a gain, whatever the gap says', () => {
    const worn = { id: 12640, name: 'Lionheart Helm', item_level: 61 };
    expect(upgradeFor(gear({ worn, gap: 4.8 }), {}, ref, priced)).toEqual({ kind: 'best' });
  });

  it('matches rings and trinkets either way round', () => {
    const ring = gear({
      slot: 'finger1',
      worn: { id: 1, name: 'A', item_level: 60 },
      pick: { ...pick, id: 2, name: 'B' },
      gap: 3,
    });
    expect(upgradeFor(ring, { finger2: 2 }, ref, priced)).toEqual({ kind: 'best' });
  });

  it('is nothing for a different pick scoring the same or lower', () => {
    expect(upgradeFor(gear({ gap: 0 }), {}, ref, priced)).toBeUndefined();
  });

  it('prices an empty slot against nothing worn', () => {
    expect(upgradeFor(gear({ worn: null, gap: 22 }), {}, ref, priced)).toMatchObject({
      kind: 'gain',
      value: 4.4,
    });
  });

  it('is "Not scored" when the worn item is unknown', () => {
    const worn = { id: 288887, name: null, item_level: null };
    expect(upgradeFor(gear({ worn, gap: null }), {}, ref, priced)).toEqual({
      kind: 'unscored',
      pick: lionheart,
    });
  });

  it('is nothing when the band has no pick', () => {
    expect(upgradeFor(gear({ pick: null, gap: null }), {}, ref, priced)).toBeUndefined();
  });
});

describe('copy', () => {
  it('writes the note once: one unknown, several, or a failed file', () => {
    expect(paperdollNote(0, false)).toBeNull();
    expect(paperdollNote(1, false)).toBe(
      '1 item is not in our item data yet. It is new in the beta, and it gets a name and icon when the client lists it. Item level shows once every item is known.',
    );
    expect(paperdollNote(3, false)).toBe(
      '3 items are not in our item data yet. They are new in the beta, and they get names and icons when the client lists them. Item level shows once every item is known.',
    );
    expect(paperdollNote(2, true)).toBe('Item icons and colours did not load. Reload the page to try again.');
  });
});
