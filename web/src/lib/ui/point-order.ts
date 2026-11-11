// web/src/lib/ui/point-order.ts
// The PointOrder primitive's data (planner spec 2026-11-08 section 3A, 4.5): the build's spend order
// as rows of ten levels, a socket per level.

import type { GameTipModel } from '../tooltips/game-tip';

/** One level's socket: a spent point, one spent past the character's level, an empty level, the
 *  next point, or `ahead`: the spec guide's next point past the build (not part of it), dimmed. */
export interface OrderSlot {
  kind: 'spent' | 'planned' | 'empty' | 'next' | 'ahead';
  /** The socket's title and accessible name (`Level 12: Endurance Training 3`, `Next point: level 31`). */
  name: string;
  iconSrc?: string;
  /** The rank the talent reached with this point. */
  rank?: number;
  /** The talent's game tooltip at the rank this point reached, shown on hover and focus in place
   *  of the plain title. */
  tip?: GameTipModel;
}

export interface PointOrderRow {
  /** `Levels 10 to 19`, `Level 60`. */
  label: string;
  slots: readonly OrderSlot[];
}

/** The faint marker over the first row: the last digit of each socket's level. */
export const ORDER_MARKERS: readonly number[] = Array.from({ length: 10 }, (_, digit) => digit);

/** A socket that holds a point (spent or planned): it carries an icon and a tab stop. */
export function holdsPoint(slot: OrderSlot): boolean {
  return slot.kind === 'spent' || slot.kind === 'planned';
}

/** A socket that draws a talent: a point of the build, or the guide's point ahead of it. */
export function drawsTalent(slot: OrderSlot): boolean {
  return holdsPoint(slot) || slot.kind === 'ahead';
}

/**
 * Each slot's place in the order (0-based, counting only the sockets that hold a point, in
 * reading order), or null for an empty or next socket: the `order-point-<index>` test id.
 */
export function orderIndices(rows: readonly PointOrderRow[]): (number | null)[][] {
  let next = 0;
  return rows.map((row) => row.slots.map((slot) => (holdsPoint(slot) ? next++ : null)));
}
