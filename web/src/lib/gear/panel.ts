// web/src/lib/gear/panel.ts
// The profile's Gear section's model (Task L8.6 ruling 1): one row per slot of GET
// /v1/characters/...'s `gear`, in the export's order, and the title line from
// `gear_gap_total`. Pure: CharacterGearPanel.svelte draws what this returns.
import { SLOT_DISPLAY_LABELS } from '../bis/slot-display-labels';
import { gearCopy } from './copy';
import type { CharacterGearSlot, GearSlotName } from './types';

/**
 * What a row shows:
 *   upgrade    worn -> pick · source · chance, the gap
 *   best       the worn item, "Best for your band" (no arrow, no gap): the worn item is the pick,
 *              or the API gave it a gap of 0 (it scores at least as much, or it is the pick of
 *              the other slot of its pair, the rings and trinkets being matched either way round)
 *   empty      "Empty" -> pick · source · chance, the gap
 *   worn-only  the worn item alone: the band has no pick for the slot
 *   none       "Empty" alone: nothing worn and no pick
 */
export type GearRowState = 'upgrade' | 'best' | 'empty' | 'worn-only' | 'none';

export interface GearItemView {
  id: number;
  /** The item's name; "Unknown item" when the export's data build does not have it. */
  name: string;
  /** False for an item the data build does not have: no tooltip, no quality, muted. */
  known: boolean;
}

export interface GearRow {
  slot: GearSlotName;
  /** The client's paperdoll label: "Head", "Ring", "Main hand". */
  label: string;
  state: GearRowState;
  /** Null when the slot is empty. */
  worn: GearItemView | null;
  /** The pick after the arrow; null on a row with no arrow (best, worn-only, none). */
  pick: GearItemView | null;
  source: string | null;
  /** "17%", "<1%"; null when the source states no rate, or no pick is shown. */
  chance: string | null;
  /** "+18.4 pts", "–" when the API could not score it; null when the row shows no gap. */
  gap: string | null;
  /** The gap read aloud: "18.4 points to gain", "Not scored". */
  gapLabel: string | null;
}

function wornView(worn: CharacterGearSlot['worn']): GearItemView | null {
  if (worn === null) return null;
  return worn.name === null
    ? { id: worn.id, name: gearCopy.unknownItem, known: false }
    : { id: worn.id, name: worn.name, known: true };
}

function rowOf(slot: CharacterGearSlot): GearRow {
  const label = SLOT_DISPLAY_LABELS[slot.slot] ?? slot.slot;
  const worn = wornView(slot.worn);
  const bare = { slot: slot.slot, label, worn, source: null, chance: null, gap: null, gapLabel: null };
  const { pick } = slot;
  if (pick === null) return { ...bare, state: worn === null ? 'none' : 'worn-only', pick: null };
  if (slot.worn !== null && (slot.worn.id === pick.id || slot.gap === 0))
    return { ...bare, state: 'best', pick: null };
  return {
    ...bare,
    state: worn === null ? 'empty' : 'upgrade',
    pick: { id: pick.id, name: pick.name, known: true },
    source: pick.source,
    chance: pick.chance === null ? null : gearCopy.chance(pick.chance),
    gap: slot.gap === null ? gearCopy.gapUnscored : gearCopy.gap(slot.gap),
    gapLabel: slot.gap === null ? gearCopy.gapUnscoredLabel : gearCopy.gapLabel(slot.gap),
  };
}

/** The section's rows, one per slot, in the order the API sends them (the export's). */
export function gearRows(gear: readonly CharacterGearSlot[]): GearRow[] {
  return gear.map(rowOf);
}

/** How many slots have something to gain: a gap above 0. */
export function gearGapCount(gear: readonly CharacterGearSlot[]): number {
  return gear.filter((slot) => slot.gap !== null && slot.gap > 0).length;
}

/** The title line: "Gear · +88 pts across 6 slots" from gear_gap_total and the count of
 *  non-zero gaps; "Gear · nothing to gain over the picks" at 0; "Gear" when the API scored no
 *  slot (a null total). */
export function gearTitle(gear: readonly CharacterGearSlot[], total: number | null): string {
  if (total === null) return gearCopy.heading;
  const slots = gearGapCount(gear);
  if (total <= 0 || slots === 0) return gearCopy.titleNoGaps;
  return gearCopy.title(total, slots);
}

/** The Gear section's id: the score card's "Most to gain: Gear" opens `#gear` (L8.6 ruling 2). */
export const GEAR_ANCHOR = 'gear';
