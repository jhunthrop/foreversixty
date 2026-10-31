// web/src/lib/gear/paperdoll-slots.ts
// Each call site's data as the paperdoll's slots (design/specs/2026-10-31-paperdoll.md §3.2):
// worn item ids read against the class's items file (the sim card, the report's "At pull"), and
// the character answer's gear panel (the character page). One mapping per data shape, so no
// call site decides a slot's state itself.
import { itemIdLabel } from '../items/unknown-item-copy';
import { SLOTS, type Item } from '../planner/types';
import type { PaperdollItemRef, PaperdollSlot, PaperdollSlotKey } from './paperdoll';
import { upgradeFor } from './paperdoll';
import type { GearItemTable } from './qualities';
import type { CharacterGearSlot } from './types';

/** The pipeline's placeholder icon: drawn as no icon, never a red "?". */
const PLACEHOLDER_ICON = 'inv_misc_questionmark';
/** Drawn only when worn (§3.1). */
const OPTIONAL_SLOTS: readonly PaperdollSlotKey[] = ['shirt', 'tabard'];

export interface WornLookup {
  /** The class's items file row for an id. */
  lookup: (id: number) => Item | undefined;
  /** A name for an id the class file leaves out (itemnames.json). */
  outsideName?: (id: number) => string | undefined;
  /** The items file failed: every worn id is drawn plain, never claimed unknown. */
  failed?: boolean;
  /** The worn copy's own item level (a log records it), over the file's. */
  itemLevel?: (slot: PaperdollSlotKey) => number | undefined;
  /** The report's detail line. */
  detail?: (slot: PaperdollSlotKey) => string | undefined;
}

function iconOf(icon: string | undefined | null): string | null {
  return icon === undefined || icon === null || icon === '' || icon === PLACEHOLDER_ICON ? null : icon;
}

function wornSlot(slot: PaperdollSlotKey, id: number, source: WornLookup): PaperdollSlot {
  const row = source.lookup(id);
  const level = source.itemLevel?.(slot);
  const detail = source.detail?.(slot);
  const base: PaperdollSlot = row
    ? {
        slot,
        state: 'item',
        item: {
          id,
          name: row.name,
          icon: iconOf(row.icon),
          quality: row.quality,
          itemLevel: level ?? row.item_level,
          twoHand: row.two_hand === true,
        },
      }
    : plainOrUnknown(slot, id, source, level);
  return detail === undefined ? base : { ...base, detail };
}

function plainOrUnknown(
  slot: PaperdollSlotKey,
  id: number,
  source: WornLookup,
  level?: number,
): PaperdollSlot {
  const outside = source.outsideName?.(id);
  if (outside === undefined && !source.failed) return { slot, state: 'unknown', unknownId: id };
  return {
    slot,
    state: 'item',
    item: {
      id,
      name: outside ?? itemIdLabel(id),
      icon: null,
      quality: null,
      itemLevel: level ?? null,
      twoHand: false,
    },
  };
}

/** The 17 slots (plus shirt and tabard when worn) from a slot -> worn item id map. */
export function slotsFromWorn(
  worn: Readonly<Partial<Record<PaperdollSlotKey, number>>>,
  source: WornLookup,
): PaperdollSlot[] {
  const keys: PaperdollSlotKey[] = [...SLOTS, ...OPTIONAL_SLOTS.filter((key) => worn[key] !== undefined)];
  return keys.map((slot) => {
    const id = worn[slot];
    return id === undefined || id <= 0 ? { slot, state: 'empty' } : wornSlot(slot, id, source);
  });
}

/** The character page's slots: the API's names and levels, the file's colours and icons, and
 *  each slot's upgrade by §3.4. `table` null or absent: colours and icons left out. */
export function slotsFromCharacterGear(
  gear: readonly CharacterGearSlot[],
  table: GearItemTable | null | undefined,
): PaperdollSlot[] {
  const facts = (id: number) => table?.get(id);
  const ref = (id: number, name: string): PaperdollItemRef => ({
    id,
    name,
    icon: facts(id)?.icon ?? null,
    quality: facts(id)?.quality ?? null,
  });
  const wornBySlot: Record<string, number | undefined> = {};
  for (const slot of gear) wornBySlot[slot.slot] = slot.worn?.id;
  return gear.map((entry): PaperdollSlot => {
    const upgrade = upgradeFor(entry, wornBySlot, ref);
    const withUpgrade = (slot: PaperdollSlot): PaperdollSlot =>
      upgrade === undefined ? slot : { ...slot, upgrade };
    const { worn } = entry;
    if (worn === null) return withUpgrade({ slot: entry.slot, state: 'empty' });
    if (worn.name === null) return withUpgrade({ slot: entry.slot, state: 'unknown', unknownId: worn.id });
    return withUpgrade({
      slot: entry.slot,
      state: 'item',
      item: {
        ...ref(worn.id, worn.name),
        itemLevel: worn.item_level,
        twoHand: facts(worn.id)?.twoHand === true,
      },
    });
  });
}
