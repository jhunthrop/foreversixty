// web/src/lib/bis/sim-link.ts
// "Open in simulator" on a BiS band: the band's whole set as the FS1 code the simulator's
// `?code=` bootstrap already reads (lib/sim/url.ts). The simulator has no `?spec=`/`?band=`
// vocabulary, so a link naming only those opened it empty; the code carries everything the
// set is: the class, the band's race, its talents (from which the simulator names the spec),
// the band's level and every slot's pick.
//
// The talents are the band's points the active trees can reach, in the order the planner's
// own band load rebuilds (`orderFromRanks`): the simulator refuses a whole code over one
// unreachable point, and a band ranked on a client whose trees have since been reshaped can
// carry one until the ranking re-spends it.
import { simCodeHref } from '../handoff-links';
import { defaultSimState, simSearch, withSimState } from '../sim/url';
import { bandTreeRanksFrom } from '../planner/band-compare';
import { encodeFS1V2, orderFromRanks, type FS1GearSlot } from '../planner/fs1';
import type { TalentIndex } from '../planner/rules';
import { talentsString } from '../sim/character';
import { SLOTS, type Gear, type Slot } from '../planner/types';
import { MAX_LEVEL } from '../sim/character';
import { isSimulatedSpec } from '../sim/spec-label';
import { presetOf, RAID_PRESET } from './presets';
import { bandTopLevel } from './rotation-view';
import type { BisBand } from './types';

function isSlot(slot: string): slot is Slot {
  return (SLOTS as readonly string[]).includes(slot);
}

/** The band's picks and their enchants in the planner's slots, so the simulator measures the
 *  set the band published; an empty slot (no pick) is left out. */
function bandGearSlots(band: BisBand): FS1GearSlot[] {
  return band.slots.flatMap((row) =>
    isSlot(row.slot) && row.item_id > 0
      ? [
          {
            slot: row.slot,
            itemId: row.item_id,
            ...(row.enchant_id ? { enchant: row.enchant_id } : {}),
          },
        ]
      : [],
  );
}

function gearOf(slots: readonly FS1GearSlot[]): Gear {
  return Object.fromEntries(slots.map((entry) => [entry.slot, entry.itemId])) as Gear;
}

/** The band's ranks, cut to the points `index`'s trees can reach. */
function reachableTreeRanks(band: BisBand, index: TalentIndex): number[][] {
  const { order } = orderFromRanks(index, bandTreeRanksFrom(band.talents));
  return bandTreeRanksFrom(talentsString(index, order));
}

/** The band's whole set as one FS1 v2 code, over the class's talent trees `index`. */
export function bandSimCode(band: BisBand, index: TalentIndex, dataBuild: string): string {
  const gearSlots = bandGearSlots(band);
  return encodeFS1V2({
    dataBuild,
    classSlug: index.file.class_slug,
    raceSlug: band.race,
    level: bandTopLevel(band.band),
    treeRanks: reachableTreeRanks(band, index),
    gear: gearOf(gearSlots),
    gearSlots,
  });
}

/** The simulator, opened on exactly this band's set, over the class's talent trees `index`. */
export function bisSimHref(band: BisBand, index: TalentIndex, dataBuild: string): string {
  return simCodeHref(bandSimCode(band, index, dataBuild));
}

/**
 * "View ranking sim": the band's set opened in the simulator and run at once. The code carries
 * the gear, enchants, talents, race and level; the simulator's defaults at that level are the
 * ranker's own preset and fight (lib/sim/preset-sets.ts, sim/leveling.RankerEncounter), so the
 * run reproduces the published figure within its error.
 */
export function rankingSimHref(band: BisBand, index: TalentIndex, dataBuild: string): string {
  const code = bandSimCode(band, index, dataBuild);
  return `/sim${simSearch(withSimState(defaultSimState(), { code, run: true }))}`;
}

/**
 * Whether the simulator can rerun the band's published figure: the level 60 raid-ready band
 * of a damage spec. /sim's defaults at 60 are exactly that preset and the ranker's fight; a
 * bare band would run with the raid buffs on, a leveling band against another fight, and a
 * tank's damage taken or a healer's healing is not a number the browser sim produces
 * (RunControl runs damage specs only).
 */
export function reproducesOnSim(band: BisBand): boolean {
  return (
    isSimulatedSpec(band.spec) && presetOf(band) === RAID_PRESET && bandTopLevel(band.band) === MAX_LEVEL
  );
}

/** "View ranking sim" where the run reproduces the band's figure (`reproducesOnSim`), else `undefined`. */
export function rankingSimHrefFor(band: BisBand, index: TalentIndex, dataBuild: string): string | undefined {
  return reproducesOnSim(band) ? rankingSimHref(band, index, dataBuild) : undefined;
}
