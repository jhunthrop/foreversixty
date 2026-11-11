// web/src/lib/bis/source-filter.ts
// The BiS page's source filters (owner MUST 2026-11-05): filtering is not hiding. With a kind
// off, a slot shows the best pick from the kinds still on. This is the twin of the ranker's
// own resolver (`sim/cmd/leveling-bis/source_filter.go`'s `resolveFilteredSlots`): the pick
// first whenever one of its kinds is on, else the lowest-ranked `by_kind` entry among the kinds
// on. A pair's second slot (finger2, trinket2, off_hand) also steers clear of whatever the first
// slot resolved to (`not_with`, falling back to `next`), and no off hand sits beside a
// two-handed main hand. Multi-source rule: an item belongs to every kind it can be had from.
import type { BisByKindEntry, BisSlot, SourceFilterKind } from './types';
import { PAIR_OF } from '../gear/slots';

/** The filter row's order. */
export const SOURCE_FILTER_KINDS: readonly SourceFilterKind[] = [
  'dungeon',
  'raid',
  'world',
  'quest',
  'crafted',
  'rep',
  'vendor',
  'pvp',
];

/** loot.json kind -> filter kind; the Go table `sourceKindToFilterKind` is the same. */
const FILTER_KIND_OF_SOURCE: Readonly<Record<string, SourceFilterKind>> = {
  dungeon: 'dungeon',
  raid: 'raid',
  world: 'world',
  world_drop: 'world',
  zone: 'world',
  quest: 'quest',
  crafted: 'crafted',
  rep: 'rep',
  vendor: 'vendor',
  pvp: 'pvp',
};

export function filterKindOfSource(sourceKind: string): SourceFilterKind | undefined {
  return FILTER_KIND_OF_SOURCE[sourceKind];
}

/** The second slot of each pair (the rings, the trinkets, the off hand), mapped to the first. */
const PAIR_FIRST: Readonly<Record<string, string>> = {
  .../* @__PURE__ */ Object.fromEntries(Object.entries(PAIR_OF).filter(([slot]) => slot.endsWith('2'))),
  off_hand: 'main_hand',
};

/** The fields of a slot the filter reads. */
export type FilterableSlot = Pick<
  BisSlot,
  'slot' | 'item_id' | 'source_kind' | 'kinds' | 'two_hand' | 'not_with' | 'by_kind'
>;

/** A slot's answer under a filter. `source`/`source_kind` are absent when it is the pick. */
export interface FilteredPick {
  item_id: number;
  two_hand: boolean;
  source?: string;
  source_kind?: string;
  /** The gap against the unfiltered pick (0 for the pick itself). */
  dps_delta: number;
}

/** What the first slot of a pair resolved to, for the second. */
export interface PairMateChoice {
  item_id: number;
  two_hand: boolean;
}

/** The pick's own filter kinds: `kinds`, else its `source_kind`'s. */
export function pickKinds(slot: Pick<BisSlot, 'kinds' | 'source_kind'>): SourceFilterKind[] {
  if (slot.kinds !== undefined && slot.kinds.length > 0) return slot.kinds;
  const kind = filterKindOfSource(slot.source_kind ?? '');
  return kind === undefined ? [] : [kind];
}

function usable(entry: BisByKindEntry | undefined, mateId: number | undefined): BisByKindEntry | undefined {
  if (entry === undefined) return undefined;
  if (mateId !== undefined && entry.not_with?.includes(mateId)) {
    const next = entry.next;
    return next !== undefined && !next.not_with?.includes(mateId) ? next : undefined;
  }
  return entry;
}

/**
 * One slot's pick with only `allowedKinds` on; `null` for "No <slot> from these sources at this
 * level". `mate` is the first slot's own filtered pick when `slot` is a pair's second slot
 * (`null` when that slot resolved to nothing).
 */
export function filteredPick(
  slot: FilterableSlot,
  allowedKinds: ReadonlySet<SourceFilterKind>,
  mate?: PairMateChoice | null,
): FilteredPick | null {
  if (slot.slot === 'off_hand' && mate?.two_hand === true) return null;
  const mateId = mate?.item_id;
  const hasPick = typeof slot.item_id === 'number' && slot.item_id !== 0;
  const pickBlocked = mateId !== undefined && (slot.not_with?.includes(mateId) ?? false);
  if (hasPick && !pickBlocked && pickKinds(slot).some((kind) => allowedKinds.has(kind))) {
    return { item_id: slot.item_id, two_hand: slot.two_hand === true, dps_delta: 0 };
  }
  let best: { entry: BisByKindEntry; kind: SourceFilterKind } | undefined;
  for (const kind of SOURCE_FILTER_KINDS) {
    if (!allowedKinds.has(kind)) continue;
    const entry = usable(slot.by_kind?.[kind], mateId);
    if (entry !== undefined && (best === undefined || entry.rank < best.entry.rank)) best = { entry, kind };
  }
  if (best === undefined) return null;
  return {
    item_id: best.entry.item_id,
    two_hand: best.entry.two_hand === true,
    source: best.entry.source,
    source_kind: best.entry.source_kind ?? best.kind,
    dps_delta: best.entry.dps_delta ?? 0,
  };
}

/** Every slot of a band under a filter, pairs resolved in order (first slot, then second). */
export function filteredBand(
  slots: readonly FilterableSlot[],
  allowedKinds: ReadonlySet<SourceFilterKind>,
): Map<string, FilteredPick | null> {
  const out = new Map<string, FilteredPick | null>();
  for (const slot of slots) {
    const first = PAIR_FIRST[slot.slot];
    const mate = first === undefined ? undefined : (out.get(first) ?? null);
    out.set(slot.slot, filteredPick(slot, allowedKinds, mate));
  }
  return out;
}
