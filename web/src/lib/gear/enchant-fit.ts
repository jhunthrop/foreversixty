// web/src/lib/gear/enchant-fit.ts
// The client's rule for which item an enchant goes on, shared by the guide's enchant table
// (lib/guides/gear-extras.ts) and Top Gear's enchant picker (lib/sim/enchants.ts). An enchant
// row carries its spell's SpellEquippedItems class, subclass mask and inventory-type mask
// (`enchants.json`, pipeline loot/gear.py); an item carries its ItemClass, ItemSubClass and
// InventoryType (`items.json`, `items/<class>.json`). A held-in-off-hand item (class 4,
// inventory type 23) takes no weapon enchant (class 2); a two-hand enchant needs a
// two-handed subclass; a shield enchant needs the shield subclass.

/** What an enchant's spell says it can go on. */
export interface EnchantRule {
  /** ItemClass the item must be; -1 for any. */
  equipped_item_class: number;
  /** ItemSubClass bits the item's subclass must be among; 0 for any. */
  equipped_item_subclass_mask: number;
  /** InventoryType bits the item's inventory type must be among; 0 for any. */
  equipped_inventory_type_mask: number;
}

/** What the client checks an item by when an enchant is applied to it. */
export interface EnchantTarget {
  class_id: number;
  subclass_id: number;
  inventory_type: number;
}

/** SpellEquippedItems' "any item class". */
const ANY_ITEM_CLASS = -1;

function maskAllows(mask: number, bit: number): boolean {
  return mask === 0 || (mask & (1 << bit)) !== 0;
}

/** The client's rule: the item's class is the enchant's, and its subclass and inventory type
 *  are among the ones the enchant's masks name (an empty mask names every one). */
export function enchantFits(rule: EnchantRule, item: EnchantTarget): boolean {
  return (
    (rule.equipped_item_class === ANY_ITEM_CLASS || rule.equipped_item_class === item.class_id) &&
    maskAllows(rule.equipped_item_subclass_mask, item.subclass_id) &&
    maskAllows(rule.equipped_inventory_type_mask, item.inventory_type)
  );
}

/** The item's shape when it carries all three columns; `undefined` for a row from a build
 *  written before they existed, which no rule can be applied to. */
export function enchantTargetOf(item: {
  class_id?: number | null;
  subclass_id?: number | null;
  inventory_type?: number | null;
}): EnchantTarget | undefined {
  const { class_id: classId, subclass_id: subclassId, inventory_type: inventoryType } = item;
  if (classId == null || subclassId == null || inventoryType == null) return undefined;
  return { class_id: classId, subclass_id: subclassId, inventory_type: inventoryType };
}

/** The enchant's rule when its row carries all three columns; `undefined` for a row from an
 *  `enchants.json` written before they existed. */
export function enchantRuleOf(row: Partial<EnchantRule>): EnchantRule | undefined {
  const {
    equipped_item_class: itemClass,
    equipped_item_subclass_mask: subclassMask,
    equipped_inventory_type_mask: inventoryTypeMask,
  } = row;
  if (itemClass === undefined || subclassMask === undefined || inventoryTypeMask === undefined)
    return undefined;
  return {
    equipped_item_class: itemClass,
    equipped_item_subclass_mask: subclassMask,
    equipped_inventory_type_mask: inventoryTypeMask,
  };
}
