// web/src/lib/home/upgrades.ts
// Home rebuild spec §3.B.2/§3.B.3: compares a signed-in character's actually-worn gear
// (`MeCharacter.build.gear`, the API lane's own extension to the contract) against a
// leveling-BiS band's picks (`data/builds/<build>/bis/<spec>.json`), slot by slot. Pure and
// synchronous -- every async concern (which band, which items file, which BiS file) is
// `lib/home/upgrades-loader.ts`'s job, so this stays trivially unit-testable against hand-
// built fixtures.
//
// `MeCharacter.build.gear`'s own slot vocabulary (`addon/ForeverSixty/Export.lua`'s
// `INVENTORY_SLOTS`: head, neck, shoulder, back, chest, wrist, hands, waist, legs, feet,
// finger1, finger2, trinket1, trinket2, main_hand, off_hand, ranged) is byte-for-byte
// `planner/types.ts`'s own `SLOTS`/`Slot` -- the same list a BiS band's own `slots[].slot`
// values use too. No slot-name mapping is needed anywhere in this module; that identity is
// the "map to the BiS json slot vocabulary" the lane brief asked to confirm, written here
// once rather than as a no-op translation table.
import type { MeCharacter } from '../account/api';
import { scoreItem, type SpecWeights } from '../addon/score';
import { hasKnownSource } from '../bis/source-cell';
import { slotScoreUnitFor, type SlotScoreUnit } from '../bis/tank-view';
import type { BisBand, BisSlot } from '../bis/types';
import { SLOTS, type Item, type Slot } from '../planner/types';

/** The slots a weapon can occupy -- `scoreItem` only sums `item.stats`, so it cannot
 *  reproduce a weapon's own damage-based score term (its published `score` bakes that in,
 *  `scoreItem` does not), which is exactly the ~13x overstatement fix round 1 caught on the
 *  Ranger Bow row. A weapon slot's gain is only ever read off a sim-verified
 *  `BisAlternative.dps_delta` (`alternativeGainDps`); never `scoreItem`-diffed. */
const WEAPON_SLOTS: ReadonlySet<string> = new Set(['main_hand', 'off_hand', 'ranged']);

/** One slot where the character's worn item differs from the band's own pick (and is not a
 *  scored tie with it -- see `tiesWithPick`). */
export interface SlotUpgrade {
  slot: Slot;
  /** The band's own recommended pick for this slot -- always a row `hasKnownSource`. */
  pick: BisSlot;
  /** The item id `build.gear` names for this slot; absent when nothing is equipped there. */
  wornItemId: number | undefined;
  /** The worn item's name, when this build's item file knows the id; absent either way
   *  nothing is equipped or the id is one this build's item table does not carry. */
  wornItemName: string | undefined;
  /** True when `wornItemId` is set but this build's item file has no row for it -- an addon
   *  export naming an id the site cannot score (a data gap, not an empty slot). */
  wornUnknown: boolean;
  /**
   * The pick's own advantage over the worn item, in real DPS, computed one of two ways (fix
   * round 1: never mix them against each other):
   *
   *   1. The worn item is a listed `alternatives` entry: `-alt.dps_delta`, the ranker's own
   *      sim-verified gap (`BisAlternative.dps_delta` is the pick's advantage over the
   *      alternative, so negating it gives the alternative's -- the worn item's -- own
   *      advantage from switching to the pick). A listed alternative with `dps_delta === 0`
   *      is a tie, handled by `tiesWithPick` before this is ever computed.
   *   2. Otherwise, for a non-weapon slot, `(pick.score - scoreItem(worn, weights)) *
   *      reference_dps_per_point` -- the band's own `reference_stat_points` score delta,
   *      converted to DPS the same way the BiS page's weight rail documents
   *      (`BisStatWeight.dps_per_point`'s own doc).
   *
   * `null` (never a fabricated number) when neither path applies: a weapon slot
   * (`notSimChecked`) with no alternative match, the band carries no
   * `reference_dps_per_point`, the pick carries no `score` at all (a sim-decided row), or the
   * worn item's id is not in this build's item table (`wornUnknown`).
   */
  gainDps: number | null;
  /** True when `gainDps` is null specifically because this is a weapon slot with no
   *  sim-verified alternative to compare against -- `scoreItem` is never used for a weapon,
   *  so this is a real "we don't have a verified number," not a data gap elsewhere. Renders
   *  as the BiS page's own muted "not sim-checked" tag (`bisCopy.notSimCheckedTag`) and is
   *  excluded from `totalGainDps`/the card line's own DPS sum. */
  notSimChecked: boolean;
}

export interface UpgradesResult {
  /** Every upgrade slot, sorted by `gainDps` descending (an unknown gain -- `null` --
   *  sorts last, after every slot with a real, measured number). */
  upgrades: SlotUpgrade[];
  /** Every slot where the character already wears the band's own pick (or a scored tie with
   *  it), for the "Already best in slot" footer line -- never claims a slot the band itself
   *  has no known source for. */
  alreadyBis: { slot: Slot; itemId: number; itemName: string }[];
  /** The sum of every known (non-null) `gainDps` across `upgrades` -- never includes a
   *  `notSimChecked` row. */
  totalGainDps: number;
  /** The unit every gain here is measured in: real DPS for a damage band, the ranker's
   *  tank score for a tank band. The card line and each row print it, so a tank never
   *  reads a score as DPS. */
  scoreUnit: SlotScoreUnit;
  /** How many `upgrades` are `notSimChecked` -- the card line's own "N slots not sim-checked"
   *  clause reads this count rather than re-deriving it. */
  notSimCheckedCount: number;
}

