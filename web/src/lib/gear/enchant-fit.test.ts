// web/src/lib/gear/enchant-fit.test.ts
import { describe, expect, it } from 'vitest';
import { enchantFits, enchantRuleOf, enchantTargetOf } from './enchant-fit';

const WEAPON = {
  equipped_item_class: 2,
  equipped_item_subclass_mask: 0xe5f3,
  equipped_inventory_type_mask: 0,
};
const HELD_IN_OFF_HAND = { class_id: 4, subclass_id: 0, inventory_type: 23 };

describe('enchantFits', () => {
  it('follows the client: class, then the subclass and inventory-type masks', () => {
    expect(enchantFits(WEAPON, { class_id: 2, subclass_id: 15, inventory_type: 13 })).toBe(true);
    expect(enchantFits(WEAPON, HELD_IN_OFF_HAND)).toBe(false);
    expect(
      enchantFits(
        { equipped_item_class: -1, equipped_item_subclass_mask: 0, equipped_inventory_type_mask: 0 },
        HELD_IN_OFF_HAND,
      ),
    ).toBe(true);
  });
});

describe('enchantTargetOf and enchantRuleOf', () => {
  it('read a row only when it carries all three columns', () => {
    expect(enchantTargetOf(HELD_IN_OFF_HAND)).toEqual(HELD_IN_OFF_HAND);
    expect(enchantTargetOf({ class_id: 4, subclass_id: null, inventory_type: 23 })).toBeUndefined();
    expect(enchantRuleOf(WEAPON)).toEqual(WEAPON);
    expect(enchantRuleOf({ equipped_item_class: 2 })).toBeUndefined();
  });
});
