// web/src/lib/planner/derive.ts
// The values the planner shows but never stores: the tree split, the level a point was
// spent at, gear stat totals, and which set bonuses are active. The API's card renderer
// computes the same numbers from the same record, so the shapes here stay simple.
import type { TalentIndex } from './rules';
import {
  BASE_LEVEL,
  FIRST_POINT_LEVEL,
  SLOTS,
  type Gear,
  type Item,
  type ItemSet,
  type SetBonus,
  type Slot,
  type StatKey,
} from './types';

/** The character level at which the point at `index` is spent. */
export function levelForIndex(index: number): number {
  return FIRST_POINT_LEVEL + index;
}

/** The level a build has reached; 9 before the first point. */
export function levelReached(order: number[]): number {
  return BASE_LEVEL + order.length;
}

export function ranksByTalent(order: number[]): Map<number, number> {
  const ranks = new Map<number, number>();
  for (const id of order) ranks.set(id, (ranks.get(id) ?? 0) + 1);
  return ranks;
}

/** Points in each tree, left to right. */
export function pointsPerTree(index: TalentIndex, order: number[]): number[] {
  const counts = new Map<number, number>(index.trees.map((tree) => [tree.id, 0]));
  for (const id of order) {
    const tree = index.treeOf.get(id);
    if (tree) counts.set(tree.id, (counts.get(tree.id) ?? 0) + 1);
  }
  return index.trees.map((tree) => counts.get(tree.id) ?? 0);
}

/** The split as the site writes it everywhere: `31/0/20`. */
export function splitLabel(index: TalentIndex, order: number[]): string {
  return pointsPerTree(index, order).join('/');
}

/** Equipped items in slot-grid order, skipping empty and unknown ids. */
export function equippedItems(items: Map<number, Item>, gear: Gear): Item[] {
  const equipped: Item[] = [];
  for (const slot of SLOTS as readonly Slot[]) {
    const id = gear[slot];
    if (id === undefined) continue;
    const item = items.get(id);
    if (item) equipped.push(item);
  }
  return equipped;
}

/** Summed stats; `armor` from the item's own armor field lands under the `armor` stat key. */
export function statTotals(equipped: Item[]): Partial<Record<StatKey, number>> {
  const totals: Partial<Record<StatKey, number>> = {};
  const add = (key: StatKey, value: number): void => {
    if (value === 0) return;
    totals[key] = (totals[key] ?? 0) + value;
  };
  for (const item of equipped) {
    add('armor', item.armor);
    for (const [key, value] of Object.entries(item.stats) as [StatKey, number][]) add(key, value);
  }
  return totals;
}

export interface ActiveSet {
  set: ItemSet;
  /** How many pieces of the set are equipped. */
  pieces: number;
  /** The bonuses those pieces unlock. */
  active: SetBonus[];
}

/** Every set with at least one piece equipped, with the bonuses that piece count unlocks. */
export function activeSetBonuses(equipped: Item[], sets: ItemSet[]): ActiveSet[] {
  const worn = new Set(equipped.map((item) => item.id));
  const result: ActiveSet[] = [];
  for (const set of sets) {
    const pieces = set.item_ids.filter((id) => worn.has(id)).length;
    if (pieces === 0) continue;
    result.push({ set, pieces, active: set.bonuses.filter((bonus) => bonus.pieces <= pieces) });
  }
  return result;
}
