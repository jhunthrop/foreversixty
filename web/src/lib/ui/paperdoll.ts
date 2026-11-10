// web/src/lib/ui/paperdoll.ts
// The shared paperdoll's cell model (web layout spec 2026-11-05 §4): every surface (BiS pick and
// compare, the character page, the sim and the report's worn sets) maps its own data to
// `DollCell`s, and ui/PaperdollFrame draws them; the mapper decides the words.
import type { GainFigure, ItemRef } from './primitives';

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
  /** The worn item's enchant in the client's words ("Crusader"), after line 2 in the game's green. */
  enchant?: string;
  /** Detail only: line 2 of the cell is the source alone (owner decision, option C). */
  note?: DollNote;
  gain?: GainFigure;
  /** The worn item's small icon beside the gain (compare mode). */
  worn?: ItemRef;
  /** An empty slot's line 1: "Nothing better yet" (a pick), "Empty" (worn). */
  emptyTitle?: string;
  /** An empty slot's reason ("First trinket comes at 40"). */
  empty?: string;
  /** A two-hander covers this off hand: the socket shows it greyed. */
  coveredBy?: ItemRef;
}

/** The order a doll opens in: head to ranged down the character panel's left, then hands to the
 *  trinkets down its right. */
const OPEN_ORDER: readonly string[] = [
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
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
];

/** Whether a cell can be selected: an empty slot cannot (§4.4 Disabled). */
export function selectable(cell: DollCell): boolean {
  return cell.item !== null;
}

/** The slot a doll opens on (§4.5): the first upgrade when there is one, else the first selectable slot. */
export function defaultSelection(cells: readonly DollCell[]): string | null {
  const bySlot = new Map(cells.map((cell) => [cell.slot, cell]));
  const order = OPEN_ORDER.flatMap((slot) => bySlot.get(slot) ?? []).filter(selectable);
  return (order.find((cell) => cell.badge === 'upgrade') ?? order[0])?.slot ?? null;
}
