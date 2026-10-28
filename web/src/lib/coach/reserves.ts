// web/src/lib/coach/reserves.ts
// The no-shift rule (experience spec 4.1): each text slot of a writing review holds a minimum
// height of `lines x 22 px`, so the model's text swaps in without moving the Then label, the
// Doing well label or the card's bottom edge. The classes are literal strings so Tailwind sees
// them: phone first, `md:` for the desktop counts.
import type { CoachReview } from './types';

export type SlotName = 'opening' | 'why' | 'drill' | 'thenWhy' | 'well' | 'chips' | 'links';

export const SLOT_RESERVE: Readonly<Record<SlotName, string>> = {
  opening: 'min-h-[154px] md:min-h-[66px]',
  why: 'min-h-[176px] md:min-h-[88px]',
  drill: 'min-h-[132px] md:min-h-[66px]',
  thenWhy: 'min-h-[110px] md:min-h-[66px]',
  well: 'min-h-[88px] md:min-h-[44px]',
  chips: 'min-h-[94px] md:min-h-[36px]',
  links: 'min-h-[80px] md:min-h-[40px]',
};

/** The slot's reserve class while the review is held at its minimums, nothing otherwise. */
export function slotClass(slot: SlotName, reserved: boolean): string {
  return reserved ? SLOT_RESERVE[slot] : '';
}

/** How many Then and Doing well slots the rule-only rendering fills, and which items carry links. */
export interface Reserve {
  thenCount: number;
  wellCount: number;
  /** Index 0 is Fix first, 1.. are the Then items. */
  links: readonly boolean[];
}

export function reserveFor(review: CoachReview): Reserve {
  return {
    thenCount: review.thenFix.length,
    wellCount: review.doingWell.length,
    links: [review.fixFirst, ...review.thenFix].map((item) => (item?.links.length ?? 0) > 0),
  };
}

/** A later answer never shrinks what an earlier one reserved. */
export function mergeReserve(previous: Reserve | null, next: Reserve): Reserve {
  if (previous === null) return next;
  const length = Math.max(previous.links.length, next.links.length);
  return {
    thenCount: Math.max(previous.thenCount, next.thenCount),
    wellCount: Math.max(previous.wellCount, next.wellCount),
    links: Array.from({ length }, (_, i) => previous.links[i] === true || next.links[i] === true),
  };
}
