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
import { bandTreeRanksFrom } from '../planner/band-compare';
import { encodeFS1V2, orderFromRanks, type FS1GearSlot } from '../planner/fs1';
import type { TalentIndex } from '../planner/rules';
import { talentsString } from '../sim/character';
import { SLOTS, type Gear, type Slot } from '../planner/types';
import { bandTopLevel } from './rotation-view';
import type { BisBand } from './types';

function isSlot(slot: string): slot is Slot {
  return (SLOTS as readonly string[]).includes(slot);
}

/** The band's picks in the planner's slots; an empty slot (no pick) is left out. */
function bandGearSlots(band: BisBand): FS1GearSlot[] {
  return band.slots.flatMap((row) =>
    isSlot(row.slot) && row.item_id > 0 ? [{ slot: row.slot, itemId: row.item_id }] : [],
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

/** The simulator, opened on exactly this band's set, over the class's talent trees `index`. */
export function bisSimHref(band: BisBand, index: TalentIndex, dataBuild: string): string {
  const gearSlots = bandGearSlots(band);
  return simCodeHref(
    encodeFS1V2({
      dataBuild,
      classSlug: index.file.class_slug,
      raceSlug: band.race,
      level: bandTopLevel(band.band),
      treeRanks: reachableTreeRanks(band, index),
      gear: gearOf(gearSlots),
      gearSlots,
    }),
  );
}
