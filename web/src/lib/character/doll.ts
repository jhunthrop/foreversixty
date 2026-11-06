// web/src/lib/character/doll.ts
// The character page's paperdoll cells (design/specs/2026-11-06-character.md C4): the worn set
// through the shared worn mapping (lib/gear/doll-cells.ts), its gains kept for the owner only; and,
// when the export carried no gear, the band's picks with their sources instead of seventeen empty
// slots.
import { dollCellsFromSlots } from '../gear/doll-cells';
import { slotLabel, type PaperdollSlot } from '../gear/paperdoll';
import type { GearItemTable } from '../gear/qualities';
import type { CharacterGearSlot } from '../gear/types';
import type { DollCell } from '../ui/paperdoll';

const PLAIN_QUALITY = 1;

/** C4: the worn cells; a visitor's carry no gain column. */
export function wornCells(slots: readonly PaperdollSlot[], classSlug: string, gains: boolean): DollCell[] {
  const cells = dollCellsFromSlots(slots, 'worn', classSlug);
  return gains ? cells : cells.map(({ gain: _gain, badge: _badge, ...cell }) => cell);
}

/** The unread state: the band's pick for every slot that has one, its source on line 2. */
export function pickCells(
  gear: readonly CharacterGearSlot[],
  table: GearItemTable | null,
  classSlug: string,
): DollCell[] {
  return gear.flatMap((slot): DollCell[] => {
    if (slot.pick === null) return [];
    const facts = table?.get(slot.pick.id);
    return [
      {
        slot: slot.slot,
        label: slotLabel(slot.slot, classSlug),
        item: {
          id: slot.pick.id,
          name: slot.pick.name,
          quality: facts?.quality ?? PLAIN_QUALITY,
          icon: facts?.icon ?? undefined,
        },
        source: slot.pick.source,
      },
    ];
  });
}
