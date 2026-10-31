// web/src/lib/gear/panel.ts
// The profile's Gear section's model (Task L8.6 and its fix round 1): one row per slot of GET
// /v1/characters/...'s `gear`, in the export's order, the title line from `gear_gap_total`, and
// the line under it that sets that total beside the score card's. Pure:
// CharacterGearPanel.svelte draws what this returns.
import { wearsListedBest } from '../bis/best-in-slot';
import { SLOT_LABELS } from '../planner/types';
import { gearCopy } from './copy';
import type { CharacterGearSlot, GearSlotName } from './types';

/**
 * What a row shows:
 *   upgrade    worn -> pick, the source and chance, the gap
 *   best       the worn item, "Already your best at this level" (no arrow, no gap): the worn
 *              item is the listed pick, or the pick of the other slot of its pair (the rings and
 *              trinkets matched either way round). The one definition the best-in-slot card counts
 *              with (lib/bis/best-in-slot.ts)
 *   empty      "Empty" -> pick, the source and chance, the gap
 *   worn-only  the worn item, "no pick for your band yet"
 *   none       "Empty", "no pick for your band yet"
 */
export type GearRowState = 'upgrade' | 'best' | 'empty' | 'worn-only' | 'none';

export interface GearItemView {
  id: number;
  /** The item's name; for a worn item the export's data build does not have (no name), the muted
   *  "Not in our item data yet" (gearCopy.unknownItem, fix round 4 item 5). */
  name: string;
  /** False for an item the data build does not have: no tooltip, no icon, muted. */
  known: boolean;
}

export interface GearRow {
  slot: GearSlotName;
  /** "Head", "Finger 1", "Trinket 2", "Main hand" (SLOT_LABELS). */
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
  /** The row has no source line: its note is the meta line's whole text, below 1024px a line of
   *  its own under the items with no dot (fix round 5 item 1). */
  inline: boolean;
  /** "+18.4 pts", "+0 pts", "Not scored"; null when the gap cell is blank. */
  gap: string | null;
  /** The gap read aloud, "18.4 points to gain"; null when the text says it all. */
  gapLabel: string | null;
  /** One of the five largest non-zero gaps: drawn in weight 700 (fix round 1 ruling 6). */
  strong: boolean;
  /** The note is the main-hand swap's, but the items file has not said (loading) or cannot say
   *  (failed) whether the worn main hand is two-handed: its line is held, not shown (fix round 4
   *  item 4). */
  noteUnknown: boolean;
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
 * (fix round 1 ruling 9); with a one-handed main hand the API's "+0 pts" stands. `null` is an
 * items file that has not answered or failed: the swap cannot be known, so that off hand shows no
 * figure (fix round 2 ruling 7) and holds the note's line (`noteUnknown`, fix round 4 item 4).
 */
export function gearRows(
  gear: readonly CharacterGearSlot[],
  twoHand: ((id: number) => boolean) | null = () => false,
): GearRow[] {
  const strong = strongSlots(gear);
  const wornBySlot: Record<string, number | undefined> = {};
  for (const slot of gear) wornBySlot[slot.slot] = slot.worn?.id;
  const mainHand = gear.find((slot) => slot.slot === 'main_hand')?.worn ?? null;
  return gear.map((slot): GearRow => {
    const label = SLOT_LABELS[slot.slot];
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
      noteUnknown: false,
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
    if (slot.worn !== null && wearsListedBest(slot.slot, pick.id, wornBySlot))
      return { ...bare, state: 'best', pick: null, note: gearCopy.best, inline: true };
    const emptyOffHandAtZero = slot.slot === 'off_hand' && slot.worn === null && slot.gap === 0;
    const swap = emptyOffHandAtZero && twoHand !== null && mainHand !== null && twoHand(mainHand.id);
    const unknown = emptyOffHandAtZero && twoHand === null && mainHand !== null;
    const blank = swap || unknown;
    const scored = slot.gap !== null;
    return {
      ...bare,
      state: worn === null ? 'empty' : 'upgrade',
      pick: { id: pick.id, name: pick.name, known: true },
      source: pick.source,
      chance: pick.chance === null ? null : gearCopy.chance(pick.chance),
      note: swap || unknown ? gearCopy.swap : null,
      noteUnknown: unknown,
      inline: false,
      gap: blank ? null : scored ? gearCopy.gap(slot.gap ?? 0) : gearCopy.gapUnscored,
      gapLabel: blank || !scored ? null : gearCopy.gapLabel(slot.gap ?? 0),
      strong: strong.has(slot.slot),
    };
  });
}

/**
 * The spec the picks are for, in muted words after the heading: "Fury · from your talents", or
 * "Arms · no points spent yet" when nothing is spent. Null when a fight named the spec (the
 * usual case, nothing to explain) or the API sent no spec. "warrior-beast-mastery" style slugs
 * lose their class: "Beast Mastery".
 */
export function gearSpecLine(
  spec: string | null | undefined,
  source: string | null | undefined,
): string | null {
  if (!spec || (source !== 'talents' && source !== 'default')) return null;
  const name = spec
    .slice(spec.indexOf('-') + 1)
    .split('-')
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
    .join(' ');
  return `${name} · ${source === 'talents' ? gearCopy.specFromTalents : gearCopy.specDefault}`;
}

/** The slots wearing their listed best, of the slots that have a pick: the card's "3 of 17 slots". */
export function gearBestCount(gear: readonly CharacterGearSlot[]): { worn: number; total: number } {
  const rows = gearRows(gear);
  return {
    worn: rows.filter((row) => row.state === 'best').length,
    total: gear.filter((slot) => slot.pick !== null).length,
  };
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
 * What follows "Gear" on the title line (fix round 1 ruling 3, fix round 2 rulings 4 and 5):
 * "+70 pts across 7 slots" from gear_gap_total and the count of non-zero gaps; at a total of 0
 * "every slot is the best you can wear" when no row shows an arrow, else "+0 pts from the picks
 * you can wear"; then " · 1 not scored" when a slot's pick has no gap, its dot glued to the word
 * before it. Null when there is nothing to say (no total, every pick scored).
 */
export function gearTitleSummary(gear: readonly CharacterGearSlot[], total: number | null): string | null {
  const unscored = gearUnscoredCount(gear);
  if (total === null) return unscored > 0 ? gearCopy.unscored(unscored) : null;
  const slots = gearGapCount(gear);
  const arrow = gearRows(gear).some((row) => row.pick !== null);
  const gain =
    total > 0 && slots > 0
      ? gearCopy.summary(total, slots)
      : arrow
        ? gearCopy.summaryZero
        : gearCopy.summaryAllWorn;
  const best = gearBestCount(gear);
  const lead =
    gain === gearCopy.summaryAllWorn || best.total === 0
      ? gain
      : `${gearCopy.bestCount(best.worn, best.total)}${gearCopy.titleSeparator}${gain}`;
  return unscored > 0 ? `${lead}${gearCopy.notScored(unscored)}` : lead;
}

/** The title line as read: "Gear · +70 pts across 7 slots · 1 not scored", or "Gear". */
export function gearTitle(gear: readonly CharacterGearSlot[], total: number | null): string {
  const summary = gearTitleSummary(gear, total);
  return summary === null ? gearCopy.heading : `${gearCopy.heading}${gearCopy.titleSeparator}${summary}`;
}

/** The Gear section's id: the score card's "Most to gain: Gear" opens `#gear` (L8.6 ruling 2). */
export const GEAR_ANCHOR = 'gear';
