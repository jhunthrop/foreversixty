// web/src/lib/gear/panel.ts
// The profile's Gear section's model (Task L8.6 and its fix round 1): one row per slot of GET
// /v1/characters/...'s `gear`, in the export's order, the title line from `gear_gap_total`, and
// the line under it that sets that total beside the score card's. Pure:
// CharacterGearPanel.svelte draws what this returns.
import { gearCopy } from './copy';
import type { CharacterGearSlot, GearSlotName } from './types';

/**
 * What a row shows:
 *   upgrade    worn -> pick, the source and chance, the gap
 *   best       the worn item, "Best for your band" (no arrow, no gap): the worn item is the pick,
 *              or the API gave it a gap of 0 (it scores at least as much, or it is the pick of
 *              the other slot of its pair, the rings and trinkets being matched either way round)
 *   empty      "Empty" -> pick, the source and chance, the gap
 *   worn-only  the worn item, "no pick for your band yet"
 *   none       "Empty", "no pick for your band yet"
 */
export type GearRowState = 'upgrade' | 'best' | 'empty' | 'worn-only' | 'none';

export interface GearItemView {
  id: number;
  /** The item's name; "Unknown item" when the export's data build does not have it. */
  name: string;
  /** False for an item the data build does not have: no tooltip, no icon, muted. */
  known: boolean;
}

export interface GearRow {
  slot: GearSlotName;
  /** "Head", "Ring 1", "Trinket 2", "Main hand". */
  label: string;
  state: GearRowState;
  /** Null when the slot is empty. */
  worn: GearItemView | null;
  /** The pick after the arrow; null on a row with no arrow (best, worn-only, none). */
  pick: GearItemView | null;
  source: string | null;
  /** "17% drop", "<1% drop"; null when the source states no rate, or no pick is shown. */
  chance: string | null;
  /** "Best for your band", "no pick for your band yet", "comes with the main-hand swap". */
  note: string | null;
  /** The row has no source line: below 1024px its note shares the items' line (two lines). */
  inline: boolean;
  /** "+18.4 pts", "+0 pts", "Not scored"; null when the gap cell is blank. */
  gap: string | null;
  /** The gap read aloud, "18.4 points to gain"; null when the text says it all. */
  gapLabel: string | null;
  /** One of the five largest non-zero gaps: drawn in weight 700 (fix round 1 ruling 6). */
  strong: boolean;
}

/** How many gaps are drawn strong. */
const STRONG_GAPS = 5;

function wornView(worn: CharacterGearSlot['worn']): GearItemView | null {
  if (worn === null) return null;
  return worn.name === null
    ? { id: worn.id, name: gearCopy.unknownItem, known: false }
    : { id: worn.id, name: worn.name, known: true };
}

/** The slots whose gap is one of the five largest above 0, ties by export order. */
function strongSlots(gear: readonly CharacterGearSlot[]): Set<GearSlotName> {
  const ranked = gear
    .map((slot, index) => ({ slot: slot.slot, gap: slot.gap, index }))
    .filter((entry): entry is { slot: GearSlotName; gap: number; index: number } => (entry.gap ?? 0) > 0)
    .sort((a, b) => b.gap - a.gap || a.index - b.index);
  return new Set(ranked.slice(0, STRONG_GAPS).map((entry) => entry.slot));
}

/**
 * The section's rows, one per slot, in the order the API sends them (the export's).
 * `twoHand` says whether an item is two-handed, from the build's items file: an empty off hand
 * under a worn two-hander, its gap 0, has its pick "come with the main-hand swap" and a blank gap
 * (fix round 1 ruling 9); without it the API's "+0 pts" stands.
 */
export function gearRows(
  gear: readonly CharacterGearSlot[],
  twoHand: (id: number) => boolean = () => false,
): GearRow[] {
  const strong = strongSlots(gear);
  const mainHand = gear.find((slot) => slot.slot === 'main_hand')?.worn ?? null;
  return gear.map((slot): GearRow => {
    const label = gearCopy.slotLabels[slot.slot] ?? slot.slot;
    const worn = wornView(slot.worn);
    const bare = {
      slot: slot.slot,
      label,
      worn,
      source: null,
      chance: null,
      gap: null,
      gapLabel: null,
      strong: false,
    };
    const { pick } = slot;
    if (pick === null)
      return {
        ...bare,
        state: worn === null ? 'none' : 'worn-only',
        pick: null,
        note: gearCopy.noPick,
        inline: true,
      };
    if (slot.worn !== null && (slot.worn.id === pick.id || slot.gap === 0))
      return { ...bare, state: 'best', pick: null, note: gearCopy.best, inline: true };
    const swap =
      slot.slot === 'off_hand' &&
      slot.worn === null &&
      slot.gap === 0 &&
      mainHand !== null &&
      twoHand(mainHand.id);
    const scored = slot.gap !== null;
    return {
      ...bare,
      state: worn === null ? 'empty' : 'upgrade',
      pick: { id: pick.id, name: pick.name, known: true },
      source: pick.source,
      chance: pick.chance === null ? null : gearCopy.chance(pick.chance),
      note: swap ? gearCopy.swap : null,
      inline: false,
      gap: swap ? null : scored ? gearCopy.gap(slot.gap ?? 0) : gearCopy.gapUnscored,
      gapLabel: swap || !scored ? null : gearCopy.gapLabel(slot.gap ?? 0),
      strong: strong.has(slot.slot),
    };
  });
}

/** How many slots have something to gain: a gap above 0. */
export function gearGapCount(gear: readonly CharacterGearSlot[]): number {
  return gear.filter((slot) => slot.gap !== null && slot.gap > 0).length;
}

/** How many slots have a pick the API could not score. */
export function gearUnscoredCount(gear: readonly CharacterGearSlot[]): number {
  return gear.filter((slot) => slot.pick !== null && slot.gap === null).length;
}

/**
 * The title line (fix round 1 ruling 3): "Gear · +70 pts across 7 slots" from gear_gap_total and
 * the count of non-zero gaps; "Gear · every slot is the best you can wear" at 0; "Gear" when the
 * API scored no slot (a null total); then " · 1 not scored" when a slot's pick has no gap.
 */
export function gearTitle(gear: readonly CharacterGearSlot[], total: number | null): string {
  const unscored = gearUnscoredCount(gear);
  const tail = unscored > 0 ? gearCopy.titleNotScored(unscored) : '';
  if (total === null) return `${gearCopy.heading}${tail}`;
  const slots = gearGapCount(gear);
  if (total <= 0 || slots === 0) return `${gearCopy.titleZero}${tail}`;
  return `${gearCopy.title(total, slots)}${tail}`;
}

/** The line under the title: the card's gear gap and gear_gap_total, what each counts. Null
 *  when the card has no gear gap (no score, or Gear at its top) or the API scored no slot. */
export function gearTotalsLine(cardGap: number | null, total: number | null): string | null {
  if (cardGap === null || cardGap <= 0 || total === null) return null;
  return gearCopy.totals(cardGap, total);
}

/** The Gear section's id: the score card's "Most to gain: Gear" opens `#gear` (L8.6 ruling 2). */
export const GEAR_ANCHOR = 'gear';
