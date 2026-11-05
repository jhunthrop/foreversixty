// web/src/lib/items/slot-words.ts
// An item's slot and type in the client's own tooltip words and casing (`Head`, `Leather`,
// `One-Hand`, `Axe`, `Held In Off-hand`), from the client's Item.InventoryType, ClassID and
// SubclassID. The API sends those three on a loot row; its raw `slot` enum is logic only and is
// never printed. A label that prints upper case does it in CSS, so the words stay the client's.

const ITEM_CLASS_WEAPON = 2;
const ITEM_CLASS_ARMOR = 4;
/** A cloak is cloth to the client, but its tooltip says only `Back`. */
const INVENTORY_TYPE_BACK = 16;

/** Item.InventoryType -> the tooltip's slot word. */
const SLOT_WORDS: Readonly<Record<number, string>> = {
  1: 'Head',
  2: 'Neck',
  3: 'Shoulder',
  4: 'Shirt',
  5: 'Chest',
  6: 'Waist',
  7: 'Legs',
  8: 'Feet',
  9: 'Wrist',
  10: 'Hands',
  11: 'Finger',
  12: 'Trinket',
  13: 'One-Hand',
  14: 'Off Hand',
  15: 'Ranged',
  16: 'Back',
  17: 'Two-Hand',
  19: 'Tabard',
  20: 'Chest',
  21: 'Main Hand',
  22: 'Off Hand',
  23: 'Held In Off-hand',
  25: 'Thrown',
  26: 'Ranged',
  28: 'Relic',
};

/** Weapon Item.SubclassID -> the tooltip's type word. */
const WEAPON_WORDS: Readonly<Record<number, string>> = {
  0: 'Axe',
  1: 'Axe',
  2: 'Bow',
  3: 'Gun',
  4: 'Mace',
  5: 'Mace',
  6: 'Polearm',
  7: 'Sword',
  8: 'Sword',
  10: 'Staff',
  13: 'Fist Weapon',
  15: 'Dagger',
  16: 'Thrown',
  18: 'Crossbow',
  19: 'Wand',
};

/** Armour Item.SubclassID -> the tooltip's type word. Miscellaneous (0) prints none. */
const ARMOR_WORDS: Readonly<Record<number, string>> = {
  1: 'Cloth',
  2: 'Leather',
  3: 'Mail',
  4: 'Plate',
  6: 'Shield',
  7: 'Libram',
  8: 'Idol',
  9: 'Totem',
};

function typeWord(inventoryType: number, itemClass: number, itemSubclass: number): string | undefined {
  if (itemClass === ITEM_CLASS_WEAPON) return WEAPON_WORDS[itemSubclass];
  if (itemClass === ITEM_CLASS_ARMOR && inventoryType !== INVENTORY_TYPE_BACK) {
    return ARMOR_WORDS[itemSubclass];
  }
  return undefined;
}

/** `Head · Leather`, `One-Hand · Axe`, `Trinket`; empty for an item with no slot. */
export function slotLine(inventoryType: number, itemClass: number, itemSubclass: number): string {
  const slot = SLOT_WORDS[inventoryType];
  if (slot === undefined) return '';
  const type = typeWord(inventoryType, itemClass, itemSubclass);
  return type === undefined ? slot : `${slot} · ${type}`;
}
