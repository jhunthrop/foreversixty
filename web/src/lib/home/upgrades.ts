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
import type { BisBand, BisSlot } from '../bis/types';
import { SLOTS, type Item, type Slot } from '../planner/types';

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
   * The pick's own score minus the worn item's score, converted from the band's
   * `reference_stat_points` units to real DPS via the band's own `reference_dps_per_point`
   * (the identical conversion the BiS page's weight rail already documents for turning a
   * normalized scale factor into "DPS per point" -- see `BisStatWeight.dps_per_point`'s own
   * doc). `null` when either half of that subtraction cannot be honestly computed: the
   * band's own pick carries no `score` at all (a sim-decided row -- see `BisSlot.score`'s
   * own doc, "absent on a row the ranker's own sim decided"), the band carries no
   * `reference_dps_per_point` (a file published before that field existed), or the worn
   * item's id is not in this build's item table (`wornUnknown`). Never a raw score-unit
   * number badged as DPS, and never a fabricated figure for an item this site cannot score.
   */
  gainDps: number | null;
}

export interface UpgradesResult {
  /** Every upgrade slot, sorted by `gainDps` descending (an unknown gain -- `null` --
   *  sorts last, after every slot with a real, measured number). */
  upgrades: SlotUpgrade[];
  /** Every slot where the character already wears the band's own pick (or a scored tie with
   *  it), for the "Already best in slot" footer line -- never claims a slot the band itself
   *  has no known source for. */
  alreadyBis: { slot: Slot; itemId: number; itemName: string }[];
  /** The sum of every known (non-null) `gainDps` across `upgrades`. */
  totalGainDps: number;
}

/** `band.weights` (an array of `{stat, weight}` rows) as the `Record<string, number>`
 *  `scoreItem` (`lib/addon/score.ts`) takes -- one place this reshaping happens, so a BiS
 *  band's own weight list and the addon panel's `stat-weights.json` shape can both feed the
 *  identical scorer without this module caring which one it was given. */
export function weightsRecordFor(band: Pick<BisBand, 'weights'>): SpecWeights {
  return Object.fromEntries(band.weights.map((weight) => [weight.stat, weight.weight]));
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
  return (pick.alternatives ?? []).some((alt) => alt.item_id === wornItemId && alt.dps_delta === 0);
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
    const gainDps = gainDpsFor(pick, wornItem, wornUnknown, band, weights);
    if (gainDps !== null) totalGainDps += gainDps;

    upgrades.push({ slot, pick, wornItemId, wornItemName: wornItem?.name, wornUnknown, gainDps });
  }

  upgrades.sort((a, b) => (b.gainDps ?? Number.NEGATIVE_INFINITY) - (a.gainDps ?? Number.NEGATIVE_INFINITY));
  return { upgrades, alreadyBis, totalGainDps };
}

function gainDpsFor(
  pick: BisSlot,
  wornItem: Item | undefined,
  wornUnknown: boolean,
  band: BisBand,
  weights: SpecWeights,
): number | null {
  if (
    pick.score === undefined ||
    band.reference_dps_per_point === undefined ||
    band.reference_dps_per_point === null
  ) {
    return null;
  }
  if (wornUnknown) return null;
  const wornScore = wornItem === undefined ? 0 : scoreItem(wornItem, weights);
  return (pick.score - wornScore) * band.reference_dps_per_point;
}
