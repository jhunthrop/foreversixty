// web/src/lib/addon/build-code.ts
// A planner build plus the item data the planner already has, as the FSB1 code the addon
// reads. Everything here is client-side: the addon code is generated in the browser from
// what the share panel is already holding, so no API route and no second copy of the item
// database is involved.
import { encodeFSB1, type FSB1GearSlot, type FSB1Point } from './fsb1';
import type { TalentIndex } from '../planner/rules';
import { SLOTS, type Gear, type Item, type Slot } from '../planner/types';

/**
 * The planner's own stat keys to parity contract 10.8's, for the three that differ.
 * `data/pipeline/normalize/gear.py` writes `healing`, `fire_res` and friends; the engine's
 * `Stat` enum -- which is what `Data.lua`'s weights are keyed by and what the addon scores
 * against -- spells them out. A key not in here passes through unchanged, which is the
 * common case.
 */
const CONTRACT_STAT_NAMES: Record<string, string> = {
  healing: 'healing_power',
  fire_res: 'fire_resistance',
  frost_res: 'frost_resistance',
  nature_res: 'nature_resistance',
  shadow_res: 'shadow_resistance',
  arcane_res: 'arcane_resistance',
};

export interface BuildCodeInput {
  dataBuild: string;
  classSlug: string;
  /** The planner's point order: one talent id per point, in the order spent. */
  order: number[];
  gear: Gear;
  talents: TalentIndex;
  items: Map<number, Item>;
}

/** The planner's talent ids as the addon's 1-based tab/tier/column triples. */
function pointsOf(talents: TalentIndex, order: readonly number[]): FSB1Point[] {
  const points: FSB1Point[] = [];
  for (const id of order) {
    const talent = talents.byId.get(id);
    const tree = talents.treeOf.get(id);
    // A talent id the loaded tree does not carry is a build from another class or another
    // data build. Dropping the point is wrong and inventing a cell is worse, so the whole
    // code omits it and the addon shows a shorter build than the site does -- which the
    // build-id check on the addon side is what actually catches.
    if (talent === undefined || tree === undefined) continue;
    points.push({ tab: tree.position + 1, tier: talent.tier + 1, column: talent.column + 1 });
  }
  return points;
}

/** One item's stats in the contract's vocabulary, armour folded in when it is non-zero. */
function statsOf(item: Item): Record<string, number> {
  const stats: Record<string, number> = {};
  for (const [key, value] of Object.entries(item.stats)) {
    if (value === undefined || value === 0) continue;
    stats[CONTRACT_STAT_NAMES[key] ?? key] = value;
  }
  // Armour is its own column on the item record, not a stat, but it is a stat to the
  // engine and to every tank weight list. Zero is left out: `armor=0` in the code would
  // read as "this item states zero armour", which is different from "no armour column".
  if (item.armor > 0) stats.armor = (stats.armor ?? 0) + item.armor;
  return stats;
}

export function addonCodeFor(input: BuildCodeInput): string {
  const gear: FSB1GearSlot[] = [];
  for (const slot of SLOTS) {
    const itemId = input.gear[slot as Slot];
    if (itemId === undefined) continue;
    const item = input.items.get(itemId);
    // An id with no item is an item this class list does not carry. A bare id with no
    // stats would score as zero in the addon and make every bag item an upgrade, so the
    // slot is left out and the addon says nothing about it.
    if (item === undefined) continue;
    gear.push({ slot: slot as Slot, itemId, stats: statsOf(item) });
  }
  return encodeFSB1({
    dataBuild: input.dataBuild,
    classSlug: input.classSlug,
    order: pointsOf(input.talents, input.order),
    gear,
  });
}
