// web/src/lib/gear/paperdoll.test.ts
import { describe, expect, it } from 'vitest';
import type { CharacterGearSlot } from './types';
import {
  averageItemLevel,
  cellLabel,
  hyphenate,
  paperdollLayout,
  paperdollNote,
  phoneUpgradeShown,
  pickLabel,
  upgradeFor,
  type PaperdollSlot,
} from './paperdoll';

const SOFT = '­';

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

describe('paperdollLayout', () => {
  it('puts head to wrist on the left, hands to trinket 2 on the right, top-aligned, without shirt and tabard', () => {
    const cells = paperdollLayout(fullSet());
    const left = cells.filter((cell) => cell.side === 'left').map((cell) => [cell.slot, cell.row]);
    const right = cells
      .filter((cell) => cell.side === 'right' && cell.row <= 8)
      .map((cell) => [cell.slot, cell.row]);
    expect(left).toEqual([
      ['head', 1],
      ['neck', 2],
      ['shoulder', 3],
      ['back', 4],
      ['chest', 5],
      ['wrist', 6],
      ['main_hand', 9],
    ]);
    expect(right).toEqual([
      ['hands', 1],
      ['waist', 2],
      ['legs', 3],
      ['feet', 4],
      ['finger1', 5],
      ['finger2', 6],
      ['trinket1', 7],
      ['trinket2', 8],
    ]);
  });

  it('places the main hand left and the off hand right on row 9, ranged centred on row 10 (sheet parity)', () => {
    const cells = paperdollLayout(fullSet());
    const at = (slot: string) => cells.find((cell) => cell.slot === slot);
    expect(at('main_hand')).toMatchObject({ side: 'left', row: 9 });
    expect(at('off_hand')).toMatchObject({ side: 'right', row: 9 });
    expect(at('ranged')).toMatchObject({ side: 'centre', row: 10 });
  });

  it('adds shirt and tabard after chest when the data has them', () => {
    const cells = paperdollLayout([...fullSet(), known('shirt', 1), known('tabard', 2)]);
    const left = cells.filter((cell) => cell.side === 'left' && cell.row <= 8).map((cell) => cell.slot);
    expect(left).toEqual(['head', 'neck', 'shoulder', 'back', 'chest', 'shirt', 'tabard', 'wrist']);
  });

  it('orders the phone frame as pairs, each followed by its upgrade rows, left slot first', () => {
    const slots = fullSet().map((slot) =>
      slot.slot === 'head' || slot.slot === 'hands' || slot.slot === 'neck'
        ? {
            ...slot,
            upgrade: { kind: 'gain' as const, points: 4, pick: { id: 9, name: 'P', icon: null, quality: 4 } },
          }
        : slot,
    );
    const cells = paperdollLayout(slots);
    const phone = [...cells].sort((a, b) => a.phoneRow - b.phoneRow || a.order - b.order);
    const sequence = phone.slice(0, 7).map((cell) => `${cell.kind}:${cell.slot}:${cell.phoneRow}`);
    expect(sequence).toEqual([
      'slot:head:2',
      'slot:hands:2',
      'upgrade:head:3',
      'upgrade:hands:4',
      'slot:neck:5',
      'slot:waist:5',
      'upgrade:neck:6',
    ]);
  });

  it('keeps the DOM order equal to the phone order', () => {
    const cells = paperdollLayout(fullSet());
    const rows = cells.map((cell) => cell.phoneRow);
    expect(rows).toEqual([...rows].sort((a, b) => a - b));
  });

  it('draws no phone upgrade row for "Best in slot", one for a gain or "Not scored"', () => {
    expect(phoneUpgradeShown({ kind: 'best' })).toBe(false);
    expect(phoneUpgradeShown(undefined)).toBe(false);
    expect(phoneUpgradeShown({ kind: 'unscored', pick: { id: 1, name: 'P', icon: null, quality: 3 } })).toBe(
      true,
    );
  });

  it('marks an empty off hand under a two-hander, and only then', () => {
    const twoHander = fullSet().map((slot) =>
      slot.slot === 'main_hand'
        ? known('main_hand', 7, 63, true)
        : slot.slot === 'off_hand'
          ? ({ slot: 'off_hand', state: 'empty' } as const)
          : slot,
    );
    expect(paperdollLayout(twoHander).find((cell) => cell.slot === 'off_hand')?.twoHandOff).toBe(true);
    const oneHander = twoHander.map((slot) => (slot.slot === 'main_hand' ? known('main_hand', 7) : slot));
    expect(paperdollLayout(oneHander).find((cell) => cell.slot === 'off_hand')?.twoHandOff).toBe(false);
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

  it('is a gain for a different, higher-scoring pick', () => {
    expect(upgradeFor(gear({}), {}, ref)).toEqual({ kind: 'gain', points: 60, pick: lionheart });
  });

  it('is "Best in slot" when the pick is worn, never a gain, whatever the gap says', () => {
    const worn = { id: 12640, name: 'Lionheart Helm', item_level: 61 };
    expect(upgradeFor(gear({ worn, gap: 4.8 }), {}, ref)).toEqual({ kind: 'best' });
  });

  it('matches rings and trinkets either way round', () => {
    const ring = gear({
      slot: 'finger1',
      worn: { id: 1, name: 'A', item_level: 60 },
      pick: { ...pick, id: 2, name: 'B' },
      gap: 3,
    });
    expect(upgradeFor(ring, { finger2: 2 }, ref)).toEqual({ kind: 'best' });
  });

  it('is nothing for a different pick scoring the same or lower', () => {
    expect(upgradeFor(gear({ gap: 0 }), {}, ref)).toBeUndefined();
  });

  it('is the pick’s whole score on an empty slot', () => {
    expect(upgradeFor(gear({ worn: null, gap: 22 }), {}, ref)).toMatchObject({ kind: 'gain', points: 22 });
  });

  it('is "Not scored" when the worn item is unknown', () => {
    const worn = { id: 288887, name: null, item_level: null };
    expect(upgradeFor(gear({ worn, gap: null }), {}, ref)).toEqual({ kind: 'unscored', pick: lionheart });
  });

  it('is nothing when the band has no pick', () => {
    expect(upgradeFor(gear({ pick: null, gap: null }), {}, ref)).toBeUndefined();
  });
});

describe('copy', () => {
  it('names a cell, an unknown cell and a pick for a screen reader', () => {
    expect(cellLabel(known('head', 1))).toBe('Head: Item 1');
    expect(cellLabel({ slot: 'waist', state: 'unknown', unknownId: 288887 })).toBe(
      'Waist: Item 288887, not in our item data yet',
    );
    expect(cellLabel({ slot: 'finger1', state: 'empty' })).toBe('Finger 1: empty');
    const pick = { id: 1, name: 'Lionheart Helm', icon: null, quality: 4 };
    expect(pickLabel('head', { kind: 'gain', points: 60, pick })).toBe(
      'Head upgrade: Lionheart Helm, 60 points to gain',
    );
    expect(pickLabel('chest', { kind: 'unscored', pick: { ...pick, name: 'Dawn Armor' } })).toBe(
      'Chest upgrade: Dawn Armor, not scored',
    );
  });

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

describe('hyphenate', () => {
  it('puts soft hyphens inside a long word, at syllable breaks, never within 3 letters of an end', () => {
    const word = hyphenate('Electromagnetic');
    expect(word.replaceAll(SOFT, '')).toBe('Electromagnetic');
    expect(word).toBe(`Elec${SOFT}tro${SOFT}mag${SOFT}ne${SOFT}tic`);
    const parts = word.split(SOFT);
    expect(parts[0]?.length).toBeGreaterThanOrEqual(3);
    expect(parts.at(-1)?.length).toBeGreaterThanOrEqual(3);
  });

  it('leaves every word that fits the narrowest line whole, and punctuation and spacing alone', () => {
    expect(hyphenate('Lionheart Helm')).toBe('Lionheart Helm');
    expect(hyphenate('Dormant Heart of the Mountain')).toBe('Dormant Heart of the Mountain');
    expect(hyphenate('Stratholme Militia Shoulderguard')).toContain(`Shoul${SOFT}der${SOFT}guard`);
    expect(hyphenate('Band of Vigor')).toBe('Band of Vigor');
    expect(hyphenate("Warder's")).toBe("Warder's");
  });
});
