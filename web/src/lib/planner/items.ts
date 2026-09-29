// web/src/lib/planner/items.ts
// Item list helpers for the gear panel. design/DESIGN-SYSTEM.md: item-rarity colours are
// the WoW standard and are never repurposed, so they come from the rarity tokens in
// styles/tokens.css rather than from a hex in a component.
import { slotsForItem } from './rules';
import type { Item, Slot } from './types';

const RARITY_CLASS: Record<number, string> = {
  0: 'text-rarity-poor',
  1: 'text-rarity-common',
  2: 'text-rarity-uncommon',
  3: 'text-rarity-rare-text',
  4: 'text-rarity-epic-text',
  5: 'text-rarity-legendary',
};

/** The raw (unlightened) rarity tokens -- tokens.css's own comment names these "for bars,
 *  borders and icons" as distinct from the lightened `-text` variants `RARITY_CLASS` above
 *  uses for on-screen text. A quality-coloured item icon border (the in-game socket look,
 *  the BiS paperdoll's own gear row and alternative chips) reads from this map instead. */
const RARITY_BORDER_VAR: Record<number, string> = {
  0: 'var(--color-rarity-poor)',
  1: 'var(--color-rarity-common)',
  2: 'var(--color-rarity-uncommon)',
  3: 'var(--color-rarity-rare)',
  4: 'var(--color-rarity-epic)',
  5: 'var(--color-rarity-legendary)',
};

/**
 * The line a worn slot shows: the item's name when the planner files carry it; the name
 * from the build's itemnames.json with "not simmed" when they do not (a keepsake ring, a
 * totem); an honest "unknown item" only when no data of ours knows the id at all.
 */
export function wornItemLabel(
  itemName: string | undefined,
  equippedId: number | undefined,
  outsideNames: Readonly<Record<string, string>>,
  copy: { itemNotSimmed: (name: string) => string; unknownItem: (id: number) => string },
): string {
  if (itemName !== undefined) return itemName;
  if (equippedId === undefined) return 'Empty';
  const outside = outsideNames[String(equippedId)];
  return outside === undefined ? copy.unknownItem(equippedId) : copy.itemNotSimmed(outside);
}

export function rarityClassFor(quality: number): string {
  return RARITY_CLASS[quality] ?? RARITY_CLASS[1];
}

/** A CSS `border-color` value for `quality` -- an inline-style pair for `rarityClassFor`
 *  wherever a border, not text, needs to carry the item's rarity (see `RARITY_BORDER_VAR`'s
 *  own doc). Falls back to common the same way `rarityClassFor` does. */
export function rarityBorderColorFor(quality: number): string {
  return RARITY_BORDER_VAR[quality] ?? RARITY_BORDER_VAR[1];
}

/** The items that fit a slot, sorted the way the spec asks: required level, then name. */
export function itemsForSlot(items: Item[], slot: Slot): Item[] {
  return items
    .filter((item) => slotsForItem(item).includes(slot))
    .sort((a, b) => a.required_level - b.required_level || a.name.localeCompare(b.name));
}

export function searchItems(items: Item[], query: string): Item[] {
  const needle = query.trim().toLowerCase();
  if (needle.length === 0) return items;
  return items.filter((item) => item.name.toLowerCase().includes(needle));
}
