// web/src/lib/sim/enchant-names.ts
// A worn enchant in the client's own words, for every worn-gear list (the sim's Gear, the saved
// sim's Character strip, a report's Gear at the pull): the build's enchant table as id -> name,
// and the slots with their enchant's name. Kept apart from gear-card.ts so a page that only needs
// the names does not load the Gear card's stats and set lines.
import type { PaperdollSlot, PaperdollSlotKey } from '../gear/paperdoll';
import type { EnchantRow } from './enchants';
import type { GearSlot } from './types';

/** The client's name for each enchant id. A row with no `display` draws no line: the recipe name
 *  (`name`, "Enchant Gloves - Greater Strength") is never shown. */
export function enchantNames(rows: readonly EnchantRow[]): ReadonlyMap<number, string> {
  return new Map(
    rows.flatMap((row): [number, string][] =>
      row.display === undefined || row.display === '' ? [] : [[row.id, row.display]],
    ),
  );
}

/** Each slot with its worn enchant's client name, from the engine's gear rows (`gear_slots`). */
export function withEnchants(
  slots: readonly PaperdollSlot[],
  gear: readonly GearSlot[],
  enchants: ReadonlyMap<number, string>,
): PaperdollSlot[] {
  const bySlot = new Map(gear.map((row) => [row.slot, row]));
  return slots.map((slot) => {
    const enchantId = bySlot.get(slot.slot as PaperdollSlotKey)?.enchant;
    const enchant = enchantId === undefined ? undefined : enchants.get(enchantId);
    return enchant === undefined ? slot : { ...slot, enchant };
  });
}
