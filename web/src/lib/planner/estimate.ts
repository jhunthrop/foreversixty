// web/src/lib/planner/estimate.ts
// What the planner's live estimate sims (spec 2026-11-08 decision P1, gaps G1 and G7): the build
// on the gear the Level build section names (that band's published best-in-slot set, with its
// enchants), at the character's own level with only the points that level earns. Before this the
// estimate ran the planner's own gear map, empty for every link and fresh build since the gear
// panel left: a naked character's damage, a third of the band's figure beside it.
import type { BisBand } from '../bis/types';
import { MAX_LEVEL, specOf, talentLevel, type SimCharacter } from '../sim/character';
import type { GearSlot } from '../sim/types';
import type { FS1GearSlot } from './fs1';
import { pointsAvailableAt } from './derive';
import type { TalentIndex } from './rules';
import { SLOTS, type Gear, type Slot } from './types';

const isSlot = (slot: string): slot is Slot => (SLOTS as readonly string[]).includes(slot);

/** A band's picks as the engine's gear list and the planner's id map; empty slots are left out. */
export function bandGear(band: BisBand): { gear: Gear; slots: GearSlot[] } {
  const slots: GearSlot[] = band.slots.flatMap((row) =>
    isSlot(row.slot) && row.item_id > 0
      ? [{ slot: row.slot, item_id: row.item_id, ...(row.enchant_id ? { enchant: row.enchant_id } : {}) }]
      : [],
  );
  const gear = Object.fromEntries(slots.map((row) => [row.slot, row.item_id])) as Gear;
  return { gear, slots };
}

/** The same picks as an FS1 code's gear (the `Sim this build ›` link sims what the head shows). */
export function bandCodeGear(band: BisBand): { gear: Gear; gearSlots: FS1GearSlot[] } {
  const { gear, slots } = bandGear(band);
  return {
    gear,
    gearSlots: slots.map((row) => ({
      slot: row.slot as Slot,
      itemId: row.item_id,
      ...(row.enchant === undefined ? {} : { enchant: row.enchant }),
    })),
  };
}

/** The points a character's level earns of `order` (all of them with no character). */
export function earnedOrder<T>(order: readonly T[], characterLevel: number | null): T[] {
  return characterLevel === null ? [...order] : order.slice(0, pointsAvailableAt(characterLevel));
}

/**
 * `base` (the planner's character, `characterFromPlanner`) on `band`'s gear when one is loaded,
 * cut to `characterLevel`'s points and simmed at that level when a character is: points planned
 * past the level (decision P4) are not the character's yet, so they never move its number.
 */
export function estimateCharacter(
  base: SimCharacter,
  index: TalentIndex,
  band: BisBand | null,
  characterLevel: number | null,
): SimCharacter {
  const order = earnedOrder(base.point_order, characterLevel);
  const worn = band === null ? null : bandGear(band);
  return {
    ...base,
    point_order: order,
    spec: specOf(index, order),
    talent_level: talentLevel(order),
    level: characterLevel === null ? MAX_LEVEL : Math.min(MAX_LEVEL, characterLevel),
    gear: worn === null ? base.gear : worn.gear,
    gear_slots: worn === null ? base.gear_slots : worn.slots,
  };
}
