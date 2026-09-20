// web/src/lib/sim/enchants.ts
// data/builds/<build>/enchants.json and suffixes.json.
//
// Contract 10.4 describes enchants as keyed by "effect_id" with `item_types` as a numeric
// `EnchantType` shape restriction and `phase` as a phase-table name. The real, shipped file
// at data/builds/1.60.1.69893/enchants.json disagrees with that prose on every one of those
// points, and the fixtures this task tests against (Task 2) were built from the real file,
// so this module follows the real file:
//   - rows are keyed by `id` (holding the effect id's value), not `effect_id`;
//   - `item_types` is a string enum (`normal` / `kit` / `shield` / `two_hand`), not numbers;
//   - `phase` is the client's own small integer content phase (0/2/4/5/6 observed), which
//     has nothing to do with this site's pre-beta/beta/launch/raids-1 phase table;
//   - `spell_id` is always present and > 0 across all 173 real rows; `item_id` is the one
//     that is sometimes 0. Both are always present, never optional.
//   - a stat row's keys are not limited to the planner's `StatKey` set. Enumerated by
//     reading every distinct key across BOTH real files
//     (data/builds/1.60.1.69893/enchants.json and suffixes.json): the real files also carry
//     `arcane_power`, `block_value`, `bonus_armor`, `fire_power`, `frost_power`, `health`,
//     `holy_power`, `mana`, `melee_haste`, `nature_power`, `shadow_power` and
//     `spell_damage` -- 12 keys outside `StatKey`. (`ranged_attack_power` was a
//     thirteenth until the data lane added it to `STAT_KEYS`; it is inside `StatKey`
//     now, so naming it again here would be redundant.) `EnchantStatKey`
//     below is `StatKey` unioned with exactly those 12, so both `EnchantRow.stats` and
//     `SuffixRow.stats` keep the same typo protection as every other stat map in the
//     codebase while still typing the real data exactly, rather than widening to
//     `Record<string, number>` and losing that protection altogether.
// `suffixes.json` matched contract 6.3's shape exactly, so `SuffixRow` is unchanged beyond
// the `stats` key widening above.
//
// Both files are optional: a build the data lane has not regenerated ships neither, and
// the enchant/suffix columns simply do not appear. Neither is large enough to want pruning,
// so both are fetched whole, per build, and cached by the browser the way items.json is.
//
// These are the browser UI's copy only. The planner inside the wasm does not read them:
// contract A9 embeds the same rows in `sim/internal/simdb`, so `Expand` needs no file at
// runtime and the two can never disagree about what fits where.
import { dataUrl, loadOptional } from '../planner/load';
import type { Item, StatKey } from '../planner/types';

/**
 * `StatKey` plus every stat key the real enchant/suffix files carry that the planner's
 * gear-stat set does not (see the file header for how this list was derived). A closed
 * union rather than `Record<string, number>` so a typo in a consumer (`stats.melee_hast`)
 * still fails to type-check.
 */
export type EnchantStatKey =
  | StatKey
  | 'arcane_power'
  | 'block_value'
  | 'bonus_armor'
  | 'fire_power'
  | 'frost_power'
  | 'health'
  | 'holy_power'
  | 'mana'
  | 'melee_haste'
  | 'nature_power'
  | 'shadow_power'
  | 'spell_damage';

export interface EnchantRow {
  /**
   * The client's enchantment effect id. This is the key (contract 10.4, as corrected above)
   * and it is what `GearSlot.enchant` and `Candidate.enchant` carry.
   */
  id: number;
  /** The spell that applies it. Always present and > 0 in the real file. */
  spell_id: number;
  /** The item that applies it, where one does; 0 when the enchant has no source item. */
  item_id: number;
  name: string;
  icon: string;
  /** Planner slot names. */
  slots: string[];
  /** The `EnchantType` shape restriction, as the fork's own string enum values. */
  item_types: string[];
  /** Empty for an enchant every class can use. */
  classes: string[];
  stats: Partial<Record<EnchantStatKey, number>>;
  /** The client's own small content-phase number, unrelated to the site's phase table. */
  phase: number;
}

export interface SuffixRow {
  id: number;
  name: string;
  stats: Partial<Record<EnchantStatKey, number>>;
}

/**
 * `Candidate.enchant` of 0 means "inherit the equipped item's enchant for this slot where
 * it fits" (contract 1.3), so 0 is the wire's own "keep current" and there is no separate
 * value for it. -1 is this page's "explicitly none": it never travels, and `candidates.ts`
 * turns it into an omitted enchant on a candidate whose slot is empty.
 */
export const NO_ENCHANT = 0;
export const KEEP_CURRENT_ENCHANT = -1;

/** Raidbots shows a per-slot selection cap; ours is four, which is a legible list. */
export const ENCHANTS_PER_SLOT_CAP = 4;

export function loadEnchants(build: string): Promise<EnchantRow[]> {
  return loadOptional<EnchantRow[]>(dataUrl(build, 'enchants.json'), []);
}

export function loadSuffixes(build: string): Promise<SuffixRow[]> {
  return loadOptional<SuffixRow[]>(dataUrl(build, 'suffixes.json'), []);
}

/**
 * The enchants this slot allows. Two gates, both the fork's own: the enchant's slot list,
 * and its class list where it has one (an empty `classes` is for everyone).
 *
 * `item_types` is deliberately NOT applied here. It is the `EnchantType` shape restriction
 * and `items.json` rows carry no armour subclass to test it against, so the browser cannot
 * evaluate it. `Expand` inside the wasm can, from simdb, and refuses a bad pairing at the
 * boundary with its own words -- which is the honest failure, and better than a filter here
 * that would have to guess.
 */
export function enchantsForSlot(rows: readonly EnchantRow[], slot: string, classSlug = ''): EnchantRow[] {
  return rows.filter((row) => {
    if (!row.slots.includes(slot)) return false;
    if (row.classes.length > 0 && classSlug !== '' && !row.classes.includes(classSlug)) return false;
    return true;
  });
}

/** The random suffixes this item rolls, in file order. */
export function suffixesForItem(rows: readonly SuffixRow[], item: Item): SuffixRow[] {
  const ids = new Set(item.suffixes ?? []);
  return rows.filter((row) => ids.has(row.id));
}
