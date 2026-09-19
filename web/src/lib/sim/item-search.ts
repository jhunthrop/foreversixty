// web/src/lib/sim/item-search.ts
// The item search behind Top Gear's "add a candidate", design 3.1.3. Entirely client-side
// over the build's own items/<class>.json, which the page already has loaded for the gear
// grid -- no API call, no second index. `searchItems` runs on every keystroke
// (`ItemSearch.svelte`), so it does one linear pass per call rather than building or
// maintaining a persistent index: the class item lists are a few hundred rows at most, and
// there is nothing here worth caching across calls.
import { slotsForItem } from '../planner/rules';
import { SLOTS, SLOT_LABELS, type Item, type Slot } from '../planner/types';

export interface ItemQuery {
  text: string;
  minItemLevel: number;
  /** A planner slot, or "" for any. */
  slot: string;
  /** A loot source id, or "" for anywhere. */
  sourceId: string;
  /** On by default, per the design. */
  usableOnly: boolean;
}

export interface SearchContext {
  /** The character's level, for the usable-only gate. */
  level: number;
  /** From `loot.ts`'s `sourcesByItem`. Empty when the build ships no loot file. */
  sourcesByItem: Map<number, string[]>;
}

/**
 * The list is class-filtered already (items/<class>.json is per class) and a long list is
 * a scroll nobody reads, so the search shows the best hundred and the caller can say how
 * many it hid.
 */
export const SEARCH_LIMIT = 100;

export function defaultItemQuery(): ItemQuery {
  return { text: '', minItemLevel: 0, slot: '', sourceId: '', usableOnly: true };
}

/** The slot filter's options, in the grid's own order. */
export function slotOptions(): { slot: Slot; label: string }[] {
  return SLOTS.map((slot) => ({ slot, label: SLOT_LABELS[slot] }));
}

/**
 * Best item level first, then name, so the top of the list is the part worth reading. The
 * design's own note stands: search can find items this character has no way to obtain, and
 * `usableOnly` is about what can be equipped, not about what can be got.
 */
export function searchItems(items: readonly Item[], query: ItemQuery, ctx: SearchContext): Item[] {
  const needle = query.text.trim().toLowerCase();
  return items
    .filter((item) => {
      if (needle !== '' && !item.name.toLowerCase().includes(needle)) return false;
      if (item.item_level < query.minItemLevel) return false;
      if (query.slot !== '' && !slotsForItem(item).includes(query.slot as Slot)) return false;
      if (query.sourceId !== '' && !(ctx.sourcesByItem.get(item.id) ?? []).includes(query.sourceId)) {
        return false;
      }
      if (query.usableOnly && item.required_level > ctx.level) return false;
      return true;
    })
    .sort((a, b) => b.item_level - a.item_level || a.name.localeCompare(b.name))
    .slice(0, SEARCH_LIMIT);
}
