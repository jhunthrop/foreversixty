// web/src/lib/gear/paperdoll.ts
// The worn paperdoll's rules (design/specs/2026-10-31-paperdoll.md, web layout spec 2026-11-05
// §4.1 item 9): each slot's state, its upgrade (the API's verdict, priced in the band's unit), the
// two-hander covering the off hand, unknown items, the item-level ruling and the panel's words.
// The sim card, the character page and the report each map their data to `PaperdollSlot`s
// (paperdoll-slots.ts); doll-cells.ts turns those into the shared doll's cells.
import { wearsListedBest } from '../bis/best-in-slot';
import { itemIdLabel } from '../items/unknown-item-copy';
import type { Slot } from '../planner/types';
import { slotLabelForClass } from '../bis/slot-display-labels';
import type { SlotScoreUnit } from '../bis/tank-view';
import { gearCopy } from './copy';
import type { CharacterGearSlot } from './types';
import { COSMETIC_SLOTS } from './slots';

export type PaperdollSlotKey = Slot | 'shirt' | 'tabard';

/** An item as the frame draws it. `icon` and `quality` are null when the items file has not
 *  said (failed): the plain socket, the name in the body colour. */
export interface PaperdollItem {
  id: number;
  name: string;
  icon: string | null;
  quality: number | null;
  itemLevel: number | null;
  twoHand: boolean;
}

/** A pick, for its icon, colour and tooltip. */
export interface PaperdollItemRef {
  id: number;
  name: string;
  icon: string | null;
  quality: number | null;
}

/** A gain in the band's own unit (`SlotScoreUnit`: DPS, HPS or a tank's score); `value` null when
 *  the rule cannot price it (lib/home/upgrades.ts `slotGainFor`: a weapon no sim checked). */
export type PaperdollUpgrade =
  | { kind: 'gain'; value: number | null; unit: SlotScoreUnit; pick: PaperdollItemRef }
  | { kind: 'best' }
  | { kind: 'unscored'; pick: PaperdollItemRef };

export interface PaperdollSlot {
  slot: PaperdollSlotKey;
  state: 'item' | 'empty' | 'unknown';
  /** With `state: 'item'`. */
  item?: PaperdollItem;
  /** With `state: 'unknown'`: the id the data does not have. */
  unknownId?: number;
  upgrade?: PaperdollUpgrade;
  /** The report's "Item level 66 · Enchanted". */
  detail?: string;
  /** The client's enchant text ("+7 Strength"), the green line under the name. */
  enchant?: string;
}

/** Shirt and tabard are drawn only when the data carries them (§3.1); the item level skips them. */

/**
 * The two-handed main hand covering an empty off hand (§4.3 "Two-hand in main hand"): the doll
 * draws it greyed in the off hand's socket, as the game's character pane does. Null otherwise.
 */
export function twoHandCover(slots: readonly PaperdollSlot[]): PaperdollItemRef | null {
  const main = slots.find((slot) => slot.slot === 'main_hand')?.item;
  const off = slots.find((slot) => slot.slot === 'off_hand');
  return off?.state !== 'item' && main?.twoHand === true ? main : null;
}

/**
 * The centre's item level (§3.5): the mean item level of every equipped slot but shirt and tabard,
 * one decimal. An empty slot (the off hand under a two-hander included) is not equipped. Null when
 * nothing is equipped, or when any equipped item is not in our data or has no known level: never a
 * partial mean.
 */
export function averageItemLevel(slots: readonly PaperdollSlot[]): string | null {
  const counted = slots.filter((slot) => !COSMETIC_SLOTS.has(slot.slot) && slot.state !== 'empty');
  if (counted.length === 0) return null;
  const levels = counted.map((slot) => (slot.state === 'item' ? (slot.item?.itemLevel ?? null) : null));
  if (levels.some((level) => level === null)) return null;
  const total = (levels as number[]).reduce((sum, level) => sum + level, 0);
  return (total / levels.length).toFixed(1);
}

/** How many equipped slots hold an item our data does not have. */
export function unknownCount(slots: readonly PaperdollSlot[]): number {
  return slots.filter((slot) => slot.state === 'unknown').length;
}

/** Prices the API's gain: the pick's gain over the worn id in `unit`, null when it cannot. */
export interface GainPricer {
  unit: SlotScoreUnit;
  price: (slot: PaperdollSlotKey, pickId: number, wornId: number | undefined) => number | null;
}

/**
 * A character-page slot's upgrade (§3.4) from the API's `CharacterGearSlot`. `wornBySlot` is the
 * worn item id per slot, for the rings and trinkets matched either way round; `ref` resolves a
 * pick's icon and quality. The API's gap decides whether there is a gain (never when the pick is
 * worn); `pricer` says how much in the band's unit (web layout spec §4.6, D4: no points).
 */
export function upgradeFor(
  gear: CharacterGearSlot,
  wornBySlot: Readonly<Record<string, number | undefined>>,
  ref: (id: number, name: string) => PaperdollItemRef,
  pricer: GainPricer,
): PaperdollUpgrade | undefined {
  const { worn, pick, gap } = gear;
  if (pick === null) return undefined;
  if (worn !== null && worn.name === null) return { kind: 'unscored', pick: ref(pick.id, pick.name) };
  if (worn !== null && wearsListedBest(gear.slot, pick.id, { ...wornBySlot, [gear.slot]: worn.id }))
    return { kind: 'best' };
  if (gap === null) return worn === null ? { kind: 'unscored', pick: ref(pick.id, pick.name) } : undefined;
  if (gap <= 0) return undefined;
  return {
    kind: 'gain',
    value: pricer.price(gear.slot, pick.id, worn?.id),
    unit: pricer.unit,
    pick: ref(pick.id, pick.name),
  };
}

/** The slot's word: `slotLabelForClass` ("Relic" for the relic classes), plus shirt and tabard. */
export function slotLabel(slot: PaperdollSlotKey, classSlug = ''): string {
  if (slot === 'shirt') return gearCopy.paperdoll.shirt;
  if (slot === 'tabard') return gearCopy.paperdoll.tabard;
  return slotLabelForClass(slot, classSlug);
}

/** "Waist · Item 288887": an unknown item always carries its slot word (§3.3). */
export function unknownName(slot: PaperdollSlotKey, id: number, classSlug = ''): string {
  return `${slotLabel(slot, classSlug)} · ${itemIdLabel(id)}`;
}

/** The note under the stage (§3.3, §6): once, only when a slot is unknown or the file failed. */
export function paperdollNote(unknown: number, failed: boolean): string | null {
  if (failed) return gearCopy.paperdoll.noteFailed;
  if (unknown === 0) return null;
  return unknown === 1 ? gearCopy.paperdoll.noteOne : gearCopy.paperdoll.noteMany(unknown);
}
