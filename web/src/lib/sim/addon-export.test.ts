// web/src/lib/sim/addon-export.test.ts
import { describe, expect, it } from 'vitest';
import { addonGearList, addonStringFor, gearEntry } from './addon-export';
import type { GearSlot } from './types';

const gear: GearSlot[] = [
  { slot: 'head', item_id: 16963, enchant: 2543 },
  { slot: 'main_hand', item_id: 12784 },
  { slot: 'trinket1', item_id: 13968, enchant: 0, suffix: 2 },
];

describe('gearEntry', () => {
  it('writes item, enchant and suffix only as far as it has to', () => {
    expect(gearEntry({ slot: 'head', item_id: 1 })).toBe('head=1');
    expect(gearEntry({ slot: 'head', item_id: 1, enchant: 2 })).toBe('head=1:2');
    expect(gearEntry({ slot: 'head', item_id: 1, suffix: 3 })).toBe('head=1:0:3');
    expect(gearEntry({ slot: 'head', item_id: 1, enchant: 2, suffix: 3 })).toBe('head=1:2:3');
  });
});

describe('addonGearList', () => {
  it('joins the slots with commas in the order given', () => {
    expect(addonGearList(gear)).toBe('head=16963:2543,main_hand=12784,trinket1=13968:0:2');
  });
});

describe('addonStringFor', () => {
  it('is an FS1 string the existing decoder reads unchanged', () => {
    const code = addonStringFor({
      dataBuild: '1.60.1',
      classSlug: 'warrior',
      raceSlug: 'orc',
      talents: '0-5530515-',
      gear,
    });
    expect(code.startsWith('FS1:1.60.1:warrior:orc:0/5530515/0:')).toBe(true);
    expect(code.endsWith(addonGearList(gear))).toBe(true);
  });
});
