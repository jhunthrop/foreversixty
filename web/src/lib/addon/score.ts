// web/src/lib/addon/score.ts
// The weighted score the gear panel shows, from data/builds/<build>/stat-weights.json.
//
// The same arithmetic the addon's Gear.lua does, on the same numbers, so a player who
// sorts by score here and then opens the in-game panel sees the two agree.
import { SPECS } from '../sim/specs';
import type { Item, StatKey } from '../planner/types';

/** Contract 10.8 stat names to weights. */
export type SpecWeights = Readonly<Record<string, number>>;

export interface WeightsEntry {
  spec: string;
  weights: SpecWeights;
  sources: { label: string; url: string; kind: string }[];
}

export type WeightsFile = readonly WeightsEntry[];

/**
 * The planner's own stat keys to contract 10.8's, for the ones that differ.
 * `data/pipeline/normalize/gear.py` writes `healing` and `fire_res`; the engine's `Stat`
 * enum -- which is what `Data.lua`'s weights are keyed by and what the addon scores
 * against -- spells them out. One table, imported by `build-code.ts` too: two copies of
 * it is two things to keep right.
 */
export const CONTRACT_STAT_NAMES: Readonly<Record<string, string>> = {
  healing: 'healing_power',
  fire_res: 'fire_resistance',
  frost_res: 'frost_resistance',
  nature_res: 'nature_resistance',
  shadow_res: 'shadow_resistance',
  arcane_res: 'arcane_resistance',
};

export function scoreItem(item: Item, weights: SpecWeights): number {
  let total = 0;
  // item.stats is Partial<Record<StatKey, number>>, so a value read back is
  // `number | undefined` even for a key Object.keys just handed us; skip it rather than
  // let a `NaN` sneak into the total.
  for (const key of Object.keys(item.stats) as StatKey[]) {
    const value = item.stats[key];
    if (value === undefined) continue;
    total += value * (weights[CONTRACT_STAT_NAMES[key] ?? key] ?? 0);
  }
  // Armour is a column on the record, not a stat, but it is a stat to every tank weight
  // list; leaving it out would score a shield at nothing.
  total += item.armor * (weights.armor ?? 0);
  return total;
}

/**
 * The spec key for a build: the tree with the most points, ties to the first tree.
 * The design's own rule, and the same one `Gear.specOf` applies in the addon.
 */
export function specKeyFor(classSlug: string, pointsPerTree: readonly number[]): string {
  const classSpecs = SPECS.filter((spec) => spec.class_slug === classSlug).sort(
    (a, b) => a.tree_index - b.tree_index,
  );
  let bestIndex = 0;
  for (let index = 1; index < classSpecs.length; index += 1) {
    if ((pointsPerTree[index] ?? 0) > (pointsPerTree[bestIndex] ?? 0)) bestIndex = index;
  }
  return classSpecs[bestIndex]?.spec ?? '';
}

export function weightsFor(file: WeightsFile, specKey: string): WeightsEntry | undefined {
  return file.find((entry) => entry.spec === specKey);
}
