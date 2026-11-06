// web/src/lib/ui/paperdoll.ts
// The shared paperdoll's pure model (web layout spec 2026-11-05 §4, owner decision D10): two
// left-aligned columns in the character panel's order, each cell two lines, one selection, and the
// keyboard model. Every surface (BiS pick and compare, the character page and sim card worn, the
// report) maps its own data to `DollCell`s; ui/Paperdoll.svelte draws them the same way in all
// four modes. The mode only names whose item a cell shows; the mapper decides the words.
import type { GainFigure, ItemRef } from './primitives';

export type DollMode = 'pick' | 'compare' | 'worn' | 'report';

/** The level flag under the slot label: `Level 20`, `Usable now` (kill), `At 25` (gold). */
export interface DollFlag {
  text: string;
  tone: 'muted' | 'kill' | 'gold';
}

/** A note the slot's detail shows (the worn item's facts, why a slot was re-picked). */
export interface DollNote {
  /** The detail section's label ("You wear", "Re-picked"). */
  label: string;
  /** Words before the item ("Next best:", "You wear"). */
  lead: string;
  item?: ItemRef;
  /** Words after the item ("−1.5 DPS", "(same DPS)", "(item level 25, +3 Mp5)"). */
  tail?: string;
  tone?: 'muted' | 'wipe' | 'kill';
}

export interface DollCell {
  /** The slot key (`head`, `finger1`, ...). */
  slot: string;
  label: string;
  flag?: DollFlag;
  /** Null for an empty slot. */
  item: ItemRef | null;
  /** The worn id is missing from our data. */
  unknown?: boolean;
  badge?: 'upgrade' | 'worn';
  /** Short stats after the name on line 1: `+8 Agi · +10 Sta`. */
  stats?: string;
  /** Line 2: where it comes from, in full; it wraps on a phone and is never truncated. */
  source?: string;
  /** A source-kind key for the line's glyph (`dungeon`, `crafted`, ...). */
  sourceKind?: string;
  /** Detail only: line 2 of the cell is the source alone (owner decision, option C). */
  note?: DollNote;
  gain?: GainFigure;
  /** The worn item's small icon beside the gain (compare mode). */
  worn?: ItemRef;
  /** An empty slot's reason ("First trinket comes at 40"). */
  empty?: string;
  /** A two-hander covers this off hand: the socket shows it greyed. */
  coveredBy?: ItemRef;
}

/** The character panel's order: head to ranged down the left, hands to trinkets down the right. */
export const DOLL_LEFT = [
  'head',
  'neck',
  'shoulder',
  'back',
  'chest',
  'shirt',
  'tabard',
  'wrist',
  'main_hand',
  'off_hand',
  'ranged',
] as const;
export const DOLL_RIGHT = [
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
] as const;

/** The cells in the doll's two columns, in reading order; slots the caller did not pass are left out. */
export function dollColumns(cells: readonly DollCell[]): { left: DollCell[]; right: DollCell[] } {
  const bySlot = new Map(cells.map((cell) => [cell.slot, cell]));
  const pick = (keys: readonly string[]) => keys.flatMap((key) => bySlot.get(key) ?? []);
  return { left: pick(DOLL_LEFT), right: pick(DOLL_RIGHT) };
}

/** Reading order: down the left column, then down the right. */
export function readingOrder(cells: readonly DollCell[]): DollCell[] {
  const { left, right } = dollColumns(cells);
  return [...left, ...right];
}

/** Whether a cell can be selected: an empty slot cannot (§4.4 Disabled). */
export function selectable(cell: DollCell): boolean {
  return cell.item !== null;
}

/**
 * The slot a key moves focus to (§4.4): down and up in reading order, left and right across to
 * the same row of the other column, Home and End to the first and last. Empty slots are skipped.
 * Null for any other key, or when nothing is selectable.
 */
export function dollKeyTarget(cells: readonly DollCell[], from: string, key: string): string | null {
  const { left, right } = dollColumns(cells);
  const order = [...left, ...right].filter(selectable);
  if (order.length === 0) return null;
  const at = Math.max(
    0,
    order.findIndex((cell) => cell.slot === from),
  );
  switch (key) {
    case 'ArrowDown':
      return order[Math.min(order.length - 1, at + 1)].slot;
    case 'ArrowUp':
      return order[Math.max(0, at - 1)].slot;
    case 'Home':
      return order[0].slot;
    case 'End':
      return order[order.length - 1].slot;
    case 'ArrowRight':
    case 'ArrowLeft': {
      const [own, other] = left.some((cell) => cell.slot === from) ? [left, right] : [right, left];
      const row = own.findIndex((cell) => cell.slot === from);
      const across =
        other
          .slice(0, Math.max(row, 0) + 1)
          .reverse()
          .find(selectable) ?? other.find(selectable);
      return across?.slot ?? from;
    }
    default:
      return null;
  }
}

/** The slot a doll opens on (§4.5): the first upgrade when there is one, else the first selectable slot. */
export function defaultSelection(cells: readonly DollCell[]): string | null {
  const order = readingOrder(cells).filter(selectable);
  return (order.find((cell) => cell.badge === 'upgrade') ?? order[0])?.slot ?? null;
}
