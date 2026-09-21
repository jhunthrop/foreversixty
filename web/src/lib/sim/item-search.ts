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
  /**
   * From `sim-items.ts`'s `knownItemIds`: the ids the engine's embedded database carries.
   * `undefined` (every existing caller before this field existed) and `null` (a build that
   * ships no `simitems.json`) both mean "nothing to filter against" -- search results
   * still show every item this class can equip. Unlike `usableOnly`, this is never a
   * player-facing toggle: an item the engine does not know about cannot be simulated no
   * matter what the character could otherwise wear, so search simply never offers it
   * (`CandidateRows.svelte`'s disabled row is for candidates that reach the grid some
   * other way -- bags, bank, a Droptimizer pick, an old pinned link -- where the item is
   * real and hiding it would look like data loss; a fresh search has no such row yet).
   */
  known?: ReadonlySet<number> | null;
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
 * A build's item file carries the client's own internal QA fixtures alongside real items
 * (dps D26): Slot = Main hand, no name filter, used to list `90 Epic Frost Staff`,
 * `90 Epic Rogue Dagger` and four more before any real weapon, on a level-60 client where
 * item level tops out around 92 -- but a ceiling check alone cannot tell those apart from
 * Atiesh, Greatstaff of the Guardian (also item level 90, and real). What every one of
 * these 687 fixture rows across every class's item file shares, and no real item ever
 * does, is a name that opens with its own bare item level: "90 Epic Frost Staff", "63 Green
 * Rogue Dagger". A real item name is never a number -- that is the one rule this needs, and
 * it is exact (verified against all nine class files on build 1.60.1.69893: 687 matches,
 * zero false positives against any shipped item), so there is no reason to also stack an
 * item-level ceiling or a no-source rule on top of it and risk hiding a real, simply
 * not-yet-sourced item (loot.json's own header: the re-itemised raid tier is thin by
 * design, not by bug).
 */
const DEV_FIXTURE_NAME = /^\d/;

export function isDevFixtureItem(item: Item): boolean {
  return DEV_FIXTURE_NAME.test(item.name);
}

/**
 * The one predicate `searchItems` and `matchCount` both filter on, so the two never drift
 * apart into two different definitions of "matches".
 */
function matchesQuery(item: Item, query: ItemQuery, ctx: SearchContext): boolean {
  if (isDevFixtureItem(item)) return false;
  const needle = query.text.trim().toLowerCase();
  if (needle !== '' && !item.name.toLowerCase().includes(needle)) return false;
  if (item.item_level < query.minItemLevel) return false;
  if (query.slot !== '' && !slotsForItem(item).includes(query.slot as Slot)) return false;
  if (query.sourceId !== '' && !(ctx.sourcesByItem.get(item.id) ?? []).includes(query.sourceId)) {
    return false;
  }
  if (query.usableOnly && item.required_level > ctx.level) return false;
  const known = ctx.known ?? null;
  if (known !== null && !known.has(item.id)) return false;
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

export interface NoResultsReason {
  /** Relaxing only the slot filter finds at least one match: the name is real, it is just
   *  not equippable in the slot currently filtered. */
  existsElsewhere: boolean;
  /** Every match `existsElsewhere` found is two-handed -- the specific, nameable reason an
   *  off-hand search finds nothing for it (dps D27: Ashkandi is two-handed, so excluding it
   *  from an off-hand search is correct; the page should say why instead of reading like
   *  the item does not exist). */
  twoHanded: boolean;
}

/**
 * Why a text search under a slot filter came back empty, when it is worth saying more than
 * `bulkCopy.searchNoResults` -- checked only once a name is typed and a slot is chosen
 * (an empty query or "any slot" already explains itself). `searchItems`' own slot filter
 * (`slotsForItem`) never lists `off_hand` for a two-handed weapon, so re-running the same
 * query with the slot relaxed is the one call this needs: it reuses `matchesQuery`'s own
 * rules (fixtures still excluded, `usableOnly` unchanged) rather than a second definition
 * of "matches" that could drift from `searchItems`'.
 */
export function noResultsReason(
  items: readonly Item[],
  query: ItemQuery,
  ctx: SearchContext,
): NoResultsReason {
  if (query.text.trim() === '' || query.slot === '') return { existsElsewhere: false, twoHanded: false };
  const elsewhere = searchItems(items, { ...query, slot: '' }, ctx);
  return {
    existsElsewhere: elsewhere.length > 0,
    twoHanded: elsewhere.length > 0 && elsewhere.every((match) => match.two_hand === true),
  };
}