/** `band.weights` (an array of `{stat, weight}` rows) as the `Record<string, number>`
 *  `scoreItem` (`lib/addon/score.ts`) takes -- one place this reshaping happens, so a BiS
 *  band's own weight list and the addon panel's `stat-weights.json` shape can both feed the
 *  identical scorer without this module caring which one it was given. */
export function weightsRecordFor(band: Pick<BisBand, 'weights'>): SpecWeights {
  return Object.fromEntries(band.weights.map((weight) => [weight.stat, weight.weight]));
}

/** The worn item's own listed-alternative row, when `pick.alternatives` names it -- `undefined`
 *  for an item the pick's own runner-up list does not carry at all (most worn items: the
 *  ranker only publishes a handful of runners-up per slot). */
function alternativeFor(pick: BisSlot, wornItemId: number | undefined) {
  if (wornItemId === undefined) return undefined;
  return (pick.alternatives ?? []).find((alt) => alt.item_id === wornItemId);
}

/**
 * True when `wornItemId` is the band's own pick, or a listed alternative with `dps_delta ===
 * 0` -- a real, measured tie the ranker's own verify pass already found, not merely "scores
 * the same" by this module's own cheaper stat-weight arithmetic (`BisAlternative.dps_delta`'s
 * own doc: 0 exactly for an alternative identical in value to the pick).
 */
function tiesWithPick(pick: BisSlot, wornItemId: number | undefined): boolean {
  if (wornItemId === undefined) return false;
  if (pick.item_id === wornItemId) return true;
  return alternativeFor(pick, wornItemId)?.dps_delta === 0;
}

/**
 * The worn-vs-BiS comparison for one character at one band (§3.B.2's "Best in slot" card and
 * §3.B.3's "Your upgrades" table share this one computation). A slot the band has no known
 * source for (`hasKnownSource` false, or missing from `band.slots` entirely) is skipped
 * outright -- neither an upgrade nor "already best in slot": the ranker itself could not
 * source anything there, so this module has no honest verdict to offer on it either.
 */
export function upgradesFor(
  character: Pick<MeCharacter, 'build'>,
  band: BisBand,
  items: ReadonlyMap<number, Item>,
): UpgradesResult {
  const weights = weightsRecordFor(band);
  const gear = character.build?.gear ?? {};
  const bySlot = new Map(band.slots.filter(hasKnownSource).map((slot) => [slot.slot, slot]));

  const upgrades: SlotUpgrade[] = [];
  const alreadyBis: UpgradesResult['alreadyBis'] = [];
  let totalGainDps = 0;
  let notSimCheckedCount = 0;

  for (const slot of SLOTS) {
    const pick = bySlot.get(slot);
    if (pick === undefined) continue;
    const wornItemId = gear[slot];

    if (tiesWithPick(pick, wornItemId)) {
      if (wornItemId !== undefined) {
        alreadyBis.push({
          slot,
          itemId: wornItemId,
          itemName: items.get(wornItemId)?.name ?? pick.item_name,
        });
      }
      continue;
    }

    const wornItem = wornItemId === undefined ? undefined : items.get(wornItemId);
    const wornUnknown = wornItemId !== undefined && wornItem === undefined;
    const { gainDps, notSimChecked } = gainFor(slot, pick, wornItemId, wornItem, wornUnknown, band, weights);
    if (gainDps !== null) totalGainDps += gainDps;
    if (notSimChecked) notSimCheckedCount += 1;

    upgrades.push({
      slot,
      pick,
      wornItemId,
      wornItemName: wornItem?.name,
      wornUnknown,
      gainDps,
      notSimChecked,
    });
  }

  upgrades.sort((a, b) => (b.gainDps ?? Number.NEGATIVE_INFINITY) - (a.gainDps ?? Number.NEGATIVE_INFINITY));
  return { upgrades, alreadyBis, totalGainDps, notSimCheckedCount, scoreUnit: slotScoreUnitFor(band) };
}

/** Rule (fix round 1): (a) a tie is filtered out before this runs; (b) a worn item listed in
 *  the pick's own `alternatives` always wins -- the ranker's own sim-verified number, never
 *  second-guessed by `scoreItem`; (c) otherwise, for a non-weapon slot, the `scoreItem`
 *  stat-weight diff; (d) otherwise (a weapon slot with no alternative match), the gain is
 *  unknown -- never diff a full-sim `score` (which bakes in weapon DPS) against a
 *  `scoreItem` approximation (which cannot). */
function gainFor(
  slot: Slot,
  pick: BisSlot,
  wornItemId: number | undefined,
  wornItem: Item | undefined,
  wornUnknown: boolean,
  band: BisBand,
  weights: SpecWeights,
): { gainDps: number | null; notSimChecked: boolean } {
  const alternative = alternativeFor(pick, wornItemId);
  if (alternative !== undefined) {
    // tiesWithPick already filtered out dps_delta === 0; every alternative reaching here is a
    // real, sim-verified gain.
    return { gainDps: -alternative.dps_delta, notSimChecked: false };
  }
  if (WEAPON_SLOTS.has(slot)) {
    return { gainDps: null, notSimChecked: true };
  }
  if (
    pick.score === undefined ||
    band.reference_dps_per_point === undefined ||
    band.reference_dps_per_point === null ||
    wornUnknown
  ) {
    return { gainDps: null, notSimChecked: false };
  }
  const wornScore = wornItem === undefined ? 0 : scoreItem(wornItem, weights);
  return { gainDps: (pick.score - wornScore) * band.reference_dps_per_point, notSimChecked: false };
}
