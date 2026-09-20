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
 * The one predicate `searchItems` and `matchCount` both filter on, so the two never drift
 * apart into two different definitions of "matches".
 */
function matchesQuery(item: Item, query: ItemQuery, ctx: SearchContext): boolean {
  const needle = query.text.trim().toLowerCase();
  if (needle !== '' && !item.name.toLowerCase().includes(needle)) return false;
  if (item.item_level < query.minItemLevel) return false;
  if (query.slot !== '' && !slotsForItem(item).includes(query.slot as Slot)) return false;
  if (query.sourceId !== '' && !(ctx.sourcesByItem.get(item.id) ?? []).includes(query.sourceId)) {
    return false;
  }
  if (query.usableOnly && item.required_level > ctx.level) return false;
  return true;
}

/**
 * Best item level first, then name, so the top of the list is the part worth reading. The
 * design's own note stands: search can find items this character has no way to obtain, and
 * `usableOnly` is about what can be equipped, not about what can be got.
 *
 * Deliberately re-scans `items` on every call rather than building or caching an index:
 * this runs against one class's `items.json` (a few hundred rows at most, called once per
 * keystroke), and a full filter/sort/slice pass over that is cheap enough that an index
 * would add invalidation logic for no measurable benefit. Revisit this if the item table a
 * caller passes ever grows to thousands of rows (e.g. an all-classes search), where a
 * per-keystroke linear scan would start to matter.
 */
export function searchItems(items: readonly Item[], query: ItemQuery, ctx: SearchContext): Item[] {
  return items
    .filter((item) => matchesQuery(item, query, ctx))
    .sort((a, b) => b.item_level - a.item_level || a.name.localeCompare(b.name))
    .slice(0, SEARCH_LIMIT);
}

/**
 * How many items match `query` before the `SEARCH_LIMIT` slice -- what `ItemSearch.svelte`
 * needs to say "showing 100 of 143" rather than silently dropping the rest with no notice.
 */
export function matchCount(items: readonly Item[], query: ItemQuery, ctx: SearchContext): number {
  return items.filter((item) => matchesQuery(item, query, ctx)).length;
}
