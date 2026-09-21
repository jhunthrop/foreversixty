// web/src/lib/sim/legendary-variants.ts
// dps D26: Top Gear's item search can return more than one item id under one exact name.
// Atiesh, Greatstaff of the Guardian ships as four ids, one per class whose own quest chain
// grants it (Mage, Warlock, Priest, Druid), and every class with the Staves weapon skill --
// which includes classes that could never obtain any of the four -- sees all four
// identical-looking rows with nothing to tell them apart.
//
// Nothing in this build's own item data says which class a variant belongs to: there is no
// AllowableClass column on a legendary staff, and Atiesh's stats are hand-carried rather
// than sourced from a column at all (data/pipeline/simdb/equip.py's own header names it as
// one of the eighteen items whose on-equip amount exists in no column the pipeline reads).
// This is the one case item-search.ts's own "not a hard-coded id list unless nothing else
// identifies them" rule allows -- nothing else does, and the four ids are Blizzard's own,
// stable since Classic.
export const LEGENDARY_VARIANT_CLASS: Readonly<Record<number, string>> = {
  22589: 'Mage',
  22630: 'Warlock',
  22631: 'Priest',
  22632: 'Druid',
};

/** The class label for a legendary-variant item id, or "" for every other item. */
export function legendaryVariantLabel(itemId: number): string {
  return LEGENDARY_VARIANT_CLASS[itemId] ?? '';
}
