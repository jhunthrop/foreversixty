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

export function rarityClassFor(quality: number): string {
  return RARITY_CLASS[quality] ?? RARITY_CLASS[1];
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
