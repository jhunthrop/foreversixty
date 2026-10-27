// web/src/lib/guild/seat-view.ts
// What the seats list decides from a seat read: which of the five §3.2.5 states a row is in, whether
// its 30-day lock still runs, and the error line a refused give or move reads as.
import { BillingApiError } from '../billing/api';
import type { SeatCandidate, SeatView } from '../billing/contract';
import { seatsCopy } from '../billing/seats-copy';

export type SeatRowKind = 'empty' | 'returned' | 'held' | 'offered' | 'accepted';

export function seatRowKind(seat: SeatView): SeatRowKind {
  if (seat.state === 'empty') return seat.returned === null ? 'empty' : 'returned';
  return seat.state;
}

/** Seats that name someone: held, offered or accepted. */
export function givenCount(seats: readonly SeatView[]): number {
  return seats.filter((seat) => seat.state !== 'empty').length;
}

export function isSeatLocked(seat: SeatView, now: Date): boolean {
  return seat.movable_at !== null && new Date(seat.movable_at).getTime() > now.getTime();
}

export function sortedByName(candidates: readonly SeatCandidate[]): SeatCandidate[] {
  return [...candidates].sort((a, b) => a.name.localeCompare(b.name));
}

/** `seat_locked` names the date it can move, `not_member` the member; anything else is the generic line. */
export function seatErrorLine(error: unknown, chosenName: string, zone?: string): string {
  if (!(error instanceof BillingApiError)) return seatsCopy.picker.saveFailed;
  const movableAt = error.fields?.movable_at;
  if (error.code === 'seat_locked' && movableAt !== undefined) {
    return seatsCopy.picker.lockedError(movableAt, zone);
  }
  if (error.code === 'not_member') return seatsCopy.picker.notMember(chosenName);
  return seatsCopy.picker.saveFailed;
}
