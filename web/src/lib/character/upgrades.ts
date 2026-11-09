// web/src/lib/character/upgrades.ts
// The character paperdoll's Upgrades list (directions mock b2-character): the band's picks with a
// priced gain over what is worn, largest first, in the frame's centre (under the slot list on a
// phone). Ties and unpriced slots are left out, as upgradeTotal leaves them out of the band's sum.
import { isTie } from '../bis/doll-view';
import { slotLabel, type PaperdollSlot } from '../gear/paperdoll';
import type { ItemRef } from '../ui/primitives';

/** An item the data does not grade reads as common, as the doll draws it. */
const PLAIN_QUALITY = 1;

/** The frame's centre holds six rows; the rest are counted under them. */
export const UPGRADES_SHOWN = 6;

export interface UpgradeRowView {
  slot: string;
  /** The slot's word ("Main hand"). */
  label: string;
  /** The band's pick: the item to get. */
  item: ItemRef;
  /** The priced gain, in the band's unit. */
  value: number;
}

export function upgradeRows(
  slots: readonly PaperdollSlot[],
  classSlug: string,
  limit: number = Number.POSITIVE_INFINITY,
): UpgradeRowView[] {
  return slots
    .flatMap((slot): UpgradeRowView[] => {
      const up = slot.upgrade;
      if (up?.kind !== 'gain') return [];
      const value = up.value;
      if (value === null || isTie(value) || value <= 0) return [];
      const { id, name, quality, icon } = up.pick;
      return [
        {
          slot: slot.slot,
          label: slotLabel(slot.slot, classSlug),
          item: { id, name, quality: quality ?? PLAIN_QUALITY, icon: icon ?? undefined },
          value,
        },
      ];
    })
    .sort((a, b) => b.value - a.value)
    .slice(0, limit);
}
