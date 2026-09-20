// web/src/lib/sim/addon-export.ts
// "Copy to addon" (design 3.3): the winning set as an export string the companion and the
// in-game addon read, so a player can see the swaps in game.
//
// A gear entry carries `item_id[:enchant[:suffix]]`, which contract 10.5 makes explicit
// for `<gear>` and `sets=` and not only for the bags and bank sections: "a version-1
// decoder reading a bare id is unaffected". A set with no enchants therefore encodes
// byte-for-byte as version 1, and every existing decoder reads it unchanged.
//
// This reuses `fs1.ts`'s own `FS1_PREFIX` rather than retyping "FS1". Everything else here
// is new: `fs1.ts`'s `encodeGearSlots` takes a `Map`-ordered, camelCase `FS1GearSlot[]` and
// always reorders by `SLOTS`, while a winning combo's gear is the envelope's own snake_case
// `GearSlot[]` and the order it already won in -- so the gear-entry grammar (identical in
// shape) is written fresh against that type, and the talents conversion (the engine's
// dash-joined tree string into FS1's slash-joined one) has no counterpart in `fs1.ts` at
// all: `encodeTrees` there encodes raw rank arrays, not an already-encoded string.
import { FS1_PREFIX } from '../planner/fs1';
import type { GearSlot } from './types';

/** `slot=item[:enchant[:suffix]]`. A suffix with no enchant writes the enchant as 0. */
export function gearEntry(slot: GearSlot): string {
  const enchant = slot.enchant ?? 0;
  const suffix = slot.suffix ?? 0;
  if (suffix > 0) return `${slot.slot}=${slot.item_id}:${enchant}:${suffix}`;
  if (enchant > 0) return `${slot.slot}=${slot.item_id}:${enchant}`;
  return `${slot.slot}=${slot.item_id}`;
}

export function addonGearList(gear: readonly GearSlot[]): string {
  return gear.map(gearEntry).join(',');
}

export interface AddonStringInput {
  dataBuild: string;
  classSlug: string;
  raceSlug: string;
  /** The engine's talents string, dash-joined; the FS1 grammar slash-joins the same trees. */
  talents: string;
  gear: readonly GearSlot[];
}

/**
 * `FS1:<build>:<class>:<race>:<t1>/<t2>/<t3>:<gear>`.
 *
 * The talents come in as the engine's dash-joined string because that is what a
 * `SimRequest` carries, and the FS1 grammar slash-joins the identical three fields;
 * nothing is re-encoded beyond the separator. A tree the engine trimmed to empty writes as
 * "0", the same way `fs1.ts`'s own `encodeTree` does, so the field count is always three.
 */
export function addonStringFor(input: AddonStringInput): string {
  const trees = input.talents.split('-');
  while (trees.length < 3) trees.push('');
  const encoded = trees.slice(0, 3).map((tree) => (tree === '' ? '0' : tree));
  return [
    FS1_PREFIX,
    input.dataBuild,
    input.classSlug,
    input.raceSlug,
    encoded.join('/'),
    addonGearList(input.gear),
  ].join(':');
}
